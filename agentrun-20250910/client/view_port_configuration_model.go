// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewPortConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetHeight(v float32) *ViewPortConfiguration
	GetHeight() *float32
	SetWidth(v float32) *ViewPortConfiguration
	GetWidth() *float32
}

type ViewPortConfiguration struct {
	// 视口高度（像素）
	//
	// This parameter is required.
	Height *float32 `json:"height,omitempty" xml:"height,omitempty"`
	// 视口宽度（像素）
	//
	// This parameter is required.
	Width *float32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s ViewPortConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ViewPortConfiguration) GoString() string {
	return s.String()
}

func (s *ViewPortConfiguration) GetHeight() *float32 {
	return s.Height
}

func (s *ViewPortConfiguration) GetWidth() *float32 {
	return s.Width
}

func (s *ViewPortConfiguration) SetHeight(v float32) *ViewPortConfiguration {
	s.Height = &v
	return s
}

func (s *ViewPortConfiguration) SetWidth(v float32) *ViewPortConfiguration {
	s.Width = &v
	return s
}

func (s *ViewPortConfiguration) Validate() error {
	return dara.Validate(s)
}
