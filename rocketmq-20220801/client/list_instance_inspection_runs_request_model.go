// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceInspectionRunsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListInstanceInspectionRunsRequest
	GetEndTime() *string
	SetMaxResults(v int32) *ListInstanceInspectionRunsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstanceInspectionRunsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListInstanceInspectionRunsRequest
	GetRegionId() *string
	SetReportStatus(v string) *ListInstanceInspectionRunsRequest
	GetReportStatus() *string
	SetResult(v string) *ListInstanceInspectionRunsRequest
	GetResult() *string
	SetSortBy(v string) *ListInstanceInspectionRunsRequest
	GetSortBy() *string
	SetStartTime(v string) *ListInstanceInspectionRunsRequest
	GetStartTime() *string
	SetStatus(v string) *ListInstanceInspectionRunsRequest
	GetStatus() *string
	SetTemplateCode(v string) *ListInstanceInspectionRunsRequest
	GetTemplateCode() *string
}

type ListInstanceInspectionRunsRequest struct {
	// example:
	//
	// 2026-09-08T00:00:00Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// Succeeded
	ReportStatus *string `json:"reportStatus,omitempty" xml:"reportStatus,omitempty"`
	// example:
	//
	// risk
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// reportFinishedAt
	SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	// example:
	//
	// 2026-09-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// rmq.consumer-health
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
}

func (s ListInstanceInspectionRunsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionRunsRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionRunsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListInstanceInspectionRunsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstanceInspectionRunsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstanceInspectionRunsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceInspectionRunsRequest) GetReportStatus() *string {
	return s.ReportStatus
}

func (s *ListInstanceInspectionRunsRequest) GetResult() *string {
	return s.Result
}

func (s *ListInstanceInspectionRunsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListInstanceInspectionRunsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListInstanceInspectionRunsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceInspectionRunsRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ListInstanceInspectionRunsRequest) SetEndTime(v string) *ListInstanceInspectionRunsRequest {
	s.EndTime = &v
	return s
}

func (s *ListInstanceInspectionRunsRequest) SetMaxResults(v int32) *ListInstanceInspectionRunsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstanceInspectionRunsRequest) SetNextToken(v string) *ListInstanceInspectionRunsRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstanceInspectionRunsRequest) SetRegionId(v string) *ListInstanceInspectionRunsRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstanceInspectionRunsRequest) SetReportStatus(v string) *ListInstanceInspectionRunsRequest {
	s.ReportStatus = &v
	return s
}

func (s *ListInstanceInspectionRunsRequest) SetResult(v string) *ListInstanceInspectionRunsRequest {
	s.Result = &v
	return s
}

func (s *ListInstanceInspectionRunsRequest) SetSortBy(v string) *ListInstanceInspectionRunsRequest {
	s.SortBy = &v
	return s
}

func (s *ListInstanceInspectionRunsRequest) SetStartTime(v string) *ListInstanceInspectionRunsRequest {
	s.StartTime = &v
	return s
}

func (s *ListInstanceInspectionRunsRequest) SetStatus(v string) *ListInstanceInspectionRunsRequest {
	s.Status = &v
	return s
}

func (s *ListInstanceInspectionRunsRequest) SetTemplateCode(v string) *ListInstanceInspectionRunsRequest {
	s.TemplateCode = &v
	return s
}

func (s *ListInstanceInspectionRunsRequest) Validate() error {
	return dara.Validate(s)
}
