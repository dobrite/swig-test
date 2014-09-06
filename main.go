package main

import (
	"fmt"
	"github.com/dobrite/swig-test/simplelib"
)

func main() {

	simpleClass := simplelib.NewSimpleClass()
	result := simpleClass.Hello()
	fmt.Println(result)

	strings := simplelib.NewStringVector()
	simpleClass.HelloString(strings)

	fmt.Println(strings.Size())
	for i := 0; int64(i) < strings.Size(); i++ {
		fmt.Println("s")
		fmt.Println(strings.Get(i))
	}

	bytes := simplelib.NewByteVector()
	simpleClass.HelloBytes(bytes)

	fmt.Println(bytes.Size())
	for i := 0; int64(i) < bytes.Size(); i++ {
		fmt.Printf("%s", string(bytes.Get(i)))
	}

}
