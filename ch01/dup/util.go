package main

import (
	"bufio"
	"os"
)

type logInfo struct {
	times    int
	filename string
}

func countLines(f *os.File, counts map[string]int, countFiles map[string][]string) {
	filename := f.Name()

	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if !arrayContains(countFiles[input.Text()], filename) {
			countFiles[input.Text()] = append(countFiles[input.Text()], filename)
		}
	}
}

func arrayContains(arr []string, name string) bool {
	for _, filename := range arr {
		if filename == name {
			return true
		}
	}
	return false
}
