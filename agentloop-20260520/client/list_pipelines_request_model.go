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
}

type ListPipelinesRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// MTIzNDU2Nzg5MA==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// my-pipeline
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// example:
	//
	// Active
	ScheduleStatus *string `json:"scheduleStatus,omitempty" xml:"scheduleStatus,omitempty"`
	// example:
	//
	// RunOnce
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
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

func (s *ListPipelinesRequest) Validate() error {
	return dara.Validate(s)
}
