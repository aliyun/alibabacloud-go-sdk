// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleFlightItemListItemListFlightRuleInfosValue interface {
	dara.Model
	String() string
	GoString() string
	SetRefundChangeRuleDesc(v string) *ModuleFlightItemListItemListFlightRuleInfosValue
	GetRefundChangeRuleDesc() *string
	SetBaggageDesc(v string) *ModuleFlightItemListItemListFlightRuleInfosValue
	GetBaggageDesc() *string
}

type ModuleFlightItemListItemListFlightRuleInfosValue struct {
	RefundChangeRuleDesc *string `json:"refund_change_rule_desc,omitempty" xml:"refund_change_rule_desc,omitempty"`
	BaggageDesc          *string `json:"baggage_desc,omitempty" xml:"baggage_desc,omitempty"`
}

func (s ModuleFlightItemListItemListFlightRuleInfosValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListItemListFlightRuleInfosValue) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListItemListFlightRuleInfosValue) GetRefundChangeRuleDesc() *string {
	return s.RefundChangeRuleDesc
}

func (s *ModuleFlightItemListItemListFlightRuleInfosValue) GetBaggageDesc() *string {
	return s.BaggageDesc
}

func (s *ModuleFlightItemListItemListFlightRuleInfosValue) SetRefundChangeRuleDesc(v string) *ModuleFlightItemListItemListFlightRuleInfosValue {
	s.RefundChangeRuleDesc = &v
	return s
}

func (s *ModuleFlightItemListItemListFlightRuleInfosValue) SetBaggageDesc(v string) *ModuleFlightItemListItemListFlightRuleInfosValue {
	s.BaggageDesc = &v
	return s
}

func (s *ModuleFlightItemListItemListFlightRuleInfosValue) Validate() error {
	return dara.Validate(s)
}
