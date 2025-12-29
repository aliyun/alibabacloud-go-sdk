// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomQARequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *ListCustomQARequest
	GetHotelId() *string
	SetKeyword(v string) *ListCustomQARequest
	GetKeyword() *string
	SetPage(v *ListCustomQARequestPage) *ListCustomQARequest
	GetPage() *ListCustomQARequestPage
}

type ListCustomQARequest struct {
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
	Page *ListCustomQARequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
}

func (s ListCustomQARequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomQARequest) GoString() string {
	return s.String()
}

func (s *ListCustomQARequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ListCustomQARequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListCustomQARequest) GetPage() *ListCustomQARequestPage {
	return s.Page
}

func (s *ListCustomQARequest) SetHotelId(v string) *ListCustomQARequest {
	s.HotelId = &v
	return s
}

func (s *ListCustomQARequest) SetKeyword(v string) *ListCustomQARequest {
	s.Keyword = &v
	return s
}

func (s *ListCustomQARequest) SetPage(v *ListCustomQARequestPage) *ListCustomQARequest {
	s.Page = v
	return s
}

func (s *ListCustomQARequest) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCustomQARequestPage struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCustomQARequestPage) String() string {
	return dara.Prettify(s)
}

func (s ListCustomQARequestPage) GoString() string {
	return s.String()
}

func (s *ListCustomQARequestPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomQARequestPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomQARequestPage) SetPageNumber(v int32) *ListCustomQARequestPage {
	s.PageNumber = &v
	return s
}

func (s *ListCustomQARequestPage) SetPageSize(v int32) *ListCustomQARequestPage {
	s.PageSize = &v
	return s
}

func (s *ListCustomQARequestPage) Validate() error {
	return dara.Validate(s)
}
