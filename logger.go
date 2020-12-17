package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	encoder zapcore.EncoderConfig
	cfg     zap.Config
	level   zapcore.Level
	l       *zap.Logger
}

func New(opts ...Option) (*Logger, error) {
	logger := &Logger{
		level: zapcore.InfoLevel,
	}

	logger.encoder = zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	logger.cfg = zap.Config{
		Level:       zap.NewAtomicLevelAt(logger.level),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    logger.encoder,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	for _, o := range opts {
		o(logger)
	}

	var err error

	logger.l, err = logger.cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
