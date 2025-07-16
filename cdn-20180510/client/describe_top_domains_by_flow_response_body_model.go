// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTopDomainsByFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainCount(v int64) *DescribeTopDomainsByFlowResponseBody
	GetDomainCount() *int64
	SetDomainOnlineCount(v int64) *DescribeTopDomainsByFlowResponseBody
	GetDomainOnlineCount() *int64
	SetEndTime(v string) *DescribeTopDomainsByFlowResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeTopDomainsByFlowResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeTopDomainsByFlowResponseBody
	GetStartTime() *string
	SetTopDomains(v *DescribeTopDomainsByFlowResponseBodyTopDomains) *DescribeTopDomainsByFlowResponseBody
	GetTopDomains() *DescribeTopDomainsByFlowResponseBodyTopDomains
}

type DescribeTopDomainsByFlowResponseBody struct {
	// The total number of accelerated domain names that belong to the current Alibaba Cloud account.
	//
	// example:
	//
	// 68
	DomainCount *int64 `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	// The total number of accelerated domain names that are in the **Enabled*	- state within the current Alibaba Cloud account.
	//
	// example:
	//
	// 68
	DomainOnlineCount *int64 `json:"DomainOnlineCount,omitempty" xml:"DomainOnlineCount,omitempty"`
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2019-12-23T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4E09C5D7-E1CF-4CAA-A45E-8727F4C8FD70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2019-12-22T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The top N domain names ranked by network traffic.
	TopDomains *DescribeTopDomainsByFlowResponseBodyTopDomains `json:"TopDomains,omitempty" xml:"TopDomains,omitempty" type:"Struct"`
}

func (s DescribeTopDomainsByFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopDomainsByFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTopDomainsByFlowResponseBody) GetDomainCount() *int64 {
	return s.DomainCount
}

func (s *DescribeTopDomainsByFlowResponseBody) GetDomainOnlineCount() *int64 {
	return s.DomainOnlineCount
}

func (s *DescribeTopDomainsByFlowResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTopDomainsByFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTopDomainsByFlowResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeTopDomainsByFlowResponseBody) GetTopDomains() *DescribeTopDomainsByFlowResponseBodyTopDomains {
	return s.TopDomains
}

func (s *DescribeTopDomainsByFlowResponseBody) SetDomainCount(v int64) *DescribeTopDomainsByFlowResponseBody {
	s.DomainCount = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBody) SetDomainOnlineCount(v int64) *DescribeTopDomainsByFlowResponseBody {
	s.DomainOnlineCount = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBody) SetEndTime(v string) *DescribeTopDomainsByFlowResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBody) SetRequestId(v string) *DescribeTopDomainsByFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBody) SetStartTime(v string) *DescribeTopDomainsByFlowResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBody) SetTopDomains(v *DescribeTopDomainsByFlowResponseBodyTopDomains) *DescribeTopDomainsByFlowResponseBody {
	s.TopDomains = v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTopDomainsByFlowResponseBodyTopDomains struct {
	TopDomain []*DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain `json:"TopDomain,omitempty" xml:"TopDomain,omitempty" type:"Repeated"`
}

func (s DescribeTopDomainsByFlowResponseBodyTopDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopDomainsByFlowResponseBodyTopDomains) GoString() string {
	return s.String()
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomains) GetTopDomain() []*DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	return s.TopDomain
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomains) SetTopDomain(v []*DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) *DescribeTopDomainsByFlowResponseBodyTopDomains {
	s.TopDomain = v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain struct {
	// The accelerated domain name.
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
	MaxBps *float32 `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	// The time when the bandwidth reached the peak value.
	//
	// example:
	//
	// 1457111400
	MaxBpsTime *string `json:"MaxBpsTime,omitempty" xml:"MaxBpsTime,omitempty"`
	// The ranking of the accelerated domain name.
	//
	// example:
	//
	// 1
	Rank *int64 `json:"Rank,omitempty" xml:"Rank,omitempty"`
	// The number of visits to the domain name.
	//
	// example:
	//
	// 107784230
	TotalAccess *int64 `json:"TotalAccess,omitempty" xml:"TotalAccess,omitempty"`
	// The total volume of traffic.
	//
	// example:
	//
	// 2043859876683.9001
	TotalTraffic *string `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
	// The proportion of network traffic consumed to access the domain name.
	//
	// example:
	//
	// 30.64191989360235
	TrafficPercent *string `json:"TrafficPercent,omitempty" xml:"TrafficPercent,omitempty"`
}

func (s DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) GoString() string {
	return s.String()
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetMaxBps() *float32 {
	return s.MaxBps
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetMaxBpsTime() *string {
	return s.MaxBpsTime
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetRank() *int64 {
	return s.Rank
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetTotalAccess() *int64 {
	return s.TotalAccess
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetTotalTraffic() *string {
	return s.TotalTraffic
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetTrafficPercent() *string {
	return s.TrafficPercent
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetDomainName(v string) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.DomainName = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBps(v float32) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBps = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBpsTime(v string) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBpsTime = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetRank(v int64) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.Rank = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalAccess(v int64) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalAccess = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalTraffic(v string) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTrafficPercent(v string) *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TrafficPercent = &v
	return s
}

func (s *DescribeTopDomainsByFlowResponseBodyTopDomainsTopDomain) Validate() error {
	return dara.Validate(s)
}
