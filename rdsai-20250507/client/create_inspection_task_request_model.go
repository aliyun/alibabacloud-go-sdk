// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInspectionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *CreateInspectionTaskRequest
	GetEndTime() *string
	SetInspectionItems(v string) *CreateInspectionTaskRequest
	GetInspectionItems() *string
	SetInstanceIds(v string) *CreateInspectionTaskRequest
	GetInstanceIds() *string
	SetRegionId(v string) *CreateInspectionTaskRequest
	GetRegionId() *string
	SetReportLanguage(v string) *CreateInspectionTaskRequest
	GetReportLanguage() *string
	SetReportRegionId(v string) *CreateInspectionTaskRequest
	GetReportRegionId() *string
	SetReportType(v string) *CreateInspectionTaskRequest
	GetReportType() *string
	SetStartTime(v string) *CreateInspectionTaskRequest
	GetStartTime() *string
}

type CreateInspectionTaskRequest struct {
	// The end of the inspection time range. The time must be in UTC and formatted as YYYY-MM-DDTHH:mm:ssZ. If StartTime and EndTime are not specified, the inspection covers the last 24 hours.
	//
	// example:
	//
	// 2026-01-30T02:10:48Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The inspection items to run, separated by commas. If this parameter is omitted, all inspection items are run.
	//
	// ### Inspection items
	//
	// - `instance_info` (instance information)
	//
	// - `resource_usage` (resource usage)
	//
	// - `connection_session_management` (connection and session management)
	//
	// - `performance_metrics` (performance metrics)
	//
	// - `slow_query_analysis` (slow query analysis)
	//
	// - `error_log_analysis` (error log analysis)
	//
	// - `lock_wait_deadlock_analysis` (lock wait and deadlock analysis)
	//
	// - `backup_recovery_analysis` (backup and recovery analysis)
	//
	// - `high_availability_disaster_recovery_analysis` (high availability and disaster recovery inspection)
	//
	// - `security_configuration_analysis` (security configuration inspection)
	//
	// - `storage_engine_analysis` (storage engine inspection)
	//
	// - `schema_object_analysis` (schema and object inspection)
	//
	// example:
	//
	// instance_info, resource_usage
	InspectionItems *string `json:"InspectionItems,omitempty" xml:"InspectionItems,omitempty"`
	// The IDs of the instances to inspect. Separate multiple instance IDs with a comma.
	//
	// example:
	//
	// rm-2ze6mk259v322****,rm-2zef3b65430j0****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The language of the inspection report. Valid values are zh-CN (Simplified Chinese) and en-US (English). The default value is en-US.
	ReportLanguage *string `json:"ReportLanguage,omitempty" xml:"ReportLanguage,omitempty"`
	ReportRegionId *string `json:"ReportRegionId,omitempty" xml:"ReportRegionId,omitempty"`
	// The format of the inspection report. Valid values are pdf and json. The default value is pdf.
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The beginning of the inspection time range. The time must be in UTC and formatted as YYYY-MM-DDTHH:mm:ssZ. If StartTime and EndTime are not specified, the inspection covers the last 24 hours.
	//
	// example:
	//
	// 2025-12-28T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateInspectionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInspectionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateInspectionTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateInspectionTaskRequest) GetInspectionItems() *string {
	return s.InspectionItems
}

func (s *CreateInspectionTaskRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *CreateInspectionTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateInspectionTaskRequest) GetReportLanguage() *string {
	return s.ReportLanguage
}

func (s *CreateInspectionTaskRequest) GetReportRegionId() *string {
	return s.ReportRegionId
}

func (s *CreateInspectionTaskRequest) GetReportType() *string {
	return s.ReportType
}

func (s *CreateInspectionTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateInspectionTaskRequest) SetEndTime(v string) *CreateInspectionTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreateInspectionTaskRequest) SetInspectionItems(v string) *CreateInspectionTaskRequest {
	s.InspectionItems = &v
	return s
}

func (s *CreateInspectionTaskRequest) SetInstanceIds(v string) *CreateInspectionTaskRequest {
	s.InstanceIds = &v
	return s
}

func (s *CreateInspectionTaskRequest) SetRegionId(v string) *CreateInspectionTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInspectionTaskRequest) SetReportLanguage(v string) *CreateInspectionTaskRequest {
	s.ReportLanguage = &v
	return s
}

func (s *CreateInspectionTaskRequest) SetReportRegionId(v string) *CreateInspectionTaskRequest {
	s.ReportRegionId = &v
	return s
}

func (s *CreateInspectionTaskRequest) SetReportType(v string) *CreateInspectionTaskRequest {
	s.ReportType = &v
	return s
}

func (s *CreateInspectionTaskRequest) SetStartTime(v string) *CreateInspectionTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateInspectionTaskRequest) Validate() error {
	return dara.Validate(s)
}
