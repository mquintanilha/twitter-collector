package main
import (
    "context"
    "fmt"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
    ID              bson.ObjectId   `bson:"_id" json:"id"`
    Time_stamp      string  `bson:"name" json:"name"`
    User            User
    Place           Place
    Entities        Entities
    Language        string  `bson:"lang" json:"lang"`
}

type User struct {
    Id              int     `bson:"id" json:"id"`
    Name            string  `bson:"name" json:"name"`
    Followers_count int     `bson:"followers_count" json:"followers_count"`
}

type Place struct  {
    Name    string  `bson:"name" json:"name"`
    Country_code    string  `bson:"country_code", json:"country_code"`
}

type Entities struct {
    Hashtags    Hashtags
}

type Hashtags struct {
    Text    string  `bson:"text" json:"text"`
}

func main() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")

}