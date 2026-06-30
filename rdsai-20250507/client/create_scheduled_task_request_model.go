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
	SetInspectionItems(v string) *CreateScheduledTaskRequest
	GetInspectionItems() *string
	SetInstanceIds(v string) *CreateScheduledTaskRequest
	GetInstanceIds() *string
	SetName(v string) *CreateScheduledTaskRequest
	GetName() *string
	SetRegionId(v string) *CreateScheduledTaskRequest
	GetRegionId() *string
	SetReportLanguage(v string) *CreateScheduledTaskRequest
	GetReportLanguage() *string
	SetReportRegionId(v string) *CreateScheduledTaskRequest
	GetReportRegionId() *string
	SetReportType(v string) *CreateScheduledTaskRequest
	GetReportType() *string
	SetStartTime(v string) *CreateScheduledTaskRequest
	GetStartTime() *string
	SetTimeRange(v string) *CreateScheduledTaskRequest
	GetTimeRange() *string
}

type CreateScheduledTaskRequest struct {
	// The description of the scheduled inspection task.
	//
	// example:
	//
	// 定时RDS实例巡检任务
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The inspection frequency. Use commas (,) to separate multiple values. The default is DAILY. Valid values:
	//
	// - DAILY: Every day
	//
	// - Monday: Monday
	//
	// - Tuesday: Tuesday
	//
	// - Wednesday: Wednesday
	//
	// - Thursday: Thursday
	//
	// - Friday: Friday
	//
	// - Saturday: Saturday
	//
	// - Sunday: Sunday
	//
	// ### Note: DAILY overrides weekly values. For example, if you enter DAILY,Monday, the system uses DAILY as the inspection frequency.
	//
	// example:
	//
	// Monday
	Frequency       *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	InspectionItems *string `json:"InspectionItems,omitempty" xml:"InspectionItems,omitempty"`
	// The IDs of the instances for the task. Use commas (,) to separate multiple IDs.
	//
	// example:
	//
	// rm-2ze6mk259v322****,rm-2zef3b65430j0****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The name of the scheduled inspection task. The maximum length is 64 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS巡检
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The report language. The default value is zh-CN. Supported values: zh-CN, zh-TW, ja-JP, and en-US.
	//
	// example:
	//
	// zh-CN
	ReportLanguage *string `json:"ReportLanguage,omitempty" xml:"ReportLanguage,omitempty"`
	// The ID of the region where the report is stored.
	ReportRegionId *string `json:"ReportRegionId,omitempty" xml:"ReportRegionId,omitempty"`
	// The type of the report.
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The execution time for the scheduled inspection task. Specify the time in the HH:mm:ssZ format (UTC time). The default is 02:00:00Z.
	//
	// example:
	//
	// 02:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time range of data to inspect, in hours. Valid values are from 1 to 168 (7 days). The default is 24.
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

func (s *CreateScheduledTaskRequest) GetInspectionItems() *string {
	return s.InspectionItems
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

func (s *CreateScheduledTaskRequest) GetReportRegionId() *string {
	return s.ReportRegionId
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

func (s *CreateScheduledTaskRequest) SetInspectionItems(v string) *CreateScheduledTaskRequest {
	s.InspectionItems = &v
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

func (s *CreateScheduledTaskRequest) SetReportRegionId(v string) *CreateScheduledTaskRequest {
	s.ReportRegionId = &v
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
