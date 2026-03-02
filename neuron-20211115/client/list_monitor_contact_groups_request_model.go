// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMonitorContactGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ListMonitorContactGroupsRequest
	GetEnterpriseId() *int64
	SetName(v string) *ListMonitorContactGroupsRequest
	GetName() *string
	SetOrderBy(v string) *ListMonitorContactGroupsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListMonitorContactGroupsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListMonitorContactGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMonitorContactGroupsRequest
	GetPageSize() *int32
}

type ListMonitorContactGroupsRequest struct {
	EnterpriseId   *int64  `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListMonitorContactGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMonitorContactGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListMonitorContactGroupsRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ListMonitorContactGroupsRequest) GetName() *string {
	return s.Name
}

func (s *ListMonitorContactGroupsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMonitorContactGroupsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListMonitorContactGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMonitorContactGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMonitorContactGroupsRequest) SetEnterpriseId(v int64) *ListMonitorContactGroupsRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ListMonitorContactGroupsRequest) SetName(v string) *ListMonitorContactGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListMonitorContactGroupsRequest) SetOrderBy(v string) *ListMonitorContactGroupsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMonitorContactGroupsRequest) SetOrderDirection(v string) *ListMonitorContactGroupsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListMonitorContactGroupsRequest) SetPageNumber(v int32) *ListMonitorContactGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMonitorContactGroupsRequest) SetPageSize(v int32) *ListMonitorContactGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMonitorContactGroupsRequest) Validate() error {
	return dara.Validate(s)
}
