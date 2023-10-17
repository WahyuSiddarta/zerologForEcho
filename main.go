package zerologForEcho

import (
	"fmt"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

const DefaultRequestIDHeader = echo.HeaderXRequestID
const DefaultRecoverMsg = "Recovered"

var defaultRecoverConfig = RecoverConfig{
	StackTraceSize: 4 << 10, // 4 KB
}

type RecoverConfig struct {

	// Size allocated on memory for stack trace.
	StackTraceSize int
	// If stack trace is enabled, this is to print stack traces of all goroutines.
	PrintStackTraceOfAllGoroutines bool

	// The panic was happened, and it was handled and logged gracefully.
	// What's next?
	//
	// This function is called to handle the error of panic.
	ErrorHandler func(c echo.Context, err error)
}

func Recover(zerolog *zerolog.Logger) echo.MiddlewareFunc {
	return RecoverWithConfig(defaultRecoverConfig, zerolog)
}

func RecoverWithConfig(config RecoverConfig, zerolog *zerolog.Logger) echo.MiddlewareFunc {
	if config.StackTraceSize == 0 {
		config.StackTraceSize = defaultRecoverConfig.StackTraceSize
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					e := func() error {
						if e, ok := err.(error); ok {
							return e
						} else {
							return fmt.Errorf("panic: %v", err)
						}
					}()

					c.Error(e)
					stack := make([]byte, config.StackTraceSize)
					stackLen := runtime.Stack(stack, config.PrintStackTraceOfAllGoroutines)
					// fmt.Println("MYFUNCTION ", string(stack[:stackLen]))
					zerolog.Error().Msgf("RECOVERY %+v", string(stack[:stackLen]))

					if config.ErrorHandler != nil {
						config.ErrorHandler(c, e)
					}
				}
			}()
			return next(c)
		}
	}
}
