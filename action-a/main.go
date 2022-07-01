package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Action A")
	fmt.Fprintf(os.Stdout, "::set-output name=time::%s\n", time.Now().Format(time.RFC3339))
}