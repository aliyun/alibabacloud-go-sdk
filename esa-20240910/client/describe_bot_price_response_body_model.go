// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBotPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceModel(v *DescribeBotPriceResponseBodyPriceModel) *DescribeBotPriceResponseBody
	GetPriceModel() *DescribeBotPriceResponseBodyPriceModel
	SetRequestId(v string) *DescribeBotPriceResponseBody
	GetRequestId() *string
}

type DescribeBotPriceResponseBody struct {
	// The price information.
	PriceModel *DescribeBotPriceResponseBodyPriceModel `json:"PriceModel,omitempty" xml:"PriceModel,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// FCF50EDF-1C2B-51E9-A372-E194D16ED350
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBotPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBotPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBotPriceResponseBody) GetPriceModel() *DescribeBotPriceResponseBodyPriceModel {
	return s.PriceModel
}

func (s *DescribeBotPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBotPriceResponseBody) SetPriceModel(v *DescribeBotPriceResponseBodyPriceModel) *DescribeBotPriceResponseBody {
	s.PriceModel = v
	return s
}

func (s *DescribeBotPriceResponseBody) SetRequestId(v string) *DescribeBotPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBotPriceResponseBody) Validate() error {
	if s.PriceModel != nil {
		if err := s.PriceModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBotPriceResponseBodyPriceModel struct {
	// The currency. Valid values:
	//
	// - JPY: Japanese Yen.
	//
	// - USD: US Dollar.
	//
	// - CNY: Chinese Yuan.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The discount amount of the order.
	//
	// example:
	//
	// 50
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The final price of the order, which is the actual transaction price.
	//
	// example:
	//
	// 100
	Price *float32                                    `json:"Price,omitempty" xml:"Price,omitempty"`
	Rule  *DescribeBotPriceResponseBodyPriceModelRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
	// The original price of the order. Original price = actual transaction price + discount amount.
	//
	// example:
	//
	// 150
	TotalPrice *float32 `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
}

func (s DescribeBotPriceResponseBodyPriceModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeBotPriceResponseBodyPriceModel) GoString() string {
	return s.String()
}

func (s *DescribeBotPriceResponseBodyPriceModel) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeBotPriceResponseBodyPriceModel) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeBotPriceResponseBodyPriceModel) GetPrice() *float32 {
	return s.Price
}

func (s *DescribeBotPriceResponseBodyPriceModel) GetRule() *DescribeBotPriceResponseBodyPriceModelRule {
	return s.Rule
}

func (s *DescribeBotPriceResponseBodyPriceModel) GetTotalPrice() *float32 {
	return s.TotalPrice
}

func (s *DescribeBotPriceResponseBodyPriceModel) SetCurrency(v string) *DescribeBotPriceResponseBodyPriceModel {
	s.Currency = &v
	return s
}

func (s *DescribeBotPriceResponseBodyPriceModel) SetDiscountPrice(v float32) *DescribeBotPriceResponseBodyPriceModel {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeBotPriceResponseBodyPriceModel) SetPrice(v float32) *DescribeBotPriceResponseBodyPriceModel {
	s.Price = &v
	return s
}

func (s *DescribeBotPriceResponseBodyPriceModel) SetRule(v *DescribeBotPriceResponseBodyPriceModelRule) *DescribeBotPriceResponseBodyPriceModel {
	s.Rule = v
	return s
}

func (s *DescribeBotPriceResponseBodyPriceModel) SetTotalPrice(v float32) *DescribeBotPriceResponseBodyPriceModel {
	s.TotalPrice = &v
	return s
}

func (s *DescribeBotPriceResponseBodyPriceModel) Validate() error {
	if s.Rule != nil {
		if err := s.Rule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBotPriceResponseBodyPriceModelRule struct {
	RuleList []*DescribeBotPriceResponseBodyPriceModelRuleRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s DescribeBotPriceResponseBodyPriceModelRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeBotPriceResponseBodyPriceModelRule) GoString() string {
	return s.String()
}

func (s *DescribeBotPriceResponseBodyPriceModelRule) GetRuleList() []*DescribeBotPriceResponseBodyPriceModelRuleRuleList {
	return s.RuleList
}

func (s *DescribeBotPriceResponseBodyPriceModelRule) SetRuleList(v []*DescribeBotPriceResponseBodyPriceModelRuleRuleList) *DescribeBotPriceResponseBodyPriceModelRule {
	s.RuleList = v
	return s
}

func (s *DescribeBotPriceResponseBodyPriceModelRule) Validate() error {
	if s.RuleList != nil {
		for _, item := range s.RuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBotPriceResponseBodyPriceModelRuleRuleList struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleDescId *int64  `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s DescribeBotPriceResponseBodyPriceModelRuleRuleList) String() string {
	return dara.Prettify(s)
}

func (s DescribeBotPriceResponseBodyPriceModelRuleRuleList) GoString() string {
	return s.String()
}

func (s *DescribeBotPriceResponseBodyPriceModelRuleRuleList) GetName() *string {
	return s.Name
}

func (s *DescribeBotPriceResponseBodyPriceModelRuleRuleList) GetRuleDescId() *int64 {
	return s.RuleDescId
}

func (s *DescribeBotPriceResponseBodyPriceModelRuleRuleList) SetName(v string) *DescribeBotPriceResponseBodyPriceModelRuleRuleList {
	s.Name = &v
	return s
}

func (s *DescribeBotPriceResponseBodyPriceModelRuleRuleList) SetRuleDescId(v int64) *DescribeBotPriceResponseBodyPriceModelRuleRuleList {
	s.RuleDescId = &v
	return s
}

func (s *DescribeBotPriceResponseBodyPriceModelRuleRuleList) Validate() error {
	return dara.Validate(s)
}
