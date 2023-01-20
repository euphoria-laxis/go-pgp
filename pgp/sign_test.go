package pgp_test

import (
	"fmt"
	"github.com/euphoria-laxis/go-pgp/pgp"
	"golang.org/x/crypto/openpgp"
	"testing"
)

func TestSignature(t *testing.T) {
	fmt.Println("Signature test: START")
	entity, err := pgp.GetEntity([]byte(TestPublicKey), []byte(TestPrivateKey))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Created private key entity.")

	var signature []byte
	signature, err = pgp.Sign(entity, []byte(TestMessage))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Created signature of test message with private key entity.")

	var publicKeyEntity *openpgp.Entity
	publicKeyEntity, err = pgp.GetEntity([]byte(TestPublicKey), []byte{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Created public key entity.")

	err = pgp.Verify(publicKeyEntity, []byte(TestMessage), signature)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Signature verified using public key entity.")
	fmt.Println("Signature test: END")
}
