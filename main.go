package main

import (
	"fmt"
	"math/rand"
	"myapp/mascot"
	"os"
	"path"
	"time"

	"rsc.io/quote"
)

func main() {
	fmt.Println(mascot.Best())
	fmt.Println(quote.Go())
	fmt.Println("ABC")
	var a int
	var s string = "abc"
	fmt.Println(a)
	fmt.Println(s)

	fmt.Println("ABC", "CDE")

	fmt.Print("ABC", "CDE ")
	var ac = false
	if ac {
		fmt.Println("True  is")
	} else {
		fmt.Println("false is ")
	}
	age := 35
	if age >= 62 {
		fmt.Println("You are retiring")
	}
	if age >= 30 && age <= 40 {
		fmt.Println("Promotion is due")
	}
	rand.Seed(time.Now().UnixNano())
	_, file := path.Split("css/main.css")
	fmt.Println(file)
	fmt.Printf("%#v\n", os.Args)
	fmt.Println(len(os.Args))
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])
}