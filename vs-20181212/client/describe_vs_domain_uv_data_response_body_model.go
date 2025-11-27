// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainUvDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVsDomainUvDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVsDomainUvDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsDomainUvDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeVsDomainUvDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVsDomainUvDataResponseBody
	GetStartTime() *string
	SetUvDataInterval(v *DescribeVsDomainUvDataResponseBodyUvDataInterval) *DescribeVsDomainUvDataResponseBody
	GetUvDataInterval() *DescribeVsDomainUvDataResponseBodyUvDataInterval
}

type DescribeVsDomainUvDataResponseBody struct {
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
	// 2015-11-30T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2015-11-29T00:00:00Z
	StartTime      *string                                           `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UvDataInterval *DescribeVsDomainUvDataResponseBodyUvDataInterval `json:"UvDataInterval,omitempty" xml:"UvDataInterval,omitempty" type:"Struct"`
}

func (s DescribeVsDomainUvDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainUvDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainUvDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVsDomainUvDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainUvDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDomainUvDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsDomainUvDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDomainUvDataResponseBody) GetUvDataInterval() *DescribeVsDomainUvDataResponseBodyUvDataInterval {
	return s.UvDataInterval
}

func (s *DescribeVsDomainUvDataResponseBody) SetDataInterval(v string) *DescribeVsDomainUvDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVsDomainUvDataResponseBody) SetDomainName(v string) *DescribeVsDomainUvDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainUvDataResponseBody) SetEndTime(v string) *DescribeVsDomainUvDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainUvDataResponseBody) SetRequestId(v string) *DescribeVsDomainUvDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainUvDataResponseBody) SetStartTime(v string) *DescribeVsDomainUvDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainUvDataResponseBody) SetUvDataInterval(v *DescribeVsDomainUvDataResponseBodyUvDataInterval) *DescribeVsDomainUvDataResponseBody {
	s.UvDataInterval = v
	return s
}

func (s *DescribeVsDomainUvDataResponseBody) Validate() error {
	if s.UvDataInterval != nil {
		if err := s.UvDataInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVsDomainUvDataResponseBodyUvDataInterval struct {
	UsageData []*DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainUvDataResponseBodyUvDataInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainUvDataResponseBodyUvDataInterval) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainUvDataResponseBodyUvDataInterval) GetUsageData() []*DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData {
	return s.UsageData
}

func (s *DescribeVsDomainUvDataResponseBodyUvDataInterval) SetUsageData(v []*DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData) *DescribeVsDomainUvDataResponseBodyUvDataInterval {
	s.UsageData = v
	return s
}

func (s *DescribeVsDomainUvDataResponseBodyUvDataInterval) Validate() error {
	if s.UsageData != nil {
		for _, item := range s.UsageData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData struct {
	// example:
	//
	// 2015-11-29T15:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// example:
	//
	// 100
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData) GetValue() *string {
	return s.Value
}

func (s *DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData) SetTimeStamp(v string) *DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData) SetValue(v string) *DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData {
	s.Value = &v
	return s
}

func (s *DescribeVsDomainUvDataResponseBodyUvDataIntervalUsageData) Validate() error {
	return dara.Validate(s)
}
