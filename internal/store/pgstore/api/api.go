package api

import (
	"net/http"

	"github.com/Giovani-O/ama-go-server.git/internal/store/pgstore/pgstore"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type apiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	a := apiHandler{
		q: q,
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Route("/rooms", func(r chi.Router) {
			r.Post("/", a.handleCreateRoom)
			r.Get("/", a.handleGetRooms)

			r.Route("/{room_id}/messages", func(r chi.Router) {
				r.Post("/", a.handleCreateRoomMessage)
				r.Get("/", a.handleGetRoomMessages)

				r.Route("/{message_id}", func(r chi.Router) {
					r.Get("/", a.handleGetRoomMessage)
					r.Patch("/react", a.handleReactToMessage)
					r.Delete("/react", a.handleRemoveReactionFromMessage)
					r.Patch("/answer", a.handleMarkMessageAsAnswered)
				})
			})
		})
	})

	a.r = r
	return a
}

func (h apiHandler) handleCreateRoom(q http.ResponseWriter, r *http.Request)                {}
func (h apiHandler) handleGetRooms(q http.ResponseWriter, r *http.Request)                  {}
func (h apiHandler) handleCreateRoomMessage(q http.ResponseWriter, r *http.Request)         {}
func (h apiHandler) handleGetRoomMessages(q http.ResponseWriter, r *http.Request)           {}
func (h apiHandler) handleGetRoomMessage(q http.ResponseWriter, r *http.Request)            {}
func (h apiHandler) handleReactToMessage(q http.ResponseWriter, r *http.Request)            {}
func (h apiHandler) handleRemoveReactionFromMessage(q http.ResponseWriter, r *http.Request) {}
func (h apiHandler) handleMarkMessageAsAnswered(q http.ResponseWriter, r *http.Request)     {}
