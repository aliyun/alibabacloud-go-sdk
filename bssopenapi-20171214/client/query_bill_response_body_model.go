// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryBillResponseBody
	GetCode() *string
	SetData(v *QueryBillResponseBodyData) *QueryBillResponseBody
	GetData() *QueryBillResponseBodyData
	SetMessage(v string) *QueryBillResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryBillResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryBillResponseBody
	GetSuccess() *bool
}

type QueryBillResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryBillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AE3F4057-DFC9-4B81-8858-F620651010C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryBillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryBillResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBillResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryBillResponseBody) GetData() *QueryBillResponseBodyData {
	return s.Data
}

func (s *QueryBillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryBillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryBillResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryBillResponseBody) SetCode(v string) *QueryBillResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBillResponseBody) SetData(v *QueryBillResponseBodyData) *QueryBillResponseBody {
	s.Data = v
	return s
}

func (s *QueryBillResponseBody) SetMessage(v string) *QueryBillResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBillResponseBody) SetRequestId(v string) *QueryBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBillResponseBody) SetSuccess(v bool) *QueryBillResponseBody {
	s.Success = &v
	return s
}

func (s *QueryBillResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryBillResponseBodyData struct {
	// The ID of the account.
	//
	// example:
	//
	// 185xxxxx489
	AccountID *string `json:"AccountID,omitempty" xml:"AccountID,omitempty"`
	// The name of the account.
	//
	// example:
	//
	// test@test.aliyunid.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The billing cycle, in the YYYY-MM format.
	//
	// example:
	//
	// 2020-02
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The details of the bills.
	Items *QueryBillResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryBillResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryBillResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBillResponseBodyData) GetAccountID() *string {
	return s.AccountID
}

func (s *QueryBillResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *QueryBillResponseBodyData) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *QueryBillResponseBodyData) GetItems() *QueryBillResponseBodyDataItems {
	return s.Items
}

func (s *QueryBillResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryBillResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryBillResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryBillResponseBodyData) SetAccountID(v string) *QueryBillResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *QueryBillResponseBodyData) SetAccountName(v string) *QueryBillResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *QueryBillResponseBodyData) SetBillingCycle(v string) *QueryBillResponseBodyData {
	s.BillingCycle = &v
	return s
}

func (s *QueryBillResponseBodyData) SetItems(v *QueryBillResponseBodyDataItems) *QueryBillResponseBodyData {
	s.Items = v
	return s
}

func (s *QueryBillResponseBodyData) SetPageNum(v int32) *QueryBillResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryBillResponseBodyData) SetPageSize(v int32) *QueryBillResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryBillResponseBodyData) SetTotalCount(v int32) *QueryBillResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryBillResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryBillResponseBodyDataItems struct {
	Item []*QueryBillResponseBodyDataItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QueryBillResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s QueryBillResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QueryBillResponseBodyDataItems) GetItem() []*QueryBillResponseBodyDataItemsItem {
	return s.Item
}

func (s *QueryBillResponseBodyDataItems) SetItem(v []*QueryBillResponseBodyDataItemsItem) *QueryBillResponseBodyDataItems {
	s.Item = v
	return s
}

func (s *QueryBillResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}

type QueryBillResponseBodyDataItemsItem struct {
	// The amount deducted by using credit refunds.
	//
	// example:
	//
	// 0
	AdjustAmount *float32 `json:"AdjustAmount,omitempty" xml:"AdjustAmount,omitempty"`
	// The amount paid after tax is deducted.
	//
	// example:
	//
	// 0
	AfterTaxAmount *float32 `json:"AfterTaxAmount,omitempty" xml:"AfterTaxAmount,omitempty"`
	// The amount paid in cash. The amount that was deducted by using credit refunds is not included.
	//
	// example:
	//
	// 0
	CashAmount *float32 `json:"CashAmount,omitempty" xml:"CashAmount,omitempty"`
	// The code of the commodity.
	//
	// example:
	//
	// rds
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The type of the currency.
	//
	// 	- CNY
	//
	// 	- USD
	//
	// 	- JPY
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The amount deducted by using vouchers.
	//
	// example:
	//
	// 0
	DeductedByCashCoupons *float32 `json:"DeductedByCashCoupons,omitempty" xml:"DeductedByCashCoupons,omitempty"`
	// The amount deducted by using coupons.
	//
	// example:
	//
	// 0
	DeductedByCoupons *float32 `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	// The amount deducted by using prepaid cards.
	//
	// example:
	//
	// 0
	DeductedByPrepaidCard *float32 `json:"DeductedByPrepaidCard,omitempty" xml:"DeductedByPrepaidCard,omitempty"`
	// The discount amount.
	//
	// example:
	//
	// 0
	InvoiceDiscount *float32 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	// The type of the bill.
	//
	// 	- SubscriptionOrder
	//
	// 	- PayAsYouGoBill
	//
	// 	- Refund
	//
	// 	- Adjustment
	//
	// example:
	//
	// SubscriptionOrder
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The unsettled amount of the bill.
	//
	// example:
	//
	// 0
	OutstandingAmount *float32 `json:"OutstandingAmount,omitempty" xml:"OutstandingAmount,omitempty"`
	// The ID of the member. This parameter is returned in a multi-account payment scenario.
	//
	// example:
	//
	// 3532535235
	OwnerID *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// The amount paid with cash.
	//
	// example:
	//
	// 0
	PaymentAmount *float32 `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
	// The currency used for payment.
	//
	// example:
	//
	// USD
	PaymentCurrency *string `json:"PaymentCurrency,omitempty" xml:"PaymentCurrency,omitempty"`
	// The time when the order was paid.
	//
	// example:
	//
	// 2020-03-11 08:00:00
	PaymentTime *string `json:"PaymentTime,omitempty" xml:"PaymentTime,omitempty"`
	// The ID of the transaction.
	//
	// example:
	//
	// 2O3AADJFEAF2PDKSFAMFCB11918
	PaymentTransactionID *string `json:"PaymentTransactionID,omitempty" xml:"PaymentTransactionID,omitempty"`
	// The code of the service.
	//
	// example:
	//
	// rds
	PipCode *string `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	// The pretax amount
	//
	// example:
	//
	// 100
	PretaxAmount *float32 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	// The pretax amount paid in local currency.
	//
	// example:
	//
	// 0
	PretaxAmountLocal *float32 `json:"PretaxAmountLocal,omitempty" xml:"PretaxAmountLocal,omitempty"`
	// The pretax gross amount.
	//
	// example:
	//
	// 0
	PretaxGrossAmount *float32 `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	// The code of the service.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The details of the service.
	//
	// example:
	//
	// ApsaraDB RDS (subscription)
	ProductDetail *string `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// ApsaraDB RDS
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The type of the service.
	//
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The ID of the order or bill.
	//
	// example:
	//
	// 2020xxxx5912
	RecordID *string `json:"RecordID,omitempty" xml:"RecordID,omitempty"`
	// The round down discount.
	//
	// example:
	//
	// 0
	RoundDownDiscount *string `json:"RoundDownDiscount,omitempty" xml:"RoundDownDiscount,omitempty"`
	// The payment status of the bill. Valid values:
	//
	// 	- PayFinish: The bill is paid.
	//
	// 	- PayUnclear: The bill is not cleared.
	//
	// 	- PayUnsettle: The bill is not settled.
	//
	// 	- NoSettle: The bill is free of settlement.
	//
	// example:
	//
	// PayFinish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the order corresponding to the bill.
	//
	// example:
	//
	// 325345254353
	SubOrderId *string `json:"SubOrderId,omitempty" xml:"SubOrderId,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription
	//
	// 	- PayAsYouGo
	//
	// example:
	//
	// PayAsYouGo
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	// The tax.
	//
	// example:
	//
	// 0
	Tax *float32 `json:"Tax,omitempty" xml:"Tax,omitempty"`
	// The time when the bill ends.
	//
	// example:
	//
	// 2020-03-11 08:00:00
	UsageEndTime *string `json:"UsageEndTime,omitempty" xml:"UsageEndTime,omitempty"`
	// The time when the bill starts.
	//
	// example:
	//
	// 2020-03-11 07:00:00
	UsageStartTime *string `json:"UsageStartTime,omitempty" xml:"UsageStartTime,omitempty"`
}

func (s QueryBillResponseBodyDataItemsItem) String() string {
	return dara.Prettify(s)
}

func (s QueryBillResponseBodyDataItemsItem) GoString() string {
	return s.String()
}

func (s *QueryBillResponseBodyDataItemsItem) GetAdjustAmount() *float32 {
	return s.AdjustAmount
}

func (s *QueryBillResponseBodyDataItemsItem) GetAfterTaxAmount() *float32 {
	return s.AfterTaxAmount
}

func (s *QueryBillResponseBodyDataItemsItem) GetCashAmount() *float32 {
	return s.CashAmount
}

func (s *QueryBillResponseBodyDataItemsItem) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *QueryBillResponseBodyDataItemsItem) GetCurrency() *string {
	return s.Currency
}

func (s *QueryBillResponseBodyDataItemsItem) GetDeductedByCashCoupons() *float32 {
	return s.DeductedByCashCoupons
}

func (s *QueryBillResponseBodyDataItemsItem) GetDeductedByCoupons() *float32 {
	return s.DeductedByCoupons
}

func (s *QueryBillResponseBodyDataItemsItem) GetDeductedByPrepaidCard() *float32 {
	return s.DeductedByPrepaidCard
}

func (s *QueryBillResponseBodyDataItemsItem) GetInvoiceDiscount() *float32 {
	return s.InvoiceDiscount
}

func (s *QueryBillResponseBodyDataItemsItem) GetItem() *string {
	return s.Item
}

func (s *QueryBillResponseBodyDataItemsItem) GetOutstandingAmount() *float32 {
	return s.OutstandingAmount
}

func (s *QueryBillResponseBodyDataItemsItem) GetOwnerID() *string {
	return s.OwnerID
}

func (s *QueryBillResponseBodyDataItemsItem) GetPaymentAmount() *float32 {
	return s.PaymentAmount
}

func (s *QueryBillResponseBodyDataItemsItem) GetPaymentCurrency() *string {
	return s.PaymentCurrency
}

func (s *QueryBillResponseBodyDataItemsItem) GetPaymentTime() *string {
	return s.PaymentTime
}

func (s *QueryBillResponseBodyDataItemsItem) GetPaymentTransactionID() *string {
	return s.PaymentTransactionID
}

func (s *QueryBillResponseBodyDataItemsItem) GetPipCode() *string {
	return s.PipCode
}

func (s *QueryBillResponseBodyDataItemsItem) GetPretaxAmount() *float32 {
	return s.PretaxAmount
}

func (s *QueryBillResponseBodyDataItemsItem) GetPretaxAmountLocal() *float32 {
	return s.PretaxAmountLocal
}

func (s *QueryBillResponseBodyDataItemsItem) GetPretaxGrossAmount() *float32 {
	return s.PretaxGrossAmount
}

func (s *QueryBillResponseBodyDataItemsItem) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryBillResponseBodyDataItemsItem) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *QueryBillResponseBodyDataItemsItem) GetProductName() *string {
	return s.ProductName
}

func (s *QueryBillResponseBodyDataItemsItem) GetProductType() *string {
	return s.ProductType
}

func (s *QueryBillResponseBodyDataItemsItem) GetRecordID() *string {
	return s.RecordID
}

func (s *QueryBillResponseBodyDataItemsItem) GetRoundDownDiscount() *string {
	return s.RoundDownDiscount
}

func (s *QueryBillResponseBodyDataItemsItem) GetStatus() *string {
	return s.Status
}

func (s *QueryBillResponseBodyDataItemsItem) GetSubOrderId() *string {
	return s.SubOrderId
}

func (s *QueryBillResponseBodyDataItemsItem) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *QueryBillResponseBodyDataItemsItem) GetTax() *float32 {
	return s.Tax
}

func (s *QueryBillResponseBodyDataItemsItem) GetUsageEndTime() *string {
	return s.UsageEndTime
}

func (s *QueryBillResponseBodyDataItemsItem) GetUsageStartTime() *string {
	return s.UsageStartTime
}

func (s *QueryBillResponseBodyDataItemsItem) SetAdjustAmount(v float32) *QueryBillResponseBodyDataItemsItem {
	s.AdjustAmount = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetAfterTaxAmount(v float32) *QueryBillResponseBodyDataItemsItem {
	s.AfterTaxAmount = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetCashAmount(v float32) *QueryBillResponseBodyDataItemsItem {
	s.CashAmount = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetCommodityCode(v string) *QueryBillResponseBodyDataItemsItem {
	s.CommodityCode = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetCurrency(v string) *QueryBillResponseBodyDataItemsItem {
	s.Currency = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetDeductedByCashCoupons(v float32) *QueryBillResponseBodyDataItemsItem {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetDeductedByCoupons(v float32) *QueryBillResponseBodyDataItemsItem {
	s.DeductedByCoupons = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetDeductedByPrepaidCard(v float32) *QueryBillResponseBodyDataItemsItem {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetInvoiceDiscount(v float32) *QueryBillResponseBodyDataItemsItem {
	s.InvoiceDiscount = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetItem(v string) *QueryBillResponseBodyDataItemsItem {
	s.Item = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetOutstandingAmount(v float32) *QueryBillResponseBodyDataItemsItem {
	s.OutstandingAmount = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetOwnerID(v string) *QueryBillResponseBodyDataItemsItem {
	s.OwnerID = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetPaymentAmount(v float32) *QueryBillResponseBodyDataItemsItem {
	s.PaymentAmount = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetPaymentCurrency(v string) *QueryBillResponseBodyDataItemsItem {
	s.PaymentCurrency = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetPaymentTime(v string) *QueryBillResponseBodyDataItemsItem {
	s.PaymentTime = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetPaymentTransactionID(v string) *QueryBillResponseBodyDataItemsItem {
	s.PaymentTransactionID = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetPipCode(v string) *QueryBillResponseBodyDataItemsItem {
	s.PipCode = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetPretaxAmount(v float32) *QueryBillResponseBodyDataItemsItem {
	s.PretaxAmount = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetPretaxAmountLocal(v float32) *QueryBillResponseBodyDataItemsItem {
	s.PretaxAmountLocal = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetPretaxGrossAmount(v float32) *QueryBillResponseBodyDataItemsItem {
	s.PretaxGrossAmount = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetProductCode(v string) *QueryBillResponseBodyDataItemsItem {
	s.ProductCode = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetProductDetail(v string) *QueryBillResponseBodyDataItemsItem {
	s.ProductDetail = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetProductName(v string) *QueryBillResponseBodyDataItemsItem {
	s.ProductName = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetProductType(v string) *QueryBillResponseBodyDataItemsItem {
	s.ProductType = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetRecordID(v string) *QueryBillResponseBodyDataItemsItem {
	s.RecordID = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetRoundDownDiscount(v string) *QueryBillResponseBodyDataItemsItem {
	s.RoundDownDiscount = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetStatus(v string) *QueryBillResponseBodyDataItemsItem {
	s.Status = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetSubOrderId(v string) *QueryBillResponseBodyDataItemsItem {
	s.SubOrderId = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetSubscriptionType(v string) *QueryBillResponseBodyDataItemsItem {
	s.SubscriptionType = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetTax(v float32) *QueryBillResponseBodyDataItemsItem {
	s.Tax = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetUsageEndTime(v string) *QueryBillResponseBodyDataItemsItem {
	s.UsageEndTime = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetUsageStartTime(v string) *QueryBillResponseBodyDataItemsItem {
	s.UsageStartTime = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) Validate() error {
	return dara.Validate(s)
}
