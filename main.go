package main

import (
	"bufio"
	"fmt"
	"goBootcamp/doctor"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	userInput, _ := reader.ReadString('\n')

	fmt.Println("User Input:", userInput)
}