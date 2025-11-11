// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindUserAuthnSourceMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProviderId(v string) *BindUserAuthnSourceMappingRequest
	GetIdentityProviderId() *string
	SetInstanceId(v string) *BindUserAuthnSourceMappingRequest
	GetInstanceId() *string
	SetUserExternalId(v string) *BindUserAuthnSourceMappingRequest
	GetUserExternalId() *string
	SetUserId(v string) *BindUserAuthnSourceMappingRequest
	GetUserId() *string
}

type BindUserAuthnSourceMappingRequest struct {
	// 来源Idp Id
	//
	// This parameter is required.
	//
	// example:
	//
	// idp_mwpcwnhrimlr2horxXXXX
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
	// user_fherbdywxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BindUserAuthnSourceMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s BindUserAuthnSourceMappingRequest) GoString() string {
	return s.String()
}

func (s *BindUserAuthnSourceMappingRequest) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *BindUserAuthnSourceMappingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BindUserAuthnSourceMappingRequest) GetUserExternalId() *string {
	return s.UserExternalId
}

func (s *BindUserAuthnSourceMappingRequest) GetUserId() *string {
	return s.UserId
}

func (s *BindUserAuthnSourceMappingRequest) SetIdentityProviderId(v string) *BindUserAuthnSourceMappingRequest {
	s.IdentityProviderId = &v
	return s
}

func (s *BindUserAuthnSourceMappingRequest) SetInstanceId(v string) *BindUserAuthnSourceMappingRequest {
	s.InstanceId = &v
	return s
}

func (s *BindUserAuthnSourceMappingRequest) SetUserExternalId(v string) *BindUserAuthnSourceMappingRequest {
	s.UserExternalId = &v
	return s
}

func (s *BindUserAuthnSourceMappingRequest) SetUserId(v string) *BindUserAuthnSourceMappingRequest {
	s.UserId = &v
	return s
}

func (s *BindUserAuthnSourceMappingRequest) Validate() error {
	return dara.Validate(s)
}
