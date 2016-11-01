package stringEncrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

//
// Code taken from http://stackoverflow.com/questions/18817336/golang-encrypting-a-string-with-aes-and-base64
//
// This is all rewritten from Intermernet's answer on the above stackoverflow page.
//

// Decrypt will take secret key string and encrypted string and return clear string
func Decrypt(key string, text string) (string, error) {
	rawkey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}
	dectext, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}
	btext := []byte(dectext)
	cipherblock, err := aes.NewCipher([]byte(rawkey))
	if err != nil {
		return "", err
	}
	if len(btext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := btext[:aes.BlockSize]
	btext = btext[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(cipherblock, iv)
	cfb.XORKeyStream(btext, btext)
	data, err := base64.StdEncoding.DecodeString(string(btext))
	if err != nil {
		return "", err
	}
	return string(data), nil
}
