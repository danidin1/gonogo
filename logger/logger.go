package logger

import ()

const (
	FATAL_LEVEL   int = 0
	ERROR_LEVEL       = 1
	ALERT_LEVEL       = 2
	WARNING_LEVEL     = 3
	INFO_LEVEL        = 4
	DEBUG_LEVEL       = 5
)

func Debug(args ...interface{}) {
	LogMessage(DEBUG_LEVEL, args...)
}

func Info(args ...interface{}) {
	LogMessage(INFO_LEVEL, args...)
}

func Warning(args ...interface{}) {
	LogMessage(WARNING_LEVEL, args...)
}

func Error(args ...interface{}) {
	LogMessage(ERROR_LEVEL, args...)
}

func Alert(args ...interface{}) {
	LogMessage(ALERT_LEVEL, args...)
}

func Fatal(args ...interface{}) {
	LogMessage(FATAL_LEVEL, args...)
}

func LogMessage(severity int, args ...interface{}) {

}

func InitLogger(logLevel int) int {
	return 0
}

func GetLogName() string {
	return "dani's log"
}

func CloseLog() {
}
