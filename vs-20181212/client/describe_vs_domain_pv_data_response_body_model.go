// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainPvDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVsDomainPvDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVsDomainPvDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsDomainPvDataResponseBody
	GetEndTime() *string
	SetPvDataInterval(v *DescribeVsDomainPvDataResponseBodyPvDataInterval) *DescribeVsDomainPvDataResponseBody
	GetPvDataInterval() *DescribeVsDomainPvDataResponseBodyPvDataInterval
	SetRequestId(v string) *DescribeVsDomainPvDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVsDomainPvDataResponseBody
	GetStartTime() *string
}

type DescribeVsDomainPvDataResponseBody struct {
	// example:
	//
	// 3600
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 2021-11-12T15:59:59Z
	EndTime        *string                                           `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PvDataInterval *DescribeVsDomainPvDataResponseBodyPvDataInterval `json:"PvDataInterval,omitempty" xml:"PvDataInterval,omitempty" type:"Struct"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2021-11-22T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainPvDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainPvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVsDomainPvDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainPvDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDomainPvDataResponseBody) GetPvDataInterval() *DescribeVsDomainPvDataResponseBodyPvDataInterval {
	return s.PvDataInterval
}

func (s *DescribeVsDomainPvDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsDomainPvDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDomainPvDataResponseBody) SetDataInterval(v string) *DescribeVsDomainPvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVsDomainPvDataResponseBody) SetDomainName(v string) *DescribeVsDomainPvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainPvDataResponseBody) SetEndTime(v string) *DescribeVsDomainPvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainPvDataResponseBody) SetPvDataInterval(v *DescribeVsDomainPvDataResponseBodyPvDataInterval) *DescribeVsDomainPvDataResponseBody {
	s.PvDataInterval = v
	return s
}

func (s *DescribeVsDomainPvDataResponseBody) SetRequestId(v string) *DescribeVsDomainPvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainPvDataResponseBody) SetStartTime(v string) *DescribeVsDomainPvDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainPvDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVsDomainPvDataResponseBodyPvDataInterval struct {
	UsageData []*DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainPvDataResponseBodyPvDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainPvDataResponseBodyPvDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvDataResponseBodyPvDataInterval) GetUsageData() []*DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData {
	return s.UsageData
}

func (s *DescribeVsDomainPvDataResponseBodyPvDataInterval) SetUsageData(v []*DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData) *DescribeVsDomainPvDataResponseBodyPvDataInterval {
	s.UsageData = v
	return s
}

func (s *DescribeVsDomainPvDataResponseBodyPvDataInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData struct {
	// example:
	//
	// 2021-11-22T00:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// example:
	//
	// 100
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData) GetValue() *string {
	return s.Value
}

func (s *DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData) SetTimeStamp(v string) *DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData) SetValue(v string) *DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData {
	s.Value = &v
	return s
}

func (s *DescribeVsDomainPvDataResponseBodyPvDataIntervalUsageData) Validate() error {
	return dara.Validate(s)
}
