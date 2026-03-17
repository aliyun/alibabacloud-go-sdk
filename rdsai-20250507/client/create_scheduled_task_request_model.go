// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateScheduledTaskRequest
	GetDescription() *string
	SetFrequency(v string) *CreateScheduledTaskRequest
	GetFrequency() *string
	SetInstanceIds(v string) *CreateScheduledTaskRequest
	GetInstanceIds() *string
	SetName(v string) *CreateScheduledTaskRequest
	GetName() *string
	SetRegionId(v string) *CreateScheduledTaskRequest
	GetRegionId() *string
	SetReportLanguage(v string) *CreateScheduledTaskRequest
	GetReportLanguage() *string
	SetReportType(v string) *CreateScheduledTaskRequest
	GetReportType() *string
	SetStartTime(v string) *CreateScheduledTaskRequest
	GetStartTime() *string
	SetTimeRange(v string) *CreateScheduledTaskRequest
	GetTimeRange() *string
}

type CreateScheduledTaskRequest struct {
	// The description of the scheduled inspection task.
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
	// 	- Saturday \\*Sunday
	//
	// ### [](#daily--dailymonday--daily-)Note: DAILY takes precedence over other values. For example, if you enter DAILY,Monday, the backend uses DAILY as the inspection frequency.
	//
	// example:
	//
	// Monday
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The IDs of the related instances. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// rm-2ze6mk259v322****,rm-2zef3b65430j0****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The name of the scheduled inspection task. The name cannot exceed 64 characters in length.
	//
	// This parameter is required.
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReportLanguage *string `json:"ReportLanguage,omitempty" xml:"ReportLanguage,omitempty"`
	ReportType     *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The time when the inspection task is executed. Specify the time in the ISO 8601 standard in the HH:mm:ssZ format. The time must be in UTC. Default value: 02:00 AM.
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

func (s CreateScheduledTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateScheduledTaskRequest) GetFrequency() *string {
	return s.Frequency
}

func (s *CreateScheduledTaskRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *CreateScheduledTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateScheduledTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateScheduledTaskRequest) GetReportLanguage() *string {
	return s.ReportLanguage
}

func (s *CreateScheduledTaskRequest) GetReportType() *string {
	return s.ReportType
}

func (s *CreateScheduledTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateScheduledTaskRequest) GetTimeRange() *string {
	return s.TimeRange
}

func (s *CreateScheduledTaskRequest) SetDescription(v string) *CreateScheduledTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetFrequency(v string) *CreateScheduledTaskRequest {
	s.Frequency = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetInstanceIds(v string) *CreateScheduledTaskRequest {
	s.InstanceIds = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetName(v string) *CreateScheduledTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetRegionId(v string) *CreateScheduledTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetReportLanguage(v string) *CreateScheduledTaskRequest {
	s.ReportLanguage = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetReportType(v string) *CreateScheduledTaskRequest {
	s.ReportType = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetStartTime(v string) *CreateScheduledTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetTimeRange(v string) *CreateScheduledTaskRequest {
	s.TimeRange = &v
	return s
}

func (s *CreateScheduledTaskRequest) Validate() error {
	return dara.Validate(s)
}
