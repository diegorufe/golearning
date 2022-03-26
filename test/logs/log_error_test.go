package test

import (
	"log"
	"testing"

	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// TestError para ver la salida de stacktrace de un error
func TestError(t *testing.T) {
	// Ejemplo
	// 	2022/03/26 18:37:42.381043 e:/trabajo/repos/go/golearning/test/logs/log_error_test.go:21: golearning/test/logs.TestError
	// 	e:/trabajo/repos/go/golearning/test/logs/log_error_test.go:17
	// 2022/03/26 18:37:42.455005 e:/trabajo/repos/go/golearning/test/logs/log_error_test.go:21: testing.tRunner
	// 	E:/programas/go-installation/go1.17/src/testing/testing.go:1259
	// 2022/03/26 18:37:42.455005 e:/trabajo/repos/go/golearning/test/logs/log_error_test.go:21: runtime.goexit
	// 	E:/programas/go-installation/go1.17/src/runtime/asm_amd64.s:1581
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	errWithoutStacTrace := errors.New("This is a error")

	if err, ok := errWithoutStacTrace.(stackTracer); ok {
		for _, f := range err.StackTrace() {
			log.Printf("%+s:%d\n", f, f)
		}
	}
}
