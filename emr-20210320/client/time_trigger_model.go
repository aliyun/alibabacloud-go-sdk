// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTimeTrigger interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *TimeTrigger
	GetEndTime() *int64
	SetLaunchExpirationTime(v int32) *TimeTrigger
	GetLaunchExpirationTime() *int32
	SetLaunchTime(v string) *TimeTrigger
	GetLaunchTime() *string
	SetRecurrenceType(v string) *TimeTrigger
	GetRecurrenceType() *string
	SetRecurrenceValue(v string) *TimeTrigger
	GetRecurrenceValue() *string
	SetStartTime(v int64) *TimeTrigger
	GetStartTime() *int64
}

type TimeTrigger struct {
	// 结束时间戳。单位为毫秒。
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
	// 启动时间。
	//
	// This parameter is required.
	LaunchTime *string `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
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
	// 开始时间戳。单位为毫秒。
	//
	// example:
	//
	// 1639714800000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s TimeTrigger) String() string {
	return dara.Prettify(s)
}

func (s TimeTrigger) GoString() string {
	return s.String()
}

func (s *TimeTrigger) GetEndTime() *int64 {
	return s.EndTime
}

func (s *TimeTrigger) GetLaunchExpirationTime() *int32 {
	return s.LaunchExpirationTime
}

func (s *TimeTrigger) GetLaunchTime() *string {
	return s.LaunchTime
}

func (s *TimeTrigger) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *TimeTrigger) GetRecurrenceValue() *string {
	return s.RecurrenceValue
}

func (s *TimeTrigger) GetStartTime() *int64 {
	return s.StartTime
}

func (s *TimeTrigger) SetEndTime(v int64) *TimeTrigger {
	s.EndTime = &v
	return s
}

func (s *TimeTrigger) SetLaunchExpirationTime(v int32) *TimeTrigger {
	s.LaunchExpirationTime = &v
	return s
}

func (s *TimeTrigger) SetLaunchTime(v string) *TimeTrigger {
	s.LaunchTime = &v
	return s
}

func (s *TimeTrigger) SetRecurrenceType(v string) *TimeTrigger {
	s.RecurrenceType = &v
	return s
}

func (s *TimeTrigger) SetRecurrenceValue(v string) *TimeTrigger {
	s.RecurrenceValue = &v
	return s
}

func (s *TimeTrigger) SetStartTime(v int64) *TimeTrigger {
	s.StartTime = &v
	return s
}

func (s *TimeTrigger) Validate() error {
	return dara.Validate(s)
}
