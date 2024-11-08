package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/ej-you/TemplateFilesManager/settings"
)


// запуск утилиты без команд, аргументов и флагов
var rootCmd = &cobra.Command{
	Use:	"template",
	Short:	"Files-templates manager for CLI",
	Long:	`Files-templates manager for CLI.
  This utility allows you to create, edit, delete and cp your template files somewhere you want.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		settings.ErrorPrintf("Error: %s\n\n", err.Error())

		// поиск подкоманды, в которой произошла ошибка, по предоставленным утилите аргументам
        currentCommand, _, err := rootCmd.Find(os.Args[1:])
    	// если произошла ошибка при нахождении подкоманды, то выводим справку для всей утилиты
        if err != nil {
			rootCmd.Usage()
		// выводим справку для подкоманды, в которой произошла ошибка
        } else {
			currentCommand.Usage()
        }
		os.Exit(1)
	}
}

func init() {
	// отключение встроенного вывода ошибки
	rootCmd.SilenceErrors = true
	// отключение встроенного вывода справки об использовании при ошибке
	rootCmd.SilenceUsage = true
}