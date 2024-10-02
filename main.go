package main 

import (
	"fmt"
	"github.com/Kassimali/cryptit/encrypt"
	"github.com/Kassimali/cryptit/decrypt"
)

func main () {
	encryptedmsg := encrypt.Numbus("Testing encryption")
	fmt.Println(encryptedmsg)
	fmt.Println(decrypt.Decrypt(encryptedmsg))
}