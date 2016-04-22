package log

import (
	syslog "log"
	"sync"
)

// Level of logging
type Level int

// Levels
const (
	LevelDebug Level = 1 + iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

// crude global log level control:
// log everything at this level and above
var logMutex sync.RWMutex
var logLevel = LevelDebug
var logLevelStack []Level

// LevelPush push current level down the stack and set  
func LevelPush( level Level) {
    logMutex.Lock()
    defer logMutex.Unlock()

    logLevelStack = append(logLevelStack, level)
    logLevel = level
}

// LevelPop pop current level from the stack   
func LevelPop() {
    logMutex.Lock()
    defer logMutex.Unlock()

    len := len(logLevelStack)
    logLevel, logLevelStack = logLevelStack[len-1], logLevelStack[:len-1]  
}

// Debug log if debug is greater than current log level
func Debug(msg ...interface{}) {
    logMutex.RLock()
    defer logMutex.RUnlock()

	if LevelDebug >= logLevel {
		syslog.Println("D:", msg)
	}
}

// Info log if info is greater than current log level
func Info(msg ...interface{}) {
    logMutex.RLock()
    defer logMutex.RUnlock()

	if LevelInfo >= logLevel {
		syslog.Println("I:", msg)
	}
}

// Warn log if warn is greater than current log level
func Warn(msg ...interface{}) {
    logMutex.RLock()
    defer logMutex.RUnlock()

	if LevelWarn >= logLevel {
		syslog.Println("W:", msg)
	}
}

// Error log if error is greater than current log level
func Error(msg ...interface{}) {
    logMutex.RLock()
    defer logMutex.RUnlock()

	if LevelError >= logLevel {
		syslog.Println("E:", msg)
	}
}

// Fatal log unconditionally and panic
func Fatal(msg ...interface{}) {
	syslog.Fatalln("F:", msg)
}
