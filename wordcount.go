package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := strings.Join(os.Args[1:], " ")
	count := len(strings.Fields(input))
	fmt.Printf("%d\n", count)
}