// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelSceneItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelId(v string) *ListHotelSceneItemsRequest
	GetHotelId() *string
	SetListHotelSceneReq(v *ListHotelSceneItemsRequestListHotelSceneReq) *ListHotelSceneItemsRequest
	GetListHotelSceneReq() *ListHotelSceneItemsRequestListHotelSceneReq
}

type ListHotelSceneItemsRequest struct {
	// hotelID
	//
	// This parameter is required.
	//
	// example:
	//
	// 80d84ea8ed9e422fbad52715c8fc56f1
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// ListHotelSceneReq
	//
	// This parameter is required.
	ListHotelSceneReq *ListHotelSceneItemsRequestListHotelSceneReq `json:"ListHotelSceneReq,omitempty" xml:"ListHotelSceneReq,omitempty" type:"Struct"`
}

func (s ListHotelSceneItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemsRequest) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ListHotelSceneItemsRequest) GetListHotelSceneReq() *ListHotelSceneItemsRequestListHotelSceneReq {
	return s.ListHotelSceneReq
}

func (s *ListHotelSceneItemsRequest) SetHotelId(v string) *ListHotelSceneItemsRequest {
	s.HotelId = &v
	return s
}

func (s *ListHotelSceneItemsRequest) SetListHotelSceneReq(v *ListHotelSceneItemsRequestListHotelSceneReq) *ListHotelSceneItemsRequest {
	s.ListHotelSceneReq = v
	return s
}

func (s *ListHotelSceneItemsRequest) Validate() error {
	if s.ListHotelSceneReq != nil {
		if err := s.ListHotelSceneReq.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHotelSceneItemsRequestListHotelSceneReq struct {
	// example:
	//
	// 客用品类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 棉签
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// This parameter is required.
	Page *ListHotelSceneItemsRequestListHotelSceneReqPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// 已添加
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// GOODS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListHotelSceneItemsRequestListHotelSceneReq) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemsRequestListHotelSceneReq) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsRequestListHotelSceneReq) GetCategory() *string {
	return s.Category
}

func (s *ListHotelSceneItemsRequestListHotelSceneReq) GetKeywords() *string {
	return s.Keywords
}

func (s *ListHotelSceneItemsRequestListHotelSceneReq) GetPage() *ListHotelSceneItemsRequestListHotelSceneReqPage {
	return s.Page
}

func (s *ListHotelSceneItemsRequestListHotelSceneReq) GetStatus() *string {
	return s.Status
}

func (s *ListHotelSceneItemsRequestListHotelSceneReq) GetType() *string {
	return s.Type
}

func (s *ListHotelSceneItemsRequestListHotelSceneReq) SetCategory(v string) *ListHotelSceneItemsRequestListHotelSceneReq {
	s.Category = &v
	return s
}

func (s *ListHotelSceneItemsRequestListHotelSceneReq) SetKeywords(v string) *ListHotelSceneItemsRequestListHotelSceneReq {
	s.Keywords = &v
	return s
}

func (s *ListHotelSceneItemsRequestListHotelSceneReq) SetPage(v *ListHotelSceneItemsRequestListHotelSceneReqPage) *ListHotelSceneItemsRequestListHotelSceneReq {
	s.Page = v
	return s
}

func (s *ListHotelSceneItemsRequestListHotelSceneReq) SetStatus(v string) *ListHotelSceneItemsRequestListHotelSceneReq {
	s.Status = &v
	return s
}

func (s *ListHotelSceneItemsRequestListHotelSceneReq) SetType(v string) *ListHotelSceneItemsRequestListHotelSceneReq {
	s.Type = &v
	return s
}

func (s *ListHotelSceneItemsRequestListHotelSceneReq) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHotelSceneItemsRequestListHotelSceneReqPage struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHotelSceneItemsRequestListHotelSceneReqPage) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemsRequestListHotelSceneReqPage) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsRequestListHotelSceneReqPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHotelSceneItemsRequestListHotelSceneReqPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHotelSceneItemsRequestListHotelSceneReqPage) SetPageNumber(v int32) *ListHotelSceneItemsRequestListHotelSceneReqPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelSceneItemsRequestListHotelSceneReqPage) SetPageSize(v int32) *ListHotelSceneItemsRequestListHotelSceneReqPage {
	s.PageSize = &v
	return s
}

func (s *ListHotelSceneItemsRequestListHotelSceneReqPage) Validate() error {
	return dara.Validate(s)
}
