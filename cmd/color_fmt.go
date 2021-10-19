package cmd

import "fmt"

func green(s string) string {
	return fmt.Sprintf("\033[32m %s \033[0m", s)
}


func red(s string) string {
	return fmt.Sprintf("\033[31m %s \033[0m", s)
}

func purple(s string) string {
	return fmt.Sprintf("\033[35m %s \033[0m", s)
}