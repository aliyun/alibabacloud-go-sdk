// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateDNSShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIntranetDomain(v string) *UpdatePrivateDNSShrinkRequest
	GetIntranetDomain() *string
	SetRecordsShrink(v string) *UpdatePrivateDNSShrinkRequest
	GetRecordsShrink() *string
	SetSecurityToken(v string) *UpdatePrivateDNSShrinkRequest
	GetSecurityToken() *string
	SetType(v string) *UpdatePrivateDNSShrinkRequest
	GetType() *string
}

type UpdatePrivateDNSShrinkRequest struct {
	// The internal domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// api.demo.com
	IntranetDomain *string `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	// The resolution records. This parameter is valid only when Type is set to A.
	//
	// This parameter is required.
	RecordsShrink *string `json:"Records,omitempty" xml:"Records,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
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

func (s UpdatePrivateDNSShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateDNSShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivateDNSShrinkRequest) GetIntranetDomain() *string {
	return s.IntranetDomain
}

func (s *UpdatePrivateDNSShrinkRequest) GetRecordsShrink() *string {
	return s.RecordsShrink
}

func (s *UpdatePrivateDNSShrinkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpdatePrivateDNSShrinkRequest) GetType() *string {
	return s.Type
}

func (s *UpdatePrivateDNSShrinkRequest) SetIntranetDomain(v string) *UpdatePrivateDNSShrinkRequest {
	s.IntranetDomain = &v
	return s
}

func (s *UpdatePrivateDNSShrinkRequest) SetRecordsShrink(v string) *UpdatePrivateDNSShrinkRequest {
	s.RecordsShrink = &v
	return s
}

func (s *UpdatePrivateDNSShrinkRequest) SetSecurityToken(v string) *UpdatePrivateDNSShrinkRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdatePrivateDNSShrinkRequest) SetType(v string) *UpdatePrivateDNSShrinkRequest {
	s.Type = &v
	return s
}

func (s *UpdatePrivateDNSShrinkRequest) Validate() error {
	return dara.Validate(s)
}
