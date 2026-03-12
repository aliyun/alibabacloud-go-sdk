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
	// The description of the new inspection configuration.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The new inspection frequency. Separate multiple values with commas (,). Default value: DAILY. Valid values:
	//
	// 	- DAILY
	//
	// 	- Monday
	//
	// 	- Tuesday
	//
	// 	- Wednesday
	//
	// 	- Thursday
	//
	// 	- Friday
	//
	// 	- Saturday
	//
	// 	- Sunday
	//
	// ### [](#daily--dailymonday--daily-)Note: DAILY takes precedence over other values. For example, if you enter DAILY,Monday, the backend will use DAILY as the inspection frequency.
	//
	// example:
	//
	// Monday
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The new list of related instances. Separate multiple instances with commas (,).
	//
	// example:
	//
	// rm-2ze6mk259v322****,rm-2zef3b65430j0****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The name of the new inspection configuration.
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
	// The new execution time of the inspection task. Specify the time in the ISO 8601 standard in the HH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 02:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The inspection time range. The default value is the latest 24 hours. Valid values: 1 to 168. The maximum value is 7 days.
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
