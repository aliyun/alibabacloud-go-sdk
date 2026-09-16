// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConnectorAuthenticationUpdateInput interface {
	dara.Model
	String() string
	GoString() string
	SetBasic(v *ConnectorAuthenticationUpdateInputBasic) *ConnectorAuthenticationUpdateInput
	GetBasic() *ConnectorAuthenticationUpdateInputBasic
	SetBotToken(v *ConnectorAuthenticationUpdateInputBotToken) *ConnectorAuthenticationUpdateInput
	GetBotToken() *ConnectorAuthenticationUpdateInputBotToken
	SetOauth(v *ConnectorAuthenticationUpdateInputOauth) *ConnectorAuthenticationUpdateInput
	GetOauth() *ConnectorAuthenticationUpdateInputOauth
	SetPatToken(v *ConnectorAuthenticationUpdateInputPatToken) *ConnectorAuthenticationUpdateInput
	GetPatToken() *ConnectorAuthenticationUpdateInputPatToken
	SetRole(v *ConnectorAuthenticationUpdateInputRole) *ConnectorAuthenticationUpdateInput
	GetRole() *ConnectorAuthenticationUpdateInputRole
	SetSatellite(v *ConnectorAuthenticationUpdateInputSatellite) *ConnectorAuthenticationUpdateInput
	GetSatellite() *ConnectorAuthenticationUpdateInputSatellite
	SetType(v string) *ConnectorAuthenticationUpdateInput
	GetType() *string
}

type ConnectorAuthenticationUpdateInput struct {
	// The configuration that uses a username and password to replace the existing authentication configuration.
	Basic *ConnectorAuthenticationUpdateInputBasic `json:"basic,omitempty" xml:"basic,omitempty" type:"Struct"`
	// The configuration that uses a bot token and signing key to replace the existing authentication configuration.
	BotToken *ConnectorAuthenticationUpdateInputBotToken `json:"botToken,omitempty" xml:"botToken,omitempty" type:"Struct"`
	// The configuration that uses an OAuth client identity to replace the existing authentication configuration.
	Oauth *ConnectorAuthenticationUpdateInputOauth `json:"oauth,omitempty" xml:"oauth,omitempty" type:"Struct"`
	// The configuration that uses a personal access token to replace the existing authentication configuration.
	PatToken *ConnectorAuthenticationUpdateInputPatToken `json:"patToken,omitempty" xml:"patToken,omitempty" type:"Struct"`
	// The configuration that uses a RAM role ARN to replace the existing authentication configuration.
	Role *ConnectorAuthenticationUpdateInputRole `json:"role,omitempty" xml:"role,omitempty" type:"Struct"`
	// The configuration that uses a local credential binding to replace the existing authentication configuration.
	Satellite *ConnectorAuthenticationUpdateInputSatellite `json:"satellite,omitempty" xml:"satellite,omitempty" type:"Struct"`
	// Authentication type
	//
	// This parameter is required.
	//
	// example:
	//
	// DEFAULT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ConnectorAuthenticationUpdateInput) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthenticationUpdateInput) GoString() string {
	return s.String()
}

func (s *ConnectorAuthenticationUpdateInput) GetBasic() *ConnectorAuthenticationUpdateInputBasic {
	return s.Basic
}

func (s *ConnectorAuthenticationUpdateInput) GetBotToken() *ConnectorAuthenticationUpdateInputBotToken {
	return s.BotToken
}

func (s *ConnectorAuthenticationUpdateInput) GetOauth() *ConnectorAuthenticationUpdateInputOauth {
	return s.Oauth
}

func (s *ConnectorAuthenticationUpdateInput) GetPatToken() *ConnectorAuthenticationUpdateInputPatToken {
	return s.PatToken
}

func (s *ConnectorAuthenticationUpdateInput) GetRole() *ConnectorAuthenticationUpdateInputRole {
	return s.Role
}

func (s *ConnectorAuthenticationUpdateInput) GetSatellite() *ConnectorAuthenticationUpdateInputSatellite {
	return s.Satellite
}

func (s *ConnectorAuthenticationUpdateInput) GetType() *string {
	return s.Type
}

func (s *ConnectorAuthenticationUpdateInput) SetBasic(v *ConnectorAuthenticationUpdateInputBasic) *ConnectorAuthenticationUpdateInput {
	s.Basic = v
	return s
}

func (s *ConnectorAuthenticationUpdateInput) SetBotToken(v *ConnectorAuthenticationUpdateInputBotToken) *ConnectorAuthenticationUpdateInput {
	s.BotToken = v
	return s
}

func (s *ConnectorAuthenticationUpdateInput) SetOauth(v *ConnectorAuthenticationUpdateInputOauth) *ConnectorAuthenticationUpdateInput {
	s.Oauth = v
	return s
}

func (s *ConnectorAuthenticationUpdateInput) SetPatToken(v *ConnectorAuthenticationUpdateInputPatToken) *ConnectorAuthenticationUpdateInput {
	s.PatToken = v
	return s
}

func (s *ConnectorAuthenticationUpdateInput) SetRole(v *ConnectorAuthenticationUpdateInputRole) *ConnectorAuthenticationUpdateInput {
	s.Role = v
	return s
}

func (s *ConnectorAuthenticationUpdateInput) SetSatellite(v *ConnectorAuthenticationUpdateInputSatellite) *ConnectorAuthenticationUpdateInput {
	s.Satellite = v
	return s
}

func (s *ConnectorAuthenticationUpdateInput) SetType(v string) *ConnectorAuthenticationUpdateInput {
	s.Type = &v
	return s
}

func (s *ConnectorAuthenticationUpdateInput) Validate() error {
	if s.Basic != nil {
		if err := s.Basic.Validate(); err != nil {
			return err
		}
	}
	if s.BotToken != nil {
		if err := s.BotToken.Validate(); err != nil {
			return err
		}
	}
	if s.Oauth != nil {
		if err := s.Oauth.Validate(); err != nil {
			return err
		}
	}
	if s.PatToken != nil {
		if err := s.PatToken.Validate(); err != nil {
			return err
		}
	}
	if s.Role != nil {
		if err := s.Role.Validate(); err != nil {
			return err
		}
	}
	if s.Satellite != nil {
		if err := s.Satellite.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConnectorAuthenticationUpdateInputBasic struct {
	// Replacement password
	//
	// example:
	//
	// example-password
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// Username
	//
	// This parameter is required.
	//
	// example:
	//
	// starops
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ConnectorAuthenticationUpdateInputBasic) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthenticationUpdateInputBasic) GoString() string {
	return s.String()
}

func (s *ConnectorAuthenticationUpdateInputBasic) GetPassword() *string {
	return s.Password
}

func (s *ConnectorAuthenticationUpdateInputBasic) GetUsername() *string {
	return s.Username
}

func (s *ConnectorAuthenticationUpdateInputBasic) SetPassword(v string) *ConnectorAuthenticationUpdateInputBasic {
	s.Password = &v
	return s
}

func (s *ConnectorAuthenticationUpdateInputBasic) SetUsername(v string) *ConnectorAuthenticationUpdateInputBasic {
	s.Username = &v
	return s
}

func (s *ConnectorAuthenticationUpdateInputBasic) Validate() error {
	return dara.Validate(s)
}

type ConnectorAuthenticationUpdateInputBotToken struct {
	// Replacement bot token
	//
	// This parameter is required.
	//
	// example:
	//
	// example-bot-token
	BotToken *string `json:"botToken,omitempty" xml:"botToken,omitempty"`
	// Replacement signing secret
	//
	// example:
	//
	// example-signing-secret
	SigningSecret *string `json:"signingSecret,omitempty" xml:"signingSecret,omitempty"`
}

func (s ConnectorAuthenticationUpdateInputBotToken) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthenticationUpdateInputBotToken) GoString() string {
	return s.String()
}

func (s *ConnectorAuthenticationUpdateInputBotToken) GetBotToken() *string {
	return s.BotToken
}

func (s *ConnectorAuthenticationUpdateInputBotToken) GetSigningSecret() *string {
	return s.SigningSecret
}

func (s *ConnectorAuthenticationUpdateInputBotToken) SetBotToken(v string) *ConnectorAuthenticationUpdateInputBotToken {
	s.BotToken = &v
	return s
}

func (s *ConnectorAuthenticationUpdateInputBotToken) SetSigningSecret(v string) *ConnectorAuthenticationUpdateInputBotToken {
	s.SigningSecret = &v
	return s
}

func (s *ConnectorAuthenticationUpdateInputBotToken) Validate() error {
	return dara.Validate(s)
}

type ConnectorAuthenticationUpdateInputOauth struct {
	// OAuth client ID
	//
	// This parameter is required.
	//
	// example:
	//
	// client-id
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// Replacement OAuth client secret
	//
	// example:
	//
	// example-client-secret
	ClientSecret *string `json:"clientSecret,omitempty" xml:"clientSecret,omitempty"`
}

func (s ConnectorAuthenticationUpdateInputOauth) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthenticationUpdateInputOauth) GoString() string {
	return s.String()
}

func (s *ConnectorAuthenticationUpdateInputOauth) GetClientId() *string {
	return s.ClientId
}

func (s *ConnectorAuthenticationUpdateInputOauth) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *ConnectorAuthenticationUpdateInputOauth) SetClientId(v string) *ConnectorAuthenticationUpdateInputOauth {
	s.ClientId = &v
	return s
}

func (s *ConnectorAuthenticationUpdateInputOauth) SetClientSecret(v string) *ConnectorAuthenticationUpdateInputOauth {
	s.ClientSecret = &v
	return s
}

func (s *ConnectorAuthenticationUpdateInputOauth) Validate() error {
	return dara.Validate(s)
}

type ConnectorAuthenticationUpdateInputPatToken struct {
	// The personal access token used to replace the existing credential.
	//
	// example:
	//
	// example-personal-access-token
	PatToken *string `json:"patToken,omitempty" xml:"patToken,omitempty"`
}

func (s ConnectorAuthenticationUpdateInputPatToken) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthenticationUpdateInputPatToken) GoString() string {
	return s.String()
}

func (s *ConnectorAuthenticationUpdateInputPatToken) GetPatToken() *string {
	return s.PatToken
}

func (s *ConnectorAuthenticationUpdateInputPatToken) SetPatToken(v string) *ConnectorAuthenticationUpdateInputPatToken {
	s.PatToken = &v
	return s
}

func (s *ConnectorAuthenticationUpdateInputPatToken) Validate() error {
	return dara.Validate(s)
}

type ConnectorAuthenticationUpdateInputRole struct {
	// Role ARN
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::1234567890123456:role/starops-reader
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
}

func (s ConnectorAuthenticationUpdateInputRole) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthenticationUpdateInputRole) GoString() string {
	return s.String()
}

func (s *ConnectorAuthenticationUpdateInputRole) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ConnectorAuthenticationUpdateInputRole) SetRoleArn(v string) *ConnectorAuthenticationUpdateInputRole {
	s.RoleArn = &v
	return s
}

func (s *ConnectorAuthenticationUpdateInputRole) Validate() error {
	return dara.Validate(s)
}

type ConnectorAuthenticationUpdateInputSatellite struct {
	// Local credential binding name
	//
	// This parameter is required.
	//
	// example:
	//
	// private-gitlab
	BindingName *string `json:"bindingName,omitempty" xml:"bindingName,omitempty"`
}

func (s ConnectorAuthenticationUpdateInputSatellite) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthenticationUpdateInputSatellite) GoString() string {
	return s.String()
}

func (s *ConnectorAuthenticationUpdateInputSatellite) GetBindingName() *string {
	return s.BindingName
}

func (s *ConnectorAuthenticationUpdateInputSatellite) SetBindingName(v string) *ConnectorAuthenticationUpdateInputSatellite {
	s.BindingName = &v
	return s
}

func (s *ConnectorAuthenticationUpdateInputSatellite) Validate() error {
	return dara.Validate(s)
}
