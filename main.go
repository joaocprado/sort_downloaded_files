package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func logMove(oldPath string, newPath string) {
	message := fmt.Sprintf("File moved: %s -> %s", oldPath, newPath)
	fmt.Println(message)
}

func main() {
	homeDir, _ := os.UserHomeDir()
	mediaDir := path.Join(homeDir, "media")
	animeDir := path.Join(mediaDir, "anime")
	movieDir := path.Join(mediaDir, "movie")
	downloadDir := path.Join(homeDir, "downloads", "complete")
	filesInDir, _ := os.ReadDir(downloadDir)
	subspleaseRegex, _ := regexp.Compile(`\[SubsPlease\]`)
	ytsRegex, _ := regexp.Compile(`\[YTS\.\w+\]`)
	tgxRegex, _ := regexp.Compile(`\[TGx\]`)
	for _, file := range filesInDir {
		oldPath := path.Join(downloadDir, file.Name())
		fmt.Println(oldPath)
		if subspleaseRegex.MatchString(file.Name()) {
			fmt.Println(file.Name())
			newPath := path.Join(animeDir, file.Name())
			err := os.Rename(oldPath, newPath)
			checkErr(err)
			logMove(oldPath, newPath)
		}
		if ytsRegex.MatchString(file.Name()) || tgxRegex.MatchString(file.Name()) {
			fmt.Println(file.Name())
			newPath := path.Join(movieDir, file.Name())
			err := os.Rename(oldPath, newPath)
			checkErr(err)
			logMove(oldPath, newPath)
		}
	}
}
