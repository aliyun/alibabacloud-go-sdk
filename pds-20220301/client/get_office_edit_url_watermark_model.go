// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOfficeEditUrlWatermark interface {
	dara.Model
	String() string
	GoString() string
	SetFillstyle(v string) *GetOfficeEditUrlWatermark
	GetFillstyle() *string
	SetFont(v string) *GetOfficeEditUrlWatermark
	GetFont() *string
	SetHorizontal(v int64) *GetOfficeEditUrlWatermark
	GetHorizontal() *int64
	SetRotate(v float64) *GetOfficeEditUrlWatermark
	GetRotate() *float64
	SetType(v int32) *GetOfficeEditUrlWatermark
	GetType() *int32
	SetValue(v string) *GetOfficeEditUrlWatermark
	GetValue() *string
	SetVertical(v int64) *GetOfficeEditUrlWatermark
	GetVertical() *int64
}

type GetOfficeEditUrlWatermark struct {
	Fillstyle  *string  `json:"fillstyle,omitempty" xml:"fillstyle,omitempty"`
	Font       *string  `json:"font,omitempty" xml:"font,omitempty"`
	Horizontal *int64   `json:"horizontal,omitempty" xml:"horizontal,omitempty"`
	Rotate     *float64 `json:"rotate,omitempty" xml:"rotate,omitempty"`
	Type       *int32   `json:"type,omitempty" xml:"type,omitempty"`
	Value      *string  `json:"value,omitempty" xml:"value,omitempty"`
	Vertical   *int64   `json:"vertical,omitempty" xml:"vertical,omitempty"`
}

func (s GetOfficeEditUrlWatermark) String() string {
	return dara.Prettify(s)
}

func (s GetOfficeEditUrlWatermark) GoString() string {
	return s.String()
}

func (s *GetOfficeEditUrlWatermark) GetFillstyle() *string {
	return s.Fillstyle
}

func (s *GetOfficeEditUrlWatermark) GetFont() *string {
	return s.Font
}

func (s *GetOfficeEditUrlWatermark) GetHorizontal() *int64 {
	return s.Horizontal
}

func (s *GetOfficeEditUrlWatermark) GetRotate() *float64 {
	return s.Rotate
}

func (s *GetOfficeEditUrlWatermark) GetType() *int32 {
	return s.Type
}

func (s *GetOfficeEditUrlWatermark) GetValue() *string {
	return s.Value
}

func (s *GetOfficeEditUrlWatermark) GetVertical() *int64 {
	return s.Vertical
}

func (s *GetOfficeEditUrlWatermark) SetFillstyle(v string) *GetOfficeEditUrlWatermark {
	s.Fillstyle = &v
	return s
}

func (s *GetOfficeEditUrlWatermark) SetFont(v string) *GetOfficeEditUrlWatermark {
	s.Font = &v
	return s
}

func (s *GetOfficeEditUrlWatermark) SetHorizontal(v int64) *GetOfficeEditUrlWatermark {
	s.Horizontal = &v
	return s
}

func (s *GetOfficeEditUrlWatermark) SetRotate(v float64) *GetOfficeEditUrlWatermark {
	s.Rotate = &v
	return s
}

func (s *GetOfficeEditUrlWatermark) SetType(v int32) *GetOfficeEditUrlWatermark {
	s.Type = &v
	return s
}

func (s *GetOfficeEditUrlWatermark) SetValue(v string) *GetOfficeEditUrlWatermark {
	s.Value = &v
	return s
}

func (s *GetOfficeEditUrlWatermark) SetVertical(v int64) *GetOfficeEditUrlWatermark {
	s.Vertical = &v
	return s
}

func (s *GetOfficeEditUrlWatermark) Validate() error {
	return dara.Validate(s)
}
