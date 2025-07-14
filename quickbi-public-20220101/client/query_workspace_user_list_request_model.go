// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorkspaceUserListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *QueryWorkspaceUserListRequest
	GetKeyword() *string
	SetPageNum(v int32) *QueryWorkspaceUserListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryWorkspaceUserListRequest
	GetPageSize() *int32
	SetWorkspaceId(v string) *QueryWorkspaceUserListRequest
	GetWorkspaceId() *string
}

type QueryWorkspaceUserListRequest struct {
	// Keyword for the username of the workspace member.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Current page number of the workspace member list:
	//
	// - Starting value: 1
	//
	// - Default value: 1
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Number of rows per page in a paginated query:
	//
	// - Default value: 10
	//
	// - Maximum value: 1000
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryWorkspaceUserListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkspaceUserListRequest) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceUserListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *QueryWorkspaceUserListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryWorkspaceUserListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryWorkspaceUserListRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryWorkspaceUserListRequest) SetKeyword(v string) *QueryWorkspaceUserListRequest {
	s.Keyword = &v
	return s
}

func (s *QueryWorkspaceUserListRequest) SetPageNum(v int32) *QueryWorkspaceUserListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryWorkspaceUserListRequest) SetPageSize(v int32) *QueryWorkspaceUserListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryWorkspaceUserListRequest) SetWorkspaceId(v string) *QueryWorkspaceUserListRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryWorkspaceUserListRequest) Validate() error {
	return dara.Validate(s)
}
