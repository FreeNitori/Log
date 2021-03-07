package log

import (
	"github.com/sirupsen/logrus"
)

// Info logs message on info level.
func Info(args ...interface{}) {
	Instance.Info(args...)
}

// Infof logs message on info level.
func Infof(format string, args ...interface{}) {
	Instance.Infof(format, args...)
}

// Debug logs message on debug level.
func Debug(args ...interface{}) {
	Instance.Debug(args...)
}

// Debugf logs message on debug level.
func Debugf(format string, args ...interface{}) {
	Instance.Debugf(format, args...)
}

// Warn logs message on warn level.
func Warn(args ...interface{}) {
	Instance.Warn(args...)
}

// Warnf logs message on warn level.
func Warnf(format string, args ...interface{}) {
	Instance.Warnf(format, args...)
}

// Error logs message on error level.
func Error(args ...interface{}) {
	Instance.Error(args...)
}

// Errorf logs message on error level.
func Errorf(format string, args ...interface{}) {
	Instance.Errorf(format, args...)
}

// Fatal logs message on fatal level.
func Fatal(args ...interface{}) {
	Instance.Fatal(args...)
}

// Fatalf logs message on fatal level.
func Fatalf(format string, args ...interface{}) {
	Instance.Fatalf(format, args...)
}

// SetLevel sets the logger level.
func SetLevel(level logrus.Level) {
	Instance.SetLevel(level)
}

// GetLevel returns the logger level.
func GetLevel() logrus.Level {
	return Instance.GetLevel()
}

// IsLevelEnabled checks if the log level of the logger is greater than the level param.
func IsLevelEnabled(level logrus.Level) bool {
	return Instance.IsLevelEnabled(level)
}
