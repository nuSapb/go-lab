package main

import (
	"log"
	"net/http"
)

func main() {

	// allowRoleAdmin := allowRole("admin")
	// allowRoleStaft := allowRole("staft")
	http.HandleFunc("/", indexHandler)
	http.Handle("/admin", allowRoles("admin")(http.HandlerFunc(adminHandler)))
	http.Handle("/staft", allowRoles("staft")(http.HandlerFunc(staftHandler)))
	http.Handle("/admin-staft", allowRoles("admin", "staft")(http.HandlerFunc(adminStaftHandler)))
	err := http.ListenAndServe(":8070", nil)
	log.Println(err)
}

type middleware func(http.Handler) http.Handler

func allowRole(role string) middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRole := r.Header.Get("Role")
			if reqRole != role {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func allowRoles(roles ...string) middleware {
	allow := make(map[string]struct{})
	for _, role := range roles {
		allow[role] = struct{}{}
	}
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRole := r.Header.Get("Role")
			if _, ok := allow[reqRole]; !ok {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello."))
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Admin."))
}

func staftHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Staft."))
}

func adminStaftHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Admin and Staft."))
}
