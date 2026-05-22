// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDDoSInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListDDoSInstancesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListDDoSInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDDoSInstancesRequest
	GetPageSize() *int32
	SetSiteInstanceId(v string) *ListDDoSInstancesRequest
	GetSiteInstanceId() *string
	SetSortBy(v string) *ListDDoSInstancesRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListDDoSInstancesRequest
	GetSortOrder() *string
	SetStatus(v string) *ListDDoSInstancesRequest
	GetStatus() *string
}

type ListDDoSInstancesRequest struct {
	// example:
	//
	// esa-ddos-b1e0l80ugfeo
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// esa-site-a71k7bw1adf
	SiteInstanceId *string `json:"SiteInstanceId,omitempty" xml:"SiteInstanceId,omitempty"`
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// asc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDDoSInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDDoSInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListDDoSInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDDoSInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDDoSInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDDoSInstancesRequest) GetSiteInstanceId() *string {
	return s.SiteInstanceId
}

func (s *ListDDoSInstancesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListDDoSInstancesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListDDoSInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDDoSInstancesRequest) SetInstanceId(v string) *ListDDoSInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDDoSInstancesRequest) SetPageNumber(v int32) *ListDDoSInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDDoSInstancesRequest) SetPageSize(v int32) *ListDDoSInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDDoSInstancesRequest) SetSiteInstanceId(v string) *ListDDoSInstancesRequest {
	s.SiteInstanceId = &v
	return s
}

func (s *ListDDoSInstancesRequest) SetSortBy(v string) *ListDDoSInstancesRequest {
	s.SortBy = &v
	return s
}

func (s *ListDDoSInstancesRequest) SetSortOrder(v string) *ListDDoSInstancesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListDDoSInstancesRequest) SetStatus(v string) *ListDDoSInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListDDoSInstancesRequest) Validate() error {
	return dara.Validate(s)
}
