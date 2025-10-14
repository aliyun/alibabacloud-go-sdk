// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDohDomainStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDohDomainStatisticsResponseBody
	GetRequestId() *string
	SetStatistics(v []*DescribeDohDomainStatisticsResponseBodyStatistics) *DescribeDohDomainStatisticsResponseBody
	GetStatistics() []*DescribeDohDomainStatisticsResponseBodyStatistics
}

type DescribeDohDomainStatisticsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0F32959D-417B-4D66-8463-68606605E3E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics list.
	Statistics []*DescribeDohDomainStatisticsResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s DescribeDohDomainStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohDomainStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDohDomainStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDohDomainStatisticsResponseBody) GetStatistics() []*DescribeDohDomainStatisticsResponseBodyStatistics {
	return s.Statistics
}

func (s *DescribeDohDomainStatisticsResponseBody) SetRequestId(v string) *DescribeDohDomainStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDohDomainStatisticsResponseBody) SetStatistics(v []*DescribeDohDomainStatisticsResponseBodyStatistics) *DescribeDohDomainStatisticsResponseBody {
	s.Statistics = v
	return s
}

func (s *DescribeDohDomainStatisticsResponseBody) Validate() error {
	if s.Statistics != nil {
		for _, item := range s.Statistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDohDomainStatisticsResponseBodyStatistics struct {
	// The timestamp.
	//
	// example:
	//
	// 1544976000000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The total number of requests.
	//
	// example:
	//
	// 3141592653
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The number of IPv4-based HTTP requests.
	//
	// example:
	//
	// 3141592653
	V4HttpCount *int64 `json:"V4HttpCount,omitempty" xml:"V4HttpCount,omitempty"`
	// The number of IPv4-based HTTPS requests.
	//
	// example:
	//
	// 3141592653
	V4HttpsCount *int64 `json:"V4HttpsCount,omitempty" xml:"V4HttpsCount,omitempty"`
	// The number of IPv6-based HTTP requests.
	//
	// example:
	//
	// 3141592653
	V6HttpCount *int64 `json:"V6HttpCount,omitempty" xml:"V6HttpCount,omitempty"`
	// The number of IPv6-based HTTPS requests.
	//
	// example:
	//
	// 3141592653
	V6HttpsCount *int64 `json:"V6HttpsCount,omitempty" xml:"V6HttpsCount,omitempty"`
}

func (s DescribeDohDomainStatisticsResponseBodyStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohDomainStatisticsResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDohDomainStatisticsResponseBodyStatistics) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeDohDomainStatisticsResponseBodyStatistics) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDohDomainStatisticsResponseBodyStatistics) GetV4HttpCount() *int64 {
	return s.V4HttpCount
}

func (s *DescribeDohDomainStatisticsResponseBodyStatistics) GetV4HttpsCount() *int64 {
	return s.V4HttpsCount
}

func (s *DescribeDohDomainStatisticsResponseBodyStatistics) GetV6HttpCount() *int64 {
	return s.V6HttpCount
}

func (s *DescribeDohDomainStatisticsResponseBodyStatistics) GetV6HttpsCount() *int64 {
	return s.V6HttpsCount
}

func (s *DescribeDohDomainStatisticsResponseBodyStatistics) SetTimestamp(v int64) *DescribeDohDomainStatisticsResponseBodyStatistics {
	s.Timestamp = &v
	return s
}

func (s *DescribeDohDomainStatisticsResponseBodyStatistics) SetTotalCount(v int64) *DescribeDohDomainStatisticsResponseBodyStatistics {
	s.TotalCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsResponseBodyStatistics) SetV4HttpCount(v int64) *DescribeDohDomainStatisticsResponseBodyStatistics {
	s.V4HttpCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsResponseBodyStatistics) SetV4HttpsCount(v int64) *DescribeDohDomainStatisticsResponseBodyStatistics {
	s.V4HttpsCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsResponseBodyStatistics) SetV6HttpCount(v int64) *DescribeDohDomainStatisticsResponseBodyStatistics {
	s.V6HttpCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsResponseBodyStatistics) SetV6HttpsCount(v int64) *DescribeDohDomainStatisticsResponseBodyStatistics {
	s.V6HttpsCount = &v
	return s
}

func (s *DescribeDohDomainStatisticsResponseBodyStatistics) Validate() error {
	return dara.Validate(s)
}
