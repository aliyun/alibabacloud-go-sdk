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
	SetPageNumber(v string) *ListGroupsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListGroupsRequest
	GetPageSize() *string
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
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *ListGroupsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListGroupsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListGroupsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListGroupsRequest) SetFilter(v string) *ListGroupsRequest {
	s.Filter = &v
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

func (s *ListGroupsRequest) SetWorkspaceId(v string) *ListGroupsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListGroupsRequest) Validate() error {
	return dara.Validate(s)
}
