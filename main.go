package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	svc := Service{}
	handler := Handler{
		Service: svc,
	}

	ginEngine := gin.Default()
	ginEngine.GET("/range-fizzbuzz", handler.FizzBuzzHandler)

	port := 8082
	address := fmt.Sprintf(":%d", port)

	server := &http.Server{
		Addr:         address,
		Handler:      ginEngine,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("shutting down gracefully....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		fmt.Println("forced to shutdown", err)
	}

	fmt.Println("shutted down")
}
