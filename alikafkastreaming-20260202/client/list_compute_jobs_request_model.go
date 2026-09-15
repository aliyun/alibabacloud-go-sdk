// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListComputeJobsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListComputeJobsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListComputeJobsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListComputeJobsRequest
	GetRegionId() *string
	SetSearch(v string) *ListComputeJobsRequest
	GetSearch() *string
	SetSortDirection(v string) *ListComputeJobsRequest
	GetSortDirection() *string
	SetSortField(v string) *ListComputeJobsRequest
	GetSortField() *string
}

type ListComputeJobsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// alikafka_streaming-cn-a1b2c3d4
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// order
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// example:
	//
	// DESC
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
	// example:
	//
	// createTime
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
}

func (s ListComputeJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComputeJobsRequest) GoString() string {
	return s.String()
}

func (s *ListComputeJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListComputeJobsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListComputeJobsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListComputeJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListComputeJobsRequest) GetSearch() *string {
	return s.Search
}

func (s *ListComputeJobsRequest) GetSortDirection() *string {
	return s.SortDirection
}

func (s *ListComputeJobsRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListComputeJobsRequest) SetInstanceId(v string) *ListComputeJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListComputeJobsRequest) SetMaxResults(v int32) *ListComputeJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListComputeJobsRequest) SetNextToken(v string) *ListComputeJobsRequest {
	s.NextToken = &v
	return s
}

func (s *ListComputeJobsRequest) SetRegionId(v string) *ListComputeJobsRequest {
	s.RegionId = &v
	return s
}

func (s *ListComputeJobsRequest) SetSearch(v string) *ListComputeJobsRequest {
	s.Search = &v
	return s
}

func (s *ListComputeJobsRequest) SetSortDirection(v string) *ListComputeJobsRequest {
	s.SortDirection = &v
	return s
}

func (s *ListComputeJobsRequest) SetSortField(v string) *ListComputeJobsRequest {
	s.SortField = &v
	return s
}

func (s *ListComputeJobsRequest) Validate() error {
	return dara.Validate(s)
}
