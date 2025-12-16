// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageSchedulerxCalendarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarName(v string) *ManageSchedulerxCalendarRequest
	GetCalendarName() *string
	SetIncremental(v bool) *ManageSchedulerxCalendarRequest
	GetIncremental() *bool
	SetMonthDaysContent(v string) *ManageSchedulerxCalendarRequest
	GetMonthDaysContent() *string
	SetRegionId(v string) *ManageSchedulerxCalendarRequest
	GetRegionId() *string
	SetYear(v int32) *ManageSchedulerxCalendarRequest
	GetYear() *int32
}

type ManageSchedulerxCalendarRequest struct {
	// The calendar name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025workday
	CalendarName *string `json:"CalendarName,omitempty" xml:"CalendarName,omitempty"`
	// Specifies whether to perform an incremental update. Default value: false.
	//
	// 	- false: Updates the specified months and removes configurations for all other months.
	//
	// 	- true: Updates only the specified months and preserves existing configurations for other months.
	//
	// example:
	//
	// false
	Incremental *bool `json:"Incremental,omitempty" xml:"Incremental,omitempty"`
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

func (s ManageSchedulerxCalendarRequest) String() string {
	return dara.Prettify(s)
}

func (s ManageSchedulerxCalendarRequest) GoString() string {
	return s.String()
}

func (s *ManageSchedulerxCalendarRequest) GetCalendarName() *string {
	return s.CalendarName
}

func (s *ManageSchedulerxCalendarRequest) GetIncremental() *bool {
	return s.Incremental
}

func (s *ManageSchedulerxCalendarRequest) GetMonthDaysContent() *string {
	return s.MonthDaysContent
}

func (s *ManageSchedulerxCalendarRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ManageSchedulerxCalendarRequest) GetYear() *int32 {
	return s.Year
}

func (s *ManageSchedulerxCalendarRequest) SetCalendarName(v string) *ManageSchedulerxCalendarRequest {
	s.CalendarName = &v
	return s
}

func (s *ManageSchedulerxCalendarRequest) SetIncremental(v bool) *ManageSchedulerxCalendarRequest {
	s.Incremental = &v
	return s
}

func (s *ManageSchedulerxCalendarRequest) SetMonthDaysContent(v string) *ManageSchedulerxCalendarRequest {
	s.MonthDaysContent = &v
	return s
}

func (s *ManageSchedulerxCalendarRequest) SetRegionId(v string) *ManageSchedulerxCalendarRequest {
	s.RegionId = &v
	return s
}

func (s *ManageSchedulerxCalendarRequest) SetYear(v int32) *ManageSchedulerxCalendarRequest {
	s.Year = &v
	return s
}

func (s *ManageSchedulerxCalendarRequest) Validate() error {
	return dara.Validate(s)
}
