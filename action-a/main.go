package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Action A")
	fmt.Printf(`echo "::set-output name=time::%s"\n`, time.Now().Format(time.RFC3339))
}