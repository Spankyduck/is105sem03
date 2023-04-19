package main

import (
	"fmt"

	"github.com/Spankyduck/is105sem03/mycrypt"
)

func main() {

	input := "Kjevik;SN39040;18.03.2022 01:50;6"
	key := mycrypt.ALF_SEM03
	keyLength := 4

	// Encrypt
	encryptedRune := mycrypt.Krypter([]rune(input), key, keyLength)
	encryptedString := string(encryptedRune)
	fmt.Println("Encrypted:", encryptedString)

	// Decrypt
	decryptedRune := mycrypt.Krypter(encryptedRune, key, len(key)-keyLength)
	decryptedString := string(decryptedRune)
	fmt.Println("Decrypted:", decryptedString)

}
