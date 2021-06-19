package main

import (
	"math"
	"testing"
)

func TestIsEquilateral(t *testing.T){
	var tests = []struct {
		a uint32
		b uint32
		c uint32
		result bool
	}{
		{1, 1, 1, true},
		{5000, 5000, 5000, true},
		{1000000, 1000000, 1000000, true},
		{1, 2, 2, false},
		{2, 1, 2, false},
		{2, 2, 1, false},
	}
	for _, test := range tests{
		if output := isEquilateral(test.a, test.b, test.c); output != test.result{
			t.Errorf("Test failed: %d,%d,%d inputted, %t expected, received: %t", test.a, test.b, test.c, test.result, output)
		}
	}
}

func TestIsIsosceles(t *testing.T){
	var tests = []struct {
		a uint32
		b uint32
		c uint32
		result bool
	}{
		{1, 1, 1, true},
		{5000, 5000, 5000, true},
		{1000000, 1000000, 1000000, true},
		{1, 2, 2, true},
		{2, 1, 2, true},
		{2, 2, 1, true},
		{1, 2, 3, false},
	}
	for _, test := range tests{
		if output := isIsosceles(test.a, test.b, test.c); output != test.result{
			t.Errorf("Test failed: %d,%d,%d inputted, %t expected, received: %t", test.a, test.b, test.c, test.result, output)
		}
	}
}

func TestIsScalene(t *testing.T){
	var tests = []struct {
		a uint32
		b uint32
		c uint32
		result bool
	}{
		{1, 1, 1, false},
		{5000, 5000, 5000, false},
		{1000000, 1000000, 1000000, false},
		{1, 2, 2, false},
		{2, 1, 2, false},
		{2, 2, 1, false},
		{1, 2, 3, true},
	}
	for _, test := range tests{
		if output := isScalene(test.a, test.b, test.c); output != test.result{
			t.Errorf("Test failed: %d,%d,%d inputted, %t expected, received: %t", test.a, test.b, test.c, test.result, output)
		}
	}
}

func TestIsExists(t *testing.T){
	var tests = []struct {
		a uint32
		b uint32
		c uint32
		result bool
	}{
		{1, 1, 1, true},
		{5000, 5000, 5000, true},
		{1000000, 1000000, 1000000, true},
		{1, 2, 2, true},
		{2, 1, 2, true},
		{2, 2, 1, true},
		{1, 2, 3, false},
	}
	for _, test := range tests{
		if output := isExists(test.a, test.b, test.c); output != test.result{
			t.Errorf("Test failed: %d,%d,%d inputted, %t expected, received: %t", test.a, test.b, test.c, test.result, output)
		}
	}
}

func TestGetTriangleType(t *testing.T){
	var tests = []struct {
		a uint32
		b uint32
		c uint32
		result Triangle
	}{
		{0, 0, 0, Invalid},
		{0, 1, 0, Invalid},
		{0, 0, 1, Invalid},
		{2, 1, 1, Invalid},
		{1, 2, 1, Invalid},
		{1, 1, 2, Invalid},
		{1, 1, 1, Equilateral},
		{5000, 5000, 5000, Equilateral},
		{1000000, 1000000, 1000000, Equilateral},
		{math.MaxUint32/2, math.MaxUint32/2, math.MaxUint32/2, Equilateral},
		{math.MaxUint32, math.MaxUint32, math.MaxUint32, Equilateral},
		{1, 2, 2, Isosceles},
		{2, 1, 2, Isosceles},
		{2, 2, 1, Isosceles},
		{1, math.MaxUint32, math.MaxUint32, Isosceles},
		{math.MaxUint32, 1, math.MaxUint32, Isosceles},
		{math.MaxUint32, math.MaxUint32, 1, Isosceles},
		{4, 2, 3, Scalene},
		{400000, 200000, 300000, Scalene},
		{400000000, 200000000, 300000000, Scalene},
		{math.MaxUint32, math.MaxUint32 - 1, math.MaxUint32 - 2, Scalene},
		{100, math.MaxUint32 - 1, math.MaxUint32 - 3, Scalene},
	}
	for _, test := range tests{
		if output := getTriangleType(test.a, test.b, test.c); output != test.result{
			t.Errorf("Test failed: %d,%d,%d inputted, %q expected, received: %q", test.a, test.b, test.c, test.result, output)
		}
	}
}