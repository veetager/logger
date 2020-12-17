package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option func(l *Logger)

func Debug() Option {
	return func(l *Logger) {
		l.level = zap.DebugLevel
		l.encoder.EncodeLevel = zapcore.CapitalLevelEncoder
		l.encoder.EncodeTime = zapcore.ISO8601TimeEncoder
		l.encoder.EncodeDuration = zapcore.StringDurationEncoder
		l.cfg.Development = true
		l.cfg.Sampling = nil
		l.cfg.Encoding = "console"
		l.cfg.Level = zap.NewAtomicLevelAt(l.level)
	}
}
