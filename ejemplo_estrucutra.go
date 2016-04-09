package main

import "fmt"

type Persona struct {
	nombre   string
	estatura float64
}

func (p *Persona) correr() string {
	return "corriendo"
}
func main() {
	p := Persona{"Brian", 1.77}
	fmt.Println(p.nombre, p.correr())
}