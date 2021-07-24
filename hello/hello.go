package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Tim")
	fmt.Println(message)
}
