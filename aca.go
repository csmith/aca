package aca

import (
	"embed"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//go:embed data/*.txt
var data embed.FS

// Generator provides methods to create adjective-colour-animal triplets.
type Generator struct {
	adjectives []string
	colours    []string
	animals    []string

	rand      *rand.Rand
	delimiter string
}

// NewDefaultGenerator creates a new Generator using "-" as a delimiter and a default source of randomness.
func NewDefaultGenerator() (*Generator, error) {
	return NewGenerator("-", rand.NewSource(time.Now().UnixNano()))
}

// NewGenerator creates a new Generator that will use the given delimiter and source of randomness.
func NewGenerator(delimiter string, source rand.Source) (*Generator, error) {
	adjectives, err := readData("data/adjectives.txt")
	if err != nil {
		return nil, err
	}

	colours, err := readData("data/colours.txt")
	if err != nil {
		return nil, err
	}

	animals, err := readData("data/animals.txt")
	if err != nil {
		return nil, err
	}

	return &Generator{
		adjectives: adjectives,
		colours:    colours,
		animals:    animals,
		rand:       rand.New(source),
		delimiter:  delimiter,
	}, nil
}

// Generate creates a new {adjective, colour, animal} tuple and formats it using the configured delimiter.
func (g *Generator) Generate() string {
	adjective := g.adjectives[g.rand.Intn(len(g.adjectives))]
	colour := g.colours[g.rand.Intn(len(g.colours))]
	animal := g.animals[g.rand.Intn(len(g.animals))]
	return fmt.Sprintf("%s%s%s%s%s", adjective, g.delimiter, colour, g.delimiter, animal)
}

func readData(name string) ([]string, error) {
	b, err := data.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(b), "\n"), nil
}
