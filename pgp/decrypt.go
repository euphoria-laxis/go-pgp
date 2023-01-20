package pgp

import (
	"bytes"
	_ "crypto/sha256"
	"errors"
	"fmt"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	_ "golang.org/x/crypto/ripemd160"
	"io/ioutil"
)

func Decrypt(entity *openpgp.Entity, encrypted []byte) ([]byte, error) {
	// Decode message
	block, err := armor.Decode(bytes.NewReader(encrypted))
	if err != nil {
		return []byte{}, fmt.Errorf("Error decoding: %v", err)
	}
	if block.Type != "Message" {
		return []byte{}, errors.New("Invalid message type")
	}

	// Decrypt message
	entityList := openpgp.EntityList{entity}
	messageReader, err := openpgp.ReadMessage(block.Body, entityList, nil, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("Error reading message: %v", err)
	}
	var read []byte
	read, err = ioutil.ReadAll(messageReader.UnverifiedBody)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading unverified body: %v", err)
	}

	// Return output - an unencrypted, and uncompressed message
	return read, nil
}
