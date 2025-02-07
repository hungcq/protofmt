package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Define the override flag.
	override := flag.Bool("o", false, "modify file in place")
	// Also allow the long version.
	flag.BoolVar(override, "override", false, "modify file in place")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [-o|--override] <proto_file>\n", os.Args[0])
		os.Exit(1)
	}
	filename := flag.Arg(0)

	// Open the file.
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", filename, err)
		os.Exit(1)
	}
	defer file.Close()

	// Read the file into lines.
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}

	// Prepare a buffer to collect the formatted output.
	var output bytes.Buffer

	// Regex to match a protobuf field line.
	fieldRegex := regexp.MustCompile(`^(\s*)([^\s].*?\S)\s+([A-Za-z_]\w*)\s*=\s*([0-9]+)(.*)$`)

	// fieldLine holds the parts of a protobuf field declaration.
	type fieldLine struct {
		indent string // leading whitespace
		ftype  string // field type or qualifiers
		name   string // field name
		number string // field number
		rest   string // remainder of the line (semicolon, comment, etc.)
	}

	var group []fieldLine
	flushGroup := func() {
		if len(group) == 0 {
			return
		}
		// Find maximum widths for the type and name columns.
		maxType, maxName := 0, 0
		for _, fl := range group {
			if len(fl.ftype) > maxType {
				maxType = len(fl.ftype)
			}
			if len(fl.name) > maxName {
				maxName = len(fl.name)
			}
		}
		// Write each field with proper padding.
		for _, fl := range group {
			pad1 := strings.Repeat(" ", maxType-len(fl.ftype))
			pad2 := strings.Repeat(" ", maxName-len(fl.name))
			fmt.Fprintf(&output, "%s%s%s %s%s = %s%s\n", fl.indent, fl.ftype, pad1, fl.name, pad2, fl.number, fl.rest)
		}
		group = nil
	}

	// Process each line.
	for _, line := range lines {
		if matches := fieldRegex.FindStringSubmatch(line); matches != nil {
			fl := fieldLine{
				indent: matches[1],
				ftype:  matches[2],
				name:   matches[3],
				number: matches[4],
				rest:   matches[5],
			}
			group = append(group, fl)
		} else {
			flushGroup()
			output.WriteString(line + "\n")
		}
	}
	flushGroup()

	// Write the output.
	if *override {
		// Write back to the file (modify in place).
		if err := os.WriteFile(filename, output.Bytes(), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file %s: %v\n", filename, err)
			os.Exit(1)
		}
	} else {
		// Write to stdout.
		io.Copy(os.Stdout, &output)
	}
}
