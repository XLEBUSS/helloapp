package main

import "fmt"

func main(){
	Hello()
	greetBlack()
	greetWhite()
}

func Hello(){
	fmt.Println("hello everyone")
}

func greetWhite(){
	fmt.Println("hello white")
}

func greetBlack(){
	fmt.Println("hello black")
}