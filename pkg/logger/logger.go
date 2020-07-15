package logger

import (
	"context"
	"io"
	"log"
)

type Level int8

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	default:
		return ""
	}
}

type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	level     Level
	fields    Fields
	callers   []string
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

func (l *Logger) WithLevel(lvl Level) *Logger {
	ll := l.clone()
	ll.level = lvl
	return ll
}

func (l *Logger) WithFields(f Fields) *Logger {
	ll := l.clone()
	if ll.fields == nil {
		ll.fields = make(Fields)
	}
	for k, v := range f {
		ll.fields[k] = v
	}
	return ll
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx
	return ll
}

//func (l *Logger) WithCaller(skip int) *Logger {
//	ll := l.clone()
//	pc, file, line, ok := runtime.Caller(skip)
//	if ok {
//		f := runtime.FuncForPC(pc)
//		ll.callers = []string{fmt.Sprintf("%s: %d %s"), file, line, f.Name()}
//	}
//}
