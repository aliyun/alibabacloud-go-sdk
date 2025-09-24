// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductAmortizedCostByAmortizationPeriodResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBody
	GetCode() *string
	SetData(v *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData) *DescribeProductAmortizedCostByAmortizationPeriodResponseBody
	GetData() *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData
	SetMessage(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeProductAmortizedCostByAmortizationPeriodResponseBody
	GetSuccess() *bool
}

type DescribeProductAmortizedCostByAmortizationPeriodResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Successful！
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeProductAmortizedCostByAmortizationPeriodResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductAmortizedCostByAmortizationPeriodResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBody) GetData() *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData {
	return s.Data
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBody) SetCode(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBody) SetData(v *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData) *DescribeProductAmortizedCostByAmortizationPeriodResponseBody {
	s.Data = v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBody) SetMessage(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBody) SetRequestId(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBody) SetSuccess(v bool) *DescribeProductAmortizedCostByAmortizationPeriodResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData struct {
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
	// The data entries returned.
	Items []*DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The position from which the results started to return. The parameter must be left empty or set to the value of the NextToken parameter returned from the last call. Otherwise, an error is returned. If this parameter is left empty, data is queried from the beginning.
	//
	// example:
	//
	// CAESEgoQCg4KCmdtdF9jcmVhdGUEARgBIkgKCQBwhGmPcAEAAAo7AzYAAAAxTDgwMDcxMjg3ZDJhNmM3ZDguTDgwMDAwMDAwMDAwMzE1MTIuTDgwMDcyZDMyZTJkYzg3N2U
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData) GetAccountID() *string {
	return s.AccountID
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData) GetItems() []*DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	return s.Items
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData) SetAccountID(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData) SetAccountName(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData) SetItems(v []*DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData {
	s.Items = v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData) SetMaxResults(v int32) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData) SetNextToken(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData) SetTotalCount(v int32) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems struct {
	AfterDiscountAmount *float64 `json:"AfterDiscountAmount,omitempty" xml:"AfterDiscountAmount,omitempty"`
	// The allocation month. Format: YYYYMM.
	//
	// example:
	//
	// 202210
	AmortizationPeriod *string `json:"AmortizationPeriod,omitempty" xml:"AmortizationPeriod,omitempty"`
	// The allocation status. Valid values:
	//
	// 	- amortized: allocated
	//
	// 	- unAmortized: not allocated
	//
	// example:
	//
	// amortized
	AmortizationStatus *string `json:"AmortizationStatus,omitempty" xml:"AmortizationStatus,omitempty"`
	// The ID of the account to which the bill belongs.
	//
	// example:
	//
	// 185xxxxx489
	BillAccountID *int64 `json:"BillAccountID,omitempty" xml:"BillAccountID,omitempty"`
	// The name of the account to which the bill belongs.
	//
	// example:
	//
	// test@test.aliyunid.com
	BillAccountName *string `json:"BillAccountName,omitempty" xml:"BillAccountName,omitempty"`
	// The ID of the account to which the resource belongs.
	//
	// example:
	//
	// 185xxxxx489
	BillOwnerID *int64 `json:"BillOwnerID,omitempty" xml:"BillOwnerID,omitempty"`
	// The name of the account to which the resource belongs.
	//
	// example:
	//
	// test@test.aliyunid.com
	BillOwnerName *string `json:"BillOwnerName,omitempty" xml:"BillOwnerName,omitempty"`
	// The business type.
	//
	// example:
	//
	// trusteeship
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The billing cycle. Format: YYYYMM.
	//
	// example:
	//
	// 202210
	ConsumePeriod                          *string  `json:"ConsumePeriod,omitempty" xml:"ConsumePeriod,omitempty"`
	CurrentAmortizationAfterDiscountAmount *float64 `json:"CurrentAmortizationAfterDiscountAmount,omitempty" xml:"CurrentAmortizationAfterDiscountAmount,omitempty"`
	// The amount deducted by using vouchers and allocated to the current allocation month.
	//
	// example:
	//
	// 0
	CurrentAmortizationDeductedByCashCoupons *float64 `json:"CurrentAmortizationDeductedByCashCoupons,omitempty" xml:"CurrentAmortizationDeductedByCashCoupons,omitempty"`
	// The amount deducted by using coupons and allocated to the current allocation month.
	//
	// example:
	//
	// 0
	CurrentAmortizationDeductedByCoupons *float64 `json:"CurrentAmortizationDeductedByCoupons,omitempty" xml:"CurrentAmortizationDeductedByCoupons,omitempty"`
	// The amount deducted by using prepaid cards and allocated to the current allocation month.
	//
	// example:
	//
	// 0
	CurrentAmortizationDeductedByPrepaidCard *float64 `json:"CurrentAmortizationDeductedByPrepaidCard,omitempty" xml:"CurrentAmortizationDeductedByPrepaidCard,omitempty"`
	// The expenditure amount allocated to the current allocation month. Invoicing is supported.
	//
	// example:
	//
	// 0
	CurrentAmortizationExpenditureAmount *float64 `json:"CurrentAmortizationExpenditureAmount,omitempty" xml:"CurrentAmortizationExpenditureAmount,omitempty"`
	// The discount amount allocated to the current allocation month.
	//
	// example:
	//
	// 0
	CurrentAmortizationInvoiceDiscount *float64 `json:"CurrentAmortizationInvoiceDiscount,omitempty" xml:"CurrentAmortizationInvoiceDiscount,omitempty"`
	// The pretax amount allocated to the current allocation month.
	//
	// example:
	//
	// 0
	CurrentAmortizationPretaxAmount *float64 `json:"CurrentAmortizationPretaxAmount,omitempty" xml:"CurrentAmortizationPretaxAmount,omitempty"`
	// The pretax gross amount allocated to the current allocation month.
	//
	// example:
	//
	// 0
	CurrentAmortizationPretaxGrossAmount *float64 `json:"CurrentAmortizationPretaxGrossAmount,omitempty" xml:"CurrentAmortizationPretaxGrossAmount,omitempty"`
	// The round-off amount allocated to the current allocation month.
	//
	// example:
	//
	// 0
	CurrentAmortizationRoundDownDiscount *float64 `json:"CurrentAmortizationRoundDownDiscount,omitempty" xml:"CurrentAmortizationRoundDownDiscount,omitempty"`
	// The amount deducted by using vouchers.
	//
	// example:
	//
	// 0
	DeductedByCashCoupons *float64 `json:"DeductedByCashCoupons,omitempty" xml:"DeductedByCashCoupons,omitempty"`
	// The amount deducted by using coupons.
	//
	// example:
	//
	// 0
	DeductedByCoupons *float64 `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	// The amount deducted by using prepaid cards.
	//
	// example:
	//
	// 0
	DeductedByPrepaidCard *float64 `json:"DeductedByPrepaidCard,omitempty" xml:"DeductedByPrepaidCard,omitempty"`
	// The expenditure amount. Invoicing is supported.
	//
	// example:
	//
	// 0
	ExpenditureAmount *float64 `json:"ExpenditureAmount,omitempty" xml:"ExpenditureAmount,omitempty"`
	// The discount amount.
	//
	// example:
	//
	// 0
	InvoiceDiscount *float64 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	// The pretax amount.
	//
	// example:
	//
	// 0
	PretaxAmount *float64 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	// The pretax gross amount.
	//
	// example:
	//
	// 0
	PretaxGrossAmount                      *float64 `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	PreviouslyAmortizedAfterDiscountAmount *float64 `json:"PreviouslyAmortizedAfterDiscountAmount,omitempty" xml:"PreviouslyAmortizedAfterDiscountAmount,omitempty"`
	// The amount deducted by using vouchers and allocated before the current allocation month.
	//
	// example:
	//
	// 0
	PreviouslyAmortizedDeductedByCashCoupons *float64 `json:"PreviouslyAmortizedDeductedByCashCoupons,omitempty" xml:"PreviouslyAmortizedDeductedByCashCoupons,omitempty"`
	// The amount deducted by using coupons and allocated before the current allocation month.
	//
	// example:
	//
	// 0
	PreviouslyAmortizedDeductedByCoupons *float64 `json:"PreviouslyAmortizedDeductedByCoupons,omitempty" xml:"PreviouslyAmortizedDeductedByCoupons,omitempty"`
	// The amount deducted by using prepaid cards and allocated before the current allocation month.
	//
	// example:
	//
	// 0
	PreviouslyAmortizedDeductedByPrepaidCard *float64 `json:"PreviouslyAmortizedDeductedByPrepaidCard,omitempty" xml:"PreviouslyAmortizedDeductedByPrepaidCard,omitempty"`
	// The expenditure amount allocated before the current allocation month. Invoicing is supported.
	//
	// example:
	//
	// 0
	PreviouslyAmortizedExpenditureAmount *float64 `json:"PreviouslyAmortizedExpenditureAmount,omitempty" xml:"PreviouslyAmortizedExpenditureAmount,omitempty"`
	// The discount amount allocated before the current allocation month.
	//
	// example:
	//
	// 0
	PreviouslyAmortizedInvoiceDiscount *float64 `json:"PreviouslyAmortizedInvoiceDiscount,omitempty" xml:"PreviouslyAmortizedInvoiceDiscount,omitempty"`
	// The pretax amount allocated before the current allocation month.
	//
	// example:
	//
	// 0
	PreviouslyAmortizedPretaxAmount *float64 `json:"PreviouslyAmortizedPretaxAmount,omitempty" xml:"PreviouslyAmortizedPretaxAmount,omitempty"`
	// The pretax gross amount allocated before the current allocation month.
	//
	// example:
	//
	// 0
	PreviouslyAmortizedPretaxGrossAmount *float64 `json:"PreviouslyAmortizedPretaxGrossAmount,omitempty" xml:"PreviouslyAmortizedPretaxGrossAmount,omitempty"`
	// The round-off amount allocated before the current allocation month.
	//
	// example:
	//
	// 0
	PreviouslyAmortizedRoundDownDiscount *float64 `json:"PreviouslyAmortizedRoundDownDiscount,omitempty" xml:"PreviouslyAmortizedRoundDownDiscount,omitempty"`
	// The code of the service.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The specific service resource.
	//
	// example:
	//
	// ApsaraDB RDS
	ProductDetail *string `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	// The code of the specific service resource.
	//
	// example:
	//
	// rds
	ProductDetailCode *string `json:"ProductDetailCode,omitempty" xml:"ProductDetailCode,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// ApsaraDB RDS
	ProductName                              *string  `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	RemainingAmortizationAfterDiscountAmount *float64 `json:"RemainingAmortizationAfterDiscountAmount,omitempty" xml:"RemainingAmortizationAfterDiscountAmount,omitempty"`
	// The amount deducted by using vouchers and to be allocated to one or more future allocation months.
	//
	// example:
	//
	// 0
	RemainingAmortizationDeductedByCashCoupons *float64 `json:"RemainingAmortizationDeductedByCashCoupons,omitempty" xml:"RemainingAmortizationDeductedByCashCoupons,omitempty"`
	// The amount deducted by using coupons and to be allocated to one or more future allocation months.
	//
	// example:
	//
	// 0
	RemainingAmortizationDeductedByCoupons *float64 `json:"RemainingAmortizationDeductedByCoupons,omitempty" xml:"RemainingAmortizationDeductedByCoupons,omitempty"`
	// The amount deducted by using prepaid cards and to be allocated to one or more future allocation months.
	//
	// example:
	//
	// 0
	RemainingAmortizationDeductedByPrepaidCard *float64 `json:"RemainingAmortizationDeductedByPrepaidCard,omitempty" xml:"RemainingAmortizationDeductedByPrepaidCard,omitempty"`
	// The expenditure amount to be allocated to one or more future allocation months. Invoicing is supported.
	//
	// example:
	//
	// 0
	RemainingAmortizationExpenditureAmount *float64 `json:"RemainingAmortizationExpenditureAmount,omitempty" xml:"RemainingAmortizationExpenditureAmount,omitempty"`
	// The discount amount to be allocated to one or more future allocation months.
	//
	// example:
	//
	// 0
	RemainingAmortizationInvoiceDiscount *float64 `json:"RemainingAmortizationInvoiceDiscount,omitempty" xml:"RemainingAmortizationInvoiceDiscount,omitempty"`
	// The pretax amount to be allocated to one or more future allocation months.
	//
	// example:
	//
	// 0
	RemainingAmortizationPretaxAmount *float64 `json:"RemainingAmortizationPretaxAmount,omitempty" xml:"RemainingAmortizationPretaxAmount,omitempty"`
	// The pretax gross amount to be allocated to one or more future allocation months.
	//
	// example:
	//
	// 0
	RemainingAmortizationPretaxGrossAmount *float64 `json:"RemainingAmortizationPretaxGrossAmount,omitempty" xml:"RemainingAmortizationPretaxGrossAmount,omitempty"`
	// The round-off amount to be allocated to one or more future allocation months.
	//
	// example:
	//
	// 0
	RemainingAmortizationRoundDownDiscount *float64 `json:"RemainingAmortizationRoundDownDiscount,omitempty" xml:"RemainingAmortizationRoundDownDiscount,omitempty"`
	// The round-off amount.
	//
	// example:
	//
	// 0
	RoundDownDiscount *float64 `json:"RoundDownDiscount,omitempty" xml:"RoundDownDiscount,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription: the subscription billing method
	//
	// 	- PayAsYouGo: the pay-as-you-go billing method
	//
	// example:
	//
	// PayAsYouGo
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetAfterDiscountAmount() *float64 {
	return s.AfterDiscountAmount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetAmortizationPeriod() *string {
	return s.AmortizationPeriod
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetAmortizationStatus() *string {
	return s.AmortizationStatus
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetBillAccountID() *int64 {
	return s.BillAccountID
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetBillAccountName() *string {
	return s.BillAccountName
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetBillOwnerID() *int64 {
	return s.BillOwnerID
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetBillOwnerName() *string {
	return s.BillOwnerName
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetBizType() *string {
	return s.BizType
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetConsumePeriod() *string {
	return s.ConsumePeriod
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationAfterDiscountAmount() *float64 {
	return s.CurrentAmortizationAfterDiscountAmount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationDeductedByCashCoupons() *float64 {
	return s.CurrentAmortizationDeductedByCashCoupons
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationDeductedByCoupons() *float64 {
	return s.CurrentAmortizationDeductedByCoupons
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationDeductedByPrepaidCard() *float64 {
	return s.CurrentAmortizationDeductedByPrepaidCard
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationExpenditureAmount() *float64 {
	return s.CurrentAmortizationExpenditureAmount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationInvoiceDiscount() *float64 {
	return s.CurrentAmortizationInvoiceDiscount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationPretaxAmount() *float64 {
	return s.CurrentAmortizationPretaxAmount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationPretaxGrossAmount() *float64 {
	return s.CurrentAmortizationPretaxGrossAmount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationRoundDownDiscount() *float64 {
	return s.CurrentAmortizationRoundDownDiscount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetDeductedByCashCoupons() *float64 {
	return s.DeductedByCashCoupons
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetDeductedByCoupons() *float64 {
	return s.DeductedByCoupons
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetDeductedByPrepaidCard() *float64 {
	return s.DeductedByPrepaidCard
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetExpenditureAmount() *float64 {
	return s.ExpenditureAmount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetInvoiceDiscount() *float64 {
	return s.InvoiceDiscount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPretaxAmount() *float64 {
	return s.PretaxAmount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPretaxGrossAmount() *float64 {
	return s.PretaxGrossAmount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPreviouslyAmortizedAfterDiscountAmount() *float64 {
	return s.PreviouslyAmortizedAfterDiscountAmount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPreviouslyAmortizedDeductedByCashCoupons() *float64 {
	return s.PreviouslyAmortizedDeductedByCashCoupons
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPreviouslyAmortizedDeductedByCoupons() *float64 {
	return s.PreviouslyAmortizedDeductedByCoupons
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPreviouslyAmortizedDeductedByPrepaidCard() *float64 {
	return s.PreviouslyAmortizedDeductedByPrepaidCard
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPreviouslyAmortizedExpenditureAmount() *float64 {
	return s.PreviouslyAmortizedExpenditureAmount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPreviouslyAmortizedInvoiceDiscount() *float64 {
	return s.PreviouslyAmortizedInvoiceDiscount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPreviouslyAmortizedPretaxAmount() *float64 {
	return s.PreviouslyAmortizedPretaxAmount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPreviouslyAmortizedPretaxGrossAmount() *float64 {
	return s.PreviouslyAmortizedPretaxGrossAmount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPreviouslyAmortizedRoundDownDiscount() *float64 {
	return s.PreviouslyAmortizedRoundDownDiscount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetProductDetailCode() *string {
	return s.ProductDetailCode
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRemainingAmortizationAfterDiscountAmount() *float64 {
	return s.RemainingAmortizationAfterDiscountAmount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRemainingAmortizationDeductedByCashCoupons() *float64 {
	return s.RemainingAmortizationDeductedByCashCoupons
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRemainingAmortizationDeductedByCoupons() *float64 {
	return s.RemainingAmortizationDeductedByCoupons
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRemainingAmortizationDeductedByPrepaidCard() *float64 {
	return s.RemainingAmortizationDeductedByPrepaidCard
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRemainingAmortizationExpenditureAmount() *float64 {
	return s.RemainingAmortizationExpenditureAmount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRemainingAmortizationInvoiceDiscount() *float64 {
	return s.RemainingAmortizationInvoiceDiscount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRemainingAmortizationPretaxAmount() *float64 {
	return s.RemainingAmortizationPretaxAmount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRemainingAmortizationPretaxGrossAmount() *float64 {
	return s.RemainingAmortizationPretaxGrossAmount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRemainingAmortizationRoundDownDiscount() *float64 {
	return s.RemainingAmortizationRoundDownDiscount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRoundDownDiscount() *float64 {
	return s.RoundDownDiscount
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetAfterDiscountAmount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.AfterDiscountAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetAmortizationPeriod(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.AmortizationPeriod = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetAmortizationStatus(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.AmortizationStatus = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetBillAccountID(v int64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.BillAccountID = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetBillAccountName(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.BillAccountName = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetBillOwnerID(v int64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.BillOwnerID = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetBillOwnerName(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.BillOwnerName = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetBizType(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.BizType = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetConsumePeriod(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ConsumePeriod = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationAfterDiscountAmount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationAfterDiscountAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationDeductedByCashCoupons(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationDeductedByCashCoupons = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationDeductedByCoupons(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationDeductedByCoupons = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationDeductedByPrepaidCard(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationDeductedByPrepaidCard = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationExpenditureAmount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationExpenditureAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationInvoiceDiscount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationInvoiceDiscount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationPretaxAmount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationPretaxAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationPretaxGrossAmount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationPretaxGrossAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationRoundDownDiscount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationRoundDownDiscount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetDeductedByCashCoupons(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetDeductedByCoupons(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.DeductedByCoupons = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetDeductedByPrepaidCard(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetExpenditureAmount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ExpenditureAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetInvoiceDiscount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.InvoiceDiscount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPretaxAmount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PretaxAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPretaxGrossAmount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PretaxGrossAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPreviouslyAmortizedAfterDiscountAmount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PreviouslyAmortizedAfterDiscountAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPreviouslyAmortizedDeductedByCashCoupons(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PreviouslyAmortizedDeductedByCashCoupons = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPreviouslyAmortizedDeductedByCoupons(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PreviouslyAmortizedDeductedByCoupons = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPreviouslyAmortizedDeductedByPrepaidCard(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PreviouslyAmortizedDeductedByPrepaidCard = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPreviouslyAmortizedExpenditureAmount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PreviouslyAmortizedExpenditureAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPreviouslyAmortizedInvoiceDiscount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PreviouslyAmortizedInvoiceDiscount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPreviouslyAmortizedPretaxAmount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PreviouslyAmortizedPretaxAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPreviouslyAmortizedPretaxGrossAmount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PreviouslyAmortizedPretaxGrossAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPreviouslyAmortizedRoundDownDiscount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PreviouslyAmortizedRoundDownDiscount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetProductCode(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetProductDetail(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ProductDetail = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetProductDetailCode(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ProductDetailCode = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetProductName(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ProductName = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRemainingAmortizationAfterDiscountAmount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RemainingAmortizationAfterDiscountAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRemainingAmortizationDeductedByCashCoupons(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RemainingAmortizationDeductedByCashCoupons = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRemainingAmortizationDeductedByCoupons(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RemainingAmortizationDeductedByCoupons = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRemainingAmortizationDeductedByPrepaidCard(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RemainingAmortizationDeductedByPrepaidCard = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRemainingAmortizationExpenditureAmount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RemainingAmortizationExpenditureAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRemainingAmortizationInvoiceDiscount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RemainingAmortizationInvoiceDiscount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRemainingAmortizationPretaxAmount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RemainingAmortizationPretaxAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRemainingAmortizationPretaxGrossAmount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RemainingAmortizationPretaxGrossAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRemainingAmortizationRoundDownDiscount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RemainingAmortizationRoundDownDiscount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRoundDownDiscount(v float64) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RoundDownDiscount = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetSubscriptionType(v string) *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
