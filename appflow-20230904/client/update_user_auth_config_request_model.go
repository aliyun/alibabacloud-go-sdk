// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserAuthConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfig(v string) *UpdateUserAuthConfigRequest
	GetAuthConfig() *string
	SetAuthConfigId(v string) *UpdateUserAuthConfigRequest
	GetAuthConfigId() *string
	SetAuthConfigName(v string) *UpdateUserAuthConfigRequest
	GetAuthConfigName() *string
	SetConnectorId(v string) *UpdateUserAuthConfigRequest
	GetConnectorId() *string
	SetConnectorVersion(v string) *UpdateUserAuthConfigRequest
	GetConnectorVersion() *string
}

type UpdateUserAuthConfigRequest struct {
	// The user authentication credential, provided as a JSON string.
	//
	// example:
	//
	// {\\"apiKey\\": \\"************\\"}
	AuthConfig *string `json:"AuthConfig,omitempty" xml:"AuthConfig,omitempty"`
	// The ID of the credential.
	//
	// This parameter is required.
	//
	// example:
	//
	// uac-42b60d53bcce466d9d08
	AuthConfigId *string `json:"AuthConfigId,omitempty" xml:"AuthConfigId,omitempty"`
	// The name of the credential.
	//
	// example:
	//
	// name
	AuthConfigName *string `json:"AuthConfigName,omitempty" xml:"AuthConfigName,omitempty"`
	// The ID of the connector.
	//
	// This parameter is required.
	//
	// example:
	//
	// connector-6b8df2297dca4a5f8f15
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The connector version.
	//
	// example:
	//
	// 1
	ConnectorVersion *string `json:"ConnectorVersion,omitempty" xml:"ConnectorVersion,omitempty"`
}

func (s UpdateUserAuthConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserAuthConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserAuthConfigRequest) GetAuthConfig() *string {
	return s.AuthConfig
}

func (s *UpdateUserAuthConfigRequest) GetAuthConfigId() *string {
	return s.AuthConfigId
}

func (s *UpdateUserAuthConfigRequest) GetAuthConfigName() *string {
	return s.AuthConfigName
}

func (s *UpdateUserAuthConfigRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *UpdateUserAuthConfigRequest) GetConnectorVersion() *string {
	return s.ConnectorVersion
}

func (s *UpdateUserAuthConfigRequest) SetAuthConfig(v string) *UpdateUserAuthConfigRequest {
	s.AuthConfig = &v
	return s
}

func (s *UpdateUserAuthConfigRequest) SetAuthConfigId(v string) *UpdateUserAuthConfigRequest {
	s.AuthConfigId = &v
	return s
}

func (s *UpdateUserAuthConfigRequest) SetAuthConfigName(v string) *UpdateUserAuthConfigRequest {
	s.AuthConfigName = &v
	return s
}

func (s *UpdateUserAuthConfigRequest) SetConnectorId(v string) *UpdateUserAuthConfigRequest {
	s.ConnectorId = &v
	return s
}

func (s *UpdateUserAuthConfigRequest) SetConnectorVersion(v string) *UpdateUserAuthConfigRequest {
	s.ConnectorVersion = &v
	return s
}

func (s *UpdateUserAuthConfigRequest) Validate() error {
	return dara.Validate(s)
}
