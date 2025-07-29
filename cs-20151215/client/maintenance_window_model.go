// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMaintenanceWindow interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v string) *MaintenanceWindow
	GetDuration() *string
	SetEnable(v bool) *MaintenanceWindow
	GetEnable() *bool
	SetMaintenanceTime(v string) *MaintenanceWindow
	GetMaintenanceTime() *string
	SetRecurrence(v string) *MaintenanceWindow
	GetRecurrence() *string
	SetWeeklyPeriod(v string) *MaintenanceWindow
	GetWeeklyPeriod() *string
}

type MaintenanceWindow struct {
	// example:
	//
	// 3h
	Duration *string `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 2020-10-15T12:31:00.000+08:00
	MaintenanceTime *string `json:"maintenance_time,omitempty" xml:"maintenance_time,omitempty"`
	// example:
	//
	// FREQ=WEEKLY;INTERVAL=4;BYDAY=MO,TU
	Recurrence *string `json:"recurrence,omitempty" xml:"recurrence,omitempty"`
	// example:
	//
	// Monday,Thursday
	WeeklyPeriod *string `json:"weekly_period,omitempty" xml:"weekly_period,omitempty"`
}

func (s MaintenanceWindow) String() string {
	return dara.Prettify(s)
}

func (s MaintenanceWindow) GoString() string {
	return s.String()
}

func (s *MaintenanceWindow) GetDuration() *string {
	return s.Duration
}

func (s *MaintenanceWindow) GetEnable() *bool {
	return s.Enable
}

func (s *MaintenanceWindow) GetMaintenanceTime() *string {
	return s.MaintenanceTime
}

func (s *MaintenanceWindow) GetRecurrence() *string {
	return s.Recurrence
}

func (s *MaintenanceWindow) GetWeeklyPeriod() *string {
	return s.WeeklyPeriod
}

func (s *MaintenanceWindow) SetDuration(v string) *MaintenanceWindow {
	s.Duration = &v
	return s
}

func (s *MaintenanceWindow) SetEnable(v bool) *MaintenanceWindow {
	s.Enable = &v
	return s
}

func (s *MaintenanceWindow) SetMaintenanceTime(v string) *MaintenanceWindow {
	s.MaintenanceTime = &v
	return s
}

func (s *MaintenanceWindow) SetRecurrence(v string) *MaintenanceWindow {
	s.Recurrence = &v
	return s
}

func (s *MaintenanceWindow) SetWeeklyPeriod(v string) *MaintenanceWindow {
	s.WeeklyPeriod = &v
	return s
}

func (s *MaintenanceWindow) Validate() error {
	return dara.Validate(s)
}
