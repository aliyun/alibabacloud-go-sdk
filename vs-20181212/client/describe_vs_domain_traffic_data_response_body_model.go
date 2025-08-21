// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVsDomainTrafficDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVsDomainTrafficDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsDomainTrafficDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeVsDomainTrafficDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVsDomainTrafficDataResponseBody
	GetStartTime() *string
	SetTrafficDataPerInterval(v *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerInterval) *DescribeVsDomainTrafficDataResponseBody
	GetTrafficDataPerInterval() *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerInterval
}

type DescribeVsDomainTrafficDataResponseBody struct {
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
	// 2021-09-20T07:10:42Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2021-10-25T16:00:00Z
	StartTime              *string                                                        `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TrafficDataPerInterval *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerInterval `json:"TrafficDataPerInterval,omitempty" xml:"TrafficDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeVsDomainTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainTrafficDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVsDomainTrafficDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDomainTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsDomainTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDomainTrafficDataResponseBody) GetTrafficDataPerInterval() *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerInterval {
	return s.TrafficDataPerInterval
}

func (s *DescribeVsDomainTrafficDataResponseBody) SetDataInterval(v string) *DescribeVsDomainTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVsDomainTrafficDataResponseBody) SetDomainName(v string) *DescribeVsDomainTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainTrafficDataResponseBody) SetEndTime(v string) *DescribeVsDomainTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainTrafficDataResponseBody) SetRequestId(v string) *DescribeVsDomainTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainTrafficDataResponseBody) SetStartTime(v string) *DescribeVsDomainTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainTrafficDataResponseBody) SetTrafficDataPerInterval(v *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerInterval) *DescribeVsDomainTrafficDataResponseBody {
	s.TrafficDataPerInterval = v
	return s
}

func (s *DescribeVsDomainTrafficDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVsDomainTrafficDataResponseBodyTrafficDataPerInterval struct {
	DataModule []*DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainTrafficDataResponseBodyTrafficDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainTrafficDataResponseBodyTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerInterval) GetDataModule() []*DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerInterval) SetDataModule(v []*DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule struct {
	// example:
	//
	// 2021-09-20T07:10:42Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// example:
	//
	// 100
	TrafficValue *string `json:"TrafficValue,omitempty" xml:"TrafficValue,omitempty"`
}

func (s DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetTrafficValue() *string {
	return s.TrafficValue
}

func (s *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTrafficValue(v string) *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TrafficValue = &v
	return s
}

func (s *DescribeVsDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
