package main

import (
	"fmt"
	"reflect"
)

type triangle struct {
	height float64
	base   float64
}
type square struct {
	height float64
	width  float64
}

type shape interface {
	getObjArea() (string, float64)
}

func main() {
	tri := triangle{
		height: 2.5,
		base:   3.5,
	}
	sq := square{
		height: 3.5,
		width:  3.5,
	}

	printArea(tri)
	printArea(sq)
}

func printArea(s shape) {
	o, a := s.getObjArea()
	fmt.Println("Shape is:", o, "and area is:", a)
}

func (t triangle) getObjArea() (string, float64) {
	e := reflect.ValueOf(&t).Elem()
	return e.Type().Name(), 0.5 * t.base * t.height
}

func (s square) getObjArea() (string, float64) {
	e := reflect.ValueOf(&s).Elem()
	return e.Type().Name(), s.height * s.width
}
