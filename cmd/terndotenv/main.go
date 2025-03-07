package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	fmt.Println("GOBID_DATABASE_HOST:", os.Getenv("GOBID_DATABASE_HOST"))
	fmt.Println("GOBID_DATABASE_USER:", os.Getenv("GOBID_DATABASE_USER"))
	fmt.Println("GOBID_DATABASE_PASSWORD:", os.Getenv("GOBID_DATABASE_PASSWORD"))
	

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/migrations",
		"--config",
		"./internal/store/migrations/tern.conf",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Command execution failed: ", err)
		fmt.Println("Output:", string(output))
		panic(err)
	}

	fmt.Println("Command executed sucessfully ", string(output))
}
