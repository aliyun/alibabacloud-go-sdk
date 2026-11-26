// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateUserGroupRequest
	GetAppToken() *string
	SetDescription(v string) *CreateUserGroupRequest
	GetDescription() *string
	SetOpsUserCode(v string) *CreateUserGroupRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateUserGroupRequest
	GetOpsUserId() *string
	SetParentId(v string) *CreateUserGroupRequest
	GetParentId() *string
	SetTenantId(v string) *CreateUserGroupRequest
	GetTenantId() *string
	SetUserGroupName(v string) *CreateUserGroupRequest
	GetUserGroupName() *string
}

type CreateUserGroupRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The description of the user group.
	//
	// example:
	//
	// Sales organization
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The ID of the parent user group. If this parameter is not specified, the user group is created under the root node.
	//
	// example:
	//
	// 7ea8973f-7a5c-4e8a-956b-4fe0e7e2eb11
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The tenant ID. This is a common parameter that can be explicitly passed by using the `--tenant-id` option in winnexo-cli.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The name of the user group. The name must be unique under the same parent node.
	//
	// This parameter is required.
	//
	// example:
	//
	// East China Sales
	UserGroupName *string `json:"userGroupName,omitempty" xml:"userGroupName,omitempty"`
}

func (s CreateUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateUserGroupRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateUserGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUserGroupRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateUserGroupRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateUserGroupRequest) GetParentId() *string {
	return s.ParentId
}

func (s *CreateUserGroupRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateUserGroupRequest) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *CreateUserGroupRequest) SetAppToken(v string) *CreateUserGroupRequest {
	s.AppToken = &v
	return s
}

func (s *CreateUserGroupRequest) SetDescription(v string) *CreateUserGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateUserGroupRequest) SetOpsUserCode(v string) *CreateUserGroupRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateUserGroupRequest) SetOpsUserId(v string) *CreateUserGroupRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateUserGroupRequest) SetParentId(v string) *CreateUserGroupRequest {
	s.ParentId = &v
	return s
}

func (s *CreateUserGroupRequest) SetTenantId(v string) *CreateUserGroupRequest {
	s.TenantId = &v
	return s
}

func (s *CreateUserGroupRequest) SetUserGroupName(v string) *CreateUserGroupRequest {
	s.UserGroupName = &v
	return s
}

func (s *CreateUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
