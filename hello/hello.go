package main

import (
	"fmt"
	"greetings"
)

func main() {
	//Get a greetings message and print it.
	message := greetings.Hello("Peter")
	fmt.Println(message)
}
