// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebAuthnAuthenticatorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticatorId(v string) *DeleteWebAuthnAuthenticatorRequest
	GetAuthenticatorId() *string
	SetInstanceId(v string) *DeleteWebAuthnAuthenticatorRequest
	GetInstanceId() *string
	SetUserId(v string) *DeleteWebAuthnAuthenticatorRequest
	GetUserId() *string
}

type DeleteWebAuthnAuthenticatorRequest struct {
	// 认证器ID
	//
	// This parameter is required.
	//
	// example:
	//
	// authn_h4x5etphqdasydr67lxxxxx
	AuthenticatorId *string `json:"AuthenticatorId,omitempty" xml:"AuthenticatorId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// UserID
	//
	// This parameter is required.
	//
	// example:
	//
	// user_53eh54zdr6iazeijyep5wcxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteWebAuthnAuthenticatorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebAuthnAuthenticatorRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebAuthnAuthenticatorRequest) GetAuthenticatorId() *string {
	return s.AuthenticatorId
}

func (s *DeleteWebAuthnAuthenticatorRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteWebAuthnAuthenticatorRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteWebAuthnAuthenticatorRequest) SetAuthenticatorId(v string) *DeleteWebAuthnAuthenticatorRequest {
	s.AuthenticatorId = &v
	return s
}

func (s *DeleteWebAuthnAuthenticatorRequest) SetInstanceId(v string) *DeleteWebAuthnAuthenticatorRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteWebAuthnAuthenticatorRequest) SetUserId(v string) *DeleteWebAuthnAuthenticatorRequest {
	s.UserId = &v
	return s
}

func (s *DeleteWebAuthnAuthenticatorRequest) Validate() error {
	return dara.Validate(s)
}
