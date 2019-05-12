package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Replacer type holds settings to rename and replace
type Replacer struct {
	toReplace   string
	replaceWith string
	matchCase   bool
}

// StringReplacer returns Replacer object with rename and replace values
func StringReplacer(toReplace string, replaceWith string, matchCase bool) *Replacer {
	return &Replacer{
		toReplace:   toReplace,
		replaceWith: replaceWith,
		matchCase:   matchCase,
	}
}

// Replace replaces string based on match condition
func (cir *Replacer) Replace(str string) string {
	if cir.matchCase {
		return strings.ReplaceAll(str, cir.toReplace, cir.replaceWith)
	} else {
		regExp := regexp.MustCompile("(?i)" + cir.toReplace)
		return regExp.ReplaceAllString(str, cir.replaceWith)
	}
}

// ShowCopyright shows product and copyright information
func ShowCopyright() {
	fmt.Println("randr - Rename And Replace [Version 1.0]")
	fmt.Println("https://github.com/firozansari/randr")
	fmt.Println("")
	fmt.Println("(c) 2019 Firoz Ansari. All rights reserved.")
	fmt.Println("")
}

// ShowUsage shows command line usage with examples
func ShowUsage() {
	fmt.Println("Usage:")
	fmt.Println("  rand -find=source -replace=target [-match=true/false] -location=directory")
	fmt.Println("")
	fmt.Println("Arguments:")
	fmt.Println("  -f, -find       find text to rename or replace")
	fmt.Println("  -r, -replace    replace text to")
	fmt.Println("  -m, -match      match case -- true: case sensitive (default); false: case insensitive")
	fmt.Println("  -l, -location   directory location to process files and sub directories")
	fmt.Println("")
	fmt.Println("  -h, -?, -help   show command usage and examples")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  randr.exe -find=worker -replace=employee -match=true -location=C:\\projects\\payroll\\")
	fmt.Println("  randr.exe -find=Notes -replace=contents -match=false -location=C:\\documents\\daily\\")
	fmt.Println("")
}

func main() {
	ShowCopyright()

	flag.Usage = func() {
		ShowUsage()
	}

	var findPtr string
	flag.StringVar(&findPtr, "find", "", "a string")
	flag.StringVar(&findPtr, "f", "", "a string")

	var replacePtr string
	flag.StringVar(&replacePtr, "replace", "", "a string")
	flag.StringVar(&replacePtr, "r", "", "a string")

	var matchPtr bool
	flag.BoolVar(&matchPtr, "match", true, "a bool")
	flag.BoolVar(&matchPtr, "m", true, "a bool")

	var locationPtr string
	flag.StringVar(&locationPtr, "location", ".", "a string")
	flag.StringVar(&locationPtr, "l", ".", "a string")

	flag.Parse()

	fmt.Println("Processing ...")
	fmt.Println("")

	var directories []string
	errDirectoryProcess := filepath.Walk(locationPtr, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			directories = append(directories, path)
		}
		return nil
	})
	if errDirectoryProcess != nil {
		panic(errDirectoryProcess)
	}

	for _, directory := range directories {
		// Rename the directory if its name contains the search text
		replacerDirectoryName := StringReplacer(findPtr, replacePtr, matchPtr)
		newName := replacerDirectoryName.Replace(directory)
		if directory != newName {
			fmt.Println("RENAME: " + directory + " -> " + newName)
			err := os.Rename(directory, newName)
			if err != nil {
				panic(err)
			}
		}
	}

	var files []string
	errFileProcess := filepath.Walk(locationPtr, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if errFileProcess != nil {
		panic(errFileProcess)
	}

	for _, file := range files {

		// Find and replace content of the file if its contains the search text
		contentPtr, errContentProcess := ioutil.ReadFile(file)
		if errContentProcess != nil {
			panic(errContentProcess)
		}
		content := string(contentPtr)
		replacerContent := StringReplacer(findPtr, replacePtr, matchPtr)
		updateContent := replacerContent.Replace(content)
		if content != updateContent {
			fmt.Println("REPLACE: " + file)
			errReplaceProcess := ioutil.WriteFile(file, []byte(updateContent), 0644)
			if errReplaceProcess != nil {
				panic(errReplaceProcess)
			}
		}

		// Rename the file if its name contains the search text
		replacerFileName := StringReplacer(findPtr, replacePtr, matchPtr)
		newName := replacerFileName.Replace(file)
		if file != newName {
			fmt.Println("RENAME: " + file + " -> " + newName)
			err := os.Rename(file, newName)
			if err != nil {
				panic(err)
			}
		}

	}

	fmt.Println("")
	fmt.Println("Completed.")

}
