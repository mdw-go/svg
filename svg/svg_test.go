package svg

import (
	"bytes"
	"testing"
)

func TestSVGWriter(t *testing.T) {
	var buffer bytes.Buffer
	writer := NewWriter(&buffer)
	writer.Header(42, 43)
	writer.Text(44, 45, "Hello, world!", "font-family:monospace")
	writer.Line(46, 47, 48, 49, "stroke:black")
	writer.Circle(50, 51, 52, "fill:black")
	writer.Rectangle(53, 54, 55, 56, "fill:white")
	writer.Footer()

	actual := buffer.String()
	if actual != expected {
		t.Errorf("\nwant:%s\ngot:\n%s", expected, actual)
	}
}

const expected = `<?xml version="1.0"?>
<svg width="42" height="43" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
<text x="44" y="45" style="font-family:monospace">Hello, world!</text>
<line x1="46" y1="47" x2="48" y2="49" style="stroke:black" />
<circle cx="50" cy="51" r="52" style="fill:black" />
<rect x="53" y="54" width="55" height="56" style="fill:white" />
</svg>
`
