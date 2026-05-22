// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginClientCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v string) *GetOriginClientCertificateResponseBody
	GetCertificate() *string
	SetRequestId(v string) *GetOriginClientCertificateResponseBody
	GetRequestId() *string
	SetResult(v *GetOriginClientCertificateResponseBodyResult) *GetOriginClientCertificateResponseBody
	GetResult() *GetOriginClientCertificateResponseBodyResult
	SetSiteId(v int64) *GetOriginClientCertificateResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *GetOriginClientCertificateResponseBody
	GetSiteName() *string
	SetStatus(v string) *GetOriginClientCertificateResponseBody
	GetStatus() *string
}

type GetOriginClientCertificateResponseBody struct {
	Certificate *string                                       `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result      *GetOriginClientCertificateResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	SiteId      *int64                                        `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName    *string                                       `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	Status      *string                                       `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetOriginClientCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOriginClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *GetOriginClientCertificateResponseBody) GetCertificate() *string {
	return s.Certificate
}

func (s *GetOriginClientCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOriginClientCertificateResponseBody) GetResult() *GetOriginClientCertificateResponseBodyResult {
	return s.Result
}

func (s *GetOriginClientCertificateResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetOriginClientCertificateResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *GetOriginClientCertificateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetOriginClientCertificateResponseBody) SetCertificate(v string) *GetOriginClientCertificateResponseBody {
	s.Certificate = &v
	return s
}

func (s *GetOriginClientCertificateResponseBody) SetRequestId(v string) *GetOriginClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOriginClientCertificateResponseBody) SetResult(v *GetOriginClientCertificateResponseBodyResult) *GetOriginClientCertificateResponseBody {
	s.Result = v
	return s
}

func (s *GetOriginClientCertificateResponseBody) SetSiteId(v int64) *GetOriginClientCertificateResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetOriginClientCertificateResponseBody) SetSiteName(v string) *GetOriginClientCertificateResponseBody {
	s.SiteName = &v
	return s
}

func (s *GetOriginClientCertificateResponseBody) SetStatus(v string) *GetOriginClientCertificateResponseBody {
	s.Status = &v
	return s
}

func (s *GetOriginClientCertificateResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOriginClientCertificateResponseBodyResult struct {
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

func (s GetOriginClientCertificateResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetOriginClientCertificateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOriginClientCertificateResponseBodyResult) GetCommonName() *string {
	return s.CommonName
}

func (s *GetOriginClientCertificateResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetOriginClientCertificateResponseBodyResult) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *GetOriginClientCertificateResponseBodyResult) GetHostnames() []*string {
	return s.Hostnames
}

func (s *GetOriginClientCertificateResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *GetOriginClientCertificateResponseBodyResult) GetIssuer() *string {
	return s.Issuer
}

func (s *GetOriginClientCertificateResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetOriginClientCertificateResponseBodyResult) GetNotAfter() *string {
	return s.NotAfter
}

func (s *GetOriginClientCertificateResponseBodyResult) GetNotBefore() *string {
	return s.NotBefore
}

func (s *GetOriginClientCertificateResponseBodyResult) GetPubkeyAlgorithm() *string {
	return s.PubkeyAlgorithm
}

func (s *GetOriginClientCertificateResponseBodyResult) GetSAN() *string {
	return s.SAN
}

func (s *GetOriginClientCertificateResponseBodyResult) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *GetOriginClientCertificateResponseBodyResult) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *GetOriginClientCertificateResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetOriginClientCertificateResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetOriginClientCertificateResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetOriginClientCertificateResponseBodyResult) SetCommonName(v string) *GetOriginClientCertificateResponseBodyResult {
	s.CommonName = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetCreateTime(v string) *GetOriginClientCertificateResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetFingerprintSha256(v string) *GetOriginClientCertificateResponseBodyResult {
	s.FingerprintSha256 = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetHostnames(v []*string) *GetOriginClientCertificateResponseBodyResult {
	s.Hostnames = v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetId(v string) *GetOriginClientCertificateResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetIssuer(v string) *GetOriginClientCertificateResponseBodyResult {
	s.Issuer = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetName(v string) *GetOriginClientCertificateResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetNotAfter(v string) *GetOriginClientCertificateResponseBodyResult {
	s.NotAfter = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetNotBefore(v string) *GetOriginClientCertificateResponseBodyResult {
	s.NotBefore = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetPubkeyAlgorithm(v string) *GetOriginClientCertificateResponseBodyResult {
	s.PubkeyAlgorithm = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetSAN(v string) *GetOriginClientCertificateResponseBodyResult {
	s.SAN = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetSerialNumber(v string) *GetOriginClientCertificateResponseBodyResult {
	s.SerialNumber = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetSignatureAlgorithm(v string) *GetOriginClientCertificateResponseBodyResult {
	s.SignatureAlgorithm = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetStatus(v string) *GetOriginClientCertificateResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetType(v string) *GetOriginClientCertificateResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) SetUpdateTime(v string) *GetOriginClientCertificateResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *GetOriginClientCertificateResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
