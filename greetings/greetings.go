package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// a function whose name starts with a capital letter can be called by a function not in the same package.
// note that the func returns multiple types now
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// := operator is a shortcut for declaring and initializing a variable in one line
	message := fmt.Sprintf(randomFormat(), name)

	// nil means no error
	return message, nil
}

// Go executes init functions automatically at program startup, after global variables have been initialized.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {

	// define a slice. A slice is like an array, except that its size changes dynamically as you add and remove items.
	formats := []string{
		"Hi, %v, Welcome",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
