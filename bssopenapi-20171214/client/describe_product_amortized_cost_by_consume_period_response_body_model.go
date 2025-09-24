// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductAmortizedCostByConsumePeriodResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBody
	GetCode() *string
	SetData(v *DescribeProductAmortizedCostByConsumePeriodResponseBodyData) *DescribeProductAmortizedCostByConsumePeriodResponseBody
	GetData() *DescribeProductAmortizedCostByConsumePeriodResponseBodyData
	SetMessage(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeProductAmortizedCostByConsumePeriodResponseBody
	GetSuccess() *bool
}

type DescribeProductAmortizedCostByConsumePeriodResponseBody struct {
	// example:
	//
	// 200
	Code *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeProductAmortizedCostByConsumePeriodResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful！
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeProductAmortizedCostByConsumePeriodResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductAmortizedCostByConsumePeriodResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBody) GetData() *DescribeProductAmortizedCostByConsumePeriodResponseBodyData {
	return s.Data
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBody) SetCode(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBody) SetData(v *DescribeProductAmortizedCostByConsumePeriodResponseBodyData) *DescribeProductAmortizedCostByConsumePeriodResponseBody {
	s.Data = v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBody) SetMessage(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBody) SetRequestId(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBody) SetSuccess(v bool) *DescribeProductAmortizedCostByConsumePeriodResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeProductAmortizedCostByConsumePeriodResponseBodyData struct {
	// example:
	//
	// 185766xxxx
	AccountID *string `json:"AccountID,omitempty" xml:"AccountID,omitempty"`
	// example:
	//
	// test@test.aliyunid.com
	AccountName *string                                                             `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Items       []*DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// CAESEgoQCg4KCmdtdF9jcmVhdGUEARgBIkgKCQBwhGmPcAEAAAo7AzYAAAAxTDgwMDcxMjg3ZDJhNmM3ZDguTDgwMDAwMDAwMDAwMzE1MTIuTDgwMDcyZDMyZTJkYzg3N2U
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProductAmortizedCostByConsumePeriodResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductAmortizedCostByConsumePeriodResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyData) GetAccountID() *string {
	return s.AccountID
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyData) GetItems() []*DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	return s.Items
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyData) SetAccountID(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyData) SetAccountName(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyData) SetItems(v []*DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) *DescribeProductAmortizedCostByConsumePeriodResponseBodyData {
	s.Items = v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyData) SetMaxResults(v int32) *DescribeProductAmortizedCostByConsumePeriodResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyData) SetNextToken(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyData) SetTotalCount(v int32) *DescribeProductAmortizedCostByConsumePeriodResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems struct {
	AfterDiscountAmount *float64 `json:"AfterDiscountAmount,omitempty" xml:"AfterDiscountAmount,omitempty"`
	// example:
	//
	// 202210
	AmortizationPeriod *string `json:"AmortizationPeriod,omitempty" xml:"AmortizationPeriod,omitempty"`
	// example:
	//
	// amortized
	AmortizationStatus *string `json:"AmortizationStatus,omitempty" xml:"AmortizationStatus,omitempty"`
	// example:
	//
	// 185xxxxx489
	BillAccountID *int64 `json:"BillAccountID,omitempty" xml:"BillAccountID,omitempty"`
	// example:
	//
	// test@test.aliyunid.com
	BillAccountName *string `json:"BillAccountName,omitempty" xml:"BillAccountName,omitempty"`
	// example:
	//
	// 185xxxxx489
	BillOwnerID *int64 `json:"BillOwnerID,omitempty" xml:"BillOwnerID,omitempty"`
	// example:
	//
	// test@test.aliyunid.com
	BillOwnerName *string `json:"BillOwnerName,omitempty" xml:"BillOwnerName,omitempty"`
	// example:
	//
	// trusteeship
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// 202210
	ConsumePeriod                          *string  `json:"ConsumePeriod,omitempty" xml:"ConsumePeriod,omitempty"`
	CurrentAmortizationAfterDiscountAmount *float64 `json:"CurrentAmortizationAfterDiscountAmount,omitempty" xml:"CurrentAmortizationAfterDiscountAmount,omitempty"`
	// example:
	//
	// 0
	CurrentAmortizationDeductedByCashCoupons *float64 `json:"CurrentAmortizationDeductedByCashCoupons,omitempty" xml:"CurrentAmortizationDeductedByCashCoupons,omitempty"`
	// example:
	//
	// 0
	CurrentAmortizationDeductedByCoupons *float64 `json:"CurrentAmortizationDeductedByCoupons,omitempty" xml:"CurrentAmortizationDeductedByCoupons,omitempty"`
	// example:
	//
	// 0
	CurrentAmortizationDeductedByPrepaidCard *float64 `json:"CurrentAmortizationDeductedByPrepaidCard,omitempty" xml:"CurrentAmortizationDeductedByPrepaidCard,omitempty"`
	// example:
	//
	// 0
	CurrentAmortizationExpenditureAmount *float64 `json:"CurrentAmortizationExpenditureAmount,omitempty" xml:"CurrentAmortizationExpenditureAmount,omitempty"`
	// example:
	//
	// 0
	CurrentAmortizationInvoiceDiscount *float64 `json:"CurrentAmortizationInvoiceDiscount,omitempty" xml:"CurrentAmortizationInvoiceDiscount,omitempty"`
	// example:
	//
	// 0
	CurrentAmortizationPretaxAmount *float64 `json:"CurrentAmortizationPretaxAmount,omitempty" xml:"CurrentAmortizationPretaxAmount,omitempty"`
	// example:
	//
	// 0
	CurrentAmortizationPretaxGrossAmount *float64 `json:"CurrentAmortizationPretaxGrossAmount,omitempty" xml:"CurrentAmortizationPretaxGrossAmount,omitempty"`
	// example:
	//
	// 0
	CurrentAmortizationRoundDownDiscount *float64 `json:"CurrentAmortizationRoundDownDiscount,omitempty" xml:"CurrentAmortizationRoundDownDiscount,omitempty"`
	// example:
	//
	// 0
	DeductedByCashCoupons *float64 `json:"DeductedByCashCoupons,omitempty" xml:"DeductedByCashCoupons,omitempty"`
	// example:
	//
	// 0
	DeductedByCoupons *float64 `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	// example:
	//
	// 0
	DeductedByPrepaidCard *float64 `json:"DeductedByPrepaidCard,omitempty" xml:"DeductedByPrepaidCard,omitempty"`
	// example:
	//
	// 0
	ExpenditureAmount *float64 `json:"ExpenditureAmount,omitempty" xml:"ExpenditureAmount,omitempty"`
	// example:
	//
	// 0
	InvoiceDiscount *float64 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	// example:
	//
	// 0
	PretaxAmount *float64 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	// example:
	//
	// 0
	PretaxGrossAmount                      *float64 `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	PreviouslyAmortizedAfterDiscountAmount *float64 `json:"PreviouslyAmortizedAfterDiscountAmount,omitempty" xml:"PreviouslyAmortizedAfterDiscountAmount,omitempty"`
	// example:
	//
	// 0
	PreviouslyAmortizedDeductedByCashCoupons *float64 `json:"PreviouslyAmortizedDeductedByCashCoupons,omitempty" xml:"PreviouslyAmortizedDeductedByCashCoupons,omitempty"`
	// example:
	//
	// 0
	PreviouslyAmortizedDeductedByCoupons *float64 `json:"PreviouslyAmortizedDeductedByCoupons,omitempty" xml:"PreviouslyAmortizedDeductedByCoupons,omitempty"`
	// example:
	//
	// 0
	PreviouslyAmortizedDeductedByPrepaidCard *float64 `json:"PreviouslyAmortizedDeductedByPrepaidCard,omitempty" xml:"PreviouslyAmortizedDeductedByPrepaidCard,omitempty"`
	// example:
	//
	// 0
	PreviouslyAmortizedExpenditureAmount *float64 `json:"PreviouslyAmortizedExpenditureAmount,omitempty" xml:"PreviouslyAmortizedExpenditureAmount,omitempty"`
	// example:
	//
	// 0
	PreviouslyAmortizedInvoiceDiscount *float64 `json:"PreviouslyAmortizedInvoiceDiscount,omitempty" xml:"PreviouslyAmortizedInvoiceDiscount,omitempty"`
	// example:
	//
	// 0
	PreviouslyAmortizedPretaxAmount *float64 `json:"PreviouslyAmortizedPretaxAmount,omitempty" xml:"PreviouslyAmortizedPretaxAmount,omitempty"`
	// example:
	//
	// 0
	PreviouslyAmortizedPretaxGrossAmount *float64 `json:"PreviouslyAmortizedPretaxGrossAmount,omitempty" xml:"PreviouslyAmortizedPretaxGrossAmount,omitempty"`
	// example:
	//
	// 0
	PreviouslyAmortizedRoundDownDiscount *float64 `json:"PreviouslyAmortizedRoundDownDiscount,omitempty" xml:"PreviouslyAmortizedRoundDownDiscount,omitempty"`
	// example:
	//
	// rds
	ProductCode   *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductDetail *string `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	// example:
	//
	// rds
	ProductDetailCode                        *string  `json:"ProductDetailCode,omitempty" xml:"ProductDetailCode,omitempty"`
	ProductName                              *string  `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	RemainingAmortizationAfterDiscountAmount *float64 `json:"RemainingAmortizationAfterDiscountAmount,omitempty" xml:"RemainingAmortizationAfterDiscountAmount,omitempty"`
	// example:
	//
	// 0
	RemainingAmortizationDeductedByCashCoupons *float64 `json:"RemainingAmortizationDeductedByCashCoupons,omitempty" xml:"RemainingAmortizationDeductedByCashCoupons,omitempty"`
	// example:
	//
	// 0
	RemainingAmortizationDeductedByCoupons *float64 `json:"RemainingAmortizationDeductedByCoupons,omitempty" xml:"RemainingAmortizationDeductedByCoupons,omitempty"`
	// example:
	//
	// 0
	RemainingAmortizationDeductedByPrepaidCard *float64 `json:"RemainingAmortizationDeductedByPrepaidCard,omitempty" xml:"RemainingAmortizationDeductedByPrepaidCard,omitempty"`
	// example:
	//
	// 0
	RemainingAmortizationExpenditureAmount *float64 `json:"RemainingAmortizationExpenditureAmount,omitempty" xml:"RemainingAmortizationExpenditureAmount,omitempty"`
	// example:
	//
	// 0
	RemainingAmortizationInvoiceDiscount *float64 `json:"RemainingAmortizationInvoiceDiscount,omitempty" xml:"RemainingAmortizationInvoiceDiscount,omitempty"`
	// example:
	//
	// 0
	RemainingAmortizationPretaxAmount *float64 `json:"RemainingAmortizationPretaxAmount,omitempty" xml:"RemainingAmortizationPretaxAmount,omitempty"`
	// example:
	//
	// 0
	RemainingAmortizationPretaxGrossAmount *float64 `json:"RemainingAmortizationPretaxGrossAmount,omitempty" xml:"RemainingAmortizationPretaxGrossAmount,omitempty"`
	// example:
	//
	// 0
	RemainingAmortizationRoundDownDiscount *float64 `json:"RemainingAmortizationRoundDownDiscount,omitempty" xml:"RemainingAmortizationRoundDownDiscount,omitempty"`
	// example:
	//
	// 0
	RoundDownDiscount *float64 `json:"RoundDownDiscount,omitempty" xml:"RoundDownDiscount,omitempty"`
	// example:
	//
	// PayAsYouGo
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetAfterDiscountAmount() *float64 {
	return s.AfterDiscountAmount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetAmortizationPeriod() *string {
	return s.AmortizationPeriod
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetAmortizationStatus() *string {
	return s.AmortizationStatus
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetBillAccountID() *int64 {
	return s.BillAccountID
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetBillAccountName() *string {
	return s.BillAccountName
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetBillOwnerID() *int64 {
	return s.BillOwnerID
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetBillOwnerName() *string {
	return s.BillOwnerName
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetBizType() *string {
	return s.BizType
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetConsumePeriod() *string {
	return s.ConsumePeriod
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetCurrentAmortizationAfterDiscountAmount() *float64 {
	return s.CurrentAmortizationAfterDiscountAmount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetCurrentAmortizationDeductedByCashCoupons() *float64 {
	return s.CurrentAmortizationDeductedByCashCoupons
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetCurrentAmortizationDeductedByCoupons() *float64 {
	return s.CurrentAmortizationDeductedByCoupons
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetCurrentAmortizationDeductedByPrepaidCard() *float64 {
	return s.CurrentAmortizationDeductedByPrepaidCard
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetCurrentAmortizationExpenditureAmount() *float64 {
	return s.CurrentAmortizationExpenditureAmount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetCurrentAmortizationInvoiceDiscount() *float64 {
	return s.CurrentAmortizationInvoiceDiscount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetCurrentAmortizationPretaxAmount() *float64 {
	return s.CurrentAmortizationPretaxAmount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetCurrentAmortizationPretaxGrossAmount() *float64 {
	return s.CurrentAmortizationPretaxGrossAmount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetCurrentAmortizationRoundDownDiscount() *float64 {
	return s.CurrentAmortizationRoundDownDiscount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetDeductedByCashCoupons() *float64 {
	return s.DeductedByCashCoupons
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetDeductedByCoupons() *float64 {
	return s.DeductedByCoupons
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetDeductedByPrepaidCard() *float64 {
	return s.DeductedByPrepaidCard
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetExpenditureAmount() *float64 {
	return s.ExpenditureAmount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetInvoiceDiscount() *float64 {
	return s.InvoiceDiscount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetPretaxAmount() *float64 {
	return s.PretaxAmount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetPretaxGrossAmount() *float64 {
	return s.PretaxGrossAmount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetPreviouslyAmortizedAfterDiscountAmount() *float64 {
	return s.PreviouslyAmortizedAfterDiscountAmount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetPreviouslyAmortizedDeductedByCashCoupons() *float64 {
	return s.PreviouslyAmortizedDeductedByCashCoupons
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetPreviouslyAmortizedDeductedByCoupons() *float64 {
	return s.PreviouslyAmortizedDeductedByCoupons
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetPreviouslyAmortizedDeductedByPrepaidCard() *float64 {
	return s.PreviouslyAmortizedDeductedByPrepaidCard
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetPreviouslyAmortizedExpenditureAmount() *float64 {
	return s.PreviouslyAmortizedExpenditureAmount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetPreviouslyAmortizedInvoiceDiscount() *float64 {
	return s.PreviouslyAmortizedInvoiceDiscount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetPreviouslyAmortizedPretaxAmount() *float64 {
	return s.PreviouslyAmortizedPretaxAmount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetPreviouslyAmortizedPretaxGrossAmount() *float64 {
	return s.PreviouslyAmortizedPretaxGrossAmount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetPreviouslyAmortizedRoundDownDiscount() *float64 {
	return s.PreviouslyAmortizedRoundDownDiscount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetProductDetailCode() *string {
	return s.ProductDetailCode
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetRemainingAmortizationAfterDiscountAmount() *float64 {
	return s.RemainingAmortizationAfterDiscountAmount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetRemainingAmortizationDeductedByCashCoupons() *float64 {
	return s.RemainingAmortizationDeductedByCashCoupons
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetRemainingAmortizationDeductedByCoupons() *float64 {
	return s.RemainingAmortizationDeductedByCoupons
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetRemainingAmortizationDeductedByPrepaidCard() *float64 {
	return s.RemainingAmortizationDeductedByPrepaidCard
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetRemainingAmortizationExpenditureAmount() *float64 {
	return s.RemainingAmortizationExpenditureAmount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetRemainingAmortizationInvoiceDiscount() *float64 {
	return s.RemainingAmortizationInvoiceDiscount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetRemainingAmortizationPretaxAmount() *float64 {
	return s.RemainingAmortizationPretaxAmount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetRemainingAmortizationPretaxGrossAmount() *float64 {
	return s.RemainingAmortizationPretaxGrossAmount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetRemainingAmortizationRoundDownDiscount() *float64 {
	return s.RemainingAmortizationRoundDownDiscount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetRoundDownDiscount() *float64 {
	return s.RoundDownDiscount
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetAfterDiscountAmount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.AfterDiscountAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetAmortizationPeriod(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.AmortizationPeriod = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetAmortizationStatus(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.AmortizationStatus = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetBillAccountID(v int64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.BillAccountID = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetBillAccountName(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.BillAccountName = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetBillOwnerID(v int64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.BillOwnerID = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetBillOwnerName(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.BillOwnerName = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetBizType(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.BizType = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetConsumePeriod(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.ConsumePeriod = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetCurrentAmortizationAfterDiscountAmount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CurrentAmortizationAfterDiscountAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetCurrentAmortizationDeductedByCashCoupons(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CurrentAmortizationDeductedByCashCoupons = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetCurrentAmortizationDeductedByCoupons(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CurrentAmortizationDeductedByCoupons = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetCurrentAmortizationDeductedByPrepaidCard(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CurrentAmortizationDeductedByPrepaidCard = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetCurrentAmortizationExpenditureAmount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CurrentAmortizationExpenditureAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetCurrentAmortizationInvoiceDiscount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CurrentAmortizationInvoiceDiscount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetCurrentAmortizationPretaxAmount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CurrentAmortizationPretaxAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetCurrentAmortizationPretaxGrossAmount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CurrentAmortizationPretaxGrossAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetCurrentAmortizationRoundDownDiscount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CurrentAmortizationRoundDownDiscount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetDeductedByCashCoupons(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetDeductedByCoupons(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.DeductedByCoupons = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetDeductedByPrepaidCard(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetExpenditureAmount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.ExpenditureAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetInvoiceDiscount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.InvoiceDiscount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetPretaxAmount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PretaxAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetPretaxGrossAmount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PretaxGrossAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetPreviouslyAmortizedAfterDiscountAmount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PreviouslyAmortizedAfterDiscountAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetPreviouslyAmortizedDeductedByCashCoupons(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PreviouslyAmortizedDeductedByCashCoupons = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetPreviouslyAmortizedDeductedByCoupons(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PreviouslyAmortizedDeductedByCoupons = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetPreviouslyAmortizedDeductedByPrepaidCard(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PreviouslyAmortizedDeductedByPrepaidCard = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetPreviouslyAmortizedExpenditureAmount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PreviouslyAmortizedExpenditureAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetPreviouslyAmortizedInvoiceDiscount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PreviouslyAmortizedInvoiceDiscount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetPreviouslyAmortizedPretaxAmount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PreviouslyAmortizedPretaxAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetPreviouslyAmortizedPretaxGrossAmount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PreviouslyAmortizedPretaxGrossAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetPreviouslyAmortizedRoundDownDiscount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PreviouslyAmortizedRoundDownDiscount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetProductCode(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetProductDetail(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.ProductDetail = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetProductDetailCode(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.ProductDetailCode = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetProductName(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.ProductName = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetRemainingAmortizationAfterDiscountAmount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RemainingAmortizationAfterDiscountAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetRemainingAmortizationDeductedByCashCoupons(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RemainingAmortizationDeductedByCashCoupons = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetRemainingAmortizationDeductedByCoupons(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RemainingAmortizationDeductedByCoupons = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetRemainingAmortizationDeductedByPrepaidCard(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RemainingAmortizationDeductedByPrepaidCard = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetRemainingAmortizationExpenditureAmount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RemainingAmortizationExpenditureAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetRemainingAmortizationInvoiceDiscount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RemainingAmortizationInvoiceDiscount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetRemainingAmortizationPretaxAmount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RemainingAmortizationPretaxAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetRemainingAmortizationPretaxGrossAmount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RemainingAmortizationPretaxGrossAmount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetRemainingAmortizationRoundDownDiscount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RemainingAmortizationRoundDownDiscount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetRoundDownDiscount(v float64) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RoundDownDiscount = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) SetSubscriptionType(v string) *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
