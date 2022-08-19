package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	friends := []string{"Harish", "Mohan", "Sanjay"}
	bunch, err := greetings.Hellos(friends)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(bunch)
	}
}
