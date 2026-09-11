// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPriceComponentsValue interface {
	dara.Model
	String() string
	GoString() string
	SetCurrency(v string) *PriceComponentsValue
	GetCurrency() *string
	SetTradePrice(v float32) *PriceComponentsValue
	GetTradePrice() *float32
	SetOriginalPrice(v float32) *PriceComponentsValue
	GetOriginalPrice() *float32
	SetModuleSum(v float32) *PriceComponentsValue
	GetModuleSum() *float32
	SetEffectiveModuleSum(v float32) *PriceComponentsValue
	GetEffectiveModuleSum() *float32
	SetQuantity(v float32) *PriceComponentsValue
	GetQuantity() *float32
	SetPricingUnit(v string) *PriceComponentsValue
	GetPricingUnit() *string
	SetModules(v []*PriceComponentsValueModules) *PriceComponentsValue
	GetModules() []*PriceComponentsValueModules
	SetUsageCharges(v []*PriceComponentsValueUsageCharges) *PriceComponentsValue
	GetUsageCharges() []*PriceComponentsValueUsageCharges
	SetChargeComposition(v string) *PriceComponentsValue
	GetChargeComposition() *string
}

type PriceComponentsValue struct {
	Currency           *string                             `json:"currency,omitempty" xml:"currency,omitempty"`
	TradePrice         *float32                            `json:"tradePrice,omitempty" xml:"tradePrice,omitempty"`
	OriginalPrice      *float32                            `json:"originalPrice,omitempty" xml:"originalPrice,omitempty"`
	ModuleSum          *float32                            `json:"moduleSum,omitempty" xml:"moduleSum,omitempty"`
	EffectiveModuleSum *float32                            `json:"effectiveModuleSum,omitempty" xml:"effectiveModuleSum,omitempty"`
	Quantity           *float32                            `json:"quantity,omitempty" xml:"quantity,omitempty"`
	PricingUnit        *string                             `json:"pricingUnit,omitempty" xml:"pricingUnit,omitempty"`
	Modules            []*PriceComponentsValueModules      `json:"modules,omitempty" xml:"modules,omitempty" type:"Repeated"`
	UsageCharges       []*PriceComponentsValueUsageCharges `json:"usageCharges,omitempty" xml:"usageCharges,omitempty" type:"Repeated"`
	ChargeComposition  *string                             `json:"chargeComposition,omitempty" xml:"chargeComposition,omitempty"`
}

func (s PriceComponentsValue) String() string {
	return dara.Prettify(s)
}

func (s PriceComponentsValue) GoString() string {
	return s.String()
}

func (s *PriceComponentsValue) GetCurrency() *string {
	return s.Currency
}

func (s *PriceComponentsValue) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *PriceComponentsValue) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *PriceComponentsValue) GetModuleSum() *float32 {
	return s.ModuleSum
}

func (s *PriceComponentsValue) GetEffectiveModuleSum() *float32 {
	return s.EffectiveModuleSum
}

func (s *PriceComponentsValue) GetQuantity() *float32 {
	return s.Quantity
}

func (s *PriceComponentsValue) GetPricingUnit() *string {
	return s.PricingUnit
}

func (s *PriceComponentsValue) GetModules() []*PriceComponentsValueModules {
	return s.Modules
}

func (s *PriceComponentsValue) GetUsageCharges() []*PriceComponentsValueUsageCharges {
	return s.UsageCharges
}

func (s *PriceComponentsValue) GetChargeComposition() *string {
	return s.ChargeComposition
}

func (s *PriceComponentsValue) SetCurrency(v string) *PriceComponentsValue {
	s.Currency = &v
	return s
}

func (s *PriceComponentsValue) SetTradePrice(v float32) *PriceComponentsValue {
	s.TradePrice = &v
	return s
}

func (s *PriceComponentsValue) SetOriginalPrice(v float32) *PriceComponentsValue {
	s.OriginalPrice = &v
	return s
}

func (s *PriceComponentsValue) SetModuleSum(v float32) *PriceComponentsValue {
	s.ModuleSum = &v
	return s
}

func (s *PriceComponentsValue) SetEffectiveModuleSum(v float32) *PriceComponentsValue {
	s.EffectiveModuleSum = &v
	return s
}

func (s *PriceComponentsValue) SetQuantity(v float32) *PriceComponentsValue {
	s.Quantity = &v
	return s
}

func (s *PriceComponentsValue) SetPricingUnit(v string) *PriceComponentsValue {
	s.PricingUnit = &v
	return s
}

func (s *PriceComponentsValue) SetModules(v []*PriceComponentsValueModules) *PriceComponentsValue {
	s.Modules = v
	return s
}

func (s *PriceComponentsValue) SetUsageCharges(v []*PriceComponentsValueUsageCharges) *PriceComponentsValue {
	s.UsageCharges = v
	return s
}

func (s *PriceComponentsValue) SetChargeComposition(v string) *PriceComponentsValue {
	s.ChargeComposition = &v
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
	ModuleCode           *string  `json:"moduleCode,omitempty" xml:"moduleCode,omitempty"`
	CostAfterDiscount    *float32 `json:"costAfterDiscount,omitempty" xml:"costAfterDiscount,omitempty"`
	OriginalCost         *float32 `json:"originalCost,omitempty" xml:"originalCost,omitempty"`
	InvoiceDiscount      *float32 `json:"invoiceDiscount,omitempty" xml:"invoiceDiscount,omitempty"`
	UnitPrice            *float32 `json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
	BillingMode          *string  `json:"billingMode,omitempty" xml:"billingMode,omitempty"`
	UsageUnit            *string  `json:"usageUnit,omitempty" xml:"usageUnit,omitempty"`
	QuantityUsedForQuote *float32 `json:"quantityUsedForQuote,omitempty" xml:"quantityUsedForQuote,omitempty"`
}

func (s PriceComponentsValueModules) String() string {
	return dara.Prettify(s)
}

func (s PriceComponentsValueModules) GoString() string {
	return s.String()
}

func (s *PriceComponentsValueModules) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *PriceComponentsValueModules) GetCostAfterDiscount() *float32 {
	return s.CostAfterDiscount
}

func (s *PriceComponentsValueModules) GetOriginalCost() *float32 {
	return s.OriginalCost
}

func (s *PriceComponentsValueModules) GetInvoiceDiscount() *float32 {
	return s.InvoiceDiscount
}

func (s *PriceComponentsValueModules) GetUnitPrice() *float32 {
	return s.UnitPrice
}

func (s *PriceComponentsValueModules) GetBillingMode() *string {
	return s.BillingMode
}

func (s *PriceComponentsValueModules) GetUsageUnit() *string {
	return s.UsageUnit
}

func (s *PriceComponentsValueModules) GetQuantityUsedForQuote() *float32 {
	return s.QuantityUsedForQuote
}

func (s *PriceComponentsValueModules) SetModuleCode(v string) *PriceComponentsValueModules {
	s.ModuleCode = &v
	return s
}

func (s *PriceComponentsValueModules) SetCostAfterDiscount(v float32) *PriceComponentsValueModules {
	s.CostAfterDiscount = &v
	return s
}

func (s *PriceComponentsValueModules) SetOriginalCost(v float32) *PriceComponentsValueModules {
	s.OriginalCost = &v
	return s
}

func (s *PriceComponentsValueModules) SetInvoiceDiscount(v float32) *PriceComponentsValueModules {
	s.InvoiceDiscount = &v
	return s
}

func (s *PriceComponentsValueModules) SetUnitPrice(v float32) *PriceComponentsValueModules {
	s.UnitPrice = &v
	return s
}

func (s *PriceComponentsValueModules) SetBillingMode(v string) *PriceComponentsValueModules {
	s.BillingMode = &v
	return s
}

func (s *PriceComponentsValueModules) SetUsageUnit(v string) *PriceComponentsValueModules {
	s.UsageUnit = &v
	return s
}

func (s *PriceComponentsValueModules) SetQuantityUsedForQuote(v float32) *PriceComponentsValueModules {
	s.QuantityUsedForQuote = &v
	return s
}

func (s *PriceComponentsValueModules) Validate() error {
	return dara.Validate(s)
}

type PriceComponentsValueUsageCharges struct {
	ModuleCode           *string  `json:"moduleCode,omitempty" xml:"moduleCode,omitempty"`
	Unit                 *string  `json:"unit,omitempty" xml:"unit,omitempty"`
	PricePerUnit         *float32 `json:"pricePerUnit,omitempty" xml:"pricePerUnit,omitempty"`
	OriginalPricePerUnit *float32 `json:"originalPricePerUnit,omitempty" xml:"originalPricePerUnit,omitempty"`
	AssumedQuantity      *float32 `json:"assumedQuantity,omitempty" xml:"assumedQuantity,omitempty"`
	AssumedQuantityCost  *float32 `json:"assumedQuantityCost,omitempty" xml:"assumedQuantityCost,omitempty"`
}

func (s PriceComponentsValueUsageCharges) String() string {
	return dara.Prettify(s)
}

func (s PriceComponentsValueUsageCharges) GoString() string {
	return s.String()
}

func (s *PriceComponentsValueUsageCharges) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *PriceComponentsValueUsageCharges) GetUnit() *string {
	return s.Unit
}

func (s *PriceComponentsValueUsageCharges) GetPricePerUnit() *float32 {
	return s.PricePerUnit
}

func (s *PriceComponentsValueUsageCharges) GetOriginalPricePerUnit() *float32 {
	return s.OriginalPricePerUnit
}

func (s *PriceComponentsValueUsageCharges) GetAssumedQuantity() *float32 {
	return s.AssumedQuantity
}

func (s *PriceComponentsValueUsageCharges) GetAssumedQuantityCost() *float32 {
	return s.AssumedQuantityCost
}

func (s *PriceComponentsValueUsageCharges) SetModuleCode(v string) *PriceComponentsValueUsageCharges {
	s.ModuleCode = &v
	return s
}

func (s *PriceComponentsValueUsageCharges) SetUnit(v string) *PriceComponentsValueUsageCharges {
	s.Unit = &v
	return s
}

func (s *PriceComponentsValueUsageCharges) SetPricePerUnit(v float32) *PriceComponentsValueUsageCharges {
	s.PricePerUnit = &v
	return s
}

func (s *PriceComponentsValueUsageCharges) SetOriginalPricePerUnit(v float32) *PriceComponentsValueUsageCharges {
	s.OriginalPricePerUnit = &v
	return s
}

func (s *PriceComponentsValueUsageCharges) SetAssumedQuantity(v float32) *PriceComponentsValueUsageCharges {
	s.AssumedQuantity = &v
	return s
}

func (s *PriceComponentsValueUsageCharges) SetAssumedQuantityCost(v float32) *PriceComponentsValueUsageCharges {
	s.AssumedQuantityCost = &v
	return s
}

func (s *PriceComponentsValueUsageCharges) Validate() error {
	return dara.Validate(s)
}
