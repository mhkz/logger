package logger

import (
	"testing"
)

func TestFileLogger(t *testing.T){
	logger := NewFileLogger(LogLevelDebug, "/Users/guofeng/Desktop/logs", "test")
	logger.Debug("user id[%d] is come from china", 324234)
	logger.Warn("test warn log")
	logger.Fatal("test Fatal log")
	logger.Close()
}