package greetings

import (
	"regexp"
	"testing"
)

// *testing.T -> a pointer to the testing package's testing.T type
func TestHelloName(t *testing.T) {
	name := "Tim"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Tim")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
