// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelSceneBookItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListHotelSceneBookItemsResponseBody
	GetCode() *int32
	SetMessage(v string) *ListHotelSceneBookItemsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHotelSceneBookItemsResponseBody
	GetRequestId() *string
	SetResult(v *ListHotelSceneBookItemsResponseBodyResult) *ListHotelSceneBookItemsResponseBody
	GetResult() *ListHotelSceneBookItemsResponseBodyResult
}

type ListHotelSceneBookItemsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListHotelSceneBookItemsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListHotelSceneBookItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneBookItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelSceneBookItemsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListHotelSceneBookItemsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHotelSceneBookItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotelSceneBookItemsResponseBody) GetResult() *ListHotelSceneBookItemsResponseBodyResult {
	return s.Result
}

func (s *ListHotelSceneBookItemsResponseBody) SetCode(v int32) *ListHotelSceneBookItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBody) SetMessage(v string) *ListHotelSceneBookItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBody) SetRequestId(v string) *ListHotelSceneBookItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBody) SetResult(v *ListHotelSceneBookItemsResponseBodyResult) *ListHotelSceneBookItemsResponseBody {
	s.Result = v
	return s
}

func (s *ListHotelSceneBookItemsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHotelSceneBookItemsResponseBodyResult struct {
	Page          *ListHotelSceneBookItemsResponseBodyResultPage            `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	SceneItemList []*ListHotelSceneBookItemsResponseBodyResultSceneItemList `json:"SceneItemList,omitempty" xml:"SceneItemList,omitempty" type:"Repeated"`
}

func (s ListHotelSceneBookItemsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneBookItemsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelSceneBookItemsResponseBodyResult) GetPage() *ListHotelSceneBookItemsResponseBodyResultPage {
	return s.Page
}

func (s *ListHotelSceneBookItemsResponseBodyResult) GetSceneItemList() []*ListHotelSceneBookItemsResponseBodyResultSceneItemList {
	return s.SceneItemList
}

func (s *ListHotelSceneBookItemsResponseBodyResult) SetPage(v *ListHotelSceneBookItemsResponseBodyResultPage) *ListHotelSceneBookItemsResponseBodyResult {
	s.Page = v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResult) SetSceneItemList(v []*ListHotelSceneBookItemsResponseBodyResultSceneItemList) *ListHotelSceneBookItemsResponseBodyResult {
	s.SceneItemList = v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResult) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	if s.SceneItemList != nil {
		for _, item := range s.SceneItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHotelSceneBookItemsResponseBodyResultPage struct {
	// example:
	//
	// True
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
	// 23
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// 3
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListHotelSceneBookItemsResponseBodyResultPage) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneBookItemsResponseBodyResultPage) GoString() string {
	return s.String()
}

func (s *ListHotelSceneBookItemsResponseBodyResultPage) GetHasNext() *bool {
	return s.HasNext
}

func (s *ListHotelSceneBookItemsResponseBodyResultPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHotelSceneBookItemsResponseBodyResultPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHotelSceneBookItemsResponseBodyResultPage) GetTotal() *int32 {
	return s.Total
}

func (s *ListHotelSceneBookItemsResponseBodyResultPage) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListHotelSceneBookItemsResponseBodyResultPage) SetHasNext(v bool) *ListHotelSceneBookItemsResponseBodyResultPage {
	s.HasNext = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultPage) SetPageNumber(v int32) *ListHotelSceneBookItemsResponseBodyResultPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultPage) SetPageSize(v int32) *ListHotelSceneBookItemsResponseBodyResultPage {
	s.PageSize = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultPage) SetTotal(v int32) *ListHotelSceneBookItemsResponseBodyResultPage {
	s.Total = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultPage) SetTotalPage(v int32) *ListHotelSceneBookItemsResponseBodyResultPage {
	s.TotalPage = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultPage) Validate() error {
	return dara.Validate(s)
}

type ListHotelSceneBookItemsResponseBodyResultSceneItemList struct {
	// example:
	//
	// https://ailabs.alibabausercontent.com/platform/28d7a91e3c05db3855725fc39e0387e7/welcome_audios/aa918294b6ca3aa115c51135bf9b80cb/l9f996sq.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 11824
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 青椒肉丝
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1850
	Price *int64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// example:
	//
	// 已添加
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// FOOD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1666161803
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListHotelSceneBookItemsResponseBodyResultSceneItemList) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneBookItemsResponseBodyResultSceneItemList) GoString() string {
	return s.String()
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) GetIcon() *string {
	return s.Icon
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) GetId() *int64 {
	return s.Id
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) GetName() *string {
	return s.Name
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) GetPrice() *int64 {
	return s.Price
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) GetStatus() *string {
	return s.Status
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) GetType() *string {
	return s.Type
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) SetIcon(v string) *ListHotelSceneBookItemsResponseBodyResultSceneItemList {
	s.Icon = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) SetId(v int64) *ListHotelSceneBookItemsResponseBodyResultSceneItemList {
	s.Id = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) SetName(v string) *ListHotelSceneBookItemsResponseBodyResultSceneItemList {
	s.Name = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) SetPrice(v int64) *ListHotelSceneBookItemsResponseBodyResultSceneItemList {
	s.Price = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) SetStatus(v string) *ListHotelSceneBookItemsResponseBodyResultSceneItemList {
	s.Status = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) SetType(v string) *ListHotelSceneBookItemsResponseBodyResultSceneItemList {
	s.Type = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) SetUpdateTime(v int64) *ListHotelSceneBookItemsResponseBodyResultSceneItemList {
	s.UpdateTime = &v
	return s
}

func (s *ListHotelSceneBookItemsResponseBodyResultSceneItemList) Validate() error {
	return dara.Validate(s)
}
