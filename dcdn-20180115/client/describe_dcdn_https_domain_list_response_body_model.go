// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnHttpsDomainListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertInfos(v *DescribeDcdnHttpsDomainListResponseBodyCertInfos) *DescribeDcdnHttpsDomainListResponseBody
	GetCertInfos() *DescribeDcdnHttpsDomainListResponseBodyCertInfos
	SetRequestId(v string) *DescribeDcdnHttpsDomainListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDcdnHttpsDomainListResponseBody
	GetTotalCount() *int32
}

type DescribeDcdnHttpsDomainListResponseBody struct {
	// The information about the certificate.
	CertInfos *DescribeDcdnHttpsDomainListResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
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

func (s DescribeDcdnHttpsDomainListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnHttpsDomainListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnHttpsDomainListResponseBody) GetCertInfos() *DescribeDcdnHttpsDomainListResponseBodyCertInfos {
	return s.CertInfos
}

func (s *DescribeDcdnHttpsDomainListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnHttpsDomainListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDcdnHttpsDomainListResponseBody) SetCertInfos(v *DescribeDcdnHttpsDomainListResponseBodyCertInfos) *DescribeDcdnHttpsDomainListResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBody) SetRequestId(v string) *DescribeDcdnHttpsDomainListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBody) SetTotalCount(v int32) *DescribeDcdnHttpsDomainListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnHttpsDomainListResponseBodyCertInfos struct {
	CertInfo []*DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeDcdnHttpsDomainListResponseBodyCertInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnHttpsDomainListResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfos) GetCertInfo() []*DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo {
	return s.CertInfo
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfos) SetCertInfo(v []*DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) *DescribeDcdnHttpsDomainListResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo struct {
	// The returned primary domain name of the certificate.
	//
	// example:
	//
	// *.com
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
	// cert
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The time at which the certificate became effective.
	//
	// example:
	//
	// 2018-11-26 14:45:09
	CertStartTime *string `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	// The status of the certificate. Valid values:
	//
	// 	- **ok**: The certificate is working as expected.
	//
	// 	- **mismatch**: The certificate does not match the specified domain name.
	//
	// 	- **expired**: The certificate has expired.
	//
	// 	- **expire_soon**: The certificate is about to expire.
	//
	// example:
	//
	// mismatch
	CertStatus *string `json:"CertStatus,omitempty" xml:"CertStatus,omitempty"`
	// The type of the certificate. Valid values:
	//
	// 	- **cas**: a certificate that is purchased by using Certificate Management Service
	//
	// 	- **upload**: a custom certificate that you upload
	//
	// example:
	//
	// upload
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
	// *.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) GetCertCommonName() *string {
	return s.CertCommonName
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) GetCertName() *string {
	return s.CertName
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) GetCertStartTime() *string {
	return s.CertStartTime
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) GetCertStatus() *string {
	return s.CertStatus
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) GetCertType() *string {
	return s.CertType
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) GetCertUpdateTime() *string {
	return s.CertUpdateTime
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertCommonName(v string) *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertCommonName = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertStartTime(v string) *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertStartTime = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertStatus(v string) *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertStatus = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) SetCertUpdateTime(v string) *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertUpdateTime = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) SetDomainName(v string) *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponseBodyCertInfosCertInfo) Validate() error {
	return dara.Validate(s)
}
