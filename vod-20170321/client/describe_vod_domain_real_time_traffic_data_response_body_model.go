// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVodDomainRealTimeTrafficDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVodDomainRealTimeTrafficDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainRealTimeTrafficDataResponseBody
	GetEndTime() *string
	SetRealTimeTrafficDataPerInterval(v *DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) *DescribeVodDomainRealTimeTrafficDataResponseBody
	GetRealTimeTrafficDataPerInterval() *DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval
	SetRequestId(v string) *DescribeVodDomainRealTimeTrafficDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodDomainRealTimeTrafficDataResponseBody
	GetStartTime() *string
}

type DescribeVodDomainRealTimeTrafficDataResponseBody struct {
	// The time interval at which data is returned. Unit: seconds.
	//
	// The returned value varies based on the time range per query. Valid values: 60 (1 minute), 300 (5 minutes), and 3600 (1 hour). For more information, see the **Time granularity*	- section in the **API documentation**.
	//
	// example:
	//
	// 60
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range.
	//
	// example:
	//
	// 2019-12-10T20:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The details of traffic data in each time interval.
	RealTimeTrafficDataPerInterval *DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval `json:"RealTimeTrafficDataPerInterval,omitempty" xml:"RealTimeTrafficDataPerInterval,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// A666D44F-19D6-490E-97CF-1A64AB962C57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range.
	//
	// example:
	//
	// 2019-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainRealTimeTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBody) GetRealTimeTrafficDataPerInterval() *DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval {
	return s.RealTimeTrafficDataPerInterval
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBody) SetDataInterval(v string) *DescribeVodDomainRealTimeTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBody) SetDomainName(v string) *DescribeVodDomainRealTimeTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBody) SetEndTime(v string) *DescribeVodDomainRealTimeTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBody) SetRealTimeTrafficDataPerInterval(v *DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) *DescribeVodDomainRealTimeTrafficDataResponseBody {
	s.RealTimeTrafficDataPerInterval = v
	return s
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBody) SetRequestId(v string) *DescribeVodDomainRealTimeTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBody) SetStartTime(v string) *DescribeVodDomainRealTimeTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBody) Validate() error {
	if s.RealTimeTrafficDataPerInterval != nil {
		if err := s.RealTimeTrafficDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval struct {
	DataModule []*DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) GetDataModule() []*DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) SetDataModule(v []*DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) *DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) Validate() error {
	if s.DataModule != nil {
		for _, item := range s.DataModule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule struct {
	// The timestamp of the data returned. The time follows the ISO 8601 standard. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-12-10T20:01:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The traffic data. Unit: bytes.
	//
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeVodDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
