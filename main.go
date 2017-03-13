package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"

	"google.golang.org/appengine"
)

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "static/playercount.html", 301)
	})
	http.HandleFunc("/setup", renderTileCounts)

	appengine.Main()
	//http.ListenAndServe(":8080", nil)
}

var tileLimits = []int{
	24, //Grain
	21, //Cheese
	18, //Wine
	15, //Wool
	12, //Brocade
}

const (
	mapSpacesFourPl  = 52
	mapSpacesThreePl = 48
	mapSpacesTwoPl   = 43

	removeTilesFourPl  = 0
	removeTilesThreePl = 6
	removeTilesTwoPl   = 12
)

type tileSetup struct {
	PlayerCount  string
	TilesUsed    []int
	TilesRemoved []int
	Note         string
}

func renderTileCounts(w http.ResponseWriter, r *http.Request) {
	//TODO:move this out later for performance. Don't need to parse on each request. Nice for rapid dev though
	var tmpl = template.Must(template.ParseFiles("assets/setup.html"))

	playerCount := r.FormValue("playerCount")
	var mapSpaces, removeTiles int
	var note string
	switch playerCount {
	case "3":
		mapSpaces = mapSpacesThreePl
		removeTiles = removeTilesThreePl
		note = "Remove 2 each of Farmer, Boatman, Craftsman Trader. Remove 3 each of Knight, Scholar, Monk"
	case "2":
		mapSpaces = mapSpacesTwoPl
		removeTiles = removeTilesTwoPl
		note = "Remove 4 each of Farmer, Boatman, Craftsman Trader. Remove 6 each of Knight, Scholar, Monk"
	default:
		playerCount = "4"
		mapSpaces = mapSpacesFourPl
		removeTiles = removeTilesFourPl
	}

	tilesUsed := calcTileNumbers(mapSpaces, tileLimits)
	tilesRemoved := calcTileNumbers(removeTiles, subtract(tileLimits, tilesUsed))

	tmpl.Execute(w, tileSetup{playerCount, tilesUsed, tilesRemoved, note})
}

// a - b elementwise
func subtract(a, b []int) []int {
	if len(a) != len(b) {
		panic("cannot subtract slices of different lengths")
	}
	result := make([]int, len(a))
	for k := range a {
		if a[k] < b[k] {
			panic(fmt.Sprintf("Trying to subract %v from %v. Index %v too large", b, a, k))
		}
		result[k] = a[k] - b[k]
	}
	return result
}

// calcTileNumbers returns a slice containing the randomised starting
// number of tiles. Each slice entry represents a different type of tile.
// 'required' is the total number of required tiles.
// 'limits' represents the pool of available tiles to draw from
func calcTileNumbers(required int, limits []int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	tilePool := make([]int, 0, required)
	for tileType := 0; tileType < len(limits); tileType++ {
		for n := 0; n < limits[tileType]; n++ {
			tilePool = append(tilePool, tileType)
		}
	}
	if len(tilePool) < required {
		panic("Too few tiles to sastisfy requirement")
	}
	randomTileIndices := r.Perm(len(tilePool))

	output := make([]int, len(limits))
	for i := 0; i < required; i++ {
		output[tilePool[randomTileIndices[i]]]++
	}
	return output
}
