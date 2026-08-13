// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentAuthModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMode(v string) *UpdateAgentAuthModeRequest
	GetAuthMode() *string
	SetOperatingObjectName(v string) *UpdateAgentAuthModeRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *UpdateAgentAuthModeRequest
	GetTenantId() *string
}

type UpdateAgentAuthModeRequest struct {
	// 使用权限授权模式：SPECIFIED_USERS=指定用户（需显式授权），ALL_USERS=所有用户（无需授权即可使用）
	//
	// This parameter is required.
	//
	// example:
	//
	// SPECIFIED_USERS
	AuthMode *string `json:"authMode,omitempty" xml:"authMode,omitempty"`
	// 数字员工名称
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateAgentAuthModeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentAuthModeRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentAuthModeRequest) GetAuthMode() *string {
	return s.AuthMode
}

func (s *UpdateAgentAuthModeRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *UpdateAgentAuthModeRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateAgentAuthModeRequest) SetAuthMode(v string) *UpdateAgentAuthModeRequest {
	s.AuthMode = &v
	return s
}

func (s *UpdateAgentAuthModeRequest) SetOperatingObjectName(v string) *UpdateAgentAuthModeRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *UpdateAgentAuthModeRequest) SetTenantId(v string) *UpdateAgentAuthModeRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateAgentAuthModeRequest) Validate() error {
	return dara.Validate(s)
}
