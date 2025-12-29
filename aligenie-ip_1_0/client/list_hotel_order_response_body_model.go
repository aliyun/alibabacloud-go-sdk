// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListHotelOrderResponseBody
	GetCode() *int32
	SetMessage(v string) *ListHotelOrderResponseBody
	GetMessage() *string
	SetPage(v *ListHotelOrderResponseBodyPage) *ListHotelOrderResponseBody
	GetPage() *ListHotelOrderResponseBodyPage
	SetRequestId(v string) *ListHotelOrderResponseBody
	GetRequestId() *string
	SetResult(v []*ListHotelOrderResponseBodyResult) *ListHotelOrderResponseBody
	GetResult() []*ListHotelOrderResponseBodyResult
}

type ListHotelOrderResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Page    *ListHotelOrderResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// 07F61FDA-606F-10A0-8ED0-C6CE62710A48
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListHotelOrderResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListHotelOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotelOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelOrderResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListHotelOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHotelOrderResponseBody) GetPage() *ListHotelOrderResponseBodyPage {
	return s.Page
}

func (s *ListHotelOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotelOrderResponseBody) GetResult() []*ListHotelOrderResponseBodyResult {
	return s.Result
}

func (s *ListHotelOrderResponseBody) SetCode(v int32) *ListHotelOrderResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelOrderResponseBody) SetMessage(v string) *ListHotelOrderResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelOrderResponseBody) SetPage(v *ListHotelOrderResponseBodyPage) *ListHotelOrderResponseBody {
	s.Page = v
	return s
}

func (s *ListHotelOrderResponseBody) SetRequestId(v string) *ListHotelOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelOrderResponseBody) SetResult(v []*ListHotelOrderResponseBodyResult) *ListHotelOrderResponseBody {
	s.Result = v
	return s
}

func (s *ListHotelOrderResponseBody) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHotelOrderResponseBodyPage struct {
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
	// 21
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// 7
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListHotelOrderResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s ListHotelOrderResponseBodyPage) GoString() string {
	return s.String()
}

func (s *ListHotelOrderResponseBodyPage) GetHasNext() *bool {
	return s.HasNext
}

func (s *ListHotelOrderResponseBodyPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHotelOrderResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHotelOrderResponseBodyPage) GetTotal() *int32 {
	return s.Total
}

func (s *ListHotelOrderResponseBodyPage) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListHotelOrderResponseBodyPage) SetHasNext(v bool) *ListHotelOrderResponseBodyPage {
	s.HasNext = &v
	return s
}

func (s *ListHotelOrderResponseBodyPage) SetPageNumber(v int32) *ListHotelOrderResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *ListHotelOrderResponseBodyPage) SetPageSize(v int32) *ListHotelOrderResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *ListHotelOrderResponseBodyPage) SetTotal(v int32) *ListHotelOrderResponseBodyPage {
	s.Total = &v
	return s
}

func (s *ListHotelOrderResponseBodyPage) SetTotalPage(v int32) *ListHotelOrderResponseBodyPage {
	s.TotalPage = &v
	return s
}

func (s *ListHotelOrderResponseBodyPage) Validate() error {
	return dara.Validate(s)
}

type ListHotelOrderResponseBodyResult struct {
	Amt *int64 `json:"Amt,omitempty" xml:"Amt,omitempty"`
	// example:
	//
	// 21.5
	ApplyAmt         *int64  `json:"ApplyAmt,omitempty" xml:"ApplyAmt,omitempty"`
	DeliveryMethod   *string `json:"DeliveryMethod,omitempty" xml:"DeliveryMethod,omitempty"`
	DeliveryRoomName *string `json:"DeliveryRoomName,omitempty" xml:"DeliveryRoomName,omitempty"`
	DeliveryTime     *int64  `json:"DeliveryTime,omitempty" xml:"DeliveryTime,omitempty"`
	// example:
	//
	// 1659952892000
	GmtCreate *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Icon      *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	ItemId    *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemType  *string `json:"ItemType,omitempty" xml:"ItemType,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 20220808180132000114508555527711
	OrderNo       *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
	OrderStatus   *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	PaymentMethod *string `json:"PaymentMethod,omitempty" xml:"PaymentMethod,omitempty"`
	// example:
	//
	// 12
	Quantity *int64 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// example:
	//
	// 2001
	RoomNo    *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// INIT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SumAmt *int64  `json:"SumAmt,omitempty" xml:"SumAmt,omitempty"`
	// example:
	//
	// REPAIR
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/changjingfenlei/shebeiweixiu.png
	TypeIconUrl *string `json:"TypeIconUrl,omitempty" xml:"TypeIconUrl,omitempty"`
	// example:
	//
	// 设备维修
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s ListHotelOrderResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListHotelOrderResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListHotelOrderResponseBodyResult) GetAmt() *int64 {
	return s.Amt
}

func (s *ListHotelOrderResponseBodyResult) GetApplyAmt() *int64 {
	return s.ApplyAmt
}

func (s *ListHotelOrderResponseBodyResult) GetDeliveryMethod() *string {
	return s.DeliveryMethod
}

func (s *ListHotelOrderResponseBodyResult) GetDeliveryRoomName() *string {
	return s.DeliveryRoomName
}

func (s *ListHotelOrderResponseBodyResult) GetDeliveryTime() *int64 {
	return s.DeliveryTime
}

func (s *ListHotelOrderResponseBodyResult) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListHotelOrderResponseBodyResult) GetIcon() *string {
	return s.Icon
}

func (s *ListHotelOrderResponseBodyResult) GetItemId() *int64 {
	return s.ItemId
}

func (s *ListHotelOrderResponseBodyResult) GetItemType() *string {
	return s.ItemType
}

func (s *ListHotelOrderResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListHotelOrderResponseBodyResult) GetOrderNo() *string {
	return s.OrderNo
}

func (s *ListHotelOrderResponseBodyResult) GetOrderStatus() *string {
	return s.OrderStatus
}

func (s *ListHotelOrderResponseBodyResult) GetPaymentMethod() *string {
	return s.PaymentMethod
}

func (s *ListHotelOrderResponseBodyResult) GetQuantity() *int64 {
	return s.Quantity
}

func (s *ListHotelOrderResponseBodyResult) GetRoomNo() *string {
	return s.RoomNo
}

func (s *ListHotelOrderResponseBodyResult) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListHotelOrderResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListHotelOrderResponseBodyResult) GetSumAmt() *int64 {
	return s.SumAmt
}

func (s *ListHotelOrderResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListHotelOrderResponseBodyResult) GetTypeIconUrl() *string {
	return s.TypeIconUrl
}

func (s *ListHotelOrderResponseBodyResult) GetTypeName() *string {
	return s.TypeName
}

func (s *ListHotelOrderResponseBodyResult) SetAmt(v int64) *ListHotelOrderResponseBodyResult {
	s.Amt = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetApplyAmt(v int64) *ListHotelOrderResponseBodyResult {
	s.ApplyAmt = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetDeliveryMethod(v string) *ListHotelOrderResponseBodyResult {
	s.DeliveryMethod = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetDeliveryRoomName(v string) *ListHotelOrderResponseBodyResult {
	s.DeliveryRoomName = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetDeliveryTime(v int64) *ListHotelOrderResponseBodyResult {
	s.DeliveryTime = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetGmtCreate(v int64) *ListHotelOrderResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetIcon(v string) *ListHotelOrderResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetItemId(v int64) *ListHotelOrderResponseBodyResult {
	s.ItemId = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetItemType(v string) *ListHotelOrderResponseBodyResult {
	s.ItemType = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetName(v string) *ListHotelOrderResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetOrderNo(v string) *ListHotelOrderResponseBodyResult {
	s.OrderNo = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetOrderStatus(v string) *ListHotelOrderResponseBodyResult {
	s.OrderStatus = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetPaymentMethod(v string) *ListHotelOrderResponseBodyResult {
	s.PaymentMethod = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetQuantity(v int64) *ListHotelOrderResponseBodyResult {
	s.Quantity = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetRoomNo(v string) *ListHotelOrderResponseBodyResult {
	s.RoomNo = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetStartTime(v int64) *ListHotelOrderResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetStatus(v string) *ListHotelOrderResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetSumAmt(v int64) *ListHotelOrderResponseBodyResult {
	s.SumAmt = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetType(v string) *ListHotelOrderResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetTypeIconUrl(v string) *ListHotelOrderResponseBodyResult {
	s.TypeIconUrl = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) SetTypeName(v string) *ListHotelOrderResponseBodyResult {
	s.TypeName = &v
	return s
}

func (s *ListHotelOrderResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
