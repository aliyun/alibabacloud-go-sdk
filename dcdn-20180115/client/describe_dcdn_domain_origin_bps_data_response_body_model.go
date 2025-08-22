// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainOriginBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainOriginBpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainOriginBpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainOriginBpsDataResponseBody
	GetEndTime() *string
	SetOriginBpsDataPerInterval(v *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval) *DescribeDcdnDomainOriginBpsDataResponseBody
	GetOriginBpsDataPerInterval() *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval
	SetRequestId(v string) *DescribeDcdnDomainOriginBpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainOriginBpsDataResponseBody
	GetStartTime() *string
}

type DescribeDcdnDomainOriginBpsDataResponseBody struct {
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
	// 2019-12-11T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The origin bandwidth data returned at each time interval. Unit: bit/s.
	OriginBpsDataPerInterval *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval `json:"OriginBpsDataPerInterval,omitempty" xml:"OriginBpsDataPerInterval,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7CBCD6AD-B016-42E5-AE0B-B3731DE8F755
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2019-12-10T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainOriginBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainOriginBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBody) GetOriginBpsDataPerInterval() *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval {
	return s.OriginBpsDataPerInterval
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainOriginBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainOriginBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainOriginBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBody) SetOriginBpsDataPerInterval(v *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval) *DescribeDcdnDomainOriginBpsDataResponseBody {
	s.OriginBpsDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainOriginBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainOriginBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval struct {
	DataModule []*DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval) GetDataModule() []*DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval) SetDataModule(v []*DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule struct {
	// The bandwidth that was consumed for fetching dynamic content from the origin over HTTP.
	//
	// example:
	//
	// 100
	DynamicHttpOriginBps *float32 `json:"DynamicHttpOriginBps,omitempty" xml:"DynamicHttpOriginBps,omitempty"`
	// The bandwidth that was consumed for fetching dynamic content from the origin over HTTPS.
	//
	// example:
	//
	// 100
	DynamicHttpsOriginBps *float32 `json:"DynamicHttpsOriginBps,omitempty" xml:"DynamicHttpsOriginBps,omitempty"`
	// The bandwidth that was consumed for fetching content from the origin.
	//
	// example:
	//
	// 100
	OriginBps *float32 `json:"OriginBps,omitempty" xml:"OriginBps,omitempty"`
	// The bandwidth that was consumed for fetching static content from the origin over HTTP.
	//
	// example:
	//
	// 100
	StaticHttpOriginBps *float32 `json:"StaticHttpOriginBps,omitempty" xml:"StaticHttpOriginBps,omitempty"`
	// The bandwidth that was consumed for fetching static content from the origin over HTTPS.
	//
	// example:
	//
	// 100
	StaticHttpsOriginBps *float32 `json:"StaticHttpsOriginBps,omitempty" xml:"StaticHttpsOriginBps,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2019-12-10T00:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) GetDynamicHttpOriginBps() *float32 {
	return s.DynamicHttpOriginBps
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) GetDynamicHttpsOriginBps() *float32 {
	return s.DynamicHttpsOriginBps
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) GetOriginBps() *float32 {
	return s.OriginBps
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) GetStaticHttpOriginBps() *float32 {
	return s.StaticHttpOriginBps
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) GetStaticHttpsOriginBps() *float32 {
	return s.StaticHttpsOriginBps
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) SetDynamicHttpOriginBps(v float32) *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule {
	s.DynamicHttpOriginBps = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) SetDynamicHttpsOriginBps(v float32) *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule {
	s.DynamicHttpsOriginBps = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) SetOriginBps(v float32) *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule {
	s.OriginBps = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) SetStaticHttpOriginBps(v float32) *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule {
	s.StaticHttpOriginBps = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) SetStaticHttpsOriginBps(v float32) *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule {
	s.StaticHttpsOriginBps = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainOriginBpsDataResponseBodyOriginBpsDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
