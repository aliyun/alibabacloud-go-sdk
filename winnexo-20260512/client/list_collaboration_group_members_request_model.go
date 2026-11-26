// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollaborationGroupMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListCollaborationGroupMembersRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *ListCollaborationGroupMembersRequest
	GetCollaborationGroupId() *string
	SetKeyword(v string) *ListCollaborationGroupMembersRequest
	GetKeyword() *string
	SetOpsUserCode(v string) *ListCollaborationGroupMembersRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListCollaborationGroupMembersRequest
	GetOpsUserId() *string
	SetPage(v int64) *ListCollaborationGroupMembersRequest
	GetPage() *int64
	SetPageSize(v int64) *ListCollaborationGroupMembersRequest
	GetPageSize() *int64
	SetTenantId(v string) *ListCollaborationGroupMembersRequest
	GetTenantId() *string
}

type ListCollaborationGroupMembersRequest struct {
	// The platform application token issued by the console. If specified, the operation is executed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
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
	// The proxy user code, which is unique within the tenant. This parameter is used only when the application credential is configured as source_uid.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The proxy user ID, which is the WINNEXO platform user ID. This parameter is used only when the application credential is configured as winnexo_uid.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The page number, starting from 1.
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

func (s *ListCollaborationGroupMembersRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListCollaborationGroupMembersRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListCollaborationGroupMembersRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListCollaborationGroupMembersRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListCollaborationGroupMembersRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *ListCollaborationGroupMembersRequest) SetAppToken(v string) *ListCollaborationGroupMembersRequest {
	s.AppToken = &v
	return s
}

func (s *ListCollaborationGroupMembersRequest) SetCollaborationGroupId(v string) *ListCollaborationGroupMembersRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListCollaborationGroupMembersRequest) SetKeyword(v string) *ListCollaborationGroupMembersRequest {
	s.Keyword = &v
	return s
}

func (s *ListCollaborationGroupMembersRequest) SetOpsUserCode(v string) *ListCollaborationGroupMembersRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListCollaborationGroupMembersRequest) SetOpsUserId(v string) *ListCollaborationGroupMembersRequest {
	s.OpsUserId = &v
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
