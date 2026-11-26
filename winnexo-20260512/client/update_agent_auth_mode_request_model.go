// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentAuthModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UpdateAgentAuthModeRequest
	GetAppToken() *string
	SetAuthMode(v string) *UpdateAgentAuthModeRequest
	GetAuthMode() *string
	SetOperatingObjectName(v string) *UpdateAgentAuthModeRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *UpdateAgentAuthModeRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateAgentAuthModeRequest
	GetOpsUserId() *string
	SetTenantId(v string) *UpdateAgentAuthModeRequest
	GetTenantId() *string
}

type UpdateAgentAuthModeRequest struct {
	// The platform application token issued by the console. If specified, the operation is executed as the user identified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The authorization mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// SPECIFIED_USERS
	AuthMode *string `json:"authMode,omitempty" xml:"authMode,omitempty"`
	// The name of the digital employee, which is the name of the operating object. This parameter is optional.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
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
	// The ID of the tenant where the operation takes effect.
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

func (s *UpdateAgentAuthModeRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UpdateAgentAuthModeRequest) GetAuthMode() *string {
	return s.AuthMode
}

func (s *UpdateAgentAuthModeRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *UpdateAgentAuthModeRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateAgentAuthModeRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UpdateAgentAuthModeRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateAgentAuthModeRequest) SetAppToken(v string) *UpdateAgentAuthModeRequest {
	s.AppToken = &v
	return s
}

func (s *UpdateAgentAuthModeRequest) SetAuthMode(v string) *UpdateAgentAuthModeRequest {
	s.AuthMode = &v
	return s
}

func (s *UpdateAgentAuthModeRequest) SetOperatingObjectName(v string) *UpdateAgentAuthModeRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *UpdateAgentAuthModeRequest) SetOpsUserCode(v string) *UpdateAgentAuthModeRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateAgentAuthModeRequest) SetOpsUserId(v string) *UpdateAgentAuthModeRequest {
	s.OpsUserId = &v
	return s
}

func (s *UpdateAgentAuthModeRequest) SetTenantId(v string) *UpdateAgentAuthModeRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateAgentAuthModeRequest) Validate() error {
	return dara.Validate(s)
}
