package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/anchepiece/suggest"
)

const (
	Name     = "suggest-autocorrect"
	FullName = "github.com/anchepiece/suggest-autocorrect"
	Usage    = "suggest-autocorrect is a console based tool to return the closest match in a list"
	Version  = "0.1.0"
	Author   = "anchepiece"
	Email    = "anchepiece@gmail.com"
)

var (
	flagQuery               = flag.String("q", "", "query term to search")
	flagCommands            = flag.String("c", "", "comma-separated list of possible suggestions")
	flagAutocorrectDisabled = flag.Bool("d", false, "disable autocorrect feature")
)

func simpleExample() {

	suggester := suggest.Suggest{}

	query := "fgerp"
	commands := []string{"cat", "mkdir", "fgrep", "history"}

	match, err := suggester.AutocorrectAgainst(query, commands)
	if err == nil {
		fmt.Println("Autocorrected to:", match) // "fgrep"
	}

	suggester.Commands = commands
	if match, err := suggester.Autocorrect(query); err == nil {
		fmt.Println("Also Autocorrected to:", match) // "fgrep"
	}

	query = "println"
	commands = []string{"Fprint", "Fprintf", "Fprintln", "Sprintf", "Print", "Printf", "Println"}
	suggester.Options.SimilarityMinimum = 8

	fmt.Printf("Searching %v for %s\n", query, commands)

	if result, err := suggester.QueryAgainst(query, commands); err == nil {
		if !result.Success() {
			fmt.Println("No close matches")
		} else {
			fmt.Println("Similar matches:", result.Matches) // [Println Fprintln]
			fmt.Println("Autocorrect:", result.Autocorrect) // Println

		}
	}
	os.Exit(0)
}

func simpleExample2() {
	suggester := suggest.Suggest{}

	query := "fgerp"
	commands := []string{"cat", "mkdir", "fgrep", "history"}

	suggester.Commands = commands
	if match, err := suggester.Autocorrect("mkdri"); err == nil {
		fmt.Println("Autocorrected to:", match) // "mkdir"
	}

	// Alternate autocorrect usage pattern
	match, _ := suggester.AutocorrectAgainst(query, commands)
	if match != "" {
		fmt.Println("Autocorrected to:", match) // "fgrep"
	}

	// Alternate usage pattern
	query = "printf"
	commands = []string{"Fprint", "Fprintf", "Fprintln", "Sprintf", "Printf", "Println"}
	suggester.Options.SimilarityMinimum = 8

	fmt.Printf("Searching %v in %s\n", query, commands)

	if result, err := suggester.QueryAgainst(query, commands); err == nil {
		if !result.Success() {
			fmt.Println("No close matches")

		} else {
			fmt.Println("Similar matches:", result.Matches)
			// [Println Fprintln]

			fmt.Println("Autocorrect:", result.Autocorrect)
			// Println
		}
	}
}

func simpleExample3() {
	// make sure our example in doc.go works
	suggester := suggest.Suggest{}
	suggester.Options.SimilarityMinimum = 7
	suggester.Options.AutocorrectDisabled = false

	query := "proflie"
	commands := []string{"perfil", "profiel", "profile", "profil", "account"}

	suggester.Commands = commands
	if result, err := suggester.Query(query); err == nil {
		if !result.Success() {
			fmt.Println("No close matches")
		} else {
			fmt.Println("Similar matches:", result.Matches) // [profile profil profiel]
			fmt.Println("Autocorrect:", result.Autocorrect) // profile
		}
	}
	os.Exit(0)
}

func main() {

	flag.Parse()
	// fmt.Println(Name)

	simpleExample3()
	simpleExample2()
	simpleExample()

	if *flagQuery == "" || *flagCommands == "" {
		fmt.Println("Error: must specify both a query (-q) and the possible commands (-c)")
		os.Exit(1)
	}

	fmt.Println("Entered query:", *flagQuery)

	var commandList []string
	for _, s := range strings.Split(*flagCommands, ",") {
		commandList = append(commandList, strings.TrimSpace(s))
	}
	if len(commandList) > 0 {
		fmt.Println("Entered possible commands:", commandList)
	}

	fmt.Println()
	s := &suggest.Suggest{
		Commands: commandList,
	}
	s.Options.AutocorrectDisabled = *flagAutocorrectDisabled
	s.Options.SimilarityMinimum = 7

	if m := s.ExactMatch(*flagQuery); m != "" {
		fmt.Println("Continuing with exact match!", m)
		return // Run the command!
	} else {
		fmt.Printf("You called a command '%s' which does not exist. \n", *flagQuery)
	}

	if !*flagAutocorrectDisabled {
		if a, err := s.Autocorrect(*flagQuery); err == nil && a != "" {
			fmt.Printf("Continuing under the assumption that you meant '%s'\n", a)
			return // Run the command!
		}
	}

	m, err := s.Query(*flagQuery)
	if err != nil {
		fmt.Println("An error occurred: %v", err)
	}

	if len(m.Matches) == 1 {
		fmt.Println("Did you mean this?")
	}

	if len(m.Matches) > 1 {
		fmt.Println("Did you mean one of these?")
	}

	for _, suggestion := range m.Matches {
		fmt.Printf("\t%s\n", suggestion)
	}
}

// scan through a passed file and append all words to available commands
func scan(r io.Reader) {
	// TODO: handle command line option to parse

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		// Match the query to input
		fmt.Println(word)
	}
}
