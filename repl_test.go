package main

import (
	"testing"
)


func TestCleanInput(t *testing.T){

	// test cases will be inside of this anonymous struct

	cases := []struct{
		input string
		expected []string
	}{
		{
			input: "  hello  world  ",
			expected: []string{"hello","world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: "Lawwrd he coming",
			expected: []string{"lawwrd", "he", "coming"},
		},
	}


	//Here I will loop through each test case
	//For each test case I will call cleanInput and store the return value in the actual variable

	for _, c := range cases {
		actual := cleanInput(c.input)

	// Next I will check that the length of the actual string slice is equal to the expect string slice
	if len(actual) != len(c.expected) {
		t.Errorf("cleanInput(%q) returned %d words, expected %d", c.input, len(actual), len(c.expected))
		continue
	}

	// now we should ensure that each word is formatted as expected

	for i := range actual {
		if actual[i] != c.expected[i]{
			t.Errorf("cleanInput(%q) return %q at position %d, expected %q", c.input, actual[i], i, c.expected[i])
		}
	}


	}
}