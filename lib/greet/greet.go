package greet

import "fmt"

type Language string

const (
	LanguageEnglish Language = "en"
	LanguageSpanish Language = "sp"
)

func (l Language) GetPrefix() string {
	switch l {
	case LanguageEnglish:
		return "Hello"

	case LanguageSpanish:
		return "Hola"
	}

	return "Hello"
}

func Greet(language Language, name string) string {
	prefix := language.GetPrefix()
	return fmt.Sprintf("%s, %s!", prefix, name)
}
