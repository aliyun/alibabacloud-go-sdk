// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScheduledTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyScheduledTaskRequest
	GetDescription() *string
	SetFrequency(v string) *ModifyScheduledTaskRequest
	GetFrequency() *string
	SetInspectionItems(v string) *ModifyScheduledTaskRequest
	GetInspectionItems() *string
	SetInstanceIds(v string) *ModifyScheduledTaskRequest
	GetInstanceIds() *string
	SetName(v string) *ModifyScheduledTaskRequest
	GetName() *string
	SetReportLanguage(v string) *ModifyScheduledTaskRequest
	GetReportLanguage() *string
	SetScheduledId(v string) *ModifyScheduledTaskRequest
	GetScheduledId() *string
	SetStartTime(v string) *ModifyScheduledTaskRequest
	GetStartTime() *string
	SetTimeRange(v string) *ModifyScheduledTaskRequest
	GetTimeRange() *string
}

type ModifyScheduledTaskRequest struct {
	// The new description of the inspection configuration.
	//
	// example:
	//
	// 定时RDS实例巡检任务
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The new inspection frequency. Separate multiple values with a comma (,). The default value is DAILY. Valid values:
	//
	// - DAILY: Every day
	//
	// - Monday: Every Monday
	//
	// - Tuesday: Every Tuesday
	//
	// - Wednesday: Every Wednesday
	//
	// - Thursday: Every Thursday
	//
	// - Friday: Every Friday
	//
	// - Saturday: Every Saturday
	//
	// - Sunday: Every Sunday
	//
	// ### Note: `DAILY` overrides all other day-of-the-week settings. For example, if you specify `DAILY,Monday`, the system uses `DAILY` as the inspection frequency.
	//
	// example:
	//
	// Monday
	Frequency       *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	InspectionItems *string `json:"InspectionItems,omitempty" xml:"InspectionItems,omitempty"`
	// The new instance IDs to associate with the task. Separate multiple IDs with a comma (,).
	//
	// example:
	//
	// rm-2ze6mk259v322****,rm-2zef3b65430j0****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The new name of the inspection configuration.
	//
	// example:
	//
	// RDS巡检任务
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ReportLanguage *string `json:"ReportLanguage,omitempty" xml:"ReportLanguage,omitempty"`
	// The ID of the scheduled inspection configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 847268a4-196f-416b-aa12-bfe0c115****
	ScheduledId *string `json:"ScheduledId,omitempty" xml:"ScheduledId,omitempty"`
	// The new time to run the inspection task. The time must be in the `HH:mm:ssZ` format and in UTC.
	//
	// example:
	//
	// 02:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The inspection time range in hours. The default is 24, which means data from the last 24 hours is inspected. Valid values: 1 to 168. The maximum supported range is 7 days.
	//
	// example:
	//
	// 24
	TimeRange *string `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
}

func (s ModifyScheduledTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyScheduledTaskRequest) GetFrequency() *string {
	return s.Frequency
}

func (s *ModifyScheduledTaskRequest) GetInspectionItems() *string {
	return s.InspectionItems
}

func (s *ModifyScheduledTaskRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *ModifyScheduledTaskRequest) GetName() *string {
	return s.Name
}

func (s *ModifyScheduledTaskRequest) GetReportLanguage() *string {
	return s.ReportLanguage
}

func (s *ModifyScheduledTaskRequest) GetScheduledId() *string {
	return s.ScheduledId
}

func (s *ModifyScheduledTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ModifyScheduledTaskRequest) GetTimeRange() *string {
	return s.TimeRange
}

func (s *ModifyScheduledTaskRequest) SetDescription(v string) *ModifyScheduledTaskRequest {
	s.Description = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetFrequency(v string) *ModifyScheduledTaskRequest {
	s.Frequency = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetInspectionItems(v string) *ModifyScheduledTaskRequest {
	s.InspectionItems = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetInstanceIds(v string) *ModifyScheduledTaskRequest {
	s.InstanceIds = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetName(v string) *ModifyScheduledTaskRequest {
	s.Name = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetReportLanguage(v string) *ModifyScheduledTaskRequest {
	s.ReportLanguage = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetScheduledId(v string) *ModifyScheduledTaskRequest {
	s.ScheduledId = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetStartTime(v string) *ModifyScheduledTaskRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetTimeRange(v string) *ModifyScheduledTaskRequest {
	s.TimeRange = &v
	return s
}

func (s *ModifyScheduledTaskRequest) Validate() error {
	return dara.Validate(s)
}
