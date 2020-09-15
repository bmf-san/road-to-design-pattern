package main

import "fmt"

type builder interface {
	makeHeader(header string)
	makeBody(body string)
	makeFooter(footer string)
}

type Article struct {
	Header string
	Body   string
	Footer string
}

type Director struct {
	builder builder
}

func NewDirector(builder builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) Construct() {
	d.builder.makeHeader("header")
	d.builder.makeBody("body")
	d.builder.makeFooter("footer")
}

type HTMLBuilder struct {
	Article *Article
}

func NewHTMLBuilder() *HTMLBuilder {
	return &HTMLBuilder{
		Article: &Article{},
	}
}

func (h *HTMLBuilder) makeHeader(header string) {
	h.Article.Header = fmt.Sprintf("<header>%s</header>", header)
}

func (h *HTMLBuilder) makeBody(body string) {
	h.Article.Body = fmt.Sprintf("<div>%s</div>", body)
}

func (h *HTMLBuilder) makeFooter(footer string) {
	h.Article.Footer = fmt.Sprintf("<footer>%s</footer>", footer)
}

func (h *HTMLBuilder) GetArticle() {
	fmt.Println(h.Article.Header)
	fmt.Println(h.Article.Body)
	fmt.Println(h.Article.Footer)
}

type MarkdownBuilder struct {
	Article *Article
}

func NewMarkdownBuilder() *MarkdownBuilder {
	return &MarkdownBuilder{
		Article: &Article{},
	}
}

func (m *MarkdownBuilder) makeHeader(header string) {
	m.Article.Header = fmt.Sprintf("#%s", header)
}

func (m *MarkdownBuilder) makeBody(body string) {
	m.Article.Body = fmt.Sprintf("-%s", body)
}

func (m *MarkdownBuilder) makeFooter(footer string) {
	m.Article.Footer = fmt.Sprintf("##%s", footer)
}

func (m *MarkdownBuilder) GetArticle() {
	fmt.Println(m.Article.Header)
	fmt.Println(m.Article.Body)
	fmt.Println(m.Article.Footer)
}

func main() {
	opt := "html"

	switch opt {
	case "html":
		builderHTML := NewHTMLBuilder()
		director := NewDirector(builderHTML)
		director.Construct()
		builderHTML.GetArticle()
	case "markdown":
		builderMarkdown := NewMarkdownBuilder()
		director := NewDirector(builderMarkdown)
		director.Construct()
		builderMarkdown.GetArticle()
	}
}
