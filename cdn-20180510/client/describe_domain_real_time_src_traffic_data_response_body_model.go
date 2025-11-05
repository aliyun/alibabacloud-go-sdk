// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeSrcTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainRealTimeSrcTrafficDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainRealTimeSrcTrafficDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainRealTimeSrcTrafficDataResponseBody
	GetEndTime() *string
	SetRealTimeSrcTrafficDataPerInterval(v *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) *DescribeDomainRealTimeSrcTrafficDataResponseBody
	GetRealTimeSrcTrafficDataPerInterval() *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval
	SetRequestId(v string) *DescribeDomainRealTimeSrcTrafficDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainRealTimeSrcTrafficDataResponseBody
	GetStartTime() *string
}

type DescribeDomainRealTimeSrcTrafficDataResponseBody struct {
	// The time interval between the data entries returned. Unit: seconds.
	//
	// Depending on the maximum time range per query, the value is 60 (1 minute), 300 (5 minutes), or 3600 (1 hour). For more information, see the "Time granularity" section in Usage notes.
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
	// The end of the time range for which the data was queried.
	//
	// example:
	//
	// 2019-12-10T20:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The amount of back-to-origin traffic returned at each interval.
	RealTimeSrcTrafficDataPerInterval *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval `json:"RealTimeSrcTrafficDataPerInterval,omitempty" xml:"RealTimeSrcTrafficDataPerInterval,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// A666D44F-19D6-490E-97CF-1A64AB962C57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range for which the data was queried.
	//
	// example:
	//
	// 2019-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainRealTimeSrcTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeSrcTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBody) GetRealTimeSrcTrafficDataPerInterval() *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval {
	return s.RealTimeSrcTrafficDataPerInterval
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBody) SetDataInterval(v string) *DescribeDomainRealTimeSrcTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBody) SetDomainName(v string) *DescribeDomainRealTimeSrcTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBody) SetEndTime(v string) *DescribeDomainRealTimeSrcTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBody) SetRealTimeSrcTrafficDataPerInterval(v *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) *DescribeDomainRealTimeSrcTrafficDataResponseBody {
	s.RealTimeSrcTrafficDataPerInterval = v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBody) SetRequestId(v string) *DescribeDomainRealTimeSrcTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBody) SetStartTime(v string) *DescribeDomainRealTimeSrcTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBody) Validate() error {
	if s.RealTimeSrcTrafficDataPerInterval != nil {
		if err := s.RealTimeSrcTrafficDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval struct {
	DataModule []*DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) GetDataModule() []*DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) SetDataModule(v []*DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) Validate() error {
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

type DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule struct {
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2019-12-10T20:01:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The amount of traffic.
	//
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
