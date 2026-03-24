// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCertificateInfoByIDResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertInfos(v *DescribeCertificateInfoByIDResponseBodyCertInfos) *DescribeCertificateInfoByIDResponseBody
	GetCertInfos() *DescribeCertificateInfoByIDResponseBodyCertInfos
	SetRequestId(v string) *DescribeCertificateInfoByIDResponseBody
	GetRequestId() *string
}

type DescribeCertificateInfoByIDResponseBody struct {
	CertInfos *DescribeCertificateInfoByIDResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 5C1E43DC-9E51-4771-82C0-7D5ECEB547A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCertificateInfoByIDResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertificateInfoByIDResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCertificateInfoByIDResponseBody) GetCertInfos() *DescribeCertificateInfoByIDResponseBodyCertInfos {
	return s.CertInfos
}

func (s *DescribeCertificateInfoByIDResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCertificateInfoByIDResponseBody) SetCertInfos(v *DescribeCertificateInfoByIDResponseBodyCertInfos) *DescribeCertificateInfoByIDResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeCertificateInfoByIDResponseBody) SetRequestId(v string) *DescribeCertificateInfoByIDResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCertificateInfoByIDResponseBody) Validate() error {
	if s.CertInfos != nil {
		if err := s.CertInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCertificateInfoByIDResponseBodyCertInfos struct {
	CertInfo []*DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeCertificateInfoByIDResponseBodyCertInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertificateInfoByIDResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfos) GetCertInfo() []*DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo {
	return s.CertInfo
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfos) SetCertInfo(v []*DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) *DescribeCertificateInfoByIDResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfos) Validate() error {
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

type DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo struct {
	CertExpireTime *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	CertId         *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertName       *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertType       *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DomainList     *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	HttpsCrt       *string `json:"HttpsCrt,omitempty" xml:"HttpsCrt,omitempty"`
}

func (s DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) GetCertId() *string {
	return s.CertId
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) GetCertName() *string {
	return s.CertName
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) GetCertType() *string {
	return s.CertType
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) GetDomainList() *string {
	return s.DomainList
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) GetHttpsCrt() *string {
	return s.HttpsCrt
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) SetCertId(v string) *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo {
	s.CertId = &v
	return s
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) SetCreateTime(v string) *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) SetDomainList(v string) *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo {
	s.DomainList = &v
	return s
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) SetHttpsCrt(v string) *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo {
	s.HttpsCrt = &v
	return s
}

func (s *DescribeCertificateInfoByIDResponseBodyCertInfosCertInfo) Validate() error {
	return dara.Validate(s)
}
