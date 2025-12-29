// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelSceneItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListHotelSceneItemsResponseBody
	GetCode() *int32
	SetMessage(v string) *ListHotelSceneItemsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHotelSceneItemsResponseBody
	GetRequestId() *string
	SetResult(v *ListHotelSceneItemsResponseBodyResult) *ListHotelSceneItemsResponseBody
	GetResult() *ListHotelSceneItemsResponseBodyResult
}

type ListHotelSceneItemsResponseBody struct {
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
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *ListHotelSceneItemsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListHotelSceneItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListHotelSceneItemsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHotelSceneItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotelSceneItemsResponseBody) GetResult() *ListHotelSceneItemsResponseBodyResult {
	return s.Result
}

func (s *ListHotelSceneItemsResponseBody) SetCode(v int32) *ListHotelSceneItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelSceneItemsResponseBody) SetMessage(v string) *ListHotelSceneItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelSceneItemsResponseBody) SetRequestId(v string) *ListHotelSceneItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelSceneItemsResponseBody) SetResult(v *ListHotelSceneItemsResponseBodyResult) *ListHotelSceneItemsResponseBody {
	s.Result = v
	return s
}

func (s *ListHotelSceneItemsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHotelSceneItemsResponseBodyResult struct {
	Page          *ListHotelSceneItemsResponseBodyResultPage            `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	SceneItemList []*ListHotelSceneItemsResponseBodyResultSceneItemList `json:"SceneItemList,omitempty" xml:"SceneItemList,omitempty" type:"Repeated"`
}

func (s ListHotelSceneItemsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsResponseBodyResult) GetPage() *ListHotelSceneItemsResponseBodyResultPage {
	return s.Page
}

func (s *ListHotelSceneItemsResponseBodyResult) GetSceneItemList() []*ListHotelSceneItemsResponseBodyResultSceneItemList {
	return s.SceneItemList
}

func (s *ListHotelSceneItemsResponseBodyResult) SetPage(v *ListHotelSceneItemsResponseBodyResultPage) *ListHotelSceneItemsResponseBodyResult {
	s.Page = v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResult) SetSceneItemList(v []*ListHotelSceneItemsResponseBodyResultSceneItemList) *ListHotelSceneItemsResponseBodyResult {
	s.SceneItemList = v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResult) Validate() error {
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

type ListHotelSceneItemsResponseBodyResultPage struct {
	// example:
	//
	// False
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

func (s ListHotelSceneItemsResponseBodyResultPage) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemsResponseBodyResultPage) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsResponseBodyResultPage) GetHasNext() *bool {
	return s.HasNext
}

func (s *ListHotelSceneItemsResponseBodyResultPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHotelSceneItemsResponseBodyResultPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHotelSceneItemsResponseBodyResultPage) GetTotal() *int32 {
	return s.Total
}

func (s *ListHotelSceneItemsResponseBodyResultPage) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListHotelSceneItemsResponseBodyResultPage) SetHasNext(v bool) *ListHotelSceneItemsResponseBodyResultPage {
	s.HasNext = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultPage) SetPageNumber(v int32) *ListHotelSceneItemsResponseBodyResultPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultPage) SetPageSize(v int32) *ListHotelSceneItemsResponseBodyResultPage {
	s.PageSize = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultPage) SetTotal(v int32) *ListHotelSceneItemsResponseBodyResultPage {
	s.Total = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultPage) SetTotalPage(v int32) *ListHotelSceneItemsResponseBodyResultPage {
	s.TotalPage = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultPage) Validate() error {
	return dara.Validate(s)
}

type ListHotelSceneItemsResponseBodyResultSceneItemList struct {
	BeyondLimitReply *string `json:"BeyondLimitReply,omitempty" xml:"BeyondLimitReply,omitempty"`
	// example:
	//
	// 客用品类
	Category       *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DeliveryMethod *string `json:"DeliveryMethod,omitempty" xml:"DeliveryMethod,omitempty"`
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/jiudianmianban_fuwushangpintu/wupin/keyongpinlei/mianqian.png
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// id
	//
	// example:
	//
	// 10336
	Id          *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	LimitNumber *int32 `json:"LimitNumber,omitempty" xml:"LimitNumber,omitempty"`
	LimitSwitch *int32 `json:"LimitSwitch,omitempty" xml:"LimitSwitch,omitempty"`
	// example:
	//
	// 棉签
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PaymentMethod *string `json:"PaymentMethod,omitempty" xml:"PaymentMethod,omitempty"`
	// example:
	//
	// 160
	Price     *int64  `json:"Price,omitempty" xml:"Price,omitempty"`
	RobotName *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
	// example:
	//
	// 已添加
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// GOODS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1666163226
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListHotelSceneItemsResponseBodyResultSceneItemList) String() string {
	return dara.Prettify(s)
}

func (s ListHotelSceneItemsResponseBodyResultSceneItemList) GoString() string {
	return s.String()
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) GetBeyondLimitReply() *string {
	return s.BeyondLimitReply
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) GetCategory() *string {
	return s.Category
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) GetDeliveryMethod() *string {
	return s.DeliveryMethod
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) GetIcon() *string {
	return s.Icon
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) GetId() *int64 {
	return s.Id
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) GetLimitNumber() *int32 {
	return s.LimitNumber
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) GetLimitSwitch() *int32 {
	return s.LimitSwitch
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) GetName() *string {
	return s.Name
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) GetPaymentMethod() *string {
	return s.PaymentMethod
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) GetPrice() *int64 {
	return s.Price
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) GetRobotName() *string {
	return s.RobotName
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) GetStatus() *string {
	return s.Status
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) GetType() *string {
	return s.Type
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetBeyondLimitReply(v string) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.BeyondLimitReply = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetCategory(v string) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.Category = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetDeliveryMethod(v string) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.DeliveryMethod = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetIcon(v string) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.Icon = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetId(v int64) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.Id = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetLimitNumber(v int32) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.LimitNumber = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetLimitSwitch(v int32) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.LimitSwitch = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetName(v string) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.Name = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetPaymentMethod(v string) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.PaymentMethod = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetPrice(v int64) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.Price = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetRobotName(v string) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.RobotName = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetStatus(v string) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.Status = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetType(v string) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.Type = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) SetUpdateTime(v int64) *ListHotelSceneItemsResponseBodyResultSceneItemList {
	s.UpdateTime = &v
	return s
}

func (s *ListHotelSceneItemsResponseBodyResultSceneItemList) Validate() error {
	return dara.Validate(s)
}
