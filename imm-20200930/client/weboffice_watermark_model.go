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
	// The color and transparency of the text watermark.
	//
	// example:
	//
	// rgba(192, 192, 192, 0.6)
	FillStyle *string `json:"FillStyle,omitempty" xml:"FillStyle,omitempty"`
	// The font of the text watermark.
	//
	// example:
	//
	// bold 20px Serif
	Font *string `json:"Font,omitempty" xml:"Font,omitempty"`
	// The horizontal spacing of the text watermark. Unit: pixel.
	//
	// example:
	//
	// 50
	Horizontal *int64 `json:"Horizontal,omitempty" xml:"Horizontal,omitempty"`
	// The rotation of the text watermark. Unit: radian.
	//
	// example:
	//
	// -0.7853982
	Rotate *float32 `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// The watermark type. Valid values:
	//
	// 	- 0: no watermark.
	//
	// 	- 1: text watermark.
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The watermark text.
	//
	// >  This parameter takes effect only if you set the Type parameter to 1.
	//
	// example:
	//
	// example
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The vertical spacing of the text watermark. Unit: pixel.
	//
	// example:
	//
	// 100
	Vertical *int64 `json:"Vertical,omitempty" xml:"Vertical,omitempty"`
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
