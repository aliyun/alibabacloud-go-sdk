// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainMultiUsageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDomainMultiUsageDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDomainMultiUsageDataResponseBody
	GetRequestId() *string
	SetRequestPerInterval(v *DescribeDomainMultiUsageDataResponseBodyRequestPerInterval) *DescribeDomainMultiUsageDataResponseBody
	GetRequestPerInterval() *DescribeDomainMultiUsageDataResponseBodyRequestPerInterval
	SetStartTime(v string) *DescribeDomainMultiUsageDataResponseBody
	GetStartTime() *string
	SetTrafficPerInterval(v *DescribeDomainMultiUsageDataResponseBodyTrafficPerInterval) *DescribeDomainMultiUsageDataResponseBody
	GetTrafficPerInterval() *DescribeDomainMultiUsageDataResponseBodyTrafficPerInterval
}

type DescribeDomainMultiUsageDataResponseBody struct {
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
	RequestPerInterval *DescribeDomainMultiUsageDataResponseBodyRequestPerInterval `json:"RequestPerInterval,omitempty" xml:"RequestPerInterval,omitempty" type:"Struct"`
	// The start of the time range that was queried.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The statistics of network traffic collected every 5 minutes.
	TrafficPerInterval *DescribeDomainMultiUsageDataResponseBodyTrafficPerInterval `json:"TrafficPerInterval,omitempty" xml:"TrafficPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainMultiUsageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainMultiUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainMultiUsageDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainMultiUsageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainMultiUsageDataResponseBody) GetRequestPerInterval() *DescribeDomainMultiUsageDataResponseBodyRequestPerInterval {
	return s.RequestPerInterval
}

func (s *DescribeDomainMultiUsageDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainMultiUsageDataResponseBody) GetTrafficPerInterval() *DescribeDomainMultiUsageDataResponseBodyTrafficPerInterval {
	return s.TrafficPerInterval
}

func (s *DescribeDomainMultiUsageDataResponseBody) SetEndTime(v string) *DescribeDomainMultiUsageDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainMultiUsageDataResponseBody) SetRequestId(v string) *DescribeDomainMultiUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainMultiUsageDataResponseBody) SetRequestPerInterval(v *DescribeDomainMultiUsageDataResponseBodyRequestPerInterval) *DescribeDomainMultiUsageDataResponseBody {
	s.RequestPerInterval = v
	return s
}

func (s *DescribeDomainMultiUsageDataResponseBody) SetStartTime(v string) *DescribeDomainMultiUsageDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainMultiUsageDataResponseBody) SetTrafficPerInterval(v *DescribeDomainMultiUsageDataResponseBodyTrafficPerInterval) *DescribeDomainMultiUsageDataResponseBody {
	s.TrafficPerInterval = v
	return s
}

func (s *DescribeDomainMultiUsageDataResponseBody) Validate() error {
	if s.RequestPerInterval != nil {
		if err := s.RequestPerInterval.Validate(); err != nil {
			return err
		}
	}
	if s.TrafficPerInterval != nil {
		if err := s.TrafficPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainMultiUsageDataResponseBodyRequestPerInterval struct {
	RequestDataModule []*DescribeDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule `json:"RequestDataModule,omitempty" xml:"RequestDataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainMultiUsageDataResponseBodyRequestPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainMultiUsageDataResponseBodyRequestPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainMultiUsageDataResponseBodyRequestPerInterval) GetRequestDataModule() []*DescribeDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule {
	return s.RequestDataModule
}

func (s *DescribeDomainMultiUsageDataResponseBodyRequestPerInterval) SetRequestDataModule(v []*DescribeDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) *DescribeDomainMultiUsageDataResponseBodyRequestPerInterval {
	s.RequestDataModule = v
	return s
}

func (s *DescribeDomainMultiUsageDataResponseBodyRequestPerInterval) Validate() error {
	if s.RequestDataModule != nil {
		for _, item := range s.RequestDataModule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule struct {
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The number of requests.
	//
	// example:
	//
	// 11288111
	Request *int64 `json:"Request,omitempty" xml:"Request,omitempty"`
	// The timestamp of the returned number of requests.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The type.
	//
	// >  The value is Simple for Alibaba Cloud CDN.
	//
	// example:
	//
	// Simple
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) GetRequest() *int64 {
	return s.Request
}

func (s *DescribeDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) GetType() *string {
	return s.Type
}

func (s *DescribeDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) SetDomain(v string) *DescribeDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule {
	s.Domain = &v
	return s
}

func (s *DescribeDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) SetRequest(v int64) *DescribeDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule {
	s.Request = &v
	return s
}

func (s *DescribeDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) SetTimeStamp(v string) *DescribeDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) SetType(v string) *DescribeDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule {
	s.Type = &v
	return s
}

func (s *DescribeDomainMultiUsageDataResponseBodyRequestPerIntervalRequestDataModule) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainMultiUsageDataResponseBodyTrafficPerInterval struct {
	TrafficDataModule []*DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule `json:"TrafficDataModule,omitempty" xml:"TrafficDataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainMultiUsageDataResponseBodyTrafficPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainMultiUsageDataResponseBodyTrafficPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainMultiUsageDataResponseBodyTrafficPerInterval) GetTrafficDataModule() []*DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule {
	return s.TrafficDataModule
}

func (s *DescribeDomainMultiUsageDataResponseBodyTrafficPerInterval) SetTrafficDataModule(v []*DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) *DescribeDomainMultiUsageDataResponseBodyTrafficPerInterval {
	s.TrafficDataModule = v
	return s
}

func (s *DescribeDomainMultiUsageDataResponseBodyTrafficPerInterval) Validate() error {
	if s.TrafficDataModule != nil {
		for _, item := range s.TrafficDataModule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule struct {
	// The name of the region.
	//
	// example:
	//
	// CN
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The bandwidth. Unit: bit/s.
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
	// The startstamp of the returned usage data.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The type of requests. Valid values:
	//
	// 	- **StaticHttps**: static HTTPS requests
	//
	// 	- **DynamicHttps**: dynamic HTTPS requests
	//
	// 	- **DynamicHttp**: dynamic HTTP requests
	//
	// 	- **StaticQuic**: static QUIC requests
	//
	// 	- **DynamicQuic**: dynamic QUIC requests
	//
	// example:
	//
	// DynamicHttp
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) GetArea() *string {
	return s.Area
}

func (s *DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) GetBps() *float32 {
	return s.Bps
}

func (s *DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) GetType() *string {
	return s.Type
}

func (s *DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) SetArea(v string) *DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule {
	s.Area = &v
	return s
}

func (s *DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) SetBps(v float32) *DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule {
	s.Bps = &v
	return s
}

func (s *DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) SetDomain(v string) *DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule {
	s.Domain = &v
	return s
}

func (s *DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) SetTimeStamp(v string) *DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) SetType(v string) *DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule {
	s.Type = &v
	return s
}

func (s *DescribeDomainMultiUsageDataResponseBodyTrafficPerIntervalTrafficDataModule) Validate() error {
	return dara.Validate(s)
}
