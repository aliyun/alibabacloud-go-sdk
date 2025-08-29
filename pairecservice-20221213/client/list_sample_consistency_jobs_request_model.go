// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSampleConsistencyJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListSampleConsistencyJobsRequest
	GetInstanceId() *string
	SetOrder(v string) *ListSampleConsistencyJobsRequest
	GetOrder() *string
	SetPageNumber(v int64) *ListSampleConsistencyJobsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListSampleConsistencyJobsRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListSampleConsistencyJobsRequest
	GetSortBy() *string
}

type ListSampleConsistencyJobsRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Order      *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListSampleConsistencyJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSampleConsistencyJobsRequest) GoString() string {
	return s.String()
}

func (s *ListSampleConsistencyJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSampleConsistencyJobsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListSampleConsistencyJobsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSampleConsistencyJobsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSampleConsistencyJobsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListSampleConsistencyJobsRequest) SetInstanceId(v string) *ListSampleConsistencyJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSampleConsistencyJobsRequest) SetOrder(v string) *ListSampleConsistencyJobsRequest {
	s.Order = &v
	return s
}

func (s *ListSampleConsistencyJobsRequest) SetPageNumber(v int64) *ListSampleConsistencyJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSampleConsistencyJobsRequest) SetPageSize(v int64) *ListSampleConsistencyJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSampleConsistencyJobsRequest) SetSortBy(v string) *ListSampleConsistencyJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListSampleConsistencyJobsRequest) Validate() error {
	return dara.Validate(s)
}
