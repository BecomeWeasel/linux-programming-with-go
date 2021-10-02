package main

import (
	"bufio"
	"fmt"
	"os"
)

const path = "testfile"

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	f, err := os.Open(path)
	checkError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		checkError(scanner.Err())
		fmt.Println(line)
	}
}
