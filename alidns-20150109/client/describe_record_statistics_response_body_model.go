// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRecordStatisticsResponseBody
	GetRequestId() *string
	SetStatistics(v *DescribeRecordStatisticsResponseBodyStatistics) *DescribeRecordStatisticsResponseBody
	GetStatistics() *DescribeRecordStatisticsResponseBodyStatistics
}

type DescribeRecordStatisticsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6AEC7A64-3CB1-4C49-8B35-0B901F1E26BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics on the DNS requests.
	Statistics *DescribeRecordStatisticsResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
}

func (s DescribeRecordStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecordStatisticsResponseBody) GetStatistics() *DescribeRecordStatisticsResponseBodyStatistics {
	return s.Statistics
}

func (s *DescribeRecordStatisticsResponseBody) SetRequestId(v string) *DescribeRecordStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordStatisticsResponseBody) SetStatistics(v *DescribeRecordStatisticsResponseBodyStatistics) *DescribeRecordStatisticsResponseBody {
	s.Statistics = v
	return s
}

func (s *DescribeRecordStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRecordStatisticsResponseBodyStatistics struct {
	Statistic []*DescribeRecordStatisticsResponseBodyStatisticsStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
}

func (s DescribeRecordStatisticsResponseBodyStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordStatisticsResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeRecordStatisticsResponseBodyStatistics) GetStatistic() []*DescribeRecordStatisticsResponseBodyStatisticsStatistic {
	return s.Statistic
}

func (s *DescribeRecordStatisticsResponseBodyStatistics) SetStatistic(v []*DescribeRecordStatisticsResponseBodyStatisticsStatistic) *DescribeRecordStatisticsResponseBodyStatistics {
	s.Statistic = v
	return s
}

func (s *DescribeRecordStatisticsResponseBodyStatistics) Validate() error {
	return dara.Validate(s)
}

type DescribeRecordStatisticsResponseBodyStatisticsStatistic struct {
	// The number of DNS requests.
	//
	// example:
	//
	// 15292887
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The statistical timestamp. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1556640000000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeRecordStatisticsResponseBodyStatisticsStatistic) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordStatisticsResponseBodyStatisticsStatistic) GoString() string {
	return s.String()
}

func (s *DescribeRecordStatisticsResponseBodyStatisticsStatistic) GetCount() *int64 {
	return s.Count
}

func (s *DescribeRecordStatisticsResponseBodyStatisticsStatistic) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeRecordStatisticsResponseBodyStatisticsStatistic) SetCount(v int64) *DescribeRecordStatisticsResponseBodyStatisticsStatistic {
	s.Count = &v
	return s
}

func (s *DescribeRecordStatisticsResponseBodyStatisticsStatistic) SetTimestamp(v int64) *DescribeRecordStatisticsResponseBodyStatisticsStatistic {
	s.Timestamp = &v
	return s
}

func (s *DescribeRecordStatisticsResponseBodyStatisticsStatistic) Validate() error {
	return dara.Validate(s)
}
