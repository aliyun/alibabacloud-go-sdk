// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v *ListJobsRequestFilter) *ListJobsRequest
	GetFilter() *ListJobsRequestFilter
	SetPageNumber(v int32) *ListJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobsRequest
	GetPageSize() *int32
	SetSortBy(v *ListJobsRequestSortBy) *ListJobsRequest
	GetSortBy() *ListJobsRequestSortBy
}

type ListJobsRequest struct {
	Filter *ListJobsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy   *ListJobsRequestSortBy `json:"SortBy,omitempty" xml:"SortBy,omitempty" type:"Struct"`
}

func (s ListJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) GetFilter() *ListJobsRequestFilter {
	return s.Filter
}

func (s *ListJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsRequest) GetSortBy() *ListJobsRequestSortBy {
	return s.SortBy
}

func (s *ListJobsRequest) SetFilter(v *ListJobsRequestFilter) *ListJobsRequest {
	s.Filter = v
	return s
}

func (s *ListJobsRequest) SetPageNumber(v int32) *ListJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsRequest) SetSortBy(v *ListJobsRequestSortBy) *ListJobsRequest {
	s.SortBy = v
	return s
}

func (s *ListJobsRequest) Validate() error {
	return dara.Validate(s)
}

type ListJobsRequestFilter struct {
	// example:
	//
	// job-xxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// testJob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1703819914
	TimeCreatedAfter *int32 `json:"TimeCreatedAfter,omitempty" xml:"TimeCreatedAfter,omitempty"`
	// example:
	//
	// 1703820113
	TimeCreatedBefore *int32 `json:"TimeCreatedBefore,omitempty" xml:"TimeCreatedBefore,omitempty"`
}

func (s ListJobsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListJobsRequestFilter) GetJobId() *string {
	return s.JobId
}

func (s *ListJobsRequestFilter) GetJobName() *string {
	return s.JobName
}

func (s *ListJobsRequestFilter) GetStatus() *string {
	return s.Status
}

func (s *ListJobsRequestFilter) GetTimeCreatedAfter() *int32 {
	return s.TimeCreatedAfter
}

func (s *ListJobsRequestFilter) GetTimeCreatedBefore() *int32 {
	return s.TimeCreatedBefore
}

func (s *ListJobsRequestFilter) SetJobId(v string) *ListJobsRequestFilter {
	s.JobId = &v
	return s
}

func (s *ListJobsRequestFilter) SetJobName(v string) *ListJobsRequestFilter {
	s.JobName = &v
	return s
}

func (s *ListJobsRequestFilter) SetStatus(v string) *ListJobsRequestFilter {
	s.Status = &v
	return s
}

func (s *ListJobsRequestFilter) SetTimeCreatedAfter(v int32) *ListJobsRequestFilter {
	s.TimeCreatedAfter = &v
	return s
}

func (s *ListJobsRequestFilter) SetTimeCreatedBefore(v int32) *ListJobsRequestFilter {
	s.TimeCreatedBefore = &v
	return s
}

func (s *ListJobsRequestFilter) Validate() error {
	return dara.Validate(s)
}

type ListJobsRequestSortBy struct {
	// example:
	//
	// time_start
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s ListJobsRequestSortBy) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequestSortBy) GoString() string {
	return s.String()
}

func (s *ListJobsRequestSortBy) GetLabel() *string {
	return s.Label
}

func (s *ListJobsRequestSortBy) GetOrder() *string {
	return s.Order
}

func (s *ListJobsRequestSortBy) SetLabel(v string) *ListJobsRequestSortBy {
	s.Label = &v
	return s
}

func (s *ListJobsRequestSortBy) SetOrder(v string) *ListJobsRequestSortBy {
	s.Order = &v
	return s
}

func (s *ListJobsRequestSortBy) Validate() error {
	return dara.Validate(s)
}
