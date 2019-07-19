package main

import (
	"github.com/gin-gonic/gin"
	"github.com/linnv/swag/example/basic/api"
	"github.com/linnv/swag/example/basic/docs"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.New()
	r.GET("/testapi/get-string-by-int/:some_id", api.GetStringByInt)
	r.GET("//testapi/get-struct-array-by-string/:some_id", api.GetStructArrayByString)
	r.POST("/testapi/upload", api.Upload)
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()

}
