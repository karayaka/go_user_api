package handlers

import (
	"encoding/json"
	. "go_user_api/dataloaders"
	. "go_user_api/models"
	"net/http"
)

func Run() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8000", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//page nesnesi
	page := Page{ID: 7, Name: "Kullanıcılar", Description: "Kullanıcı Listesi", URI: "/users"}
	//dataloaders
	users := LoadUsers()
	print("denene")
	interests := LoadInterest()
	interestMappings := LoadInterestMapping()

	var newUsers []User
	for _, user := range users {
		for _, interestMapping := range interestMappings {
			if user.ID == interestMapping.UserID {
				for _, interest := range interests {
					if interestMapping.InterestID == interest.ID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
	}
	viewModel := UserViewModel{Page: page, Users: newUsers}
	data, _ := json.Marshal(viewModel)
	w.Write([]byte(data))
}
