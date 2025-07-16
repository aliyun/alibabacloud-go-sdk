// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainHitRateDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainHitRateDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainHitRateDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainHitRateDataResponseBody
	GetEndTime() *string
	SetHitRateInterval(v *DescribeDomainHitRateDataResponseBodyHitRateInterval) *DescribeDomainHitRateDataResponseBody
	GetHitRateInterval() *DescribeDomainHitRateDataResponseBodyHitRateInterval
	SetRequestId(v string) *DescribeDomainHitRateDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainHitRateDataResponseBody
	GetStartTime() *string
}

type DescribeDomainHitRateDataResponseBody struct {
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
	// 2019-12-30T08:10:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The byte hit ratio at each time interval. The byte hit ratio is measured in percentage.
	HitRateInterval *DescribeDomainHitRateDataResponseBodyHitRateInterval `json:"HitRateInterval,omitempty" xml:"HitRateInterval,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2019-12-30T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainHitRateDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainHitRateDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainHitRateDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainHitRateDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainHitRateDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainHitRateDataResponseBody) GetHitRateInterval() *DescribeDomainHitRateDataResponseBodyHitRateInterval {
	return s.HitRateInterval
}

func (s *DescribeDomainHitRateDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainHitRateDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainHitRateDataResponseBody) SetDataInterval(v string) *DescribeDomainHitRateDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBody) SetDomainName(v string) *DescribeDomainHitRateDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBody) SetEndTime(v string) *DescribeDomainHitRateDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBody) SetHitRateInterval(v *DescribeDomainHitRateDataResponseBodyHitRateInterval) *DescribeDomainHitRateDataResponseBody {
	s.HitRateInterval = v
	return s
}

func (s *DescribeDomainHitRateDataResponseBody) SetRequestId(v string) *DescribeDomainHitRateDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBody) SetStartTime(v string) *DescribeDomainHitRateDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainHitRateDataResponseBodyHitRateInterval struct {
	DataModule []*DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainHitRateDataResponseBodyHitRateInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainHitRateDataResponseBodyHitRateInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainHitRateDataResponseBodyHitRateInterval) GetDataModule() []*DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDomainHitRateDataResponseBodyHitRateInterval) SetDataModule(v []*DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) *DescribeDomainHitRateDataResponseBodyHitRateInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDomainHitRateDataResponseBodyHitRateInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule struct {
	// The byte hit ratio of HTTPS requests.
	//
	// example:
	//
	// 50.0
	HttpsValue *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2019-12-30T08:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The byte hit ratio.
	//
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) GetHttpsValue() *string {
	return s.HttpsValue
}

func (s *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) SetHttpsValue(v string) *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule {
	s.HttpsValue = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) SetTimeStamp(v string) *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) SetValue(v string) *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainHitRateDataResponseBodyHitRateIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
