// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertificatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListCertificatesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListCertificatesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListCertificatesResponseBody
	GetRequestId() *string
	SetResult(v []*ListCertificatesResponseBodyResult) *ListCertificatesResponseBody
	GetResult() []*ListCertificatesResponseBodyResult
	SetSiteId(v int64) *ListCertificatesResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *ListCertificatesResponseBody
	GetSiteName() *string
	SetTotalCount(v int64) *ListCertificatesResponseBody
	GetTotalCount() *int64
}

type ListCertificatesResponseBody struct {
	PageNumber *int64                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     []*ListCertificatesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	SiteId     *int64                                `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName   *string                               `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCertificatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCertificatesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCertificatesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCertificatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCertificatesResponseBody) GetResult() []*ListCertificatesResponseBodyResult {
	return s.Result
}

func (s *ListCertificatesResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListCertificatesResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *ListCertificatesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCertificatesResponseBody) SetPageNumber(v int64) *ListCertificatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCertificatesResponseBody) SetPageSize(v int64) *ListCertificatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCertificatesResponseBody) SetRequestId(v string) *ListCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCertificatesResponseBody) SetResult(v []*ListCertificatesResponseBodyResult) *ListCertificatesResponseBody {
	s.Result = v
	return s
}

func (s *ListCertificatesResponseBody) SetSiteId(v int64) *ListCertificatesResponseBody {
	s.SiteId = &v
	return s
}

func (s *ListCertificatesResponseBody) SetSiteName(v string) *ListCertificatesResponseBody {
	s.SiteName = &v
	return s
}

func (s *ListCertificatesResponseBody) SetTotalCount(v int64) *ListCertificatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCertificatesResponseBody) Validate() error {
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

type ListCertificatesResponseBodyResult struct {
	ApplyCode         *int64                                   `json:"ApplyCode,omitempty" xml:"ApplyCode,omitempty"`
	ApplyMessage      *string                                  `json:"ApplyMessage,omitempty" xml:"ApplyMessage,omitempty"`
	CasId             *string                                  `json:"CasId,omitempty" xml:"CasId,omitempty"`
	CommonName        *string                                  `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CreateTime        *string                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DCV               []*ListCertificatesResponseBodyResultDCV `json:"DCV,omitempty" xml:"DCV,omitempty" type:"Repeated"`
	FingerprintSha256 *string                                  `json:"FingerprintSha256,omitempty" xml:"FingerprintSha256,omitempty"`
	Id                *string                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	Issuer            *string                                  `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	IssuerCN          *string                                  `json:"IssuerCN,omitempty" xml:"IssuerCN,omitempty"`
	KeyServerId       *string                                  `json:"KeyServerId,omitempty" xml:"KeyServerId,omitempty"`
	Name              *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	NotAfter          *string                                  `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	NotBefore         *string                                  `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	PubAlg            *string                                  `json:"PubAlg,omitempty" xml:"PubAlg,omitempty"`
	Region            *string                                  `json:"Region,omitempty" xml:"Region,omitempty"`
	SAN               *string                                  `json:"SAN,omitempty" xml:"SAN,omitempty"`
	SerialNumber      *string                                  `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SigAlg            *string                                  `json:"SigAlg,omitempty" xml:"SigAlg,omitempty"`
	Status            *string                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type              *string                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime        *string                                  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListCertificatesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListCertificatesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCertificatesResponseBodyResult) GetApplyCode() *int64 {
	return s.ApplyCode
}

func (s *ListCertificatesResponseBodyResult) GetApplyMessage() *string {
	return s.ApplyMessage
}

func (s *ListCertificatesResponseBodyResult) GetCasId() *string {
	return s.CasId
}

func (s *ListCertificatesResponseBodyResult) GetCommonName() *string {
	return s.CommonName
}

func (s *ListCertificatesResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCertificatesResponseBodyResult) GetDCV() []*ListCertificatesResponseBodyResultDCV {
	return s.DCV
}

func (s *ListCertificatesResponseBodyResult) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *ListCertificatesResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListCertificatesResponseBodyResult) GetIssuer() *string {
	return s.Issuer
}

func (s *ListCertificatesResponseBodyResult) GetIssuerCN() *string {
	return s.IssuerCN
}

func (s *ListCertificatesResponseBodyResult) GetKeyServerId() *string {
	return s.KeyServerId
}

func (s *ListCertificatesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListCertificatesResponseBodyResult) GetNotAfter() *string {
	return s.NotAfter
}

func (s *ListCertificatesResponseBodyResult) GetNotBefore() *string {
	return s.NotBefore
}

func (s *ListCertificatesResponseBodyResult) GetPubAlg() *string {
	return s.PubAlg
}

func (s *ListCertificatesResponseBodyResult) GetRegion() *string {
	return s.Region
}

func (s *ListCertificatesResponseBodyResult) GetSAN() *string {
	return s.SAN
}

func (s *ListCertificatesResponseBodyResult) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ListCertificatesResponseBodyResult) GetSigAlg() *string {
	return s.SigAlg
}

func (s *ListCertificatesResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListCertificatesResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListCertificatesResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListCertificatesResponseBodyResult) SetApplyCode(v int64) *ListCertificatesResponseBodyResult {
	s.ApplyCode = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetApplyMessage(v string) *ListCertificatesResponseBodyResult {
	s.ApplyMessage = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetCasId(v string) *ListCertificatesResponseBodyResult {
	s.CasId = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetCommonName(v string) *ListCertificatesResponseBodyResult {
	s.CommonName = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetCreateTime(v string) *ListCertificatesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetDCV(v []*ListCertificatesResponseBodyResultDCV) *ListCertificatesResponseBodyResult {
	s.DCV = v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetFingerprintSha256(v string) *ListCertificatesResponseBodyResult {
	s.FingerprintSha256 = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetId(v string) *ListCertificatesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetIssuer(v string) *ListCertificatesResponseBodyResult {
	s.Issuer = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetIssuerCN(v string) *ListCertificatesResponseBodyResult {
	s.IssuerCN = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetKeyServerId(v string) *ListCertificatesResponseBodyResult {
	s.KeyServerId = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetName(v string) *ListCertificatesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetNotAfter(v string) *ListCertificatesResponseBodyResult {
	s.NotAfter = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetNotBefore(v string) *ListCertificatesResponseBodyResult {
	s.NotBefore = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetPubAlg(v string) *ListCertificatesResponseBodyResult {
	s.PubAlg = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetRegion(v string) *ListCertificatesResponseBodyResult {
	s.Region = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetSAN(v string) *ListCertificatesResponseBodyResult {
	s.SAN = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetSerialNumber(v string) *ListCertificatesResponseBodyResult {
	s.SerialNumber = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetSigAlg(v string) *ListCertificatesResponseBodyResult {
	s.SigAlg = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetStatus(v string) *ListCertificatesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetType(v string) *ListCertificatesResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) SetUpdateTime(v string) *ListCertificatesResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *ListCertificatesResponseBodyResult) Validate() error {
	if s.DCV != nil {
		for _, item := range s.DCV {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCertificatesResponseBodyResultDCV struct {
	Id     *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Key    *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value  *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListCertificatesResponseBodyResultDCV) String() string {
	return dara.Prettify(s)
}

func (s ListCertificatesResponseBodyResultDCV) GoString() string {
	return s.String()
}

func (s *ListCertificatesResponseBodyResultDCV) GetId() *string {
	return s.Id
}

func (s *ListCertificatesResponseBodyResultDCV) GetKey() *string {
	return s.Key
}

func (s *ListCertificatesResponseBodyResultDCV) GetStatus() *string {
	return s.Status
}

func (s *ListCertificatesResponseBodyResultDCV) GetType() *string {
	return s.Type
}

func (s *ListCertificatesResponseBodyResultDCV) GetValue() *string {
	return s.Value
}

func (s *ListCertificatesResponseBodyResultDCV) SetId(v string) *ListCertificatesResponseBodyResultDCV {
	s.Id = &v
	return s
}

func (s *ListCertificatesResponseBodyResultDCV) SetKey(v string) *ListCertificatesResponseBodyResultDCV {
	s.Key = &v
	return s
}

func (s *ListCertificatesResponseBodyResultDCV) SetStatus(v string) *ListCertificatesResponseBodyResultDCV {
	s.Status = &v
	return s
}

func (s *ListCertificatesResponseBodyResultDCV) SetType(v string) *ListCertificatesResponseBodyResultDCV {
	s.Type = &v
	return s
}

func (s *ListCertificatesResponseBodyResultDCV) SetValue(v string) *ListCertificatesResponseBodyResultDCV {
	s.Value = &v
	return s
}

func (s *ListCertificatesResponseBodyResultDCV) Validate() error {
	return dara.Validate(s)
}
