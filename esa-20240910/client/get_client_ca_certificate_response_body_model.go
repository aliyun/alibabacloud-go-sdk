// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientCaCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v string) *GetClientCaCertificateResponseBody
	GetCertificate() *string
	SetRequestId(v string) *GetClientCaCertificateResponseBody
	GetRequestId() *string
	SetResult(v *GetClientCaCertificateResponseBodyResult) *GetClientCaCertificateResponseBody
	GetResult() *GetClientCaCertificateResponseBodyResult
	SetSiteId(v int64) *GetClientCaCertificateResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *GetClientCaCertificateResponseBody
	GetSiteName() *string
	SetStatus(v string) *GetClientCaCertificateResponseBody
	GetStatus() *string
}

type GetClientCaCertificateResponseBody struct {
	// The certificate content.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The certificate information.
	Result *GetClientCaCertificateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The website ID.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The website name.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// The certificate status.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetClientCaCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClientCaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *GetClientCaCertificateResponseBody) GetCertificate() *string {
	return s.Certificate
}

func (s *GetClientCaCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClientCaCertificateResponseBody) GetResult() *GetClientCaCertificateResponseBodyResult {
	return s.Result
}

func (s *GetClientCaCertificateResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetClientCaCertificateResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *GetClientCaCertificateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetClientCaCertificateResponseBody) SetCertificate(v string) *GetClientCaCertificateResponseBody {
	s.Certificate = &v
	return s
}

func (s *GetClientCaCertificateResponseBody) SetRequestId(v string) *GetClientCaCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClientCaCertificateResponseBody) SetResult(v *GetClientCaCertificateResponseBodyResult) *GetClientCaCertificateResponseBody {
	s.Result = v
	return s
}

func (s *GetClientCaCertificateResponseBody) SetSiteId(v int64) *GetClientCaCertificateResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetClientCaCertificateResponseBody) SetSiteName(v string) *GetClientCaCertificateResponseBody {
	s.SiteName = &v
	return s
}

func (s *GetClientCaCertificateResponseBody) SetStatus(v string) *GetClientCaCertificateResponseBody {
	s.Status = &v
	return s
}

func (s *GetClientCaCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetClientCaCertificateResponseBodyResult struct {
	// The Common Name of the certificate.
	//
	// example:
	//
	// www.example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The time when the certificate was created.
	//
	// example:
	//
	// 2024-03-05 18:24:04
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FingerprintSha256 *string `json:"FingerprintSha256,omitempty" xml:"FingerprintSha256,omitempty"`
	// The certificate ID.
	//
	// example:
	//
	// babab9db65ee5efcca9f3d41d4b5****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The certificate authority (CA) that issued the certificate.
	//
	// example:
	//
	// GlobalSign nv-sa
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The certificate name.
	//
	// example:
	//
	// yourCertName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the certificate expires.
	//
	// example:
	//
	// 2024-03-31 02:08:00
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The time when the certificate takes effect.
	//
	// example:
	//
	// 2023-03-31 02:08:00
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The public-key algorithm of the certificate.
	//
	// example:
	//
	// RSA
	PubkeyAlgorithm *string `json:"PubkeyAlgorithm,omitempty" xml:"PubkeyAlgorithm,omitempty"`
	// The Subject Alternative Name (SAN) of the certificate.
	//
	// example:
	//
	// www.example.com,*.example.com
	SAN          *string `json:"SAN,omitempty" xml:"SAN,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The signature algorithm of the certificate.
	//
	// example:
	//
	// SHA256-RSA
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitempty" xml:"SignatureAlgorithm,omitempty"`
	// The certificate status.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The certificate type.
	//
	// example:
	//
	// upload
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the certificate was updated.
	//
	// example:
	//
	// 2024-03-05 18:24:04
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetClientCaCertificateResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetClientCaCertificateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetClientCaCertificateResponseBodyResult) GetCommonName() *string {
	return s.CommonName
}

func (s *GetClientCaCertificateResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetClientCaCertificateResponseBodyResult) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *GetClientCaCertificateResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *GetClientCaCertificateResponseBodyResult) GetIssuer() *string {
	return s.Issuer
}

func (s *GetClientCaCertificateResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetClientCaCertificateResponseBodyResult) GetNotAfter() *string {
	return s.NotAfter
}

func (s *GetClientCaCertificateResponseBodyResult) GetNotBefore() *string {
	return s.NotBefore
}

func (s *GetClientCaCertificateResponseBodyResult) GetPubkeyAlgorithm() *string {
	return s.PubkeyAlgorithm
}

func (s *GetClientCaCertificateResponseBodyResult) GetSAN() *string {
	return s.SAN
}

func (s *GetClientCaCertificateResponseBodyResult) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *GetClientCaCertificateResponseBodyResult) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *GetClientCaCertificateResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetClientCaCertificateResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetClientCaCertificateResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetClientCaCertificateResponseBodyResult) SetCommonName(v string) *GetClientCaCertificateResponseBodyResult {
	s.CommonName = &v
	return s
}

func (s *GetClientCaCertificateResponseBodyResult) SetCreateTime(v string) *GetClientCaCertificateResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetClientCaCertificateResponseBodyResult) SetFingerprintSha256(v string) *GetClientCaCertificateResponseBodyResult {
	s.FingerprintSha256 = &v
	return s
}

func (s *GetClientCaCertificateResponseBodyResult) SetId(v string) *GetClientCaCertificateResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetClientCaCertificateResponseBodyResult) SetIssuer(v string) *GetClientCaCertificateResponseBodyResult {
	s.Issuer = &v
	return s
}

func (s *GetClientCaCertificateResponseBodyResult) SetName(v string) *GetClientCaCertificateResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetClientCaCertificateResponseBodyResult) SetNotAfter(v string) *GetClientCaCertificateResponseBodyResult {
	s.NotAfter = &v
	return s
}

func (s *GetClientCaCertificateResponseBodyResult) SetNotBefore(v string) *GetClientCaCertificateResponseBodyResult {
	s.NotBefore = &v
	return s
}

func (s *GetClientCaCertificateResponseBodyResult) SetPubkeyAlgorithm(v string) *GetClientCaCertificateResponseBodyResult {
	s.PubkeyAlgorithm = &v
	return s
}

func (s *GetClientCaCertificateResponseBodyResult) SetSAN(v string) *GetClientCaCertificateResponseBodyResult {
	s.SAN = &v
	return s
}

func (s *GetClientCaCertificateResponseBodyResult) SetSerialNumber(v string) *GetClientCaCertificateResponseBodyResult {
	s.SerialNumber = &v
	return s
}

func (s *GetClientCaCertificateResponseBodyResult) SetSignatureAlgorithm(v string) *GetClientCaCertificateResponseBodyResult {
	s.SignatureAlgorithm = &v
	return s
}

func (s *GetClientCaCertificateResponseBodyResult) SetStatus(v string) *GetClientCaCertificateResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetClientCaCertificateResponseBodyResult) SetType(v string) *GetClientCaCertificateResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetClientCaCertificateResponseBodyResult) SetUpdateTime(v string) *GetClientCaCertificateResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *GetClientCaCertificateResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
