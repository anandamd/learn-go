package main

import (
	"flag"
	"fmt"
)

func main() {
	window := flag.String("window", "", "The window to check")
	flag.Parse()
	fmt.Println(*window)
	fmt.P
}
