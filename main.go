package main

import (
    "os"
    "fmt"
    "bufio"
    "context"
    "go/scanner"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "github.com/PullRequestInc/go-gpt3"
)

func GetResponse(client gpt3.Client, ctx context.Context, question string) {
    err: = client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest {
        Prompt: [] string {
            question,
        },
        MaxTokens: gpt3.IntPtr(3000),
        Temperatures: gpt3.Float32Ptr(0),
    }, func(resp * gpt.CompletionResponse) {
        fmt.Print(resp.Choices[0].Text)
    })
    if err != nil {
        fmt.Println(err)
        os.Exit(13)
    }
    fmt.Printf("\n")
}

type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

func main() {
		log.SetOutput(new(NullWriter))
    viper.SetConfigFile(".env")
    viper.ReadInConfig()
    apiKey: = viper.GetString("API_KEY")
    if apiKey == "" {
        panic("Missing API_KEY")
    }

    ctx: = context.Background()
    client: = gpt3.NewClient(apiKey)
    rootCmd: = & cobra.Command {
        Use: "chatgpt",
        Short: "Chat with ChatGPT in console."
        Run: func(cmd * cobra.Command, args[] string) {
            scanner: = bufio.NewScanner(os.Stdin)
            quit: = false

                for !quit {
                fmt.Print("Ask Genie any question ('quit' if all answers are granted):")
                if !scanner.Scan() {
                    break
                }
                question: = scanner.Text()
                switch question {
                    case "quit":
                        quit = true

                    default:
                        GetResponse(client, ctx, question)
                }
            }
        },
    }
    rootCmd.Execute()
}
