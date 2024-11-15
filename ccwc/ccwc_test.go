package main

import (
	"testing"
	"testing/fstest"
)

func assertEquals(t *testing.T, actual any, expected any) {
	if actual != expected {
		t.Errorf("expected \"%s\" result but got \"%s\"", expected, actual)
	}
}

func assertHandleArgsReturnsError(t *testing.T, actual error, expectedMessage string) {
	if actual == nil {
		t.Error("Expected error but got nil")
	} else if actual.Error() != expectedMessage {
		t.Errorf("expected error message \"%s\" but got \"%s\"", expectedMessage, actual.Error())
	}
}

type HandleArgsErrorTestParameter struct {
	Title         string
	Arguments     []string
	ExpectedError string
}

func TestHandleArgsErrors(t *testing.T) {
	tests := []HandleArgsErrorTestParameter{
		{
			"it returns an error when no file name is given",
			[]string{"-c", "test", ".txt"},
			"invalid number of arguments",
		},
		{
			"it returns an error when no arguments are given",
			[]string{},
			"invalid number of arguments",
		},
		{
			"it returns an error when a bad flag is passed",
			[]string{"-x", "test"},
			"invalid flag -x",
		},
		{
			"it returns an error when the file given does not exist",
			[]string{"-c", "fileThatDoesntExist"},
			"failed to open file",
		},
	}

	for _, p := range tests {
		t.Run(p.Title, func(t *testing.T) {
			result, err := handleArgs(p.Arguments, fstest.MapFS{
				"test": {Data: []byte("test")},
			})

			assertEquals(t, result, "")
			assertHandleArgsReturnsError(t, err, p.ExpectedError)
		})
	}
}

type CountTestParameters struct {
	Title    string
	Data     string
	Expected string
}

func TestCountsBytesInString(t *testing.T) {
	tests := []CountTestParameters{
		{
			"it returns the correct value for an example string",
			"This is a test",
			"14",
		},
		{
			"it returns 0 when the string is empty",
			"",
			"0",
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			data := []byte(test.Data)
			result := countBytes(data)
			assertEquals(t, result, test.Expected)
		})
	}
}

func TestCountLinesInString(t *testing.T) {
	tests := []CountTestParameters{
		{
			"it returns the correct number of lines",
			"1 Line \n 2 Lines \n 3 Lines\n",
			"3",
		},
		{
			"it returns 1 line when there is only one line",
			"1 Line\n",
			"1",
		},
		{
			"it returns 0 when there is no contents",
			"",
			"0",
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			data := []byte(test.Data)
			result := countLines(data)
			assertEquals(t, result, test.Expected)
		})
	}
}

func TestCountWordsInString(t *testing.T) {
	tests := []CountTestParameters{
		{
			"it returns the correct number of words in a string",
			"This is a string",
			"4",
		},
		{
			"it returns the correct number of words even when there are multiple spaces",
			"This  is a string",
			"4",
		},
		{
			"it returns 0 when given an empty string",
			"",
			"0",
		},
		{
			"it returns 1 when there is one word with no spaces",
			"This",
			"1",
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			data := []byte(test.Data)
			result := countWords(data)
			assertEquals(t, result, test.Expected)
		})
	}
}
