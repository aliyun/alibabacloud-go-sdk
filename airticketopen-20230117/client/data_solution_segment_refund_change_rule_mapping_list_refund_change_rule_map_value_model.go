// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetRefundRuleAllUnusedList(v []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue
	GetRefundRuleAllUnusedList() []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList
	SetRefundRulePartUnusedList(v []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue
	GetRefundRulePartUnusedList() []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList
	SetChangeRuleInUnusedList(v []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue
	GetChangeRuleInUnusedList() []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList
	SetChangeRuleOutUnusedList(v []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue
	GetChangeRuleOutUnusedList() []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList
}

type DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue struct {
	// refund rule for fully-unused tickets
	RefundRuleAllUnusedList []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList `json:"refund_rule_all_unused_list,omitempty" xml:"refund_rule_all_unused_list,omitempty" type:"Repeated"`
	// refund rule for partially-used tickets
	RefundRulePartUnusedList []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList `json:"refund_rule_part_unused_list,omitempty" xml:"refund_rule_part_unused_list,omitempty" type:"Repeated"`
	// change rule for inbound segment unused tickets
	ChangeRuleInUnusedList []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList `json:"change_rule_in_unused_list,omitempty" xml:"change_rule_in_unused_list,omitempty" type:"Repeated"`
	// change rule for outbound-flight-unused tickets
	ChangeRuleOutUnusedList []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList `json:"change_rule_out_unused_list,omitempty" xml:"change_rule_out_unused_list,omitempty" type:"Repeated"`
}

func (s DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) String() string {
	return dara.Prettify(s)
}

func (s DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) GoString() string {
	return s.String()
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) GetRefundRuleAllUnusedList() []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList {
	return s.RefundRuleAllUnusedList
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) GetRefundRulePartUnusedList() []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList {
	return s.RefundRulePartUnusedList
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) GetChangeRuleInUnusedList() []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList {
	return s.ChangeRuleInUnusedList
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) GetChangeRuleOutUnusedList() []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList {
	return s.ChangeRuleOutUnusedList
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) SetRefundRuleAllUnusedList(v []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue {
	s.RefundRuleAllUnusedList = v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) SetRefundRulePartUnusedList(v []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue {
	s.RefundRulePartUnusedList = v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) SetChangeRuleInUnusedList(v []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue {
	s.ChangeRuleInUnusedList = v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) SetChangeRuleOutUnusedList(v []*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue {
	s.ChangeRuleOutUnusedList = v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) Validate() error {
	if s.RefundRuleAllUnusedList != nil {
		for _, item := range s.RefundRuleAllUnusedList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RefundRulePartUnusedList != nil {
		for _, item := range s.RefundRulePartUnusedList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ChangeRuleInUnusedList != nil {
		for _, item := range s.ChangeRuleInUnusedList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ChangeRuleOutUnusedList != nil {
		for _, item := range s.ChangeRuleOutUnusedList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList struct {
	// type: 0 - fully-unused ticket; 1 - partially used ticket
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// time unit: day/hour
	//
	// example:
	//
	// hour
	TimeUnit *string `json:"time_unit,omitempty" xml:"time_unit,omitempty"`
	// applicable refund rule start time, time unit (day/hour)
	//
	// example:
	//
	// 36
	RuleStartTime *int32 `json:"rule_start_time,omitempty" xml:"rule_start_time,omitempty"`
	// applicable refund rule end time, time unit (day/hour)
	//
	// example:
	//
	// 12
	RuleEndTime *int32 `json:"rule_end_time,omitempty" xml:"rule_end_time,omitempty"`
	// whether refundable X-Y hour(day) before departure
	//
	// example:
	//
	// true
	CanRefund *bool `json:"can_refund,omitempty" xml:"can_refund,omitempty"`
	// refund fee X-Y hour(day) before departure
	//
	// example:
	//
	// 200
	RefundFee *float64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// whether tax is fully refundable X-Y hour(day) before departure
	//
	// example:
	//
	// true
	CanReturnAllTax *bool `json:"can_return_all_tax,omitempty" xml:"can_return_all_tax,omitempty"`
	// tax amount refundable X-Y hour(day) before departure
	//
	// example:
	//
	// 100
	ReturnPartTaxFee *float64 `json:"return_part_tax_fee,omitempty" xml:"return_part_tax_fee,omitempty"`
}

func (s DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) String() string {
	return dara.Prettify(s)
}

func (s DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GoString() string {
	return s.String()
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GetType() *int32 {
	return s.Type
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GetRuleStartTime() *int32 {
	return s.RuleStartTime
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GetRuleEndTime() *int32 {
	return s.RuleEndTime
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GetCanRefund() *bool {
	return s.CanRefund
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GetRefundFee() *float64 {
	return s.RefundFee
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GetCanReturnAllTax() *bool {
	return s.CanReturnAllTax
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GetReturnPartTaxFee() *float64 {
	return s.ReturnPartTaxFee
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) SetType(v int32) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.Type = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) SetTimeUnit(v string) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.TimeUnit = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) SetRuleStartTime(v int32) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.RuleStartTime = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) SetRuleEndTime(v int32) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.RuleEndTime = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) SetCanRefund(v bool) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.CanRefund = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) SetRefundFee(v float64) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.RefundFee = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) SetCanReturnAllTax(v bool) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.CanReturnAllTax = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) SetReturnPartTaxFee(v float64) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.ReturnPartTaxFee = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) Validate() error {
	return dara.Validate(s)
}

type DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList struct {
	// type: 0 - fully-unused ticket; 1 - partially used ticket
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// time unit: day/hour
	//
	// example:
	//
	// hour
	TimeUnit *string `json:"time_unit,omitempty" xml:"time_unit,omitempty"`
	// applicable refund rule start time, time unit (day/hour)
	//
	// example:
	//
	// 36
	RuleStartTime *int32 `json:"rule_start_time,omitempty" xml:"rule_start_time,omitempty"`
	// applicable refund rule end time, time unit (day/hour)
	//
	// example:
	//
	// 12
	RuleEndTime *int32 `json:"rule_end_time,omitempty" xml:"rule_end_time,omitempty"`
	// whether refundable X-Y hour(day) before departure
	//
	// example:
	//
	// true
	CanRefund *bool `json:"can_refund,omitempty" xml:"can_refund,omitempty"`
	// refund fee X-Y hour(day) before departure
	//
	// example:
	//
	// 200
	RefundFee *float64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// whether tax is fully refundable X-Y hour(day) before departure
	//
	// example:
	//
	// true
	CanReturnAllTax *bool `json:"can_return_all_tax,omitempty" xml:"can_return_all_tax,omitempty"`
	// tax amount refundable X-Y hour(day) before departure
	//
	// example:
	//
	// 100
	ReturnPartTaxFee *float64 `json:"return_part_tax_fee,omitempty" xml:"return_part_tax_fee,omitempty"`
}

func (s DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) String() string {
	return dara.Prettify(s)
}

func (s DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GoString() string {
	return s.String()
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GetType() *int32 {
	return s.Type
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GetRuleStartTime() *int32 {
	return s.RuleStartTime
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GetRuleEndTime() *int32 {
	return s.RuleEndTime
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GetCanRefund() *bool {
	return s.CanRefund
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GetRefundFee() *float64 {
	return s.RefundFee
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GetCanReturnAllTax() *bool {
	return s.CanReturnAllTax
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GetReturnPartTaxFee() *float64 {
	return s.ReturnPartTaxFee
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) SetType(v int32) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.Type = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) SetTimeUnit(v string) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.TimeUnit = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) SetRuleStartTime(v int32) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.RuleStartTime = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) SetRuleEndTime(v int32) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.RuleEndTime = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) SetCanRefund(v bool) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.CanRefund = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) SetRefundFee(v float64) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.RefundFee = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) SetCanReturnAllTax(v bool) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.CanReturnAllTax = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) SetReturnPartTaxFee(v float64) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.ReturnPartTaxFee = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) Validate() error {
	return dara.Validate(s)
}

type DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList struct {
	// type: 2 - outbound segment unused; 3 - inbound segment unused
	//
	// example:
	//
	// 2
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// time unit: day/hour
	//
	// example:
	//
	// hour
	TimeUnit *string `json:"time_unit,omitempty" xml:"time_unit,omitempty"`
	// applicable change rule start time, time unit (day/hour)
	//
	// example:
	//
	// 36
	RuleStartTime *int32 `json:"rule_start_time,omitempty" xml:"rule_start_time,omitempty"`
	// applicable change rule end time, time unit (day/hour)
	//
	// example:
	//
	// 12
	RuleEndTime *int32 `json:"rule_end_time,omitempty" xml:"rule_end_time,omitempty"`
	// whether changeable X-Y hour(day) before departure
	//
	// example:
	//
	// true
	CanChange *bool `json:"can_change,omitempty" xml:"can_change,omitempty"`
	// change fee X-Y hour(day) before departure
	//
	// example:
	//
	// 100
	ChangeFee *float64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
}

func (s DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) String() string {
	return dara.Prettify(s)
}

func (s DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) GoString() string {
	return s.String()
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) GetType() *int32 {
	return s.Type
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) GetRuleStartTime() *int32 {
	return s.RuleStartTime
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) GetRuleEndTime() *int32 {
	return s.RuleEndTime
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) GetCanChange() *bool {
	return s.CanChange
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) GetChangeFee() *float64 {
	return s.ChangeFee
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) SetType(v int32) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList {
	s.Type = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) SetTimeUnit(v string) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList {
	s.TimeUnit = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) SetRuleStartTime(v int32) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList {
	s.RuleStartTime = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) SetRuleEndTime(v int32) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList {
	s.RuleEndTime = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) SetCanChange(v bool) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList {
	s.CanChange = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) SetChangeFee(v float64) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList {
	s.ChangeFee = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) Validate() error {
	return dara.Validate(s)
}

type DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList struct {
	// type: 2 - outbound segment unused; 3 - inbound segment unused
	//
	// example:
	//
	// 2
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// time unit: day/hour
	//
	// example:
	//
	// hour
	TimeUnit *string `json:"time_unit,omitempty" xml:"time_unit,omitempty"`
	// applicable change rule start time, time unit (day/hour)
	//
	// example:
	//
	// 36
	RuleStartTime *int32 `json:"rule_start_time,omitempty" xml:"rule_start_time,omitempty"`
	// applicable change rule end time, time unit (day/hour)
	//
	// example:
	//
	// 12
	RuleEndTime *int32 `json:"rule_end_time,omitempty" xml:"rule_end_time,omitempty"`
	// whether changeable X-Y hour(day) before departure
	//
	// example:
	//
	// true
	CanChange *bool `json:"can_change,omitempty" xml:"can_change,omitempty"`
	// change fee X-Y hour(day) before departure
	//
	// example:
	//
	// 100
	ChangeFee *float64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
}

func (s DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) String() string {
	return dara.Prettify(s)
}

func (s DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) GoString() string {
	return s.String()
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) GetType() *int32 {
	return s.Type
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) GetRuleStartTime() *int32 {
	return s.RuleStartTime
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) GetRuleEndTime() *int32 {
	return s.RuleEndTime
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) GetCanChange() *bool {
	return s.CanChange
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) GetChangeFee() *float64 {
	return s.ChangeFee
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) SetType(v int32) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList {
	s.Type = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) SetTimeUnit(v string) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList {
	s.TimeUnit = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) SetRuleStartTime(v int32) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList {
	s.RuleStartTime = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) SetRuleEndTime(v int32) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList {
	s.RuleEndTime = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) SetCanChange(v bool) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList {
	s.CanChange = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) SetChangeFee(v float64) *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList {
	s.ChangeFee = &v
	return s
}

func (s *DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) Validate() error {
	return dara.Validate(s)
}
