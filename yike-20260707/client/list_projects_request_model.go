// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListProjectsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListProjectsRequest
	GetPageSize() *int32
	SetWorkspaceId(v string) *ListProjectsRequest
	GetWorkspaceId() *string
}

type ListProjectsRequest struct {
	// The page number. Pages start from 1. Specify a valid positive integer for both this parameter and PageSize. Otherwise, the query defaults to page 1 with 10 entries per page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page. Specify a valid positive integer for both this parameter and PageNo. Otherwise, the query defaults to page 1 with 10 entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filters projects by workspace ID.
	//
	// example:
	//
	// ws_zna577pdximv***
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListProjectsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListProjectsRequest) SetPageNo(v int32) *ListProjectsRequest {
	s.PageNo = &v
	return s
}

func (s *ListProjectsRequest) SetPageSize(v int32) *ListProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectsRequest) SetWorkspaceId(v string) *ListProjectsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListProjectsRequest) Validate() error {
	return dara.Validate(s)
}
