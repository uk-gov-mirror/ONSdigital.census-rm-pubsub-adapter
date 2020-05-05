package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"reflect"
	"testing"
)

func TestGetZapLevel(t *testing.T) {
	t.Run("Test INFO", testGetZapLevel("INFO", zap.InfoLevel))
	t.Run("Test DEBUG", testGetZapLevel("DEBUG", zap.DebugLevel))
	t.Run("Test ERROR", testGetZapLevel("ERROR", zap.ErrorLevel))
	t.Run("Test WARN", testGetZapLevel("WARN", zap.WarnLevel))
}

func testGetZapLevel(levelString string, expectedLevel zapcore.Level) func(*testing.T) {
	return func(t *testing.T) {
		level, err := getZapLevel(levelString)
		if err != nil {
			t.Error(err)
		}
		if reflect.DeepEqual(level, expectedLevel) {
			t.Errorf("Level: %s doesn't match expected level: %s", level, expectedLevel)
		}
	}
}
