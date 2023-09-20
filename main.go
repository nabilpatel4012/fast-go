package main

import (
	"fmt"
)

func main() {

  var myVar = "Hello"
  myVar2 := "Hii"
  name := "Nabil"
  var num int = 4
  const myConstant = 99
  fmt.Println(myVar)
  fmt.Println(myVar2)
  fmt.Println(myConstant)
  fmt.Printf("Hello World!\n")

  printMe()
  myFunctionWithInput(myVar)
  squaredValue := myFunctionWithReturnType(num)
  fmt.Println(squaredValue)
  myFunctionWithMultipleReturnType(num,name)

}


//Functions

//Without return-type and input arguments
func printMe() {
  fmt.Println("I am printed!!!")
}

func myFunctionWithInput(value string) {
  fmt.Println(value)
}

func myFunctionWithReturnType(value int) int {
  var square = value*value
  return square
}

func myFunctionWithMultipleReturnType(value int, name string) (int, string) {
  var square = value*value
  myMessage := "Hello Nabil"
  return square,myMessage
}

