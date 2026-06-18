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
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of certificate details.
	Result []*ListCertificatesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The ID of the site.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The name of the site.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The error code returned when a certificate application fails.
	//
	// example:
	//
	// 2
	ApplyCode *int64 `json:"ApplyCode,omitempty" xml:"ApplyCode,omitempty"`
	// The error message returned when a certificate application fails.
	//
	// example:
	//
	// canceled
	ApplyMessage *string `json:"ApplyMessage,omitempty" xml:"ApplyMessage,omitempty"`
	// The ID of the cloud certificate.
	//
	// example:
	//
	// 30000569
	CasId *string `json:"CasId,omitempty" xml:"CasId,omitempty"`
	// The common name of the certificate.
	//
	// example:
	//
	// www.example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The time when the certificate was created.
	//
	// example:
	//
	// 2022-06-24 07:48:51
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The domain control validation (DCV) information.
	DCV []*ListCertificatesResponseBodyResultDCV `json:"DCV,omitempty" xml:"DCV,omitempty" type:"Repeated"`
	// The SHA-256 fingerprint of the certificate.
	//
	// example:
	//
	// 1dc5fc9af4eead2570c70d94b416130baeb6d4429b51fd3557379588456a****
	FingerprintSha256 *string `json:"FingerprintSha256,omitempty" xml:"FingerprintSha256,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// baba39055622c008b90285a8838e****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The issuer of the certificate.
	//
	// example:
	//
	// GlobalSign nv-sa
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The common name of the certificate issuer.
	//
	// example:
	//
	// GlobalSign Organization Validation CA - SHA256 - G3
	IssuerCN *string `json:"IssuerCN,omitempty" xml:"IssuerCN,omitempty"`
	// The ID of the keyless server. This parameter is valid only when `Type` is set to `keyless`.
	//
	// example:
	//
	// baba39055622c008b90285a8838e****
	KeyServerId *string `json:"KeyServerId,omitempty" xml:"KeyServerId,omitempty"`
	// The name of the certificate.
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
	// The time when the certificate becomes valid.
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
	PubAlg *string `json:"PubAlg,omitempty" xml:"PubAlg,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The subject alternative name (SAN) of the certificate.
	//
	// example:
	//
	// www.example.com,*.example.com
	SAN *string `json:"SAN,omitempty" xml:"SAN,omitempty"`
	// The serial number of the certificate.
	//
	// example:
	//
	// babab022c5e9b27bf9c64d7f4b16****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The signature algorithm of the certificate.
	//
	// example:
	//
	// SHA256-RSA
	SigAlg *string `json:"SigAlg,omitempty" xml:"SigAlg,omitempty"`
	// The status of the certificate. Valid values:
	//
	// - `OK`: The certificate is valid.
	//
	// - `Expired`: The certificate has expired.
	//
	// - `Expiring`: The certificate is due to expire within 30 days.
	//
	// - `Issued`: The free certificate has been issued.
	//
	// - `Applying`: The free certificate application is in progress.
	//
	// - `ApplyFailed`: The free certificate application has failed.
	//
	// - `Canceled`: The free certificate application has been canceled.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the certificate. Valid values:
	//
	// - `cas`: An Alibaba Cloud certificate.
	//
	// - `upload`: A user-uploaded certificate.
	//
	// - `free`: A free certificate.
	//
	// - `keyless`: A keyless certificate.
	//
	// example:
	//
	// free
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the certificate was last updated.
	//
	// example:
	//
	// 2023-04-20 06:18:42
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	// The ID of the DCV record.
	//
	// example:
	//
	// bababf7cdd1546a2ad04c0def1f4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DCV key. For the `DNS` method, this is the name of the TXT record. For the `HTTP` method, this is the URL.
	//
	// example:
	//
	// http://www.example.com/.well-known/acme-challenge/pH20CqwS5L3ZnvkhI436DCzadKFuG7QcUcvB_4KsAow
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The validation status.
	//
	// example:
	//
	// pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The DCV method. Valid values: `DNS` and `HTTP`.
	//
	// example:
	//
	// HTTP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The DCV value.
	//
	// example:
	//
	// pH20CqwS5L3ZnvkhI436DCzadKFuG7QcUcvB_4KsAow.KfzYo4LH3EgOt7a73G-RqZkbR0eYtLfEUmtmqGmr4FQ
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
