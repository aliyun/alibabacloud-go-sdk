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
	SetExclusions(v []*MaintenanceWindowExclusions) *MaintenanceWindow
	GetExclusions() []*MaintenanceWindowExclusions
	SetMaintenanceTime(v string) *MaintenanceWindow
	GetMaintenanceTime() *string
	SetRecurrence(v string) *MaintenanceWindow
	GetRecurrence() *string
	SetWeeklyPeriod(v string) *MaintenanceWindow
	GetWeeklyPeriod() *string
}

type MaintenanceWindow struct {
	// The maintenance duration. Valid values: 1 to 24. Unit: hours.
	//
	// Default value: 3h.
	//
	// example:
	//
	// 3h
	Duration *string `json:"duration,omitempty" xml:"duration,omitempty"`
	// Specifies whether to enable the maintenance window. Valid values:
	//
	// - `true`: Enables the maintenance window.
	//
	// - `false`: Does not enable the maintenance window.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The list of exclusions for the maintenance window. A maximum of 10 exclusions are supported. Exclusions take precedence over the maintenance window. If an active exclusion covers an O&M type, the O&M operation is not performed even if the current time is within the maintenance window. If multiple exclusions are active at the same time, the union of the exclusions is applied. The total number of NoOperations and NoUpgrades exclusions cannot exceed 3. Only one-time periods are supported. Periodic configurations are not supported.
	Exclusions []*MaintenanceWindowExclusions `json:"exclusions,omitempty" xml:"exclusions,omitempty" type:"Repeated"`
	// The start time of the maintenance window. The time must be in the RFC 3339 standard format.
	//
	// example:
	//
	// 2020-10-15T12:31:00.000+08:00
	MaintenanceTime *string `json:"maintenance_time,omitempty" xml:"maintenance_time,omitempty"`
	// The recurrence rule for the maintenance window, defined by using the RFC 5545 Recurrence Rule syntax. Only FREQ=WEEKLY is supported. COUNT and UNTIL are not supported.
	//
	// example:
	//
	// FREQ=WEEKLY;INTERVAL=4;BYDAY=MO,TU
	Recurrence *string `json:"recurrence,omitempty" xml:"recurrence,omitempty"`
	// The weekly maintenance period. Separate multiple values with commas (,). Valid values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, and Sunday.
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

func (s *MaintenanceWindow) GetExclusions() []*MaintenanceWindowExclusions {
	return s.Exclusions
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

func (s *MaintenanceWindow) SetExclusions(v []*MaintenanceWindowExclusions) *MaintenanceWindow {
	s.Exclusions = v
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
	if s.Exclusions != nil {
		for _, item := range s.Exclusions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MaintenanceWindowExclusions struct {
	// The description of the exclusion. The description can be up to 128 characters in length and does not support 4-byte characters such as emojis.
	//
	// example:
	//
	// No changes are allowed during the change freeze period
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The end time of the exclusion. The time must be in the RFC 3339 standard format, later than the start_time, and later than the current time. Maximum duration for a single exclusion: 30 days for NoOperations, NoUpgrades, and NoCVEPatch, and 180 days for NoMinorUpgrades.
	//
	// example:
	//
	// 2026-11-15T00:00:00+08:00
	EndTime *string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// The name of the exclusion. The name must be 1 to 63 characters in length and can contain digits, underscores (_), and hyphens (-). It must start with an uppercase or lowercase letter, a Chinese character, or a digit. The name must be unique within the same maintenance window. If you do not specify this parameter, a default name in the format of exclusion-{random string} is generated.
	//
	// example:
	//
	// exclusion-holiday
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The scope of the exclusion, which specifies the proactive O&M types that are prohibited during this time period. Valid values (case-sensitive):
	//
	// - NoOperations: Prohibits all proactive O&M operations, including cluster upgrades, component upgrades, CVE patching, and OS image updates.
	//
	// - NoUpgrades: Prohibits cluster upgrades and component upgrades, but still allows CVE patching and responses to scheduled ECS events.
	//
	// - NoMinorUpgrades: Prohibits only minor version upgrades for clusters, but still allows patch upgrades, component upgrades, and CVE patching.
	//
	// - NoCVEPatch: Prohibits automatic CVE patching.
	//
	// example:
	//
	// NoOperations
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The start time of the exclusion. The time must be in the RFC 3339 standard format and use the same time zone offset as the end_time and maintenance_time.
	//
	// example:
	//
	// 2026-11-01T00:00:00+08:00
	StartTime *string `json:"start_time,omitempty" xml:"start_time,omitempty"`
}

func (s MaintenanceWindowExclusions) String() string {
	return dara.Prettify(s)
}

func (s MaintenanceWindowExclusions) GoString() string {
	return s.String()
}

func (s *MaintenanceWindowExclusions) GetDescription() *string {
	return s.Description
}

func (s *MaintenanceWindowExclusions) GetEndTime() *string {
	return s.EndTime
}

func (s *MaintenanceWindowExclusions) GetName() *string {
	return s.Name
}

func (s *MaintenanceWindowExclusions) GetScope() *string {
	return s.Scope
}

func (s *MaintenanceWindowExclusions) GetStartTime() *string {
	return s.StartTime
}

func (s *MaintenanceWindowExclusions) SetDescription(v string) *MaintenanceWindowExclusions {
	s.Description = &v
	return s
}

func (s *MaintenanceWindowExclusions) SetEndTime(v string) *MaintenanceWindowExclusions {
	s.EndTime = &v
	return s
}

func (s *MaintenanceWindowExclusions) SetName(v string) *MaintenanceWindowExclusions {
	s.Name = &v
	return s
}

func (s *MaintenanceWindowExclusions) SetScope(v string) *MaintenanceWindowExclusions {
	s.Scope = &v
	return s
}

func (s *MaintenanceWindowExclusions) SetStartTime(v string) *MaintenanceWindowExclusions {
	s.StartTime = &v
	return s
}

func (s *MaintenanceWindowExclusions) Validate() error {
	return dara.Validate(s)
}
