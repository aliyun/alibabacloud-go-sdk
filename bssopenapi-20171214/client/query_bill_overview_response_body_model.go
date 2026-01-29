// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBillOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryBillOverviewResponseBody
	GetCode() *string
	SetData(v *QueryBillOverviewResponseBodyData) *QueryBillOverviewResponseBody
	GetData() *QueryBillOverviewResponseBodyData
	SetMessage(v string) *QueryBillOverviewResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryBillOverviewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryBillOverviewResponseBody
	GetSuccess() *bool
}

type QueryBillOverviewResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryBillOverviewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// BCB1E1BC-05BF-4217-80EE-BF09A29407BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryBillOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryBillOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBillOverviewResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryBillOverviewResponseBody) GetData() *QueryBillOverviewResponseBodyData {
	return s.Data
}

func (s *QueryBillOverviewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryBillOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryBillOverviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryBillOverviewResponseBody) SetCode(v string) *QueryBillOverviewResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBillOverviewResponseBody) SetData(v *QueryBillOverviewResponseBodyData) *QueryBillOverviewResponseBody {
	s.Data = v
	return s
}

func (s *QueryBillOverviewResponseBody) SetMessage(v string) *QueryBillOverviewResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBillOverviewResponseBody) SetRequestId(v string) *QueryBillOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBillOverviewResponseBody) SetSuccess(v bool) *QueryBillOverviewResponseBody {
	s.Success = &v
	return s
}

func (s *QueryBillOverviewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryBillOverviewResponseBodyData struct {
	// The ID of the account.
	//
	// example:
	//
	// 185766xxxx
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
	// 2020-03
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The details of the bills.
	Items *QueryBillOverviewResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s QueryBillOverviewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryBillOverviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBillOverviewResponseBodyData) GetAccountID() *string {
	return s.AccountID
}

func (s *QueryBillOverviewResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *QueryBillOverviewResponseBodyData) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *QueryBillOverviewResponseBodyData) GetItems() *QueryBillOverviewResponseBodyDataItems {
	return s.Items
}

func (s *QueryBillOverviewResponseBodyData) SetAccountID(v string) *QueryBillOverviewResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *QueryBillOverviewResponseBodyData) SetAccountName(v string) *QueryBillOverviewResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *QueryBillOverviewResponseBodyData) SetBillingCycle(v string) *QueryBillOverviewResponseBodyData {
	s.BillingCycle = &v
	return s
}

func (s *QueryBillOverviewResponseBodyData) SetItems(v *QueryBillOverviewResponseBodyDataItems) *QueryBillOverviewResponseBodyData {
	s.Items = v
	return s
}

func (s *QueryBillOverviewResponseBodyData) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryBillOverviewResponseBodyDataItems struct {
	Item []*QueryBillOverviewResponseBodyDataItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QueryBillOverviewResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s QueryBillOverviewResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QueryBillOverviewResponseBodyDataItems) GetItem() []*QueryBillOverviewResponseBodyDataItemsItem {
	return s.Item
}

func (s *QueryBillOverviewResponseBodyDataItems) SetItem(v []*QueryBillOverviewResponseBodyDataItemsItem) *QueryBillOverviewResponseBodyDataItems {
	s.Item = v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItems) Validate() error {
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

type QueryBillOverviewResponseBodyDataItemsItem struct {
	// The amount deducted by using credit refunds.
	//
	// example:
	//
	// 0
	AdjustAmount *float32 `json:"AdjustAmount,omitempty" xml:"AdjustAmount,omitempty"`
	// The amount paid after the tax is deducted.
	//
	// example:
	//
	// 0
	AfterTaxAmount *float32 `json:"AfterTaxAmount,omitempty" xml:"AfterTaxAmount,omitempty"`
	// The ID of the account to which the bill belongs.
	//
	// example:
	//
	// 185766xxxx
	BillAccountID *string `json:"BillAccountID,omitempty" xml:"BillAccountID,omitempty"`
	// The name of the account to which the bill belongs.
	//
	// example:
	//
	// test@test.aliyunid.com
	BillAccountName *string `json:"BillAccountName,omitempty" xml:"BillAccountName,omitempty"`
	// The business type.
	//
	// example:
	//
	// trusteeship
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The amount paid in cash. The amount that was deducted by using credit refunds is not included.
	//
	// example:
	//
	// 0
	CashAmount *float32 `json:"CashAmount,omitempty" xml:"CashAmount,omitempty"`
	// The code of the commodity. The commodity code is the same as that displayed in User Center.
	//
	// example:
	//
	// rds
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The type of the currency. Valid values:
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
	// The type of the bill. Valid values:
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
	// PayAsYouGoBill
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The unsettled amount or the amount deducted by using credits. This may be an unsettled amount of a regular user or credits consumed by a credit user.
	//
	// example:
	//
	// 0
	OutstandingAmount *float32 `json:"OutstandingAmount,omitempty" xml:"OutstandingAmount,omitempty"`
	// The ID of the bill owner.
	//
	// example:
	//
	// 1222
	OwnerID *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// The amount paid in cash. The amount that was deducted by using credit refunds is included.
	//
	// example:
	//
	// 100
	PaymentAmount *float32 `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
	// The currency used for payment.
	//
	// example:
	//
	// USD
	PaymentCurrency *string `json:"PaymentCurrency,omitempty" xml:"PaymentCurrency,omitempty"`
	// The code of the service. The service code is the same as that displayed in User Center.
	//
	// example:
	//
	// rds
	PipCode *string `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	// The pretax amount.
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
	// 100
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
	// NAT Gateway (Pay-As-You-Go)
	ProductDetail *string `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// NAT Gateway
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The type of the service.
	//
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The round down discount.
	//
	// example:
	//
	// 0
	RoundDownDiscount *string `json:"RoundDownDiscount,omitempty" xml:"RoundDownDiscount,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription: the subscription billing method
	//
	// 	- PayAsYouGo: the pay-as-you-go billing method
	//
	// example:
	//
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	// The tax.
	//
	// example:
	//
	// 0
	Tax *float32 `json:"Tax,omitempty" xml:"Tax,omitempty"`
}

func (s QueryBillOverviewResponseBodyDataItemsItem) String() string {
	return dara.Prettify(s)
}

func (s QueryBillOverviewResponseBodyDataItemsItem) GoString() string {
	return s.String()
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetAdjustAmount() *float32 {
	return s.AdjustAmount
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetAfterTaxAmount() *float32 {
	return s.AfterTaxAmount
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetBillAccountID() *string {
	return s.BillAccountID
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetBillAccountName() *string {
	return s.BillAccountName
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetBizType() *string {
	return s.BizType
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetCashAmount() *float32 {
	return s.CashAmount
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetCurrency() *string {
	return s.Currency
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetDeductedByCashCoupons() *float32 {
	return s.DeductedByCashCoupons
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetDeductedByCoupons() *float32 {
	return s.DeductedByCoupons
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetDeductedByPrepaidCard() *float32 {
	return s.DeductedByPrepaidCard
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetInvoiceDiscount() *float32 {
	return s.InvoiceDiscount
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetItem() *string {
	return s.Item
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetOutstandingAmount() *float32 {
	return s.OutstandingAmount
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetOwnerID() *string {
	return s.OwnerID
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetPaymentAmount() *float32 {
	return s.PaymentAmount
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetPaymentCurrency() *string {
	return s.PaymentCurrency
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetPipCode() *string {
	return s.PipCode
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetPretaxAmount() *float32 {
	return s.PretaxAmount
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetPretaxAmountLocal() *float32 {
	return s.PretaxAmountLocal
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetPretaxGrossAmount() *float32 {
	return s.PretaxGrossAmount
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetProductName() *string {
	return s.ProductName
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetProductType() *string {
	return s.ProductType
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetRoundDownDiscount() *string {
	return s.RoundDownDiscount
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) GetTax() *float32 {
	return s.Tax
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetAdjustAmount(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.AdjustAmount = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetAfterTaxAmount(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.AfterTaxAmount = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetBillAccountID(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.BillAccountID = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetBillAccountName(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.BillAccountName = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetBizType(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.BizType = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetCashAmount(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.CashAmount = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetCommodityCode(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.CommodityCode = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetCurrency(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.Currency = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetDeductedByCashCoupons(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetDeductedByCoupons(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.DeductedByCoupons = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetDeductedByPrepaidCard(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetInvoiceDiscount(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.InvoiceDiscount = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetItem(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.Item = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetOutstandingAmount(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.OutstandingAmount = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetOwnerID(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.OwnerID = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetPaymentAmount(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.PaymentAmount = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetPaymentCurrency(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.PaymentCurrency = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetPipCode(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.PipCode = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetPretaxAmount(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.PretaxAmount = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetPretaxAmountLocal(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.PretaxAmountLocal = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetPretaxGrossAmount(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.PretaxGrossAmount = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetProductCode(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.ProductCode = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetProductDetail(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.ProductDetail = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetProductName(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.ProductName = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetProductType(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.ProductType = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetRoundDownDiscount(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.RoundDownDiscount = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetSubscriptionType(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.SubscriptionType = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetTax(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.Tax = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) Validate() error {
	return dara.Validate(s)
}
