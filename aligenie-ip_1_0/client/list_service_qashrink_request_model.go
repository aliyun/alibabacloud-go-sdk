// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceQAShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *ListServiceQAShrinkRequest
	GetActive() *bool
	SetHotelId(v string) *ListServiceQAShrinkRequest
	GetHotelId() *string
	SetKeyword(v string) *ListServiceQAShrinkRequest
	GetKeyword() *string
	SetPageShrink(v string) *ListServiceQAShrinkRequest
	GetPageShrink() *string
}

type ListServiceQAShrinkRequest struct {
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// ***
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageShrink *string `json:"Page,omitempty" xml:"Page,omitempty"`
}

func (s ListServiceQAShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceQAShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListServiceQAShrinkRequest) GetActive() *bool {
	return s.Active
}

func (s *ListServiceQAShrinkRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ListServiceQAShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListServiceQAShrinkRequest) GetPageShrink() *string {
	return s.PageShrink
}

func (s *ListServiceQAShrinkRequest) SetActive(v bool) *ListServiceQAShrinkRequest {
	s.Active = &v
	return s
}

func (s *ListServiceQAShrinkRequest) SetHotelId(v string) *ListServiceQAShrinkRequest {
	s.HotelId = &v
	return s
}

func (s *ListServiceQAShrinkRequest) SetKeyword(v string) *ListServiceQAShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListServiceQAShrinkRequest) SetPageShrink(v string) *ListServiceQAShrinkRequest {
	s.PageShrink = &v
	return s
}

func (s *ListServiceQAShrinkRequest) Validate() error {
	return dara.Validate(s)
}
