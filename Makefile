build:
	@echo "Generating Templ..."
	templ generate
	@echo "Building CSS..."
	npx @tailwindcss/cli -i ./internal/assets/css/input.css -o ./public/css/tw.css --minify
	@echo "Generating static HTML pages..."
	go run ./cmd/kebabspeedy/main.go