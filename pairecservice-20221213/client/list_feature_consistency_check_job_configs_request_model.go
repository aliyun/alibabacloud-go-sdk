// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureConsistencyCheckJobConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListFeatureConsistencyCheckJobConfigsRequest
	GetInstanceId() *string
	SetOrder(v string) *ListFeatureConsistencyCheckJobConfigsRequest
	GetOrder() *string
	SetPageNumber(v string) *ListFeatureConsistencyCheckJobConfigsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListFeatureConsistencyCheckJobConfigsRequest
	GetPageSize() *string
	SetSortBy(v string) *ListFeatureConsistencyCheckJobConfigsRequest
	GetSortBy() *string
}

type ListFeatureConsistencyCheckJobConfigsRequest struct {
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
}

func (s ListFeatureConsistencyCheckJobConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobConfigsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFeatureConsistencyCheckJobConfigsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListFeatureConsistencyCheckJobConfigsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListFeatureConsistencyCheckJobConfigsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListFeatureConsistencyCheckJobConfigsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListFeatureConsistencyCheckJobConfigsRequest) SetInstanceId(v string) *ListFeatureConsistencyCheckJobConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsRequest) SetOrder(v string) *ListFeatureConsistencyCheckJobConfigsRequest {
	s.Order = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsRequest) SetPageNumber(v string) *ListFeatureConsistencyCheckJobConfigsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsRequest) SetPageSize(v string) *ListFeatureConsistencyCheckJobConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsRequest) SetSortBy(v string) *ListFeatureConsistencyCheckJobConfigsRequest {
	s.SortBy = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsRequest) Validate() error {
	return dara.Validate(s)
}
