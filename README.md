# Adjective-colour-animal generator

Very simple library for generating random human-friendly identifiers. These are useful when the user may need to
share the code with others, such as when inviting them to play a game.

The word lists are heavily based on those used by the [unique-names-generator](https://github.com/andreasonny83/unique-names-generator)
npm module, with some manual changes to exclude words that are hard to spell, hard to pronounce, or are biased or
contentious in some way. To make up for the removals, a number of colours and fictional animals have been added.

Example identifiers:

```
beautiful-cerulean-timelord
inevitable-grey-toad
satisfied-yellow-tapir
weak-crimson-aardwolf
coherent-blush-guineafowl
poised-blue-vicuna
female-turquoise-dwarf
```

There are 1172 adjectives, 132 colours and 419 animals giving rise to just over 64 million unique combinations of words.

The longest adjective is 15 characters, colour is 11 characters, and animal is 13, giving rise to a maximum string
length of 41 characters if using a single character delimiter like "-".

## Example usage

```go
package main

import (
	"github.com/csmith/aca"
	"math/rand"
	"time"
)

func main() {
	// Using the default generator
	generator, err := aca.NewDefaultGenerator()
	if err != nil {
		panic(err)
	}
	println(generator.Generate()) // prints something like "beautiful-cerulean-timelord"

	// Using a custom delimiter
	generator, err = aca.NewGenerator(", ", rand.NewSource(time.Now().UnixNano()))
	println(generator.Generate()) // prints something like "beautiful, cerulean, timelord"
}
```