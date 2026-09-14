// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnoseReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDiagnoseReportRequest
	GetClientToken() *string
	SetDiagnoseType(v string) *CreateDiagnoseReportRequest
	GetDiagnoseType() *string
	SetEndTime(v string) *CreateDiagnoseReportRequest
	GetEndTime() *string
	SetRegionId(v string) *CreateDiagnoseReportRequest
	GetRegionId() *string
	SetResourceId(v string) *CreateDiagnoseReportRequest
	GetResourceId() *string
	SetResourceType(v string) *CreateDiagnoseReportRequest
	GetResourceType() *string
	SetStartTime(v string) *CreateDiagnoseReportRequest
	GetStartTime() *string
}

type CreateDiagnoseReportRequest struct {
	// A client-generated token to ensure request idempotency. This lets you safely retry the request without creating a duplicate diagnostic report. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The diagnosis type. The only valid value is:
	//
	// - Performance: performance diagnosis
	//
	// This parameter is required.
	//
	// example:
	//
	// Performance
	DiagnoseType *string `json:"DiagnoseType,omitempty" xml:"DiagnoseType,omitempty"`
	// The end time for the diagnosis, in UTC. The time must be in the ISO 8601 format (yyyy-MM-ddTHH:mm:ssZ).
	//
	// example:
	//
	// 2024-09-07T16:49:25Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The region ID. Call the [DescribeRegions](https://help.aliyun.com/zh/ecs/developer-reference/api-ebs-2021-07-30-describeregions?spm=a2c4g.11186623.0.i7) operation to find all regions supported by EBS Data Insight.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// d-asb1s8***
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. The only valid value is:
	//
	// - Disk: a cloud disk
	//
	// example:
	//
	// Disk
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The start time for the diagnosis, in UTC. The time must be in the ISO 8601 format (yyyy-MM-ddTHH:mm:ssZ).
	//
	// example:
	//
	// 2024-09-01T02:26:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateDiagnoseReportRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnoseReportRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnoseReportRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDiagnoseReportRequest) GetDiagnoseType() *string {
	return s.DiagnoseType
}

func (s *CreateDiagnoseReportRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateDiagnoseReportRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDiagnoseReportRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateDiagnoseReportRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateDiagnoseReportRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateDiagnoseReportRequest) SetClientToken(v string) *CreateDiagnoseReportRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDiagnoseReportRequest) SetDiagnoseType(v string) *CreateDiagnoseReportRequest {
	s.DiagnoseType = &v
	return s
}

func (s *CreateDiagnoseReportRequest) SetEndTime(v string) *CreateDiagnoseReportRequest {
	s.EndTime = &v
	return s
}

func (s *CreateDiagnoseReportRequest) SetRegionId(v string) *CreateDiagnoseReportRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDiagnoseReportRequest) SetResourceId(v string) *CreateDiagnoseReportRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateDiagnoseReportRequest) SetResourceType(v string) *CreateDiagnoseReportRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateDiagnoseReportRequest) SetStartTime(v string) *CreateDiagnoseReportRequest {
	s.StartTime = &v
	return s
}

func (s *CreateDiagnoseReportRequest) Validate() error {
	return dara.Validate(s)
}
