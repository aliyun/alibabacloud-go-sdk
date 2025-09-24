// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceBillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryInstanceBillResponseBody
	GetCode() *string
	SetData(v *QueryInstanceBillResponseBodyData) *QueryInstanceBillResponseBody
	GetData() *QueryInstanceBillResponseBodyData
	SetMessage(v string) *QueryInstanceBillResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryInstanceBillResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryInstanceBillResponseBody
	GetSuccess() *bool
}

type QueryInstanceBillResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryInstanceBillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryInstanceBillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceBillResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstanceBillResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryInstanceBillResponseBody) GetData() *QueryInstanceBillResponseBodyData {
	return s.Data
}

func (s *QueryInstanceBillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryInstanceBillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInstanceBillResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryInstanceBillResponseBody) SetCode(v string) *QueryInstanceBillResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInstanceBillResponseBody) SetData(v *QueryInstanceBillResponseBodyData) *QueryInstanceBillResponseBody {
	s.Data = v
	return s
}

func (s *QueryInstanceBillResponseBody) SetMessage(v string) *QueryInstanceBillResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInstanceBillResponseBody) SetRequestId(v string) *QueryInstanceBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInstanceBillResponseBody) SetSuccess(v bool) *QueryInstanceBillResponseBody {
	s.Success = &v
	return s
}

func (s *QueryInstanceBillResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryInstanceBillResponseBodyData struct {
	// The ID of the account.
	//
	// example:
	//
	// 185xxxx3489
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
	Items *QueryInstanceBillResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
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
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryInstanceBillResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceBillResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryInstanceBillResponseBodyData) GetAccountID() *string {
	return s.AccountID
}

func (s *QueryInstanceBillResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *QueryInstanceBillResponseBodyData) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *QueryInstanceBillResponseBodyData) GetItems() *QueryInstanceBillResponseBodyDataItems {
	return s.Items
}

func (s *QueryInstanceBillResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryInstanceBillResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryInstanceBillResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryInstanceBillResponseBodyData) SetAccountID(v string) *QueryInstanceBillResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *QueryInstanceBillResponseBodyData) SetAccountName(v string) *QueryInstanceBillResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *QueryInstanceBillResponseBodyData) SetBillingCycle(v string) *QueryInstanceBillResponseBodyData {
	s.BillingCycle = &v
	return s
}

func (s *QueryInstanceBillResponseBodyData) SetItems(v *QueryInstanceBillResponseBodyDataItems) *QueryInstanceBillResponseBodyData {
	s.Items = v
	return s
}

func (s *QueryInstanceBillResponseBodyData) SetPageNum(v int32) *QueryInstanceBillResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryInstanceBillResponseBodyData) SetPageSize(v int32) *QueryInstanceBillResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryInstanceBillResponseBodyData) SetTotalCount(v int32) *QueryInstanceBillResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryInstanceBillResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryInstanceBillResponseBodyDataItems struct {
	Item []*QueryInstanceBillResponseBodyDataItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QueryInstanceBillResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceBillResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QueryInstanceBillResponseBodyDataItems) GetItem() []*QueryInstanceBillResponseBodyDataItemsItem {
	return s.Item
}

func (s *QueryInstanceBillResponseBodyDataItems) SetItem(v []*QueryInstanceBillResponseBodyDataItemsItem) *QueryInstanceBillResponseBodyDataItems {
	s.Item = v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}

type QueryInstanceBillResponseBodyDataItemsItem struct {
	// The amount deducted by using credit refunds.
	//
	// example:
	//
	// 0
	AdjustAmount *float32 `json:"AdjustAmount,omitempty" xml:"AdjustAmount,omitempty"`
	// The billing date. This parameter is returned only if the Granularity parameter is set to DAILY. Format: YYYY-MM-DD.
	//
	// example:
	//
	// 2020-03
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The billable item. This parameter is returned only if the IsBillingItem parameter is set to true.
	//
	// example:
	//
	// Other
	BillingItem *string `json:"BillingItem,omitempty" xml:"BillingItem,omitempty"`
	// The billing type.
	//
	// example:
	//
	// Other
	BillingType *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
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
	// The cost center.
	//
	// example:
	//
	// Not allocated
	CostUnit *string `json:"CostUnit,omitempty" xml:"CostUnit,omitempty"`
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
	// The amount deducted by using resource plans.
	//
	// example:
	//
	// NULL
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
	// i-kjhdskjgshfdlkjfdh
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The instance type of the instance.
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
	// 0
	InvoiceDiscount *float32 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	// The type of the bill.
	//
	// 	- SubscriptionOrder: subscription order
	//
	// 	- PayAsYouGoBill: pay-as-you-go bill
	//
	// 	- Refund: refund
	//
	// 	- Adjustment: reconciliation
	//
	// example:
	//
	// PayAsYouGoBill
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The unit price. This parameter is returned only if the isBillingItem parameter is set to true.
	//
	// example:
	//
	// 100
	ListPrice *string `json:"ListPrice,omitempty" xml:"ListPrice,omitempty"`
	// The unit of the unit price. This parameter is returned only if the isBillingItem parameter is set to true.
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
	// 0
	OutstandingAmount *float32 `json:"OutstandingAmount,omitempty" xml:"OutstandingAmount,omitempty"`
	// The ID of the member account. This parameter is returned in a multi-account payment scenario.
	//
	// example:
	//
	// 169***013
	OwnerID *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// The amount paid in cash.
	//
	// example:
	//
	// 0
	PaymentAmount *float32 `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
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
	// 0
	PretaxAmount *float32 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
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
	// The region.
	//
	// example:
	//
	// China (Hangzhou)
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The resource group.
	//
	// example:
	//
	// Default resource group
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The validity period.
	//
	// example:
	//
	// 10
	ServicePeriod *string `json:"ServicePeriod,omitempty" xml:"ServicePeriod,omitempty"`
	// The unit of the validity period.
	//
	// example:
	//
	// Seconds
	ServicePeriodUnit *string `json:"ServicePeriodUnit,omitempty" xml:"ServicePeriodUnit,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription: the subscription billing method
	//
	// 	- PayAsYouGo: the pay-as-you-go billing method
	//
	// **
	//
	// ****This parameter is returned together with the ProductCode parameter.
	//
	// example:
	//
	// PayAsYouGo
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	// The tag.
	//
	// example:
	//
	// tag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The usage of the billable item. This parameter is returned only if the isBillingItem parameter is set to true.
	//
	// example:
	//
	// 100
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// The unit of usage. This parameter is returned only if the isBillingItem parameter is set to true.
	//
	// example:
	//
	// GB
	UsageUnit *string `json:"UsageUnit,omitempty" xml:"UsageUnit,omitempty"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-hangzhou-h
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s QueryInstanceBillResponseBodyDataItemsItem) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceBillResponseBodyDataItemsItem) GoString() string {
	return s.String()
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetAdjustAmount() *float32 {
	return s.AdjustAmount
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetBillingDate() *string {
	return s.BillingDate
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetBillingItem() *string {
	return s.BillingItem
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetBillingType() *string {
	return s.BillingType
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetCashAmount() *float32 {
	return s.CashAmount
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetCostUnit() *string {
	return s.CostUnit
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetCurrency() *string {
	return s.Currency
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetDeductedByCashCoupons() *float32 {
	return s.DeductedByCashCoupons
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetDeductedByCoupons() *float32 {
	return s.DeductedByCoupons
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetDeductedByPrepaidCard() *float32 {
	return s.DeductedByPrepaidCard
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetDeductedByResourcePackage() *string {
	return s.DeductedByResourcePackage
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetInstanceConfig() *string {
	return s.InstanceConfig
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetInstanceID() *string {
	return s.InstanceID
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetInternetIP() *string {
	return s.InternetIP
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetIntranetIP() *string {
	return s.IntranetIP
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetInvoiceDiscount() *float32 {
	return s.InvoiceDiscount
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetItem() *string {
	return s.Item
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetListPrice() *string {
	return s.ListPrice
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetListPriceUnit() *string {
	return s.ListPriceUnit
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetNickName() *string {
	return s.NickName
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetOutstandingAmount() *float32 {
	return s.OutstandingAmount
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetOwnerID() *string {
	return s.OwnerID
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetPaymentAmount() *float32 {
	return s.PaymentAmount
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetPipCode() *string {
	return s.PipCode
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetPretaxAmount() *float32 {
	return s.PretaxAmount
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetPretaxGrossAmount() *float32 {
	return s.PretaxGrossAmount
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetProductName() *string {
	return s.ProductName
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetProductType() *string {
	return s.ProductType
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetRegion() *string {
	return s.Region
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetServicePeriod() *string {
	return s.ServicePeriod
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetServicePeriodUnit() *string {
	return s.ServicePeriodUnit
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetTag() *string {
	return s.Tag
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetUsage() *string {
	return s.Usage
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetUsageUnit() *string {
	return s.UsageUnit
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) GetZone() *string {
	return s.Zone
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetAdjustAmount(v float32) *QueryInstanceBillResponseBodyDataItemsItem {
	s.AdjustAmount = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetBillingDate(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.BillingDate = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetBillingItem(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.BillingItem = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetBillingType(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.BillingType = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetCashAmount(v float32) *QueryInstanceBillResponseBodyDataItemsItem {
	s.CashAmount = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetCommodityCode(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.CommodityCode = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetCostUnit(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.CostUnit = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetCurrency(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.Currency = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetDeductedByCashCoupons(v float32) *QueryInstanceBillResponseBodyDataItemsItem {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetDeductedByCoupons(v float32) *QueryInstanceBillResponseBodyDataItemsItem {
	s.DeductedByCoupons = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetDeductedByPrepaidCard(v float32) *QueryInstanceBillResponseBodyDataItemsItem {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetDeductedByResourcePackage(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.DeductedByResourcePackage = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetInstanceConfig(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.InstanceConfig = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetInstanceID(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.InstanceID = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetInstanceSpec(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.InstanceSpec = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetInternetIP(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.InternetIP = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetIntranetIP(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.IntranetIP = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetInvoiceDiscount(v float32) *QueryInstanceBillResponseBodyDataItemsItem {
	s.InvoiceDiscount = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetItem(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.Item = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetListPrice(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.ListPrice = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetListPriceUnit(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.ListPriceUnit = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetNickName(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.NickName = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetOutstandingAmount(v float32) *QueryInstanceBillResponseBodyDataItemsItem {
	s.OutstandingAmount = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetOwnerID(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.OwnerID = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetPaymentAmount(v float32) *QueryInstanceBillResponseBodyDataItemsItem {
	s.PaymentAmount = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetPipCode(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.PipCode = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetPretaxAmount(v float32) *QueryInstanceBillResponseBodyDataItemsItem {
	s.PretaxAmount = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetPretaxGrossAmount(v float32) *QueryInstanceBillResponseBodyDataItemsItem {
	s.PretaxGrossAmount = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetProductCode(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.ProductCode = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetProductDetail(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.ProductDetail = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetProductName(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.ProductName = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetProductType(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.ProductType = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetRegion(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.Region = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetResourceGroup(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.ResourceGroup = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetServicePeriod(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.ServicePeriod = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetServicePeriodUnit(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.ServicePeriodUnit = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetSubscriptionType(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.SubscriptionType = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetTag(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.Tag = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetUsage(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.Usage = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetUsageUnit(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.UsageUnit = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetZone(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.Zone = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) Validate() error {
	return dara.Validate(s)
}
