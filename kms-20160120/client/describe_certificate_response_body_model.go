// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *DescribeCertificateResponseBody
	GetArn() *string
	SetCertificateId(v string) *DescribeCertificateResponseBody
	GetCertificateId() *string
	SetCreatedAt(v string) *DescribeCertificateResponseBody
	GetCreatedAt() *string
	SetExportablePrivateKey(v bool) *DescribeCertificateResponseBody
	GetExportablePrivateKey() *bool
	SetIssuer(v string) *DescribeCertificateResponseBody
	GetIssuer() *string
	SetKeySpec(v string) *DescribeCertificateResponseBody
	GetKeySpec() *string
	SetNotAfter(v string) *DescribeCertificateResponseBody
	GetNotAfter() *string
	SetNotBefore(v string) *DescribeCertificateResponseBody
	GetNotBefore() *string
	SetRequestId(v string) *DescribeCertificateResponseBody
	GetRequestId() *string
	SetSerial(v string) *DescribeCertificateResponseBody
	GetSerial() *string
	SetSignatureAlgorithm(v string) *DescribeCertificateResponseBody
	GetSignatureAlgorithm() *string
	SetStatus(v string) *DescribeCertificateResponseBody
	GetStatus() *string
	SetSubject(v string) *DescribeCertificateResponseBody
	GetSubject() *string
	SetSubjectAlternativeNames(v []*string) *DescribeCertificateResponseBody
	GetSubjectAlternativeNames() []*string
	SetSubjectKeyIdentifier(v string) *DescribeCertificateResponseBody
	GetSubjectKeyIdentifier() *string
	SetSubjectPublicKey(v string) *DescribeCertificateResponseBody
	GetSubjectPublicKey() *string
	SetTags(v map[string]interface{}) *DescribeCertificateResponseBody
	GetTags() map[string]interface{}
	SetUpdatedAt(v string) *DescribeCertificateResponseBody
	GetUpdatedAt() *string
}

type DescribeCertificateResponseBody struct {
	// The Alibaba Cloud Resource Name (ARN) of the certificate.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:159498693826****:certificate/9a28de48-8d8b-484d-a766-dec4****"
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the certificate. The ID must be globally unique in Certificates Manager.
	//
	// example:
	//
	// 9a28de48-8d8b-484d-a766-dec4****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The time when the certificate was created.
	//
	// example:
	//
	// 2020-10-13T03:05:03Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// Indicates whether the private key of the certificate can be exported for use. Valid values:
	//
	// 	- true: The private key of the certificate can be exported for use. This is the default value.
	//
	// 	- false: The private key of the certificate cannot be exported for use.
	//
	// example:
	//
	// true
	ExportablePrivateKey *bool `json:"ExportablePrivateKey,omitempty" xml:"ExportablePrivateKey,omitempty"`
	// The certificate issuer in the distinguished name (DN) format.
	//
	// example:
	//
	// CN=testCA,OU=kms,O=aliyun,C=CN
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The type of the key.
	//
	// example:
	//
	// RSA_2048
	KeySpec *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	// The end of the validity period of the certificate.
	//
	// example:
	//
	// 2022-10-13T03:09:00Z
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The beginning of the validity period of the certificate.
	//
	// example:
	//
	// 2020-10-13T03:09:00Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// edb671a3-c5a1-4ebe-a1de-d748b640bdf2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the certificate.
	//
	// example:
	//
	// 12345678
	Serial *string `json:"Serial,omitempty" xml:"Serial,omitempty"`
	// The signature algorithm of the certificate. Valid values:
	//
	// 	- RSA2048-SHA256
	//
	// 	- ECDSA-SHA256
	//
	// 	- SM2-SM3
	//
	// example:
	//
	// ECDSA-SHA256
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitempty" xml:"SignatureAlgorithm,omitempty"`
	// The status of the certificate. Valid values:
	//
	// 	- PENDING: The certificate is to be imported.
	//
	// 	- ACTIVE: The certificate is enabled.
	//
	// 	- INACTIVE: The certificate is disabled.
	//
	// 	- REVOKED: The certificate is revoked.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The subject of the certificate, which is in the DN format.
	//
	// example:
	//
	// CN=userName,OU=aliyun,O=aliyun,C=CN
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The alias of the certificate subject.
	//
	// A domain name list is supported. A maximum of 10 domain names are supported.
	SubjectAlternativeNames []*string `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty" type:"Repeated"`
	// The public key identifier of the certificate subject.
	//
	// example:
	//
	// 79 36 26 DE 9F F5 15 E3 56 DC ****
	SubjectKeyIdentifier *string `json:"SubjectKeyIdentifier,omitempty" xml:"SubjectKeyIdentifier,omitempty"`
	// The public key of the certificate.
	//
	// example:
	//
	// -----BEGIN PUBLIC KEY----- MIIBIjA -----END PUBLIC KEY-----
	SubjectPublicKey *string `json:"SubjectPublicKey,omitempty" xml:"SubjectPublicKey,omitempty"`
	// The tag of the certificate.
	//
	// example:
	//
	// [{\\"TagKey\\":\\"S1key1\\",\\"TagValue\\":\\"S1val1\\"},{\\"TagKey\\":\\"S1key2\\",\\"TagValue\\":\\"S2val2\\"}]
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The time when the certificate was updated.
	//
	// example:
	//
	// 2020-12-23T06:10:13Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s DescribeCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCertificateResponseBody) GetArn() *string {
	return s.Arn
}

func (s *DescribeCertificateResponseBody) GetCertificateId() *string {
	return s.CertificateId
}

func (s *DescribeCertificateResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DescribeCertificateResponseBody) GetExportablePrivateKey() *bool {
	return s.ExportablePrivateKey
}

func (s *DescribeCertificateResponseBody) GetIssuer() *string {
	return s.Issuer
}

func (s *DescribeCertificateResponseBody) GetKeySpec() *string {
	return s.KeySpec
}

func (s *DescribeCertificateResponseBody) GetNotAfter() *string {
	return s.NotAfter
}

func (s *DescribeCertificateResponseBody) GetNotBefore() *string {
	return s.NotBefore
}

func (s *DescribeCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCertificateResponseBody) GetSerial() *string {
	return s.Serial
}

func (s *DescribeCertificateResponseBody) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *DescribeCertificateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeCertificateResponseBody) GetSubject() *string {
	return s.Subject
}

func (s *DescribeCertificateResponseBody) GetSubjectAlternativeNames() []*string {
	return s.SubjectAlternativeNames
}

func (s *DescribeCertificateResponseBody) GetSubjectKeyIdentifier() *string {
	return s.SubjectKeyIdentifier
}

func (s *DescribeCertificateResponseBody) GetSubjectPublicKey() *string {
	return s.SubjectPublicKey
}

func (s *DescribeCertificateResponseBody) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *DescribeCertificateResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *DescribeCertificateResponseBody) SetArn(v string) *DescribeCertificateResponseBody {
	s.Arn = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetCertificateId(v string) *DescribeCertificateResponseBody {
	s.CertificateId = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetCreatedAt(v string) *DescribeCertificateResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetExportablePrivateKey(v bool) *DescribeCertificateResponseBody {
	s.ExportablePrivateKey = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetIssuer(v string) *DescribeCertificateResponseBody {
	s.Issuer = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetKeySpec(v string) *DescribeCertificateResponseBody {
	s.KeySpec = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetNotAfter(v string) *DescribeCertificateResponseBody {
	s.NotAfter = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetNotBefore(v string) *DescribeCertificateResponseBody {
	s.NotBefore = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetRequestId(v string) *DescribeCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetSerial(v string) *DescribeCertificateResponseBody {
	s.Serial = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetSignatureAlgorithm(v string) *DescribeCertificateResponseBody {
	s.SignatureAlgorithm = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetStatus(v string) *DescribeCertificateResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetSubject(v string) *DescribeCertificateResponseBody {
	s.Subject = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetSubjectAlternativeNames(v []*string) *DescribeCertificateResponseBody {
	s.SubjectAlternativeNames = v
	return s
}

func (s *DescribeCertificateResponseBody) SetSubjectKeyIdentifier(v string) *DescribeCertificateResponseBody {
	s.SubjectKeyIdentifier = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetSubjectPublicKey(v string) *DescribeCertificateResponseBody {
	s.SubjectPublicKey = &v
	return s
}

func (s *DescribeCertificateResponseBody) SetTags(v map[string]interface{}) *DescribeCertificateResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeCertificateResponseBody) SetUpdatedAt(v string) *DescribeCertificateResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
