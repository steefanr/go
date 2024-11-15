package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	result, err := handleArgs(args, os.DirFS("."))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func handleArgs(args []string, fs fs.FS) (result string, err error) {
	result = ""
	err = nil

	if len(args) == 0 || len(args) > 2 {
		return result, errors.New("invalid number of arguments")
	}

	filePath := args[0]

	if len(args) == 2 {
		filePath = args[1]
	}

	f, err1 := fs.Open(filePath)

	if err1 != nil {
		return result, errors.New("failed to open file")
	}

	defer f.Close()
	fileInfo, _ := f.Stat()
	data, err := io.ReadAll(f)

	if err != nil {
		return result, errors.New("failed to read file")
	}

	count := ""

	if len(args) == 1 {
		count = handleAllCount(data)
	} else {
		count, err = handleSingleCount(args[0], data)

		if err != nil {
			return result, err
		}
	}

	return count + " " + fileInfo.Name(), nil
}

func handleAllCount(data []byte) string {
	byteCount := countBytes(data)
	wordCount := countWords(data)
	lineCount := countLines(data)

	return fmt.Sprintf("%8s %8s %8s", lineCount, wordCount, byteCount)
}

func handleSingleCount(flag string, data []byte) (string, error) {
	count := ""

	switch flag {
	case "-c":
		count = countBytes(data)
	case "-l":
		count = countLines(data)
	case "-w":
		count = countWords(data)
	default:
		return "", errors.New("invalid flag " + flag)
	}

	return fmt.Sprintf("%10s", count), nil
}

func countBytes(data []byte) string {
	return strconv.Itoa(len(data))
}

func countLines(data []byte) string {
	return strconv.Itoa(bytes.Count(data, []byte("\n")))
}

func countWords(data []byte) string {
	return strconv.Itoa(len(strings.Fields(string(data))))
}
