// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelSceneItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListHotelSceneItemResponseBody
	GetCode() *int32
	SetMessage(v string) *ListHotelSceneItemResponseBody
	GetMessage() *string
	SetPage(v *ListHotelSceneItemResponseBodyPage) *ListHotelSceneItemResponseBody
	GetPage() *ListHotelSceneItemResponseBodyPage
	SetRequestId(v string) *ListHotelSceneItemResponseBody
	GetRequestId() *string
	SetResult(v *ListHotelSceneItemResponseBodyResult) *ListHotelSceneItemResponseBody
	GetResult() *ListHotelSceneItemResponseBodyResult
}

type ListHotelSceneItemResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Page    *ListHotelSceneItemResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// CEADB586-51CB-1B6B-95BD-AB85A7A08E97
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListHotelSceneItemResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListHotelSceneItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListHotelSceneItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHotelSceneItemResponseBody) GetPage() *ListHotelSceneItemResponseBodyPage {
	return s.Page
}

func (s *ListHotelSceneItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotelSceneItemResponseBody) GetResult() *ListHotelSceneItemResponseBodyResult {
	return s.Result
}

func (s *ListHotelSceneItemResponseBody) SetCode(v int32) *ListHotelSceneItemResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelSceneItemResponseBody) SetMessage(v string) *ListHotelSceneItemResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelSceneItemResponseBody) SetPage(v *ListHotelSceneItemResponseBodyPage) *ListHotelSceneItemResponseBody {
	s.Page = v
	return s
}

func (s *ListHotelSceneItemResponseBody) SetRequestId(v string) *ListHotelSceneItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelSceneItemResponseBody) SetResult(v *ListHotelSceneItemResponseBodyResult) *ListHotelSceneItemResponseBody {
	s.Result = v
	return s
}

func (s *ListHotelSceneItemResponseBody) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHotelSceneItemResponseBodyPage struct {
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 12
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// 6
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListHotelSceneItemResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemResponseBodyPage) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemResponseBodyPage) GetHasNext() *bool {
	return s.HasNext
}

func (s *ListHotelSceneItemResponseBodyPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHotelSceneItemResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHotelSceneItemResponseBodyPage) GetTotal() *int32 {
	return s.Total
}

func (s *ListHotelSceneItemResponseBodyPage) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListHotelSceneItemResponseBodyPage) SetHasNext(v bool) *ListHotelSceneItemResponseBodyPage {
	s.HasNext = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyPage) SetPageNumber(v int32) *ListHotelSceneItemResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyPage) SetPageSize(v int32) *ListHotelSceneItemResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyPage) SetTotal(v int32) *ListHotelSceneItemResponseBodyPage {
	s.Total = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyPage) SetTotalPage(v int32) *ListHotelSceneItemResponseBodyPage {
	s.TotalPage = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyPage) Validate() error {
	return dara.Validate(s)
}

type ListHotelSceneItemResponseBodyResult struct {
	SecondCategoryList []*ListHotelSceneItemResponseBodyResultSecondCategoryList `json:"SecondCategoryList,omitempty" xml:"SecondCategoryList,omitempty" type:"Repeated"`
}

func (s ListHotelSceneItemResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemResponseBodyResult) GetSecondCategoryList() []*ListHotelSceneItemResponseBodyResultSecondCategoryList {
	return s.SecondCategoryList
}

func (s *ListHotelSceneItemResponseBodyResult) SetSecondCategoryList(v []*ListHotelSceneItemResponseBodyResultSecondCategoryList) *ListHotelSceneItemResponseBodyResult {
	s.SecondCategoryList = v
	return s
}

func (s *ListHotelSceneItemResponseBodyResult) Validate() error {
	if s.SecondCategoryList != nil {
		for _, item := range s.SecondCategoryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHotelSceneItemResponseBodyResultSecondCategoryList struct {
	ItemList []*ListHotelSceneItemResponseBodyResultSecondCategoryListItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
	// example:
	//
	// 客用品类
	SecondCategoryName *string `json:"SecondCategoryName,omitempty" xml:"SecondCategoryName,omitempty"`
}

func (s ListHotelSceneItemResponseBodyResultSecondCategoryList) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemResponseBodyResultSecondCategoryList) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryList) GetItemList() []*ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	return s.ItemList
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryList) GetSecondCategoryName() *string {
	return s.SecondCategoryName
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryList) SetItemList(v []*ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) *ListHotelSceneItemResponseBodyResultSecondCategoryList {
	s.ItemList = v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryList) SetSecondCategoryName(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryList {
	s.SecondCategoryName = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryList) Validate() error {
	if s.ItemList != nil {
		for _, item := range s.ItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHotelSceneItemResponseBodyResultSecondCategoryListItemList struct {
	// example:
	//
	// 客用品类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/jiudianmianban_fuwushangpintu/wupin/keyongpinlei/mianqian.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 152860
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 棉签
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Price        *int64 `json:"Price,omitempty" xml:"Price,omitempty"`
	ResidueLimit *int64 `json:"ResidueLimit,omitempty" xml:"ResidueLimit,omitempty"`
	// example:
	//
	// 已添加
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// GOODS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) GetCategory() *string {
	return s.Category
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) GetIcon() *string {
	return s.Icon
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) GetId() *string {
	return s.Id
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) GetName() *string {
	return s.Name
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) GetPrice() *int64 {
	return s.Price
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) GetResidueLimit() *int64 {
	return s.ResidueLimit
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) GetStatus() *string {
	return s.Status
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) GetType() *string {
	return s.Type
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetCategory(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Category = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetIcon(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Icon = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetId(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Id = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetName(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Name = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetPrice(v int64) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Price = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetResidueLimit(v int64) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.ResidueLimit = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetStatus(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Status = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) SetType(v string) *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList {
	s.Type = &v
	return s
}

func (s *ListHotelSceneItemResponseBodyResultSecondCategoryListItemList) Validate() error {
	return dara.Validate(s)
}
