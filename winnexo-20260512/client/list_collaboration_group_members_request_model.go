// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollaborationGroupMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *ListCollaborationGroupMembersRequest
	GetCollaborationGroupId() *string
	SetKeyword(v string) *ListCollaborationGroupMembersRequest
	GetKeyword() *string
	SetPage(v int64) *ListCollaborationGroupMembersRequest
	GetPage() *int64
	SetPageSize(v int64) *ListCollaborationGroupMembersRequest
	GetPageSize() *int64
	SetTenantId(v string) *ListCollaborationGroupMembersRequest
	GetTenantId() *string
}

type ListCollaborationGroupMembersRequest struct {
	// The collaboration workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The search keyword. This parameter is optional and can be up to 256 characters in length.
	//
	// example:
	//
	// Sales
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
	// The tenant ID. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListCollaborationGroupMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCollaborationGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *ListCollaborationGroupMembersRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListCollaborationGroupMembersRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListCollaborationGroupMembersRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListCollaborationGroupMembersRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCollaborationGroupMembersRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListCollaborationGroupMembersRequest) SetCollaborationGroupId(v string) *ListCollaborationGroupMembersRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListCollaborationGroupMembersRequest) SetKeyword(v string) *ListCollaborationGroupMembersRequest {
	s.Keyword = &v
	return s
}

func (s *ListCollaborationGroupMembersRequest) SetPage(v int64) *ListCollaborationGroupMembersRequest {
	s.Page = &v
	return s
}

func (s *ListCollaborationGroupMembersRequest) SetPageSize(v int64) *ListCollaborationGroupMembersRequest {
	s.PageSize = &v
	return s
}

func (s *ListCollaborationGroupMembersRequest) SetTenantId(v string) *ListCollaborationGroupMembersRequest {
	s.TenantId = &v
	return s
}

func (s *ListCollaborationGroupMembersRequest) Validate() error {
	return dara.Validate(s)
}
