package main

import "snippetbox.seifeddinetouj.ifb/internal/models"

type templateData struct {
    Snippet models.Snippet
    Snippets []models.Snippet
}