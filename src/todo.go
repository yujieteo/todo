package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Get folder name from user
	fmt.Print("Enter folder name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	folderName := scanner.Text()

	// Create folder
	err := os.Mkdir(folderName, 0755)
	if err != nil {
		fmt.Println("Failed to create folder:", err)
		return
	}

	// Initialize git repository
	cmd := exec.Command("git", "init", folderName)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to initialize git repository:", err)
		return
	}

	// Create and write to todo file
	todoFile := folderName + "/todo.txt"
	file, err := os.Create(todoFile)
	if err != nil {
		fmt.Println("Failed to create todo file:", err)
		return
	}
	defer file.Close()

	fmt.Println("Enter todo items (press Ctrl+D or Ctrl+Z + Enter to finish):")

	scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		item := scanner.Text()

		// Trim leading whitespace and check if the item starts with a valid symbol
		trimmedItem := strings.TrimLeft(item, " \t")
		if len(trimmedItem) == 0 {
			fmt.Println("Invalid symbol:", trimmedItem)
			continue
		}

		symbol := trimmedItem[0]
		if symbol != '-' && symbol != 'x' && symbol != '?' && symbol != '\t' {
			fmt.Println("Invalid symbol:", string(symbol))
			continue
		}

		// Write the item to the file
		_, err := file.WriteString(item + "\n")
		if err != nil {
			fmt.Println("Failed to write to todo file:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("Todo file created successfully:", todoFile)
}
