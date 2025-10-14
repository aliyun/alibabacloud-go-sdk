// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDohSubDomainStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDohSubDomainStatisticsResponseBody
	GetRequestId() *string
	SetStatistics(v []*DescribeDohSubDomainStatisticsResponseBodyStatistics) *DescribeDohSubDomainStatisticsResponseBody
	GetStatistics() []*DescribeDohSubDomainStatisticsResponseBodyStatistics
}

type DescribeDohSubDomainStatisticsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0F32959D-417B-4D66-8463-68606605E3E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics list.
	Statistics []*DescribeDohSubDomainStatisticsResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s DescribeDohSubDomainStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohSubDomainStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDohSubDomainStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDohSubDomainStatisticsResponseBody) GetStatistics() []*DescribeDohSubDomainStatisticsResponseBodyStatistics {
	return s.Statistics
}

func (s *DescribeDohSubDomainStatisticsResponseBody) SetRequestId(v string) *DescribeDohSubDomainStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsResponseBody) SetStatistics(v []*DescribeDohSubDomainStatisticsResponseBodyStatistics) *DescribeDohSubDomainStatisticsResponseBody {
	s.Statistics = v
	return s
}

func (s *DescribeDohSubDomainStatisticsResponseBody) Validate() error {
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

type DescribeDohSubDomainStatisticsResponseBodyStatistics struct {
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

func (s DescribeDohSubDomainStatisticsResponseBodyStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohSubDomainStatisticsResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDohSubDomainStatisticsResponseBodyStatistics) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeDohSubDomainStatisticsResponseBodyStatistics) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDohSubDomainStatisticsResponseBodyStatistics) GetV4HttpCount() *int64 {
	return s.V4HttpCount
}

func (s *DescribeDohSubDomainStatisticsResponseBodyStatistics) GetV4HttpsCount() *int64 {
	return s.V4HttpsCount
}

func (s *DescribeDohSubDomainStatisticsResponseBodyStatistics) GetV6HttpCount() *int64 {
	return s.V6HttpCount
}

func (s *DescribeDohSubDomainStatisticsResponseBodyStatistics) GetV6HttpsCount() *int64 {
	return s.V6HttpsCount
}

func (s *DescribeDohSubDomainStatisticsResponseBodyStatistics) SetTimestamp(v int64) *DescribeDohSubDomainStatisticsResponseBodyStatistics {
	s.Timestamp = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsResponseBodyStatistics) SetTotalCount(v int64) *DescribeDohSubDomainStatisticsResponseBodyStatistics {
	s.TotalCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsResponseBodyStatistics) SetV4HttpCount(v int64) *DescribeDohSubDomainStatisticsResponseBodyStatistics {
	s.V4HttpCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsResponseBodyStatistics) SetV4HttpsCount(v int64) *DescribeDohSubDomainStatisticsResponseBodyStatistics {
	s.V4HttpsCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsResponseBodyStatistics) SetV6HttpCount(v int64) *DescribeDohSubDomainStatisticsResponseBodyStatistics {
	s.V6HttpCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsResponseBodyStatistics) SetV6HttpsCount(v int64) *DescribeDohSubDomainStatisticsResponseBodyStatistics {
	s.V6HttpsCount = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsResponseBodyStatistics) Validate() error {
	return dara.Validate(s)
}
