package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessageEncryptDecrypt(t *testing.T) {
	message := NewMessage()
	const msg = "hey-ho-ho-ho"
	message.Plaintext = msg
	encryptedMessage, err := message.Encrypt("secret")
	if assert.NoError(t, err) {
		assert.NotEqual(t, message.Plaintext, encryptedMessage)
		assert.Equal(t, message.Encrypted, encryptedMessage)
	}

	decryptedMessage, err := message.Decrypt("secret")
	if assert.NoError(t, err) {
		assert.Equal(t, message.Plaintext, decryptedMessage)
		assert.Equal(t, msg, decryptedMessage)
	}
}