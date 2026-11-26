// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserWithGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateUserWithGroupsRequest
	GetAppToken() *string
	SetDisplayName(v string) *CreateUserWithGroupsRequest
	GetDisplayName() *string
	SetOpsUserCode(v string) *CreateUserWithGroupsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateUserWithGroupsRequest
	GetOpsUserId() *string
	SetPasswordEncrypted(v string) *CreateUserWithGroupsRequest
	GetPasswordEncrypted() *string
	SetRoleCodes(v []*string) *CreateUserWithGroupsRequest
	GetRoleCodes() []*string
	SetTenantId(v string) *CreateUserWithGroupsRequest
	GetTenantId() *string
	SetUserGroupIds(v []*string) *CreateUserWithGroupsRequest
	GetUserGroupIds() []*string
	SetWnAccountId(v string) *CreateUserWithGroupsRequest
	GetWnAccountId() *string
}

type CreateUserWithGroupsRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The display name of the user. The name must be unique within the tenant and cannot exceed 100 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
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
	// The Base64-encoded password ciphertext encrypted by using the RSA-OAEP-SHA256 algorithm.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	PasswordEncrypted *string `json:"passwordEncrypted,omitempty" xml:"passwordEncrypted,omitempty"`
	// The list of initial system role codes. If this parameter is not specified, the `APPLICATION_USER` role is assigned by default.
	//
	// example:
	//
	// string_value
	RoleCodes []*string `json:"roleCodes,omitempty" xml:"roleCodes,omitempty" type:"Repeated"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this parameter explicitly by using `--tenant-id`.
	//
	// example:
	//
	// string_value
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The list of initial user group IDs. A maximum of 100 user group IDs can be specified. All user groups must belong to the current tenant.
	//
	// example:
	//
	// string_value
	UserGroupIds []*string `json:"userGroupIds,omitempty" xml:"userGroupIds,omitempty" type:"Repeated"`
	// The WINNEXO logon account. This parameter is a unique identifier and cannot be empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleAccountId
	WnAccountId *string `json:"wnAccountId,omitempty" xml:"wnAccountId,omitempty"`
}

func (s CreateUserWithGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserWithGroupsRequest) GoString() string {
	return s.String()
}

func (s *CreateUserWithGroupsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateUserWithGroupsRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserWithGroupsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateUserWithGroupsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateUserWithGroupsRequest) GetPasswordEncrypted() *string {
	return s.PasswordEncrypted
}

func (s *CreateUserWithGroupsRequest) GetRoleCodes() []*string {
	return s.RoleCodes
}

func (s *CreateUserWithGroupsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateUserWithGroupsRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *CreateUserWithGroupsRequest) GetWnAccountId() *string {
	return s.WnAccountId
}

func (s *CreateUserWithGroupsRequest) SetAppToken(v string) *CreateUserWithGroupsRequest {
	s.AppToken = &v
	return s
}

func (s *CreateUserWithGroupsRequest) SetDisplayName(v string) *CreateUserWithGroupsRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserWithGroupsRequest) SetOpsUserCode(v string) *CreateUserWithGroupsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateUserWithGroupsRequest) SetOpsUserId(v string) *CreateUserWithGroupsRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateUserWithGroupsRequest) SetPasswordEncrypted(v string) *CreateUserWithGroupsRequest {
	s.PasswordEncrypted = &v
	return s
}

func (s *CreateUserWithGroupsRequest) SetRoleCodes(v []*string) *CreateUserWithGroupsRequest {
	s.RoleCodes = v
	return s
}

func (s *CreateUserWithGroupsRequest) SetTenantId(v string) *CreateUserWithGroupsRequest {
	s.TenantId = &v
	return s
}

func (s *CreateUserWithGroupsRequest) SetUserGroupIds(v []*string) *CreateUserWithGroupsRequest {
	s.UserGroupIds = v
	return s
}

func (s *CreateUserWithGroupsRequest) SetWnAccountId(v string) *CreateUserWithGroupsRequest {
	s.WnAccountId = &v
	return s
}

func (s *CreateUserWithGroupsRequest) Validate() error {
	return dara.Validate(s)
}
