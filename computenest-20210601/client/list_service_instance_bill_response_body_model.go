// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceBillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItem(v []*ListServiceInstanceBillResponseBodyItem) *ListServiceInstanceBillResponseBody
	GetItem() []*ListServiceInstanceBillResponseBodyItem
	SetMaxResults(v int32) *ListServiceInstanceBillResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceInstanceBillResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServiceInstanceBillResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListServiceInstanceBillResponseBody
	GetTotalCount() *int64
}

type ListServiceInstanceBillResponseBody struct {
	// The billing information of the backup schedule.
	Item []*ListServiceInstanceBillResponseBodyItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbAx7BkQzyYC+ONO+WudHGKEdB0uWSY7AGnM3qCgm/Ynge7zU6NWdbj0Tegyajyqyc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2849EE73-AFFA-5AFD-9575-12FA886451DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceInstanceBillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceBillResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceBillResponseBody) GetItem() []*ListServiceInstanceBillResponseBodyItem {
	return s.Item
}

func (s *ListServiceInstanceBillResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceInstanceBillResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceInstanceBillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceInstanceBillResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListServiceInstanceBillResponseBody) SetItem(v []*ListServiceInstanceBillResponseBodyItem) *ListServiceInstanceBillResponseBody {
	s.Item = v
	return s
}

func (s *ListServiceInstanceBillResponseBody) SetMaxResults(v int32) *ListServiceInstanceBillResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceBillResponseBody) SetNextToken(v string) *ListServiceInstanceBillResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceBillResponseBody) SetRequestId(v string) *ListServiceInstanceBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceBillResponseBody) SetTotalCount(v int64) *ListServiceInstanceBillResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceInstanceBillResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListServiceInstanceBillResponseBodyItem struct {
	// The billing cycle. Format: YYYY-MM.
	//
	// example:
	//
	// 2025-02
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The billing date. This parameter is required only if the **Granularity*	- parameter is set to DAILY. Format: YYYY-MM-DD.
	//
	// example:
	//
	// 2024-10-23
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The billable item.
	//
	// example:
	//
	// Bandwidth
	BillingItem *string `json:"BillingItem,omitempty" xml:"BillingItem,omitempty"`
	// The code of the billable item.
	//
	// example:
	//
	// disk
	BillingItemCode *string `json:"BillingItemCode,omitempty" xml:"BillingItemCode,omitempty"`
	// The currency unit.
	//
	// 	- China site: **CNY**.
	//
	// 	- International site: **USD**.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The amount deducted with resource plans.
	//
	// example:
	//
	// 0
	DeductedByResourcePackage *string `json:"DeductedByResourcePackage,omitempty" xml:"DeductedByResourcePackage,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// rm-bp1z88pb48487907u
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The discount amount.
	//
	// example:
	//
	// 0
	InvoiceDiscount *string `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	// The unit price.
	//
	// example:
	//
	// 0.12
	ListPrice *string `json:"ListPrice,omitempty" xml:"ListPrice,omitempty"`
	// The unit of the unit price.
	//
	// example:
	//
	// CNY/GB
	ListPriceUnit *string `json:"ListPriceUnit,omitempty" xml:"ListPriceUnit,omitempty"`
	// The pretax amount.
	//
	// example:
	//
	// 0
	PretaxAmount *string `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	// The pretax gross amount.
	//
	// example:
	//
	// 0
	PretaxGrossAmount *string `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	// The code of the service.
	//
	// example:
	//
	// sls
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The specific service resource.
	//
	// example:
	//
	// sls
	ProductDetail *string `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	// The name of the cloud service or the name of the service-linked role with which the cloud service is associated.
	//
	// example:
	//
	// NLB
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The billing cycle in which the bill is split.
	//
	// example:
	//
	// 2021-07
	SplitBillingCycle *string `json:"SplitBillingCycle,omitempty" xml:"SplitBillingCycle,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription: the subscription billing method.
	//
	// 	- PayAsYouGo: the pay-as-you-go billing method.
	//
	// example:
	//
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	// The amount of resource usage.
	//
	// example:
	//
	// {\\"EmbeddingTokens\\": 314}
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// The unit of usage.
	//
	// example:
	//
	// GB
	UsageUnit *string `json:"UsageUnit,omitempty" xml:"UsageUnit,omitempty"`
}

func (s ListServiceInstanceBillResponseBodyItem) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceBillResponseBodyItem) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceBillResponseBodyItem) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *ListServiceInstanceBillResponseBodyItem) GetBillingDate() *string {
	return s.BillingDate
}

func (s *ListServiceInstanceBillResponseBodyItem) GetBillingItem() *string {
	return s.BillingItem
}

func (s *ListServiceInstanceBillResponseBodyItem) GetBillingItemCode() *string {
	return s.BillingItemCode
}

func (s *ListServiceInstanceBillResponseBodyItem) GetCurrency() *string {
	return s.Currency
}

func (s *ListServiceInstanceBillResponseBodyItem) GetDeductedByResourcePackage() *string {
	return s.DeductedByResourcePackage
}

func (s *ListServiceInstanceBillResponseBodyItem) GetInstanceID() *string {
	return s.InstanceID
}

func (s *ListServiceInstanceBillResponseBodyItem) GetInvoiceDiscount() *string {
	return s.InvoiceDiscount
}

func (s *ListServiceInstanceBillResponseBodyItem) GetListPrice() *string {
	return s.ListPrice
}

func (s *ListServiceInstanceBillResponseBodyItem) GetListPriceUnit() *string {
	return s.ListPriceUnit
}

func (s *ListServiceInstanceBillResponseBodyItem) GetPretaxAmount() *string {
	return s.PretaxAmount
}

func (s *ListServiceInstanceBillResponseBodyItem) GetPretaxGrossAmount() *string {
	return s.PretaxGrossAmount
}

func (s *ListServiceInstanceBillResponseBodyItem) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListServiceInstanceBillResponseBodyItem) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *ListServiceInstanceBillResponseBodyItem) GetProductName() *string {
	return s.ProductName
}

func (s *ListServiceInstanceBillResponseBodyItem) GetSplitBillingCycle() *string {
	return s.SplitBillingCycle
}

func (s *ListServiceInstanceBillResponseBodyItem) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *ListServiceInstanceBillResponseBodyItem) GetUsage() *string {
	return s.Usage
}

func (s *ListServiceInstanceBillResponseBodyItem) GetUsageUnit() *string {
	return s.UsageUnit
}

func (s *ListServiceInstanceBillResponseBodyItem) SetBillingCycle(v string) *ListServiceInstanceBillResponseBodyItem {
	s.BillingCycle = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetBillingDate(v string) *ListServiceInstanceBillResponseBodyItem {
	s.BillingDate = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetBillingItem(v string) *ListServiceInstanceBillResponseBodyItem {
	s.BillingItem = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetBillingItemCode(v string) *ListServiceInstanceBillResponseBodyItem {
	s.BillingItemCode = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetCurrency(v string) *ListServiceInstanceBillResponseBodyItem {
	s.Currency = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetDeductedByResourcePackage(v string) *ListServiceInstanceBillResponseBodyItem {
	s.DeductedByResourcePackage = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetInstanceID(v string) *ListServiceInstanceBillResponseBodyItem {
	s.InstanceID = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetInvoiceDiscount(v string) *ListServiceInstanceBillResponseBodyItem {
	s.InvoiceDiscount = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetListPrice(v string) *ListServiceInstanceBillResponseBodyItem {
	s.ListPrice = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetListPriceUnit(v string) *ListServiceInstanceBillResponseBodyItem {
	s.ListPriceUnit = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetPretaxAmount(v string) *ListServiceInstanceBillResponseBodyItem {
	s.PretaxAmount = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetPretaxGrossAmount(v string) *ListServiceInstanceBillResponseBodyItem {
	s.PretaxGrossAmount = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetProductCode(v string) *ListServiceInstanceBillResponseBodyItem {
	s.ProductCode = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetProductDetail(v string) *ListServiceInstanceBillResponseBodyItem {
	s.ProductDetail = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetProductName(v string) *ListServiceInstanceBillResponseBodyItem {
	s.ProductName = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetSplitBillingCycle(v string) *ListServiceInstanceBillResponseBodyItem {
	s.SplitBillingCycle = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetSubscriptionType(v string) *ListServiceInstanceBillResponseBodyItem {
	s.SubscriptionType = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetUsage(v string) *ListServiceInstanceBillResponseBodyItem {
	s.Usage = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetUsageUnit(v string) *ListServiceInstanceBillResponseBodyItem {
	s.UsageUnit = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) Validate() error {
	return dara.Validate(s)
}
