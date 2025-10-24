// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachGatewayDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomDomain(v *DetachGatewayDomainRequestCustomDomain) *DetachGatewayDomainRequest
	GetCustomDomain() *DetachGatewayDomainRequestCustomDomain
}

type DetachGatewayDomainRequest struct {
	// The custom domain name information.
	//
	// This parameter is required.
	CustomDomain *DetachGatewayDomainRequestCustomDomain `json:"CustomDomain,omitempty" xml:"CustomDomain,omitempty" type:"Struct"`
}

func (s DetachGatewayDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachGatewayDomainRequest) GoString() string {
	return s.String()
}

func (s *DetachGatewayDomainRequest) GetCustomDomain() *DetachGatewayDomainRequestCustomDomain {
	return s.CustomDomain
}

func (s *DetachGatewayDomainRequest) SetCustomDomain(v *DetachGatewayDomainRequestCustomDomain) *DetachGatewayDomainRequest {
	s.CustomDomain = v
	return s
}

func (s *DetachGatewayDomainRequest) Validate() error {
	if s.CustomDomain != nil {
		if err := s.CustomDomain.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetachGatewayDomainRequestCustomDomain struct {
	// The custom domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The domain name type.
	//
	// Valid value:
	//
	// 	- intranet: internal network.
	//
	// 	- internet: public network.
	//
	// This parameter is required.
	//
	// example:
	//
	// intranet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DetachGatewayDomainRequestCustomDomain) String() string {
	return dara.Prettify(s)
}

func (s DetachGatewayDomainRequestCustomDomain) GoString() string {
	return s.String()
}

func (s *DetachGatewayDomainRequestCustomDomain) GetDomain() *string {
	return s.Domain
}

func (s *DetachGatewayDomainRequestCustomDomain) GetType() *string {
	return s.Type
}

func (s *DetachGatewayDomainRequestCustomDomain) SetDomain(v string) *DetachGatewayDomainRequestCustomDomain {
	s.Domain = &v
	return s
}

func (s *DetachGatewayDomainRequestCustomDomain) SetType(v string) *DetachGatewayDomainRequestCustomDomain {
	s.Type = &v
	return s
}

func (s *DetachGatewayDomainRequestCustomDomain) Validate() error {
	return dara.Validate(s)
}
