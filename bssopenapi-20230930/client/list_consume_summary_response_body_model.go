// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumeSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListConsumeSummaryResponseBodyData) *ListConsumeSummaryResponseBody
	GetData() []*ListConsumeSummaryResponseBodyData
	SetMaxResults(v int32) *ListConsumeSummaryResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListConsumeSummaryResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListConsumeSummaryResponseBody
	GetRequestId() *string
}

type ListConsumeSummaryResponseBody struct {
	// The consumption summary records returned in this request. An empty array is returned if no matching records are found.
	//
	// example:
	//
	// []
	Data []*ListConsumeSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The maximum number of records per page used in this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token for the next page. If a non-empty value is returned, pass it as is to the next request and keep other query conditions unchanged. If it is not returned or is empty, pagination ends.
	//
	// example:
	//
	// example-next-token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The unique identifier of the request, used for troubleshooting.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConsumeSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConsumeSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsumeSummaryResponseBody) GetData() []*ListConsumeSummaryResponseBodyData {
	return s.Data
}

func (s *ListConsumeSummaryResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListConsumeSummaryResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConsumeSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConsumeSummaryResponseBody) SetData(v []*ListConsumeSummaryResponseBodyData) *ListConsumeSummaryResponseBody {
	s.Data = v
	return s
}

func (s *ListConsumeSummaryResponseBody) SetMaxResults(v int32) *ListConsumeSummaryResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListConsumeSummaryResponseBody) SetNextToken(v string) *ListConsumeSummaryResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListConsumeSummaryResponseBody) SetRequestId(v string) *ListConsumeSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsumeSummaryResponseBody) Validate() error {
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

type ListConsumeSummaryResponseBodyData struct {
	// The amount after coupon deduction, which is the result of deducting coupons from the amount after discount. When all related fields are available, it equals AmountAfterPromotionDeduction - CouponAmount.
	//
	// example:
	//
	// 65
	AmountAfterCouponDeduction *string `json:"AmountAfterCouponDeduction,omitempty" xml:"AmountAfterCouponDeduction,omitempty"`
	// The amount after discount, which represents the amount after applying discounts such as website promotions and contract discounts. Coupons are not deducted at this point. When all related fields are available, it equals AmountAfterSavingPlanDeduction - PromotionAmount.
	//
	// example:
	//
	// 70
	AmountAfterPromotionDeduction *string `json:"AmountAfterPromotionDeduction,omitempty" xml:"AmountAfterPromotionDeduction,omitempty"`
	// The amount after savings plan deduction, which is the remaining amount after deducting the list price portion covered by the savings plan. When all related fields are available, it equals ListPrice - SavingPlanDeductListPrice.
	//
	// example:
	//
	// 80
	AmountAfterSavingPlanDeduction *string `json:"AmountAfterSavingPlanDeduction,omitempty" xml:"AmountAfterSavingPlanDeduction,omitempty"`
	// The billing date, which is the consumption date of the cloud product, in the format of yyyy-MM-dd. It has a value only when summarized by day.
	//
	// example:
	//
	// 2026-10-01
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The billing formula, which describes how the total list price is calculated. It should be understood in conjunction with the list price on the Alibaba Cloud website, usage, duration, and other billing factors.
	//
	// example:
	//
	// List price × Usage × Duration
	BillingFormula *string `json:"BillingFormula,omitempty" xml:"BillingFormula,omitempty"`
	// The billing item code, which identifies the billing item to which the fee belongs.
	//
	// example:
	//
	// example_billing_item
	BillingItemCode *string `json:"BillingItemCode,omitempty" xml:"BillingItemCode,omitempty"`
	// The billing item name, which is the display name of the billing item to which the fee belongs.
	//
	// example:
	//
	// Sample billing item
	BillingItemName *string `json:"BillingItemName,omitempty" xml:"BillingItemName,omitempty"`
	// The billing month, which represents the billing month to which the fee belongs, in the format of yyyy-MM.
	//
	// example:
	//
	// 2026-10
	BillingMonth *string `json:"BillingMonth,omitempty" xml:"BillingMonth,omitempty"`
	// The usage that is ultimately used for fee calculation. If usage deduction is applied, this value is the usage before deduction minus the usage deducted by resource plans, free quotas, and other deductions.
	//
	// example:
	//
	// 100
	BillingQuantity *string `json:"BillingQuantity,omitempty" xml:"BillingQuantity,omitempty"`
	// The usage unit, used in conjunction with BillingQuantity.
	//
	// example:
	//
	// GB
	BillingQuantityUnit *string `json:"BillingQuantityUnit,omitempty" xml:"BillingQuantityUnit,omitempty"`
	// The duration, which represents the purchase or usage duration involved in billing. It has a value only when the billing formula involves time. If duration deduction is applied, this value is the duration before deduction minus the deducted duration.
	//
	// example:
	//
	// 1
	ChargeDuration *string `json:"ChargeDuration,omitempty" xml:"ChargeDuration,omitempty"`
	// The duration unit, used in conjunction with ChargeDuration.
	//
	// example:
	//
	// Hour
	ChargeDurationUnit *string `json:"ChargeDurationUnit,omitempty" xml:"ChargeDurationUnit,omitempty"`
	// The billing method. The display name such as prepaid or pay-as-you-go is returned.
	//
	// example:
	//
	// Pay-as-you-go
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The commodity code, used to uniquely identify a specific commodity under a cloud product.
	//
	// example:
	//
	// ecs
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The commodity name, which is the display name of a specific commodity under a cloud product.
	//
	// example:
	//
	// Sample commodity
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	// The transaction type, which describes the transaction behavior that generated the bill, such as new purchase, renewal, upgrade, downgrade, unsubscription, pay-as-you-go, or adjustment. The display name is returned.
	//
	// example:
	//
	// Pay-as-you-go
	ConsumeType *string `json:"ConsumeType,omitempty" xml:"ConsumeType,omitempty"`
	// The coupon deduction amount, which is the fee covered by coupons. This amount cannot be used for invoicing.
	//
	// example:
	//
	// 5
	CouponAmount *string `json:"CouponAmount,omitempty" xml:"CouponAmount,omitempty"`
	// The currency, which represents the currency used for the list price and amounts in this record. For example, CNY indicates Chinese Yuan.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The summary record identifier, used as an opaque string. It can be empty.
	//
	// example:
	//
	// example-summary-key
	HashKey *string `json:"HashKey,omitempty" xml:"HashKey,omitempty"`
	// The instance ID at the billing granularity, which represents the instance identity used to aggregate fees when a cloud product generates a bill. The billing granularity may differ from the resource management granularity, so it may be different from ResourceId.
	//
	// example:
	//
	// i-example001
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The resource name, which is the name set by the user for the cloud resource. Whether it is returned depends on whether the product supports querying and whether the resource information is available.
	//
	// example:
	//
	// example-resource
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The specification, used to describe the specification information corresponding to the billing item.
	//
	// example:
	//
	// example.spec
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The line item type, which describes the nature of the fee, such as subscription prepaid fee, subscription postpaid fee, subscription unsubscription fee, pay-as-you-go fee for cloud resources, adjustment fee, or rounding fee. The display name is returned.
	//
	// example:
	//
	// Pay-as-you-go
	LineItemType *string `json:"LineItemType,omitempty" xml:"LineItemType,omitempty"`
	// The total list price, which represents the total amount calculated based on the list price for the purchased cloud resource configuration. Savings plan deductions, discounts, and coupons are not deducted yet. Note that ListPrice in this API is the total amount. For the list unit price, see ListUnitPrice.
	//
	// example:
	//
	// 100
	ListPrice *string `json:"ListPrice,omitempty" xml:"ListPrice,omitempty"`
	// The price unit, which is the unit corresponding to the list unit price ListUnitPrice, such as CNY/GB.
	//
	// example:
	//
	// CNY/GB/Hour
	ListPriceUnit *string `json:"ListPriceUnit,omitempty" xml:"ListPriceUnit,omitempty"`
	// The list price on the Alibaba Cloud website, which is the list unit price of the commodity. It differs in meaning from the total list price ListPrice. In scenarios such as tiered pricing, you cannot recalculate the fee by simply multiplying a single unit price by the total usage.
	//
	// example:
	//
	// 1
	ListUnitPrice *string `json:"ListUnitPrice,omitempty" xml:"ListUnitPrice,omitempty"`
	// The product code, used to uniquely identify a specific cloud product.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The product name, which is the display name of the cloud product.
	//
	// example:
	//
	// Elastic Compute Service
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The discount amount, which represents the total discounts such as website promotions and contract discounts. Coupon deductions are separately recorded in CouponAmount.
	//
	// example:
	//
	// 10
	PromotionAmount *string `json:"PromotionAmount,omitempty" xml:"PromotionAmount,omitempty"`
	// The configuration, used to describe the configuration options corresponding to the billing item.
	//
	// example:
	//
	// Sample configuration
	PurchaseConfiguration *string `json:"PurchaseConfiguration,omitempty" xml:"PurchaseConfiguration,omitempty"`
	// The region code, which represents the code of the logical region where the cloud resource is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The region, which represents the display name of the region where the cloud resource is located. For products sold by region, it usually corresponds to the region selected when placing the order.
	//
	// example:
	//
	// China (Hangzhou)
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The payable amount, which is the fee to be paid after savings plan deductions, discounts, and coupon deductions. In the new version of bills, this corresponds to the payable amount (tax-inclusive) definition. For fee verification, refer to the actual return value of this field.
	//
	// example:
	//
	// 65
	RequireAmount *string `json:"RequireAmount,omitempty" xml:"RequireAmount,omitempty"`
	// The resource ID, used to uniquely identify a cloud resource in the corresponding cloud product or service.
	//
	// example:
	//
	// i-example001
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource ownership account ID, which is the ID of the account that actually uses the cloud resource instance. It may be different from the purchasing account.
	//
	// example:
	//
	// 1234567890123456
	ResourceOwnerAccountId *string `json:"ResourceOwnerAccountId,omitempty" xml:"ResourceOwnerAccountId,omitempty"`
	// The resource ownership account, which is the name of the account that actually uses the cloud resource instance. It may be different from the purchasing account.
	//
	// example:
	//
	// owner@example.com
	ResourceOwnerAccountName *string `json:"ResourceOwnerAccountName,omitempty" xml:"ResourceOwnerAccountName,omitempty"`
	// The ID of the resource purchasing account, which is the ID of the account that placed the order to purchase the cloud resource instance.
	//
	// example:
	//
	// 1234567890123456
	ResourcePurchaseAccountId *string `json:"ResourcePurchaseAccountId,omitempty" xml:"ResourcePurchaseAccountId,omitempty"`
	// The resource purchasing account, which is the name of the account that placed the order to purchase the cloud resource instance.
	//
	// example:
	//
	// buyer@example.com
	ResourcePurchaseAccountName *string `json:"ResourcePurchaseAccountName,omitempty" xml:"ResourcePurchaseAccountName,omitempty"`
	// The savings plan deduction amount, which is the committed spend of the savings plan consumed to complete this deduction. It is not the list price amount of the deducted resource and cannot replace SavingPlanDeductListPrice for deduction from the total list price.
	//
	// example:
	//
	// 8
	SavingPlanDeductAmount *string `json:"SavingPlanDeductAmount,omitempty" xml:"SavingPlanDeductAmount,omitempty"`
	// The savings plan deduction list price, which is the amount calculated based on the list price for the pay-as-you-go resource fees deducted by the savings plan.
	//
	// example:
	//
	// 20
	SavingPlanDeductListPrice *string `json:"SavingPlanDeductListPrice,omitempty" xml:"SavingPlanDeductListPrice,omitempty"`
	// The service entity, which is the name of the company that actually provides the product or service. Different products may correspond to different companies.
	//
	// example:
	//
	// Sample service entity
	ServiceEntity *string `json:"ServiceEntity,omitempty" xml:"ServiceEntity,omitempty"`
}

func (s ListConsumeSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListConsumeSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConsumeSummaryResponseBodyData) GetAmountAfterCouponDeduction() *string {
	return s.AmountAfterCouponDeduction
}

func (s *ListConsumeSummaryResponseBodyData) GetAmountAfterPromotionDeduction() *string {
	return s.AmountAfterPromotionDeduction
}

func (s *ListConsumeSummaryResponseBodyData) GetAmountAfterSavingPlanDeduction() *string {
	return s.AmountAfterSavingPlanDeduction
}

func (s *ListConsumeSummaryResponseBodyData) GetBillingDate() *string {
	return s.BillingDate
}

func (s *ListConsumeSummaryResponseBodyData) GetBillingFormula() *string {
	return s.BillingFormula
}

func (s *ListConsumeSummaryResponseBodyData) GetBillingItemCode() *string {
	return s.BillingItemCode
}

func (s *ListConsumeSummaryResponseBodyData) GetBillingItemName() *string {
	return s.BillingItemName
}

func (s *ListConsumeSummaryResponseBodyData) GetBillingMonth() *string {
	return s.BillingMonth
}

func (s *ListConsumeSummaryResponseBodyData) GetBillingQuantity() *string {
	return s.BillingQuantity
}

func (s *ListConsumeSummaryResponseBodyData) GetBillingQuantityUnit() *string {
	return s.BillingQuantityUnit
}

func (s *ListConsumeSummaryResponseBodyData) GetChargeDuration() *string {
	return s.ChargeDuration
}

func (s *ListConsumeSummaryResponseBodyData) GetChargeDurationUnit() *string {
	return s.ChargeDurationUnit
}

func (s *ListConsumeSummaryResponseBodyData) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListConsumeSummaryResponseBodyData) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListConsumeSummaryResponseBodyData) GetCommodityName() *string {
	return s.CommodityName
}

func (s *ListConsumeSummaryResponseBodyData) GetConsumeType() *string {
	return s.ConsumeType
}

func (s *ListConsumeSummaryResponseBodyData) GetCouponAmount() *string {
	return s.CouponAmount
}

func (s *ListConsumeSummaryResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *ListConsumeSummaryResponseBodyData) GetHashKey() *string {
	return s.HashKey
}

func (s *ListConsumeSummaryResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListConsumeSummaryResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListConsumeSummaryResponseBodyData) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *ListConsumeSummaryResponseBodyData) GetLineItemType() *string {
	return s.LineItemType
}

func (s *ListConsumeSummaryResponseBodyData) GetListPrice() *string {
	return s.ListPrice
}

func (s *ListConsumeSummaryResponseBodyData) GetListPriceUnit() *string {
	return s.ListPriceUnit
}

func (s *ListConsumeSummaryResponseBodyData) GetListUnitPrice() *string {
	return s.ListUnitPrice
}

func (s *ListConsumeSummaryResponseBodyData) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListConsumeSummaryResponseBodyData) GetProductName() *string {
	return s.ProductName
}

func (s *ListConsumeSummaryResponseBodyData) GetPromotionAmount() *string {
	return s.PromotionAmount
}

func (s *ListConsumeSummaryResponseBodyData) GetPurchaseConfiguration() *string {
	return s.PurchaseConfiguration
}

func (s *ListConsumeSummaryResponseBodyData) GetRegionCode() *string {
	return s.RegionCode
}

func (s *ListConsumeSummaryResponseBodyData) GetRegionName() *string {
	return s.RegionName
}

func (s *ListConsumeSummaryResponseBodyData) GetRequireAmount() *string {
	return s.RequireAmount
}

func (s *ListConsumeSummaryResponseBodyData) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListConsumeSummaryResponseBodyData) GetResourceOwnerAccountId() *string {
	return s.ResourceOwnerAccountId
}

func (s *ListConsumeSummaryResponseBodyData) GetResourceOwnerAccountName() *string {
	return s.ResourceOwnerAccountName
}

func (s *ListConsumeSummaryResponseBodyData) GetResourcePurchaseAccountId() *string {
	return s.ResourcePurchaseAccountId
}

func (s *ListConsumeSummaryResponseBodyData) GetResourcePurchaseAccountName() *string {
	return s.ResourcePurchaseAccountName
}

func (s *ListConsumeSummaryResponseBodyData) GetSavingPlanDeductAmount() *string {
	return s.SavingPlanDeductAmount
}

func (s *ListConsumeSummaryResponseBodyData) GetSavingPlanDeductListPrice() *string {
	return s.SavingPlanDeductListPrice
}

func (s *ListConsumeSummaryResponseBodyData) GetServiceEntity() *string {
	return s.ServiceEntity
}

func (s *ListConsumeSummaryResponseBodyData) SetAmountAfterCouponDeduction(v string) *ListConsumeSummaryResponseBodyData {
	s.AmountAfterCouponDeduction = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetAmountAfterPromotionDeduction(v string) *ListConsumeSummaryResponseBodyData {
	s.AmountAfterPromotionDeduction = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetAmountAfterSavingPlanDeduction(v string) *ListConsumeSummaryResponseBodyData {
	s.AmountAfterSavingPlanDeduction = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetBillingDate(v string) *ListConsumeSummaryResponseBodyData {
	s.BillingDate = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetBillingFormula(v string) *ListConsumeSummaryResponseBodyData {
	s.BillingFormula = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetBillingItemCode(v string) *ListConsumeSummaryResponseBodyData {
	s.BillingItemCode = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetBillingItemName(v string) *ListConsumeSummaryResponseBodyData {
	s.BillingItemName = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetBillingMonth(v string) *ListConsumeSummaryResponseBodyData {
	s.BillingMonth = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetBillingQuantity(v string) *ListConsumeSummaryResponseBodyData {
	s.BillingQuantity = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetBillingQuantityUnit(v string) *ListConsumeSummaryResponseBodyData {
	s.BillingQuantityUnit = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetChargeDuration(v string) *ListConsumeSummaryResponseBodyData {
	s.ChargeDuration = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetChargeDurationUnit(v string) *ListConsumeSummaryResponseBodyData {
	s.ChargeDurationUnit = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetChargeType(v string) *ListConsumeSummaryResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetCommodityCode(v string) *ListConsumeSummaryResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetCommodityName(v string) *ListConsumeSummaryResponseBodyData {
	s.CommodityName = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetConsumeType(v string) *ListConsumeSummaryResponseBodyData {
	s.ConsumeType = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetCouponAmount(v string) *ListConsumeSummaryResponseBodyData {
	s.CouponAmount = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetCurrency(v string) *ListConsumeSummaryResponseBodyData {
	s.Currency = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetHashKey(v string) *ListConsumeSummaryResponseBodyData {
	s.HashKey = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetInstanceId(v string) *ListConsumeSummaryResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetInstanceName(v string) *ListConsumeSummaryResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetInstanceSpec(v string) *ListConsumeSummaryResponseBodyData {
	s.InstanceSpec = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetLineItemType(v string) *ListConsumeSummaryResponseBodyData {
	s.LineItemType = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetListPrice(v string) *ListConsumeSummaryResponseBodyData {
	s.ListPrice = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetListPriceUnit(v string) *ListConsumeSummaryResponseBodyData {
	s.ListPriceUnit = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetListUnitPrice(v string) *ListConsumeSummaryResponseBodyData {
	s.ListUnitPrice = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetProductCode(v string) *ListConsumeSummaryResponseBodyData {
	s.ProductCode = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetProductName(v string) *ListConsumeSummaryResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetPromotionAmount(v string) *ListConsumeSummaryResponseBodyData {
	s.PromotionAmount = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetPurchaseConfiguration(v string) *ListConsumeSummaryResponseBodyData {
	s.PurchaseConfiguration = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetRegionCode(v string) *ListConsumeSummaryResponseBodyData {
	s.RegionCode = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetRegionName(v string) *ListConsumeSummaryResponseBodyData {
	s.RegionName = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetRequireAmount(v string) *ListConsumeSummaryResponseBodyData {
	s.RequireAmount = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetResourceId(v string) *ListConsumeSummaryResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetResourceOwnerAccountId(v string) *ListConsumeSummaryResponseBodyData {
	s.ResourceOwnerAccountId = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetResourceOwnerAccountName(v string) *ListConsumeSummaryResponseBodyData {
	s.ResourceOwnerAccountName = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetResourcePurchaseAccountId(v string) *ListConsumeSummaryResponseBodyData {
	s.ResourcePurchaseAccountId = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetResourcePurchaseAccountName(v string) *ListConsumeSummaryResponseBodyData {
	s.ResourcePurchaseAccountName = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetSavingPlanDeductAmount(v string) *ListConsumeSummaryResponseBodyData {
	s.SavingPlanDeductAmount = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetSavingPlanDeductListPrice(v string) *ListConsumeSummaryResponseBodyData {
	s.SavingPlanDeductListPrice = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) SetServiceEntity(v string) *ListConsumeSummaryResponseBodyData {
	s.ServiceEntity = &v
	return s
}

func (s *ListConsumeSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
