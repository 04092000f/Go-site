package main

import (
    "github.com/gin-gonic/contrib/renders/multitemplate"
)

func makeTemplates() multitemplate.Render {
    templates := multitemplate.New()

    templates.AddFromFiles("templates",
        "index.html",
        "page1.html",
    )

    return templates
}