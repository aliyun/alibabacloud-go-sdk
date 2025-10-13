// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigurationPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeConfigurationPriceResponseBody
	GetCode() *string
	SetData(v *DescribeConfigurationPriceResponseBodyData) *DescribeConfigurationPriceResponseBody
	GetData() *DescribeConfigurationPriceResponseBodyData
	SetErrorCode(v string) *DescribeConfigurationPriceResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeConfigurationPriceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeConfigurationPriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeConfigurationPriceResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeConfigurationPriceResponseBody
	GetTraceId() *string
}

type DescribeConfigurationPriceResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The request was successful.
	//
	// 	- **3xx**: The request was redirected.
	//
	// 	- **4xx**: The request failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The price.
	Data *DescribeConfigurationPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// 	- If the request was successful, **ErrorCode*	- is not returned.
	//
	// 	- If the request failed, **ErrorCode*	- is returned. For more information, see **Error codes*	- in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned. Valid values:
	//
	// 	- If the request was successful, **success*	- is returned.
	//
	// 	- If the request failed, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ADCEC067-86AD-19E2-BD43-E83F3841****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the configuration price was obtained.
	//
	// 	- **true**: The price was obtained.
	//
	// 	- **false**: The price failed to be queried.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace.
	//
	// example:
	//
	// 1a0dcc771722848598056771******
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeConfigurationPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigurationPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeConfigurationPriceResponseBody) GetData() *DescribeConfigurationPriceResponseBodyData {
	return s.Data
}

func (s *DescribeConfigurationPriceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeConfigurationPriceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeConfigurationPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConfigurationPriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeConfigurationPriceResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeConfigurationPriceResponseBody) SetCode(v string) *DescribeConfigurationPriceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBody) SetData(v *DescribeConfigurationPriceResponseBodyData) *DescribeConfigurationPriceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeConfigurationPriceResponseBody) SetErrorCode(v string) *DescribeConfigurationPriceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBody) SetMessage(v string) *DescribeConfigurationPriceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBody) SetRequestId(v string) *DescribeConfigurationPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBody) SetSuccess(v bool) *DescribeConfigurationPriceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBody) SetTraceId(v string) *DescribeConfigurationPriceResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeConfigurationPriceResponseBodyData struct {
	// The remaining capacity of the resource plan.
	BagUsage *DescribeConfigurationPriceResponseBodyDataBagUsage `json:"BagUsage,omitempty" xml:"BagUsage,omitempty" type:"Struct"`
	// The price of CPU and memory.
	CpuMemPrice *DescribeConfigurationPriceResponseBodyDataCpuMemPrice `json:"CpuMemPrice,omitempty" xml:"CpuMemPrice,omitempty" type:"Struct"`
	// The information about pricing.
	Order *DescribeConfigurationPriceResponseBodyDataOrder `json:"Order,omitempty" xml:"Order,omitempty" type:"Struct"`
	// The price based on the number of requests.
	RequestPrice *DescribeConfigurationPriceResponseBodyDataRequestPrice `json:"RequestPrice,omitempty" xml:"RequestPrice,omitempty" type:"Struct"`
	// The promotion rules.
	Rules []*DescribeConfigurationPriceResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The traffic price.
	TrafficPrice *DescribeConfigurationPriceResponseBodyDataTrafficPrice `json:"TrafficPrice,omitempty" xml:"TrafficPrice,omitempty" type:"Struct"`
}

func (s DescribeConfigurationPriceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigurationPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponseBodyData) GetBagUsage() *DescribeConfigurationPriceResponseBodyDataBagUsage {
	return s.BagUsage
}

func (s *DescribeConfigurationPriceResponseBodyData) GetCpuMemPrice() *DescribeConfigurationPriceResponseBodyDataCpuMemPrice {
	return s.CpuMemPrice
}

func (s *DescribeConfigurationPriceResponseBodyData) GetOrder() *DescribeConfigurationPriceResponseBodyDataOrder {
	return s.Order
}

func (s *DescribeConfigurationPriceResponseBodyData) GetRequestPrice() *DescribeConfigurationPriceResponseBodyDataRequestPrice {
	return s.RequestPrice
}

func (s *DescribeConfigurationPriceResponseBodyData) GetRules() []*DescribeConfigurationPriceResponseBodyDataRules {
	return s.Rules
}

func (s *DescribeConfigurationPriceResponseBodyData) GetTrafficPrice() *DescribeConfigurationPriceResponseBodyDataTrafficPrice {
	return s.TrafficPrice
}

func (s *DescribeConfigurationPriceResponseBodyData) SetBagUsage(v *DescribeConfigurationPriceResponseBodyDataBagUsage) *DescribeConfigurationPriceResponseBodyData {
	s.BagUsage = v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyData) SetCpuMemPrice(v *DescribeConfigurationPriceResponseBodyDataCpuMemPrice) *DescribeConfigurationPriceResponseBodyData {
	s.CpuMemPrice = v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyData) SetOrder(v *DescribeConfigurationPriceResponseBodyDataOrder) *DescribeConfigurationPriceResponseBodyData {
	s.Order = v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyData) SetRequestPrice(v *DescribeConfigurationPriceResponseBodyDataRequestPrice) *DescribeConfigurationPriceResponseBodyData {
	s.RequestPrice = v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyData) SetRules(v []*DescribeConfigurationPriceResponseBodyDataRules) *DescribeConfigurationPriceResponseBodyData {
	s.Rules = v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyData) SetTrafficPrice(v *DescribeConfigurationPriceResponseBodyDataTrafficPrice) *DescribeConfigurationPriceResponseBodyData {
	s.TrafficPrice = v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyData) Validate() error {
	if s.BagUsage != nil {
		if err := s.BagUsage.Validate(); err != nil {
			return err
		}
	}
	if s.CpuMemPrice != nil {
		if err := s.CpuMemPrice.Validate(); err != nil {
			return err
		}
	}
	if s.Order != nil {
		if err := s.Order.Validate(); err != nil {
			return err
		}
	}
	if s.RequestPrice != nil {
		if err := s.RequestPrice.Validate(); err != nil {
			return err
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
	if s.TrafficPrice != nil {
		if err := s.TrafficPrice.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeConfigurationPriceResponseBodyDataBagUsage struct {
	// The available CPU capacity. Unit: cores \\*.
	//
	// example:
	//
	// 497570.450009
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Cu  *float32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The available memory size. Unit: GiB ×.
	//
	// example:
	//
	// 989802.563546
	Mem *float32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
}

func (s DescribeConfigurationPriceResponseBodyDataBagUsage) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigurationPriceResponseBodyDataBagUsage) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponseBodyDataBagUsage) GetCpu() *float32 {
	return s.Cpu
}

func (s *DescribeConfigurationPriceResponseBodyDataBagUsage) GetCu() *float32 {
	return s.Cu
}

func (s *DescribeConfigurationPriceResponseBodyDataBagUsage) GetMem() *float32 {
	return s.Mem
}

func (s *DescribeConfigurationPriceResponseBodyDataBagUsage) SetCpu(v float32) *DescribeConfigurationPriceResponseBodyDataBagUsage {
	s.Cpu = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataBagUsage) SetCu(v float32) *DescribeConfigurationPriceResponseBodyDataBagUsage {
	s.Cu = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataBagUsage) SetMem(v float32) *DescribeConfigurationPriceResponseBodyDataBagUsage {
	s.Mem = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataBagUsage) Validate() error {
	return dara.Validate(s)
}

type DescribeConfigurationPriceResponseBodyDataCpuMemPrice struct {
	// The information about pricing.
	Order *DescribeConfigurationPriceResponseBodyDataCpuMemPriceOrder `json:"Order,omitempty" xml:"Order,omitempty" type:"Struct"`
	// The discount rules.
	Rules []*DescribeConfigurationPriceResponseBodyDataCpuMemPriceRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeConfigurationPriceResponseBodyDataCpuMemPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigurationPriceResponseBodyDataCpuMemPrice) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponseBodyDataCpuMemPrice) GetOrder() *DescribeConfigurationPriceResponseBodyDataCpuMemPriceOrder {
	return s.Order
}

func (s *DescribeConfigurationPriceResponseBodyDataCpuMemPrice) GetRules() []*DescribeConfigurationPriceResponseBodyDataCpuMemPriceRules {
	return s.Rules
}

func (s *DescribeConfigurationPriceResponseBodyDataCpuMemPrice) SetOrder(v *DescribeConfigurationPriceResponseBodyDataCpuMemPriceOrder) *DescribeConfigurationPriceResponseBodyDataCpuMemPrice {
	s.Order = v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataCpuMemPrice) SetRules(v []*DescribeConfigurationPriceResponseBodyDataCpuMemPriceRules) *DescribeConfigurationPriceResponseBodyDataCpuMemPrice {
	s.Rules = v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataCpuMemPrice) Validate() error {
	if s.Order != nil {
		if err := s.Order.Validate(); err != nil {
			return err
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

type DescribeConfigurationPriceResponseBodyDataCpuMemPriceOrder struct {
	// The discount amount.
	//
	// example:
	//
	// 0.0009259
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The original price.
	//
	// example:
	//
	// 0.0046296
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// The ID of the discount rule.
	RuleIds []*string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
	// The final price of the order.
	//
	// example:
	//
	// 0.0037037
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s DescribeConfigurationPriceResponseBodyDataCpuMemPriceOrder) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigurationPriceResponseBodyDataCpuMemPriceOrder) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponseBodyDataCpuMemPriceOrder) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *DescribeConfigurationPriceResponseBodyDataCpuMemPriceOrder) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *DescribeConfigurationPriceResponseBodyDataCpuMemPriceOrder) GetRuleIds() []*string {
	return s.RuleIds
}

func (s *DescribeConfigurationPriceResponseBodyDataCpuMemPriceOrder) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *DescribeConfigurationPriceResponseBodyDataCpuMemPriceOrder) SetDiscountAmount(v float32) *DescribeConfigurationPriceResponseBodyDataCpuMemPriceOrder {
	s.DiscountAmount = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataCpuMemPriceOrder) SetOriginalAmount(v float32) *DescribeConfigurationPriceResponseBodyDataCpuMemPriceOrder {
	s.OriginalAmount = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataCpuMemPriceOrder) SetRuleIds(v []*string) *DescribeConfigurationPriceResponseBodyDataCpuMemPriceOrder {
	s.RuleIds = v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataCpuMemPriceOrder) SetTradeAmount(v float32) *DescribeConfigurationPriceResponseBodyDataCpuMemPriceOrder {
	s.TradeAmount = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataCpuMemPriceOrder) Validate() error {
	return dara.Validate(s)
}

type DescribeConfigurationPriceResponseBodyDataCpuMemPriceRules struct {
	// The name of discount rule.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the discount rule.
	//
	// example:
	//
	// 2000010******
	RuleDescId *float32 `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s DescribeConfigurationPriceResponseBodyDataCpuMemPriceRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigurationPriceResponseBodyDataCpuMemPriceRules) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponseBodyDataCpuMemPriceRules) GetName() *string {
	return s.Name
}

func (s *DescribeConfigurationPriceResponseBodyDataCpuMemPriceRules) GetRuleDescId() *float32 {
	return s.RuleDescId
}

func (s *DescribeConfigurationPriceResponseBodyDataCpuMemPriceRules) SetName(v string) *DescribeConfigurationPriceResponseBodyDataCpuMemPriceRules {
	s.Name = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataCpuMemPriceRules) SetRuleDescId(v float32) *DescribeConfigurationPriceResponseBodyDataCpuMemPriceRules {
	s.RuleDescId = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataCpuMemPriceRules) Validate() error {
	return dara.Validate(s)
}

type DescribeConfigurationPriceResponseBodyDataOrder struct {
	// The discount amount.
	//
	// example:
	//
	// 0.0018518
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The original price of the order.
	//
	// example:
	//
	// 0.0092592
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// The ID of the promotion rule.
	RuleIds []*string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
	// The transaction price.
	//
	// example:
	//
	// 0.0074074
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s DescribeConfigurationPriceResponseBodyDataOrder) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigurationPriceResponseBodyDataOrder) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponseBodyDataOrder) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *DescribeConfigurationPriceResponseBodyDataOrder) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *DescribeConfigurationPriceResponseBodyDataOrder) GetRuleIds() []*string {
	return s.RuleIds
}

func (s *DescribeConfigurationPriceResponseBodyDataOrder) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *DescribeConfigurationPriceResponseBodyDataOrder) SetDiscountAmount(v float32) *DescribeConfigurationPriceResponseBodyDataOrder {
	s.DiscountAmount = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataOrder) SetOriginalAmount(v float32) *DescribeConfigurationPriceResponseBodyDataOrder {
	s.OriginalAmount = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataOrder) SetRuleIds(v []*string) *DescribeConfigurationPriceResponseBodyDataOrder {
	s.RuleIds = v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataOrder) SetTradeAmount(v float32) *DescribeConfigurationPriceResponseBodyDataOrder {
	s.TradeAmount = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataOrder) Validate() error {
	return dara.Validate(s)
}

type DescribeConfigurationPriceResponseBodyDataRequestPrice struct {
	// The information about pricing.
	Order *DescribeConfigurationPriceResponseBodyDataRequestPriceOrder `json:"Order,omitempty" xml:"Order,omitempty" type:"Struct"`
	// The discount rule.
	Rules []*DescribeConfigurationPriceResponseBodyDataRequestPriceRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeConfigurationPriceResponseBodyDataRequestPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigurationPriceResponseBodyDataRequestPrice) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponseBodyDataRequestPrice) GetOrder() *DescribeConfigurationPriceResponseBodyDataRequestPriceOrder {
	return s.Order
}

func (s *DescribeConfigurationPriceResponseBodyDataRequestPrice) GetRules() []*DescribeConfigurationPriceResponseBodyDataRequestPriceRules {
	return s.Rules
}

func (s *DescribeConfigurationPriceResponseBodyDataRequestPrice) SetOrder(v *DescribeConfigurationPriceResponseBodyDataRequestPriceOrder) *DescribeConfigurationPriceResponseBodyDataRequestPrice {
	s.Order = v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataRequestPrice) SetRules(v []*DescribeConfigurationPriceResponseBodyDataRequestPriceRules) *DescribeConfigurationPriceResponseBodyDataRequestPrice {
	s.Rules = v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataRequestPrice) Validate() error {
	if s.Order != nil {
		if err := s.Order.Validate(); err != nil {
			return err
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

type DescribeConfigurationPriceResponseBodyDataRequestPriceOrder struct {
	// The discount amount.
	//
	// example:
	//
	// 0.0009259
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The original price of the order.
	//
	// example:
	//
	// 0.0046296
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// The ID of the discount rule.
	RuleIds []*string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
	// The actual price of the order.
	//
	// example:
	//
	// 0.0037037
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s DescribeConfigurationPriceResponseBodyDataRequestPriceOrder) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigurationPriceResponseBodyDataRequestPriceOrder) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponseBodyDataRequestPriceOrder) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *DescribeConfigurationPriceResponseBodyDataRequestPriceOrder) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *DescribeConfigurationPriceResponseBodyDataRequestPriceOrder) GetRuleIds() []*string {
	return s.RuleIds
}

func (s *DescribeConfigurationPriceResponseBodyDataRequestPriceOrder) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *DescribeConfigurationPriceResponseBodyDataRequestPriceOrder) SetDiscountAmount(v float32) *DescribeConfigurationPriceResponseBodyDataRequestPriceOrder {
	s.DiscountAmount = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataRequestPriceOrder) SetOriginalAmount(v float32) *DescribeConfigurationPriceResponseBodyDataRequestPriceOrder {
	s.OriginalAmount = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataRequestPriceOrder) SetRuleIds(v []*string) *DescribeConfigurationPriceResponseBodyDataRequestPriceOrder {
	s.RuleIds = v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataRequestPriceOrder) SetTradeAmount(v float32) *DescribeConfigurationPriceResponseBodyDataRequestPriceOrder {
	s.TradeAmount = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataRequestPriceOrder) Validate() error {
	return dara.Validate(s)
}

type DescribeConfigurationPriceResponseBodyDataRequestPriceRules struct {
	// The name of the discount rule.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the discount policy.
	//
	// example:
	//
	// 2000010******
	RuleDescId *int64 `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s DescribeConfigurationPriceResponseBodyDataRequestPriceRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigurationPriceResponseBodyDataRequestPriceRules) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponseBodyDataRequestPriceRules) GetName() *string {
	return s.Name
}

func (s *DescribeConfigurationPriceResponseBodyDataRequestPriceRules) GetRuleDescId() *int64 {
	return s.RuleDescId
}

func (s *DescribeConfigurationPriceResponseBodyDataRequestPriceRules) SetName(v string) *DescribeConfigurationPriceResponseBodyDataRequestPriceRules {
	s.Name = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataRequestPriceRules) SetRuleDescId(v int64) *DescribeConfigurationPriceResponseBodyDataRequestPriceRules {
	s.RuleDescId = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataRequestPriceRules) Validate() error {
	return dara.Validate(s)
}

type DescribeConfigurationPriceResponseBodyDataRules struct {
	// The name of the promotion rule.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the promotion rule.
	//
	// example:
	//
	// 2000010******
	RuleDescId *int64 `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s DescribeConfigurationPriceResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigurationPriceResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponseBodyDataRules) GetName() *string {
	return s.Name
}

func (s *DescribeConfigurationPriceResponseBodyDataRules) GetRuleDescId() *int64 {
	return s.RuleDescId
}

func (s *DescribeConfigurationPriceResponseBodyDataRules) SetName(v string) *DescribeConfigurationPriceResponseBodyDataRules {
	s.Name = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataRules) SetRuleDescId(v int64) *DescribeConfigurationPriceResponseBodyDataRules {
	s.RuleDescId = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}

type DescribeConfigurationPriceResponseBodyDataTrafficPrice struct {
	// The information about pricing.
	Order *DescribeConfigurationPriceResponseBodyDataTrafficPriceOrder `json:"Order,omitempty" xml:"Order,omitempty" type:"Struct"`
	// The discount rule.
	Rules []*DescribeConfigurationPriceResponseBodyDataTrafficPriceRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeConfigurationPriceResponseBodyDataTrafficPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigurationPriceResponseBodyDataTrafficPrice) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponseBodyDataTrafficPrice) GetOrder() *DescribeConfigurationPriceResponseBodyDataTrafficPriceOrder {
	return s.Order
}

func (s *DescribeConfigurationPriceResponseBodyDataTrafficPrice) GetRules() []*DescribeConfigurationPriceResponseBodyDataTrafficPriceRules {
	return s.Rules
}

func (s *DescribeConfigurationPriceResponseBodyDataTrafficPrice) SetOrder(v *DescribeConfigurationPriceResponseBodyDataTrafficPriceOrder) *DescribeConfigurationPriceResponseBodyDataTrafficPrice {
	s.Order = v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataTrafficPrice) SetRules(v []*DescribeConfigurationPriceResponseBodyDataTrafficPriceRules) *DescribeConfigurationPriceResponseBodyDataTrafficPrice {
	s.Rules = v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataTrafficPrice) Validate() error {
	if s.Order != nil {
		if err := s.Order.Validate(); err != nil {
			return err
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

type DescribeConfigurationPriceResponseBodyDataTrafficPriceOrder struct {
	// The discount amount.
	//
	// example:
	//
	// 0.0009259
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The original price of the order.
	//
	// example:
	//
	// 0.0046296
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// The ID of the discount rule.
	RuleIds []*string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
	// The final price of the order.
	//
	// example:
	//
	// 0.0037037
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s DescribeConfigurationPriceResponseBodyDataTrafficPriceOrder) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigurationPriceResponseBodyDataTrafficPriceOrder) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponseBodyDataTrafficPriceOrder) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *DescribeConfigurationPriceResponseBodyDataTrafficPriceOrder) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *DescribeConfigurationPriceResponseBodyDataTrafficPriceOrder) GetRuleIds() []*string {
	return s.RuleIds
}

func (s *DescribeConfigurationPriceResponseBodyDataTrafficPriceOrder) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *DescribeConfigurationPriceResponseBodyDataTrafficPriceOrder) SetDiscountAmount(v float32) *DescribeConfigurationPriceResponseBodyDataTrafficPriceOrder {
	s.DiscountAmount = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataTrafficPriceOrder) SetOriginalAmount(v float32) *DescribeConfigurationPriceResponseBodyDataTrafficPriceOrder {
	s.OriginalAmount = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataTrafficPriceOrder) SetRuleIds(v []*string) *DescribeConfigurationPriceResponseBodyDataTrafficPriceOrder {
	s.RuleIds = v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataTrafficPriceOrder) SetTradeAmount(v float32) *DescribeConfigurationPriceResponseBodyDataTrafficPriceOrder {
	s.TradeAmount = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataTrafficPriceOrder) Validate() error {
	return dara.Validate(s)
}

type DescribeConfigurationPriceResponseBodyDataTrafficPriceRules struct {
	// The name of the discount rule.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the discount rule.
	//
	// example:
	//
	// 2000010******
	RuleDescId *float32 `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s DescribeConfigurationPriceResponseBodyDataTrafficPriceRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigurationPriceResponseBodyDataTrafficPriceRules) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponseBodyDataTrafficPriceRules) GetName() *string {
	return s.Name
}

func (s *DescribeConfigurationPriceResponseBodyDataTrafficPriceRules) GetRuleDescId() *float32 {
	return s.RuleDescId
}

func (s *DescribeConfigurationPriceResponseBodyDataTrafficPriceRules) SetName(v string) *DescribeConfigurationPriceResponseBodyDataTrafficPriceRules {
	s.Name = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataTrafficPriceRules) SetRuleDescId(v float32) *DescribeConfigurationPriceResponseBodyDataTrafficPriceRules {
	s.RuleDescId = &v
	return s
}

func (s *DescribeConfigurationPriceResponseBodyDataTrafficPriceRules) Validate() error {
	return dara.Validate(s)
}
