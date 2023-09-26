package main

import (
	"fmt"

	"github.com/ankit-verma-209171/cryptit/decrypt"
	"github.com/ankit-verma-209171/cryptit/encrypt"
)

func main() {
	encyrptedStr := encrypt.Nimbus("zabcd")
	fmt.Println(encyrptedStr)
	fmt.Println(decrypt.Nimbus(encyrptedStr))
}
