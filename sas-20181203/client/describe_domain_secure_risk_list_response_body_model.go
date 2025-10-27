// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecureRiskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNoSslCount(v int32) *DescribeDomainSecureRiskListResponseBody
	GetNoSslCount() *int32
	SetRequestId(v string) *DescribeDomainSecureRiskListResponseBody
	GetRequestId() *string
	SetRiskCount(v int32) *DescribeDomainSecureRiskListResponseBody
	GetRiskCount() *int32
	SetRiskList(v []*DescribeDomainSecureRiskListResponseBodyRiskList) *DescribeDomainSecureRiskListResponseBody
	GetRiskList() []*DescribeDomainSecureRiskListResponseBodyRiskList
}

type DescribeDomainSecureRiskListResponseBody struct {
	// The number of the websites for which no certificates are installed.
	//
	// example:
	//
	// 1
	NoSslCount *int32 `json:"NoSslCount,omitempty" xml:"NoSslCount,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// CE500770-42D3-442E-9DDD-156E0F9F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of risks.
	//
	// example:
	//
	// 1
	RiskCount *int32 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The risks.
	RiskList []*DescribeDomainSecureRiskListResponseBodyRiskList `json:"RiskList,omitempty" xml:"RiskList,omitempty" type:"Repeated"`
}

func (s DescribeDomainSecureRiskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureRiskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureRiskListResponseBody) GetNoSslCount() *int32 {
	return s.NoSslCount
}

func (s *DescribeDomainSecureRiskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainSecureRiskListResponseBody) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *DescribeDomainSecureRiskListResponseBody) GetRiskList() []*DescribeDomainSecureRiskListResponseBodyRiskList {
	return s.RiskList
}

func (s *DescribeDomainSecureRiskListResponseBody) SetNoSslCount(v int32) *DescribeDomainSecureRiskListResponseBody {
	s.NoSslCount = &v
	return s
}

func (s *DescribeDomainSecureRiskListResponseBody) SetRequestId(v string) *DescribeDomainSecureRiskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSecureRiskListResponseBody) SetRiskCount(v int32) *DescribeDomainSecureRiskListResponseBody {
	s.RiskCount = &v
	return s
}

func (s *DescribeDomainSecureRiskListResponseBody) SetRiskList(v []*DescribeDomainSecureRiskListResponseBodyRiskList) *DescribeDomainSecureRiskListResponseBody {
	s.RiskList = v
	return s
}

func (s *DescribeDomainSecureRiskListResponseBody) Validate() error {
	if s.RiskList != nil {
		for _, item := range s.RiskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainSecureRiskListResponseBodyRiskList struct {
	// The number of alerts.
	//
	// example:
	//
	// 1
	AlarmCount *int32 `json:"AlarmCount,omitempty" xml:"AlarmCount,omitempty"`
	// The domain name.
	//
	// example:
	//
	// test.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The issuer of the certificate.
	//
	// example:
	//
	// globalsign
	SslBrand *string `json:"SslBrand,omitempty" xml:"SslBrand,omitempty"`
	// Indicates whether the certificate is configured. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	SslStatus *int32 `json:"SslStatus,omitempty" xml:"SslStatus,omitempty"`
	// The UUIDs of the backend servers of the website.
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
	// The number of vulnerabilities.
	//
	// example:
	//
	// 1
	VulCount *int32 `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
}

func (s DescribeDomainSecureRiskListResponseBodyRiskList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureRiskListResponseBodyRiskList) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureRiskListResponseBodyRiskList) GetAlarmCount() *int32 {
	return s.AlarmCount
}

func (s *DescribeDomainSecureRiskListResponseBodyRiskList) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainSecureRiskListResponseBodyRiskList) GetSslBrand() *string {
	return s.SslBrand
}

func (s *DescribeDomainSecureRiskListResponseBodyRiskList) GetSslStatus() *int32 {
	return s.SslStatus
}

func (s *DescribeDomainSecureRiskListResponseBodyRiskList) GetUuidList() []*string {
	return s.UuidList
}

func (s *DescribeDomainSecureRiskListResponseBodyRiskList) GetVulCount() *int32 {
	return s.VulCount
}

func (s *DescribeDomainSecureRiskListResponseBodyRiskList) SetAlarmCount(v int32) *DescribeDomainSecureRiskListResponseBodyRiskList {
	s.AlarmCount = &v
	return s
}

func (s *DescribeDomainSecureRiskListResponseBodyRiskList) SetDomain(v string) *DescribeDomainSecureRiskListResponseBodyRiskList {
	s.Domain = &v
	return s
}

func (s *DescribeDomainSecureRiskListResponseBodyRiskList) SetSslBrand(v string) *DescribeDomainSecureRiskListResponseBodyRiskList {
	s.SslBrand = &v
	return s
}

func (s *DescribeDomainSecureRiskListResponseBodyRiskList) SetSslStatus(v int32) *DescribeDomainSecureRiskListResponseBodyRiskList {
	s.SslStatus = &v
	return s
}

func (s *DescribeDomainSecureRiskListResponseBodyRiskList) SetUuidList(v []*string) *DescribeDomainSecureRiskListResponseBodyRiskList {
	s.UuidList = v
	return s
}

func (s *DescribeDomainSecureRiskListResponseBodyRiskList) SetVulCount(v int32) *DescribeDomainSecureRiskListResponseBodyRiskList {
	s.VulCount = &v
	return s
}

func (s *DescribeDomainSecureRiskListResponseBodyRiskList) Validate() error {
	return dara.Validate(s)
}
