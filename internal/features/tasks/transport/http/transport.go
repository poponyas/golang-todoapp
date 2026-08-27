package tasks_transport_http

import (
	"context"
	"net/http"

	"github.com/poponyas/golang-todoapp/internal/core/domain"
	core_http_server "github.com/poponyas/golang-todoapp/internal/core/transport/http/server"
)

type TasksHTTPHandler struct {
	tasksService TasksService
}

type TasksService interface {
	CreateTask(
		ctx context.Context,
		task domain.Task,
	) (domain.Task, error)

	GetTasks(
		ctx context.Context,
		userID *int,
		limit *int,
		offset *int,
	) ([]domain.Task, error)

	GetTask(
		ctx context.Context,
		id int,
	) (domain.Task, error)

	DeleteTask(
		ctx context.Context,
		id int,
	) error

	PatchTask(
		ctx context.Context,
		id int,
		patch domain.TaskPatch,
	) (domain.Task, error)
}

func NewTasksHTTPHandler(tasksService TasksService) *TasksHTTPHandler {
	return &TasksHTTPHandler{
		tasksService: tasksService,
	}
}

func (h *TasksHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method: http.MethodPost,
			Path:   "/tasks",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				h.CreateTask(w, r)
			},
		},
		{
			Method: http.MethodGet,
			Path:   "/tasks",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				h.GetTasks(w, r)
			},
		},
		{
			Method: http.MethodGet,
			Path:   "/tasks/{id}",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				h.GetTask(w, r)
			},
		},
		{
			Method: http.MethodDelete,
			Path:   "/tasks/{id}",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				h.DeleteTask(w, r)
			},
		},
		{
			Method: http.MethodPatch,
			Path:   "/tasks/{id}",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				h.PatchTask(w, r)
			},
		},
	}
}
