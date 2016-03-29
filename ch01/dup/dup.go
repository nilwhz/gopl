package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// dupv1()
	dupv2()
	// dupv3()
}

func dupv1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	fmt.Println()
	for line, times := range counts {
		if times > 1 {
			fmt.Printf("%d\t%s\n", times, line)
		}
	}
}

func dupv2() {
	counts := make(map[string]int)
	countFiles := make(map[string][]string)

	fiels := os.Args[1:]
	if len(fiels) == 0 {
		countLines(os.Stdin, counts, countFiles)
	} else {
		for _, arg := range fiels {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
				continue
			}
			defer f.Close()
			countLines(f, counts, countFiles)
		}
	}
	for line, num := range counts {
		if num > 1 {
			fmt.Printf("%s\t%d\t%v\n", line, num, countFiles[line])
		}
	}
}

func dupv3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, num := range counts {
		fmt.Printf("%d\t%s\n", num, line)
	}
}
