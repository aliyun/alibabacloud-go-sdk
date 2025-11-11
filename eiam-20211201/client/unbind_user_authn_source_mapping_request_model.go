// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindUserAuthnSourceMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderId(v string) *UnbindUserAuthnSourceMappingRequest
	GetIdentityProviderId() *string
	SetInstanceId(v string) *UnbindUserAuthnSourceMappingRequest
	GetInstanceId() *string
	SetUserExternalId(v string) *UnbindUserAuthnSourceMappingRequest
	GetUserExternalId() *string
	SetUserId(v string) *UnbindUserAuthnSourceMappingRequest
	GetUserId() *string
}

type UnbindUserAuthnSourceMappingRequest struct {
	// 来源ID
	//
	// This parameter is required.
	//
	// example:
	//
	// idp_my664lwkhpicbyzirog3xxxxx
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 外部关联ID
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxxxx
	UserExternalId *string `json:"UserExternalId,omitempty" xml:"UserExternalId,omitempty"`
	// 用户ID
	//
	// This parameter is required.
	//
	// example:
	//
	// user_ue2jvisn35exxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UnbindUserAuthnSourceMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindUserAuthnSourceMappingRequest) GoString() string {
	return s.String()
}

func (s *UnbindUserAuthnSourceMappingRequest) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *UnbindUserAuthnSourceMappingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnbindUserAuthnSourceMappingRequest) GetUserExternalId() *string {
	return s.UserExternalId
}

func (s *UnbindUserAuthnSourceMappingRequest) GetUserId() *string {
	return s.UserId
}

func (s *UnbindUserAuthnSourceMappingRequest) SetIdentityProviderId(v string) *UnbindUserAuthnSourceMappingRequest {
	s.IdentityProviderId = &v
	return s
}

func (s *UnbindUserAuthnSourceMappingRequest) SetInstanceId(v string) *UnbindUserAuthnSourceMappingRequest {
	s.InstanceId = &v
	return s
}

func (s *UnbindUserAuthnSourceMappingRequest) SetUserExternalId(v string) *UnbindUserAuthnSourceMappingRequest {
	s.UserExternalId = &v
	return s
}

func (s *UnbindUserAuthnSourceMappingRequest) SetUserId(v string) *UnbindUserAuthnSourceMappingRequest {
	s.UserId = &v
	return s
}

func (s *UnbindUserAuthnSourceMappingRequest) Validate() error {
	return dara.Validate(s)
}
