# Petitions

A Go package that uses Markov chains to generate fake UK government petitions.
The first time the code is run, it will fetch all of the current UK petitions and generate a Markov chain with this data. The chain is saved to a file called chain.json to prevent it having to fetch the data again.

## Installation

`go get github.com/harrystevenson/petitions`

## Usage

```go
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
