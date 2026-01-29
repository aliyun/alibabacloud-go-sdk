// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAmortizedCostByAmortizationPeriodResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody
	GetCode() *string
	SetData(v *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody
	GetData() *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData
	SetMessage(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody
	GetSuccess() *bool
}

type DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody) GetData() *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData {
	return s.Data
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody) SetCode(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody) SetData(v *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody) SetMessage(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody) SetRequestId(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody) SetSuccess(v bool) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData struct {
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
	Items []*DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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

func (s DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData) GetAccountID() *string {
	return s.AccountID
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData) GetItems() []*DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	return s.Items
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData) SetAccountID(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData) SetAccountName(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData) SetItems(v []*DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData {
	s.Items = v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData) SetMaxResults(v int32) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData) SetNextToken(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData) SetTotalCount(v int32) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyData) Validate() error {
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

type DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems struct {
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
	ConsumePeriod *string `json:"ConsumePeriod,omitempty" xml:"ConsumePeriod,omitempty"`
	// The cost center.
	//
	// example:
	//
	// Not allocated
	CostUnit *string `json:"CostUnit,omitempty" xml:"CostUnit,omitempty"`
	// The code of the cost center.
	//
	// example:
	//
	// 1234
	CostUnitCode                           *string  `json:"CostUnitCode,omitempty" xml:"CostUnitCode,omitempty"`
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
	// The ID of the instance.
	//
	// example:
	//
	// i-kjhdskjgshfdlkjfdh
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The public IP address of the instance.
	//
	// example:
	//
	// 34.xx.x.x
	InternetIP *string `json:"InternetIP,omitempty" xml:"InternetIP,omitempty"`
	// The private IP address of the instance.
	//
	// example:
	//
	// 192.xx.xx.xx
	IntranetIP *string `json:"IntranetIP,omitempty" xml:"IntranetIP,omitempty"`
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
	// The expenditure amount allocated before the allocation month. Invoicing is supported.
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
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The region.
	//
	// example:
	//
	// China (Hangzhou)
	Region                                   *string  `json:"Region,omitempty" xml:"Region,omitempty"`
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
	// The name of the resource group.
	//
	// example:
	//
	// Default resource group
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The round-off amount.
	//
	// example:
	//
	// 0
	RoundDownDiscount *float64 `json:"RoundDownDiscount,omitempty" xml:"RoundDownDiscount,omitempty"`
	// The name of the account to which the split item belongs.
	//
	// example:
	//
	// test**1122
	SplitAccountName *string `json:"SplitAccountName,omitempty" xml:"SplitAccountName,omitempty"`
	// The ID of the split item.
	//
	// example:
	//
	// i-28bycvyb4
	SplitItemID *string `json:"SplitItemID,omitempty" xml:"SplitItemID,omitempty"`
	// The name of the split item.
	//
	// example:
	//
	// iZ28bycvyb4Z
	SplitItemName *string `json:"SplitItemName,omitempty" xml:"SplitItemName,omitempty"`
	// The name of the specific service resource to which the split item belongs.
	//
	// example:
	//
	// rds
	SplitProductDetail *string `json:"SplitProductDetail,omitempty" xml:"SplitProductDetail,omitempty"`
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
	// The tag of the instance.
	//
	// example:
	//
	// tag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-hangzhou
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetAfterDiscountAmount() *float64 {
	return s.AfterDiscountAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetAmortizationPeriod() *string {
	return s.AmortizationPeriod
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetAmortizationStatus() *string {
	return s.AmortizationStatus
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetBillAccountID() *int64 {
	return s.BillAccountID
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetBillAccountName() *string {
	return s.BillAccountName
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetBillOwnerID() *int64 {
	return s.BillOwnerID
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetBillOwnerName() *string {
	return s.BillOwnerName
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetBizType() *string {
	return s.BizType
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetConsumePeriod() *string {
	return s.ConsumePeriod
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCostUnit() *string {
	return s.CostUnit
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCostUnitCode() *string {
	return s.CostUnitCode
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationAfterDiscountAmount() *float64 {
	return s.CurrentAmortizationAfterDiscountAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationDeductedByCashCoupons() *float64 {
	return s.CurrentAmortizationDeductedByCashCoupons
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationDeductedByCoupons() *float64 {
	return s.CurrentAmortizationDeductedByCoupons
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationDeductedByPrepaidCard() *float64 {
	return s.CurrentAmortizationDeductedByPrepaidCard
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationExpenditureAmount() *float64 {
	return s.CurrentAmortizationExpenditureAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationInvoiceDiscount() *float64 {
	return s.CurrentAmortizationInvoiceDiscount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationPretaxAmount() *float64 {
	return s.CurrentAmortizationPretaxAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationPretaxGrossAmount() *float64 {
	return s.CurrentAmortizationPretaxGrossAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationRoundDownDiscount() *float64 {
	return s.CurrentAmortizationRoundDownDiscount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetDeductedByCashCoupons() *float64 {
	return s.DeductedByCashCoupons
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetDeductedByCoupons() *float64 {
	return s.DeductedByCoupons
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetDeductedByPrepaidCard() *float64 {
	return s.DeductedByPrepaidCard
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetExpenditureAmount() *float64 {
	return s.ExpenditureAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetInternetIP() *string {
	return s.InternetIP
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetIntranetIP() *string {
	return s.IntranetIP
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetInvoiceDiscount() *float64 {
	return s.InvoiceDiscount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPretaxAmount() *float64 {
	return s.PretaxAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPretaxGrossAmount() *float64 {
	return s.PretaxGrossAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPreviouslyAmortizedAfterDiscountAmount() *float64 {
	return s.PreviouslyAmortizedAfterDiscountAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPreviouslyAmortizedDeductedByCashCoupons() *float64 {
	return s.PreviouslyAmortizedDeductedByCashCoupons
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPreviouslyAmortizedDeductedByCoupons() *float64 {
	return s.PreviouslyAmortizedDeductedByCoupons
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPreviouslyAmortizedDeductedByPrepaidCard() *float64 {
	return s.PreviouslyAmortizedDeductedByPrepaidCard
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPreviouslyAmortizedExpenditureAmount() *float64 {
	return s.PreviouslyAmortizedExpenditureAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPreviouslyAmortizedInvoiceDiscount() *float64 {
	return s.PreviouslyAmortizedInvoiceDiscount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPreviouslyAmortizedPretaxAmount() *float64 {
	return s.PreviouslyAmortizedPretaxAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPreviouslyAmortizedPretaxGrossAmount() *float64 {
	return s.PreviouslyAmortizedPretaxGrossAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPreviouslyAmortizedRoundDownDiscount() *float64 {
	return s.PreviouslyAmortizedRoundDownDiscount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetProductDetailCode() *string {
	return s.ProductDetailCode
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRegion() *string {
	return s.Region
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRemainingAmortizationAfterDiscountAmount() *float64 {
	return s.RemainingAmortizationAfterDiscountAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRemainingAmortizationDeductedByCashCoupons() *float64 {
	return s.RemainingAmortizationDeductedByCashCoupons
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRemainingAmortizationDeductedByCoupons() *float64 {
	return s.RemainingAmortizationDeductedByCoupons
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRemainingAmortizationDeductedByPrepaidCard() *float64 {
	return s.RemainingAmortizationDeductedByPrepaidCard
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRemainingAmortizationExpenditureAmount() *float64 {
	return s.RemainingAmortizationExpenditureAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRemainingAmortizationInvoiceDiscount() *float64 {
	return s.RemainingAmortizationInvoiceDiscount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRemainingAmortizationPretaxAmount() *float64 {
	return s.RemainingAmortizationPretaxAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRemainingAmortizationPretaxGrossAmount() *float64 {
	return s.RemainingAmortizationPretaxGrossAmount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRemainingAmortizationRoundDownDiscount() *float64 {
	return s.RemainingAmortizationRoundDownDiscount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRoundDownDiscount() *float64 {
	return s.RoundDownDiscount
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetSplitAccountName() *string {
	return s.SplitAccountName
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetSplitItemID() *string {
	return s.SplitItemID
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetSplitItemName() *string {
	return s.SplitItemName
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetSplitProductDetail() *string {
	return s.SplitProductDetail
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetTag() *string {
	return s.Tag
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetZone() *string {
	return s.Zone
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetAfterDiscountAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.AfterDiscountAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetAmortizationPeriod(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.AmortizationPeriod = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetAmortizationStatus(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.AmortizationStatus = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetBillAccountID(v int64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.BillAccountID = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetBillAccountName(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.BillAccountName = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetBillOwnerID(v int64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.BillOwnerID = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetBillOwnerName(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.BillOwnerName = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetBizType(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.BizType = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetConsumePeriod(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ConsumePeriod = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCostUnit(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CostUnit = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCostUnitCode(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CostUnitCode = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationAfterDiscountAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationAfterDiscountAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationDeductedByCashCoupons(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationDeductedByCashCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationDeductedByCoupons(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationDeductedByCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationDeductedByPrepaidCard(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationDeductedByPrepaidCard = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationExpenditureAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationExpenditureAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationInvoiceDiscount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationInvoiceDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationPretaxAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationPretaxAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationPretaxGrossAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationPretaxGrossAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationRoundDownDiscount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationRoundDownDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetDeductedByCashCoupons(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetDeductedByCoupons(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.DeductedByCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetDeductedByPrepaidCard(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetExpenditureAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ExpenditureAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetInstanceID(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.InstanceID = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetInternetIP(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.InternetIP = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetIntranetIP(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.IntranetIP = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetInvoiceDiscount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.InvoiceDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPretaxAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PretaxAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPretaxGrossAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PretaxGrossAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPreviouslyAmortizedAfterDiscountAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PreviouslyAmortizedAfterDiscountAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPreviouslyAmortizedDeductedByCashCoupons(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PreviouslyAmortizedDeductedByCashCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPreviouslyAmortizedDeductedByCoupons(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PreviouslyAmortizedDeductedByCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPreviouslyAmortizedDeductedByPrepaidCard(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PreviouslyAmortizedDeductedByPrepaidCard = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPreviouslyAmortizedExpenditureAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PreviouslyAmortizedExpenditureAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPreviouslyAmortizedInvoiceDiscount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PreviouslyAmortizedInvoiceDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPreviouslyAmortizedPretaxAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PreviouslyAmortizedPretaxAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPreviouslyAmortizedPretaxGrossAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PreviouslyAmortizedPretaxGrossAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPreviouslyAmortizedRoundDownDiscount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PreviouslyAmortizedRoundDownDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetProductCode(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetProductDetail(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ProductDetail = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetProductDetailCode(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ProductDetailCode = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetProductName(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ProductName = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRegion(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.Region = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRemainingAmortizationAfterDiscountAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RemainingAmortizationAfterDiscountAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRemainingAmortizationDeductedByCashCoupons(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RemainingAmortizationDeductedByCashCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRemainingAmortizationDeductedByCoupons(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RemainingAmortizationDeductedByCoupons = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRemainingAmortizationDeductedByPrepaidCard(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RemainingAmortizationDeductedByPrepaidCard = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRemainingAmortizationExpenditureAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RemainingAmortizationExpenditureAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRemainingAmortizationInvoiceDiscount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RemainingAmortizationInvoiceDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRemainingAmortizationPretaxAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RemainingAmortizationPretaxAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRemainingAmortizationPretaxGrossAmount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RemainingAmortizationPretaxGrossAmount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRemainingAmortizationRoundDownDiscount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RemainingAmortizationRoundDownDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetResourceGroup(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRoundDownDiscount(v float64) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RoundDownDiscount = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetSplitAccountName(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.SplitAccountName = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetSplitItemID(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.SplitItemID = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetSplitItemName(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.SplitItemName = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetSplitProductDetail(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.SplitProductDetail = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetSubscriptionType(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetTag(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.Tag = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetZone(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.Zone = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
