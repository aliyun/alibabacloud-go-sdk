// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iByTimeScalingRule interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ByTimeScalingRule
	GetEndTime() *int64
	SetLaunchExpirationTime(v int32) *ByTimeScalingRule
	GetLaunchExpirationTime() *int32
	SetLaunchTime(v int64) *ByTimeScalingRule
	GetLaunchTime() *int64
	SetRecurrenceType(v string) *ByTimeScalingRule
	GetRecurrenceType() *string
	SetRecurrenceValue(v string) *ByTimeScalingRule
	GetRecurrenceValue() *string
}

type ByTimeScalingRule struct {
	// 重复执行定时任务的结束时间戳。单位为毫秒。
	//
	// example:
	//
	// 1639714800000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 定时任务触发操作失败后，在此时间内重试。单位为秒，取值范围：0~3600。
	//
	// example:
	//
	// 600
	LaunchExpirationTime *int32 `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
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
	// - recurrenceType 取 MINUTELY 时，只能填一个数值，取值范围：1~1440。
	//
	// - recurrenceType 取 HOURLY 时，只能填一个数值，取值范围：1~24。
	//
	// - recurrenceType 取 DAILY 时，只能填一个数值，取值范围：1~31。
	//
	// - recurrenceType 取 WEEKLY 时，可以填入多个值，填多个值时使用英文逗号（,）分隔。周一到周天分别用MON，TUE，WED，THU，FRI，SAT，SUN代替。 比如 MON,FRI,SUN 代表周一、周五、周天。
	//
	// - recurrenceType 取 MONTHLY 时，格式为A-B或者A,B。A、B的取值范围为1~31，如果使用A-B时B必须大于A。
	//
	// example:
	//
	// MON,FRI,SUN
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
}

func (s ByTimeScalingRule) String() string {
	return dara.Prettify(s)
}

func (s ByTimeScalingRule) GoString() string {
	return s.String()
}

func (s *ByTimeScalingRule) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ByTimeScalingRule) GetLaunchExpirationTime() *int32 {
	return s.LaunchExpirationTime
}

func (s *ByTimeScalingRule) GetLaunchTime() *int64 {
	return s.LaunchTime
}

func (s *ByTimeScalingRule) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *ByTimeScalingRule) GetRecurrenceValue() *string {
	return s.RecurrenceValue
}

func (s *ByTimeScalingRule) SetEndTime(v int64) *ByTimeScalingRule {
	s.EndTime = &v
	return s
}

func (s *ByTimeScalingRule) SetLaunchExpirationTime(v int32) *ByTimeScalingRule {
	s.LaunchExpirationTime = &v
	return s
}

func (s *ByTimeScalingRule) SetLaunchTime(v int64) *ByTimeScalingRule {
	s.LaunchTime = &v
	return s
}

func (s *ByTimeScalingRule) SetRecurrenceType(v string) *ByTimeScalingRule {
	s.RecurrenceType = &v
	return s
}

func (s *ByTimeScalingRule) SetRecurrenceValue(v string) *ByTimeScalingRule {
	s.RecurrenceValue = &v
	return s
}

func (s *ByTimeScalingRule) Validate() error {
	return dara.Validate(s)
}
