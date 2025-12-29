// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelRequestShrink(v string) *ListHotelsShrinkRequest
	GetHotelRequestShrink() *string
	SetPageShrink(v string) *ListHotelsShrinkRequest
	GetPageShrink() *string
	SetStatus(v int32) *ListHotelsShrinkRequest
	GetStatus() *int32
}

type ListHotelsShrinkRequest struct {
	// if can be null:
	// true
	HotelRequestShrink *string `json:"HotelRequest,omitempty" xml:"HotelRequest,omitempty"`
	// This parameter is required.
	PageShrink *string `json:"Page,omitempty" xml:"Page,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListHotelsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotelsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotelsShrinkRequest) GetHotelRequestShrink() *string {
	return s.HotelRequestShrink
}

func (s *ListHotelsShrinkRequest) GetPageShrink() *string {
	return s.PageShrink
}

func (s *ListHotelsShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListHotelsShrinkRequest) SetHotelRequestShrink(v string) *ListHotelsShrinkRequest {
	s.HotelRequestShrink = &v
	return s
}

func (s *ListHotelsShrinkRequest) SetPageShrink(v string) *ListHotelsShrinkRequest {
	s.PageShrink = &v
	return s
}

func (s *ListHotelsShrinkRequest) SetStatus(v int32) *ListHotelsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListHotelsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
