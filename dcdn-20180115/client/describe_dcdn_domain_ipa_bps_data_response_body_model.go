// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainIpaBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBpsDataPerInterval(v *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerInterval) *DescribeDcdnDomainIpaBpsDataResponseBody
	GetBpsDataPerInterval() *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerInterval
	SetDataInterval(v string) *DescribeDcdnDomainIpaBpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainIpaBpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainIpaBpsDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDcdnDomainIpaBpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainIpaBpsDataResponseBody
	GetStartTime() *string
}

type DescribeDcdnDomainIpaBpsDataResponseBody struct {
	// The bandwidth data returned at each interval.
	BpsDataPerInterval *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" type:"Struct"`
	// The time interval at which data was collected. Unit: seconds.
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
	// The beginning of the time range during which data was queried.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainIpaBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIpaBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBody) GetBpsDataPerInterval() *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerInterval {
	return s.BpsDataPerInterval
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBody) SetBpsDataPerInterval(v *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerInterval) *DescribeDcdnDomainIpaBpsDataResponseBody {
	s.BpsDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainIpaBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainIpaBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainIpaBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainIpaBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainIpaBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerInterval struct {
	DataModule []*DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerInterval) GetDataModule() []*DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerInterval) SetDataModule(v []*DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule) *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule struct {
	// The bandwidth value. Unit: bit/s.
	//
	// example:
	//
	// 11288111
	IpaBps *float32 `json:"IpaBps,omitempty" xml:"IpaBps,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule) GetIpaBps() *float32 {
	return s.IpaBps
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule) SetIpaBps(v float32) *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.IpaBps = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainIpaBpsDataResponseBodyBpsDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
