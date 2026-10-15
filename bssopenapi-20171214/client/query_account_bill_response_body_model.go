// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountBillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAccountBillResponseBody
	GetCode() *string
	SetData(v *QueryAccountBillResponseBodyData) *QueryAccountBillResponseBody
	GetData() *QueryAccountBillResponseBodyData
	SetMessage(v string) *QueryAccountBillResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAccountBillResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAccountBillResponseBody
	GetSuccess() *bool
}

type QueryAccountBillResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryAccountBillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3BFC23FE-A182-4D96-A1E4-7521B30B8E43
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAccountBillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountBillResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccountBillResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAccountBillResponseBody) GetData() *QueryAccountBillResponseBodyData {
	return s.Data
}

func (s *QueryAccountBillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAccountBillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAccountBillResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAccountBillResponseBody) SetCode(v string) *QueryAccountBillResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAccountBillResponseBody) SetData(v *QueryAccountBillResponseBodyData) *QueryAccountBillResponseBody {
	s.Data = v
	return s
}

func (s *QueryAccountBillResponseBody) SetMessage(v string) *QueryAccountBillResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAccountBillResponseBody) SetRequestId(v string) *QueryAccountBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAccountBillResponseBody) SetSuccess(v bool) *QueryAccountBillResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAccountBillResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAccountBillResponseBodyData struct {
	// The account ID.
	//
	// example:
	//
	// 1857xxxxx489
	AccountID *string `json:"AccountID,omitempty" xml:"AccountID,omitempty"`
	// The account name.
	//
	// example:
	//
	// test@test.aliyunid.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The billing cycle. Format: YYYY-MM.
	//
	// example:
	//
	// 2020-03
	BillingCycle *string                                `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	Items        *QueryAccountBillResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries displayed per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryAccountBillResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountBillResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAccountBillResponseBodyData) GetAccountID() *string {
	return s.AccountID
}

func (s *QueryAccountBillResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *QueryAccountBillResponseBodyData) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *QueryAccountBillResponseBodyData) GetItems() *QueryAccountBillResponseBodyDataItems {
	return s.Items
}

func (s *QueryAccountBillResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryAccountBillResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryAccountBillResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryAccountBillResponseBodyData) SetAccountID(v string) *QueryAccountBillResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *QueryAccountBillResponseBodyData) SetAccountName(v string) *QueryAccountBillResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *QueryAccountBillResponseBodyData) SetBillingCycle(v string) *QueryAccountBillResponseBodyData {
	s.BillingCycle = &v
	return s
}

func (s *QueryAccountBillResponseBodyData) SetItems(v *QueryAccountBillResponseBodyDataItems) *QueryAccountBillResponseBodyData {
	s.Items = v
	return s
}

func (s *QueryAccountBillResponseBodyData) SetPageNum(v int32) *QueryAccountBillResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryAccountBillResponseBodyData) SetPageSize(v int32) *QueryAccountBillResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryAccountBillResponseBodyData) SetTotalCount(v int32) *QueryAccountBillResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryAccountBillResponseBodyData) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAccountBillResponseBodyDataItems struct {
	Item []*QueryAccountBillResponseBodyDataItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QueryAccountBillResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountBillResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QueryAccountBillResponseBodyDataItems) GetItem() []*QueryAccountBillResponseBodyDataItemsItem {
	return s.Item
}

func (s *QueryAccountBillResponseBodyDataItems) SetItem(v []*QueryAccountBillResponseBodyDataItemsItem) *QueryAccountBillResponseBodyDataItems {
	s.Item = v
	return s
}

func (s *QueryAccountBillResponseBodyDataItems) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAccountBillResponseBodyDataItemsItem struct {
	AdjustAmount          *float32 `json:"AdjustAmount,omitempty" xml:"AdjustAmount,omitempty"`
	BillAccountID         *string  `json:"BillAccountID,omitempty" xml:"BillAccountID,omitempty"`
	BillAccountName       *string  `json:"BillAccountName,omitempty" xml:"BillAccountName,omitempty"`
	BillingDate           *string  `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	BizType               *string  `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CashAmount            *float32 `json:"CashAmount,omitempty" xml:"CashAmount,omitempty"`
	CostUnit              *string  `json:"CostUnit,omitempty" xml:"CostUnit,omitempty"`
	Currency              *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DeductedByCashCoupons *float32 `json:"DeductedByCashCoupons,omitempty" xml:"DeductedByCashCoupons,omitempty"`
	DeductedByCoupons     *float32 `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	DeductedByPrepaidCard *float32 `json:"DeductedByPrepaidCard,omitempty" xml:"DeductedByPrepaidCard,omitempty"`
	InvoiceDiscount       *float32 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	OutstandingAmount     *float32 `json:"OutstandingAmount,omitempty" xml:"OutstandingAmount,omitempty"`
	OwnerID               *string  `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	OwnerName             *string  `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	PaymentAmount         *float32 `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
	PipCode               *string  `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	PretaxAmount          *float32 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	PretaxGrossAmount     *float32 `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	ProductCode           *string  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName           *string  `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	SubscriptionType      *string  `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s QueryAccountBillResponseBodyDataItemsItem) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountBillResponseBodyDataItemsItem) GoString() string {
	return s.String()
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetAdjustAmount() *float32 {
	return s.AdjustAmount
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetBillAccountID() *string {
	return s.BillAccountID
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetBillAccountName() *string {
	return s.BillAccountName
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetBillingDate() *string {
	return s.BillingDate
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetBizType() *string {
	return s.BizType
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetCashAmount() *float32 {
	return s.CashAmount
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetCostUnit() *string {
	return s.CostUnit
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetCurrency() *string {
	return s.Currency
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetDeductedByCashCoupons() *float32 {
	return s.DeductedByCashCoupons
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetDeductedByCoupons() *float32 {
	return s.DeductedByCoupons
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetDeductedByPrepaidCard() *float32 {
	return s.DeductedByPrepaidCard
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetInvoiceDiscount() *float32 {
	return s.InvoiceDiscount
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetOutstandingAmount() *float32 {
	return s.OutstandingAmount
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetOwnerID() *string {
	return s.OwnerID
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetOwnerName() *string {
	return s.OwnerName
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetPaymentAmount() *float32 {
	return s.PaymentAmount
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetPipCode() *string {
	return s.PipCode
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetPretaxAmount() *float32 {
	return s.PretaxAmount
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetPretaxGrossAmount() *float32 {
	return s.PretaxGrossAmount
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetProductName() *string {
	return s.ProductName
}

func (s *QueryAccountBillResponseBodyDataItemsItem) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetAdjustAmount(v float32) *QueryAccountBillResponseBodyDataItemsItem {
	s.AdjustAmount = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetBillAccountID(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.BillAccountID = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetBillAccountName(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.BillAccountName = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetBillingDate(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.BillingDate = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetBizType(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.BizType = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetCashAmount(v float32) *QueryAccountBillResponseBodyDataItemsItem {
	s.CashAmount = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetCostUnit(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.CostUnit = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetCurrency(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.Currency = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetDeductedByCashCoupons(v float32) *QueryAccountBillResponseBodyDataItemsItem {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetDeductedByCoupons(v float32) *QueryAccountBillResponseBodyDataItemsItem {
	s.DeductedByCoupons = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetDeductedByPrepaidCard(v float32) *QueryAccountBillResponseBodyDataItemsItem {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetInvoiceDiscount(v float32) *QueryAccountBillResponseBodyDataItemsItem {
	s.InvoiceDiscount = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetOutstandingAmount(v float32) *QueryAccountBillResponseBodyDataItemsItem {
	s.OutstandingAmount = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetOwnerID(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.OwnerID = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetOwnerName(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.OwnerName = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetPaymentAmount(v float32) *QueryAccountBillResponseBodyDataItemsItem {
	s.PaymentAmount = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetPipCode(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.PipCode = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetPretaxAmount(v float32) *QueryAccountBillResponseBodyDataItemsItem {
	s.PretaxAmount = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetPretaxGrossAmount(v float32) *QueryAccountBillResponseBodyDataItemsItem {
	s.PretaxGrossAmount = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetProductCode(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.ProductCode = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetProductName(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.ProductName = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetSubscriptionType(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.SubscriptionType = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) Validate() error {
	return dara.Validate(s)
}
