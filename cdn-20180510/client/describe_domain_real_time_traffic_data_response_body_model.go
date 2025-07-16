// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainRealTimeTrafficDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainRealTimeTrafficDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainRealTimeTrafficDataResponseBody
	GetEndTime() *string
	SetRealTimeTrafficDataPerInterval(v *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) *DescribeDomainRealTimeTrafficDataResponseBody
	GetRealTimeTrafficDataPerInterval() *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval
	SetRequestId(v string) *DescribeDomainRealTimeTrafficDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainRealTimeTrafficDataResponseBody
	GetStartTime() *string
}

type DescribeDomainRealTimeTrafficDataResponseBody struct {
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
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2019-12-10T20:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The network traffic returned at each time interval. Unit: bytes.
	RealTimeTrafficDataPerInterval *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval `json:"RealTimeTrafficDataPerInterval,omitempty" xml:"RealTimeTrafficDataPerInterval,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// A666D44F-19D6-490E-97CF-1A64AB962C57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2019-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainRealTimeTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeTrafficDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainRealTimeTrafficDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainRealTimeTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainRealTimeTrafficDataResponseBody) GetRealTimeTrafficDataPerInterval() *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval {
	return s.RealTimeTrafficDataPerInterval
}

func (s *DescribeDomainRealTimeTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainRealTimeTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainRealTimeTrafficDataResponseBody) SetDataInterval(v string) *DescribeDomainRealTimeTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataResponseBody) SetDomainName(v string) *DescribeDomainRealTimeTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataResponseBody) SetEndTime(v string) *DescribeDomainRealTimeTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataResponseBody) SetRealTimeTrafficDataPerInterval(v *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) *DescribeDomainRealTimeTrafficDataResponseBody {
	s.RealTimeTrafficDataPerInterval = v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataResponseBody) SetRequestId(v string) *DescribeDomainRealTimeTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataResponseBody) SetStartTime(v string) *DescribeDomainRealTimeTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval struct {
	DataModule []*DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) GetDataModule() []*DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) SetDataModule(v []*DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule struct {
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2019-12-10T20:01:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The traffic value at each time interval.
	//
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
