// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainUvDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainUvDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainUvDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainUvDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDcdnDomainUvDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainUvDataResponseBody
	GetStartTime() *string
	SetUvDataInterval(v *DescribeDcdnDomainUvDataResponseBodyUvDataInterval) *DescribeDcdnDomainUvDataResponseBody
	GetUvDataInterval() *DescribeDcdnDomainUvDataResponseBodyUvDataInterval
}

type DescribeDcdnDomainUvDataResponseBody struct {
	// The time interval between the data entries returned. Unit: seconds.
	//
	// example:
	//
	// 3600
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range that was queried.
	//
	// example:
	//
	// 2015-11-30T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E9D3257A-1B7C-414C-90C1-8D07AC47BCAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range that was queried.
	//
	// example:
	//
	// 2015-11-29T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The number of UVs at each interval.
	UvDataInterval *DescribeDcdnDomainUvDataResponseBodyUvDataInterval `json:"UvDataInterval,omitempty" xml:"UvDataInterval,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainUvDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainUvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUvDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainUvDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainUvDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainUvDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainUvDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainUvDataResponseBody) GetUvDataInterval() *DescribeDcdnDomainUvDataResponseBodyUvDataInterval {
	return s.UvDataInterval
}

func (s *DescribeDcdnDomainUvDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainUvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainUvDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainUvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainUvDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainUvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainUvDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainUvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainUvDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainUvDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainUvDataResponseBody) SetUvDataInterval(v *DescribeDcdnDomainUvDataResponseBodyUvDataInterval) *DescribeDcdnDomainUvDataResponseBody {
	s.UvDataInterval = v
	return s
}

func (s *DescribeDcdnDomainUvDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainUvDataResponseBodyUvDataInterval struct {
	UsageData []*DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainUvDataResponseBodyUvDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainUvDataResponseBodyUvDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUvDataResponseBodyUvDataInterval) GetUsageData() []*DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData {
	return s.UsageData
}

func (s *DescribeDcdnDomainUvDataResponseBodyUvDataInterval) SetUsageData(v []*DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData) *DescribeDcdnDomainUvDataResponseBodyUvDataInterval {
	s.UsageData = v
	return s
}

func (s *DescribeDcdnDomainUvDataResponseBodyUvDataInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData struct {
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2015-11-29T00:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The number of UVs.
	//
	// example:
	//
	// 326
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData) GetValue() *string {
	return s.Value
}

func (s *DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData) SetTimeStamp(v string) *DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData) SetValue(v string) *DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData {
	s.Value = &v
	return s
}

func (s *DescribeDcdnDomainUvDataResponseBodyUvDataIntervalUsageData) Validate() error {
	return dara.Validate(s)
}
