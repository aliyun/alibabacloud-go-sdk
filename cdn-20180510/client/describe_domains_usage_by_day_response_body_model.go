// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainsUsageByDayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainsUsageByDayResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainsUsageByDayResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainsUsageByDayResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDomainsUsageByDayResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainsUsageByDayResponseBody
	GetStartTime() *string
	SetUsageByDays(v *DescribeDomainsUsageByDayResponseBodyUsageByDays) *DescribeDomainsUsageByDayResponseBody
	GetUsageByDays() *DescribeDomainsUsageByDayResponseBodyUsageByDays
	SetUsageTotal(v *DescribeDomainsUsageByDayResponseBodyUsageTotal) *DescribeDomainsUsageByDayResponseBody
	GetUsageTotal() *DescribeDomainsUsageByDayResponseBodyUsageTotal
}

type DescribeDomainsUsageByDayResponseBody struct {
	// The time interval between the data entries returned. Unit: seconds.
	//
	// example:
	//
	// 86400
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2019-12-23T09:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C88EF8ED-72F0-45EA-9E86-95114E224FC5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2019-12-22T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The monitoring data collected at each time interval.
	UsageByDays *DescribeDomainsUsageByDayResponseBodyUsageByDays `json:"UsageByDays,omitempty" xml:"UsageByDays,omitempty" type:"Struct"`
	// The summarized monitoring data.
	UsageTotal *DescribeDomainsUsageByDayResponseBodyUsageTotal `json:"UsageTotal,omitempty" xml:"UsageTotal,omitempty" type:"Struct"`
}

func (s DescribeDomainsUsageByDayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsUsageByDayResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainsUsageByDayResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainsUsageByDayResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainsUsageByDayResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainsUsageByDayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainsUsageByDayResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainsUsageByDayResponseBody) GetUsageByDays() *DescribeDomainsUsageByDayResponseBodyUsageByDays {
	return s.UsageByDays
}

func (s *DescribeDomainsUsageByDayResponseBody) GetUsageTotal() *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	return s.UsageTotal
}

func (s *DescribeDomainsUsageByDayResponseBody) SetDataInterval(v string) *DescribeDomainsUsageByDayResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBody) SetDomainName(v string) *DescribeDomainsUsageByDayResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBody) SetEndTime(v string) *DescribeDomainsUsageByDayResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBody) SetRequestId(v string) *DescribeDomainsUsageByDayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBody) SetStartTime(v string) *DescribeDomainsUsageByDayResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBody) SetUsageByDays(v *DescribeDomainsUsageByDayResponseBodyUsageByDays) *DescribeDomainsUsageByDayResponseBody {
	s.UsageByDays = v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBody) SetUsageTotal(v *DescribeDomainsUsageByDayResponseBodyUsageTotal) *DescribeDomainsUsageByDayResponseBody {
	s.UsageTotal = v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainsUsageByDayResponseBodyUsageByDays struct {
	UsageByDay []*DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay `json:"UsageByDay,omitempty" xml:"UsageByDay,omitempty" type:"Repeated"`
}

func (s DescribeDomainsUsageByDayResponseBodyUsageByDays) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsUsageByDayResponseBodyUsageByDays) GoString() string {
	return s.String()
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDays) GetUsageByDay() []*DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	return s.UsageByDay
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDays) SetUsageByDay(v []*DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) *DescribeDomainsUsageByDayResponseBodyUsageByDays {
	s.UsageByDay = v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDays) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay struct {
	// The byte hit ratio. The byte hit ratio is measured in percentage.
	//
	// example:
	//
	// 97.46250599529726
	BytesHitRate *string `json:"BytesHitRate,omitempty" xml:"BytesHitRate,omitempty"`
	// The peak bandwidth value. Unit: bit/s.
	//
	// example:
	//
	// 306747.76
	MaxBps *string `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	// The time when the bandwidth reached the peak value.
	//
	// example:
	//
	// 2019-12-23 10:55:00
	MaxBpsTime *string `json:"MaxBpsTime,omitempty" xml:"MaxBpsTime,omitempty"`
	// The peak bandwidth value during back-to-origin routing. Unit: bit/s.
	//
	// example:
	//
	// 72584.072
	MaxSrcBps *string `json:"MaxSrcBps,omitempty" xml:"MaxSrcBps,omitempty"`
	// The time when the bandwidth during back-to-origin routing reached the peak value.
	//
	// example:
	//
	// 2019-12-23 11:45:00
	MaxSrcBpsTime *string `json:"MaxSrcBpsTime,omitempty" xml:"MaxSrcBpsTime,omitempty"`
	// The number of queries per second (QPS).
	//
	// example:
	//
	// 7.466354166666667
	Qps *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The cache hit ratio that is calculated based on requests. The cache hit ratio is measured in percentage.
	//
	// example:
	//
	// 70.24770071912111
	RequestHitRate *string `json:"RequestHitRate,omitempty" xml:"RequestHitRate,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2019-12-22
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The total amount of requests.
	//
	// example:
	//
	// 645093
	TotalAccess *string `json:"TotalAccess,omitempty" xml:"TotalAccess,omitempty"`
	// The total amount of network traffic. Unit: bytes.
	//
	// example:
	//
	// 564300099309
	TotalTraffic *string `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
}

func (s DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GoString() string {
	return s.String()
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetBytesHitRate() *string {
	return s.BytesHitRate
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetMaxBps() *string {
	return s.MaxBps
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetMaxBpsTime() *string {
	return s.MaxBpsTime
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetMaxSrcBps() *string {
	return s.MaxSrcBps
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetMaxSrcBpsTime() *string {
	return s.MaxSrcBpsTime
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetQps() *string {
	return s.Qps
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetRequestHitRate() *string {
	return s.RequestHitRate
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetTotalAccess() *string {
	return s.TotalAccess
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetTotalTraffic() *string {
	return s.TotalTraffic
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetBytesHitRate(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.BytesHitRate = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetMaxBps(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.MaxBps = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetMaxBpsTime(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.MaxBpsTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetMaxSrcBps(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.MaxSrcBps = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetMaxSrcBpsTime(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.MaxSrcBpsTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetQps(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.Qps = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetRequestHitRate(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.RequestHitRate = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetTimeStamp(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetTotalAccess(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.TotalAccess = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetTotalTraffic(v string) *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainsUsageByDayResponseBodyUsageTotal struct {
	// The byte hit ratio. The byte hit ratio is measured in percentage.
	//
	// example:
	//
	// 97.03110726801242
	BytesHitRate *string `json:"BytesHitRate,omitempty" xml:"BytesHitRate,omitempty"`
	// The peak bandwidth value. Unit: bit/s.
	//
	// example:
	//
	// 1.0747912780000001E8
	MaxBps *string `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	// The time when the bandwidth reached the peak value.
	//
	// example:
	//
	// 2019-12-23 10:55:00
	MaxBpsTime *string `json:"MaxBpsTime,omitempty" xml:"MaxBpsTime,omitempty"`
	// The peak bandwidth value during back-to-origin routing. Unit: bit/s.
	//
	// example:
	//
	// 72584.072
	MaxSrcBps *string `json:"MaxSrcBps,omitempty" xml:"MaxSrcBps,omitempty"`
	// The time when the bandwidth during back-to-origin routing reached the peak value.
	//
	// example:
	//
	// 2019-12-23 11:45:00
	MaxSrcBpsTime *string `json:"MaxSrcBpsTime,omitempty" xml:"MaxSrcBpsTime,omitempty"`
	// The cache hit ratio that is calculated based on requests. The cache hit ratio is measured in percentage.
	//
	// example:
	//
	// 69.92610837438424
	RequestHitRate *string `json:"RequestHitRate,omitempty" xml:"RequestHitRate,omitempty"`
	// The total amount of requests.
	//
	// example:
	//
	// 1319500
	TotalAccess *string `json:"TotalAccess,omitempty" xml:"TotalAccess,omitempty"`
	// The total amount of network traffic. Unit: bytes.
	//
	// example:
	//
	// 1117711832100
	TotalTraffic *string `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
}

func (s DescribeDomainsUsageByDayResponseBodyUsageTotal) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsUsageByDayResponseBodyUsageTotal) GoString() string {
	return s.String()
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) GetBytesHitRate() *string {
	return s.BytesHitRate
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) GetMaxBps() *string {
	return s.MaxBps
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) GetMaxBpsTime() *string {
	return s.MaxBpsTime
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) GetMaxSrcBps() *string {
	return s.MaxSrcBps
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) GetMaxSrcBpsTime() *string {
	return s.MaxSrcBpsTime
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) GetRequestHitRate() *string {
	return s.RequestHitRate
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) GetTotalAccess() *string {
	return s.TotalAccess
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) GetTotalTraffic() *string {
	return s.TotalTraffic
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetBytesHitRate(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.BytesHitRate = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetMaxBps(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.MaxBps = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetMaxBpsTime(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.MaxBpsTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetMaxSrcBps(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.MaxSrcBps = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetMaxSrcBpsTime(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.MaxSrcBpsTime = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetRequestHitRate(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.RequestHitRate = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetTotalAccess(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.TotalAccess = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) SetTotalTraffic(v string) *DescribeDomainsUsageByDayResponseBodyUsageTotal {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeDomainsUsageByDayResponseBodyUsageTotal) Validate() error {
	return dara.Validate(s)
}
