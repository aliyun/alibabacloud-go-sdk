// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineRunsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListPipelineRunsRequest
	GetEndTime() *int64
	SetMaxResults(v int32) *ListPipelineRunsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPipelineRunsRequest
	GetNextToken() *string
	SetStartTime(v int64) *ListPipelineRunsRequest
	GetStartTime() *int64
	SetStatus(v string) *ListPipelineRunsRequest
	GetStatus() *string
	SetTriggerType(v string) *ListPipelineRunsRequest
	GetTriggerType() *string
}

type ListPipelineRunsRequest struct {
	// example:
	//
	// 1735660800
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The maximum number of entries to return. Default value: 50. Maximum value: 200.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Set this parameter to the nextToken value returned in the previous response to retrieve the next page. Do not specify this parameter for the first request.
	//
	// example:
	//
	// MTIzNDU2Nzg5MA==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1735574400
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// Filters by run status. Valid values: Pending, Running, Succeeded, Failed, and Cancelled. If this parameter is not specified, no filtering is applied.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Filters by trigger type. Valid values: Manual, Scheduled, and RunOnce. If this parameter is not specified, no filtering is applied.
	//
	// example:
	//
	// Scheduled
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s ListPipelineRunsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRunsRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListPipelineRunsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPipelineRunsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPipelineRunsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListPipelineRunsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListPipelineRunsRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListPipelineRunsRequest) SetEndTime(v int64) *ListPipelineRunsRequest {
	s.EndTime = &v
	return s
}

func (s *ListPipelineRunsRequest) SetMaxResults(v int32) *ListPipelineRunsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPipelineRunsRequest) SetNextToken(v string) *ListPipelineRunsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPipelineRunsRequest) SetStartTime(v int64) *ListPipelineRunsRequest {
	s.StartTime = &v
	return s
}

func (s *ListPipelineRunsRequest) SetStatus(v string) *ListPipelineRunsRequest {
	s.Status = &v
	return s
}

func (s *ListPipelineRunsRequest) SetTriggerType(v string) *ListPipelineRunsRequest {
	s.TriggerType = &v
	return s
}

func (s *ListPipelineRunsRequest) Validate() error {
	return dara.Validate(s)
}
