package decrypt

func Nimbus(str string) string {
	decryptedString := ""
	for _, ch := range str {
		decryptedString += string(ch + 3)
	}

	return decryptedString
}
