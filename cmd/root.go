package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gerador-de-token",
	Short: "Token Generator CLI",
	Long:  "Uma CLI para gerar tokens em formato HEX ou BASE64 de um determinado tamanho.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'gerador-de-token generate --help' para ver mais informações.")
	},
}

// Execute roda o comando raiz
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
