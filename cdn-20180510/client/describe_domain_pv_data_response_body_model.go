// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainPvDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainPvDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainPvDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainPvDataResponseBody
	GetEndTime() *string
	SetPvDataInterval(v *DescribeDomainPvDataResponseBodyPvDataInterval) *DescribeDomainPvDataResponseBody
	GetPvDataInterval() *DescribeDomainPvDataResponseBodyPvDataInterval
	SetRequestId(v string) *DescribeDomainPvDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainPvDataResponseBody
	GetStartTime() *string
}

type DescribeDomainPvDataResponseBody struct {
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
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2015-11-28T04:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of PVs at each interval.
	PvDataInterval *DescribeDomainPvDataResponseBodyPvDataInterval `json:"PvDataInterval,omitempty" xml:"PvDataInterval,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// BCD7D917-76F1-442F-BB75-C810DE34C761
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2015-11-28T03:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainPvDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainPvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainPvDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainPvDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainPvDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainPvDataResponseBody) GetPvDataInterval() *DescribeDomainPvDataResponseBodyPvDataInterval {
	return s.PvDataInterval
}

func (s *DescribeDomainPvDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainPvDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainPvDataResponseBody) SetDataInterval(v string) *DescribeDomainPvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainPvDataResponseBody) SetDomainName(v string) *DescribeDomainPvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainPvDataResponseBody) SetEndTime(v string) *DescribeDomainPvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainPvDataResponseBody) SetPvDataInterval(v *DescribeDomainPvDataResponseBodyPvDataInterval) *DescribeDomainPvDataResponseBody {
	s.PvDataInterval = v
	return s
}

func (s *DescribeDomainPvDataResponseBody) SetRequestId(v string) *DescribeDomainPvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainPvDataResponseBody) SetStartTime(v string) *DescribeDomainPvDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainPvDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainPvDataResponseBodyPvDataInterval struct {
	UsageData []*DescribeDomainPvDataResponseBodyPvDataIntervalUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDomainPvDataResponseBodyPvDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainPvDataResponseBodyPvDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainPvDataResponseBodyPvDataInterval) GetUsageData() []*DescribeDomainPvDataResponseBodyPvDataIntervalUsageData {
	return s.UsageData
}

func (s *DescribeDomainPvDataResponseBodyPvDataInterval) SetUsageData(v []*DescribeDomainPvDataResponseBodyPvDataIntervalUsageData) *DescribeDomainPvDataResponseBodyPvDataInterval {
	s.UsageData = v
	return s
}

func (s *DescribeDomainPvDataResponseBodyPvDataInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainPvDataResponseBodyPvDataIntervalUsageData struct {
	// The timestamp of the returned data.
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

func (s DescribeDomainPvDataResponseBodyPvDataIntervalUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainPvDataResponseBodyPvDataIntervalUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDomainPvDataResponseBodyPvDataIntervalUsageData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainPvDataResponseBodyPvDataIntervalUsageData) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainPvDataResponseBodyPvDataIntervalUsageData) SetTimeStamp(v string) *DescribeDomainPvDataResponseBodyPvDataIntervalUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainPvDataResponseBodyPvDataIntervalUsageData) SetValue(v string) *DescribeDomainPvDataResponseBodyPvDataIntervalUsageData {
	s.Value = &v
	return s
}

func (s *DescribeDomainPvDataResponseBodyPvDataIntervalUsageData) Validate() error {
	return dara.Validate(s)
}
