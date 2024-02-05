package main

import (
	"fmt"
	"greetings"
)

func main() {
	message := greetings.Hello("Hey Gophers!")
	fmt.Println(message)

}
