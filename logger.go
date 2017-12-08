package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	LevelError = iota
	LevelWarning
	LevelInformational
	LevelDebug
)

type Logger struct {
	level int
	err   *log.Logger
	warn  *log.Logger
	info  *log.Logger
	debug *log.Logger
	depth int
}

var loggerInstance *Logger

func GetLoggerInstance() *Logger {
	if loggerInstance != nil {
		return loggerInstance
	}

	loggerInstance = new(Logger)
	loggerInstance.Init()

	return loggerInstance
}

func (logger *Logger) Init() {
	writer := os.Stdout
	flag := log.Lshortfile

	logger.err = log.New(writer, "[E] ", flag)
	logger.warn = log.New(writer, "[W] ", flag)
	logger.info = log.New(writer, "[I] ", flag)
	logger.debug = log.New(writer, "[D] ", flag)

	logger.SetLevel(LevelDebug)
	logger.depth = 2
}

func (logger *Logger) SetLevel(level int) {
	logger.level = level
}

func (logger *Logger) Error(format string, v ...interface{}) {
	if LevelError > logger.level {
		return
	}
	logger.err.Output(logger.depth, fmt.Sprintf(format, v...))
}

func (logger *Logger) Warning(format string, v ...interface{}) {
	if LevelWarning > logger.level {
		return
	}
	logger.warn.Output(logger.depth, fmt.Sprintf(format, v...))
}

func (logger *Logger) Information(format string, v ...interface{}) {
	if LevelInformational > logger.level {
		return
	}
	logger.info.Output(logger.depth, fmt.Sprintf(format, v...))
}

func (logger *Logger) Debug(format string, v ...interface{}) {
	if LevelDebug > logger.level {
		return
	}
	logger.debug.Output(logger.depth, fmt.Sprintf(format, v...))
}
