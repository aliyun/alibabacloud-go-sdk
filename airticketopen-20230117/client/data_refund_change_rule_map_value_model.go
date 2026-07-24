// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataRefundChangeRuleMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetRefundRuleAllUnusedList(v []*DataRefundChangeRuleMapValueRefundRuleAllUnusedList) *DataRefundChangeRuleMapValue
	GetRefundRuleAllUnusedList() []*DataRefundChangeRuleMapValueRefundRuleAllUnusedList
	SetRefundRulePartUnusedList(v []*DataRefundChangeRuleMapValueRefundRulePartUnusedList) *DataRefundChangeRuleMapValue
	GetRefundRulePartUnusedList() []*DataRefundChangeRuleMapValueRefundRulePartUnusedList
	SetChangeRuleInUnusedList(v []*DataRefundChangeRuleMapValueChangeRuleInUnusedList) *DataRefundChangeRuleMapValue
	GetChangeRuleInUnusedList() []*DataRefundChangeRuleMapValueChangeRuleInUnusedList
	SetChangeRuleOutUnusedList(v []*DataRefundChangeRuleMapValueChangeRuleOutUnusedList) *DataRefundChangeRuleMapValue
	GetChangeRuleOutUnusedList() []*DataRefundChangeRuleMapValueChangeRuleOutUnusedList
}

type DataRefundChangeRuleMapValue struct {
	// 全程未使用退票规则
	RefundRuleAllUnusedList []*DataRefundChangeRuleMapValueRefundRuleAllUnusedList `json:"refund_rule_all_unused_list,omitempty" xml:"refund_rule_all_unused_list,omitempty" type:"Repeated"`
	// 部分未使用退票规则
	RefundRulePartUnusedList []*DataRefundChangeRuleMapValueRefundRulePartUnusedList `json:"refund_rule_part_unused_list,omitempty" xml:"refund_rule_part_unused_list,omitempty" type:"Repeated"`
	// 回程未使用改签规则
	ChangeRuleInUnusedList []*DataRefundChangeRuleMapValueChangeRuleInUnusedList `json:"change_rule_in_unused_list,omitempty" xml:"change_rule_in_unused_list,omitempty" type:"Repeated"`
	// 去程未使用改签规则
	ChangeRuleOutUnusedList []*DataRefundChangeRuleMapValueChangeRuleOutUnusedList `json:"change_rule_out_unused_list,omitempty" xml:"change_rule_out_unused_list,omitempty" type:"Repeated"`
}

func (s DataRefundChangeRuleMapValue) String() string {
	return dara.Prettify(s)
}

func (s DataRefundChangeRuleMapValue) GoString() string {
	return s.String()
}

func (s *DataRefundChangeRuleMapValue) GetRefundRuleAllUnusedList() []*DataRefundChangeRuleMapValueRefundRuleAllUnusedList {
	return s.RefundRuleAllUnusedList
}

func (s *DataRefundChangeRuleMapValue) GetRefundRulePartUnusedList() []*DataRefundChangeRuleMapValueRefundRulePartUnusedList {
	return s.RefundRulePartUnusedList
}

func (s *DataRefundChangeRuleMapValue) GetChangeRuleInUnusedList() []*DataRefundChangeRuleMapValueChangeRuleInUnusedList {
	return s.ChangeRuleInUnusedList
}

func (s *DataRefundChangeRuleMapValue) GetChangeRuleOutUnusedList() []*DataRefundChangeRuleMapValueChangeRuleOutUnusedList {
	return s.ChangeRuleOutUnusedList
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

func (s *DataRefundChangeRuleMapValue) Validate() error {
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

type DataRefundChangeRuleMapValueRefundRuleAllUnusedList struct {
	// 该退票规则适用航程使用类型 0:全程未使用; 1:部分未使用
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// 时间单位：day/hour
	//
	// example:
	//
	// hour
	TimeUnit *string `json:"time_unit,omitempty" xml:"time_unit,omitempty"`
	// 该退票规则适用于的退票时间区间开始时间，单位（天/小时）
	//
	// example:
	//
	// 36
	RuleStartTime *int32 `json:"rule_start_time,omitempty" xml:"rule_start_time,omitempty"`
	// 该退票规则适用于的退票时间区间结束时间，单位（天/小时）
	//
	// example:
	//
	// 12
	RuleEndTime *int32 `json:"rule_end_time,omitempty" xml:"rule_end_time,omitempty"`
	// 起飞前X-Y小时(天) 可否退票
	//
	// example:
	//
	// true
	CanRefund *bool `json:"can_refund,omitempty" xml:"can_refund,omitempty"`
	// 起飞前X-Y小时(天) 退票手续费
	//
	// example:
	//
	// 200
	RefundFee *float64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 起飞前X-Y小时(天) 是否全额退税
	//
	// example:
	//
	// true
	CanReturnAllTax *bool `json:"can_return_all_tax,omitempty" xml:"can_return_all_tax,omitempty"`
	// 起飞前X-Y小时(天) 可部分退税金额
	//
	// example:
	//
	// 100
	ReturnPartTaxFee *float64 `json:"return_part_tax_fee,omitempty" xml:"return_part_tax_fee,omitempty"`
}

func (s DataRefundChangeRuleMapValueRefundRuleAllUnusedList) String() string {
	return dara.Prettify(s)
}

func (s DataRefundChangeRuleMapValueRefundRuleAllUnusedList) GoString() string {
	return s.String()
}

func (s *DataRefundChangeRuleMapValueRefundRuleAllUnusedList) GetType() *int32 {
	return s.Type
}

func (s *DataRefundChangeRuleMapValueRefundRuleAllUnusedList) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *DataRefundChangeRuleMapValueRefundRuleAllUnusedList) GetRuleStartTime() *int32 {
	return s.RuleStartTime
}

func (s *DataRefundChangeRuleMapValueRefundRuleAllUnusedList) GetRuleEndTime() *int32 {
	return s.RuleEndTime
}

func (s *DataRefundChangeRuleMapValueRefundRuleAllUnusedList) GetCanRefund() *bool {
	return s.CanRefund
}

func (s *DataRefundChangeRuleMapValueRefundRuleAllUnusedList) GetRefundFee() *float64 {
	return s.RefundFee
}

func (s *DataRefundChangeRuleMapValueRefundRuleAllUnusedList) GetCanReturnAllTax() *bool {
	return s.CanReturnAllTax
}

func (s *DataRefundChangeRuleMapValueRefundRuleAllUnusedList) GetReturnPartTaxFee() *float64 {
	return s.ReturnPartTaxFee
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

func (s *DataRefundChangeRuleMapValueRefundRuleAllUnusedList) Validate() error {
	return dara.Validate(s)
}

type DataRefundChangeRuleMapValueRefundRulePartUnusedList struct {
	// 该退票规则适用航程使用类型 0:全程未使用; 1:部分未使用
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// 时间单位：day/hour
	//
	// example:
	//
	// hour
	TimeUnit *string `json:"time_unit,omitempty" xml:"time_unit,omitempty"`
	// 该退票规则适用于的退票时间区间开始时间，单位（天/小时）
	//
	// example:
	//
	// 36
	RuleStartTime *int32 `json:"rule_start_time,omitempty" xml:"rule_start_time,omitempty"`
	// 该退票规则适用于的退票时间区间结束时间，单位（天/小时）
	//
	// example:
	//
	// 12
	RuleEndTime *int32 `json:"rule_end_time,omitempty" xml:"rule_end_time,omitempty"`
	// 起飞前X-Y小时(天) 可否退票
	//
	// example:
	//
	// true
	CanRefund *bool `json:"can_refund,omitempty" xml:"can_refund,omitempty"`
	// 起飞前X-Y小时(天) 退票手续费
	//
	// example:
	//
	// 200
	RefundFee *float64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 起飞前X-Y小时(天) 是否全额退税
	//
	// example:
	//
	// true
	CanReturnAllTax *bool `json:"can_return_all_tax,omitempty" xml:"can_return_all_tax,omitempty"`
	// 起飞前X-Y小时(天) 可部分退税金额
	//
	// example:
	//
	// 100
	ReturnPartTaxFee *float64 `json:"return_part_tax_fee,omitempty" xml:"return_part_tax_fee,omitempty"`
}

func (s DataRefundChangeRuleMapValueRefundRulePartUnusedList) String() string {
	return dara.Prettify(s)
}

func (s DataRefundChangeRuleMapValueRefundRulePartUnusedList) GoString() string {
	return s.String()
}

func (s *DataRefundChangeRuleMapValueRefundRulePartUnusedList) GetType() *int32 {
	return s.Type
}

func (s *DataRefundChangeRuleMapValueRefundRulePartUnusedList) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *DataRefundChangeRuleMapValueRefundRulePartUnusedList) GetRuleStartTime() *int32 {
	return s.RuleStartTime
}

func (s *DataRefundChangeRuleMapValueRefundRulePartUnusedList) GetRuleEndTime() *int32 {
	return s.RuleEndTime
}

func (s *DataRefundChangeRuleMapValueRefundRulePartUnusedList) GetCanRefund() *bool {
	return s.CanRefund
}

func (s *DataRefundChangeRuleMapValueRefundRulePartUnusedList) GetRefundFee() *float64 {
	return s.RefundFee
}

func (s *DataRefundChangeRuleMapValueRefundRulePartUnusedList) GetCanReturnAllTax() *bool {
	return s.CanReturnAllTax
}

func (s *DataRefundChangeRuleMapValueRefundRulePartUnusedList) GetReturnPartTaxFee() *float64 {
	return s.ReturnPartTaxFee
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

func (s *DataRefundChangeRuleMapValueRefundRulePartUnusedList) Validate() error {
	return dara.Validate(s)
}

type DataRefundChangeRuleMapValueChangeRuleInUnusedList struct {
	// 该改签规则适用航程使用类型 2:去程未使用; 3:回程未使用
	//
	// example:
	//
	// 2
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// 时间单位：day/hour
	//
	// example:
	//
	// hour
	TimeUnit *string `json:"time_unit,omitempty" xml:"time_unit,omitempty"`
	// 该退票规则适用于的退票时间区间开始时间，单位（天/小时）
	//
	// example:
	//
	// 36
	RuleStartTime *int32 `json:"rule_start_time,omitempty" xml:"rule_start_time,omitempty"`
	// 该退票规则适用于的退票时间区间结束时间，单位（天/小时）
	//
	// example:
	//
	// 12
	RuleEndTime *int32 `json:"rule_end_time,omitempty" xml:"rule_end_time,omitempty"`
	// 起飞前X-Y小时(天) 可否可以改签
	//
	// example:
	//
	// true
	CanChange *bool `json:"can_change,omitempty" xml:"can_change,omitempty"`
	// 起飞前X-Y小时(天) 改签费
	//
	// example:
	//
	// 100
	ChangeFee *float64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
}

func (s DataRefundChangeRuleMapValueChangeRuleInUnusedList) String() string {
	return dara.Prettify(s)
}

func (s DataRefundChangeRuleMapValueChangeRuleInUnusedList) GoString() string {
	return s.String()
}

func (s *DataRefundChangeRuleMapValueChangeRuleInUnusedList) GetType() *int32 {
	return s.Type
}

func (s *DataRefundChangeRuleMapValueChangeRuleInUnusedList) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *DataRefundChangeRuleMapValueChangeRuleInUnusedList) GetRuleStartTime() *int32 {
	return s.RuleStartTime
}

func (s *DataRefundChangeRuleMapValueChangeRuleInUnusedList) GetRuleEndTime() *int32 {
	return s.RuleEndTime
}

func (s *DataRefundChangeRuleMapValueChangeRuleInUnusedList) GetCanChange() *bool {
	return s.CanChange
}

func (s *DataRefundChangeRuleMapValueChangeRuleInUnusedList) GetChangeFee() *float64 {
	return s.ChangeFee
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

func (s *DataRefundChangeRuleMapValueChangeRuleInUnusedList) Validate() error {
	return dara.Validate(s)
}

type DataRefundChangeRuleMapValueChangeRuleOutUnusedList struct {
	// 该改签规则适用航程使用类型 2:去程未使用; 3:回程未使用
	//
	// example:
	//
	// 2
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// 时间单位：day/hour
	//
	// example:
	//
	// hour
	TimeUnit *string `json:"time_unit,omitempty" xml:"time_unit,omitempty"`
	// 该退票规则适用于的退票时间区间开始时间，单位（天/小时）
	//
	// example:
	//
	// 36
	RuleStartTime *int32 `json:"rule_start_time,omitempty" xml:"rule_start_time,omitempty"`
	// 该退票规则适用于的退票时间区间结束时间，单位（天/小时）
	//
	// example:
	//
	// 12
	RuleEndTime *int32 `json:"rule_end_time,omitempty" xml:"rule_end_time,omitempty"`
	// 起飞前X-Y小时(天) 可否可以改签
	//
	// example:
	//
	// true
	CanChange *bool `json:"can_change,omitempty" xml:"can_change,omitempty"`
	// 起飞前X-Y小时(天) 改签费
	//
	// example:
	//
	// 100
	ChangeFee *float64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
}

func (s DataRefundChangeRuleMapValueChangeRuleOutUnusedList) String() string {
	return dara.Prettify(s)
}

func (s DataRefundChangeRuleMapValueChangeRuleOutUnusedList) GoString() string {
	return s.String()
}

func (s *DataRefundChangeRuleMapValueChangeRuleOutUnusedList) GetType() *int32 {
	return s.Type
}

func (s *DataRefundChangeRuleMapValueChangeRuleOutUnusedList) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *DataRefundChangeRuleMapValueChangeRuleOutUnusedList) GetRuleStartTime() *int32 {
	return s.RuleStartTime
}

func (s *DataRefundChangeRuleMapValueChangeRuleOutUnusedList) GetRuleEndTime() *int32 {
	return s.RuleEndTime
}

func (s *DataRefundChangeRuleMapValueChangeRuleOutUnusedList) GetCanChange() *bool {
	return s.CanChange
}

func (s *DataRefundChangeRuleMapValueChangeRuleOutUnusedList) GetChangeFee() *float64 {
	return s.ChangeFee
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

func (s *DataRefundChangeRuleMapValueChangeRuleOutUnusedList) Validate() error {
	return dara.Validate(s)
}
