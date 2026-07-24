// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListComputeJobsRequest
	GetCurrentPage() *int64
	SetInstanceId(v string) *ListComputeJobsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListComputeJobsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListComputeJobsRequest
	GetNextToken() *string
	SetPageSize(v int64) *ListComputeJobsRequest
	GetPageSize() *int64
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
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Search        *string `json:"Search,omitempty" xml:"Search,omitempty"`
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
	SortField     *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
}

func (s ListComputeJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComputeJobsRequest) GoString() string {
	return s.String()
}

func (s *ListComputeJobsRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
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

func (s *ListComputeJobsRequest) GetPageSize() *int64 {
	return s.PageSize
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

func (s *ListComputeJobsRequest) SetCurrentPage(v int64) *ListComputeJobsRequest {
	s.CurrentPage = &v
	return s
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

func (s *ListComputeJobsRequest) SetPageSize(v int64) *ListComputeJobsRequest {
	s.PageSize = &v
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
