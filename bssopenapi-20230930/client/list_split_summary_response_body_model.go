// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSplitSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListSplitSummaryResponseBodyData) *ListSplitSummaryResponseBody
	GetData() []*ListSplitSummaryResponseBodyData
	SetMaxResults(v int32) *ListSplitSummaryResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListSplitSummaryResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSplitSummaryResponseBody
	GetRequestId() *string
}

type ListSplitSummaryResponseBody struct {
	// The array of split bill summary records returned in this request. This array may be empty if no matching records are found.
	//
	// example:
	//
	// []
	Data []*ListSplitSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The maximum number of records per page used in this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for querying the next page. If it is not empty, use it as the NextToken for the next request. If it is empty or not returned, the pagination ends. Even if the number of returned records is less than MaxResults or the current page is empty, you must continue to query as long as the token is not empty.
	//
	// example:
	//
	// opaque-token-example
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The unique identifier of the request, which is used to locate the request.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSplitSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSplitSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *ListSplitSummaryResponseBody) GetData() []*ListSplitSummaryResponseBodyData {
	return s.Data
}

func (s *ListSplitSummaryResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSplitSummaryResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSplitSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSplitSummaryResponseBody) SetData(v []*ListSplitSummaryResponseBodyData) *ListSplitSummaryResponseBody {
	s.Data = v
	return s
}

func (s *ListSplitSummaryResponseBody) SetMaxResults(v int32) *ListSplitSummaryResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSplitSummaryResponseBody) SetNextToken(v string) *ListSplitSummaryResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSplitSummaryResponseBody) SetRequestId(v string) *ListSplitSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSplitSummaryResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSplitSummaryResponseBodyData struct {
	// The amount after coupon deduction, which equals AmountAfterPromotionDeduction minus CouponAmount.
	//
	// example:
	//
	// 85
	AmountAfterCouponDeduction *string `json:"AmountAfterCouponDeduction,omitempty" xml:"AmountAfterCouponDeduction,omitempty"`
	// The amount after discount, which is the amount after official website discounts and contract discounts are applied, but before coupon deductions. When the savings plan deduction data is complete, this equals ListPrice minus SavingPlanDeductListPrice minus PromotionAmount. When the savings plan deduction list price is unavailable, this is calculated as ListPrice minus PromotionAmount.
	//
	// example:
	//
	// 90
	AmountAfterPromotionDeduction *string `json:"AmountAfterPromotionDeduction,omitempty" xml:"AmountAfterPromotionDeduction,omitempty"`
	// The amount after savings plan deduction, which equals ListPrice minus SavingPlanDeductListPrice. This may be empty when there is no savings plan deduction data.
	//
	// example:
	//
	// 100
	AmountAfterSavingPlanDeduction *string `json:"AmountAfterSavingPlanDeduction,omitempty" xml:"AmountAfterSavingPlanDeduction,omitempty"`
	// The billing date, which is the date when the user consumed the cloud service. The format is yyyy-MM-dd. This has a value only when the data is summarized by day.
	//
	// example:
	//
	// 2026-09-01
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The billing item code, which is the billing item to which the cloud resource expense belongs. Whether this field has a value depends on the split bill records and summary data.
	//
	// example:
	//
	// example_billing_item
	BillingItemCode *string `json:"BillingItemCode,omitempty" xml:"BillingItemCode,omitempty"`
	// The billing item name.
	//
	// example:
	//
	// Sample billing item
	BillingItemName *string `json:"BillingItemName,omitempty" xml:"BillingItemName,omitempty"`
	// The billing month, which is the month to which the expense belongs. The format is yyyy-MM.
	//
	// example:
	//
	// 2026-09
	BillingMonth *string `json:"BillingMonth,omitempty" xml:"BillingMonth,omitempty"`
	// The usage, which is the total usage that is ultimately billed. This must be understood in combination with the product, billing item, and unit. Do not directly add up the values if the units or billing rules are different.
	//
	// example:
	//
	// 100
	BillingQuantity *string `json:"BillingQuantity,omitempty" xml:"BillingQuantity,omitempty"`
	// The unit of usage, used in conjunction with BillingQuantity.
	//
	// example:
	//
	// GB
	BillingQuantityUnit *string `json:"BillingQuantityUnit,omitempty" xml:"BillingQuantityUnit,omitempty"`
	// The total duration of purchasing or using cloud products. This field is used when the corresponding billing method includes a duration.
	//
	// example:
	//
	// 24
	ChargeDuration *string `json:"ChargeDuration,omitempty" xml:"ChargeDuration,omitempty"`
	// The unit of duration, used in conjunction with ChargeDuration.
	//
	// example:
	//
	// Hour
	ChargeDurationUnit *string `json:"ChargeDurationUnit,omitempty" xml:"ChargeDurationUnit,omitempty"`
	// The display name of the billing method, returned based on the request language, such as Subscription or Pay-as-you-go. Use Subscription or PayAsYouGo for filtering.
	//
	// example:
	//
	// Pay-as-you-go
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The commodity code, which is the unique identifier of a specific commodity under a cloud service.
	//
	// example:
	//
	// oss
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The commodity name.
	//
	// example:
	//
	// Sample commodity
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	// The display name of the transaction type, which identifies the transaction corresponding to the bill, such as new purchase, renewal, upgrade, downgrade, unsubscription, pay-as-you-go, and adjustment. Do not use the display name as the ConsumeType filter code.
	//
	// example:
	//
	// Pay-as-you-go
	ConsumeType *string `json:"ConsumeType,omitempty" xml:"ConsumeType,omitempty"`
	// The cost category, which indicates the cost category and category value to which the record belongs. It is returned as a string. The specific content is subject to the bill record and can be empty.
	//
	// example:
	//
	// {"business":"commerce"}
	CostCategories *string `json:"CostCategories,omitempty" xml:"CostCategories,omitempty"`
	// The coupon deduction amount. This amount cannot be used for invoicing.
	//
	// example:
	//
	// 5
	CouponAmount *string `json:"CouponAmount,omitempty" xml:"CouponAmount,omitempty"`
	// The currency corresponding to the amounts and prices.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The name of the financial unit. A financial unit is a unit used to group or classify expenses and must be maintained by the user.
	//
	// example:
	//
	// engineering
	FinanceUnitName *string `json:"FinanceUnitName,omitempty" xml:"FinanceUnitName,omitempty"`
	// The identifier of the summary record, used as an opaque string. This can be empty. The identifier may change when summary dimensions such as resource groups and tags change.
	//
	// example:
	//
	// example-split-summary-key
	HashKey *string `json:"HashKey,omitempty" xml:"HashKey,omitempty"`
	// The instance ID at the billing granularity, which is the dimension for expense statistics and summary when a cloud service generates a bill. This may be different from the resource ID.
	//
	// example:
	//
	// oss-cn-hangzhou-example
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The resource name, which is the custom name given to the cloud resource by the user. Whether this field has a value depends on whether the product and resource information can be obtained.
	//
	// example:
	//
	// example-resource
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The specification, which describes the specification of the resource or billing item associated with this record.
	//
	// example:
	//
	// example.spec
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The public IP address. This is returned only when the product supports it and the relevant resource information can be obtained.
	//
	// example:
	//
	// 192.0.2.10
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address. This is returned only when the product supports it and the relevant resource information can be obtained.
	//
	// example:
	//
	// 10.0.0.10
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The display name of the expense type, which identifies the nature of the expense generated by the bill, such as subscription prepaid expense, subscription unsubscription expense, pay-as-you-go expense for cloud resources, adjustment expense, and round-down expense. Do not use the display name as the LineItemType filter code.
	//
	// example:
	//
	// Pay-as-you-go expense for cloud resources
	LineItemType *string `json:"LineItemType,omitempty" xml:"LineItemType,omitempty"`
	// The total list price, which is the total list price of the selected configuration.
	//
	// example:
	//
	// 100
	ListPrice *string `json:"ListPrice,omitempty" xml:"ListPrice,omitempty"`
	// The price unit, used in conjunction with ListUnitPrice.
	//
	// example:
	//
	// CNY/GB
	ListPriceUnit *string `json:"ListPriceUnit,omitempty" xml:"ListPriceUnit,omitempty"`
	// The list price on the official website, which is the list unit price of the commodity. Whether this field has a value depends on the split bill and summary data. Do not multiply the total usage by a unit price to recalculate all bills.
	//
	// example:
	//
	// 1
	ListUnitPrice *string `json:"ListUnitPrice,omitempty" xml:"ListUnitPrice,omitempty"`
	// The product code, which is the unique identifier of a specific cloud service on Alibaba Cloud.
	//
	// example:
	//
	// oss
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The product name.
	//
	// example:
	//
	// Object Storage Service (OSS)
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The discount amount, which is the total discount amount on the bill, including official website discounts and contract discounts.
	//
	// example:
	//
	// 10
	PromotionAmount *string `json:"PromotionAmount,omitempty" xml:"PromotionAmount,omitempty"`
	// The purchase configuration of the billing item, which describes the configurations related to the billing item.
	//
	// example:
	//
	// Sample configuration
	PurchaseConfiguration *string `json:"PurchaseConfiguration,omitempty" xml:"PurchaseConfiguration,omitempty"`
	// The region code, which is the logical region code where the cloud resource is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The region name, which indicates the region selected when the order was placed.
	//
	// example:
	//
	// China (Hangzhou)
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The payable amount (tax-inclusive), which is the actual amount the user needs to pay after various discounts, savings plan deductions, and coupon deductions are applied. When all amount fields are complete, this equals ListPrice minus SavingPlanDeductListPrice minus PromotionAmount minus CouponAmount. Use this field for expense verification. If other amount fields are missing, do not treat empty values as zero to calculate the payable amount.
	//
	// example:
	//
	// 85
	RequireAmount *string `json:"RequireAmount,omitempty" xml:"RequireAmount,omitempty"`
	// The resource group, which is the group where cloud resources are located as configured by the user in the Resource Management console.
	//
	// example:
	//
	// production
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The resource ID, which is the identifier used to uniquely identify a resource in a cloud service.
	//
	// example:
	//
	// example-bucket
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the resource ownership account, which is the ID of the account that actually uses the resource and may differ from the purchasing account.
	//
	// example:
	//
	// 1234567890123456
	ResourceOwnerAccountId *string `json:"ResourceOwnerAccountId,omitempty" xml:"ResourceOwnerAccountId,omitempty"`
	// The name of the resource ownership account.
	//
	// example:
	//
	// owner@example.com
	ResourceOwnerAccountName *string `json:"ResourceOwnerAccountName,omitempty" xml:"ResourceOwnerAccountName,omitempty"`
	// The ID of the resource purchaser account, which is the ID of the account that purchases the resource.
	//
	// example:
	//
	// 1234567890123456
	ResourcePurchaseAccountId *string `json:"ResourcePurchaseAccountId,omitempty" xml:"ResourcePurchaseAccountId,omitempty"`
	// The name of the resource purchaser account.
	//
	// example:
	//
	// buyer@example.com
	ResourcePurchaseAccountName *string `json:"ResourcePurchaseAccountName,omitempty" xml:"ResourcePurchaseAccountName,omitempty"`
	// The resource tag, which is the custom tag added by the user to Alibaba Cloud resources. It is returned as a string and can be empty.
	//
	// example:
	//
	// env prod
	ResourceTag *string `json:"ResourceTag,omitempty" xml:"ResourceTag,omitempty"`
	// The savings plan deduction amount, which is the committed amount consumed by the savings plan itself. This has a different meaning from SavingPlanDeductListPrice, cannot be used interchangeably, and should not be deducted from the payable amount again.
	//
	// example:
	//
	// 0
	SavingPlanDeductAmount *string `json:"SavingPlanDeductAmount,omitempty" xml:"SavingPlanDeductAmount,omitempty"`
	// The total list price deducted by the savings plan, which is the total list price of the pay-as-you-go cloud resources deducted when a savings plan is applied.
	//
	// example:
	//
	// 0
	SavingPlanDeductListPrice *string `json:"SavingPlanDeductListPrice,omitempty" xml:"SavingPlanDeductListPrice,omitempty"`
	// The service entity, which is the name of the company that provides the product or service. Alibaba Cloud sells a wide range of products, and some products may be provided by different companies.
	//
	// example:
	//
	// Sample service entity
	ServiceEntity *string `json:"ServiceEntity,omitempty" xml:"ServiceEntity,omitempty"`
	// The split item ID, which is the code corresponding to the specific split granularity of the bill. For example, when OSS is split by bucket, this is the bucket code. This field may be empty for records that are not split.
	//
	// example:
	//
	// example-bucket
	SplitItemId *string `json:"SplitItemId,omitempty" xml:"SplitItemId,omitempty"`
	// The split item name, which is the name corresponding to the specific split granularity of the bill.
	//
	// example:
	//
	// example-bucket
	SplitItemName *string `json:"SplitItemName,omitempty" xml:"SplitItemName,omitempty"`
	// The zone. The display name may be returned.
	//
	// example:
	//
	// China (Hangzhou) Zone I
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s ListSplitSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSplitSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSplitSummaryResponseBodyData) GetAmountAfterCouponDeduction() *string {
	return s.AmountAfterCouponDeduction
}

func (s *ListSplitSummaryResponseBodyData) GetAmountAfterPromotionDeduction() *string {
	return s.AmountAfterPromotionDeduction
}

func (s *ListSplitSummaryResponseBodyData) GetAmountAfterSavingPlanDeduction() *string {
	return s.AmountAfterSavingPlanDeduction
}

func (s *ListSplitSummaryResponseBodyData) GetBillingDate() *string {
	return s.BillingDate
}

func (s *ListSplitSummaryResponseBodyData) GetBillingItemCode() *string {
	return s.BillingItemCode
}

func (s *ListSplitSummaryResponseBodyData) GetBillingItemName() *string {
	return s.BillingItemName
}

func (s *ListSplitSummaryResponseBodyData) GetBillingMonth() *string {
	return s.BillingMonth
}

func (s *ListSplitSummaryResponseBodyData) GetBillingQuantity() *string {
	return s.BillingQuantity
}

func (s *ListSplitSummaryResponseBodyData) GetBillingQuantityUnit() *string {
	return s.BillingQuantityUnit
}

func (s *ListSplitSummaryResponseBodyData) GetChargeDuration() *string {
	return s.ChargeDuration
}

func (s *ListSplitSummaryResponseBodyData) GetChargeDurationUnit() *string {
	return s.ChargeDurationUnit
}

func (s *ListSplitSummaryResponseBodyData) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListSplitSummaryResponseBodyData) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListSplitSummaryResponseBodyData) GetCommodityName() *string {
	return s.CommodityName
}

func (s *ListSplitSummaryResponseBodyData) GetConsumeType() *string {
	return s.ConsumeType
}

func (s *ListSplitSummaryResponseBodyData) GetCostCategories() *string {
	return s.CostCategories
}

func (s *ListSplitSummaryResponseBodyData) GetCouponAmount() *string {
	return s.CouponAmount
}

func (s *ListSplitSummaryResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *ListSplitSummaryResponseBodyData) GetFinanceUnitName() *string {
	return s.FinanceUnitName
}

func (s *ListSplitSummaryResponseBodyData) GetHashKey() *string {
	return s.HashKey
}

func (s *ListSplitSummaryResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSplitSummaryResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListSplitSummaryResponseBodyData) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *ListSplitSummaryResponseBodyData) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListSplitSummaryResponseBodyData) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListSplitSummaryResponseBodyData) GetLineItemType() *string {
	return s.LineItemType
}

func (s *ListSplitSummaryResponseBodyData) GetListPrice() *string {
	return s.ListPrice
}

func (s *ListSplitSummaryResponseBodyData) GetListPriceUnit() *string {
	return s.ListPriceUnit
}

func (s *ListSplitSummaryResponseBodyData) GetListUnitPrice() *string {
	return s.ListUnitPrice
}

func (s *ListSplitSummaryResponseBodyData) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListSplitSummaryResponseBodyData) GetProductName() *string {
	return s.ProductName
}

func (s *ListSplitSummaryResponseBodyData) GetPromotionAmount() *string {
	return s.PromotionAmount
}

func (s *ListSplitSummaryResponseBodyData) GetPurchaseConfiguration() *string {
	return s.PurchaseConfiguration
}

func (s *ListSplitSummaryResponseBodyData) GetRegionCode() *string {
	return s.RegionCode
}

func (s *ListSplitSummaryResponseBodyData) GetRegionName() *string {
	return s.RegionName
}

func (s *ListSplitSummaryResponseBodyData) GetRequireAmount() *string {
	return s.RequireAmount
}

func (s *ListSplitSummaryResponseBodyData) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *ListSplitSummaryResponseBodyData) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListSplitSummaryResponseBodyData) GetResourceOwnerAccountId() *string {
	return s.ResourceOwnerAccountId
}

func (s *ListSplitSummaryResponseBodyData) GetResourceOwnerAccountName() *string {
	return s.ResourceOwnerAccountName
}

func (s *ListSplitSummaryResponseBodyData) GetResourcePurchaseAccountId() *string {
	return s.ResourcePurchaseAccountId
}

func (s *ListSplitSummaryResponseBodyData) GetResourcePurchaseAccountName() *string {
	return s.ResourcePurchaseAccountName
}

func (s *ListSplitSummaryResponseBodyData) GetResourceTag() *string {
	return s.ResourceTag
}

func (s *ListSplitSummaryResponseBodyData) GetSavingPlanDeductAmount() *string {
	return s.SavingPlanDeductAmount
}

func (s *ListSplitSummaryResponseBodyData) GetSavingPlanDeductListPrice() *string {
	return s.SavingPlanDeductListPrice
}

func (s *ListSplitSummaryResponseBodyData) GetServiceEntity() *string {
	return s.ServiceEntity
}

func (s *ListSplitSummaryResponseBodyData) GetSplitItemId() *string {
	return s.SplitItemId
}

func (s *ListSplitSummaryResponseBodyData) GetSplitItemName() *string {
	return s.SplitItemName
}

func (s *ListSplitSummaryResponseBodyData) GetZone() *string {
	return s.Zone
}

func (s *ListSplitSummaryResponseBodyData) SetAmountAfterCouponDeduction(v string) *ListSplitSummaryResponseBodyData {
	s.AmountAfterCouponDeduction = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetAmountAfterPromotionDeduction(v string) *ListSplitSummaryResponseBodyData {
	s.AmountAfterPromotionDeduction = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetAmountAfterSavingPlanDeduction(v string) *ListSplitSummaryResponseBodyData {
	s.AmountAfterSavingPlanDeduction = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetBillingDate(v string) *ListSplitSummaryResponseBodyData {
	s.BillingDate = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetBillingItemCode(v string) *ListSplitSummaryResponseBodyData {
	s.BillingItemCode = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetBillingItemName(v string) *ListSplitSummaryResponseBodyData {
	s.BillingItemName = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetBillingMonth(v string) *ListSplitSummaryResponseBodyData {
	s.BillingMonth = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetBillingQuantity(v string) *ListSplitSummaryResponseBodyData {
	s.BillingQuantity = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetBillingQuantityUnit(v string) *ListSplitSummaryResponseBodyData {
	s.BillingQuantityUnit = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetChargeDuration(v string) *ListSplitSummaryResponseBodyData {
	s.ChargeDuration = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetChargeDurationUnit(v string) *ListSplitSummaryResponseBodyData {
	s.ChargeDurationUnit = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetChargeType(v string) *ListSplitSummaryResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetCommodityCode(v string) *ListSplitSummaryResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetCommodityName(v string) *ListSplitSummaryResponseBodyData {
	s.CommodityName = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetConsumeType(v string) *ListSplitSummaryResponseBodyData {
	s.ConsumeType = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetCostCategories(v string) *ListSplitSummaryResponseBodyData {
	s.CostCategories = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetCouponAmount(v string) *ListSplitSummaryResponseBodyData {
	s.CouponAmount = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetCurrency(v string) *ListSplitSummaryResponseBodyData {
	s.Currency = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetFinanceUnitName(v string) *ListSplitSummaryResponseBodyData {
	s.FinanceUnitName = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetHashKey(v string) *ListSplitSummaryResponseBodyData {
	s.HashKey = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetInstanceId(v string) *ListSplitSummaryResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetInstanceName(v string) *ListSplitSummaryResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetInstanceSpec(v string) *ListSplitSummaryResponseBodyData {
	s.InstanceSpec = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetInternetIp(v string) *ListSplitSummaryResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetIntranetIp(v string) *ListSplitSummaryResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetLineItemType(v string) *ListSplitSummaryResponseBodyData {
	s.LineItemType = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetListPrice(v string) *ListSplitSummaryResponseBodyData {
	s.ListPrice = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetListPriceUnit(v string) *ListSplitSummaryResponseBodyData {
	s.ListPriceUnit = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetListUnitPrice(v string) *ListSplitSummaryResponseBodyData {
	s.ListUnitPrice = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetProductCode(v string) *ListSplitSummaryResponseBodyData {
	s.ProductCode = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetProductName(v string) *ListSplitSummaryResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetPromotionAmount(v string) *ListSplitSummaryResponseBodyData {
	s.PromotionAmount = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetPurchaseConfiguration(v string) *ListSplitSummaryResponseBodyData {
	s.PurchaseConfiguration = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetRegionCode(v string) *ListSplitSummaryResponseBodyData {
	s.RegionCode = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetRegionName(v string) *ListSplitSummaryResponseBodyData {
	s.RegionName = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetRequireAmount(v string) *ListSplitSummaryResponseBodyData {
	s.RequireAmount = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetResourceGroup(v string) *ListSplitSummaryResponseBodyData {
	s.ResourceGroup = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetResourceId(v string) *ListSplitSummaryResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetResourceOwnerAccountId(v string) *ListSplitSummaryResponseBodyData {
	s.ResourceOwnerAccountId = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetResourceOwnerAccountName(v string) *ListSplitSummaryResponseBodyData {
	s.ResourceOwnerAccountName = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetResourcePurchaseAccountId(v string) *ListSplitSummaryResponseBodyData {
	s.ResourcePurchaseAccountId = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetResourcePurchaseAccountName(v string) *ListSplitSummaryResponseBodyData {
	s.ResourcePurchaseAccountName = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetResourceTag(v string) *ListSplitSummaryResponseBodyData {
	s.ResourceTag = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetSavingPlanDeductAmount(v string) *ListSplitSummaryResponseBodyData {
	s.SavingPlanDeductAmount = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetSavingPlanDeductListPrice(v string) *ListSplitSummaryResponseBodyData {
	s.SavingPlanDeductListPrice = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetServiceEntity(v string) *ListSplitSummaryResponseBodyData {
	s.ServiceEntity = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetSplitItemId(v string) *ListSplitSummaryResponseBodyData {
	s.SplitItemId = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetSplitItemName(v string) *ListSplitSummaryResponseBodyData {
	s.SplitItemName = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) SetZone(v string) *ListSplitSummaryResponseBodyData {
	s.Zone = &v
	return s
}

func (s *ListSplitSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
