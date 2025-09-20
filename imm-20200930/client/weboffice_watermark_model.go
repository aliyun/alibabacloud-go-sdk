// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebofficeWatermark interface {
	dara.Model
	String() string
	GoString() string
	SetFillStyle(v string) *WebofficeWatermark
	GetFillStyle() *string
	SetFont(v string) *WebofficeWatermark
	GetFont() *string
	SetHorizontal(v int64) *WebofficeWatermark
	GetHorizontal() *int64
	SetRotate(v float32) *WebofficeWatermark
	GetRotate() *float32
	SetType(v int64) *WebofficeWatermark
	GetType() *int64
	SetValue(v string) *WebofficeWatermark
	GetValue() *string
	SetVertical(v int64) *WebofficeWatermark
	GetVertical() *int64
}

type WebofficeWatermark struct {
	// example:
	//
	// rgba(192, 192, 192, 0.6)
	FillStyle *string `json:"FillStyle,omitempty" xml:"FillStyle,omitempty"`
	// example:
	//
	// bold 20px Serif
	Font       *string  `json:"Font,omitempty" xml:"Font,omitempty"`
	Horizontal *int64   `json:"Horizontal,omitempty" xml:"Horizontal,omitempty"`
	Rotate     *float32 `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	Type       *int64   `json:"Type,omitempty" xml:"Type,omitempty"`
	Value      *string  `json:"Value,omitempty" xml:"Value,omitempty"`
	Vertical   *int64   `json:"Vertical,omitempty" xml:"Vertical,omitempty"`
}

func (s WebofficeWatermark) String() string {
	return dara.Prettify(s)
}

func (s WebofficeWatermark) GoString() string {
	return s.String()
}

func (s *WebofficeWatermark) GetFillStyle() *string {
	return s.FillStyle
}

func (s *WebofficeWatermark) GetFont() *string {
	return s.Font
}

func (s *WebofficeWatermark) GetHorizontal() *int64 {
	return s.Horizontal
}

func (s *WebofficeWatermark) GetRotate() *float32 {
	return s.Rotate
}

func (s *WebofficeWatermark) GetType() *int64 {
	return s.Type
}

func (s *WebofficeWatermark) GetValue() *string {
	return s.Value
}

func (s *WebofficeWatermark) GetVertical() *int64 {
	return s.Vertical
}

func (s *WebofficeWatermark) SetFillStyle(v string) *WebofficeWatermark {
	s.FillStyle = &v
	return s
}

func (s *WebofficeWatermark) SetFont(v string) *WebofficeWatermark {
	s.Font = &v
	return s
}

func (s *WebofficeWatermark) SetHorizontal(v int64) *WebofficeWatermark {
	s.Horizontal = &v
	return s
}

func (s *WebofficeWatermark) SetRotate(v float32) *WebofficeWatermark {
	s.Rotate = &v
	return s
}

func (s *WebofficeWatermark) SetType(v int64) *WebofficeWatermark {
	s.Type = &v
	return s
}

func (s *WebofficeWatermark) SetValue(v string) *WebofficeWatermark {
	s.Value = &v
	return s
}

func (s *WebofficeWatermark) SetVertical(v int64) *WebofficeWatermark {
	s.Vertical = &v
	return s
}

func (s *WebofficeWatermark) Validate() error {
	return dara.Validate(s)
}
