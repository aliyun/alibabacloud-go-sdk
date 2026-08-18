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
	// The duration of the cluster maintenance window.
	//
	// Valid values: 1 to 24.
	//
	// Unit: hours.
	//
	// Default value: 3.
	//
	// example:
	//
	// 3h
	Duration *string `json:"duration,omitempty" xml:"duration,omitempty"`
	// Specifies whether to enable the cluster maintenance window. Valid values:
	//
	// 	- `true`: enables the cluster maintenance window.
	//
	// 	- `false`: disables the cluster maintenance window.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The start time of the cluster maintenance window. The value follows a standard time format in Golang. Example: 15:04:05Z.
	//
	// example:
	//
	// 03:00:00Z
	MaintenanceTime *string `json:"maintenance_time,omitempty" xml:"maintenance_time,omitempty"`
	// Defines a maintenance window recurrence rule by using the RFC5545 recurrence rule. Currently, only FREQ=WEEKLY is supported, and you cannot specify COUNT or UNTIL.
	//
	// example:
	//
	// FREQ=WEEKLY;INTERVAL=4;BYDAY=MO,TU
	Recurrence *string `json:"recurrence,omitempty" xml:"recurrence,omitempty"`
	// The day of the week when maintenance is performed. Separate multiple days with commas (,).
	//
	// Valid values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, and Sunday.
	//
	// Default value: `Thursday`.
	//
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
