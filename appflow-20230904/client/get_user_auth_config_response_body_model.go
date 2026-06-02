// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAuthConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserAuthConfigResponseBody
	GetRequestId() *string
	SetUserAuthConfig(v *GetUserAuthConfigResponseBodyUserAuthConfig) *GetUserAuthConfigResponseBody
	GetUserAuthConfig() *GetUserAuthConfigResponseBodyUserAuthConfig
}

type GetUserAuthConfigResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 458CB9FE-8C71-58A8-AD49-97EF28D58FAB
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserAuthConfig *GetUserAuthConfigResponseBodyUserAuthConfig `json:"UserAuthConfig,omitempty" xml:"UserAuthConfig,omitempty" type:"Struct"`
}

func (s GetUserAuthConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserAuthConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserAuthConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserAuthConfigResponseBody) GetUserAuthConfig() *GetUserAuthConfigResponseBodyUserAuthConfig {
	return s.UserAuthConfig
}

func (s *GetUserAuthConfigResponseBody) SetRequestId(v string) *GetUserAuthConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserAuthConfigResponseBody) SetUserAuthConfig(v *GetUserAuthConfigResponseBodyUserAuthConfig) *GetUserAuthConfigResponseBody {
	s.UserAuthConfig = v
	return s
}

func (s *GetUserAuthConfigResponseBody) Validate() error {
	if s.UserAuthConfig != nil {
		if err := s.UserAuthConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserAuthConfigResponseBodyUserAuthConfig struct {
	// example:
	//
	// {\\"apiKey":\\"xxxxxxxxxx\\"}
	AuthConfig *string `json:"AuthConfig,omitempty" xml:"AuthConfig,omitempty"`
	// example:
	//
	// uac-111111111
	AuthConfigId *string `json:"AuthConfigId,omitempty" xml:"AuthConfigId,omitempty"`
	// example:
	//
	// dingtlak_name
	AuthConfigName *string `json:"AuthConfigName,omitempty" xml:"AuthConfigName,omitempty"`
	// example:
	//
	// ApiKey
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// connector-172176821387
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// example:
	//
	// 1
	ConnectorVersion *string `json:"ConnectorVersion,omitempty" xml:"ConnectorVersion,omitempty"`
}

func (s GetUserAuthConfigResponseBodyUserAuthConfig) String() string {
	return dara.Prettify(s)
}

func (s GetUserAuthConfigResponseBodyUserAuthConfig) GoString() string {
	return s.String()
}

func (s *GetUserAuthConfigResponseBodyUserAuthConfig) GetAuthConfig() *string {
	return s.AuthConfig
}

func (s *GetUserAuthConfigResponseBodyUserAuthConfig) GetAuthConfigId() *string {
	return s.AuthConfigId
}

func (s *GetUserAuthConfigResponseBodyUserAuthConfig) GetAuthConfigName() *string {
	return s.AuthConfigName
}

func (s *GetUserAuthConfigResponseBodyUserAuthConfig) GetAuthType() *string {
	return s.AuthType
}

func (s *GetUserAuthConfigResponseBodyUserAuthConfig) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *GetUserAuthConfigResponseBodyUserAuthConfig) GetConnectorVersion() *string {
	return s.ConnectorVersion
}

func (s *GetUserAuthConfigResponseBodyUserAuthConfig) SetAuthConfig(v string) *GetUserAuthConfigResponseBodyUserAuthConfig {
	s.AuthConfig = &v
	return s
}

func (s *GetUserAuthConfigResponseBodyUserAuthConfig) SetAuthConfigId(v string) *GetUserAuthConfigResponseBodyUserAuthConfig {
	s.AuthConfigId = &v
	return s
}

func (s *GetUserAuthConfigResponseBodyUserAuthConfig) SetAuthConfigName(v string) *GetUserAuthConfigResponseBodyUserAuthConfig {
	s.AuthConfigName = &v
	return s
}

func (s *GetUserAuthConfigResponseBodyUserAuthConfig) SetAuthType(v string) *GetUserAuthConfigResponseBodyUserAuthConfig {
	s.AuthType = &v
	return s
}

func (s *GetUserAuthConfigResponseBodyUserAuthConfig) SetConnectorId(v string) *GetUserAuthConfigResponseBodyUserAuthConfig {
	s.ConnectorId = &v
	return s
}

func (s *GetUserAuthConfigResponseBodyUserAuthConfig) SetConnectorVersion(v string) *GetUserAuthConfigResponseBodyUserAuthConfig {
	s.ConnectorVersion = &v
	return s
}

func (s *GetUserAuthConfigResponseBodyUserAuthConfig) Validate() error {
	return dara.Validate(s)
}
