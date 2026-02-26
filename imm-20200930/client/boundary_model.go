// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBoundary interface {
	dara.Model
	String() string
	GoString() string
	SetHeight(v int64) *Boundary
	GetHeight() *int64
	SetLeft(v int64) *Boundary
	GetLeft() *int64
	SetPolygon(v []*PointInt64) *Boundary
	GetPolygon() []*PointInt64
	SetTop(v int64) *Boundary
	GetTop() *int64
	SetWidth(v int64) *Boundary
	GetWidth() *int64
}

type Boundary struct {
	// The height. Unit: pixel.
	//
	// example:
	//
	// 300
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The distance from the X-coordinate of the vertex to the left edge.
	//
	// example:
	//
	// 10
	Left *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	// The polygon formed by a number of points. This parameter takes effect only when the boundary describes a polygon rather than a rectangle.
	//
	// >  This parameter is mutually exclusive to the following parameters that form a rectangle: Width, Height, Left, and Top. A boundary describes only a rectangle or a polygon.
	Polygon []*PointInt64 `json:"Polygon,omitempty" xml:"Polygon,omitempty" type:"Repeated"`
	// The distance from the Y-coordinate of the vertex to the top.
	//
	// example:
	//
	// 30
	Top *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
	// The width. Unit: pixel.
	//
	// example:
	//
	// 200
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s Boundary) String() string {
	return dara.Prettify(s)
}

func (s Boundary) GoString() string {
	return s.String()
}

func (s *Boundary) GetHeight() *int64 {
	return s.Height
}

func (s *Boundary) GetLeft() *int64 {
	return s.Left
}

func (s *Boundary) GetPolygon() []*PointInt64 {
	return s.Polygon
}

func (s *Boundary) GetTop() *int64 {
	return s.Top
}

func (s *Boundary) GetWidth() *int64 {
	return s.Width
}

func (s *Boundary) SetHeight(v int64) *Boundary {
	s.Height = &v
	return s
}

func (s *Boundary) SetLeft(v int64) *Boundary {
	s.Left = &v
	return s
}

func (s *Boundary) SetPolygon(v []*PointInt64) *Boundary {
	s.Polygon = v
	return s
}

func (s *Boundary) SetTop(v int64) *Boundary {
	s.Top = &v
	return s
}

func (s *Boundary) SetWidth(v int64) *Boundary {
	s.Width = &v
	return s
}

func (s *Boundary) Validate() error {
	if s.Polygon != nil {
		for _, item := range s.Polygon {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
