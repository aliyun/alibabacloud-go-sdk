// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivateDNSShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIntranetDomain(v string) *CreatePrivateDNSShrinkRequest
	GetIntranetDomain() *string
	SetRecordsShrink(v string) *CreatePrivateDNSShrinkRequest
	GetRecordsShrink() *string
	SetSecurityToken(v string) *CreatePrivateDNSShrinkRequest
	GetSecurityToken() *string
	SetType(v string) *CreatePrivateDNSShrinkRequest
	GetType() *string
}

type CreatePrivateDNSShrinkRequest struct {
	// The internal domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// api.demo.com
	IntranetDomain *string `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	// The resolution records. This parameter is valid only when Type is set to A.
	RecordsShrink *string `json:"Records,omitempty" xml:"Records,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The internal domain name resolution type. Valid values:
	//
	// 	- VPC: resolution for VPC access authorizations. A resolution of this type can be bound only to traditional dedicated instances.
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

func (s CreatePrivateDNSShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateDNSShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePrivateDNSShrinkRequest) GetIntranetDomain() *string {
	return s.IntranetDomain
}

func (s *CreatePrivateDNSShrinkRequest) GetRecordsShrink() *string {
	return s.RecordsShrink
}

func (s *CreatePrivateDNSShrinkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreatePrivateDNSShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreatePrivateDNSShrinkRequest) SetIntranetDomain(v string) *CreatePrivateDNSShrinkRequest {
	s.IntranetDomain = &v
	return s
}

func (s *CreatePrivateDNSShrinkRequest) SetRecordsShrink(v string) *CreatePrivateDNSShrinkRequest {
	s.RecordsShrink = &v
	return s
}

func (s *CreatePrivateDNSShrinkRequest) SetSecurityToken(v string) *CreatePrivateDNSShrinkRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreatePrivateDNSShrinkRequest) SetType(v string) *CreatePrivateDNSShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreatePrivateDNSShrinkRequest) Validate() error {
	return dara.Validate(s)
}
