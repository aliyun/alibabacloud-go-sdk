// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateEndTime(v int64) *ListPipelinesRequest
	GetCreateEndTime() *int64
	SetCreateStartTime(v int64) *ListPipelinesRequest
	GetCreateStartTime() *int64
	SetCreatorAccountIds(v string) *ListPipelinesRequest
	GetCreatorAccountIds() *string
	SetExecuteAccountIds(v string) *ListPipelinesRequest
	GetExecuteAccountIds() *string
	SetExecuteEndTime(v int64) *ListPipelinesRequest
	GetExecuteEndTime() *int64
	SetExecuteStartTime(v int64) *ListPipelinesRequest
	GetExecuteStartTime() *int64
	SetMaxResults(v int64) *ListPipelinesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListPipelinesRequest
	GetNextToken() *string
	SetPipelineName(v string) *ListPipelinesRequest
	GetPipelineName() *string
	SetStatusList(v string) *ListPipelinesRequest
	GetStatusList() *string
}

type ListPipelinesRequest struct {
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
	// 112122121
	CreatorAccountIds *string `json:"creatorAccountIds,omitempty" xml:"creatorAccountIds,omitempty"`
	// example:
	//
	// 1111111
	ExecuteAccountIds *string `json:"executeAccountIds,omitempty" xml:"executeAccountIds,omitempty"`
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
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// aaaaaaaaaa
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 流水线
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// example:
	//
	// RUNNING,SUCCESS
	StatusList *string `json:"statusList,omitempty" xml:"statusList,omitempty"`
}

func (s ListPipelinesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesRequest) GoString() string {
	return s.String()
}

func (s *ListPipelinesRequest) GetCreateEndTime() *int64 {
	return s.CreateEndTime
}

func (s *ListPipelinesRequest) GetCreateStartTime() *int64 {
	return s.CreateStartTime
}

func (s *ListPipelinesRequest) GetCreatorAccountIds() *string {
	return s.CreatorAccountIds
}

func (s *ListPipelinesRequest) GetExecuteAccountIds() *string {
	return s.ExecuteAccountIds
}

func (s *ListPipelinesRequest) GetExecuteEndTime() *int64 {
	return s.ExecuteEndTime
}

func (s *ListPipelinesRequest) GetExecuteStartTime() *int64 {
	return s.ExecuteStartTime
}

func (s *ListPipelinesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListPipelinesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPipelinesRequest) GetPipelineName() *string {
	return s.PipelineName
}

func (s *ListPipelinesRequest) GetStatusList() *string {
	return s.StatusList
}

func (s *ListPipelinesRequest) SetCreateEndTime(v int64) *ListPipelinesRequest {
	s.CreateEndTime = &v
	return s
}

func (s *ListPipelinesRequest) SetCreateStartTime(v int64) *ListPipelinesRequest {
	s.CreateStartTime = &v
	return s
}

func (s *ListPipelinesRequest) SetCreatorAccountIds(v string) *ListPipelinesRequest {
	s.CreatorAccountIds = &v
	return s
}

func (s *ListPipelinesRequest) SetExecuteAccountIds(v string) *ListPipelinesRequest {
	s.ExecuteAccountIds = &v
	return s
}

func (s *ListPipelinesRequest) SetExecuteEndTime(v int64) *ListPipelinesRequest {
	s.ExecuteEndTime = &v
	return s
}

func (s *ListPipelinesRequest) SetExecuteStartTime(v int64) *ListPipelinesRequest {
	s.ExecuteStartTime = &v
	return s
}

func (s *ListPipelinesRequest) SetMaxResults(v int64) *ListPipelinesRequest {
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

func (s *ListPipelinesRequest) SetStatusList(v string) *ListPipelinesRequest {
	s.StatusList = &v
	return s
}

func (s *ListPipelinesRequest) Validate() error {
	return dara.Validate(s)
}
