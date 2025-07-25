// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDomainStatisticsResponseBody
	GetRequestId() *string
	SetStatistics(v *DescribeDomainStatisticsResponseBodyStatistics) *DescribeDomainStatisticsResponseBody
	GetStatistics() *DescribeDomainStatisticsResponseBodyStatistics
}

type DescribeDomainStatisticsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6AEC7A64-3CB1-4C49-8B35-0B901F1E26BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics on the Domain Name System (DNS) requests.
	Statistics *DescribeDomainStatisticsResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
}

func (s DescribeDomainStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainStatisticsResponseBody) GetStatistics() *DescribeDomainStatisticsResponseBodyStatistics {
	return s.Statistics
}

func (s *DescribeDomainStatisticsResponseBody) SetRequestId(v string) *DescribeDomainStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainStatisticsResponseBody) SetStatistics(v *DescribeDomainStatisticsResponseBodyStatistics) *DescribeDomainStatisticsResponseBody {
	s.Statistics = v
	return s
}

func (s *DescribeDomainStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainStatisticsResponseBodyStatistics struct {
	Statistic []*DescribeDomainStatisticsResponseBodyStatisticsStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Repeated"`
}

func (s DescribeDomainStatisticsResponseBodyStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainStatisticsResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatisticsResponseBodyStatistics) GetStatistic() []*DescribeDomainStatisticsResponseBodyStatisticsStatistic {
	return s.Statistic
}

func (s *DescribeDomainStatisticsResponseBodyStatistics) SetStatistic(v []*DescribeDomainStatisticsResponseBodyStatisticsStatistic) *DescribeDomainStatisticsResponseBodyStatistics {
	s.Statistic = v
	return s
}

func (s *DescribeDomainStatisticsResponseBodyStatistics) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainStatisticsResponseBodyStatisticsStatistic struct {
	// The number of DNS requests.
	//
	// example:
	//
	// 15292887
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The statistical timestamp. Unit: milliseconds. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1556640000000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeDomainStatisticsResponseBodyStatisticsStatistic) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainStatisticsResponseBodyStatisticsStatistic) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatisticsResponseBodyStatisticsStatistic) GetCount() *int64 {
	return s.Count
}

func (s *DescribeDomainStatisticsResponseBodyStatisticsStatistic) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainStatisticsResponseBodyStatisticsStatistic) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeDomainStatisticsResponseBodyStatisticsStatistic) SetCount(v int64) *DescribeDomainStatisticsResponseBodyStatisticsStatistic {
	s.Count = &v
	return s
}

func (s *DescribeDomainStatisticsResponseBodyStatisticsStatistic) SetDomainName(v string) *DescribeDomainStatisticsResponseBodyStatisticsStatistic {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainStatisticsResponseBodyStatisticsStatistic) SetTimestamp(v int64) *DescribeDomainStatisticsResponseBodyStatisticsStatistic {
	s.Timestamp = &v
	return s
}

func (s *DescribeDomainStatisticsResponseBodyStatisticsStatistic) Validate() error {
	return dara.Validate(s)
}
