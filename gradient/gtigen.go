// Code generated by "goki generate ./..."; DO NOT EDIT.

package gradient

import (
	"goki.dev/colors"
	"goki.dev/gti"
	"goki.dev/mat32/v2"
	"goki.dev/ordmap"
)

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/colors/gradient.Base",
	ShortName: "gradient.Base",
	IDName:    "base",
	Doc:       "Base contains the data and logic common to all gradient types.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{"-setters"}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Stops", &gti.Field{Name: "Stops", Type: "[]goki.dev/colors/gradient.Stop", LocalType: "[]Stop", Doc: "the stops for the gradient; use AddStop to add stops", Directives: gti.Directives{}, Tag: "set:\"-\""}},
		{"Spread", &gti.Field{Name: "Spread", Type: "goki.dev/colors/gradient.Spreads", LocalType: "Spreads", Doc: "the spread method used for the gradient if it stops before the end", Directives: gti.Directives{}, Tag: ""}},
		{"Blend", &gti.Field{Name: "Blend", Type: "goki.dev/colors.BlendTypes", LocalType: "colors.BlendTypes", Doc: "the colorspace algorithm to use for blending colors", Directives: gti.Directives{}, Tag: ""}},
		{"Units", &gti.Field{Name: "Units", Type: "goki.dev/colors/gradient.Units", LocalType: "Units", Doc: "the units to use for the gradient", Directives: gti.Directives{}, Tag: ""}},
		{"Box", &gti.Field{Name: "Box", Type: "goki.dev/mat32/v2.Box2", LocalType: "mat32.Box2", Doc: "the bounding box of the object with the gradient; this is used when rendering\ngradients with [Units] of [ObjectBoundingBox]; this must be set using SetTransform\nto recompute necessary values.", Directives: gti.Directives{}, Tag: "set:\"-\""}},
		{"Transform", &gti.Field{Name: "Transform", Type: "goki.dev/mat32/v2.Mat2", LocalType: "mat32.Mat2", Doc: "Transform is the transformation matrix applied to the gradient's points;\nit must be set using SetTransform to recompute necessary values.", Directives: gti.Directives{}, Tag: "set:\"-\""}},
		{"ObjectMatrix", &gti.Field{Name: "ObjectMatrix", Type: "goki.dev/mat32/v2.Mat2", LocalType: "mat32.Mat2", Doc: "ObjectMatrix is the effective object transformation matrix for a gradient\nwith [Units] of [ObjectBoundingBox]. It should not be set by end users, but\nmust be updated using [Base.ComputeObjectMatrix] whenever [Base.Box] or\n[Base.Transform] is updated, which happens automatically in SetBox and SetTransform.", Directives: gti.Directives{}, Tag: "set:\"-\""}},
	}),
	Embeds:  ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

// SetSpread sets the [Base.Spread]:
// the spread method used for the gradient if it stops before the end
func (t *Base) SetSpread(v Spreads) *Base {
	t.Spread = v
	return t
}

// SetBlend sets the [Base.Blend]:
// the colorspace algorithm to use for blending colors
func (t *Base) SetBlend(v colors.BlendTypes) *Base {
	t.Blend = v
	return t
}

// SetUnits sets the [Base.Units]:
// the units to use for the gradient
func (t *Base) SetUnits(v Units) *Base {
	t.Units = v
	return t
}

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/colors/gradient.Linear",
	ShortName: "gradient.Linear",
	IDName:    "linear",
	Doc:       "Linear represents a linear gradient. It implements the [image.Image] interface.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{"-setters"}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Start", &gti.Field{Name: "Start", Type: "goki.dev/mat32/v2.Vec2", LocalType: "mat32.Vec2", Doc: "the starting point of the gradient (x1 and y1 in SVG)", Directives: gti.Directives{}, Tag: ""}},
		{"End", &gti.Field{Name: "End", Type: "goki.dev/mat32/v2.Vec2", LocalType: "mat32.Vec2", Doc: "the ending point of the gradient (x2 and y2 in SVG)", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Base", &gti.Field{Name: "Base", Type: "goki.dev/colors/gradient.Base", LocalType: "Base", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

// SetStart sets the [Linear.Start]:
// the starting point of the gradient (x1 and y1 in SVG)
func (t *Linear) SetStart(v mat32.Vec2) *Linear {
	t.Start = v
	return t
}

// SetEnd sets the [Linear.End]:
// the ending point of the gradient (x2 and y2 in SVG)
func (t *Linear) SetEnd(v mat32.Vec2) *Linear {
	t.End = v
	return t
}

// SetSpread sets the [Linear.Spread]
func (t *Linear) SetSpread(v Spreads) *Linear {
	t.Spread = v
	return t
}

// SetBlend sets the [Linear.Blend]
func (t *Linear) SetBlend(v colors.BlendTypes) *Linear {
	t.Blend = v
	return t
}

// SetUnits sets the [Linear.Units]
func (t *Linear) SetUnits(v Units) *Linear {
	t.Units = v
	return t
}

var _ = gti.AddType(&gti.Type{
	Name:      "goki.dev/colors/gradient.Radial",
	ShortName: "gradient.Radial",
	IDName:    "radial",
	Doc:       "Radial represents a radial gradient. It implements the [image.Image] interface.",
	Directives: gti.Directives{
		&gti.Directive{Tool: "gti", Directive: "add", Args: []string{"-setters"}},
	},
	Fields: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Center", &gti.Field{Name: "Center", Type: "goki.dev/mat32/v2.Vec2", LocalType: "mat32.Vec2", Doc: "the center point of the gradient (cx and cy in SVG)", Directives: gti.Directives{}, Tag: ""}},
		{"Focal", &gti.Field{Name: "Focal", Type: "goki.dev/mat32/v2.Vec2", LocalType: "mat32.Vec2", Doc: "the focal point of the gradient (fx and fy in SVG)", Directives: gti.Directives{}, Tag: ""}},
		{"Radius", &gti.Field{Name: "Radius", Type: "goki.dev/mat32/v2.Vec2", LocalType: "mat32.Vec2", Doc: "the radius of the gradient (rx and ry in SVG)", Directives: gti.Directives{}, Tag: ""}},
	}),
	Embeds: ordmap.Make([]ordmap.KeyVal[string, *gti.Field]{
		{"Base", &gti.Field{Name: "Base", Type: "goki.dev/colors/gradient.Base", LocalType: "Base", Doc: "", Directives: gti.Directives{}, Tag: ""}},
	}),
	Methods: ordmap.Make([]ordmap.KeyVal[string, *gti.Method]{}),
})

// SetCenter sets the [Radial.Center]:
// the center point of the gradient (cx and cy in SVG)
func (t *Radial) SetCenter(v mat32.Vec2) *Radial {
	t.Center = v
	return t
}

// SetFocal sets the [Radial.Focal]:
// the focal point of the gradient (fx and fy in SVG)
func (t *Radial) SetFocal(v mat32.Vec2) *Radial {
	t.Focal = v
	return t
}

// SetRadius sets the [Radial.Radius]:
// the radius of the gradient (rx and ry in SVG)
func (t *Radial) SetRadius(v mat32.Vec2) *Radial {
	t.Radius = v
	return t
}

// SetSpread sets the [Radial.Spread]
func (t *Radial) SetSpread(v Spreads) *Radial {
	t.Spread = v
	return t
}

// SetBlend sets the [Radial.Blend]
func (t *Radial) SetBlend(v colors.BlendTypes) *Radial {
	t.Blend = v
	return t
}

// SetUnits sets the [Radial.Units]
func (t *Radial) SetUnits(v Units) *Radial {
	t.Units = v
	return t
}
