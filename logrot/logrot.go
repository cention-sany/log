// Package logrot handles log rotation on SIGHUP.
package logrot

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/cention-sany/log"
)

var lg log.IFLogger

func mustOpenFileForAppend(name string) *os.File {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		lg.Fatal("Error: ", err)
	}
	return f
}

// LogRot represents log file that will be reopened on a given signal.
type LogRot struct {
	lg      log.OutSetter
	name    string
	LogFile *os.File
	signal  os.Signal
	quit    chan struct{}
}

// WriteTo sets the log output to the given file and reopen the file on SIGHUP.
func WriteTo(name string) *LogRot {
	return rotateOn(name, syscall.SIGHUP, log.StdLib())
}

func WriteToWithLog(name string, l log.OutSetter) *LogRot {
	return rotateOn(name, syscall.SIGHUP, l)
}

// WriteAllTo sets the log output, os.Stdout and os.Stderr to the given file and reopen the file on SIGHUP.
func WriteAllTo(name string) *LogRot {
	lr := WriteTo(name)
	os.Stdout = lr.LogFile
	os.Stderr = lr.LogFile
	return lr
}

func WriteAllToWithLog(name string, l log.OutSetter) *LogRot {
	lr := WriteToWithLog(name, l)
	os.Stdout = lr.LogFile
	os.Stderr = lr.LogFile
	return lr
}

// rotateOn rotates the log file on the given signals
func rotateOn(name string, sig os.Signal, l log.OutSetter) *LogRot {
	rl := &LogRot{
		lg:      l,
		name:    name,
		signal:  sig,
		LogFile: mustOpenFileForAppend(name),
		quit:    make(chan struct{}),
	}
	rl.lg.SetOutput(rl.LogFile)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, sig)
	go func() {
		for {
			select {
			case s := <-sigs:
				if s == rl.signal {
					lg.Printf("%s received - rotating log file handle on %s\n", s, rl.name)
					rl.rotate()
				}
			case <-rl.quit:
				return
			}
		}
	}()
	return rl
}

func (rl *LogRot) Close() {
	if rl != nil && rl.LogFile != nil {
		rl.quit <- struct{}{}
		rl.LogFile.Close()
	}
}

func (rl *LogRot) rotate() {
	oldLog := rl.LogFile
	rl.LogFile = mustOpenFileForAppend(rl.name)
	rl.lg.SetOutput(rl.LogFile)
	oldLog.Close()
}

func init() {
	if log.Level == log.LvlNoLog {
		lg = log.NoLog()
	} else {
		lg = log.StdLog()
	}
}
