// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExportablePrivateKey(v bool) *CreateCertificateRequest
	GetExportablePrivateKey() *bool
	SetKeySpec(v string) *CreateCertificateRequest
	GetKeySpec() *string
	SetSubject(v string) *CreateCertificateRequest
	GetSubject() *string
	SetSubjectAlternativeNames(v map[string]interface{}) *CreateCertificateRequest
	GetSubjectAlternativeNames() map[string]interface{}
}

type CreateCertificateRequest struct {
	// Specifies whether the private key of the certificate can be exported for use. Valid values:
	//
	// 	- true: The private key of the certificate can be exported for use. This is the default value.
	//
	// 	- false: The private key of the certificate cannot be exported for use. We recommend that you set this parameter to false to protect keys with a higher security level.
	//
	// example:
	//
	// true
	ExportablePrivateKey *bool `json:"ExportablePrivateKey,omitempty" xml:"ExportablePrivateKey,omitempty"`
	// The type of the key. Valid values:
	//
	// 	- RSA_2048
	//
	// 	- EC_P256
	//
	// 	- EC_SM2
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA_2048
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The certificate subject, which is the owner of the certificate.
	//
	// Specify the value in the distinguished name (DN) format, as defined in [RFC 2253](https://tools.ietf.org/html/rfc2253?spm=a2c4g.11186623.2.13.265f1a1cGFCn3Q). A DN is a sequence of relative distinguished names (RDNs).
	//
	// RDNs are key-value pairs in the format of `attribute1=value1,attribute2=value2`. Separate multiple RDNs with commas (,).
	//
	// The Subject parameter consists of the following fields:
	//
	// 	- CN: required. The name of the certificate subject.
	//
	// 	- C: required. The two-character country or region code in the [ISO 3166-1](https://www.iso.org/obp/ui/#search/code/) standard. For example, CN indicates China.
	//
	// 	- O: required. The legal name of the enterprise, company, organization, or institution.
	//
	// 	- OU: required. The name of the department.
	//
	// 	- ST: optional. The name of the province, municipality, autonomous region, or special administrative region.
	//
	// 	- L: optional. The name of the city.
	//
	// This parameter is required.
	//
	// example:
	//
	// CN=userName,OU=kms,O=aliyun,C=CN
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The subject alternative names.
	//
	// A domain name list is supported. A maximum of 10 domain names are supported.
	//
	// example:
	//
	// ["test1.example.com","test2.example.com"]
	SubjectAlternativeNames map[string]interface{} `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty"`
}

func (s CreateCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateCertificateRequest) GetExportablePrivateKey() *bool {
	return s.ExportablePrivateKey
}

func (s *CreateCertificateRequest) GetKeySpec() *string {
	return s.KeySpec
}

func (s *CreateCertificateRequest) GetSubject() *string {
	return s.Subject
}

func (s *CreateCertificateRequest) GetSubjectAlternativeNames() map[string]interface{} {
	return s.SubjectAlternativeNames
}

func (s *CreateCertificateRequest) SetExportablePrivateKey(v bool) *CreateCertificateRequest {
	s.ExportablePrivateKey = &v
	return s
}

func (s *CreateCertificateRequest) SetKeySpec(v string) *CreateCertificateRequest {
	s.KeySpec = &v
	return s
}

func (s *CreateCertificateRequest) SetSubject(v string) *CreateCertificateRequest {
	s.Subject = &v
	return s
}

func (s *CreateCertificateRequest) SetSubjectAlternativeNames(v map[string]interface{}) *CreateCertificateRequest {
	s.SubjectAlternativeNames = v
	return s
}

func (s *CreateCertificateRequest) Validate() error {
	return dara.Validate(s)
}
