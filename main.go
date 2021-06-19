package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Triangle byte

const (
	Equilateral Triangle = 1 << iota
	Isosceles
	Scalene
	Invalid
)

func (triangle Triangle) String() string {
	return [...]string{"Equilateral", "Isosceles", "Scalene", "Invalid"}[triangle]
}

func isExists(a uint32, b uint32, c uint32) bool {
	a64 := uint64(a)
	b64 := uint64(b)
	c64 := uint64(c)
	return (a64 + b64 > c64) && (a64 + c64 > b64) && (b64 + c64 > a64)
}

func isEquilateral(a uint32, b uint32, c uint32) bool {
	return (a == b) && (b == c)
}

func isIsosceles(a uint32, b uint32, c uint32) bool {
	return (a == b) || (a == c) || (b == c)
}

func isScalene(a uint32, b uint32, c uint32) bool {
	return (a != b) && (b != c) && (a != c)
}

func getTriangleType(a uint32, b uint32, c uint32) Triangle {
	if isExists(a, b, c) {
		if isEquilateral(a, b, c) {
			return Equilateral
		} else if isIsosceles(a, b, c) {
			return Isosceles
		} else if isScalene(a, b, c){
			return Scalene
		}
	}
	return Invalid
}

func triangleSideInput(name string) (result uint32) {
	for result == 0{
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Enter side %s: \n", name)
		scanner.Scan()
		if input, err := strconv.ParseInt(scanner.Text(), 10, 64); err == nil && input > 0 && input <= math.MaxUint32{
			result = uint32(input)
		} else {
			fmt.Printf("Length of side should be [1 .. %d]\n", math.MaxUint32)
		}
	}
	return
}

func main() {
	fmt.Println("Hello! This program tells a type of triangle based on it's 3 sides")
	fmt.Printf("Length of sides should be [1 .. %d]\n", math.MaxUint32)
	a := triangleSideInput("A")
	b := triangleSideInput("B")
	c := triangleSideInput("C")
	triangleType := getTriangleType(a, b, c)
	fmt.Println(triangleType)
	if triangleType == Invalid{
		fmt.Println("(Valid has (A + B > C) and (A + C > B) and (B + C > A))")
	}
}