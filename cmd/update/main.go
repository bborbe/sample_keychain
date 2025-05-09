package main

import (
	"bytes"
	"fmt"
	"github.com/keybase/go-keychain"
)

const (
	service = "myapp"
	account = "user@example.com"
	label   = "My App Secret"
	value   = "new secret value"
)

func main() {
	queryItem := keychain.NewItem()
	queryItem.SetSecClass(keychain.SecClassGenericPassword)
	queryItem.SetService(service)
	queryItem.SetAccount(account)
	queryItem.SetReturnData(true)
	results, err := keychain.QueryItem(queryItem)
	if err != nil {
		fmt.Println("Error retrieving secret:", err)
		return
	}

	updateItem := keychain.NewItem()
	updateItem.SetSecClass(keychain.SecClassGenericPassword)
	updateItem.SetService(service)
	updateItem.SetAccount(account)
	updateItem.SetLabel(label)
	updateItem.SetData([]byte(value))
	updateItem.SetAccessible(keychain.AccessibleWhenUnlocked)
	updateItem.SetSynchronizable(keychain.SynchronizableNo)
	switch len(results) {
	case 0:
		if err := keychain.AddItem(updateItem); err != nil {
			fmt.Println("Error storing secret:", err)
			return
		}
		fmt.Println("secret created")
		return
	case 1:
		if bytes.Compare(results[0].Data, []byte(value)) == 0 {
			fmt.Println("secret already uptodate")
			return
		}
		if err := keychain.UpdateItem(queryItem, updateItem); err != nil {
			fmt.Println("Error updating secret:", err)
			return
		}
		fmt.Println("secret updated")
		return
	default:
		fmt.Println("Too many results retrieved:", len(results))
		return
	}
}
