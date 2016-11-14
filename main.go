package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/yuya-takeyama/argf"
)

const usage = `Usage: url64 [-d] <file>

URL-safe Base64 encode or decode <file> or standard input, to standard output.
`

func decode(reader io.Reader) error {
	_, err := io.Copy(os.Stdout, base64.NewDecoder(base64.RawURLEncoding, reader))
	return err
}

func encode(reader io.Reader) error {
	encoder := base64.NewEncoder(base64.RawURLEncoding, os.Stdout)
	_, err := io.Copy(encoder, os.Stdin)
	if err != nil {
		return err
	}
	err = encoder.Close()
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write([]byte("\n"))
	if err != nil {
		return err
	}
	return nil
}

func run() error {
	var displayHelp bool
	var shouldDecode bool
	flag.BoolVar(&displayHelp, "h", false, "display help")
	flag.BoolVar(&shouldDecode, "d", false, "decode")
	flag.Parse()
	reader, err := argf.From(flag.Args())

	if err != nil {
		return err
	}

	if displayHelp {
		fmt.Println()

		return nil
	}

	if shouldDecode {
		return decode(reader)
	}

	return encode(reader)
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
