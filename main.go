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

	for {
		userInput, _ := reader.ReadString('\n')
		
		if userInput == "quit" {
			break
		} else {
			response := doctor.Response(userInput)
			fmt.Println(response)
		}
	}
}