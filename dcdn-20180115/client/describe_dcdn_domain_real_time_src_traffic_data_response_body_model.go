// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeSrcTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody
	GetEndTime() *string
	SetRealTimeSrcTrafficDataPerInterval(v *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody
	GetRealTimeSrcTrafficDataPerInterval() *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval
	SetRequestId(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody
	GetStartTime() *string
}

type DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody struct {
	// The time interval between the data entries returned. Unit: seconds.
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
	// The amount of origin traffic returned at each time interval. Unit: bytes.
	RealTimeSrcTrafficDataPerInterval *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval `json:"RealTimeSrcTrafficDataPerInterval,omitempty" xml:"RealTimeSrcTrafficDataPerInterval,omitempty" type:"Struct"`
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
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) GetRealTimeSrcTrafficDataPerInterval() *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval {
	return s.RealTimeSrcTrafficDataPerInterval
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) SetRealTimeSrcTrafficDataPerInterval(v *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody {
	s.RealTimeSrcTrafficDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval struct {
	DataModule []*DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) GetDataModule() []*DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) SetDataModule(v []*DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule struct {
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The traffic value at each time interval.
	//
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcTrafficDataResponseBodyRealTimeSrcTrafficDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
