// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelRequest(v *ListHotelsRequestHotelRequest) *ListHotelsRequest
	GetHotelRequest() *ListHotelsRequestHotelRequest
	SetPage(v *ListHotelsRequestPage) *ListHotelsRequest
	GetPage() *ListHotelsRequestPage
	SetStatus(v int32) *ListHotelsRequest
	GetStatus() *int32
}

type ListHotelsRequest struct {
	// if can be null:
	// true
	HotelRequest *ListHotelsRequestHotelRequest `json:"HotelRequest,omitempty" xml:"HotelRequest,omitempty" type:"Struct"`
	// This parameter is required.
	Page *ListHotelsRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListHotelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotelsRequest) GoString() string {
	return s.String()
}

func (s *ListHotelsRequest) GetHotelRequest() *ListHotelsRequestHotelRequest {
	return s.HotelRequest
}

func (s *ListHotelsRequest) GetPage() *ListHotelsRequestPage {
	return s.Page
}

func (s *ListHotelsRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListHotelsRequest) SetHotelRequest(v *ListHotelsRequestHotelRequest) *ListHotelsRequest {
	s.HotelRequest = v
	return s
}

func (s *ListHotelsRequest) SetPage(v *ListHotelsRequestPage) *ListHotelsRequest {
	s.Page = v
	return s
}

func (s *ListHotelsRequest) SetStatus(v int32) *ListHotelsRequest {
	s.Status = &v
	return s
}

func (s *ListHotelsRequest) Validate() error {
	if s.HotelRequest != nil {
		if err := s.HotelRequest.Validate(); err != nil {
			return err
		}
	}
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHotelsRequestHotelRequest struct {
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
}

func (s ListHotelsRequestHotelRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotelsRequestHotelRequest) GoString() string {
	return s.String()
}

func (s *ListHotelsRequestHotelRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ListHotelsRequestHotelRequest) SetHotelId(v string) *ListHotelsRequestHotelRequest {
	s.HotelId = &v
	return s
}

func (s *ListHotelsRequestHotelRequest) Validate() error {
	return dara.Validate(s)
}

type ListHotelsRequestPage struct {
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

func (s ListHotelsRequestPage) String() string {
	return dara.Prettify(s)
}

func (s ListHotelsRequestPage) GoString() string {
	return s.String()
}

func (s *ListHotelsRequestPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHotelsRequestPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHotelsRequestPage) SetPageNumber(v int32) *ListHotelsRequestPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelsRequestPage) SetPageSize(v int32) *ListHotelsRequestPage {
	s.PageSize = &v
	return s
}

func (s *ListHotelsRequestPage) Validate() error {
	return dara.Validate(s)
}
