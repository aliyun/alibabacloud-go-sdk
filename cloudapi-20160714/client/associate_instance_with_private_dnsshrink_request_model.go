// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateInstanceWithPrivateDNSShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AssociateInstanceWithPrivateDNSShrinkRequest
	GetInstanceId() *string
	SetIntranetDomainsShrink(v string) *AssociateInstanceWithPrivateDNSShrinkRequest
	GetIntranetDomainsShrink() *string
	SetSecurityToken(v string) *AssociateInstanceWithPrivateDNSShrinkRequest
	GetSecurityToken() *string
}

type AssociateInstanceWithPrivateDNSShrinkRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// apigateway-hz-ead4f4b0bac8
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The internal domain names included in the resolution.
	//
	// This parameter is required.
	IntranetDomainsShrink *string `json:"IntranetDomains,omitempty" xml:"IntranetDomains,omitempty"`
	SecurityToken         *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AssociateInstanceWithPrivateDNSShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateInstanceWithPrivateDNSShrinkRequest) GoString() string {
	return s.String()
}

func (s *AssociateInstanceWithPrivateDNSShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AssociateInstanceWithPrivateDNSShrinkRequest) GetIntranetDomainsShrink() *string {
	return s.IntranetDomainsShrink
}

func (s *AssociateInstanceWithPrivateDNSShrinkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AssociateInstanceWithPrivateDNSShrinkRequest) SetInstanceId(v string) *AssociateInstanceWithPrivateDNSShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateInstanceWithPrivateDNSShrinkRequest) SetIntranetDomainsShrink(v string) *AssociateInstanceWithPrivateDNSShrinkRequest {
	s.IntranetDomainsShrink = &v
	return s
}

func (s *AssociateInstanceWithPrivateDNSShrinkRequest) SetSecurityToken(v string) *AssociateInstanceWithPrivateDNSShrinkRequest {
	s.SecurityToken = &v
	return s
}

func (s *AssociateInstanceWithPrivateDNSShrinkRequest) Validate() error {
	return dara.Validate(s)
}
