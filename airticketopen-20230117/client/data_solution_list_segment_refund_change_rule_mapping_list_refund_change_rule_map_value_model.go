// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetRefundRuleAllUnusedList(v []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue
	GetRefundRuleAllUnusedList() []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList
	SetRefundRulePartUnusedList(v []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue
	GetRefundRulePartUnusedList() []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList
	SetChangeRuleInUnusedList(v []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue
	GetChangeRuleInUnusedList() []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList
	SetChangeRuleOutUnusedList(v []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue
	GetChangeRuleOutUnusedList() []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList
}

type DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue struct {
	// refund rule for fully-unused tickets
	RefundRuleAllUnusedList []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList `json:"refund_rule_all_unused_list,omitempty" xml:"refund_rule_all_unused_list,omitempty" type:"Repeated"`
	// refund rule for partially-used tickets
	RefundRulePartUnusedList []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList `json:"refund_rule_part_unused_list,omitempty" xml:"refund_rule_part_unused_list,omitempty" type:"Repeated"`
	// change rule for inbound segment unused tickets
	ChangeRuleInUnusedList []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList `json:"change_rule_in_unused_list,omitempty" xml:"change_rule_in_unused_list,omitempty" type:"Repeated"`
	// change rule for outbound segment unused tickets
	ChangeRuleOutUnusedList []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList `json:"change_rule_out_unused_list,omitempty" xml:"change_rule_out_unused_list,omitempty" type:"Repeated"`
}

func (s DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) String() string {
	return dara.Prettify(s)
}

func (s DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) GoString() string {
	return s.String()
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) GetRefundRuleAllUnusedList() []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList {
	return s.RefundRuleAllUnusedList
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) GetRefundRulePartUnusedList() []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList {
	return s.RefundRulePartUnusedList
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) GetChangeRuleInUnusedList() []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList {
	return s.ChangeRuleInUnusedList
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) GetChangeRuleOutUnusedList() []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList {
	return s.ChangeRuleOutUnusedList
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) SetRefundRuleAllUnusedList(v []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue {
	s.RefundRuleAllUnusedList = v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) SetRefundRulePartUnusedList(v []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue {
	s.RefundRulePartUnusedList = v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) SetChangeRuleInUnusedList(v []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue {
	s.ChangeRuleInUnusedList = v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) SetChangeRuleOutUnusedList(v []*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue {
	s.ChangeRuleOutUnusedList = v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) Validate() error {
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

type DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList struct {
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
	// 20
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
	// 20
	ReturnPartTaxFee *float64 `json:"return_part_tax_fee,omitempty" xml:"return_part_tax_fee,omitempty"`
}

func (s DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) String() string {
	return dara.Prettify(s)
}

func (s DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GoString() string {
	return s.String()
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GetType() *int32 {
	return s.Type
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GetRuleStartTime() *int32 {
	return s.RuleStartTime
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GetRuleEndTime() *int32 {
	return s.RuleEndTime
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GetCanRefund() *bool {
	return s.CanRefund
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GetRefundFee() *float64 {
	return s.RefundFee
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GetCanReturnAllTax() *bool {
	return s.CanReturnAllTax
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GetReturnPartTaxFee() *float64 {
	return s.ReturnPartTaxFee
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) SetType(v int32) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.Type = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) SetTimeUnit(v string) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.TimeUnit = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) SetRuleStartTime(v int32) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.RuleStartTime = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) SetRuleEndTime(v int32) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.RuleEndTime = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) SetCanRefund(v bool) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.CanRefund = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) SetRefundFee(v float64) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.RefundFee = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) SetCanReturnAllTax(v bool) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.CanReturnAllTax = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) SetReturnPartTaxFee(v float64) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.ReturnPartTaxFee = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) Validate() error {
	return dara.Validate(s)
}

type DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList struct {
	// type: 0 - fully-unused ticket; 1 - partially used ticket
	//
	// example:
	//
	// 1
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
	// 20
	RefundFee *float64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// whether tax is fully refundable X-Y hour(day) before departure
	//
	// example:
	//
	// 30
	CanReturnAllTax *bool `json:"can_return_all_tax,omitempty" xml:"can_return_all_tax,omitempty"`
	// tax amount refundable X-Y hour(day) before departure
	//
	// example:
	//
	// 20
	ReturnPartTaxFee *float64 `json:"return_part_tax_fee,omitempty" xml:"return_part_tax_fee,omitempty"`
}

func (s DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) String() string {
	return dara.Prettify(s)
}

func (s DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GoString() string {
	return s.String()
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GetType() *int32 {
	return s.Type
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GetRuleStartTime() *int32 {
	return s.RuleStartTime
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GetRuleEndTime() *int32 {
	return s.RuleEndTime
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GetCanRefund() *bool {
	return s.CanRefund
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GetRefundFee() *float64 {
	return s.RefundFee
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GetCanReturnAllTax() *bool {
	return s.CanReturnAllTax
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GetReturnPartTaxFee() *float64 {
	return s.ReturnPartTaxFee
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) SetType(v int32) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.Type = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) SetTimeUnit(v string) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.TimeUnit = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) SetRuleStartTime(v int32) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.RuleStartTime = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) SetRuleEndTime(v int32) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.RuleEndTime = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) SetCanRefund(v bool) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.CanRefund = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) SetRefundFee(v float64) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.RefundFee = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) SetCanReturnAllTax(v bool) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.CanReturnAllTax = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) SetReturnPartTaxFee(v float64) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.ReturnPartTaxFee = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) Validate() error {
	return dara.Validate(s)
}

type DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList struct {
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
	// 20
	ChangeFee *float64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
}

func (s DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) String() string {
	return dara.Prettify(s)
}

func (s DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) GoString() string {
	return s.String()
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) GetType() *int32 {
	return s.Type
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) GetRuleStartTime() *int32 {
	return s.RuleStartTime
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) GetRuleEndTime() *int32 {
	return s.RuleEndTime
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) GetCanChange() *bool {
	return s.CanChange
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) GetChangeFee() *float64 {
	return s.ChangeFee
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) SetType(v int32) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList {
	s.Type = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) SetTimeUnit(v string) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList {
	s.TimeUnit = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) SetRuleStartTime(v int32) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList {
	s.RuleStartTime = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) SetRuleEndTime(v int32) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList {
	s.RuleEndTime = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) SetCanChange(v bool) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList {
	s.CanChange = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) SetChangeFee(v float64) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList {
	s.ChangeFee = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) Validate() error {
	return dara.Validate(s)
}

type DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList struct {
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
	// 10
	ChangeFee *float64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
}

func (s DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) String() string {
	return dara.Prettify(s)
}

func (s DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) GoString() string {
	return s.String()
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) GetType() *int32 {
	return s.Type
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) GetRuleStartTime() *int32 {
	return s.RuleStartTime
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) GetRuleEndTime() *int32 {
	return s.RuleEndTime
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) GetCanChange() *bool {
	return s.CanChange
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) GetChangeFee() *float64 {
	return s.ChangeFee
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) SetType(v int32) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList {
	s.Type = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) SetTimeUnit(v string) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList {
	s.TimeUnit = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) SetRuleStartTime(v int32) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList {
	s.RuleStartTime = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) SetRuleEndTime(v int32) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList {
	s.RuleEndTime = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) SetCanChange(v bool) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList {
	s.CanChange = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) SetChangeFee(v float64) *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList {
	s.ChangeFee = &v
	return s
}

func (s *DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) Validate() error {
	return dara.Validate(s)
}
