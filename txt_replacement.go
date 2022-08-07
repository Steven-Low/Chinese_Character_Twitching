package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("D:\\Documents\\go_workspace\\mandarin_text_replacement\\Chinese_Character_Twitching\\full.txt")
	if err != nil {
		log.Fatalf("Error while opening file. Err: %s", err)
	}
	defer file.Close()

	fileBuffer := new(bytes.Buffer)
	fileBuffer.ReadFrom(file)
	fileString := fileBuffer.String()

	fmt.Print(fileString)

	// (``) for mulitiple line strings
	input_string := fileString
	old1 := "李清雨"
	new1 := "韩筱雪"
	old2 := "清雨"
	new2 := "筱雪"
	n := -1

	new_string := strings.Replace(input_string, old1, new1, n)

	new_string2 := strings.Replace(new_string, old2, new2, n)

	fmt.Println(new_string2)
}
