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
	SetMaxResults(v int64) *ListPipelineRunsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListPipelineRunsRequest
	GetNextToken() *string
	SetStartTime(v int64) *ListPipelineRunsRequest
	GetStartTime() *int64
	SetStatus(v string) *ListPipelineRunsRequest
	GetStatus() *string
	SetTriggerMode(v int32) *ListPipelineRunsRequest
	GetTriggerMode() *int32
}

type ListPipelineRunsRequest struct {
	// example:
	//
	// 1586863220000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// aaaaaa
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1586863220000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1
	TriggerMode *int32 `json:"triggerMode,omitempty" xml:"triggerMode,omitempty"`
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

func (s *ListPipelineRunsRequest) GetMaxResults() *int64 {
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

func (s *ListPipelineRunsRequest) GetTriggerMode() *int32 {
	return s.TriggerMode
}

func (s *ListPipelineRunsRequest) SetEndTime(v int64) *ListPipelineRunsRequest {
	s.EndTime = &v
	return s
}

func (s *ListPipelineRunsRequest) SetMaxResults(v int64) *ListPipelineRunsRequest {
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

func (s *ListPipelineRunsRequest) SetTriggerMode(v int32) *ListPipelineRunsRequest {
	s.TriggerMode = &v
	return s
}

func (s *ListPipelineRunsRequest) Validate() error {
	return dara.Validate(s)
}
