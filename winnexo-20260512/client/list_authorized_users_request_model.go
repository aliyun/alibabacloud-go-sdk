// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListAuthorizedUsersRequest
	GetAppToken() *string
	SetGranteeType(v string) *ListAuthorizedUsersRequest
	GetGranteeType() *string
	SetKeyword(v string) *ListAuthorizedUsersRequest
	GetKeyword() *string
	SetOperatingObjectName(v string) *ListAuthorizedUsersRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *ListAuthorizedUsersRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListAuthorizedUsersRequest
	GetOpsUserId() *string
	SetPermission(v string) *ListAuthorizedUsersRequest
	GetPermission() *string
	SetTenantId(v string) *ListAuthorizedUsersRequest
	GetTenantId() *string
}

type ListAuthorizedUsersRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The filter type. Valid values:
	//
	// - USER
	//
	// - USER_GROUP
	//
	// If this parameter is not specified, all types are returned.
	//
	// example:
	//
	// USER
	GranteeType *string `json:"granteeType,omitempty" xml:"granteeType,omitempty"`
	// The search keyword.
	//
	// example:
	//
	// SampleKeyword
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The name of the digital employee.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
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
	// Permission
	//
	// example:
	//
	// USE
	Permission *string `json:"permission,omitempty" xml:"permission,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListAuthorizedUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedUsersRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedUsersRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListAuthorizedUsersRequest) GetGranteeType() *string {
	return s.GranteeType
}

func (s *ListAuthorizedUsersRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAuthorizedUsersRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListAuthorizedUsersRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListAuthorizedUsersRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListAuthorizedUsersRequest) GetPermission() *string {
	return s.Permission
}

func (s *ListAuthorizedUsersRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAuthorizedUsersRequest) SetAppToken(v string) *ListAuthorizedUsersRequest {
	s.AppToken = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetGranteeType(v string) *ListAuthorizedUsersRequest {
	s.GranteeType = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetKeyword(v string) *ListAuthorizedUsersRequest {
	s.Keyword = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetOperatingObjectName(v string) *ListAuthorizedUsersRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetOpsUserCode(v string) *ListAuthorizedUsersRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetOpsUserId(v string) *ListAuthorizedUsersRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetPermission(v string) *ListAuthorizedUsersRequest {
	s.Permission = &v
	return s
}

func (s *ListAuthorizedUsersRequest) SetTenantId(v string) *ListAuthorizedUsersRequest {
	s.TenantId = &v
	return s
}

func (s *ListAuthorizedUsersRequest) Validate() error {
	return dara.Validate(s)
}
