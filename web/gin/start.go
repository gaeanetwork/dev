package ginserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	//
	_ "net/http/pprof"

	"github.com/gaeanetwork/gaea-core/api"
	"github.com/gaeanetwork/gaea-core/api/auth"
	"github.com/gaeanetwork/gaea-core/api/tee"
	"github.com/gaeanetwork/gaea-core/common/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Start gin web http
func Start() {
	r := setupRouter()
	gracefulStart(r)
}

func setupRouter() *gin.Engine {
	// Switch to "release" mode in production.
	gin.SetMode(gin.ReleaseMode)

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Serving static files
	r.Static("/site", "/dist")

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Allows all origins
	r.Use(cors.Default())

	apiRG := r.Group("/api")
	{
		// common module - Ping test - curl http://localhost:8080/api/ping
		apiRG.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

		// user module
		api.RegisterTransmissionAPI(apiRG) // Transmission api register
		auth.RegisterAPI(apiRG)            // Register / Login / Logout

		// tee data share module
		tee.RegisterSharedDataAPI(apiRG)
	}
	return r
}

func gracefulStart(router *gin.Engine) {
	srv := &http.Server{
		Addr:    config.ListenAddr,
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Start profiling http endpoint if enabled
	if config.ProfileEnabled {
		go func() {
			log.Printf("Starting profiling server with listenAddress = %s\n", config.PProfAddr)
			if profileErr := http.ListenAndServe(config.PProfAddr, nil); profileErr != nil {
				log.Panicf("Error starting profiler: %v\n", profileErr)
			}
		}()
	}

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
