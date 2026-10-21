// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPriceComponentsValue interface {
	dara.Model
	String() string
	GoString() string
	SetChargeComposition(v string) *PriceComponentsValue
	GetChargeComposition() *string
	SetCurrency(v string) *PriceComponentsValue
	GetCurrency() *string
	SetEffectiveModuleSum(v float32) *PriceComponentsValue
	GetEffectiveModuleSum() *float32
	SetModuleSum(v float32) *PriceComponentsValue
	GetModuleSum() *float32
	SetModules(v []*PriceComponentsValueModules) *PriceComponentsValue
	GetModules() []*PriceComponentsValueModules
	SetOriginalPrice(v float32) *PriceComponentsValue
	GetOriginalPrice() *float32
	SetPricingUnit(v string) *PriceComponentsValue
	GetPricingUnit() *string
	SetQuantity(v float32) *PriceComponentsValue
	GetQuantity() *float32
	SetTradePrice(v float32) *PriceComponentsValue
	GetTradePrice() *float32
	SetUsageCharges(v []*PriceComponentsValueUsageCharges) *PriceComponentsValue
	GetUsageCharges() []*PriceComponentsValueUsageCharges
}

type PriceComponentsValue struct {
	// The billing composition. Valid values: recurring_only (all time-based fees), usage_only (all usage-based fees, where a totalAmount of 0 indicates no fixed fees rather than free of charge, and actual costs depend on the unit price and actual usage in usageCharges), or mixed (both, which should be displayed separately). This allows the caller to choose the display text without iterating through modules to determine the composition.
	ChargeComposition  *string                        `json:"chargeComposition,omitempty" xml:"chargeComposition,omitempty"`
	Currency           *string                        `json:"currency,omitempty" xml:"currency,omitempty"`
	EffectiveModuleSum *float32                       `json:"effectiveModuleSum,omitempty" xml:"effectiveModuleSum,omitempty"`
	ModuleSum          *float32                       `json:"moduleSum,omitempty" xml:"moduleSum,omitempty"`
	Modules            []*PriceComponentsValueModules `json:"modules,omitempty" xml:"modules,omitempty" type:"Repeated"`
	// The total original price before discounts for this pricing inquiry. When originalPrice differs from tradePrice or moduleSum, it indicates that a discount was applied, such as a first-purchase discount or a promotional campaign.
	OriginalPrice *float32 `json:"originalPrice,omitempty" xml:"originalPrice,omitempty"`
	// The billing cycle unit, such as \\"Hour\\", \\"1 Month\\", \\"12 Month\\", or \\"GB·Month\\", which indicates the time dimension of moduleSum. This field is not returned in the BSS response and must be explicitly declared by the mapping maintainer in pricingTarget.pricingUnit. It is optional for subscription scenarios because the CLI infers it from Period and PeriodUnit. However, it is required for pay-as-you-go scenarios in the mapping; otherwise, the amount has no dimension.
	PricingUnit *string `json:"pricingUnit,omitempty" xml:"pricingUnit,omitempty"`
	// The pricing quantity from OpenAPI input parameters, such as Amount in ECS RunInstances. The moduleSum is the unit price per instance, and effectiveModuleSum = moduleSum × quantity is the total price. The caller should present the actual total price for N instances to the user based on this value to avoid mistaking the unit price for the total price.
	Quantity   *float32 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	TradePrice *float32 `json:"tradePrice,omitempty" xml:"tradePrice,omitempty"`
	// The list of unit prices for pay-as-you-go, usage-based modules, such as CNY/GB or CNY/CU. The amounts for these modules are not included in moduleSum or effectiveModuleSum. This field is empty when there are no usage-based modules.
	UsageCharges []*PriceComponentsValueUsageCharges `json:"usageCharges,omitempty" xml:"usageCharges,omitempty" type:"Repeated"`
}

func (s PriceComponentsValue) String() string {
	return dara.Prettify(s)
}

func (s PriceComponentsValue) GoString() string {
	return s.String()
}

func (s *PriceComponentsValue) GetChargeComposition() *string {
	return s.ChargeComposition
}

func (s *PriceComponentsValue) GetCurrency() *string {
	return s.Currency
}

func (s *PriceComponentsValue) GetEffectiveModuleSum() *float32 {
	return s.EffectiveModuleSum
}

func (s *PriceComponentsValue) GetModuleSum() *float32 {
	return s.ModuleSum
}

func (s *PriceComponentsValue) GetModules() []*PriceComponentsValueModules {
	return s.Modules
}

func (s *PriceComponentsValue) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *PriceComponentsValue) GetPricingUnit() *string {
	return s.PricingUnit
}

func (s *PriceComponentsValue) GetQuantity() *float32 {
	return s.Quantity
}

func (s *PriceComponentsValue) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *PriceComponentsValue) GetUsageCharges() []*PriceComponentsValueUsageCharges {
	return s.UsageCharges
}

func (s *PriceComponentsValue) SetChargeComposition(v string) *PriceComponentsValue {
	s.ChargeComposition = &v
	return s
}

func (s *PriceComponentsValue) SetCurrency(v string) *PriceComponentsValue {
	s.Currency = &v
	return s
}

func (s *PriceComponentsValue) SetEffectiveModuleSum(v float32) *PriceComponentsValue {
	s.EffectiveModuleSum = &v
	return s
}

func (s *PriceComponentsValue) SetModuleSum(v float32) *PriceComponentsValue {
	s.ModuleSum = &v
	return s
}

func (s *PriceComponentsValue) SetModules(v []*PriceComponentsValueModules) *PriceComponentsValue {
	s.Modules = v
	return s
}

func (s *PriceComponentsValue) SetOriginalPrice(v float32) *PriceComponentsValue {
	s.OriginalPrice = &v
	return s
}

func (s *PriceComponentsValue) SetPricingUnit(v string) *PriceComponentsValue {
	s.PricingUnit = &v
	return s
}

func (s *PriceComponentsValue) SetQuantity(v float32) *PriceComponentsValue {
	s.Quantity = &v
	return s
}

func (s *PriceComponentsValue) SetTradePrice(v float32) *PriceComponentsValue {
	s.TradePrice = &v
	return s
}

func (s *PriceComponentsValue) SetUsageCharges(v []*PriceComponentsValueUsageCharges) *PriceComponentsValue {
	s.UsageCharges = v
	return s
}

func (s *PriceComponentsValue) Validate() error {
	if s.Modules != nil {
		for _, item := range s.Modules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UsageCharges != nil {
		for _, item := range s.UsageCharges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PriceComponentsValueModules struct {
	// The billing dimension of this module, which is always explicitly output. Valid values: \\"recurring\\" for time-based fees, which are included in moduleSum and totalAmount; \\"usage\\" for usage-based fees, where costAfterDiscount is the usage unit price (CNY/usageUnit) and is not included in moduleSum, and the unit price is also provided in usageCharges.
	BillingMode       *string  `json:"billingMode,omitempty" xml:"billingMode,omitempty"`
	CostAfterDiscount *float32 `json:"costAfterDiscount,omitempty" xml:"costAfterDiscount,omitempty"`
	// The discount amount for this module. The value is calculated as invoiceDiscount = originalCost - costAfterDiscount and is precisely provided by the upstream BSS. Manual subtraction in tax-inclusive scenarios may cause precision loss.
	InvoiceDiscount *float32 `json:"invoiceDiscount,omitempty" xml:"invoiceDiscount,omitempty"`
	ModuleCode      *string  `json:"moduleCode,omitempty" xml:"moduleCode,omitempty"`
	OriginalCost    *float32 `json:"originalCost,omitempty" xml:"originalCost,omitempty"`
	// The estimated usage used for pricing this module. This field has a value only when the caller passes a usage greater than 1 through PricingContext. In this case, costAfterDiscount is the total cost for this usage rather than the unit price. Refer to usageCharges[].pricePerUnit for the normalized unit price. This field is empty when no estimated usage is passed.
	QuantityUsedForQuote *float32 `json:"quantityUsedForQuote,omitempty" xml:"quantityUsedForQuote,omitempty"`
	// The unit price of this module per billing unit, such as CNY/hour, CNY/month, or CNY/unit. This is used to explain the composition of costAfterDiscount in multi-quantity or multi-period scenarios.
	UnitPrice *float32 `json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
	// The usage unit, such as GB or CU. This field has a value when billingMode is \\"usage\\" and indicates the dimension of costAfterDiscount.
	UsageUnit *string `json:"usageUnit,omitempty" xml:"usageUnit,omitempty"`
}

func (s PriceComponentsValueModules) String() string {
	return dara.Prettify(s)
}

func (s PriceComponentsValueModules) GoString() string {
	return s.String()
}

func (s *PriceComponentsValueModules) GetBillingMode() *string {
	return s.BillingMode
}

func (s *PriceComponentsValueModules) GetCostAfterDiscount() *float32 {
	return s.CostAfterDiscount
}

func (s *PriceComponentsValueModules) GetInvoiceDiscount() *float32 {
	return s.InvoiceDiscount
}

func (s *PriceComponentsValueModules) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *PriceComponentsValueModules) GetOriginalCost() *float32 {
	return s.OriginalCost
}

func (s *PriceComponentsValueModules) GetQuantityUsedForQuote() *float32 {
	return s.QuantityUsedForQuote
}

func (s *PriceComponentsValueModules) GetUnitPrice() *float32 {
	return s.UnitPrice
}

func (s *PriceComponentsValueModules) GetUsageUnit() *string {
	return s.UsageUnit
}

func (s *PriceComponentsValueModules) SetBillingMode(v string) *PriceComponentsValueModules {
	s.BillingMode = &v
	return s
}

func (s *PriceComponentsValueModules) SetCostAfterDiscount(v float32) *PriceComponentsValueModules {
	s.CostAfterDiscount = &v
	return s
}

func (s *PriceComponentsValueModules) SetInvoiceDiscount(v float32) *PriceComponentsValueModules {
	s.InvoiceDiscount = &v
	return s
}

func (s *PriceComponentsValueModules) SetModuleCode(v string) *PriceComponentsValueModules {
	s.ModuleCode = &v
	return s
}

func (s *PriceComponentsValueModules) SetOriginalCost(v float32) *PriceComponentsValueModules {
	s.OriginalCost = &v
	return s
}

func (s *PriceComponentsValueModules) SetQuantityUsedForQuote(v float32) *PriceComponentsValueModules {
	s.QuantityUsedForQuote = &v
	return s
}

func (s *PriceComponentsValueModules) SetUnitPrice(v float32) *PriceComponentsValueModules {
	s.UnitPrice = &v
	return s
}

func (s *PriceComponentsValueModules) SetUsageUnit(v string) *PriceComponentsValueModules {
	s.UsageUnit = &v
	return s
}

func (s *PriceComponentsValueModules) Validate() error {
	return dara.Validate(s)
}

type PriceComponentsValueUsageCharges struct {
	// The estimated usage passed by the caller through PricingContext. This field has a value only when it is greater than 1. In this case, this entry also provides the assumedQuantityCost calculated based on this assumed usage.
	AssumedQuantity *float32 `json:"assumedQuantity,omitempty" xml:"assumedQuantity,omitempty"`
	// The module cost calculated based on assumedQuantity, which equals pricePerUnit × assumedQuantity. This cost is not included in totalAmount because it is based on the assumed usage reported by the caller and is not a cost necessarily incurred by this call.
	AssumedQuantityCost *float32 `json:"assumedQuantityCost,omitempty" xml:"assumedQuantityCost,omitempty"`
	// The billing module code, such as TrafficOut or NatCu. This corresponds to the moduleCode in modules.
	ModuleCode *string `json:"moduleCode,omitempty" xml:"moduleCode,omitempty"`
	// The original unit price before discounts, in CNY/unit, which is also normalized. This field is empty when not provided by the upstream. The difference between this value and pricePerUnit is the unit price discount for this module.
	OriginalPricePerUnit *float32 `json:"originalPricePerUnit,omitempty" xml:"originalPricePerUnit,omitempty"`
	// The discounted unit price, in CNY/unit. The actual cost equals pricePerUnit × actual usage. Because the usage is unknown during pricing, it is not included in totalAmount. If the caller passed an estimated usage, the total returned by the upstream is normalized back to the unit price based on that usage.
	PricePerUnit *float32 `json:"pricePerUnit,omitempty" xml:"pricePerUnit,omitempty"`
	// The usage unit, such as GB, CU, or 10,000 requests. This comes from the module-level billing.unit annotation in the mapping. The dimension of pricePerUnit is CNY/unit.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s PriceComponentsValueUsageCharges) String() string {
	return dara.Prettify(s)
}

func (s PriceComponentsValueUsageCharges) GoString() string {
	return s.String()
}

func (s *PriceComponentsValueUsageCharges) GetAssumedQuantity() *float32 {
	return s.AssumedQuantity
}

func (s *PriceComponentsValueUsageCharges) GetAssumedQuantityCost() *float32 {
	return s.AssumedQuantityCost
}

func (s *PriceComponentsValueUsageCharges) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *PriceComponentsValueUsageCharges) GetOriginalPricePerUnit() *float32 {
	return s.OriginalPricePerUnit
}

func (s *PriceComponentsValueUsageCharges) GetPricePerUnit() *float32 {
	return s.PricePerUnit
}

func (s *PriceComponentsValueUsageCharges) GetUnit() *string {
	return s.Unit
}

func (s *PriceComponentsValueUsageCharges) SetAssumedQuantity(v float32) *PriceComponentsValueUsageCharges {
	s.AssumedQuantity = &v
	return s
}

func (s *PriceComponentsValueUsageCharges) SetAssumedQuantityCost(v float32) *PriceComponentsValueUsageCharges {
	s.AssumedQuantityCost = &v
	return s
}

func (s *PriceComponentsValueUsageCharges) SetModuleCode(v string) *PriceComponentsValueUsageCharges {
	s.ModuleCode = &v
	return s
}

func (s *PriceComponentsValueUsageCharges) SetOriginalPricePerUnit(v float32) *PriceComponentsValueUsageCharges {
	s.OriginalPricePerUnit = &v
	return s
}

func (s *PriceComponentsValueUsageCharges) SetPricePerUnit(v float32) *PriceComponentsValueUsageCharges {
	s.PricePerUnit = &v
	return s
}

func (s *PriceComponentsValueUsageCharges) SetUnit(v string) *PriceComponentsValueUsageCharges {
	s.Unit = &v
	return s
}

func (s *PriceComponentsValueUsageCharges) Validate() error {
	return dara.Validate(s)
}
