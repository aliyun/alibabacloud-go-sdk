// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainCertificateInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertInfos(v *DescribeDomainCertificateInfoResponseBodyCertInfos) *DescribeDomainCertificateInfoResponseBody
	GetCertInfos() *DescribeDomainCertificateInfoResponseBodyCertInfos
	SetRequestId(v string) *DescribeDomainCertificateInfoResponseBody
	GetRequestId() *string
}

type DescribeDomainCertificateInfoResponseBody struct {
	CertInfos *DescribeDomainCertificateInfoResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 5C1E43DC-9E51-4771-82C0-7D5ECEB547A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainCertificateInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCertificateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainCertificateInfoResponseBody) GetCertInfos() *DescribeDomainCertificateInfoResponseBodyCertInfos {
	return s.CertInfos
}

func (s *DescribeDomainCertificateInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainCertificateInfoResponseBody) SetCertInfos(v *DescribeDomainCertificateInfoResponseBodyCertInfos) *DescribeDomainCertificateInfoResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBody) SetRequestId(v string) *DescribeDomainCertificateInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBody) Validate() error {
	if s.CertInfos != nil {
		if err := s.CertInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainCertificateInfoResponseBodyCertInfos struct {
	CertInfo []*DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeDomainCertificateInfoResponseBodyCertInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCertificateInfoResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfos) GetCertInfo() []*DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	return s.CertInfo
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfos) SetCertInfo(v []*DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) *DescribeDomainCertificateInfoResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfos) Validate() error {
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

type DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo struct {
	CertDomainName          *string `json:"CertDomainName,omitempty" xml:"CertDomainName,omitempty"`
	CertExpireTime          *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	CertId                  *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertLife                *string `json:"CertLife,omitempty" xml:"CertLife,omitempty"`
	CertName                *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertOrg                 *string `json:"CertOrg,omitempty" xml:"CertOrg,omitempty"`
	CertRegion              *string `json:"CertRegion,omitempty" xml:"CertRegion,omitempty"`
	CertStartTime           *string `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	CertType                *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertUpdateTime          *string `json:"CertUpdateTime,omitempty" xml:"CertUpdateTime,omitempty"`
	DomainCnameStatus       *string `json:"DomainCnameStatus,omitempty" xml:"DomainCnameStatus,omitempty"`
	DomainName              *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ServerCertificate       *string `json:"ServerCertificate,omitempty" xml:"ServerCertificate,omitempty"`
	ServerCertificateStatus *string `json:"ServerCertificateStatus,omitempty" xml:"ServerCertificateStatus,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertDomainName() *string {
	return s.CertDomainName
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertId() *string {
	return s.CertId
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertLife() *string {
	return s.CertLife
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertName() *string {
	return s.CertName
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertOrg() *string {
	return s.CertOrg
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertRegion() *string {
	return s.CertRegion
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertStartTime() *string {
	return s.CertStartTime
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertType() *string {
	return s.CertType
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertUpdateTime() *string {
	return s.CertUpdateTime
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetDomainCnameStatus() *string {
	return s.DomainCnameStatus
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetServerCertificate() *string {
	return s.ServerCertificate
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetServerCertificateStatus() *string {
	return s.ServerCertificateStatus
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertDomainName(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertDomainName = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertId(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertId = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertLife(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertLife = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertOrg(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertOrg = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertRegion(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertRegion = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertStartTime(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertStartTime = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertUpdateTime(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertUpdateTime = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetDomainCnameStatus(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.DomainCnameStatus = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetDomainName(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetServerCertificate(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.ServerCertificate = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetServerCertificateStatus(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.ServerCertificateStatus = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetStatus(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.Status = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) Validate() error {
	return dara.Validate(s)
}
