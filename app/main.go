package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello world")

	// files are located here
	//fs := http.FileServer(http.Dir("./static"))
	// handler def
	//http.Handle("/", fs)

	var router *gin.Engine
	router = gin.Default()
	router.Static("/", "./static")
	router.Run()

	log.Println("starting port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
