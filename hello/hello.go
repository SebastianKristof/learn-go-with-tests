package main

import "fmt"

const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"
const english = "English"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
    }
    
	return greetHelloPrefix(language) + name
    
}
func greetHelloPrefix(language string) (prefix string) {
    
    switch language {
    case french:
        prefix = frenchHelloPrefix
    case spanish:
        prefix = spanishHelloPrefix
    case english:
        prefix = helloPrefix
    default:
        prefix = helloPrefix
    }
    return
}

func main() {
	fmt.Println(Hello("Seb", ""))
}
