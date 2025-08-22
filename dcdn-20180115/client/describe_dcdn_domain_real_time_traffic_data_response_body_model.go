// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainRealTimeTrafficDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainRealTimeTrafficDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainRealTimeTrafficDataResponseBody
	GetEndTime() *string
	SetRealTimeTrafficDataPerInterval(v *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) *DescribeDcdnDomainRealTimeTrafficDataResponseBody
	GetRealTimeTrafficDataPerInterval() *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval
	SetRequestId(v string) *DescribeDcdnDomainRealTimeTrafficDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainRealTimeTrafficDataResponseBody
	GetStartTime() *string
}

type DescribeDcdnDomainRealTimeTrafficDataResponseBody struct {
	// The time interval between the data entries. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 60 (1 minute), 300 (5 minutes), and 3600(1 hour). For more information, see **Usage notes**.
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
	// 2015-12-10T20:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The amount of back-to-origin traffic returned at each interval.
	RealTimeTrafficDataPerInterval *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval `json:"RealTimeTrafficDataPerInterval,omitempty" xml:"RealTimeTrafficDataPerInterval,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// A666D44F-19D6-490E-97CF-1A64AB962C57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range during which data was queried.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBody) GetRealTimeTrafficDataPerInterval() *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval {
	return s.RealTimeTrafficDataPerInterval
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainRealTimeTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainRealTimeTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainRealTimeTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBody) SetRealTimeTrafficDataPerInterval(v *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) *DescribeDcdnDomainRealTimeTrafficDataResponseBody {
	s.RealTimeTrafficDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainRealTimeTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval struct {
	DataModule []*DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) GetDataModule() []*DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) SetDataModule(v []*DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule struct {
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The traffic value at each time interval.
	//
	// > The network traffic is measured in bytes.
	//
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
