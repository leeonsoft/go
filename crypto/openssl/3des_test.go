package openssl

import (
	"encoding/base64"
	"testing"

	"github.com/yangtizi/go/crypto/assert"
)

func TestDes3Encrypt(t *testing.T) {
	src := []byte("123456")

	// 3DES-ECB, PKCS7PADDING
	key := []byte("123456789012345678901234")
	dst, err := Des3ECBEncrypt(src, key, PKCS7PADDING)
	assert.NoError(t, err)
	t.Log(base64.StdEncoding.EncodeToString(dst))
	assert.Equal(t, base64.StdEncoding.EncodeToString(dst), "zvq8DAMSQaM=")
}

func TestDes3ECBDecrypt(t *testing.T) {
	src, err := base64.StdEncoding.DecodeString("zvq8DAMSQaM=")
	assert.NoError(t, err)

	// 3DES-ECB, PKCS7PADDING
	key := []byte("123456789012345678901234")
	dst, err := Des3ECBDecrypt(src, key, PKCS7PADDING)
	assert.NoError(t, err)
	t.Log(string(dst))
	assert.Equal(t, dst, []byte("123456"))
}

func TestDes3CBCEncrypt(t *testing.T) {
	src := []byte("123456")
	iv := []byte("67890678")
	// 3DES-ECB, PKCS7PADDING
	key := []byte("123456789012345678901234")
	dst, err := Des3CBCEncrypt(src, key, iv, PKCS7PADDING)
	assert.NoError(t, err)
	t.Log(base64.StdEncoding.EncodeToString(dst))
	assert.Equal(t, base64.StdEncoding.EncodeToString(dst), "elcFR372FXU=")
}

func TestDes3CBCDecrypt(t *testing.T) {
	src, err := base64.StdEncoding.DecodeString("elcFR372FXU=")
	assert.NoError(t, err)

	iv := []byte("67890678")

	// 3DES-ECB, PKCS7PADDING
	key := []byte("123456789012345678901234")
	dst, err := Des3CBCDecrypt(src, key, iv, PKCS7PADDING)
	assert.NoError(t, err)
	t.Log(string(dst))
	assert.Equal(t, dst, []byte("123456"))
}
