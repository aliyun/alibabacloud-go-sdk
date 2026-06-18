// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCacheReservePriceGapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceModel(v *DescribeCacheReservePriceGapResponseBodyPriceModel) *DescribeCacheReservePriceGapResponseBody
	GetPriceModel() *DescribeCacheReservePriceGapResponseBodyPriceModel
	SetRequestId(v string) *DescribeCacheReservePriceGapResponseBody
	GetRequestId() *string
}

type DescribeCacheReservePriceGapResponseBody struct {
	PriceModel *DescribeCacheReservePriceGapResponseBodyPriceModel `json:"PriceModel,omitempty" xml:"PriceModel,omitempty" type:"Struct"`
	// example:
	//
	// 40423A7F-A83D-1E24-B80E-86DD25790759
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCacheReservePriceGapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheReservePriceGapResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCacheReservePriceGapResponseBody) GetPriceModel() *DescribeCacheReservePriceGapResponseBodyPriceModel {
	return s.PriceModel
}

func (s *DescribeCacheReservePriceGapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCacheReservePriceGapResponseBody) SetPriceModel(v *DescribeCacheReservePriceGapResponseBodyPriceModel) *DescribeCacheReservePriceGapResponseBody {
	s.PriceModel = v
	return s
}

func (s *DescribeCacheReservePriceGapResponseBody) SetRequestId(v string) *DescribeCacheReservePriceGapResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCacheReservePriceGapResponseBody) Validate() error {
	if s.PriceModel != nil {
		if err := s.PriceModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCacheReservePriceGapResponseBodyPriceModel struct {
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 1
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// esa-cr-9tuv*********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	Price *float32                                                `json:"Price,omitempty" xml:"Price,omitempty"`
	Rule  *DescribeCacheReservePriceGapResponseBodyPriceModelRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
	// example:
	//
	// 2
	TotalPrice *float32 `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
}

func (s DescribeCacheReservePriceGapResponseBodyPriceModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheReservePriceGapResponseBodyPriceModel) GoString() string {
	return s.String()
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModel) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModel) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModel) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModel) GetPrice() *float32 {
	return s.Price
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModel) GetRule() *DescribeCacheReservePriceGapResponseBodyPriceModelRule {
	return s.Rule
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModel) GetTotalPrice() *float32 {
	return s.TotalPrice
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModel) SetCurrency(v string) *DescribeCacheReservePriceGapResponseBodyPriceModel {
	s.Currency = &v
	return s
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModel) SetDiscountPrice(v float32) *DescribeCacheReservePriceGapResponseBodyPriceModel {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModel) SetInstanceId(v string) *DescribeCacheReservePriceGapResponseBodyPriceModel {
	s.InstanceId = &v
	return s
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModel) SetPrice(v float32) *DescribeCacheReservePriceGapResponseBodyPriceModel {
	s.Price = &v
	return s
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModel) SetRule(v *DescribeCacheReservePriceGapResponseBodyPriceModelRule) *DescribeCacheReservePriceGapResponseBodyPriceModel {
	s.Rule = v
	return s
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModel) SetTotalPrice(v float32) *DescribeCacheReservePriceGapResponseBodyPriceModel {
	s.TotalPrice = &v
	return s
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModel) Validate() error {
	if s.Rule != nil {
		if err := s.Rule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCacheReservePriceGapResponseBodyPriceModelRule struct {
	RuleList []*DescribeCacheReservePriceGapResponseBodyPriceModelRuleRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s DescribeCacheReservePriceGapResponseBodyPriceModelRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheReservePriceGapResponseBodyPriceModelRule) GoString() string {
	return s.String()
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModelRule) GetRuleList() []*DescribeCacheReservePriceGapResponseBodyPriceModelRuleRuleList {
	return s.RuleList
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModelRule) SetRuleList(v []*DescribeCacheReservePriceGapResponseBodyPriceModelRuleRuleList) *DescribeCacheReservePriceGapResponseBodyPriceModelRule {
	s.RuleList = v
	return s
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModelRule) Validate() error {
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

type DescribeCacheReservePriceGapResponseBodyPriceModelRuleRuleList struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleDescId *int64  `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s DescribeCacheReservePriceGapResponseBodyPriceModelRuleRuleList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheReservePriceGapResponseBodyPriceModelRuleRuleList) GoString() string {
	return s.String()
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModelRuleRuleList) GetName() *string {
	return s.Name
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModelRuleRuleList) GetRuleDescId() *int64 {
	return s.RuleDescId
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModelRuleRuleList) SetName(v string) *DescribeCacheReservePriceGapResponseBodyPriceModelRuleRuleList {
	s.Name = &v
	return s
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModelRuleRuleList) SetRuleDescId(v int64) *DescribeCacheReservePriceGapResponseBodyPriceModelRuleRuleList {
	s.RuleDescId = &v
	return s
}

func (s *DescribeCacheReservePriceGapResponseBodyPriceModelRuleRuleList) Validate() error {
	return dara.Validate(s)
}
