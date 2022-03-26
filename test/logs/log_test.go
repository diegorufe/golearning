package test

import (
	"log"
	"testing"
)

// Test de log con salida en pantalla con flags
func TestLog(t *testing.T) {
	// Ejemplo:  2022/03/26 18:37:08.642943 e:/trabajo/repos/go/golearning/test/logs/log_test.go:11: TestLog
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("TestLog")
}
