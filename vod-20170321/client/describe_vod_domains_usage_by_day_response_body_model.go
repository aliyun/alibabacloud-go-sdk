// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainsUsageByDayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVodDomainsUsageByDayResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVodDomainsUsageByDayResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainsUsageByDayResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeVodDomainsUsageByDayResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodDomainsUsageByDayResponseBody
	GetStartTime() *string
	SetUsageByDays(v *DescribeVodDomainsUsageByDayResponseBodyUsageByDays) *DescribeVodDomainsUsageByDayResponseBody
	GetUsageByDays() *DescribeVodDomainsUsageByDayResponseBodyUsageByDays
	SetUsageTotal(v *DescribeVodDomainsUsageByDayResponseBodyUsageTotal) *DescribeVodDomainsUsageByDayResponseBody
	GetUsageTotal() *DescribeVodDomainsUsageByDayResponseBodyUsageTotal
}

type DescribeVodDomainsUsageByDayResponseBody struct {
	DataInterval *string                                              `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName   *string                                              `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime      *string                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId    *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UsageByDays  *DescribeVodDomainsUsageByDayResponseBodyUsageByDays `json:"UsageByDays,omitempty" xml:"UsageByDays,omitempty" type:"Struct"`
	UsageTotal   *DescribeVodDomainsUsageByDayResponseBodyUsageTotal  `json:"UsageTotal,omitempty" xml:"UsageTotal,omitempty" type:"Struct"`
}

func (s DescribeVodDomainsUsageByDayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainsUsageByDayResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainsUsageByDayResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodDomainsUsageByDayResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainsUsageByDayResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainsUsageByDayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainsUsageByDayResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainsUsageByDayResponseBody) GetUsageByDays() *DescribeVodDomainsUsageByDayResponseBodyUsageByDays {
	return s.UsageByDays
}

func (s *DescribeVodDomainsUsageByDayResponseBody) GetUsageTotal() *DescribeVodDomainsUsageByDayResponseBodyUsageTotal {
	return s.UsageTotal
}

func (s *DescribeVodDomainsUsageByDayResponseBody) SetDataInterval(v string) *DescribeVodDomainsUsageByDayResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBody) SetDomainName(v string) *DescribeVodDomainsUsageByDayResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBody) SetEndTime(v string) *DescribeVodDomainsUsageByDayResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBody) SetRequestId(v string) *DescribeVodDomainsUsageByDayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBody) SetStartTime(v string) *DescribeVodDomainsUsageByDayResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBody) SetUsageByDays(v *DescribeVodDomainsUsageByDayResponseBodyUsageByDays) *DescribeVodDomainsUsageByDayResponseBody {
	s.UsageByDays = v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBody) SetUsageTotal(v *DescribeVodDomainsUsageByDayResponseBodyUsageTotal) *DescribeVodDomainsUsageByDayResponseBody {
	s.UsageTotal = v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainsUsageByDayResponseBodyUsageByDays struct {
	UsageByDay []*DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay `json:"UsageByDay,omitempty" xml:"UsageByDay,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainsUsageByDayResponseBodyUsageByDays) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainsUsageByDayResponseBodyUsageByDays) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDays) GetUsageByDay() []*DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	return s.UsageByDay
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDays) SetUsageByDay(v []*DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) *DescribeVodDomainsUsageByDayResponseBodyUsageByDays {
	s.UsageByDay = v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDays) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay struct {
	BytesHitRate   *string `json:"BytesHitRate,omitempty" xml:"BytesHitRate,omitempty"`
	MaxBps         *string `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	MaxBpsTime     *string `json:"MaxBpsTime,omitempty" xml:"MaxBpsTime,omitempty"`
	MaxSrcBps      *string `json:"MaxSrcBps,omitempty" xml:"MaxSrcBps,omitempty"`
	MaxSrcBpsTime  *string `json:"MaxSrcBpsTime,omitempty" xml:"MaxSrcBpsTime,omitempty"`
	Qps            *string `json:"Qps,omitempty" xml:"Qps,omitempty"`
	RequestHitRate *string `json:"RequestHitRate,omitempty" xml:"RequestHitRate,omitempty"`
	TimeStamp      *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	TotalAccess    *string `json:"TotalAccess,omitempty" xml:"TotalAccess,omitempty"`
	TotalTraffic   *string `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
}

func (s DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetBytesHitRate() *string {
	return s.BytesHitRate
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetMaxBps() *string {
	return s.MaxBps
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetMaxBpsTime() *string {
	return s.MaxBpsTime
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetMaxSrcBps() *string {
	return s.MaxSrcBps
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetMaxSrcBpsTime() *string {
	return s.MaxSrcBpsTime
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetQps() *string {
	return s.Qps
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetRequestHitRate() *string {
	return s.RequestHitRate
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetTotalAccess() *string {
	return s.TotalAccess
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) GetTotalTraffic() *string {
	return s.TotalTraffic
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetBytesHitRate(v string) *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.BytesHitRate = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetMaxBps(v string) *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.MaxBps = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetMaxBpsTime(v string) *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.MaxBpsTime = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetMaxSrcBps(v string) *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.MaxSrcBps = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetMaxSrcBpsTime(v string) *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.MaxSrcBpsTime = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetQps(v string) *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.Qps = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetRequestHitRate(v string) *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.RequestHitRate = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetTimeStamp(v string) *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetTotalAccess(v string) *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.TotalAccess = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) SetTotalTraffic(v string) *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageByDaysUsageByDay) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainsUsageByDayResponseBodyUsageTotal struct {
	BytesHitRate   *string `json:"BytesHitRate,omitempty" xml:"BytesHitRate,omitempty"`
	MaxBps         *string `json:"MaxBps,omitempty" xml:"MaxBps,omitempty"`
	MaxBpsTime     *string `json:"MaxBpsTime,omitempty" xml:"MaxBpsTime,omitempty"`
	MaxSrcBps      *string `json:"MaxSrcBps,omitempty" xml:"MaxSrcBps,omitempty"`
	MaxSrcBpsTime  *string `json:"MaxSrcBpsTime,omitempty" xml:"MaxSrcBpsTime,omitempty"`
	RequestHitRate *string `json:"RequestHitRate,omitempty" xml:"RequestHitRate,omitempty"`
	TotalAccess    *string `json:"TotalAccess,omitempty" xml:"TotalAccess,omitempty"`
	TotalTraffic   *string `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
}

func (s DescribeVodDomainsUsageByDayResponseBodyUsageTotal) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainsUsageByDayResponseBodyUsageTotal) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageTotal) GetBytesHitRate() *string {
	return s.BytesHitRate
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageTotal) GetMaxBps() *string {
	return s.MaxBps
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageTotal) GetMaxBpsTime() *string {
	return s.MaxBpsTime
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageTotal) GetMaxSrcBps() *string {
	return s.MaxSrcBps
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageTotal) GetMaxSrcBpsTime() *string {
	return s.MaxSrcBpsTime
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageTotal) GetRequestHitRate() *string {
	return s.RequestHitRate
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageTotal) GetTotalAccess() *string {
	return s.TotalAccess
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageTotal) GetTotalTraffic() *string {
	return s.TotalTraffic
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageTotal) SetBytesHitRate(v string) *DescribeVodDomainsUsageByDayResponseBodyUsageTotal {
	s.BytesHitRate = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageTotal) SetMaxBps(v string) *DescribeVodDomainsUsageByDayResponseBodyUsageTotal {
	s.MaxBps = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageTotal) SetMaxBpsTime(v string) *DescribeVodDomainsUsageByDayResponseBodyUsageTotal {
	s.MaxBpsTime = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageTotal) SetMaxSrcBps(v string) *DescribeVodDomainsUsageByDayResponseBodyUsageTotal {
	s.MaxSrcBps = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageTotal) SetMaxSrcBpsTime(v string) *DescribeVodDomainsUsageByDayResponseBodyUsageTotal {
	s.MaxSrcBpsTime = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageTotal) SetRequestHitRate(v string) *DescribeVodDomainsUsageByDayResponseBodyUsageTotal {
	s.RequestHitRate = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageTotal) SetTotalAccess(v string) *DescribeVodDomainsUsageByDayResponseBodyUsageTotal {
	s.TotalAccess = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageTotal) SetTotalTraffic(v string) *DescribeVodDomainsUsageByDayResponseBodyUsageTotal {
	s.TotalTraffic = &v
	return s
}

func (s *DescribeVodDomainsUsageByDayResponseBodyUsageTotal) Validate() error {
	return dara.Validate(s)
}
