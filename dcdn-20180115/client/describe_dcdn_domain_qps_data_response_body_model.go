// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainQpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainQpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainQpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainQpsDataResponseBody
	GetEndTime() *string
	SetQpsDataPerInterval(v *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerInterval) *DescribeDcdnDomainQpsDataResponseBody
	GetQpsDataPerInterval() *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerInterval
	SetRequestId(v string) *DescribeDcdnDomainQpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainQpsDataResponseBody
	GetStartTime() *string
}

type DescribeDcdnDomainQpsDataResponseBody struct {
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
	// The QPS returned at each time interval.
	QpsDataPerInterval *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerInterval `json:"QpsDataPerInterval,omitempty" xml:"QpsDataPerInterval,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainQpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainQpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainQpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainQpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainQpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainQpsDataResponseBody) GetQpsDataPerInterval() *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerInterval {
	return s.QpsDataPerInterval
}

func (s *DescribeDcdnDomainQpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainQpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainQpsDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainQpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainQpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainQpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBody) SetQpsDataPerInterval(v *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerInterval) *DescribeDcdnDomainQpsDataResponseBody {
	s.QpsDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainQpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainQpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainQpsDataResponseBodyQpsDataPerInterval struct {
	DataModule []*DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainQpsDataResponseBodyQpsDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainQpsDataResponseBodyQpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerInterval) GetDataModule() []*DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerInterval) SetDataModule(v []*DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule struct {
	// The total number of requests.
	//
	// example:
	//
	// 100
	Acc *float32 `json:"Acc,omitempty" xml:"Acc,omitempty"`
	// The number of requests for dynamic content delivery over HTTP.
	//
	// example:
	//
	// 0
	DynamicHttpAcc *float32 `json:"DynamicHttpAcc,omitempty" xml:"DynamicHttpAcc,omitempty"`
	// The QPS for dynamic content delivery over HTTP.
	//
	// example:
	//
	// 0
	DynamicHttpQps *float32 `json:"DynamicHttpQps,omitempty" xml:"DynamicHttpQps,omitempty"`
	// The number of requests for dynamic content delivery over HTTPS.
	//
	// example:
	//
	// 0
	DynamicHttpsAcc *float32 `json:"DynamicHttpsAcc,omitempty" xml:"DynamicHttpsAcc,omitempty"`
	// The QPS for dynamic content delivery over HTTPS.
	//
	// example:
	//
	// 0
	DynamicHttpsQps *float32 `json:"DynamicHttpsQps,omitempty" xml:"DynamicHttpsQps,omitempty"`
	// The total QPS.
	//
	// example:
	//
	// 0.56
	Qps *float32 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The number of requests for static content delivery over HTTP.
	//
	// example:
	//
	// 0
	StaticHttpAcc *float32 `json:"StaticHttpAcc,omitempty" xml:"StaticHttpAcc,omitempty"`
	// The QPS for static content delivery over HTTP.
	//
	// example:
	//
	// 0
	StaticHttpQps *float32 `json:"StaticHttpQps,omitempty" xml:"StaticHttpQps,omitempty"`
	// The number of requests for static content delivery over HTTPS.
	//
	// example:
	//
	// 0
	StaticHttpsAcc *float32 `json:"StaticHttpsAcc,omitempty" xml:"StaticHttpsAcc,omitempty"`
	// The QPS for static content delivery over HTTPS.
	//
	// example:
	//
	// 0
	StaticHttpsQps *float32 `json:"StaticHttpsQps,omitempty" xml:"StaticHttpsQps,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2017-12-10T21:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) GetAcc() *float32 {
	return s.Acc
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) GetDynamicHttpAcc() *float32 {
	return s.DynamicHttpAcc
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) GetDynamicHttpQps() *float32 {
	return s.DynamicHttpQps
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) GetDynamicHttpsAcc() *float32 {
	return s.DynamicHttpsAcc
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) GetDynamicHttpsQps() *float32 {
	return s.DynamicHttpsQps
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) GetQps() *float32 {
	return s.Qps
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) GetStaticHttpAcc() *float32 {
	return s.StaticHttpAcc
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) GetStaticHttpQps() *float32 {
	return s.StaticHttpQps
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) GetStaticHttpsAcc() *float32 {
	return s.StaticHttpsAcc
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) GetStaticHttpsQps() *float32 {
	return s.StaticHttpsQps
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetAcc(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.Acc = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetDynamicHttpAcc(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.DynamicHttpAcc = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetDynamicHttpQps(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.DynamicHttpQps = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetDynamicHttpsAcc(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.DynamicHttpsAcc = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetDynamicHttpsQps(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.DynamicHttpsQps = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetQps(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.Qps = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetStaticHttpAcc(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.StaticHttpAcc = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetStaticHttpQps(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.StaticHttpQps = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetStaticHttpsAcc(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.StaticHttpsAcc = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetStaticHttpsQps(v float32) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.StaticHttpsQps = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainQpsDataResponseBodyQpsDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
