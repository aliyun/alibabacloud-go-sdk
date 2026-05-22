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
	PageNumber *int64                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     []*ListOriginClientCertificatesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	SiteId     *int64                                            `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName   *string                                           `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	TotalCount *int64                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

type ListOriginClientCertificatesResponseBodyResult struct {
	CommonName         *string   `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CreateTime         *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FingerprintSha256  *string   `json:"FingerprintSha256,omitempty" xml:"FingerprintSha256,omitempty"`
	Hostnames          []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	Id                 *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	Issuer             *string   `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	Name               *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NotAfter           *string   `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	NotBefore          *string   `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	PubkeyAlgorithm    *string   `json:"PubkeyAlgorithm,omitempty" xml:"PubkeyAlgorithm,omitempty"`
	SAN                *string   `json:"SAN,omitempty" xml:"SAN,omitempty"`
	SerialNumber       *string   `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SignatureAlgorithm *string   `json:"SignatureAlgorithm,omitempty" xml:"SignatureAlgorithm,omitempty"`
	Status             *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	Type               *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime         *string   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
