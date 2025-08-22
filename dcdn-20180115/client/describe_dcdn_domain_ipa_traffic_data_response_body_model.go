// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainIpaTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainIpaTrafficDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainIpaTrafficDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainIpaTrafficDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDcdnDomainIpaTrafficDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainIpaTrafficDataResponseBody
	GetStartTime() *string
	SetTrafficDataPerInterval(v *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerInterval) *DescribeDcdnDomainIpaTrafficDataResponseBody
	GetTrafficDataPerInterval() *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerInterval
}

type DescribeDcdnDomainIpaTrafficDataResponseBody struct {
	// The time interval between the data entries returned. Unit: seconds.
	//
	// example:
	//
	// 300
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
	// 2017-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B955107D-E658-4E77-B913-E0AC3D31693E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The network traffic that was collected at each interval.
	TrafficDataPerInterval *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerInterval `json:"TrafficDataPerInterval,omitempty" xml:"TrafficDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainIpaTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIpaTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBody) GetTrafficDataPerInterval() *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerInterval {
	return s.TrafficDataPerInterval
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainIpaTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainIpaTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainIpaTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainIpaTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainIpaTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBody) SetTrafficDataPerInterval(v *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerInterval) *DescribeDcdnDomainIpaTrafficDataResponseBody {
	s.TrafficDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerInterval struct {
	DataModule []*DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerInterval) GetDataModule() []*DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerInterval) SetDataModule(v []*DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule) *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule struct {
	// The total amount of network traffic.
	//
	// example:
	//
	// 423304182
	IpaTraffic *float32 `json:"IpaTraffic,omitempty" xml:"IpaTraffic,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetIpaTraffic() *float32 {
	return s.IpaTraffic
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetIpaTraffic(v float32) *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.IpaTraffic = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainIpaTrafficDataResponseBodyTrafficDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
