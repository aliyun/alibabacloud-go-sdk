// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomQAShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *ListCustomQAShrinkRequest
	GetHotelId() *string
	SetKeyword(v string) *ListCustomQAShrinkRequest
	GetKeyword() *string
	SetPageShrink(v string) *ListCustomQAShrinkRequest
	GetPageShrink() *string
}

type ListCustomQAShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 520a0c0***5eb
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// ***
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// This parameter is required.
	PageShrink *string `json:"Page,omitempty" xml:"Page,omitempty"`
}

func (s ListCustomQAShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomQAShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCustomQAShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ListCustomQAShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListCustomQAShrinkRequest) GetPageShrink() *string {
	return s.PageShrink
}

func (s *ListCustomQAShrinkRequest) SetHotelId(v string) *ListCustomQAShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ListCustomQAShrinkRequest) SetKeyword(v string) *ListCustomQAShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListCustomQAShrinkRequest) SetPageShrink(v string) *ListCustomQAShrinkRequest {
	s.PageShrink = &v
	return s
}

func (s *ListCustomQAShrinkRequest) Validate() error {
	return dara.Validate(s)
}
