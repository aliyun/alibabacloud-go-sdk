// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelSceneBookItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *ListHotelSceneBookItemsRequest
	GetHotelId() *string
	SetPage(v *ListHotelSceneBookItemsRequestPage) *ListHotelSceneBookItemsRequest
	GetPage() *ListHotelSceneBookItemsRequestPage
	SetType(v string) *ListHotelSceneBookItemsRequest
	GetType() *string
}

type ListHotelSceneBookItemsRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// This parameter is required.
	Page *ListHotelSceneBookItemsRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// FOOD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListHotelSceneBookItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneBookItemsRequest) GoString() string {
	return s.String()
}

func (s *ListHotelSceneBookItemsRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ListHotelSceneBookItemsRequest) GetPage() *ListHotelSceneBookItemsRequestPage {
	return s.Page
}

func (s *ListHotelSceneBookItemsRequest) GetType() *string {
	return s.Type
}

func (s *ListHotelSceneBookItemsRequest) SetHotelId(v string) *ListHotelSceneBookItemsRequest {
	s.HotelId = &v
	return s
}

func (s *ListHotelSceneBookItemsRequest) SetPage(v *ListHotelSceneBookItemsRequestPage) *ListHotelSceneBookItemsRequest {
	s.Page = v
	return s
}

func (s *ListHotelSceneBookItemsRequest) SetType(v string) *ListHotelSceneBookItemsRequest {
	s.Type = &v
	return s
}

func (s *ListHotelSceneBookItemsRequest) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHotelSceneBookItemsRequestPage struct {
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

func (s ListHotelSceneBookItemsRequestPage) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneBookItemsRequestPage) GoString() string {
	return s.String()
}

func (s *ListHotelSceneBookItemsRequestPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHotelSceneBookItemsRequestPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHotelSceneBookItemsRequestPage) SetPageNumber(v int32) *ListHotelSceneBookItemsRequestPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelSceneBookItemsRequestPage) SetPageSize(v int32) *ListHotelSceneBookItemsRequestPage {
	s.PageSize = &v
	return s
}

func (s *ListHotelSceneBookItemsRequestPage) Validate() error {
	return dara.Validate(s)
}
