package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

const (
	prefix         = "DUMMY_SRV"
	defaultNetAddr = "0.0.0.0"
	defaultNetPort = "8080"
)

var (
	config            *Config
	xForwardedHeaders = []string{
		"Forwarded",
		"X-Forwarded-For",
		"X-Forwarded-Host",
		"X-Forwarded-Port",
		"X-Forwarded-Proto",
		"X-Forwarded-Proto-Version",
		"X-Real-Ip",
		"Host",
	}
)

func init() {
	var err error
	config, err = NewConfig(prefix)
	if err != nil {
		log.Fatal(err)
	}

}

type Config struct {
	env *viper.Viper
}

func NewConfig(prefix string) (*Config, error) {
	config := &Config{}
	env := viper.New()
	env.SetDefault("net_addr", defaultNetAddr)
	env.SetDefault("net_port", defaultNetPort)
	env.AllowEmptyEnv(false)
	env.SetEnvPrefix(prefix)
	env.AutomaticEnv()
	config.env = env
	return config, nil
}

func (config *Config) NetAddr() string {
	return config.env.GetString("net_addr")
}
func (config *Config) NetPort() string {
	return config.env.GetString("net_port")
}

func (config *Config) ListenerAddr() string {
	return fmt.Sprintf("%s:%s", config.NetAddr(), config.NetPort())
}

func headersInfoHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var headers = make(map[string]interface{}, 0)
		for _, headerName := range xForwardedHeaders {
			headerVal := ctx.GetHeader(headerName)
			if headerVal != "" {
				headers[headerName] = headerVal
			}
		}
		r := map[string]interface{}{"headers": headers, "client_ip": ctx.ClientIP()}
		ctx.JSON(200, r)
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/headers", headersInfoHandler())
	if err := r.Run(config.ListenerAddr()); err != nil {
		log.Fatal(err)
	}
}
