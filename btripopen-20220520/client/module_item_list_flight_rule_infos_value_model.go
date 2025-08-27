// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleItemListFlightRuleInfosValue interface {
	dara.Model
	String() string
	GoString() string
	SetRefundChangeRuleDesc(v string) *ModuleItemListFlightRuleInfosValue
	GetRefundChangeRuleDesc() *string
	SetBaggageDesc(v string) *ModuleItemListFlightRuleInfosValue
	GetBaggageDesc() *string
}

type ModuleItemListFlightRuleInfosValue struct {
	RefundChangeRuleDesc *string `json:"refund_change_rule_desc,omitempty" xml:"refund_change_rule_desc,omitempty"`
	BaggageDesc          *string `json:"baggage_desc,omitempty" xml:"baggage_desc,omitempty"`
}

func (s ModuleItemListFlightRuleInfosValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListFlightRuleInfosValue) GoString() string {
	return s.String()
}

func (s *ModuleItemListFlightRuleInfosValue) GetRefundChangeRuleDesc() *string {
	return s.RefundChangeRuleDesc
}

func (s *ModuleItemListFlightRuleInfosValue) GetBaggageDesc() *string {
	return s.BaggageDesc
}

func (s *ModuleItemListFlightRuleInfosValue) SetRefundChangeRuleDesc(v string) *ModuleItemListFlightRuleInfosValue {
	s.RefundChangeRuleDesc = &v
	return s
}

func (s *ModuleItemListFlightRuleInfosValue) SetBaggageDesc(v string) *ModuleItemListFlightRuleInfosValue {
	s.BaggageDesc = &v
	return s
}

func (s *ModuleItemListFlightRuleInfosValue) Validate() error {
	return dara.Validate(s)
}
