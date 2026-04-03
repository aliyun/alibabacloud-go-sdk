// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccessDeniedDetail interface {
	dara.Model
	String() string
	GoString() string
	SetAuthAction(v string) *AccessDeniedDetail
	GetAuthAction() *string
	SetAuthPrincipalDisplayName(v string) *AccessDeniedDetail
	GetAuthPrincipalDisplayName() *string
	SetAuthPrincipalOwnerId(v string) *AccessDeniedDetail
	GetAuthPrincipalOwnerId() *string
	SetAuthPrincipalType(v string) *AccessDeniedDetail
	GetAuthPrincipalType() *string
	SetEncodedDiagnosticInfo(v string) *AccessDeniedDetail
	GetEncodedDiagnosticInfo() *string
	SetNoPermissionType(v string) *AccessDeniedDetail
	GetNoPermissionType() *string
	SetPolicyType(v string) *AccessDeniedDetail
	GetPolicyType() *string
}

type AccessDeniedDetail struct {
	// 被拒绝的 RAM action，如 agentrun:ListTemplates
	//
	// if can be null:
	// true
	AuthAction *string `json:"authAction,omitempty" xml:"authAction,omitempty"`
	// 鉴权主体展示名
	//
	// if can be null:
	// true
	AuthPrincipalDisplayName *string `json:"authPrincipalDisplayName,omitempty" xml:"authPrincipalDisplayName,omitempty"`
	// 鉴权主体所属账号 ID
	//
	// if can be null:
	// true
	AuthPrincipalOwnerId *string `json:"authPrincipalOwnerId,omitempty" xml:"authPrincipalOwnerId,omitempty"`
	// 鉴权主体类型，如 SubUser、AssumedRoleUser
	//
	// if can be null:
	// true
	AuthPrincipalType *string `json:"authPrincipalType,omitempty" xml:"authPrincipalType,omitempty"`
	// Base64 编码的诊断信息，可用于 RAM 控制台自诊断
	//
	// if can be null:
	// true
	EncodedDiagnosticInfo *string `json:"encodedDiagnosticInfo,omitempty" xml:"encodedDiagnosticInfo,omitempty"`
	// 无权限类型：ImplicitDeny 或 ExplicitDeny
	//
	// if can be null:
	// true
	NoPermissionType *string `json:"noPermissionType,omitempty" xml:"noPermissionType,omitempty"`
	// 策略类型，如 ResourceBasedPolicy、IdentityBasedPolicy
	//
	// if can be null:
	// true
	PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty"`
}

func (s AccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s AccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *AccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *AccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *AccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *AccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *AccessDeniedDetail) GetEncodedDiagnosticInfo() *string {
	return s.EncodedDiagnosticInfo
}

func (s *AccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *AccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AccessDeniedDetail) SetAuthAction(v string) *AccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *AccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *AccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *AccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *AccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *AccessDeniedDetail) SetAuthPrincipalType(v string) *AccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *AccessDeniedDetail) SetEncodedDiagnosticInfo(v string) *AccessDeniedDetail {
	s.EncodedDiagnosticInfo = &v
	return s
}

func (s *AccessDeniedDetail) SetNoPermissionType(v string) *AccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *AccessDeniedDetail) SetPolicyType(v string) *AccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *AccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
