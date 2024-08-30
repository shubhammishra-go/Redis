package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "name1", "Shubham Mishra", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "name1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name1 :: ", val)

	// Output: for key2
	// key2 does not exist

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	// storing JSON keys values

	type Person struct {
		ID         string `json: "id"`
		Name       string `json: "name"`
		Age        int    `json: "age"`
		Occupation string `json: "occupation"`
	}

	jsonString, err := json.Marshal(

		Person{
			ID:         uuid.NewString(),
			Name:       "Shubham Mishra",
			Age:        24,
			Occupation: "Software Engineer",
		})

	if err != nil {
		fmt.Println(err)
		return
	}

	err = rdb.Set(ctx, "person1", jsonString, 0).Err()

	if err != nil {
		fmt.Println(err)
		return
	}

	val, err = rdb.Get(ctx, "person1").Result()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Person1 ::: ", val)

}
