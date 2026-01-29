// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAmortizedCostByConsumePeriodResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBody
	GetCode() *string
	SetData(v *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData) *DescribeInstanceAmortizedCostByConsumePeriodResponseBody
	GetData() *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData
	SetMessage(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstanceAmortizedCostByConsumePeriodResponseBody
	GetSuccess() *bool
}

type DescribeInstanceAmortizedCostByConsumePeriodResponseBody struct {
	// example:
	//
	// 200
	Code *string                                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DescribeInstanceAmortizedCostByConsumePeriodResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAmortizedCostByConsumePeriodResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBody) GetData() *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData {
	return s.Data
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBody) SetCode(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBody) SetData(v *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData) *DescribeInstanceAmortizedCostByConsumePeriodResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBody) SetMessage(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBody) SetRequestId(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBody) SetSuccess(v bool) *DescribeInstanceAmortizedCostByConsumePeriodResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData struct {
	// example:
	//
	// 185766xxxx
	AccountID *string `json:"AccountID,omitempty" xml:"AccountID,omitempty"`
	// example:
	//
	// test@test.aliyunid.com
	AccountName *string                                                              `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Items       []*DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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

func (s DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData) GetAccountID() *string {
	return s.AccountID
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData) GetItems() []*DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	return s.Items
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData) SetAccountID(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData) SetAccountName(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData) SetItems(v []*DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData {
	s.Items = v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData) SetMaxResults(v int32) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData) SetNextToken(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData) SetTotalCount(v int32) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems struct {
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
	ConsumePeriod *string `json:"ConsumePeriod,omitempty" xml:"ConsumePeriod,omitempty"`
	CostUnit      *string `json:"CostUnit,omitempty" xml:"CostUnit,omitempty"`
	// example:
	//
	// 1234
	CostUnitCode                           *string  `json:"CostUnitCode,omitempty" xml:"CostUnitCode,omitempty"`
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
	// i-kjhdskjgshfdlkjfdh
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// example:
	//
	// 34.xx.x.x
	InternetIP *string `json:"InternetIP,omitempty" xml:"InternetIP,omitempty"`
	// example:
	//
	// 192.xx.xx.xx
	IntranetIP *string `json:"IntranetIP,omitempty" xml:"IntranetIP,omitempty"`
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
	Region                                   *string  `json:"Region,omitempty" xml:"Region,omitempty"`
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
	ResourceGroup                          *string  `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// example:
	//
	// 0
	RoundDownDiscount *float64 `json:"RoundDownDiscount,omitempty" xml:"RoundDownDiscount,omitempty"`
	// example:
	//
	// test**1122
	SplitAccountName *string `json:"SplitAccountName,omitempty" xml:"SplitAccountName,omitempty"`
	// example:
	//
	// i-28bycvyb4
	SplitItemID *string `json:"SplitItemID,omitempty" xml:"SplitItemID,omitempty"`
	// example:
	//
	// iZ28bycvyb4Z
	SplitItemName *string `json:"SplitItemName,omitempty" xml:"SplitItemName,omitempty"`
	// example:
	//
	// rds
	SplitProductDetail *string `json:"SplitProductDetail,omitempty" xml:"SplitProductDetail,omitempty"`
	// example:
	//
	// PayAsYouGo
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	// example:
	//
	// tag
	Tag  *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetAfterDiscountAmount() *float64 {
	return s.AfterDiscountAmount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetAmortizationPeriod() *string {
	return s.AmortizationPeriod
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetAmortizationStatus() *string {
	return s.AmortizationStatus
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetBillAccountID() *int64 {
	return s.BillAccountID
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetBillAccountName() *string {
	return s.BillAccountName
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetBillOwnerID() *int64 {
	return s.BillOwnerID
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetBillOwnerName() *string {
	return s.BillOwnerName
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetBizType() *string {
	return s.BizType
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetConsumePeriod() *string {
	return s.ConsumePeriod
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetCostUnit() *string {
	return s.CostUnit
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetCostUnitCode() *string {
	return s.CostUnitCode
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetCurrentAmortizationAfterDiscountAmount() *float64 {
	return s.CurrentAmortizationAfterDiscountAmount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetCurrentAmortizationDeductedByCashCoupons() *float64 {
	return s.CurrentAmortizationDeductedByCashCoupons
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetCurrentAmortizationDeductedByCoupons() *float64 {
	return s.CurrentAmortizationDeductedByCoupons
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetCurrentAmortizationDeductedByPrepaidCard() *float64 {
	return s.CurrentAmortizationDeductedByPrepaidCard
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetCurrentAmortizationExpenditureAmount() *float64 {
	return s.CurrentAmortizationExpenditureAmount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetCurrentAmortizationInvoiceDiscount() *float64 {
	return s.CurrentAmortizationInvoiceDiscount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetCurrentAmortizationPretaxAmount() *float64 {
	return s.CurrentAmortizationPretaxAmount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetCurrentAmortizationPretaxGrossAmount() *float64 {
	return s.CurrentAmortizationPretaxGrossAmount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetCurrentAmortizationRoundDownDiscount() *float64 {
	return s.CurrentAmortizationRoundDownDiscount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetDeductedByCashCoupons() *float64 {
	return s.DeductedByCashCoupons
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetDeductedByCoupons() *float64 {
	return s.DeductedByCoupons
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetDeductedByPrepaidCard() *float64 {
	return s.DeductedByPrepaidCard
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetExpenditureAmount() *float64 {
	return s.ExpenditureAmount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetInternetIP() *string {
	return s.InternetIP
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetIntranetIP() *string {
	return s.IntranetIP
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetInvoiceDiscount() *float64 {
	return s.InvoiceDiscount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetPretaxAmount() *float64 {
	return s.PretaxAmount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetPretaxGrossAmount() *float64 {
	return s.PretaxGrossAmount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetPreviouslyAmortizedAfterDiscountAmount() *float64 {
	return s.PreviouslyAmortizedAfterDiscountAmount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetPreviouslyAmortizedDeductedByCashCoupons() *float64 {
	return s.PreviouslyAmortizedDeductedByCashCoupons
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetPreviouslyAmortizedDeductedByCoupons() *float64 {
	return s.PreviouslyAmortizedDeductedByCoupons
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetPreviouslyAmortizedDeductedByPrepaidCard() *float64 {
	return s.PreviouslyAmortizedDeductedByPrepaidCard
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetPreviouslyAmortizedExpenditureAmount() *float64 {
	return s.PreviouslyAmortizedExpenditureAmount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetPreviouslyAmortizedInvoiceDiscount() *float64 {
	return s.PreviouslyAmortizedInvoiceDiscount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetPreviouslyAmortizedPretaxAmount() *float64 {
	return s.PreviouslyAmortizedPretaxAmount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetPreviouslyAmortizedPretaxGrossAmount() *float64 {
	return s.PreviouslyAmortizedPretaxGrossAmount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetPreviouslyAmortizedRoundDownDiscount() *float64 {
	return s.PreviouslyAmortizedRoundDownDiscount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetProductDetailCode() *string {
	return s.ProductDetailCode
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetRegion() *string {
	return s.Region
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetRemainingAmortizationAfterDiscountAmount() *float64 {
	return s.RemainingAmortizationAfterDiscountAmount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetRemainingAmortizationDeductedByCashCoupons() *float64 {
	return s.RemainingAmortizationDeductedByCashCoupons
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetRemainingAmortizationDeductedByCoupons() *float64 {
	return s.RemainingAmortizationDeductedByCoupons
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetRemainingAmortizationDeductedByPrepaidCard() *float64 {
	return s.RemainingAmortizationDeductedByPrepaidCard
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetRemainingAmortizationExpenditureAmount() *float64 {
	return s.RemainingAmortizationExpenditureAmount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetRemainingAmortizationInvoiceDiscount() *float64 {
	return s.RemainingAmortizationInvoiceDiscount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetRemainingAmortizationPretaxAmount() *float64 {
	return s.RemainingAmortizationPretaxAmount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetRemainingAmortizationPretaxGrossAmount() *float64 {
	return s.RemainingAmortizationPretaxGrossAmount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetRemainingAmortizationRoundDownDiscount() *float64 {
	return s.RemainingAmortizationRoundDownDiscount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetRoundDownDiscount() *float64 {
	return s.RoundDownDiscount
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetSplitAccountName() *string {
	return s.SplitAccountName
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetSplitItemID() *string {
	return s.SplitItemID
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetSplitItemName() *string {
	return s.SplitItemName
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetSplitProductDetail() *string {
	return s.SplitProductDetail
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetTag() *string {
	return s.Tag
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) GetZone() *string {
	return s.Zone
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetAfterDiscountAmount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.AfterDiscountAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetAmortizationPeriod(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.AmortizationPeriod = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetAmortizationStatus(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.AmortizationStatus = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetBillAccountID(v int64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.BillAccountID = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetBillAccountName(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.BillAccountName = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetBillOwnerID(v int64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.BillOwnerID = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetBillOwnerName(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.BillOwnerName = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetBizType(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.BizType = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetConsumePeriod(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.ConsumePeriod = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetCostUnit(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CostUnit = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetCostUnitCode(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CostUnitCode = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetCurrentAmortizationAfterDiscountAmount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CurrentAmortizationAfterDiscountAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetCurrentAmortizationDeductedByCashCoupons(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CurrentAmortizationDeductedByCashCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetCurrentAmortizationDeductedByCoupons(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CurrentAmortizationDeductedByCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetCurrentAmortizationDeductedByPrepaidCard(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CurrentAmortizationDeductedByPrepaidCard = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetCurrentAmortizationExpenditureAmount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CurrentAmortizationExpenditureAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetCurrentAmortizationInvoiceDiscount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CurrentAmortizationInvoiceDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetCurrentAmortizationPretaxAmount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CurrentAmortizationPretaxAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetCurrentAmortizationPretaxGrossAmount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CurrentAmortizationPretaxGrossAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetCurrentAmortizationRoundDownDiscount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.CurrentAmortizationRoundDownDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetDeductedByCashCoupons(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetDeductedByCoupons(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.DeductedByCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetDeductedByPrepaidCard(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetExpenditureAmount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.ExpenditureAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetInstanceID(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.InstanceID = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetInternetIP(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.InternetIP = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetIntranetIP(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.IntranetIP = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetInvoiceDiscount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.InvoiceDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetPretaxAmount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PretaxAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetPretaxGrossAmount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PretaxGrossAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetPreviouslyAmortizedAfterDiscountAmount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PreviouslyAmortizedAfterDiscountAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetPreviouslyAmortizedDeductedByCashCoupons(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PreviouslyAmortizedDeductedByCashCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetPreviouslyAmortizedDeductedByCoupons(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PreviouslyAmortizedDeductedByCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetPreviouslyAmortizedDeductedByPrepaidCard(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PreviouslyAmortizedDeductedByPrepaidCard = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetPreviouslyAmortizedExpenditureAmount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PreviouslyAmortizedExpenditureAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetPreviouslyAmortizedInvoiceDiscount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PreviouslyAmortizedInvoiceDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetPreviouslyAmortizedPretaxAmount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PreviouslyAmortizedPretaxAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetPreviouslyAmortizedPretaxGrossAmount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PreviouslyAmortizedPretaxGrossAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetPreviouslyAmortizedRoundDownDiscount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.PreviouslyAmortizedRoundDownDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetProductCode(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetProductDetail(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.ProductDetail = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetProductDetailCode(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.ProductDetailCode = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetProductName(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.ProductName = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetRegion(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.Region = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetRemainingAmortizationAfterDiscountAmount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RemainingAmortizationAfterDiscountAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetRemainingAmortizationDeductedByCashCoupons(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RemainingAmortizationDeductedByCashCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetRemainingAmortizationDeductedByCoupons(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RemainingAmortizationDeductedByCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetRemainingAmortizationDeductedByPrepaidCard(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RemainingAmortizationDeductedByPrepaidCard = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetRemainingAmortizationExpenditureAmount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RemainingAmortizationExpenditureAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetRemainingAmortizationInvoiceDiscount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RemainingAmortizationInvoiceDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetRemainingAmortizationPretaxAmount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RemainingAmortizationPretaxAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetRemainingAmortizationPretaxGrossAmount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RemainingAmortizationPretaxGrossAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetRemainingAmortizationRoundDownDiscount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RemainingAmortizationRoundDownDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetResourceGroup(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetRoundDownDiscount(v float64) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.RoundDownDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetSplitAccountName(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.SplitAccountName = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetSplitItemID(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.SplitItemID = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetSplitItemName(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.SplitItemName = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetSplitProductDetail(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.SplitProductDetail = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetSubscriptionType(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetTag(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.Tag = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) SetZone(v string) *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems {
	s.Zone = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
