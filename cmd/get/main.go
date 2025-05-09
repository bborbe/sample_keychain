package main

import (
	"fmt"
	"github.com/keybase/go-keychain"
)

const (
	service = "myapp"
	account = "user@example.com"
)

func main() {
	// Retrieve a secret
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetService(service)
	query.SetAccount(account)
	query.SetReturnData(true)
	results, err := keychain.QueryItem(query)
	if err != nil {
		fmt.Println("Error retrieving secret:", err)
		return
	}
	if len(results) > 0 {
		fmt.Println("Retrieved secret:", string(results[0].Data))
	} else {
		fmt.Println("No secret found")
	}
}
