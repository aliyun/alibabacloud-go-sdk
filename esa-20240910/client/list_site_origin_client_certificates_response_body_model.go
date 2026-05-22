// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSiteOriginClientCertificatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListSiteOriginClientCertificatesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListSiteOriginClientCertificatesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListSiteOriginClientCertificatesResponseBody
	GetRequestId() *string
	SetResult(v []*ListSiteOriginClientCertificatesResponseBodyResult) *ListSiteOriginClientCertificatesResponseBody
	GetResult() []*ListSiteOriginClientCertificatesResponseBodyResult
	SetSiteId(v int64) *ListSiteOriginClientCertificatesResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *ListSiteOriginClientCertificatesResponseBody
	GetSiteName() *string
	SetTotalCount(v int64) *ListSiteOriginClientCertificatesResponseBody
	GetTotalCount() *int64
}

type ListSiteOriginClientCertificatesResponseBody struct {
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
	// 1024
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the certificates.
	Result []*ListSiteOriginClientCertificatesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
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
	// The total number of certificate.
	//
	// example:
	//
	// 90
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSiteOriginClientCertificatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSiteOriginClientCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSiteOriginClientCertificatesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSiteOriginClientCertificatesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSiteOriginClientCertificatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSiteOriginClientCertificatesResponseBody) GetResult() []*ListSiteOriginClientCertificatesResponseBodyResult {
	return s.Result
}

func (s *ListSiteOriginClientCertificatesResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListSiteOriginClientCertificatesResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *ListSiteOriginClientCertificatesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSiteOriginClientCertificatesResponseBody) SetPageNumber(v int64) *ListSiteOriginClientCertificatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBody) SetPageSize(v int64) *ListSiteOriginClientCertificatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBody) SetRequestId(v string) *ListSiteOriginClientCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBody) SetResult(v []*ListSiteOriginClientCertificatesResponseBodyResult) *ListSiteOriginClientCertificatesResponseBody {
	s.Result = v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBody) SetSiteId(v int64) *ListSiteOriginClientCertificatesResponseBody {
	s.SiteId = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBody) SetSiteName(v string) *ListSiteOriginClientCertificatesResponseBody {
	s.SiteName = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBody) SetTotalCount(v int64) *ListSiteOriginClientCertificatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBody) Validate() error {
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

type ListSiteOriginClientCertificatesResponseBodyResult struct {
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
	// The certificate ID on ESA.
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
	// The public key algorithm of the certificate.
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
	// example:
	//
	// upload
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the certificate was updated.
	//
	// example:
	//
	// 2022-09-22 05:33:13
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListSiteOriginClientCertificatesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListSiteOriginClientCertificatesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) GetCommonName() *string {
	return s.CommonName
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) GetIssuer() *string {
	return s.Issuer
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) GetNotAfter() *string {
	return s.NotAfter
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) GetNotBefore() *string {
	return s.NotBefore
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) GetPubkeyAlgorithm() *string {
	return s.PubkeyAlgorithm
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) GetSAN() *string {
	return s.SAN
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) SetCommonName(v string) *ListSiteOriginClientCertificatesResponseBodyResult {
	s.CommonName = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) SetCreateTime(v string) *ListSiteOriginClientCertificatesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) SetFingerprintSha256(v string) *ListSiteOriginClientCertificatesResponseBodyResult {
	s.FingerprintSha256 = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) SetId(v string) *ListSiteOriginClientCertificatesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) SetIssuer(v string) *ListSiteOriginClientCertificatesResponseBodyResult {
	s.Issuer = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) SetName(v string) *ListSiteOriginClientCertificatesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) SetNotAfter(v string) *ListSiteOriginClientCertificatesResponseBodyResult {
	s.NotAfter = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) SetNotBefore(v string) *ListSiteOriginClientCertificatesResponseBodyResult {
	s.NotBefore = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) SetPubkeyAlgorithm(v string) *ListSiteOriginClientCertificatesResponseBodyResult {
	s.PubkeyAlgorithm = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) SetSAN(v string) *ListSiteOriginClientCertificatesResponseBodyResult {
	s.SAN = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) SetSerialNumber(v string) *ListSiteOriginClientCertificatesResponseBodyResult {
	s.SerialNumber = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) SetSignatureAlgorithm(v string) *ListSiteOriginClientCertificatesResponseBodyResult {
	s.SignatureAlgorithm = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) SetStatus(v string) *ListSiteOriginClientCertificatesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) SetType(v string) *ListSiteOriginClientCertificatesResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) SetUpdateTime(v string) *ListSiteOriginClientCertificatesResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *ListSiteOriginClientCertificatesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
