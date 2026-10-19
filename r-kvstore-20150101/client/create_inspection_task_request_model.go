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
	SetInstanceId(v string) *CreateInspectionTaskRequest
	GetInstanceId() *string
	SetInstanceIds(v string) *CreateInspectionTaskRequest
	GetInstanceIds() *string
	SetReportLanguage(v string) *CreateInspectionTaskRequest
	GetReportLanguage() *string
	SetStartTime(v string) *CreateInspectionTaskRequest
	GetStartTime() *string
}

type CreateInspectionTaskRequest struct {
	// The end time of the inspection time range.
	//
	// The time is in the YYYY-MM-DDTHH:mm:ssZ format and in UTC.
	//
	// example:
	//
	// 2026-07-29T06:59:26Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The specified inspection items. Valid values:
	//
	// - RESOURCE_USAGE
	//
	// - BIGKEY
	//
	// - HOTKEY
	//
	// - PERFORMANCE_METRICS
	//
	// - SLOWLOG
	//
	// This parameter is required.
	//
	// example:
	//
	// RESOURCE_USAGE
	InspectionItems *string `json:"InspectionItems,omitempty" xml:"InspectionItems,omitempty"`
	// The ID of the assistant instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ta-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of associated instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The language of the report.
	//
	// example:
	//
	// zh-CN
	ReportLanguage *string `json:"ReportLanguage,omitempty" xml:"ReportLanguage,omitempty"`
	// The start time of the inspection time range.
	//
	// The time is in the YYYY-MM-DDTHH:mm:ssZ format and in UTC.
	//
	// example:
	//
	// 2026-07-29T05:59:26Z
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

func (s *CreateInspectionTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInspectionTaskRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *CreateInspectionTaskRequest) GetReportLanguage() *string {
	return s.ReportLanguage
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

func (s *CreateInspectionTaskRequest) SetInstanceId(v string) *CreateInspectionTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateInspectionTaskRequest) SetInstanceIds(v string) *CreateInspectionTaskRequest {
	s.InstanceIds = &v
	return s
}

func (s *CreateInspectionTaskRequest) SetReportLanguage(v string) *CreateInspectionTaskRequest {
	s.ReportLanguage = &v
	return s
}

func (s *CreateInspectionTaskRequest) SetStartTime(v string) *CreateInspectionTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateInspectionTaskRequest) Validate() error {
	return dara.Validate(s)
}
