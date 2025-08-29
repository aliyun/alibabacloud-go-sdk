// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureConsistencyCheckJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListFeatureConsistencyCheckJobsRequest
	GetInstanceId() *string
	SetOrder(v string) *ListFeatureConsistencyCheckJobsRequest
	GetOrder() *string
	SetPageNumber(v string) *ListFeatureConsistencyCheckJobsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListFeatureConsistencyCheckJobsRequest
	GetPageSize() *string
	SetSortBy(v string) *ListFeatureConsistencyCheckJobsRequest
	GetSortBy() *string
	SetStatus(v string) *ListFeatureConsistencyCheckJobsRequest
	GetStatus() *string
}

type ListFeatureConsistencyCheckJobsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFeatureConsistencyCheckJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobsRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFeatureConsistencyCheckJobsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListFeatureConsistencyCheckJobsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListFeatureConsistencyCheckJobsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListFeatureConsistencyCheckJobsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListFeatureConsistencyCheckJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListFeatureConsistencyCheckJobsRequest) SetInstanceId(v string) *ListFeatureConsistencyCheckJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsRequest) SetOrder(v string) *ListFeatureConsistencyCheckJobsRequest {
	s.Order = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsRequest) SetPageNumber(v string) *ListFeatureConsistencyCheckJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsRequest) SetPageSize(v string) *ListFeatureConsistencyCheckJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsRequest) SetSortBy(v string) *ListFeatureConsistencyCheckJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsRequest) SetStatus(v string) *ListFeatureConsistencyCheckJobsRequest {
	s.Status = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsRequest) Validate() error {
	return dara.Validate(s)
}
