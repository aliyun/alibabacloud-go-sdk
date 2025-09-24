// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceBillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeInstanceBillResponseBody
	GetCode() *string
	SetData(v *DescribeInstanceBillResponseBodyData) *DescribeInstanceBillResponseBody
	GetData() *DescribeInstanceBillResponseBodyData
	SetMessage(v string) *DescribeInstanceBillResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeInstanceBillResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstanceBillResponseBody
	GetSuccess() *bool
}

type DescribeInstanceBillResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DescribeInstanceBillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInstanceBillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceBillResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceBillResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeInstanceBillResponseBody) GetData() *DescribeInstanceBillResponseBodyData {
	return s.Data
}

func (s *DescribeInstanceBillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeInstanceBillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceBillResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstanceBillResponseBody) SetCode(v string) *DescribeInstanceBillResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceBillResponseBody) SetData(v *DescribeInstanceBillResponseBodyData) *DescribeInstanceBillResponseBody {
	s.Data = v
	return s
}

func (s *DescribeInstanceBillResponseBody) SetMessage(v string) *DescribeInstanceBillResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceBillResponseBody) SetRequestId(v string) *DescribeInstanceBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceBillResponseBody) SetSuccess(v bool) *DescribeInstanceBillResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceBillResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceBillResponseBodyData struct {
	// The ID of the account.
	//
	// example:
	//
	// 122
	AccountID *string `json:"AccountID,omitempty" xml:"AccountID,omitempty"`
	// The name of the account.
	//
	// example:
	//
	// test@test.aliyunid.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The billing cycle in the YYYY-MM format.
	//
	// example:
	//
	// 2020-03
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The details of the bill.
	Items []*DescribeInstanceBillResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The position where the query stopped. If this parameter is left empty, all the results are returned. If you perform another call, you must set the NextToken parameter to the value of this parameter.
	//
	// example:
	//
	// CAESEgoQCg4KCm
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceBillResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceBillResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceBillResponseBodyData) GetAccountID() *string {
	return s.AccountID
}

func (s *DescribeInstanceBillResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeInstanceBillResponseBodyData) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *DescribeInstanceBillResponseBodyData) GetItems() []*DescribeInstanceBillResponseBodyDataItems {
	return s.Items
}

func (s *DescribeInstanceBillResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInstanceBillResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceBillResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstanceBillResponseBodyData) SetAccountID(v string) *DescribeInstanceBillResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyData) SetAccountName(v string) *DescribeInstanceBillResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyData) SetBillingCycle(v string) *DescribeInstanceBillResponseBodyData {
	s.BillingCycle = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyData) SetItems(v []*DescribeInstanceBillResponseBodyDataItems) *DescribeInstanceBillResponseBodyData {
	s.Items = v
	return s
}

func (s *DescribeInstanceBillResponseBodyData) SetMaxResults(v int32) *DescribeInstanceBillResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyData) SetNextToken(v string) *DescribeInstanceBillResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyData) SetTotalCount(v int32) *DescribeInstanceBillResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceBillResponseBodyDataItems struct {
	// The amount deducted with credit refund.
	//
	// example:
	//
	// 0
	AdjustAmount        *float32 `json:"AdjustAmount,omitempty" xml:"AdjustAmount,omitempty"`
	AfterDiscountAmount *float32 `json:"AfterDiscountAmount,omitempty" xml:"AfterDiscountAmount,omitempty"`
	// The ID of the account to which the bill belongs.
	//
	// example:
	//
	// 122
	BillAccountID *string `json:"BillAccountID,omitempty" xml:"BillAccountID,omitempty"`
	// The name of the account to which the bill belongs.
	//
	// example:
	//
	// test@test.aliyunid.com
	BillAccountName *string `json:"BillAccountName,omitempty" xml:"BillAccountName,omitempty"`
	// The billing date. This parameter is returned only if the Granularity parameter is set to DAILY. Format: YYYY-MM-DD.
	//
	// example:
	//
	// 2020-03-20
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The billable item. A value is returned only if the IsBillingItem parameter is set to true.
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
	// The billing method.
	//
	// example:
	//
	// Other
	BillingType *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	// The type of business.
	//
	// example:
	//
	// trusteeship
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The amount paid in cash. The amount deducted with credit refund is not included.
	//
	// example:
	//
	// 0
	CashAmount *float32 `json:"CashAmount,omitempty" xml:"CashAmount,omitempty"`
	// The code of the commodity. The code is the same as that in Cost Center.
	//
	// example:
	//
	// rds
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The cost center.
	//
	// example:
	//
	// Not allocated
	CostUnit *string `json:"CostUnit,omitempty" xml:"CostUnit,omitempty"`
	// The type of currency. Valid values:
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
	// The amount deducted with vouchers.
	//
	// example:
	//
	// 0.1
	DeductedByCashCoupons *float32 `json:"DeductedByCashCoupons,omitempty" xml:"DeductedByCashCoupons,omitempty"`
	// The amount deducted with coupons.
	//
	// example:
	//
	// 0.1
	DeductedByCoupons *float32 `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	// The amount deducted with prepaid cards.
	//
	// example:
	//
	// 0.1
	DeductedByPrepaidCard *float32 `json:"DeductedByPrepaidCard,omitempty" xml:"DeductedByPrepaidCard,omitempty"`
	// The amount deducted with resource plans. This parameter is valid only when the isBillingItem parameter is set to true.
	//
	// example:
	//
	// 0.1
	DeductedByResourcePackage *string `json:"DeductedByResourcePackage,omitempty" xml:"DeductedByResourcePackage,omitempty"`
	// The configurations of the instance.
	//
	// example:
	//
	// CPU:12
	InstanceConfig *string `json:"InstanceConfig,omitempty" xml:"InstanceConfig,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-dadada
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The type of the instance.
	//
	// example:
	//
	// ecs.sn1ne.3xlarge
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 34.xx.x.x
	InternetIP *string `json:"InternetIP,omitempty" xml:"InternetIP,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 192.xx.xx.xx
	IntranetIP *string `json:"IntranetIP,omitempty" xml:"IntranetIP,omitempty"`
	// The discount amount.
	//
	// example:
	//
	// 0.1
	InvoiceDiscount *float32 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	// The type of the bill.
	//
	// 	- SubscriptionOrder: the subscription bill.
	//
	// 	- PayAsYouGoBill: the pay-as-you-go bill,
	//
	// 	- Refund: the refund.
	//
	// 	- Adjustment: the adjustment bill.
	//
	// example:
	//
	// PayAsYouGoBill
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// iZ28bycvyb4Z
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// The unit price of the service. This parameter is valid only when the isBillingItem parameter is set to true.
	//
	// example:
	//
	// 100
	ListPrice *string `json:"ListPrice,omitempty" xml:"ListPrice,omitempty"`
	// The unit of the unit price. This parameter is valid only when the isBillingItem parameter is set to true.
	//
	// example:
	//
	// CNY
	ListPriceUnit *string `json:"ListPriceUnit,omitempty" xml:"ListPriceUnit,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// test
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The unsettled amount.
	//
	// example:
	//
	// 0.1
	OutstandingAmount *float32 `json:"OutstandingAmount,omitempty" xml:"OutstandingAmount,omitempty"`
	// The ID of the account that owns the resource. This parameter is returned in multi-account payment scenario.
	//
	// example:
	//
	// 123
	OwnerID *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// The amount paid in cash. The amount deducted with credit refund is not included.
	//
	// example:
	//
	// 0.1
	PaymentAmount *float32 `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
	// The code of the service. The code is the same as that in Cost Center.
	//
	// example:
	//
	// rds
	PipCode *string `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	// The pretax amount.
	//
	// example:
	//
	// 0.1
	PretaxAmount *float32 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	// The pretax gross amount.
	//
	// example:
	//
	// 0.1
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
	// ApsaraDB RDS
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
	// The ID of the region.
	//
	// example:
	//
	// China (Hangzhou)
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// Default resource group
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The service duration.
	//
	// example:
	//
	// 3600
	ServicePeriod *string `json:"ServicePeriod,omitempty" xml:"ServicePeriod,omitempty"`
	// The unit of the service duration.
	//
	// example:
	//
	// Second
	ServicePeriodUnit *string `json:"ServicePeriodUnit,omitempty" xml:"ServicePeriodUnit,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription: the subscription billing method.
	//
	// 	- PayAsYouGo: the pay-as-you-go billing method.
	//
	// example:
	//
	// PayAsYouGo
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	// The tag of the resource.
	//
	// example:
	//
	// key:testKey value:testValue; key:testKey1 value:testValues1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The usage of the service. This parameter is valid only when the isBillingItem parameter is set to true. The usage is the total usage in all bills in the billing cycle, not the amount that you purchase. For example, if 1 GB of storage is used and bills are generated every hour, the usage is 1 GB per hour. In this case, the usage is 24 GB per day.
	//
	// example:
	//
	// 100
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// The unit of usage. This parameter is valid only when the isBillingItem parameter is set to true.
	//
	// example:
	//
	// GB
	UsageUnit *string `json:"UsageUnit,omitempty" xml:"UsageUnit,omitempty"`
	// The ID of the zone.
	//
	// example:
	//
	// Hangzhou Zone B
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s DescribeInstanceBillResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceBillResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetAdjustAmount() *float32 {
	return s.AdjustAmount
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetAfterDiscountAmount() *float32 {
	return s.AfterDiscountAmount
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetBillAccountID() *string {
	return s.BillAccountID
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetBillAccountName() *string {
	return s.BillAccountName
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetBillingDate() *string {
	return s.BillingDate
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetBillingItem() *string {
	return s.BillingItem
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetBillingItemCode() *string {
	return s.BillingItemCode
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetBillingType() *string {
	return s.BillingType
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetBizType() *string {
	return s.BizType
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetCashAmount() *float32 {
	return s.CashAmount
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetCostUnit() *string {
	return s.CostUnit
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetDeductedByCashCoupons() *float32 {
	return s.DeductedByCashCoupons
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetDeductedByCoupons() *float32 {
	return s.DeductedByCoupons
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetDeductedByPrepaidCard() *float32 {
	return s.DeductedByPrepaidCard
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetDeductedByResourcePackage() *string {
	return s.DeductedByResourcePackage
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetInstanceConfig() *string {
	return s.InstanceConfig
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetInternetIP() *string {
	return s.InternetIP
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetIntranetIP() *string {
	return s.IntranetIP
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetInvoiceDiscount() *float32 {
	return s.InvoiceDiscount
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetItem() *string {
	return s.Item
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetItemName() *string {
	return s.ItemName
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetListPrice() *string {
	return s.ListPrice
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetListPriceUnit() *string {
	return s.ListPriceUnit
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetNickName() *string {
	return s.NickName
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetOutstandingAmount() *float32 {
	return s.OutstandingAmount
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetOwnerID() *string {
	return s.OwnerID
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetPaymentAmount() *float32 {
	return s.PaymentAmount
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetPipCode() *string {
	return s.PipCode
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetPretaxAmount() *float32 {
	return s.PretaxAmount
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetPretaxGrossAmount() *float32 {
	return s.PretaxGrossAmount
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetRegion() *string {
	return s.Region
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetServicePeriod() *string {
	return s.ServicePeriod
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetServicePeriodUnit() *string {
	return s.ServicePeriodUnit
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetTag() *string {
	return s.Tag
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetUsage() *string {
	return s.Usage
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetUsageUnit() *string {
	return s.UsageUnit
}

func (s *DescribeInstanceBillResponseBodyDataItems) GetZone() *string {
	return s.Zone
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetAdjustAmount(v float32) *DescribeInstanceBillResponseBodyDataItems {
	s.AdjustAmount = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetAfterDiscountAmount(v float32) *DescribeInstanceBillResponseBodyDataItems {
	s.AfterDiscountAmount = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetBillAccountID(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.BillAccountID = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetBillAccountName(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.BillAccountName = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetBillingDate(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.BillingDate = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetBillingItem(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.BillingItem = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetBillingItemCode(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.BillingItemCode = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetBillingType(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.BillingType = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetBizType(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.BizType = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetCashAmount(v float32) *DescribeInstanceBillResponseBodyDataItems {
	s.CashAmount = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetCommodityCode(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.CommodityCode = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetCostUnit(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.CostUnit = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetCurrency(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.Currency = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetDeductedByCashCoupons(v float32) *DescribeInstanceBillResponseBodyDataItems {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetDeductedByCoupons(v float32) *DescribeInstanceBillResponseBodyDataItems {
	s.DeductedByCoupons = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetDeductedByPrepaidCard(v float32) *DescribeInstanceBillResponseBodyDataItems {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetDeductedByResourcePackage(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.DeductedByResourcePackage = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetInstanceConfig(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.InstanceConfig = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetInstanceID(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.InstanceID = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetInstanceSpec(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetInternetIP(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.InternetIP = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetIntranetIP(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.IntranetIP = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetInvoiceDiscount(v float32) *DescribeInstanceBillResponseBodyDataItems {
	s.InvoiceDiscount = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetItem(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.Item = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetItemName(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.ItemName = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetListPrice(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.ListPrice = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetListPriceUnit(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.ListPriceUnit = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetNickName(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.NickName = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetOutstandingAmount(v float32) *DescribeInstanceBillResponseBodyDataItems {
	s.OutstandingAmount = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetOwnerID(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.OwnerID = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetPaymentAmount(v float32) *DescribeInstanceBillResponseBodyDataItems {
	s.PaymentAmount = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetPipCode(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.PipCode = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetPretaxAmount(v float32) *DescribeInstanceBillResponseBodyDataItems {
	s.PretaxAmount = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetPretaxGrossAmount(v float32) *DescribeInstanceBillResponseBodyDataItems {
	s.PretaxGrossAmount = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetProductCode(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetProductDetail(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.ProductDetail = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetProductName(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.ProductName = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetProductType(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.ProductType = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetRegion(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.Region = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetResourceGroup(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetServicePeriod(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.ServicePeriod = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetServicePeriodUnit(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.ServicePeriodUnit = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetSubscriptionType(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetTag(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.Tag = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetUsage(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.Usage = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetUsageUnit(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.UsageUnit = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) SetZone(v string) *DescribeInstanceBillResponseBodyDataItems {
	s.Zone = &v
	return s
}

func (s *DescribeInstanceBillResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
