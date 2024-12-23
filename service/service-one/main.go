package main

import (
	"fmt"

	"github.com/project/lib/greet"
)

func main() {
	result := greet.Greet(greet.LanguageEnglish, "Service One")
	fmt.Println(result)
}
