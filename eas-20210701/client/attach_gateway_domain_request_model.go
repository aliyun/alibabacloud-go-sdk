// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachGatewayDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomDomain(v *AttachGatewayDomainRequestCustomDomain) *AttachGatewayDomainRequest
	GetCustomDomain() *AttachGatewayDomainRequestCustomDomain
}

type AttachGatewayDomainRequest struct {
	// The custom domain name information.
	//
	// This parameter is required.
	CustomDomain *AttachGatewayDomainRequestCustomDomain `json:"CustomDomain,omitempty" xml:"CustomDomain,omitempty" type:"Struct"`
}

func (s AttachGatewayDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachGatewayDomainRequest) GoString() string {
	return s.String()
}

func (s *AttachGatewayDomainRequest) GetCustomDomain() *AttachGatewayDomainRequestCustomDomain {
	return s.CustomDomain
}

func (s *AttachGatewayDomainRequest) SetCustomDomain(v *AttachGatewayDomainRequestCustomDomain) *AttachGatewayDomainRequest {
	s.CustomDomain = v
	return s
}

func (s *AttachGatewayDomainRequest) Validate() error {
	return dara.Validate(s)
}

type AttachGatewayDomainRequestCustomDomain struct {
	// The ID of the SSL certificate bound to the domain name. Obtain the certificate ID after you upload or purchase a certificate in the [Certificate Management Service](https://yundunnext.console.aliyun.com/?spm=5176.2020520163.console-base_help.2.4b3baJixaJixOc\\&p=cas) console.
	//
	// example:
	//
	// 1473**25
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
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

func (s AttachGatewayDomainRequestCustomDomain) String() string {
	return dara.Prettify(s)
}

func (s AttachGatewayDomainRequestCustomDomain) GoString() string {
	return s.String()
}

func (s *AttachGatewayDomainRequestCustomDomain) GetCertificateId() *string {
	return s.CertificateId
}

func (s *AttachGatewayDomainRequestCustomDomain) GetDomain() *string {
	return s.Domain
}

func (s *AttachGatewayDomainRequestCustomDomain) GetType() *string {
	return s.Type
}

func (s *AttachGatewayDomainRequestCustomDomain) SetCertificateId(v string) *AttachGatewayDomainRequestCustomDomain {
	s.CertificateId = &v
	return s
}

func (s *AttachGatewayDomainRequestCustomDomain) SetDomain(v string) *AttachGatewayDomainRequestCustomDomain {
	s.Domain = &v
	return s
}

func (s *AttachGatewayDomainRequestCustomDomain) SetType(v string) *AttachGatewayDomainRequestCustomDomain {
	s.Type = &v
	return s
}

func (s *AttachGatewayDomainRequestCustomDomain) Validate() error {
	return dara.Validate(s)
}
