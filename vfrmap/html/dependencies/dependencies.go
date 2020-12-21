package dependencies

import (
	"log"
	"net/http"
)

//go:generate go-bindata -modtime 1 -pkg dependencies -o bindata.go -prefix "_vendor/bindata" "_vendor/bindata" "_vendor/bindata/images"

type FS struct {
}

// MustAsset loads an assets or fails execution if it cannot be loaded
func MustAsset(name string) []byte {
	asset, err := Asset(name)
	if err != nil {
		log.Fatalf("Could not load asset %s", name)
	}

	return asset
}

func (_ FS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "alpine.min.js":
		w.Header().Set("Content-Type", "text/css")
		w.Write(MustAsset("alpine.min.js"))
	case "images/layers-2x.png":
		w.Header().Set("Content-Type", "image/png")
		w.Write(MustAsset("images/layers-2x.png"))
	case "images/layers.png":
		w.Header().Set("Content-Type", "image/png")
		w.Write(MustAsset("images/layers.png"))
	case "images/marker-icon-2x.png":
		w.Header().Set("Content-Type", "image/png")
		w.Write(MustAsset("images/marker-icon-2x.png"))
	case "images/marker-icon.png":
		w.Header().Set("Content-Type", "image/png")
		w.Write(MustAsset("images/marker-icon.png"))
	case "images/marker-shadow.png":
		w.Header().Set("Content-Type", "image/png")
		w.Write(MustAsset("images/marker-shadow.png"))
	case "leaflet.css":
		w.Header().Set("Content-Type", "text/css")
		w.Write(MustAsset("leaflet.css"))
	case "leaflet.js":
		w.Header().Set("Content-Type", "text/javascript")
		w.Write(MustAsset("leaflet.js"))
	case "leaflet.rotatedMarker.js":
		w.Header().Set("Content-Type", "text/javascript")
		w.Write(MustAsset("leaflet.rotatedMarker.js"))
	case "tailwind.min.css":
		w.Header().Set("Content-Type", "text/css")
		w.Write(MustAsset("tailwind.min.css"))
	}
}
