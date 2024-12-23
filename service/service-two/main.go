package main

import (
	"fmt"

	"github.com/project/lib/greet"
)

func main() {
	result := greet.Greet(greet.LanguageSpanish, "Service Two")
	fmt.Println(result)
}
