package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-fastapi"
)

var ( // Create by GoRelease at compile time
	version = "dev"
	commit  = "none"
	//date    = "unknown"
)

type EchoInput struct {
	Phrase string `json:"phrase"`
}

type EchoOutput struct {
	OriginalInput EchoInput `json:"original_input"`
}

func EchoHandler(ctx *gin.Context, in EchoInput) (out EchoOutput, err error) {
	out.OriginalInput = in
	return
}

func mainApp() error {
	r := gin.Default()
	myRouter := fastapi.NewRouter()
	myRouter.AddCall("/echo", EchoHandler)
	r.POST("/api/*path", myRouter.GinHandler) // Must have *path parameter
	return r.Run()                            // Return the error from r.Run()
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	if err := mainApp(); err != nil { // Handle the returned error
		log.Fatal(err)
	}
}
