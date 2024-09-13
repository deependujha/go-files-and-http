package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filname string) {
	databyte, err := os.ReadFile(filname)
	checkNilErr(err)

	fmt.Println("Text data inside the file is: `", string(databyte), "`")

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
