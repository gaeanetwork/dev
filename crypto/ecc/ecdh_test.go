package ecc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/gaeanetwork/gaea-core/crypto"
	"github.com/gaeanetwork/gaea-core/crypto/keyagreement"
	"github.com/stretchr/testify/assert"
)

var (
	privHexForTests = "308187020100301306072a8648ce3d020106082a8648ce3d030107046d306b02010104202d130ea6dac76fcae718fbd20bf146643aa66fe6e5902975d2c5ed6ab3bcb5e2a144034200048f03f8321b00a4466f4bf4be51c91898cd50d8cc64c6ecf53e73443e348d5925a16f88c8952b78ebac2dc277a2cc54c77b4c3c07830f49629b689edf63086293"
	pubHexForTests  = "048f03f8321b00a4466f4bf4be51c91898cd50d8cc64c6ecf53e73443e348d5925a16f88c8952b78ebac2dc277a2cc54c77b4c3c07830f49629b689edf63086293"
)

func Test_ECDH(t *testing.T) {
	priv1, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)
	priv2, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)
	assert.NotEqual(t, priv1.D, priv2.D)

	ecdh := &ECDH{}
	secretKey1, err := ecdh.GenerateSharedSecret(priv1, &priv2.PublicKey)
	assert.NoError(t, err)
	secretKey2, err := ecdh.GenerateSharedSecret(priv2, &priv1.PublicKey)
	assert.NoError(t, err)
	assert.Equal(t, secretKey1, secretKey2)
}

func Test_ECDH_DifferentCurve(t *testing.T) {
	priv1, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)
	priv3, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	assert.NoError(t, err)

	ecdh := &ECDH{}
	secretKey1, err := ecdh.GenerateSharedSecret(priv1, &priv3.PublicKey)
	assert.NoError(t, err)
	secretKey2, err := ecdh.GenerateSharedSecret(priv3, &priv1.PublicKey)
	assert.NoError(t, err)
	assert.NotEqual(t, secretKey1, secretKey2)
}

func Test_ECDH_InvalidPriv(t *testing.T) {
	priv1, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)

	// point type error
	ecdh := &ECDH{}
	_, err = ecdh.GenerateSharedSecret(*priv1, &priv1.PublicKey)
	assert.Error(t, err)

	// invalid private key error
	priv2, err := rsa.GenerateKey(rand.Reader, 32)
	assert.NoError(t, err)
	_, err = ecdh.GenerateSharedSecret(priv2, &priv1.PublicKey)
	assert.Error(t, err)
}

func Test_ECDH_InvalidPub(t *testing.T) {
	priv1, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)

	// point type error
	ecdh := &ECDH{}
	_, err = ecdh.GenerateSharedSecret(priv1, priv1.PublicKey)
	assert.Error(t, err)

	// invalid private key error
	priv2, err := rsa.GenerateKey(rand.Reader, 32)
	assert.NoError(t, err)
	_, err = ecdh.GenerateSharedSecret(priv1, &priv2.PublicKey)
	assert.Error(t, err)
}

func Test_GetAlgorithm(t *testing.T) {
	ecdh := &ECDH{}
	assert.Equal(t, keyagreement.ECDH, ecdh.GetAlgorithm())
}

func Test_Android_ECDH(t *testing.T) {
	androidPubKey := "04f2ca2888417bac66b5e7bcdbcbaefe1771f45e8ac29eef23ddc84157ab16e005bd7ca457632658220a6aa721d326961e4014dae8c789030c82640bd083f3daae"
	pubBytes, err := hex.DecodeString(androidPubKey)
	if err != nil {
		t.Fatal(err)
	}

	x, y := elliptic.Unmarshal(elliptic.P256(), pubBytes)
	pubkey := &ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}

	privBytes, err := hex.DecodeString(privHexForTests)
	if err != nil {
		t.Fatal(err)
	}

	priv, err := FromPrivBytes(privBytes)
	if err != nil {
		t.Fatal(err)
	}

	ecdh := &ECDH{}
	secretkey, err := ecdh.GenerateSharedSecret(priv, pubkey)
	assert.NoError(t, err)
	assert.Equal(t, "c987c7fd47a40da41f8cf5c485b78c38a2052409e8fd6aeac185cbc02c9117f7", hex.EncodeToString(secretkey))

	data := []byte("Hello World!")
	ciphertext, err := crypto.AesEncrypt(data, secretkey)
	assert.NoError(t, err)
	assert.Equal(t, "fb0a9a743252b042db5b918346fb0bdc", hex.EncodeToString(ciphertext))
}

func Test_SHA256(t *testing.T) {
	data := []byte("04be8ac2b0cc27d92b102b7fa25fc2d5aeb9ea5c4dfb88c74d4f8532c1ece317c8a47c6f7232f676c6c1ec46b8ab2a6687c7575b9892ae815a5f84248a946564f2")
	hash := sha256.Sum256(data)
	assert.Equal(t, "b088cf414cbab06fff85602bbc27a3e24c96a757ee29c78c48b9eaa198686a12", hex.EncodeToString(hash[:]))
}

func Test_JSECDH(t *testing.T) {
	jsPubKey := "0495d820c781208cdd0389ea0d54cdbe1b64c0d91e9541816c55608057e5be07a81cb46eea388ee05a30107fd512705301b8e6b0d83284d54d1babb63aafc879b4"
	pubBytes, err := hex.DecodeString(jsPubKey)
	if err != nil {
		t.Fatal(err)
	}

	x, y := elliptic.Unmarshal(elliptic.P256(), pubBytes)
	pubkey := &ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}

	privBytes, err := hex.DecodeString(privHexForTests)
	if err != nil {
		t.Fatal(err)
	}

	priv, err := FromPrivBytes(privBytes)
	if err != nil {
		t.Fatal(err)
	}

	ecdh := &ECDH{}
	secretkey, err := ecdh.GenerateSharedSecret(priv, pubkey)
	assert.NoError(t, err)
	assert.Equal(t, "a514b562b21ce3075148fbedbd6cc0ead9b33094488b701ca835a9ec3c9f81d8", hex.EncodeToString(secretkey))

	// js private key
	d := "d233a716bf371afc597636a9b00342603759ab9f39ab5954e6d51a996cd2bfdd"
	b, _ := hex.DecodeString(d)
	D := big.NewInt(0).SetBytes(b)
	privkey := &ecdsa.PrivateKey{
		PublicKey: *pubkey,
		D:         D,
	}
	pub, err := FromPubHex(pubHexForTests)
	if err != nil {
		panic(err)
	}
	jsx, _ := pub.ScalarMult(pub.X, pub.Y, privkey.D.Bytes())
	assert.Equal(t, "99629494961789099446600506511571181974131020151128428048581066925321839516601", jsx.String())
	jssecretkey := sha256.Sum256([]byte(jsx.String()))
	assert.Equal(t, "a514b562b21ce3075148fbedbd6cc0ead9b33094488b701ca835a9ec3c9f81d8", hex.EncodeToString(jssecretkey[:]))

}
