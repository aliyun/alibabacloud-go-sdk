// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v string) *GetCertificateResponseBody
	GetCertificate() *string
	SetRequestId(v string) *GetCertificateResponseBody
	GetRequestId() *string
	SetResult(v *GetCertificateResponseBodyResult) *GetCertificateResponseBody
	GetResult() *GetCertificateResponseBodyResult
	SetSiteId(v int64) *GetCertificateResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *GetCertificateResponseBody
	GetSiteName() *string
	SetStatus(v string) *GetCertificateResponseBody
	GetStatus() *string
}

type GetCertificateResponseBody struct {
	// Certificate content.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The certificate information.
	Result *GetCertificateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Site ID.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Site name.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// Certificate status.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *GetCertificateResponseBody) GetCertificate() *string {
	return s.Certificate
}

func (s *GetCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCertificateResponseBody) GetResult() *GetCertificateResponseBodyResult {
	return s.Result
}

func (s *GetCertificateResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetCertificateResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *GetCertificateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetCertificateResponseBody) SetCertificate(v string) *GetCertificateResponseBody {
	s.Certificate = &v
	return s
}

func (s *GetCertificateResponseBody) SetRequestId(v string) *GetCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCertificateResponseBody) SetResult(v *GetCertificateResponseBodyResult) *GetCertificateResponseBody {
	s.Result = v
	return s
}

func (s *GetCertificateResponseBody) SetSiteId(v int64) *GetCertificateResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetCertificateResponseBody) SetSiteName(v string) *GetCertificateResponseBody {
	s.SiteName = &v
	return s
}

func (s *GetCertificateResponseBody) SetStatus(v string) *GetCertificateResponseBody {
	s.Status = &v
	return s
}

func (s *GetCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCertificateResponseBodyResult struct {
	// The error code returned for certificate application.
	//
	// example:
	//
	// 2
	ApplyCode *int64 `json:"ApplyCode,omitempty" xml:"ApplyCode,omitempty"`
	// The error message returned for certificate application.
	//
	// example:
	//
	// canceled
	ApplyMessage *string `json:"ApplyMessage,omitempty" xml:"ApplyMessage,omitempty"`
	// Cloud certificate ID.
	//
	// example:
	//
	// 30000478
	CasId *string `json:"CasId,omitempty" xml:"CasId,omitempty"`
	// Common Name (CN) field of the certificate.
	//
	// example:
	//
	// www.example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 2020-05-12 02:00:53
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The Domain Control Validation (DCV) information.
	DCV []*GetCertificateResponseBodyResultDCV `json:"DCV,omitempty" xml:"DCV,omitempty" type:"Repeated"`
	// SHA256 fingerprint of the certificate.
	//
	// example:
	//
	// 1dc5fc9af4eead2570c70d94b416130baeb6d4429b51fd3557379588456aca66
	FingerprintSha256 *string `json:"FingerprintSha256,omitempty" xml:"FingerprintSha256,omitempty"`
	// Certificate ID.
	//
	// example:
	//
	// babaded901474b9693acf530e0fb1d95
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Certificate issuer.
	//
	// example:
	//
	// DigiCert
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// Certificate issuing authority.
	//
	// example:
	//
	// DigiCert Global Root CA
	IssuerCN *string `json:"IssuerCN,omitempty" xml:"IssuerCN,omitempty"`
	// Certificate name.
	//
	// example:
	//
	// yourCertName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// End time of the certificate validity period.
	//
	// example:
	//
	// 2023-11-26T16:00:00Z
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// Start time of the certificate validity period.
	//
	// example:
	//
	// 2023-11-26T16:00:00Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// Certificate public key algorithm.
	//
	// example:
	//
	// ECDSA
	PubAlg *string `json:"PubAlg,omitempty" xml:"PubAlg,omitempty"`
	// Region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Subject Alternative Name (SAN) of the certificate.
	//
	// example:
	//
	// www.example.com,*.example.com
	SAN *string `json:"SAN,omitempty" xml:"SAN,omitempty"`
	// Serial number of the certificate.
	//
	// example:
	//
	// babaded901474b9693acf530e0fb1daa
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// Certificate signature algorithm.
	//
	// example:
	//
	// ECDSA-SHA1
	SigAlg *string `json:"SigAlg,omitempty" xml:"SigAlg,omitempty"`
	// Certificate status.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Certificate type.
	//
	// example:
	//
	// free
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Update time.
	//
	// example:
	//
	// 2022-09-22 05:33:13
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetCertificateResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCertificateResponseBodyResult) GetApplyCode() *int64 {
	return s.ApplyCode
}

func (s *GetCertificateResponseBodyResult) GetApplyMessage() *string {
	return s.ApplyMessage
}

func (s *GetCertificateResponseBodyResult) GetCasId() *string {
	return s.CasId
}

func (s *GetCertificateResponseBodyResult) GetCommonName() *string {
	return s.CommonName
}

func (s *GetCertificateResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetCertificateResponseBodyResult) GetDCV() []*GetCertificateResponseBodyResultDCV {
	return s.DCV
}

func (s *GetCertificateResponseBodyResult) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *GetCertificateResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *GetCertificateResponseBodyResult) GetIssuer() *string {
	return s.Issuer
}

func (s *GetCertificateResponseBodyResult) GetIssuerCN() *string {
	return s.IssuerCN
}

func (s *GetCertificateResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetCertificateResponseBodyResult) GetNotAfter() *string {
	return s.NotAfter
}

func (s *GetCertificateResponseBodyResult) GetNotBefore() *string {
	return s.NotBefore
}

func (s *GetCertificateResponseBodyResult) GetPubAlg() *string {
	return s.PubAlg
}

func (s *GetCertificateResponseBodyResult) GetRegion() *string {
	return s.Region
}

func (s *GetCertificateResponseBodyResult) GetSAN() *string {
	return s.SAN
}

func (s *GetCertificateResponseBodyResult) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *GetCertificateResponseBodyResult) GetSigAlg() *string {
	return s.SigAlg
}

func (s *GetCertificateResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetCertificateResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetCertificateResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetCertificateResponseBodyResult) SetApplyCode(v int64) *GetCertificateResponseBodyResult {
	s.ApplyCode = &v
	return s
}

func (s *GetCertificateResponseBodyResult) SetApplyMessage(v string) *GetCertificateResponseBodyResult {
	s.ApplyMessage = &v
	return s
}

func (s *GetCertificateResponseBodyResult) SetCasId(v string) *GetCertificateResponseBodyResult {
	s.CasId = &v
	return s
}

func (s *GetCertificateResponseBodyResult) SetCommonName(v string) *GetCertificateResponseBodyResult {
	s.CommonName = &v
	return s
}

func (s *GetCertificateResponseBodyResult) SetCreateTime(v string) *GetCertificateResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetCertificateResponseBodyResult) SetDCV(v []*GetCertificateResponseBodyResultDCV) *GetCertificateResponseBodyResult {
	s.DCV = v
	return s
}

func (s *GetCertificateResponseBodyResult) SetFingerprintSha256(v string) *GetCertificateResponseBodyResult {
	s.FingerprintSha256 = &v
	return s
}

func (s *GetCertificateResponseBodyResult) SetId(v string) *GetCertificateResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetCertificateResponseBodyResult) SetIssuer(v string) *GetCertificateResponseBodyResult {
	s.Issuer = &v
	return s
}

func (s *GetCertificateResponseBodyResult) SetIssuerCN(v string) *GetCertificateResponseBodyResult {
	s.IssuerCN = &v
	return s
}

func (s *GetCertificateResponseBodyResult) SetName(v string) *GetCertificateResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetCertificateResponseBodyResult) SetNotAfter(v string) *GetCertificateResponseBodyResult {
	s.NotAfter = &v
	return s
}

func (s *GetCertificateResponseBodyResult) SetNotBefore(v string) *GetCertificateResponseBodyResult {
	s.NotBefore = &v
	return s
}

func (s *GetCertificateResponseBodyResult) SetPubAlg(v string) *GetCertificateResponseBodyResult {
	s.PubAlg = &v
	return s
}

func (s *GetCertificateResponseBodyResult) SetRegion(v string) *GetCertificateResponseBodyResult {
	s.Region = &v
	return s
}

func (s *GetCertificateResponseBodyResult) SetSAN(v string) *GetCertificateResponseBodyResult {
	s.SAN = &v
	return s
}

func (s *GetCertificateResponseBodyResult) SetSerialNumber(v string) *GetCertificateResponseBodyResult {
	s.SerialNumber = &v
	return s
}

func (s *GetCertificateResponseBodyResult) SetSigAlg(v string) *GetCertificateResponseBodyResult {
	s.SigAlg = &v
	return s
}

func (s *GetCertificateResponseBodyResult) SetStatus(v string) *GetCertificateResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetCertificateResponseBodyResult) SetType(v string) *GetCertificateResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetCertificateResponseBodyResult) SetUpdateTime(v string) *GetCertificateResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *GetCertificateResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetCertificateResponseBodyResultDCV struct {
	// The DCV ID.
	//
	// example:
	//
	// bababf7cdd1546a2ad04c0def1f4c980
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DCV name. It is a TXT record name if Type is DNS or URL if Type is HTTP.
	//
	// example:
	//
	// http://www.example.com/.well-known/acme-challenge/pH20CqwS5L3ZnvkhI436DCzadKFuG7QcUcvB_4KsAow
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The verification status.
	//
	// example:
	//
	// pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The DCV type. Valid values: DNS and HTTP.
	//
	// example:
	//
	// HTTP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The DCV content.
	//
	// example:
	//
	// pH20CqwS5L3ZnvkhI436DCzadKFuG7QcUcvB_4KsAow.KfzYo4LH3EgOt7a73G-RqZkbR0eYtLfEUmtmqGmr4FQ
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCertificateResponseBodyResultDCV) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateResponseBodyResultDCV) GoString() string {
	return s.String()
}

func (s *GetCertificateResponseBodyResultDCV) GetId() *string {
	return s.Id
}

func (s *GetCertificateResponseBodyResultDCV) GetKey() *string {
	return s.Key
}

func (s *GetCertificateResponseBodyResultDCV) GetStatus() *string {
	return s.Status
}

func (s *GetCertificateResponseBodyResultDCV) GetType() *string {
	return s.Type
}

func (s *GetCertificateResponseBodyResultDCV) GetValue() *string {
	return s.Value
}

func (s *GetCertificateResponseBodyResultDCV) SetId(v string) *GetCertificateResponseBodyResultDCV {
	s.Id = &v
	return s
}

func (s *GetCertificateResponseBodyResultDCV) SetKey(v string) *GetCertificateResponseBodyResultDCV {
	s.Key = &v
	return s
}

func (s *GetCertificateResponseBodyResultDCV) SetStatus(v string) *GetCertificateResponseBodyResultDCV {
	s.Status = &v
	return s
}

func (s *GetCertificateResponseBodyResultDCV) SetType(v string) *GetCertificateResponseBodyResultDCV {
	s.Type = &v
	return s
}

func (s *GetCertificateResponseBodyResultDCV) SetValue(v string) *GetCertificateResponseBodyResultDCV {
	s.Value = &v
	return s
}

func (s *GetCertificateResponseBodyResultDCV) Validate() error {
	return dara.Validate(s)
}
