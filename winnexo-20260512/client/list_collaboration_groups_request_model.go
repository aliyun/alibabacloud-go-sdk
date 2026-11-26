// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCollaborationGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListCollaborationGroupsRequest
	GetAppToken() *string
	SetKeyword(v string) *ListCollaborationGroupsRequest
	GetKeyword() *string
	SetOpsUserCode(v string) *ListCollaborationGroupsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListCollaborationGroupsRequest
	GetOpsUserId() *string
	SetPage(v int64) *ListCollaborationGroupsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListCollaborationGroupsRequest
	GetPageSize() *int64
	SetTenantId(v string) *ListCollaborationGroupsRequest
	GetTenantId() *string
}

type ListCollaborationGroupsRequest struct {
	// The platform application token issued by the console. If specified, the operation is executed with the user identity specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The search keyword. This parameter is optional and can be up to 256 characters in length.
	//
	// example:
	//
	// Workflow
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

func (s ListCollaborationGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCollaborationGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListCollaborationGroupsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListCollaborationGroupsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListCollaborationGroupsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListCollaborationGroupsRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *ListCollaborationGroupsRequest) SetAppToken(v string) *ListCollaborationGroupsRequest {
	s.AppToken = &v
	return s
}

func (s *ListCollaborationGroupsRequest) SetKeyword(v string) *ListCollaborationGroupsRequest {
	s.Keyword = &v
	return s
}

func (s *ListCollaborationGroupsRequest) SetOpsUserCode(v string) *ListCollaborationGroupsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListCollaborationGroupsRequest) SetOpsUserId(v string) *ListCollaborationGroupsRequest {
	s.OpsUserId = &v
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
