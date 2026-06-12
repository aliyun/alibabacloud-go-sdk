// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceInstanceSubscriptionEstimateCostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetServiceInstanceSubscriptionEstimateCostResponseBody
	GetRequestId() *string
	SetResourcePrices(v []*GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) *GetServiceInstanceSubscriptionEstimateCostResponseBody
	GetResourcePrices() []*GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices
}

type GetServiceInstanceSubscriptionEstimateCostResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 464C8CB6-A548-5206-B83C-D32A8E43EC21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of resource price information.
	ResourcePrices []*GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices `json:"ResourcePrices,omitempty" xml:"ResourcePrices,omitempty" type:"Repeated"`
}

func (s GetServiceInstanceSubscriptionEstimateCostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceInstanceSubscriptionEstimateCostResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBody) GetResourcePrices() []*GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices {
	return s.ResourcePrices
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBody) SetRequestId(v string) *GetServiceInstanceSubscriptionEstimateCostResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBody) SetResourcePrices(v []*GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) *GetServiceInstanceSubscriptionEstimateCostResponseBody {
	s.ResourcePrices = v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBody) Validate() error {
	if s.ResourcePrices != nil {
		for _, item := range s.ResourcePrices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices struct {
	// The currency. Valid values:
	//
	// - CNY: Chinese Yuan.
	//
	// - USD: US Dollar.
	//
	// - JPY: Japanese Yen.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The price details of the pricing module.
	DetailInfos []*GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos `json:"DetailInfos,omitempty" xml:"DetailInfos,omitempty" type:"Repeated"`
	// The discount.
	//
	// example:
	//
	// 100
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The original price.
	//
	// example:
	//
	// 900
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// The renewal duration. The unit is specified by PeriodUnit.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal duration. This is the unit for the Period parameter. Valid values: Month and Year. Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ARN of the resource.
	//
	// example:
	//
	// acs:ecs:cn-hongkong:1488317743351199:instance/i-j6c6f3lbky38o8rpeqw2
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// The discount details.
	Rules []*GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The discounted price.
	//
	// example:
	//
	// 500
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) String() string {
	return dara.Prettify(s)
}

func (s GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) GetCurrency() *string {
	return s.Currency
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) GetDetailInfos() []*GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos {
	return s.DetailInfos
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) GetPeriod() *int32 {
	return s.Period
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) GetRules() []*GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules {
	return s.Rules
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) SetCurrency(v string) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices {
	s.Currency = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) SetDetailInfos(v []*GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices {
	s.DetailInfos = v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) SetDiscountAmount(v float32) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices {
	s.DiscountAmount = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) SetOriginalAmount(v float32) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices {
	s.OriginalAmount = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) SetPeriod(v int32) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices {
	s.Period = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) SetPeriodUnit(v string) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices {
	s.PeriodUnit = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) SetResourceArn(v string) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices {
	s.ResourceArn = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) SetRules(v []*GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices {
	s.Rules = v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) SetTradeAmount(v float32) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices {
	s.TradeAmount = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) Validate() error {
	if s.DetailInfos != nil {
		for _, item := range s.DetailInfos {
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

type GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos struct {
	// The discount.
	//
	// example:
	//
	// 100
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The original price.
	//
	// example:
	//
	// 900
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// The identifier of the pricing module.
	//
	// example:
	//
	// instance
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The discounted price.
	//
	// example:
	//
	// 500
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos) String() string {
	return dara.Prettify(s)
}

func (s GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos) GetResource() *string {
	return s.Resource
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos) SetDiscountAmount(v float32) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos {
	s.DiscountAmount = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos) SetOriginalAmount(v float32) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos {
	s.OriginalAmount = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos) SetResource(v string) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos {
	s.Resource = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos) SetTradeAmount(v float32) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos {
	s.TradeAmount = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos) Validate() error {
	return dara.Validate(s)
}

type GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules struct {
	// The description of the discount.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the discount.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The discount ID.
	//
	// example:
	//
	// 1021199213
	RuleDescId *int64 `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules) String() string {
	return dara.Prettify(s)
}

func (s GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules) GetDescription() *string {
	return s.Description
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules) GetName() *string {
	return s.Name
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules) GetRuleDescId() *int64 {
	return s.RuleDescId
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules) SetDescription(v string) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules {
	s.Description = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules) SetName(v string) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules {
	s.Name = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules) SetRuleDescId(v int64) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules {
	s.RuleDescId = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules) Validate() error {
	return dara.Validate(s)
}
