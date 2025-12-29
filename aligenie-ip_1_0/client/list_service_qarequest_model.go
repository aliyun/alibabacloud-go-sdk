// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceQARequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *ListServiceQARequest
	GetActive() *bool
	SetHotelId(v string) *ListServiceQARequest
	GetHotelId() *string
	SetKeyword(v string) *ListServiceQARequest
	GetKeyword() *string
	SetPage(v *ListServiceQARequestPage) *ListServiceQARequest
	GetPage() *ListServiceQARequestPage
}

type ListServiceQARequest struct {
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
	Keyword *string                   `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Page    *ListServiceQARequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
}

func (s ListServiceQARequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceQARequest) GoString() string {
	return s.String()
}

func (s *ListServiceQARequest) GetActive() *bool {
	return s.Active
}

func (s *ListServiceQARequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ListServiceQARequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListServiceQARequest) GetPage() *ListServiceQARequestPage {
	return s.Page
}

func (s *ListServiceQARequest) SetActive(v bool) *ListServiceQARequest {
	s.Active = &v
	return s
}

func (s *ListServiceQARequest) SetHotelId(v string) *ListServiceQARequest {
	s.HotelId = &v
	return s
}

func (s *ListServiceQARequest) SetKeyword(v string) *ListServiceQARequest {
	s.Keyword = &v
	return s
}

func (s *ListServiceQARequest) SetPage(v *ListServiceQARequestPage) *ListServiceQARequest {
	s.Page = v
	return s
}

func (s *ListServiceQARequest) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListServiceQARequestPage struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListServiceQARequestPage) String() string {
	return dara.Prettify(s)
}

func (s ListServiceQARequestPage) GoString() string {
	return s.String()
}

func (s *ListServiceQARequestPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListServiceQARequestPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListServiceQARequestPage) SetPageNumber(v int32) *ListServiceQARequestPage {
	s.PageNumber = &v
	return s
}

func (s *ListServiceQARequestPage) SetPageSize(v int32) *ListServiceQARequestPage {
	s.PageSize = &v
	return s
}

func (s *ListServiceQARequestPage) Validate() error {
	return dara.Validate(s)
}
