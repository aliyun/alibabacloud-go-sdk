// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSchedulerxCalendarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarName(v string) *CreateSchedulerxCalendarRequest
	GetCalendarName() *string
	SetMonthDaysContent(v string) *CreateSchedulerxCalendarRequest
	GetMonthDaysContent() *string
	SetRegionId(v string) *CreateSchedulerxCalendarRequest
	GetRegionId() *string
	SetYear(v int32) *CreateSchedulerxCalendarRequest
	GetYear() *int32
}

type CreateSchedulerxCalendarRequest struct {
	// The calendar name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025workday
	CalendarName *string `json:"CalendarName,omitempty" xml:"CalendarName,omitempty"`
	// The month and days.
	//
	// This parameter is required.
	//
	// example:
	//
	// [
	//
	//   {"month":1,"days":[3,4,5,6,9,10,11,12,13,16,17,18,19,20,28,29,30,31]},
	//
	//   {"month":2,"days":[1,2,3,6,7,8,9,10,13,14,15,16,17,20,21,22,23,24,27,28]},
	//
	//   {"month":3,"days":[1,2,3,6,7,8,9,10,13,14,15,16,17,20,21,22,23,24,27,28,29,30,31]},
	//
	//   {"month":4,"days":[3,4,6,7,10,11,12,13,14,17,18,19,20,21,23,24,25,26,27,28]},
	//
	//   {"month":5,"days":[4,5,6,8,9,10,11,12,15,16,17,18,19,22,23,24,25,26,29,30,31]},
	//
	//   {"month":6,"days":[1,2,5,6,7,8,9,12,13,14,15,16,19,20,21,25,26,27,28,29,30]},
	//
	//   {"month":7,"days":[3,4,5,6,7,10,11,12,13,14,17,18,19,20,21,24,25,26,27,28,31]},
	//
	//   {"month":8,"days":[1,2,3,4,7,8,9,10,11,14,15,16,17,18,21,22,23,24,25,28,29,30,31]},
	//
	//   {"month":9,"days":[1,4,5,6,7,8,11,12,13,14,15,18,19,20,21,22,25,26,27,28]},
	//
	//   {"month":10,"days":[7,8,9,10,11,12,13,16,17,18,19,20,23,24,25,26,27,30,31]},
	//
	//   {"month":11,"days":[1,2,3,6,7,8,9,10,13,14,15,16,17,20,21,22,23,24,27,28,29,30]},
	//
	//   {"month":12,"days":[1,4,5,6,7,8,11,12,13,14,15,18,19,20,21,22,25,26,27,28,29]}
	//
	// ]
	MonthDaysContent *string `json:"MonthDaysContent,omitempty" xml:"MonthDaysContent,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The year.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s CreateSchedulerxCalendarRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSchedulerxCalendarRequest) GoString() string {
	return s.String()
}

func (s *CreateSchedulerxCalendarRequest) GetCalendarName() *string {
	return s.CalendarName
}

func (s *CreateSchedulerxCalendarRequest) GetMonthDaysContent() *string {
	return s.MonthDaysContent
}

func (s *CreateSchedulerxCalendarRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSchedulerxCalendarRequest) GetYear() *int32 {
	return s.Year
}

func (s *CreateSchedulerxCalendarRequest) SetCalendarName(v string) *CreateSchedulerxCalendarRequest {
	s.CalendarName = &v
	return s
}

func (s *CreateSchedulerxCalendarRequest) SetMonthDaysContent(v string) *CreateSchedulerxCalendarRequest {
	s.MonthDaysContent = &v
	return s
}

func (s *CreateSchedulerxCalendarRequest) SetRegionId(v string) *CreateSchedulerxCalendarRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSchedulerxCalendarRequest) SetYear(v int32) *CreateSchedulerxCalendarRequest {
	s.Year = &v
	return s
}

func (s *CreateSchedulerxCalendarRequest) Validate() error {
	return dara.Validate(s)
}
