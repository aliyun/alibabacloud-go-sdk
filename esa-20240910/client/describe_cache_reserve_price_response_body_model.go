// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCacheReservePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceModel(v *DescribeCacheReservePriceResponseBodyPriceModel) *DescribeCacheReservePriceResponseBody
	GetPriceModel() *DescribeCacheReservePriceResponseBodyPriceModel
	SetRequestId(v string) *DescribeCacheReservePriceResponseBody
	GetRequestId() *string
}

type DescribeCacheReservePriceResponseBody struct {
	// The price information.
	PriceModel *DescribeCacheReservePriceResponseBodyPriceModel `json:"PriceModel,omitempty" xml:"PriceModel,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2430E05E-1340-5773-B5E1-B743929F46F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCacheReservePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheReservePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCacheReservePriceResponseBody) GetPriceModel() *DescribeCacheReservePriceResponseBodyPriceModel {
	return s.PriceModel
}

func (s *DescribeCacheReservePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCacheReservePriceResponseBody) SetPriceModel(v *DescribeCacheReservePriceResponseBodyPriceModel) *DescribeCacheReservePriceResponseBody {
	s.PriceModel = v
	return s
}

func (s *DescribeCacheReservePriceResponseBody) SetRequestId(v string) *DescribeCacheReservePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCacheReservePriceResponseBody) Validate() error {
	if s.PriceModel != nil {
		if err := s.PriceModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCacheReservePriceResponseBodyPriceModel struct {
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
	// 1
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The final order price, which is the actual transaction price.
	//
	// example:
	//
	// 1
	Price *float32                                             `json:"Price,omitempty" xml:"Price,omitempty"`
	Rule  *DescribeCacheReservePriceResponseBodyPriceModelRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
	// The original order price. Original order price = actual transaction price + discount amount.
	//
	// example:
	//
	// 2
	TotalPrice *float32 `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
}

func (s DescribeCacheReservePriceResponseBodyPriceModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheReservePriceResponseBodyPriceModel) GoString() string {
	return s.String()
}

func (s *DescribeCacheReservePriceResponseBodyPriceModel) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeCacheReservePriceResponseBodyPriceModel) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeCacheReservePriceResponseBodyPriceModel) GetPrice() *float32 {
	return s.Price
}

func (s *DescribeCacheReservePriceResponseBodyPriceModel) GetRule() *DescribeCacheReservePriceResponseBodyPriceModelRule {
	return s.Rule
}

func (s *DescribeCacheReservePriceResponseBodyPriceModel) GetTotalPrice() *float32 {
	return s.TotalPrice
}

func (s *DescribeCacheReservePriceResponseBodyPriceModel) SetCurrency(v string) *DescribeCacheReservePriceResponseBodyPriceModel {
	s.Currency = &v
	return s
}

func (s *DescribeCacheReservePriceResponseBodyPriceModel) SetDiscountPrice(v float32) *DescribeCacheReservePriceResponseBodyPriceModel {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeCacheReservePriceResponseBodyPriceModel) SetPrice(v float32) *DescribeCacheReservePriceResponseBodyPriceModel {
	s.Price = &v
	return s
}

func (s *DescribeCacheReservePriceResponseBodyPriceModel) SetRule(v *DescribeCacheReservePriceResponseBodyPriceModelRule) *DescribeCacheReservePriceResponseBodyPriceModel {
	s.Rule = v
	return s
}

func (s *DescribeCacheReservePriceResponseBodyPriceModel) SetTotalPrice(v float32) *DescribeCacheReservePriceResponseBodyPriceModel {
	s.TotalPrice = &v
	return s
}

func (s *DescribeCacheReservePriceResponseBodyPriceModel) Validate() error {
	if s.Rule != nil {
		if err := s.Rule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCacheReservePriceResponseBodyPriceModelRule struct {
	RuleList []*DescribeCacheReservePriceResponseBodyPriceModelRuleRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s DescribeCacheReservePriceResponseBodyPriceModelRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheReservePriceResponseBodyPriceModelRule) GoString() string {
	return s.String()
}

func (s *DescribeCacheReservePriceResponseBodyPriceModelRule) GetRuleList() []*DescribeCacheReservePriceResponseBodyPriceModelRuleRuleList {
	return s.RuleList
}

func (s *DescribeCacheReservePriceResponseBodyPriceModelRule) SetRuleList(v []*DescribeCacheReservePriceResponseBodyPriceModelRuleRuleList) *DescribeCacheReservePriceResponseBodyPriceModelRule {
	s.RuleList = v
	return s
}

func (s *DescribeCacheReservePriceResponseBodyPriceModelRule) Validate() error {
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

type DescribeCacheReservePriceResponseBodyPriceModelRuleRuleList struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleDescId *int64  `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s DescribeCacheReservePriceResponseBodyPriceModelRuleRuleList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheReservePriceResponseBodyPriceModelRuleRuleList) GoString() string {
	return s.String()
}

func (s *DescribeCacheReservePriceResponseBodyPriceModelRuleRuleList) GetName() *string {
	return s.Name
}

func (s *DescribeCacheReservePriceResponseBodyPriceModelRuleRuleList) GetRuleDescId() *int64 {
	return s.RuleDescId
}

func (s *DescribeCacheReservePriceResponseBodyPriceModelRuleRuleList) SetName(v string) *DescribeCacheReservePriceResponseBodyPriceModelRuleRuleList {
	s.Name = &v
	return s
}

func (s *DescribeCacheReservePriceResponseBodyPriceModelRuleRuleList) SetRuleDescId(v int64) *DescribeCacheReservePriceResponseBodyPriceModelRuleRuleList {
	s.RuleDescId = &v
	return s
}

func (s *DescribeCacheReservePriceResponseBodyPriceModelRuleRuleList) Validate() error {
	return dara.Validate(s)
}
