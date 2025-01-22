// serialization_test.go
package main

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/mailru/easyjson"
)

func BenchmarkSerialization(b *testing.B) {
	user := User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		CreatedAt: time.Now(),
	}

	b.Run("Standard JSON Marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = json.Marshal(user)
		}
	})

	b.Run("easyjson Marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = easyjson.Marshal(user)
		}
	})

	b.Run("Standard JSON Unmarshal", func(b *testing.B) {
		data, _ := json.Marshal(user)
		var u User
		for i := 0; i < b.N; i++ {
			_ = json.Unmarshal(data, &u)
		}
	})

	b.Run("easyjson Unmarshal", func(b *testing.B) {
		data, _ := easyjson.Marshal(user)
		var u User
		for i := 0; i < b.N; i++ {
			_ = easyjson.Unmarshal(data, &u)
		}
	})
}
