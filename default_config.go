package brain_decoder

import (
	"bytes"
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"strings"
)

var (
	OpenAIClient  *openai.Client
	DefaultPrompt = EnglishPrompt
	DefaultAI     = GPT3Dot5Turbo
)

func EnglishPrompt(source string, schema string) string {
	return fmt.Sprintf("Read %s and output it as JSON according to the JSONScheme %s specification", source, schema)
}

func ChinesePrompt(source string, schema string) string {
	return fmt.Sprintf("阅读 %s ，然后按照JSONSchema %s 规范输出为 JSON", source, schema)
}

func GPT3Dot5Turbo(prompt string) ([]byte, error) {
	completion, err := OpenAIClient.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	output := completion.Choices[0].Message.Content
	output = strings.TrimPrefix(output, "```json")
	output = strings.TrimSuffix(output, "```")
	return []byte(output), nil
}

func GPT3Dot5Turbo16K(prompt string) ([]byte, error) {
	completion, err := OpenAIClient.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo16K,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	output := completion.Choices[0].Message.Content
	output = strings.TrimPrefix(output, "```json")
	output = strings.TrimSuffix(output, "```")
	return []byte(output), nil
}

func GPT4(prompt string) ([]byte, error) {
	completion, err := OpenAIClient.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	output := completion.Choices[0].Message.Content
	output = strings.TrimPrefix(output, "```json")
	output = strings.TrimSuffix(output, "```")
	return []byte(output), nil
}

func Unmarshal(source []byte, v any) error {
	return NewDecoder(bytes.NewReader(source), DefaultAI, DefaultPrompt, 3).Encode(v)
}
