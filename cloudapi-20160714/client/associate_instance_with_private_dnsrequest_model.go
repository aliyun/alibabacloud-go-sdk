// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateInstanceWithPrivateDNSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AssociateInstanceWithPrivateDNSRequest
	GetInstanceId() *string
	SetIntranetDomains(v []*string) *AssociateInstanceWithPrivateDNSRequest
	GetIntranetDomains() []*string
	SetSecurityToken(v string) *AssociateInstanceWithPrivateDNSRequest
	GetSecurityToken() *string
}

type AssociateInstanceWithPrivateDNSRequest struct {
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
	IntranetDomains []*string `json:"IntranetDomains,omitempty" xml:"IntranetDomains,omitempty" type:"Repeated"`
	SecurityToken   *string   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AssociateInstanceWithPrivateDNSRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateInstanceWithPrivateDNSRequest) GoString() string {
	return s.String()
}

func (s *AssociateInstanceWithPrivateDNSRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AssociateInstanceWithPrivateDNSRequest) GetIntranetDomains() []*string {
	return s.IntranetDomains
}

func (s *AssociateInstanceWithPrivateDNSRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AssociateInstanceWithPrivateDNSRequest) SetInstanceId(v string) *AssociateInstanceWithPrivateDNSRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateInstanceWithPrivateDNSRequest) SetIntranetDomains(v []*string) *AssociateInstanceWithPrivateDNSRequest {
	s.IntranetDomains = v
	return s
}

func (s *AssociateInstanceWithPrivateDNSRequest) SetSecurityToken(v string) *AssociateInstanceWithPrivateDNSRequest {
	s.SecurityToken = &v
	return s
}

func (s *AssociateInstanceWithPrivateDNSRequest) Validate() error {
	return dara.Validate(s)
}
