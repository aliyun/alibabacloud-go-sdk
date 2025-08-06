// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodTopDomainsByFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainCount(v int64) *DescribeVodTopDomainsByFlowResponseBody
	GetDomainCount() *int64
	SetDomainOnlineCount(v int64) *DescribeVodTopDomainsByFlowResponseBody
	GetDomainOnlineCount() *int64
	SetEndTime(v string) *DescribeVodTopDomainsByFlowResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeVodTopDomainsByFlowResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodTopDomainsByFlowResponseBody
	GetStartTime() *string
	SetTopDomains(v *DescribeVodTopDomainsByFlowResponseBodyTopDomains) *DescribeVodTopDomainsByFlowResponseBody
	GetTopDomains() *DescribeVodTopDomainsByFlowResponseBodyTopDomains
}

type DescribeVodTopDomainsByFlowResponseBody struct {
	DomainCount       *int64                                             `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	DomainOnlineCount *int64                                             `json:"DomainOnlineCount,omitempty" xml:"DomainOnlineCount,omitempty"`
	EndTime           *string                                            `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId         *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime         *string                                            `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TopDomains        *DescribeVodTopDomainsByFlowResponseBodyTopDomains `json:"TopDomains,omitempty" xml:"TopDomains,omitempty" type:"Struct"`
}

func (s DescribeVodTopDomainsByFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTopDomainsByFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodTopDomainsByFlowResponseBody) GetDomainCount() *int64 {
	return s.DomainCount
}

func (s *DescribeVodTopDomainsByFlowResponseBody) GetDomainOnlineCount() *int64 {
	return s.DomainOnlineCount
}

func (s *DescribeVodTopDomainsByFlowResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodTopDomainsByFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodTopDomainsByFlowResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodTopDomainsByFlowResponseBody) GetTopDomains() *DescribeVodTopDomainsByFlowResponseBodyTopDomains {
	return s.TopDomains
}

func (s *DescribeVodTopDomainsByFlowResponseBody) SetDomainCount(v int64) *DescribeVodTopDomainsByFlowResponseBody {
	s.DomainCount = &v
	return s
}

func (s *DescribeVodTopDomainsByFlowResponseBody) SetDomainOnlineCount(v int64) *DescribeVodTopDomainsByFlowResponseBody {
	s.DomainOnlineCount = &v
	return s
}

func (s *DescribeVodTopDomainsByFlowResponseBody) SetEndTime(v string) *DescribeVodTopDomainsByFlowResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodTopDomainsByFlowResponseBody) SetRequestId(v string) *DescribeVodTopDomainsByFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodTopDomainsByFlowResponseBody) SetStartTime(v string) *DescribeVodTopDomainsByFlowResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodTopDomainsByFlowResponseBody) SetTopDomains(v *DescribeVodTopDomainsByFlowResponseBodyTopDomains) *DescribeVodTopDomainsByFlowResponseBody {
	s.TopDomains = v
	return s
}

func (s *DescribeVodTopDomainsByFlowResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodTopDomainsByFlowResponseBodyTopDomains struct {
	TopDomain []*DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain `json:"TopDomain,omitempty" xml:"TopDomain,omitempty" type:"Repeated"`
}

func (s DescribeVodTopDomainsByFlowResponseBodyTopDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTopDomainsByFlowResponseBodyTopDomains) GoString() string {
	return s.String()
}

func (s *DescribeVodTopDomainsByFlowResponseBodyTopDomains) GetTopDomain() []*DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	return s.TopDomain
}

func (s *DescribeVodTopDomainsByFlowResponseBodyTopDomains) SetTopDomain(v []*DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain) *DescribeVodTopDomainsByFlowResponseBodyTopDomains {
	s.TopDomain = v
	return s
}

func (s *DescribeVodTopDomainsByFlowResponseBodyTopDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	MaxBps         *int64  `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	MaxBpsTime     *string `json:"MaxBpsTime,omitempty" xml:"MaxBpsTime,omitempty"`
	Rank           *int64  `json:"Rank,omitempty" xml:"Rank,omitempty"`
	TotalAccess    *int64  `json:"TotalAccess,omitempty" xml:"TotalAccess,omitempty"`
	TotalTraffic   *string `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
	TrafficPercent *string `json:"TrafficPercent,omitempty" xml:"TrafficPercent,omitempty"`
}

func (s DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain) GoString() string {
	return s.String()
}

func (s *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetMaxBps() *int64 {
	return s.MaxBps
}

func (s *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetMaxBpsTime() *string {
	return s.MaxBpsTime
}

func (s *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetRank() *int64 {
	return s.Rank
}

func (s *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetTotalAccess() *int64 {
	return s.TotalAccess
}

func (s *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetTotalTraffic() *string {
	return s.TotalTraffic
}

func (s *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetTrafficPercent() *string {
	return s.TrafficPercent
}

func (s *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetDomainName(v string) *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.DomainName = &v
	return s
}

func (s *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBps(v int64) *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBps = &v
	return s
}

func (s *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBpsTime(v string) *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBpsTime = &v
	return s
}

func (s *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetRank(v int64) *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.Rank = &v
	return s
}

func (s *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalAccess(v int64) *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalAccess = &v
	return s
}

func (s *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalTraffic(v string) *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTrafficPercent(v string) *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TrafficPercent = &v
	return s
}

func (s *DescribeVodTopDomainsByFlowResponseBodyTopDomainsTopDomain) Validate() error {
	return dara.Validate(s)
}
