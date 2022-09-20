package main

import (
	"base58"
	"crypto/ripemd160"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// 计算sha256
func calSha256(s []byte) []byte {
	hashBytes := sha256.Sum256(s)
	hashString := hashBytes[:]
	return hashString
}

// 计算ripemd160
func calRipemd160(s []byte) []byte {
	a := ripemd160.New()
	a.Write(s)
	b := a.Sum(nil)
	return b
}

func main() {

	var version byte = 111
	var publicKey string
	fmt.Scanf("%s", &publicKey)

	key, _ := hex.DecodeString(publicKey)

	// HASH160
	c := calSha256(key)
	fmt.Printf("%x\n", c)
	fingerprint := calRipemd160(c)

	var version_byte []byte
	version_byte = append(version_byte, version)
	str := append(version_byte, fingerprint...)

	// HASH256
	checkSum := calSha256(calSha256(str))
	result := append(str, checkSum[0:4]...)

	// base58
	address := base58.Encode(result, base58.BitcoinAlphabet)
	fmt.Print(address)
}
