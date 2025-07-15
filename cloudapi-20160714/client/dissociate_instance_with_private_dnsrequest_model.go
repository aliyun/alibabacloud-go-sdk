// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateInstanceWithPrivateDNSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DissociateInstanceWithPrivateDNSRequest
	GetInstanceId() *string
	SetIntranetDomains(v []*string) *DissociateInstanceWithPrivateDNSRequest
	GetIntranetDomains() []*string
	SetSecurityToken(v string) *DissociateInstanceWithPrivateDNSRequest
	GetSecurityToken() *string
}

type DissociateInstanceWithPrivateDNSRequest struct {
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

func (s DissociateInstanceWithPrivateDNSRequest) String() string {
	return dara.Prettify(s)
}

func (s DissociateInstanceWithPrivateDNSRequest) GoString() string {
	return s.String()
}

func (s *DissociateInstanceWithPrivateDNSRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DissociateInstanceWithPrivateDNSRequest) GetIntranetDomains() []*string {
	return s.IntranetDomains
}

func (s *DissociateInstanceWithPrivateDNSRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DissociateInstanceWithPrivateDNSRequest) SetInstanceId(v string) *DissociateInstanceWithPrivateDNSRequest {
	s.InstanceId = &v
	return s
}

func (s *DissociateInstanceWithPrivateDNSRequest) SetIntranetDomains(v []*string) *DissociateInstanceWithPrivateDNSRequest {
	s.IntranetDomains = v
	return s
}

func (s *DissociateInstanceWithPrivateDNSRequest) SetSecurityToken(v string) *DissociateInstanceWithPrivateDNSRequest {
	s.SecurityToken = &v
	return s
}

func (s *DissociateInstanceWithPrivateDNSRequest) Validate() error {
	return dara.Validate(s)
}
