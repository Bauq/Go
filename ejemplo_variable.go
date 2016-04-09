package main

import "fmt"

func main() {
  fmt.Print("Ingrese un texto: ")
  var input string
  fmt.Scanf("%s", &input)

  fmt.Println(input)
}
