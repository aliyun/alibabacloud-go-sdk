// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertificatesByRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListCertificatesByRecordResponseBody
	GetRequestId() *string
	SetResult(v []*ListCertificatesByRecordResponseBodyResult) *ListCertificatesByRecordResponseBody
	GetResult() []*ListCertificatesByRecordResponseBodyResult
	SetSiteId(v int64) *ListCertificatesByRecordResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *ListCertificatesByRecordResponseBody
	GetSiteName() *string
	SetTotalCount(v int64) *ListCertificatesByRecordResponseBody
	GetTotalCount() *int64
}

type ListCertificatesByRecordResponseBody struct {
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result     []*ListCertificatesByRecordResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	SiteId     *int64                                        `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName   *string                                       `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	TotalCount *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCertificatesByRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCertificatesByRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListCertificatesByRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCertificatesByRecordResponseBody) GetResult() []*ListCertificatesByRecordResponseBodyResult {
	return s.Result
}

func (s *ListCertificatesByRecordResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListCertificatesByRecordResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *ListCertificatesByRecordResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCertificatesByRecordResponseBody) SetRequestId(v string) *ListCertificatesByRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCertificatesByRecordResponseBody) SetResult(v []*ListCertificatesByRecordResponseBodyResult) *ListCertificatesByRecordResponseBody {
	s.Result = v
	return s
}

func (s *ListCertificatesByRecordResponseBody) SetSiteId(v int64) *ListCertificatesByRecordResponseBody {
	s.SiteId = &v
	return s
}

func (s *ListCertificatesByRecordResponseBody) SetSiteName(v string) *ListCertificatesByRecordResponseBody {
	s.SiteName = &v
	return s
}

func (s *ListCertificatesByRecordResponseBody) SetTotalCount(v int64) *ListCertificatesByRecordResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCertificatesByRecordResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCertificatesByRecordResponseBodyResult struct {
	ApplylingCount *int64                                                    `json:"ApplylingCount,omitempty" xml:"ApplylingCount,omitempty"`
	Certificates   []*ListCertificatesByRecordResponseBodyResultCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	Count          *int64                                                    `json:"Count,omitempty" xml:"Count,omitempty"`
	RecordName     *string                                                   `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	Status         *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCertificatesByRecordResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListCertificatesByRecordResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCertificatesByRecordResponseBodyResult) GetApplylingCount() *int64 {
	return s.ApplylingCount
}

func (s *ListCertificatesByRecordResponseBodyResult) GetCertificates() []*ListCertificatesByRecordResponseBodyResultCertificates {
	return s.Certificates
}

func (s *ListCertificatesByRecordResponseBodyResult) GetCount() *int64 {
	return s.Count
}

func (s *ListCertificatesByRecordResponseBodyResult) GetRecordName() *string {
	return s.RecordName
}

func (s *ListCertificatesByRecordResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListCertificatesByRecordResponseBodyResult) SetApplylingCount(v int64) *ListCertificatesByRecordResponseBodyResult {
	s.ApplylingCount = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResult) SetCertificates(v []*ListCertificatesByRecordResponseBodyResultCertificates) *ListCertificatesByRecordResponseBodyResult {
	s.Certificates = v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResult) SetCount(v int64) *ListCertificatesByRecordResponseBodyResult {
	s.Count = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResult) SetRecordName(v string) *ListCertificatesByRecordResponseBodyResult {
	s.RecordName = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResult) SetStatus(v string) *ListCertificatesByRecordResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListCertificatesByRecordResponseBodyResultCertificates struct {
	// example:
	//
	// 30000137
	CasId *string `json:"CasId,omitempty" xml:"CasId,omitempty"`
	// example:
	//
	// www.example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// example:
	//
	// 2023-02-28 06:17:11
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1dc5fc9af4eead2570c70d94b416130baeb6d4429b51fd3557379588456aca66
	FingerprintSha256 *string `json:"FingerprintSha256,omitempty" xml:"FingerprintSha256,omitempty"`
	Id                *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// GlobalSign nv-sa
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// example:
	//
	// GlobalSign Organization Validation CA - SHA256 - G3
	IssuerCN *string `json:"IssuerCN,omitempty" xml:"IssuerCN,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2024-02-28 06:17:11
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// example:
	//
	// 2023-02-28 06:17:11
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// example:
	//
	// RSA
	PubAlg *string `json:"PubAlg,omitempty" xml:"PubAlg,omitempty"`
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// www.example.com,*.example.com
	SAN *string `json:"SAN,omitempty" xml:"SAN,omitempty"`
	// example:
	//
	// baba39055622c008b90285a8838ed09a
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// example:
	//
	// SHA256-RSA
	SigAlg *string `json:"SigAlg,omitempty" xml:"SigAlg,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// free
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2023-02-28 06:17:11
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListCertificatesByRecordResponseBodyResultCertificates) String() string {
	return dara.Prettify(s)
}

func (s ListCertificatesByRecordResponseBodyResultCertificates) GoString() string {
	return s.String()
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) GetCasId() *string {
	return s.CasId
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) GetCommonName() *string {
	return s.CommonName
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) GetId() *string {
	return s.Id
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) GetIssuer() *string {
	return s.Issuer
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) GetIssuerCN() *string {
	return s.IssuerCN
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) GetName() *string {
	return s.Name
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) GetNotAfter() *string {
	return s.NotAfter
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) GetNotBefore() *string {
	return s.NotBefore
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) GetPubAlg() *string {
	return s.PubAlg
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) GetRegion() *string {
	return s.Region
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) GetSAN() *string {
	return s.SAN
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) GetSigAlg() *string {
	return s.SigAlg
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) GetStatus() *string {
	return s.Status
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) GetType() *string {
	return s.Type
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) SetCasId(v string) *ListCertificatesByRecordResponseBodyResultCertificates {
	s.CasId = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) SetCommonName(v string) *ListCertificatesByRecordResponseBodyResultCertificates {
	s.CommonName = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) SetCreateTime(v string) *ListCertificatesByRecordResponseBodyResultCertificates {
	s.CreateTime = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) SetFingerprintSha256(v string) *ListCertificatesByRecordResponseBodyResultCertificates {
	s.FingerprintSha256 = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) SetId(v string) *ListCertificatesByRecordResponseBodyResultCertificates {
	s.Id = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) SetIssuer(v string) *ListCertificatesByRecordResponseBodyResultCertificates {
	s.Issuer = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) SetIssuerCN(v string) *ListCertificatesByRecordResponseBodyResultCertificates {
	s.IssuerCN = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) SetName(v string) *ListCertificatesByRecordResponseBodyResultCertificates {
	s.Name = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) SetNotAfter(v string) *ListCertificatesByRecordResponseBodyResultCertificates {
	s.NotAfter = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) SetNotBefore(v string) *ListCertificatesByRecordResponseBodyResultCertificates {
	s.NotBefore = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) SetPubAlg(v string) *ListCertificatesByRecordResponseBodyResultCertificates {
	s.PubAlg = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) SetRegion(v string) *ListCertificatesByRecordResponseBodyResultCertificates {
	s.Region = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) SetSAN(v string) *ListCertificatesByRecordResponseBodyResultCertificates {
	s.SAN = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) SetSerialNumber(v string) *ListCertificatesByRecordResponseBodyResultCertificates {
	s.SerialNumber = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) SetSigAlg(v string) *ListCertificatesByRecordResponseBodyResultCertificates {
	s.SigAlg = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) SetStatus(v string) *ListCertificatesByRecordResponseBodyResultCertificates {
	s.Status = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) SetType(v string) *ListCertificatesByRecordResponseBodyResultCertificates {
	s.Type = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) SetUpdateTime(v string) *ListCertificatesByRecordResponseBodyResultCertificates {
	s.UpdateTime = &v
	return s
}

func (s *ListCertificatesByRecordResponseBodyResultCertificates) Validate() error {
	return dara.Validate(s)
}
