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
	SetTemplateId(v string) *CreateScheduledTaskRequest
	GetTemplateId() *string
	SetTimeRange(v string) *CreateScheduledTaskRequest
	GetTimeRange() *string
}

type CreateScheduledTaskRequest struct {
	// The description of the scheduled inspection task.
	//
	// example:
	//
	// Scheduled RDS instance inspection task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The inspection frequency. Separate multiple values with commas (,). Default value: DAILY. Valid values:
	//
	// 	- DAILY: every day.
	//
	// 	- Monday: Monday.
	//
	// 	- Tuesday: Tuesday.
	//
	// 	- Wednesday: Wednesday.
	//
	// 	- Thursday: Thursday.
	//
	// 	- Friday: Friday.
	//
	// 	- Saturday: Saturday.
	//
	// 	- Sunday: Sunday.
	//
	// ### Note: DAILY overrides weekly values. For example, if you specify DAILY,Monday, the system uses DAILY as the inspection frequency.
	//
	// example:
	//
	// Monday
	Frequency       *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	InspectionItems *string `json:"InspectionItems,omitempty" xml:"InspectionItems,omitempty"`
	// The list of associated instance IDs. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// rm-2ze6mk259v322****,rm-2zef3b65430j0****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The name of the scheduled inspection task. The name can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS Inspection
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The report language. Default value: zh-CN. Valid values: zh-CN, zh-TW, ja-JP, and en-US.
	//
	// example:
	//
	// zh-CN
	ReportLanguage *string `json:"ReportLanguage,omitempty" xml:"ReportLanguage,omitempty"`
	ReportRegionId *string `json:"ReportRegionId,omitempty" xml:"ReportRegionId,omitempty"`
	ReportType     *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The time to run the inspection task. Format: HH:mm:ssZ (UTC). Default value: 02:00:00Z.
	//
	// example:
	//
	// 02:00:00Z
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The time range for the inspection. Default value: the last 24 hours. Valid values: 1 to 168 (up to 7 days).
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

func (s *CreateScheduledTaskRequest) GetTemplateId() *string {
	return s.TemplateId
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

func (s *CreateScheduledTaskRequest) SetTemplateId(v string) *CreateScheduledTaskRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetTimeRange(v string) *CreateScheduledTaskRequest {
	s.TimeRange = &v
	return s
}

func (s *CreateScheduledTaskRequest) Validate() error {
	return dara.Validate(s)
}
