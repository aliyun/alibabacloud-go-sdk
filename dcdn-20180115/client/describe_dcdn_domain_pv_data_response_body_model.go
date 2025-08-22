// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainPvDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainPvDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainPvDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainPvDataResponseBody
	GetEndTime() *string
	SetPvDataInterval(v *DescribeDcdnDomainPvDataResponseBodyPvDataInterval) *DescribeDcdnDomainPvDataResponseBody
	GetPvDataInterval() *DescribeDcdnDomainPvDataResponseBodyPvDataInterval
	SetRequestId(v string) *DescribeDcdnDomainPvDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainPvDataResponseBody
	GetStartTime() *string
}

type DescribeDcdnDomainPvDataResponseBody struct {
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
	// 2019-11-29T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of PVs at each interval.
	PvDataInterval *DescribeDcdnDomainPvDataResponseBodyPvDataInterval `json:"PvDataInterval,omitempty" xml:"PvDataInterval,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BCD7D917-76F1-442F-BB75-C810DE34C761
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range that was queried.
	//
	// example:
	//
	// 2019-11-28T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainPvDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainPvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainPvDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainPvDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainPvDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainPvDataResponseBody) GetPvDataInterval() *DescribeDcdnDomainPvDataResponseBodyPvDataInterval {
	return s.PvDataInterval
}

func (s *DescribeDcdnDomainPvDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainPvDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainPvDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainPvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainPvDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainPvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainPvDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainPvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainPvDataResponseBody) SetPvDataInterval(v *DescribeDcdnDomainPvDataResponseBodyPvDataInterval) *DescribeDcdnDomainPvDataResponseBody {
	s.PvDataInterval = v
	return s
}

func (s *DescribeDcdnDomainPvDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainPvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainPvDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainPvDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainPvDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainPvDataResponseBodyPvDataInterval struct {
	UsageData []*DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainPvDataResponseBodyPvDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainPvDataResponseBodyPvDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainPvDataResponseBodyPvDataInterval) GetUsageData() []*DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData {
	return s.UsageData
}

func (s *DescribeDcdnDomainPvDataResponseBodyPvDataInterval) SetUsageData(v []*DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData) *DescribeDcdnDomainPvDataResponseBodyPvDataInterval {
	s.UsageData = v
	return s
}

func (s *DescribeDcdnDomainPvDataResponseBodyPvDataInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData struct {
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2015-11-28T03:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The number of PVs.
	//
	// example:
	//
	// 9292
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData) GetValue() *string {
	return s.Value
}

func (s *DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData) SetTimeStamp(v string) *DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData) SetValue(v string) *DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData {
	s.Value = &v
	return s
}

func (s *DescribeDcdnDomainPvDataResponseBodyPvDataIntervalUsageData) Validate() error {
	return dara.Validate(s)
}
