// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainMultiUsageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDcdnDomainMultiUsageDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDcdnDomainMultiUsageDataResponseBody
	GetRequestId() *string
	SetRequestPerInterval(v *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerInterval) *DescribeDcdnDomainMultiUsageDataResponseBody
	GetRequestPerInterval() *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerInterval
	SetStartTime(v string) *DescribeDcdnDomainMultiUsageDataResponseBody
	GetStartTime() *string
	SetTrafficPerInterval(v *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerInterval) *DescribeDcdnDomainMultiUsageDataResponseBody
	GetTrafficPerInterval() *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerInterval
}

type DescribeDcdnDomainMultiUsageDataResponseBody struct {
	// The end of the time range that was queried.
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
	// The information about requests collected every 5 minutes.
	RequestPerInterval *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerInterval `json:"RequestPerInterval,omitempty" xml:"RequestPerInterval,omitempty" type:"Struct"`
	// The beginning of the time range that was queried.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The statistics of network traffic collected every 5 minutes.
	TrafficPerInterval *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerInterval `json:"TrafficPerInterval,omitempty" xml:"TrafficPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainMultiUsageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainMultiUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBody) GetRequestPerInterval() *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerInterval {
	return s.RequestPerInterval
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBody) GetTrafficPerInterval() *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerInterval {
	return s.TrafficPerInterval
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainMultiUsageDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainMultiUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBody) SetRequestPerInterval(v *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerInterval) *DescribeDcdnDomainMultiUsageDataResponseBody {
	s.RequestPerInterval = v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainMultiUsageDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBody) SetTrafficPerInterval(v *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerInterval) *DescribeDcdnDomainMultiUsageDataResponseBody {
	s.TrafficPerInterval = v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerInterval struct {
	RequestDataModule []*DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule `json:"RequestDataModule,omitempty" xml:"RequestDataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerInterval) GetRequestDataModule() []*DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule {
	return s.RequestDataModule
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerInterval) SetRequestDataModule(v []*DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerInterval {
	s.RequestDataModule = v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule struct {
	// The domain name.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The number of requests.
	//
	// example:
	//
	// 1128
	Request *int64 `json:"Request,omitempty" xml:"Request,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The type of the requests. Valid values: StaticHttps, DynamicHttps, DynamicHttp, StaticQuic, and DynamicQuic.
	//
	// example:
	//
	// DynamicHttp
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) GetRequest() *int64 {
	return s.Request
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) GetType() *string {
	return s.Type
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) SetDomain(v string) *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule {
	s.Domain = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) SetRequest(v int64) *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule {
	s.Request = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) SetTimeStamp(v string) *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) SetType(v string) *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule {
	s.Type = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerInterval struct {
	TrafficDataModule []*DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule `json:"TrafficDataModule,omitempty" xml:"TrafficDataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerInterval) GetTrafficDataModule() []*DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule {
	return s.TrafficDataModule
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerInterval) SetTrafficDataModule(v []*DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerInterval {
	s.TrafficDataModule = v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule struct {
	// The name of the region.
	//
	// example:
	//
	// CN
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The number of bits per second.
	//
	// example:
	//
	// 11288111.1
	Bps *float32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The type of the network traffic. Valid values: Simple, IPA, and WebSocket.
	//
	// example:
	//
	// Simple
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) GetArea() *string {
	return s.Area
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) GetBps() *float32 {
	return s.Bps
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) GetType() *string {
	return s.Type
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) SetArea(v string) *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule {
	s.Area = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) SetBps(v float32) *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule {
	s.Bps = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) SetDomain(v string) *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule {
	s.Domain = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) SetTimeStamp(v string) *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) SetType(v string) *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule {
	s.Type = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) Validate() error {
	return dara.Validate(s)
}
