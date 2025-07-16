// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListDatasourcesRequest
	GetName() *string
	SetOrder(v string) *ListDatasourcesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListDatasourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDatasourcesRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListDatasourcesRequest
	GetSortBy() *string
	SetType(v string) *ListDatasourcesRequest
	GetType() *string
	SetWorkspaceId(v string) *ListDatasourcesRequest
	GetWorkspaceId() *string
}

type ListDatasourcesRequest struct {
	// example:
	//
	// datasource1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtModifiedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// MaxCompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDatasourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasourcesRequest) GoString() string {
	return s.String()
}

func (s *ListDatasourcesRequest) GetName() *string {
	return s.Name
}

func (s *ListDatasourcesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDatasourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatasourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasourcesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListDatasourcesRequest) GetType() *string {
	return s.Type
}

func (s *ListDatasourcesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDatasourcesRequest) SetName(v string) *ListDatasourcesRequest {
	s.Name = &v
	return s
}

func (s *ListDatasourcesRequest) SetOrder(v string) *ListDatasourcesRequest {
	s.Order = &v
	return s
}

func (s *ListDatasourcesRequest) SetPageNumber(v int32) *ListDatasourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatasourcesRequest) SetPageSize(v int32) *ListDatasourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatasourcesRequest) SetSortBy(v string) *ListDatasourcesRequest {
	s.SortBy = &v
	return s
}

func (s *ListDatasourcesRequest) SetType(v string) *ListDatasourcesRequest {
	s.Type = &v
	return s
}

func (s *ListDatasourcesRequest) SetWorkspaceId(v string) *ListDatasourcesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDatasourcesRequest) Validate() error {
	return dara.Validate(s)
}
