package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

// Person model
type Person struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string             `bson:"name" json:"name"`
	Email string             `bson:"email" json:"email"`
	CPF   string             `bson:"cpf" json:"cpf"`
}

func main() {
	// Setup MongoDB connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	collection = client.Database("testdb").Collection("people")

	// Setup Gin router
	r := gin.Default()

	r.POST("/people", createPerson)
	r.GET("/people", getPeople)
	r.GET("/people/:id", getPerson)
	r.PUT("/people/:id", updatePerson)
	r.DELETE("/people/:id", deletePerson)

	r.Run(":8080")
}

// createPerson handles adding a new person
func createPerson(c *gin.Context) {
	var person Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	person.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, person)
}

// getPeople retrieves all people from the database
func getPeople(c *gin.Context) {
	var people []Person
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var person Person
		if err := cursor.Decode(&person); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		people = append(people, person)
	}
	c.JSON(http.StatusOK, people)
}

// getPerson retrieves a single person by ID
func getPerson(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var person Person
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&person)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, person)
}

// updatePerson updates a person by ID
func updatePerson(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var person Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": person})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, person)
}

// deletePerson removes a person by ID
func deletePerson(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Person deleted"})
}
