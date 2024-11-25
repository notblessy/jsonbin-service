package model

import (
	"encoding/json"
	"time"
)

type PublicJSON struct {
	ID        string          `json:"id" gorm:"primaryKey;unique"`
	Data      json.RawMessage `json:"data"`
	CreatedAt time.Time       `json:"created_at" gorm:"autoCreateTime"`
}

func (p PublicJSON) TableName() string {
	return "public_json"
}
