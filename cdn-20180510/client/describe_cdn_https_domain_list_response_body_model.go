// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnHttpsDomainListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertInfos(v *DescribeCdnHttpsDomainListResponseBodyCertInfos) *DescribeCdnHttpsDomainListResponseBody
	GetCertInfos() *DescribeCdnHttpsDomainListResponseBodyCertInfos
	SetRequestId(v string) *DescribeCdnHttpsDomainListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCdnHttpsDomainListResponseBody
	GetTotalCount() *int32
}

type DescribeCdnHttpsDomainListResponseBody struct {
	// The information about the certificate.
	CertInfos *DescribeCdnHttpsDomainListResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F5E8DF64-7175-4186-9B06-F002C0BBD0C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCdnHttpsDomainListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnHttpsDomainListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnHttpsDomainListResponseBody) GetCertInfos() *DescribeCdnHttpsDomainListResponseBodyCertInfos {
	return s.CertInfos
}

func (s *DescribeCdnHttpsDomainListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnHttpsDomainListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCdnHttpsDomainListResponseBody) SetCertInfos(v *DescribeCdnHttpsDomainListResponseBodyCertInfos) *DescribeCdnHttpsDomainListResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBody) SetRequestId(v string) *DescribeCdnHttpsDomainListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBody) SetTotalCount(v int32) *DescribeCdnHttpsDomainListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnHttpsDomainListResponseBodyCertInfos struct {
	CertInfo []*DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeCdnHttpsDomainListResponseBodyCertInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnHttpsDomainListResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfos) GetCertInfo() []*DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo {
	return s.CertInfo
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfos) SetCertInfo(v []*DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) *DescribeCdnHttpsDomainListResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo struct {
	// The returned primary domain name of the certificate.
	//
	// example:
	//
	// example.org
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
	// The time at which the certificate expires.
	//
	// example:
	//
	// 2018-12-26 14:45:09
	CertExpireTime *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// test
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The time at which the certificate became effective.
	//
	// example:
	//
	// 2018-11-26 14:45:09
	CertStartTime *string `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	// The status of the certificate.
	//
	// 	- **ok**: The certificate is working as expected.
	//
	// 	- **mismatch**: The certificate does not match the specified domain name.
	//
	// 	- **expired**: The certificate has expired.
	//
	// 	- **expire_soon**: The certificate will expire soon.
	//
	// example:
	//
	// mismatch
	CertStatus *string `json:"CertStatus,omitempty" xml:"CertStatus,omitempty"`
	// The type of the certificate.
	//
	// 	- **free**: a free certificate.
	//
	// 	- **cas**: a certificate that is purchased from Alibaba Cloud SSL Certificates Service.
	//
	// 	- **upload**: a certificate that is uploaded by the user.
	//
	// example:
	//
	// free
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The time at which the certificate was updated.
	//
	// example:
	//
	// 2019-01-08 18:33:16
	CertUpdateTime *string `json:"CertUpdateTime,omitempty" xml:"CertUpdateTime,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) GetCertCommonName() *string {
	return s.CertCommonName
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) GetCertName() *string {
	return s.CertName
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) GetCertStartTime() *string {
	return s.CertStartTime
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) GetCertStatus() *string {
	return s.CertStatus
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) GetCertType() *string {
	return s.CertType
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) GetCertUpdateTime() *string {
	return s.CertUpdateTime
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertCommonName(v string) *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertCommonName = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertStartTime(v string) *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertStartTime = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertStatus(v string) *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertStatus = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertUpdateTime(v string) *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertUpdateTime = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) SetDomainName(v string) *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponseBodyCertInfosCertInfo) Validate() error {
	return dara.Validate(s)
}
