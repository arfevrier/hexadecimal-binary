package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

func main() {
	mode := flag.String("mode", "encode", "Mode: 'encode' for hex encoding, 'decode' for binary decoding")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	buffer := make([]byte, 4096)

	switch *mode {
	case "encode":
		for {
			n, err := reader.Read(buffer)
			if n > 0 {
				fmt.Fprint(writer, hex.EncodeToString(buffer[:n]))
				writer.Flush()
			}
			if err != nil {
				break
			}
		}
	case "decode":
		for {
			n, err := reader.Read(buffer)
			if n > 0 {
				decoded, decodeErr := hex.DecodeString(string(buffer[:n]))
				if decodeErr != nil {
					fmt.Fprintln(os.Stderr, "Error decoding hex:", decodeErr)
					break
				}
				writer.Write(decoded)
				writer.Flush()
			}
			if err != nil {
				break
			}
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid mode. Use 'encode' or 'decode'.")
	}
}
