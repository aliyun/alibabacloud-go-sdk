// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindTotpAuthenticatorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UnbindTotpAuthenticatorRequest
	GetInstanceId() *string
	SetUserId(v string) *UnbindTotpAuthenticatorRequest
	GetUserId() *string
}

type UnbindTotpAuthenticatorRequest struct {
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
	// user_1asdfghjmnbcxxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UnbindTotpAuthenticatorRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindTotpAuthenticatorRequest) GoString() string {
	return s.String()
}

func (s *UnbindTotpAuthenticatorRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnbindTotpAuthenticatorRequest) GetUserId() *string {
	return s.UserId
}

func (s *UnbindTotpAuthenticatorRequest) SetInstanceId(v string) *UnbindTotpAuthenticatorRequest {
	s.InstanceId = &v
	return s
}

func (s *UnbindTotpAuthenticatorRequest) SetUserId(v string) *UnbindTotpAuthenticatorRequest {
	s.UserId = &v
	return s
}

func (s *UnbindTotpAuthenticatorRequest) Validate() error {
	return dara.Validate(s)
}
