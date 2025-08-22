// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBpsDataPerInterval(v *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerInterval) *DescribeDcdnDomainBpsDataResponseBody
	GetBpsDataPerInterval() *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerInterval
	SetDataInterval(v string) *DescribeDcdnDomainBpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainBpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainBpsDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDcdnDomainBpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainBpsDataResponseBody
	GetStartTime() *string
}

type DescribeDcdnDomainBpsDataResponseBody struct {
	// The bandwidth data returned at each interval.
	BpsDataPerInterval *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" type:"Struct"`
	// The time interval between the data entries returned.
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
	// 2017-12-10T10:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainBpsDataResponseBody) GetBpsDataPerInterval() *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerInterval {
	return s.BpsDataPerInterval
}

func (s *DescribeDcdnDomainBpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainBpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainBpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainBpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainBpsDataResponseBody) SetBpsDataPerInterval(v *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerInterval) *DescribeDcdnDomainBpsDataResponseBody {
	s.BpsDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainBpsDataResponseBodyBpsDataPerInterval struct {
	DataModule []*DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainBpsDataResponseBodyBpsDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainBpsDataResponseBodyBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerInterval) GetDataModule() []*DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerInterval) SetDataModule(v []*DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule struct {
	// The bandwidth value. Unit: bit/s.
	//
	// example:
	//
	// 11286
	Bps *float32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The bandwidth that was consumed to deliver dynamic content over HTTP. Unit: bit/s.
	//
	// example:
	//
	// 11286111
	DynamicHttpBps *float32 `json:"DynamicHttpBps,omitempty" xml:"DynamicHttpBps,omitempty"`
	// The bandwidth that was consumed to deliver dynamic content over HTTPS. Unit: bit/s.
	//
	// example:
	//
	// 12312
	DynamicHttpsBps *float32 `json:"DynamicHttpsBps,omitempty" xml:"DynamicHttpsBps,omitempty"`
	// The bandwidth that was consumed to deliver static content over HTTP. Unit: bit/s.
	//
	// example:
	//
	// 123
	StaticHttpBps *float32 `json:"StaticHttpBps,omitempty" xml:"StaticHttpBps,omitempty"`
	// The bandwidth that was consumed to deliver static content over HTTPS. Unit: bit/s.
	//
	// example:
	//
	// 123
	StaticHttpsBps *float32 `json:"StaticHttpsBps,omitempty" xml:"StaticHttpsBps,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetBps() *float32 {
	return s.Bps
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetDynamicHttpBps() *float32 {
	return s.DynamicHttpBps
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetDynamicHttpsBps() *float32 {
	return s.DynamicHttpsBps
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetStaticHttpBps() *float32 {
	return s.StaticHttpBps
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetStaticHttpsBps() *float32 {
	return s.StaticHttpsBps
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetBps(v float32) *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.Bps = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDynamicHttpBps(v float32) *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DynamicHttpBps = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDynamicHttpsBps(v float32) *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DynamicHttpsBps = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetStaticHttpBps(v float32) *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.StaticHttpBps = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetStaticHttpsBps(v float32) *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.StaticHttpsBps = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
