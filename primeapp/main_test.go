package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime!"},
		{"prime", 0, false, "0 is not a prime, by the definition!"},
		{"prime", -1, false, "Negative numbers are not prime, by definition!"},
		{"prime", 1, false, "1 is not a prime, by the definition!"},
		{"not prime", 8, false, "8 is not a prime, because it is divisible by 2"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}
		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}
		if msg != e.msg {
			t.Errorf("%s: Expected message to be %s, but got %s", e.name, e.msg, msg)
		}

	}

}
func Test_prompt(t *testing.T) {
	// clone os.Stdout
	oldOut := os.Stdout

	// create a pipe to read from
	r, w, _ := os.Pipe()
	os.Stdout = w

	prompt()

	// close the writer
	w.Close()

	// reset os.Stdout
	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if string(out) != "-> " {
		t.Errorf("Expected prompt to be '-> ', but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	// clone os.Stdout
	oldOut := os.Stdout

	// create a pipe to read from
	r, w, _ := os.Pipe()
	os.Stdout = w

	intro()

	// close the writer
	w.Close()

	// reset os.Stdout
	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if strings.Contains(string(out), " Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.") {
		t.Errorf("Expected intro to contain ' Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.', but got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number or q to quit"},
		{name: "q", input: "q", expected: ""},
		{name: "7", input: "7", expected: "7 is a prime!"},
		{name: "0", input: "0", expected: "0 is not a prime, by the definition!"},
		{name: "-1", input: "-1", expected: "Negative numbers are not prime, by definition!"},
		{name: "1", input: "1", expected: "1 is not a prime, by the definition!"},
		{name: "8", input: "8", expected: "8 is not a prime, because it is divisible by 2"},
		// test with 2
		{name: "2", input: "2", expected: "2 is a prime!"},
		// try with decimal
		{name: "decimal", input: "1.5", expected: "Please enter a whole number or q to quit"},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)
		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: Expected %s, but got %s", e.name, e.expected, res)
		}
	}
}

func Test_readerUserInput(t *testing.T) {
	doneChan := make(chan bool)
	var stdin bytes.Buffer
	stdin.Write([]byte("1\nq\n"))
	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
