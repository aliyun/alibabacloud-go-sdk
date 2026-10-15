// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySettleBillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySettleBillResponseBody
	GetCode() *string
	SetData(v *QuerySettleBillResponseBodyData) *QuerySettleBillResponseBody
	GetData() *QuerySettleBillResponseBodyData
	SetMessage(v string) *QuerySettleBillResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySettleBillResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySettleBillResponseBody
	GetSuccess() *bool
}

type QuerySettleBillResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QuerySettleBillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// AE3F4057-DFC9-4B81-8858-F620651010C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySettleBillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySettleBillResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySettleBillResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySettleBillResponseBody) GetData() *QuerySettleBillResponseBodyData {
	return s.Data
}

func (s *QuerySettleBillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySettleBillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySettleBillResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySettleBillResponseBody) SetCode(v string) *QuerySettleBillResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySettleBillResponseBody) SetData(v *QuerySettleBillResponseBodyData) *QuerySettleBillResponseBody {
	s.Data = v
	return s
}

func (s *QuerySettleBillResponseBody) SetMessage(v string) *QuerySettleBillResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySettleBillResponseBody) SetRequestId(v string) *QuerySettleBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySettleBillResponseBody) SetSuccess(v bool) *QuerySettleBillResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySettleBillResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySettleBillResponseBodyData struct {
	// The account ID.
	//
	// example:
	//
	// 185xxxxx489
	AccountID *string `json:"AccountID,omitempty" xml:"AccountID,omitempty"`
	// The account name.
	//
	// example:
	//
	// test@test.aliyunid.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The billing cycle, in the format of YYYY-MM.
	//
	// example:
	//
	// 2020-02
	BillingCycle *string                               `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	Items        *QuerySettleBillResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The maximum number of records returned in this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that indicates the position where the current call reads to. An empty value indicates that all data has been read.
	//
	// example:
	//
	// CAESEgoQCg4KCmdtdF9jcmVhdGUQARgBIkgKCQBwhGmPcAEAAAo7AzYAAAAxTDgwMDcxMjg3ZDJhNmM3ZDguTDgwMDAwMDAwMDAwMzE1MTIuTDgwMDcyZDMyZTJkYzg3N2U
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QuerySettleBillResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySettleBillResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySettleBillResponseBodyData) GetAccountID() *string {
	return s.AccountID
}

func (s *QuerySettleBillResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *QuerySettleBillResponseBodyData) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *QuerySettleBillResponseBodyData) GetItems() *QuerySettleBillResponseBodyDataItems {
	return s.Items
}

func (s *QuerySettleBillResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QuerySettleBillResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *QuerySettleBillResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QuerySettleBillResponseBodyData) SetAccountID(v string) *QuerySettleBillResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *QuerySettleBillResponseBodyData) SetAccountName(v string) *QuerySettleBillResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *QuerySettleBillResponseBodyData) SetBillingCycle(v string) *QuerySettleBillResponseBodyData {
	s.BillingCycle = &v
	return s
}

func (s *QuerySettleBillResponseBodyData) SetItems(v *QuerySettleBillResponseBodyDataItems) *QuerySettleBillResponseBodyData {
	s.Items = v
	return s
}

func (s *QuerySettleBillResponseBodyData) SetMaxResults(v int32) *QuerySettleBillResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *QuerySettleBillResponseBodyData) SetNextToken(v string) *QuerySettleBillResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *QuerySettleBillResponseBodyData) SetTotalCount(v int32) *QuerySettleBillResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QuerySettleBillResponseBodyData) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySettleBillResponseBodyDataItems struct {
	Item []*QuerySettleBillResponseBodyDataItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QuerySettleBillResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s QuerySettleBillResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QuerySettleBillResponseBodyDataItems) GetItem() []*QuerySettleBillResponseBodyDataItemsItem {
	return s.Item
}

func (s *QuerySettleBillResponseBodyDataItems) SetItem(v []*QuerySettleBillResponseBodyDataItemsItem) *QuerySettleBillResponseBodyDataItems {
	s.Item = v
	return s
}

func (s *QuerySettleBillResponseBodyDataItems) Validate() error {
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

type QuerySettleBillResponseBodyDataItemsItem struct {
	AdjustAmount          *float32 `json:"AdjustAmount,omitempty" xml:"AdjustAmount,omitempty"`
	AfterTaxAmount        *float32 `json:"AfterTaxAmount,omitempty" xml:"AfterTaxAmount,omitempty"`
	BillAccountID         *string  `json:"BillAccountID,omitempty" xml:"BillAccountID,omitempty"`
	BillAccountName       *string  `json:"BillAccountName,omitempty" xml:"BillAccountName,omitempty"`
	BizType               *string  `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CashAmount            *float32 `json:"CashAmount,omitempty" xml:"CashAmount,omitempty"`
	CommodityCode         *string  `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Currency              *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DeductedByCashCoupons *float32 `json:"DeductedByCashCoupons,omitempty" xml:"DeductedByCashCoupons,omitempty"`
	DeductedByCoupons     *float32 `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	DeductedByPrepaidCard *float32 `json:"DeductedByPrepaidCard,omitempty" xml:"DeductedByPrepaidCard,omitempty"`
	InvoiceDiscount       *float32 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	Item                  *string  `json:"Item,omitempty" xml:"Item,omitempty"`
	OutstandingAmount     *float32 `json:"OutstandingAmount,omitempty" xml:"OutstandingAmount,omitempty"`
	OwnerID               *string  `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	PaymentAmount         *float32 `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
	PaymentCurrency       *string  `json:"PaymentCurrency,omitempty" xml:"PaymentCurrency,omitempty"`
	PaymentTime           *string  `json:"PaymentTime,omitempty" xml:"PaymentTime,omitempty"`
	PaymentTransactionID  *string  `json:"PaymentTransactionID,omitempty" xml:"PaymentTransactionID,omitempty"`
	PipCode               *string  `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	PretaxAmount          *float32 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	PretaxAmountLocal     *float32 `json:"PretaxAmountLocal,omitempty" xml:"PretaxAmountLocal,omitempty"`
	PretaxGrossAmount     *float32 `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	ProductCode           *string  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductDetail         *string  `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	ProductName           *string  `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductType           *string  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RecordID              *string  `json:"RecordID,omitempty" xml:"RecordID,omitempty"`
	RoundDownDiscount     *string  `json:"RoundDownDiscount,omitempty" xml:"RoundDownDiscount,omitempty"`
	Status                *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	SubOrderId            *string  `json:"SubOrderId,omitempty" xml:"SubOrderId,omitempty"`
	SubscriptionType      *string  `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	Tax                   *float32 `json:"Tax,omitempty" xml:"Tax,omitempty"`
	UsageEndTime          *string  `json:"UsageEndTime,omitempty" xml:"UsageEndTime,omitempty"`
	UsageStartTime        *string  `json:"UsageStartTime,omitempty" xml:"UsageStartTime,omitempty"`
}

func (s QuerySettleBillResponseBodyDataItemsItem) String() string {
	return dara.Prettify(s)
}

func (s QuerySettleBillResponseBodyDataItemsItem) GoString() string {
	return s.String()
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetAdjustAmount() *float32 {
	return s.AdjustAmount
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetAfterTaxAmount() *float32 {
	return s.AfterTaxAmount
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetBillAccountID() *string {
	return s.BillAccountID
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetBillAccountName() *string {
	return s.BillAccountName
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetBizType() *string {
	return s.BizType
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetCashAmount() *float32 {
	return s.CashAmount
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetCurrency() *string {
	return s.Currency
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetDeductedByCashCoupons() *float32 {
	return s.DeductedByCashCoupons
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetDeductedByCoupons() *float32 {
	return s.DeductedByCoupons
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetDeductedByPrepaidCard() *float32 {
	return s.DeductedByPrepaidCard
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetInvoiceDiscount() *float32 {
	return s.InvoiceDiscount
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetItem() *string {
	return s.Item
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetOutstandingAmount() *float32 {
	return s.OutstandingAmount
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetOwnerID() *string {
	return s.OwnerID
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetPaymentAmount() *float32 {
	return s.PaymentAmount
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetPaymentCurrency() *string {
	return s.PaymentCurrency
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetPaymentTime() *string {
	return s.PaymentTime
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetPaymentTransactionID() *string {
	return s.PaymentTransactionID
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetPipCode() *string {
	return s.PipCode
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetPretaxAmount() *float32 {
	return s.PretaxAmount
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetPretaxAmountLocal() *float32 {
	return s.PretaxAmountLocal
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetPretaxGrossAmount() *float32 {
	return s.PretaxGrossAmount
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetProductCode() *string {
	return s.ProductCode
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetProductName() *string {
	return s.ProductName
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetProductType() *string {
	return s.ProductType
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetRecordID() *string {
	return s.RecordID
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetRoundDownDiscount() *string {
	return s.RoundDownDiscount
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetStatus() *string {
	return s.Status
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetSubOrderId() *string {
	return s.SubOrderId
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetTax() *float32 {
	return s.Tax
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetUsageEndTime() *string {
	return s.UsageEndTime
}

func (s *QuerySettleBillResponseBodyDataItemsItem) GetUsageStartTime() *string {
	return s.UsageStartTime
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetAdjustAmount(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.AdjustAmount = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetAfterTaxAmount(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.AfterTaxAmount = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetBillAccountID(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.BillAccountID = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetBillAccountName(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.BillAccountName = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetBizType(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.BizType = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetCashAmount(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.CashAmount = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetCommodityCode(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.CommodityCode = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetCurrency(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.Currency = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetDeductedByCashCoupons(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetDeductedByCoupons(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.DeductedByCoupons = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetDeductedByPrepaidCard(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetInvoiceDiscount(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.InvoiceDiscount = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetItem(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.Item = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetOutstandingAmount(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.OutstandingAmount = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetOwnerID(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.OwnerID = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetPaymentAmount(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.PaymentAmount = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetPaymentCurrency(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.PaymentCurrency = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetPaymentTime(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.PaymentTime = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetPaymentTransactionID(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.PaymentTransactionID = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetPipCode(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.PipCode = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetPretaxAmount(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.PretaxAmount = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetPretaxAmountLocal(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.PretaxAmountLocal = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetPretaxGrossAmount(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.PretaxGrossAmount = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetProductCode(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.ProductCode = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetProductDetail(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.ProductDetail = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetProductName(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.ProductName = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetProductType(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.ProductType = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetRecordID(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.RecordID = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetRoundDownDiscount(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.RoundDownDiscount = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetStatus(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.Status = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetSubOrderId(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.SubOrderId = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetSubscriptionType(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.SubscriptionType = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetTax(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.Tax = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetUsageEndTime(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.UsageEndTime = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetUsageStartTime(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.UsageStartTime = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) Validate() error {
	return dara.Validate(s)
}
