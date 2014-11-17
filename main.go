package main

import (
	"github.com/voxelbrain/goptions"
	"github.com/yarbelk/slimbox/lib"
	"log"
	"os"
	"path/filepath"
)

var catOptions cat.CatOptions

var numberLines, numberNonBlankLines, help bool
var lineNumber int = 1

func main() {
	catOptions = *cat.NewCatOptions()

	options := struct {
		Verb goptions.Verbs
		Cat  struct {
			EoL    bool          `goptions:"-E, --show-ends, description='display $ at end of each line'"`
			Tabs   bool          `goptions:"-T, --show-tabs, description='display TAB characters as ^I'"`
			Number bool          `goptions:"-n, --number, description='number all output lines'"`
			Blank  bool          `goptions:"-b, --number-nonblank, description='number non-blank output lines, overrides -n'"`
			Help   goptions.Help `goptions:"-h, --help, description='print this message'"`
			Files  goptions.Remainder
		} `goptions:"cat"`
	}{}

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
		catOptions.Blank = options.Cat.Blank
		catOptions.Number = options.Cat.Number
		catOptions.EoL = options.Cat.EoL
		catOptions.Tabs = options.Cat.Tabs

		if catOptions.Blank {
			catOptions.Number = true
		}
		if len(options.Cat.Files) == 0 {
			err := catOptions.Cat(os.Stdin, os.Stdout)
			if err != nil {
				log.Fatal(err)
			}
		}
		for _, file := range options.Cat.Files {
			func() {
				fi, err := os.Open(file)
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
	} else {
		flagset.PrintHelp(os.Stderr)
	}
}
