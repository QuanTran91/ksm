package aes

import (
	"github.com/QuanTran91/toast/cipher"
	"github.com/QuanTran91/toast/crypto"
)

func EncryptWithECB(key []byte, plainText []byte) ([]byte, error) {
	mode := cipher.NewECBMode()
	cipher, err := crypto.NewAESWith(key, mode)
	if err != nil {
		return nil, err
	}

	ciphertext := cipher.Encrypt(plainText)
	return ciphertext, nil
}
