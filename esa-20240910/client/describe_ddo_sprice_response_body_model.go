// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceModel(v *DescribeDDoSPriceResponseBodyPriceModel) *DescribeDDoSPriceResponseBody
	GetPriceModel() *DescribeDDoSPriceResponseBodyPriceModel
	SetRequestId(v string) *DescribeDDoSPriceResponseBody
	GetRequestId() *string
}

type DescribeDDoSPriceResponseBody struct {
	// The price information.
	PriceModel *DescribeDDoSPriceResponseBodyPriceModel `json:"PriceModel,omitempty" xml:"PriceModel,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDDoSPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDoSPriceResponseBody) GetPriceModel() *DescribeDDoSPriceResponseBodyPriceModel {
	return s.PriceModel
}

func (s *DescribeDDoSPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDDoSPriceResponseBody) SetPriceModel(v *DescribeDDoSPriceResponseBodyPriceModel) *DescribeDDoSPriceResponseBody {
	s.PriceModel = v
	return s
}

func (s *DescribeDDoSPriceResponseBody) SetRequestId(v string) *DescribeDDoSPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDoSPriceResponseBody) Validate() error {
	if s.PriceModel != nil {
		if err := s.PriceModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDDoSPriceResponseBodyPriceModel struct {
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
	// 40
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The final price of the order, which is the actual transaction price.
	//
	// example:
	//
	// 60
	Price *float32                                     `json:"Price,omitempty" xml:"Price,omitempty"`
	Rule  *DescribeDDoSPriceResponseBodyPriceModelRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
	// The original price of the order. Original price = actual transaction price + discount amount.
	//
	// example:
	//
	// 100
	TotalPrice *float32 `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
}

func (s DescribeDDoSPriceResponseBodyPriceModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSPriceResponseBodyPriceModel) GoString() string {
	return s.String()
}

func (s *DescribeDDoSPriceResponseBodyPriceModel) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeDDoSPriceResponseBodyPriceModel) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeDDoSPriceResponseBodyPriceModel) GetPrice() *float32 {
	return s.Price
}

func (s *DescribeDDoSPriceResponseBodyPriceModel) GetRule() *DescribeDDoSPriceResponseBodyPriceModelRule {
	return s.Rule
}

func (s *DescribeDDoSPriceResponseBodyPriceModel) GetTotalPrice() *float32 {
	return s.TotalPrice
}

func (s *DescribeDDoSPriceResponseBodyPriceModel) SetCurrency(v string) *DescribeDDoSPriceResponseBodyPriceModel {
	s.Currency = &v
	return s
}

func (s *DescribeDDoSPriceResponseBodyPriceModel) SetDiscountPrice(v float32) *DescribeDDoSPriceResponseBodyPriceModel {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeDDoSPriceResponseBodyPriceModel) SetPrice(v float32) *DescribeDDoSPriceResponseBodyPriceModel {
	s.Price = &v
	return s
}

func (s *DescribeDDoSPriceResponseBodyPriceModel) SetRule(v *DescribeDDoSPriceResponseBodyPriceModelRule) *DescribeDDoSPriceResponseBodyPriceModel {
	s.Rule = v
	return s
}

func (s *DescribeDDoSPriceResponseBodyPriceModel) SetTotalPrice(v float32) *DescribeDDoSPriceResponseBodyPriceModel {
	s.TotalPrice = &v
	return s
}

func (s *DescribeDDoSPriceResponseBodyPriceModel) Validate() error {
	if s.Rule != nil {
		if err := s.Rule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDDoSPriceResponseBodyPriceModelRule struct {
	RuleList []*DescribeDDoSPriceResponseBodyPriceModelRuleRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s DescribeDDoSPriceResponseBodyPriceModelRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSPriceResponseBodyPriceModelRule) GoString() string {
	return s.String()
}

func (s *DescribeDDoSPriceResponseBodyPriceModelRule) GetRuleList() []*DescribeDDoSPriceResponseBodyPriceModelRuleRuleList {
	return s.RuleList
}

func (s *DescribeDDoSPriceResponseBodyPriceModelRule) SetRuleList(v []*DescribeDDoSPriceResponseBodyPriceModelRuleRuleList) *DescribeDDoSPriceResponseBodyPriceModelRule {
	s.RuleList = v
	return s
}

func (s *DescribeDDoSPriceResponseBodyPriceModelRule) Validate() error {
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

type DescribeDDoSPriceResponseBodyPriceModelRuleRuleList struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleDescId *int64  `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s DescribeDDoSPriceResponseBodyPriceModelRuleRuleList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSPriceResponseBodyPriceModelRuleRuleList) GoString() string {
	return s.String()
}

func (s *DescribeDDoSPriceResponseBodyPriceModelRuleRuleList) GetName() *string {
	return s.Name
}

func (s *DescribeDDoSPriceResponseBodyPriceModelRuleRuleList) GetRuleDescId() *int64 {
	return s.RuleDescId
}

func (s *DescribeDDoSPriceResponseBodyPriceModelRuleRuleList) SetName(v string) *DescribeDDoSPriceResponseBodyPriceModelRuleRuleList {
	s.Name = &v
	return s
}

func (s *DescribeDDoSPriceResponseBodyPriceModelRuleRuleList) SetRuleDescId(v int64) *DescribeDDoSPriceResponseBodyPriceModelRuleRuleList {
	s.RuleDescId = &v
	return s
}

func (s *DescribeDDoSPriceResponseBodyPriceModelRuleRuleList) Validate() error {
	return dara.Validate(s)
}
