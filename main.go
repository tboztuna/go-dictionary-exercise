package main

import (
	"fmt"
)

func main() {
	colors := make(map[string]string)
	colors["white"] = "#F0000"

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
