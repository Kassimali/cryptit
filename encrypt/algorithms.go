package encrypt 

func Numbus(str string) string {

	encryptedStr := ""
	for _, ch := range str{
		ascii_code := int(ch)
		character := string(ascii_code + 3)
		encryptedStr += character
	}
	return encryptedStr
}