package minigoscript_test

import (
	"fmt"
	"testing"

	"github.com/ivoras/minigoscript"
)

func TestParser(t *testing.T) {
	actions, err := minigoscript.DefaultParser.Parse(`
	let a = 1
	let b = true
	let c = 'hello'
	print c
	true
	`)

	if err != nil {
		t.Error(err)
		return
	}
	for _, a := range actions {
		fmt.Println(a.Action, a.Args)
	}
}

func TestParserLang(t *testing.T) {
	actions, err := minigoscript.DefaultParser.Parse(`
	let a = 1
	let b = true
	let c = 'Hello'
	print c
	print "World"
	print c "World"
	`)

	if err != nil {
		t.Error(err)
		return
	}
	symbolMap := map[string]interface{}{}

	for i, a := range actions {
		if a.Action == "let" {
			if len(a.Args) < 3 {
				t.Error("Not enough args in line", i)
				continue
			}
			if !a.Args[0].IsIdentifier() {
				t.Error("Expecting identifier in line", i)
				continue
			}
			if !a.Args[1].IsOperator() || a.Args[1].MustGetOperator() != "=" {
				t.Error("Expecting operator = in line", i)
				continue
			}
			symbolMap[a.Args[0].MustGetIdentifier()] = a.Args[2].Value()
		} else if a.Action == "print" {
			for _, a := range a.Args {
				if a.IsIdentifier() {
					fmt.Print(symbolMap[a.MustGetIdentifier()])
				} else {
					fmt.Print(a.Value())
				}
				fmt.Print(" ")
			}
			fmt.Println()
		}
	}
}
