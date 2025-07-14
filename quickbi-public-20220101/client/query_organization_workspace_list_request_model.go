// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrganizationWorkspaceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *QueryOrganizationWorkspaceListRequest
	GetKeyword() *string
	SetPageNum(v int32) *QueryOrganizationWorkspaceListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryOrganizationWorkspaceListRequest
	GetPageSize() *int32
	SetUserId(v string) *QueryOrganizationWorkspaceListRequest
	GetUserId() *string
}

type QueryOrganizationWorkspaceListRequest struct {
	// Keyword for the workspace name.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Current page number of the workspace list:
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
	// User ID in Quick BI.
	//
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryOrganizationWorkspaceListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryOrganizationWorkspaceListRequest) GoString() string {
	return s.String()
}

func (s *QueryOrganizationWorkspaceListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *QueryOrganizationWorkspaceListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryOrganizationWorkspaceListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryOrganizationWorkspaceListRequest) GetUserId() *string {
	return s.UserId
}

func (s *QueryOrganizationWorkspaceListRequest) SetKeyword(v string) *QueryOrganizationWorkspaceListRequest {
	s.Keyword = &v
	return s
}

func (s *QueryOrganizationWorkspaceListRequest) SetPageNum(v int32) *QueryOrganizationWorkspaceListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryOrganizationWorkspaceListRequest) SetPageSize(v int32) *QueryOrganizationWorkspaceListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryOrganizationWorkspaceListRequest) SetUserId(v string) *QueryOrganizationWorkspaceListRequest {
	s.UserId = &v
	return s
}

func (s *QueryOrganizationWorkspaceListRequest) Validate() error {
	return dara.Validate(s)
}
