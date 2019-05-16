# Petitions

A go package that uses markov chains to generate fake UK government petitions.

## Installation

`go get github.com/harrystevenson/petitions`

## Usage

```
    import (
        "fmt"
        "github.com/harrystevenson/petitions"
    )

    func main() {
	    chain := petitions.CreateChain()
	    for i := 0; i < 10; i++ {
		    fmt.Println(petitions.GeneratePetition(chain))
	    }
    }

```
