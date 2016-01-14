package log

// g-log is just helper utility for log interface to link
// between custom logger and stdlib log package.
import (
	"io"
)

type Prefixer interface {
	Prefix() string
	SetPrefix(prefix string)
}

type OutSetter interface {
	SetOutput(w io.Writer)
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
	Prefixer
	PanicLogger
	PrintLogger
}

type IFLogger interface {
	OutSetter
	Prefixer
	FatalLogger
	PanicLogger
	PrintLogger
}

// StdLog is standard library global logging.
// This is created to make StdLog bind to Logger interface.
// Thus stdlib logger bind to Logger interface.
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
