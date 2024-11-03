package log

import (
	"log/slog"
	"os"
	"time"
)

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

func New() Logger {
	return slog.New(slog.NewJSONHandler(os.Stderr, nil))
}

func NewNop() Logger {
	return &Nop{}
}

type Nop struct{}

func (Nop) Debug(string, ...any) {}
func (Nop) Info(string, ...any)  {}
func (Nop) Warn(string, ...any)  {}
func (Nop) Error(string, ...any) {}

func Error(err error) any {
	return slog.String("err", err.Error())
}

func String(key string, value string) any {
	return slog.String(key, value)
}

func Bool(key string, value bool) any {
	return slog.Bool(key, value)
}

func Int(key string, value int) any {
	return slog.Int64(key, int64(value))
}

func Int64(key string, value int64) any {
	return slog.Int64(key, value)
}

func Uint(key string, value uint) any {
	return slog.Uint64(key, uint64(value))
}

func Uint64(key string, value uint64) any {
	return slog.Uint64(key, value)
}

func Duration(key string, value time.Duration) any {
	return slog.Duration(key, value)
}

func Time(key string, value time.Time) any {
	return slog.Time(key, value)
}
