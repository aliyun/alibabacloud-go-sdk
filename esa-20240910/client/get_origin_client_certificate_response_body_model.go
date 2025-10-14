// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginClientCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v string) *GetOriginClientCertificateResponseBody
	GetCertificate() *string
	SetRequestId(v string) *GetOriginClientCertificateResponseBody
	GetRequestId() *string
	SetResult(v *GetOriginClientCertificateResponseBodyResult) *GetOriginClientCertificateResponseBody
	GetResult() *GetOriginClientCertificateResponseBodyResult
	SetSiteId(v int64) *GetOriginClientCertificateResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *GetOriginClientCertificateResponseBody
	GetSiteName() *string
	SetStatus(v string) *GetOriginClientCertificateResponseBody
	GetStatus() *string
}

type GetOriginClientCertificateResponseBody struct {
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
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The certificate information.
	Result *GetOriginClientCertificateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s GetOriginClientCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOriginClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *GetOriginClientCertificateResponseBody) GetCertificate() *string {
	return s.Certificate
}

func (s *GetOriginClientCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOriginClientCertificateResponseBody) GetResult() *GetOriginClientCertificateResponseBodyResult {
	return s.Result
}

func (s *GetOriginClientCertificateResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetOriginClientCertificateResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *GetOriginClientCertificateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetOriginClientCertificateResponseBody) SetCertificate(v string) *GetOriginClientCertificateResponseBody {
	s.Certificate = &v
	return s
}

func (s *GetOriginClientCertificateResponseBody) SetRequestId(v string) *GetOriginClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOriginClientCertificateResponseBody) SetResult(v *GetOriginClientCertificateResponseBodyResult) *GetOriginClientCertificateResponseBody {
	s.Result = v
	return s
}

func (s *GetOriginClientCertificateResponseBody) SetSiteId(v int64) *GetOriginClientCertificateResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetOriginClientCertificateResponseBody) SetSiteName(v string) *GetOriginClientCertificateResponseBody {
	s.SiteName = &v
	return s
}

func (s *GetOriginClientCertificateResponseBody) SetStatus(v string) *GetOriginClientCertificateResponseBody {
	s.Status = &v
	return s
}

func (s *GetOriginClientCertificateResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOriginClientCertificateResponseBodyResult struct {
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
	// 2020-05-12 02:00:53
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The SHA-256 fingerprint of the certificate.
	//
	// example:
	//
	// 1dc5fc9af4eead2570c70d94b416130baeb6d4429b51fd3557379588456a****
	FingerprintSha256 *string `json:"FingerprintSha256,omitempty" xml:"FingerprintSha256,omitempty"`
	// The domain names to associate.
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
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

func (s GetOriginClientCertificateResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetOriginClientCertificateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOriginClientCertificateResponseBodyResult) GetCommonName() *string {
	return s.CommonName
}

func (s *GetOriginClientCertificateResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetOriginClientCertificateResponseBodyResult) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *GetOriginClientCertificateResponseBodyResult) GetHostnames() []*string {
	return s.Hostnames
}

func (s *GetOriginClientCertificateResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *GetOriginClientCertificateResponseBodyResult) GetIssuer() *string {
	return s.Issuer
}

func (s *GetOriginClientCertificateResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetOriginClientCertificateResponseBodyResult) GetNotAfter() *string {
	return s.NotAfter
}

func (s *GetOriginClientCertificateResponseBodyResult) GetNotBefore() *string {
	return s.NotBefore
}

func (s *GetOriginClientCertificateResponseBodyResult) GetPubkeyAlgorithm() *string {
	return s.PubkeyAlgorithm
}

func (s *GetOriginClientCertificateResponseBodyResult) GetSAN() *string {
	return s.SAN
}

func (s *GetOriginClientCertificateResponseBodyResult) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *GetOriginClientCertificateResponseBodyResult) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *GetOriginClientCertificateResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetOriginClientCertificateResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetOriginClientCertificateResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetOriginClientCertificateResponseBodyResult) SetCommonName(v string) *GetOriginClientCertificateResponseBodyResult {
	s.CommonName = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetCreateTime(v string) *GetOriginClientCertificateResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetFingerprintSha256(v string) *GetOriginClientCertificateResponseBodyResult {
	s.FingerprintSha256 = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetHostnames(v []*string) *GetOriginClientCertificateResponseBodyResult {
	s.Hostnames = v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetId(v string) *GetOriginClientCertificateResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetIssuer(v string) *GetOriginClientCertificateResponseBodyResult {
	s.Issuer = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetName(v string) *GetOriginClientCertificateResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetNotAfter(v string) *GetOriginClientCertificateResponseBodyResult {
	s.NotAfter = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetNotBefore(v string) *GetOriginClientCertificateResponseBodyResult {
	s.NotBefore = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetPubkeyAlgorithm(v string) *GetOriginClientCertificateResponseBodyResult {
	s.PubkeyAlgorithm = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetSAN(v string) *GetOriginClientCertificateResponseBodyResult {
	s.SAN = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetSerialNumber(v string) *GetOriginClientCertificateResponseBodyResult {
	s.SerialNumber = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetSignatureAlgorithm(v string) *GetOriginClientCertificateResponseBodyResult {
	s.SignatureAlgorithm = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetStatus(v string) *GetOriginClientCertificateResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetType(v string) *GetOriginClientCertificateResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetUpdateTime(v string) *GetOriginClientCertificateResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
