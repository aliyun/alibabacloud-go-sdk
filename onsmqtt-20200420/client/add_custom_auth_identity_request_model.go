// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomAuthIdentityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *AddCustomAuthIdentityRequest
	GetClientId() *string
	SetIdentityType(v string) *AddCustomAuthIdentityRequest
	GetIdentityType() *string
	SetInstanceId(v string) *AddCustomAuthIdentityRequest
	GetInstanceId() *string
	SetSecret(v string) *AddCustomAuthIdentityRequest
	GetSecret() *string
	SetSignMode(v string) *AddCustomAuthIdentityRequest
	GetSignMode() *string
	SetUsername(v string) *AddCustomAuthIdentityRequest
	GetUsername() *string
}

type AddCustomAuthIdentityRequest struct {
	// The client ID of the device. Specify this parameter when the identity type is CLIENT.
	//
	// example:
	//
	// GID_test@@@test
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The identity type.
	//
	// - USER: Username only.
	//
	// - CLIENT: Specific to a client.
	//
	// This parameter is required.
	//
	// example:
	//
	// USER
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mqtt-cn-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxxx
	Secret *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	// The signature verification method. Valid values:
	//
	// - ORIGIN: Original value comparison. The password is compared with the secret.
	//
	// - SIGNED: Signature verification. An HmacSHA1 signature is generated for the clientId by using the secret, and then the signature is compared with the password.
	//
	// This parameter is required.
	//
	// example:
	//
	// SIGNED
	SignMode *string `json:"SignMode,omitempty" xml:"SignMode,omitempty"`
	// The username.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s AddCustomAuthIdentityRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCustomAuthIdentityRequest) GoString() string {
	return s.String()
}

func (s *AddCustomAuthIdentityRequest) GetClientId() *string {
	return s.ClientId
}

func (s *AddCustomAuthIdentityRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *AddCustomAuthIdentityRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddCustomAuthIdentityRequest) GetSecret() *string {
	return s.Secret
}

func (s *AddCustomAuthIdentityRequest) GetSignMode() *string {
	return s.SignMode
}

func (s *AddCustomAuthIdentityRequest) GetUsername() *string {
	return s.Username
}

func (s *AddCustomAuthIdentityRequest) SetClientId(v string) *AddCustomAuthIdentityRequest {
	s.ClientId = &v
	return s
}

func (s *AddCustomAuthIdentityRequest) SetIdentityType(v string) *AddCustomAuthIdentityRequest {
	s.IdentityType = &v
	return s
}

func (s *AddCustomAuthIdentityRequest) SetInstanceId(v string) *AddCustomAuthIdentityRequest {
	s.InstanceId = &v
	return s
}

func (s *AddCustomAuthIdentityRequest) SetSecret(v string) *AddCustomAuthIdentityRequest {
	s.Secret = &v
	return s
}

func (s *AddCustomAuthIdentityRequest) SetSignMode(v string) *AddCustomAuthIdentityRequest {
	s.SignMode = &v
	return s
}

func (s *AddCustomAuthIdentityRequest) SetUsername(v string) *AddCustomAuthIdentityRequest {
	s.Username = &v
	return s
}

func (s *AddCustomAuthIdentityRequest) Validate() error {
	return dara.Validate(s)
}
