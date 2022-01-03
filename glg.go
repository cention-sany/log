package log

// g-log is just helper utility for log interface to link
// between custom logger and stdlib log package.
import (
	"io"
	"log"
)

const (
	LvlNoLog = iota
	LvlDebug
)

// not thread safe - safe it once only
var Level = LvlNoLog

type Prefixer interface {
	Prefix() string
	SetPrefix(prefix string)
}

type OutSetter interface {
	SetOutput(w io.Writer)
}

type Outputter interface {
	Output(calldepth int, s string) error
}

type FlagsSetter interface {
	SetFlags(int)
}

type FlagsGetter interface {
	Flags() int
}

type PanicLogger interface {
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
}

type FatalLogger interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
}

type PrintLogger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

type PLogger interface {
	PanicLogger
	PrintLogger
}

type IFLogger interface {
	OutSetter
	Outputter
	FlagsSetter
	FlagsGetter
	Prefixer
	FatalLogger
	PanicLogger
	PrintLogger
}

// StdLog is singleton object for this log package. It is equivalent to
// standard library log object but NOT equal it. Use the StdLib for accessing
// standard library log singleton.
func StdLog() IFLogger {
	return std
}

// NoLog as name imply no logging.
type nolog struct{}

func NoLog() IFLogger {
	return &nolog{}
}

func (nolog) Fatal(v ...interface{}) {}

func (nolog) Fatalf(format string, v ...interface{}) {}

func (nolog) Fatalln(v ...interface{}) {}

func (nolog) Panic(v ...interface{}) {}

func (nolog) Panicf(format string, v ...interface{}) {}

func (nolog) Panicln(v ...interface{}) {}

func (nolog) Prefix() string {
	return ""
}

func (nolog) Print(v ...interface{}) {}

func (nolog) Printf(format string, v ...interface{}) {}

func (nolog) Println(v ...interface{}) {}

func (nolog) SetPrefix(prefix string) {}

func (nolog) SetOutput(w io.Writer) {}

func (nolog) Output(calldepth int, s string) error { return nil }

func (nolog) SetFlags(int) {}

func (nolog) Flags() int { return 0 }

// access to stdlib log
type stdLibLog struct{}

// StdLib allows access to global singleton standard library log object.
func StdLib() IFLogger {
	return &stdLibLog{}
}

func (stdLibLog) Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func (stdLibLog) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

func (stdLibLog) Fatalln(v ...interface{}) {
	log.Fatalln(v...)
}

func (stdLibLog) Panic(v ...interface{}) {
	log.Panic(v...)
}

func (stdLibLog) Panicf(format string, v ...interface{}) {
	log.Panicf(format, v...)
}

func (stdLibLog) Panicln(v ...interface{}) {
	log.Panicln(v...)
}

func (stdLibLog) Prefix() string {
	return log.Prefix()
}

func (stdLibLog) Print(v ...interface{}) {
	log.Print(v...)
}

func (stdLibLog) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (stdLibLog) Println(v ...interface{}) {
	log.Println(v...)
}

func (stdLibLog) SetPrefix(prefix string) {
	log.SetPrefix(prefix)
}

func (stdLibLog) SetOutput(w io.Writer) {
	log.SetOutput(w)
}

func (stdLibLog) Output(calldepth int, s string) error {
	return log.Output(calldepth, s)
}

func (stdLibLog) SetFlags(f int) {
	log.SetFlags(f)
}

func (stdLibLog) Flags() int {
	return log.Flags()
}
