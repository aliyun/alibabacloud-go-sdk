// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomAuthIdentityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *UpdateCustomAuthIdentityRequest
	GetClientId() *string
	SetIdentityType(v string) *UpdateCustomAuthIdentityRequest
	GetIdentityType() *string
	SetInstanceId(v string) *UpdateCustomAuthIdentityRequest
	GetInstanceId() *string
	SetSecret(v string) *UpdateCustomAuthIdentityRequest
	GetSecret() *string
	SetSignMode(v string) *UpdateCustomAuthIdentityRequest
	GetSignMode() *string
	SetUsername(v string) *UpdateCustomAuthIdentityRequest
	GetUsername() *string
}

type UpdateCustomAuthIdentityRequest struct {
	// The client ID of the device. This parameter is required when the identity type is CLIENT.
	//
	// example:
	//
	// GID_test@@@test
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The identity type. Valid values:
	//
	// - USER: Username only.
	//
	// - CLIENT: Specific to the client.
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
	// post-111****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	Secret *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	// The signature verification method. Valid values:
	//
	// - ORIGIN: Original value comparison. The password is directly compared with the secret.
	//
	// - SIGNED: Signature verification. The clientId is signed by using HmacSHA1 and the secret, and then the signature is compared with the password.
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

func (s UpdateCustomAuthIdentityRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAuthIdentityRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomAuthIdentityRequest) GetClientId() *string {
	return s.ClientId
}

func (s *UpdateCustomAuthIdentityRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *UpdateCustomAuthIdentityRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCustomAuthIdentityRequest) GetSecret() *string {
	return s.Secret
}

func (s *UpdateCustomAuthIdentityRequest) GetSignMode() *string {
	return s.SignMode
}

func (s *UpdateCustomAuthIdentityRequest) GetUsername() *string {
	return s.Username
}

func (s *UpdateCustomAuthIdentityRequest) SetClientId(v string) *UpdateCustomAuthIdentityRequest {
	s.ClientId = &v
	return s
}

func (s *UpdateCustomAuthIdentityRequest) SetIdentityType(v string) *UpdateCustomAuthIdentityRequest {
	s.IdentityType = &v
	return s
}

func (s *UpdateCustomAuthIdentityRequest) SetInstanceId(v string) *UpdateCustomAuthIdentityRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCustomAuthIdentityRequest) SetSecret(v string) *UpdateCustomAuthIdentityRequest {
	s.Secret = &v
	return s
}

func (s *UpdateCustomAuthIdentityRequest) SetSignMode(v string) *UpdateCustomAuthIdentityRequest {
	s.SignMode = &v
	return s
}

func (s *UpdateCustomAuthIdentityRequest) SetUsername(v string) *UpdateCustomAuthIdentityRequest {
	s.Username = &v
	return s
}

func (s *UpdateCustomAuthIdentityRequest) Validate() error {
	return dara.Validate(s)
}
