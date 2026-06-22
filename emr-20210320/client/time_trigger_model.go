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
	// The timestamp that specifies the end time. Unit: milliseconds.
	//
	// example:
	//
	// 1639714800000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time period during which a scheduled task can be retried after it fails. Unit: seconds. Valid values: 0 to 3600.
	//
	// example:
	//
	// 600
	LaunchExpirationTime *int32 `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	// The execution time of the scaling rule. This parameter is required. The value is a string in the HH:MM format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 17:30
	LaunchTime *string `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	// The frequency of executing the specified rule whose trigger mode is scaling by time. Valid values:
	//
	// 	- DAILY
	//
	// 	- WEEKLY
	//
	// 	- MONTHLY
	//
	// example:
	//
	// WEEKLY
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The number of recurrences of the scheduled task. The value of this parameter depends on the value of RecurrenceType.
	//
	// 	- If the RecurrenceType parameter is set to DAILY, you can specify only one value for this parameter. Valid values: 1 to 31.
	//
	// 	- If the RecurrenceType parameter is set to WEEKLY, you can specify multiple values for this parameter and separate them with commas (,). The values MON, TUE, WED, THU, FRI, SAT, and SUN indicate the days from Monday to Sunday. For example, the value MON,FRI,SUN stands for Monday, Friday, and Sunday.
	//
	// 	- If the RecurrenceType parameter is set to MONTHLY, the value of this parameter is in the A-B or A,B format. The values of A and B are both in the range of 1 to 31. If you use the A-B format, the value of B must be greater than the value of A.
	//
	// example:
	//
	// MON,FRI,SUN
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	// The timestamp that specifies the start time. This parameter is required. Unit: milliseconds.
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
