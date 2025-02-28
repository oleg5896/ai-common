package logger

import (
	"github.com/rs/zerolog"
	"os"
)

var log = zerolog.New(os.Stdout).With().Timestamp().Logger()

func Info(msg string) {
	log.Info().Msg(msg)
}

func Error(msg string, err error) {
	log.Error().Err(err).Msg(msg)
}
