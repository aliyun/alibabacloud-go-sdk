// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncludeImage interface {
	dara.Model
	String() string
	GoString() string
	SetHeight(v int32) *IncludeImage
	GetHeight() *int32
	SetImageLink(v string) *IncludeImage
	GetImageLink() *string
	SetWidth(v int32) *IncludeImage
	GetWidth() *int32
}

type IncludeImage struct {
	// example:
	//
	// 324
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// example:
	//
	// http://k.sinaimg.cn/n/sinakd20121/594/w2048h946/20240328/74cf-32c0d62e843df76567d760b4459d57c1.jpg/w700d1q75cms.jpg
	ImageLink *string `json:"imageLink,omitempty" xml:"imageLink,omitempty"`
	// example:
	//
	// 700
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s IncludeImage) String() string {
	return dara.Prettify(s)
}

func (s IncludeImage) GoString() string {
	return s.String()
}

func (s *IncludeImage) GetHeight() *int32 {
	return s.Height
}

func (s *IncludeImage) GetImageLink() *string {
	return s.ImageLink
}

func (s *IncludeImage) GetWidth() *int32 {
	return s.Width
}

func (s *IncludeImage) SetHeight(v int32) *IncludeImage {
	s.Height = &v
	return s
}

func (s *IncludeImage) SetImageLink(v string) *IncludeImage {
	s.ImageLink = &v
	return s
}

func (s *IncludeImage) SetWidth(v int32) *IncludeImage {
	s.Width = &v
	return s
}

func (s *IncludeImage) Validate() error {
	return dara.Validate(s)
}
