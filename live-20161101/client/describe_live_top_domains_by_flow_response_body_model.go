// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveTopDomainsByFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainCount(v int64) *DescribeLiveTopDomainsByFlowResponseBody
	GetDomainCount() *int64
	SetDomainOnlineCount(v int64) *DescribeLiveTopDomainsByFlowResponseBody
	GetDomainOnlineCount() *int64
	SetEndTime(v string) *DescribeLiveTopDomainsByFlowResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeLiveTopDomainsByFlowResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLiveTopDomainsByFlowResponseBody
	GetStartTime() *string
	SetTopDomains(v *DescribeLiveTopDomainsByFlowResponseBodyTopDomains) *DescribeLiveTopDomainsByFlowResponseBody
	GetTopDomains() *DescribeLiveTopDomainsByFlowResponseBodyTopDomains
}

type DescribeLiveTopDomainsByFlowResponseBody struct {
	// The total number of domain names in your account.
	//
	// example:
	//
	// 1
	DomainCount *int64 `json:"DomainCount,omitempty" xml:"DomainCount,omitempty"`
	// The total number of domain names that are in the Enabled state in your account.
	//
	// example:
	//
	// 1
	DomainOnlineCount *int64 `json:"DomainOnlineCount,omitempty" xml:"DomainOnlineCount,omitempty"`
	// The end of the time range for which data was queried.
	//
	// example:
	//
	// 2018-03-20T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 33834C3E-388E-5FFE-A8AE-63575035C064
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range for which data was queried.
	//
	// example:
	//
	// 2018-03-17T16:00:00Z
	StartTime  *string                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TopDomains *DescribeLiveTopDomainsByFlowResponseBodyTopDomains `json:"TopDomains,omitempty" xml:"TopDomains,omitempty" type:"Struct"`
}

func (s DescribeLiveTopDomainsByFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveTopDomainsByFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveTopDomainsByFlowResponseBody) GetDomainCount() *int64 {
	return s.DomainCount
}

func (s *DescribeLiveTopDomainsByFlowResponseBody) GetDomainOnlineCount() *int64 {
	return s.DomainOnlineCount
}

func (s *DescribeLiveTopDomainsByFlowResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveTopDomainsByFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveTopDomainsByFlowResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveTopDomainsByFlowResponseBody) GetTopDomains() *DescribeLiveTopDomainsByFlowResponseBodyTopDomains {
	return s.TopDomains
}

func (s *DescribeLiveTopDomainsByFlowResponseBody) SetDomainCount(v int64) *DescribeLiveTopDomainsByFlowResponseBody {
	s.DomainCount = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseBody) SetDomainOnlineCount(v int64) *DescribeLiveTopDomainsByFlowResponseBody {
	s.DomainOnlineCount = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseBody) SetEndTime(v string) *DescribeLiveTopDomainsByFlowResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseBody) SetRequestId(v string) *DescribeLiveTopDomainsByFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseBody) SetStartTime(v string) *DescribeLiveTopDomainsByFlowResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseBody) SetTopDomains(v *DescribeLiveTopDomainsByFlowResponseBodyTopDomains) *DescribeLiveTopDomainsByFlowResponseBody {
	s.TopDomains = v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseBody) Validate() error {
	if s.TopDomains != nil {
		if err := s.TopDomains.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveTopDomainsByFlowResponseBodyTopDomains struct {
	TopDomain []*DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain `json:"TopDomain,omitempty" xml:"TopDomain,omitempty" type:"Repeated"`
}

func (s DescribeLiveTopDomainsByFlowResponseBodyTopDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveTopDomainsByFlowResponseBodyTopDomains) GoString() string {
	return s.String()
}

func (s *DescribeLiveTopDomainsByFlowResponseBodyTopDomains) GetTopDomain() []*DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	return s.TopDomain
}

func (s *DescribeLiveTopDomainsByFlowResponseBodyTopDomains) SetTopDomain(v []*DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain) *DescribeLiveTopDomainsByFlowResponseBodyTopDomains {
	s.TopDomain = v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseBodyTopDomains) Validate() error {
	if s.TopDomain != nil {
		for _, item := range s.TopDomain {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	MaxBps         *int64  `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	MaxBpsTime     *string `json:"MaxBpsTime,omitempty" xml:"MaxBpsTime,omitempty"`
	Rank           *int64  `json:"Rank,omitempty" xml:"Rank,omitempty"`
	TotalAccess    *int64  `json:"TotalAccess,omitempty" xml:"TotalAccess,omitempty"`
	TotalTraffic   *string `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
	TrafficPercent *string `json:"TrafficPercent,omitempty" xml:"TrafficPercent,omitempty"`
}

func (s DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain) GoString() string {
	return s.String()
}

func (s *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetMaxBps() *int64 {
	return s.MaxBps
}

func (s *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetMaxBpsTime() *string {
	return s.MaxBpsTime
}

func (s *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetRank() *int64 {
	return s.Rank
}

func (s *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetTotalAccess() *int64 {
	return s.TotalAccess
}

func (s *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetTotalTraffic() *string {
	return s.TotalTraffic
}

func (s *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain) GetTrafficPercent() *string {
	return s.TrafficPercent
}

func (s *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetDomainName(v string) *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBps(v int64) *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBps = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetMaxBpsTime(v string) *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.MaxBpsTime = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetRank(v int64) *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.Rank = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalAccess(v int64) *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalAccess = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTotalTraffic(v string) *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain) SetTrafficPercent(v string) *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain {
	s.TrafficPercent = &v
	return s
}

func (s *DescribeLiveTopDomainsByFlowResponseBodyTopDomainsTopDomain) Validate() error {
	return dara.Validate(s)
}
