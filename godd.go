package godd

import (
	"fmt"
	"os"
	"reflect"
)

func DD(v interface{}) {
	t := reflect.TypeOf(v)
	
	if t == nil {
		printInterface()
		os.Exit(0)
	}
	
	k := t.Kind()
	val := reflect.ValueOf(v)
	kf := fmt.Sprintf("%v", k)
	
	if kf == "struct" {
		printType(t)
		printBaseType(k)
		printValue(val)
		os.Exit(0)
	} else if kf == "func" {
		printType(t)
		printBaseType(k)
		os.Exit(0)
	} else if kf == "chan" {
		printType(t)
		printBaseType(k)
		printValue(val)
		os.Exit(0)
	} else {
		printType(t)
		printValue(val)
		os.Exit(0)
	}
}

func printInterface() {
	fmt.Printf("#Type: interface")
}

func printType(t interface{}) {
	fmt.Printf("#Type: %v\n", t)
}

func printBaseType(t interface{}) {
	fmt.Printf("#Underlying Type: %v\n", t)
}

func printValue(v interface{}) {
	fmt.Printf("#Value: %v\n", v)
}
