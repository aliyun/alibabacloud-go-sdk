// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientCaCertificatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListClientCaCertificatesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListClientCaCertificatesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListClientCaCertificatesResponseBody
	GetRequestId() *string
	SetResult(v []*ListClientCaCertificatesResponseBodyResult) *ListClientCaCertificatesResponseBody
	GetResult() []*ListClientCaCertificatesResponseBodyResult
	SetSiteId(v int64) *ListClientCaCertificatesResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *ListClientCaCertificatesResponseBody
	GetSiteName() *string
	SetTotalCount(v int64) *ListClientCaCertificatesResponseBody
	GetTotalCount() *int64
}

type ListClientCaCertificatesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried client CA certificates.
	Result []*ListClientCaCertificatesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
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
	// The total number of entries.
	//
	// example:
	//
	// 16
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClientCaCertificatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClientCaCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClientCaCertificatesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListClientCaCertificatesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListClientCaCertificatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClientCaCertificatesResponseBody) GetResult() []*ListClientCaCertificatesResponseBodyResult {
	return s.Result
}

func (s *ListClientCaCertificatesResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListClientCaCertificatesResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *ListClientCaCertificatesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListClientCaCertificatesResponseBody) SetPageNumber(v int64) *ListClientCaCertificatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClientCaCertificatesResponseBody) SetPageSize(v int64) *ListClientCaCertificatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClientCaCertificatesResponseBody) SetRequestId(v string) *ListClientCaCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClientCaCertificatesResponseBody) SetResult(v []*ListClientCaCertificatesResponseBodyResult) *ListClientCaCertificatesResponseBody {
	s.Result = v
	return s
}

func (s *ListClientCaCertificatesResponseBody) SetSiteId(v int64) *ListClientCaCertificatesResponseBody {
	s.SiteId = &v
	return s
}

func (s *ListClientCaCertificatesResponseBody) SetSiteName(v string) *ListClientCaCertificatesResponseBody {
	s.SiteName = &v
	return s
}

func (s *ListClientCaCertificatesResponseBody) SetTotalCount(v int64) *ListClientCaCertificatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListClientCaCertificatesResponseBody) Validate() error {
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

type ListClientCaCertificatesResponseBodyResult struct {
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
	// babab9db65ee5efcca9f3d41d4b5****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The CA that issued the certificate.
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
	// 2024-07-20 06:18:42
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListClientCaCertificatesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListClientCaCertificatesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListClientCaCertificatesResponseBodyResult) GetCommonName() *string {
	return s.CommonName
}

func (s *ListClientCaCertificatesResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListClientCaCertificatesResponseBodyResult) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *ListClientCaCertificatesResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListClientCaCertificatesResponseBodyResult) GetIssuer() *string {
	return s.Issuer
}

func (s *ListClientCaCertificatesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListClientCaCertificatesResponseBodyResult) GetNotAfter() *string {
	return s.NotAfter
}

func (s *ListClientCaCertificatesResponseBodyResult) GetNotBefore() *string {
	return s.NotBefore
}

func (s *ListClientCaCertificatesResponseBodyResult) GetPubkeyAlgorithm() *string {
	return s.PubkeyAlgorithm
}

func (s *ListClientCaCertificatesResponseBodyResult) GetSAN() *string {
	return s.SAN
}

func (s *ListClientCaCertificatesResponseBodyResult) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ListClientCaCertificatesResponseBodyResult) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *ListClientCaCertificatesResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListClientCaCertificatesResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListClientCaCertificatesResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListClientCaCertificatesResponseBodyResult) SetCommonName(v string) *ListClientCaCertificatesResponseBodyResult {
	s.CommonName = &v
	return s
}

func (s *ListClientCaCertificatesResponseBodyResult) SetCreateTime(v string) *ListClientCaCertificatesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListClientCaCertificatesResponseBodyResult) SetFingerprintSha256(v string) *ListClientCaCertificatesResponseBodyResult {
	s.FingerprintSha256 = &v
	return s
}

func (s *ListClientCaCertificatesResponseBodyResult) SetId(v string) *ListClientCaCertificatesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListClientCaCertificatesResponseBodyResult) SetIssuer(v string) *ListClientCaCertificatesResponseBodyResult {
	s.Issuer = &v
	return s
}

func (s *ListClientCaCertificatesResponseBodyResult) SetName(v string) *ListClientCaCertificatesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListClientCaCertificatesResponseBodyResult) SetNotAfter(v string) *ListClientCaCertificatesResponseBodyResult {
	s.NotAfter = &v
	return s
}

func (s *ListClientCaCertificatesResponseBodyResult) SetNotBefore(v string) *ListClientCaCertificatesResponseBodyResult {
	s.NotBefore = &v
	return s
}

func (s *ListClientCaCertificatesResponseBodyResult) SetPubkeyAlgorithm(v string) *ListClientCaCertificatesResponseBodyResult {
	s.PubkeyAlgorithm = &v
	return s
}

func (s *ListClientCaCertificatesResponseBodyResult) SetSAN(v string) *ListClientCaCertificatesResponseBodyResult {
	s.SAN = &v
	return s
}

func (s *ListClientCaCertificatesResponseBodyResult) SetSerialNumber(v string) *ListClientCaCertificatesResponseBodyResult {
	s.SerialNumber = &v
	return s
}

func (s *ListClientCaCertificatesResponseBodyResult) SetSignatureAlgorithm(v string) *ListClientCaCertificatesResponseBodyResult {
	s.SignatureAlgorithm = &v
	return s
}

func (s *ListClientCaCertificatesResponseBodyResult) SetStatus(v string) *ListClientCaCertificatesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListClientCaCertificatesResponseBodyResult) SetType(v string) *ListClientCaCertificatesResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListClientCaCertificatesResponseBodyResult) SetUpdateTime(v string) *ListClientCaCertificatesResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *ListClientCaCertificatesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
