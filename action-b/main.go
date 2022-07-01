package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	time := strings.Join(os.Args[1:], " ")
	fmt.Printf("Action B: %s\n", time)
}