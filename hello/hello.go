package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println("Hello, World!")
	// fmt.Println(quote.Go())

	// chamada simles
	// message, err := greetings.Hello("Nana")

	names := []string{"Anna", "Anne", "Diego", "Thiago"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
