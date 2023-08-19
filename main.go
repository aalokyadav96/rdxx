package main

import (
    "log"
    "os"
    "context"
    "github.com/redis/go-redis/v9"
    "fmt"

    "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  redis_url := os.Getenv("REDIS_URL")
	log.Println(redis_url)
	ExampleClient(redis_url)
  // now do something with s3 or whatever
}

var ctx = context.Background()

func ExampleClient(url string) {
    rdb := redis.NewClient(&redis.Options{
        Addr:     url,
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    err := rdb.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := rdb.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)

    val2, err := rdb.Get(ctx, "key2").Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exist")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("key2", val2)
    }
    // Output: key value
    // key2 does not exist
}