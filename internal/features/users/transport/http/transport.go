package users_transport_http

import (
	"context"
	"net/http"

	"github.com/poponyas/golang-todoapp/internal/core/domain"
	core_http_server "github.com/poponyas/golang-todoapp/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	usersService UserService
}

type UserService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)

	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)

	GetUser(
		ctx context.Context,
		id int,
	) (domain.User, error)

	DeleteUser(
		ctx context.Context,
		id int,
	) error

	PatchUser(
		ctx context.Context,
		id int,
		patch domain.UserPatch,
	) (domain.User, error)
}

func NewUsersHTTPHandler(
	usersService UserService,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: usersService,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method: http.MethodPost,
			Path:   "/users",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				h.CreateUser(w, r)
			},
		},
		{
			Method: http.MethodGet,
			Path:   "/users",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				h.GetUsers(w, r)
			},
		},
		{
			Method: http.MethodGet,
			Path:   "/users/{id}",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				h.GetUser(w, r)
			},
		},
		{
			Method: http.MethodDelete,
			Path:   "/users/{id}",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				h.DeleteUser(w, r)
			},
		},
		{
			Method: http.MethodPatch,
			Path:   "/users/{id}",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				h.PatchUser(w, r)
			},
		},
	}
}
