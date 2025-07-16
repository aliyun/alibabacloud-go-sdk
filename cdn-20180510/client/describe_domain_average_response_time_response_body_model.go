// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainAverageResponseTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvgRTPerInterval(v *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerInterval) *DescribeDomainAverageResponseTimeResponseBody
	GetAvgRTPerInterval() *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerInterval
	SetDataInterval(v string) *DescribeDomainAverageResponseTimeResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainAverageResponseTimeResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainAverageResponseTimeResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDomainAverageResponseTimeResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainAverageResponseTimeResponseBody
	GetStartTime() *string
}

type DescribeDomainAverageResponseTimeResponseBody struct {
	// The average response time data for time intervals.
	AvgRTPerInterval *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerInterval `json:"AvgRTPerInterval,omitempty" xml:"AvgRTPerInterval,omitempty" type:"Struct"`
	// The time interval between the data entries returned.
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
	// 2019-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3C6CCEC4-6B88-4D4A-93E4-D47B3D92CF8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2019-11-30T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainAverageResponseTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAverageResponseTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainAverageResponseTimeResponseBody) GetAvgRTPerInterval() *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerInterval {
	return s.AvgRTPerInterval
}

func (s *DescribeDomainAverageResponseTimeResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainAverageResponseTimeResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainAverageResponseTimeResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainAverageResponseTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainAverageResponseTimeResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainAverageResponseTimeResponseBody) SetAvgRTPerInterval(v *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerInterval) *DescribeDomainAverageResponseTimeResponseBody {
	s.AvgRTPerInterval = v
	return s
}

func (s *DescribeDomainAverageResponseTimeResponseBody) SetDataInterval(v string) *DescribeDomainAverageResponseTimeResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeResponseBody) SetDomainName(v string) *DescribeDomainAverageResponseTimeResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeResponseBody) SetEndTime(v string) *DescribeDomainAverageResponseTimeResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeResponseBody) SetRequestId(v string) *DescribeDomainAverageResponseTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeResponseBody) SetStartTime(v string) *DescribeDomainAverageResponseTimeResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainAverageResponseTimeResponseBodyAvgRTPerInterval struct {
	DataModule []*DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainAverageResponseTimeResponseBodyAvgRTPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAverageResponseTimeResponseBodyAvgRTPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerInterval) GetDataModule() []*DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerInterval) SetDataModule(v []*DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule) *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule struct {
	// The timestamp of the data returned.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The average response time.
	//
	// example:
	//
	// 3
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule) SetValue(v string) *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainAverageResponseTimeResponseBodyAvgRTPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
