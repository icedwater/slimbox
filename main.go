package main

import (
	"github.com/voxelbrain/goptions"
	"github.com/yarbelk/slimbox/lib"
	"github.com/yarbelk/slimbox/lib/cat"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

var catOptions cat.CatOptions

var numberLines, numberNonBlankLines, help bool
var lineNumber int = 1

type Options struct {
	Verb goptions.Verbs
	Cat  cat.CatOptions `goptions:"cat"`
}

func runCat(options *Options) {
	catOptions = *cat.NewCatOptions()
	catOptions.Blank = options.Cat.Blank
	catOptions.Number = options.Cat.Number
	catOptions.EoL = options.Cat.EoL
	catOptions.Tabs = options.Cat.Tabs

	if catOptions.Blank {
		catOptions.Number = true
	}
	if len(options.Cat.Files) == 0 {
		infile := os.NewFile(uintptr(syscall.Stdin), "/dev/stdin")
		err := catOptions.Cat(infile, os.Stdout)
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, file := range options.Cat.Files {
		func() {
			_, fi, err := lib.ParseFiles(file)
			if err != nil {
				log.Fatal(err)
			}
			defer fi.Close()
			err = catOptions.Cat(fi, os.Stdout)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

}

func main() {
	options := Options{}

	flagset := goptions.NewFlagSet(filepath.Base(os.Args[0]), &options)
	err := flagset.Parse(os.Args[1:])

	if err != nil {
		if options.Verb == "cat" {
			var catHelp goptions.HelpFunc = goptions.NewTemplatedHelpFunc(cat.CAT_HELP)
			flagset.Verbs["cat"].HelpFunc = catHelp
			flagset.Verbs["cat"].PrintHelp(os.Stderr)
			os.Exit(1)
		}
		flagset.PrintHelp(os.Stderr)
		os.Exit(1)
	}

	if options.Verb == "cat" {
		runCat(&options)
	} else {
		flagset.PrintHelp(os.Stderr)
	}
}
