package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const german = "German"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const germanHelloPrefix = "Moin, "

func Hello(name string, language string) string {
    if name == "" {
        name = "World"
    }

    return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
    switch language {
    case german:
        prefix = germanHelloPrefix
    case french:
        prefix = frenchHelloPrefix
    case spanish:
        prefix = spanishHelloPrefix
    default:
        prefix = englishHelloPrefix
    }

    return 
}


func main() {
    fmt.Println(Hello("world", ""))
}
