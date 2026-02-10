// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainCertificateInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertInfos(v *DescribeLiveDomainCertificateInfoResponseBodyCertInfos) *DescribeLiveDomainCertificateInfoResponseBody
	GetCertInfos() *DescribeLiveDomainCertificateInfoResponseBodyCertInfos
	SetRequestId(v string) *DescribeLiveDomainCertificateInfoResponseBody
	GetRequestId() *string
}

type DescribeLiveDomainCertificateInfoResponseBody struct {
	CertInfos *DescribeLiveDomainCertificateInfoResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5C1E43DC-9E51-4771-82C0-7D5ECEB547A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveDomainCertificateInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainCertificateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainCertificateInfoResponseBody) GetCertInfos() *DescribeLiveDomainCertificateInfoResponseBodyCertInfos {
	return s.CertInfos
}

func (s *DescribeLiveDomainCertificateInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainCertificateInfoResponseBody) SetCertInfos(v *DescribeLiveDomainCertificateInfoResponseBodyCertInfos) *DescribeLiveDomainCertificateInfoResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBody) SetRequestId(v string) *DescribeLiveDomainCertificateInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBody) Validate() error {
	if s.CertInfos != nil {
		if err := s.CertInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainCertificateInfoResponseBodyCertInfos struct {
	CertInfo []*DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainCertificateInfoResponseBodyCertInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainCertificateInfoResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfos) GetCertInfo() []*DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	return s.CertInfo
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfos) SetCertInfo(v []*DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) *DescribeLiveDomainCertificateInfoResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfos) Validate() error {
	if s.CertInfo != nil {
		for _, item := range s.CertInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo struct {
	CertDomainName *string `json:"CertDomainName,omitempty" xml:"CertDomainName,omitempty"`
	CertExpireTime *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	CertLife       *string `json:"CertLife,omitempty" xml:"CertLife,omitempty"`
	CertName       *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertOrg        *string `json:"CertOrg,omitempty" xml:"CertOrg,omitempty"`
	CertType       *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	SSLProtocol    *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	SSLPub         *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertDomainName() *string {
	return s.CertDomainName
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertLife() *string {
	return s.CertLife
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertName() *string {
	return s.CertName
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertOrg() *string {
	return s.CertOrg
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertType() *string {
	return s.CertType
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetSSLProtocol() *string {
	return s.SSLProtocol
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetSSLPub() *string {
	return s.SSLPub
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertDomainName(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertDomainName = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertLife(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertLife = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertOrg(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertOrg = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetDomainName(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetSSLProtocol(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetSSLPub(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.SSLPub = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetStatus(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.Status = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) Validate() error {
	return dara.Validate(s)
}
