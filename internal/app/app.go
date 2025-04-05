package app

import (
	"database/sql"
	"go-beginner/internal/api"
	"go-beginner/internal/store"
	"log"
	"net/http"
	"os"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	Db             *sql.DB
}

func NewApplication() (*Application, error) {
	pgDb, err := store.Open()
	if err != nil {
		return nil, err
	}

	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	app := &Application{
		Logger:         logger,
		WorkoutHandler: api.NewWorkoutHandler(),
		Db:             pgDb,
	}

	return app, nil
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
	app.Logger.Println("Health check passed")
}
