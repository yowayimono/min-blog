package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"log"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

type encrypt struct{}

var Encryptor = new(encrypt)

// 使用 scrypt 对密码进行加密生成一个哈希值
func (*encrypt) ScryptHash(pasword string) string {
	const KeyLen = 10
	salt := []byte{12, 32, 4, 6, 66, 22, 222, 11} // 随便写

	hashPwd, err := scrypt.Key([]byte(pasword), salt, 1<<15, 8, 1, KeyLen)
	if err != nil {
		log.Fatal("加密失败: ", err)
	}
	return base64.StdEncoding.EncodeToString(hashPwd)
}

// 使用 scrypt 对比 明文密码 和 数据库中哈希值
func (c *encrypt) ScryptCheck(password, hash string) bool {
	return c.ScryptHash(password) == hash
}

// 使用 bcrypt 对密码进行加密生成一个哈希值
func (*encrypt) BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// 使用 bcrypt 对比 明文密码 和 数据库中哈希值
func (*encrypt) BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// MD5 加密
func (*encrypt) MD5(str string, b ...byte) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(b))
}

// TODO
func UUID() string {
	return ""
}