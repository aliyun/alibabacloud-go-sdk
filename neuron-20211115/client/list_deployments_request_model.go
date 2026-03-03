// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExcludeStatus(v []*string) *ListDeploymentsRequest
	GetExcludeStatus() []*string
	SetOrderBy(v string) *ListDeploymentsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListDeploymentsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListDeploymentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDeploymentsRequest
	GetPageSize() *int32
	SetServiceGroupId(v int64) *ListDeploymentsRequest
	GetServiceGroupId() *int64
	SetStatus(v []*string) *ListDeploymentsRequest
	GetStatus() []*string
}

type ListDeploymentsRequest struct {
	ExcludeStatus []*string `json:"excludeStatus,omitempty" xml:"excludeStatus,omitempty" type:"Repeated"`
	// example:
	//
	// gmtCreated
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// desc
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceGroupId *int64    `json:"serviceGroupId,omitempty" xml:"serviceGroupId,omitempty"`
	Status         []*string `json:"status,omitempty" xml:"status,omitempty" type:"Repeated"`
}

func (s ListDeploymentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentsRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentsRequest) GetExcludeStatus() []*string {
	return s.ExcludeStatus
}

func (s *ListDeploymentsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDeploymentsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListDeploymentsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDeploymentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeploymentsRequest) GetServiceGroupId() *int64 {
	return s.ServiceGroupId
}

func (s *ListDeploymentsRequest) GetStatus() []*string {
	return s.Status
}

func (s *ListDeploymentsRequest) SetExcludeStatus(v []*string) *ListDeploymentsRequest {
	s.ExcludeStatus = v
	return s
}

func (s *ListDeploymentsRequest) SetOrderBy(v string) *ListDeploymentsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDeploymentsRequest) SetOrderDirection(v string) *ListDeploymentsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListDeploymentsRequest) SetPageNumber(v int32) *ListDeploymentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDeploymentsRequest) SetPageSize(v int32) *ListDeploymentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentsRequest) SetServiceGroupId(v int64) *ListDeploymentsRequest {
	s.ServiceGroupId = &v
	return s
}

func (s *ListDeploymentsRequest) SetStatus(v []*string) *ListDeploymentsRequest {
	s.Status = v
	return s
}

func (s *ListDeploymentsRequest) Validate() error {
	return dara.Validate(s)
}
