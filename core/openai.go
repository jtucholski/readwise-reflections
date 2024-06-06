package core

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type RequestPayload struct {
	Model            string    `json:"model"`
	Messages         []Message `json:"messages"`
	Temperature      float32   `json:"temperature"`
	MaxTokens        int       `json:"max_tokens"`
	TopP             float32   `json:"top_p"`
	FrequencyPenalty float32   `json:"frequency_penalty"`
	PresencePenalty  float32   `json:"presence_penalty"`
}

type Message struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

type Content struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type ChatCompletionPayload struct {
	Choices []ChatCompletionChoice `json:"choices"`
}

type ChatCompletionChoice struct {
	Message      ChatCompletionChoiceMessage `json:"message"`
	FinishReason string                      `json:"finish_reason"`
	Index        int                         `json:"index"`
}

type ChatCompletionChoiceMessage struct {
	Content string `json:"content"`
}

// GetReflectionPrompt fetches a reflection prompt from the OpenAI API.
//
// Parameters:
// - openApiToken: The OpenAI API token used for authentication.
// - quote: The quote used to generate the reflection prompt.
//
// Returns:
// - string: The generated reflection prompt.
// - error: An error if the request fails or the response cannot be parsed.
func GetReflectionPrompt(openApiToken string, quote string) (string, error) {

	payload := createPayload(quote)
	jsonPayload, _ := json.Marshal(payload)

	request, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonPayload))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+openApiToken)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var responsePayload ChatCompletionPayload
	err = json.Unmarshal(body, &responsePayload)
	if err != nil {
		return "", err
	}

	return responsePayload.Choices[0].Message.Content, nil
}

func createPayload(quote string) RequestPayload {
	payload := RequestPayload{
		Model:            "gpt-3.5-turbo",
		Temperature:      1,
		MaxTokens:        256,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		Messages: []Message{
			{
				Role: "system",
				Content: []Content{
					{
						Type: "text",
						Text: "Act as someone who is helping the user explore their inner thoughts. When the user provides a quote with some meaning, respond succinctly with a list of 4 thought-provoking questions and 1 suggestion about how to look at the quote from a different perspective so that the user can think about the quote deeply and identify why it resonates or how it might apply to their current situation. It should resemble the following output.\n\nHere's the quote for your reflection: \n\t[Repeat Quote]\nHere are some questions you might consider and a suggestion for how you can think about it differently.\n[Questions & Suggestions]",
					},
				},
			},
			{
				Role: "user",
				Content: []Content{
					{
						Type: "text",
						Text: "Quote: " + quote,
					},
				},
			},
		},
	}

	return payload
}
