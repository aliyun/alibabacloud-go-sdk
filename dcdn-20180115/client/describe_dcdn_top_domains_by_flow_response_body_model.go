// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnTopDomainsByFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainCount(v int64) *DescribeDcdnTopDomainsByFlowResponseBody
	GetDomainCount() *int64
	SetDomainOnlineCount(v int64) *DescribeDcdnTopDomainsByFlowResponseBody
	GetDomainOnlineCount() *int64
	SetEndTime(v string) *DescribeDcdnTopDomainsByFlowResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDcdnTopDomainsByFlowResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnTopDomainsByFlowResponseBody
	GetStartTime() *string
	SetTopDomains(v *DescribeDcdnTopDomainsByFlowResponseBodyTopDomains) *DescribeDcdnTopDomainsByFlowResponseBody
	GetTopDomains() *DescribeDcdnTopDomainsByFlowResponseBodyTopDomains
}

type DescribeDcdnTopDomainsByFlowResponseBody struct {
	// The total number of accelerated domains under your account.
	//
	// example:
	//
	// 68
	DomainCount *int64 `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	// The total number of accelerated domains that are in the **Enabled*	- state under your account.
	//
	// example:
	//
	// 68
	DomainOnlineCount *int64 `json:"DomainOnlineCount,omitempty" xml:"DomainOnlineCount,omitempty"`
	// The end of the reporting period.
	//
	// example:
	//
	// 2016-03-14T07:34:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4E09C5D7-E1CF-4CAA-A45E-8727F4C8FD70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the reporting period.
	//
	// example:
	//
	// 2016-03-14T06:34:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The top N domain names ranked by network traffic.
	TopDomains *DescribeDcdnTopDomainsByFlowResponseBodyTopDomains `json:"TopDomains,omitempty" xml:"TopDomains,omitempty" type:"Struct"`
}

func (s DescribeDcdnTopDomainsByFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnTopDomainsByFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTopDomainsByFlowResponseBody) GetDomainCount() *int64 {
	return s.DomainCount
}

func (s *DescribeDcdnTopDomainsByFlowResponseBody) GetDomainOnlineCount() *int64 {
	return s.DomainOnlineCount
}

func (s *DescribeDcdnTopDomainsByFlowResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnTopDomainsByFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnTopDomainsByFlowResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnTopDomainsByFlowResponseBody) GetTopDomains() *DescribeDcdnTopDomainsByFlowResponseBodyTopDomains {
	return s.TopDomains
}

func (s *DescribeDcdnTopDomainsByFlowResponseBody) SetDomainCount(v int64) *DescribeDcdnTopDomainsByFlowResponseBody {
	s.DomainCount = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBody) SetDomainOnlineCount(v int64) *DescribeDcdnTopDomainsByFlowResponseBody {
	s.DomainOnlineCount = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBody) SetEndTime(v string) *DescribeDcdnTopDomainsByFlowResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBody) SetRequestId(v string) *DescribeDcdnTopDomainsByFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBody) SetStartTime(v string) *DescribeDcdnTopDomainsByFlowResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBody) SetTopDomains(v *DescribeDcdnTopDomainsByFlowResponseBodyTopDomains) *DescribeDcdnTopDomainsByFlowResponseBody {
	s.TopDomains = v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnTopDomainsByFlowResponseBodyTopDomains struct {
	TopDomain []*DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain `json:"TopDomain,omitempty" xml:"TopDomain,omitempty" type:"Repeated"`
}

func (s DescribeDcdnTopDomainsByFlowResponseBodyTopDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnTopDomainsByFlowResponseBodyTopDomains) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomains) GetTopDomain() []*DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	return s.TopDomain
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomains) SetTopDomain(v []*DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) *DescribeDcdnTopDomainsByFlowResponseBodyTopDomains {
	s.TopDomain = v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain struct {
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The peak bandwidth value.
	//
	// example:
	//
	// 22139626
	MaxBps *int64 `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	// The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-11-01T08:10:00Z
	MaxBpsTime *string `json:"MaxBpsTime,omitempty" xml:"MaxBpsTime,omitempty"`
	// The ranking of the accelerated domain name.
	//
	// example:
	//
	// 1
	Rank *int64 `json:"Rank,omitempty" xml:"Rank,omitempty"`
	// The number of visits.
	//
	// example:
	//
	// 3
	TotalAccess *int64 `json:"TotalAccess,omitempty" xml:"TotalAccess,omitempty"`
	// The total amount of network traffic.
	//
	// example:
	//
	// 123
	TotalTraffic *string `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
	// The proportion of network traffic consumed to access the URL.
	//
	// example:
	//
	// 21.686305274906182
	TrafficPercent *string `json:"TrafficPercent,omitempty" xml:"TrafficPercent,omitempty"`
}

func (s DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) GoString() string {
	return s.String()
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetMaxBps() *int64 {
	return s.MaxBps
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetMaxBpsTime() *string {
	return s.MaxBpsTime
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetRank() *int64 {
	return s.Rank
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetTotalAccess() *int64 {
	return s.TotalAccess
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetTotalTraffic() *string {
	return s.TotalTraffic
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetTrafficPercent() *string {
	return s.TrafficPercent
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetDomainName(v string) *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBps(v int64) *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBps = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBpsTime(v string) *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBpsTime = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetRank(v int64) *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.Rank = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalAccess(v int64) *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalAccess = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalTraffic(v string) *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTrafficPercent(v string) *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TrafficPercent = &v
	return s
}

func (s *DescribeDcdnTopDomainsByFlowResponseBodyTopDomainsTopDomain) Validate() error {
	return dara.Validate(s)
}
