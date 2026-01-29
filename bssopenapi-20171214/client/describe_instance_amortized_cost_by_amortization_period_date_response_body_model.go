// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody
	GetCode() *string
	SetData(v *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody
	GetData() *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData
	SetMessage(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody
	GetSuccess() *bool
}

type DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody struct {
	// example:
	//
	// 200
	Code *string                                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody) GetData() *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData {
	return s.Data
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody) SetCode(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody) SetData(v *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody) SetMessage(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody) SetRequestId(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody) SetSuccess(v bool) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData struct {
	// example:
	//
	// 185766xxxx
	AccountID *string `json:"AccountID,omitempty" xml:"AccountID,omitempty"`
	// example:
	//
	// test@test.aliyunid.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// cn
	HostId *string                                                                       `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Items  []*DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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

func (s DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData) GetAccountID() *string {
	return s.AccountID
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData) GetHostId() *string {
	return s.HostId
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData) GetItems() []*DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	return s.Items
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData) SetAccountID(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData) SetAccountName(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData) SetHostId(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData {
	s.HostId = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData) SetItems(v []*DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData {
	s.Items = v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData) SetMaxResults(v int32) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData) SetNextToken(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData) SetTotalCount(v int32) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyData) Validate() error {
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

type DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems struct {
	// example:
	//
	// 2022-10
	AmortizationPeriod *string `json:"AmortizationPeriod,omitempty" xml:"AmortizationPeriod,omitempty"`
	// example:
	//
	// 2022-10-10
	AmortizationPeriodDay *string `json:"AmortizationPeriodDay,omitempty" xml:"AmortizationPeriodDay,omitempty"`
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
	// 2022-10
	ConsumePeriod *string `json:"ConsumePeriod,omitempty" xml:"ConsumePeriod,omitempty"`
	// example:
	//
	// 2022-10-10
	ConsumePeriodDay *string `json:"ConsumePeriodDay,omitempty" xml:"ConsumePeriodDay,omitempty"`
	CostUnit         *string `json:"CostUnit,omitempty" xml:"CostUnit,omitempty"`
	// example:
	//
	// 1234
	CostUnitCode *string `json:"CostUnitCode,omitempty" xml:"CostUnitCode,omitempty"`
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
	PretaxGrossAmount *float64 `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
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
	ProductDetailCode *string `json:"ProductDetailCode,omitempty" xml:"ProductDetailCode,omitempty"`
	ProductName       *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	Region            *string `json:"Region,omitempty" xml:"Region,omitempty"`
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

func (s DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetAmortizationPeriod() *string {
	return s.AmortizationPeriod
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetAmortizationPeriodDay() *string {
	return s.AmortizationPeriodDay
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetAmortizationStatus() *string {
	return s.AmortizationStatus
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetBillAccountID() *int64 {
	return s.BillAccountID
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetBillAccountName() *string {
	return s.BillAccountName
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetBillOwnerID() *int64 {
	return s.BillOwnerID
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetBillOwnerName() *string {
	return s.BillOwnerName
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetBizType() *string {
	return s.BizType
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetConsumePeriod() *string {
	return s.ConsumePeriod
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetConsumePeriodDay() *string {
	return s.ConsumePeriodDay
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetCostUnit() *string {
	return s.CostUnit
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetCostUnitCode() *string {
	return s.CostUnitCode
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetCurrentAmortizationDeductedByCashCoupons() *float64 {
	return s.CurrentAmortizationDeductedByCashCoupons
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetCurrentAmortizationDeductedByCoupons() *float64 {
	return s.CurrentAmortizationDeductedByCoupons
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetCurrentAmortizationDeductedByPrepaidCard() *float64 {
	return s.CurrentAmortizationDeductedByPrepaidCard
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetCurrentAmortizationExpenditureAmount() *float64 {
	return s.CurrentAmortizationExpenditureAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetCurrentAmortizationInvoiceDiscount() *float64 {
	return s.CurrentAmortizationInvoiceDiscount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetCurrentAmortizationPretaxAmount() *float64 {
	return s.CurrentAmortizationPretaxAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetCurrentAmortizationPretaxGrossAmount() *float64 {
	return s.CurrentAmortizationPretaxGrossAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetCurrentAmortizationRoundDownDiscount() *float64 {
	return s.CurrentAmortizationRoundDownDiscount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetDeductedByCashCoupons() *float64 {
	return s.DeductedByCashCoupons
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetDeductedByCoupons() *float64 {
	return s.DeductedByCoupons
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetDeductedByPrepaidCard() *float64 {
	return s.DeductedByPrepaidCard
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetExpenditureAmount() *float64 {
	return s.ExpenditureAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetInternetIP() *string {
	return s.InternetIP
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetIntranetIP() *string {
	return s.IntranetIP
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetInvoiceDiscount() *float64 {
	return s.InvoiceDiscount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetPretaxAmount() *float64 {
	return s.PretaxAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetPretaxGrossAmount() *float64 {
	return s.PretaxGrossAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetPreviouslyAmortizedDeductedByCashCoupons() *float64 {
	return s.PreviouslyAmortizedDeductedByCashCoupons
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetPreviouslyAmortizedDeductedByCoupons() *float64 {
	return s.PreviouslyAmortizedDeductedByCoupons
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetPreviouslyAmortizedDeductedByPrepaidCard() *float64 {
	return s.PreviouslyAmortizedDeductedByPrepaidCard
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetPreviouslyAmortizedExpenditureAmount() *float64 {
	return s.PreviouslyAmortizedExpenditureAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetPreviouslyAmortizedInvoiceDiscount() *float64 {
	return s.PreviouslyAmortizedInvoiceDiscount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetPreviouslyAmortizedPretaxAmount() *float64 {
	return s.PreviouslyAmortizedPretaxAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetPreviouslyAmortizedPretaxGrossAmount() *float64 {
	return s.PreviouslyAmortizedPretaxGrossAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetPreviouslyAmortizedRoundDownDiscount() *float64 {
	return s.PreviouslyAmortizedRoundDownDiscount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetProductDetailCode() *string {
	return s.ProductDetailCode
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetRegion() *string {
	return s.Region
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetRemainingAmortizationDeductedByCashCoupons() *float64 {
	return s.RemainingAmortizationDeductedByCashCoupons
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetRemainingAmortizationDeductedByCoupons() *float64 {
	return s.RemainingAmortizationDeductedByCoupons
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetRemainingAmortizationDeductedByPrepaidCard() *float64 {
	return s.RemainingAmortizationDeductedByPrepaidCard
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetRemainingAmortizationExpenditureAmount() *float64 {
	return s.RemainingAmortizationExpenditureAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetRemainingAmortizationInvoiceDiscount() *float64 {
	return s.RemainingAmortizationInvoiceDiscount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetRemainingAmortizationPretaxAmount() *float64 {
	return s.RemainingAmortizationPretaxAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetRemainingAmortizationPretaxGrossAmount() *float64 {
	return s.RemainingAmortizationPretaxGrossAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetRemainingAmortizationRoundDownDiscount() *float64 {
	return s.RemainingAmortizationRoundDownDiscount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetRoundDownDiscount() *float64 {
	return s.RoundDownDiscount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetSplitAccountName() *string {
	return s.SplitAccountName
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetSplitItemID() *string {
	return s.SplitItemID
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetSplitItemName() *string {
	return s.SplitItemName
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetSplitProductDetail() *string {
	return s.SplitProductDetail
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetTag() *string {
	return s.Tag
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) GetZone() *string {
	return s.Zone
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetAmortizationPeriod(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.AmortizationPeriod = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetAmortizationPeriodDay(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.AmortizationPeriodDay = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetAmortizationStatus(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.AmortizationStatus = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetBillAccountID(v int64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.BillAccountID = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetBillAccountName(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.BillAccountName = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetBillOwnerID(v int64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.BillOwnerID = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetBillOwnerName(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.BillOwnerName = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetBizType(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.BizType = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetConsumePeriod(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.ConsumePeriod = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetConsumePeriodDay(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.ConsumePeriodDay = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetCostUnit(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.CostUnit = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetCostUnitCode(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.CostUnitCode = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetCurrentAmortizationDeductedByCashCoupons(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.CurrentAmortizationDeductedByCashCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetCurrentAmortizationDeductedByCoupons(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.CurrentAmortizationDeductedByCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetCurrentAmortizationDeductedByPrepaidCard(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.CurrentAmortizationDeductedByPrepaidCard = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetCurrentAmortizationExpenditureAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.CurrentAmortizationExpenditureAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetCurrentAmortizationInvoiceDiscount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.CurrentAmortizationInvoiceDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetCurrentAmortizationPretaxAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.CurrentAmortizationPretaxAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetCurrentAmortizationPretaxGrossAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.CurrentAmortizationPretaxGrossAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetCurrentAmortizationRoundDownDiscount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.CurrentAmortizationRoundDownDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetDeductedByCashCoupons(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetDeductedByCoupons(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.DeductedByCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetDeductedByPrepaidCard(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetExpenditureAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.ExpenditureAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetInstanceID(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.InstanceID = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetInternetIP(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.InternetIP = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetIntranetIP(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.IntranetIP = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetInvoiceDiscount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.InvoiceDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetPretaxAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.PretaxAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetPretaxGrossAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.PretaxGrossAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetPreviouslyAmortizedDeductedByCashCoupons(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.PreviouslyAmortizedDeductedByCashCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetPreviouslyAmortizedDeductedByCoupons(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.PreviouslyAmortizedDeductedByCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetPreviouslyAmortizedDeductedByPrepaidCard(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.PreviouslyAmortizedDeductedByPrepaidCard = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetPreviouslyAmortizedExpenditureAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.PreviouslyAmortizedExpenditureAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetPreviouslyAmortizedInvoiceDiscount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.PreviouslyAmortizedInvoiceDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetPreviouslyAmortizedPretaxAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.PreviouslyAmortizedPretaxAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetPreviouslyAmortizedPretaxGrossAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.PreviouslyAmortizedPretaxGrossAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetPreviouslyAmortizedRoundDownDiscount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.PreviouslyAmortizedRoundDownDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetProductCode(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetProductDetail(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.ProductDetail = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetProductDetailCode(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.ProductDetailCode = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetProductName(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.ProductName = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetRegion(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.Region = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetRemainingAmortizationDeductedByCashCoupons(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.RemainingAmortizationDeductedByCashCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetRemainingAmortizationDeductedByCoupons(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.RemainingAmortizationDeductedByCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetRemainingAmortizationDeductedByPrepaidCard(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.RemainingAmortizationDeductedByPrepaidCard = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetRemainingAmortizationExpenditureAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.RemainingAmortizationExpenditureAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetRemainingAmortizationInvoiceDiscount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.RemainingAmortizationInvoiceDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetRemainingAmortizationPretaxAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.RemainingAmortizationPretaxAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetRemainingAmortizationPretaxGrossAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.RemainingAmortizationPretaxGrossAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetRemainingAmortizationRoundDownDiscount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.RemainingAmortizationRoundDownDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetResourceGroup(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetRoundDownDiscount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.RoundDownDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetSplitAccountName(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.SplitAccountName = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetSplitItemID(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.SplitItemID = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetSplitItemName(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.SplitItemName = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetSplitProductDetail(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.SplitProductDetail = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetSubscriptionType(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetTag(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.Tag = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) SetZone(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems {
	s.Zone = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
