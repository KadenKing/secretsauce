package main

import (
	"fmt"

	"github.com/KadenKing/secretsauce/complicated"
)

func main() {
	// does not compile
	// bad := complicated.ComplicatedType{}

	good := complicated.NewComplicatedType()
	good.DoComplicatedThings()
	fmt.Println("everything went okay")
}
