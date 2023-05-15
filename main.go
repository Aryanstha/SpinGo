package main

import (
	"fmt"
	"time"
)

func spinner() {

	loadingChars := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

	for i := 0; i < 20; i++ {
		loadingChar := loadingChars[i%len(loadingChars)]
		colorCode := i%8 + 30 // Choose a color code (30-37) for foreground color

		// Format the output using escape sequences
		fmt.Printf("\033[0;%dm%s Loading...\033[0m\r", colorCode, loadingChar)
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println()
	fmt.Println("Loading complete!")

}
func main() {
	spinner()
}
