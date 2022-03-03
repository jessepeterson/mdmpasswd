package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/groob/plist"
	"github.com/micromdm/micromdm/pkg/crypto/password"
)

func main() {
	var (
		flB64      = flag.Bool("b64", false, "Output base64-encoded Plist")
		flPassword = flag.String("password", "", "Password to hash")
	)
	flag.Parse()
	if *flPassword == "" {
		log.Fatal(errors.New("no password supplied"))
	}
	p, err := password.SaltedSHA512PBKDF2(*flPassword)
	if err != nil {
		log.Fatal(err)
	}
	hash := struct {
		SaltedSHA512PBKDF2 password.SaltedSHA512PBKDF2Dictionary `plist:"SALTED-SHA512-PBKDF2"`
	}{
		SaltedSHA512PBKDF2: p,
	}
	plist, err := plist.MarshalIndent(hash, "  ")
	if err != nil {
		log.Fatal(err)
	}
	if *flB64 {
		fmt.Print(base64.StdEncoding.EncodeToString(plist))
	} else {
		fmt.Print(string(plist))
	}
}
