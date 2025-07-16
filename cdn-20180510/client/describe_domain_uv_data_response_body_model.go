// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainUvDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainUvDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainUvDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainUvDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDomainUvDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainUvDataResponseBody
	GetStartTime() *string
	SetUvDataInterval(v *DescribeDomainUvDataResponseBodyUvDataInterval) *DescribeDomainUvDataResponseBody
	GetUvDataInterval() *DescribeDomainUvDataResponseBodyUvDataInterval
}

type DescribeDomainUvDataResponseBody struct {
	// The time interval. Unit: seconds.
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
	// 2019-11-29T04:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E9D3257A-1B7C-414C-90C1-8D07AC47BCAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range that was queried.
	//
	// example:
	//
	// 2019-11-29T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The number of UVs at each interval.
	UvDataInterval *DescribeDomainUvDataResponseBodyUvDataInterval `json:"UvDataInterval,omitempty" xml:"UvDataInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainUvDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainUvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainUvDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainUvDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainUvDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainUvDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainUvDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainUvDataResponseBody) GetUvDataInterval() *DescribeDomainUvDataResponseBodyUvDataInterval {
	return s.UvDataInterval
}

func (s *DescribeDomainUvDataResponseBody) SetDataInterval(v string) *DescribeDomainUvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainUvDataResponseBody) SetDomainName(v string) *DescribeDomainUvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainUvDataResponseBody) SetEndTime(v string) *DescribeDomainUvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainUvDataResponseBody) SetRequestId(v string) *DescribeDomainUvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainUvDataResponseBody) SetStartTime(v string) *DescribeDomainUvDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainUvDataResponseBody) SetUvDataInterval(v *DescribeDomainUvDataResponseBodyUvDataInterval) *DescribeDomainUvDataResponseBody {
	s.UvDataInterval = v
	return s
}

func (s *DescribeDomainUvDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainUvDataResponseBodyUvDataInterval struct {
	UsageData []*DescribeDomainUvDataResponseBodyUvDataIntervalUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDomainUvDataResponseBodyUvDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainUvDataResponseBodyUvDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainUvDataResponseBodyUvDataInterval) GetUsageData() []*DescribeDomainUvDataResponseBodyUvDataIntervalUsageData {
	return s.UsageData
}

func (s *DescribeDomainUvDataResponseBodyUvDataInterval) SetUsageData(v []*DescribeDomainUvDataResponseBodyUvDataIntervalUsageData) *DescribeDomainUvDataResponseBodyUvDataInterval {
	s.UsageData = v
	return s
}

func (s *DescribeDomainUvDataResponseBodyUvDataInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainUvDataResponseBodyUvDataIntervalUsageData struct {
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2019-11-29T00:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The number of UVs.
	//
	// example:
	//
	// 318
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainUvDataResponseBodyUvDataIntervalUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainUvDataResponseBodyUvDataIntervalUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDomainUvDataResponseBodyUvDataIntervalUsageData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainUvDataResponseBodyUvDataIntervalUsageData) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainUvDataResponseBodyUvDataIntervalUsageData) SetTimeStamp(v string) *DescribeDomainUvDataResponseBodyUvDataIntervalUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainUvDataResponseBodyUvDataIntervalUsageData) SetValue(v string) *DescribeDomainUvDataResponseBodyUvDataIntervalUsageData {
	s.Value = &v
	return s
}

func (s *DescribeDomainUvDataResponseBodyUvDataIntervalUsageData) Validate() error {
	return dara.Validate(s)
}
