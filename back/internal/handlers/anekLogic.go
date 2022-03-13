package anekLogic

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ubahwin/anekdots-site/internal/models"
	"net/http"
)

func GetAnekdot(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "project.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM anek ORDER BY RANDOM() LIMIT 1")
	anek := models.Anekdot{}
	row.Scan(&anek.Id, &anek.CategoryId, &anek.Text)

	anekdotDto := models.AnekdotDto{}
	db.QueryRow("SELECT name FROM category WHERE id = ?", anek.CategoryId).Scan(&anekdotDto.CategoryName)
	anekdotDto.Id = anek.Id
	anekdotDto.Text = anek.Text

	json.NewEncoder(w).Encode(anekdotDto)
}
