// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOriginClientCertificatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListOriginClientCertificatesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListOriginClientCertificatesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListOriginClientCertificatesResponseBody
	GetRequestId() *string
	SetResult(v []*ListOriginClientCertificatesResponseBodyResult) *ListOriginClientCertificatesResponseBody
	GetResult() []*ListOriginClientCertificatesResponseBodyResult
	SetSiteId(v int64) *ListOriginClientCertificatesResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *ListOriginClientCertificatesResponseBody
	GetSiteName() *string
	SetTotalCount(v int64) *ListOriginClientCertificatesResponseBody
	GetTotalCount() *int64
}

type ListOriginClientCertificatesResponseBody struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3558df77-8a7a-4060-a900-2d7949403836
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The certificate information.
	Result []*ListOriginClientCertificatesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
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
	// The total number of entries.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOriginClientCertificatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOriginClientCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOriginClientCertificatesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListOriginClientCertificatesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListOriginClientCertificatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOriginClientCertificatesResponseBody) GetResult() []*ListOriginClientCertificatesResponseBodyResult {
	return s.Result
}

func (s *ListOriginClientCertificatesResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListOriginClientCertificatesResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *ListOriginClientCertificatesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListOriginClientCertificatesResponseBody) SetPageNumber(v int64) *ListOriginClientCertificatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBody) SetPageSize(v int64) *ListOriginClientCertificatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBody) SetRequestId(v string) *ListOriginClientCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBody) SetResult(v []*ListOriginClientCertificatesResponseBodyResult) *ListOriginClientCertificatesResponseBody {
	s.Result = v
	return s
}

func (s *ListOriginClientCertificatesResponseBody) SetSiteId(v int64) *ListOriginClientCertificatesResponseBody {
	s.SiteId = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBody) SetSiteName(v string) *ListOriginClientCertificatesResponseBody {
	s.SiteName = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBody) SetTotalCount(v int64) *ListOriginClientCertificatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOriginClientCertificatesResponseBodyResult struct {
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
	// The certificate status.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The certificate type.
	//
	// 	- upload: custom certificate that you upload
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

func (s ListOriginClientCertificatesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListOriginClientCertificatesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListOriginClientCertificatesResponseBodyResult) GetCommonName() *string {
	return s.CommonName
}

func (s *ListOriginClientCertificatesResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListOriginClientCertificatesResponseBodyResult) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *ListOriginClientCertificatesResponseBodyResult) GetHostnames() []*string {
	return s.Hostnames
}

func (s *ListOriginClientCertificatesResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListOriginClientCertificatesResponseBodyResult) GetIssuer() *string {
	return s.Issuer
}

func (s *ListOriginClientCertificatesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListOriginClientCertificatesResponseBodyResult) GetNotAfter() *string {
	return s.NotAfter
}

func (s *ListOriginClientCertificatesResponseBodyResult) GetNotBefore() *string {
	return s.NotBefore
}

func (s *ListOriginClientCertificatesResponseBodyResult) GetPubkeyAlgorithm() *string {
	return s.PubkeyAlgorithm
}

func (s *ListOriginClientCertificatesResponseBodyResult) GetSAN() *string {
	return s.SAN
}

func (s *ListOriginClientCertificatesResponseBodyResult) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ListOriginClientCertificatesResponseBodyResult) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *ListOriginClientCertificatesResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListOriginClientCertificatesResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListOriginClientCertificatesResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListOriginClientCertificatesResponseBodyResult) SetCommonName(v string) *ListOriginClientCertificatesResponseBodyResult {
	s.CommonName = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBodyResult) SetCreateTime(v string) *ListOriginClientCertificatesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBodyResult) SetFingerprintSha256(v string) *ListOriginClientCertificatesResponseBodyResult {
	s.FingerprintSha256 = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBodyResult) SetHostnames(v []*string) *ListOriginClientCertificatesResponseBodyResult {
	s.Hostnames = v
	return s
}

func (s *ListOriginClientCertificatesResponseBodyResult) SetId(v string) *ListOriginClientCertificatesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBodyResult) SetIssuer(v string) *ListOriginClientCertificatesResponseBodyResult {
	s.Issuer = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBodyResult) SetName(v string) *ListOriginClientCertificatesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBodyResult) SetNotAfter(v string) *ListOriginClientCertificatesResponseBodyResult {
	s.NotAfter = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBodyResult) SetNotBefore(v string) *ListOriginClientCertificatesResponseBodyResult {
	s.NotBefore = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBodyResult) SetPubkeyAlgorithm(v string) *ListOriginClientCertificatesResponseBodyResult {
	s.PubkeyAlgorithm = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBodyResult) SetSAN(v string) *ListOriginClientCertificatesResponseBodyResult {
	s.SAN = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBodyResult) SetSerialNumber(v string) *ListOriginClientCertificatesResponseBodyResult {
	s.SerialNumber = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBodyResult) SetSignatureAlgorithm(v string) *ListOriginClientCertificatesResponseBodyResult {
	s.SignatureAlgorithm = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBodyResult) SetStatus(v string) *ListOriginClientCertificatesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBodyResult) SetType(v string) *ListOriginClientCertificatesResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBodyResult) SetUpdateTime(v string) *ListOriginClientCertificatesResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *ListOriginClientCertificatesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
