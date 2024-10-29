package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkoukk/tiktoken-go"
)

func printHelp() {
    fmt.Println(`gotok: Token counting utility
Usage: 
  gotok [options] < [input]

Options:
  --model string        Use the embedding for this model (default: gpt-3.5-turbo)
  --encoding string     Use this specific encoding (overrides model if set)
  --input string        Path to input file
  --output string       Output destination: stderr, stdout, or a file path (default: stderr)
  --passthrough         Pass the input to stdout (default: false)
  --quiet               Quiet mode: output token count only, suppresses passthrough
  --list                List available models and encodings

Notes:
  - If both --input and stdin are provided, contents are concatenated.
  - You can use '<' for stdin input, e.g., gotok < input.txt.
`)
}

func main() {
    // Override the default usage message
    flag.Usage = printHelp

    // Define flags
    model := flag.String("model", "gpt-3.5-turbo", "Use the embedding for this model.")
    encoding := flag.String("encoding", "", "Use this specific encoding.")
    inputFile := flag.String("input", "", "Path to input file.")
    outputFile := flag.String("output", "stderr", "Output destination: stderr, stdout, or a file path.")
    passthrough := flag.Bool("passthrough", false, "If true, pass the input to stdout.")
    quiet := flag.Bool("quiet", false, "Quiet mode: output token count only.")
    list := flag.Bool("list", false, "List available models and encodings.")

    // Parse flags
    flag.Parse()

    // Display help if no arguments are provided and no input is provided on stdin
    fileInfo, _ := os.Stdin.Stat()
    if len(os.Args) == 1 && (fileInfo.Mode()&os.ModeCharDevice) != 0 {
        printHelp()
        os.Exit(0)
    }


    // Display available models and encodings if --list is provided
    if *list {
        fmt.Println("Available models and encodings:")
        for m, tke := range tiktoken.MODEL_TO_ENCODING {
            fmt.Printf("%-30v (%s)\n", m, tke)
        }
        return
    }

    // Read input from stdin and input file if provided
    var input []byte
    if *inputFile != "" {
        fileData, err := ioutil.ReadFile(*inputFile)
        if err != nil {
            fmt.Fprintln(os.Stderr, "Error reading input file:", err)
            os.Exit(1)
        }
        input = append(input, fileData...)
    }
    stdinData, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
        os.Exit(1)
    }
    input = append(input, stdinData...)

    // Initialize tokenizer based on model or encoding
    var enc *tiktoken.Tiktoken
    var encName string
    if *encoding != "" {
        enc, err = tiktoken.GetEncoding(*encoding)
        encName = *encoding
    } else {
        enc, err = tiktoken.EncodingForModel(*model)
        encName = tiktoken.MODEL_TO_ENCODING[*model]
    }
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error initializing tokenizer:", err)
        os.Exit(1)
    }

    // Tokenize input and count tokens
    tokens := enc.Encode(string(input), nil, nil)
    tokenCount := len(tokens)
    charCount := len(input)

    // Construct output message
    outputMsg := fmt.Sprintf("gotok: encoding=%s; chars=%d; tokens=%d\n", encName, charCount, tokenCount)

    // Handle --quiet mode
    if *quiet {
        fmt.Printf("%d\n", tokenCount)
        return
    }

    // Output result
    if *outputFile == "stderr" {
        fmt.Fprint(os.Stderr, outputMsg)
    } else if *outputFile == "stdout" {
        fmt.Print(outputMsg)
    } else {
        // Write to specified file
        err := ioutil.WriteFile(*outputFile, []byte(outputMsg), 0644)
        if err != nil {
            fmt.Fprintln(os.Stderr, "Error writing to file:", err)
            os.Exit(1)
        }
    }

    // Conditionally pass input to stdout based on --passthrough
    if *passthrough && !*quiet {
        fmt.Print(string(input))
    }
}
