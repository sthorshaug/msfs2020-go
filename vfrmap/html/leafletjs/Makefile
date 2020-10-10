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

	# copy all files needed for go-bindata into a separate directory
	mkdir -p _vendor/leafletjs_bindata
	cp _vendor/rotatedMarker/Leaflet.RotatedMarker-master/leaflet.rotatedMarker.js _vendor/leafletjs_bindata/
	cp -r _vendor/leafletjs/leaflet.js _vendor/leafletjs/leaflet.css _vendor/leafletjs/images   _vendor/leafletjs_bindata/


.PHONY: generate
generate:
	go generate