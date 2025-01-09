// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue struct {
	// checked baggage quantity
	//
	// example:
	//
	// 22
	BaggageAmount *int32 `json:"baggage_amount,omitempty" xml:"baggage_amount,omitempty"`
	// checked baggage weight
	//
	// example:
	//
	// 2
	BaggageWeight *int32 `json:"baggage_weight,omitempty" xml:"baggage_weight,omitempty"`
	// checked baggage weight unit
	//
	// example:
	//
	// kg
	BaggageWeightUnit *string `json:"baggage_weight_unit,omitempty" xml:"baggage_weight_unit,omitempty"`
	// Whether the weight is for all baggages
	//
	// example:
	//
	// true
	IsAllWeight *bool `json:"is_all_weight,omitempty" xml:"is_all_weight,omitempty"`
	// carry-on baggage quantity
	//
	// example:
	//
	// 1
	CarryOnAmount *int32 `json:"carry_on_amount,omitempty" xml:"carry_on_amount,omitempty"`
	// carry-on baggage weight
	//
	// example:
	//
	// 2
	CarryOnWeight *int32 `json:"carry_on_weight,omitempty" xml:"carry_on_weight,omitempty"`
	// carry-on baggage weight unit
	//
	// example:
	//
	// kg
	CarryOnWeightUnit *string `json:"carry_on_weight_unit,omitempty" xml:"carry_on_weight_unit,omitempty"`
	// Whether the weight is for all baggages
	//
	// example:
	//
	// true
	IsAllCarryOnWeight *bool `json:"is_all_carry_on_weight,omitempty" xml:"is_all_carry_on_weight,omitempty"`
}

func (s DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) String() string {
	return tea.Prettify(s)
}

func (s DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GoString() string {
	return s.String()
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetBaggageAmount(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.BaggageAmount = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetBaggageWeight(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.BaggageWeight = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetBaggageWeightUnit(v string) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.BaggageWeightUnit = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetIsAllWeight(v bool) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.IsAllWeight = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarryOnAmount(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarryOnAmount = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarryOnWeight(v int32) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarryOnWeight = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarryOnWeightUnit(v string) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarryOnWeightUnit = &v
	return s
}

func (s *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetIsAllCarryOnWeight(v bool) *DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.IsAllCarryOnWeight = &v
	return s
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
	return tea.Prettify(s)
}

func (s DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) GoString() string {
	return s.String()
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

type DataBaggageAllowanceMapValue struct {
	// checked baggage quantity
	//
	// example:
	//
	// 1
	BaggageAmount *int32 `json:"baggage_amount,omitempty" xml:"baggage_amount,omitempty"`
	// checked baggage weight
	//
	// example:
	//
	// 10
	BaggageWeight *int32 `json:"baggage_weight,omitempty" xml:"baggage_weight,omitempty"`
	// checked baggage weight unit
	//
	// example:
	//
	// KG
	BaggageWeightUnit *string `json:"baggage_weight_unit,omitempty" xml:"baggage_weight_unit,omitempty"`
	// Whether the weight is for all baggages
	//
	// example:
	//
	// true
	IsAllWeight *bool `json:"is_all_weight,omitempty" xml:"is_all_weight,omitempty"`
	// carry-on baggage quantity
	//
	// example:
	//
	// 1
	CarryOnAmount *int32 `json:"carry_on_amount,omitempty" xml:"carry_on_amount,omitempty"`
	// carry-on baggage weight
	//
	// example:
	//
	// 5
	CarryOnWeight *int32 `json:"carry_on_weight,omitempty" xml:"carry_on_weight,omitempty"`
	// carry-on baggage weight unit
	//
	// example:
	//
	// KG
	CarryOnWeightUnit *string `json:"carry_on_weight_unit,omitempty" xml:"carry_on_weight_unit,omitempty"`
	// Whether the weight is for all baggages
	//
	// example:
	//
	// true
	IsAllCarryOnWeight *bool `json:"is_all_carry_on_weight,omitempty" xml:"is_all_carry_on_weight,omitempty"`
}

func (s DataBaggageAllowanceMapValue) String() string {
	return tea.Prettify(s)
}

func (s DataBaggageAllowanceMapValue) GoString() string {
	return s.String()
}

func (s *DataBaggageAllowanceMapValue) SetBaggageAmount(v int32) *DataBaggageAllowanceMapValue {
	s.BaggageAmount = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetBaggageWeight(v int32) *DataBaggageAllowanceMapValue {
	s.BaggageWeight = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetBaggageWeightUnit(v string) *DataBaggageAllowanceMapValue {
	s.BaggageWeightUnit = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetIsAllWeight(v bool) *DataBaggageAllowanceMapValue {
	s.IsAllWeight = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetCarryOnAmount(v int32) *DataBaggageAllowanceMapValue {
	s.CarryOnAmount = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetCarryOnWeight(v int32) *DataBaggageAllowanceMapValue {
	s.CarryOnWeight = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetCarryOnWeightUnit(v string) *DataBaggageAllowanceMapValue {
	s.CarryOnWeightUnit = &v
	return s
}

func (s *DataBaggageAllowanceMapValue) SetIsAllCarryOnWeight(v bool) *DataBaggageAllowanceMapValue {
	s.IsAllCarryOnWeight = &v
	return s
}

type DataRefundChangeRuleMapValue struct {
	// refund rule for fully-unused tickets
	RefundRuleAllUnusedList []*DataRefundChangeRuleMapValueRefundRuleAllUnusedList `json:"refund_rule_all_unused_list,omitempty" xml:"refund_rule_all_unused_list,omitempty" type:"Repeated"`
	// refund rule for partially-used tickets
	RefundRulePartUnusedList []*DataRefundChangeRuleMapValueRefundRulePartUnusedList `json:"refund_rule_part_unused_list,omitempty" xml:"refund_rule_part_unused_list,omitempty" type:"Repeated"`
	// change rule for inbound segment unused tickets
	ChangeRuleInUnusedList []*DataRefundChangeRuleMapValueChangeRuleInUnusedList `json:"change_rule_in_unused_list,omitempty" xml:"change_rule_in_unused_list,omitempty" type:"Repeated"`
	// change rule for outbound-flight-unused tickets
	ChangeRuleOutUnusedList []*DataRefundChangeRuleMapValueChangeRuleOutUnusedList `json:"change_rule_out_unused_list,omitempty" xml:"change_rule_out_unused_list,omitempty" type:"Repeated"`
}

func (s DataRefundChangeRuleMapValue) String() string {
	return tea.Prettify(s)
}

func (s DataRefundChangeRuleMapValue) GoString() string {
	return s.String()
}

func (s *DataRefundChangeRuleMapValue) SetRefundRuleAllUnusedList(v []*DataRefundChangeRuleMapValueRefundRuleAllUnusedList) *DataRefundChangeRuleMapValue {
	s.RefundRuleAllUnusedList = v
	return s
}

func (s *DataRefundChangeRuleMapValue) SetRefundRulePartUnusedList(v []*DataRefundChangeRuleMapValueRefundRulePartUnusedList) *DataRefundChangeRuleMapValue {
	s.RefundRulePartUnusedList = v
	return s
}

func (s *DataRefundChangeRuleMapValue) SetChangeRuleInUnusedList(v []*DataRefundChangeRuleMapValueChangeRuleInUnusedList) *DataRefundChangeRuleMapValue {
	s.ChangeRuleInUnusedList = v
	return s
}

func (s *DataRefundChangeRuleMapValue) SetChangeRuleOutUnusedList(v []*DataRefundChangeRuleMapValueChangeRuleOutUnusedList) *DataRefundChangeRuleMapValue {
	s.ChangeRuleOutUnusedList = v
	return s
}

type DataRefundChangeRuleMapValueRefundRuleAllUnusedList struct {
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

func (s DataRefundChangeRuleMapValueRefundRuleAllUnusedList) String() string {
	return tea.Prettify(s)
}

func (s DataRefundChangeRuleMapValueRefundRuleAllUnusedList) GoString() string {
	return s.String()
}

func (s *DataRefundChangeRuleMapValueRefundRuleAllUnusedList) SetType(v int32) *DataRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.Type = &v
	return s
}

func (s *DataRefundChangeRuleMapValueRefundRuleAllUnusedList) SetTimeUnit(v string) *DataRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.TimeUnit = &v
	return s
}

func (s *DataRefundChangeRuleMapValueRefundRuleAllUnusedList) SetRuleStartTime(v int32) *DataRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.RuleStartTime = &v
	return s
}

func (s *DataRefundChangeRuleMapValueRefundRuleAllUnusedList) SetRuleEndTime(v int32) *DataRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.RuleEndTime = &v
	return s
}

func (s *DataRefundChangeRuleMapValueRefundRuleAllUnusedList) SetCanRefund(v bool) *DataRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.CanRefund = &v
	return s
}

func (s *DataRefundChangeRuleMapValueRefundRuleAllUnusedList) SetRefundFee(v float64) *DataRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.RefundFee = &v
	return s
}

func (s *DataRefundChangeRuleMapValueRefundRuleAllUnusedList) SetCanReturnAllTax(v bool) *DataRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.CanReturnAllTax = &v
	return s
}

func (s *DataRefundChangeRuleMapValueRefundRuleAllUnusedList) SetReturnPartTaxFee(v float64) *DataRefundChangeRuleMapValueRefundRuleAllUnusedList {
	s.ReturnPartTaxFee = &v
	return s
}

type DataRefundChangeRuleMapValueRefundRulePartUnusedList struct {
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

func (s DataRefundChangeRuleMapValueRefundRulePartUnusedList) String() string {
	return tea.Prettify(s)
}

func (s DataRefundChangeRuleMapValueRefundRulePartUnusedList) GoString() string {
	return s.String()
}

func (s *DataRefundChangeRuleMapValueRefundRulePartUnusedList) SetType(v int32) *DataRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.Type = &v
	return s
}

func (s *DataRefundChangeRuleMapValueRefundRulePartUnusedList) SetTimeUnit(v string) *DataRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.TimeUnit = &v
	return s
}

func (s *DataRefundChangeRuleMapValueRefundRulePartUnusedList) SetRuleStartTime(v int32) *DataRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.RuleStartTime = &v
	return s
}

func (s *DataRefundChangeRuleMapValueRefundRulePartUnusedList) SetRuleEndTime(v int32) *DataRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.RuleEndTime = &v
	return s
}

func (s *DataRefundChangeRuleMapValueRefundRulePartUnusedList) SetCanRefund(v bool) *DataRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.CanRefund = &v
	return s
}

func (s *DataRefundChangeRuleMapValueRefundRulePartUnusedList) SetRefundFee(v float64) *DataRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.RefundFee = &v
	return s
}

func (s *DataRefundChangeRuleMapValueRefundRulePartUnusedList) SetCanReturnAllTax(v bool) *DataRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.CanReturnAllTax = &v
	return s
}

func (s *DataRefundChangeRuleMapValueRefundRulePartUnusedList) SetReturnPartTaxFee(v float64) *DataRefundChangeRuleMapValueRefundRulePartUnusedList {
	s.ReturnPartTaxFee = &v
	return s
}

type DataRefundChangeRuleMapValueChangeRuleInUnusedList struct {
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

func (s DataRefundChangeRuleMapValueChangeRuleInUnusedList) String() string {
	return tea.Prettify(s)
}

func (s DataRefundChangeRuleMapValueChangeRuleInUnusedList) GoString() string {
	return s.String()
}

func (s *DataRefundChangeRuleMapValueChangeRuleInUnusedList) SetType(v int32) *DataRefundChangeRuleMapValueChangeRuleInUnusedList {
	s.Type = &v
	return s
}

func (s *DataRefundChangeRuleMapValueChangeRuleInUnusedList) SetTimeUnit(v string) *DataRefundChangeRuleMapValueChangeRuleInUnusedList {
	s.TimeUnit = &v
	return s
}

func (s *DataRefundChangeRuleMapValueChangeRuleInUnusedList) SetRuleStartTime(v int32) *DataRefundChangeRuleMapValueChangeRuleInUnusedList {
	s.RuleStartTime = &v
	return s
}

func (s *DataRefundChangeRuleMapValueChangeRuleInUnusedList) SetRuleEndTime(v int32) *DataRefundChangeRuleMapValueChangeRuleInUnusedList {
	s.RuleEndTime = &v
	return s
}

func (s *DataRefundChangeRuleMapValueChangeRuleInUnusedList) SetCanChange(v bool) *DataRefundChangeRuleMapValueChangeRuleInUnusedList {
	s.CanChange = &v
	return s
}

func (s *DataRefundChangeRuleMapValueChangeRuleInUnusedList) SetChangeFee(v float64) *DataRefundChangeRuleMapValueChangeRuleInUnusedList {
	s.ChangeFee = &v
	return s
}

type DataRefundChangeRuleMapValueChangeRuleOutUnusedList struct {
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

func (s DataRefundChangeRuleMapValueChangeRuleOutUnusedList) String() string {
	return tea.Prettify(s)
}

func (s DataRefundChangeRuleMapValueChangeRuleOutUnusedList) GoString() string {
	return s.String()
}

func (s *DataRefundChangeRuleMapValueChangeRuleOutUnusedList) SetType(v int32) *DataRefundChangeRuleMapValueChangeRuleOutUnusedList {
	s.Type = &v
	return s
}

func (s *DataRefundChangeRuleMapValueChangeRuleOutUnusedList) SetTimeUnit(v string) *DataRefundChangeRuleMapValueChangeRuleOutUnusedList {
	s.TimeUnit = &v
	return s
}

func (s *DataRefundChangeRuleMapValueChangeRuleOutUnusedList) SetRuleStartTime(v int32) *DataRefundChangeRuleMapValueChangeRuleOutUnusedList {
	s.RuleStartTime = &v
	return s
}

func (s *DataRefundChangeRuleMapValueChangeRuleOutUnusedList) SetRuleEndTime(v int32) *DataRefundChangeRuleMapValueChangeRuleOutUnusedList {
	s.RuleEndTime = &v
	return s
}

func (s *DataRefundChangeRuleMapValueChangeRuleOutUnusedList) SetCanChange(v bool) *DataRefundChangeRuleMapValueChangeRuleOutUnusedList {
	s.CanChange = &v
	return s
}

func (s *DataRefundChangeRuleMapValueChangeRuleOutUnusedList) SetChangeFee(v float64) *DataRefundChangeRuleMapValueChangeRuleOutUnusedList {
	s.ChangeFee = &v
	return s
}

type DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue struct {
	// checked baggage quantity
	//
	// example:
	//
	// 1
	BaggageAmount *int32 `json:"baggage_amount,omitempty" xml:"baggage_amount,omitempty"`
	// checked baggage weight
	//
	// example:
	//
	// 10
	BaggageWeight *int32 `json:"baggage_weight,omitempty" xml:"baggage_weight,omitempty"`
	// checked baggage weight unit
	//
	// example:
	//
	// KG
	BaggageWeightUnit *string `json:"baggage_weight_unit,omitempty" xml:"baggage_weight_unit,omitempty"`
	// Whether the weight is for all baggages
	//
	// example:
	//
	// true
	IsAllWeight *bool `json:"is_all_weight,omitempty" xml:"is_all_weight,omitempty"`
	// carry-on baggage quantity
	//
	// example:
	//
	// 1
	CarryOnAmount *int32 `json:"carry_on_amount,omitempty" xml:"carry_on_amount,omitempty"`
	// carry-on baggage weight
	//
	// example:
	//
	// 5
	CarryOnWeight *int32 `json:"carry_on_weight,omitempty" xml:"carry_on_weight,omitempty"`
	// carry-on baggage weight unit
	//
	// example:
	//
	// KG
	CarryOnWeightUnit *string `json:"carry_on_weight_unit,omitempty" xml:"carry_on_weight_unit,omitempty"`
	// Whether the weight is for all baggages
	//
	// example:
	//
	// true
	IsAllCarryOnWeight *bool `json:"is_all_carry_on_weight,omitempty" xml:"is_all_carry_on_weight,omitempty"`
}

func (s DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) String() string {
	return tea.Prettify(s)
}

func (s DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) GoString() string {
	return s.String()
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetBaggageAmount(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.BaggageAmount = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetBaggageWeight(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.BaggageWeight = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetBaggageWeightUnit(v string) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.BaggageWeightUnit = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetIsAllWeight(v bool) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.IsAllWeight = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarryOnAmount(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarryOnAmount = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarryOnWeight(v int32) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarryOnWeight = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetCarryOnWeightUnit(v string) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.CarryOnWeightUnit = &v
	return s
}

func (s *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) SetIsAllCarryOnWeight(v bool) *DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	s.IsAllCarryOnWeight = &v
	return s
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
	return tea.Prettify(s)
}

func (s DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRuleAllUnusedList) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueRefundRulePartUnusedList) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleInUnusedList) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValueChangeRuleOutUnusedList) GoString() string {
	return s.String()
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

type AccountFlowListHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access_token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s AccountFlowListHeaders) String() string {
	return tea.Prettify(s)
}

func (s AccountFlowListHeaders) GoString() string {
	return s.String()
}

func (s *AccountFlowListHeaders) SetCommonHeaders(v map[string]*string) *AccountFlowListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AccountFlowListHeaders) SetXAcsAirticketAccessToken(v string) *AccountFlowListHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *AccountFlowListHeaders) SetXAcsAirticketLanguage(v string) *AccountFlowListHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type AccountFlowListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0
	DayNum *int32 `json:"day_num,omitempty" xml:"day_num,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1677427200000
	UtcBeginTime *int64 `json:"utc_begin_time,omitempty" xml:"utc_begin_time,omitempty"`
}

func (s AccountFlowListRequest) String() string {
	return tea.Prettify(s)
}

func (s AccountFlowListRequest) GoString() string {
	return s.String()
}

func (s *AccountFlowListRequest) SetDayNum(v int32) *AccountFlowListRequest {
	s.DayNum = &v
	return s
}

func (s *AccountFlowListRequest) SetPageIndex(v int32) *AccountFlowListRequest {
	s.PageIndex = &v
	return s
}

func (s *AccountFlowListRequest) SetPageSize(v int32) *AccountFlowListRequest {
	s.PageSize = &v
	return s
}

func (s *AccountFlowListRequest) SetUtcBeginTime(v int64) *AccountFlowListRequest {
	s.UtcBeginTime = &v
	return s
}

type AccountFlowListResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AccountFlowListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AccountFlowListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AccountFlowListResponseBody) GoString() string {
	return s.String()
}

func (s *AccountFlowListResponseBody) SetRequestId(v string) *AccountFlowListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AccountFlowListResponseBody) SetData(v *AccountFlowListResponseBodyData) *AccountFlowListResponseBody {
	s.Data = v
	return s
}

func (s *AccountFlowListResponseBody) SetErrorCode(v string) *AccountFlowListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AccountFlowListResponseBody) SetErrorData(v interface{}) *AccountFlowListResponseBody {
	s.ErrorData = v
	return s
}

func (s *AccountFlowListResponseBody) SetErrorMsg(v string) *AccountFlowListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *AccountFlowListResponseBody) SetStatus(v int32) *AccountFlowListResponseBody {
	s.Status = &v
	return s
}

func (s *AccountFlowListResponseBody) SetSuccess(v bool) *AccountFlowListResponseBody {
	s.Success = &v
	return s
}

type AccountFlowListResponseBodyData struct {
	List       []*AccountFlowListResponseBodyDataList     `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	Pagination *AccountFlowListResponseBodyDataPagination `json:"pagination,omitempty" xml:"pagination,omitempty" type:"Struct"`
}

func (s AccountFlowListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AccountFlowListResponseBodyData) GoString() string {
	return s.String()
}

func (s *AccountFlowListResponseBodyData) SetList(v []*AccountFlowListResponseBodyDataList) *AccountFlowListResponseBodyData {
	s.List = v
	return s
}

func (s *AccountFlowListResponseBodyData) SetPagination(v *AccountFlowListResponseBodyDataPagination) *AccountFlowListResponseBodyData {
	s.Pagination = v
	return s
}

type AccountFlowListResponseBodyDataList struct {
	// example:
	//
	// 1000
	AfterAvailableAmount *float64 `json:"after_available_amount,omitempty" xml:"after_available_amount,omitempty"`
	// example:
	//
	// 1950.5
	BeforeAvailableAmount *float64 `json:"before_available_amount,omitempty" xml:"before_available_amount,omitempty"`
	// example:
	//
	// 49880***971
	ChangeOrderNum *int64 `json:"change_order_num,omitempty" xml:"change_order_num,omitempty"`
	// example:
	//
	// 1627239841225842666
	FlowId *int64 `json:"flow_id,omitempty" xml:"flow_id,omitempty"`
	// example:
	//
	// 1676799185000
	GmtCreate *int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 1676966530000
	GmtModified *int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// example:
	//
	// 950.5
	OpAmount *float64 `json:"op_amount,omitempty" xml:"op_amount,omitempty"`
	// example:
	//
	// 2
	OpType *int32 `json:"op_type,omitempty" xml:"op_type,omitempty"`
	// example:
	//
	// 4988430***971
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// example:
	//
	// 1
	OrderType *int32 `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// example:
	//
	// 4988430***971
	OutOrderNum *string `json:"out_order_num,omitempty" xml:"out_order_num,omitempty"`
	// example:
	//
	// 48430***971
	RefundOrderNum *int64 `json:"refund_order_num,omitempty" xml:"refund_order_num,omitempty"`
}

func (s AccountFlowListResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s AccountFlowListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *AccountFlowListResponseBodyDataList) SetAfterAvailableAmount(v float64) *AccountFlowListResponseBodyDataList {
	s.AfterAvailableAmount = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetBeforeAvailableAmount(v float64) *AccountFlowListResponseBodyDataList {
	s.BeforeAvailableAmount = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetChangeOrderNum(v int64) *AccountFlowListResponseBodyDataList {
	s.ChangeOrderNum = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetFlowId(v int64) *AccountFlowListResponseBodyDataList {
	s.FlowId = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetGmtCreate(v int64) *AccountFlowListResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetGmtModified(v int64) *AccountFlowListResponseBodyDataList {
	s.GmtModified = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetOpAmount(v float64) *AccountFlowListResponseBodyDataList {
	s.OpAmount = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetOpType(v int32) *AccountFlowListResponseBodyDataList {
	s.OpType = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetOrderNum(v int64) *AccountFlowListResponseBodyDataList {
	s.OrderNum = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetOrderType(v int32) *AccountFlowListResponseBodyDataList {
	s.OrderType = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetOutOrderNum(v string) *AccountFlowListResponseBodyDataList {
	s.OutOrderNum = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetRefundOrderNum(v int64) *AccountFlowListResponseBodyDataList {
	s.RefundOrderNum = &v
	return s
}

type AccountFlowListResponseBodyDataPagination struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

func (s AccountFlowListResponseBodyDataPagination) String() string {
	return tea.Prettify(s)
}

func (s AccountFlowListResponseBodyDataPagination) GoString() string {
	return s.String()
}

func (s *AccountFlowListResponseBodyDataPagination) SetCurrentPage(v int32) *AccountFlowListResponseBodyDataPagination {
	s.CurrentPage = &v
	return s
}

func (s *AccountFlowListResponseBodyDataPagination) SetPageSize(v int32) *AccountFlowListResponseBodyDataPagination {
	s.PageSize = &v
	return s
}

func (s *AccountFlowListResponseBodyDataPagination) SetTotalCount(v int32) *AccountFlowListResponseBodyDataPagination {
	s.TotalCount = &v
	return s
}

func (s *AccountFlowListResponseBodyDataPagination) SetTotalPage(v int32) *AccountFlowListResponseBodyDataPagination {
	s.TotalPage = &v
	return s
}

type AccountFlowListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AccountFlowListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AccountFlowListResponse) String() string {
	return tea.Prettify(s)
}

func (s AccountFlowListResponse) GoString() string {
	return s.String()
}

func (s *AccountFlowListResponse) SetHeaders(v map[string]*string) *AccountFlowListResponse {
	s.Headers = v
	return s
}

func (s *AccountFlowListResponse) SetStatusCode(v int32) *AccountFlowListResponse {
	s.StatusCode = &v
	return s
}

func (s *AccountFlowListResponse) SetBody(v *AccountFlowListResponseBody) *AccountFlowListResponse {
	s.Body = v
	return s
}

type AncillarySuggestHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access_token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// Language Code(refer to ISO_639)
	//
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s AncillarySuggestHeaders) String() string {
	return tea.Prettify(s)
}

func (s AncillarySuggestHeaders) GoString() string {
	return s.String()
}

func (s *AncillarySuggestHeaders) SetCommonHeaders(v map[string]*string) *AncillarySuggestHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AncillarySuggestHeaders) SetXAcsAirticketAccessToken(v string) *AncillarySuggestHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *AncillarySuggestHeaders) SetXAcsAirticketLanguage(v string) *AncillarySuggestHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type AncillarySuggestRequest struct {
	// solution_id returned by enrich
	//
	// This parameter is required.
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s AncillarySuggestRequest) String() string {
	return tea.Prettify(s)
}

func (s AncillarySuggestRequest) GoString() string {
	return s.String()
}

func (s *AncillarySuggestRequest) SetSolutionId(v string) *AncillarySuggestRequest {
	s.SolutionId = &v
	return s
}

type AncillarySuggestResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Properly processed return data
	Data *AncillarySuggestResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// Data carried in error handling
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// Error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// http request successful, status value is always 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true represents success, false represents failure
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AncillarySuggestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AncillarySuggestResponseBody) GoString() string {
	return s.String()
}

func (s *AncillarySuggestResponseBody) SetRequestId(v string) *AncillarySuggestResponseBody {
	s.RequestId = &v
	return s
}

func (s *AncillarySuggestResponseBody) SetData(v *AncillarySuggestResponseBodyData) *AncillarySuggestResponseBody {
	s.Data = v
	return s
}

func (s *AncillarySuggestResponseBody) SetErrorCode(v string) *AncillarySuggestResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AncillarySuggestResponseBody) SetErrorData(v interface{}) *AncillarySuggestResponseBody {
	s.ErrorData = v
	return s
}

func (s *AncillarySuggestResponseBody) SetErrorMsg(v string) *AncillarySuggestResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *AncillarySuggestResponseBody) SetStatus(v int32) *AncillarySuggestResponseBody {
	s.Status = &v
	return s
}

func (s *AncillarySuggestResponseBody) SetSuccess(v bool) *AncillarySuggestResponseBody {
	s.Success = &v
	return s
}

type AncillarySuggestResponseBodyData struct {
	// ancillary detail list
	SegAncillaryMapList []*AncillarySuggestResponseBodyDataSegAncillaryMapList `json:"seg_ancillary_map_list,omitempty" xml:"seg_ancillary_map_list,omitempty" type:"Repeated"`
	// solution_id, equals to solution_id in request
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s AncillarySuggestResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AncillarySuggestResponseBodyData) GoString() string {
	return s.String()
}

func (s *AncillarySuggestResponseBodyData) SetSegAncillaryMapList(v []*AncillarySuggestResponseBodyDataSegAncillaryMapList) *AncillarySuggestResponseBodyData {
	s.SegAncillaryMapList = v
	return s
}

func (s *AncillarySuggestResponseBodyData) SetSolutionId(v string) *AncillarySuggestResponseBodyData {
	s.SolutionId = &v
	return s
}

type AncillarySuggestResponseBodyDataSegAncillaryMapList struct {
	// Ancillary product
	Ancillary *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary `json:"ancillary,omitempty" xml:"ancillary,omitempty" type:"Struct"`
	// Segment ID list, these segments share the same ancillary
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s AncillarySuggestResponseBodyDataSegAncillaryMapList) String() string {
	return tea.Prettify(s)
}

func (s AncillarySuggestResponseBodyDataSegAncillaryMapList) GoString() string {
	return s.String()
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapList) SetAncillary(v *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary) *AncillarySuggestResponseBodyDataSegAncillaryMapList {
	s.Ancillary = v
	return s
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapList) SetSegmentIdList(v []*string) *AncillarySuggestResponseBodyDataSegAncillaryMapList {
	s.SegmentIdList = v
	return s
}

type AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary struct {
	// Ancillary product ID
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU
	AncillaryId *string `json:"ancillary_id,omitempty" xml:"ancillary_id,omitempty"`
	// Ancillary product type. currently supports 4: paid luggage
	//
	// example:
	//
	// 4
	AncillaryType *int32 `json:"ancillary_type,omitempty" xml:"ancillary_type,omitempty"`
	// Baggage details
	BaggageAncillary *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary `json:"baggage_ancillary,omitempty" xml:"baggage_ancillary,omitempty" type:"Struct"`
}

func (s AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary) String() string {
	return tea.Prettify(s)
}

func (s AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary) GoString() string {
	return s.String()
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary) SetAncillaryId(v string) *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary {
	s.AncillaryId = &v
	return s
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary) SetAncillaryType(v int32) *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary {
	s.AncillaryType = &v
	return s
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary) SetBaggageAncillary(v *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillary {
	s.BaggageAncillary = v
	return s
}

type AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary struct {
	// baggage quantity, values such as: 3, 2, 1, 0, -2.     -2 indicates weight-based
	//
	// example:
	//
	// 0
	BaggageAmount *int32 `json:"baggage_amount,omitempty" xml:"baggage_amount,omitempty"`
	// Baggage weight, 0-50. When isAllWeight=true, it represents the total weight of all baggages.
	//
	// example:
	//
	// 0
	BaggageWeight *int32 `json:"baggage_weight,omitempty" xml:"baggage_weight,omitempty"`
	// Unit of baggage weight
	//
	// example:
	//
	// KG
	BaggageWeightUnit *string `json:"baggage_weight_unit,omitempty" xml:"baggage_weight_unit,omitempty"`
	// Whether the weight is for all baggages
	//
	// example:
	//
	// true
	IsAllWeight *bool `json:"is_all_weight,omitempty" xml:"is_all_weight,omitempty"`
	// Total price
	//
	// example:
	//
	// 10.0
	Price *float64 `json:"price,omitempty" xml:"price,omitempty"`
}

func (s AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) String() string {
	return tea.Prettify(s)
}

func (s AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) GoString() string {
	return s.String()
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) SetBaggageAmount(v int32) *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary {
	s.BaggageAmount = &v
	return s
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) SetBaggageWeight(v int32) *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary {
	s.BaggageWeight = &v
	return s
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) SetBaggageWeightUnit(v string) *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary {
	s.BaggageWeightUnit = &v
	return s
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) SetIsAllWeight(v bool) *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary {
	s.IsAllWeight = &v
	return s
}

func (s *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary) SetPrice(v float64) *AncillarySuggestResponseBodyDataSegAncillaryMapListAncillaryBaggageAncillary {
	s.Price = &v
	return s
}

type AncillarySuggestResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AncillarySuggestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AncillarySuggestResponse) String() string {
	return tea.Prettify(s)
}

func (s AncillarySuggestResponse) GoString() string {
	return s.String()
}

func (s *AncillarySuggestResponse) SetHeaders(v map[string]*string) *AncillarySuggestResponse {
	s.Headers = v
	return s
}

func (s *AncillarySuggestResponse) SetStatusCode(v int32) *AncillarySuggestResponse {
	s.StatusCode = &v
	return s
}

func (s *AncillarySuggestResponse) SetBody(v *AncillarySuggestResponseBody) *AncillarySuggestResponse {
	s.Body = v
	return s
}

type BookHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// language code(refer to ISO_639), defaults to the buyer\\"s account configuration
	//
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s BookHeaders) String() string {
	return tea.Prettify(s)
}

func (s BookHeaders) GoString() string {
	return s.String()
}

func (s *BookHeaders) SetCommonHeaders(v map[string]*string) *BookHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BookHeaders) SetXAcsAirticketAccessToken(v string) *BookHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *BookHeaders) SetXAcsAirticketLanguage(v string) *BookHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type BookRequest struct {
	// contact information
	//
	// This parameter is required.
	Contact *BookRequestContact `json:"contact,omitempty" xml:"contact,omitempty" type:"Struct"`
	// external order number(buyer customization)
	//
	// This parameter is required.
	//
	// example:
	//
	// x091-2023-0220-j-0001
	OutOrderNum *string `json:"out_order_num,omitempty" xml:"out_order_num,omitempty"`
	// passenger-ancillary purchase relationship
	PassengerAncillaryPurchaseMapList []*BookRequestPassengerAncillaryPurchaseMapList `json:"passenger_ancillary_purchase_map_list,omitempty" xml:"passenger_ancillary_purchase_map_list,omitempty" type:"Repeated"`
	// passenger list
	//
	// This parameter is required.
	PassengerList []*BookRequestPassengerList `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
	// solution_id returned by Enrich
	//
	// This parameter is required.
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s BookRequest) String() string {
	return tea.Prettify(s)
}

func (s BookRequest) GoString() string {
	return s.String()
}

func (s *BookRequest) SetContact(v *BookRequestContact) *BookRequest {
	s.Contact = v
	return s
}

func (s *BookRequest) SetOutOrderNum(v string) *BookRequest {
	s.OutOrderNum = &v
	return s
}

func (s *BookRequest) SetPassengerAncillaryPurchaseMapList(v []*BookRequestPassengerAncillaryPurchaseMapList) *BookRequest {
	s.PassengerAncillaryPurchaseMapList = v
	return s
}

func (s *BookRequest) SetPassengerList(v []*BookRequestPassengerList) *BookRequest {
	s.PassengerList = v
	return s
}

func (s *BookRequest) SetSolutionId(v string) *BookRequest {
	s.SolutionId = &v
	return s
}

type BookRequestContact struct {
	// email address
	//
	// example:
	//
	// gao******@gmail.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// first name
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// last name
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// country code
	//
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// mobile phone number
	//
	// example:
	//
	// 183******96
	MobilePhoneNum *string `json:"mobile_phone_num,omitempty" xml:"mobile_phone_num,omitempty"`
}

func (s BookRequestContact) String() string {
	return tea.Prettify(s)
}

func (s BookRequestContact) GoString() string {
	return s.String()
}

func (s *BookRequestContact) SetEmail(v string) *BookRequestContact {
	s.Email = &v
	return s
}

func (s *BookRequestContact) SetFirstName(v string) *BookRequestContact {
	s.FirstName = &v
	return s
}

func (s *BookRequestContact) SetLastName(v string) *BookRequestContact {
	s.LastName = &v
	return s
}

func (s *BookRequestContact) SetMobileCountryCode(v string) *BookRequestContact {
	s.MobileCountryCode = &v
	return s
}

func (s *BookRequestContact) SetMobilePhoneNum(v string) *BookRequestContact {
	s.MobilePhoneNum = &v
	return s
}

type BookRequestPassengerAncillaryPurchaseMapList struct {
	// ancillary information
	BookAncillaryReqItem *BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem `json:"book_ancillary_req_item,omitempty" xml:"book_ancillary_req_item,omitempty" type:"Struct"`
	// passenger list for unified ancillary purchases
	PassengerList []*BookRequestPassengerAncillaryPurchaseMapListPassengerList `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
}

func (s BookRequestPassengerAncillaryPurchaseMapList) String() string {
	return tea.Prettify(s)
}

func (s BookRequestPassengerAncillaryPurchaseMapList) GoString() string {
	return s.String()
}

func (s *BookRequestPassengerAncillaryPurchaseMapList) SetBookAncillaryReqItem(v *BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem) *BookRequestPassengerAncillaryPurchaseMapList {
	s.BookAncillaryReqItem = v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapList) SetPassengerList(v []*BookRequestPassengerAncillaryPurchaseMapListPassengerList) *BookRequestPassengerAncillaryPurchaseMapList {
	s.PassengerList = v
	return s
}

type BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem struct {
	// ancillary product ID, returned by AncillarySuggest.
	//
	// example:
	//
	// MDY2NTAxLCJleHAiOjE2NxNzM3MDEsIm5ix
	AncillaryId *string `json:"ancillary_id,omitempty" xml:"ancillary_id,omitempty"`
	// type of ancillary product, only support "4"(4 means paid baggage) currently.
	//
	// example:
	//
	// 4
	AncillaryType *int32 `json:"ancillary_type,omitempty" xml:"ancillary_type,omitempty"`
}

func (s BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem) String() string {
	return tea.Prettify(s)
}

func (s BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem) GoString() string {
	return s.String()
}

func (s *BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem) SetAncillaryId(v string) *BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem {
	s.AncillaryId = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem) SetAncillaryType(v int32) *BookRequestPassengerAncillaryPurchaseMapListBookAncillaryReqItem {
	s.AncillaryType = &v
	return s
}

type BookRequestPassengerAncillaryPurchaseMapListPassengerList struct {
	// date of birth (yyyyMMdd)
	//
	// example:
	//
	// 20020320
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// travel document
	Credential *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential `json:"credential,omitempty" xml:"credential,omitempty" type:"Struct"`
	// first name
	//
	// This parameter is required.
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// gender 0: male; 1: female
	//
	// example:
	//
	// 1
	Gender *int32 `json:"gender,omitempty" xml:"gender,omitempty"`
	// last name
	//
	// This parameter is required.
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// country code for mobile phone number
	//
	// This parameter is required.
	//
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// mobile phone number
	//
	// This parameter is required.
	//
	// example:
	//
	// 182******92
	MobilePhoneNumber *string `json:"mobile_phone_number,omitempty" xml:"mobile_phone_number,omitempty"`
	// nationality
	//
	// example:
	//
	// CN
	Nationality *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// passenger type 0: adult; 1: child; 8: Infant
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s BookRequestPassengerAncillaryPurchaseMapListPassengerList) String() string {
	return tea.Prettify(s)
}

func (s BookRequestPassengerAncillaryPurchaseMapListPassengerList) GoString() string {
	return s.String()
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) SetBirthday(v string) *BookRequestPassengerAncillaryPurchaseMapListPassengerList {
	s.Birthday = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) SetCredential(v *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential) *BookRequestPassengerAncillaryPurchaseMapListPassengerList {
	s.Credential = v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) SetFirstName(v string) *BookRequestPassengerAncillaryPurchaseMapListPassengerList {
	s.FirstName = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) SetGender(v int32) *BookRequestPassengerAncillaryPurchaseMapListPassengerList {
	s.Gender = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) SetLastName(v string) *BookRequestPassengerAncillaryPurchaseMapListPassengerList {
	s.LastName = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) SetMobileCountryCode(v string) *BookRequestPassengerAncillaryPurchaseMapListPassengerList {
	s.MobileCountryCode = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) SetMobilePhoneNumber(v string) *BookRequestPassengerAncillaryPurchaseMapListPassengerList {
	s.MobilePhoneNumber = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) SetNationality(v string) *BookRequestPassengerAncillaryPurchaseMapListPassengerList {
	s.Nationality = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerList) SetType(v int32) *BookRequestPassengerAncillaryPurchaseMapListPassengerList {
	s.Type = &v
	return s
}

type BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential struct {
	// place of issue, two-letter code
	//
	// example:
	//
	// CN
	CertIssuePlace *string `json:"cert_issue_place,omitempty" xml:"cert_issue_place,omitempty"`
	// travel document number
	//
	// example:
	//
	// E1***5673
	CredentialNum *string `json:"credential_num,omitempty" xml:"credential_num,omitempty"`
	// travel document type , only support "1"(1 means passport) currently
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CredentialType *int32 `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	// expiration date
	//
	// example:
	//
	// 20290102
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
}

func (s BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential) String() string {
	return tea.Prettify(s)
}

func (s BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential) GoString() string {
	return s.String()
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential) SetCertIssuePlace(v string) *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential {
	s.CertIssuePlace = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential) SetCredentialNum(v string) *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential {
	s.CredentialNum = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential) SetCredentialType(v int32) *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential {
	s.CredentialType = &v
	return s
}

func (s *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential) SetExpireDate(v string) *BookRequestPassengerAncillaryPurchaseMapListPassengerListCredential {
	s.ExpireDate = &v
	return s
}

type BookRequestPassengerList struct {
	// date of birth (yyyyMMdd)
	//
	// example:
	//
	// 20200320
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// travel document
	Credential *BookRequestPassengerListCredential `json:"credential,omitempty" xml:"credential,omitempty" type:"Struct"`
	// first name
	//
	// This parameter is required.
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// gender 0: MALE; 1: FEMALE
	//
	// example:
	//
	// 0
	Gender *int32 `json:"gender,omitempty" xml:"gender,omitempty"`
	// last name
	//
	// This parameter is required.
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// mobile country code
	//
	// This parameter is required.
	//
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// mobile phone number
	//
	// This parameter is required.
	//
	// example:
	//
	// 183******95
	MobilePhoneNumber *string `json:"mobile_phone_number,omitempty" xml:"mobile_phone_number,omitempty"`
	// nationality (two-letter code)
	//
	// example:
	//
	// CN
	Nationality *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// passenger type 0: adult; 1: child; 8: infant
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s BookRequestPassengerList) String() string {
	return tea.Prettify(s)
}

func (s BookRequestPassengerList) GoString() string {
	return s.String()
}

func (s *BookRequestPassengerList) SetBirthday(v string) *BookRequestPassengerList {
	s.Birthday = &v
	return s
}

func (s *BookRequestPassengerList) SetCredential(v *BookRequestPassengerListCredential) *BookRequestPassengerList {
	s.Credential = v
	return s
}

func (s *BookRequestPassengerList) SetFirstName(v string) *BookRequestPassengerList {
	s.FirstName = &v
	return s
}

func (s *BookRequestPassengerList) SetGender(v int32) *BookRequestPassengerList {
	s.Gender = &v
	return s
}

func (s *BookRequestPassengerList) SetLastName(v string) *BookRequestPassengerList {
	s.LastName = &v
	return s
}

func (s *BookRequestPassengerList) SetMobileCountryCode(v string) *BookRequestPassengerList {
	s.MobileCountryCode = &v
	return s
}

func (s *BookRequestPassengerList) SetMobilePhoneNumber(v string) *BookRequestPassengerList {
	s.MobilePhoneNumber = &v
	return s
}

func (s *BookRequestPassengerList) SetNationality(v string) *BookRequestPassengerList {
	s.Nationality = &v
	return s
}

func (s *BookRequestPassengerList) SetType(v int32) *BookRequestPassengerList {
	s.Type = &v
	return s
}

type BookRequestPassengerListCredential struct {
	// place of issue, two-letter code
	//
	// example:
	//
	// CN
	CertIssuePlace *string `json:"cert_issue_place,omitempty" xml:"cert_issue_place,omitempty"`
	// travel document number
	//
	// example:
	//
	// E1***5674
	CredentialNum *string `json:"credential_num,omitempty" xml:"credential_num,omitempty"`
	// travel document type , only support "1"(1 means passport) currently.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CredentialType *int32 `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	// expiration date
	//
	// example:
	//
	// 20290101
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
}

func (s BookRequestPassengerListCredential) String() string {
	return tea.Prettify(s)
}

func (s BookRequestPassengerListCredential) GoString() string {
	return s.String()
}

func (s *BookRequestPassengerListCredential) SetCertIssuePlace(v string) *BookRequestPassengerListCredential {
	s.CertIssuePlace = &v
	return s
}

func (s *BookRequestPassengerListCredential) SetCredentialNum(v string) *BookRequestPassengerListCredential {
	s.CredentialNum = &v
	return s
}

func (s *BookRequestPassengerListCredential) SetCredentialType(v int32) *BookRequestPassengerListCredential {
	s.CredentialType = &v
	return s
}

func (s *BookRequestPassengerListCredential) SetExpireDate(v string) *BookRequestPassengerListCredential {
	s.ExpireDate = &v
	return s
}

type BookShrinkRequest struct {
	// contact information
	//
	// This parameter is required.
	ContactShrink *string `json:"contact,omitempty" xml:"contact,omitempty"`
	// external order number(buyer customization)
	//
	// This parameter is required.
	//
	// example:
	//
	// x091-2023-0220-j-0001
	OutOrderNum *string `json:"out_order_num,omitempty" xml:"out_order_num,omitempty"`
	// passenger-ancillary purchase relationship
	PassengerAncillaryPurchaseMapListShrink *string `json:"passenger_ancillary_purchase_map_list,omitempty" xml:"passenger_ancillary_purchase_map_list,omitempty"`
	// passenger list
	//
	// This parameter is required.
	PassengerListShrink *string `json:"passenger_list,omitempty" xml:"passenger_list,omitempty"`
	// solution_id returned by Enrich
	//
	// This parameter is required.
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s BookShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BookShrinkRequest) GoString() string {
	return s.String()
}

func (s *BookShrinkRequest) SetContactShrink(v string) *BookShrinkRequest {
	s.ContactShrink = &v
	return s
}

func (s *BookShrinkRequest) SetOutOrderNum(v string) *BookShrinkRequest {
	s.OutOrderNum = &v
	return s
}

func (s *BookShrinkRequest) SetPassengerAncillaryPurchaseMapListShrink(v string) *BookShrinkRequest {
	s.PassengerAncillaryPurchaseMapListShrink = &v
	return s
}

func (s *BookShrinkRequest) SetPassengerListShrink(v string) *BookShrinkRequest {
	s.PassengerListShrink = &v
	return s
}

func (s *BookShrinkRequest) SetSolutionId(v string) *BookShrinkRequest {
	s.SolutionId = &v
	return s
}

type BookResponseBody struct {
	// request ID
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *BookResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error data
	//
	// example:
	//
	// null
	ErrorData *BookResponseBodyErrorData `json:"error_data,omitempty" xml:"error_data,omitempty" type:"Struct"`
	// error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// http reqeust has been processed successfully，status code is 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true represents success, false represents failure
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BookResponseBody) GoString() string {
	return s.String()
}

func (s *BookResponseBody) SetRequestId(v string) *BookResponseBody {
	s.RequestId = &v
	return s
}

func (s *BookResponseBody) SetData(v *BookResponseBodyData) *BookResponseBody {
	s.Data = v
	return s
}

func (s *BookResponseBody) SetErrorCode(v string) *BookResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BookResponseBody) SetErrorData(v *BookResponseBodyErrorData) *BookResponseBody {
	s.ErrorData = v
	return s
}

func (s *BookResponseBody) SetErrorMsg(v string) *BookResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *BookResponseBody) SetStatus(v int32) *BookResponseBody {
	s.Status = &v
	return s
}

func (s *BookResponseBody) SetSuccess(v bool) *BookResponseBody {
	s.Success = &v
	return s
}

type BookResponseBodyData struct {
	// order information list
	OrderList []*BookResponseBodyDataOrderList `json:"order_list,omitempty" xml:"order_list,omitempty" type:"Repeated"`
}

func (s BookResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BookResponseBodyData) GoString() string {
	return s.String()
}

func (s *BookResponseBodyData) SetOrderList(v []*BookResponseBodyDataOrderList) *BookResponseBodyData {
	s.OrderList = v
	return s
}

type BookResponseBodyDataOrderList struct {
	// order number
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s BookResponseBodyDataOrderList) String() string {
	return tea.Prettify(s)
}

func (s BookResponseBodyDataOrderList) GoString() string {
	return s.String()
}

func (s *BookResponseBodyDataOrderList) SetOrderNum(v int64) *BookResponseBodyDataOrderList {
	s.OrderNum = &v
	return s
}

type BookResponseBodyErrorData struct {
	// order information list. When the same input parameters are used to repeat a Book, if the booking has already been successful, the order number will be returned.
	OrderList []*BookResponseBodyErrorDataOrderList `json:"order_list,omitempty" xml:"order_list,omitempty" type:"Repeated"`
}

func (s BookResponseBodyErrorData) String() string {
	return tea.Prettify(s)
}

func (s BookResponseBodyErrorData) GoString() string {
	return s.String()
}

func (s *BookResponseBodyErrorData) SetOrderList(v []*BookResponseBodyErrorDataOrderList) *BookResponseBodyErrorData {
	s.OrderList = v
	return s
}

type BookResponseBodyErrorDataOrderList struct {
	// order number
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s BookResponseBodyErrorDataOrderList) String() string {
	return tea.Prettify(s)
}

func (s BookResponseBodyErrorDataOrderList) GoString() string {
	return s.String()
}

func (s *BookResponseBodyErrorDataOrderList) SetOrderNum(v int64) *BookResponseBodyErrorDataOrderList {
	s.OrderNum = &v
	return s
}

type BookResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BookResponseBody  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BookResponse) String() string {
	return tea.Prettify(s)
}

func (s BookResponse) GoString() string {
	return s.String()
}

func (s *BookResponse) SetHeaders(v map[string]*string) *BookResponse {
	s.Headers = v
	return s
}

func (s *BookResponse) SetStatusCode(v int32) *BookResponse {
	s.StatusCode = &v
	return s
}

func (s *BookResponse) SetBody(v *BookResponseBody) *BookResponse {
	s.Body = v
	return s
}

type CancelHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// language code(refer to ISO_639), defaults to buyer account configuration
	//
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s CancelHeaders) String() string {
	return tea.Prettify(s)
}

func (s CancelHeaders) GoString() string {
	return s.String()
}

func (s *CancelHeaders) SetCommonHeaders(v map[string]*string) *CancelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CancelHeaders) SetXAcsAirticketAccessToken(v string) *CancelHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *CancelHeaders) SetXAcsAirticketLanguage(v string) *CancelHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type CancelRequest struct {
	// order number
	//
	// This parameter is required.
	//
	// example:
	//
	// 496***2617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s CancelRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelRequest) GoString() string {
	return s.String()
}

func (s *CancelRequest) SetOrderNum(v int64) *CancelRequest {
	s.OrderNum = &v
	return s
}

type CancelResponseBody struct {
	// request ID
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *CancelResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error data
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// http reqeust has been processed successfully，status code is 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true represents success, false represents failure
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CancelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelResponseBody) GoString() string {
	return s.String()
}

func (s *CancelResponseBody) SetRequestId(v string) *CancelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelResponseBody) SetData(v *CancelResponseBodyData) *CancelResponseBody {
	s.Data = v
	return s
}

func (s *CancelResponseBody) SetErrorCode(v string) *CancelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CancelResponseBody) SetErrorData(v interface{}) *CancelResponseBody {
	s.ErrorData = v
	return s
}

func (s *CancelResponseBody) SetErrorMsg(v string) *CancelResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CancelResponseBody) SetStatus(v int32) *CancelResponseBody {
	s.Status = &v
	return s
}

func (s *CancelResponseBody) SetSuccess(v bool) *CancelResponseBody {
	s.Success = &v
	return s
}

type CancelResponseBodyData struct {
	// order number
	//
	// example:
	//
	// 496***2617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s CancelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CancelResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelResponseBodyData) SetOrderNum(v int64) *CancelResponseBodyData {
	s.OrderNum = &v
	return s
}

type CancelResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelResponse) GoString() string {
	return s.String()
}

func (s *CancelResponse) SetHeaders(v map[string]*string) *CancelResponse {
	s.Headers = v
	return s
}

func (s *CancelResponse) SetStatusCode(v int32) *CancelResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelResponse) SetBody(v *CancelResponseBody) *CancelResponse {
	s.Body = v
	return s
}

type ChangeApplyHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access_token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s ChangeApplyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ChangeApplyHeaders) GoString() string {
	return s.String()
}

func (s *ChangeApplyHeaders) SetCommonHeaders(v map[string]*string) *ChangeApplyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChangeApplyHeaders) SetXAcsAirticketAccessToken(v string) *ChangeApplyHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *ChangeApplyHeaders) SetXAcsAirticketLanguage(v string) *ChangeApplyHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type ChangeApplyRequest struct {
	// This parameter is required.
	ChangePassengerList []*ChangeApplyRequestChangePassengerList `json:"change_passenger_list,omitempty" xml:"change_passenger_list,omitempty" type:"Repeated"`
	// This parameter is required.
	ChangedJourneys []*ChangeApplyRequestChangedJourneys `json:"changed_journeys,omitempty" xml:"changed_journeys,omitempty" type:"Repeated"`
	// This parameter is required.
	Contact *ChangeApplyRequestContact `json:"contact,omitempty" xml:"contact,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 4988430***950
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// example:
	//
	// remark desc
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ChangeApplyRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeApplyRequest) GoString() string {
	return s.String()
}

func (s *ChangeApplyRequest) SetChangePassengerList(v []*ChangeApplyRequestChangePassengerList) *ChangeApplyRequest {
	s.ChangePassengerList = v
	return s
}

func (s *ChangeApplyRequest) SetChangedJourneys(v []*ChangeApplyRequestChangedJourneys) *ChangeApplyRequest {
	s.ChangedJourneys = v
	return s
}

func (s *ChangeApplyRequest) SetContact(v *ChangeApplyRequestContact) *ChangeApplyRequest {
	s.Contact = v
	return s
}

func (s *ChangeApplyRequest) SetOrderNum(v int64) *ChangeApplyRequest {
	s.OrderNum = &v
	return s
}

func (s *ChangeApplyRequest) SetRemark(v string) *ChangeApplyRequest {
	s.Remark = &v
	return s
}

func (s *ChangeApplyRequest) SetType(v int32) *ChangeApplyRequest {
	s.Type = &v
	return s
}

type ChangeApplyRequestChangePassengerList struct {
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s ChangeApplyRequestChangePassengerList) String() string {
	return tea.Prettify(s)
}

func (s ChangeApplyRequestChangePassengerList) GoString() string {
	return s.String()
}

func (s *ChangeApplyRequestChangePassengerList) SetDocument(v string) *ChangeApplyRequestChangePassengerList {
	s.Document = &v
	return s
}

func (s *ChangeApplyRequestChangePassengerList) SetFirstName(v string) *ChangeApplyRequestChangePassengerList {
	s.FirstName = &v
	return s
}

func (s *ChangeApplyRequestChangePassengerList) SetLastName(v string) *ChangeApplyRequestChangePassengerList {
	s.LastName = &v
	return s
}

type ChangeApplyRequestChangedJourneys struct {
	SegmentList []*ChangeApplyRequestChangedJourneysSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
}

func (s ChangeApplyRequestChangedJourneys) String() string {
	return tea.Prettify(s)
}

func (s ChangeApplyRequestChangedJourneys) GoString() string {
	return s.String()
}

func (s *ChangeApplyRequestChangedJourneys) SetSegmentList(v []*ChangeApplyRequestChangedJourneysSegmentList) *ChangeApplyRequestChangedJourneys {
	s.SegmentList = v
	return s
}

type ChangeApplyRequestChangedJourneysSegmentList struct {
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// example:
	//
	// T1
	ArriveTerminal *string `json:"arrive_terminal,omitempty" xml:"arrive_terminal,omitempty"`
	// example:
	//
	// 1677232999000
	ArriveTime    *int64  `json:"arrive_time,omitempty" xml:"arrive_time,omitempty"`
	ArriveTimeStr *string `json:"arrive_time_str,omitempty" xml:"arrive_time_str,omitempty"`
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20230320
	DepartureDate *string `json:"departure_date,omitempty" xml:"departure_date,omitempty"`
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// example:
	//
	// 1677232998000
	DepartureTime    *int64  `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	DepartureTimeStr *string `json:"departure_time_str,omitempty" xml:"departure_time_str,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
}

func (s ChangeApplyRequestChangedJourneysSegmentList) String() string {
	return tea.Prettify(s)
}

func (s ChangeApplyRequestChangedJourneysSegmentList) GoString() string {
	return s.String()
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetArrivalAirport(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetArrivalCity(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetArriveTerminal(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.ArriveTerminal = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetArriveTime(v int64) *ChangeApplyRequestChangedJourneysSegmentList {
	s.ArriveTime = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetArriveTimeStr(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.ArriveTimeStr = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetCodeShare(v bool) *ChangeApplyRequestChangedJourneysSegmentList {
	s.CodeShare = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetDepartureAirport(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetDepartureCity(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetDepartureDate(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.DepartureDate = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetDepartureTerminal(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetDepartureTime(v int64) *ChangeApplyRequestChangedJourneysSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetDepartureTimeStr(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.DepartureTimeStr = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetMarketingFlightNo(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *ChangeApplyRequestChangedJourneysSegmentList) SetOperatingFlightNo(v string) *ChangeApplyRequestChangedJourneysSegmentList {
	s.OperatingFlightNo = &v
	return s
}

type ChangeApplyRequestContact struct {
	// example:
	//
	// gao******@gmail.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// example:
	//
	// 183*****92
	MobilePhoneNum *string `json:"mobile_phone_num,omitempty" xml:"mobile_phone_num,omitempty"`
}

func (s ChangeApplyRequestContact) String() string {
	return tea.Prettify(s)
}

func (s ChangeApplyRequestContact) GoString() string {
	return s.String()
}

func (s *ChangeApplyRequestContact) SetEmail(v string) *ChangeApplyRequestContact {
	s.Email = &v
	return s
}

func (s *ChangeApplyRequestContact) SetMobileCountryCode(v string) *ChangeApplyRequestContact {
	s.MobileCountryCode = &v
	return s
}

func (s *ChangeApplyRequestContact) SetMobilePhoneNum(v string) *ChangeApplyRequestContact {
	s.MobilePhoneNum = &v
	return s
}

type ChangeApplyShrinkRequest struct {
	// This parameter is required.
	ChangePassengerListShrink *string `json:"change_passenger_list,omitempty" xml:"change_passenger_list,omitempty"`
	// This parameter is required.
	ChangedJourneysShrink *string `json:"changed_journeys,omitempty" xml:"changed_journeys,omitempty"`
	// This parameter is required.
	ContactShrink *string `json:"contact,omitempty" xml:"contact,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4988430***950
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// example:
	//
	// remark desc
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ChangeApplyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeApplyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ChangeApplyShrinkRequest) SetChangePassengerListShrink(v string) *ChangeApplyShrinkRequest {
	s.ChangePassengerListShrink = &v
	return s
}

func (s *ChangeApplyShrinkRequest) SetChangedJourneysShrink(v string) *ChangeApplyShrinkRequest {
	s.ChangedJourneysShrink = &v
	return s
}

func (s *ChangeApplyShrinkRequest) SetContactShrink(v string) *ChangeApplyShrinkRequest {
	s.ContactShrink = &v
	return s
}

func (s *ChangeApplyShrinkRequest) SetOrderNum(v int64) *ChangeApplyShrinkRequest {
	s.OrderNum = &v
	return s
}

func (s *ChangeApplyShrinkRequest) SetRemark(v string) *ChangeApplyShrinkRequest {
	s.Remark = &v
	return s
}

func (s *ChangeApplyShrinkRequest) SetType(v int32) *ChangeApplyShrinkRequest {
	s.Type = &v
	return s
}

type ChangeApplyResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ChangeApplyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ChangeApplyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeApplyResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeApplyResponseBody) SetRequestId(v string) *ChangeApplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeApplyResponseBody) SetData(v *ChangeApplyResponseBodyData) *ChangeApplyResponseBody {
	s.Data = v
	return s
}

func (s *ChangeApplyResponseBody) SetErrorCode(v string) *ChangeApplyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ChangeApplyResponseBody) SetErrorData(v interface{}) *ChangeApplyResponseBody {
	s.ErrorData = v
	return s
}

func (s *ChangeApplyResponseBody) SetErrorMsg(v string) *ChangeApplyResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ChangeApplyResponseBody) SetStatus(v int32) *ChangeApplyResponseBody {
	s.Status = &v
	return s
}

func (s *ChangeApplyResponseBody) SetSuccess(v bool) *ChangeApplyResponseBody {
	s.Success = &v
	return s
}

type ChangeApplyResponseBodyData struct {
	ChangeOrders []*ChangeApplyResponseBodyDataChangeOrders `json:"change_orders,omitempty" xml:"change_orders,omitempty" type:"Repeated"`
	// example:
	//
	// 4988430***950
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s ChangeApplyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ChangeApplyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeApplyResponseBodyData) SetChangeOrders(v []*ChangeApplyResponseBodyDataChangeOrders) *ChangeApplyResponseBodyData {
	s.ChangeOrders = v
	return s
}

func (s *ChangeApplyResponseBodyData) SetOrderNum(v int64) *ChangeApplyResponseBodyData {
	s.OrderNum = &v
	return s
}

type ChangeApplyResponseBodyDataChangeOrders struct {
	// example:
	//
	// 49884*****950
	ChangeOrderNum *int64 `json:"change_order_num,omitempty" xml:"change_order_num,omitempty"`
	// example:
	//
	// 0
	ChangeOrderStatus *int32 `json:"change_order_status,omitempty" xml:"change_order_status,omitempty"`
	// example:
	//
	// desc reason
	FailReason *string                                              `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	Passengers []*ChangeApplyResponseBodyDataChangeOrdersPassengers `json:"passengers,omitempty" xml:"passengers,omitempty" type:"Repeated"`
}

func (s ChangeApplyResponseBodyDataChangeOrders) String() string {
	return tea.Prettify(s)
}

func (s ChangeApplyResponseBodyDataChangeOrders) GoString() string {
	return s.String()
}

func (s *ChangeApplyResponseBodyDataChangeOrders) SetChangeOrderNum(v int64) *ChangeApplyResponseBodyDataChangeOrders {
	s.ChangeOrderNum = &v
	return s
}

func (s *ChangeApplyResponseBodyDataChangeOrders) SetChangeOrderStatus(v int32) *ChangeApplyResponseBodyDataChangeOrders {
	s.ChangeOrderStatus = &v
	return s
}

func (s *ChangeApplyResponseBodyDataChangeOrders) SetFailReason(v string) *ChangeApplyResponseBodyDataChangeOrders {
	s.FailReason = &v
	return s
}

func (s *ChangeApplyResponseBodyDataChangeOrders) SetPassengers(v []*ChangeApplyResponseBodyDataChangeOrdersPassengers) *ChangeApplyResponseBodyDataChangeOrders {
	s.Passengers = v
	return s
}

type ChangeApplyResponseBodyDataChangeOrdersPassengers struct {
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s ChangeApplyResponseBodyDataChangeOrdersPassengers) String() string {
	return tea.Prettify(s)
}

func (s ChangeApplyResponseBodyDataChangeOrdersPassengers) GoString() string {
	return s.String()
}

func (s *ChangeApplyResponseBodyDataChangeOrdersPassengers) SetDocument(v string) *ChangeApplyResponseBodyDataChangeOrdersPassengers {
	s.Document = &v
	return s
}

func (s *ChangeApplyResponseBodyDataChangeOrdersPassengers) SetFirstName(v string) *ChangeApplyResponseBodyDataChangeOrdersPassengers {
	s.FirstName = &v
	return s
}

func (s *ChangeApplyResponseBodyDataChangeOrdersPassengers) SetLastName(v string) *ChangeApplyResponseBodyDataChangeOrdersPassengers {
	s.LastName = &v
	return s
}

type ChangeApplyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeApplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeApplyResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeApplyResponse) GoString() string {
	return s.String()
}

func (s *ChangeApplyResponse) SetHeaders(v map[string]*string) *ChangeApplyResponse {
	s.Headers = v
	return s
}

func (s *ChangeApplyResponse) SetStatusCode(v int32) *ChangeApplyResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeApplyResponse) SetBody(v *ChangeApplyResponseBody) *ChangeApplyResponse {
	s.Body = v
	return s
}

type ChangeCancelHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access_token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s ChangeCancelHeaders) String() string {
	return tea.Prettify(s)
}

func (s ChangeCancelHeaders) GoString() string {
	return s.String()
}

func (s *ChangeCancelHeaders) SetCommonHeaders(v map[string]*string) *ChangeCancelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChangeCancelHeaders) SetXAcsAirticketAccessToken(v string) *ChangeCancelHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *ChangeCancelHeaders) SetXAcsAirticketLanguage(v string) *ChangeCancelHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type ChangeCancelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4988430***950
	ChangeOrderNum *int64 `json:"change_order_num,omitempty" xml:"change_order_num,omitempty"`
}

func (s ChangeCancelRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeCancelRequest) GoString() string {
	return s.String()
}

func (s *ChangeCancelRequest) SetChangeOrderNum(v int64) *ChangeCancelRequest {
	s.ChangeOrderNum = &v
	return s
}

type ChangeCancelResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// null
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ChangeCancelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeCancelResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeCancelResponseBody) SetRequestId(v string) *ChangeCancelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeCancelResponseBody) SetData(v interface{}) *ChangeCancelResponseBody {
	s.Data = v
	return s
}

func (s *ChangeCancelResponseBody) SetErrorCode(v string) *ChangeCancelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ChangeCancelResponseBody) SetErrorData(v interface{}) *ChangeCancelResponseBody {
	s.ErrorData = v
	return s
}

func (s *ChangeCancelResponseBody) SetErrorMsg(v string) *ChangeCancelResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ChangeCancelResponseBody) SetStatus(v int32) *ChangeCancelResponseBody {
	s.Status = &v
	return s
}

func (s *ChangeCancelResponseBody) SetSuccess(v bool) *ChangeCancelResponseBody {
	s.Success = &v
	return s
}

type ChangeCancelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeCancelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeCancelResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeCancelResponse) GoString() string {
	return s.String()
}

func (s *ChangeCancelResponse) SetHeaders(v map[string]*string) *ChangeCancelResponse {
	s.Headers = v
	return s
}

func (s *ChangeCancelResponse) SetStatusCode(v int32) *ChangeCancelResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeCancelResponse) SetBody(v *ChangeCancelResponseBody) *ChangeCancelResponse {
	s.Body = v
	return s
}

type ChangeConfirmHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access_token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s ChangeConfirmHeaders) String() string {
	return tea.Prettify(s)
}

func (s ChangeConfirmHeaders) GoString() string {
	return s.String()
}

func (s *ChangeConfirmHeaders) SetCommonHeaders(v map[string]*string) *ChangeConfirmHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChangeConfirmHeaders) SetXAcsAirticketAccessToken(v string) *ChangeConfirmHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *ChangeConfirmHeaders) SetXAcsAirticketLanguage(v string) *ChangeConfirmHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type ChangeConfirmRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4988430***950
	ChangeOrderNum *int64 `json:"change_order_num,omitempty" xml:"change_order_num,omitempty"`
}

func (s ChangeConfirmRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeConfirmRequest) GoString() string {
	return s.String()
}

func (s *ChangeConfirmRequest) SetChangeOrderNum(v int64) *ChangeConfirmRequest {
	s.ChangeOrderNum = &v
	return s
}

type ChangeConfirmResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ChangeConfirmResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ChangeConfirmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeConfirmResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeConfirmResponseBody) SetRequestId(v string) *ChangeConfirmResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeConfirmResponseBody) SetData(v *ChangeConfirmResponseBodyData) *ChangeConfirmResponseBody {
	s.Data = v
	return s
}

func (s *ChangeConfirmResponseBody) SetErrorCode(v string) *ChangeConfirmResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ChangeConfirmResponseBody) SetErrorData(v interface{}) *ChangeConfirmResponseBody {
	s.ErrorData = v
	return s
}

func (s *ChangeConfirmResponseBody) SetErrorMsg(v string) *ChangeConfirmResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ChangeConfirmResponseBody) SetStatus(v int32) *ChangeConfirmResponseBody {
	s.Status = &v
	return s
}

func (s *ChangeConfirmResponseBody) SetSuccess(v bool) *ChangeConfirmResponseBody {
	s.Success = &v
	return s
}

type ChangeConfirmResponseBodyData struct {
	// example:
	//
	// 30
	PayAmount *float64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	// example:
	//
	// hkduendkd-2023-dj0
	TransactionNo *string `json:"transaction_no,omitempty" xml:"transaction_no,omitempty"`
}

func (s ChangeConfirmResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ChangeConfirmResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeConfirmResponseBodyData) SetPayAmount(v float64) *ChangeConfirmResponseBodyData {
	s.PayAmount = &v
	return s
}

func (s *ChangeConfirmResponseBodyData) SetTransactionNo(v string) *ChangeConfirmResponseBodyData {
	s.TransactionNo = &v
	return s
}

type ChangeConfirmResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeConfirmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeConfirmResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeConfirmResponse) GoString() string {
	return s.String()
}

func (s *ChangeConfirmResponse) SetHeaders(v map[string]*string) *ChangeConfirmResponse {
	s.Headers = v
	return s
}

func (s *ChangeConfirmResponse) SetStatusCode(v int32) *ChangeConfirmResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeConfirmResponse) SetBody(v *ChangeConfirmResponseBody) *ChangeConfirmResponse {
	s.Body = v
	return s
}

type ChangeDetailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access_token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s ChangeDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailHeaders) GoString() string {
	return s.String()
}

func (s *ChangeDetailHeaders) SetCommonHeaders(v map[string]*string) *ChangeDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChangeDetailHeaders) SetXAcsAirticketAccessToken(v string) *ChangeDetailHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *ChangeDetailHeaders) SetXAcsAirticketLanguage(v string) *ChangeDetailHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type ChangeDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4988430***950
	ChangeOrderNum *int64 `json:"change_order_num,omitempty" xml:"change_order_num,omitempty"`
}

func (s ChangeDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailRequest) GoString() string {
	return s.String()
}

func (s *ChangeDetailRequest) SetChangeOrderNum(v int64) *ChangeDetailRequest {
	s.ChangeOrderNum = &v
	return s
}

type ChangeDetailResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ChangeDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ChangeDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBody) SetRequestId(v string) *ChangeDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeDetailResponseBody) SetData(v *ChangeDetailResponseBodyData) *ChangeDetailResponseBody {
	s.Data = v
	return s
}

func (s *ChangeDetailResponseBody) SetErrorCode(v string) *ChangeDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ChangeDetailResponseBody) SetErrorData(v interface{}) *ChangeDetailResponseBody {
	s.ErrorData = v
	return s
}

func (s *ChangeDetailResponseBody) SetErrorMsg(v string) *ChangeDetailResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ChangeDetailResponseBody) SetStatus(v int32) *ChangeDetailResponseBody {
	s.Status = &v
	return s
}

func (s *ChangeDetailResponseBody) SetSuccess(v bool) *ChangeDetailResponseBody {
	s.Success = &v
	return s
}

type ChangeDetailResponseBodyData struct {
	ChangeFeeDetails []*ChangeDetailResponseBodyDataChangeFeeDetails `json:"change_fee_details,omitempty" xml:"change_fee_details,omitempty" type:"Repeated"`
	// example:
	//
	// 4988430***950
	ChangeOrderNum   *int64                                          `json:"change_order_num,omitempty" xml:"change_order_num,omitempty"`
	ChangePassengers []*ChangeDetailResponseBodyDataChangePassengers `json:"change_passengers,omitempty" xml:"change_passengers,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	ChangeReasonType *int32                                         `json:"change_reason_type,omitempty" xml:"change_reason_type,omitempty"`
	ChangedJourneys  []*ChangeDetailResponseBodyDataChangedJourneys `json:"changed_journeys,omitempty" xml:"changed_journeys,omitempty" type:"Repeated"`
	// example:
	//
	// reason desc
	CloseReason *string `json:"close_reason,omitempty" xml:"close_reason,omitempty"`
	// example:
	//
	// 1677415244000
	CloseUtcTime *int64                               `json:"close_utc_time,omitempty" xml:"close_utc_time,omitempty"`
	Contact      *ChangeDetailResponseBodyDataContact `json:"contact,omitempty" xml:"contact,omitempty" type:"Struct"`
	// example:
	//
	// 1677415276000
	CreateUtcTime *int64 `json:"create_utc_time,omitempty" xml:"create_utc_time,omitempty"`
	// example:
	//
	// 1677415278000
	LastConfirmUtcTime *int64                                      `json:"last_confirm_utc_time,omitempty" xml:"last_confirm_utc_time,omitempty"`
	LastJourneys       []*ChangeDetailResponseBodyDataLastJourneys `json:"last_journeys,omitempty" xml:"last_journeys,omitempty" type:"Repeated"`
	// example:
	//
	// 5988430***541
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// example:
	//
	// 2
	OrderStatus      *int32                                          `json:"order_status,omitempty" xml:"order_status,omitempty"`
	OriginalJourneys []*ChangeDetailResponseBodyDataOriginalJourneys `json:"original_journeys,omitempty" xml:"original_journeys,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PayStatus *int32 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// example:
	//
	// 1677415255000
	PaySuccessUtcTime *int64 `json:"pay_success_utc_time,omitempty" xml:"pay_success_utc_time,omitempty"`
	// example:
	//
	// 300
	TotalAmount *float64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// example:
	//
	// hkduendkd-2023-dj0
	TransactionNo *string `json:"transaction_no,omitempty" xml:"transaction_no,omitempty"`
}

func (s ChangeDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyData) SetChangeFeeDetails(v []*ChangeDetailResponseBodyDataChangeFeeDetails) *ChangeDetailResponseBodyData {
	s.ChangeFeeDetails = v
	return s
}

func (s *ChangeDetailResponseBodyData) SetChangeOrderNum(v int64) *ChangeDetailResponseBodyData {
	s.ChangeOrderNum = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetChangePassengers(v []*ChangeDetailResponseBodyDataChangePassengers) *ChangeDetailResponseBodyData {
	s.ChangePassengers = v
	return s
}

func (s *ChangeDetailResponseBodyData) SetChangeReasonType(v int32) *ChangeDetailResponseBodyData {
	s.ChangeReasonType = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetChangedJourneys(v []*ChangeDetailResponseBodyDataChangedJourneys) *ChangeDetailResponseBodyData {
	s.ChangedJourneys = v
	return s
}

func (s *ChangeDetailResponseBodyData) SetCloseReason(v string) *ChangeDetailResponseBodyData {
	s.CloseReason = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetCloseUtcTime(v int64) *ChangeDetailResponseBodyData {
	s.CloseUtcTime = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetContact(v *ChangeDetailResponseBodyDataContact) *ChangeDetailResponseBodyData {
	s.Contact = v
	return s
}

func (s *ChangeDetailResponseBodyData) SetCreateUtcTime(v int64) *ChangeDetailResponseBodyData {
	s.CreateUtcTime = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetLastConfirmUtcTime(v int64) *ChangeDetailResponseBodyData {
	s.LastConfirmUtcTime = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetLastJourneys(v []*ChangeDetailResponseBodyDataLastJourneys) *ChangeDetailResponseBodyData {
	s.LastJourneys = v
	return s
}

func (s *ChangeDetailResponseBodyData) SetOrderNum(v int64) *ChangeDetailResponseBodyData {
	s.OrderNum = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetOrderStatus(v int32) *ChangeDetailResponseBodyData {
	s.OrderStatus = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetOriginalJourneys(v []*ChangeDetailResponseBodyDataOriginalJourneys) *ChangeDetailResponseBodyData {
	s.OriginalJourneys = v
	return s
}

func (s *ChangeDetailResponseBodyData) SetPayStatus(v int32) *ChangeDetailResponseBodyData {
	s.PayStatus = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetPaySuccessUtcTime(v int64) *ChangeDetailResponseBodyData {
	s.PaySuccessUtcTime = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetTotalAmount(v float64) *ChangeDetailResponseBodyData {
	s.TotalAmount = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetTransactionNo(v string) *ChangeDetailResponseBodyData {
	s.TransactionNo = &v
	return s
}

type ChangeDetailResponseBodyDataChangeFeeDetails struct {
	ChangeFee *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee `json:"change_fee,omitempty" xml:"change_fee,omitempty" type:"Struct"`
	Passenger *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger `json:"passenger,omitempty" xml:"passenger,omitempty" type:"Struct"`
}

func (s ChangeDetailResponseBodyDataChangeFeeDetails) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailResponseBodyDataChangeFeeDetails) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetails) SetChangeFee(v *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee) *ChangeDetailResponseBodyDataChangeFeeDetails {
	s.ChangeFee = v
	return s
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetails) SetPassenger(v *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger) *ChangeDetailResponseBodyDataChangeFeeDetails {
	s.Passenger = v
	return s
}

type ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee struct {
	// example:
	//
	// 50
	ServiceFee *float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// example:
	//
	// 20
	TaxFee *float64 `json:"tax_fee,omitempty" xml:"tax_fee,omitempty"`
	// example:
	//
	// 30
	UpgradeFee *float64 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
}

func (s ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee) SetServiceFee(v float64) *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee {
	s.ServiceFee = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee) SetTaxFee(v float64) *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee {
	s.TaxFee = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee) SetUpgradeFee(v float64) *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee {
	s.UpgradeFee = &v
	return s
}

type ChangeDetailResponseBodyDataChangeFeeDetailsPassenger struct {
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s ChangeDetailResponseBodyDataChangeFeeDetailsPassenger) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailResponseBodyDataChangeFeeDetailsPassenger) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger) SetDocument(v string) *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger {
	s.Document = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger) SetFirstName(v string) *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger {
	s.FirstName = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger) SetLastName(v string) *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger {
	s.LastName = &v
	return s
}

type ChangeDetailResponseBodyDataChangePassengers struct {
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s ChangeDetailResponseBodyDataChangePassengers) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailResponseBodyDataChangePassengers) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataChangePassengers) SetDocument(v string) *ChangeDetailResponseBodyDataChangePassengers {
	s.Document = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangePassengers) SetFirstName(v string) *ChangeDetailResponseBodyDataChangePassengers {
	s.FirstName = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangePassengers) SetLastName(v string) *ChangeDetailResponseBodyDataChangePassengers {
	s.LastName = &v
	return s
}

type ChangeDetailResponseBodyDataChangedJourneys struct {
	SegmentList []*ChangeDetailResponseBodyDataChangedJourneysSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s ChangeDetailResponseBodyDataChangedJourneys) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailResponseBodyDataChangedJourneys) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataChangedJourneys) SetSegmentList(v []*ChangeDetailResponseBodyDataChangedJourneysSegmentList) *ChangeDetailResponseBodyDataChangedJourneys {
	s.SegmentList = v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneys) SetTransferCount(v int32) *ChangeDetailResponseBodyDataChangedJourneys {
	s.TransferCount = &v
	return s
}

type ChangeDetailResponseBodyDataChangedJourneysSegmentList struct {
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// example:
	//
	// 2023-03-10 10:40:00
	ArrivalTime *string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// example:
	//
	// 7
	Availability *string `json:"availability,omitempty" xml:"availability,omitempty"`
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// example:
	//
	// 2023-03-10 07:55:00
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// example:
	//
	// 32Q
	EquipType *string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// example:
	//
	// 1295
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// example:
	//
	// SEL,HKG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// example:
	//
	// 0
	StopQuantity *int32 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s ChangeDetailResponseBodyDataChangedJourneysSegmentList) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailResponseBodyDataChangedJourneysSegmentList) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetArrivalAirport(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetArrivalCity(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetArrivalTerminal(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetArrivalTime(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetAvailability(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.Availability = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetCabin(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.Cabin = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetCabinClass(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.CabinClass = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetCodeShare(v bool) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.CodeShare = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetDepartureAirport(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetDepartureCity(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetDepartureTerminal(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetDepartureTime(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetEquipType(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.EquipType = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetFlightDuration(v int32) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetMarketingAirline(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetMarketingFlightNo(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetMarketingFlightNoInt(v int32) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetOperatingAirline(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetOperatingFlightNo(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetSegmentId(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.SegmentId = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetStopCityList(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.StopCityList = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetStopQuantity(v int32) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.StopQuantity = &v
	return s
}

type ChangeDetailResponseBodyDataContact struct {
	// example:
	//
	// gao******@gmail.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// example:
	//
	// 183*****92
	MobilePhoneNum *string `json:"mobile_phone_num,omitempty" xml:"mobile_phone_num,omitempty"`
}

func (s ChangeDetailResponseBodyDataContact) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailResponseBodyDataContact) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataContact) SetEmail(v string) *ChangeDetailResponseBodyDataContact {
	s.Email = &v
	return s
}

func (s *ChangeDetailResponseBodyDataContact) SetMobileCountryCode(v string) *ChangeDetailResponseBodyDataContact {
	s.MobileCountryCode = &v
	return s
}

func (s *ChangeDetailResponseBodyDataContact) SetMobilePhoneNum(v string) *ChangeDetailResponseBodyDataContact {
	s.MobilePhoneNum = &v
	return s
}

type ChangeDetailResponseBodyDataLastJourneys struct {
	SegmentList []*ChangeDetailResponseBodyDataLastJourneysSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s ChangeDetailResponseBodyDataLastJourneys) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailResponseBodyDataLastJourneys) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataLastJourneys) SetSegmentList(v []*ChangeDetailResponseBodyDataLastJourneysSegmentList) *ChangeDetailResponseBodyDataLastJourneys {
	s.SegmentList = v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneys) SetTransferCount(v int32) *ChangeDetailResponseBodyDataLastJourneys {
	s.TransferCount = &v
	return s
}

type ChangeDetailResponseBodyDataLastJourneysSegmentList struct {
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// example:
	//
	// 2023-03-10 10:40:00
	ArrivalTime *string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// example:
	//
	// 7
	Availability *string `json:"availability,omitempty" xml:"availability,omitempty"`
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// example:
	//
	// 2023-03-10 07:55:00
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// example:
	//
	// 32Q
	EquipType *string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// example:
	//
	// 1295
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// example:
	//
	// SEL,HKG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// example:
	//
	// 0
	StopQuantity *int32 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s ChangeDetailResponseBodyDataLastJourneysSegmentList) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailResponseBodyDataLastJourneysSegmentList) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetArrivalAirport(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetArrivalCity(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetArrivalTerminal(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetArrivalTime(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetAvailability(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.Availability = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetCabin(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.Cabin = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetCabinClass(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.CabinClass = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetCodeShare(v bool) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.CodeShare = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetDepartureAirport(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetDepartureCity(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetDepartureTerminal(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetDepartureTime(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetEquipType(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.EquipType = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetFlightDuration(v int32) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetMarketingAirline(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetMarketingFlightNo(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetMarketingFlightNoInt(v int32) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetOperatingAirline(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetOperatingFlightNo(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetSegmentId(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.SegmentId = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetStopCityList(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.StopCityList = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetStopQuantity(v int32) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.StopQuantity = &v
	return s
}

type ChangeDetailResponseBodyDataOriginalJourneys struct {
	SegmentList []*ChangeDetailResponseBodyDataOriginalJourneysSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s ChangeDetailResponseBodyDataOriginalJourneys) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailResponseBodyDataOriginalJourneys) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataOriginalJourneys) SetSegmentList(v []*ChangeDetailResponseBodyDataOriginalJourneysSegmentList) *ChangeDetailResponseBodyDataOriginalJourneys {
	s.SegmentList = v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneys) SetTransferCount(v int32) *ChangeDetailResponseBodyDataOriginalJourneys {
	s.TransferCount = &v
	return s
}

type ChangeDetailResponseBodyDataOriginalJourneysSegmentList struct {
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// example:
	//
	// 2023-03-10 10:40:00
	ArrivalTime *string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// example:
	//
	// 7
	Availability *string `json:"availability,omitempty" xml:"availability,omitempty"`
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// example:
	//
	// 2023-03-10 07:55:00
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// example:
	//
	// 32Q
	EquipType *string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// example:
	//
	// 1295
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// example:
	//
	// SEL,HKG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// example:
	//
	// 0
	StopQuantity *int32 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s ChangeDetailResponseBodyDataOriginalJourneysSegmentList) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetArrivalAirport(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetArrivalCity(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetArrivalTerminal(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetArrivalTime(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetAvailability(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.Availability = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetCabin(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.Cabin = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetCabinClass(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.CabinClass = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetCodeShare(v bool) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.CodeShare = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetDepartureAirport(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetDepartureCity(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetDepartureTerminal(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetDepartureTime(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetEquipType(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.EquipType = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetFlightDuration(v int32) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetMarketingAirline(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetMarketingFlightNo(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetMarketingFlightNoInt(v int32) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetOperatingAirline(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetOperatingFlightNo(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetSegmentId(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.SegmentId = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetStopCityList(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.StopCityList = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetStopQuantity(v int32) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.StopQuantity = &v
	return s
}

type ChangeDetailResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailResponse) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponse) SetHeaders(v map[string]*string) *ChangeDetailResponse {
	s.Headers = v
	return s
}

func (s *ChangeDetailResponse) SetStatusCode(v int32) *ChangeDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeDetailResponse) SetBody(v *ChangeDetailResponseBody) *ChangeDetailResponse {
	s.Body = v
	return s
}

type ChangeDetailListOfBuyerHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access_token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s ChangeDetailListOfBuyerHeaders) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfBuyerHeaders) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfBuyerHeaders) SetCommonHeaders(v map[string]*string) *ChangeDetailListOfBuyerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChangeDetailListOfBuyerHeaders) SetXAcsAirticketAccessToken(v string) *ChangeDetailListOfBuyerHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *ChangeDetailListOfBuyerHeaders) SetXAcsAirticketLanguage(v string) *ChangeDetailListOfBuyerHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type ChangeDetailListOfBuyerRequest struct {
	// example:
	//
	// 1
	PageIndex *int32 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 1677415276000
	UtcCreateBegin *int64 `json:"utc_create_begin,omitempty" xml:"utc_create_begin,omitempty"`
	// example:
	//
	// 1677415279000
	UtcCreateEnd *int64 `json:"utc_create_end,omitempty" xml:"utc_create_end,omitempty"`
}

func (s ChangeDetailListOfBuyerRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfBuyerRequest) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfBuyerRequest) SetPageIndex(v int32) *ChangeDetailListOfBuyerRequest {
	s.PageIndex = &v
	return s
}

func (s *ChangeDetailListOfBuyerRequest) SetPageSize(v int32) *ChangeDetailListOfBuyerRequest {
	s.PageSize = &v
	return s
}

func (s *ChangeDetailListOfBuyerRequest) SetUtcCreateBegin(v int64) *ChangeDetailListOfBuyerRequest {
	s.UtcCreateBegin = &v
	return s
}

func (s *ChangeDetailListOfBuyerRequest) SetUtcCreateEnd(v int64) *ChangeDetailListOfBuyerRequest {
	s.UtcCreateEnd = &v
	return s
}

type ChangeDetailListOfBuyerResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ChangeDetailListOfBuyerResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ChangeDetailListOfBuyerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfBuyerResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfBuyerResponseBody) SetRequestId(v string) *ChangeDetailListOfBuyerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBody) SetData(v *ChangeDetailListOfBuyerResponseBodyData) *ChangeDetailListOfBuyerResponseBody {
	s.Data = v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBody) SetErrorCode(v string) *ChangeDetailListOfBuyerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBody) SetErrorData(v interface{}) *ChangeDetailListOfBuyerResponseBody {
	s.ErrorData = v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBody) SetErrorMsg(v string) *ChangeDetailListOfBuyerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBody) SetStatus(v int32) *ChangeDetailListOfBuyerResponseBody {
	s.Status = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBody) SetSuccess(v bool) *ChangeDetailListOfBuyerResponseBody {
	s.Success = &v
	return s
}

type ChangeDetailListOfBuyerResponseBodyData struct {
	List       []*ChangeDetailListOfBuyerResponseBodyDataList     `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	Pagination *ChangeDetailListOfBuyerResponseBodyDataPagination `json:"pagination,omitempty" xml:"pagination,omitempty" type:"Struct"`
}

func (s ChangeDetailListOfBuyerResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfBuyerResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfBuyerResponseBodyData) SetList(v []*ChangeDetailListOfBuyerResponseBodyDataList) *ChangeDetailListOfBuyerResponseBodyData {
	s.List = v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyData) SetPagination(v *ChangeDetailListOfBuyerResponseBodyDataPagination) *ChangeDetailListOfBuyerResponseBodyData {
	s.Pagination = v
	return s
}

type ChangeDetailListOfBuyerResponseBodyDataList struct {
	// example:
	//
	// 4988430***950
	ChangeOrderNum *int64 `json:"change_order_num,omitempty" xml:"change_order_num,omitempty"`
	// example:
	//
	// 4988430***971
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// example:
	//
	// 2
	OrderStatus *int32 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// example:
	//
	// 2
	PayStatus *int32 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// example:
	//
	// hkduendkd-2023-dj0
	TransactionNo *string `json:"transaction_no,omitempty" xml:"transaction_no,omitempty"`
	// example:
	//
	// 1677415274000
	UtcCreateTime *int64 `json:"utc_create_time,omitempty" xml:"utc_create_time,omitempty"`
}

func (s ChangeDetailListOfBuyerResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfBuyerResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfBuyerResponseBodyDataList) SetChangeOrderNum(v int64) *ChangeDetailListOfBuyerResponseBodyDataList {
	s.ChangeOrderNum = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyDataList) SetOrderNum(v int64) *ChangeDetailListOfBuyerResponseBodyDataList {
	s.OrderNum = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyDataList) SetOrderStatus(v int32) *ChangeDetailListOfBuyerResponseBodyDataList {
	s.OrderStatus = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyDataList) SetPayStatus(v int32) *ChangeDetailListOfBuyerResponseBodyDataList {
	s.PayStatus = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyDataList) SetTransactionNo(v string) *ChangeDetailListOfBuyerResponseBodyDataList {
	s.TransactionNo = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyDataList) SetUtcCreateTime(v int64) *ChangeDetailListOfBuyerResponseBodyDataList {
	s.UtcCreateTime = &v
	return s
}

type ChangeDetailListOfBuyerResponseBodyDataPagination struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

func (s ChangeDetailListOfBuyerResponseBodyDataPagination) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfBuyerResponseBodyDataPagination) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfBuyerResponseBodyDataPagination) SetCurrentPage(v int32) *ChangeDetailListOfBuyerResponseBodyDataPagination {
	s.CurrentPage = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyDataPagination) SetPageSize(v int32) *ChangeDetailListOfBuyerResponseBodyDataPagination {
	s.PageSize = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyDataPagination) SetTotalCount(v int32) *ChangeDetailListOfBuyerResponseBodyDataPagination {
	s.TotalCount = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyDataPagination) SetTotalPage(v int32) *ChangeDetailListOfBuyerResponseBodyDataPagination {
	s.TotalPage = &v
	return s
}

type ChangeDetailListOfBuyerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeDetailListOfBuyerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeDetailListOfBuyerResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfBuyerResponse) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfBuyerResponse) SetHeaders(v map[string]*string) *ChangeDetailListOfBuyerResponse {
	s.Headers = v
	return s
}

func (s *ChangeDetailListOfBuyerResponse) SetStatusCode(v int32) *ChangeDetailListOfBuyerResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponse) SetBody(v *ChangeDetailListOfBuyerResponseBody) *ChangeDetailListOfBuyerResponse {
	s.Body = v
	return s
}

type ChangeDetailListOfOrderNumHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access_token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s ChangeDetailListOfOrderNumHeaders) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfOrderNumHeaders) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumHeaders) SetCommonHeaders(v map[string]*string) *ChangeDetailListOfOrderNumHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChangeDetailListOfOrderNumHeaders) SetXAcsAirticketAccessToken(v string) *ChangeDetailListOfOrderNumHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *ChangeDetailListOfOrderNumHeaders) SetXAcsAirticketLanguage(v string) *ChangeDetailListOfOrderNumHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type ChangeDetailListOfOrderNumRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4988430***700
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

func (s ChangeDetailListOfOrderNumRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfOrderNumRequest) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumRequest) SetOrderNum(v int64) *ChangeDetailListOfOrderNumRequest {
	s.OrderNum = &v
	return s
}

func (s *ChangeDetailListOfOrderNumRequest) SetPageIndex(v int32) *ChangeDetailListOfOrderNumRequest {
	s.PageIndex = &v
	return s
}

func (s *ChangeDetailListOfOrderNumRequest) SetPageSize(v int32) *ChangeDetailListOfOrderNumRequest {
	s.PageSize = &v
	return s
}

type ChangeDetailListOfOrderNumResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ChangeDetailListOfOrderNumResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBody) SetRequestId(v string) *ChangeDetailListOfOrderNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBody) SetData(v *ChangeDetailListOfOrderNumResponseBodyData) *ChangeDetailListOfOrderNumResponseBody {
	s.Data = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBody) SetErrorCode(v string) *ChangeDetailListOfOrderNumResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBody) SetErrorData(v interface{}) *ChangeDetailListOfOrderNumResponseBody {
	s.ErrorData = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBody) SetErrorMsg(v string) *ChangeDetailListOfOrderNumResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBody) SetStatus(v int32) *ChangeDetailListOfOrderNumResponseBody {
	s.Status = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBody) SetSuccess(v bool) *ChangeDetailListOfOrderNumResponseBody {
	s.Success = &v
	return s
}

type ChangeDetailListOfOrderNumResponseBodyData struct {
	List       []*ChangeDetailListOfOrderNumResponseBodyDataList     `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	Pagination *ChangeDetailListOfOrderNumResponseBodyDataPagination `json:"pagination,omitempty" xml:"pagination,omitempty" type:"Struct"`
}

func (s ChangeDetailListOfOrderNumResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyData) SetList(v []*ChangeDetailListOfOrderNumResponseBodyDataList) *ChangeDetailListOfOrderNumResponseBodyData {
	s.List = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyData) SetPagination(v *ChangeDetailListOfOrderNumResponseBodyDataPagination) *ChangeDetailListOfOrderNumResponseBodyData {
	s.Pagination = v
	return s
}

type ChangeDetailListOfOrderNumResponseBodyDataList struct {
	ChangeFeeDetails []*ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails `json:"change_fee_details,omitempty" xml:"change_fee_details,omitempty" type:"Repeated"`
	// example:
	//
	// 4988430***950
	ChangeOrderNum   *int64                                                            `json:"change_order_num,omitempty" xml:"change_order_num,omitempty"`
	ChangePassengers []*ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers `json:"change_passengers,omitempty" xml:"change_passengers,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	ChangeReasonType *int32                                                           `json:"change_reason_type,omitempty" xml:"change_reason_type,omitempty"`
	ChangedJourneys  []*ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys `json:"changed_journeys,omitempty" xml:"changed_journeys,omitempty" type:"Repeated"`
	// example:
	//
	// reason desc
	CloseReason *string `json:"close_reason,omitempty" xml:"close_reason,omitempty"`
	// example:
	//
	// 1677415244000
	CloseUtcTime *int64                                                 `json:"close_utc_time,omitempty" xml:"close_utc_time,omitempty"`
	Contact      *ChangeDetailListOfOrderNumResponseBodyDataListContact `json:"contact,omitempty" xml:"contact,omitempty" type:"Struct"`
	// example:
	//
	// 1677415276000
	CreateUtcTime *int64 `json:"create_utc_time,omitempty" xml:"create_utc_time,omitempty"`
	// example:
	//
	// 1677415278000
	LastConfirmUtcTime *int64                                                        `json:"last_confirm_utc_time,omitempty" xml:"last_confirm_utc_time,omitempty"`
	LastJourneys       []*ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys `json:"last_journeys,omitempty" xml:"last_journeys,omitempty" type:"Repeated"`
	// example:
	//
	// 5988430***541
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// example:
	//
	// 2
	OrderStatus      *int32                                                            `json:"order_status,omitempty" xml:"order_status,omitempty"`
	OriginalJourneys []*ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys `json:"original_journeys,omitempty" xml:"original_journeys,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PayStatus *int32 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// example:
	//
	// 1677415255000
	PaySuccessUtcTime *int64 `json:"pay_success_utc_time,omitempty" xml:"pay_success_utc_time,omitempty"`
	// example:
	//
	// 300
	TotalAmount *float64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// example:
	//
	// hkduendkd-2023-dj0
	TransactionNo *string `json:"transaction_no,omitempty" xml:"transaction_no,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetChangeFeeDetails(v []*ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.ChangeFeeDetails = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetChangeOrderNum(v int64) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.ChangeOrderNum = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetChangePassengers(v []*ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.ChangePassengers = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetChangeReasonType(v int32) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.ChangeReasonType = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetChangedJourneys(v []*ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.ChangedJourneys = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetCloseReason(v string) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.CloseReason = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetCloseUtcTime(v int64) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.CloseUtcTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetContact(v *ChangeDetailListOfOrderNumResponseBodyDataListContact) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.Contact = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetCreateUtcTime(v int64) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.CreateUtcTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetLastConfirmUtcTime(v int64) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.LastConfirmUtcTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetLastJourneys(v []*ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.LastJourneys = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetOrderNum(v int64) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.OrderNum = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetOrderStatus(v int32) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.OrderStatus = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetOriginalJourneys(v []*ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.OriginalJourneys = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetPayStatus(v int32) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.PayStatus = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetPaySuccessUtcTime(v int64) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.PaySuccessUtcTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetTotalAmount(v float64) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.TotalAmount = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetTransactionNo(v string) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.TransactionNo = &v
	return s
}

type ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails struct {
	ChangeFee *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee `json:"change_fee,omitempty" xml:"change_fee,omitempty" type:"Struct"`
	Passenger *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger `json:"passenger,omitempty" xml:"passenger,omitempty" type:"Struct"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails) SetChangeFee(v *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee) *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails {
	s.ChangeFee = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails) SetPassenger(v *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger) *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails {
	s.Passenger = v
	return s
}

type ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee struct {
	// example:
	//
	// 50
	ServiceFee *float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// example:
	//
	// 20
	TaxFee *float64 `json:"tax_fee,omitempty" xml:"tax_fee,omitempty"`
	// example:
	//
	// 30
	UpgradeFee *float64 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee) SetServiceFee(v float64) *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee {
	s.ServiceFee = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee) SetTaxFee(v float64) *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee {
	s.TaxFee = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee) SetUpgradeFee(v float64) *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee {
	s.UpgradeFee = &v
	return s
}

type ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger struct {
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger) SetDocument(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger {
	s.Document = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger) SetFirstName(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger {
	s.FirstName = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger) SetLastName(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger {
	s.LastName = &v
	return s
}

type ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers struct {
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers) SetDocument(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers {
	s.Document = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers) SetFirstName(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers {
	s.FirstName = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers) SetLastName(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers {
	s.LastName = &v
	return s
}

type ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys struct {
	SegmentList []*ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys) SetSegmentList(v []*ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys {
	s.SegmentList = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys) SetTransferCount(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys {
	s.TransferCount = &v
	return s
}

type ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList struct {
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// example:
	//
	// 2023-03-10 10:40:00
	ArrivalTime *string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// example:
	//
	// 7
	Availability *string `json:"availability,omitempty" xml:"availability,omitempty"`
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// example:
	//
	// 2023-03-10 07:55:00
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// example:
	//
	// 32Q
	EquipType *string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// example:
	//
	// 1295
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// example:
	//
	// SEL,HKG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// example:
	//
	// 0
	StopQuantity *int32 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetArrivalAirport(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetArrivalCity(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetArrivalTerminal(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetArrivalTime(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetAvailability(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.Availability = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetCabin(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.Cabin = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetCabinClass(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.CabinClass = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetCodeShare(v bool) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.CodeShare = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetDepartureAirport(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetDepartureCity(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetDepartureTerminal(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetDepartureTime(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetEquipType(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.EquipType = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetFlightDuration(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetMarketingAirline(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetMarketingFlightNo(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetMarketingFlightNoInt(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetOperatingAirline(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetOperatingFlightNo(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetSegmentId(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.SegmentId = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetStopCityList(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.StopCityList = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetStopQuantity(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.StopQuantity = &v
	return s
}

type ChangeDetailListOfOrderNumResponseBodyDataListContact struct {
	// example:
	//
	// gao******@gmail.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// example:
	//
	// 183*****92
	MobilePhoneNum *string `json:"mobile_phone_num,omitempty" xml:"mobile_phone_num,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListContact) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListContact) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListContact) SetEmail(v string) *ChangeDetailListOfOrderNumResponseBodyDataListContact {
	s.Email = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListContact) SetMobileCountryCode(v string) *ChangeDetailListOfOrderNumResponseBodyDataListContact {
	s.MobileCountryCode = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListContact) SetMobilePhoneNum(v string) *ChangeDetailListOfOrderNumResponseBodyDataListContact {
	s.MobilePhoneNum = &v
	return s
}

type ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys struct {
	SegmentList []*ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys) SetSegmentList(v []*ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys {
	s.SegmentList = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys) SetTransferCount(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys {
	s.TransferCount = &v
	return s
}

type ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList struct {
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// example:
	//
	// 2023-03-10 10:40:00
	ArrivalTime *string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// example:
	//
	// 7
	Availability *string `json:"availability,omitempty" xml:"availability,omitempty"`
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// example:
	//
	// 2023-03-10 07:55:00
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// example:
	//
	// 32Q
	EquipType *string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// example:
	//
	// 1295
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// example:
	//
	// SEL,HKG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// example:
	//
	// 0
	StopQuantity *int32 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetArrivalAirport(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetArrivalCity(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetArrivalTerminal(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetArrivalTime(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetAvailability(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.Availability = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetCabin(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.Cabin = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetCabinClass(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.CabinClass = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetCodeShare(v bool) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.CodeShare = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetDepartureAirport(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetDepartureCity(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetDepartureTerminal(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetDepartureTime(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetEquipType(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.EquipType = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetFlightDuration(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetMarketingAirline(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetMarketingFlightNo(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetMarketingFlightNoInt(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetOperatingAirline(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetOperatingFlightNo(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetSegmentId(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.SegmentId = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetStopCityList(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.StopCityList = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetStopQuantity(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.StopQuantity = &v
	return s
}

type ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys struct {
	SegmentList []*ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys) SetSegmentList(v []*ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys {
	s.SegmentList = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys) SetTransferCount(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys {
	s.TransferCount = &v
	return s
}

type ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList struct {
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// example:
	//
	// 2023-03-10 10:40:00
	ArrivalTime *string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// example:
	//
	// 7
	Availability *string `json:"availability,omitempty" xml:"availability,omitempty"`
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// example:
	//
	// 2023-03-10 07:55:00
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// example:
	//
	// 32Q
	EquipType *string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// example:
	//
	// 1295
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// example:
	//
	// SEL,HKG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// example:
	//
	// 0
	StopQuantity *int32 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetArrivalAirport(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetArrivalCity(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetArrivalTerminal(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetArrivalTime(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetAvailability(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.Availability = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetCabin(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.Cabin = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetCabinClass(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.CabinClass = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetCodeShare(v bool) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.CodeShare = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetDepartureAirport(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetDepartureCity(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetDepartureTerminal(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetDepartureTime(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetEquipType(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.EquipType = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetFlightDuration(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetMarketingAirline(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetMarketingFlightNo(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetMarketingFlightNoInt(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetOperatingAirline(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetOperatingFlightNo(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetSegmentId(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.SegmentId = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetStopCityList(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.StopCityList = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetStopQuantity(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.StopQuantity = &v
	return s
}

type ChangeDetailListOfOrderNumResponseBodyDataPagination struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataPagination) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataPagination) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataPagination) SetCurrentPage(v int32) *ChangeDetailListOfOrderNumResponseBodyDataPagination {
	s.CurrentPage = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataPagination) SetPageSize(v int32) *ChangeDetailListOfOrderNumResponseBodyDataPagination {
	s.PageSize = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataPagination) SetTotalCount(v int32) *ChangeDetailListOfOrderNumResponseBodyDataPagination {
	s.TotalCount = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataPagination) SetTotalPage(v int32) *ChangeDetailListOfOrderNumResponseBodyDataPagination {
	s.TotalPage = &v
	return s
}

type ChangeDetailListOfOrderNumResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeDetailListOfOrderNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponse) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponse) SetHeaders(v map[string]*string) *ChangeDetailListOfOrderNumResponse {
	s.Headers = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponse) SetStatusCode(v int32) *ChangeDetailListOfOrderNumResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponse) SetBody(v *ChangeDetailListOfOrderNumResponseBody) *ChangeDetailListOfOrderNumResponse {
	s.Body = v
	return s
}

type CollectFlightLowestPriceHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access_token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s CollectFlightLowestPriceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollectFlightLowestPriceHeaders) GoString() string {
	return s.String()
}

func (s *CollectFlightLowestPriceHeaders) SetCommonHeaders(v map[string]*string) *CollectFlightLowestPriceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollectFlightLowestPriceHeaders) SetXAcsAirticketAccessToken(v string) *CollectFlightLowestPriceHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *CollectFlightLowestPriceHeaders) SetXAcsAirticketLanguage(v string) *CollectFlightLowestPriceHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type CollectFlightLowestPriceRequest struct {
	// This parameter is required.
	LowestPriceFlightInfoList []*CollectFlightLowestPriceRequestLowestPriceFlightInfoList `json:"lowest_price_flight_info_list,omitempty" xml:"lowest_price_flight_info_list,omitempty" type:"Repeated"`
}

func (s CollectFlightLowestPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s CollectFlightLowestPriceRequest) GoString() string {
	return s.String()
}

func (s *CollectFlightLowestPriceRequest) SetLowestPriceFlightInfoList(v []*CollectFlightLowestPriceRequestLowestPriceFlightInfoList) *CollectFlightLowestPriceRequest {
	s.LowestPriceFlightInfoList = v
	return s
}

type CollectFlightLowestPriceRequestLowestPriceFlightInfoList struct {
	// This parameter is required.
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-11-11
	DepartureDate *string `json:"departure_date,omitempty" xml:"departure_date,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CA123,CA456
	DepartureFlightNumber *string `json:"departure_flight_number,omitempty" xml:"departure_flight_number,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100.0
	MarketTotalPrice *float64 `json:"market_total_price,omitempty" xml:"market_total_price,omitempty"`
	// example:
	//
	// 123456789dacd
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// example:
	//
	// 2024-11-11
	ReturnDate *string `json:"return_date,omitempty" xml:"return_date,omitempty"`
	// example:
	//
	// CA123,CA456
	ReturnFlightNumber *string `json:"return_flight_number,omitempty" xml:"return_flight_number,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100.1
	SuezTotalPrice *float64 `json:"suez_total_price,omitempty" xml:"suez_total_price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TripType *int32 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

func (s CollectFlightLowestPriceRequestLowestPriceFlightInfoList) String() string {
	return tea.Prettify(s)
}

func (s CollectFlightLowestPriceRequestLowestPriceFlightInfoList) GoString() string {
	return s.String()
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetArrivalCity(v string) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.ArrivalCity = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetDepartureCity(v string) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.DepartureCity = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetDepartureDate(v string) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.DepartureDate = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetDepartureFlightNumber(v string) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.DepartureFlightNumber = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetMarketTotalPrice(v float64) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.MarketTotalPrice = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetRequestId(v string) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.RequestId = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetReturnDate(v string) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.ReturnDate = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetReturnFlightNumber(v string) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.ReturnFlightNumber = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetSolutionId(v string) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.SolutionId = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetSuezTotalPrice(v float64) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.SuezTotalPrice = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetTripType(v int32) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.TripType = &v
	return s
}

type CollectFlightLowestPriceShrinkRequest struct {
	// This parameter is required.
	LowestPriceFlightInfoListShrink *string `json:"lowest_price_flight_info_list,omitempty" xml:"lowest_price_flight_info_list,omitempty"`
}

func (s CollectFlightLowestPriceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CollectFlightLowestPriceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CollectFlightLowestPriceShrinkRequest) SetLowestPriceFlightInfoListShrink(v string) *CollectFlightLowestPriceShrinkRequest {
	s.LowestPriceFlightInfoListShrink = &v
	return s
}

type CollectFlightLowestPriceResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// null
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CollectFlightLowestPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollectFlightLowestPriceResponseBody) GoString() string {
	return s.String()
}

func (s *CollectFlightLowestPriceResponseBody) SetRequestId(v string) *CollectFlightLowestPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CollectFlightLowestPriceResponseBody) SetData(v map[string]interface{}) *CollectFlightLowestPriceResponseBody {
	s.Data = v
	return s
}

func (s *CollectFlightLowestPriceResponseBody) SetErrorCode(v string) *CollectFlightLowestPriceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CollectFlightLowestPriceResponseBody) SetErrorData(v interface{}) *CollectFlightLowestPriceResponseBody {
	s.ErrorData = v
	return s
}

func (s *CollectFlightLowestPriceResponseBody) SetErrorMsg(v string) *CollectFlightLowestPriceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CollectFlightLowestPriceResponseBody) SetStatus(v int32) *CollectFlightLowestPriceResponseBody {
	s.Status = &v
	return s
}

func (s *CollectFlightLowestPriceResponseBody) SetSuccess(v bool) *CollectFlightLowestPriceResponseBody {
	s.Success = &v
	return s
}

type CollectFlightLowestPriceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CollectFlightLowestPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CollectFlightLowestPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s CollectFlightLowestPriceResponse) GoString() string {
	return s.String()
}

func (s *CollectFlightLowestPriceResponse) SetHeaders(v map[string]*string) *CollectFlightLowestPriceResponse {
	s.Headers = v
	return s
}

func (s *CollectFlightLowestPriceResponse) SetStatusCode(v int32) *CollectFlightLowestPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *CollectFlightLowestPriceResponse) SetBody(v *CollectFlightLowestPriceResponseBody) *CollectFlightLowestPriceResponse {
	s.Body = v
	return s
}

type EnrichHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// language code(refer to ISO_639), defaults to the buyer\\"s account configuration
	//
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s EnrichHeaders) String() string {
	return tea.Prettify(s)
}

func (s EnrichHeaders) GoString() string {
	return s.String()
}

func (s *EnrichHeaders) SetCommonHeaders(v map[string]*string) *EnrichHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EnrichHeaders) SetXAcsAirticketAccessToken(v string) *EnrichHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *EnrichHeaders) SetXAcsAirticketLanguage(v string) *EnrichHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type EnrichRequest struct {
	// adult passenger amount 1-9
	//
	// example:
	//
	// 1
	Adults *int32 `json:"adults,omitempty" xml:"adults,omitempty"`
	// cabin class
	//
	// 1. **ALL_CABIN*	- : all cabin class
	//
	// 2. **Y*	- : economy class
	//
	// 3. **FC*	- : first class and business class
	//
	// 4. **S*	- : premium economy class
	//
	// 5. **YS*	- : economy class and premium economy class
	//
	// 6. **YSC*	- : economy class, premium economy class and business class
	//
	// example:
	//
	// ALL_CABIN
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// child passenger amount 0-9
	//
	// example:
	//
	// 1
	Children *int32 `json:"children,omitempty" xml:"children,omitempty"`
	// infant passenger amount 0-9
	//
	// example:
	//
	// 1
	Infants *int32 `json:"infants,omitempty" xml:"infants,omitempty"`
	// journey list
	JourneyParamList []*EnrichRequestJourneyParamList `json:"journey_param_list,omitempty" xml:"journey_param_list,omitempty" type:"Repeated"`
	// solution_id returned by Search
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s EnrichRequest) String() string {
	return tea.Prettify(s)
}

func (s EnrichRequest) GoString() string {
	return s.String()
}

func (s *EnrichRequest) SetAdults(v int32) *EnrichRequest {
	s.Adults = &v
	return s
}

func (s *EnrichRequest) SetCabinClass(v string) *EnrichRequest {
	s.CabinClass = &v
	return s
}

func (s *EnrichRequest) SetChildren(v int32) *EnrichRequest {
	s.Children = &v
	return s
}

func (s *EnrichRequest) SetInfants(v int32) *EnrichRequest {
	s.Infants = &v
	return s
}

func (s *EnrichRequest) SetJourneyParamList(v []*EnrichRequestJourneyParamList) *EnrichRequest {
	s.JourneyParamList = v
	return s
}

func (s *EnrichRequest) SetSolutionId(v string) *EnrichRequest {
	s.SolutionId = &v
	return s
}

type EnrichRequestJourneyParamList struct {
	// arrival city code (capitalized)
	//
	// This parameter is required.
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// departure city code (capitalized)
	//
	// This parameter is required.
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// departure date (eg: yyyyMMdd)
	//
	// This parameter is required.
	//
	// example:
	//
	// 20230310
	DepartureDate *string `json:"departure_date,omitempty" xml:"departure_date,omitempty"`
	// segement param list
	//
	// This parameter is required.
	SegmentParamList []*EnrichRequestJourneyParamListSegmentParamList `json:"segment_param_list,omitempty" xml:"segment_param_list,omitempty" type:"Repeated"`
}

func (s EnrichRequestJourneyParamList) String() string {
	return tea.Prettify(s)
}

func (s EnrichRequestJourneyParamList) GoString() string {
	return s.String()
}

func (s *EnrichRequestJourneyParamList) SetArrivalCity(v string) *EnrichRequestJourneyParamList {
	s.ArrivalCity = &v
	return s
}

func (s *EnrichRequestJourneyParamList) SetDepartureCity(v string) *EnrichRequestJourneyParamList {
	s.DepartureCity = &v
	return s
}

func (s *EnrichRequestJourneyParamList) SetDepartureDate(v string) *EnrichRequestJourneyParamList {
	s.DepartureDate = &v
	return s
}

func (s *EnrichRequestJourneyParamList) SetSegmentParamList(v []*EnrichRequestJourneyParamListSegmentParamList) *EnrichRequestJourneyParamList {
	s.SegmentParamList = v
	return s
}

type EnrichRequestJourneyParamListSegmentParamList struct {
	// arrival airport code (capitalized)
	//
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// arrival city code (capitalized)
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// RBD
	//
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// child RBD
	//
	// example:
	//
	// E
	ChildCabin *string `json:"child_cabin,omitempty" xml:"child_cabin,omitempty"`
	// departure airport code (capitalized)
	//
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// departure city code (capitalized)
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// departure time in string format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 07:55:00
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// marketing flight no. (eg: KA5809)
	//
	// This parameter is required.
	//
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
}

func (s EnrichRequestJourneyParamListSegmentParamList) String() string {
	return tea.Prettify(s)
}

func (s EnrichRequestJourneyParamListSegmentParamList) GoString() string {
	return s.String()
}

func (s *EnrichRequestJourneyParamListSegmentParamList) SetArrivalAirport(v string) *EnrichRequestJourneyParamListSegmentParamList {
	s.ArrivalAirport = &v
	return s
}

func (s *EnrichRequestJourneyParamListSegmentParamList) SetArrivalCity(v string) *EnrichRequestJourneyParamListSegmentParamList {
	s.ArrivalCity = &v
	return s
}

func (s *EnrichRequestJourneyParamListSegmentParamList) SetCabin(v string) *EnrichRequestJourneyParamListSegmentParamList {
	s.Cabin = &v
	return s
}

func (s *EnrichRequestJourneyParamListSegmentParamList) SetChildCabin(v string) *EnrichRequestJourneyParamListSegmentParamList {
	s.ChildCabin = &v
	return s
}

func (s *EnrichRequestJourneyParamListSegmentParamList) SetDepartureAirport(v string) *EnrichRequestJourneyParamListSegmentParamList {
	s.DepartureAirport = &v
	return s
}

func (s *EnrichRequestJourneyParamListSegmentParamList) SetDepartureCity(v string) *EnrichRequestJourneyParamListSegmentParamList {
	s.DepartureCity = &v
	return s
}

func (s *EnrichRequestJourneyParamListSegmentParamList) SetDepartureTime(v string) *EnrichRequestJourneyParamListSegmentParamList {
	s.DepartureTime = &v
	return s
}

func (s *EnrichRequestJourneyParamListSegmentParamList) SetMarketingFlightNo(v string) *EnrichRequestJourneyParamListSegmentParamList {
	s.MarketingFlightNo = &v
	return s
}

type EnrichShrinkRequest struct {
	// adult passenger amount 1-9
	//
	// example:
	//
	// 1
	Adults *int32 `json:"adults,omitempty" xml:"adults,omitempty"`
	// cabin class
	//
	// 1. **ALL_CABIN*	- : all cabin class
	//
	// 2. **Y*	- : economy class
	//
	// 3. **FC*	- : first class and business class
	//
	// 4. **S*	- : premium economy class
	//
	// 5. **YS*	- : economy class and premium economy class
	//
	// 6. **YSC*	- : economy class, premium economy class and business class
	//
	// example:
	//
	// ALL_CABIN
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// child passenger amount 0-9
	//
	// example:
	//
	// 1
	Children *int32 `json:"children,omitempty" xml:"children,omitempty"`
	// infant passenger amount 0-9
	//
	// example:
	//
	// 1
	Infants *int32 `json:"infants,omitempty" xml:"infants,omitempty"`
	// journey list
	JourneyParamListShrink *string `json:"journey_param_list,omitempty" xml:"journey_param_list,omitempty"`
	// solution_id returned by Search
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s EnrichShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s EnrichShrinkRequest) GoString() string {
	return s.String()
}

func (s *EnrichShrinkRequest) SetAdults(v int32) *EnrichShrinkRequest {
	s.Adults = &v
	return s
}

func (s *EnrichShrinkRequest) SetCabinClass(v string) *EnrichShrinkRequest {
	s.CabinClass = &v
	return s
}

func (s *EnrichShrinkRequest) SetChildren(v int32) *EnrichShrinkRequest {
	s.Children = &v
	return s
}

func (s *EnrichShrinkRequest) SetInfants(v int32) *EnrichShrinkRequest {
	s.Infants = &v
	return s
}

func (s *EnrichShrinkRequest) SetJourneyParamListShrink(v string) *EnrichShrinkRequest {
	s.JourneyParamListShrink = &v
	return s
}

func (s *EnrichShrinkRequest) SetSolutionId(v string) *EnrichShrinkRequest {
	s.SolutionId = &v
	return s
}

type EnrichResponseBody struct {
	// request ID
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *EnrichResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error data
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// http reqeust has been processed successfully，status code is 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true represents success, false represents failure
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EnrichResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnrichResponseBody) GoString() string {
	return s.String()
}

func (s *EnrichResponseBody) SetRequestId(v string) *EnrichResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnrichResponseBody) SetData(v *EnrichResponseBodyData) *EnrichResponseBody {
	s.Data = v
	return s
}

func (s *EnrichResponseBody) SetErrorCode(v string) *EnrichResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *EnrichResponseBody) SetErrorData(v interface{}) *EnrichResponseBody {
	s.ErrorData = v
	return s
}

func (s *EnrichResponseBody) SetErrorMsg(v string) *EnrichResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *EnrichResponseBody) SetStatus(v int32) *EnrichResponseBody {
	s.Status = &v
	return s
}

func (s *EnrichResponseBody) SetSuccess(v bool) *EnrichResponseBody {
	s.Success = &v
	return s
}

type EnrichResponseBodyData struct {
	// solution list
	SolutionList []*EnrichResponseBodyDataSolutionList `json:"solution_list,omitempty" xml:"solution_list,omitempty" type:"Repeated"`
}

func (s EnrichResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnrichResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnrichResponseBodyData) SetSolutionList(v []*EnrichResponseBodyDataSolutionList) *EnrichResponseBodyData {
	s.SolutionList = v
	return s
}

type EnrichResponseBodyDataSolutionList struct {
	// adult fare
	//
	// example:
	//
	// 500
	AdultPrice *float64 `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	// adult tax
	//
	// example:
	//
	// 100
	AdultTax *float64 `json:"adult_tax,omitempty" xml:"adult_tax,omitempty"`
	// child fare
	//
	// example:
	//
	// 100
	ChildPrice *float64 `json:"child_price,omitempty" xml:"child_price,omitempty"`
	// child tax
	//
	// example:
	//
	// 100
	ChildTax *float64 `json:"child_tax,omitempty" xml:"child_tax,omitempty"`
	// infant fare
	//
	// example:
	//
	// 500
	InfantPrice *float64 `json:"infant_price,omitempty" xml:"infant_price,omitempty"`
	// infant tax
	//
	// example:
	//
	// 100
	InfantTax *float64 `json:"infant_tax,omitempty" xml:"infant_tax,omitempty"`
	// journey list
	JourneyList []*EnrichResponseBodyDataSolutionListJourneyList `json:"journey_list,omitempty" xml:"journey_list,omitempty" type:"Repeated"`
	// product type description
	//
	// example:
	//
	// ""
	ProductTypeDescription *string `json:"product_type_description,omitempty" xml:"product_type_description,omitempty"`
	// refund airline coupon description
	//
	// example:
	//
	// ""
	RefundTicketCouponDescription *string `json:"refund_ticket_coupon_description,omitempty" xml:"refund_ticket_coupon_description,omitempty"`
	// through check-in baggage  policy
	SegmentBaggageCheckInInfoList []*EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList `json:"segment_baggage_check_in_info_list,omitempty" xml:"segment_baggage_check_in_info_list,omitempty" type:"Repeated"`
	// baggage rule
	SegmentBaggageMappingList []*EnrichResponseBodyDataSolutionListSegmentBaggageMappingList `json:"segment_baggage_mapping_list,omitempty" xml:"segment_baggage_mapping_list,omitempty" type:"Repeated"`
	// change and refund policy
	SegmentRefundChangeRuleMappingList []*EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList `json:"segment_refund_change_rule_mapping_list,omitempty" xml:"segment_refund_change_rule_mapping_list,omitempty" type:"Repeated"`
	// solution ID
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s EnrichResponseBodyDataSolutionList) String() string {
	return tea.Prettify(s)
}

func (s EnrichResponseBodyDataSolutionList) GoString() string {
	return s.String()
}

func (s *EnrichResponseBodyDataSolutionList) SetAdultPrice(v float64) *EnrichResponseBodyDataSolutionList {
	s.AdultPrice = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionList) SetAdultTax(v float64) *EnrichResponseBodyDataSolutionList {
	s.AdultTax = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionList) SetChildPrice(v float64) *EnrichResponseBodyDataSolutionList {
	s.ChildPrice = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionList) SetChildTax(v float64) *EnrichResponseBodyDataSolutionList {
	s.ChildTax = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionList) SetInfantPrice(v float64) *EnrichResponseBodyDataSolutionList {
	s.InfantPrice = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionList) SetInfantTax(v float64) *EnrichResponseBodyDataSolutionList {
	s.InfantTax = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionList) SetJourneyList(v []*EnrichResponseBodyDataSolutionListJourneyList) *EnrichResponseBodyDataSolutionList {
	s.JourneyList = v
	return s
}

func (s *EnrichResponseBodyDataSolutionList) SetProductTypeDescription(v string) *EnrichResponseBodyDataSolutionList {
	s.ProductTypeDescription = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionList) SetRefundTicketCouponDescription(v string) *EnrichResponseBodyDataSolutionList {
	s.RefundTicketCouponDescription = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionList) SetSegmentBaggageCheckInInfoList(v []*EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) *EnrichResponseBodyDataSolutionList {
	s.SegmentBaggageCheckInInfoList = v
	return s
}

func (s *EnrichResponseBodyDataSolutionList) SetSegmentBaggageMappingList(v []*EnrichResponseBodyDataSolutionListSegmentBaggageMappingList) *EnrichResponseBodyDataSolutionList {
	s.SegmentBaggageMappingList = v
	return s
}

func (s *EnrichResponseBodyDataSolutionList) SetSegmentRefundChangeRuleMappingList(v []*EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) *EnrichResponseBodyDataSolutionList {
	s.SegmentRefundChangeRuleMappingList = v
	return s
}

func (s *EnrichResponseBodyDataSolutionList) SetSolutionId(v string) *EnrichResponseBodyDataSolutionList {
	s.SolutionId = &v
	return s
}

type EnrichResponseBodyDataSolutionListJourneyList struct {
	// segment Info
	SegmentList []*EnrichResponseBodyDataSolutionListJourneyListSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// number of transfers
	//
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s EnrichResponseBodyDataSolutionListJourneyList) String() string {
	return tea.Prettify(s)
}

func (s EnrichResponseBodyDataSolutionListJourneyList) GoString() string {
	return s.String()
}

func (s *EnrichResponseBodyDataSolutionListJourneyList) SetSegmentList(v []*EnrichResponseBodyDataSolutionListJourneyListSegmentList) *EnrichResponseBodyDataSolutionListJourneyList {
	s.SegmentList = v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyList) SetTransferCount(v int32) *EnrichResponseBodyDataSolutionListJourneyList {
	s.TransferCount = &v
	return s
}

type EnrichResponseBodyDataSolutionListJourneyListSegmentList struct {
	// arrival airport code (capitalized)
	//
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// arrival city code (capitalized)
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// arrival terminal
	//
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// arrival time in string format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 10:40:00
	ArrivalTime *string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// available seats (for reference only)
	//
	// example:
	//
	// 7
	Availability *string `json:"availability,omitempty" xml:"availability,omitempty"`
	// RBD
	//
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// cabin class
	//
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// code share or not
	//
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// departure airport code (capitalized)
	//
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// departure city code (capitalized)
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// departure terminal
	//
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// departure time in string format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 07:55:00
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// equipment type
	//
	// example:
	//
	// 32Q
	EquipType *string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// flight time, unit: minute
	//
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// marketing airline code (eg: KA)
	//
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// marketing airline flight no. (eg: KA5809)
	//
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// marketing airline integer flight no. (eg: 5809)
	//
	// example:
	//
	// 1295
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// operating airline code (eg: CX)
	//
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// operating airline flight no. (eg: CX601)
	//
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// segment ID format: flight no.+departure airport[IATA airport code]+arrival airport[IATA airport code]+departure time(MMdd)
	//
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// stop city list.
	//
	// when stop_quantity > 1 , use “,” for seperation
	//
	// example:
	//
	// MFM,PVG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// number of stops
	//
	// example:
	//
	// 0
	StopQuantity *int32 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s EnrichResponseBodyDataSolutionListJourneyListSegmentList) String() string {
	return tea.Prettify(s)
}

func (s EnrichResponseBodyDataSolutionListJourneyListSegmentList) GoString() string {
	return s.String()
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalAirport(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalCity(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalTerminal(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalTime(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetAvailability(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.Availability = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetCabin(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.Cabin = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetCabinClass(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.CabinClass = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetCodeShare(v bool) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.CodeShare = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureAirport(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureCity(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureTerminal(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureTime(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetEquipType(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.EquipType = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetFlightDuration(v int32) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetMarketingAirline(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetMarketingFlightNo(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetMarketingFlightNoInt(v int32) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetOperatingAirline(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetOperatingFlightNo(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetSegmentId(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.SegmentId = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetStopCityList(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.StopCityList = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetStopQuantity(v int32) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
	s.StopQuantity = &v
	return s
}

type EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList struct {
	// through check-in baggage policy type
	//
	// 1. baggage through check-in between segments
	//
	// 2. baggage re-check-in needed between segments
	//
	// 4. baggage through check-in at stop city ( applies for stop flight )
	//
	// 3. baggage re-checkin needed at stop city ( applies for stop flight )
	//
	// example:
	//
	// 1
	LuggageDirectInfoType *int32 `json:"luggage_direct_info_type,omitempty" xml:"luggage_direct_info_type,omitempty"`
	// segment id list. all the listed segment ids share the same baggage through check-in policy
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) String() string {
	return tea.Prettify(s)
}

func (s EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) GoString() string {
	return s.String()
}

func (s *EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) SetLuggageDirectInfoType(v int32) *EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList {
	s.LuggageDirectInfoType = &v
	return s
}

func (s *EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) SetSegmentIdList(v []*string) *EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList {
	s.SegmentIdList = v
	return s
}

type EnrichResponseBodyDataSolutionListSegmentBaggageMappingList struct {
	// baggage rule mapping, key is passenger type, value is baggage allowance details
	PassengerBaggageAllowanceMapping map[string]*DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue `json:"passenger_baggage_allowance_mapping,omitempty" xml:"passenger_baggage_allowance_mapping,omitempty"`
	// segment id list.
	//
	// all the listed segment ids share the same baggage rule
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s EnrichResponseBodyDataSolutionListSegmentBaggageMappingList) String() string {
	return tea.Prettify(s)
}

func (s EnrichResponseBodyDataSolutionListSegmentBaggageMappingList) GoString() string {
	return s.String()
}

func (s *EnrichResponseBodyDataSolutionListSegmentBaggageMappingList) SetPassengerBaggageAllowanceMapping(v map[string]*DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) *EnrichResponseBodyDataSolutionListSegmentBaggageMappingList {
	s.PassengerBaggageAllowanceMapping = v
	return s
}

func (s *EnrichResponseBodyDataSolutionListSegmentBaggageMappingList) SetSegmentIdList(v []*string) *EnrichResponseBodyDataSolutionListSegmentBaggageMappingList {
	s.SegmentIdList = v
	return s
}

type EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList struct {
	// change and refund policy mapping, key is passenger type, value is change and refund policy detail
	RefundChangeRuleMap map[string]*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue `json:"refund_change_rule_map,omitempty" xml:"refund_change_rule_map,omitempty"`
	// segment id list. all the listed segment ids share the same change and refund policy
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) String() string {
	return tea.Prettify(s)
}

func (s EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) GoString() string {
	return s.String()
}

func (s *EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) SetRefundChangeRuleMap(v map[string]*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) *EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList {
	s.RefundChangeRuleMap = v
	return s
}

func (s *EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) SetSegmentIdList(v []*string) *EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList {
	s.SegmentIdList = v
	return s
}

type EnrichResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnrichResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnrichResponse) String() string {
	return tea.Prettify(s)
}

func (s EnrichResponse) GoString() string {
	return s.String()
}

func (s *EnrichResponse) SetHeaders(v map[string]*string) *EnrichResponse {
	s.Headers = v
	return s
}

func (s *EnrichResponse) SetStatusCode(v int32) *EnrichResponse {
	s.StatusCode = &v
	return s
}

func (s *EnrichResponse) SetBody(v *EnrichResponseBody) *EnrichResponse {
	s.Body = v
	return s
}

type FileUploadHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access_token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s FileUploadHeaders) String() string {
	return tea.Prettify(s)
}

func (s FileUploadHeaders) GoString() string {
	return s.String()
}

func (s *FileUploadHeaders) SetCommonHeaders(v map[string]*string) *FileUploadHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FileUploadHeaders) SetXAcsAirticketAccessToken(v string) *FileUploadHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *FileUploadHeaders) SetXAcsAirticketLanguage(v string) *FileUploadHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type FileUploadRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0r2LSuIsHlxEoGZcnGe34U1njBOR83Q4HNSvMDGrDPK5J71VjcGdRIWz2x3+tFxvQaduwHB46Z9K
	//
	// dbIoDN8xPQ5PWlky8rKOPmAqSZfIRyPmAwvPvTJFwr8bRgHPPaq2VO8kHJ6jFIpJJ5I7Zqd1BjGS
	//
	// SR/kULQZHsDDd2zgA9RRTsEQF2OSxFFFx2P/2Q==
	FileContent *string `json:"file_content,omitempty" xml:"file_content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 496***2617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s FileUploadRequest) String() string {
	return tea.Prettify(s)
}

func (s FileUploadRequest) GoString() string {
	return s.String()
}

func (s *FileUploadRequest) SetFileContent(v string) *FileUploadRequest {
	s.FileContent = &v
	return s
}

func (s *FileUploadRequest) SetOrderNum(v int64) *FileUploadRequest {
	s.OrderNum = &v
	return s
}

type FileUploadResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *FileUploadResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s FileUploadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FileUploadResponseBody) GoString() string {
	return s.String()
}

func (s *FileUploadResponseBody) SetRequestId(v string) *FileUploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *FileUploadResponseBody) SetData(v *FileUploadResponseBodyData) *FileUploadResponseBody {
	s.Data = v
	return s
}

func (s *FileUploadResponseBody) SetErrorCode(v string) *FileUploadResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *FileUploadResponseBody) SetErrorData(v interface{}) *FileUploadResponseBody {
	s.ErrorData = v
	return s
}

func (s *FileUploadResponseBody) SetErrorMsg(v string) *FileUploadResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *FileUploadResponseBody) SetStatus(v int32) *FileUploadResponseBody {
	s.Status = &v
	return s
}

func (s *FileUploadResponseBody) SetSuccess(v bool) *FileUploadResponseBody {
	s.Success = &v
	return s
}

type FileUploadResponseBodyData struct {
	// example:
	//
	// https://fliggy-flight-jinghang-bucket.oss-cn-zhangjiakou.aliyuncs.com/suez/flight_suez_9a634376****47.jpeg
	UploadedFileUrl *string `json:"uploaded_file_url,omitempty" xml:"uploaded_file_url,omitempty"`
}

func (s FileUploadResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FileUploadResponseBodyData) GoString() string {
	return s.String()
}

func (s *FileUploadResponseBodyData) SetUploadedFileUrl(v string) *FileUploadResponseBodyData {
	s.UploadedFileUrl = &v
	return s
}

type FileUploadResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FileUploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FileUploadResponse) String() string {
	return tea.Prettify(s)
}

func (s FileUploadResponse) GoString() string {
	return s.String()
}

func (s *FileUploadResponse) SetHeaders(v map[string]*string) *FileUploadResponse {
	s.Headers = v
	return s
}

func (s *FileUploadResponse) SetStatusCode(v int32) *FileUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *FileUploadResponse) SetBody(v *FileUploadResponseBody) *FileUploadResponse {
	s.Body = v
	return s
}

type FlightChangeOfOrderHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access_token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s FlightChangeOfOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s FlightChangeOfOrderHeaders) GoString() string {
	return s.String()
}

func (s *FlightChangeOfOrderHeaders) SetCommonHeaders(v map[string]*string) *FlightChangeOfOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FlightChangeOfOrderHeaders) SetXAcsAirticketAccessToken(v string) *FlightChangeOfOrderHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *FlightChangeOfOrderHeaders) SetXAcsAirticketLanguage(v string) *FlightChangeOfOrderHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type FlightChangeOfOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s FlightChangeOfOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s FlightChangeOfOrderRequest) GoString() string {
	return s.String()
}

func (s *FlightChangeOfOrderRequest) SetOrderNum(v int64) *FlightChangeOfOrderRequest {
	s.OrderNum = &v
	return s
}

type FlightChangeOfOrderResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*FlightChangeOfOrderResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s FlightChangeOfOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FlightChangeOfOrderResponseBody) GoString() string {
	return s.String()
}

func (s *FlightChangeOfOrderResponseBody) SetRequestId(v string) *FlightChangeOfOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightChangeOfOrderResponseBody) SetData(v []*FlightChangeOfOrderResponseBodyData) *FlightChangeOfOrderResponseBody {
	s.Data = v
	return s
}

func (s *FlightChangeOfOrderResponseBody) SetErrorCode(v string) *FlightChangeOfOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *FlightChangeOfOrderResponseBody) SetErrorData(v interface{}) *FlightChangeOfOrderResponseBody {
	s.ErrorData = v
	return s
}

func (s *FlightChangeOfOrderResponseBody) SetErrorMsg(v string) *FlightChangeOfOrderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *FlightChangeOfOrderResponseBody) SetStatus(v int32) *FlightChangeOfOrderResponseBody {
	s.Status = &v
	return s
}

func (s *FlightChangeOfOrderResponseBody) SetSuccess(v bool) *FlightChangeOfOrderResponseBody {
	s.Success = &v
	return s
}

type FlightChangeOfOrderResponseBodyData struct {
	FlightChangeDetail *FlightChangeOfOrderResponseBodyDataFlightChangeDetail `json:"flight_change_detail,omitempty" xml:"flight_change_detail,omitempty" type:"Struct"`
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s FlightChangeOfOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FlightChangeOfOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *FlightChangeOfOrderResponseBodyData) SetFlightChangeDetail(v *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) *FlightChangeOfOrderResponseBodyData {
	s.FlightChangeDetail = v
	return s
}

func (s *FlightChangeOfOrderResponseBodyData) SetOrderNum(v int64) *FlightChangeOfOrderResponseBodyData {
	s.OrderNum = &v
	return s
}

type FlightChangeOfOrderResponseBodyDataFlightChangeDetail struct {
	// example:
	//
	// 天气
	ChangeReason *string `json:"change_reason,omitempty" xml:"change_reason,omitempty"`
	// example:
	//
	// 2023-02-01 10:01:00
	ChangeTime *string `json:"change_time,omitempty" xml:"change_time,omitempty"`
	// example:
	//
	// 1
	ChangeType *int32 `json:"change_type,omitempty" xml:"change_type,omitempty"`
	// example:
	//
	// MFM
	NewArrivalAirport *string `json:"new_arrival_airport,omitempty" xml:"new_arrival_airport,omitempty"`
	// example:
	//
	// 2023-02-01 15:01:00
	NewArrivalTime *string `json:"new_arrival_time,omitempty" xml:"new_arrival_time,omitempty"`
	// example:
	//
	// PVG
	NewDepartureAirport *string `json:"new_departure_airport,omitempty" xml:"new_departure_airport,omitempty"`
	// example:
	//
	// 2023-02-01 13:01:00
	NewDepartureTime *string `json:"new_departure_time,omitempty" xml:"new_departure_time,omitempty"`
	// example:
	//
	// HO1295
	NewFlightNo *string `json:"new_flight_no,omitempty" xml:"new_flight_no,omitempty"`
	// example:
	//
	// MFM
	OldArrivalAirport *string `json:"old_arrival_airport,omitempty" xml:"old_arrival_airport,omitempty"`
	// example:
	//
	// 023-02-01 14:01:00
	OldArrivalTime *string `json:"old_arrival_time,omitempty" xml:"old_arrival_time,omitempty"`
	// example:
	//
	// PVG
	OldDepartureAirport *string `json:"old_departure_airport,omitempty" xml:"old_departure_airport,omitempty"`
	// example:
	//
	// 2023-02-01 12:01:00
	OldDepartureTime *string `json:"old_departure_time,omitempty" xml:"old_departure_time,omitempty"`
	// example:
	//
	// HO1295
	OldFlightNo *string `json:"old_flight_no,omitempty" xml:"old_flight_no,omitempty"`
}

func (s FlightChangeOfOrderResponseBodyDataFlightChangeDetail) String() string {
	return tea.Prettify(s)
}

func (s FlightChangeOfOrderResponseBodyDataFlightChangeDetail) GoString() string {
	return s.String()
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetChangeReason(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.ChangeReason = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetChangeTime(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.ChangeTime = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetChangeType(v int32) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.ChangeType = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetNewArrivalAirport(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.NewArrivalAirport = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetNewArrivalTime(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.NewArrivalTime = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetNewDepartureAirport(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.NewDepartureAirport = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetNewDepartureTime(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.NewDepartureTime = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetNewFlightNo(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.NewFlightNo = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetOldArrivalAirport(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.OldArrivalAirport = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetOldArrivalTime(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.OldArrivalTime = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetOldDepartureAirport(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.OldDepartureAirport = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetOldDepartureTime(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.OldDepartureTime = &v
	return s
}

func (s *FlightChangeOfOrderResponseBodyDataFlightChangeDetail) SetOldFlightNo(v string) *FlightChangeOfOrderResponseBodyDataFlightChangeDetail {
	s.OldFlightNo = &v
	return s
}

type FlightChangeOfOrderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlightChangeOfOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlightChangeOfOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s FlightChangeOfOrderResponse) GoString() string {
	return s.String()
}

func (s *FlightChangeOfOrderResponse) SetHeaders(v map[string]*string) *FlightChangeOfOrderResponse {
	s.Headers = v
	return s
}

func (s *FlightChangeOfOrderResponse) SetStatusCode(v int32) *FlightChangeOfOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *FlightChangeOfOrderResponse) SetBody(v *FlightChangeOfOrderResponseBody) *FlightChangeOfOrderResponse {
	s.Body = v
	return s
}

type GetTokenRequest struct {
	// appKey
	//
	// This parameter is required.
	//
	// example:
	//
	// fu1bltcu3400iurywuri
	AppKey *string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// appSecret
	//
	// This parameter is required.
	//
	// example:
	//
	// ZzQ3MW1mb3E1ODAwI2ldUjYlWUdJn5YI
	AppSecret *string `json:"app_secret,omitempty" xml:"app_secret,omitempty"`
}

func (s GetTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) SetAppKey(v string) *GetTokenRequest {
	s.AppKey = &v
	return s
}

func (s *GetTokenRequest) SetAppSecret(v string) *GetTokenRequest {
	s.AppSecret = &v
	return s
}

type GetTokenResponseBody struct {
	// Request RequestId
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Correctly processed return data
	Data *GetTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Business error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// Data carried during error handling
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// Error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// When the HTTP request is successful, the status value is 200.
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// Whether it is correct
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBody) SetRequestId(v string) *GetTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenResponseBody) SetData(v *GetTokenResponseBodyData) *GetTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetTokenResponseBody) SetErrorCode(v string) *GetTokenResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTokenResponseBody) SetErrorData(v interface{}) *GetTokenResponseBody {
	s.ErrorData = v
	return s
}

func (s *GetTokenResponseBody) SetErrorMsg(v string) *GetTokenResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetTokenResponseBody) SetStatus(v int32) *GetTokenResponseBody {
	s.Status = &v
	return s
}

func (s *GetTokenResponseBody) SetSuccess(v bool) *GetTokenResponseBody {
	s.Success = &v
	return s
}

type GetTokenResponseBodyData struct {
	// Remaining valid time of the token in seconds
	//
	// example:
	//
	// 7200
	ExpireTime *int64 `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// Timestamp of token generation in seconds
	//
	// example:
	//
	// 1677055176
	GenerateTime *int64 `json:"generate_time,omitempty" xml:"generate_time,omitempty"`
	// token
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiIyUlRjY0Ezc1puSndpYU11R1ctTkVRIiwiaWF0IjoxNjc3MTU1Njg3LCJleHAiOjE2NzcxNjI4ODcsIm5iZiI6MTY3NzE1NTYyN30.bd8qBedJ7R24NC8VpMtM4Ni5OR-Cc0utPiKSx8fjoj9taalt7zXBF8uIVTETw1N-Fx9reLflwVXrbDHky7ZKqlE5o_B5Bkx-crQKlJL-NzasEmNnuJNb5kmmbCy3mvIrQfo5Q5Y0ZaQ110pXK4qd1shRbklvuQXn8lPueJtNqo8VdIOAPGG_rPwOw2P767I0fyFHcome8FR4ST1UrwxeApNKMB_BkpCsUZLgpm9h9trhKbB-3qtk6UK1GKmfw6qlWpL3PQN7FAObNruS0r0CGh3Muc9PaGsuu8Xu5on21h9WmI7L0-jatZkS55p4PEerU56XpkwJfa35_hltKNTauu
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GetTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBodyData) SetExpireTime(v int64) *GetTokenResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetTokenResponseBodyData) SetGenerateTime(v int64) *GetTokenResponseBodyData {
	s.GenerateTime = &v
	return s
}

func (s *GetTokenResponseBodyData) SetToken(v string) *GetTokenResponseBodyData {
	s.Token = &v
	return s
}

type GetTokenResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponse) GoString() string {
	return s.String()
}

func (s *GetTokenResponse) SetHeaders(v map[string]*string) *GetTokenResponse {
	s.Headers = v
	return s
}

func (s *GetTokenResponse) SetStatusCode(v int32) *GetTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTokenResponse) SetBody(v *GetTokenResponseBody) *GetTokenResponse {
	s.Body = v
	return s
}

type LuggageDirectHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access_token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s LuggageDirectHeaders) String() string {
	return tea.Prettify(s)
}

func (s LuggageDirectHeaders) GoString() string {
	return s.String()
}

func (s *LuggageDirectHeaders) SetCommonHeaders(v map[string]*string) *LuggageDirectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LuggageDirectHeaders) SetXAcsAirticketAccessToken(v string) *LuggageDirectHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *LuggageDirectHeaders) SetXAcsAirticketLanguage(v string) *LuggageDirectHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type LuggageDirectRequest struct {
	FlightSegmentParamList []*LuggageDirectRequestFlightSegmentParamList `json:"flight_segment_param_list,omitempty" xml:"flight_segment_param_list,omitempty" type:"Repeated"`
}

func (s LuggageDirectRequest) String() string {
	return tea.Prettify(s)
}

func (s LuggageDirectRequest) GoString() string {
	return s.String()
}

func (s *LuggageDirectRequest) SetFlightSegmentParamList(v []*LuggageDirectRequestFlightSegmentParamList) *LuggageDirectRequest {
	s.FlightSegmentParamList = v
	return s
}

type LuggageDirectRequestFlightSegmentParamList struct {
	// This parameter is required.
	//
	// example:
	//
	// SIN
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1705285430445
	ArrivalTime *int64 `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HGH
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// example:
	//
	// T1
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1705285430445
	DepartureTime *int64 `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CZ
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CZ616
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// example:
	//
	// CZ
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// example:
	//
	// SEL,HKG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// example:
	//
	// CZ
	TicketingAirline *string `json:"ticketing_airline,omitempty" xml:"ticketing_airline,omitempty"`
}

func (s LuggageDirectRequestFlightSegmentParamList) String() string {
	return tea.Prettify(s)
}

func (s LuggageDirectRequestFlightSegmentParamList) GoString() string {
	return s.String()
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetArrivalAirport(v string) *LuggageDirectRequestFlightSegmentParamList {
	s.ArrivalAirport = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetArrivalTerminal(v string) *LuggageDirectRequestFlightSegmentParamList {
	s.ArrivalTerminal = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetArrivalTime(v int64) *LuggageDirectRequestFlightSegmentParamList {
	s.ArrivalTime = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetCodeShare(v bool) *LuggageDirectRequestFlightSegmentParamList {
	s.CodeShare = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetDepartureAirport(v string) *LuggageDirectRequestFlightSegmentParamList {
	s.DepartureAirport = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetDepartureTerminal(v string) *LuggageDirectRequestFlightSegmentParamList {
	s.DepartureTerminal = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetDepartureTime(v int64) *LuggageDirectRequestFlightSegmentParamList {
	s.DepartureTime = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetMarketingAirline(v string) *LuggageDirectRequestFlightSegmentParamList {
	s.MarketingAirline = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetMarketingFlightNo(v string) *LuggageDirectRequestFlightSegmentParamList {
	s.MarketingFlightNo = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetOperatingAirline(v string) *LuggageDirectRequestFlightSegmentParamList {
	s.OperatingAirline = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetStopCityList(v string) *LuggageDirectRequestFlightSegmentParamList {
	s.StopCityList = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetTicketingAirline(v string) *LuggageDirectRequestFlightSegmentParamList {
	s.TicketingAirline = &v
	return s
}

type LuggageDirectShrinkRequest struct {
	FlightSegmentParamListShrink *string `json:"flight_segment_param_list,omitempty" xml:"flight_segment_param_list,omitempty"`
}

func (s LuggageDirectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s LuggageDirectShrinkRequest) GoString() string {
	return s.String()
}

func (s *LuggageDirectShrinkRequest) SetFlightSegmentParamListShrink(v string) *LuggageDirectShrinkRequest {
	s.FlightSegmentParamListShrink = &v
	return s
}

type LuggageDirectResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*LuggageDirectResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s LuggageDirectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LuggageDirectResponseBody) GoString() string {
	return s.String()
}

func (s *LuggageDirectResponseBody) SetRequestId(v string) *LuggageDirectResponseBody {
	s.RequestId = &v
	return s
}

func (s *LuggageDirectResponseBody) SetData(v []*LuggageDirectResponseBodyData) *LuggageDirectResponseBody {
	s.Data = v
	return s
}

func (s *LuggageDirectResponseBody) SetErrorCode(v string) *LuggageDirectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *LuggageDirectResponseBody) SetErrorData(v interface{}) *LuggageDirectResponseBody {
	s.ErrorData = v
	return s
}

func (s *LuggageDirectResponseBody) SetErrorMsg(v string) *LuggageDirectResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *LuggageDirectResponseBody) SetStatus(v int32) *LuggageDirectResponseBody {
	s.Status = &v
	return s
}

func (s *LuggageDirectResponseBody) SetSuccess(v bool) *LuggageDirectResponseBody {
	s.Success = &v
	return s
}

type LuggageDirectResponseBodyData struct {
	// example:
	//
	// BJS
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// example:
	//
	// 1
	DirectType *int32 `json:"direct_type,omitempty" xml:"direct_type,omitempty"`
}

func (s LuggageDirectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s LuggageDirectResponseBodyData) GoString() string {
	return s.String()
}

func (s *LuggageDirectResponseBodyData) SetCityCode(v string) *LuggageDirectResponseBodyData {
	s.CityCode = &v
	return s
}

func (s *LuggageDirectResponseBodyData) SetDirectType(v int32) *LuggageDirectResponseBodyData {
	s.DirectType = &v
	return s
}

type LuggageDirectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LuggageDirectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LuggageDirectResponse) String() string {
	return tea.Prettify(s)
}

func (s LuggageDirectResponse) GoString() string {
	return s.String()
}

func (s *LuggageDirectResponse) SetHeaders(v map[string]*string) *LuggageDirectResponse {
	s.Headers = v
	return s
}

func (s *LuggageDirectResponse) SetStatusCode(v int32) *LuggageDirectResponse {
	s.StatusCode = &v
	return s
}

func (s *LuggageDirectResponse) SetBody(v *LuggageDirectResponseBody) *LuggageDirectResponse {
	s.Body = v
	return s
}

type OrderDetailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// language code(refer to ISO_639), defaults to the buyer\\"s account configuration
	//
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s OrderDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailHeaders) GoString() string {
	return s.String()
}

func (s *OrderDetailHeaders) SetCommonHeaders(v map[string]*string) *OrderDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OrderDetailHeaders) SetXAcsAirticketAccessToken(v string) *OrderDetailHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *OrderDetailHeaders) SetXAcsAirticketLanguage(v string) *OrderDetailHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type OrderDetailRequest struct {
	// order number created by book
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// external order number(customized by buyer when book)
	//
	// example:
	//
	// x091-2023-0220-j-0001
	OutOrderNum *string `json:"out_order_num,omitempty" xml:"out_order_num,omitempty"`
}

func (s OrderDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailRequest) GoString() string {
	return s.String()
}

func (s *OrderDetailRequest) SetOrderNum(v int64) *OrderDetailRequest {
	s.OrderNum = &v
	return s
}

func (s *OrderDetailRequest) SetOutOrderNum(v string) *OrderDetailRequest {
	s.OutOrderNum = &v
	return s
}

type OrderDetailResponseBody struct {
	// request ID
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *OrderDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error data
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// http reqeust has been processed successfully，status code is 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true represents success, false represents failure
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OrderDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBody) SetRequestId(v string) *OrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *OrderDetailResponseBody) SetData(v *OrderDetailResponseBodyData) *OrderDetailResponseBody {
	s.Data = v
	return s
}

func (s *OrderDetailResponseBody) SetErrorCode(v string) *OrderDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *OrderDetailResponseBody) SetErrorData(v interface{}) *OrderDetailResponseBody {
	s.ErrorData = v
	return s
}

func (s *OrderDetailResponseBody) SetErrorMsg(v string) *OrderDetailResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *OrderDetailResponseBody) SetStatus(v int32) *OrderDetailResponseBody {
	s.Status = &v
	return s
}

func (s *OrderDetailResponseBody) SetSuccess(v bool) *OrderDetailResponseBody {
	s.Success = &v
	return s
}

type OrderDetailResponseBodyData struct {
	// ancillary product fulfillment details
	AncillaryItemDetailList []*OrderDetailResponseBodyDataAncillaryItemDetailList `json:"ancillary_item_detail_list,omitempty" xml:"ancillary_item_detail_list,omitempty" type:"Repeated"`
	// baggage rule mapping, key is passenger type, value is baggage allowance details
	BaggageAllowanceMap map[string]*DataBaggageAllowanceMapValue `json:"baggage_allowance_map,omitempty" xml:"baggage_allowance_map,omitempty"`
	// book time(timestamp)
	//
	// example:
	//
	// 1677210784000
	BookTime *int64 `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// flight product fulfillment details
	FlightItemDetailList []*OrderDetailResponseBodyDataFlightItemDetailList `json:"flight_item_detail_list,omitempty" xml:"flight_item_detail_list,omitempty" type:"Repeated"`
	// order number created by book
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// order status
	//
	// 1: order creation in process;
	//
	// 2: order creation successful;
	//
	// 3: order paid;
	//
	// 4: order successful;
	//
	// 5: order closed
	//
	// example:
	//
	// 4
	OrderStatus *int32 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// external order number(customized by buyer when book)
	//
	// example:
	//
	// x091-2023-0220-j-0001
	OutOrderNum *string `json:"out_order_num,omitempty" xml:"out_order_num,omitempty"`
	// the information about all passenger of current order
	PassengerList []*OrderDetailResponseBodyDataPassengerList `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
	// payment status
	//
	// 1: payment in process;
	//
	// 2: deduction successful;
	//
	// 3: paid to the seller;
	//
	// 4: transaction closed
	//
	// example:
	//
	// 2
	PayStatus *int32 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// pay time(timestamp)
	//
	// example:
	//
	// 1677210788000
	PayTime *int64 `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// discount amount
	//
	// example:
	//
	// 10
	PromotionPrice *float64 `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// actual payment amount
	//
	// example:
	//
	// 3000
	RealPayPrice *float64 `json:"real_pay_price,omitempty" xml:"real_pay_price,omitempty"`
	// change and refund policy mapping, key is passenger type, value is change and refund policy details
	RefundChangeRuleMap map[string]*DataRefundChangeRuleMapValue `json:"refund_change_rule_map,omitempty" xml:"refund_change_rule_map,omitempty"`
	// buyer nickname
	//
	// example:
	//
	// nick
	SessionNick *string `json:"session_nick,omitempty" xml:"session_nick,omitempty"`
	// the solution buyer booked
	Solution *OrderDetailResponseBodyDataSolution `json:"solution,omitempty" xml:"solution,omitempty" type:"Struct"`
	// order success time(timestamp)
	//
	// example:
	//
	// 1677210786000
	SucceedTime *int64 `json:"succeed_time,omitempty" xml:"succeed_time,omitempty"`
	// total price of current order
	//
	// example:
	//
	// 3000
	TotalPrice *float64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// transaction number
	//
	// example:
	//
	// hkduendkd-2023-dj0
	TransactionNo *string `json:"transaction_no,omitempty" xml:"transaction_no,omitempty"`
}

func (s OrderDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyData) SetAncillaryItemDetailList(v []*OrderDetailResponseBodyDataAncillaryItemDetailList) *OrderDetailResponseBodyData {
	s.AncillaryItemDetailList = v
	return s
}

func (s *OrderDetailResponseBodyData) SetBaggageAllowanceMap(v map[string]*DataBaggageAllowanceMapValue) *OrderDetailResponseBodyData {
	s.BaggageAllowanceMap = v
	return s
}

func (s *OrderDetailResponseBodyData) SetBookTime(v int64) *OrderDetailResponseBodyData {
	s.BookTime = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetFlightItemDetailList(v []*OrderDetailResponseBodyDataFlightItemDetailList) *OrderDetailResponseBodyData {
	s.FlightItemDetailList = v
	return s
}

func (s *OrderDetailResponseBodyData) SetOrderNum(v int64) *OrderDetailResponseBodyData {
	s.OrderNum = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetOrderStatus(v int32) *OrderDetailResponseBodyData {
	s.OrderStatus = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetOutOrderNum(v string) *OrderDetailResponseBodyData {
	s.OutOrderNum = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetPassengerList(v []*OrderDetailResponseBodyDataPassengerList) *OrderDetailResponseBodyData {
	s.PassengerList = v
	return s
}

func (s *OrderDetailResponseBodyData) SetPayStatus(v int32) *OrderDetailResponseBodyData {
	s.PayStatus = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetPayTime(v int64) *OrderDetailResponseBodyData {
	s.PayTime = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetPromotionPrice(v float64) *OrderDetailResponseBodyData {
	s.PromotionPrice = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetRealPayPrice(v float64) *OrderDetailResponseBodyData {
	s.RealPayPrice = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetRefundChangeRuleMap(v map[string]*DataRefundChangeRuleMapValue) *OrderDetailResponseBodyData {
	s.RefundChangeRuleMap = v
	return s
}

func (s *OrderDetailResponseBodyData) SetSessionNick(v string) *OrderDetailResponseBodyData {
	s.SessionNick = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetSolution(v *OrderDetailResponseBodyDataSolution) *OrderDetailResponseBodyData {
	s.Solution = v
	return s
}

func (s *OrderDetailResponseBodyData) SetSucceedTime(v int64) *OrderDetailResponseBodyData {
	s.SucceedTime = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetTotalPrice(v float64) *OrderDetailResponseBodyData {
	s.TotalPrice = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetTransactionNo(v string) *OrderDetailResponseBodyData {
	s.TransactionNo = &v
	return s
}

type OrderDetailResponseBodyDataAncillaryItemDetailList struct {
	// the ancillary buyer booked
	Ancillary *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary `json:"ancillary,omitempty" xml:"ancillary,omitempty" type:"Struct"`
	// passenger
	Passenger *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger `json:"passenger,omitempty" xml:"passenger,omitempty" type:"Struct"`
	// segment ID list, these segments share the same ancillary
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailList) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailList) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailList) SetAncillary(v *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary) *OrderDetailResponseBodyDataAncillaryItemDetailList {
	s.Ancillary = v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailList) SetPassenger(v *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) *OrderDetailResponseBodyDataAncillaryItemDetailList {
	s.Passenger = v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailList) SetSegmentIdList(v []*string) *OrderDetailResponseBodyDataAncillaryItemDetailList {
	s.SegmentIdList = v
	return s
}

type OrderDetailResponseBodyDataAncillaryItemDetailListAncillary struct {
	// ancillay_id
	//
	// example:
	//
	// MDY2NTAxLCJleHAiOjE2NxNzM3MDEsIm5ix
	AncillaryId *string `json:"ancillary_id,omitempty" xml:"ancillary_id,omitempty"`
	// ancillary product type currently supports 4: paid luggage
	//
	// example:
	//
	// 4
	AncillaryType    *int32                                                                       `json:"ancillary_type,omitempty" xml:"ancillary_type,omitempty"`
	BaggageAncillary *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary `json:"baggage_ancillary,omitempty" xml:"baggage_ancillary,omitempty" type:"Struct"`
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailListAncillary) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailListAncillary) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary) SetAncillaryId(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary {
	s.AncillaryId = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary) SetAncillaryType(v int32) *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary {
	s.AncillaryType = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary) SetBaggageAncillary(v *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary {
	s.BaggageAncillary = v
	return s
}

type OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary struct {
	// example:
	//
	// 0
	BaggageAmount *int32 `json:"baggage_amount,omitempty" xml:"baggage_amount,omitempty"`
	// example:
	//
	// 0
	BaggageWeight *int32 `json:"baggage_weight,omitempty" xml:"baggage_weight,omitempty"`
	// example:
	//
	// KG
	BaggageWeightUnit *string `json:"baggage_weight_unit,omitempty" xml:"baggage_weight_unit,omitempty"`
	IsAllWeight       *bool   `json:"is_all_weight,omitempty" xml:"is_all_weight,omitempty"`
	// example:
	//
	// 10.0
	Price *float64 `json:"price,omitempty" xml:"price,omitempty"`
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) SetBaggageAmount(v int32) *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary {
	s.BaggageAmount = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) SetBaggageWeight(v int32) *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary {
	s.BaggageWeight = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) SetBaggageWeightUnit(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary {
	s.BaggageWeightUnit = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) SetIsAllWeight(v bool) *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary {
	s.IsAllWeight = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) SetPrice(v float64) *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary {
	s.Price = &v
	return s
}

type OrderDetailResponseBodyDataAncillaryItemDetailListPassenger struct {
	// date of birth (yyyyMMdd)
	//
	// example:
	//
	// 20020301
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// travel document
	Credential *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential `json:"credential,omitempty" xml:"credential,omitempty" type:"Struct"`
	// first name
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// gender 0: male; 1: female
	//
	// example:
	//
	// 0
	Gender *int32 `json:"gender,omitempty" xml:"gender,omitempty"`
	// last name
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// country code for mobile phone number
	//
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// mobile phone number
	//
	// example:
	//
	// 183******96
	MobilePhoneNumber *string `json:"mobile_phone_number,omitempty" xml:"mobile_phone_number,omitempty"`
	// nationality
	//
	// example:
	//
	// CN
	Nationality *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// passenger type 0: adult; 1: child; 8: infant
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) SetBirthday(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger {
	s.Birthday = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) SetCredential(v *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential) *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger {
	s.Credential = v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) SetFirstName(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger {
	s.FirstName = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) SetGender(v int32) *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger {
	s.Gender = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) SetLastName(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger {
	s.LastName = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) SetMobileCountryCode(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger {
	s.MobileCountryCode = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) SetMobilePhoneNumber(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger {
	s.MobilePhoneNumber = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) SetNationality(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger {
	s.Nationality = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) SetType(v int32) *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger {
	s.Type = &v
	return s
}

type OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential struct {
	// place of issue, two-letter code
	//
	// example:
	//
	// CN
	CertIssuePlace *string `json:"cert_issue_place,omitempty" xml:"cert_issue_place,omitempty"`
	// travel document number
	//
	// example:
	//
	// E1***5674
	CredentialNum *string `json:"credential_num,omitempty" xml:"credential_num,omitempty"`
	// travel document type , only support "1"(1 means passport) currently.
	//
	// example:
	//
	// 1
	CredentialType *int32 `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	// expiration date
	//
	// example:
	//
	// 20290101
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential) SetCertIssuePlace(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential {
	s.CertIssuePlace = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential) SetCredentialNum(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential {
	s.CredentialNum = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential) SetCredentialType(v int32) *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential {
	s.CredentialType = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential) SetExpireDate(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential {
	s.ExpireDate = &v
	return s
}

type OrderDetailResponseBodyDataFlightItemDetailList struct {
	// b PNR(airline booking code) list
	BPnrList []*string `json:"b_pnr_list,omitempty" xml:"b_pnr_list,omitempty" type:"Repeated"`
	// c PNR(airline booking code) list
	CPnrList []*string `json:"c_pnr_list,omitempty" xml:"c_pnr_list,omitempty" type:"Repeated"`
	// flight price information for current passenger
	FlightPrice *OrderDetailResponseBodyDataFlightItemDetailListFlightPrice `json:"flight_price,omitempty" xml:"flight_price,omitempty" type:"Struct"`
	// RBD information in flight segment dimension
	FlightSegmentCabinRelation []*OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation `json:"flight_segment_cabin_relation,omitempty" xml:"flight_segment_cabin_relation,omitempty" type:"Repeated"`
	// passenger
	Passenger *OrderDetailResponseBodyDataFlightItemDetailListPassenger `json:"passenger,omitempty" xml:"passenger,omitempty" type:"Struct"`
	// ticketing airline
	//
	// example:
	//
	// HO
	TicketAirLine *string `json:"ticket_air_line,omitempty" xml:"ticket_air_line,omitempty"`
	// ticket number list
	TicketNos []*string `json:"ticket_nos,omitempty" xml:"ticket_nos,omitempty" type:"Repeated"`
}

func (s OrderDetailResponseBodyDataFlightItemDetailList) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBodyDataFlightItemDetailList) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) SetBPnrList(v []*string) *OrderDetailResponseBodyDataFlightItemDetailList {
	s.BPnrList = v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) SetCPnrList(v []*string) *OrderDetailResponseBodyDataFlightItemDetailList {
	s.CPnrList = v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) SetFlightPrice(v *OrderDetailResponseBodyDataFlightItemDetailListFlightPrice) *OrderDetailResponseBodyDataFlightItemDetailList {
	s.FlightPrice = v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) SetFlightSegmentCabinRelation(v []*OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) *OrderDetailResponseBodyDataFlightItemDetailList {
	s.FlightSegmentCabinRelation = v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) SetPassenger(v *OrderDetailResponseBodyDataFlightItemDetailListPassenger) *OrderDetailResponseBodyDataFlightItemDetailList {
	s.Passenger = v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) SetTicketAirLine(v string) *OrderDetailResponseBodyDataFlightItemDetailList {
	s.TicketAirLine = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) SetTicketNos(v []*string) *OrderDetailResponseBodyDataFlightItemDetailList {
	s.TicketNos = v
	return s
}

type OrderDetailResponseBodyDataFlightItemDetailListFlightPrice struct {
	// selling price
	//
	// example:
	//
	// 300
	SellPrice *float64 `json:"sell_price,omitempty" xml:"sell_price,omitempty"`
	// tax
	//
	// example:
	//
	// 10
	Tax *float64 `json:"tax,omitempty" xml:"tax,omitempty"`
}

func (s OrderDetailResponseBodyDataFlightItemDetailListFlightPrice) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBodyDataFlightItemDetailListFlightPrice) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightPrice) SetSellPrice(v float64) *OrderDetailResponseBodyDataFlightItemDetailListFlightPrice {
	s.SellPrice = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightPrice) SetTax(v float64) *OrderDetailResponseBodyDataFlightItemDetailListFlightPrice {
	s.Tax = &v
	return s
}

type OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation struct {
	// RBD
	//
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// cabin class
	//
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// cabin class description
	//
	// example:
	//
	// economy class
	CabinClassName *string `json:"cabin_class_name,omitempty" xml:"cabin_class_name,omitempty"`
	// availability
	//
	// example:
	//
	// A
	CabinQuantity *string `json:"cabin_quantity,omitempty" xml:"cabin_quantity,omitempty"`
	// segment ID format: flight no.+departure airport[IATA airport code]+arrival airport[IATA airport code]+departure time(yyMMddHHmm)
	//
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
}

func (s OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) SetCabin(v string) *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation {
	s.Cabin = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) SetCabinClass(v string) *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation {
	s.CabinClass = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) SetCabinClassName(v string) *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation {
	s.CabinClassName = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) SetCabinQuantity(v string) *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation {
	s.CabinQuantity = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) SetSegmentId(v string) *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation {
	s.SegmentId = &v
	return s
}

type OrderDetailResponseBodyDataFlightItemDetailListPassenger struct {
	// date of birth (yyyyMMdd)
	//
	// example:
	//
	// 20020301
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// travel document
	Credential *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential `json:"credential,omitempty" xml:"credential,omitempty" type:"Struct"`
	// first name
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// gender 0: MALE; 1: FEMALE
	//
	// example:
	//
	// 0
	Gender *int32 `json:"gender,omitempty" xml:"gender,omitempty"`
	// last name
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// mobile country code
	//
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// mobile phone number
	//
	// example:
	//
	// 183******96
	MobilePhoneNumber *string `json:"mobile_phone_number,omitempty" xml:"mobile_phone_number,omitempty"`
	// nationality (two-letter code)
	//
	// example:
	//
	// CN
	Nationality *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// passenger type 0: adult; 1: child; 8: infant
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s OrderDetailResponseBodyDataFlightItemDetailListPassenger) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBodyDataFlightItemDetailListPassenger) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) SetBirthday(v string) *OrderDetailResponseBodyDataFlightItemDetailListPassenger {
	s.Birthday = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) SetCredential(v *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential) *OrderDetailResponseBodyDataFlightItemDetailListPassenger {
	s.Credential = v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) SetFirstName(v string) *OrderDetailResponseBodyDataFlightItemDetailListPassenger {
	s.FirstName = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) SetGender(v int32) *OrderDetailResponseBodyDataFlightItemDetailListPassenger {
	s.Gender = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) SetLastName(v string) *OrderDetailResponseBodyDataFlightItemDetailListPassenger {
	s.LastName = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) SetMobileCountryCode(v string) *OrderDetailResponseBodyDataFlightItemDetailListPassenger {
	s.MobileCountryCode = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) SetMobilePhoneNumber(v string) *OrderDetailResponseBodyDataFlightItemDetailListPassenger {
	s.MobilePhoneNumber = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) SetNationality(v string) *OrderDetailResponseBodyDataFlightItemDetailListPassenger {
	s.Nationality = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) SetType(v int32) *OrderDetailResponseBodyDataFlightItemDetailListPassenger {
	s.Type = &v
	return s
}

type OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential struct {
	// place of issue, two-letter code
	//
	// example:
	//
	// CN
	CertIssuePlace *string `json:"cert_issue_place,omitempty" xml:"cert_issue_place,omitempty"`
	// travel document number
	//
	// example:
	//
	// E1***5674
	CredentialNum *string `json:"credential_num,omitempty" xml:"credential_num,omitempty"`
	// travel document type , only support "1"(1 means passport) currently.
	//
	// example:
	//
	// 1
	CredentialType *int32 `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	// expiration date
	//
	// example:
	//
	// 20290101
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
}

func (s OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential) SetCertIssuePlace(v string) *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential {
	s.CertIssuePlace = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential) SetCredentialNum(v string) *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential {
	s.CredentialNum = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential) SetCredentialType(v int32) *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential {
	s.CredentialType = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential) SetExpireDate(v string) *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential {
	s.ExpireDate = &v
	return s
}

type OrderDetailResponseBodyDataPassengerList struct {
	// date of birth (yyyyMMdd)
	//
	// example:
	//
	// 20020301
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// travel document
	Credential *OrderDetailResponseBodyDataPassengerListCredential `json:"credential,omitempty" xml:"credential,omitempty" type:"Struct"`
	// first name
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// gender 0: MALE; 1: FEMALE
	//
	// example:
	//
	// 0
	Gender *int32 `json:"gender,omitempty" xml:"gender,omitempty"`
	// last name
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// mobile country code
	//
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// mobile phone number
	//
	// example:
	//
	// 183******96
	MobilePhoneNumber *string `json:"mobile_phone_number,omitempty" xml:"mobile_phone_number,omitempty"`
	// nationality (two-letter code)
	//
	// example:
	//
	// CN
	Nationality *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// passenger type 0: adult; 1: child; 8: infant
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s OrderDetailResponseBodyDataPassengerList) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBodyDataPassengerList) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataPassengerList) SetBirthday(v string) *OrderDetailResponseBodyDataPassengerList {
	s.Birthday = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerList) SetCredential(v *OrderDetailResponseBodyDataPassengerListCredential) *OrderDetailResponseBodyDataPassengerList {
	s.Credential = v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerList) SetFirstName(v string) *OrderDetailResponseBodyDataPassengerList {
	s.FirstName = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerList) SetGender(v int32) *OrderDetailResponseBodyDataPassengerList {
	s.Gender = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerList) SetLastName(v string) *OrderDetailResponseBodyDataPassengerList {
	s.LastName = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerList) SetMobileCountryCode(v string) *OrderDetailResponseBodyDataPassengerList {
	s.MobileCountryCode = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerList) SetMobilePhoneNumber(v string) *OrderDetailResponseBodyDataPassengerList {
	s.MobilePhoneNumber = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerList) SetNationality(v string) *OrderDetailResponseBodyDataPassengerList {
	s.Nationality = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerList) SetType(v int32) *OrderDetailResponseBodyDataPassengerList {
	s.Type = &v
	return s
}

type OrderDetailResponseBodyDataPassengerListCredential struct {
	// place of issue, two-letter code
	//
	// example:
	//
	// CN
	CertIssuePlace *string `json:"cert_issue_place,omitempty" xml:"cert_issue_place,omitempty"`
	// travel document number
	//
	// example:
	//
	// E1***5674
	CredentialNum *string `json:"credential_num,omitempty" xml:"credential_num,omitempty"`
	// travel document type , only support "1"(1 means passport) currently.
	//
	// example:
	//
	// 1
	CredentialType *int32 `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	// expiration date
	//
	// example:
	//
	// 20290101
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
}

func (s OrderDetailResponseBodyDataPassengerListCredential) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBodyDataPassengerListCredential) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataPassengerListCredential) SetCertIssuePlace(v string) *OrderDetailResponseBodyDataPassengerListCredential {
	s.CertIssuePlace = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerListCredential) SetCredentialNum(v string) *OrderDetailResponseBodyDataPassengerListCredential {
	s.CredentialNum = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerListCredential) SetCredentialType(v int32) *OrderDetailResponseBodyDataPassengerListCredential {
	s.CredentialType = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerListCredential) SetExpireDate(v string) *OrderDetailResponseBodyDataPassengerListCredential {
	s.ExpireDate = &v
	return s
}

type OrderDetailResponseBodyDataSolution struct {
	// adult fare
	//
	// example:
	//
	// 300
	AdultPrice *float64 `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	// adult tax
	//
	// example:
	//
	// 30
	AdultTax *float64 `json:"adult_tax,omitempty" xml:"adult_tax,omitempty"`
	// child fare
	//
	// example:
	//
	// 200
	ChildPrice *float64 `json:"child_price,omitempty" xml:"child_price,omitempty"`
	// child tax
	//
	// example:
	//
	// 20
	ChildTax *float64 `json:"child_tax,omitempty" xml:"child_tax,omitempty"`
	// infant fare
	//
	// example:
	//
	// 100
	InfantPrice *float64 `json:"infant_price,omitempty" xml:"infant_price,omitempty"`
	// infant tax
	//
	// example:
	//
	// 10
	InfantTax *float64 `json:"infant_tax,omitempty" xml:"infant_tax,omitempty"`
	// journey list
	JourneyList []*OrderDetailResponseBodyDataSolutionJourneyList `json:"journey_list,omitempty" xml:"journey_list,omitempty" type:"Repeated"`
	// product type description
	//
	// example:
	//
	// description
	ProductTypeDescription *string `json:"product_type_description,omitempty" xml:"product_type_description,omitempty"`
	// refund coupon description
	//
	// example:
	//
	// description
	RefundTicketCouponDescription *string `json:"refund_ticket_coupon_description,omitempty" xml:"refund_ticket_coupon_description,omitempty"`
	// through check-in baggage policy
	SegmentBaggageCheckInInfoList []*OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList `json:"segment_baggage_check_in_info_list,omitempty" xml:"segment_baggage_check_in_info_list,omitempty" type:"Repeated"`
	// baggage rule list
	SegmentBaggageMappingList []*OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList `json:"segment_baggage_mapping_list,omitempty" xml:"segment_baggage_mapping_list,omitempty" type:"Repeated"`
	// change and refund policy
	SegmentRefundChangeRuleMappingList []*OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList `json:"segment_refund_change_rule_mapping_list,omitempty" xml:"segment_refund_change_rule_mapping_list,omitempty" type:"Repeated"`
	// solution_id
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TUxxx
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s OrderDetailResponseBodyDataSolution) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBodyDataSolution) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataSolution) SetAdultPrice(v float64) *OrderDetailResponseBodyDataSolution {
	s.AdultPrice = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetAdultTax(v float64) *OrderDetailResponseBodyDataSolution {
	s.AdultTax = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetChildPrice(v float64) *OrderDetailResponseBodyDataSolution {
	s.ChildPrice = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetChildTax(v float64) *OrderDetailResponseBodyDataSolution {
	s.ChildTax = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetInfantPrice(v float64) *OrderDetailResponseBodyDataSolution {
	s.InfantPrice = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetInfantTax(v float64) *OrderDetailResponseBodyDataSolution {
	s.InfantTax = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetJourneyList(v []*OrderDetailResponseBodyDataSolutionJourneyList) *OrderDetailResponseBodyDataSolution {
	s.JourneyList = v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetProductTypeDescription(v string) *OrderDetailResponseBodyDataSolution {
	s.ProductTypeDescription = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetRefundTicketCouponDescription(v string) *OrderDetailResponseBodyDataSolution {
	s.RefundTicketCouponDescription = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetSegmentBaggageCheckInInfoList(v []*OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList) *OrderDetailResponseBodyDataSolution {
	s.SegmentBaggageCheckInInfoList = v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetSegmentBaggageMappingList(v []*OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList) *OrderDetailResponseBodyDataSolution {
	s.SegmentBaggageMappingList = v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetSegmentRefundChangeRuleMappingList(v []*OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) *OrderDetailResponseBodyDataSolution {
	s.SegmentRefundChangeRuleMappingList = v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetSolutionId(v string) *OrderDetailResponseBodyDataSolution {
	s.SolutionId = &v
	return s
}

type OrderDetailResponseBodyDataSolutionJourneyList struct {
	// segment list
	SegmentList []*OrderDetailResponseBodyDataSolutionJourneyListSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// number of transfers
	//
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s OrderDetailResponseBodyDataSolutionJourneyList) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBodyDataSolutionJourneyList) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataSolutionJourneyList) SetSegmentList(v []*OrderDetailResponseBodyDataSolutionJourneyListSegmentList) *OrderDetailResponseBodyDataSolutionJourneyList {
	s.SegmentList = v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyList) SetTransferCount(v int32) *OrderDetailResponseBodyDataSolutionJourneyList {
	s.TransferCount = &v
	return s
}

type OrderDetailResponseBodyDataSolutionJourneyListSegmentList struct {
	// arrival airport code (capitalized)
	//
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// arrival city code (capitalized)
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// arrival terminal
	//
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// arrival time in string format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 10:40:00
	ArrivalTime *string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// availability
	//
	// example:
	//
	// 7
	Availability *string `json:"availability,omitempty" xml:"availability,omitempty"`
	// RBD
	//
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// cabin class
	//
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// code share or not
	//
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// departure airport code (capitalized)
	//
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// departure city code (capitalized)
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// departure terminal
	//
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// departure time in string format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 07:55:00
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// equipment type
	//
	// example:
	//
	// 32Q
	EquipType *string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// flight time, unit: minute
	//
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// marketing airline code (eg: KA)
	//
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// marketing airline flight no. (eg: KA5809)
	//
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// marketing airline flight no. (eg: 5809)
	//
	// example:
	//
	// 1295
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// operating airline code (eg: CX)
	//
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// operating airline flight no. (eg: CX601)
	//
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// segment ID format: flight no.+departure airport[IATA airport code]+arrival airport[IATA airport code]+departure time(MMdd)
	//
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// stop city list when stop_quantity > 1 , use “,” for seperation
	//
	// example:
	//
	// SEL,HKG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// number of stops
	//
	// example:
	//
	// 0
	StopQuantity *int32 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s OrderDetailResponseBodyDataSolutionJourneyListSegmentList) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetArrivalAirport(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetArrivalCity(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetArrivalTerminal(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetArrivalTime(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetAvailability(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.Availability = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetCabin(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.Cabin = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetCabinClass(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.CabinClass = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetCodeShare(v bool) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.CodeShare = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetDepartureAirport(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetDepartureCity(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetDepartureTerminal(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetDepartureTime(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetEquipType(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.EquipType = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetFlightDuration(v int32) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetMarketingAirline(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetMarketingFlightNo(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetMarketingFlightNoInt(v int32) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetOperatingAirline(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetOperatingFlightNo(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetSegmentId(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.SegmentId = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetStopCityList(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.StopCityList = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetStopQuantity(v int32) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.StopQuantity = &v
	return s
}

type OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList struct {
	// through check-in baggage policy type
	//
	// 1. baggage through check-in between segments
	//
	// 2. baggage re-check-in needed between segments
	//
	// 3. baggage through check-in at stop city ( applies for stop flight )
	//
	// 4. baggage re-checkin needed at stop city ( applies for stop flight )
	//
	// example:
	//
	// 1
	LuggageDirectInfoType *int32 `json:"luggage_direct_info_type,omitempty" xml:"luggage_direct_info_type,omitempty"`
	// segment id list. all the listed segment ids share the same baggage through check-in policy
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList) SetLuggageDirectInfoType(v int32) *OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList {
	s.LuggageDirectInfoType = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList) SetSegmentIdList(v []*string) *OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList {
	s.SegmentIdList = v
	return s
}

type OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList struct {
	// baggage rule mapping, key is passenger type, value is baggage allowance details
	PassengerBaggageAllowanceMapping map[string]*DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue `json:"passenger_baggage_allowance_mapping,omitempty" xml:"passenger_baggage_allowance_mapping,omitempty"`
	// segment id list. all the listed segment ids share the same baggage rule
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList) SetPassengerBaggageAllowanceMapping(v map[string]*DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) *OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList {
	s.PassengerBaggageAllowanceMapping = v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList) SetSegmentIdList(v []*string) *OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList {
	s.SegmentIdList = v
	return s
}

type OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList struct {
	// change and refund policy mapping, key is passenger type, value is change and refund policy details
	RefundChangeRuleMap map[string]*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue `json:"refund_change_rule_map,omitempty" xml:"refund_change_rule_map,omitempty"`
	// segment id list. all the listed segment ids share the same change and refund policy
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) SetRefundChangeRuleMap(v map[string]*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) *OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList {
	s.RefundChangeRuleMap = v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) SetSegmentIdList(v []*string) *OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList {
	s.SegmentIdList = v
	return s
}

type OrderDetailResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OrderDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s OrderDetailResponse) GoString() string {
	return s.String()
}

func (s *OrderDetailResponse) SetHeaders(v map[string]*string) *OrderDetailResponse {
	s.Headers = v
	return s
}

func (s *OrderDetailResponse) SetStatusCode(v int32) *OrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *OrderDetailResponse) SetBody(v *OrderDetailResponseBody) *OrderDetailResponse {
	s.Body = v
	return s
}

type OrderListHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// language code(refer to ISO_639), defaults to the buyer\\"s account configuration
	//
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s OrderListHeaders) String() string {
	return tea.Prettify(s)
}

func (s OrderListHeaders) GoString() string {
	return s.String()
}

func (s *OrderListHeaders) SetCommonHeaders(v map[string]*string) *OrderListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OrderListHeaders) SetXAcsAirticketAccessToken(v string) *OrderListHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *OrderListHeaders) SetXAcsAirticketLanguage(v string) *OrderListHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type OrderListRequest struct {
	// latest booking time (timestamp)
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-02 11:20:00
	BookTimeEnd *int64 `json:"book_time_end,omitempty" xml:"book_time_end,omitempty"`
	// earliest book time(timestamp)
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-01 11:20:00
	BookTimeStart *int64 `json:"book_time_start,omitempty" xml:"book_time_start,omitempty"`
	// pagination query parameters, from which page to start querying
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// pagination query parameters, how many orders to return
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// which order status will be query
	//
	// 1: order reservation in process
	//
	// 2: order reservation successful
	//
	// 3: order paid
	//
	// 4: order successful
	//
	// 5: order closed
	//
	// example:
	//
	// 4
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s OrderListRequest) String() string {
	return tea.Prettify(s)
}

func (s OrderListRequest) GoString() string {
	return s.String()
}

func (s *OrderListRequest) SetBookTimeEnd(v int64) *OrderListRequest {
	s.BookTimeEnd = &v
	return s
}

func (s *OrderListRequest) SetBookTimeStart(v int64) *OrderListRequest {
	s.BookTimeStart = &v
	return s
}

func (s *OrderListRequest) SetPageIndex(v int32) *OrderListRequest {
	s.PageIndex = &v
	return s
}

func (s *OrderListRequest) SetPageSize(v int32) *OrderListRequest {
	s.PageSize = &v
	return s
}

func (s *OrderListRequest) SetStatus(v int32) *OrderListRequest {
	s.Status = &v
	return s
}

type OrderListResponseBody struct {
	// request ID
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *OrderListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error data
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// http reqeust has been processed successfully，status code is 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true represents success, false represents failure
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OrderListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OrderListResponseBody) GoString() string {
	return s.String()
}

func (s *OrderListResponseBody) SetRequestId(v string) *OrderListResponseBody {
	s.RequestId = &v
	return s
}

func (s *OrderListResponseBody) SetData(v *OrderListResponseBodyData) *OrderListResponseBody {
	s.Data = v
	return s
}

func (s *OrderListResponseBody) SetErrorCode(v string) *OrderListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *OrderListResponseBody) SetErrorData(v interface{}) *OrderListResponseBody {
	s.ErrorData = v
	return s
}

func (s *OrderListResponseBody) SetErrorMsg(v string) *OrderListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *OrderListResponseBody) SetStatus(v int32) *OrderListResponseBody {
	s.Status = &v
	return s
}

func (s *OrderListResponseBody) SetSuccess(v bool) *OrderListResponseBody {
	s.Success = &v
	return s
}

type OrderListResponseBodyData struct {
	// order list
	List []*OrderListResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// information of pagination
	Pagination *OrderListResponseBodyDataPagination `json:"pagination,omitempty" xml:"pagination,omitempty" type:"Struct"`
}

func (s OrderListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s OrderListResponseBodyData) GoString() string {
	return s.String()
}

func (s *OrderListResponseBodyData) SetList(v []*OrderListResponseBodyDataList) *OrderListResponseBodyData {
	s.List = v
	return s
}

func (s *OrderListResponseBodyData) SetPagination(v *OrderListResponseBodyDataPagination) *OrderListResponseBodyData {
	s.Pagination = v
	return s
}

type OrderListResponseBodyDataList struct {
	// book time(timestamp)
	//
	// example:
	//
	// 1677210784000
	BookTime *int64 `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// order number created by book
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// order status
	//
	// 1: order reservation in process
	//
	// 2: order reservation successful
	//
	// 3: order paid
	//
	// 4: order successful
	//
	// 5: order closed
	//
	// example:
	//
	// 4
	OrderStatus *string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// external order number(customized by buyer when book)
	//
	// example:
	//
	// x091-2023-0220-j-0001
	OutOrderNum *string `json:"out_order_num,omitempty" xml:"out_order_num,omitempty"`
	// the information about all passenger of current order
	PassengerList []*OrderListResponseBodyDataListPassengerList `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
	// payment status
	//
	// 1: payment in process
	//
	// 2: deduction successful
	//
	// 3: paid to the seller
	//
	// 4: transaction closed
	//
	// example:
	//
	// 2
	PayStatus *string `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// pay time(timestamp)
	//
	// example:
	//
	// 1677210788000
	PayTime *int64 `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// discount amount
	//
	// example:
	//
	// 10
	PromotionPrice *float64 `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// actual payment amount
	//
	// example:
	//
	// 3000
	RealPayPrice *float64 `json:"real_pay_price,omitempty" xml:"real_pay_price,omitempty"`
	// buyer nickname
	//
	// example:
	//
	// nick
	SessionNick *string `json:"session_nick,omitempty" xml:"session_nick,omitempty"`
	// order success time(timestamp)
	//
	// example:
	//
	// 1677210786000
	SucceedTime *int64 `json:"succeed_time,omitempty" xml:"succeed_time,omitempty"`
	// total price of current order
	//
	// example:
	//
	// 3000
	TotalPrice *float64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// transaction number
	//
	// example:
	//
	// hkduendkd-2023-dj0
	TransactionNo *string `json:"transaction_no,omitempty" xml:"transaction_no,omitempty"`
}

func (s OrderListResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s OrderListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *OrderListResponseBodyDataList) SetBookTime(v int64) *OrderListResponseBodyDataList {
	s.BookTime = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetOrderNum(v int64) *OrderListResponseBodyDataList {
	s.OrderNum = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetOrderStatus(v string) *OrderListResponseBodyDataList {
	s.OrderStatus = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetOutOrderNum(v string) *OrderListResponseBodyDataList {
	s.OutOrderNum = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetPassengerList(v []*OrderListResponseBodyDataListPassengerList) *OrderListResponseBodyDataList {
	s.PassengerList = v
	return s
}

func (s *OrderListResponseBodyDataList) SetPayStatus(v string) *OrderListResponseBodyDataList {
	s.PayStatus = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetPayTime(v int64) *OrderListResponseBodyDataList {
	s.PayTime = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetPromotionPrice(v float64) *OrderListResponseBodyDataList {
	s.PromotionPrice = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetRealPayPrice(v float64) *OrderListResponseBodyDataList {
	s.RealPayPrice = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetSessionNick(v string) *OrderListResponseBodyDataList {
	s.SessionNick = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetSucceedTime(v int64) *OrderListResponseBodyDataList {
	s.SucceedTime = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetTotalPrice(v float64) *OrderListResponseBodyDataList {
	s.TotalPrice = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetTransactionNo(v string) *OrderListResponseBodyDataList {
	s.TransactionNo = &v
	return s
}

type OrderListResponseBodyDataListPassengerList struct {
	// date of birth (yyyyMMdd)
	//
	// example:
	//
	// 20020301
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// credential
	Credential *OrderListResponseBodyDataListPassengerListCredential `json:"credential,omitempty" xml:"credential,omitempty" type:"Struct"`
	// first name
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// gender 0: MALE; 1: FEMALE
	//
	// example:
	//
	// 0
	Gender *int32 `json:"gender,omitempty" xml:"gender,omitempty"`
	// last name
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// mobile country code
	//
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// mobile phone number
	//
	// example:
	//
	// 183******96
	MobilePhoneNumber *string `json:"mobile_phone_number,omitempty" xml:"mobile_phone_number,omitempty"`
	// nationality (two-letter code)
	//
	// example:
	//
	// CN
	Nationality *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// passenger type 0: adult; 1: child; 8: infant
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s OrderListResponseBodyDataListPassengerList) String() string {
	return tea.Prettify(s)
}

func (s OrderListResponseBodyDataListPassengerList) GoString() string {
	return s.String()
}

func (s *OrderListResponseBodyDataListPassengerList) SetBirthday(v string) *OrderListResponseBodyDataListPassengerList {
	s.Birthday = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerList) SetCredential(v *OrderListResponseBodyDataListPassengerListCredential) *OrderListResponseBodyDataListPassengerList {
	s.Credential = v
	return s
}

func (s *OrderListResponseBodyDataListPassengerList) SetFirstName(v string) *OrderListResponseBodyDataListPassengerList {
	s.FirstName = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerList) SetGender(v int32) *OrderListResponseBodyDataListPassengerList {
	s.Gender = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerList) SetLastName(v string) *OrderListResponseBodyDataListPassengerList {
	s.LastName = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerList) SetMobileCountryCode(v string) *OrderListResponseBodyDataListPassengerList {
	s.MobileCountryCode = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerList) SetMobilePhoneNumber(v string) *OrderListResponseBodyDataListPassengerList {
	s.MobilePhoneNumber = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerList) SetNationality(v string) *OrderListResponseBodyDataListPassengerList {
	s.Nationality = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerList) SetType(v int32) *OrderListResponseBodyDataListPassengerList {
	s.Type = &v
	return s
}

type OrderListResponseBodyDataListPassengerListCredential struct {
	// issuing place (two-letter code)
	//
	// example:
	//
	// CN
	CertIssuePlace *string `json:"cert_issue_place,omitempty" xml:"cert_issue_place,omitempty"`
	// 
	//
	// example:
	//
	// E1***5674
	CredentialNum *string `json:"credential_num,omitempty" xml:"credential_num,omitempty"`
	// credential type , only support "1"(1 means passport) currently.
	//
	// example:
	//
	// 1
	CredentialType *int32 `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	// credential expiration date
	//
	// example:
	//
	// 20290101
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
}

func (s OrderListResponseBodyDataListPassengerListCredential) String() string {
	return tea.Prettify(s)
}

func (s OrderListResponseBodyDataListPassengerListCredential) GoString() string {
	return s.String()
}

func (s *OrderListResponseBodyDataListPassengerListCredential) SetCertIssuePlace(v string) *OrderListResponseBodyDataListPassengerListCredential {
	s.CertIssuePlace = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerListCredential) SetCredentialNum(v string) *OrderListResponseBodyDataListPassengerListCredential {
	s.CredentialNum = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerListCredential) SetCredentialType(v int32) *OrderListResponseBodyDataListPassengerListCredential {
	s.CredentialType = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerListCredential) SetExpireDate(v string) *OrderListResponseBodyDataListPassengerListCredential {
	s.ExpireDate = &v
	return s
}

type OrderListResponseBodyDataPagination struct {
	// current page index
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// page size
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// the number of total orders
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// the number of total pages
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

func (s OrderListResponseBodyDataPagination) String() string {
	return tea.Prettify(s)
}

func (s OrderListResponseBodyDataPagination) GoString() string {
	return s.String()
}

func (s *OrderListResponseBodyDataPagination) SetCurrentPage(v int32) *OrderListResponseBodyDataPagination {
	s.CurrentPage = &v
	return s
}

func (s *OrderListResponseBodyDataPagination) SetPageSize(v int32) *OrderListResponseBodyDataPagination {
	s.PageSize = &v
	return s
}

func (s *OrderListResponseBodyDataPagination) SetTotalCount(v int32) *OrderListResponseBodyDataPagination {
	s.TotalCount = &v
	return s
}

func (s *OrderListResponseBodyDataPagination) SetTotalPage(v int32) *OrderListResponseBodyDataPagination {
	s.TotalPage = &v
	return s
}

type OrderListResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OrderListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OrderListResponse) String() string {
	return tea.Prettify(s)
}

func (s OrderListResponse) GoString() string {
	return s.String()
}

func (s *OrderListResponse) SetHeaders(v map[string]*string) *OrderListResponse {
	s.Headers = v
	return s
}

func (s *OrderListResponse) SetStatusCode(v int32) *OrderListResponse {
	s.StatusCode = &v
	return s
}

func (s *OrderListResponse) SetBody(v *OrderListResponseBody) *OrderListResponse {
	s.Body = v
	return s
}

type PricingHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// language code(refer to ISO_639), defaults to the buyer\\"s account configuration
	//
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s PricingHeaders) String() string {
	return tea.Prettify(s)
}

func (s PricingHeaders) GoString() string {
	return s.String()
}

func (s *PricingHeaders) SetCommonHeaders(v map[string]*string) *PricingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PricingHeaders) SetXAcsAirticketAccessToken(v string) *PricingHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *PricingHeaders) SetXAcsAirticketLanguage(v string) *PricingHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type PricingRequest struct {
	// solution_id returned by Enrich
	//
	// This parameter is required.
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s PricingRequest) String() string {
	return tea.Prettify(s)
}

func (s PricingRequest) GoString() string {
	return s.String()
}

func (s *PricingRequest) SetSolutionId(v string) *PricingRequest {
	s.SolutionId = &v
	return s
}

type PricingResponseBody struct {
	// request ID
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *PricingResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error data
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// http reqeust has been processed successfully，status code is 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true represents success, false represents failure
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PricingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PricingResponseBody) GoString() string {
	return s.String()
}

func (s *PricingResponseBody) SetRequestId(v string) *PricingResponseBody {
	s.RequestId = &v
	return s
}

func (s *PricingResponseBody) SetData(v *PricingResponseBodyData) *PricingResponseBody {
	s.Data = v
	return s
}

func (s *PricingResponseBody) SetErrorCode(v string) *PricingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PricingResponseBody) SetErrorData(v interface{}) *PricingResponseBody {
	s.ErrorData = v
	return s
}

func (s *PricingResponseBody) SetErrorMsg(v string) *PricingResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *PricingResponseBody) SetStatus(v int32) *PricingResponseBody {
	s.Status = &v
	return s
}

func (s *PricingResponseBody) SetSuccess(v bool) *PricingResponseBody {
	s.Success = &v
	return s
}

type PricingResponseBodyData struct {
	// price information after the price change
	ChangedPriceInfo *PricingResponseBodyDataChangedPriceInfo `json:"changed_price_info,omitempty" xml:"changed_price_info,omitempty" type:"Struct"`
	// whether the price has changed
	//
	// example:
	//
	// true
	IsChanged *bool `json:"is_changed,omitempty" xml:"is_changed,omitempty"`
	// the price information before the change, only available when is_changed = true
	OriginalPriceInfo *PricingResponseBodyDataOriginalPriceInfo `json:"original_price_info,omitempty" xml:"original_price_info,omitempty" type:"Struct"`
	// remaining seats: A indicates more than 9, 0-9 represents the specific number
	//
	// example:
	//
	// A
	RemainSeats *string `json:"remain_seats,omitempty" xml:"remain_seats,omitempty"`
	// the solution represented by the solution_id in request
	Solution *PricingResponseBodyDataSolution `json:"solution,omitempty" xml:"solution,omitempty" type:"Struct"`
}

func (s PricingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PricingResponseBodyData) GoString() string {
	return s.String()
}

func (s *PricingResponseBodyData) SetChangedPriceInfo(v *PricingResponseBodyDataChangedPriceInfo) *PricingResponseBodyData {
	s.ChangedPriceInfo = v
	return s
}

func (s *PricingResponseBodyData) SetIsChanged(v bool) *PricingResponseBodyData {
	s.IsChanged = &v
	return s
}

func (s *PricingResponseBodyData) SetOriginalPriceInfo(v *PricingResponseBodyDataOriginalPriceInfo) *PricingResponseBodyData {
	s.OriginalPriceInfo = v
	return s
}

func (s *PricingResponseBodyData) SetRemainSeats(v string) *PricingResponseBodyData {
	s.RemainSeats = &v
	return s
}

func (s *PricingResponseBodyData) SetSolution(v *PricingResponseBodyDataSolution) *PricingResponseBodyData {
	s.Solution = v
	return s
}

type PricingResponseBodyDataChangedPriceInfo struct {
	// adult fare
	//
	// example:
	//
	// 100
	AdultPrice *float64 `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	// adult tax
	//
	// example:
	//
	// 10
	AdultTax *float64 `json:"adult_tax,omitempty" xml:"adult_tax,omitempty"`
	// child fare
	//
	// example:
	//
	// 100
	ChildPrice *float64 `json:"child_price,omitempty" xml:"child_price,omitempty"`
	// child tax
	//
	// example:
	//
	// 10
	ChildTax *float64 `json:"child_tax,omitempty" xml:"child_tax,omitempty"`
	// infant fare
	//
	// example:
	//
	// 100
	InfantPrice *float64 `json:"infant_price,omitempty" xml:"infant_price,omitempty"`
	// infant tax
	//
	// example:
	//
	// 10
	InfantTax *float64 `json:"infant_tax,omitempty" xml:"infant_tax,omitempty"`
}

func (s PricingResponseBodyDataChangedPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s PricingResponseBodyDataChangedPriceInfo) GoString() string {
	return s.String()
}

func (s *PricingResponseBodyDataChangedPriceInfo) SetAdultPrice(v float64) *PricingResponseBodyDataChangedPriceInfo {
	s.AdultPrice = &v
	return s
}

func (s *PricingResponseBodyDataChangedPriceInfo) SetAdultTax(v float64) *PricingResponseBodyDataChangedPriceInfo {
	s.AdultTax = &v
	return s
}

func (s *PricingResponseBodyDataChangedPriceInfo) SetChildPrice(v float64) *PricingResponseBodyDataChangedPriceInfo {
	s.ChildPrice = &v
	return s
}

func (s *PricingResponseBodyDataChangedPriceInfo) SetChildTax(v float64) *PricingResponseBodyDataChangedPriceInfo {
	s.ChildTax = &v
	return s
}

func (s *PricingResponseBodyDataChangedPriceInfo) SetInfantPrice(v float64) *PricingResponseBodyDataChangedPriceInfo {
	s.InfantPrice = &v
	return s
}

func (s *PricingResponseBodyDataChangedPriceInfo) SetInfantTax(v float64) *PricingResponseBodyDataChangedPriceInfo {
	s.InfantTax = &v
	return s
}

type PricingResponseBodyDataOriginalPriceInfo struct {
	// adult fare
	//
	// example:
	//
	// 200
	AdultPrice *float64 `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	// adult tax
	//
	// example:
	//
	// 20
	AdultTax *float64 `json:"adult_tax,omitempty" xml:"adult_tax,omitempty"`
	// child fare
	//
	// example:
	//
	// 200
	ChildPrice *float64 `json:"child_price,omitempty" xml:"child_price,omitempty"`
	// child tax
	//
	// example:
	//
	// 20
	ChildTax *float64 `json:"child_tax,omitempty" xml:"child_tax,omitempty"`
	// infant fare
	//
	// example:
	//
	// 200
	InfantPrice *float64 `json:"infant_price,omitempty" xml:"infant_price,omitempty"`
	// infant tax
	//
	// example:
	//
	// 20
	InfantTax *float64 `json:"infant_tax,omitempty" xml:"infant_tax,omitempty"`
}

func (s PricingResponseBodyDataOriginalPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s PricingResponseBodyDataOriginalPriceInfo) GoString() string {
	return s.String()
}

func (s *PricingResponseBodyDataOriginalPriceInfo) SetAdultPrice(v float64) *PricingResponseBodyDataOriginalPriceInfo {
	s.AdultPrice = &v
	return s
}

func (s *PricingResponseBodyDataOriginalPriceInfo) SetAdultTax(v float64) *PricingResponseBodyDataOriginalPriceInfo {
	s.AdultTax = &v
	return s
}

func (s *PricingResponseBodyDataOriginalPriceInfo) SetChildPrice(v float64) *PricingResponseBodyDataOriginalPriceInfo {
	s.ChildPrice = &v
	return s
}

func (s *PricingResponseBodyDataOriginalPriceInfo) SetChildTax(v float64) *PricingResponseBodyDataOriginalPriceInfo {
	s.ChildTax = &v
	return s
}

func (s *PricingResponseBodyDataOriginalPriceInfo) SetInfantPrice(v float64) *PricingResponseBodyDataOriginalPriceInfo {
	s.InfantPrice = &v
	return s
}

func (s *PricingResponseBodyDataOriginalPriceInfo) SetInfantTax(v float64) *PricingResponseBodyDataOriginalPriceInfo {
	s.InfantTax = &v
	return s
}

type PricingResponseBodyDataSolution struct {
	// adult fare
	//
	// example:
	//
	// 300
	AdultPrice *float64 `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	// adult tax
	//
	// example:
	//
	// 30
	AdultTax *float64 `json:"adult_tax,omitempty" xml:"adult_tax,omitempty"`
	// child fare
	//
	// example:
	//
	// 200
	ChildPrice *float64 `json:"child_price,omitempty" xml:"child_price,omitempty"`
	// child tax
	//
	// example:
	//
	// 20
	ChildTax *float64 `json:"child_tax,omitempty" xml:"child_tax,omitempty"`
	// infant fare
	//
	// example:
	//
	// 200
	InfantPrice *float64 `json:"infant_price,omitempty" xml:"infant_price,omitempty"`
	// infant tax
	//
	// example:
	//
	// 10
	InfantTax *float64 `json:"infant_tax,omitempty" xml:"infant_tax,omitempty"`
	// journey list
	JourneyList []*PricingResponseBodyDataSolutionJourneyList `json:"journey_list,omitempty" xml:"journey_list,omitempty" type:"Repeated"`
	// product type description
	//
	// example:
	//
	// description
	ProductTypeDescription *string `json:"product_type_description,omitempty" xml:"product_type_description,omitempty"`
	// refund coupon description
	//
	// example:
	//
	// description
	RefundTicketCouponDescription *string `json:"refund_ticket_coupon_description,omitempty" xml:"refund_ticket_coupon_description,omitempty"`
	// through check-in baggage policy
	SegmentBaggageCheckInInfoList []*PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList `json:"segment_baggage_check_in_info_list,omitempty" xml:"segment_baggage_check_in_info_list,omitempty" type:"Repeated"`
	// baggage rule list
	SegmentBaggageMappingList []*PricingResponseBodyDataSolutionSegmentBaggageMappingList `json:"segment_baggage_mapping_list,omitempty" xml:"segment_baggage_mapping_list,omitempty" type:"Repeated"`
	// change and refund policy
	SegmentRefundChangeRuleMappingList []*PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList `json:"segment_refund_change_rule_mapping_list,omitempty" xml:"segment_refund_change_rule_mapping_list,omitempty" type:"Repeated"`
	// solution_id, equals to solution_id in request
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s PricingResponseBodyDataSolution) String() string {
	return tea.Prettify(s)
}

func (s PricingResponseBodyDataSolution) GoString() string {
	return s.String()
}

func (s *PricingResponseBodyDataSolution) SetAdultPrice(v float64) *PricingResponseBodyDataSolution {
	s.AdultPrice = &v
	return s
}

func (s *PricingResponseBodyDataSolution) SetAdultTax(v float64) *PricingResponseBodyDataSolution {
	s.AdultTax = &v
	return s
}

func (s *PricingResponseBodyDataSolution) SetChildPrice(v float64) *PricingResponseBodyDataSolution {
	s.ChildPrice = &v
	return s
}

func (s *PricingResponseBodyDataSolution) SetChildTax(v float64) *PricingResponseBodyDataSolution {
	s.ChildTax = &v
	return s
}

func (s *PricingResponseBodyDataSolution) SetInfantPrice(v float64) *PricingResponseBodyDataSolution {
	s.InfantPrice = &v
	return s
}

func (s *PricingResponseBodyDataSolution) SetInfantTax(v float64) *PricingResponseBodyDataSolution {
	s.InfantTax = &v
	return s
}

func (s *PricingResponseBodyDataSolution) SetJourneyList(v []*PricingResponseBodyDataSolutionJourneyList) *PricingResponseBodyDataSolution {
	s.JourneyList = v
	return s
}

func (s *PricingResponseBodyDataSolution) SetProductTypeDescription(v string) *PricingResponseBodyDataSolution {
	s.ProductTypeDescription = &v
	return s
}

func (s *PricingResponseBodyDataSolution) SetRefundTicketCouponDescription(v string) *PricingResponseBodyDataSolution {
	s.RefundTicketCouponDescription = &v
	return s
}

func (s *PricingResponseBodyDataSolution) SetSegmentBaggageCheckInInfoList(v []*PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList) *PricingResponseBodyDataSolution {
	s.SegmentBaggageCheckInInfoList = v
	return s
}

func (s *PricingResponseBodyDataSolution) SetSegmentBaggageMappingList(v []*PricingResponseBodyDataSolutionSegmentBaggageMappingList) *PricingResponseBodyDataSolution {
	s.SegmentBaggageMappingList = v
	return s
}

func (s *PricingResponseBodyDataSolution) SetSegmentRefundChangeRuleMappingList(v []*PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) *PricingResponseBodyDataSolution {
	s.SegmentRefundChangeRuleMappingList = v
	return s
}

func (s *PricingResponseBodyDataSolution) SetSolutionId(v string) *PricingResponseBodyDataSolution {
	s.SolutionId = &v
	return s
}

type PricingResponseBodyDataSolutionJourneyList struct {
	// segment list
	SegmentList []*PricingResponseBodyDataSolutionJourneyListSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// number of transfers
	//
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s PricingResponseBodyDataSolutionJourneyList) String() string {
	return tea.Prettify(s)
}

func (s PricingResponseBodyDataSolutionJourneyList) GoString() string {
	return s.String()
}

func (s *PricingResponseBodyDataSolutionJourneyList) SetSegmentList(v []*PricingResponseBodyDataSolutionJourneyListSegmentList) *PricingResponseBodyDataSolutionJourneyList {
	s.SegmentList = v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyList) SetTransferCount(v int32) *PricingResponseBodyDataSolutionJourneyList {
	s.TransferCount = &v
	return s
}

type PricingResponseBodyDataSolutionJourneyListSegmentList struct {
	// arrival airport code (capitalized)
	//
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// arrival city code (capitalized)
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// arrival terminal
	//
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// arrival time in string format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 10:40:00
	ArrivalTime *string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// available seats (for reference only)
	//
	// example:
	//
	// 7
	Availability *string `json:"availability,omitempty" xml:"availability,omitempty"`
	// RBD
	//
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// cabin class
	//
	// example:
	//
	// ALL_CABIN
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// code share or not
	//
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// departure airport code (capitalized)
	//
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// departure city code (capitalized)
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// departure terminal
	//
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// departure time in string format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 07:55:00
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// equipment type
	//
	// example:
	//
	// 32Q
	EquipType *string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// flight time, unit: minute
	//
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// marketing airline code (eg: KA)
	//
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// marketing airline flight no. (eg: KA5809)
	//
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// marketing airline flight no. (eg: 5809)
	//
	// example:
	//
	// 1295
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// operating airline code (eg: CX)
	//
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// operating airline flight no. (eg: CX601)
	//
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// segment ID format: flight no.+departure airport[IATA airport code]+arrival airport[IATA airport code]+departure time(MMdd)
	//
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// stop city list. when stop_quantity > 1 , use “,” for seperation
	//
	// example:
	//
	// MFM,PVG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// number of stops
	//
	// example:
	//
	// 0
	StopQuantity *int32 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s PricingResponseBodyDataSolutionJourneyListSegmentList) String() string {
	return tea.Prettify(s)
}

func (s PricingResponseBodyDataSolutionJourneyListSegmentList) GoString() string {
	return s.String()
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetArrivalAirport(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetArrivalCity(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetArrivalTerminal(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetArrivalTime(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetAvailability(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.Availability = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetCabin(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.Cabin = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetCabinClass(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.CabinClass = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetCodeShare(v bool) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.CodeShare = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetDepartureAirport(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetDepartureCity(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetDepartureTerminal(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetDepartureTime(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetEquipType(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.EquipType = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetFlightDuration(v int32) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetMarketingAirline(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetMarketingFlightNo(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetMarketingFlightNoInt(v int32) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetOperatingAirline(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetOperatingFlightNo(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetSegmentId(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.SegmentId = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetStopCityList(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.StopCityList = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetStopQuantity(v int32) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.StopQuantity = &v
	return s
}

type PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList struct {
	// through check-in baggage policy type
	//
	// 1. baggage through check-in between segments
	//
	// 2. baggage re-check-in needed between segments
	//
	// 3. baggage through check-in at stop city ( applies for stop flight )
	//
	// 4. baggage re-checkin needed at stop city ( applies for stop flight )
	//
	// example:
	//
	// 1
	LuggageDirectInfoType *int32 `json:"luggage_direct_info_type,omitempty" xml:"luggage_direct_info_type,omitempty"`
	// segment id list. all the listed segment ids share the same baggage through check-in policy
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList) String() string {
	return tea.Prettify(s)
}

func (s PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList) GoString() string {
	return s.String()
}

func (s *PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList) SetLuggageDirectInfoType(v int32) *PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList {
	s.LuggageDirectInfoType = &v
	return s
}

func (s *PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList) SetSegmentIdList(v []*string) *PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList {
	s.SegmentIdList = v
	return s
}

type PricingResponseBodyDataSolutionSegmentBaggageMappingList struct {
	// baggage rule mapping, key is passenger type, value is baggage allowance details
	PassengerBaggageAllowanceMapping map[string]*DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue `json:"passenger_baggage_allowance_mapping,omitempty" xml:"passenger_baggage_allowance_mapping,omitempty"`
	// segment id list all the listed segment id share the same baggage rule
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s PricingResponseBodyDataSolutionSegmentBaggageMappingList) String() string {
	return tea.Prettify(s)
}

func (s PricingResponseBodyDataSolutionSegmentBaggageMappingList) GoString() string {
	return s.String()
}

func (s *PricingResponseBodyDataSolutionSegmentBaggageMappingList) SetPassengerBaggageAllowanceMapping(v map[string]*DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) *PricingResponseBodyDataSolutionSegmentBaggageMappingList {
	s.PassengerBaggageAllowanceMapping = v
	return s
}

func (s *PricingResponseBodyDataSolutionSegmentBaggageMappingList) SetSegmentIdList(v []*string) *PricingResponseBodyDataSolutionSegmentBaggageMappingList {
	s.SegmentIdList = v
	return s
}

type PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList struct {
	// change and refund policy mapping, key is passenger type, value is change and refund policy detail
	RefundChangeRuleMap map[string]*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue `json:"refund_change_rule_map,omitempty" xml:"refund_change_rule_map,omitempty"`
	// segment id list. all the listed segment ids share the same change and refund policy
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) String() string {
	return tea.Prettify(s)
}

func (s PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) GoString() string {
	return s.String()
}

func (s *PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) SetRefundChangeRuleMap(v map[string]*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) *PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList {
	s.RefundChangeRuleMap = v
	return s
}

func (s *PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) SetSegmentIdList(v []*string) *PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList {
	s.SegmentIdList = v
	return s
}

type PricingResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PricingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PricingResponse) String() string {
	return tea.Prettify(s)
}

func (s PricingResponse) GoString() string {
	return s.String()
}

func (s *PricingResponse) SetHeaders(v map[string]*string) *PricingResponse {
	s.Headers = v
	return s
}

func (s *PricingResponse) SetStatusCode(v int32) *PricingResponse {
	s.StatusCode = &v
	return s
}

func (s *PricingResponse) SetBody(v *PricingResponseBody) *PricingResponse {
	s.Body = v
	return s
}

type RefundApplyHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// language code(refer to ISO_639), defaults to the buyer\\"s account configuration
	//
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s RefundApplyHeaders) String() string {
	return tea.Prettify(s)
}

func (s RefundApplyHeaders) GoString() string {
	return s.String()
}

func (s *RefundApplyHeaders) SetCommonHeaders(v map[string]*string) *RefundApplyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RefundApplyHeaders) SetXAcsAirticketAccessToken(v string) *RefundApplyHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *RefundApplyHeaders) SetXAcsAirticketLanguage(v string) *RefundApplyHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type RefundApplyRequest struct {
	// order number
	//
	// This parameter is required.
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// journeys for which a refund is being requested
	//
	// This parameter is required.
	RefundJourneys []*RefundApplyRequestRefundJourneys `json:"refund_journeys,omitempty" xml:"refund_journeys,omitempty" type:"Repeated"`
	// passengers that applying for a refund
	//
	// This parameter is required.
	RefundPassengerList []*RefundApplyRequestRefundPassengerList `json:"refund_passenger_list,omitempty" xml:"refund_passenger_list,omitempty" type:"Repeated"`
	// refund type and attachments
	//
	// This parameter is required.
	RefundType *RefundApplyRequestRefundType `json:"refund_type,omitempty" xml:"refund_type,omitempty" type:"Struct"`
}

func (s RefundApplyRequest) String() string {
	return tea.Prettify(s)
}

func (s RefundApplyRequest) GoString() string {
	return s.String()
}

func (s *RefundApplyRequest) SetOrderNum(v int64) *RefundApplyRequest {
	s.OrderNum = &v
	return s
}

func (s *RefundApplyRequest) SetRefundJourneys(v []*RefundApplyRequestRefundJourneys) *RefundApplyRequest {
	s.RefundJourneys = v
	return s
}

func (s *RefundApplyRequest) SetRefundPassengerList(v []*RefundApplyRequestRefundPassengerList) *RefundApplyRequest {
	s.RefundPassengerList = v
	return s
}

func (s *RefundApplyRequest) SetRefundType(v *RefundApplyRequestRefundType) *RefundApplyRequest {
	s.RefundType = v
	return s
}

type RefundApplyRequestRefundJourneys struct {
	// segment list
	//
	// This parameter is required.
	SegmentList []*RefundApplyRequestRefundJourneysSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
}

func (s RefundApplyRequestRefundJourneys) String() string {
	return tea.Prettify(s)
}

func (s RefundApplyRequestRefundJourneys) GoString() string {
	return s.String()
}

func (s *RefundApplyRequestRefundJourneys) SetSegmentList(v []*RefundApplyRequestRefundJourneysSegmentList) *RefundApplyRequestRefundJourneys {
	s.SegmentList = v
	return s
}

type RefundApplyRequestRefundJourneysSegmentList struct {
	// arrival airport code (capitalized)
	//
	// This parameter is required.
	//
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// arrival city code (capitalized)
	//
	// This parameter is required.
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// departure airport code (capitalized)
	//
	// This parameter is required.
	//
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// departure city code (capitalized)
	//
	// This parameter is required.
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
}

func (s RefundApplyRequestRefundJourneysSegmentList) String() string {
	return tea.Prettify(s)
}

func (s RefundApplyRequestRefundJourneysSegmentList) GoString() string {
	return s.String()
}

func (s *RefundApplyRequestRefundJourneysSegmentList) SetArrivalAirport(v string) *RefundApplyRequestRefundJourneysSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *RefundApplyRequestRefundJourneysSegmentList) SetArrivalCity(v string) *RefundApplyRequestRefundJourneysSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *RefundApplyRequestRefundJourneysSegmentList) SetDepartureAirport(v string) *RefundApplyRequestRefundJourneysSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *RefundApplyRequestRefundJourneysSegmentList) SetDepartureCity(v string) *RefundApplyRequestRefundJourneysSegmentList {
	s.DepartureCity = &v
	return s
}

type RefundApplyRequestRefundPassengerList struct {
	// credential number
	//
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// first name
	//
	// This parameter is required.
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// last name
	//
	// This parameter is required.
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s RefundApplyRequestRefundPassengerList) String() string {
	return tea.Prettify(s)
}

func (s RefundApplyRequestRefundPassengerList) GoString() string {
	return s.String()
}

func (s *RefundApplyRequestRefundPassengerList) SetDocument(v string) *RefundApplyRequestRefundPassengerList {
	s.Document = &v
	return s
}

func (s *RefundApplyRequestRefundPassengerList) SetFirstName(v string) *RefundApplyRequestRefundPassengerList {
	s.FirstName = &v
	return s
}

func (s *RefundApplyRequestRefundPassengerList) SetLastName(v string) *RefundApplyRequestRefundPassengerList {
	s.LastName = &v
	return s
}

type RefundApplyRequestRefundType struct {
	// attachment file URLs
	//
	// (note: upload the files using a separate file upload interface to get the file URLs)
	//
	// example:
	//
	// [xxx,yyy]
	File []*string `json:"file,omitempty" xml:"file,omitempty" type:"Repeated"`
	// refund type
	//
	// 2: voluntary (I want to change my travel plan/I don\\"t want to fly)
	//
	// 5: involuntary, due to flight delay or cancellation, schedule changes, etc., by the airline
	//
	// 6: involuntary, due to health reasons with a certificate from a hospital of at least secondary level A or above
	//
	//  (note: attachments are not mandatory, but it is recommended to provide them for involuntary refunds as they can increase the success rate)
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	RefundTypeId *int32 `json:"refund_type_id,omitempty" xml:"refund_type_id,omitempty"`
	// remark
	//
	// example:
	//
	// remark desc
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s RefundApplyRequestRefundType) String() string {
	return tea.Prettify(s)
}

func (s RefundApplyRequestRefundType) GoString() string {
	return s.String()
}

func (s *RefundApplyRequestRefundType) SetFile(v []*string) *RefundApplyRequestRefundType {
	s.File = v
	return s
}

func (s *RefundApplyRequestRefundType) SetRefundTypeId(v int32) *RefundApplyRequestRefundType {
	s.RefundTypeId = &v
	return s
}

func (s *RefundApplyRequestRefundType) SetRemark(v string) *RefundApplyRequestRefundType {
	s.Remark = &v
	return s
}

type RefundApplyShrinkRequest struct {
	// order number
	//
	// This parameter is required.
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// journeys for which a refund is being requested
	//
	// This parameter is required.
	RefundJourneysShrink *string `json:"refund_journeys,omitempty" xml:"refund_journeys,omitempty"`
	// passengers that applying for a refund
	//
	// This parameter is required.
	RefundPassengerListShrink *string `json:"refund_passenger_list,omitempty" xml:"refund_passenger_list,omitempty"`
	// refund type and attachments
	//
	// This parameter is required.
	RefundTypeShrink *string `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
}

func (s RefundApplyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RefundApplyShrinkRequest) GoString() string {
	return s.String()
}

func (s *RefundApplyShrinkRequest) SetOrderNum(v int64) *RefundApplyShrinkRequest {
	s.OrderNum = &v
	return s
}

func (s *RefundApplyShrinkRequest) SetRefundJourneysShrink(v string) *RefundApplyShrinkRequest {
	s.RefundJourneysShrink = &v
	return s
}

func (s *RefundApplyShrinkRequest) SetRefundPassengerListShrink(v string) *RefundApplyShrinkRequest {
	s.RefundPassengerListShrink = &v
	return s
}

func (s *RefundApplyShrinkRequest) SetRefundTypeShrink(v string) *RefundApplyShrinkRequest {
	s.RefundTypeShrink = &v
	return s
}

type RefundApplyResponseBody struct {
	// request ID
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *RefundApplyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error data
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// http reqeust has been processed successfully，status code is 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true represents success, false represents failure
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RefundApplyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefundApplyResponseBody) GoString() string {
	return s.String()
}

func (s *RefundApplyResponseBody) SetRequestId(v string) *RefundApplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefundApplyResponseBody) SetData(v *RefundApplyResponseBodyData) *RefundApplyResponseBody {
	s.Data = v
	return s
}

func (s *RefundApplyResponseBody) SetErrorCode(v string) *RefundApplyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RefundApplyResponseBody) SetErrorData(v interface{}) *RefundApplyResponseBody {
	s.ErrorData = v
	return s
}

func (s *RefundApplyResponseBody) SetErrorMsg(v string) *RefundApplyResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RefundApplyResponseBody) SetStatus(v int32) *RefundApplyResponseBody {
	s.Status = &v
	return s
}

func (s *RefundApplyResponseBody) SetSuccess(v bool) *RefundApplyResponseBody {
	s.Success = &v
	return s
}

type RefundApplyResponseBodyData struct {
	// order number
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// refund results
	RefundResults []*RefundApplyResponseBodyDataRefundResults `json:"refund_results,omitempty" xml:"refund_results,omitempty" type:"Repeated"`
}

func (s RefundApplyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RefundApplyResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefundApplyResponseBodyData) SetOrderNum(v int64) *RefundApplyResponseBodyData {
	s.OrderNum = &v
	return s
}

func (s *RefundApplyResponseBodyData) SetRefundResults(v []*RefundApplyResponseBodyDataRefundResults) *RefundApplyResponseBodyData {
	s.RefundResults = v
	return s
}

type RefundApplyResponseBodyDataRefundResults struct {
	// reason for refund application failure
	//
	// example:
	//
	// desc reason
	FailReason *string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	// refund order number
	//
	// example:
	//
	// 4966***617202
	RefundOrderNum *int64 `json:"refund_order_num,omitempty" xml:"refund_order_num,omitempty"`
	// passengers of current refund order
	RefundPassengers []*RefundApplyResponseBodyDataRefundResultsRefundPassengers `json:"refund_passengers,omitempty" xml:"refund_passengers,omitempty" type:"Repeated"`
	// refund order status
	//
	// 0: refund order created successfully
	//
	// 1: refund order creation failed
	//
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s RefundApplyResponseBodyDataRefundResults) String() string {
	return tea.Prettify(s)
}

func (s RefundApplyResponseBodyDataRefundResults) GoString() string {
	return s.String()
}

func (s *RefundApplyResponseBodyDataRefundResults) SetFailReason(v string) *RefundApplyResponseBodyDataRefundResults {
	s.FailReason = &v
	return s
}

func (s *RefundApplyResponseBodyDataRefundResults) SetRefundOrderNum(v int64) *RefundApplyResponseBodyDataRefundResults {
	s.RefundOrderNum = &v
	return s
}

func (s *RefundApplyResponseBodyDataRefundResults) SetRefundPassengers(v []*RefundApplyResponseBodyDataRefundResultsRefundPassengers) *RefundApplyResponseBodyDataRefundResults {
	s.RefundPassengers = v
	return s
}

func (s *RefundApplyResponseBodyDataRefundResults) SetStatus(v int32) *RefundApplyResponseBodyDataRefundResults {
	s.Status = &v
	return s
}

type RefundApplyResponseBodyDataRefundResultsRefundPassengers struct {
	// credential number
	//
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// first name
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// last name
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s RefundApplyResponseBodyDataRefundResultsRefundPassengers) String() string {
	return tea.Prettify(s)
}

func (s RefundApplyResponseBodyDataRefundResultsRefundPassengers) GoString() string {
	return s.String()
}

func (s *RefundApplyResponseBodyDataRefundResultsRefundPassengers) SetDocument(v string) *RefundApplyResponseBodyDataRefundResultsRefundPassengers {
	s.Document = &v
	return s
}

func (s *RefundApplyResponseBodyDataRefundResultsRefundPassengers) SetFirstName(v string) *RefundApplyResponseBodyDataRefundResultsRefundPassengers {
	s.FirstName = &v
	return s
}

func (s *RefundApplyResponseBodyDataRefundResultsRefundPassengers) SetLastName(v string) *RefundApplyResponseBodyDataRefundResultsRefundPassengers {
	s.LastName = &v
	return s
}

type RefundApplyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundApplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefundApplyResponse) String() string {
	return tea.Prettify(s)
}

func (s RefundApplyResponse) GoString() string {
	return s.String()
}

func (s *RefundApplyResponse) SetHeaders(v map[string]*string) *RefundApplyResponse {
	s.Headers = v
	return s
}

func (s *RefundApplyResponse) SetStatusCode(v int32) *RefundApplyResponse {
	s.StatusCode = &v
	return s
}

func (s *RefundApplyResponse) SetBody(v *RefundApplyResponseBody) *RefundApplyResponse {
	s.Body = v
	return s
}

type RefundDetailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// language code(refer to ISO_639), defaults to the buyer\\"s account configuration
	//
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s RefundDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailHeaders) GoString() string {
	return s.String()
}

func (s *RefundDetailHeaders) SetCommonHeaders(v map[string]*string) *RefundDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RefundDetailHeaders) SetXAcsAirticketAccessToken(v string) *RefundDetailHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *RefundDetailHeaders) SetXAcsAirticketLanguage(v string) *RefundDetailHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type RefundDetailRequest struct {
	// refund order number
	//
	// This parameter is required.
	//
	// example:
	//
	// 4966***617732
	RefundOrderNum *int64 `json:"refund_order_num,omitempty" xml:"refund_order_num,omitempty"`
}

func (s RefundDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailRequest) GoString() string {
	return s.String()
}

func (s *RefundDetailRequest) SetRefundOrderNum(v int64) *RefundDetailRequest {
	s.RefundOrderNum = &v
	return s
}

type RefundDetailResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *RefundDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error data
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// http reqeust has been processed successfully，status code is 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true represents success, false represents failure
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RefundDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailResponseBody) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBody) SetRequestId(v string) *RefundDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefundDetailResponseBody) SetData(v *RefundDetailResponseBodyData) *RefundDetailResponseBody {
	s.Data = v
	return s
}

func (s *RefundDetailResponseBody) SetErrorCode(v string) *RefundDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RefundDetailResponseBody) SetErrorData(v interface{}) *RefundDetailResponseBody {
	s.ErrorData = v
	return s
}

func (s *RefundDetailResponseBody) SetErrorMsg(v string) *RefundDetailResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RefundDetailResponseBody) SetStatus(v int32) *RefundDetailResponseBody {
	s.Status = &v
	return s
}

func (s *RefundDetailResponseBody) SetSuccess(v bool) *RefundDetailResponseBody {
	s.Success = &v
	return s
}

type RefundDetailResponseBodyData struct {
	// whether it is a supplementary refund order (if the refund amount is not enough, you can use RefundApply to create a supplementary refund order)
	//
	// example:
	//
	// false
	ContainMultiRefund *bool `json:"contain_multi_refund,omitempty" xml:"contain_multi_refund,omitempty"`
	// supplementary refund orders
	MultiRefundDetails []*RefundDetailResponseBodyDataMultiRefundDetails `json:"multi_refund_details,omitempty" xml:"multi_refund_details,omitempty" type:"Repeated"`
	// order number that returned by Book
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// refund details by passenger dimension
	PassengerRefundDetails []*RefundDetailResponseBodyDataPassengerRefundDetails `json:"passenger_refund_details,omitempty" xml:"passenger_refund_details,omitempty" type:"Repeated"`
	// refund completed time(timestamp)
	//
	// example:
	//
	// 1677229005000
	PaySuccessUtcTime *int64 `json:"pay_success_utc_time,omitempty" xml:"pay_success_utc_time,omitempty"`
	// URLs for refund attachments
	//
	// example:
	//
	// [zzz,yyy]
	RefundAttachmentUrls []*string `json:"refund_attachment_urls,omitempty" xml:"refund_attachment_urls,omitempty" type:"Repeated"`
	// refunded journey
	RefundJourneys []*RefundDetailResponseBodyDataRefundJourneys `json:"refund_journeys,omitempty" xml:"refund_journeys,omitempty" type:"Repeated"`
	// refund order number that returned by RefundApply
	//
	// example:
	//
	// 4966***617654
	RefundOrderNum *int64 `json:"refund_order_num,omitempty" xml:"refund_order_num,omitempty"`
	// reason for refund
	//
	// example:
	//
	// desc reason
	RefundReason *string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// refund type
	//
	// 2: voluntary application
	//
	// 5: flight delay or cancellation, flight schedule change, etc., due to airline reasons
	//
	// 6: health reasons with a certificate from a secondary class A hospital or above
	//
	// 7: non-voluntary confirmed guidance
	//
	// 100: non-voluntary non-confirmed guidance
	//
	// example:
	//
	// 5
	RefundType *int32 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
	// reason for refund rejection
	//
	// example:
	//
	// refuse reason
	RefuseReason *string `json:"refuse_reason,omitempty" xml:"refuse_reason,omitempty"`
	// refund order status
	//
	// 0: refund application
	//
	// 1: refund in progress
	//
	// 2: refund failed
	//
	// 3: refund succeeded
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// transaction number
	//
	// example:
	//
	// 1677229005000
	TransactionNo *string `json:"transaction_no,omitempty" xml:"transaction_no,omitempty"`
	// refund order created time(timestamp)
	//
	// example:
	//
	// 1677229002000
	UtcCreateTime *int64 `json:"utc_create_time,omitempty" xml:"utc_create_time,omitempty"`
}

func (s RefundDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBodyData) SetContainMultiRefund(v bool) *RefundDetailResponseBodyData {
	s.ContainMultiRefund = &v
	return s
}

func (s *RefundDetailResponseBodyData) SetMultiRefundDetails(v []*RefundDetailResponseBodyDataMultiRefundDetails) *RefundDetailResponseBodyData {
	s.MultiRefundDetails = v
	return s
}

func (s *RefundDetailResponseBodyData) SetOrderNum(v int64) *RefundDetailResponseBodyData {
	s.OrderNum = &v
	return s
}

func (s *RefundDetailResponseBodyData) SetPassengerRefundDetails(v []*RefundDetailResponseBodyDataPassengerRefundDetails) *RefundDetailResponseBodyData {
	s.PassengerRefundDetails = v
	return s
}

func (s *RefundDetailResponseBodyData) SetPaySuccessUtcTime(v int64) *RefundDetailResponseBodyData {
	s.PaySuccessUtcTime = &v
	return s
}

func (s *RefundDetailResponseBodyData) SetRefundAttachmentUrls(v []*string) *RefundDetailResponseBodyData {
	s.RefundAttachmentUrls = v
	return s
}

func (s *RefundDetailResponseBodyData) SetRefundJourneys(v []*RefundDetailResponseBodyDataRefundJourneys) *RefundDetailResponseBodyData {
	s.RefundJourneys = v
	return s
}

func (s *RefundDetailResponseBodyData) SetRefundOrderNum(v int64) *RefundDetailResponseBodyData {
	s.RefundOrderNum = &v
	return s
}

func (s *RefundDetailResponseBodyData) SetRefundReason(v string) *RefundDetailResponseBodyData {
	s.RefundReason = &v
	return s
}

func (s *RefundDetailResponseBodyData) SetRefundType(v int32) *RefundDetailResponseBodyData {
	s.RefundType = &v
	return s
}

func (s *RefundDetailResponseBodyData) SetRefuseReason(v string) *RefundDetailResponseBodyData {
	s.RefuseReason = &v
	return s
}

func (s *RefundDetailResponseBodyData) SetStatus(v int32) *RefundDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *RefundDetailResponseBodyData) SetTransactionNo(v string) *RefundDetailResponseBodyData {
	s.TransactionNo = &v
	return s
}

func (s *RefundDetailResponseBodyData) SetUtcCreateTime(v int64) *RefundDetailResponseBodyData {
	s.UtcCreateTime = &v
	return s
}

type RefundDetailResponseBodyDataMultiRefundDetails struct {
	// supplementary refund order number
	//
	// example:
	//
	// 498843***6950
	MultiRefundOrderNum *int64 `json:"multi_refund_order_num,omitempty" xml:"multi_refund_order_num,omitempty"`
	// transaction number of the supplementary refund order
	//
	// example:
	//
	// 498843***6950
	MultiRefundTransactionNo *string `json:"multi_refund_transaction_no,omitempty" xml:"multi_refund_transaction_no,omitempty"`
	// supplementary refund details in passenger dimension
	PassengerMultiRefundDetails []*RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails `json:"passenger_multi_refund_details,omitempty" xml:"passenger_multi_refund_details,omitempty" type:"Repeated"`
}

func (s RefundDetailResponseBodyDataMultiRefundDetails) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailResponseBodyDataMultiRefundDetails) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBodyDataMultiRefundDetails) SetMultiRefundOrderNum(v int64) *RefundDetailResponseBodyDataMultiRefundDetails {
	s.MultiRefundOrderNum = &v
	return s
}

func (s *RefundDetailResponseBodyDataMultiRefundDetails) SetMultiRefundTransactionNo(v string) *RefundDetailResponseBodyDataMultiRefundDetails {
	s.MultiRefundTransactionNo = &v
	return s
}

func (s *RefundDetailResponseBodyDataMultiRefundDetails) SetPassengerMultiRefundDetails(v []*RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails) *RefundDetailResponseBodyDataMultiRefundDetails {
	s.PassengerMultiRefundDetails = v
	return s
}

type RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails struct {
	// amount of the supplementary refund for the change order
	//
	// example:
	//
	// 30
	ChangeOrderRefundFee *float64 `json:"change_order_refund_fee,omitempty" xml:"change_order_refund_fee,omitempty"`
	// amount of the supplementary refund for the original order
	//
	// example:
	//
	// 30
	OriginalOrderRefundFee *float64 `json:"original_order_refund_fee,omitempty" xml:"original_order_refund_fee,omitempty"`
	// passenger for the refund
	Passenger *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger `json:"passenger,omitempty" xml:"passenger,omitempty" type:"Struct"`
}

func (s RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails) SetChangeOrderRefundFee(v float64) *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails {
	s.ChangeOrderRefundFee = &v
	return s
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails) SetOriginalOrderRefundFee(v float64) *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails {
	s.OriginalOrderRefundFee = &v
	return s
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails) SetPassenger(v *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger) *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails {
	s.Passenger = v
	return s
}

type RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger struct {
	// credential number
	//
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// first name
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// last name
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger) SetDocument(v string) *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger {
	s.Document = &v
	return s
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger) SetFirstName(v string) *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger {
	s.FirstName = &v
	return s
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger) SetLastName(v string) *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger {
	s.LastName = &v
	return s
}

type RefundDetailResponseBodyDataPassengerRefundDetails struct {
	// information of the passenger applying for a refund
	Passenger *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger `json:"passenger,omitempty" xml:"passenger,omitempty" type:"Struct"`
	// details of the refund fee
	RefundFee *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee `json:"refund_fee,omitempty" xml:"refund_fee,omitempty" type:"Struct"`
}

func (s RefundDetailResponseBodyDataPassengerRefundDetails) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailResponseBodyDataPassengerRefundDetails) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetails) SetPassenger(v *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger) *RefundDetailResponseBodyDataPassengerRefundDetails {
	s.Passenger = v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetails) SetRefundFee(v *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) *RefundDetailResponseBodyDataPassengerRefundDetails {
	s.RefundFee = v
	return s
}

type RefundDetailResponseBodyDataPassengerRefundDetailsPassenger struct {
	// credential number
	//
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// first name
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// first name
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s RefundDetailResponseBodyDataPassengerRefundDetailsPassenger) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailResponseBodyDataPassengerRefundDetailsPassenger) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger) SetDocument(v string) *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger {
	s.Document = &v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger) SetFirstName(v string) *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger {
	s.FirstName = &v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger) SetLastName(v string) *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger {
	s.LastName = &v
	return s
}

type RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee struct {
	// total price of the used flight tickets
	//
	// example:
	//
	// 30
	AlreadyUsedTotalFee *float64 `json:"already_used_total_fee,omitempty" xml:"already_used_total_fee,omitempty"`
	// amount refunded to the user after change (only available when has change order)
	//
	// example:
	//
	// 30
	ModifyRefundToBuyerMoney *float64 `json:"modify_refund_to_buyer_money,omitempty" xml:"modify_refund_to_buyer_money,omitempty"`
	// non-refundable change service fee
	//
	// example:
	//
	// 30
	NonRefundableChangeServiceFee *float64 `json:"non_refundable_change_service_fee,omitempty" xml:"non_refundable_change_service_fee,omitempty"`
	// non-refundable upgrade fee
	//
	// example:
	//
	// 30
	NonRefundableChangeUpgradeFee *float64 `json:"non_refundable_change_upgrade_fee,omitempty" xml:"non_refundable_change_upgrade_fee,omitempty"`
	// non-refundable tax amount, i.e., tax refund fee
	//
	// example:
	//
	// 30
	NonRefundableTaxFee *float64 `json:"non_refundable_tax_fee,omitempty" xml:"non_refundable_tax_fee,omitempty"`
	// non-refundable ticket amount, i.e., ticket refund fee
	//
	// example:
	//
	// 30
	NonRefundableTicketFee *float64 `json:"non_refundable_ticket_fee,omitempty" xml:"non_refundable_ticket_fee,omitempty"`
	// amount refundable to the user from the original ticket (fare + tax - non_refundable_ticket_fee - non_refundable_tax_fee - already_used_total_fee - discount)
	//
	// example:
	//
	// 30
	RefundToBuyerMoney *float64 `json:"refund_to_buyer_money,omitempty" xml:"refund_to_buyer_money,omitempty"`
}

func (s RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) SetAlreadyUsedTotalFee(v float64) *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee {
	s.AlreadyUsedTotalFee = &v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) SetModifyRefundToBuyerMoney(v float64) *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee {
	s.ModifyRefundToBuyerMoney = &v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) SetNonRefundableChangeServiceFee(v float64) *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee {
	s.NonRefundableChangeServiceFee = &v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) SetNonRefundableChangeUpgradeFee(v float64) *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee {
	s.NonRefundableChangeUpgradeFee = &v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) SetNonRefundableTaxFee(v float64) *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee {
	s.NonRefundableTaxFee = &v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) SetNonRefundableTicketFee(v float64) *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee {
	s.NonRefundableTicketFee = &v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) SetRefundToBuyerMoney(v float64) *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee {
	s.RefundToBuyerMoney = &v
	return s
}

type RefundDetailResponseBodyDataRefundJourneys struct {
	// segment list
	SegmentList []*RefundDetailResponseBodyDataRefundJourneysSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// number of transfer
	//
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s RefundDetailResponseBodyDataRefundJourneys) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailResponseBodyDataRefundJourneys) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBodyDataRefundJourneys) SetSegmentList(v []*RefundDetailResponseBodyDataRefundJourneysSegmentList) *RefundDetailResponseBodyDataRefundJourneys {
	s.SegmentList = v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneys) SetTransferCount(v int32) *RefundDetailResponseBodyDataRefundJourneys {
	s.TransferCount = &v
	return s
}

type RefundDetailResponseBodyDataRefundJourneysSegmentList struct {
	// arrival airport code (capitalized)
	//
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// arrival city code (capitalized)
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// arrival terminal
	//
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// arrival time in string format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 10:40:00
	ArrivalTime *string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// field deprecated
	//
	// example:
	//
	// 7
	Availability *string `json:"availability,omitempty" xml:"availability,omitempty"`
	// RBD
	//
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// cabin class
	//
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// code share or not
	//
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// departure airport code (capitalized)
	//
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// departure city code (capitalized)
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// departure terminal
	//
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// departure time in string format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 07:55:00
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// equipment type
	//
	// example:
	//
	// 32Q
	EquipType *string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// flight time, unit: minute
	//
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// marketing airline code (eg: KA)
	//
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// marketing flight no. (eg: KA5809)
	//
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// marketing flight no. (eg: 5809)
	//
	// example:
	//
	// 1295
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// operating airline code (eg: CX)
	//
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// operating flight no. (eg: CX601)
	//
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// segment ID format: flight no.+departure airport[IATA airport code]+arrival airport[IATA airport code]+departure time(MMdd)
	//
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// stopover city list when stop_quantity > 0 , use “,” for seperation use
	//
	// example:
	//
	// SEL,HKG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// number of stopover
	//
	// example:
	//
	// 0
	StopQuantity *int32 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s RefundDetailResponseBodyDataRefundJourneysSegmentList) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailResponseBodyDataRefundJourneysSegmentList) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetArrivalAirport(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetArrivalCity(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetArrivalTerminal(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetArrivalTime(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetAvailability(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.Availability = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetCabin(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.Cabin = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetCabinClass(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.CabinClass = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetCodeShare(v bool) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.CodeShare = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetDepartureAirport(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetDepartureCity(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetDepartureTerminal(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetDepartureTime(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetEquipType(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.EquipType = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetFlightDuration(v int32) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetMarketingAirline(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetMarketingFlightNo(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetMarketingFlightNoInt(v int32) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetOperatingAirline(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetOperatingFlightNo(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetSegmentId(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.SegmentId = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetStopCityList(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.StopCityList = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetStopQuantity(v int32) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.StopQuantity = &v
	return s
}

type RefundDetailResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefundDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailResponse) GoString() string {
	return s.String()
}

func (s *RefundDetailResponse) SetHeaders(v map[string]*string) *RefundDetailResponse {
	s.Headers = v
	return s
}

func (s *RefundDetailResponse) SetStatusCode(v int32) *RefundDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *RefundDetailResponse) SetBody(v *RefundDetailResponseBody) *RefundDetailResponse {
	s.Body = v
	return s
}

type RefundDetailListHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// language code(refer to ISO_639), defaults to the buyer\\"s account configuration
	//
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s RefundDetailListHeaders) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailListHeaders) GoString() string {
	return s.String()
}

func (s *RefundDetailListHeaders) SetCommonHeaders(v map[string]*string) *RefundDetailListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RefundDetailListHeaders) SetXAcsAirticketAccessToken(v string) *RefundDetailListHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *RefundDetailListHeaders) SetXAcsAirticketLanguage(v string) *RefundDetailListHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type RefundDetailListRequest struct {
	// order number returned by Book
	//
	// example:
	//
	// 49884*****950
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// pagination query parameters, from which page to start querying
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// pagination query parameters, how many orders to return
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// the earliest time(timestamp) of refund order creation
	//
	// This parameter is required.
	//
	// example:
	//
	// 1677229002000
	RefundCreateBeginTime *int64 `json:"refund_create_begin_time,omitempty" xml:"refund_create_begin_time,omitempty"`
	// the latest time(timestamp) of refund order creation
	//
	// This parameter is required.
	//
	// example:
	//
	// 1677229005000
	RefundCreateEndTime *int64 `json:"refund_create_end_time,omitempty" xml:"refund_create_end_time,omitempty"`
}

func (s RefundDetailListRequest) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailListRequest) GoString() string {
	return s.String()
}

func (s *RefundDetailListRequest) SetOrderNum(v int64) *RefundDetailListRequest {
	s.OrderNum = &v
	return s
}

func (s *RefundDetailListRequest) SetPageIndex(v int32) *RefundDetailListRequest {
	s.PageIndex = &v
	return s
}

func (s *RefundDetailListRequest) SetPageSize(v int32) *RefundDetailListRequest {
	s.PageSize = &v
	return s
}

func (s *RefundDetailListRequest) SetRefundCreateBeginTime(v int64) *RefundDetailListRequest {
	s.RefundCreateBeginTime = &v
	return s
}

func (s *RefundDetailListRequest) SetRefundCreateEndTime(v int64) *RefundDetailListRequest {
	s.RefundCreateEndTime = &v
	return s
}

type RefundDetailListResponseBody struct {
	// request ID
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *RefundDetailListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error data
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// http reqeust has been processed successfully，status code is 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true represents success, false represents failure
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RefundDetailListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailListResponseBody) GoString() string {
	return s.String()
}

func (s *RefundDetailListResponseBody) SetRequestId(v string) *RefundDetailListResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefundDetailListResponseBody) SetData(v *RefundDetailListResponseBodyData) *RefundDetailListResponseBody {
	s.Data = v
	return s
}

func (s *RefundDetailListResponseBody) SetErrorCode(v string) *RefundDetailListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RefundDetailListResponseBody) SetErrorData(v interface{}) *RefundDetailListResponseBody {
	s.ErrorData = v
	return s
}

func (s *RefundDetailListResponseBody) SetErrorMsg(v string) *RefundDetailListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RefundDetailListResponseBody) SetStatus(v int32) *RefundDetailListResponseBody {
	s.Status = &v
	return s
}

func (s *RefundDetailListResponseBody) SetSuccess(v bool) *RefundDetailListResponseBody {
	s.Success = &v
	return s
}

type RefundDetailListResponseBodyData struct {
	// refund order list
	List []*RefundDetailListResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// information of pagination
	Pagination *RefundDetailListResponseBodyDataPagination `json:"pagination,omitempty" xml:"pagination,omitempty" type:"Struct"`
}

func (s RefundDetailListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailListResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefundDetailListResponseBodyData) SetList(v []*RefundDetailListResponseBodyDataList) *RefundDetailListResponseBodyData {
	s.List = v
	return s
}

func (s *RefundDetailListResponseBodyData) SetPagination(v *RefundDetailListResponseBodyDataPagination) *RefundDetailListResponseBodyData {
	s.Pagination = v
	return s
}

type RefundDetailListResponseBodyDataList struct {
	// whether it is a supplementary refund order (if the refund amount is not enough, you can use RefundApply to create a supplementary refund order)
	//
	// example:
	//
	// true
	IsMultiRefund *bool `json:"is_multi_refund,omitempty" xml:"is_multi_refund,omitempty"`
	// order number that returned by Book
	//
	// example:
	//
	// 49884*****2345
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// refund order number that returned by RefundApply
	//
	// example:
	//
	// 49884*****950
	RefundOrderNum *int64 `json:"refund_order_num,omitempty" xml:"refund_order_num,omitempty"`
	// refund order status
	//
	// 0: refund application
	//
	// 1: refund in progress
	//
	// 2: refund failed
	//
	// 3: refund succeeded
	//
	// example:
	//
	// 1
	RefundOrderStatus *int32 `json:"refund_order_status,omitempty" xml:"refund_order_status,omitempty"`
	// the original refund order number associated with this supplementary refund order, only avaliable when is_multi_refund=true
	//
	// example:
	//
	// 49884*****2387
	RelatedRefundOrderNum *string `json:"related_refund_order_num,omitempty" xml:"related_refund_order_num,omitempty"`
	// transaction number
	//
	// example:
	//
	// 49884**tde-95za
	TransactionNo *string `json:"transaction_no,omitempty" xml:"transaction_no,omitempty"`
	// refund time(timestamp)
	//
	// example:
	//
	// 1677229002000
	UtcCreateTime *int64 `json:"utc_create_time,omitempty" xml:"utc_create_time,omitempty"`
}

func (s RefundDetailListResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *RefundDetailListResponseBodyDataList) SetIsMultiRefund(v bool) *RefundDetailListResponseBodyDataList {
	s.IsMultiRefund = &v
	return s
}

func (s *RefundDetailListResponseBodyDataList) SetOrderNum(v int64) *RefundDetailListResponseBodyDataList {
	s.OrderNum = &v
	return s
}

func (s *RefundDetailListResponseBodyDataList) SetRefundOrderNum(v int64) *RefundDetailListResponseBodyDataList {
	s.RefundOrderNum = &v
	return s
}

func (s *RefundDetailListResponseBodyDataList) SetRefundOrderStatus(v int32) *RefundDetailListResponseBodyDataList {
	s.RefundOrderStatus = &v
	return s
}

func (s *RefundDetailListResponseBodyDataList) SetRelatedRefundOrderNum(v string) *RefundDetailListResponseBodyDataList {
	s.RelatedRefundOrderNum = &v
	return s
}

func (s *RefundDetailListResponseBodyDataList) SetTransactionNo(v string) *RefundDetailListResponseBodyDataList {
	s.TransactionNo = &v
	return s
}

func (s *RefundDetailListResponseBodyDataList) SetUtcCreateTime(v int64) *RefundDetailListResponseBodyDataList {
	s.UtcCreateTime = &v
	return s
}

type RefundDetailListResponseBodyDataPagination struct {
	// current page index
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// page size
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// the number of total refund orders
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// the number of total pages
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

func (s RefundDetailListResponseBodyDataPagination) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailListResponseBodyDataPagination) GoString() string {
	return s.String()
}

func (s *RefundDetailListResponseBodyDataPagination) SetCurrentPage(v int32) *RefundDetailListResponseBodyDataPagination {
	s.CurrentPage = &v
	return s
}

func (s *RefundDetailListResponseBodyDataPagination) SetPageSize(v int32) *RefundDetailListResponseBodyDataPagination {
	s.PageSize = &v
	return s
}

func (s *RefundDetailListResponseBodyDataPagination) SetTotalCount(v int32) *RefundDetailListResponseBodyDataPagination {
	s.TotalCount = &v
	return s
}

func (s *RefundDetailListResponseBodyDataPagination) SetTotalPage(v int32) *RefundDetailListResponseBodyDataPagination {
	s.TotalPage = &v
	return s
}

type RefundDetailListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundDetailListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefundDetailListResponse) String() string {
	return tea.Prettify(s)
}

func (s RefundDetailListResponse) GoString() string {
	return s.String()
}

func (s *RefundDetailListResponse) SetHeaders(v map[string]*string) *RefundDetailListResponse {
	s.Headers = v
	return s
}

func (s *RefundDetailListResponse) SetStatusCode(v int32) *RefundDetailListResponse {
	s.StatusCode = &v
	return s
}

func (s *RefundDetailListResponse) SetBody(v *RefundDetailListResponseBody) *RefundDetailListResponse {
	s.Body = v
	return s
}

type SearchHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// language code(refer to ISO_639), defaults to the buyer\\"s account configuration
	//
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s SearchHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchHeaders) GoString() string {
	return s.String()
}

func (s *SearchHeaders) SetCommonHeaders(v map[string]*string) *SearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchHeaders) SetXAcsAirticketAccessToken(v string) *SearchHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *SearchHeaders) SetXAcsAirticketLanguage(v string) *SearchHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type SearchRequest struct {
	// adult passenger amount 1-9
	//
	// example:
	//
	// 2
	Adults *int32 `json:"adults,omitempty" xml:"adults,omitempty"`
	// itinerary list
	//
	// This parameter is required.
	AirLegs []*SearchRequestAirLegs `json:"air_legs,omitempty" xml:"air_legs,omitempty" type:"Repeated"`
	// cabin class
	//
	// 1. **ALL_CABIN*	- : all cabin class
	//
	// 2. **Y*	- : economy class
	//
	// 3. **FC*	- : first class and business class
	//
	// 4. **S*	- : premium economy class
	//
	// 5. **YS*	- : economy class and premium economy class
	//
	// 6. **YSC*	- : economy class, premium economy class and business class
	//
	// example:
	//
	// ALL_CABIN
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// child passenger amount 0-9
	//
	// example:
	//
	// 1
	Children *int32 `json:"children,omitempty" xml:"children,omitempty"`
	// infant passenger amount 0-9
	//
	// example:
	//
	// 1
	Infants *int32 `json:"infants,omitempty" xml:"infants,omitempty"`
	// search controls
	SearchControlOptions *SearchRequestSearchControlOptions `json:"search_control_options,omitempty" xml:"search_control_options,omitempty" type:"Struct"`
}

func (s SearchRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchRequest) GoString() string {
	return s.String()
}

func (s *SearchRequest) SetAdults(v int32) *SearchRequest {
	s.Adults = &v
	return s
}

func (s *SearchRequest) SetAirLegs(v []*SearchRequestAirLegs) *SearchRequest {
	s.AirLegs = v
	return s
}

func (s *SearchRequest) SetCabinClass(v string) *SearchRequest {
	s.CabinClass = &v
	return s
}

func (s *SearchRequest) SetChildren(v int32) *SearchRequest {
	s.Children = &v
	return s
}

func (s *SearchRequest) SetInfants(v int32) *SearchRequest {
	s.Infants = &v
	return s
}

func (s *SearchRequest) SetSearchControlOptions(v *SearchRequestSearchControlOptions) *SearchRequest {
	s.SearchControlOptions = v
	return s
}

type SearchRequestAirLegs struct {
	// arrival airport [IATA airport code] list
	//
	// example:
	//
	// MFM
	ArrivalAirportList []*string `json:"arrival_airport_list,omitempty" xml:"arrival_airport_list,omitempty" type:"Repeated"`
	// arrival city code
	//
	// This parameter is required.
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// departure airport [IATA airport code] list
	//
	// example:
	//
	// PVG
	DepartureAirportList []*string `json:"departure_airport_list,omitempty" xml:"departure_airport_list,omitempty" type:"Repeated"`
	// departure city code
	//
	// This parameter is required.
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// departure date (eg: yyyyMMdd)
	//
	// This parameter is required.
	//
	// example:
	//
	// 20230310
	DepartureDate *string `json:"departure_date,omitempty" xml:"departure_date,omitempty"`
}

func (s SearchRequestAirLegs) String() string {
	return tea.Prettify(s)
}

func (s SearchRequestAirLegs) GoString() string {
	return s.String()
}

func (s *SearchRequestAirLegs) SetArrivalAirportList(v []*string) *SearchRequestAirLegs {
	s.ArrivalAirportList = v
	return s
}

func (s *SearchRequestAirLegs) SetArrivalCity(v string) *SearchRequestAirLegs {
	s.ArrivalCity = &v
	return s
}

func (s *SearchRequestAirLegs) SetDepartureAirportList(v []*string) *SearchRequestAirLegs {
	s.DepartureAirportList = v
	return s
}

func (s *SearchRequestAirLegs) SetDepartureCity(v string) *SearchRequestAirLegs {
	s.DepartureCity = &v
	return s
}

func (s *SearchRequestAirLegs) SetDepartureDate(v string) *SearchRequestAirLegs {
	s.DepartureDate = &v
	return s
}

type SearchRequestSearchControlOptions struct {
	// excluded airlines list
	AirlineExcludedList []*string `json:"airline_excluded_list,omitempty" xml:"airline_excluded_list,omitempty" type:"Repeated"`
	// preferred airlines list
	AirlinePreferList []*string `json:"airline_prefer_list,omitempty" xml:"airline_prefer_list,omitempty" type:"Repeated"`
	// example:
	//
	// A1
	ServiceQuality *string `json:"service_quality,omitempty" xml:"service_quality,omitempty"`
}

func (s SearchRequestSearchControlOptions) String() string {
	return tea.Prettify(s)
}

func (s SearchRequestSearchControlOptions) GoString() string {
	return s.String()
}

func (s *SearchRequestSearchControlOptions) SetAirlineExcludedList(v []*string) *SearchRequestSearchControlOptions {
	s.AirlineExcludedList = v
	return s
}

func (s *SearchRequestSearchControlOptions) SetAirlinePreferList(v []*string) *SearchRequestSearchControlOptions {
	s.AirlinePreferList = v
	return s
}

func (s *SearchRequestSearchControlOptions) SetServiceQuality(v string) *SearchRequestSearchControlOptions {
	s.ServiceQuality = &v
	return s
}

type SearchShrinkRequest struct {
	// adult passenger amount 1-9
	//
	// example:
	//
	// 2
	Adults *int32 `json:"adults,omitempty" xml:"adults,omitempty"`
	// itinerary list
	//
	// This parameter is required.
	AirLegsShrink *string `json:"air_legs,omitempty" xml:"air_legs,omitempty"`
	// cabin class
	//
	// 1. **ALL_CABIN*	- : all cabin class
	//
	// 2. **Y*	- : economy class
	//
	// 3. **FC*	- : first class and business class
	//
	// 4. **S*	- : premium economy class
	//
	// 5. **YS*	- : economy class and premium economy class
	//
	// 6. **YSC*	- : economy class, premium economy class and business class
	//
	// example:
	//
	// ALL_CABIN
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// child passenger amount 0-9
	//
	// example:
	//
	// 1
	Children *int32 `json:"children,omitempty" xml:"children,omitempty"`
	// infant passenger amount 0-9
	//
	// example:
	//
	// 1
	Infants *int32 `json:"infants,omitempty" xml:"infants,omitempty"`
	// search controls
	SearchControlOptionsShrink *string `json:"search_control_options,omitempty" xml:"search_control_options,omitempty"`
}

func (s SearchShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchShrinkRequest) SetAdults(v int32) *SearchShrinkRequest {
	s.Adults = &v
	return s
}

func (s *SearchShrinkRequest) SetAirLegsShrink(v string) *SearchShrinkRequest {
	s.AirLegsShrink = &v
	return s
}

func (s *SearchShrinkRequest) SetCabinClass(v string) *SearchShrinkRequest {
	s.CabinClass = &v
	return s
}

func (s *SearchShrinkRequest) SetChildren(v int32) *SearchShrinkRequest {
	s.Children = &v
	return s
}

func (s *SearchShrinkRequest) SetInfants(v int32) *SearchShrinkRequest {
	s.Infants = &v
	return s
}

func (s *SearchShrinkRequest) SetSearchControlOptionsShrink(v string) *SearchShrinkRequest {
	s.SearchControlOptionsShrink = &v
	return s
}

type SearchResponseBody struct {
	// request ID
	//
	// example:
	//
	// 2236993B-7BE7-5F92-B179-21FF08570165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *SearchResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error data
	//
	// example:
	//
	// {}
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// error message
	//
	// example:
	//
	// ""
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// http reqeust has been processed successfully，status code is 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true represents success, false represents failure
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBody) GoString() string {
	return s.String()
}

func (s *SearchResponseBody) SetRequestId(v string) *SearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchResponseBody) SetData(v *SearchResponseBodyData) *SearchResponseBody {
	s.Data = v
	return s
}

func (s *SearchResponseBody) SetErrorCode(v string) *SearchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchResponseBody) SetErrorData(v interface{}) *SearchResponseBody {
	s.ErrorData = v
	return s
}

func (s *SearchResponseBody) SetErrorMsg(v string) *SearchResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SearchResponseBody) SetStatus(v int32) *SearchResponseBody {
	s.Status = &v
	return s
}

func (s *SearchResponseBody) SetSuccess(v bool) *SearchResponseBody {
	s.Success = &v
	return s
}

type SearchResponseBodyData struct {
	// solution list
	SolutionList []*SearchResponseBodyDataSolutionList `json:"solution_list,omitempty" xml:"solution_list,omitempty" type:"Repeated"`
}

func (s SearchResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyData) SetSolutionList(v []*SearchResponseBodyDataSolutionList) *SearchResponseBodyData {
	s.SolutionList = v
	return s
}

type SearchResponseBodyDataSolutionList struct {
	// adult fare
	//
	// example:
	//
	// 600
	AdultPrice *float64 `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	// adult tax
	//
	// example:
	//
	// 11
	AdultTax *float64 `json:"adult_tax,omitempty" xml:"adult_tax,omitempty"`
	// child fare
	//
	// example:
	//
	// 500
	ChildPrice *float64 `json:"child_price,omitempty" xml:"child_price,omitempty"`
	// child tax
	//
	// example:
	//
	// 10
	ChildTax *float64 `json:"child_tax,omitempty" xml:"child_tax,omitempty"`
	// infant fare
	//
	// example:
	//
	// 400
	InfantPrice *float64 `json:"infant_price,omitempty" xml:"infant_price,omitempty"`
	// infant tax
	//
	// example:
	//
	// 9
	InfantTax *float64 `json:"infant_tax,omitempty" xml:"infant_tax,omitempty"`
	// segment list
	JourneyList []*SearchResponseBodyDataSolutionListJourneyList `json:"journey_list,omitempty" xml:"journey_list,omitempty" type:"Repeated"`
	// product type description
	//
	// example:
	//
	// ""
	ProductTypeDescription *string `json:"product_type_description,omitempty" xml:"product_type_description,omitempty"`
	// refund airline coupon description
	//
	// example:
	//
	// ""
	RefundTicketCouponDescription *string `json:"refund_ticket_coupon_description,omitempty" xml:"refund_ticket_coupon_description,omitempty"`
	// through check-in baggage policy
	SegmentBaggageCheckInInfoList []*SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList `json:"segment_baggage_check_in_info_list,omitempty" xml:"segment_baggage_check_in_info_list,omitempty" type:"Repeated"`
	// baggage rule
	SegmentBaggageMappingList []*SearchResponseBodyDataSolutionListSegmentBaggageMappingList `json:"segment_baggage_mapping_list,omitempty" xml:"segment_baggage_mapping_list,omitempty" type:"Repeated"`
	// change and refund policy
	SegmentRefundChangeRuleMappingList []*SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList `json:"segment_refund_change_rule_mapping_list,omitempty" xml:"segment_refund_change_rule_mapping_list,omitempty" type:"Repeated"`
	// solution ID
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s SearchResponseBodyDataSolutionList) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodyDataSolutionList) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyDataSolutionList) SetAdultPrice(v float64) *SearchResponseBodyDataSolutionList {
	s.AdultPrice = &v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetAdultTax(v float64) *SearchResponseBodyDataSolutionList {
	s.AdultTax = &v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetChildPrice(v float64) *SearchResponseBodyDataSolutionList {
	s.ChildPrice = &v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetChildTax(v float64) *SearchResponseBodyDataSolutionList {
	s.ChildTax = &v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetInfantPrice(v float64) *SearchResponseBodyDataSolutionList {
	s.InfantPrice = &v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetInfantTax(v float64) *SearchResponseBodyDataSolutionList {
	s.InfantTax = &v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetJourneyList(v []*SearchResponseBodyDataSolutionListJourneyList) *SearchResponseBodyDataSolutionList {
	s.JourneyList = v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetProductTypeDescription(v string) *SearchResponseBodyDataSolutionList {
	s.ProductTypeDescription = &v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetRefundTicketCouponDescription(v string) *SearchResponseBodyDataSolutionList {
	s.RefundTicketCouponDescription = &v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetSegmentBaggageCheckInInfoList(v []*SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) *SearchResponseBodyDataSolutionList {
	s.SegmentBaggageCheckInInfoList = v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetSegmentBaggageMappingList(v []*SearchResponseBodyDataSolutionListSegmentBaggageMappingList) *SearchResponseBodyDataSolutionList {
	s.SegmentBaggageMappingList = v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetSegmentRefundChangeRuleMappingList(v []*SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) *SearchResponseBodyDataSolutionList {
	s.SegmentRefundChangeRuleMappingList = v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetSolutionId(v string) *SearchResponseBodyDataSolutionList {
	s.SolutionId = &v
	return s
}

type SearchResponseBodyDataSolutionListJourneyList struct {
	// segment Info
	SegmentList []*SearchResponseBodyDataSolutionListJourneyListSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// number of transfers
	//
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s SearchResponseBodyDataSolutionListJourneyList) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodyDataSolutionListJourneyList) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyDataSolutionListJourneyList) SetSegmentList(v []*SearchResponseBodyDataSolutionListJourneyListSegmentList) *SearchResponseBodyDataSolutionListJourneyList {
	s.SegmentList = v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyList) SetTransferCount(v int32) *SearchResponseBodyDataSolutionListJourneyList {
	s.TransferCount = &v
	return s
}

type SearchResponseBodyDataSolutionListJourneyListSegmentList struct {
	// arrival airport code (capitalized)
	//
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// arrival city code (capitalized)
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// arrival terminal
	//
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// arrival time in string format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 10:40:00
	ArrivalTime *string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// available seats (for reference only)
	//
	// example:
	//
	// 7
	Availability *string `json:"availability,omitempty" xml:"availability,omitempty"`
	// RBD
	//
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// cabin class
	//
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// code share or not
	//
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// departure airport code (capitalized)
	//
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// departure city code (capitalized)
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// departure terminal
	//
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// departure time in string format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 07:55:00
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// equipment type
	//
	// example:
	//
	// 32Q
	EquipType *string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// flight time, unit: minute
	//
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// marketing airline code (ex.: KA)
	//
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// marketing airline flight no. (ex.: KA5809)
	//
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// marketing airline integer flight no. (ex.: 5809)
	//
	// example:
	//
	// 1259
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// operating airline code (ex.: CX)
	//
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// operating airline flight no. (ex.: CX601)
	//
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// segment ID format: flight no.+departure airport[IATA airport code]+arrival airport[IATA airport code]+departure time(MMdd)
	//
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// stop city list
	//
	// when stop_quantity > 1, use “,” for seperation
	//
	// example:
	//
	// MFM,PVG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// number of stops
	//
	// example:
	//
	// 0
	StopQuantity *int32 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s SearchResponseBodyDataSolutionListJourneyListSegmentList) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodyDataSolutionListJourneyListSegmentList) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalAirport(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalCity(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalTerminal(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalTime(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetAvailability(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.Availability = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetCabin(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.Cabin = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetCabinClass(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.CabinClass = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetCodeShare(v bool) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.CodeShare = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureAirport(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureCity(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureTerminal(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureTime(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetEquipType(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.EquipType = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetFlightDuration(v int32) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetMarketingAirline(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetMarketingFlightNo(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetMarketingFlightNoInt(v int32) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetOperatingAirline(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetOperatingFlightNo(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetSegmentId(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.SegmentId = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetStopCityList(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.StopCityList = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetStopQuantity(v int32) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.StopQuantity = &v
	return s
}

type SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList struct {
	// through check-in baggage policy type
	//
	// 1. baggage through check-in between segments
	//
	// 2. baggage re-check-in needed between segments
	//
	// 3. baggage through check-in at stop city ( applies for stop flight )
	//
	// 4. baggage re-checkin needed at stop city ( applies for stop flight )
	//
	// example:
	//
	// 1
	LuggageDirectInfoType *int32 `json:"luggage_direct_info_type,omitempty" xml:"luggage_direct_info_type,omitempty"`
	// segment id list.
	//
	// all the listed segment ids share the same baggage through check-in  policy
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) SetLuggageDirectInfoType(v int32) *SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList {
	s.LuggageDirectInfoType = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) SetSegmentIdList(v []*string) *SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList {
	s.SegmentIdList = v
	return s
}

type SearchResponseBodyDataSolutionListSegmentBaggageMappingList struct {
	// baggage rule mapping, key is passenger type, value is baggage allowance details
	PassengerBaggageAllowanceMapping map[string]*DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue `json:"passenger_baggage_allowance_mapping,omitempty" xml:"passenger_baggage_allowance_mapping,omitempty"`
	// segment id list.
	//
	// all the listed segment id share the same baggage rule
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s SearchResponseBodyDataSolutionListSegmentBaggageMappingList) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodyDataSolutionListSegmentBaggageMappingList) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyDataSolutionListSegmentBaggageMappingList) SetPassengerBaggageAllowanceMapping(v map[string]*DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) *SearchResponseBodyDataSolutionListSegmentBaggageMappingList {
	s.PassengerBaggageAllowanceMapping = v
	return s
}

func (s *SearchResponseBodyDataSolutionListSegmentBaggageMappingList) SetSegmentIdList(v []*string) *SearchResponseBodyDataSolutionListSegmentBaggageMappingList {
	s.SegmentIdList = v
	return s
}

type SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList struct {
	// change and refund policy mapping, key is passenger type, value is change and refund policy details
	RefundChangeRuleMap map[string]*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue `json:"refund_change_rule_map,omitempty" xml:"refund_change_rule_map,omitempty"`
	// segment id list.
	//
	// all the listed segment ids share the same change and refund policy
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) SetRefundChangeRuleMap(v map[string]*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) *SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList {
	s.RefundChangeRuleMap = v
	return s
}

func (s *SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) SetSegmentIdList(v []*string) *SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList {
	s.SegmentIdList = v
	return s
}

type SearchResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchResponse) GoString() string {
	return s.String()
}

func (s *SearchResponse) SetHeaders(v map[string]*string) *SearchResponse {
	s.Headers = v
	return s
}

func (s *SearchResponse) SetStatusCode(v int32) *SearchResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchResponse) SetBody(v *SearchResponseBody) *SearchResponse {
	s.Body = v
	return s
}

type TicketingHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// language code(refer to ISO_639), defaults to the buyer\\"s account configuration
	//
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s TicketingHeaders) String() string {
	return tea.Prettify(s)
}

func (s TicketingHeaders) GoString() string {
	return s.String()
}

func (s *TicketingHeaders) SetCommonHeaders(v map[string]*string) *TicketingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TicketingHeaders) SetXAcsAirticketAccessToken(v string) *TicketingHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *TicketingHeaders) SetXAcsAirticketLanguage(v string) *TicketingHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type TicketingRequest struct {
	// order number
	//
	// This parameter is required.
	//
	// example:
	//
	// 496***2617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s TicketingRequest) String() string {
	return tea.Prettify(s)
}

func (s TicketingRequest) GoString() string {
	return s.String()
}

func (s *TicketingRequest) SetOrderNum(v int64) *TicketingRequest {
	s.OrderNum = &v
	return s
}

type TicketingResponseBody struct {
	// request ID
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *TicketingResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error data
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// http reqeust has been processed successfully，status code is 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true represents success, false represents failure
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TicketingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TicketingResponseBody) GoString() string {
	return s.String()
}

func (s *TicketingResponseBody) SetRequestId(v string) *TicketingResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketingResponseBody) SetData(v *TicketingResponseBodyData) *TicketingResponseBody {
	s.Data = v
	return s
}

func (s *TicketingResponseBody) SetErrorCode(v string) *TicketingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TicketingResponseBody) SetErrorData(v interface{}) *TicketingResponseBody {
	s.ErrorData = v
	return s
}

func (s *TicketingResponseBody) SetErrorMsg(v string) *TicketingResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TicketingResponseBody) SetStatus(v int32) *TicketingResponseBody {
	s.Status = &v
	return s
}

func (s *TicketingResponseBody) SetSuccess(v bool) *TicketingResponseBody {
	s.Success = &v
	return s
}

type TicketingResponseBodyData struct {
	// order number
	//
	// example:
	//
	// 496***2617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// transaction serial number
	//
	// example:
	//
	// hkduendkd-2023-dj0
	TransactionNo *string `json:"transaction_no,omitempty" xml:"transaction_no,omitempty"`
}

func (s TicketingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TicketingResponseBodyData) GoString() string {
	return s.String()
}

func (s *TicketingResponseBodyData) SetOrderNum(v int64) *TicketingResponseBodyData {
	s.OrderNum = &v
	return s
}

func (s *TicketingResponseBodyData) SetTransactionNo(v string) *TicketingResponseBodyData {
	s.TransactionNo = &v
	return s
}

type TicketingResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketingResponse) String() string {
	return tea.Prettify(s)
}

func (s TicketingResponse) GoString() string {
	return s.String()
}

func (s *TicketingResponse) SetHeaders(v map[string]*string) *TicketingResponse {
	s.Headers = v
	return s
}

func (s *TicketingResponse) SetStatusCode(v int32) *TicketingResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketingResponse) SetBody(v *TicketingResponseBody) *TicketingResponse {
	s.Body = v
	return s
}

type TicketingCheckHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// language Code(refer to ISO_639), defaults to the buyer\\"s account configuration
	//
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s TicketingCheckHeaders) String() string {
	return tea.Prettify(s)
}

func (s TicketingCheckHeaders) GoString() string {
	return s.String()
}

func (s *TicketingCheckHeaders) SetCommonHeaders(v map[string]*string) *TicketingCheckHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TicketingCheckHeaders) SetXAcsAirticketAccessToken(v string) *TicketingCheckHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *TicketingCheckHeaders) SetXAcsAirticketLanguage(v string) *TicketingCheckHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type TicketingCheckRequest struct {
	// order number
	//
	// This parameter is required.
	//
	// example:
	//
	// 496***2617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s TicketingCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s TicketingCheckRequest) GoString() string {
	return s.String()
}

func (s *TicketingCheckRequest) SetOrderNum(v int64) *TicketingCheckRequest {
	s.OrderNum = &v
	return s
}

type TicketingCheckResponseBody struct {
	// request ID
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *TicketingCheckResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error data
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// http reqeust has been processed successfully，status code is 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true represents success, false represents failure
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TicketingCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TicketingCheckResponseBody) GoString() string {
	return s.String()
}

func (s *TicketingCheckResponseBody) SetRequestId(v string) *TicketingCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketingCheckResponseBody) SetData(v *TicketingCheckResponseBodyData) *TicketingCheckResponseBody {
	s.Data = v
	return s
}

func (s *TicketingCheckResponseBody) SetErrorCode(v string) *TicketingCheckResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TicketingCheckResponseBody) SetErrorData(v interface{}) *TicketingCheckResponseBody {
	s.ErrorData = v
	return s
}

func (s *TicketingCheckResponseBody) SetErrorMsg(v string) *TicketingCheckResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TicketingCheckResponseBody) SetStatus(v int32) *TicketingCheckResponseBody {
	s.Status = &v
	return s
}

func (s *TicketingCheckResponseBody) SetSuccess(v bool) *TicketingCheckResponseBody {
	s.Success = &v
	return s
}

type TicketingCheckResponseBodyData struct {
	// order number
	//
	// example:
	//
	// 496***2617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s TicketingCheckResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TicketingCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *TicketingCheckResponseBodyData) SetOrderNum(v int64) *TicketingCheckResponseBodyData {
	s.OrderNum = &v
	return s
}

type TicketingCheckResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketingCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketingCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s TicketingCheckResponse) GoString() string {
	return s.String()
}

func (s *TicketingCheckResponse) SetHeaders(v map[string]*string) *TicketingCheckResponse {
	s.Headers = v
	return s
}

func (s *TicketingCheckResponse) SetStatusCode(v int32) *TicketingCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketingCheckResponse) SetBody(v *TicketingCheckResponseBody) *TicketingCheckResponse {
	s.Body = v
	return s
}

type TransitVisaHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// access_token
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsImtpZCI6InN1ZXpfa2V5aWQifQ.eyJqdGkiOiJCQldMaWIzN0VxbC0xMjhhR2N5elJ3IiwiaWF0IjoxNjc3MDY2NTAxLCJleHAiOjE2NzcwNzM3MDEsIm5iZiI6MTY3NzA2NjQ0MX0.AF0DxsZK4Edyg0C6ObRQFUo36R1VYrb5IYmak25TmL1OfR5RkIUc3PpqFuQKNLKXf5fOtVQaKjaexzwodVeWZQDKEG_RPt_Ybb99EnEm6vPKs6e3pWFbKiBq71WleLHhVrdFb4YPowRKjc7bG0jyGUxiQ2iXy0RWDj9tIjfI-KEdzNp5oVnX7j4p3H12DwQrRPmd1nz3BciAQNINvDpzqusuIUw8JXyLFCz838Y0NhwB1_bYZyctxRLSzrGZuI5rrWtItgupqMsOlJ3RNy1QrIbQ2g6nPmzl-atOqcQ4Nw0HeDLR8dhM1OsIcFLbKXBUtwXofflhzAQrkDxhwYiXii
	XAcsAirticketAccessToken *string `json:"x-acs-airticket-access-token,omitempty" xml:"x-acs-airticket-access-token,omitempty"`
	// example:
	//
	// en_US
	XAcsAirticketLanguage *string `json:"x-acs-airticket-language,omitempty" xml:"x-acs-airticket-language,omitempty"`
}

func (s TransitVisaHeaders) String() string {
	return tea.Prettify(s)
}

func (s TransitVisaHeaders) GoString() string {
	return s.String()
}

func (s *TransitVisaHeaders) SetCommonHeaders(v map[string]*string) *TransitVisaHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TransitVisaHeaders) SetXAcsAirticketAccessToken(v string) *TransitVisaHeaders {
	s.XAcsAirticketAccessToken = &v
	return s
}

func (s *TransitVisaHeaders) SetXAcsAirticketLanguage(v string) *TransitVisaHeaders {
	s.XAcsAirticketLanguage = &v
	return s
}

type TransitVisaRequest struct {
	FlightSegmentParamList []*TransitVisaRequestFlightSegmentParamList `json:"flight_segment_param_list,omitempty" xml:"flight_segment_param_list,omitempty" type:"Repeated"`
}

func (s TransitVisaRequest) String() string {
	return tea.Prettify(s)
}

func (s TransitVisaRequest) GoString() string {
	return s.String()
}

func (s *TransitVisaRequest) SetFlightSegmentParamList(v []*TransitVisaRequestFlightSegmentParamList) *TransitVisaRequest {
	s.FlightSegmentParamList = v
	return s
}

type TransitVisaRequestFlightSegmentParamList struct {
	// This parameter is required.
	//
	// example:
	//
	// SIN
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1705285430445
	ArrivalTime *int64 `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HGH
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// example:
	//
	// T1
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1705285430445
	DepartureTime *int64 `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CZ
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CZ616
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// example:
	//
	// CZ
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// example:
	//
	// SEL,HKG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// example:
	//
	// CZ
	TicketingAirline *string `json:"ticketing_airline,omitempty" xml:"ticketing_airline,omitempty"`
}

func (s TransitVisaRequestFlightSegmentParamList) String() string {
	return tea.Prettify(s)
}

func (s TransitVisaRequestFlightSegmentParamList) GoString() string {
	return s.String()
}

func (s *TransitVisaRequestFlightSegmentParamList) SetArrivalAirport(v string) *TransitVisaRequestFlightSegmentParamList {
	s.ArrivalAirport = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetArrivalTerminal(v string) *TransitVisaRequestFlightSegmentParamList {
	s.ArrivalTerminal = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetArrivalTime(v int64) *TransitVisaRequestFlightSegmentParamList {
	s.ArrivalTime = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetCodeShare(v bool) *TransitVisaRequestFlightSegmentParamList {
	s.CodeShare = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetDepartureAirport(v string) *TransitVisaRequestFlightSegmentParamList {
	s.DepartureAirport = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetDepartureTerminal(v string) *TransitVisaRequestFlightSegmentParamList {
	s.DepartureTerminal = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetDepartureTime(v int64) *TransitVisaRequestFlightSegmentParamList {
	s.DepartureTime = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetMarketingAirline(v string) *TransitVisaRequestFlightSegmentParamList {
	s.MarketingAirline = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetMarketingFlightNo(v string) *TransitVisaRequestFlightSegmentParamList {
	s.MarketingFlightNo = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetOperatingAirline(v string) *TransitVisaRequestFlightSegmentParamList {
	s.OperatingAirline = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetStopCityList(v string) *TransitVisaRequestFlightSegmentParamList {
	s.StopCityList = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetTicketingAirline(v string) *TransitVisaRequestFlightSegmentParamList {
	s.TicketingAirline = &v
	return s
}

type TransitVisaShrinkRequest struct {
	FlightSegmentParamListShrink *string `json:"flight_segment_param_list,omitempty" xml:"flight_segment_param_list,omitempty"`
}

func (s TransitVisaShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s TransitVisaShrinkRequest) GoString() string {
	return s.String()
}

func (s *TransitVisaShrinkRequest) SetFlightSegmentParamListShrink(v string) *TransitVisaShrinkRequest {
	s.FlightSegmentParamListShrink = &v
	return s
}

type TransitVisaResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*TransitVisaResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TransitVisaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransitVisaResponseBody) GoString() string {
	return s.String()
}

func (s *TransitVisaResponseBody) SetRequestId(v string) *TransitVisaResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransitVisaResponseBody) SetData(v []*TransitVisaResponseBodyData) *TransitVisaResponseBody {
	s.Data = v
	return s
}

func (s *TransitVisaResponseBody) SetErrorCode(v string) *TransitVisaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TransitVisaResponseBody) SetErrorData(v interface{}) *TransitVisaResponseBody {
	s.ErrorData = v
	return s
}

func (s *TransitVisaResponseBody) SetErrorMsg(v string) *TransitVisaResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TransitVisaResponseBody) SetStatus(v int32) *TransitVisaResponseBody {
	s.Status = &v
	return s
}

func (s *TransitVisaResponseBody) SetSuccess(v bool) *TransitVisaResponseBody {
	s.Success = &v
	return s
}

type TransitVisaResponseBodyData struct {
	// example:
	//
	// HGH
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// example:
	//
	// 1
	VisaType *int32 `json:"visa_type,omitempty" xml:"visa_type,omitempty"`
}

func (s TransitVisaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TransitVisaResponseBodyData) GoString() string {
	return s.String()
}

func (s *TransitVisaResponseBodyData) SetCityCode(v string) *TransitVisaResponseBodyData {
	s.CityCode = &v
	return s
}

func (s *TransitVisaResponseBodyData) SetVisaType(v int32) *TransitVisaResponseBodyData {
	s.VisaType = &v
	return s
}

type TransitVisaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransitVisaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransitVisaResponse) String() string {
	return tea.Prettify(s)
}

func (s TransitVisaResponse) GoString() string {
	return s.String()
}

func (s *TransitVisaResponse) SetHeaders(v map[string]*string) *TransitVisaResponse {
	s.Headers = v
	return s
}

func (s *TransitVisaResponse) SetStatusCode(v int32) *TransitVisaResponse {
	s.StatusCode = &v
	return s
}

func (s *TransitVisaResponse) SetBody(v *TransitVisaResponseBody) *TransitVisaResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("airticketopen"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 账号资金流水
//
// @param request - AccountFlowListRequest
//
// @param headers - AccountFlowListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AccountFlowListResponse
func (client *Client) AccountFlowListWithOptions(request *AccountFlowListRequest, headers *AccountFlowListHeaders, runtime *util.RuntimeOptions) (_result *AccountFlowListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DayNum)) {
		query["day_num"] = request.DayNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["page_index"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["page_size"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UtcBeginTime)) {
		query["utc_begin_time"] = request.UtcBeginTime
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AccountFlowList"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/account/flow-list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AccountFlowListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 账号资金流水
//
// @param request - AccountFlowListRequest
//
// @return AccountFlowListResponse
func (client *Client) AccountFlowList(request *AccountFlowListRequest) (_result *AccountFlowListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AccountFlowListHeaders{}
	_result = &AccountFlowListResponse{}
	_body, _err := client.AccountFlowListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Ancillary - Suggestion
//
// Description:
//
// search ancillary for selected solution, you should enter the solution_id returned by enrich.
//
// @param request - AncillarySuggestRequest
//
// @param headers - AncillarySuggestHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AncillarySuggestResponse
func (client *Client) AncillarySuggestWithOptions(request *AncillarySuggestRequest, headers *AncillarySuggestHeaders, runtime *util.RuntimeOptions) (_result *AncillarySuggestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SolutionId)) {
		body["solution_id"] = request.SolutionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AncillarySuggest"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/ancillary/action-suggest"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AncillarySuggestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Ancillary - Suggestion
//
// Description:
//
// search ancillary for selected solution, you should enter the solution_id returned by enrich.
//
// @param request - AncillarySuggestRequest
//
// @return AncillarySuggestResponse
func (client *Client) AncillarySuggest(request *AncillarySuggestRequest) (_result *AncillarySuggestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AncillarySuggestHeaders{}
	_result = &AncillarySuggestResponse{}
	_body, _err := client.AncillarySuggestWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Transaction-Reservation
//
// Description:
//
// Enter solution_id returned by enrich, ancillary_id returned by ancillarySuggest(optional), passengers information and contact information, the book interface will create an order wait for pay.
//
// There are two issues should be noticed:
//
// 1. the solution_id must be processed by pricing.
//
// 2. the order created by book interface should be pay within 30 minutes, otherwise the order will be closed.
//
// @param tmpReq - BookRequest
//
// @param headers - BookHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BookResponse
func (client *Client) BookWithOptions(tmpReq *BookRequest, headers *BookHeaders, runtime *util.RuntimeOptions) (_result *BookResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BookShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Contact)) {
		request.ContactShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Contact, tea.String("contact"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.PassengerAncillaryPurchaseMapList)) {
		request.PassengerAncillaryPurchaseMapListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PassengerAncillaryPurchaseMapList, tea.String("passenger_ancillary_purchase_map_list"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.PassengerList)) {
		request.PassengerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PassengerList, tea.String("passenger_list"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactShrink)) {
		body["contact"] = request.ContactShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OutOrderNum)) {
		body["out_order_num"] = request.OutOrderNum
	}

	if !tea.BoolValue(util.IsUnset(request.PassengerAncillaryPurchaseMapListShrink)) {
		body["passenger_ancillary_purchase_map_list"] = request.PassengerAncillaryPurchaseMapListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PassengerListShrink)) {
		body["passenger_list"] = request.PassengerListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SolutionId)) {
		body["solution_id"] = request.SolutionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Book"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/trade/action-book"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Transaction-Reservation
//
// Description:
//
// Enter solution_id returned by enrich, ancillary_id returned by ancillarySuggest(optional), passengers information and contact information, the book interface will create an order wait for pay.
//
// There are two issues should be noticed:
//
// 1. the solution_id must be processed by pricing.
//
// 2. the order created by book interface should be pay within 30 minutes, otherwise the order will be closed.
//
// @param request - BookRequest
//
// @return BookResponse
func (client *Client) Book(request *BookRequest) (_result *BookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BookHeaders{}
	_result = &BookResponse{}
	_body, _err := client.BookWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Transaction - Unpaid Cancellation
//
// Description:
//
// close an unpaid order
//
// @param request - CancelRequest
//
// @param headers - CancelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelResponse
func (client *Client) CancelWithOptions(request *CancelRequest, headers *CancelHeaders, runtime *util.RuntimeOptions) (_result *CancelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderNum)) {
		body["order_num"] = request.OrderNum
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Cancel"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/trade/action-cancel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Transaction - Unpaid Cancellation
//
// Description:
//
// close an unpaid order
//
// @param request - CancelRequest
//
// @return CancelResponse
func (client *Client) Cancel(request *CancelRequest) (_result *CancelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CancelHeaders{}
	_result = &CancelResponse{}
	_body, _err := client.CancelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 改签-Apply
//
// @param tmpReq - ChangeApplyRequest
//
// @param headers - ChangeApplyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeApplyResponse
func (client *Client) ChangeApplyWithOptions(tmpReq *ChangeApplyRequest, headers *ChangeApplyHeaders, runtime *util.RuntimeOptions) (_result *ChangeApplyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ChangeApplyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ChangePassengerList)) {
		request.ChangePassengerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChangePassengerList, tea.String("change_passenger_list"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ChangedJourneys)) {
		request.ChangedJourneysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChangedJourneys, tea.String("changed_journeys"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Contact)) {
		request.ContactShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Contact, tea.String("contact"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangePassengerListShrink)) {
		body["change_passenger_list"] = request.ChangePassengerListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ChangedJourneysShrink)) {
		body["changed_journeys"] = request.ChangedJourneysShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ContactShrink)) {
		body["contact"] = request.ContactShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNum)) {
		body["order_num"] = request.OrderNum
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeApply"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/change/action-apply"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeApplyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 改签-Apply
//
// @param request - ChangeApplyRequest
//
// @return ChangeApplyResponse
func (client *Client) ChangeApply(request *ChangeApplyRequest) (_result *ChangeApplyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ChangeApplyHeaders{}
	_result = &ChangeApplyResponse{}
	_body, _err := client.ChangeApplyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 改签-取消
//
// @param request - ChangeCancelRequest
//
// @param headers - ChangeCancelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeCancelResponse
func (client *Client) ChangeCancelWithOptions(request *ChangeCancelRequest, headers *ChangeCancelHeaders, runtime *util.RuntimeOptions) (_result *ChangeCancelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeOrderNum)) {
		body["change_order_num"] = request.ChangeOrderNum
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeCancel"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/change/action-cancel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeCancelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 改签-取消
//
// @param request - ChangeCancelRequest
//
// @return ChangeCancelResponse
func (client *Client) ChangeCancel(request *ChangeCancelRequest) (_result *ChangeCancelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ChangeCancelHeaders{}
	_result = &ChangeCancelResponse{}
	_body, _err := client.ChangeCancelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 改签-确认
//
// @param request - ChangeConfirmRequest
//
// @param headers - ChangeConfirmHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeConfirmResponse
func (client *Client) ChangeConfirmWithOptions(request *ChangeConfirmRequest, headers *ChangeConfirmHeaders, runtime *util.RuntimeOptions) (_result *ChangeConfirmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeOrderNum)) {
		body["change_order_num"] = request.ChangeOrderNum
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeConfirm"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/change/action-confirm"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeConfirmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 改签-确认
//
// @param request - ChangeConfirmRequest
//
// @return ChangeConfirmResponse
func (client *Client) ChangeConfirm(request *ChangeConfirmRequest) (_result *ChangeConfirmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ChangeConfirmHeaders{}
	_result = &ChangeConfirmResponse{}
	_body, _err := client.ChangeConfirmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 改签-Detail
//
// @param request - ChangeDetailRequest
//
// @param headers - ChangeDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeDetailResponse
func (client *Client) ChangeDetailWithOptions(request *ChangeDetailRequest, headers *ChangeDetailHeaders, runtime *util.RuntimeOptions) (_result *ChangeDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeOrderNum)) {
		query["change_order_num"] = request.ChangeOrderNum
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeDetail"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/change/detail"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 改签-Detail
//
// @param request - ChangeDetailRequest
//
// @return ChangeDetailResponse
func (client *Client) ChangeDetail(request *ChangeDetailRequest) (_result *ChangeDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ChangeDetailHeaders{}
	_result = &ChangeDetailResponse{}
	_body, _err := client.ChangeDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 改签单列表-关于买家账号
//
// @param request - ChangeDetailListOfBuyerRequest
//
// @param headers - ChangeDetailListOfBuyerHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeDetailListOfBuyerResponse
func (client *Client) ChangeDetailListOfBuyerWithOptions(request *ChangeDetailListOfBuyerRequest, headers *ChangeDetailListOfBuyerHeaders, runtime *util.RuntimeOptions) (_result *ChangeDetailListOfBuyerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["page_index"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["page_size"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UtcCreateBegin)) {
		query["utc_create_begin"] = request.UtcCreateBegin
	}

	if !tea.BoolValue(util.IsUnset(request.UtcCreateEnd)) {
		query["utc_create_end"] = request.UtcCreateEnd
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeDetailListOfBuyer"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/change/buyer/detail-list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeDetailListOfBuyerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 改签单列表-关于买家账号
//
// @param request - ChangeDetailListOfBuyerRequest
//
// @return ChangeDetailListOfBuyerResponse
func (client *Client) ChangeDetailListOfBuyer(request *ChangeDetailListOfBuyerRequest) (_result *ChangeDetailListOfBuyerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ChangeDetailListOfBuyerHeaders{}
	_result = &ChangeDetailListOfBuyerResponse{}
	_body, _err := client.ChangeDetailListOfBuyerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 改签单列表-关于正向订单
//
// @param request - ChangeDetailListOfOrderNumRequest
//
// @param headers - ChangeDetailListOfOrderNumHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeDetailListOfOrderNumResponse
func (client *Client) ChangeDetailListOfOrderNumWithOptions(request *ChangeDetailListOfOrderNumRequest, headers *ChangeDetailListOfOrderNumHeaders, runtime *util.RuntimeOptions) (_result *ChangeDetailListOfOrderNumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderNum)) {
		query["order_num"] = request.OrderNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["page_index"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["page_size"] = request.PageSize
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeDetailListOfOrderNum"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/change/order-num/detail-list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeDetailListOfOrderNumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 改签单列表-关于正向订单
//
// @param request - ChangeDetailListOfOrderNumRequest
//
// @return ChangeDetailListOfOrderNumResponse
func (client *Client) ChangeDetailListOfOrderNum(request *ChangeDetailListOfOrderNumRequest) (_result *ChangeDetailListOfOrderNumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ChangeDetailListOfOrderNumHeaders{}
	_result = &ChangeDetailListOfOrderNumResponse{}
	_body, _err := client.ChangeDetailListOfOrderNumWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 数据收集-低价航班信息
//
// @param tmpReq - CollectFlightLowestPriceRequest
//
// @param headers - CollectFlightLowestPriceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CollectFlightLowestPriceResponse
func (client *Client) CollectFlightLowestPriceWithOptions(tmpReq *CollectFlightLowestPriceRequest, headers *CollectFlightLowestPriceHeaders, runtime *util.RuntimeOptions) (_result *CollectFlightLowestPriceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CollectFlightLowestPriceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LowestPriceFlightInfoList)) {
		request.LowestPriceFlightInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LowestPriceFlightInfoList, tea.String("lowest_price_flight_info_list"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LowestPriceFlightInfoListShrink)) {
		body["lowest_price_flight_info_list"] = request.LowestPriceFlightInfoListShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CollectFlightLowestPrice"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/data-collect/flight-lowest-price"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CollectFlightLowestPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据收集-低价航班信息
//
// @param request - CollectFlightLowestPriceRequest
//
// @return CollectFlightLowestPriceResponse
func (client *Client) CollectFlightLowestPrice(request *CollectFlightLowestPriceRequest) (_result *CollectFlightLowestPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollectFlightLowestPriceHeaders{}
	_result = &CollectFlightLowestPriceResponse{}
	_body, _err := client.CollectFlightLowestPriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Search-Enrich
//
// Description:
//
// Enrich supports two modes:
//
// 1. mode1: enter solution_id returned by Search.
//
// 2. mode2: enter journeyParamList.
//
// If you already confirm which flight to fly with, then you can use mode2, otherwise, use mode1(search first, then chose one solution_ID and Enrich).
//
// @param tmpReq - EnrichRequest
//
// @param headers - EnrichHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnrichResponse
func (client *Client) EnrichWithOptions(tmpReq *EnrichRequest, headers *EnrichHeaders, runtime *util.RuntimeOptions) (_result *EnrichResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &EnrichShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.JourneyParamList)) {
		request.JourneyParamListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JourneyParamList, tea.String("journey_param_list"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Adults)) {
		body["adults"] = request.Adults
	}

	if !tea.BoolValue(util.IsUnset(request.CabinClass)) {
		body["cabin_class"] = request.CabinClass
	}

	if !tea.BoolValue(util.IsUnset(request.Children)) {
		body["children"] = request.Children
	}

	if !tea.BoolValue(util.IsUnset(request.Infants)) {
		body["infants"] = request.Infants
	}

	if !tea.BoolValue(util.IsUnset(request.JourneyParamListShrink)) {
		body["journey_param_list"] = request.JourneyParamListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SolutionId)) {
		body["solution_id"] = request.SolutionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Enrich"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/trade/action-enrich"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnrichResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Search-Enrich
//
// Description:
//
// Enrich supports two modes:
//
// 1. mode1: enter solution_id returned by Search.
//
// 2. mode2: enter journeyParamList.
//
// If you already confirm which flight to fly with, then you can use mode2, otherwise, use mode1(search first, then chose one solution_ID and Enrich).
//
// @param request - EnrichRequest
//
// @return EnrichResponse
func (client *Client) Enrich(request *EnrichRequest) (_result *EnrichResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EnrichHeaders{}
	_result = &EnrichResponse{}
	_body, _err := client.EnrichWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 附件上传
//
// @param request - FileUploadRequest
//
// @param headers - FileUploadHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileUploadResponse
func (client *Client) FileUploadWithOptions(request *FileUploadRequest, headers *FileUploadHeaders, runtime *util.RuntimeOptions) (_result *FileUploadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileContent)) {
		body["file_content"] = request.FileContent
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNum)) {
		body["order_num"] = request.OrderNum
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FileUpload"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/attachment/action-upload"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FileUploadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 附件上传
//
// @param request - FileUploadRequest
//
// @return FileUploadResponse
func (client *Client) FileUpload(request *FileUploadRequest) (_result *FileUploadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FileUploadHeaders{}
	_result = &FileUploadResponse{}
	_body, _err := client.FileUploadWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 航变信息-关于订单
//
// @param request - FlightChangeOfOrderRequest
//
// @param headers - FlightChangeOfOrderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FlightChangeOfOrderResponse
func (client *Client) FlightChangeOfOrderWithOptions(request *FlightChangeOfOrderRequest, headers *FlightChangeOfOrderHeaders, runtime *util.RuntimeOptions) (_result *FlightChangeOfOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderNum)) {
		query["order_num"] = request.OrderNum
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FlightChangeOfOrder"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/flightchange/of-order"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FlightChangeOfOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 航变信息-关于订单
//
// @param request - FlightChangeOfOrderRequest
//
// @return FlightChangeOfOrderResponse
func (client *Client) FlightChangeOfOrder(request *FlightChangeOfOrderRequest) (_result *FlightChangeOfOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FlightChangeOfOrderHeaders{}
	_result = &FlightChangeOfOrderResponse{}
	_body, _err := client.FlightChangeOfOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get Token
//
// @param request - GetTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenResponse
func (client *Client) GetTokenWithOptions(request *GetTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["app_key"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.AppSecret)) {
		query["app_secret"] = request.AppSecret
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetToken"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/token"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get Token
//
// @param request - GetTokenRequest
//
// @return GetTokenResponse
func (client *Client) GetToken(request *GetTokenRequest) (_result *GetTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 航程行李直挂
//
// @param tmpReq - LuggageDirectRequest
//
// @param headers - LuggageDirectHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LuggageDirectResponse
func (client *Client) LuggageDirectWithOptions(tmpReq *LuggageDirectRequest, headers *LuggageDirectHeaders, runtime *util.RuntimeOptions) (_result *LuggageDirectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &LuggageDirectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FlightSegmentParamList)) {
		request.FlightSegmentParamListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FlightSegmentParamList, tea.String("flight_segment_param_list"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FlightSegmentParamListShrink)) {
		query["flight_segment_param_list"] = request.FlightSegmentParamListShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LuggageDirect"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/flight-data/luggage-direct"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &LuggageDirectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 航程行李直挂
//
// @param request - LuggageDirectRequest
//
// @return LuggageDirectResponse
func (client *Client) LuggageDirect(request *LuggageDirectRequest) (_result *LuggageDirectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LuggageDirectHeaders{}
	_result = &LuggageDirectResponse{}
	_body, _err := client.LuggageDirectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Trade-Order Details
//
// Description:
//
// query order detail
//
// @param request - OrderDetailRequest
//
// @param headers - OrderDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OrderDetailResponse
func (client *Client) OrderDetailWithOptions(request *OrderDetailRequest, headers *OrderDetailHeaders, runtime *util.RuntimeOptions) (_result *OrderDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderNum)) {
		query["order_num"] = request.OrderNum
	}

	if !tea.BoolValue(util.IsUnset(request.OutOrderNum)) {
		query["out_order_num"] = request.OutOrderNum
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OrderDetail"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/trade/order-detail"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &OrderDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Trade-Order Details
//
// Description:
//
// query order detail
//
// @param request - OrderDetailRequest
//
// @return OrderDetailResponse
func (client *Client) OrderDetail(request *OrderDetailRequest) (_result *OrderDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OrderDetailHeaders{}
	_result = &OrderDetailResponse{}
	_body, _err := client.OrderDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Trade - Order List
//
// Description:
//
// query order list
//
// @param request - OrderListRequest
//
// @param headers - OrderListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OrderListResponse
func (client *Client) OrderListWithOptions(request *OrderListRequest, headers *OrderListHeaders, runtime *util.RuntimeOptions) (_result *OrderListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BookTimeEnd)) {
		query["book_time_end"] = request.BookTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.BookTimeStart)) {
		query["book_time_start"] = request.BookTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["page_index"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["page_size"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OrderList"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/trade/order-list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &OrderListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Trade - Order List
//
// Description:
//
// query order list
//
// @param request - OrderListRequest
//
// @return OrderListResponse
func (client *Client) OrderList(request *OrderListRequest) (_result *OrderListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OrderListHeaders{}
	_result = &OrderListResponse{}
	_body, _err := client.OrderListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Trade - Seat and Price Verification
//
// Description:
//
// Check is price and remaining seats of solution you selected has changed. You should enter the solution_id returned by enrich.
//
// @param request - PricingRequest
//
// @param headers - PricingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PricingResponse
func (client *Client) PricingWithOptions(request *PricingRequest, headers *PricingHeaders, runtime *util.RuntimeOptions) (_result *PricingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SolutionId)) {
		body["solution_id"] = request.SolutionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Pricing"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/trade/action-pricing"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PricingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Trade - Seat and Price Verification
//
// Description:
//
// Check is price and remaining seats of solution you selected has changed. You should enter the solution_id returned by enrich.
//
// @param request - PricingRequest
//
// @return PricingResponse
func (client *Client) Pricing(request *PricingRequest) (_result *PricingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PricingHeaders{}
	_result = &PricingResponse{}
	_body, _err := client.PricingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 退票-申请
//
// Description:
//
// Apply for a refund and generate a refund order.
//
// @param tmpReq - RefundApplyRequest
//
// @param headers - RefundApplyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefundApplyResponse
func (client *Client) RefundApplyWithOptions(tmpReq *RefundApplyRequest, headers *RefundApplyHeaders, runtime *util.RuntimeOptions) (_result *RefundApplyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RefundApplyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RefundJourneys)) {
		request.RefundJourneysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RefundJourneys, tea.String("refund_journeys"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RefundPassengerList)) {
		request.RefundPassengerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RefundPassengerList, tea.String("refund_passenger_list"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RefundType)) {
		request.RefundTypeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RefundType, tea.String("refund_type"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderNum)) {
		body["order_num"] = request.OrderNum
	}

	if !tea.BoolValue(util.IsUnset(request.RefundJourneysShrink)) {
		body["refund_journeys"] = request.RefundJourneysShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RefundPassengerListShrink)) {
		body["refund_passenger_list"] = request.RefundPassengerListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RefundTypeShrink)) {
		body["refund_type"] = request.RefundTypeShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RefundApply"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/refund/action-apply"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefundApplyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 退票-申请
//
// Description:
//
// Apply for a refund and generate a refund order.
//
// @param request - RefundApplyRequest
//
// @return RefundApplyResponse
func (client *Client) RefundApply(request *RefundApplyRequest) (_result *RefundApplyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RefundApplyHeaders{}
	_result = &RefundApplyResponse{}
	_body, _err := client.RefundApplyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Refund - Detail
//
// Description:
//
// Query refund order detail.
//
// @param request - RefundDetailRequest
//
// @param headers - RefundDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefundDetailResponse
func (client *Client) RefundDetailWithOptions(request *RefundDetailRequest, headers *RefundDetailHeaders, runtime *util.RuntimeOptions) (_result *RefundDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RefundOrderNum)) {
		query["refund_order_num"] = request.RefundOrderNum
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefundDetail"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/refund/detail"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RefundDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Refund - Detail
//
// Description:
//
// Query refund order detail.
//
// @param request - RefundDetailRequest
//
// @return RefundDetailResponse
func (client *Client) RefundDetail(request *RefundDetailRequest) (_result *RefundDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RefundDetailHeaders{}
	_result = &RefundDetailResponse{}
	_body, _err := client.RefundDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Refund - Detail List
//
// Description:
//
// Query refund order detail.
//
// @param request - RefundDetailListRequest
//
// @param headers - RefundDetailListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefundDetailListResponse
func (client *Client) RefundDetailListWithOptions(request *RefundDetailListRequest, headers *RefundDetailListHeaders, runtime *util.RuntimeOptions) (_result *RefundDetailListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderNum)) {
		query["order_num"] = request.OrderNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		query["page_index"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["page_size"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RefundCreateBeginTime)) {
		query["refund_create_begin_time"] = request.RefundCreateBeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.RefundCreateEndTime)) {
		query["refund_create_end_time"] = request.RefundCreateEndTime
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefundDetailList"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/refund/detail-list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RefundDetailListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Refund - Detail List
//
// Description:
//
// Query refund order detail.
//
// @param request - RefundDetailListRequest
//
// @return RefundDetailListResponse
func (client *Client) RefundDetailList(request *RefundDetailListRequest) (_result *RefundDetailListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RefundDetailListHeaders{}
	_result = &RefundDetailListResponse{}
	_body, _err := client.RefundDetailListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// search
//
// Description:
//
// Enter the information of departure, arrival, departure date, passenger number and cabin, return the lowest price for each flight.
//
// @param tmpReq - SearchRequest
//
// @param headers - SearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchResponse
func (client *Client) SearchWithOptions(tmpReq *SearchRequest, headers *SearchHeaders, runtime *util.RuntimeOptions) (_result *SearchResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AirLegs)) {
		request.AirLegsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AirLegs, tea.String("air_legs"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SearchControlOptions)) {
		request.SearchControlOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SearchControlOptions, tea.String("search_control_options"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Adults)) {
		body["adults"] = request.Adults
	}

	if !tea.BoolValue(util.IsUnset(request.AirLegsShrink)) {
		body["air_legs"] = request.AirLegsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CabinClass)) {
		body["cabin_class"] = request.CabinClass
	}

	if !tea.BoolValue(util.IsUnset(request.Children)) {
		body["children"] = request.Children
	}

	if !tea.BoolValue(util.IsUnset(request.Infants)) {
		body["infants"] = request.Infants
	}

	if !tea.BoolValue(util.IsUnset(request.SearchControlOptionsShrink)) {
		body["search_control_options"] = request.SearchControlOptionsShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Search"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/trade/action-search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// search
//
// Description:
//
// Enter the information of departure, arrival, departure date, passenger number and cabin, return the lowest price for each flight.
//
// @param request - SearchRequest
//
// @return SearchResponse
func (client *Client) Search(request *SearchRequest) (_result *SearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchHeaders{}
	_result = &SearchResponse{}
	_body, _err := client.SearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Transaction - Payment and Ticket Issuance
//
// @param request - TicketingRequest
//
// @param headers - TicketingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TicketingResponse
func (client *Client) TicketingWithOptions(request *TicketingRequest, headers *TicketingHeaders, runtime *util.RuntimeOptions) (_result *TicketingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderNum)) {
		body["order_num"] = request.OrderNum
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Ticketing"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/trade/action-ticketing"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TicketingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Transaction - Payment and Ticket Issuance
//
// @param request - TicketingRequest
//
// @return TicketingResponse
func (client *Client) Ticketing(request *TicketingRequest) (_result *TicketingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TicketingHeaders{}
	_result = &TicketingResponse{}
	_body, _err := client.TicketingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Transaction - Pre-payment verification
//
// Description:
//
// Pre-check for Ticketing, this interface is optional to use.
//
// @param request - TicketingCheckRequest
//
// @param headers - TicketingCheckHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TicketingCheckResponse
func (client *Client) TicketingCheckWithOptions(request *TicketingCheckRequest, headers *TicketingCheckHeaders, runtime *util.RuntimeOptions) (_result *TicketingCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderNum)) {
		body["order_num"] = request.OrderNum
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TicketingCheck"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/trade/action-ticketing-check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TicketingCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Transaction - Pre-payment verification
//
// Description:
//
// Pre-check for Ticketing, this interface is optional to use.
//
// @param request - TicketingCheckRequest
//
// @return TicketingCheckResponse
func (client *Client) TicketingCheck(request *TicketingCheckRequest) (_result *TicketingCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TicketingCheckHeaders{}
	_result = &TicketingCheckResponse{}
	_body, _err := client.TicketingCheckWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 航程过境签
//
// @param tmpReq - TransitVisaRequest
//
// @param headers - TransitVisaHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransitVisaResponse
func (client *Client) TransitVisaWithOptions(tmpReq *TransitVisaRequest, headers *TransitVisaHeaders, runtime *util.RuntimeOptions) (_result *TransitVisaResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &TransitVisaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FlightSegmentParamList)) {
		request.FlightSegmentParamListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FlightSegmentParamList, tea.String("flight_segment_param_list"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FlightSegmentParamListShrink)) {
		query["flight_segment_param_list"] = request.FlightSegmentParamListShrink
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketAccessToken)) {
		realHeaders["x-acs-airticket-access-token"] = util.ToJSONString(headers.XAcsAirticketAccessToken)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsAirticketLanguage)) {
		realHeaders["x-acs-airticket-language"] = util.ToJSONString(headers.XAcsAirticketLanguage)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TransitVisa"),
		Version:     tea.String("2023-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/airticket/v1/flight-data/transit-visa"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TransitVisaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 航程过境签
//
// @param request - TransitVisaRequest
//
// @return TransitVisaResponse
func (client *Client) TransitVisa(request *TransitVisaRequest) (_result *TransitVisaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TransitVisaHeaders{}
	_result = &TransitVisaResponse{}
	_body, _err := client.TransitVisaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
