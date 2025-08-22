// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeSrcBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody
	GetEndTime() *string
	SetRealTimeSrcBpsDataPerInterval(v *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody
	GetRealTimeSrcBpsDataPerInterval() *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval
	SetRequestId(v string) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody
	GetStartTime() *string
}

type DescribeDcdnDomainRealTimeSrcBpsDataResponseBody struct {
	// The time interval between the data entries returned. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. The returned value is 60 (1 minute), 300 (5 minutes), or 3600(1 hour). For more information, see **Usage notes**.
	//
	// example:
	//
	// 60
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
	// 2015-12-10T20:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The origin bandwidth data returned at each time interval. Unit: bit/s.
	RealTimeSrcBpsDataPerInterval *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval `json:"RealTimeSrcBpsDataPerInterval,omitempty" xml:"RealTimeSrcBpsDataPerInterval,omitempty" type:"Struct"`
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
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) GetRealTimeSrcBpsDataPerInterval() *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval {
	return s.RealTimeSrcBpsDataPerInterval
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) SetRealTimeSrcBpsDataPerInterval(v *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody {
	s.RealTimeSrcBpsDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval struct {
	DataModule []*DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) GetDataModule() []*DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) SetDataModule(v []*DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule struct {
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2015-12-10T20:01:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The bandwidth value.
	//
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) SetValue(v string) *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
