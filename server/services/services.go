package services

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)



func OpenAI(userMsg string, opts) *openai.ChatCompletion {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("WTF Mati... Where is .env file?")
	}
	api_key := os.Getenv("OPEN_AI_API_KEY")

	client := openai.NewClient(option.WithAPIKey(api_key))
	ctx := context.TODO()

	param := openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(opts),
		}),
		Seed:  openai.Int(1),
		Model: openai.F(openai.ChatModelGPT4o),
	}

	completion, err := client.Chat.Completions.New(ctx, param)

	if err != nil {
		log.Fatal(err)
	}
	param.Messages.Value = append(param.Messages.Value, completion.Choices[0].Message)
	param.Messages.Value = append(param.Messages.Value, openai.UserMessage(userMsg))

	completion, err = client.Chat.Completions.New(ctx, param)

	if err != nil {
		log.Fatal(err)
	}
	return completion
}
