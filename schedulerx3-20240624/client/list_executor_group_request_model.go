// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutorGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListExecutorGroupRequest
	GetClusterId() *string
	SetMaxResults(v int32) *ListExecutorGroupRequest
	GetMaxResults() *int32
	SetName(v string) *ListExecutorGroupRequest
	GetName() *string
	SetNextToken(v string) *ListExecutorGroupRequest
	GetNextToken() *string
	SetPageNum(v int32) *ListExecutorGroupRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListExecutorGroupRequest
	GetPageSize() *int32
	SetWorkerType(v string) *ListExecutorGroupRequest
	GetWorkerType() *string
}

type ListExecutorGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// eCKqVlS5FKF5EWGGOo8EgQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// openclaw
	WorkerType *string `json:"WorkerType,omitempty" xml:"WorkerType,omitempty"`
}

func (s ListExecutorGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorGroupRequest) GoString() string {
	return s.String()
}

func (s *ListExecutorGroupRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListExecutorGroupRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExecutorGroupRequest) GetName() *string {
	return s.Name
}

func (s *ListExecutorGroupRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExecutorGroupRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListExecutorGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExecutorGroupRequest) GetWorkerType() *string {
	return s.WorkerType
}

func (s *ListExecutorGroupRequest) SetClusterId(v string) *ListExecutorGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *ListExecutorGroupRequest) SetMaxResults(v int32) *ListExecutorGroupRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExecutorGroupRequest) SetName(v string) *ListExecutorGroupRequest {
	s.Name = &v
	return s
}

func (s *ListExecutorGroupRequest) SetNextToken(v string) *ListExecutorGroupRequest {
	s.NextToken = &v
	return s
}

func (s *ListExecutorGroupRequest) SetPageNum(v int32) *ListExecutorGroupRequest {
	s.PageNum = &v
	return s
}

func (s *ListExecutorGroupRequest) SetPageSize(v int32) *ListExecutorGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListExecutorGroupRequest) SetWorkerType(v string) *ListExecutorGroupRequest {
	s.WorkerType = &v
	return s
}

func (s *ListExecutorGroupRequest) Validate() error {
	return dara.Validate(s)
}
