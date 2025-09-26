// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBrowserViewPort interface {
	dara.Model
	String() string
	GoString() string
	SetHeight(v float64) *BrowserViewPort
	GetHeight() *float64
	SetWidth(v float64) *BrowserViewPort
	GetWidth() *float64
}

type BrowserViewPort struct {
	Height *float64 `json:"height,omitempty" xml:"height,omitempty"`
	Width  *float64 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s BrowserViewPort) String() string {
	return dara.Prettify(s)
}

func (s BrowserViewPort) GoString() string {
	return s.String()
}

func (s *BrowserViewPort) GetHeight() *float64 {
	return s.Height
}

func (s *BrowserViewPort) GetWidth() *float64 {
	return s.Width
}

func (s *BrowserViewPort) SetHeight(v float64) *BrowserViewPort {
	s.Height = &v
	return s
}

func (s *BrowserViewPort) SetWidth(v float64) *BrowserViewPort {
	s.Width = &v
	return s
}

func (s *BrowserViewPort) Validate() error {
	return dara.Validate(s)
}
