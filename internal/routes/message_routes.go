/*
Copyright © 2024 Cranom Technologies Limited, Beingana Jim Junior and Contributors
*/
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/open-ug/conveyor/internal/handlers"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func DriverRoutes(app *fiber.App, db *mongo.Database, redisClient *redis.Client) {
	applicationPrefix := app.Group("/drivers")
	applicationHandler := handlers.NewMessageHandler(db, redisClient)

	applicationPrefix.Post("/broadcast-message", applicationHandler.PublishMessage)
}
