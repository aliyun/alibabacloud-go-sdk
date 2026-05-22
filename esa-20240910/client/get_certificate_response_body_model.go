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
	Certificate *string                           `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	RequestId   *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result      *GetCertificateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	SiteId      *int64                            `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName    *string                           `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	Status      *string                           `json:"Status,omitempty" xml:"Status,omitempty"`
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
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCertificateResponseBodyResult struct {
	ApplyCode         *int64                                 `json:"ApplyCode,omitempty" xml:"ApplyCode,omitempty"`
	ApplyMessage      *string                                `json:"ApplyMessage,omitempty" xml:"ApplyMessage,omitempty"`
	CasId             *string                                `json:"CasId,omitempty" xml:"CasId,omitempty"`
	CommonName        *string                                `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	CreateTime        *string                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DCV               []*GetCertificateResponseBodyResultDCV `json:"DCV,omitempty" xml:"DCV,omitempty" type:"Repeated"`
	FingerprintSha256 *string                                `json:"FingerprintSha256,omitempty" xml:"FingerprintSha256,omitempty"`
	Id                *string                                `json:"Id,omitempty" xml:"Id,omitempty"`
	Issuer            *string                                `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	IssuerCN          *string                                `json:"IssuerCN,omitempty" xml:"IssuerCN,omitempty"`
	KeyServerId       *string                                `json:"KeyServerId,omitempty" xml:"KeyServerId,omitempty"`
	Name              *string                                `json:"Name,omitempty" xml:"Name,omitempty"`
	NotAfter          *string                                `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	NotBefore         *string                                `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	PubAlg            *string                                `json:"PubAlg,omitempty" xml:"PubAlg,omitempty"`
	Region            *string                                `json:"Region,omitempty" xml:"Region,omitempty"`
	SAN               *string                                `json:"SAN,omitempty" xml:"SAN,omitempty"`
	SerialNumber      *string                                `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SigAlg            *string                                `json:"SigAlg,omitempty" xml:"SigAlg,omitempty"`
	Status            *string                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Type              *string                                `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime        *string                                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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

func (s *GetCertificateResponseBodyResult) GetKeyServerId() *string {
	return s.KeyServerId
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

func (s *GetCertificateResponseBodyResult) SetKeyServerId(v string) *GetCertificateResponseBodyResult {
	s.KeyServerId = &v
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

type GetCertificateResponseBodyResultDCV struct {
	Id     *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Key    *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value  *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
