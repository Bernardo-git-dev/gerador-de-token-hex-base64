package cmd

import (
	"fmt"

	"github.com/Bernardo-git-dev/gerador-de-token-hex-base64/utils"
	"github.com/spf13/cobra"
)

// Flags para o comando
var format string
var size int

// Comando generate
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Gere um token em formato HEX ou BASE64",
	Long:  "Gera um token criptograficamente seguro em formato HEX ou BASE64 com o tamanho especificado.",
	Run: func(cmd *cobra.Command, args []string) {
		if size <= 0 {
			fmt.Println("Tamanho inválido. Por favor, insira um número inteiro positivo.")
			return
		}

		var token string
		var err error

		switch format {
		case "hex":
			token, err = utils.GenerateHexToken(size)
		case "base64":
			token, err = utils.GenerateBase64Token(size)
		default:
			fmt.Println("Formato inválido. Use 'hex' ou 'base64'.")
			return
		}

		if err != nil {
			fmt.Printf("Erro ao gerar o token: %v\n", err)
			return
		}

		fmt.Println("\nToken gerado com sucesso:")
		fmt.Println("------------------------------------------------------------------")
		fmt.Println(token)
		fmt.Println("------------------------------------------------------------------")
	},
}

func init() {
	// Adicione o comando "generate" ao comando raiz
	rootCmd.AddCommand(generateCmd)

	// Configure as flags
	generateCmd.Flags().StringVarP(&format, "format", "f", "hex", "Formato do token (hex ou base64)")
	generateCmd.Flags().IntVarP(&size, "size", "s", 32, "Tamanho do token em bytes")
}
