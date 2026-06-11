// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubPartnerOrderListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSubPartnerOrderListResponseBody
	GetCode() *string
	SetData(v []*GetSubPartnerOrderListResponseBodyData) *GetSubPartnerOrderListResponseBody
	GetData() []*GetSubPartnerOrderListResponseBodyData
	SetMessage(v string) *GetSubPartnerOrderListResponseBody
	GetMessage() *string
	SetPageNo(v int32) *GetSubPartnerOrderListResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *GetSubPartnerOrderListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetSubPartnerOrderListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSubPartnerOrderListResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *GetSubPartnerOrderListResponseBody
	GetTotal() *int32
}

type GetSubPartnerOrderListResponseBody struct {
	// Status code
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data
	Data []*GetSubPartnerOrderListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Message
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Paging size
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// TotalCount indicates the total amount of data under the current request conditions. This parameter is optional and is not returned by default.
	//
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSubPartnerOrderListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSubPartnerOrderListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubPartnerOrderListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSubPartnerOrderListResponseBody) GetData() []*GetSubPartnerOrderListResponseBodyData {
	return s.Data
}

func (s *GetSubPartnerOrderListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSubPartnerOrderListResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetSubPartnerOrderListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSubPartnerOrderListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSubPartnerOrderListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSubPartnerOrderListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetSubPartnerOrderListResponseBody) SetCode(v string) *GetSubPartnerOrderListResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBody) SetData(v []*GetSubPartnerOrderListResponseBodyData) *GetSubPartnerOrderListResponseBody {
	s.Data = v
	return s
}

func (s *GetSubPartnerOrderListResponseBody) SetMessage(v string) *GetSubPartnerOrderListResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBody) SetPageNo(v int32) *GetSubPartnerOrderListResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBody) SetPageSize(v int32) *GetSubPartnerOrderListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBody) SetRequestId(v string) *GetSubPartnerOrderListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBody) SetSuccess(v bool) *GetSubPartnerOrderListResponseBody {
	s.Success = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBody) SetTotal(v int32) *GetSubPartnerOrderListResponseBody {
	s.Total = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSubPartnerOrderListResponseBodyData struct {
	// Order discount
	//
	// example:
	//
	// 0.9
	AmountDiscount *float64 `json:"AmountDiscount,omitempty" xml:"AmountDiscount,omitempty"`
	// Actual payment amount
	//
	// example:
	//
	// 3750
	AmountDue *float64 `json:"AmountDue,omitempty" xml:"AmountDue,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 2024-07-07 07:52:22
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// Customer classification
	//
	// example:
	//
	// C类
	CustomerClassification *string `json:"CustomerClassification,omitempty" xml:"CustomerClassification,omitempty"`
	// Coupon amount
	//
	// example:
	//
	// 0
	DeductedAmountByCoupons *float64 `json:"DeductedAmountByCoupons,omitempty" xml:"DeductedAmountByCoupons,omitempty"`
	// Discounted price
	//
	// example:
	//
	// 3750
	DiscountedPrice *float64 `json:"DiscountedPrice,omitempty" xml:"DiscountedPrice,omitempty"`
	// Order ID
	//
	// example:
	//
	// 236414227150922
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Order status:
	//
	// - 1 Unpaid
	//
	// - 2 Abandoned
	//
	// - 3 Paid
	//
	// example:
	//
	// 3
	OrderStatus *int32 `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// Order type:
	//
	// - BUY: New purchase
	//
	// - UPGRADE: Upgrade
	//
	// - DOWNGRADE: Downgrade
	//
	// - RENEW: Renewal
	//
	// - REFUND: Refund
	//
	// - OTHERS: Others
	//
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// Payment Time
	//
	// example:
	//
	// 2024-07-07 07:52:22
	PaidAt *string `json:"PaidAt,omitempty" xml:"PaidAt,omitempty"`
	// Payment type:
	//
	// 1: Non-agent payment
	//
	// 2: Agent payment
	//
	// example:
	//
	// 1
	PayType *int32 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Original price/List price
	//
	// example:
	//
	// 3750
	Price *float64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// Product code
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Product name.
	//
	// example:
	//
	// 弹性计算
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// Opportunity ID
	//
	// example:
	//
	// 202502233443
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Sub-partner Name
	//
	// example:
	//
	// xxx有限公司
	SubPartnerName *string `json:"SubPartnerName,omitempty" xml:"SubPartnerName,omitempty"`
	// Secondary partner UID
	//
	// example:
	//
	// 1123132
	SubPartnerUid *int64 `json:"SubPartnerUid,omitempty" xml:"SubPartnerUid,omitempty"`
}

func (s GetSubPartnerOrderListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSubPartnerOrderListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSubPartnerOrderListResponseBodyData) GetAmountDiscount() *float64 {
	return s.AmountDiscount
}

func (s *GetSubPartnerOrderListResponseBodyData) GetAmountDue() *float64 {
	return s.AmountDue
}

func (s *GetSubPartnerOrderListResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetSubPartnerOrderListResponseBodyData) GetCustomerClassification() *string {
	return s.CustomerClassification
}

func (s *GetSubPartnerOrderListResponseBodyData) GetDeductedAmountByCoupons() *float64 {
	return s.DeductedAmountByCoupons
}

func (s *GetSubPartnerOrderListResponseBodyData) GetDiscountedPrice() *float64 {
	return s.DiscountedPrice
}

func (s *GetSubPartnerOrderListResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetSubPartnerOrderListResponseBodyData) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *GetSubPartnerOrderListResponseBodyData) GetOrderType() *string {
	return s.OrderType
}

func (s *GetSubPartnerOrderListResponseBodyData) GetPaidAt() *string {
	return s.PaidAt
}

func (s *GetSubPartnerOrderListResponseBodyData) GetPayType() *int32 {
	return s.PayType
}

func (s *GetSubPartnerOrderListResponseBodyData) GetPrice() *float64 {
	return s.Price
}

func (s *GetSubPartnerOrderListResponseBodyData) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetSubPartnerOrderListResponseBodyData) GetProductName() *string {
	return s.ProductName
}

func (s *GetSubPartnerOrderListResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetSubPartnerOrderListResponseBodyData) GetSubPartnerName() *string {
	return s.SubPartnerName
}

func (s *GetSubPartnerOrderListResponseBodyData) GetSubPartnerUid() *int64 {
	return s.SubPartnerUid
}

func (s *GetSubPartnerOrderListResponseBodyData) SetAmountDiscount(v float64) *GetSubPartnerOrderListResponseBodyData {
	s.AmountDiscount = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetAmountDue(v float64) *GetSubPartnerOrderListResponseBodyData {
	s.AmountDue = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetCreatedAt(v string) *GetSubPartnerOrderListResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetCustomerClassification(v string) *GetSubPartnerOrderListResponseBodyData {
	s.CustomerClassification = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetDeductedAmountByCoupons(v float64) *GetSubPartnerOrderListResponseBodyData {
	s.DeductedAmountByCoupons = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetDiscountedPrice(v float64) *GetSubPartnerOrderListResponseBodyData {
	s.DiscountedPrice = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetOrderId(v int64) *GetSubPartnerOrderListResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetOrderStatus(v int32) *GetSubPartnerOrderListResponseBodyData {
	s.OrderStatus = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetOrderType(v string) *GetSubPartnerOrderListResponseBodyData {
	s.OrderType = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetPaidAt(v string) *GetSubPartnerOrderListResponseBodyData {
	s.PaidAt = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetPayType(v int32) *GetSubPartnerOrderListResponseBodyData {
	s.PayType = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetPrice(v float64) *GetSubPartnerOrderListResponseBodyData {
	s.Price = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetProductCode(v string) *GetSubPartnerOrderListResponseBodyData {
	s.ProductCode = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetProductName(v string) *GetSubPartnerOrderListResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetProjectId(v int64) *GetSubPartnerOrderListResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetSubPartnerName(v string) *GetSubPartnerOrderListResponseBodyData {
	s.SubPartnerName = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) SetSubPartnerUid(v int64) *GetSubPartnerOrderListResponseBodyData {
	s.SubPartnerUid = &v
	return s
}

func (s *GetSubPartnerOrderListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
