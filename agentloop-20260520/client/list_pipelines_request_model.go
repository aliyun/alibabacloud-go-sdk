// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPipelinesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPipelinesRequest
	GetNextToken() *string
	SetPipelineName(v string) *ListPipelinesRequest
	GetPipelineName() *string
	SetScheduleStatus(v string) *ListPipelinesRequest
	GetScheduleStatus() *string
	SetScheduleType(v string) *ListPipelinesRequest
	GetScheduleType() *string
	SetSinkName(v string) *ListPipelinesRequest
	GetSinkName() *string
}

type ListPipelinesRequest struct {
	// The maximum number of pipelines to return in the current request. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Pass the nextToken value from the previous response to retrieve the next page of results. Do not specify this parameter for the first request.
	//
	// example:
	//
	// MTIzNDU2Nzg5MA==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The name of the pipeline. Supports exact matching. If not specified, the results are not filtered by pipeline name.
	//
	// example:
	//
	// my-pipeline
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// The schedule status. Valid values: None (no schedule), Active (active), Paused (paused), and Terminated (terminated). If not specified, the results are not filtered by schedule status.
	//
	// example:
	//
	// Active
	ScheduleStatus *string `json:"scheduleStatus,omitempty" xml:"scheduleStatus,omitempty"`
	// The schedule type. Valid values: RunOnce (single execution), Scheduled (periodic scheduling), and Continuous (continuous execution driven by trace source signals). If not specified, the results are not filtered by schedule type.
	//
	// example:
	//
	// RunOnce
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
	// The name of the output dataset. Supports exact matching and applies to standard sinks, conditional branch sinks, and default sinks. If not specified, the results are not filtered by sink name.
	//
	// example:
	//
	// my-dataset
	SinkName *string `json:"sinkName,omitempty" xml:"sinkName,omitempty"`
}

func (s ListPipelinesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesRequest) GoString() string {
	return s.String()
}

func (s *ListPipelinesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPipelinesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPipelinesRequest) GetPipelineName() *string {
	return s.PipelineName
}

func (s *ListPipelinesRequest) GetScheduleStatus() *string {
	return s.ScheduleStatus
}

func (s *ListPipelinesRequest) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *ListPipelinesRequest) GetSinkName() *string {
	return s.SinkName
}

func (s *ListPipelinesRequest) SetMaxResults(v int32) *ListPipelinesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPipelinesRequest) SetNextToken(v string) *ListPipelinesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPipelinesRequest) SetPipelineName(v string) *ListPipelinesRequest {
	s.PipelineName = &v
	return s
}

func (s *ListPipelinesRequest) SetScheduleStatus(v string) *ListPipelinesRequest {
	s.ScheduleStatus = &v
	return s
}

func (s *ListPipelinesRequest) SetScheduleType(v string) *ListPipelinesRequest {
	s.ScheduleType = &v
	return s
}

func (s *ListPipelinesRequest) SetSinkName(v string) *ListPipelinesRequest {
	s.SinkName = &v
	return s
}

func (s *ListPipelinesRequest) Validate() error {
	return dara.Validate(s)
}
