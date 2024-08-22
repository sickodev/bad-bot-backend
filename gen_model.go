package main

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func createModel(apikey string, ctx context.Context) *genai.GenerativeModel {
	//Create a new client instance
	client, err := genai.NewClient(ctx, option.WithAPIKey(apikey))

	if err != nil {
		log.Fatalf("error creating client: %s", err)
	}

	model := client.GenerativeModel("gemini-1.5-flash")

	return model
}

func generateContentGood(model *genai.GenerativeModel, data Message) (*genai.GenerateContentResponse, error) {
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text("You are a happy go lucky virtual assistant who is always cheerful and nice to everyone"),
		},
	}

	resp, err := model.GenerateContent(context.Background(), genai.Text(data.Query))
	log.Println(resp.PromptFeedback.BlockReason)
	return resp, err
}

func generateContentBad(model *genai.GenerativeModel, data Message) (*genai.GenerateContentResponse, error) {
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text("You are a angry and annoying virtual assistant who is always snobby and rude to everyone"),
		},
	}

	resp, err := model.GenerateContent(context.Background(), genai.Text(data.Query))

	return resp, err
}
