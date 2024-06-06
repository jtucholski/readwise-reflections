# README

This is my first application written in Go. It's a command-line application that integrates with two APIs to enhance the user's experience with insightful prompts:
* [Readwise API](https://readwise.io/api_deets) - For retrieving book highlights.
* [OpenAI API](https://platform.openai.com/) - For generating thought-provoking prompts based on the quotes.

The application retrieves a random favorited quote from Readwise's API and then uses the OpenAI API to generate a thought-provoking prompt based on the quote. The prompt is displayed to the user and copied to their clipboard so that they can use it in their note-taking software.

## How to use

To run this application, you need to set up environment variables for API access:
1. `READWISE_TOKEN` - Your token for the Readwise API.
2. `OPENAI_TOKEN` - Your token for the OpenAI API.

You can set these variables manually, or use an `.env` file for easier management:

```bash
cp .env.example ./cmd/.env
```

Navigate to the `cmd` directory and run the application using Go:

```bash
cd cmd
go run .
```

## Known issues

As this is my first application written in Go, I'm sure there are some issues that I don't know about. 

Here are some known limitations:
- **Inefficiency:** The current method retrieves all highlights and filteres for favorites post-fetch. This is due to the API limitation, as it only returns 1000 highlights per request and offers no way to filter the response server-side.

## Potential Enhancements

- **User input:** Allow users to choose between retrieving a random highlight or entering their own.
- **Model selection:** Allow users to select which model to use for the OpenAI API.