// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollaborationGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListCollaborationGroupsRequest
	GetKeyword() *string
	SetPage(v int64) *ListCollaborationGroupsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListCollaborationGroupsRequest
	GetPageSize() *int64
	SetTenantId(v string) *ListCollaborationGroupsRequest
	GetTenantId() *string
}

type ListCollaborationGroupsRequest struct {
	// The search keyword. This parameter is optional and can contain up to 256 characters.
	//
	// example:
	//
	// Workflow
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The page number. Pages start from 1.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The tenant ID. If you do not specify this parameter, the default tenant of the caller is used.
	//
	// example:
	//
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListCollaborationGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCollaborationGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListCollaborationGroupsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListCollaborationGroupsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListCollaborationGroupsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCollaborationGroupsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListCollaborationGroupsRequest) SetKeyword(v string) *ListCollaborationGroupsRequest {
	s.Keyword = &v
	return s
}

func (s *ListCollaborationGroupsRequest) SetPage(v int64) *ListCollaborationGroupsRequest {
	s.Page = &v
	return s
}

func (s *ListCollaborationGroupsRequest) SetPageSize(v int64) *ListCollaborationGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCollaborationGroupsRequest) SetTenantId(v string) *ListCollaborationGroupsRequest {
	s.TenantId = &v
	return s
}

func (s *ListCollaborationGroupsRequest) Validate() error {
	return dara.Validate(s)
}
