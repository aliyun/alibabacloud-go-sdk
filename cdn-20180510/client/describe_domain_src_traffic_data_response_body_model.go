// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSrcTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainSrcTrafficDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainSrcTrafficDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainSrcTrafficDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDomainSrcTrafficDataResponseBody
	GetRequestId() *string
	SetSrcTrafficDataPerInterval(v *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval) *DescribeDomainSrcTrafficDataResponseBody
	GetSrcTrafficDataPerInterval() *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval
	SetStartTime(v string) *DescribeDomainSrcTrafficDataResponseBody
	GetStartTime() *string
}

type DescribeDomainSrcTrafficDataResponseBody struct {
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
	// 2015-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A666D44F-19D6-490E-97CF-1A64AB962C57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The amount of origin traffic returned at each time interval. Unit: bytes.
	SrcTrafficDataPerInterval *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval `json:"SrcTrafficDataPerInterval,omitempty" xml:"SrcTrafficDataPerInterval,omitempty" type:"Struct"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainSrcTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTrafficDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainSrcTrafficDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainSrcTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainSrcTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainSrcTrafficDataResponseBody) GetSrcTrafficDataPerInterval() *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval {
	return s.SrcTrafficDataPerInterval
}

func (s *DescribeDomainSrcTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainSrcTrafficDataResponseBody) SetDataInterval(v string) *DescribeDomainSrcTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponseBody) SetDomainName(v string) *DescribeDomainSrcTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponseBody) SetEndTime(v string) *DescribeDomainSrcTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponseBody) SetRequestId(v string) *DescribeDomainSrcTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponseBody) SetSrcTrafficDataPerInterval(v *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval) *DescribeDomainSrcTrafficDataResponseBody {
	s.SrcTrafficDataPerInterval = v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponseBody) SetStartTime(v string) *DescribeDomainSrcTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval struct {
	DataModule []*DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval) GetDataModule() []*DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval) SetDataModule(v []*DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule struct {
	// The amount of traffic generated by origin HTTPS requests.
	//
	// example:
	//
	// 0
	HttpsValue *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2015-12-10T20:35:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The traffic value at each time interval.
	//
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) GetHttpsValue() *string {
	return s.HttpsValue
}

func (s *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) SetHttpsValue(v string) *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule {
	s.HttpsValue = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponseBodySrcTrafficDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
