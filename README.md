# Art Encoder and Decoder

This is a simple Go program that can encode and decode strings into a specific art format.

## Getting Started

To run the program, use the following command:

``` go run .```

## Usage

The program supports several command-line flags:

-h: Show help information.
-d: Decode an art string.
-e: Encode a string into art format.
-multi: Enable multi-line mode.

## Decoding

To decode an art string, use the -d flag followed by the string in quotes. For example:

``` go run . -d "[1 #][2 -_]-[3 #]" ```

## Encoding

To encode a string into art format, use the -e flag followed by the string. For example:

``` go run . -e AAABBC ```

## Multi-line Mode

To enable multi-line mode, add the -multi flag. After running the command, you can type in your input. Once you have finished, press 'Ctrl+d'. For example:

``` go run . -d -multi ```
``` [1 _]/[1 ?][1 _] ```
``` [1 _]/[1 ?][1 _] ```

## Error Handling

If there is an error during encoding or decoding, the program will print an error message and exit.