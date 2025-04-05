package routes

import (
	"go-beginner/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	// Health check route
	r.Get("/health", app.HealthCheck)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutById)

	// Add more routes here as needed
	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)

	return r
}
