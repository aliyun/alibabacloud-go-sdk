// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListGroupsRequest
	GetFilter() *string
	SetOrder(v string) *ListGroupsRequest
	GetOrder() *string
	SetPageNumber(v string) *ListGroupsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListGroupsRequest
	GetPageSize() *string
	SetSort(v string) *ListGroupsRequest
	GetSort() *string
	SetTrafficMode(v string) *ListGroupsRequest
	GetTrafficMode() *string
	SetWorkspaceId(v string) *ListGroupsRequest
	GetWorkspaceId() *string
}

type ListGroupsRequest struct {
	// The name of the filter that is used to filter out unwanted service groups. Fuzzy match is supported.
	//
	// example:
	//
	// foo
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	Order  *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 20
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Sort        *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	TrafficMode *string `json:"TrafficMode,omitempty" xml:"TrafficMode,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListGroupsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListGroupsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListGroupsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListGroupsRequest) GetSort() *string {
	return s.Sort
}

func (s *ListGroupsRequest) GetTrafficMode() *string {
	return s.TrafficMode
}

func (s *ListGroupsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListGroupsRequest) SetFilter(v string) *ListGroupsRequest {
	s.Filter = &v
	return s
}

func (s *ListGroupsRequest) SetOrder(v string) *ListGroupsRequest {
	s.Order = &v
	return s
}

func (s *ListGroupsRequest) SetPageNumber(v string) *ListGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGroupsRequest) SetPageSize(v string) *ListGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupsRequest) SetSort(v string) *ListGroupsRequest {
	s.Sort = &v
	return s
}

func (s *ListGroupsRequest) SetTrafficMode(v string) *ListGroupsRequest {
	s.TrafficMode = &v
	return s
}

func (s *ListGroupsRequest) SetWorkspaceId(v string) *ListGroupsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListGroupsRequest) Validate() error {
	return dara.Validate(s)
}
