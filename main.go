package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"card-detector-go/detector"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a card number: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read input:", err)
		return
	}

	cardNumber := strings.TrimSpace(input)
	network := detector.DetectNetwork(cardNumber)
	fmt.Println("Network:", network)
}
