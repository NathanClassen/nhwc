# nhwc - nathanhayesWordCount

### A full-featured clone of the Unix tool `wc` written in Go!
Written to solve the first challenge over at [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-wc)
___

## Build

`make build`

This will create a binary at `bin/nhwc`

## Usage

Invoke `nhwc` passing in desired options followed by 0 or more files. `nhwc` will output how many lines, words, bytes, or character each file contains, based on options specified.

If no files are provided, `nhwc` will read from stdin until EOF (or ^D) and take this as input to the program.



#### nhwc [-c -l -m -w] [file ...]


#### `-c` will write the number of bytes in each input file to the stdout – this option will cancel out any prior usage of the `-m` option.

#### `-l` will write the number of lines in each input file to the stdout.

#### `-m` will write the number of characters in each input file to stdout.  If the current locale does not support multibyte characters, this is equivalent to the -c option.  This will cancel out any prior usage of the -c option.

#### `-w` will write the number of words in each input file to the stdout.
