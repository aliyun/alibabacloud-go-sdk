// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMonitorContactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ListMonitorContactsRequest
	GetEnterpriseId() *int64
	SetGroupId(v int64) *ListMonitorContactsRequest
	GetGroupId() *int64
	SetName(v string) *ListMonitorContactsRequest
	GetName() *string
	SetOrderBy(v string) *ListMonitorContactsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListMonitorContactsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListMonitorContactsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMonitorContactsRequest
	GetPageSize() *int32
}

type ListMonitorContactsRequest struct {
	// example:
	//
	// 4
	EnterpriseId *int64 `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	// example:
	//
	// 1445
	GroupId *int64  `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// id
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// DESC
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListMonitorContactsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMonitorContactsRequest) GoString() string {
	return s.String()
}

func (s *ListMonitorContactsRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ListMonitorContactsRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListMonitorContactsRequest) GetName() *string {
	return s.Name
}

func (s *ListMonitorContactsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMonitorContactsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListMonitorContactsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMonitorContactsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMonitorContactsRequest) SetEnterpriseId(v int64) *ListMonitorContactsRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ListMonitorContactsRequest) SetGroupId(v int64) *ListMonitorContactsRequest {
	s.GroupId = &v
	return s
}

func (s *ListMonitorContactsRequest) SetName(v string) *ListMonitorContactsRequest {
	s.Name = &v
	return s
}

func (s *ListMonitorContactsRequest) SetOrderBy(v string) *ListMonitorContactsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMonitorContactsRequest) SetOrderDirection(v string) *ListMonitorContactsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListMonitorContactsRequest) SetPageNumber(v int32) *ListMonitorContactsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMonitorContactsRequest) SetPageSize(v int32) *ListMonitorContactsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMonitorContactsRequest) Validate() error {
	return dara.Validate(s)
}
