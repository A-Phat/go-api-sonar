package api

import (
	"encoding/json"
	"go-api-sonar/internal/math"
	"net/http"
	"strconv"
)

func MathHandler(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))
	result := math.Add(a, b)
	result = math.DuplicateAdd1(a, b)
	result = math.DuplicateAdd2(a, b)

	json.NewEncoder(w).Encode(map[string]int{"result": result})
}
