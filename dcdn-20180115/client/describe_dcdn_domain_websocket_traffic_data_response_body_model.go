// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainWebsocketTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainWebsocketTrafficDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainWebsocketTrafficDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainWebsocketTrafficDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDcdnDomainWebsocketTrafficDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainWebsocketTrafficDataResponseBody
	GetStartTime() *string
	SetTrafficDataPerInterval(v *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerInterval) *DescribeDcdnDomainWebsocketTrafficDataResponseBody
	GetTrafficDataPerInterval() *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerInterval
}

type DescribeDcdnDomainWebsocketTrafficDataResponseBody struct {
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
	TrafficDataPerInterval *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerInterval `json:"TrafficDataPerInterval,omitempty" xml:"TrafficDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainWebsocketTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBody) GetTrafficDataPerInterval() *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerInterval {
	return s.TrafficDataPerInterval
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainWebsocketTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainWebsocketTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainWebsocketTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainWebsocketTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainWebsocketTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBody) SetTrafficDataPerInterval(v *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerInterval) *DescribeDcdnDomainWebsocketTrafficDataResponseBody {
	s.TrafficDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerInterval struct {
	DataModule []*DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerInterval) GetDataModule() []*DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerInterval) SetDataModule(v []*DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule) *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule struct {
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The total volume of traffic.
	//
	// example:
	//
	// 423304182
	WebsocketTraffic *float32 `json:"WebsocketTraffic,omitempty" xml:"WebsocketTraffic,omitempty"`
}

func (s DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetWebsocketTraffic() *float32 {
	return s.WebsocketTraffic
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetWebsocketTraffic(v float32) *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.WebsocketTraffic = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketTrafficDataResponseBodyTrafficDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
