// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerOrderListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetCustomerOrderListResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetCustomerOrderListResponseBody
	GetCode() *string
	SetData(v []*GetCustomerOrderListResponseBodyData) *GetCustomerOrderListResponseBody
	GetData() []*GetCustomerOrderListResponseBodyData
	SetHttpStatusCode(v int32) *GetCustomerOrderListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCustomerOrderListResponseBody
	GetMessage() *string
	SetPageNo(v int32) *GetCustomerOrderListResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *GetCustomerOrderListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetCustomerOrderListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomerOrderListResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *GetCustomerOrderListResponseBody
	GetTotal() *int32
}

type GetCustomerOrderListResponseBody struct {
	// Access denied details
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Status Code
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data
	Data []*GetCustomerOrderListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Message
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 2103a30617045934095083027d88c5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total number of entries
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetCustomerOrderListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerOrderListResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomerOrderListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetCustomerOrderListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCustomerOrderListResponseBody) GetData() []*GetCustomerOrderListResponseBodyData {
	return s.Data
}

func (s *GetCustomerOrderListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCustomerOrderListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCustomerOrderListResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetCustomerOrderListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetCustomerOrderListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomerOrderListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomerOrderListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetCustomerOrderListResponseBody) SetAccessDeniedDetail(v string) *GetCustomerOrderListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetCustomerOrderListResponseBody) SetCode(v string) *GetCustomerOrderListResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomerOrderListResponseBody) SetData(v []*GetCustomerOrderListResponseBodyData) *GetCustomerOrderListResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomerOrderListResponseBody) SetHttpStatusCode(v int32) *GetCustomerOrderListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCustomerOrderListResponseBody) SetMessage(v string) *GetCustomerOrderListResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomerOrderListResponseBody) SetPageNo(v int32) *GetCustomerOrderListResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetCustomerOrderListResponseBody) SetPageSize(v int32) *GetCustomerOrderListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetCustomerOrderListResponseBody) SetRequestId(v string) *GetCustomerOrderListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomerOrderListResponseBody) SetSuccess(v bool) *GetCustomerOrderListResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomerOrderListResponseBody) SetTotal(v int32) *GetCustomerOrderListResponseBody {
	s.Total = &v
	return s
}

func (s *GetCustomerOrderListResponseBody) Validate() error {
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

type GetCustomerOrderListResponseBodyData struct {
	// Order discount
	//
	// example:
	//
	// 1
	AmountDiscount *float64 `json:"AmountDiscount,omitempty" xml:"AmountDiscount,omitempty"`
	// Actual payment amount
	//
	// example:
	//
	// 29137
	AmountDue *float64 `json:"AmountDue,omitempty" xml:"AmountDue,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 2019-01-24 14:20:40
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// Customer Account
	//
	// example:
	//
	// test_123
	CustomerAccount *string `json:"CustomerAccount,omitempty" xml:"CustomerAccount,omitempty"`
	// Customer categorization
	//
	// example:
	//
	// C类
	CustomerClassification *string `json:"CustomerClassification,omitempty" xml:"CustomerClassification,omitempty"`
	// Customer UID
	//
	// example:
	//
	// 123456
	CustomerUid *int64 `json:"CustomerUid,omitempty" xml:"CustomerUid,omitempty"`
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
	// 29137
	DiscountedPrice *float64 `json:"DiscountedPrice,omitempty" xml:"DiscountedPrice,omitempty"`
	// Order ID
	//
	// example:
	//
	// 236414227150922
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Order status. Values include:
	//
	// 1: Unpaid
	//
	// 2: Paid
	//
	// 3: Voided
	//
	// example:
	//
	// 3
	OrderStatus *int32 `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// Order type. Values include: BUY, UPGRADE, DOWNGRADE, RENEW, REFUND, OTHERS
	//
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// Payment Time
	//
	// example:
	//
	// 2019-01-24 14:20:40
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
	// Original Price/List Price
	//
	// example:
	//
	// 29137
	Price *float64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// Product code
	//
	// example:
	//
	// slb
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Product name
	//
	// example:
	//
	// slb
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// Opportunity ID
	//
	// example:
	//
	// 202502230421
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Customer-facing staff
	RamAccountForCustomerManagers []*string `json:"RamAccountForCustomerManagers,omitempty" xml:"RamAccountForCustomerManagers,omitempty" type:"Repeated"`
}

func (s GetCustomerOrderListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerOrderListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomerOrderListResponseBodyData) GetAmountDiscount() *float64 {
	return s.AmountDiscount
}

func (s *GetCustomerOrderListResponseBodyData) GetAmountDue() *float64 {
	return s.AmountDue
}

func (s *GetCustomerOrderListResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetCustomerOrderListResponseBodyData) GetCustomerAccount() *string {
	return s.CustomerAccount
}

func (s *GetCustomerOrderListResponseBodyData) GetCustomerClassification() *string {
	return s.CustomerClassification
}

func (s *GetCustomerOrderListResponseBodyData) GetCustomerUid() *int64 {
	return s.CustomerUid
}

func (s *GetCustomerOrderListResponseBodyData) GetDeductedAmountByCoupons() *float64 {
	return s.DeductedAmountByCoupons
}

func (s *GetCustomerOrderListResponseBodyData) GetDiscountedPrice() *float64 {
	return s.DiscountedPrice
}

func (s *GetCustomerOrderListResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetCustomerOrderListResponseBodyData) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *GetCustomerOrderListResponseBodyData) GetOrderType() *string {
	return s.OrderType
}

func (s *GetCustomerOrderListResponseBodyData) GetPaidAt() *string {
	return s.PaidAt
}

func (s *GetCustomerOrderListResponseBodyData) GetPayType() *int32 {
	return s.PayType
}

func (s *GetCustomerOrderListResponseBodyData) GetPrice() *float64 {
	return s.Price
}

func (s *GetCustomerOrderListResponseBodyData) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetCustomerOrderListResponseBodyData) GetProductName() *string {
	return s.ProductName
}

func (s *GetCustomerOrderListResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetCustomerOrderListResponseBodyData) GetRamAccountForCustomerManagers() []*string {
	return s.RamAccountForCustomerManagers
}

func (s *GetCustomerOrderListResponseBodyData) SetAmountDiscount(v float64) *GetCustomerOrderListResponseBodyData {
	s.AmountDiscount = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetAmountDue(v float64) *GetCustomerOrderListResponseBodyData {
	s.AmountDue = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetCreatedAt(v string) *GetCustomerOrderListResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetCustomerAccount(v string) *GetCustomerOrderListResponseBodyData {
	s.CustomerAccount = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetCustomerClassification(v string) *GetCustomerOrderListResponseBodyData {
	s.CustomerClassification = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetCustomerUid(v int64) *GetCustomerOrderListResponseBodyData {
	s.CustomerUid = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetDeductedAmountByCoupons(v float64) *GetCustomerOrderListResponseBodyData {
	s.DeductedAmountByCoupons = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetDiscountedPrice(v float64) *GetCustomerOrderListResponseBodyData {
	s.DiscountedPrice = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetOrderId(v int64) *GetCustomerOrderListResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetOrderStatus(v int32) *GetCustomerOrderListResponseBodyData {
	s.OrderStatus = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetOrderType(v string) *GetCustomerOrderListResponseBodyData {
	s.OrderType = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetPaidAt(v string) *GetCustomerOrderListResponseBodyData {
	s.PaidAt = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetPayType(v int32) *GetCustomerOrderListResponseBodyData {
	s.PayType = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetPrice(v float64) *GetCustomerOrderListResponseBodyData {
	s.Price = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetProductCode(v string) *GetCustomerOrderListResponseBodyData {
	s.ProductCode = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetProductName(v string) *GetCustomerOrderListResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetProjectId(v int64) *GetCustomerOrderListResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) SetRamAccountForCustomerManagers(v []*string) *GetCustomerOrderListResponseBodyData {
	s.RamAccountForCustomerManagers = v
	return s
}

func (s *GetCustomerOrderListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
