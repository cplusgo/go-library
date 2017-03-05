package main

import (
	"fmt"
	"github.com/cplusgo/go-library"
)

type Demo struct {
	name string
}

func (demo Demo) ToString2() {

}

func (demo Demo) ToString() {

}

func main()  {
	demo := &Demo{}
	fmt.Println(go_library.HasMethod(*demo, "ToString"))
}
