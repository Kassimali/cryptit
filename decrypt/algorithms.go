package decrypt 

func Decrypt(str string) string{

	decryptedString := ""
	for _, ch := range str {
		ascii_code := int(ch)
		character := string(ascii_code -3)
		decryptedString += character
	}

	return decryptedString
}