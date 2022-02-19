package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"hrhounkanrin.com/dictionary/dictionary"
)

func main() {
	action := flag.String("action", "list", "Action to perform on the dictionary")
	d, err := dictionary.New("./badger")
	handleError(err)
	defer d.Close()
	flag.Parse()
	switch *action {
	case "list":
		actionList(d)
	case "add":
		actionAdd(d, flag.Args())
	case "define":
		actionDefine(d, flag.Args())
	case "remove":
		actionRemove(d, flag.Args())
	default:
		fmt.Printf("Unknown action: %v\n", *action)
	}
}
func actionList(d *dictionary.Dictionary) {
	words, entries, _ := d.List()
	for _, word := range words {
		fmt.Printf("word: %v - def: %v\n", word, entries[word])
	}
}

func actionRemove(d *dictionary.Dictionary, args []string) {
	if len(args) == 0 {
		handleError(errors.New("No word provided"))
	}
	word := args[0]
	err := d.Remove(word)
	handleError(err)
	fmt.Printf("wor: %v successfuly remove from dictionary\n", word)
}

func actionAdd(d *dictionary.Dictionary, args []string) {
	if len(args) < 2 {
		handleError(errors.New("Some arguments may be missing"))
	}
	word := args[0]
	definition := args[1]
	err := d.Add(word, definition)
	handleError(err)
	fmt.Printf("wor: %v successfuly added to dictionary\n", word)
}

func actionDefine(d *dictionary.Dictionary, args []string) {
	if len(args) == 0 {
		handleError(errors.New("No word provided"))
	}
	word := args[0]
	entry, err := d.Get(word)
	handleError(err)
	fmt.Printf("word: %v - def: %v\n", word, entry)

}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Dictionary error: %v\n", err)
		os.Exit(1)
	}
}
