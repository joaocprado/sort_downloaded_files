package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	homeDir, _ := os.UserHomeDir()
	downloadDir := fmt.Sprintf("%s/downloads/complete", homeDir)
	// downloadDir = "/home/joao/code_project/sort_downloaded_files/downloads" // Test
	filesInDir, _ := os.ReadDir(downloadDir)
	subspleaseRegex, _ := regexp.Compile(`\[SubsPlease\]`)
	ytsRegex, _ := regexp.Compile(`\[YTS\.\w+\]`)
	tgxRegex, _ := regexp.Compile(`\[TGx\]`)
	for _, file := range filesInDir {
		if subspleaseRegex.MatchString(file.Name()) {
			fmt.Println(file.Name())
		}
		if ytsRegex.MatchString(file.Name()) {
			fmt.Println(file.Name())
		}
		if tgxRegex.MatchString(file.Name()) {
			fmt.Println(file.Name())
		}
	}
}
