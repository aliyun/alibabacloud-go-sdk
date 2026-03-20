// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateWebCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCert(v string) *AssociateWebCertRequest
	GetCert() *string
	SetCertId(v int32) *AssociateWebCertRequest
	GetCertId() *int32
	SetCertIdentifier(v string) *AssociateWebCertRequest
	GetCertIdentifier() *string
	SetCertName(v string) *AssociateWebCertRequest
	GetCertName() *string
	SetCertRegion(v string) *AssociateWebCertRequest
	GetCertRegion() *string
	SetDomain(v string) *AssociateWebCertRequest
	GetDomain() *string
	SetKey(v string) *AssociateWebCertRequest
	GetKey() *string
}

type AssociateWebCertRequest struct {
	// The public key of the certificate that you want to associate. This parameter must be used together with the CertName and Key parameters.
	//
	// >  If you specify a value for the CertName, Cert, and Key parameters, you do not need to specify a value for the CertId parameter.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- 62EcYPWd2Oy1vs6MTXcJSfN9Z7rZ9fmxWr2BFN2XbahgnsSXM48ixZJ4krc+1M+j2kcubVpsE2cgHdj4v8H6jUz9Ji4mr7vMNS6dXv8PUkl/qoDeNGCNdyTS5NIL5ir+g92cL8IGOkjgvhlqt9vc65Cgb4mL+n5+DV9uOyTZTW/MojmlgfUekC2xiXa54nxJf17Y1TADGSbyJbsC0Q9nIrHsPl8YKkvRWvIAqYxXZ7wRwWWmv4TMxFhWRiNY7yZIo2ZUhl02SIDNggIEeg== -----END CERTIFICATE-----
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// The ID of the certificate.
	//
	// >  If you specify a certificate ID, you do not need to specify a value for the CertName, Cert, and Key parameters. You can specify only one of this parameter and the CertIdentifier parameter.
	//
	// example:
	//
	// 2404693
	CertId *int32 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The globally unique ID of the certificate. The value is in the "Certificate ID-cn-hangzhou" format. For example, if the ID of the certificate is 123, the value of the CertIdentifier parameter is 123-cn-hangzhou.
	//
	// >  You can specify only one of this parameter and the CertId parameter.
	//
	// example:
	//
	// 9430680-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The name of the certificate.
	//
	// >  You can specify the name of the certificate that you want to associate. This parameter must be used together with the Cert and Key parameters.
	//
	// example:
	//
	// example-cert
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The region of the certificate. Valid values: **cn-hangzhou*	- and **ap-southeast-1**. Default value: **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	CertRegion *string `json:"CertRegion,omitempty" xml:"CertRegion,omitempty"`
	// The domain name of the website.
	//
	// >  A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The private key of the certificate that you want to associate. This parameter must be used together with the CertName and Cert parameters.
	//
	// >  If you specify a value for the CertName, Cert, and Key parameters, you do not need to specify a value for the CertId parameter.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY----- DADTPZoOHd9WtZ3UKHJTRgNQmioPQn2bqdKHop+B/dn/4VZL7Jt8zSDGM9sTMThLyvsmLQKBgQCr+ujntC1kN6pGBj2Fw2l/EA/W3rYEce2tyhjgmG7rZ+A/jVE9fld5sQra6ZdwBcQJaiygoIYoaMF2EjRwc0qwHaluq0C15f6ujSoHh2e+D5zdmkTg/3NKNjqNv6xA2gYpinVDzFdZ9Zujxvuh9o4Vqf0YF8bv5UK5G04RtKadOw== -----END RSA PRIVATE KEY-----
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s AssociateWebCertRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateWebCertRequest) GoString() string {
	return s.String()
}

func (s *AssociateWebCertRequest) GetCert() *string {
	return s.Cert
}

func (s *AssociateWebCertRequest) GetCertId() *int32 {
	return s.CertId
}

func (s *AssociateWebCertRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *AssociateWebCertRequest) GetCertName() *string {
	return s.CertName
}

func (s *AssociateWebCertRequest) GetCertRegion() *string {
	return s.CertRegion
}

func (s *AssociateWebCertRequest) GetDomain() *string {
	return s.Domain
}

func (s *AssociateWebCertRequest) GetKey() *string {
	return s.Key
}

func (s *AssociateWebCertRequest) SetCert(v string) *AssociateWebCertRequest {
	s.Cert = &v
	return s
}

func (s *AssociateWebCertRequest) SetCertId(v int32) *AssociateWebCertRequest {
	s.CertId = &v
	return s
}

func (s *AssociateWebCertRequest) SetCertIdentifier(v string) *AssociateWebCertRequest {
	s.CertIdentifier = &v
	return s
}

func (s *AssociateWebCertRequest) SetCertName(v string) *AssociateWebCertRequest {
	s.CertName = &v
	return s
}

func (s *AssociateWebCertRequest) SetCertRegion(v string) *AssociateWebCertRequest {
	s.CertRegion = &v
	return s
}

func (s *AssociateWebCertRequest) SetDomain(v string) *AssociateWebCertRequest {
	s.Domain = &v
	return s
}

func (s *AssociateWebCertRequest) SetKey(v string) *AssociateWebCertRequest {
	s.Key = &v
	return s
}

func (s *AssociateWebCertRequest) Validate() error {
	return dara.Validate(s)
}
