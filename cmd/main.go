package main

import (
	"os"

	"github.com/taylormonacelli/anyjuly"
)

func main() {
	code := anyjuly.Execute()
	os.Exit(code)
}
