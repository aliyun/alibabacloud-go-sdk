// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateDNSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *DeletePrivateDNSRequest
	GetForce() *bool
	SetIntranetDomain(v string) *DeletePrivateDNSRequest
	GetIntranetDomain() *string
	SetSecurityToken(v string) *DeletePrivateDNSRequest
	GetSecurityToken() *string
	SetType(v string) *DeletePrivateDNSRequest
	GetType() *string
}

type DeletePrivateDNSRequest struct {
	// Specifies whether to force delete the resolution.
	//
	// 	- true: force deletes the resolution if the resolution is associated with an instance.
	//
	// 	- false: does not force delete the resolution if the resolution is associated with an instance.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The internal domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// api.demo.com
	IntranetDomain *string `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The internal domain name resolution type. Valid values:
	//
	// 	- VPC: resolution for virtual private cloud (VPC) access authorizations. A resolution of this type can be bound only to traditional dedicated instances.
	//
	// 	- A: resolution that supports A records. A resolution of this type can be bound only to VPC integration dedicated instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeletePrivateDNSRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateDNSRequest) GoString() string {
	return s.String()
}

func (s *DeletePrivateDNSRequest) GetForce() *bool {
	return s.Force
}

func (s *DeletePrivateDNSRequest) GetIntranetDomain() *string {
	return s.IntranetDomain
}

func (s *DeletePrivateDNSRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeletePrivateDNSRequest) GetType() *string {
	return s.Type
}

func (s *DeletePrivateDNSRequest) SetForce(v bool) *DeletePrivateDNSRequest {
	s.Force = &v
	return s
}

func (s *DeletePrivateDNSRequest) SetIntranetDomain(v string) *DeletePrivateDNSRequest {
	s.IntranetDomain = &v
	return s
}

func (s *DeletePrivateDNSRequest) SetSecurityToken(v string) *DeletePrivateDNSRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeletePrivateDNSRequest) SetType(v string) *DeletePrivateDNSRequest {
	s.Type = &v
	return s
}

func (s *DeletePrivateDNSRequest) Validate() error {
	return dara.Validate(s)
}
