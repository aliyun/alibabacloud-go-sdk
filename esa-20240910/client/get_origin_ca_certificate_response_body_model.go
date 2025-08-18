// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginCaCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v string) *GetOriginCaCertificateResponseBody
	GetCertificate() *string
	SetRequestId(v string) *GetOriginCaCertificateResponseBody
	GetRequestId() *string
	SetResult(v *GetOriginCaCertificateResponseBodyResult) *GetOriginCaCertificateResponseBody
	GetResult() *GetOriginCaCertificateResponseBodyResult
	SetSiteId(v int64) *GetOriginCaCertificateResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *GetOriginCaCertificateResponseBody
	GetSiteName() *string
	SetStatus(v string) *GetOriginCaCertificateResponseBody
	GetStatus() *string
}

type GetOriginCaCertificateResponseBody struct {
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
	// 3558df77-8a7a-4060-a900-2d7949403836
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The certificate information.
	Result *GetOriginCaCertificateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The website ID.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The website name.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// The status of the certificate.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetOriginCaCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOriginCaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *GetOriginCaCertificateResponseBody) GetCertificate() *string {
	return s.Certificate
}

func (s *GetOriginCaCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOriginCaCertificateResponseBody) GetResult() *GetOriginCaCertificateResponseBodyResult {
	return s.Result
}

func (s *GetOriginCaCertificateResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetOriginCaCertificateResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *GetOriginCaCertificateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetOriginCaCertificateResponseBody) SetCertificate(v string) *GetOriginCaCertificateResponseBody {
	s.Certificate = &v
	return s
}

func (s *GetOriginCaCertificateResponseBody) SetRequestId(v string) *GetOriginCaCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOriginCaCertificateResponseBody) SetResult(v *GetOriginCaCertificateResponseBodyResult) *GetOriginCaCertificateResponseBody {
	s.Result = v
	return s
}

func (s *GetOriginCaCertificateResponseBody) SetSiteId(v int64) *GetOriginCaCertificateResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetOriginCaCertificateResponseBody) SetSiteName(v string) *GetOriginCaCertificateResponseBody {
	s.SiteName = &v
	return s
}

func (s *GetOriginCaCertificateResponseBody) SetStatus(v string) *GetOriginCaCertificateResponseBody {
	s.Status = &v
	return s
}

func (s *GetOriginCaCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetOriginCaCertificateResponseBodyResult struct {
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
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The SHA-256 fingerprint of the certificate.
	//
	// example:
	//
	// 1dc5fc9af4eead2570c70d94b416130baeb6d4429b51fd3557379588456a****
	FingerprintSha256 *string `json:"FingerprintSha256,omitempty" xml:"FingerprintSha256,omitempty"`
	// The certificate ID.
	//
	// example:
	//
	// babaabcd****
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
	SAN *string `json:"SAN,omitempty" xml:"SAN,omitempty"`
	// The serial number of the certificate.
	//
	// example:
	//
	// babaded901474b9693acf530e0fb1d**
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The signature algorithm of the certificate.
	//
	// example:
	//
	// SHA256-RSA
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitempty" xml:"SignatureAlgorithm,omitempty"`
	// The status of the certificate.
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

func (s GetOriginCaCertificateResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetOriginCaCertificateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOriginCaCertificateResponseBodyResult) GetCommonName() *string {
	return s.CommonName
}

func (s *GetOriginCaCertificateResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetOriginCaCertificateResponseBodyResult) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *GetOriginCaCertificateResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *GetOriginCaCertificateResponseBodyResult) GetIssuer() *string {
	return s.Issuer
}

func (s *GetOriginCaCertificateResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetOriginCaCertificateResponseBodyResult) GetNotAfter() *string {
	return s.NotAfter
}

func (s *GetOriginCaCertificateResponseBodyResult) GetNotBefore() *string {
	return s.NotBefore
}

func (s *GetOriginCaCertificateResponseBodyResult) GetPubkeyAlgorithm() *string {
	return s.PubkeyAlgorithm
}

func (s *GetOriginCaCertificateResponseBodyResult) GetSAN() *string {
	return s.SAN
}

func (s *GetOriginCaCertificateResponseBodyResult) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *GetOriginCaCertificateResponseBodyResult) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *GetOriginCaCertificateResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetOriginCaCertificateResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetOriginCaCertificateResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetOriginCaCertificateResponseBodyResult) SetCommonName(v string) *GetOriginCaCertificateResponseBodyResult {
	s.CommonName = &v
	return s
}

func (s *GetOriginCaCertificateResponseBodyResult) SetCreateTime(v string) *GetOriginCaCertificateResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetOriginCaCertificateResponseBodyResult) SetFingerprintSha256(v string) *GetOriginCaCertificateResponseBodyResult {
	s.FingerprintSha256 = &v
	return s
}

func (s *GetOriginCaCertificateResponseBodyResult) SetId(v string) *GetOriginCaCertificateResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetOriginCaCertificateResponseBodyResult) SetIssuer(v string) *GetOriginCaCertificateResponseBodyResult {
	s.Issuer = &v
	return s
}

func (s *GetOriginCaCertificateResponseBodyResult) SetName(v string) *GetOriginCaCertificateResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetOriginCaCertificateResponseBodyResult) SetNotAfter(v string) *GetOriginCaCertificateResponseBodyResult {
	s.NotAfter = &v
	return s
}

func (s *GetOriginCaCertificateResponseBodyResult) SetNotBefore(v string) *GetOriginCaCertificateResponseBodyResult {
	s.NotBefore = &v
	return s
}

func (s *GetOriginCaCertificateResponseBodyResult) SetPubkeyAlgorithm(v string) *GetOriginCaCertificateResponseBodyResult {
	s.PubkeyAlgorithm = &v
	return s
}

func (s *GetOriginCaCertificateResponseBodyResult) SetSAN(v string) *GetOriginCaCertificateResponseBodyResult {
	s.SAN = &v
	return s
}

func (s *GetOriginCaCertificateResponseBodyResult) SetSerialNumber(v string) *GetOriginCaCertificateResponseBodyResult {
	s.SerialNumber = &v
	return s
}

func (s *GetOriginCaCertificateResponseBodyResult) SetSignatureAlgorithm(v string) *GetOriginCaCertificateResponseBodyResult {
	s.SignatureAlgorithm = &v
	return s
}

func (s *GetOriginCaCertificateResponseBodyResult) SetStatus(v string) *GetOriginCaCertificateResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetOriginCaCertificateResponseBodyResult) SetType(v string) *GetOriginCaCertificateResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetOriginCaCertificateResponseBodyResult) SetUpdateTime(v string) *GetOriginCaCertificateResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *GetOriginCaCertificateResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
