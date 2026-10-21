// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetApiPriceResponseBody
	GetCode() *string
	SetMessage(v string) *GetApiPriceResponseBody
	GetMessage() *string
	SetPrice(v *GetApiPriceResponseBodyPrice) *GetApiPriceResponseBody
	GetPrice() *GetApiPriceResponseBodyPrice
	SetRequestId(v string) *GetApiPriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetApiPriceResponseBody
	GetSuccess() *bool
}

type GetApiPriceResponseBody struct {
	// The pricing result code. Valid values: Success, PricingNotRequired, PricingNotSupported, or an upstream pricing failure code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The server-side pricing result description. The language is specified by the language parameter in the request body and defaults to Chinese.
	//
	// example:
	//
	// Pricing succeeded
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pricing data. This field is omitted for PricingNotRequired and PricingNotSupported results that do not include pricing.
	//
	// if can be null:
	// true
	Price *GetApiPriceResponseBodyPrice `json:"price,omitempty" xml:"price,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether a valid pricing result is obtained. The value is true for normal pricing and when no pricing inquiry is required, and false when pricing is not supported or the pricing inquiry fails.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetApiPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApiPriceResponseBody) GoString() string {
	return s.String()
}

func (s *GetApiPriceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetApiPriceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetApiPriceResponseBody) GetPrice() *GetApiPriceResponseBodyPrice {
	return s.Price
}

func (s *GetApiPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApiPriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetApiPriceResponseBody) SetCode(v string) *GetApiPriceResponseBody {
	s.Code = &v
	return s
}

func (s *GetApiPriceResponseBody) SetMessage(v string) *GetApiPriceResponseBody {
	s.Message = &v
	return s
}

func (s *GetApiPriceResponseBody) SetPrice(v *GetApiPriceResponseBodyPrice) *GetApiPriceResponseBody {
	s.Price = v
	return s
}

func (s *GetApiPriceResponseBody) SetRequestId(v string) *GetApiPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApiPriceResponseBody) SetSuccess(v bool) *GetApiPriceResponseBody {
	s.Success = &v
	return s
}

func (s *GetApiPriceResponseBody) Validate() error {
	if s.Price != nil {
		if err := s.Price.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApiPriceResponseBodyPrice struct {
	// The common billing dimension for fixed fees, including totalAmount, originalAmount, and discountAmount. This is a canonical field, such as \\"Hour\\" or \\"1 Month\\", derived from the pricingUnit declared in the mapping. It is empty if not declared. This field describes only the dimension for fixed fees. The pay-as-you-go billing method components each have their own unit (see usageCharges[].unit).
	BillingUnit      *string  `json:"billingUnit,omitempty" xml:"billingUnit,omitempty"`
	CalculatedAmount *float32 `json:"calculatedAmount,omitempty" xml:"calculatedAmount,omitempty"`
	// The billing composition. This is a canonical field. Valid values: recurring_only (all time-based fees), usage_only (all usage-based fees, where a totalAmount of 0 indicates no fixed fees rather than free of charge, and actual costs depend on the unit price and actual usage in usageCharges), or mixed (both, which should be displayed separately as \\"X CNY/hour + Y CNY/GB\\"). This allows the caller to choose the display text without iterating through modules to determine the composition.
	ChargeComposition *string `json:"chargeComposition,omitempty" xml:"chargeComposition,omitempty"`
	// The details of each pricing call in composite or delta mode. The key is the priceCall ID, such as prepaid, postpaid, or a custom ID. This field is empty in single mode. The calculatedAmount is the aggregated result of each component based on the priceCalculation expression. The components field exposes the line items so that the caller can view the composition, such as instance fees versus storage fees.
	Components map[string]*PriceComponentsValue `json:"components,omitempty" xml:"components,omitempty"`
	Currency   *string                          `json:"currency,omitempty" xml:"currency,omitempty"`
	// The total discount amount for fixed fees. This is a canonical field calculated as originalAmount - totalAmount. It includes only discounts for fixed fees and excludes the pay-as-you-go billing method. The value is 0 when there is no discount and empty when the original price is missing.
	DiscountAmount *float32 `json:"discountAmount,omitempty" xml:"discountAmount,omitempty"`
	// The error code when upstream pricing fails, such as InvalidParameter or COMMODITY.INVALID_COMPONENT. This field is empty when success is true.
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error description when upstream pricing fails. This facilitates self-troubleshooting of specific parameter issues by the caller. This field is empty when success is true.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The total original price before discounts for fixed fees. This is a canonical field with the same scope as totalAmount. It includes only fixed fees and excludes the pay-as-you-go billing method. Refer to usageCharges[].originalPricePerUnit for the pre-discount unit price of the pay-as-you-go billing method. The difference between this value and totalAmount is the total discount for this inquiry. This field is empty when the upstream does not provide original price information.
	OriginalAmount *float32                                  `json:"originalAmount,omitempty" xml:"originalAmount,omitempty"`
	PriceSummary   *GetApiPriceResponseBodyPricePriceSummary `json:"priceSummary,omitempty" xml:"priceSummary,omitempty" type:"Struct"`
	PricingMode    *string                                   `json:"pricingMode,omitempty" xml:"pricingMode,omitempty"`
	Success        *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
	// The periodic payable amount for fixed fees in this pricing inquiry (canonical field). For the unit of measurement, see billingUnit. **This field includes only fixed fees**. Usage-based charges (such as traffic in CNY/GB or CUs in CNY/CU) have different units of measurement and cannot be added to this field. For the unit prices of usage-based charges, see usageCharges. To check whether both types of charges exist, see chargeComposition. In single mode, this value is derived from priceSummary.effectiveModuleSum (falling back to tradePrice if missing). In composite or delta mode, this value is derived from calculatedAmount. This field always has a value when success is true.
	TotalAmount *float32 `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
	// The request ID of the upstream pricing API (such as the RequestId of BSS GetSubscriptionPrice). This field is populated regardless of whether the request succeeds or fails, and is used for audit reconciliation and upstream log tracing.
	UpstreamRequestId *string `json:"upstreamRequestId,omitempty" xml:"upstreamRequestId,omitempty"`
	// The list of unit prices for usage-based billing modules (canonical field). Modules billed by actual usage, such as traffic and NAT CUs, have different units of measurement (such as CNY/GB or CNY/CU) from billingUnit and are not included in totalAmount. This field provides the unit price for each item. In single mode, the values are derived from priceSummary. In composite or delta mode, the values are concatenated from each component in the order of priceCall. This field is empty when the mapping contains no usage-based modules.
	UsageCharges []*GetApiPriceResponseBodyPriceUsageCharges `json:"usageCharges,omitempty" xml:"usageCharges,omitempty" type:"Repeated"`
}

func (s GetApiPriceResponseBodyPrice) String() string {
	return dara.Prettify(s)
}

func (s GetApiPriceResponseBodyPrice) GoString() string {
	return s.String()
}

func (s *GetApiPriceResponseBodyPrice) GetBillingUnit() *string {
	return s.BillingUnit
}

func (s *GetApiPriceResponseBodyPrice) GetCalculatedAmount() *float32 {
	return s.CalculatedAmount
}

func (s *GetApiPriceResponseBodyPrice) GetChargeComposition() *string {
	return s.ChargeComposition
}

func (s *GetApiPriceResponseBodyPrice) GetComponents() map[string]*PriceComponentsValue {
	return s.Components
}

func (s *GetApiPriceResponseBodyPrice) GetCurrency() *string {
	return s.Currency
}

func (s *GetApiPriceResponseBodyPrice) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *GetApiPriceResponseBodyPrice) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetApiPriceResponseBodyPrice) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetApiPriceResponseBodyPrice) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *GetApiPriceResponseBodyPrice) GetPriceSummary() *GetApiPriceResponseBodyPricePriceSummary {
	return s.PriceSummary
}

func (s *GetApiPriceResponseBodyPrice) GetPricingMode() *string {
	return s.PricingMode
}

func (s *GetApiPriceResponseBodyPrice) GetSuccess() *bool {
	return s.Success
}

func (s *GetApiPriceResponseBodyPrice) GetTotalAmount() *float32 {
	return s.TotalAmount
}

func (s *GetApiPriceResponseBodyPrice) GetUpstreamRequestId() *string {
	return s.UpstreamRequestId
}

func (s *GetApiPriceResponseBodyPrice) GetUsageCharges() []*GetApiPriceResponseBodyPriceUsageCharges {
	return s.UsageCharges
}

func (s *GetApiPriceResponseBodyPrice) SetBillingUnit(v string) *GetApiPriceResponseBodyPrice {
	s.BillingUnit = &v
	return s
}

func (s *GetApiPriceResponseBodyPrice) SetCalculatedAmount(v float32) *GetApiPriceResponseBodyPrice {
	s.CalculatedAmount = &v
	return s
}

func (s *GetApiPriceResponseBodyPrice) SetChargeComposition(v string) *GetApiPriceResponseBodyPrice {
	s.ChargeComposition = &v
	return s
}

func (s *GetApiPriceResponseBodyPrice) SetComponents(v map[string]*PriceComponentsValue) *GetApiPriceResponseBodyPrice {
	s.Components = v
	return s
}

func (s *GetApiPriceResponseBodyPrice) SetCurrency(v string) *GetApiPriceResponseBodyPrice {
	s.Currency = &v
	return s
}

func (s *GetApiPriceResponseBodyPrice) SetDiscountAmount(v float32) *GetApiPriceResponseBodyPrice {
	s.DiscountAmount = &v
	return s
}

func (s *GetApiPriceResponseBodyPrice) SetErrorCode(v string) *GetApiPriceResponseBodyPrice {
	s.ErrorCode = &v
	return s
}

func (s *GetApiPriceResponseBodyPrice) SetErrorMessage(v string) *GetApiPriceResponseBodyPrice {
	s.ErrorMessage = &v
	return s
}

func (s *GetApiPriceResponseBodyPrice) SetOriginalAmount(v float32) *GetApiPriceResponseBodyPrice {
	s.OriginalAmount = &v
	return s
}

func (s *GetApiPriceResponseBodyPrice) SetPriceSummary(v *GetApiPriceResponseBodyPricePriceSummary) *GetApiPriceResponseBodyPrice {
	s.PriceSummary = v
	return s
}

func (s *GetApiPriceResponseBodyPrice) SetPricingMode(v string) *GetApiPriceResponseBodyPrice {
	s.PricingMode = &v
	return s
}

func (s *GetApiPriceResponseBodyPrice) SetSuccess(v bool) *GetApiPriceResponseBodyPrice {
	s.Success = &v
	return s
}

func (s *GetApiPriceResponseBodyPrice) SetTotalAmount(v float32) *GetApiPriceResponseBodyPrice {
	s.TotalAmount = &v
	return s
}

func (s *GetApiPriceResponseBodyPrice) SetUpstreamRequestId(v string) *GetApiPriceResponseBodyPrice {
	s.UpstreamRequestId = &v
	return s
}

func (s *GetApiPriceResponseBodyPrice) SetUsageCharges(v []*GetApiPriceResponseBodyPriceUsageCharges) *GetApiPriceResponseBodyPrice {
	s.UsageCharges = v
	return s
}

func (s *GetApiPriceResponseBodyPrice) Validate() error {
	if s.PriceSummary != nil {
		if err := s.PriceSummary.Validate(); err != nil {
			return err
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

type GetApiPriceResponseBodyPricePriceSummary struct {
	// The billing composition. Valid values: recurring_only (all time-based fees), usage_only (all usage-based fees, where a totalAmount of 0 indicates no fixed fees rather than free of charge, and actual costs depend on the unit price and actual usage in usageCharges), or mixed (both, which should be displayed separately). This allows the caller to choose the display text without iterating through modules to determine the composition.
	ChargeComposition  *string                                            `json:"chargeComposition,omitempty" xml:"chargeComposition,omitempty"`
	Currency           *string                                            `json:"currency,omitempty" xml:"currency,omitempty"`
	EffectiveModuleSum *float32                                           `json:"effectiveModuleSum,omitempty" xml:"effectiveModuleSum,omitempty"`
	ModuleSum          *float32                                           `json:"moduleSum,omitempty" xml:"moduleSum,omitempty"`
	Modules            []*GetApiPriceResponseBodyPricePriceSummaryModules `json:"modules,omitempty" xml:"modules,omitempty" type:"Repeated"`
	// The total original price before discounts for this pricing inquiry. When originalPrice differs from tradePrice or moduleSum, it indicates that a discount was applied, such as a first-purchase discount or a promotional campaign.
	OriginalPrice *float32 `json:"originalPrice,omitempty" xml:"originalPrice,omitempty"`
	// The billing cycle unit, such as \\"Hour\\", \\"1 Month\\", \\"12 Month\\", or \\"GB·Month\\", which indicates the time dimension of moduleSum. This field is not returned in the BSS response and must be explicitly declared by the mapping maintainer in pricingTarget.pricingUnit. It is optional for subscription scenarios because the CLI infers it from Period and PeriodUnit. However, it is required for pay-as-you-go scenarios in the mapping; otherwise, the amount has no dimension.
	PricingUnit *string `json:"pricingUnit,omitempty" xml:"pricingUnit,omitempty"`
	// The pricing quantity from OpenAPI input parameters, such as Amount in ECS RunInstances. The moduleSum is the unit price per instance, and effectiveModuleSum = moduleSum × quantity is the total price. The caller should present the actual total price for N instances to the user based on this value to avoid mistaking the unit price for the total price.
	Quantity   *float32 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	TradePrice *float32 `json:"tradePrice,omitempty" xml:"tradePrice,omitempty"`
	// The list of unit prices for pay-as-you-go, usage-based modules, such as CNY/GB or CNY/CU. The amounts for these modules are not included in moduleSum or effectiveModuleSum. This field is empty when there are no usage-based modules.
	UsageCharges []*GetApiPriceResponseBodyPricePriceSummaryUsageCharges `json:"usageCharges,omitempty" xml:"usageCharges,omitempty" type:"Repeated"`
}

func (s GetApiPriceResponseBodyPricePriceSummary) String() string {
	return dara.Prettify(s)
}

func (s GetApiPriceResponseBodyPricePriceSummary) GoString() string {
	return s.String()
}

func (s *GetApiPriceResponseBodyPricePriceSummary) GetChargeComposition() *string {
	return s.ChargeComposition
}

func (s *GetApiPriceResponseBodyPricePriceSummary) GetCurrency() *string {
	return s.Currency
}

func (s *GetApiPriceResponseBodyPricePriceSummary) GetEffectiveModuleSum() *float32 {
	return s.EffectiveModuleSum
}

func (s *GetApiPriceResponseBodyPricePriceSummary) GetModuleSum() *float32 {
	return s.ModuleSum
}

func (s *GetApiPriceResponseBodyPricePriceSummary) GetModules() []*GetApiPriceResponseBodyPricePriceSummaryModules {
	return s.Modules
}

func (s *GetApiPriceResponseBodyPricePriceSummary) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *GetApiPriceResponseBodyPricePriceSummary) GetPricingUnit() *string {
	return s.PricingUnit
}

func (s *GetApiPriceResponseBodyPricePriceSummary) GetQuantity() *float32 {
	return s.Quantity
}

func (s *GetApiPriceResponseBodyPricePriceSummary) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *GetApiPriceResponseBodyPricePriceSummary) GetUsageCharges() []*GetApiPriceResponseBodyPricePriceSummaryUsageCharges {
	return s.UsageCharges
}

func (s *GetApiPriceResponseBodyPricePriceSummary) SetChargeComposition(v string) *GetApiPriceResponseBodyPricePriceSummary {
	s.ChargeComposition = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummary) SetCurrency(v string) *GetApiPriceResponseBodyPricePriceSummary {
	s.Currency = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummary) SetEffectiveModuleSum(v float32) *GetApiPriceResponseBodyPricePriceSummary {
	s.EffectiveModuleSum = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummary) SetModuleSum(v float32) *GetApiPriceResponseBodyPricePriceSummary {
	s.ModuleSum = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummary) SetModules(v []*GetApiPriceResponseBodyPricePriceSummaryModules) *GetApiPriceResponseBodyPricePriceSummary {
	s.Modules = v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummary) SetOriginalPrice(v float32) *GetApiPriceResponseBodyPricePriceSummary {
	s.OriginalPrice = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummary) SetPricingUnit(v string) *GetApiPriceResponseBodyPricePriceSummary {
	s.PricingUnit = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummary) SetQuantity(v float32) *GetApiPriceResponseBodyPricePriceSummary {
	s.Quantity = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummary) SetTradePrice(v float32) *GetApiPriceResponseBodyPricePriceSummary {
	s.TradePrice = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummary) SetUsageCharges(v []*GetApiPriceResponseBodyPricePriceSummaryUsageCharges) *GetApiPriceResponseBodyPricePriceSummary {
	s.UsageCharges = v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummary) Validate() error {
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

type GetApiPriceResponseBodyPricePriceSummaryModules struct {
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

func (s GetApiPriceResponseBodyPricePriceSummaryModules) String() string {
	return dara.Prettify(s)
}

func (s GetApiPriceResponseBodyPricePriceSummaryModules) GoString() string {
	return s.String()
}

func (s *GetApiPriceResponseBodyPricePriceSummaryModules) GetBillingMode() *string {
	return s.BillingMode
}

func (s *GetApiPriceResponseBodyPricePriceSummaryModules) GetCostAfterDiscount() *float32 {
	return s.CostAfterDiscount
}

func (s *GetApiPriceResponseBodyPricePriceSummaryModules) GetInvoiceDiscount() *float32 {
	return s.InvoiceDiscount
}

func (s *GetApiPriceResponseBodyPricePriceSummaryModules) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetApiPriceResponseBodyPricePriceSummaryModules) GetOriginalCost() *float32 {
	return s.OriginalCost
}

func (s *GetApiPriceResponseBodyPricePriceSummaryModules) GetQuantityUsedForQuote() *float32 {
	return s.QuantityUsedForQuote
}

func (s *GetApiPriceResponseBodyPricePriceSummaryModules) GetUnitPrice() *float32 {
	return s.UnitPrice
}

func (s *GetApiPriceResponseBodyPricePriceSummaryModules) GetUsageUnit() *string {
	return s.UsageUnit
}

func (s *GetApiPriceResponseBodyPricePriceSummaryModules) SetBillingMode(v string) *GetApiPriceResponseBodyPricePriceSummaryModules {
	s.BillingMode = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummaryModules) SetCostAfterDiscount(v float32) *GetApiPriceResponseBodyPricePriceSummaryModules {
	s.CostAfterDiscount = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummaryModules) SetInvoiceDiscount(v float32) *GetApiPriceResponseBodyPricePriceSummaryModules {
	s.InvoiceDiscount = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummaryModules) SetModuleCode(v string) *GetApiPriceResponseBodyPricePriceSummaryModules {
	s.ModuleCode = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummaryModules) SetOriginalCost(v float32) *GetApiPriceResponseBodyPricePriceSummaryModules {
	s.OriginalCost = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummaryModules) SetQuantityUsedForQuote(v float32) *GetApiPriceResponseBodyPricePriceSummaryModules {
	s.QuantityUsedForQuote = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummaryModules) SetUnitPrice(v float32) *GetApiPriceResponseBodyPricePriceSummaryModules {
	s.UnitPrice = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummaryModules) SetUsageUnit(v string) *GetApiPriceResponseBodyPricePriceSummaryModules {
	s.UsageUnit = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummaryModules) Validate() error {
	return dara.Validate(s)
}

type GetApiPriceResponseBodyPricePriceSummaryUsageCharges struct {
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

func (s GetApiPriceResponseBodyPricePriceSummaryUsageCharges) String() string {
	return dara.Prettify(s)
}

func (s GetApiPriceResponseBodyPricePriceSummaryUsageCharges) GoString() string {
	return s.String()
}

func (s *GetApiPriceResponseBodyPricePriceSummaryUsageCharges) GetAssumedQuantity() *float32 {
	return s.AssumedQuantity
}

func (s *GetApiPriceResponseBodyPricePriceSummaryUsageCharges) GetAssumedQuantityCost() *float32 {
	return s.AssumedQuantityCost
}

func (s *GetApiPriceResponseBodyPricePriceSummaryUsageCharges) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetApiPriceResponseBodyPricePriceSummaryUsageCharges) GetOriginalPricePerUnit() *float32 {
	return s.OriginalPricePerUnit
}

func (s *GetApiPriceResponseBodyPricePriceSummaryUsageCharges) GetPricePerUnit() *float32 {
	return s.PricePerUnit
}

func (s *GetApiPriceResponseBodyPricePriceSummaryUsageCharges) GetUnit() *string {
	return s.Unit
}

func (s *GetApiPriceResponseBodyPricePriceSummaryUsageCharges) SetAssumedQuantity(v float32) *GetApiPriceResponseBodyPricePriceSummaryUsageCharges {
	s.AssumedQuantity = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummaryUsageCharges) SetAssumedQuantityCost(v float32) *GetApiPriceResponseBodyPricePriceSummaryUsageCharges {
	s.AssumedQuantityCost = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummaryUsageCharges) SetModuleCode(v string) *GetApiPriceResponseBodyPricePriceSummaryUsageCharges {
	s.ModuleCode = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummaryUsageCharges) SetOriginalPricePerUnit(v float32) *GetApiPriceResponseBodyPricePriceSummaryUsageCharges {
	s.OriginalPricePerUnit = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummaryUsageCharges) SetPricePerUnit(v float32) *GetApiPriceResponseBodyPricePriceSummaryUsageCharges {
	s.PricePerUnit = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummaryUsageCharges) SetUnit(v string) *GetApiPriceResponseBodyPricePriceSummaryUsageCharges {
	s.Unit = &v
	return s
}

func (s *GetApiPriceResponseBodyPricePriceSummaryUsageCharges) Validate() error {
	return dara.Validate(s)
}

type GetApiPriceResponseBodyPriceUsageCharges struct {
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

func (s GetApiPriceResponseBodyPriceUsageCharges) String() string {
	return dara.Prettify(s)
}

func (s GetApiPriceResponseBodyPriceUsageCharges) GoString() string {
	return s.String()
}

func (s *GetApiPriceResponseBodyPriceUsageCharges) GetAssumedQuantity() *float32 {
	return s.AssumedQuantity
}

func (s *GetApiPriceResponseBodyPriceUsageCharges) GetAssumedQuantityCost() *float32 {
	return s.AssumedQuantityCost
}

func (s *GetApiPriceResponseBodyPriceUsageCharges) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetApiPriceResponseBodyPriceUsageCharges) GetOriginalPricePerUnit() *float32 {
	return s.OriginalPricePerUnit
}

func (s *GetApiPriceResponseBodyPriceUsageCharges) GetPricePerUnit() *float32 {
	return s.PricePerUnit
}

func (s *GetApiPriceResponseBodyPriceUsageCharges) GetUnit() *string {
	return s.Unit
}

func (s *GetApiPriceResponseBodyPriceUsageCharges) SetAssumedQuantity(v float32) *GetApiPriceResponseBodyPriceUsageCharges {
	s.AssumedQuantity = &v
	return s
}

func (s *GetApiPriceResponseBodyPriceUsageCharges) SetAssumedQuantityCost(v float32) *GetApiPriceResponseBodyPriceUsageCharges {
	s.AssumedQuantityCost = &v
	return s
}

func (s *GetApiPriceResponseBodyPriceUsageCharges) SetModuleCode(v string) *GetApiPriceResponseBodyPriceUsageCharges {
	s.ModuleCode = &v
	return s
}

func (s *GetApiPriceResponseBodyPriceUsageCharges) SetOriginalPricePerUnit(v float32) *GetApiPriceResponseBodyPriceUsageCharges {
	s.OriginalPricePerUnit = &v
	return s
}

func (s *GetApiPriceResponseBodyPriceUsageCharges) SetPricePerUnit(v float32) *GetApiPriceResponseBodyPriceUsageCharges {
	s.PricePerUnit = &v
	return s
}

func (s *GetApiPriceResponseBodyPriceUsageCharges) SetUnit(v string) *GetApiPriceResponseBodyPriceUsageCharges {
	s.Unit = &v
	return s
}

func (s *GetApiPriceResponseBodyPriceUsageCharges) Validate() error {
	return dara.Validate(s)
}
