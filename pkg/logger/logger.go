package logger

import (
	log "github.com/sirupsen/logrus"
)

// Logger is a base struct that could eventually maintain connections to something like bugsnag or logging tools
type Logger struct {
	serviceName string
}

// NewLogger creates a new instance of the custom logger struct and returns it
func NewLogger(serviceName string) *Logger {
	var l = new(Logger)
	l.serviceName = serviceName

	return l
}

// LogInfo is a publicly exposed info log that passes the message along correctly
func (l *Logger) LogInfo(messages ...interface{}) {
	log.WithField("service-name", l.serviceName).Info(messages)
}

// LogWarning is a publicly exposed info log that passes the message along correctly
func (l *Logger) LogWarning(messages ...interface{}) {
	log.WithField("service-name", l.serviceName).Warn(messages)
}

// LogError is a publicly exposed info log that passes the message along correctly
func (l *Logger) LogError(messages ...interface{}) {
	log.WithField("service-name", l.serviceName).Error(messages)
}
