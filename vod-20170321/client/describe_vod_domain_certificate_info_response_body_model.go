// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainCertificateInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertInfos(v *DescribeVodDomainCertificateInfoResponseBodyCertInfos) *DescribeVodDomainCertificateInfoResponseBody
	GetCertInfos() *DescribeVodDomainCertificateInfoResponseBodyCertInfos
	SetRequestId(v string) *DescribeVodDomainCertificateInfoResponseBody
	GetRequestId() *string
}

type DescribeVodDomainCertificateInfoResponseBody struct {
	CertInfos *DescribeVodDomainCertificateInfoResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 5C1E43DC-9E51-4771-****-7D5ECEB547A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodDomainCertificateInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainCertificateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainCertificateInfoResponseBody) GetCertInfos() *DescribeVodDomainCertificateInfoResponseBodyCertInfos {
	return s.CertInfos
}

func (s *DescribeVodDomainCertificateInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainCertificateInfoResponseBody) SetCertInfos(v *DescribeVodDomainCertificateInfoResponseBodyCertInfos) *DescribeVodDomainCertificateInfoResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBody) SetRequestId(v string) *DescribeVodDomainCertificateInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBody) Validate() error {
	if s.CertInfos != nil {
		if err := s.CertInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVodDomainCertificateInfoResponseBodyCertInfos struct {
	CertInfo []*DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainCertificateInfoResponseBodyCertInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainCertificateInfoResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfos) GetCertInfo() []*DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	return s.CertInfo
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfos) SetCertInfo(v []*DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) *DescribeVodDomainCertificateInfoResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfos) Validate() error {
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

type DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo struct {
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

func (s DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertDomainName() *string {
	return s.CertDomainName
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertId() *string {
	return s.CertId
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertLife() *string {
	return s.CertLife
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertName() *string {
	return s.CertName
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertOrg() *string {
	return s.CertOrg
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertRegion() *string {
	return s.CertRegion
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertStartTime() *string {
	return s.CertStartTime
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertType() *string {
	return s.CertType
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertUpdateTime() *string {
	return s.CertUpdateTime
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetDomainCnameStatus() *string {
	return s.DomainCnameStatus
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetServerCertificate() *string {
	return s.ServerCertificate
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetServerCertificateStatus() *string {
	return s.ServerCertificateStatus
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertDomainName(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertDomainName = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertId(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertId = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertLife(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertLife = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertOrg(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertOrg = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertRegion(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertRegion = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertStartTime(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertStartTime = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertUpdateTime(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertUpdateTime = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetDomainCnameStatus(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.DomainCnameStatus = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetDomainName(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetServerCertificate(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.ServerCertificate = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetServerCertificateStatus(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.ServerCertificateStatus = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetStatus(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.Status = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) Validate() error {
	return dara.Validate(s)
}
