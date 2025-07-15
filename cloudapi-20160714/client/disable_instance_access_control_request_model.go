// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInstanceAccessControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *DisableInstanceAccessControlRequest
	GetAclId() *string
	SetAddressIPVersion(v string) *DisableInstanceAccessControlRequest
	GetAddressIPVersion() *string
	SetInstanceId(v string) *DisableInstanceAccessControlRequest
	GetInstanceId() *string
	SetSecurityToken(v string) *DisableInstanceAccessControlRequest
	GetSecurityToken() *string
}

type DisableInstanceAccessControlRequest struct {
	// The ID of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// acl-grgqc0mnuhkqciwtam
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// The IP version. Valid values: **ipv4*	- and **ipv6**.
	//
	// example:
	//
	// ipv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// apigateway-cn-v6419k43xxxxx
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DisableInstanceAccessControlRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableInstanceAccessControlRequest) GoString() string {
	return s.String()
}

func (s *DisableInstanceAccessControlRequest) GetAclId() *string {
	return s.AclId
}

func (s *DisableInstanceAccessControlRequest) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *DisableInstanceAccessControlRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableInstanceAccessControlRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DisableInstanceAccessControlRequest) SetAclId(v string) *DisableInstanceAccessControlRequest {
	s.AclId = &v
	return s
}

func (s *DisableInstanceAccessControlRequest) SetAddressIPVersion(v string) *DisableInstanceAccessControlRequest {
	s.AddressIPVersion = &v
	return s
}

func (s *DisableInstanceAccessControlRequest) SetInstanceId(v string) *DisableInstanceAccessControlRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableInstanceAccessControlRequest) SetSecurityToken(v string) *DisableInstanceAccessControlRequest {
	s.SecurityToken = &v
	return s
}

func (s *DisableInstanceAccessControlRequest) Validate() error {
	return dara.Validate(s)
}
