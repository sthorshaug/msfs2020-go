.PHONY: vendor
vendor:
	# download leafletjs
	rm -rf _vendor
	mkdir -p _vendor/leafletjs
	curl http://cdn.leafletjs.com/leaflet/v1.7.1/leaflet.zip --output _vendor/leafletjs/leaflet.zip
	unzip _vendor/leafletjs/leaflet.zip -d _vendor/leafletjs
	rm _vendor/leafletjs/leaflet.zip

	# download rotatedMarker plugin for leaflet
	mkdir -p _vendor/rotatedMarker
	curl https://codeload.github.com/bbecquet/Leaflet.RotatedMarker/zip/master --output _vendor/rotatedMarker/source.zip
	unzip _vendor/rotatedMarker/source.zip -d _vendor/RotatedMarker

	# download tailwindcss
	mkdir -p _vendor/tailwind
	curl --location https://unpkg.com/tailwindcss@^1.0/dist/tailwind.min.css --output _vendor/tailwind/tailwind.min.css

	# download alpine.js
	mkdir -p _vendor/alpinejs
	curl --location https://cdn.jsdelivr.net/gh/alpinejs/alpine@v2.7.0/dist/alpine.min.js --output _vendor/alpinejs/alpine.min.js

	# copy all files needed for go-bindata into a separate directory
	mkdir -p _vendor/bindata
	cp _vendor/RotatedMarker/Leaflet.RotatedMarker-master/leaflet.rotatedMarker.js _vendor/bindata/
	cp -r _vendor/leafletjs/leaflet.js _vendor/leafletjs/leaflet.css _vendor/leafletjs/images _vendor/bindata/
	cp _vendor/tailwind/tailwind.min.css _vendor/bindata/
	cp _vendor/alpinejs/alpine.min.js _vendor/bindata/


.PHONY: generate
generate:
	go generate
