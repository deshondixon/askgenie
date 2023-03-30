package main

import (
	"context"
	"go/scanner"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
		viper.SetConfigFile(".env")
		viper.ReadInConfig()
		apiKey := viper.GetString("API_KEY")
		if apiKey == "" {
				panic("Missing API_KEY")
		}

		ctx := context.Background()
		client := gpt3.NewClient(apiKey)
		rootCmd := &cobra.Command{
				Use: "chatgpt", 
				Short: "Chat with ChatGPT in console."
				Run: func(cmd *cobra.Command, args []string) {
					scanner := bufio.NewScanner(os.Stdin)
					quit := false

					for  !quit{
							 fmt.Print("Ask Genie any question ('quit' if all answers are granted):")
							 if !scanner.Scan() {
								 break
							 }
							 question := scanner.Text()
							 switch question {
								case "quit":
							 }
					}
		}
	}
}
