package md5

import (
	mmdd55 "crypto/md5"
	"fmt"
)

// ToBin 计算MD5的16进制值
func ToBin(str string) [16]byte {
	data := []byte(str)
	has := mmdd55.Sum(data)
	return has
}

// ToHex 计算小写md5
func ToHex(str string) string {
	data := []byte(str)
	has := mmdd55.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

// ToHEX 计算大写md5
func ToHEX(str string) string {
	data := []byte(str)
	has := mmdd55.Sum(data)
	md5str := fmt.Sprintf("%X", has)
	return md5str
}
