package encrypt

func Nimbus(str string) string {
	encryptedString := ""
	for _, ch := range str {
		encryptedString += string(ch + 3)
	}

	return encryptedString
}
