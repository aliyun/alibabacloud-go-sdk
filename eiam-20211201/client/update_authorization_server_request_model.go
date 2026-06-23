// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationServerId(v string) *UpdateAuthorizationServerRequest
	GetAuthorizationServerId() *string
	SetAuthorizationServerName(v string) *UpdateAuthorizationServerRequest
	GetAuthorizationServerName() *string
	SetClientToken(v string) *UpdateAuthorizationServerRequest
	GetClientToken() *string
	SetInstanceId(v string) *UpdateAuthorizationServerRequest
	GetInstanceId() *string
	SetIssuerDomain(v string) *UpdateAuthorizationServerRequest
	GetIssuerDomain() *string
	SetIssuerMode(v string) *UpdateAuthorizationServerRequest
	GetIssuerMode() *string
}

type UpdateAuthorizationServerRequest struct {
	// The authorization server ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// iauths_system
	AuthorizationServerId *string `json:"AuthorizationServerId,omitempty" xml:"AuthorizationServerId,omitempty"`
	// The name of the authorization server.
	//
	// example:
	//
	// my_authorization_server
	AuthorizationServerName *string `json:"AuthorizationServerName,omitempty" xml:"AuthorizationServerName,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate a parameter value, but you must make sure that the value is unique among different requests. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see References: [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// This parameter is required.
	//
	// example:
	//
	// abc123xyz
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The domain name used by the issuer.
	//
	// example:
	//
	// xxxx.aliyunidaas.com
	IssuerDomain *string `json:"IssuerDomain,omitempty" xml:"IssuerDomain,omitempty"`
	// The issuer mode.
	//
	// example:
	//
	// dynamic
	IssuerMode *string `json:"IssuerMode,omitempty" xml:"IssuerMode,omitempty"`
}

func (s UpdateAuthorizationServerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationServerRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationServerRequest) GetAuthorizationServerId() *string {
	return s.AuthorizationServerId
}

func (s *UpdateAuthorizationServerRequest) GetAuthorizationServerName() *string {
	return s.AuthorizationServerName
}

func (s *UpdateAuthorizationServerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAuthorizationServerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAuthorizationServerRequest) GetIssuerDomain() *string {
	return s.IssuerDomain
}

func (s *UpdateAuthorizationServerRequest) GetIssuerMode() *string {
	return s.IssuerMode
}

func (s *UpdateAuthorizationServerRequest) SetAuthorizationServerId(v string) *UpdateAuthorizationServerRequest {
	s.AuthorizationServerId = &v
	return s
}

func (s *UpdateAuthorizationServerRequest) SetAuthorizationServerName(v string) *UpdateAuthorizationServerRequest {
	s.AuthorizationServerName = &v
	return s
}

func (s *UpdateAuthorizationServerRequest) SetClientToken(v string) *UpdateAuthorizationServerRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAuthorizationServerRequest) SetInstanceId(v string) *UpdateAuthorizationServerRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAuthorizationServerRequest) SetIssuerDomain(v string) *UpdateAuthorizationServerRequest {
	s.IssuerDomain = &v
	return s
}

func (s *UpdateAuthorizationServerRequest) SetIssuerMode(v string) *UpdateAuthorizationServerRequest {
	s.IssuerMode = &v
	return s
}

func (s *UpdateAuthorizationServerRequest) Validate() error {
	return dara.Validate(s)
}
