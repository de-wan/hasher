package hasher

import (
	"bytes"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func HashValue(valueToHash []byte) (hash []byte, err error) {
	salt := []byte("secret-salt")
	hashByte := argon2.IDKey(valueToHash, salt, 1, 64*1024, 1, 32)

	var hashBase64 = make([]byte, base64.StdEncoding.EncodedLen(len(hashByte)))
	base64.StdEncoding.Encode(hashBase64, hashByte)

	hash = hashBase64

	return hash, err
}

func VerifySecret(secret []byte, hash []byte) (bool, error) {
	hashedSecret, err := HashValue(secret)
	if err != nil {
		return false, err
	}

	return bytes.Equal(hashedSecret, hash), nil
}
