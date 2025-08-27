// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleFlightItemListBestPriceItemFlightRuleInfosValue interface {
	dara.Model
	String() string
	GoString() string
	SetRefundChangeRuleDesc(v string) *ModuleFlightItemListBestPriceItemFlightRuleInfosValue
	GetRefundChangeRuleDesc() *string
	SetBaggageDesc(v string) *ModuleFlightItemListBestPriceItemFlightRuleInfosValue
	GetBaggageDesc() *string
}

type ModuleFlightItemListBestPriceItemFlightRuleInfosValue struct {
	RefundChangeRuleDesc *string `json:"refund_change_rule_desc,omitempty" xml:"refund_change_rule_desc,omitempty"`
	BaggageDesc          *string `json:"baggage_desc,omitempty" xml:"baggage_desc,omitempty"`
}

func (s ModuleFlightItemListBestPriceItemFlightRuleInfosValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemFlightRuleInfosValue) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemFlightRuleInfosValue) GetRefundChangeRuleDesc() *string {
	return s.RefundChangeRuleDesc
}

func (s *ModuleFlightItemListBestPriceItemFlightRuleInfosValue) GetBaggageDesc() *string {
	return s.BaggageDesc
}

func (s *ModuleFlightItemListBestPriceItemFlightRuleInfosValue) SetRefundChangeRuleDesc(v string) *ModuleFlightItemListBestPriceItemFlightRuleInfosValue {
	s.RefundChangeRuleDesc = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemFlightRuleInfosValue) SetBaggageDesc(v string) *ModuleFlightItemListBestPriceItemFlightRuleInfosValue {
	s.BaggageDesc = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemFlightRuleInfosValue) Validate() error {
	return dara.Validate(s)
}
