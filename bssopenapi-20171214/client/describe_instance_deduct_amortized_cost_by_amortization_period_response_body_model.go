// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody
	GetCode() *string
	SetData(v *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody
	GetData() *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData
	SetMessage(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody
	GetSuccess() *bool
}

type DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody struct {
	// example:
	//
	// 200
	Code *string                                                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful！
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EAE08A27-386C-579E-966D-8853EC3C5D0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody) GetData() *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData {
	return s.Data
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody) SetCode(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody) SetData(v *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody) SetMessage(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody) SetRequestId(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody) SetSuccess(v bool) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData struct {
	// example:
	//
	// 185766xxxx
	AccountID *string `json:"AccountID,omitempty" xml:"AccountID,omitempty"`
	// example:
	//
	// test@test.aliyunid.com
	AccountName *string                                                                         `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Items       []*DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// CAESEgoQCg4KCmdtdF9jcmVhdGUEARgBIkgKCQBwhGmPcAEAAAo7AzYAAAAxTDgwMDcxMjg3ZDJhNmM3ZDguTDgwMDAwMDAwMDAwMzE1MTIuTDgwMDcyZDMyZTJkYzg3N2U
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 400
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData) GetAccountID() *string {
	return s.AccountID
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData) GetItems() []*DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	return s.Items
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData) SetAccountID(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData) SetAccountName(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData) SetItems(v []*DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData {
	s.Items = v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData) SetMaxResults(v int32) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData) SetNextToken(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData) SetTotalCount(v int32) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyData) Validate() error {
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

type DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems struct {
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
	CostUnit      *string `json:"CostUnit,omitempty" xml:"CostUnit,omitempty"`
	// example:
	//
	// 1234
	CostUnitCode *string `json:"CostUnitCode,omitempty" xml:"CostUnitCode,omitempty"`
	// example:
	//
	// 0
	CurrentAmortizationDeductedByCoupons *float64 `json:"CurrentAmortizationDeductedByCoupons,omitempty" xml:"CurrentAmortizationDeductedByCoupons,omitempty"`
	// example:
	//
	// 0
	CurrentAmortizationInvoiceDiscount *float64 `json:"CurrentAmortizationInvoiceDiscount,omitempty" xml:"CurrentAmortizationInvoiceDiscount,omitempty"`
	// example:
	//
	// 10
	CurrentAmortizationPretaxAmount *float64 `json:"CurrentAmortizationPretaxAmount,omitempty" xml:"CurrentAmortizationPretaxAmount,omitempty"`
	// example:
	//
	// 10
	CurrentAmortizationPretaxGrossAmount *float64 `json:"CurrentAmortizationPretaxGrossAmount,omitempty" xml:"CurrentAmortizationPretaxGrossAmount,omitempty"`
	// example:
	//
	// 0
	CurrentAmortizationRoundDownDiscount *float64 `json:"CurrentAmortizationRoundDownDiscount,omitempty" xml:"CurrentAmortizationRoundDownDiscount,omitempty"`
	// example:
	//
	// 0
	DeductedByCoupons *float64 `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	// example:
	//
	// CPU:12
	InstanceConfig *string `json:"InstanceConfig,omitempty" xml:"InstanceConfig,omitempty"`
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
	// 100
	PretaxAmount *float64 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	// example:
	//
	// 100
	PretaxGrossAmount *float64 `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
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
	// example:
	//
	// spn-001
	ReferFrInstanceID *string `json:"ReferFrInstanceID,omitempty" xml:"ReferFrInstanceID,omitempty"`
	// example:
	//
	// 185xxxxx489
	ReferFrOwnerID *string `json:"ReferFrOwnerID,omitempty" xml:"ReferFrOwnerID,omitempty"`
	// example:
	//
	// savingplan_common_public_cn
	ReferFrProductDetailCode *string `json:"ReferFrProductDetailCode,omitempty" xml:"ReferFrProductDetailCode,omitempty"`
	Region                   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroup            *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// example:
	//
	// 0
	RoundDownDiscount *float64 `json:"RoundDownDiscount,omitempty" xml:"RoundDownDiscount,omitempty"`
	// example:
	//
	// 12@test.com
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
	// key:testKey value:testValue; key:testKey1 value:testValues1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// 0
	UnusedAmortizationDeductedByCoupons *float64 `json:"UnusedAmortizationDeductedByCoupons,omitempty" xml:"UnusedAmortizationDeductedByCoupons,omitempty"`
	// example:
	//
	// 0
	UnusedAmortizationInvoiceDiscount *float64 `json:"UnusedAmortizationInvoiceDiscount,omitempty" xml:"UnusedAmortizationInvoiceDiscount,omitempty"`
	// example:
	//
	// 0
	UnusedAmortizationPretaxAmount *float64 `json:"UnusedAmortizationPretaxAmount,omitempty" xml:"UnusedAmortizationPretaxAmount,omitempty"`
	// example:
	//
	// 0
	UnusedAmortizationPretaxGrossAmount *float64 `json:"UnusedAmortizationPretaxGrossAmount,omitempty" xml:"UnusedAmortizationPretaxGrossAmount,omitempty"`
	// example:
	//
	// 0
	UnusedAmortizationRoundDownDiscount *float64 `json:"UnusedAmortizationRoundDownDiscount,omitempty" xml:"UnusedAmortizationRoundDownDiscount,omitempty"`
	Zone                                *string  `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetAmortizationPeriod() *string {
	return s.AmortizationPeriod
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetAmortizationPeriodDay() *string {
	return s.AmortizationPeriodDay
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetAmortizationStatus() *string {
	return s.AmortizationStatus
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetBillAccountID() *int64 {
	return s.BillAccountID
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetBillAccountName() *string {
	return s.BillAccountName
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetBillOwnerID() *int64 {
	return s.BillOwnerID
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetBillOwnerName() *string {
	return s.BillOwnerName
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetBizType() *string {
	return s.BizType
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetConsumePeriod() *string {
	return s.ConsumePeriod
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCostUnit() *string {
	return s.CostUnit
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCostUnitCode() *string {
	return s.CostUnitCode
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationDeductedByCoupons() *float64 {
	return s.CurrentAmortizationDeductedByCoupons
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationInvoiceDiscount() *float64 {
	return s.CurrentAmortizationInvoiceDiscount
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationPretaxAmount() *float64 {
	return s.CurrentAmortizationPretaxAmount
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationPretaxGrossAmount() *float64 {
	return s.CurrentAmortizationPretaxGrossAmount
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetCurrentAmortizationRoundDownDiscount() *float64 {
	return s.CurrentAmortizationRoundDownDiscount
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetDeductedByCoupons() *float64 {
	return s.DeductedByCoupons
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetInstanceConfig() *string {
	return s.InstanceConfig
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetInternetIP() *string {
	return s.InternetIP
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetIntranetIP() *string {
	return s.IntranetIP
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetInvoiceDiscount() *float64 {
	return s.InvoiceDiscount
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPretaxAmount() *float64 {
	return s.PretaxAmount
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetPretaxGrossAmount() *float64 {
	return s.PretaxGrossAmount
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetProductDetailCode() *string {
	return s.ProductDetailCode
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetReferFrInstanceID() *string {
	return s.ReferFrInstanceID
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetReferFrOwnerID() *string {
	return s.ReferFrOwnerID
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetReferFrProductDetailCode() *string {
	return s.ReferFrProductDetailCode
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRegion() *string {
	return s.Region
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetRoundDownDiscount() *float64 {
	return s.RoundDownDiscount
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetSplitAccountName() *string {
	return s.SplitAccountName
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetSplitItemID() *string {
	return s.SplitItemID
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetSplitItemName() *string {
	return s.SplitItemName
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetSplitProductDetail() *string {
	return s.SplitProductDetail
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetTag() *string {
	return s.Tag
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetUnusedAmortizationDeductedByCoupons() *float64 {
	return s.UnusedAmortizationDeductedByCoupons
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetUnusedAmortizationInvoiceDiscount() *float64 {
	return s.UnusedAmortizationInvoiceDiscount
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetUnusedAmortizationPretaxAmount() *float64 {
	return s.UnusedAmortizationPretaxAmount
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetUnusedAmortizationPretaxGrossAmount() *float64 {
	return s.UnusedAmortizationPretaxGrossAmount
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetUnusedAmortizationRoundDownDiscount() *float64 {
	return s.UnusedAmortizationRoundDownDiscount
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) GetZone() *string {
	return s.Zone
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetAmortizationPeriod(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.AmortizationPeriod = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetAmortizationPeriodDay(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.AmortizationPeriodDay = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetAmortizationStatus(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.AmortizationStatus = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetBillAccountID(v int64) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.BillAccountID = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetBillAccountName(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.BillAccountName = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetBillOwnerID(v int64) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.BillOwnerID = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetBillOwnerName(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.BillOwnerName = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetBizType(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.BizType = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetConsumePeriod(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ConsumePeriod = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCostUnit(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CostUnit = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCostUnitCode(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CostUnitCode = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationDeductedByCoupons(v float64) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationDeductedByCoupons = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationInvoiceDiscount(v float64) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationInvoiceDiscount = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationPretaxAmount(v float64) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationPretaxAmount = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationPretaxGrossAmount(v float64) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationPretaxGrossAmount = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetCurrentAmortizationRoundDownDiscount(v float64) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.CurrentAmortizationRoundDownDiscount = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetDeductedByCoupons(v float64) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.DeductedByCoupons = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetInstanceConfig(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.InstanceConfig = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetInstanceID(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.InstanceID = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetInternetIP(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.InternetIP = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetIntranetIP(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.IntranetIP = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetInvoiceDiscount(v float64) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.InvoiceDiscount = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPretaxAmount(v float64) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PretaxAmount = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetPretaxGrossAmount(v float64) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.PretaxGrossAmount = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetProductCode(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetProductDetail(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ProductDetail = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetProductDetailCode(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ProductDetailCode = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetProductName(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ProductName = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetReferFrInstanceID(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ReferFrInstanceID = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetReferFrOwnerID(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ReferFrOwnerID = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetReferFrProductDetailCode(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ReferFrProductDetailCode = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRegion(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.Region = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetResourceGroup(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetRoundDownDiscount(v float64) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.RoundDownDiscount = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetSplitAccountName(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.SplitAccountName = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetSplitItemID(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.SplitItemID = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetSplitItemName(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.SplitItemName = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetSplitProductDetail(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.SplitProductDetail = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetSubscriptionType(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetTag(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.Tag = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetUnusedAmortizationDeductedByCoupons(v float64) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.UnusedAmortizationDeductedByCoupons = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetUnusedAmortizationInvoiceDiscount(v float64) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.UnusedAmortizationInvoiceDiscount = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetUnusedAmortizationPretaxAmount(v float64) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.UnusedAmortizationPretaxAmount = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetUnusedAmortizationPretaxGrossAmount(v float64) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.UnusedAmortizationPretaxGrossAmount = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetUnusedAmortizationRoundDownDiscount(v float64) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.UnusedAmortizationRoundDownDiscount = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) SetZone(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems {
	s.Zone = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
