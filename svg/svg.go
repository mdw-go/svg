package svg

import (
	"encoding/xml"
	"fmt"
	"io"
)

type Writer struct{ inner io.Writer }

func NewWriter(writer io.Writer) *Writer { return &Writer{inner: writer} }

func (this *Writer) Header(width int, height int) *Writer {
	_, _ = fmt.Fprintf(this.inner, `<?xml version="1.0"?>`+"\n")
	_, _ = fmt.Fprintf(this.inner,
		`<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">`+"\n",
		width, height)
	return this
}
func (this *Writer) Footer() {
	_, _ = fmt.Fprintln(this.inner, "</svg>")
}
func (this *Writer) Circle(x, y, r int, style string) *Writer {
	_, _ = fmt.Fprintf(this.inner, `<circle cx="%d" cy="%d" r="%d" style="%s" />`+"\n", x, y, r, style)
	return this
}
func (this *Writer) Line(x1, y1, x2, y2 int, style string) *Writer {
	_, _ = fmt.Fprintf(this.inner, `<line x1="%d" y1="%d" x2="%d" y2="%d" style="%s" />`+"\n", x1, y1, x2, y2, style)
	return this
}
func (this *Writer) Rectangle(x, y, w, h int, style string) *Writer {
	_, _ = fmt.Fprintf(this.inner, `<rect x="%d" y="%d" width="%d" height="%d" style="%s" />`+"\n", x, y, w, h, style)
	return this
}
func (this *Writer) Text(x, y int, text, style string) *Writer {
	_, _ = fmt.Fprintf(this.inner, `<text x="%d" y="%d" style="%s">`, x, y, style)
	xml.Escape(this.inner, []byte(text))
	_, _ = fmt.Fprintln(this.inner, `</text>`)
	return this
}
