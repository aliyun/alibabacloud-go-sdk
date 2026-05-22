// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOriginCaCertificatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListOriginCaCertificatesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListOriginCaCertificatesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListOriginCaCertificatesResponseBody
	GetRequestId() *string
	SetResult(v []*ListOriginCaCertificatesResponseBodyResult) *ListOriginCaCertificatesResponseBody
	GetResult() []*ListOriginCaCertificatesResponseBodyResult
	SetSiteId(v int64) *ListOriginCaCertificatesResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *ListOriginCaCertificatesResponseBody
	GetSiteName() *string
	SetTotalCount(v int64) *ListOriginCaCertificatesResponseBody
	GetTotalCount() *int64
}

type ListOriginCaCertificatesResponseBody struct {
	// Page number, default is 1 if not provided.
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
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the certificates.
	Result []*ListOriginCaCertificatesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
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

func (s ListOriginCaCertificatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOriginCaCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOriginCaCertificatesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListOriginCaCertificatesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListOriginCaCertificatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOriginCaCertificatesResponseBody) GetResult() []*ListOriginCaCertificatesResponseBodyResult {
	return s.Result
}

func (s *ListOriginCaCertificatesResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListOriginCaCertificatesResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *ListOriginCaCertificatesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListOriginCaCertificatesResponseBody) SetPageNumber(v int64) *ListOriginCaCertificatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBody) SetPageSize(v int64) *ListOriginCaCertificatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBody) SetRequestId(v string) *ListOriginCaCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBody) SetResult(v []*ListOriginCaCertificatesResponseBodyResult) *ListOriginCaCertificatesResponseBody {
	s.Result = v
	return s
}

func (s *ListOriginCaCertificatesResponseBody) SetSiteId(v int64) *ListOriginCaCertificatesResponseBody {
	s.SiteId = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBody) SetSiteName(v string) *ListOriginCaCertificatesResponseBody {
	s.SiteName = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBody) SetTotalCount(v int64) *ListOriginCaCertificatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOriginCaCertificatesResponseBodyResult struct {
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
	// 2023-11-26T16:00:00Z
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

func (s ListOriginCaCertificatesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListOriginCaCertificatesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListOriginCaCertificatesResponseBodyResult) GetCommonName() *string {
	return s.CommonName
}

func (s *ListOriginCaCertificatesResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListOriginCaCertificatesResponseBodyResult) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *ListOriginCaCertificatesResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListOriginCaCertificatesResponseBodyResult) GetIssuer() *string {
	return s.Issuer
}

func (s *ListOriginCaCertificatesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListOriginCaCertificatesResponseBodyResult) GetNotAfter() *string {
	return s.NotAfter
}

func (s *ListOriginCaCertificatesResponseBodyResult) GetNotBefore() *string {
	return s.NotBefore
}

func (s *ListOriginCaCertificatesResponseBodyResult) GetPubkeyAlgorithm() *string {
	return s.PubkeyAlgorithm
}

func (s *ListOriginCaCertificatesResponseBodyResult) GetSAN() *string {
	return s.SAN
}

func (s *ListOriginCaCertificatesResponseBodyResult) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ListOriginCaCertificatesResponseBodyResult) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *ListOriginCaCertificatesResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListOriginCaCertificatesResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListOriginCaCertificatesResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListOriginCaCertificatesResponseBodyResult) SetCommonName(v string) *ListOriginCaCertificatesResponseBodyResult {
	s.CommonName = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBodyResult) SetCreateTime(v string) *ListOriginCaCertificatesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBodyResult) SetFingerprintSha256(v string) *ListOriginCaCertificatesResponseBodyResult {
	s.FingerprintSha256 = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBodyResult) SetId(v string) *ListOriginCaCertificatesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBodyResult) SetIssuer(v string) *ListOriginCaCertificatesResponseBodyResult {
	s.Issuer = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBodyResult) SetName(v string) *ListOriginCaCertificatesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBodyResult) SetNotAfter(v string) *ListOriginCaCertificatesResponseBodyResult {
	s.NotAfter = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBodyResult) SetNotBefore(v string) *ListOriginCaCertificatesResponseBodyResult {
	s.NotBefore = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBodyResult) SetPubkeyAlgorithm(v string) *ListOriginCaCertificatesResponseBodyResult {
	s.PubkeyAlgorithm = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBodyResult) SetSAN(v string) *ListOriginCaCertificatesResponseBodyResult {
	s.SAN = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBodyResult) SetSerialNumber(v string) *ListOriginCaCertificatesResponseBodyResult {
	s.SerialNumber = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBodyResult) SetSignatureAlgorithm(v string) *ListOriginCaCertificatesResponseBodyResult {
	s.SignatureAlgorithm = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBodyResult) SetStatus(v string) *ListOriginCaCertificatesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBodyResult) SetType(v string) *ListOriginCaCertificatesResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBodyResult) SetUpdateTime(v string) *ListOriginCaCertificatesResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *ListOriginCaCertificatesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
