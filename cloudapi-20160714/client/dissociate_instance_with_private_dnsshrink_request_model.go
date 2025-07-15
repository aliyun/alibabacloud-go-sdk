// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateInstanceWithPrivateDNSShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DissociateInstanceWithPrivateDNSShrinkRequest
	GetInstanceId() *string
	SetIntranetDomainsShrink(v string) *DissociateInstanceWithPrivateDNSShrinkRequest
	GetIntranetDomainsShrink() *string
	SetSecurityToken(v string) *DissociateInstanceWithPrivateDNSShrinkRequest
	GetSecurityToken() *string
}

type DissociateInstanceWithPrivateDNSShrinkRequest struct {
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

func (s DissociateInstanceWithPrivateDNSShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DissociateInstanceWithPrivateDNSShrinkRequest) GoString() string {
	return s.String()
}

func (s *DissociateInstanceWithPrivateDNSShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DissociateInstanceWithPrivateDNSShrinkRequest) GetIntranetDomainsShrink() *string {
	return s.IntranetDomainsShrink
}

func (s *DissociateInstanceWithPrivateDNSShrinkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DissociateInstanceWithPrivateDNSShrinkRequest) SetInstanceId(v string) *DissociateInstanceWithPrivateDNSShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DissociateInstanceWithPrivateDNSShrinkRequest) SetIntranetDomainsShrink(v string) *DissociateInstanceWithPrivateDNSShrinkRequest {
	s.IntranetDomainsShrink = &v
	return s
}

func (s *DissociateInstanceWithPrivateDNSShrinkRequest) SetSecurityToken(v string) *DissociateInstanceWithPrivateDNSShrinkRequest {
	s.SecurityToken = &v
	return s
}

func (s *DissociateInstanceWithPrivateDNSShrinkRequest) Validate() error {
	return dara.Validate(s)
}
