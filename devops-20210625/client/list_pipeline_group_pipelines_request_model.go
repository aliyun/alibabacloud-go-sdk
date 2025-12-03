// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineGroupPipelinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateEndTime(v int64) *ListPipelineGroupPipelinesRequest
	GetCreateEndTime() *int64
	SetCreateStartTime(v int64) *ListPipelineGroupPipelinesRequest
	GetCreateStartTime() *int64
	SetExecuteEndTime(v int64) *ListPipelineGroupPipelinesRequest
	GetExecuteEndTime() *int64
	SetExecuteStartTime(v int64) *ListPipelineGroupPipelinesRequest
	GetExecuteStartTime() *int64
	SetMaxResults(v int64) *ListPipelineGroupPipelinesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListPipelineGroupPipelinesRequest
	GetNextToken() *string
	SetPipelineName(v string) *ListPipelineGroupPipelinesRequest
	GetPipelineName() *string
	SetResultStatusList(v string) *ListPipelineGroupPipelinesRequest
	GetResultStatusList() *string
}

type ListPipelineGroupPipelinesRequest struct {
	// example:
	//
	// 1586863220000
	CreateEndTime *int64 `json:"createEndTime,omitempty" xml:"createEndTime,omitempty"`
	// example:
	//
	// 1586863220000
	CreateStartTime *int64 `json:"createStartTime,omitempty" xml:"createStartTime,omitempty"`
	// example:
	//
	// 1586863220000
	ExecuteEndTime *int64 `json:"executeEndTime,omitempty" xml:"executeEndTime,omitempty"`
	// example:
	//
	// 1586863220000
	ExecuteStartTime *int64 `json:"executeStartTime,omitempty" xml:"executeStartTime,omitempty"`
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// aaaa
	NextToken    *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// example:
	//
	// RUNNING,SUCCESS
	ResultStatusList *string `json:"resultStatusList,omitempty" xml:"resultStatusList,omitempty"`
}

func (s ListPipelineGroupPipelinesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineGroupPipelinesRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineGroupPipelinesRequest) GetCreateEndTime() *int64 {
	return s.CreateEndTime
}

func (s *ListPipelineGroupPipelinesRequest) GetCreateStartTime() *int64 {
	return s.CreateStartTime
}

func (s *ListPipelineGroupPipelinesRequest) GetExecuteEndTime() *int64 {
	return s.ExecuteEndTime
}

func (s *ListPipelineGroupPipelinesRequest) GetExecuteStartTime() *int64 {
	return s.ExecuteStartTime
}

func (s *ListPipelineGroupPipelinesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListPipelineGroupPipelinesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPipelineGroupPipelinesRequest) GetPipelineName() *string {
	return s.PipelineName
}

func (s *ListPipelineGroupPipelinesRequest) GetResultStatusList() *string {
	return s.ResultStatusList
}

func (s *ListPipelineGroupPipelinesRequest) SetCreateEndTime(v int64) *ListPipelineGroupPipelinesRequest {
	s.CreateEndTime = &v
	return s
}

func (s *ListPipelineGroupPipelinesRequest) SetCreateStartTime(v int64) *ListPipelineGroupPipelinesRequest {
	s.CreateStartTime = &v
	return s
}

func (s *ListPipelineGroupPipelinesRequest) SetExecuteEndTime(v int64) *ListPipelineGroupPipelinesRequest {
	s.ExecuteEndTime = &v
	return s
}

func (s *ListPipelineGroupPipelinesRequest) SetExecuteStartTime(v int64) *ListPipelineGroupPipelinesRequest {
	s.ExecuteStartTime = &v
	return s
}

func (s *ListPipelineGroupPipelinesRequest) SetMaxResults(v int64) *ListPipelineGroupPipelinesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPipelineGroupPipelinesRequest) SetNextToken(v string) *ListPipelineGroupPipelinesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPipelineGroupPipelinesRequest) SetPipelineName(v string) *ListPipelineGroupPipelinesRequest {
	s.PipelineName = &v
	return s
}

func (s *ListPipelineGroupPipelinesRequest) SetResultStatusList(v string) *ListPipelineGroupPipelinesRequest {
	s.ResultStatusList = &v
	return s
}

func (s *ListPipelineGroupPipelinesRequest) Validate() error {
	return dara.Validate(s)
}
