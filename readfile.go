package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	filename:="./plaintext.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("File contents: %s", content)
}