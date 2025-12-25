// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPriceInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCurrency(v string) *PriceInfo
	GetCurrency() *string
	SetDiscountAmount(v string) *PriceInfo
	GetDiscountAmount() *string
	SetOptionalPromotions(v []*PromotionInfo) *PriceInfo
	GetOptionalPromotions() []*PromotionInfo
	SetOriginalAmount(v string) *PriceInfo
	GetOriginalAmount() *string
	SetPriceModules(v []*PriceInfoPriceModules) *PriceInfo
	GetPriceModules() []*PriceInfoPriceModules
	SetRules(v []*PriceInfoRules) *PriceInfo
	GetRules() []*PriceInfoRules
	SetTaxAmount(v string) *PriceInfo
	GetTaxAmount() *string
	SetTradeAmount(v string) *PriceInfo
	GetTradeAmount() *string
}

type PriceInfo struct {
	Currency           *string                  `json:"currency,omitempty" xml:"currency,omitempty"`
	DiscountAmount     *string                  `json:"discountAmount,omitempty" xml:"discountAmount,omitempty"`
	OptionalPromotions []*PromotionInfo         `json:"optionalPromotions,omitempty" xml:"optionalPromotions,omitempty" type:"Repeated"`
	OriginalAmount     *string                  `json:"originalAmount,omitempty" xml:"originalAmount,omitempty"`
	PriceModules       []*PriceInfoPriceModules `json:"priceModules,omitempty" xml:"priceModules,omitempty" type:"Repeated"`
	Rules              []*PriceInfoRules        `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
	TaxAmount          *string                  `json:"taxAmount,omitempty" xml:"taxAmount,omitempty"`
	TradeAmount        *string                  `json:"tradeAmount,omitempty" xml:"tradeAmount,omitempty"`
}

func (s PriceInfo) String() string {
	return dara.Prettify(s)
}

func (s PriceInfo) GoString() string {
	return s.String()
}

func (s *PriceInfo) GetCurrency() *string {
	return s.Currency
}

func (s *PriceInfo) GetDiscountAmount() *string {
	return s.DiscountAmount
}

func (s *PriceInfo) GetOptionalPromotions() []*PromotionInfo {
	return s.OptionalPromotions
}

func (s *PriceInfo) GetOriginalAmount() *string {
	return s.OriginalAmount
}

func (s *PriceInfo) GetPriceModules() []*PriceInfoPriceModules {
	return s.PriceModules
}

func (s *PriceInfo) GetRules() []*PriceInfoRules {
	return s.Rules
}

func (s *PriceInfo) GetTaxAmount() *string {
	return s.TaxAmount
}

func (s *PriceInfo) GetTradeAmount() *string {
	return s.TradeAmount
}

func (s *PriceInfo) SetCurrency(v string) *PriceInfo {
	s.Currency = &v
	return s
}

func (s *PriceInfo) SetDiscountAmount(v string) *PriceInfo {
	s.DiscountAmount = &v
	return s
}

func (s *PriceInfo) SetOptionalPromotions(v []*PromotionInfo) *PriceInfo {
	s.OptionalPromotions = v
	return s
}

func (s *PriceInfo) SetOriginalAmount(v string) *PriceInfo {
	s.OriginalAmount = &v
	return s
}

func (s *PriceInfo) SetPriceModules(v []*PriceInfoPriceModules) *PriceInfo {
	s.PriceModules = v
	return s
}

func (s *PriceInfo) SetRules(v []*PriceInfoRules) *PriceInfo {
	s.Rules = v
	return s
}

func (s *PriceInfo) SetTaxAmount(v string) *PriceInfo {
	s.TaxAmount = &v
	return s
}

func (s *PriceInfo) SetTradeAmount(v string) *PriceInfo {
	s.TradeAmount = &v
	return s
}

func (s *PriceInfo) Validate() error {
	if s.OptionalPromotions != nil {
		for _, item := range s.OptionalPromotions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PriceModules != nil {
		for _, item := range s.PriceModules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PriceInfoPriceModules struct {
	OriginalAmount *string `json:"originalAmount,omitempty" xml:"originalAmount,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PriceInfoPriceModules) String() string {
	return dara.Prettify(s)
}

func (s PriceInfoPriceModules) GoString() string {
	return s.String()
}

func (s *PriceInfoPriceModules) GetOriginalAmount() *string {
	return s.OriginalAmount
}

func (s *PriceInfoPriceModules) GetType() *string {
	return s.Type
}

func (s *PriceInfoPriceModules) SetOriginalAmount(v string) *PriceInfoPriceModules {
	s.OriginalAmount = &v
	return s
}

func (s *PriceInfoPriceModules) SetType(v string) *PriceInfoPriceModules {
	s.Type = &v
	return s
}

func (s *PriceInfoPriceModules) Validate() error {
	return dara.Validate(s)
}

type PriceInfoRules struct {
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
}

func (s PriceInfoRules) String() string {
	return dara.Prettify(s)
}

func (s PriceInfoRules) GoString() string {
	return s.String()
}

func (s *PriceInfoRules) GetAmount() *string {
	return s.Amount
}

func (s *PriceInfoRules) GetName() *string {
	return s.Name
}

func (s *PriceInfoRules) GetRuleId() *string {
	return s.RuleId
}

func (s *PriceInfoRules) SetAmount(v string) *PriceInfoRules {
	s.Amount = &v
	return s
}

func (s *PriceInfoRules) SetName(v string) *PriceInfoRules {
	s.Name = &v
	return s
}

func (s *PriceInfoRules) SetRuleId(v string) *PriceInfoRules {
	s.RuleId = &v
	return s
}

func (s *PriceInfoRules) Validate() error {
	return dara.Validate(s)
}
