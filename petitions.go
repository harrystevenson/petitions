package petitions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/mb-14/gomarkov"
)

// CreateChain checks to see if we have a saved version of our chain, and creates it if not
func CreateChain() *gomarkov.Chain {
	// The file holding the chain
	_, err := os.Stat("chain.json")

	// If we don't have the saved chain, create it
	if os.IsNotExist(err) {
		// Get the petitions
		petitions := fetchPetitionData()
		// Add them to our chain
		chain := gomarkov.NewChain(1)
		for i := 0; i < len(petitions); i++ {
			chain.Add(strings.Split(petitions[i], " "))
		}
		// Save the chain to a file
		chainData, _ := json.Marshal(chain)
		ioutil.WriteFile("chain.json", chainData, 0644)
		return chain
	}

	// Read our chain from file and return it
	var chain gomarkov.Chain
	data, _ := ioutil.ReadFile("chain.json")
	json.Unmarshal(data, &chain)
	return &chain

}

// Requests the data from the Petitions API. Returns a list of all petition names.
func fetchPetitionData() []string {
	petitionPageCount := 412
	petitions := []string{}

	rootURL := "https://petition.parliament.uk/petitions.json?state=all&page="
	for i := 1; i < petitionPageCount+1; i++ {
		fmt.Println("Fetching Page", i)
		response, err := http.Get(rootURL + strconv.Itoa(i))
		if err != nil {
			panic(err)
		}

		defer response.Body.Close()

		// Holds our response body (quite bodgy but saves making lots of structs)
		var data map[string][]map[string]map[string]string

		body, _ := ioutil.ReadAll(response.Body)

		_ = json.Unmarshal(body, &data)

		for j := 0; j < len(data["data"]); j++ {
			title := data["data"][j]["attributes"]["action"]
			petitions = append(petitions, title)
		}

	}

	return petitions

}

// GeneratePetition creates a petition name using the markov chain.
func GeneratePetition(chain *gomarkov.Chain) string {
	tokens := []string{gomarkov.StartToken}
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, _ := chain.Generate(tokens[(len(tokens) - 1):])
		tokens = append(tokens, next)
	}
	return strings.Join(tokens[1:len(tokens)-1], " ")
}
