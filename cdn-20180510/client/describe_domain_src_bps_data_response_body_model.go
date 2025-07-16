// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSrcBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainSrcBpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainSrcBpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainSrcBpsDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDomainSrcBpsDataResponseBody
	GetRequestId() *string
	SetSrcBpsDataPerInterval(v *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) *DescribeDomainSrcBpsDataResponseBody
	GetSrcBpsDataPerInterval() *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval
	SetStartTime(v string) *DescribeDomainSrcBpsDataResponseBody
	GetStartTime() *string
}

type DescribeDomainSrcBpsDataResponseBody struct {
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
	// 2019-12-10T20:30:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The origin bandwidth data at each time interval. Unit: bit/s.
	SrcBpsDataPerInterval *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval `json:"SrcBpsDataPerInterval,omitempty" xml:"SrcBpsDataPerInterval,omitempty" type:"Struct"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2019-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainSrcBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcBpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainSrcBpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainSrcBpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainSrcBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainSrcBpsDataResponseBody) GetSrcBpsDataPerInterval() *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval {
	return s.SrcBpsDataPerInterval
}

func (s *DescribeDomainSrcBpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainSrcBpsDataResponseBody) SetDataInterval(v string) *DescribeDomainSrcBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBody) SetDomainName(v string) *DescribeDomainSrcBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBody) SetEndTime(v string) *DescribeDomainSrcBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBody) SetRequestId(v string) *DescribeDomainSrcBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBody) SetSrcBpsDataPerInterval(v *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) *DescribeDomainSrcBpsDataResponseBody {
	s.SrcBpsDataPerInterval = v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBody) SetStartTime(v string) *DescribeDomainSrcBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval struct {
	DataModule []*DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) GetDataModule() []*DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) SetDataModule(v []*DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule struct {
	// The bandwidth values of origin HTTPS requests.
	//
	// example:
	//
	// 10
	HttpsValue *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2019-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The traffic value at each time interval.
	//
	// example:
	//
	// 500
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) GetHttpsValue() *string {
	return s.HttpsValue
}

func (s *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) SetHttpsValue(v string) *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule {
	s.HttpsValue = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) SetValue(v string) *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainSrcBpsDataResponseBodySrcBpsDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
