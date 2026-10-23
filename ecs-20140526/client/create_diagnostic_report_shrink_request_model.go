// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnosticReportShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalOptionsShrink(v string) *CreateDiagnosticReportShrinkRequest
	GetAdditionalOptionsShrink() *string
	SetEndTime(v string) *CreateDiagnosticReportShrinkRequest
	GetEndTime() *string
	SetMetricSetId(v string) *CreateDiagnosticReportShrinkRequest
	GetMetricSetId() *string
	SetRegionId(v string) *CreateDiagnosticReportShrinkRequest
	GetRegionId() *string
	SetResourceId(v string) *CreateDiagnosticReportShrinkRequest
	GetResourceId() *string
	SetStartTime(v string) *CreateDiagnosticReportShrinkRequest
	GetStartTime() *string
}

type CreateDiagnosticReportShrinkRequest struct {
	// The extension in JSON format.
	AdditionalOptionsShrink *string `json:"AdditionalOptions,omitempty" xml:"AdditionalOptions,omitempty"`
	// The end time. The time follows the ISO 8601 standard and is displayed in UTC. This parameter takes effect only for diagnostic metrics that do not require Cloud Assistant commands to be executed in the guest OS.
	//
	// example:
	//
	// 2022-07-11T14:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The diagnostic metric set ID. If this parameter is left empty, the default diagnostic metric set dms-instancedefault for ECS instances is used.
	//
	// example:
	//
	// dms-uf6i0tv2refv8wz*****
	MetricSetId *string `json:"MetricSetId,omitempty" xml:"MetricSetId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the latest Alibaba Cloud region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-uf6i0tv2refv8wz*****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The start time. The time follows the ISO 8601 standard and is displayed in UTC. This parameter takes effect only for diagnostic metrics that do not require Cloud Assistant commands to be executed in the guest OS.
	//
	// example:
	//
	// 2022-07-11T12:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateDiagnosticReportShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticReportShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticReportShrinkRequest) GetAdditionalOptionsShrink() *string {
	return s.AdditionalOptionsShrink
}

func (s *CreateDiagnosticReportShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateDiagnosticReportShrinkRequest) GetMetricSetId() *string {
	return s.MetricSetId
}

func (s *CreateDiagnosticReportShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDiagnosticReportShrinkRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateDiagnosticReportShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateDiagnosticReportShrinkRequest) SetAdditionalOptionsShrink(v string) *CreateDiagnosticReportShrinkRequest {
	s.AdditionalOptionsShrink = &v
	return s
}

func (s *CreateDiagnosticReportShrinkRequest) SetEndTime(v string) *CreateDiagnosticReportShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *CreateDiagnosticReportShrinkRequest) SetMetricSetId(v string) *CreateDiagnosticReportShrinkRequest {
	s.MetricSetId = &v
	return s
}

func (s *CreateDiagnosticReportShrinkRequest) SetRegionId(v string) *CreateDiagnosticReportShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDiagnosticReportShrinkRequest) SetResourceId(v string) *CreateDiagnosticReportShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateDiagnosticReportShrinkRequest) SetStartTime(v string) *CreateDiagnosticReportShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *CreateDiagnosticReportShrinkRequest) Validate() error {
	return dara.Validate(s)
}
