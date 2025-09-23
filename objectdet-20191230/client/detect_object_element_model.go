// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectObjectElement interface {
	dara.Model
	String() string
	GoString() string
	SetHeight(v int64) *DetectObjectElement
	GetHeight() *int64
	SetScore(v float32) *DetectObjectElement
	GetScore() *float32
	SetType(v string) *DetectObjectElement
	GetType() *string
	SetWidth(v int64) *DetectObjectElement
	GetWidth() *int64
	SetX(v int64) *DetectObjectElement
	GetX() *int64
	SetY(v int64) *DetectObjectElement
	GetY() *int64
}

type DetectObjectElement struct {
	// example:
	//
	// 200
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 0.68225745
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// VEHICLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 100
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 5
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 10
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DetectObjectElement) String() string {
	return dara.Prettify(s)
}

func (s DetectObjectElement) GoString() string {
	return s.String()
}

func (s *DetectObjectElement) GetHeight() *int64 {
	return s.Height
}

func (s *DetectObjectElement) GetScore() *float32 {
	return s.Score
}

func (s *DetectObjectElement) GetType() *string {
	return s.Type
}

func (s *DetectObjectElement) GetWidth() *int64 {
	return s.Width
}

func (s *DetectObjectElement) GetX() *int64 {
	return s.X
}

func (s *DetectObjectElement) GetY() *int64 {
	return s.Y
}

func (s *DetectObjectElement) SetHeight(v int64) *DetectObjectElement {
	s.Height = &v
	return s
}

func (s *DetectObjectElement) SetScore(v float32) *DetectObjectElement {
	s.Score = &v
	return s
}

func (s *DetectObjectElement) SetType(v string) *DetectObjectElement {
	s.Type = &v
	return s
}

func (s *DetectObjectElement) SetWidth(v int64) *DetectObjectElement {
	s.Width = &v
	return s
}

func (s *DetectObjectElement) SetX(v int64) *DetectObjectElement {
	s.X = &v
	return s
}

func (s *DetectObjectElement) SetY(v int64) *DetectObjectElement {
	s.Y = &v
	return s
}

func (s *DetectObjectElement) Validate() error {
	return dara.Validate(s)
}
