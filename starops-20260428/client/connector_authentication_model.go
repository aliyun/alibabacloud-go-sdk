// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConnectorAuthentication interface {
	dara.Model
	String() string
	GoString() string
	SetBasic(v *ConnectorAuthenticationBasic) *ConnectorAuthentication
	GetBasic() *ConnectorAuthenticationBasic
	SetBotToken(v *ConnectorAuthenticationBotToken) *ConnectorAuthentication
	GetBotToken() *ConnectorAuthenticationBotToken
	SetOauth(v *ConnectorAuthenticationOauth) *ConnectorAuthentication
	GetOauth() *ConnectorAuthenticationOauth
	SetRole(v *ConnectorAuthenticationRole) *ConnectorAuthentication
	GetRole() *ConnectorAuthenticationRole
	SetSatellite(v *ConnectorAuthenticationSatellite) *ConnectorAuthentication
	GetSatellite() *ConnectorAuthenticationSatellite
	SetType(v string) *ConnectorAuthentication
	GetType() *string
}

type ConnectorAuthentication struct {
	// The security identity information for basic authentication, excluding the password.
	Basic *ConnectorAuthenticationBasic `json:"basic,omitempty" xml:"basic,omitempty" type:"Struct"`
	// The security identity information for the bot, excluding the token and signing key.
	BotToken *ConnectorAuthenticationBotToken `json:"botToken,omitempty" xml:"botToken,omitempty" type:"Struct"`
	// The security identity information for OAuth, excluding the client secret.
	Oauth *ConnectorAuthenticationOauth `json:"oauth,omitempty" xml:"oauth,omitempty" type:"Struct"`
	// The security identity information based on the RAM role ARN.
	Role *ConnectorAuthenticationRole `json:"role,omitempty" xml:"role,omitempty" type:"Struct"`
	// The security identity information based on local credential binding.
	Satellite *ConnectorAuthenticationSatellite `json:"satellite,omitempty" xml:"satellite,omitempty" type:"Struct"`
	// Authentication type
	//
	// This parameter is required.
	//
	// example:
	//
	// DEFAULT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ConnectorAuthentication) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthentication) GoString() string {
	return s.String()
}

func (s *ConnectorAuthentication) GetBasic() *ConnectorAuthenticationBasic {
	return s.Basic
}

func (s *ConnectorAuthentication) GetBotToken() *ConnectorAuthenticationBotToken {
	return s.BotToken
}

func (s *ConnectorAuthentication) GetOauth() *ConnectorAuthenticationOauth {
	return s.Oauth
}

func (s *ConnectorAuthentication) GetRole() *ConnectorAuthenticationRole {
	return s.Role
}

func (s *ConnectorAuthentication) GetSatellite() *ConnectorAuthenticationSatellite {
	return s.Satellite
}

func (s *ConnectorAuthentication) GetType() *string {
	return s.Type
}

func (s *ConnectorAuthentication) SetBasic(v *ConnectorAuthenticationBasic) *ConnectorAuthentication {
	s.Basic = v
	return s
}

func (s *ConnectorAuthentication) SetBotToken(v *ConnectorAuthenticationBotToken) *ConnectorAuthentication {
	s.BotToken = v
	return s
}

func (s *ConnectorAuthentication) SetOauth(v *ConnectorAuthenticationOauth) *ConnectorAuthentication {
	s.Oauth = v
	return s
}

func (s *ConnectorAuthentication) SetRole(v *ConnectorAuthenticationRole) *ConnectorAuthentication {
	s.Role = v
	return s
}

func (s *ConnectorAuthentication) SetSatellite(v *ConnectorAuthenticationSatellite) *ConnectorAuthentication {
	s.Satellite = v
	return s
}

func (s *ConnectorAuthentication) SetType(v string) *ConnectorAuthentication {
	s.Type = &v
	return s
}

func (s *ConnectorAuthentication) Validate() error {
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

type ConnectorAuthenticationBasic struct {
	// Username
	//
	// This parameter is required.
	//
	// example:
	//
	// starops
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ConnectorAuthenticationBasic) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthenticationBasic) GoString() string {
	return s.String()
}

func (s *ConnectorAuthenticationBasic) GetUsername() *string {
	return s.Username
}

func (s *ConnectorAuthenticationBasic) SetUsername(v string) *ConnectorAuthenticationBasic {
	s.Username = &v
	return s
}

func (s *ConnectorAuthenticationBasic) Validate() error {
	return dara.Validate(s)
}

type ConnectorAuthenticationBotToken struct {
	// Bot ID
	//
	// This parameter is required.
	//
	// example:
	//
	// bot-123456
	BotId *string `json:"botId,omitempty" xml:"botId,omitempty"`
}

func (s ConnectorAuthenticationBotToken) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthenticationBotToken) GoString() string {
	return s.String()
}

func (s *ConnectorAuthenticationBotToken) GetBotId() *string {
	return s.BotId
}

func (s *ConnectorAuthenticationBotToken) SetBotId(v string) *ConnectorAuthenticationBotToken {
	s.BotId = &v
	return s
}

func (s *ConnectorAuthenticationBotToken) Validate() error {
	return dara.Validate(s)
}

type ConnectorAuthenticationOauth struct {
	// OAuth client ID
	//
	// This parameter is required.
	//
	// example:
	//
	// client-id
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
}

func (s ConnectorAuthenticationOauth) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthenticationOauth) GoString() string {
	return s.String()
}

func (s *ConnectorAuthenticationOauth) GetClientId() *string {
	return s.ClientId
}

func (s *ConnectorAuthenticationOauth) SetClientId(v string) *ConnectorAuthenticationOauth {
	s.ClientId = &v
	return s
}

func (s *ConnectorAuthenticationOauth) Validate() error {
	return dara.Validate(s)
}

type ConnectorAuthenticationRole struct {
	// Role ARN
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::1234567890123456:role/starops-reader
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
}

func (s ConnectorAuthenticationRole) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthenticationRole) GoString() string {
	return s.String()
}

func (s *ConnectorAuthenticationRole) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ConnectorAuthenticationRole) SetRoleArn(v string) *ConnectorAuthenticationRole {
	s.RoleArn = &v
	return s
}

func (s *ConnectorAuthenticationRole) Validate() error {
	return dara.Validate(s)
}

type ConnectorAuthenticationSatellite struct {
	// Local credential binding name
	//
	// This parameter is required.
	//
	// example:
	//
	// private-gitlab
	BindingName *string `json:"bindingName,omitempty" xml:"bindingName,omitempty"`
}

func (s ConnectorAuthenticationSatellite) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthenticationSatellite) GoString() string {
	return s.String()
}

func (s *ConnectorAuthenticationSatellite) GetBindingName() *string {
	return s.BindingName
}

func (s *ConnectorAuthenticationSatellite) SetBindingName(v string) *ConnectorAuthenticationSatellite {
	s.BindingName = &v
	return s
}

func (s *ConnectorAuthenticationSatellite) Validate() error {
	return dara.Validate(s)
}
