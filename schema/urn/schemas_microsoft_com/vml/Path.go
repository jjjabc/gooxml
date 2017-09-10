// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml

import (
	"encoding/xml"
	"fmt"
)

type Path struct {
	CT_Path
}

func NewPath() *Path {
	ret := &Path{}
	ret.CT_Path = *NewCT_Path()
	return ret
}

func (m *Path) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_Path.MarshalXML(e, start)
}

func (m *Path) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Path = *NewCT_Path()
	for _, attr := range start.Attr {
		if attr.Name.Local == "v" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.VAttr = &parsed
		}
		if attr.Name.Local == "limo" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LimoAttr = &parsed
		}
		if attr.Name.Local == "textboxrect" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TextboxrectAttr = &parsed
		}
		if attr.Name.Local == "fillok" {
			m.FillokAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "strokeok" {
			m.StrokeokAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "shadowok" {
			m.ShadowokAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "arrowok" {
			m.ArrowokAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "gradientshapeok" {
			m.GradientshapeokAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "textpathok" {
			m.TextpathokAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "insetpenok" {
			m.InsetpenokAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "connecttype" {
			m.ConnecttypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "connectlocs" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ConnectlocsAttr = &parsed
		}
		if attr.Name.Local == "connectangles" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ConnectanglesAttr = &parsed
		}
		if attr.Name.Local == "extrusionok" {
			m.ExtrusionokAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Path: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Path and its children
func (m *Path) Validate() error {
	return m.ValidateWithPath("Path")
}

// ValidateWithPath validates the Path and its children, prefixing error messages with path
func (m *Path) ValidateWithPath(path string) error {
	if err := m.CT_Path.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
