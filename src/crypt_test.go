package main

import (
	"fmt"
	"internal/crypt"
)

func main() {
	fmt.Printf("client main!\n")
	data := "kylewbanks.com"
	key := "KCQ"
	e := crypt.Encrypt(data, key)
	fmt.Printf(e + "\n")
	d := crypt.Decrypt(e, key)
	fmt.Printf(d + "\n")
}
