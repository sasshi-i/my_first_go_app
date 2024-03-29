package main

import (
	"net/http"
)

func authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParthFrom()
	user, _ = data.UserbyEmail(r.PostFromValue("email"))
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session := user.CreateSession()
		cookie := http.Cookie{
			Name:			"_cookie",
			Value:		session.Uuid,
			Httponly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "login", 302)
	}
}