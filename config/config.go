package config

import (
	"errors"
)

var (
	logger *Logger
)

func Init() error {
	return errors.New("Fake erro")
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}
