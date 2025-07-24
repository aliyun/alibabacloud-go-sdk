// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iByTimeScalingRuleSpec interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ByTimeScalingRuleSpec
	GetEndTime() *int64
	SetLaunchTime(v int64) *ByTimeScalingRuleSpec
	GetLaunchTime() *int64
	SetRecurrenceType(v string) *ByTimeScalingRuleSpec
	GetRecurrenceType() *string
	SetRecurrenceValue(v string) *ByTimeScalingRuleSpec
	GetRecurrenceValue() *string
}

type ByTimeScalingRuleSpec struct {
	// 重复执行定时任务的结束时间戳。单位为毫秒。
	//
	// example:
	//
	// 1639714800000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 启动时间戳。单位为毫秒。
	//
	// This parameter is required.
	//
	// example:
	//
	// 1639714634819
	LaunchTime *int64 `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	// 指定时间规则的执行类型。
	//
	// example:
	//
	// WEEKLY
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// 重复执行定时任务的数值。具体取值取决于 recurrenceType 设置。
	//
	// example:
	//
	// MON,FRI,SUN
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
}

func (s ByTimeScalingRuleSpec) String() string {
	return dara.Prettify(s)
}

func (s ByTimeScalingRuleSpec) GoString() string {
	return s.String()
}

func (s *ByTimeScalingRuleSpec) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ByTimeScalingRuleSpec) GetLaunchTime() *int64 {
	return s.LaunchTime
}

func (s *ByTimeScalingRuleSpec) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *ByTimeScalingRuleSpec) GetRecurrenceValue() *string {
	return s.RecurrenceValue
}

func (s *ByTimeScalingRuleSpec) SetEndTime(v int64) *ByTimeScalingRuleSpec {
	s.EndTime = &v
	return s
}

func (s *ByTimeScalingRuleSpec) SetLaunchTime(v int64) *ByTimeScalingRuleSpec {
	s.LaunchTime = &v
	return s
}

func (s *ByTimeScalingRuleSpec) SetRecurrenceType(v string) *ByTimeScalingRuleSpec {
	s.RecurrenceType = &v
	return s
}

func (s *ByTimeScalingRuleSpec) SetRecurrenceValue(v string) *ByTimeScalingRuleSpec {
	s.RecurrenceValue = &v
	return s
}

func (s *ByTimeScalingRuleSpec) Validate() error {
	return dara.Validate(s)
}
