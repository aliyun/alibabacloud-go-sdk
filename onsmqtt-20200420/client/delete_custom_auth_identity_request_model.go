// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomAuthIdentityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *DeleteCustomAuthIdentityRequest
	GetClientId() *string
	SetIdentityType(v string) *DeleteCustomAuthIdentityRequest
	GetIdentityType() *string
	SetInstanceId(v string) *DeleteCustomAuthIdentityRequest
	GetInstanceId() *string
	SetUsername(v string) *DeleteCustomAuthIdentityRequest
	GetUsername() *string
}

type DeleteCustomAuthIdentityRequest struct {
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
	// The ID of the MSMQ for MQTT instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// post-cn-0pp12gl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The username.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DeleteCustomAuthIdentityRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAuthIdentityRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomAuthIdentityRequest) GetClientId() *string {
	return s.ClientId
}

func (s *DeleteCustomAuthIdentityRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *DeleteCustomAuthIdentityRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCustomAuthIdentityRequest) GetUsername() *string {
	return s.Username
}

func (s *DeleteCustomAuthIdentityRequest) SetClientId(v string) *DeleteCustomAuthIdentityRequest {
	s.ClientId = &v
	return s
}

func (s *DeleteCustomAuthIdentityRequest) SetIdentityType(v string) *DeleteCustomAuthIdentityRequest {
	s.IdentityType = &v
	return s
}

func (s *DeleteCustomAuthIdentityRequest) SetInstanceId(v string) *DeleteCustomAuthIdentityRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCustomAuthIdentityRequest) SetUsername(v string) *DeleteCustomAuthIdentityRequest {
	s.Username = &v
	return s
}

func (s *DeleteCustomAuthIdentityRequest) Validate() error {
	return dara.Validate(s)
}
