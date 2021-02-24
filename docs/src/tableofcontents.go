package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type tableOfContents struct {
	app.Compo

	Ilinks   []string
	selected string
}

func newTableOfContents() *tableOfContents {
	return &tableOfContents{}
}

func (t *tableOfContents) Links(v ...string) *tableOfContents {
	t.Ilinks = v
	return t
}

func (t *tableOfContents) OnNav(ctx app.Context) {
	t.selected = "#" + ctx.Page.URL().Fragment
	t.Update()
}

func (t *tableOfContents) Render() app.UI {
	return app.Aside().
		Class("pane").
		Class("index").
		Body(
			app.H1().Text("Index"),
			app.Section().Body(
				app.Range(t.Ilinks).Slice(func(i int) app.UI {
					link := t.Ilinks[i]

					return &tableOfContentLink{
						Title: link,
						Focus: t.selected == githubIndex(link),
					}
				}),
			),

			app.Section().Body(
				&tableOfContentLink{
					Title: "Report issue",
					Focus: t.selected == "#report-issue",
				},
				&tableOfContentLink{
					Title: "Support go-app",
					Focus: t.selected == "#support-go-app",
				},
			),
		)
}

type tableOfContentLink struct {
	app.Compo

	Title string
	Focus bool
}

func (l *tableOfContentLink) Render() app.UI {
	focus := ""
	if l.Focus {
		focus = "focus"
	}

	return app.A().
		Class(focus).
		Class("index-link").
		Href(githubIndex(l.Title)).
		Text(l.Title)
}
