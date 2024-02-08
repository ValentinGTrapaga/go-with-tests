package main

import (
	"fmt"
)

const (
	french  = "French"
	spanish = "Spanish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHeloPrefix   = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return getPrefixByLanguage(language) + name

}

// prefix is a named return value, defaults to the zero value, in this case ""
func getPrefixByLanguage(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHeloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}
