// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainWebsocketBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBpsDataPerInterval(v *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerInterval) *DescribeDcdnDomainWebsocketBpsDataResponseBody
	GetBpsDataPerInterval() *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerInterval
	SetDataInterval(v string) *DescribeDcdnDomainWebsocketBpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainWebsocketBpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainWebsocketBpsDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDcdnDomainWebsocketBpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainWebsocketBpsDataResponseBody
	GetStartTime() *string
}

type DescribeDcdnDomainWebsocketBpsDataResponseBody struct {
	// The bandwidth data returned at each interval.
	BpsDataPerInterval *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" type:"Struct"`
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
	// 3C6CCEC4-6B88-4D4A-93E4-D47B3D92CF8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainWebsocketBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBody) GetBpsDataPerInterval() *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerInterval {
	return s.BpsDataPerInterval
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBody) SetBpsDataPerInterval(v *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerInterval) *DescribeDcdnDomainWebsocketBpsDataResponseBody {
	s.BpsDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainWebsocketBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainWebsocketBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainWebsocketBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainWebsocketBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainWebsocketBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerInterval struct {
	DataModule []*DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerInterval) GetDataModule() []*DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerInterval) SetDataModule(v []*DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule) *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule struct {
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The bandwidth value. Unit: bit/s.
	//
	// example:
	//
	// 11288111
	WebsocketBps *float32 `json:"WebsocketBps,omitempty" xml:"WebsocketBps,omitempty"`
}

func (s DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule) GetWebsocketBps() *float32 {
	return s.WebsocketBps
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule) SetWebsocketBps(v float32) *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.WebsocketBps = &v
	return s
}

func (s *DescribeDcdnDomainWebsocketBpsDataResponseBodyBpsDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
