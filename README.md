# gotok - Token Counting Utility

gotok is a lightweight command-line utility designed to count tokens in text using OpenAI's tokenization standards, making it easy to estimate input sizes for models like GPT-4o and older.

## Features
- Supports multiple OpenAI models and encodings.
- Flexible input and output options.
- Quiet mode for streamlined token counting.

## Installation

### Prerequisites
- Go 1.19 or higher

### Install
Clone the repository and build the binary using Go:

```sh
git clone https://github.com/mattjoyce/gotok.git
cd gotok
go install
```

This will compile the `gotok` binary and place it in your `$GOPATH/bin` directory, making it accessible from your command line.

## Usage

```sh
gotok [options] < [input]
```

### Options

- `--model string`  
  Specifies the model to use for token embedding (default: `gpt-4o`). If set, the model's default encoding will be applied.

- `--encoding string`  
  Sets a specific encoding manually, overriding the model's default. Must be one of the listed encodings (e.g., `cl100k_base`).

- `--input string`  
  File path to the input text file. If both `--input` and stdin are provided, `gotok` will concatenate their contents.

- `--output string`  
  Designates where to send the output. Choose between `stderr`, `stdout`, or specify a file path (default: `stderr`).

- `--passthrough`  
  Outputs the original input text to `stdout` in addition to the token count. Set to `false` by default.

- `--quiet`  
  Quiet mode suppresses all output except for the token count, overriding `--passthrough`.

- `--list`  
  Lists the available models and encodings, providing an easy reference for compatible options.

## Available Models and Encodings
Here is a list of available models and their associated encodings that can be used with gotok:

- `code-search-ada-code-001`       (r50k_base)
- `code-davinci-001`               (p50k_base)
- `text-embedding-ada-002`         (cl100k_base)
- `text-similarity-curie-001`      (r50k_base)
- `code-search-babbage-code-001`   (r50k_base)
- `text-ada-001`                   (r50k_base)
- `code-davinci-002`               (p50k_base)
- `text-embedding-3-small`         (cl100k_base)
- `text-davinci-002`               (p50k_base)
- `ada`                            (r50k_base)
- `cushman-codex`                  (p50k_base)
- `code-davinci-edit-001`          (p50k_edit)
- `text-search-davinci-doc-001`    (r50k_base)
- `text-babbage-001`               (r50k_base)
- `code-cushman-002`               (p50k_base)
- `code-cushman-001`               (p50k_base)
- `text-similarity-davinci-001`    (r50k_base)
- `text-similarity-babbage-001`    (r50k_base)
- `text-similarity-ada-001`        (r50k_base)
- `text-search-ada-doc-001`        (r50k_base)
- `gpt2`                           (gpt2)
- `davinci`                        (r50k_base)
- `babbage`                        (r50k_base)
- `text-search-curie-doc-001`      (r50k_base)
- `curie`                          (r50k_base)
- `davinci-codex`                  (p50k_base)
- `text-davinci-edit-001`          (p50k_edit)
- `text-embedding-3-large`         (cl100k_base)
- `gpt-4`                          (cl100k_base)
- `gpt-3.5-turbo`                  (cl100k_base)
- `text-curie-001`                 (r50k_base)
- `text-search-babbage-doc-001`    (r50k_base)
- `gpt-4o`                         (o200k_base)
- `text-davinci-003`               (p50k_base)
- `text-davinci-001`               (r50k_base)

## Examples

1. **Basic Usage**  
   Count tokens in `input.txt`, displaying output in `stderr` (default):
   ```sh
   gotok --input input.txt
   ```

2. **Using Stdin Input**  
   Pipe input text from stdin:
   ```sh
   gotok < input.txt
   ```

3. **Specify Encoding and Model**  
   Specify a particular encoding and model:
   ```sh
   gotok --encoding "cl100k_base" --model "gpt-4" < input.txt
   ```

4. **Quiet Mode**  
   Count tokens only (suppress all other output):
   ```sh
   gotok --quiet --input input.txt
   ```

5. **List Available Models and Encodings**  
   ```sh
   gotok --list
   ```

## Notes
- When both `--input` and stdin input are provided, the content from both sources will be concatenated before processing.
- Use the `--passthrough` option to print the original input text to `stdout` while viewing token counts.

This utility simplifies text preprocessing tasks for token-based models by giving you quick insight into input size, which helps in better managing prompt limits and structuring inputs for optimal model performance.

