// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientCertificatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListClientCertificatesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListClientCertificatesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListClientCertificatesResponseBody
	GetRequestId() *string
	SetResult(v []*ListClientCertificatesResponseBodyResult) *ListClientCertificatesResponseBody
	GetResult() []*ListClientCertificatesResponseBodyResult
	SetSiteId(v int64) *ListClientCertificatesResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *ListClientCertificatesResponseBody
	GetSiteName() *string
	SetTotalCount(v int64) *ListClientCertificatesResponseBody
	GetTotalCount() *int64
}

type ListClientCertificatesResponseBody struct {
	PageNumber *int64                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     []*ListClientCertificatesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	SiteId     *int64                                      `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName   *string                                     `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClientCertificatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClientCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClientCertificatesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListClientCertificatesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListClientCertificatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClientCertificatesResponseBody) GetResult() []*ListClientCertificatesResponseBodyResult {
	return s.Result
}

func (s *ListClientCertificatesResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListClientCertificatesResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *ListClientCertificatesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListClientCertificatesResponseBody) SetPageNumber(v int64) *ListClientCertificatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClientCertificatesResponseBody) SetPageSize(v int64) *ListClientCertificatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClientCertificatesResponseBody) SetRequestId(v string) *ListClientCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClientCertificatesResponseBody) SetResult(v []*ListClientCertificatesResponseBodyResult) *ListClientCertificatesResponseBody {
	s.Result = v
	return s
}

func (s *ListClientCertificatesResponseBody) SetSiteId(v int64) *ListClientCertificatesResponseBody {
	s.SiteId = &v
	return s
}

func (s *ListClientCertificatesResponseBody) SetSiteName(v string) *ListClientCertificatesResponseBody {
	s.SiteName = &v
	return s
}

func (s *ListClientCertificatesResponseBody) SetTotalCount(v int64) *ListClientCertificatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListClientCertificatesResponseBody) Validate() error {
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

type ListClientCertificatesResponseBodyResult struct {
	CACertificateId    *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	CommonName         *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FingerprintSha256  *string `json:"FingerprintSha256,omitempty" xml:"FingerprintSha256,omitempty"`
	Id                 *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Issuer             *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NotAfter           *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	NotBefore          *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	PubkeyAlgorithm    *string `json:"PubkeyAlgorithm,omitempty" xml:"PubkeyAlgorithm,omitempty"`
	SAN                *string `json:"SAN,omitempty" xml:"SAN,omitempty"`
	SerialNumber       *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitempty" xml:"SignatureAlgorithm,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type               *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime         *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListClientCertificatesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListClientCertificatesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListClientCertificatesResponseBodyResult) GetCACertificateId() *string {
	return s.CACertificateId
}

func (s *ListClientCertificatesResponseBodyResult) GetCommonName() *string {
	return s.CommonName
}

func (s *ListClientCertificatesResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListClientCertificatesResponseBodyResult) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *ListClientCertificatesResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListClientCertificatesResponseBodyResult) GetIssuer() *string {
	return s.Issuer
}

func (s *ListClientCertificatesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListClientCertificatesResponseBodyResult) GetNotAfter() *string {
	return s.NotAfter
}

func (s *ListClientCertificatesResponseBodyResult) GetNotBefore() *string {
	return s.NotBefore
}

func (s *ListClientCertificatesResponseBodyResult) GetPubkeyAlgorithm() *string {
	return s.PubkeyAlgorithm
}

func (s *ListClientCertificatesResponseBodyResult) GetSAN() *string {
	return s.SAN
}

func (s *ListClientCertificatesResponseBodyResult) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ListClientCertificatesResponseBodyResult) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *ListClientCertificatesResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListClientCertificatesResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListClientCertificatesResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListClientCertificatesResponseBodyResult) SetCACertificateId(v string) *ListClientCertificatesResponseBodyResult {
	s.CACertificateId = &v
	return s
}

func (s *ListClientCertificatesResponseBodyResult) SetCommonName(v string) *ListClientCertificatesResponseBodyResult {
	s.CommonName = &v
	return s
}

func (s *ListClientCertificatesResponseBodyResult) SetCreateTime(v string) *ListClientCertificatesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListClientCertificatesResponseBodyResult) SetFingerprintSha256(v string) *ListClientCertificatesResponseBodyResult {
	s.FingerprintSha256 = &v
	return s
}

func (s *ListClientCertificatesResponseBodyResult) SetId(v string) *ListClientCertificatesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListClientCertificatesResponseBodyResult) SetIssuer(v string) *ListClientCertificatesResponseBodyResult {
	s.Issuer = &v
	return s
}

func (s *ListClientCertificatesResponseBodyResult) SetName(v string) *ListClientCertificatesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListClientCertificatesResponseBodyResult) SetNotAfter(v string) *ListClientCertificatesResponseBodyResult {
	s.NotAfter = &v
	return s
}

func (s *ListClientCertificatesResponseBodyResult) SetNotBefore(v string) *ListClientCertificatesResponseBodyResult {
	s.NotBefore = &v
	return s
}

func (s *ListClientCertificatesResponseBodyResult) SetPubkeyAlgorithm(v string) *ListClientCertificatesResponseBodyResult {
	s.PubkeyAlgorithm = &v
	return s
}

func (s *ListClientCertificatesResponseBodyResult) SetSAN(v string) *ListClientCertificatesResponseBodyResult {
	s.SAN = &v
	return s
}

func (s *ListClientCertificatesResponseBodyResult) SetSerialNumber(v string) *ListClientCertificatesResponseBodyResult {
	s.SerialNumber = &v
	return s
}

func (s *ListClientCertificatesResponseBodyResult) SetSignatureAlgorithm(v string) *ListClientCertificatesResponseBodyResult {
	s.SignatureAlgorithm = &v
	return s
}

func (s *ListClientCertificatesResponseBodyResult) SetStatus(v string) *ListClientCertificatesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListClientCertificatesResponseBodyResult) SetType(v string) *ListClientCertificatesResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListClientCertificatesResponseBodyResult) SetUpdateTime(v string) *ListClientCertificatesResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *ListClientCertificatesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
