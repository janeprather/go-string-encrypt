package stringEncrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

//
// Code taken from http://stackoverflow.com/questions/18817336/golang-encrypting-a-string-with-aes-and-base64
//
// This is all rewritten from Intermernet's answer on the above stackoverflow page.
//

// Encrypt takes secret encryption key string and text string, returns encrypted string
func Encrypt(key string, text string) (string, error) {
	rawkey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}
	cipherblock, err := aes.NewCipher([]byte(rawkey))
	if err != nil {
		return "", err
	}
	b64text := base64.StdEncoding.EncodeToString([]byte(text))
	ciphertext := make([]byte, aes.BlockSize+len(b64text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(cipherblock, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b64text))
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
