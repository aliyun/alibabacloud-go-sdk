// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferPortraitStyleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHeight(v int32) *TransferPortraitStyleRequest
	GetHeight() *int32
	SetImageUrl(v string) *TransferPortraitStyleRequest
	GetImageUrl() *string
	SetNumbers(v int32) *TransferPortraitStyleRequest
	GetNumbers() *int32
	SetRedrawAmplitude(v int32) *TransferPortraitStyleRequest
	GetRedrawAmplitude() *int32
	SetStyle(v int32) *TransferPortraitStyleRequest
	GetStyle() *int32
	SetWidth(v int32) *TransferPortraitStyleRequest
	GetWidth() *int32
}

type TransferPortraitStyleRequest struct {
	// example:
	//
	// 500
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// example:
	//
	// WWW
	ImageUrl *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	// example:
	//
	// 4
	Numbers *int32 `json:"numbers,omitempty" xml:"numbers,omitempty"`
	// example:
	//
	// 1
	RedrawAmplitude *int32 `json:"redrawAmplitude,omitempty" xml:"redrawAmplitude,omitempty"`
	// example:
	//
	// 1
	Style *int32 `json:"style,omitempty" xml:"style,omitempty"`
	// example:
	//
	// 500
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s TransferPortraitStyleRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferPortraitStyleRequest) GoString() string {
	return s.String()
}

func (s *TransferPortraitStyleRequest) GetHeight() *int32 {
	return s.Height
}

func (s *TransferPortraitStyleRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *TransferPortraitStyleRequest) GetNumbers() *int32 {
	return s.Numbers
}

func (s *TransferPortraitStyleRequest) GetRedrawAmplitude() *int32 {
	return s.RedrawAmplitude
}

func (s *TransferPortraitStyleRequest) GetStyle() *int32 {
	return s.Style
}

func (s *TransferPortraitStyleRequest) GetWidth() *int32 {
	return s.Width
}

func (s *TransferPortraitStyleRequest) SetHeight(v int32) *TransferPortraitStyleRequest {
	s.Height = &v
	return s
}

func (s *TransferPortraitStyleRequest) SetImageUrl(v string) *TransferPortraitStyleRequest {
	s.ImageUrl = &v
	return s
}

func (s *TransferPortraitStyleRequest) SetNumbers(v int32) *TransferPortraitStyleRequest {
	s.Numbers = &v
	return s
}

func (s *TransferPortraitStyleRequest) SetRedrawAmplitude(v int32) *TransferPortraitStyleRequest {
	s.RedrawAmplitude = &v
	return s
}

func (s *TransferPortraitStyleRequest) SetStyle(v int32) *TransferPortraitStyleRequest {
	s.Style = &v
	return s
}

func (s *TransferPortraitStyleRequest) SetWidth(v int32) *TransferPortraitStyleRequest {
	s.Width = &v
	return s
}

func (s *TransferPortraitStyleRequest) Validate() error {
	return dara.Validate(s)
}
