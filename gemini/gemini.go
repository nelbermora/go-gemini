package gemini

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func Init() {
	/*ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	var err error
	Client, err = genai.NewClient(ctx, option.WithAPIKey("AIzaSyDxrBg5kH9z1-IVpbOd7CBZ0AwOzDRaJkw"))
	if err != nil {
		log.Fatal(err)
	}
	//defer Client.Close()

	Model = Client.GenerativeModel("gemini-1.0-pro")
	*/
}

func AskForData(input, key string) ([]string, error) {
	ctx := context.Background()
	output := []string{}

	client, err := genai.NewClient(ctx, option.WithAPIKey(key))
	if err != nil {
		return output, err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.0-pro")

	resp, err := model.GenerateContent(ctx, genai.Text(input))
	if err != nil {
		return output, err
	}

	output = printResponse(resp)
	/*
		for _, cand := range resp.Candidates {
			if cand.Content != nil {
				for _, part := range cand.Content.Parts {
					output = append(output, fmt.Sprint(part))
				}
			}
		}*/
	return output, nil
}

func printResponse(resp *genai.GenerateContentResponse) []string {
	var otp []string
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
				otps := strings.Split(fmt.Sprint(part), "\n")
				otp = append(otp, otps...)
			}
		}
	}
	fmt.Println("---")
	return otp
}
