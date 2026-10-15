// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSplitItemBillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSplitItemBillResponseBody
	GetCode() *string
	SetData(v *DescribeSplitItemBillResponseBodyData) *DescribeSplitItemBillResponseBody
	GetData() *DescribeSplitItemBillResponseBodyData
	SetMessage(v string) *DescribeSplitItemBillResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSplitItemBillResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSplitItemBillResponseBody
	GetSuccess() *bool
}

type DescribeSplitItemBillResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *DescribeSplitItemBillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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

func (s DescribeSplitItemBillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSplitItemBillResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSplitItemBillResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSplitItemBillResponseBody) GetData() *DescribeSplitItemBillResponseBodyData {
	return s.Data
}

func (s *DescribeSplitItemBillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSplitItemBillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSplitItemBillResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSplitItemBillResponseBody) SetCode(v string) *DescribeSplitItemBillResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSplitItemBillResponseBody) SetData(v *DescribeSplitItemBillResponseBodyData) *DescribeSplitItemBillResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSplitItemBillResponseBody) SetMessage(v string) *DescribeSplitItemBillResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSplitItemBillResponseBody) SetRequestId(v string) *DescribeSplitItemBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSplitItemBillResponseBody) SetSuccess(v bool) *DescribeSplitItemBillResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSplitItemBillResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSplitItemBillResponseBodyData struct {
	// The account ID.
	//
	// example:
	//
	// 185xxxx3489
	AccountID *string `json:"AccountID,omitempty" xml:"AccountID,omitempty"`
	// The user account.
	//
	// example:
	//
	// test@test.aliyunid.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The billing cycle. Format: YYYY-MM.
	//
	// example:
	//
	// 2020-03
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The bill details.
	Items []*DescribeSplitItemBillResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The maximum number of entries returned for the current request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that indicates the position where the current call stops reading. An empty value indicates that all data has been read. Set this value to the NextToken request parameter in the next call.
	//
	// example:
	//
	// CAESEgoQCg4K
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSplitItemBillResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSplitItemBillResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSplitItemBillResponseBodyData) GetAccountID() *string {
	return s.AccountID
}

func (s *DescribeSplitItemBillResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeSplitItemBillResponseBodyData) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *DescribeSplitItemBillResponseBodyData) GetItems() []*DescribeSplitItemBillResponseBodyDataItems {
	return s.Items
}

func (s *DescribeSplitItemBillResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSplitItemBillResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSplitItemBillResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSplitItemBillResponseBodyData) SetAccountID(v string) *DescribeSplitItemBillResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyData) SetAccountName(v string) *DescribeSplitItemBillResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyData) SetBillingCycle(v string) *DescribeSplitItemBillResponseBodyData {
	s.BillingCycle = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyData) SetItems(v []*DescribeSplitItemBillResponseBodyDataItems) *DescribeSplitItemBillResponseBodyData {
	s.Items = v
	return s
}

func (s *DescribeSplitItemBillResponseBodyData) SetMaxResults(v int32) *DescribeSplitItemBillResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyData) SetNextToken(v string) *DescribeSplitItemBillResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyData) SetTotalCount(v int32) *DescribeSplitItemBillResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyData) Validate() error {
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

type DescribeSplitItemBillResponseBodyDataItems struct {
	AdjustAmount *float32 `json:"AdjustAmount,omitempty" xml:"AdjustAmount,omitempty"`
	// The amount after discount, which includes the payable amount after coupon deduction. Calculation rule: Amount after discount = List price - Discount amount.
	//
	// example:
	//
	// 0
	AfterDiscountAmount *float32 `json:"AfterDiscountAmount,omitempty" xml:"AfterDiscountAmount,omitempty"`
	// The ID of the account to which the bill belongs.
	//
	// example:
	//
	// 185xxxx3489
	BillAccountID *string `json:"BillAccountID,omitempty" xml:"BillAccountID,omitempty"`
	// The name of the account to which the bill belongs.
	//
	// example:
	//
	// test@test.aliyunid.com
	BillAccountName *string `json:"BillAccountName,omitempty" xml:"BillAccountName,omitempty"`
	// The billing date in the YYYY-MM-DD format. This parameter is not supported.
	//
	// example:
	//
	// 2020-01-20
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The billing item.
	//
	// example:
	//
	// Bandwidth
	BillingItem *string `json:"BillingItem,omitempty" xml:"BillingItem,omitempty"`
	// The billing item code.
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
	// The business type.
	//
	// example:
	//
	// trusteeship
	BizType    *string  `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CashAmount *float32 `json:"CashAmount,omitempty" xml:"CashAmount,omitempty"`
	// The commodity code, which is the same as the product detail code in the User Center.
	//
	// example:
	//
	// rds
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The cost unit.
	//
	// example:
	//
	// Unallocated
	CostUnit *string `json:"CostUnit,omitempty" xml:"CostUnit,omitempty"`
	// The currency. Valid values:
	//
	// - CNY
	//
	// - USD
	//
	// - JPY
	//
	// example:
	//
	// CNY
	Currency              *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DeductedByCashCoupons *float32 `json:"DeductedByCashCoupons,omitempty" xml:"DeductedByCashCoupons,omitempty"`
	// The coupon discount amount.
	//
	// example:
	//
	// 0
	DeductedByCoupons     *float32 `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	DeductedByPrepaidCard *float32 `json:"DeductedByPrepaidCard,omitempty" xml:"DeductedByPrepaidCard,omitempty"`
	// The amount deducted by resource plans.
	//
	// example:
	//
	// 0
	DeductedByResourcePackage *string `json:"DeductedByResourcePackage,omitempty" xml:"DeductedByResourcePackage,omitempty"`
	// The detailed instance configuration.
	//
	// example:
	//
	// CPU:12
	InstanceConfig *string `json:"InstanceConfig,omitempty" xml:"InstanceConfig,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-kjhdskjgshfdlkjfdh
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The instance specifications.
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
	// The internal IP address.
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
	// The bill type. Valid values:
	//
	// - SubscriptionOrder: Subscription order.
	//
	// - PayAsYouGoBill: Pay-as-you-go bill.
	//
	// - Refund: Refund.
	//
	// - Adjustment: Adjustment.
	//
	// example:
	//
	// PayAsYouGoBill
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The item name.
	//
	// example:
	//
	// iZ28bycvyb4Z
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
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
	// The instance nickname.
	//
	// example:
	//
	// nick
	NickName          *string  `json:"NickName,omitempty" xml:"NickName,omitempty"`
	OutstandingAmount *float32 `json:"OutstandingAmount,omitempty" xml:"OutstandingAmount,omitempty"`
	// The account ID of the resource owner in multi-account payment scenarios.
	//
	// example:
	//
	// 169***013
	OwnerID       *string  `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	PaymentAmount *float32 `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
	// The product code, which is the same as the bill product code in the User Center.
	//
	// example:
	//
	// rds
	PipCode *string `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	// The payable amount.
	//
	// example:
	//
	// 0
	PretaxAmount *float32 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	// The original amount.
	//
	// example:
	//
	// 0
	PretaxGrossAmount *float32 `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	// The product code.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The product detail.
	//
	// example:
	//
	// ApsaraDB RDS
	ProductDetail *string `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	// The product name.
	//
	// example:
	//
	// ApsaraDB RDS
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The product type.
	//
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The region.
	//
	// example:
	//
	// Hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The resource group.
	//
	// example:
	//
	// Default resource group
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The service duration.
	//
	// example:
	//
	// 20
	ServicePeriod *string `json:"ServicePeriod,omitempty" xml:"ServicePeriod,omitempty"`
	// The unit of the service duration.
	//
	// example:
	//
	// Hour
	ServicePeriodUnit *string `json:"ServicePeriodUnit,omitempty" xml:"ServicePeriodUnit,omitempty"`
	// The account ID of the attached resource.
	//
	// example:
	//
	// 12**122
	SplitAccountID *string `json:"SplitAccountID,omitempty" xml:"SplitAccountID,omitempty"`
	// The account name of the attached resource.
	//
	// example:
	//
	// test**1122
	SplitAccountName *string `json:"SplitAccountName,omitempty" xml:"SplitAccountName,omitempty"`
	// The split bill month.
	//
	// example:
	//
	// 2021-06
	SplitBillingCycle *string `json:"SplitBillingCycle,omitempty" xml:"SplitBillingCycle,omitempty"`
	// The split bill date.
	//
	// example:
	//
	// 2021-06-01
	SplitBillingDate *string `json:"SplitBillingDate,omitempty" xml:"SplitBillingDate,omitempty"`
	// The commodity code of the attached resource.
	//
	// example:
	//
	// rds
	SplitCommodityCode *string `json:"SplitCommodityCode,omitempty" xml:"SplitCommodityCode,omitempty"`
	// The ID of the attached resource.
	//
	// example:
	//
	// i-28bycvyb4
	SplitItemID *string `json:"SplitItemID,omitempty" xml:"SplitItemID,omitempty"`
	// The name of the attached resource.
	//
	// example:
	//
	// iZ28bycvyb4Z
	SplitItemName *string `json:"SplitItemName,omitempty" xml:"SplitItemName,omitempty"`
	// The product detail name of the attached resource.
	//
	// example:
	//
	// ApsaraDB
	SplitProductDetail *string `json:"SplitProductDetail,omitempty" xml:"SplitProductDetail,omitempty"`
	// The subscription type. Valid values:
	//
	// - Subscription: Subscription.
	//
	// - PayAsYouGo: Pay-as-you-go.
	//
	// example:
	//
	// PayAsYouGo
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	// The resource tag. If a tag on a resource is changed, the bills generated during the validity period will include this tag.
	//
	// example:
	//
	// key:testKey value:testValue; key:testKey1 value:testValues1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The usage.
	//
	// example:
	//
	// 100
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// The usage unit.
	//
	// example:
	//
	// GB
	UsageUnit *string `json:"UsageUnit,omitempty" xml:"UsageUnit,omitempty"`
	// The zone.
	//
	// example:
	//
	// China North 1
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s DescribeSplitItemBillResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSplitItemBillResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetAdjustAmount() *float32 {
	return s.AdjustAmount
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetAfterDiscountAmount() *float32 {
	return s.AfterDiscountAmount
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetBillAccountID() *string {
	return s.BillAccountID
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetBillAccountName() *string {
	return s.BillAccountName
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetBillingDate() *string {
	return s.BillingDate
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetBillingItem() *string {
	return s.BillingItem
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetBillingItemCode() *string {
	return s.BillingItemCode
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetBillingType() *string {
	return s.BillingType
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetBizType() *string {
	return s.BizType
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetCashAmount() *float32 {
	return s.CashAmount
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetCostUnit() *string {
	return s.CostUnit
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetDeductedByCashCoupons() *float32 {
	return s.DeductedByCashCoupons
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetDeductedByCoupons() *float32 {
	return s.DeductedByCoupons
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetDeductedByPrepaidCard() *float32 {
	return s.DeductedByPrepaidCard
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetDeductedByResourcePackage() *string {
	return s.DeductedByResourcePackage
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetInstanceConfig() *string {
	return s.InstanceConfig
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetInternetIP() *string {
	return s.InternetIP
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetIntranetIP() *string {
	return s.IntranetIP
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetInvoiceDiscount() *float32 {
	return s.InvoiceDiscount
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetItem() *string {
	return s.Item
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetItemName() *string {
	return s.ItemName
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetListPrice() *string {
	return s.ListPrice
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetListPriceUnit() *string {
	return s.ListPriceUnit
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetNickName() *string {
	return s.NickName
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetOutstandingAmount() *float32 {
	return s.OutstandingAmount
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetOwnerID() *string {
	return s.OwnerID
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetPaymentAmount() *float32 {
	return s.PaymentAmount
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetPipCode() *string {
	return s.PipCode
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetPretaxAmount() *float32 {
	return s.PretaxAmount
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetPretaxGrossAmount() *float32 {
	return s.PretaxGrossAmount
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetRegion() *string {
	return s.Region
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetServicePeriod() *string {
	return s.ServicePeriod
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetServicePeriodUnit() *string {
	return s.ServicePeriodUnit
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetSplitAccountID() *string {
	return s.SplitAccountID
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetSplitAccountName() *string {
	return s.SplitAccountName
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetSplitBillingCycle() *string {
	return s.SplitBillingCycle
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetSplitBillingDate() *string {
	return s.SplitBillingDate
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetSplitCommodityCode() *string {
	return s.SplitCommodityCode
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetSplitItemID() *string {
	return s.SplitItemID
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetSplitItemName() *string {
	return s.SplitItemName
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetSplitProductDetail() *string {
	return s.SplitProductDetail
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetTag() *string {
	return s.Tag
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetUsage() *string {
	return s.Usage
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetUsageUnit() *string {
	return s.UsageUnit
}

func (s *DescribeSplitItemBillResponseBodyDataItems) GetZone() *string {
	return s.Zone
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetAdjustAmount(v float32) *DescribeSplitItemBillResponseBodyDataItems {
	s.AdjustAmount = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetAfterDiscountAmount(v float32) *DescribeSplitItemBillResponseBodyDataItems {
	s.AfterDiscountAmount = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetBillAccountID(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.BillAccountID = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetBillAccountName(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.BillAccountName = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetBillingDate(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.BillingDate = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetBillingItem(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.BillingItem = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetBillingItemCode(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.BillingItemCode = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetBillingType(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.BillingType = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetBizType(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.BizType = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetCashAmount(v float32) *DescribeSplitItemBillResponseBodyDataItems {
	s.CashAmount = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetCommodityCode(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.CommodityCode = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetCostUnit(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.CostUnit = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetCurrency(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.Currency = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetDeductedByCashCoupons(v float32) *DescribeSplitItemBillResponseBodyDataItems {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetDeductedByCoupons(v float32) *DescribeSplitItemBillResponseBodyDataItems {
	s.DeductedByCoupons = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetDeductedByPrepaidCard(v float32) *DescribeSplitItemBillResponseBodyDataItems {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetDeductedByResourcePackage(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.DeductedByResourcePackage = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetInstanceConfig(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.InstanceConfig = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetInstanceID(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.InstanceID = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetInstanceSpec(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetInternetIP(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.InternetIP = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetIntranetIP(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.IntranetIP = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetInvoiceDiscount(v float32) *DescribeSplitItemBillResponseBodyDataItems {
	s.InvoiceDiscount = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetItem(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.Item = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetItemName(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.ItemName = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetListPrice(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.ListPrice = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetListPriceUnit(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.ListPriceUnit = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetNickName(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.NickName = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetOutstandingAmount(v float32) *DescribeSplitItemBillResponseBodyDataItems {
	s.OutstandingAmount = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetOwnerID(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.OwnerID = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetPaymentAmount(v float32) *DescribeSplitItemBillResponseBodyDataItems {
	s.PaymentAmount = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetPipCode(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.PipCode = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetPretaxAmount(v float32) *DescribeSplitItemBillResponseBodyDataItems {
	s.PretaxAmount = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetPretaxGrossAmount(v float32) *DescribeSplitItemBillResponseBodyDataItems {
	s.PretaxGrossAmount = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetProductCode(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetProductDetail(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.ProductDetail = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetProductName(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.ProductName = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetProductType(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.ProductType = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetRegion(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.Region = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetResourceGroup(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetServicePeriod(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.ServicePeriod = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetServicePeriodUnit(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.ServicePeriodUnit = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetSplitAccountID(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.SplitAccountID = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetSplitAccountName(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.SplitAccountName = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetSplitBillingCycle(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.SplitBillingCycle = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetSplitBillingDate(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.SplitBillingDate = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetSplitCommodityCode(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.SplitCommodityCode = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetSplitItemID(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.SplitItemID = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetSplitItemName(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.SplitItemName = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetSplitProductDetail(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.SplitProductDetail = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetSubscriptionType(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetTag(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.Tag = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetUsage(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.Usage = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetUsageUnit(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.UsageUnit = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetZone(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.Zone = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
