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
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type OfcOLEObject struct {
	OfcCT_OLEObject
}

func NewOfcOLEObject() *OfcOLEObject {
	ret := &OfcOLEObject{}
	ret.OfcCT_OLEObject = *NewOfcCT_OLEObject()
	return ret
}

func (m *OfcOLEObject) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:o"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:v"}, Value: "urn:schemas-microsoft-com:vml"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "o:OLEObject"
	return m.OfcCT_OLEObject.MarshalXML(e, start)
}

func (m *OfcOLEObject) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.OfcCT_OLEObject = *NewOfcCT_OLEObject()
	for _, attr := range start.Attr {
		if attr.Name.Local == "Type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "ProgID" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ProgIDAttr = &parsed
		}
		if attr.Name.Local == "ShapeID" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ShapeIDAttr = &parsed
		}
		if attr.Name.Local == "DrawAspect" {
			m.DrawAspectAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "ObjectID" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ObjectIDAttr = &parsed
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
		}
		if attr.Name.Local == "UpdateMode" {
			m.UpdateModeAttr.UnmarshalXMLAttr(attr)
		}
	}
lOfcOLEObject:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "LinkType"}:
				m.LinkType = new(string)
				if err := d.DecodeElement(m.LinkType, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "LockedField"}:
				m.LockedField = sharedTypes.ST_TrueFalseBlankUnset
				if err := d.DecodeElement(&m.LockedField, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "FieldCodes"}:
				m.FieldCodes = new(string)
				if err := d.DecodeElement(m.FieldCodes, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on OfcOLEObject %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lOfcOLEObject
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the OfcOLEObject and its children
func (m *OfcOLEObject) Validate() error {
	return m.ValidateWithPath("OfcOLEObject")
}

// ValidateWithPath validates the OfcOLEObject and its children, prefixing error messages with path
func (m *OfcOLEObject) ValidateWithPath(path string) error {
	if err := m.OfcCT_OLEObject.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
