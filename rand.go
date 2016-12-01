package stringEncrypt

import (
	"crypto/rand"
	"encoding/base64"
)

//
// Code taken from http://stackoverflow.com/questions/32349807/how-can-i-generate-a-random-int-using-the-crypto-rand-package
//
// This was almost entirely taken from elithrar's answer on stackoverflow.
//

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.StdEncoding.EncodeToString(b), err
}

// GenerateKey will return a base64 representation of a 32byte string
// suitable for use as a key
func GenerateKey() (string, error) {
	return generateRandomString(32)
}
