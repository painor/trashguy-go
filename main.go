package main

import (
	"./src"
	"fmt"
)

func main() {
	trashy := src.TrashGuy{"a b c", src.DEFAULT_OPTIONS, src.Sprite{}, src.Slice{}, -1, -1, -1}
	trashy.Init()
	fmt.Printf(trashy.Animate())

}
