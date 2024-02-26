package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/relucant1/go-encrypt-picture/filecrypt"
	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("run excrypt to encrypt a file")
		os.Exit(1)

	}

}
func printHelp() {
	fmt.Println("file encryption")
}
func encryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("missing the path to ghe file")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")

	}
	password := getPassword()
	fmt.Println("\n Encrypting")
	filecrypt.Encrypt(file, password)
	fmt.Println("\nfile successfully protected")
}

func decryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("missing the path to ghe file")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")

	}
	fmt.Print("enter password:")
	password, _ := term.ReadPassword(0)
	fmt.Println("\n Dcrypting ...")
	filecrypt.Decrypt(file, password)
	fmt.Println("\nfile successfully protected")
}

func getPassword() []byte {
	fmt.Print("enter password")
	password, _ := term.ReadPassword(0)
	fmt.Print("confirm password")
	password2, _ := term.ReadPassword(0)
	if !validatePassword(password, password2) {
		fmt.Print("\nPassword do not match, please try again")
		return getPassword()
	}
	return password
}

func validatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	}
	return true

}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
