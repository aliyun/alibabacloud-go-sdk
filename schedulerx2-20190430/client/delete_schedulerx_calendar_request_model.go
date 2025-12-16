// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchedulerxCalendarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarName(v string) *DeleteSchedulerxCalendarRequest
	GetCalendarName() *string
	SetRegionId(v string) *DeleteSchedulerxCalendarRequest
	GetRegionId() *string
	SetYear(v int32) *DeleteSchedulerxCalendarRequest
	GetYear() *int32
}

type DeleteSchedulerxCalendarRequest struct {
	// The calendar name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025workday
	CalendarName *string `json:"CalendarName,omitempty" xml:"CalendarName,omitempty"`
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

func (s DeleteSchedulerxCalendarRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchedulerxCalendarRequest) GoString() string {
	return s.String()
}

func (s *DeleteSchedulerxCalendarRequest) GetCalendarName() *string {
	return s.CalendarName
}

func (s *DeleteSchedulerxCalendarRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSchedulerxCalendarRequest) GetYear() *int32 {
	return s.Year
}

func (s *DeleteSchedulerxCalendarRequest) SetCalendarName(v string) *DeleteSchedulerxCalendarRequest {
	s.CalendarName = &v
	return s
}

func (s *DeleteSchedulerxCalendarRequest) SetRegionId(v string) *DeleteSchedulerxCalendarRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSchedulerxCalendarRequest) SetYear(v int32) *DeleteSchedulerxCalendarRequest {
	s.Year = &v
	return s
}

func (s *DeleteSchedulerxCalendarRequest) Validate() error {
	return dara.Validate(s)
}
