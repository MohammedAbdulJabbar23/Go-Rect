package main

import(
  "example.com/event-booking/db"
  "example.com/event-booking/routes"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
)

func main() {
  db.InitDB();
  server := gin.Default();
  config := cors.DefaultConfig()
  config.AllowOrigins = []string{"http://localhost:3000"} // Update with your React app's origin
  config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
  server.Use(cors.New(config))
  routes.RegisterRoutes(server);
  server.Run(":8080");
}


 
