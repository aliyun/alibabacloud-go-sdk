// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteOriginClientCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v string) *GetSiteOriginClientCertificateResponseBody
	GetCertificate() *string
	SetRequestId(v string) *GetSiteOriginClientCertificateResponseBody
	GetRequestId() *string
	SetResult(v *GetSiteOriginClientCertificateResponseBodyResult) *GetSiteOriginClientCertificateResponseBody
	GetResult() *GetSiteOriginClientCertificateResponseBodyResult
	SetSiteId(v int64) *GetSiteOriginClientCertificateResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *GetSiteOriginClientCertificateResponseBody
	GetSiteName() *string
	SetStatus(v string) *GetSiteOriginClientCertificateResponseBody
	GetStatus() *string
}

type GetSiteOriginClientCertificateResponseBody struct {
	// The certificate content.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The certificate information.
	Result *GetSiteOriginClientCertificateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
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

func (s GetSiteOriginClientCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSiteOriginClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *GetSiteOriginClientCertificateResponseBody) GetCertificate() *string {
	return s.Certificate
}

func (s *GetSiteOriginClientCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSiteOriginClientCertificateResponseBody) GetResult() *GetSiteOriginClientCertificateResponseBodyResult {
	return s.Result
}

func (s *GetSiteOriginClientCertificateResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetSiteOriginClientCertificateResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *GetSiteOriginClientCertificateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetSiteOriginClientCertificateResponseBody) SetCertificate(v string) *GetSiteOriginClientCertificateResponseBody {
	s.Certificate = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBody) SetRequestId(v string) *GetSiteOriginClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBody) SetResult(v *GetSiteOriginClientCertificateResponseBodyResult) *GetSiteOriginClientCertificateResponseBody {
	s.Result = v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBody) SetSiteId(v int64) *GetSiteOriginClientCertificateResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBody) SetSiteName(v string) *GetSiteOriginClientCertificateResponseBody {
	s.SiteName = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBody) SetStatus(v string) *GetSiteOriginClientCertificateResponseBody {
	s.Status = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSiteOriginClientCertificateResponseBodyResult struct {
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
	// DigiCert
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
	// babab022c5e9b27bf9c64d7f4b16****
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
	// 2024-07-20 06:18:42
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetSiteOriginClientCertificateResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetSiteOriginClientCertificateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetCommonName() *string {
	return s.CommonName
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetIssuer() *string {
	return s.Issuer
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetNotAfter() *string {
	return s.NotAfter
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetNotBefore() *string {
	return s.NotBefore
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetPubkeyAlgorithm() *string {
	return s.PubkeyAlgorithm
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetSAN() *string {
	return s.SAN
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetCommonName(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.CommonName = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetCreateTime(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetFingerprintSha256(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.FingerprintSha256 = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetId(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetIssuer(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.Issuer = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetName(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetNotAfter(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.NotAfter = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetNotBefore(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.NotBefore = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetPubkeyAlgorithm(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.PubkeyAlgorithm = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetSAN(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.SAN = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetSerialNumber(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.SerialNumber = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetSignatureAlgorithm(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.SignatureAlgorithm = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetStatus(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetType(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetUpdateTime(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
