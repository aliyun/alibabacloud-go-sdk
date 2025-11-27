// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainCertificateInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertInfos(v *DescribeVsDomainCertificateInfoResponseBodyCertInfos) *DescribeVsDomainCertificateInfoResponseBody
	GetCertInfos() *DescribeVsDomainCertificateInfoResponseBodyCertInfos
	SetRequestId(v string) *DescribeVsDomainCertificateInfoResponseBody
	GetRequestId() *string
}

type DescribeVsDomainCertificateInfoResponseBody struct {
	CertInfos *DescribeVsDomainCertificateInfoResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsDomainCertificateInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainCertificateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainCertificateInfoResponseBody) GetCertInfos() *DescribeVsDomainCertificateInfoResponseBodyCertInfos {
	return s.CertInfos
}

func (s *DescribeVsDomainCertificateInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsDomainCertificateInfoResponseBody) SetCertInfos(v *DescribeVsDomainCertificateInfoResponseBodyCertInfos) *DescribeVsDomainCertificateInfoResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBody) SetRequestId(v string) *DescribeVsDomainCertificateInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBody) Validate() error {
	if s.CertInfos != nil {
		if err := s.CertInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVsDomainCertificateInfoResponseBodyCertInfos struct {
	CertInfo []*DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainCertificateInfoResponseBodyCertInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainCertificateInfoResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfos) GetCertInfo() []*DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	return s.CertInfo
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfos) SetCertInfo(v []*DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) *DescribeVsDomainCertificateInfoResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfos) Validate() error {
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

type DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo struct {
	// example:
	//
	// example.com
	CertDomainName *string `json:"CertDomainName,omitempty" xml:"CertDomainName,omitempty"`
	// example:
	//
	// 2018-06-03T22:03:39Z
	CertExpireTime *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	// example:
	//
	// months
	CertLife *string `json:"CertLife,omitempty" xml:"CertLife,omitempty"`
	// example:
	//
	// myname
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// example:
	//
	// Let\\"s Encrypt
	CertOrg *string `json:"CertOrg,omitempty" xml:"CertOrg,omitempty"`
	// example:
	//
	// free
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// asdadaxxxx
	SSLPub *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	// example:
	//
	// on
	ServerCertificateStatus *string `json:"ServerCertificateStatus,omitempty" xml:"ServerCertificateStatus,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertDomainName() *string {
	return s.CertDomainName
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertLife() *string {
	return s.CertLife
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertName() *string {
	return s.CertName
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertOrg() *string {
	return s.CertOrg
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertType() *string {
	return s.CertType
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) GetSSLPub() *string {
	return s.SSLPub
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) GetServerCertificateStatus() *string {
	return s.ServerCertificateStatus
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertDomainName(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertDomainName = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertLife(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertLife = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertOrg(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertOrg = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetDomainName(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetSSLPub(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.SSLPub = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetServerCertificateStatus(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.ServerCertificateStatus = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) SetStatus(v string) *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.Status = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponseBodyCertInfosCertInfo) Validate() error {
	return dara.Validate(s)
}
