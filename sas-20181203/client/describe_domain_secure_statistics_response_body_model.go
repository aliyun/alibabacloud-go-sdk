// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecureStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmCount(v int32) *DescribeDomainSecureStatisticsResponseBody
	GetAlarmCount() *int32
	SetNoSslCount(v int32) *DescribeDomainSecureStatisticsResponseBody
	GetNoSslCount() *int32
	SetRequestId(v string) *DescribeDomainSecureStatisticsResponseBody
	GetRequestId() *string
	SetRiskCount(v int32) *DescribeDomainSecureStatisticsResponseBody
	GetRiskCount() *int32
	SetTotalDomainCount(v int32) *DescribeDomainSecureStatisticsResponseBody
	GetTotalDomainCount() *int32
	SetVulCount(v int32) *DescribeDomainSecureStatisticsResponseBody
	GetVulCount() *int32
}

type DescribeDomainSecureStatisticsResponseBody struct {
	// The number of domain names that trigger security alerts.
	//
	// example:
	//
	// 2
	AlarmCount *int32 `json:"AlarmCount,omitempty" xml:"AlarmCount,omitempty"`
	// The number of the websites for which no certificates are installed.
	//
	// example:
	//
	// 1
	NoSslCount *int32 `json:"NoSslCount,omitempty" xml:"NoSslCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1EE7B150-D67E-53FD-A52D-3E8E669A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of the domain names that have security risks.
	//
	// example:
	//
	// 1
	RiskCount *int32 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The total number of domain names.
	//
	// example:
	//
	// 72
	TotalDomainCount *int32 `json:"TotalDomainCount,omitempty" xml:"TotalDomainCount,omitempty"`
	// The number of the domain names that have vulnerabilities.
	//
	// example:
	//
	// 2
	VulCount *int32 `json:"VulCount,omitempty" xml:"VulCount,omitempty"`
}

func (s DescribeDomainSecureStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureStatisticsResponseBody) GetAlarmCount() *int32 {
	return s.AlarmCount
}

func (s *DescribeDomainSecureStatisticsResponseBody) GetNoSslCount() *int32 {
	return s.NoSslCount
}

func (s *DescribeDomainSecureStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainSecureStatisticsResponseBody) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *DescribeDomainSecureStatisticsResponseBody) GetTotalDomainCount() *int32 {
	return s.TotalDomainCount
}

func (s *DescribeDomainSecureStatisticsResponseBody) GetVulCount() *int32 {
	return s.VulCount
}

func (s *DescribeDomainSecureStatisticsResponseBody) SetAlarmCount(v int32) *DescribeDomainSecureStatisticsResponseBody {
	s.AlarmCount = &v
	return s
}

func (s *DescribeDomainSecureStatisticsResponseBody) SetNoSslCount(v int32) *DescribeDomainSecureStatisticsResponseBody {
	s.NoSslCount = &v
	return s
}

func (s *DescribeDomainSecureStatisticsResponseBody) SetRequestId(v string) *DescribeDomainSecureStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSecureStatisticsResponseBody) SetRiskCount(v int32) *DescribeDomainSecureStatisticsResponseBody {
	s.RiskCount = &v
	return s
}

func (s *DescribeDomainSecureStatisticsResponseBody) SetTotalDomainCount(v int32) *DescribeDomainSecureStatisticsResponseBody {
	s.TotalDomainCount = &v
	return s
}

func (s *DescribeDomainSecureStatisticsResponseBody) SetVulCount(v int32) *DescribeDomainSecureStatisticsResponseBody {
	s.VulCount = &v
	return s
}

func (s *DescribeDomainSecureStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}
