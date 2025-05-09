package main

import (
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
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(service)
	item.SetAccount(account)
	item.SetLabel(label)
	item.SetData([]byte(value))
	item.SetAccessible(keychain.AccessibleWhenUnlocked)
	item.SetSynchronizable(keychain.SynchronizableNo)
	err := keychain.AddItem(item)
	if err != nil {
		fmt.Println("Error storing secret:", err)
		return
	}
}
