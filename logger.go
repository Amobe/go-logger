package logger

import (
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	LevelError = iota
	LevelWarning
	LevelInformational
	LevelDebug
)

type LogType int

const (
	LogTypeStdout LogType = iota
	LogTypeFile
)

var mLoggerInstance *Logger

type Logger struct {
	level   int
	logType LogType
	err     *log.Logger
	warn    *log.Logger
	info    *log.Logger
	debug   *log.Logger
	depth   int
	isInit  bool
}

func (logger *Logger) Init(logType LogType, filePath string) {
	var writer io.Writer

	if logger.isInit {
		return
	}

	logger.logType = logType
	if logType == LogTypeStdout {
		writer = os.Stdout
	} else {
		writer = &lumberjack.Logger{
			Filename:   filePath,
			MaxSize:    256,
			MaxBackups: 3,
			MaxAge:     28,
		}
	}

	flag := log.Lshortfile
	logger.err = log.New(writer, "[E] ", flag)
	logger.warn = log.New(writer, "[W] ", flag)
	logger.info = log.New(writer, "[I] ", flag)
	logger.debug = log.New(writer, "[D] ", flag)

	logger.SetLevel(LevelDebug)
	logger.depth = 2
	logger.isInit = true
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

func SetLoggerInstance(logger *Logger) {
	mLoggerInstance = logger
}

func GetLoggerInstance() *Logger {
	return mLoggerInstance
}
