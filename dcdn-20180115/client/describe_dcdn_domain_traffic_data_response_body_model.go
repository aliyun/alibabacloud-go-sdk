// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainTrafficDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainTrafficDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainTrafficDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDcdnDomainTrafficDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainTrafficDataResponseBody
	GetStartTime() *string
	SetTrafficDataPerInterval(v *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerInterval) *DescribeDcdnDomainTrafficDataResponseBody
	GetTrafficDataPerInterval() *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerInterval
}

type DescribeDcdnDomainTrafficDataResponseBody struct {
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
	// The network traffic returned at each time interval. Unit: bytes.
	TrafficDataPerInterval *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerInterval `json:"TrafficDataPerInterval,omitempty" xml:"TrafficDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTrafficDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainTrafficDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainTrafficDataResponseBody) GetTrafficDataPerInterval() *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerInterval {
	return s.TrafficDataPerInterval
}

func (s *DescribeDcdnDomainTrafficDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBody) SetTrafficDataPerInterval(v *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerInterval) *DescribeDcdnDomainTrafficDataResponseBody {
	s.TrafficDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerInterval struct {
	DataModule []*DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerInterval) GetDataModule() []*DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerInterval) SetDataModule(v []*DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule struct {
	// The network traffic that was consumed to deliver dynamic content over HTTP.
	//
	// example:
	//
	// 0
	DynamicHttpTraffic *float32 `json:"DynamicHttpTraffic,omitempty" xml:"DynamicHttpTraffic,omitempty"`
	// The network traffic that was consumed to deliver dynamic content over HTTPS.
	//
	// example:
	//
	// 0
	DynamicHttpsTraffic *float32 `json:"DynamicHttpsTraffic,omitempty" xml:"DynamicHttpsTraffic,omitempty"`
	// The network traffic that was consumed to deliver static content over HTTP.
	//
	// example:
	//
	// 123
	StaticHttpTraffic *float32 `json:"StaticHttpTraffic,omitempty" xml:"StaticHttpTraffic,omitempty"`
	// The network traffic that was consumed to deliver static content over HTTPS.
	//
	// example:
	//
	// 132
	StaticHttpsTraffic *float32 `json:"StaticHttpsTraffic,omitempty" xml:"StaticHttpsTraffic,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The total amount of network traffic.
	//
	// example:
	//
	// 0
	Traffic *float32 `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
}

func (s DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetDynamicHttpTraffic() *float32 {
	return s.DynamicHttpTraffic
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetDynamicHttpsTraffic() *float32 {
	return s.DynamicHttpsTraffic
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetStaticHttpTraffic() *float32 {
	return s.StaticHttpTraffic
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetStaticHttpsTraffic() *float32 {
	return s.StaticHttpsTraffic
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetTraffic() *float32 {
	return s.Traffic
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetDynamicHttpTraffic(v float32) *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.DynamicHttpTraffic = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetDynamicHttpsTraffic(v float32) *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.DynamicHttpsTraffic = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetStaticHttpTraffic(v float32) *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.StaticHttpTraffic = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetStaticHttpsTraffic(v float32) *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.StaticHttpsTraffic = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTraffic(v float32) *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.Traffic = &v
	return s
}

func (s *DescribeDcdnDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
