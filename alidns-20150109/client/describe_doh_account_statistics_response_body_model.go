// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDohAccountStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDohAccountStatisticsResponseBody
	GetRequestId() *string
	SetStatistics(v []*DescribeDohAccountStatisticsResponseBodyStatistics) *DescribeDohAccountStatisticsResponseBody
	GetStatistics() []*DescribeDohAccountStatisticsResponseBodyStatistics
}

type DescribeDohAccountStatisticsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0F32959D-417B-4D66-8463-68606605E3E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics list.
	Statistics []*DescribeDohAccountStatisticsResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s DescribeDohAccountStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohAccountStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDohAccountStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDohAccountStatisticsResponseBody) GetStatistics() []*DescribeDohAccountStatisticsResponseBodyStatistics {
	return s.Statistics
}

func (s *DescribeDohAccountStatisticsResponseBody) SetRequestId(v string) *DescribeDohAccountStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDohAccountStatisticsResponseBody) SetStatistics(v []*DescribeDohAccountStatisticsResponseBodyStatistics) *DescribeDohAccountStatisticsResponseBody {
	s.Statistics = v
	return s
}

func (s *DescribeDohAccountStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDohAccountStatisticsResponseBodyStatistics struct {
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

func (s DescribeDohAccountStatisticsResponseBodyStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohAccountStatisticsResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDohAccountStatisticsResponseBodyStatistics) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeDohAccountStatisticsResponseBodyStatistics) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDohAccountStatisticsResponseBodyStatistics) GetV4HttpCount() *int64 {
	return s.V4HttpCount
}

func (s *DescribeDohAccountStatisticsResponseBodyStatistics) GetV4HttpsCount() *int64 {
	return s.V4HttpsCount
}

func (s *DescribeDohAccountStatisticsResponseBodyStatistics) GetV6HttpCount() *int64 {
	return s.V6HttpCount
}

func (s *DescribeDohAccountStatisticsResponseBodyStatistics) GetV6HttpsCount() *int64 {
	return s.V6HttpsCount
}

func (s *DescribeDohAccountStatisticsResponseBodyStatistics) SetTimestamp(v int64) *DescribeDohAccountStatisticsResponseBodyStatistics {
	s.Timestamp = &v
	return s
}

func (s *DescribeDohAccountStatisticsResponseBodyStatistics) SetTotalCount(v int64) *DescribeDohAccountStatisticsResponseBodyStatistics {
	s.TotalCount = &v
	return s
}

func (s *DescribeDohAccountStatisticsResponseBodyStatistics) SetV4HttpCount(v int64) *DescribeDohAccountStatisticsResponseBodyStatistics {
	s.V4HttpCount = &v
	return s
}

func (s *DescribeDohAccountStatisticsResponseBodyStatistics) SetV4HttpsCount(v int64) *DescribeDohAccountStatisticsResponseBodyStatistics {
	s.V4HttpsCount = &v
	return s
}

func (s *DescribeDohAccountStatisticsResponseBodyStatistics) SetV6HttpCount(v int64) *DescribeDohAccountStatisticsResponseBodyStatistics {
	s.V6HttpCount = &v
	return s
}

func (s *DescribeDohAccountStatisticsResponseBodyStatistics) SetV6HttpsCount(v int64) *DescribeDohAccountStatisticsResponseBodyStatistics {
	s.V6HttpsCount = &v
	return s
}

func (s *DescribeDohAccountStatisticsResponseBodyStatistics) Validate() error {
	return dara.Validate(s)
}
