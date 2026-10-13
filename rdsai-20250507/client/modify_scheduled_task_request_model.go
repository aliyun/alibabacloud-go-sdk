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
	SetTemplateId(v string) *ModifyScheduledTaskRequest
	GetTemplateId() *string
	SetTimeRange(v string) *ModifyScheduledTaskRequest
	GetTimeRange() *string
}

type ModifyScheduledTaskRequest struct {
	// The new description of the inspection configuration.
	//
	// example:
	//
	// Scheduled RDS instance inspection task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The new inspection frequency. Separate multiple values with commas (,). Default value: DAILY. Valid values:
	//
	// 	- DAILY: every day
	//
	// 	- Monday: Monday
	//
	// 	- Tuesday: Tuesday
	//
	// 	- Wednesday: Wednesday
	//
	// 	- Thursday: Thursday
	//
	// 	- Friday: Friday
	//
	// 	- Saturday: Saturday
	//
	// 	- Sunday: Sunday
	//
	// ### Note: DAILY overrides weekly values. For example, if you specify DAILY,Monday, the backend uses DAILY as the inspection frequency.
	//
	// example:
	//
	// Monday
	Frequency       *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	InspectionItems *string `json:"InspectionItems,omitempty" xml:"InspectionItems,omitempty"`
	// The new list of associated instance IDs. Separate multiple values with commas (,).
	//
	// example:
	//
	// rm-2ze6mk259v322****,rm-2zef3b65430j0****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The new name of the inspection configuration.
	//
	// example:
	//
	// RDS inspection task
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
	// The new time to execute the inspection task. Format: HH:mm:ssZ (UTC).
	//
	// example:
	//
	// 02:00:00Z
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The inspection time range. Default value: the last 24 hours. Valid values: 1 to 168 (up to 7 days).
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

func (s *ModifyScheduledTaskRequest) GetTemplateId() *string {
	return s.TemplateId
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

func (s *ModifyScheduledTaskRequest) SetTemplateId(v string) *ModifyScheduledTaskRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetTimeRange(v string) *ModifyScheduledTaskRequest {
	s.TimeRange = &v
	return s
}

func (s *ModifyScheduledTaskRequest) Validate() error {
	return dara.Validate(s)
}
