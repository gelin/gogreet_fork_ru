package greet

import (
	"fmt"
	"testing"
)

func TestGreetingFor(t *testing.T) {
	result := GreetingFor("Мир")
	fmt.Println(result)
	if result != "Здравствуй, Мир!" {
		t.Fatal(result)
	}
}
