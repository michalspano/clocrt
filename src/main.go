package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	/*
		--------------------------------------------------------------------------------
		Usage with explanation:
		1. --print | -pr; we print the redefined table to the standard output
		2. --cell-align=<align> | -ca=<align>; we align the cells in the table
		  according to the given alignment: left, center, right
		3. --output-path=<path> | -op=<path>; we save the redefined table to the
		  an `.md` file at the given path
		4. --help | -h; we print the usage and exit the program
		  the help flag is automatically executed if no command-line arguments are given
		--------------------------------------------------------------------------------
	*/

	if len(os.Args) == 1 || os.Args[1] == "--help" || os.Args[1] == "-h" {
		func() {
			fmt.Println("Usage: clocrt \"`[CLOC-PATTERN]`\" [OPTIONS]")
			fmt.Println("OPTIONS:\n\t--cell-align|-ca=<align>\t center|c (default), left|l, right|r")
			fmt.Println("\t--output-path|-op=<path>\t output file path (default: out.md)")
			fmt.Println("\t--print|-pr\t\t\t print the result to stdout")
			fmt.Println("BUGS:\n\t Don't forget to quote the pattern per the example: \"`[CLOC-PATTERN]`\"")
		}()
		os.Exit(1)
	}

	// create buffers for optional flags and parse them from the command-line per the documentation
	var (
		args                        []string = os.Args[1:]
		printOutput, startIter      bool     = false, false
		cellAlignOption, outputPath string   = "", "out.md"
	)

	func() {
		for _, arg := range args {
			if strings.Contains(arg, "--print") || strings.Contains(arg, "-pr") {
				printOutput = true
			} else if strings.Contains(arg, "--cell-align") || strings.Contains(arg, "-ca") {
				cellAlignOption = strings.Split(arg, "=")[1]
			} else if strings.Contains(arg, "--output-path") || strings.Contains(arg, "-op") {
				outputPath = strings.Split(arg, "=")[1]
			}
		}
	}()

	/*
		--------------------------------------------------------------------------------
		Reading the CLOC-PATTERN from the command-line:
		1. We split the pattern by the newline character;
		2. We append the lines, just then the string 'github' is found in the line, then
		 we chcek if the line is a separator or if empty - we won't append those;
		 Note: this is just a matter of convention, we can change it to any string that is
		 found in any of the lines of the CLOC-PATTERN.
		 - TODO: add more options for the trimming of the CLOC-PATTERN
		3. We split each line by the whitespace character, albeit we need to remove indexes
		 that are empty;
		4. We populate the 2D array with the parsed CLOC-PATTERN;
		--------------------------------------------------------------------------------
	*/

	parsedData := [][]string{}
	lines := strings.Split(args[0], "\n")

	func() {
		for _, line := range lines {
			if strings.Contains(line, "github") {
				startIter = true
			}
			if startIter && line != "" && line[0] != '-' {
				lineTrim := strings.Split(line, " ")

				rowBuffer := []string{}
				for _, val := range lineTrim {
					if val != "" {
						rowBuffer = append(rowBuffer, val)
					}
				}
				parsedData = append(parsedData, rowBuffer)
			}
		}
	}()

	/*
		--------------------------------------------------------------------------------
		Redefining the CLOC-PATTERN:
		1. We create a new 2D array that will hold the redefined CLOC-PATTERN;
		2. We iterate over the CLOC-PATTERN and for each line we check for non-alphanumeric
		 characters, we append them to a buffer.
		3. Then, we populate the row with the remaining corresponding numerical values.
		Note: the reason for the followin - some names from within the CLOC-PATTERN may be
		 made from more words, therefore being interpreted as separate data fields. We carry
		 this process to avoid the aforementioned issue.
		Note2: the first 2 lines needn't be checked.
		TODO: add additional spacing to a column with any name of length > 1
		--------------------------------------------------------------------------------
	*/

	tempParsedData := [][]string{}
	func() {
		for n, row := range parsedData {
			if n <= 1 {
				tempParsedData = append(tempParsedData, row)
				continue
			}
			idx := 0
			var nonAlphaNumeric, temp []string
			for _, str := range row {
				if !strings.ContainsAny(str, "0123456789") {
					nonAlphaNumeric = append(nonAlphaNumeric, str)
					idx++
				}
			}
			// append the remaining numerical values to the row
			// (with the respect to the redefined name)
			temp = append(temp, strings.Join(nonAlphaNumeric, " "))
			for i := idx; i < len(row); i++ {
				temp = append(temp, row[i])
			}
			tempParsedData = append(tempParsedData, temp)
		}
	}()
	parsedData = tempParsedData // update the parsedData with the new data

	// ensure that data were found - correct CLOC-PATTERN given
	if len(parsedData) == 0 {
		fmt.Println("No data found")
		os.Exit(1)
	}

	// output file
	file, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writeCurrentRow := func(row, header []string) {
		for idx, val := range row {
			file.WriteString(fmt.Sprintf("| %s ", val))

			// the additional spacing is only required for the body rows
			if header == nil {
				continue
			}

			// we add the additional spacing with the respect to the length of
			// the corresponding cell of the header row

			headerCellLen := len(header[idx])
			for i := 0; i < headerCellLen-len(val); i++ {
				file.WriteString(" ")
			}
		}
		file.WriteString("|\n")
	}

	// write the heading (not the header row) to the table
	heading := strings.Join(parsedData[0], " ")
	file.WriteString("### " + heading + "\n")

	// write the header row of the table
	headerRow := parsedData[1]
	writeCurrentRow(headerRow, nil)

	/*
		--------------------------------------------------------------------------------
		Markdown [MD] cell alignment syntax
		Docs: https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet#tables
		Note: We add extra whitespaces within the cells of each row for better readability.
		--------------------------------------------------------------------------------
	*/

	tableSeparator := func(fillSize, n int) {
		for i := 0; i < fillSize-n; i++ {
			file.WriteString("-")
		}
	}
	func() {
		for i := 0; i < len(headerRow); i++ {
			headerCellSize := len(headerRow[i])

			switch cellAlignOption {
			case "r", "right":
				file.WriteString("| ")
				tableSeparator(headerCellSize, 1)
				file.WriteString(": ")

			case "l", "left":
				file.WriteString("| :")
				tableSeparator(headerCellSize, 1)
				file.WriteString(" ")

			default: // center
				file.WriteString("| :")
				tableSeparator(headerCellSize, 2)
				file.WriteString(": ")
			}
		}
		file.WriteString("|\n")
	}()

	// write the body of the table
	for _, item := range parsedData[2:] {
		writeCurrentRow(item, headerRow)
	}

	/*
		--------------------------------------------------------------------------------
		Print to the standard output if the --print | -p flag is set for brevity, we use
		the `cat` command to read the output file we won't interate once again over the
		content of the file in the execution of the program; the `cat` command is a simple
		and more efficient way to read the content of the file than reading it line by
		line once again.
		--------------------------------------------------------------------------------
	*/

	if printOutput {
		func() {
			cmd := exec.Command("cat", file.Name())
			cmd.Stdout = os.Stdout
			cmd.Run()
		}()
	}
}
