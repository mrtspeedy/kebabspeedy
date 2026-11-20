package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/a-h/templ"
	"github.com/mrtspeedy/kebabspeedy/internal/views/pages"
)

func renderPage(
	ctx context.Context,
	path string,
	component templ.Component,
) error {

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	// Create the output HTML file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	// Make sure file is closed when done writing
	defer file.Close()

	// Render the component to the file
	return component.Render(ctx, file)
}

func main() {
	ctx := context.Background()

	// Define every page to generate a HTML file for
	pages := []struct {
		path string
		component templ.Component
	}{
		{
			path:      "public/index.html",
			component: pages.HomePage(),
		},
		{
			path: "public/about/index.html",
			component: pages.AboutPage(),
		},
		{
			path: "public/404.html",
			component: pages.NotFoundPage(),
		},
	}

	// Render every page
	for _, page := range pages {
		if err := renderPage(ctx, page.path, page.component); err != nil {
			log.Fatalf("Failed to render %s: %v", page.path, err)
		}

		log.Printf("Rendered: %s", page.path)
	}

	log.Println("Generation complete.")
}
