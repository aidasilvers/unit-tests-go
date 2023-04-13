package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func TestPrompt(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	old := os.Stdout
	os.Stdout = w

	prompt()

	w.Close()
	os.Stdout = old

	output, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	expected := "-> "
	if string(output) != expected {
		t.Fatalf("expected %q but got %q", expected, string(output))
	}
}

func TestCheckNumbers(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
		done     bool
	}{
		{
			name:     "Invalid input",
			input:    "hello",
			expected: "Please enter a whole number!",
			done:     false,
		},
		{
			name:     "Zero",
			input:    "0",
			expected: "0 is not prime, by definition!",
			done:     false,
		},
		{
			name:     "One",
			input:    "1",
			expected: "1 is not prime, by definition!",
			done:     false,
		},
		{
			name:     "Negative number",
			input:    "-5",
			expected: "Negative numbers are not prime, by definition!",
			done:     false,
		},
		{
			name:     "Prime number",
			input:    "7",
			expected: "7 is a prime number!",
			done:     false,
		},
		{
			name:     "Quit",
			input:    "q",
			expected: "",
			done:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)
			scanner := bufio.NewScanner(r)

			output, done := checkNumbers(scanner)

			if output != tc.expected {
				t.Fatalf("expected %q but got %q", tc.expected, output)
			}

			if done != tc.done {
				t.Fatalf("expected done to be %v but got %v", tc.done, done)
			}
		})
	}
}

func TestIntro(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	old := os.Stdout
	os.Stdout = w

	intro()

	w.Close()
	os.Stdout = old

	output, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	expected := "Is it Prime?\n------------\nEnter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.\n-> "
	if string(output) != expected {
		t.Fatalf("expected %q but got %q", expected, string(output))
	}
}

func TestReadUserInput(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	old := os.Stdin
	os.Stdin = r

	doneChan := make(chan bool)
	go readUserInput(os.Stdin, doneChan)

	// Simulate user input
	input := "5\nq\n"
	_, err = fmt.Fprint(w, input)
	if err != nil {
		t.Fatal(err)
	}
	w.Close()

	<-doneChan

	os.Stdin = old
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
	// the best
}
