// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainOriginTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainOriginTrafficDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainOriginTrafficDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainOriginTrafficDataResponseBody
	GetEndTime() *string
	SetOriginTrafficDataPerInterval(v *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval) *DescribeDcdnDomainOriginTrafficDataResponseBody
	GetOriginTrafficDataPerInterval() *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval
	SetRequestId(v string) *DescribeDcdnDomainOriginTrafficDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainOriginTrafficDataResponseBody
	GetStartTime() *string
}

type DescribeDcdnDomainOriginTrafficDataResponseBody struct {
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
	// The amount of back-to-origin traffic returned at each time interval. Unit: bytes.
	OriginTrafficDataPerInterval *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval `json:"OriginTrafficDataPerInterval,omitempty" xml:"OriginTrafficDataPerInterval,omitempty" type:"Struct"`
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
	// 2017-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainOriginTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainOriginTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBody) GetOriginTrafficDataPerInterval() *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval {
	return s.OriginTrafficDataPerInterval
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainOriginTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainOriginTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainOriginTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBody) SetOriginTrafficDataPerInterval(v *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval) *DescribeDcdnDomainOriginTrafficDataResponseBody {
	s.OriginTrafficDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainOriginTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainOriginTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval struct {
	DataModule []*DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval) GetDataModule() []*DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval) SetDataModule(v []*DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule struct {
	// The amount of back-to-origin traffic that was consumed to deliver dynamic content over HTTP.
	//
	// example:
	//
	// 1000
	DynamicHttpOriginTraffic *float32 `json:"DynamicHttpOriginTraffic,omitempty" xml:"DynamicHttpOriginTraffic,omitempty"`
	// The amount of back-to-origin traffic that was consumed to deliver dynamic content over HTTPS.
	//
	// example:
	//
	// 500
	DynamicHttpsOriginTraffic *float32 `json:"DynamicHttpsOriginTraffic,omitempty" xml:"DynamicHttpsOriginTraffic,omitempty"`
	// The amount of back-to-origin traffic.
	//
	// example:
	//
	// 100
	OriginTraffic *float32 `json:"OriginTraffic,omitempty" xml:"OriginTraffic,omitempty"`
	// The amount of back-to-origin traffic that was consumed to deliver static content over HTTP.
	//
	// example:
	//
	// 0
	StaticHttpOriginTraffic *float32 `json:"StaticHttpOriginTraffic,omitempty" xml:"StaticHttpOriginTraffic,omitempty"`
	// The amount of back-to-origin traffic that was consumed to deliver static content over HTTPS.
	//
	// example:
	//
	// 100
	StaticHttpsOriginTraffic *float32 `json:"StaticHttpsOriginTraffic,omitempty" xml:"StaticHttpsOriginTraffic,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2017-12-10T21:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) GetDynamicHttpOriginTraffic() *float32 {
	return s.DynamicHttpOriginTraffic
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) GetDynamicHttpsOriginTraffic() *float32 {
	return s.DynamicHttpsOriginTraffic
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) GetOriginTraffic() *float32 {
	return s.OriginTraffic
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) GetStaticHttpOriginTraffic() *float32 {
	return s.StaticHttpOriginTraffic
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) GetStaticHttpsOriginTraffic() *float32 {
	return s.StaticHttpsOriginTraffic
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) SetDynamicHttpOriginTraffic(v float32) *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule {
	s.DynamicHttpOriginTraffic = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) SetDynamicHttpsOriginTraffic(v float32) *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule {
	s.DynamicHttpsOriginTraffic = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) SetOriginTraffic(v float32) *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule {
	s.OriginTraffic = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) SetStaticHttpOriginTraffic(v float32) *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule {
	s.StaticHttpOriginTraffic = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) SetStaticHttpsOriginTraffic(v float32) *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule {
	s.StaticHttpsOriginTraffic = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainOriginTrafficDataResponseBodyOriginTrafficDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
