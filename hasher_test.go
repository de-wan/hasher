package hasher

import (
	"testing"
)

func TestHashValue(t *testing.T) {
	hashByte, err := HashValue([]byte("username"))
	if err != nil {
		t.Errorf(err.Error())
	}

	hashStr := string(hashByte)
	t.Log("Hash: " + hashStr)
	if hashStr == "" {
		t.Errorf("Error: Returned empty hash")
	}
}

func TestVerifyHash(t *testing.T) {
	secret := "username"
	hash := "tY/J59lpmvydnd59PM0+FKqXPLCP3j/NDsydesy3kWY="

	isValid, err := VerifySecret([]byte(secret), []byte(hash))
	if err != nil {
		t.Error(err)
	}
	if isValid == false {
		t.Error("Wrong secret or hash")
	}
}
