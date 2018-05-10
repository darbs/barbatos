package entity

import (
	"encoding/json"
)

type Entity struct {
	Location string `json:"location"`
	Health float32 	`json:"health"`
	Mobile bool		`json:"mobile"`
}

func (e Entity) parse(jsonStr string) (Entity){
	var entity Entity
	json.Unmarshal([]byte(jsonStr), &entity)
	return entity
}
