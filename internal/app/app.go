package app

import (
	"database/sql"
	"go-beginner/internal/api"
	"go-beginner/internal/store"
	"go-beginner/migrations"
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

	err = store.MigrateFS(pgDb, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	workoutStore := store.NewPostgresWorkoutStore(pgDb)

	app := &Application{
		Logger:         logger,
		WorkoutHandler: api.NewWorkoutHandler(workoutStore),
		Db:             pgDb,
	}

	return app, nil
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
	app.Logger.Println("Health check passed")
}
