// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveHttpsDomainListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertInfos(v *DescribeLiveHttpsDomainListResponseBodyCertInfos) *DescribeLiveHttpsDomainListResponseBody
	GetCertInfos() *DescribeLiveHttpsDomainListResponseBodyCertInfos
	SetRequestId(v string) *DescribeLiveHttpsDomainListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeLiveHttpsDomainListResponseBody
	GetTotalCount() *int32
}

type DescribeLiveHttpsDomainListResponseBody struct {
	CertInfos *DescribeLiveHttpsDomainListResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	// The request ID.
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

func (s DescribeLiveHttpsDomainListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveHttpsDomainListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveHttpsDomainListResponseBody) GetCertInfos() *DescribeLiveHttpsDomainListResponseBodyCertInfos {
	return s.CertInfos
}

func (s *DescribeLiveHttpsDomainListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveHttpsDomainListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeLiveHttpsDomainListResponseBody) SetCertInfos(v *DescribeLiveHttpsDomainListResponseBodyCertInfos) *DescribeLiveHttpsDomainListResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeLiveHttpsDomainListResponseBody) SetRequestId(v string) *DescribeLiveHttpsDomainListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveHttpsDomainListResponseBody) SetTotalCount(v int32) *DescribeLiveHttpsDomainListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLiveHttpsDomainListResponseBody) Validate() error {
	if s.CertInfos != nil {
		if err := s.CertInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveHttpsDomainListResponseBodyCertInfos struct {
	CertInfo []*DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveHttpsDomainListResponseBodyCertInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveHttpsDomainListResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfos) GetCertInfo() []*DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo {
	return s.CertInfo
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfos) SetCertInfo(v []*DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) *DescribeLiveHttpsDomainListResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfos) Validate() error {
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

type DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo struct {
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
	CertExpireTime *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	CertName       *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertStartTime  *string `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	CertStatus     *string `json:"CertStatus,omitempty" xml:"CertStatus,omitempty"`
	CertType       *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	CertUpdateTime *string `json:"CertUpdateTime,omitempty" xml:"CertUpdateTime,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) GetCertCommonName() *string {
	return s.CertCommonName
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) GetCertName() *string {
	return s.CertName
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) GetCertStartTime() *string {
	return s.CertStartTime
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) GetCertStatus() *string {
	return s.CertStatus
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) GetCertType() *string {
	return s.CertType
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) GetCertUpdateTime() *string {
	return s.CertUpdateTime
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) SetCertCommonName(v string) *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertCommonName = &v
	return s
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) SetCertStartTime(v string) *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertStartTime = &v
	return s
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) SetCertStatus(v string) *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertStatus = &v
	return s
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) SetCertUpdateTime(v string) *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo {
	s.CertUpdateTime = &v
	return s
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) SetDomainName(v string) *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveHttpsDomainListResponseBodyCertInfosCertInfo) Validate() error {
	return dara.Validate(s)
}
