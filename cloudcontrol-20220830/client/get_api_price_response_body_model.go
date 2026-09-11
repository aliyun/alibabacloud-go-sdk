// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrice(v *GetApiPriceResponseBodyPrice) *GetApiPriceResponseBody
	GetPrice() *GetApiPriceResponseBodyPrice
	SetRequestId(v string) *GetApiPriceResponseBody
	GetRequestId() *string
}

type GetApiPriceResponseBody struct {
	Price *GetApiPriceResponseBodyPrice `json:"price,omitempty" xml:"price,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetApiPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApiPriceResponseBody) GoString() string {
	return s.String()
}

func (s *GetApiPriceResponseBody) GetPrice() *GetApiPriceResponseBodyPrice {
	return s.Price
}

func (s *GetApiPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApiPriceResponseBody) SetPrice(v *GetApiPriceResponseBodyPrice) *GetApiPriceResponseBody {
	s.Price = v
	return s
}

func (s *GetApiPriceResponseBody) SetRequestId(v string) *GetApiPriceResponseBody {
	s.RequestId = &v
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
	BillingUnit       *string                                     `json:"billingUnit,omitempty" xml:"billingUnit,omitempty"`
	CalculatedAmount  *float32                                    `json:"calculatedAmount,omitempty" xml:"calculatedAmount,omitempty"`
	ChargeComposition *string                                     `json:"chargeComposition,omitempty" xml:"chargeComposition,omitempty"`
	Components        map[string]*PriceComponentsValue            `json:"components,omitempty" xml:"components,omitempty"`
	Currency          *string                                     `json:"currency,omitempty" xml:"currency,omitempty"`
	DiscountAmount    *float32                                    `json:"discountAmount,omitempty" xml:"discountAmount,omitempty"`
	ErrorCode         *string                                     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage      *string                                     `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	OriginalAmount    *float32                                    `json:"originalAmount,omitempty" xml:"originalAmount,omitempty"`
	PriceSummary      *GetApiPriceResponseBodyPricePriceSummary   `json:"priceSummary,omitempty" xml:"priceSummary,omitempty" type:"Struct"`
	PricingMode       *string                                     `json:"pricingMode,omitempty" xml:"pricingMode,omitempty"`
	Success           *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
	TotalAmount       *float32                                    `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
	UpstreamRequestId *string                                     `json:"upstreamRequestId,omitempty" xml:"upstreamRequestId,omitempty"`
	UsageCharges      []*GetApiPriceResponseBodyPriceUsageCharges `json:"usageCharges,omitempty" xml:"usageCharges,omitempty" type:"Repeated"`
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
	ChargeComposition  *string                                                 `json:"chargeComposition,omitempty" xml:"chargeComposition,omitempty"`
	Currency           *string                                                 `json:"currency,omitempty" xml:"currency,omitempty"`
	EffectiveModuleSum *float32                                                `json:"effectiveModuleSum,omitempty" xml:"effectiveModuleSum,omitempty"`
	ModuleSum          *float32                                                `json:"moduleSum,omitempty" xml:"moduleSum,omitempty"`
	Modules            []*GetApiPriceResponseBodyPricePriceSummaryModules      `json:"modules,omitempty" xml:"modules,omitempty" type:"Repeated"`
	OriginalPrice      *float32                                                `json:"originalPrice,omitempty" xml:"originalPrice,omitempty"`
	PricingUnit        *string                                                 `json:"pricingUnit,omitempty" xml:"pricingUnit,omitempty"`
	Quantity           *float32                                                `json:"quantity,omitempty" xml:"quantity,omitempty"`
	TradePrice         *float32                                                `json:"tradePrice,omitempty" xml:"tradePrice,omitempty"`
	UsageCharges       []*GetApiPriceResponseBodyPricePriceSummaryUsageCharges `json:"usageCharges,omitempty" xml:"usageCharges,omitempty" type:"Repeated"`
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
	BillingMode          *string  `json:"billingMode,omitempty" xml:"billingMode,omitempty"`
	CostAfterDiscount    *float32 `json:"costAfterDiscount,omitempty" xml:"costAfterDiscount,omitempty"`
	InvoiceDiscount      *float32 `json:"invoiceDiscount,omitempty" xml:"invoiceDiscount,omitempty"`
	ModuleCode           *string  `json:"moduleCode,omitempty" xml:"moduleCode,omitempty"`
	OriginalCost         *float32 `json:"originalCost,omitempty" xml:"originalCost,omitempty"`
	QuantityUsedForQuote *float32 `json:"quantityUsedForQuote,omitempty" xml:"quantityUsedForQuote,omitempty"`
	UnitPrice            *float32 `json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
	UsageUnit            *string  `json:"usageUnit,omitempty" xml:"usageUnit,omitempty"`
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
	AssumedQuantity      *float32 `json:"assumedQuantity,omitempty" xml:"assumedQuantity,omitempty"`
	AssumedQuantityCost  *float32 `json:"assumedQuantityCost,omitempty" xml:"assumedQuantityCost,omitempty"`
	ModuleCode           *string  `json:"moduleCode,omitempty" xml:"moduleCode,omitempty"`
	OriginalPricePerUnit *float32 `json:"originalPricePerUnit,omitempty" xml:"originalPricePerUnit,omitempty"`
	PricePerUnit         *float32 `json:"pricePerUnit,omitempty" xml:"pricePerUnit,omitempty"`
	Unit                 *string  `json:"unit,omitempty" xml:"unit,omitempty"`
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
	AssumedQuantity      *float32 `json:"assumedQuantity,omitempty" xml:"assumedQuantity,omitempty"`
	AssumedQuantityCost  *float32 `json:"assumedQuantityCost,omitempty" xml:"assumedQuantityCost,omitempty"`
	ModuleCode           *string  `json:"moduleCode,omitempty" xml:"moduleCode,omitempty"`
	OriginalPricePerUnit *float32 `json:"originalPricePerUnit,omitempty" xml:"originalPricePerUnit,omitempty"`
	PricePerUnit         *float32 `json:"pricePerUnit,omitempty" xml:"pricePerUnit,omitempty"`
	Unit                 *string  `json:"unit,omitempty" xml:"unit,omitempty"`
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
