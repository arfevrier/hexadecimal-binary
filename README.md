# Unbuffered hexadecimal converter

The only benefit over `xxd` is that the binary has no buffer. Input is immediately transcoded to its hexadecimal representation. And vice versa.

## Build OS
Golang build:
- Linux amd64
- Windows amd64
## Usage

Encode binary input from *stdin* to *stdout* string representation of hexadecimal: 

```bash
$ echo -n "Hello world!" | ./hexadecimal-binary -mode encode
48656c6c6f20776f726c6421
```

Decode string representation of hexadecimal from *stdin* to binary output on *stdout* : 
```bash
$ echo -n "48656c6c6f20776f726c6421" | ./hexadecimal-binary -mode decode
Hello world!
```
