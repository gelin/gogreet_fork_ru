package greet

import (
	"fmt"
	"github.com/gelin/gogreet_fork_ru/format"
)

func GreetingFor(name string) string {
	return fmt.Sprintf(format.GreetingFormat(), name)
}
