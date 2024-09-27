package dataloaders

import (
	"encoding/json"
	. "go_user_api/models"
	util "go_user_api/utils"
)

func LoadUsers() []User {
	bytes, _ := util.ReadFile("json/users.json")
	var data []User
	json.Unmarshal(bytes, &data)
	return data
}

func LoadInterest() []Interest {
	bytes, _ := util.ReadFile("json/interests.json")
	var data []Interest
	json.Unmarshal(bytes, &data)
	return data
}

func LoadInterestMapping() []InterestMapping {
	bytes, _ := util.ReadFile("json/userInterestMappings.json")
	var data []InterestMapping
	json.Unmarshal(bytes, &data)
	return data
}
