// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserAuthConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfig(v string) *CreateUserAuthConfigRequest
	GetAuthConfig() *string
	SetAuthConfigName(v string) *CreateUserAuthConfigRequest
	GetAuthConfigName() *string
	SetAuthType(v string) *CreateUserAuthConfigRequest
	GetAuthType() *string
	SetConnectorId(v string) *CreateUserAuthConfigRequest
	GetConnectorId() *string
	SetConnectorVersion(v string) *CreateUserAuthConfigRequest
	GetConnectorVersion() *string
}

type CreateUserAuthConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ***
	AuthConfig *string `json:"AuthConfig,omitempty" xml:"AuthConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bailian-01ce5586-420f-4cbf-9392-7001a1c33bc1
	AuthConfigName *string `json:"AuthConfigName,omitempty" xml:"AuthConfigName,omitempty"`
	// example:
	//
	// DingdingAccessToken
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// connector-3c60c6e123e146fbb6f8
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// example:
	//
	// 1
	ConnectorVersion *string `json:"ConnectorVersion,omitempty" xml:"ConnectorVersion,omitempty"`
}

func (s CreateUserAuthConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserAuthConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateUserAuthConfigRequest) GetAuthConfig() *string {
	return s.AuthConfig
}

func (s *CreateUserAuthConfigRequest) GetAuthConfigName() *string {
	return s.AuthConfigName
}

func (s *CreateUserAuthConfigRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *CreateUserAuthConfigRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *CreateUserAuthConfigRequest) GetConnectorVersion() *string {
	return s.ConnectorVersion
}

func (s *CreateUserAuthConfigRequest) SetAuthConfig(v string) *CreateUserAuthConfigRequest {
	s.AuthConfig = &v
	return s
}

func (s *CreateUserAuthConfigRequest) SetAuthConfigName(v string) *CreateUserAuthConfigRequest {
	s.AuthConfigName = &v
	return s
}

func (s *CreateUserAuthConfigRequest) SetAuthType(v string) *CreateUserAuthConfigRequest {
	s.AuthType = &v
	return s
}

func (s *CreateUserAuthConfigRequest) SetConnectorId(v string) *CreateUserAuthConfigRequest {
	s.ConnectorId = &v
	return s
}

func (s *CreateUserAuthConfigRequest) SetConnectorVersion(v string) *CreateUserAuthConfigRequest {
	s.ConnectorVersion = &v
	return s
}

func (s *CreateUserAuthConfigRequest) Validate() error {
	return dara.Validate(s)
}
