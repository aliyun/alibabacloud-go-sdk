// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRatePlanPriceGapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceModel(v *DescribeRatePlanPriceGapResponseBodyPriceModel) *DescribeRatePlanPriceGapResponseBody
	GetPriceModel() *DescribeRatePlanPriceGapResponseBodyPriceModel
	SetRequestId(v string) *DescribeRatePlanPriceGapResponseBody
	GetRequestId() *string
}

type DescribeRatePlanPriceGapResponseBody struct {
	// The price information.
	PriceModel *DescribeRatePlanPriceGapResponseBodyPriceModel `json:"PriceModel,omitempty" xml:"PriceModel,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 40423A7F-A83D-1E24-B80E-86DD25790759
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRatePlanPriceGapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRatePlanPriceGapResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRatePlanPriceGapResponseBody) GetPriceModel() *DescribeRatePlanPriceGapResponseBodyPriceModel {
	return s.PriceModel
}

func (s *DescribeRatePlanPriceGapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRatePlanPriceGapResponseBody) SetPriceModel(v *DescribeRatePlanPriceGapResponseBodyPriceModel) *DescribeRatePlanPriceGapResponseBody {
	s.PriceModel = v
	return s
}

func (s *DescribeRatePlanPriceGapResponseBody) SetRequestId(v string) *DescribeRatePlanPriceGapResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRatePlanPriceGapResponseBody) Validate() error {
	if s.PriceModel != nil {
		if err := s.PriceModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRatePlanPriceGapResponseBodyPriceModel struct {
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
	// The instance ID.
	//
	// example:
	//
	// xcdn-91fknmb80f0g
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The final order price, which is the actual transaction price.
	//
	// example:
	//
	// 1
	Price *float32                                            `json:"Price,omitempty" xml:"Price,omitempty"`
	Rule  *DescribeRatePlanPriceGapResponseBodyPriceModelRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
	// The target plan code.
	//
	// example:
	//
	// entranceplan
	TargetPlanCode *string `json:"TargetPlanCode,omitempty" xml:"TargetPlanCode,omitempty"`
	// The original order price, which equals the actual transaction price plus the discount amount.
	//
	// example:
	//
	// 2
	TotalPrice *float32 `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
}

func (s DescribeRatePlanPriceGapResponseBodyPriceModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeRatePlanPriceGapResponseBodyPriceModel) GoString() string {
	return s.String()
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModel) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModel) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModel) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModel) GetPrice() *float32 {
	return s.Price
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModel) GetRule() *DescribeRatePlanPriceGapResponseBodyPriceModelRule {
	return s.Rule
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModel) GetTargetPlanCode() *string {
	return s.TargetPlanCode
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModel) GetTotalPrice() *float32 {
	return s.TotalPrice
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModel) SetCurrency(v string) *DescribeRatePlanPriceGapResponseBodyPriceModel {
	s.Currency = &v
	return s
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModel) SetDiscountPrice(v float32) *DescribeRatePlanPriceGapResponseBodyPriceModel {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModel) SetInstanceId(v string) *DescribeRatePlanPriceGapResponseBodyPriceModel {
	s.InstanceId = &v
	return s
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModel) SetPrice(v float32) *DescribeRatePlanPriceGapResponseBodyPriceModel {
	s.Price = &v
	return s
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModel) SetRule(v *DescribeRatePlanPriceGapResponseBodyPriceModelRule) *DescribeRatePlanPriceGapResponseBodyPriceModel {
	s.Rule = v
	return s
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModel) SetTargetPlanCode(v string) *DescribeRatePlanPriceGapResponseBodyPriceModel {
	s.TargetPlanCode = &v
	return s
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModel) SetTotalPrice(v float32) *DescribeRatePlanPriceGapResponseBodyPriceModel {
	s.TotalPrice = &v
	return s
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModel) Validate() error {
	if s.Rule != nil {
		if err := s.Rule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRatePlanPriceGapResponseBodyPriceModelRule struct {
	RuleList []*DescribeRatePlanPriceGapResponseBodyPriceModelRuleRuleList `json:"RuleList,omitempty" xml:"RuleList,omitempty" type:"Repeated"`
}

func (s DescribeRatePlanPriceGapResponseBodyPriceModelRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeRatePlanPriceGapResponseBodyPriceModelRule) GoString() string {
	return s.String()
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModelRule) GetRuleList() []*DescribeRatePlanPriceGapResponseBodyPriceModelRuleRuleList {
	return s.RuleList
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModelRule) SetRuleList(v []*DescribeRatePlanPriceGapResponseBodyPriceModelRuleRuleList) *DescribeRatePlanPriceGapResponseBodyPriceModelRule {
	s.RuleList = v
	return s
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModelRule) Validate() error {
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

type DescribeRatePlanPriceGapResponseBodyPriceModelRuleRuleList struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleDescId *int64  `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s DescribeRatePlanPriceGapResponseBodyPriceModelRuleRuleList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRatePlanPriceGapResponseBodyPriceModelRuleRuleList) GoString() string {
	return s.String()
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModelRuleRuleList) GetName() *string {
	return s.Name
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModelRuleRuleList) GetRuleDescId() *int64 {
	return s.RuleDescId
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModelRuleRuleList) SetName(v string) *DescribeRatePlanPriceGapResponseBodyPriceModelRuleRuleList {
	s.Name = &v
	return s
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModelRuleRuleList) SetRuleDescId(v int64) *DescribeRatePlanPriceGapResponseBodyPriceModelRuleRuleList {
	s.RuleDescId = &v
	return s
}

func (s *DescribeRatePlanPriceGapResponseBodyPriceModelRuleRuleList) Validate() error {
	return dara.Validate(s)
}
