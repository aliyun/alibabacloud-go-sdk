// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobExecutorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutorType(v string) *ListJobExecutorsRequest
	GetExecutorType() *string
	SetMaxResults(v int32) *ListJobExecutorsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListJobExecutorsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListJobExecutorsRequest
	GetRegionId() *string
	SetStatus(v string) *ListJobExecutorsRequest
	GetStatus() *string
}

type ListJobExecutorsRequest struct {
	// example:
	//
	// driver
	ExecutorType *string `json:"executorType,omitempty" xml:"executorType,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 2
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// Dead
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListJobExecutorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobExecutorsRequest) GoString() string {
	return s.String()
}

func (s *ListJobExecutorsRequest) GetExecutorType() *string {
	return s.ExecutorType
}

func (s *ListJobExecutorsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListJobExecutorsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListJobExecutorsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListJobExecutorsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListJobExecutorsRequest) SetExecutorType(v string) *ListJobExecutorsRequest {
	s.ExecutorType = &v
	return s
}

func (s *ListJobExecutorsRequest) SetMaxResults(v int32) *ListJobExecutorsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListJobExecutorsRequest) SetNextToken(v string) *ListJobExecutorsRequest {
	s.NextToken = &v
	return s
}

func (s *ListJobExecutorsRequest) SetRegionId(v string) *ListJobExecutorsRequest {
	s.RegionId = &v
	return s
}

func (s *ListJobExecutorsRequest) SetStatus(v string) *ListJobExecutorsRequest {
	s.Status = &v
	return s
}

func (s *ListJobExecutorsRequest) Validate() error {
	return dara.Validate(s)
}
