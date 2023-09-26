package decrypt

func Nimbus(str string) string {
	decryptedStr := ""
	for _, ch := range str {
		asciiCode := int(ch)
		character := string(asciiCode - 3)
		decryptedStr += character
	}
	return decryptedStr
}
