# Go PGP

Created from [https://github.com/jchavannes/go-pgp.git](https://github.com/jchavannes/go-pgp.git).

The project was updated to be used as a go package. A part of the code
was changed and the tests were updated.

Layer on top of `golang.org/x/crypto/openpgp` to handle a few PGP use 
cases.

## Examples

### Encryption

[pgp/encrypt_test.go](pgp/encrypt_test.go)

#### Encrypt

```go
// Create public key entity
publicKeyPacket, _ := pgp.GetPublicKeyPacket([]byte(TestPublicKey))
pubEntity, _ := pgp.CreateEntityFromKeys(publicKeyPacket, nil)

// Encrypt message
encrypted, _ := pgp.Encrypt(pubEntity, []byte(TestMessage))
```

#### Decrypt

```go
// Create private key entity
privEntity, _ := pgp.GetEntity([]byte(TestPublicKey), []byte(TestPrivateKey))

// Decrypt message
decrypted, _ := pgp.Decrypt(privEntity, encrypted)
```

### Signing

[pgp/sign_test.go](pgp/sign_test.go)

#### Sign

```go
// Create private key entity
entity, _ := pgp.GetEntity([]byte(TestPublicKey), []byte(TestPrivateKey))

// Sign message
signature, _ := pgp.Sign(entity, []byte(TestMessage))
```

#### Verify

```go
// Create public key packet
pubKeyPacket, _ := pgp.GetPublicKeyPacket([]byte(TestPublicKey))

// Verify signature
err = pgp.Verify(pubKeyPacket, []byte(TestMessage), signature)
if err == nil {
    fmt.Println("Signature verified.")
}
```

## Credits

**Jason Chavannes <jason.chavannes@gmail.com>** for the original
repository I forked to create this package.

## License

This project is under [MIT License](./LICENSE.md) and can be used for
your projects.
