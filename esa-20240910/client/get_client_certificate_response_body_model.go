// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v string) *GetClientCertificateResponseBody
	GetCertificate() *string
	SetRequestId(v string) *GetClientCertificateResponseBody
	GetRequestId() *string
	SetResult(v *GetClientCertificateResponseBodyResult) *GetClientCertificateResponseBody
	GetResult() *GetClientCertificateResponseBodyResult
	SetSiteId(v int64) *GetClientCertificateResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *GetClientCertificateResponseBody
	GetSiteName() *string
	SetStatus(v string) *GetClientCertificateResponseBody
	GetStatus() *string
}

type GetClientCertificateResponseBody struct {
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
	Result *GetClientCertificateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetClientCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *GetClientCertificateResponseBody) GetCertificate() *string {
	return s.Certificate
}

func (s *GetClientCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClientCertificateResponseBody) GetResult() *GetClientCertificateResponseBodyResult {
	return s.Result
}

func (s *GetClientCertificateResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetClientCertificateResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *GetClientCertificateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetClientCertificateResponseBody) SetCertificate(v string) *GetClientCertificateResponseBody {
	s.Certificate = &v
	return s
}

func (s *GetClientCertificateResponseBody) SetRequestId(v string) *GetClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClientCertificateResponseBody) SetResult(v *GetClientCertificateResponseBodyResult) *GetClientCertificateResponseBody {
	s.Result = v
	return s
}

func (s *GetClientCertificateResponseBody) SetSiteId(v int64) *GetClientCertificateResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetClientCertificateResponseBody) SetSiteName(v string) *GetClientCertificateResponseBody {
	s.SiteName = &v
	return s
}

func (s *GetClientCertificateResponseBody) SetStatus(v string) *GetClientCertificateResponseBody {
	s.Status = &v
	return s
}

func (s *GetClientCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetClientCertificateResponseBodyResult struct {
	// The ID of the CA certificate.
	//
	// example:
	//
	// babab9db65ee5efcca9f3d41d4b50d66
	CACertificateId *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
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
	// 2024-06-24 07:48:51
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FingerprintSha256 *string `json:"FingerprintSha256,omitempty" xml:"FingerprintSha256,omitempty"`
	// The certificate ID.
	//
	// example:
	//
	// baba39055622c008b90285a8838ed09a
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
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The certificate type.
	//
	// example:
	//
	// dcdn
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the certificate was updated.
	//
	// example:
	//
	// 2024-09-22 05:33:13
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetClientCertificateResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetClientCertificateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetClientCertificateResponseBodyResult) GetCACertificateId() *string {
	return s.CACertificateId
}

func (s *GetClientCertificateResponseBodyResult) GetCommonName() *string {
	return s.CommonName
}

func (s *GetClientCertificateResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetClientCertificateResponseBodyResult) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *GetClientCertificateResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *GetClientCertificateResponseBodyResult) GetIssuer() *string {
	return s.Issuer
}

func (s *GetClientCertificateResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetClientCertificateResponseBodyResult) GetNotAfter() *string {
	return s.NotAfter
}

func (s *GetClientCertificateResponseBodyResult) GetNotBefore() *string {
	return s.NotBefore
}

func (s *GetClientCertificateResponseBodyResult) GetPubkeyAlgorithm() *string {
	return s.PubkeyAlgorithm
}

func (s *GetClientCertificateResponseBodyResult) GetSAN() *string {
	return s.SAN
}

func (s *GetClientCertificateResponseBodyResult) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *GetClientCertificateResponseBodyResult) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *GetClientCertificateResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetClientCertificateResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetClientCertificateResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetClientCertificateResponseBodyResult) SetCACertificateId(v string) *GetClientCertificateResponseBodyResult {
	s.CACertificateId = &v
	return s
}

func (s *GetClientCertificateResponseBodyResult) SetCommonName(v string) *GetClientCertificateResponseBodyResult {
	s.CommonName = &v
	return s
}

func (s *GetClientCertificateResponseBodyResult) SetCreateTime(v string) *GetClientCertificateResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetClientCertificateResponseBodyResult) SetFingerprintSha256(v string) *GetClientCertificateResponseBodyResult {
	s.FingerprintSha256 = &v
	return s
}

func (s *GetClientCertificateResponseBodyResult) SetId(v string) *GetClientCertificateResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetClientCertificateResponseBodyResult) SetIssuer(v string) *GetClientCertificateResponseBodyResult {
	s.Issuer = &v
	return s
}

func (s *GetClientCertificateResponseBodyResult) SetName(v string) *GetClientCertificateResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetClientCertificateResponseBodyResult) SetNotAfter(v string) *GetClientCertificateResponseBodyResult {
	s.NotAfter = &v
	return s
}

func (s *GetClientCertificateResponseBodyResult) SetNotBefore(v string) *GetClientCertificateResponseBodyResult {
	s.NotBefore = &v
	return s
}

func (s *GetClientCertificateResponseBodyResult) SetPubkeyAlgorithm(v string) *GetClientCertificateResponseBodyResult {
	s.PubkeyAlgorithm = &v
	return s
}

func (s *GetClientCertificateResponseBodyResult) SetSAN(v string) *GetClientCertificateResponseBodyResult {
	s.SAN = &v
	return s
}

func (s *GetClientCertificateResponseBodyResult) SetSerialNumber(v string) *GetClientCertificateResponseBodyResult {
	s.SerialNumber = &v
	return s
}

func (s *GetClientCertificateResponseBodyResult) SetSignatureAlgorithm(v string) *GetClientCertificateResponseBodyResult {
	s.SignatureAlgorithm = &v
	return s
}

func (s *GetClientCertificateResponseBodyResult) SetStatus(v string) *GetClientCertificateResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetClientCertificateResponseBodyResult) SetType(v string) *GetClientCertificateResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetClientCertificateResponseBodyResult) SetUpdateTime(v string) *GetClientCertificateResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *GetClientCertificateResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
