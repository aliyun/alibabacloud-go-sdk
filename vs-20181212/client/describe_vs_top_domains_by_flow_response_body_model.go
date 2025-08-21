// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsTopDomainsByFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainCount(v int64) *DescribeVsTopDomainsByFlowResponseBody
	GetDomainCount() *int64
	SetDomainOnlineCount(v int64) *DescribeVsTopDomainsByFlowResponseBody
	GetDomainOnlineCount() *int64
	SetEndTime(v string) *DescribeVsTopDomainsByFlowResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeVsTopDomainsByFlowResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVsTopDomainsByFlowResponseBody
	GetStartTime() *string
	SetTopDomains(v *DescribeVsTopDomainsByFlowResponseBodyTopDomains) *DescribeVsTopDomainsByFlowResponseBody
	GetTopDomains() *DescribeVsTopDomainsByFlowResponseBodyTopDomains
}

type DescribeVsTopDomainsByFlowResponseBody struct {
	// example:
	//
	// 20
	DomainCount *int64 `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	// example:
	//
	// 12
	DomainOnlineCount *int64 `json:"DomainOnlineCount,omitempty" xml:"DomainOnlineCount,omitempty"`
	// example:
	//
	// 2018-12-10T18:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2021-12-10T10:00:00Z
	StartTime  *string                                           `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TopDomains *DescribeVsTopDomainsByFlowResponseBodyTopDomains `json:"TopDomains,omitempty" xml:"TopDomains,omitempty" type:"Struct"`
}

func (s DescribeVsTopDomainsByFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsTopDomainsByFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsTopDomainsByFlowResponseBody) GetDomainCount() *int64 {
	return s.DomainCount
}

func (s *DescribeVsTopDomainsByFlowResponseBody) GetDomainOnlineCount() *int64 {
	return s.DomainOnlineCount
}

func (s *DescribeVsTopDomainsByFlowResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsTopDomainsByFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsTopDomainsByFlowResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsTopDomainsByFlowResponseBody) GetTopDomains() *DescribeVsTopDomainsByFlowResponseBodyTopDomains {
	return s.TopDomains
}

func (s *DescribeVsTopDomainsByFlowResponseBody) SetDomainCount(v int64) *DescribeVsTopDomainsByFlowResponseBody {
	s.DomainCount = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBody) SetDomainOnlineCount(v int64) *DescribeVsTopDomainsByFlowResponseBody {
	s.DomainOnlineCount = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBody) SetEndTime(v string) *DescribeVsTopDomainsByFlowResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBody) SetRequestId(v string) *DescribeVsTopDomainsByFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBody) SetStartTime(v string) *DescribeVsTopDomainsByFlowResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBody) SetTopDomains(v *DescribeVsTopDomainsByFlowResponseBodyTopDomains) *DescribeVsTopDomainsByFlowResponseBody {
	s.TopDomains = v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVsTopDomainsByFlowResponseBodyTopDomains struct {
	TopDomain []*DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain `json:"TopDomain,omitempty" xml:"TopDomain,omitempty" type:"Repeated"`
}

func (s DescribeVsTopDomainsByFlowResponseBodyTopDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsTopDomainsByFlowResponseBodyTopDomains) GoString() string {
	return s.String()
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomains) GetTopDomain() []*DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	return s.TopDomain
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomains) SetTopDomain(v []*DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) *DescribeVsTopDomainsByFlowResponseBodyTopDomains {
	s.TopDomain = v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain struct {
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 100
	MaxBps *int64 `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	// example:
	//
	// 1457111400
	MaxBpsTime *string `json:"MaxBpsTime,omitempty" xml:"MaxBpsTime,omitempty"`
	// example:
	//
	// 1
	Rank *int64 `json:"Rank,omitempty" xml:"Rank,omitempty"`
	// example:
	//
	// 100
	TotalAccess *int64 `json:"TotalAccess,omitempty" xml:"TotalAccess,omitempty"`
	// example:
	//
	// 100
	TotalTraffic *string `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
	// example:
	//
	// 30.64191989360235
	TrafficPercent *string `json:"TrafficPercent,omitempty" xml:"TrafficPercent,omitempty"`
}

func (s DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) GoString() string {
	return s.String()
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetMaxBps() *int64 {
	return s.MaxBps
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetMaxBpsTime() *string {
	return s.MaxBpsTime
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetRank() *int64 {
	return s.Rank
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetTotalAccess() *int64 {
	return s.TotalAccess
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetTotalTraffic() *string {
	return s.TotalTraffic
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetTrafficPercent() *string {
	return s.TrafficPercent
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetDomainName(v string) *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.DomainName = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBps(v int64) *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBps = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBpsTime(v string) *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBpsTime = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetRank(v int64) *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.Rank = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalAccess(v int64) *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalAccess = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalTraffic(v string) *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTrafficPercent(v string) *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TrafficPercent = &v
	return s
}

func (s *DescribeVsTopDomainsByFlowResponseBodyTopDomainsTopDomain) Validate() error {
	return dara.Validate(s)
}
