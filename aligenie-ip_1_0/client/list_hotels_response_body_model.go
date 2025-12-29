// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListHotelsResponseBody
	GetCode() *int32
	SetMessage(v string) *ListHotelsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHotelsResponseBody
	GetRequestId() *string
	SetResult(v *ListHotelsResponseBodyResult) *ListHotelsResponseBody
	GetResult() *ListHotelsResponseBodyResult
}

type ListHotelsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListHotelsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListHotelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListHotelsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHotelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotelsResponseBody) GetResult() *ListHotelsResponseBodyResult {
	return s.Result
}

func (s *ListHotelsResponseBody) SetCode(v int32) *ListHotelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelsResponseBody) SetMessage(v string) *ListHotelsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelsResponseBody) SetRequestId(v string) *ListHotelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelsResponseBody) SetResult(v *ListHotelsResponseBodyResult) *ListHotelsResponseBody {
	s.Result = v
	return s
}

func (s *ListHotelsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHotelsResponseBodyResult struct {
	HotelInfoList []*ListHotelsResponseBodyResultHotelInfoList `json:"HotelInfoList,omitempty" xml:"HotelInfoList,omitempty" type:"Repeated"`
	Page          *ListHotelsResponseBodyResultPage            `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
}

func (s ListHotelsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListHotelsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelsResponseBodyResult) GetHotelInfoList() []*ListHotelsResponseBodyResultHotelInfoList {
	return s.HotelInfoList
}

func (s *ListHotelsResponseBodyResult) GetPage() *ListHotelsResponseBodyResultPage {
	return s.Page
}

func (s *ListHotelsResponseBodyResult) SetHotelInfoList(v []*ListHotelsResponseBodyResultHotelInfoList) *ListHotelsResponseBodyResult {
	s.HotelInfoList = v
	return s
}

func (s *ListHotelsResponseBodyResult) SetPage(v *ListHotelsResponseBodyResultPage) *ListHotelsResponseBodyResult {
	s.Page = v
	return s
}

func (s *ListHotelsResponseBodyResult) Validate() error {
	if s.HotelInfoList != nil {
		for _, item := range s.HotelInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHotelsResponseBodyResultHotelInfoList struct {
	AccountNames []*string `json:"AccountNames,omitempty" xml:"AccountNames,omitempty" type:"Repeated"`
	// example:
	//
	// 1654568802000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 酒店地址
	HotelAddress *string `json:"HotelAddress,omitempty" xml:"HotelAddress,omitempty"`
	// example:
	//
	// 73ab1b03018d4da69b5bef17095f569b
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// 酒店名称
	HotelName *string `json:"HotelName,omitempty" xml:"HotelName,omitempty"`
	// example:
	//
	// 酒店
	IndustryType *string `json:"IndustryType,omitempty" xml:"IndustryType,omitempty"`
	// example:
	//
	// 13312340987
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// 测试产品
	RelatedProductName *string `json:"RelatedProductName,omitempty" xml:"RelatedProductName,omitempty"`
	// example:
	//
	// 12
	RoomNum *int32 `json:"RoomNum,omitempty" xml:"RoomNum,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListHotelsResponseBodyResultHotelInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListHotelsResponseBodyResultHotelInfoList) GoString() string {
	return s.String()
}

func (s *ListHotelsResponseBodyResultHotelInfoList) GetAccountNames() []*string {
	return s.AccountNames
}

func (s *ListHotelsResponseBodyResultHotelInfoList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListHotelsResponseBodyResultHotelInfoList) GetHotelAddress() *string {
	return s.HotelAddress
}

func (s *ListHotelsResponseBodyResultHotelInfoList) GetHotelId() *string {
	return s.HotelId
}

func (s *ListHotelsResponseBodyResultHotelInfoList) GetHotelName() *string {
	return s.HotelName
}

func (s *ListHotelsResponseBodyResultHotelInfoList) GetIndustryType() *string {
	return s.IndustryType
}

func (s *ListHotelsResponseBodyResultHotelInfoList) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ListHotelsResponseBodyResultHotelInfoList) GetRelatedProductName() *string {
	return s.RelatedProductName
}

func (s *ListHotelsResponseBodyResultHotelInfoList) GetRoomNum() *int32 {
	return s.RoomNum
}

func (s *ListHotelsResponseBodyResultHotelInfoList) GetStatus() *int32 {
	return s.Status
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetAccountNames(v []*string) *ListHotelsResponseBodyResultHotelInfoList {
	s.AccountNames = v
	return s
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetCreateTime(v int64) *ListHotelsResponseBodyResultHotelInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetHotelAddress(v string) *ListHotelsResponseBodyResultHotelInfoList {
	s.HotelAddress = &v
	return s
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetHotelId(v string) *ListHotelsResponseBodyResultHotelInfoList {
	s.HotelId = &v
	return s
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetHotelName(v string) *ListHotelsResponseBodyResultHotelInfoList {
	s.HotelName = &v
	return s
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetIndustryType(v string) *ListHotelsResponseBodyResultHotelInfoList {
	s.IndustryType = &v
	return s
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetPhoneNumber(v string) *ListHotelsResponseBodyResultHotelInfoList {
	s.PhoneNumber = &v
	return s
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetRelatedProductName(v string) *ListHotelsResponseBodyResultHotelInfoList {
	s.RelatedProductName = &v
	return s
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetRoomNum(v int32) *ListHotelsResponseBodyResultHotelInfoList {
	s.RoomNum = &v
	return s
}

func (s *ListHotelsResponseBodyResultHotelInfoList) SetStatus(v int32) *ListHotelsResponseBodyResultHotelInfoList {
	s.Status = &v
	return s
}

func (s *ListHotelsResponseBodyResultHotelInfoList) Validate() error {
	return dara.Validate(s)
}

type ListHotelsResponseBodyResultPage struct {
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

func (s ListHotelsResponseBodyResultPage) String() string {
	return dara.Prettify(s)
}

func (s ListHotelsResponseBodyResultPage) GoString() string {
	return s.String()
}

func (s *ListHotelsResponseBodyResultPage) GetHasNext() *bool {
	return s.HasNext
}

func (s *ListHotelsResponseBodyResultPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHotelsResponseBodyResultPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHotelsResponseBodyResultPage) GetTotal() *int32 {
	return s.Total
}

func (s *ListHotelsResponseBodyResultPage) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListHotelsResponseBodyResultPage) SetHasNext(v bool) *ListHotelsResponseBodyResultPage {
	s.HasNext = &v
	return s
}

func (s *ListHotelsResponseBodyResultPage) SetPageNumber(v int32) *ListHotelsResponseBodyResultPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelsResponseBodyResultPage) SetPageSize(v int32) *ListHotelsResponseBodyResultPage {
	s.PageSize = &v
	return s
}

func (s *ListHotelsResponseBodyResultPage) SetTotal(v int32) *ListHotelsResponseBodyResultPage {
	s.Total = &v
	return s
}

func (s *ListHotelsResponseBodyResultPage) SetTotalPage(v int32) *ListHotelsResponseBodyResultPage {
	s.TotalPage = &v
	return s
}

func (s *ListHotelsResponseBodyResultPage) Validate() error {
	return dara.Validate(s)
}
