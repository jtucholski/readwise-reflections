package core

import (
	"encoding/json"
	"io"

	"math/rand"
	"net/http"
	"net/url"
)

type Highlight struct {
	Text string `json:"text"`
	Id   int    `json:"id"`
	Tags []Tag  `json:"tags"`
}

type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Page struct {
	Results []Highlight `json:"results"`
	Next    string      `json:"next"`
}

// GetFavoriteQuote fetches a random favorite quote from the Readwise API.
//
// Parameters:
// - readwiseToken: The Readwise API token used for authentication.
//
// Returns:
// - string: The text of the randomly selected favorite quote.
// - error: An error if the request fails or the response cannot be parsed.
func GetFavoriteQuote(readwiseToken string) (string, error) {
	highlights, favorites := []Highlight{}, []Highlight{}

	url, _ := url.Parse("https://readwise.io/api/v2/highlights?page_size=500")
	request, _ := http.NewRequest("GET", url.String(), nil)
	request.Header.Add("Authorization", "Token "+readwiseToken)

	for {
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			return "", err
		}

		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return "", err
		}

		var page Page
		err = json.Unmarshal(body, &page)
		if err != nil {
			return "", err
		}

		highlights = append(highlights, page.Results...)

		if page.Next == "" {
			break
		}

		request.URL, _ = url.Parse(page.Next)
	}

	for _, highlight := range highlights {
		for _, tag := range highlight.Tags {
			if tag.Name == "favorite" {
				favorites = append(favorites, highlight)
			}
		}
	}

	if len(favorites) == 0 {
		return "", nil
	}

	randIndex := rand.Intn(len(favorites))

	return favorites[randIndex].Text, nil
}
