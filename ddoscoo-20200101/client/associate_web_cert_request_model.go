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
	Cert   *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	CertId *int32  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The globally unique ID of the certificate. The value is in the "Certificate ID-cn-hangzhou" format. For example, if the ID of the certificate is 123, the value of the CertIdentifier parameter is 123-cn-hangzhou.
	//
	// >  You can specify only one of this parameter and the CertId parameter.
	//
	// example:
	//
	// 9430680-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	CertName       *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The region of the certificate. Valid values: **cn-hangzhou*	- and **ap-southeast-1**. Default value: **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	CertRegion *string `json:"CertRegion,omitempty" xml:"CertRegion,omitempty"`
	// This parameter is required.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Key    *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
