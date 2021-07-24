package greetings

import (
	"errors"
	"fmt"
)

// a function whose name starts with a capital letter can be called by a function not in the same package.
// note that the func returns multiple types now
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// := operator is a shortcut for declaring and initializing a variable in one line
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	// nil means no error
	return message, nil
}
