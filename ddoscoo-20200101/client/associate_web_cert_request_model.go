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
	// The public key of the certificate to associate. This parameter must be used together with CertName and Key.
	//
	// > If you specify CertName, Cert, and Key, you do not need to specify CertId.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- 62EcYPWd2Oy1vs6MTXcJSf*************************************************************************************************************************************************************************ZIo2ZUhl02SIDNggIEeg== -----END CERTIFICATE-----
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// The certificate ID.
	//
	// > If you specify the certificate ID, you do not need to specify CertName, Cert, or Key. This parameter cannot be used together with CertIdentifier.
	//
	// example:
	//
	// 2404693
	CertId *int32 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The global certificate ID, which is in the format of certificate ID + "-cn-hangzhou". For example, if the certificate ID is 123, set CertIdentifier to "123-cn-hangzhou".
	//
	// > This parameter cannot be used together with CertId.
	//
	// example:
	//
	// 9430680-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The certificate name.
	//
	// > The name of the certificate to associate. This parameter must be used together with Cert and Key.
	//
	// example:
	//
	// example-cert
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The region where the certificate is stored. Valid values: **cn-hangzhou*	- and **ap-southeast-1**. Default value: **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	CertRegion *string `json:"CertRegion,omitempty" xml:"CertRegion,omitempty"`
	// The domain name of the website service.
	//
	// > The domain name must have a website forwarding rule configured. You can call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The private key of the certificate to associate. This parameter must be used together with CertName and Cert.
	//
	// > If you specify CertName, Cert, and Key, you do not need to specify CertId.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY----- DADTPZoOHd9WtZ3UKHJT**************************************************************************************************************************************************************************qf0YF8bv5UK5G04RtKadOw== -----END RSA PRIVATE KEY-----
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
