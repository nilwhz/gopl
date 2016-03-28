package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// echov1()
	// echov2()
	// echov3()
	// echov4()
	// echov5()
}

func echov1() {
	var s, step string
	for _, arg := range os.Args[1:] {
		s += step + arg
		step = " "
	}
	fmt.Println(s)
}

func echov2() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func echov3() {
	fmt.Println(os.Args[1:])
}

func echov4() {
	fmt.Println(strings.Join(os.Args[:], " "))
}

func echov5() {
	for index, value := range os.Args[1:] {
		fmt.Println(fmt.Sprintf("index:%d\tvalue:%s", index, value))
	}
}
