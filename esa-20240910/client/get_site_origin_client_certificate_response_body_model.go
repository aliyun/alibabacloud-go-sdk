// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteOriginClientCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v string) *GetSiteOriginClientCertificateResponseBody
	GetCertificate() *string
	SetRequestId(v string) *GetSiteOriginClientCertificateResponseBody
	GetRequestId() *string
	SetResult(v *GetSiteOriginClientCertificateResponseBodyResult) *GetSiteOriginClientCertificateResponseBody
	GetResult() *GetSiteOriginClientCertificateResponseBodyResult
	SetSiteId(v int64) *GetSiteOriginClientCertificateResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *GetSiteOriginClientCertificateResponseBody
	GetSiteName() *string
	SetStatus(v string) *GetSiteOriginClientCertificateResponseBody
	GetStatus() *string
}

type GetSiteOriginClientCertificateResponseBody struct {
	Certificate *string                                           `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result      *GetSiteOriginClientCertificateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	SiteId      *int64                                            `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName    *string                                           `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	Status      *string                                           `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetSiteOriginClientCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSiteOriginClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *GetSiteOriginClientCertificateResponseBody) GetCertificate() *string {
	return s.Certificate
}

func (s *GetSiteOriginClientCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSiteOriginClientCertificateResponseBody) GetResult() *GetSiteOriginClientCertificateResponseBodyResult {
	return s.Result
}

func (s *GetSiteOriginClientCertificateResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetSiteOriginClientCertificateResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *GetSiteOriginClientCertificateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetSiteOriginClientCertificateResponseBody) SetCertificate(v string) *GetSiteOriginClientCertificateResponseBody {
	s.Certificate = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBody) SetRequestId(v string) *GetSiteOriginClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBody) SetResult(v *GetSiteOriginClientCertificateResponseBodyResult) *GetSiteOriginClientCertificateResponseBody {
	s.Result = v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBody) SetSiteId(v int64) *GetSiteOriginClientCertificateResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBody) SetSiteName(v string) *GetSiteOriginClientCertificateResponseBody {
	s.SiteName = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBody) SetStatus(v string) *GetSiteOriginClientCertificateResponseBody {
	s.Status = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSiteOriginClientCertificateResponseBodyResult struct {
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

func (s GetSiteOriginClientCertificateResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetSiteOriginClientCertificateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetCommonName() *string {
	return s.CommonName
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetIssuer() *string {
	return s.Issuer
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetNotAfter() *string {
	return s.NotAfter
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetNotBefore() *string {
	return s.NotBefore
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetPubkeyAlgorithm() *string {
	return s.PubkeyAlgorithm
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetSAN() *string {
	return s.SAN
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetCommonName(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.CommonName = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetCreateTime(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetFingerprintSha256(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.FingerprintSha256 = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetId(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetIssuer(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.Issuer = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetName(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetNotAfter(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.NotAfter = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetNotBefore(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.NotBefore = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetPubkeyAlgorithm(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.PubkeyAlgorithm = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetSAN(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.SAN = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetSerialNumber(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.SerialNumber = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetSignatureAlgorithm(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.SignatureAlgorithm = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetStatus(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetType(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) SetUpdateTime(v string) *GetSiteOriginClientCertificateResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *GetSiteOriginClientCertificateResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
