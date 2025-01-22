// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/mailru/easyjson"
)

func main() {
	// Create a sample user for testing.
	user := User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		CreatedAt: time.Now(),
	}

	// Test serialization and deserialization.
	testSerialization(user)
}

// testSerialization tests serialization and deserialization for both standard JSON and easyjson.
func testSerialization(user User) {
	// Standard JSON Serialization.
	standardJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Standard JSON Marshal failed: %v", err)
	}
	fmt.Println("Standard JSON:", string(standardJSON))

	// Standard JSON Deserialization.
	var userStandard User
	err = json.Unmarshal(standardJSON, &userStandard)
	if err != nil {
		log.Fatalf("Standard JSON Unmarshal failed: %v", err)
	}
	fmt.Println("Standard Deserialized User:", userStandard)

	// easyjson Serialization.
	easyJSON, err := easyjson.Marshal(user)
	if err != nil {
		log.Fatalf("easyjson Marshal failed: %v", err)
	}
	fmt.Println("easyjson:", string(easyJSON))

	// easyjson Deserialization.
	var userEasy User
	err = easyjson.Unmarshal(easyJSON, &userEasy)
	if err != nil {
		log.Fatalf("easyjson Unmarshal failed: %v", err)
	}
	fmt.Println("easyjson Deserialized User:", userEasy)
}
