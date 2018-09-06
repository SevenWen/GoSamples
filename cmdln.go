package main

import (
	"fmt"
	"flag" //package to parse user input
	"os"
)

func main() {
	fmt.Printf("It's a simple command line echo\n")

	rawinput:=os.Args//get the raw user input as a set
	fmt.Println(rawinput)

	var userInput string
	flag.StringVar(&userInput,"input","...\n","user input a string") //define a flag
	flag.Parse()//parse user input
	fmt.Printf(userInput)
}
