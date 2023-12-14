package main

import (
	"fmt"

	"github.com/nilanjanb3/cryptit/decrypt"
	"github.com/nilanjanb3/cryptit/encrypt"
)

func main() {

	text := "Nilanjan"

	encryptedString := encrypt.Nimbus(text)
	fmt.Println(encryptedString)
	fmt.Println(decrypt.Nimbus(encryptedString))

}
