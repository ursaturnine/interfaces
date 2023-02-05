package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// this function is of type, bot
func (englishBot) getGreeting() string {
	// very custom logic for English greeting
	return "Hi, There!"
}

// this function is of type, bot
func (spanishBot) getGreeting() string {
	// very custom logic for Spanish greeting
	return "Hola!"
}
