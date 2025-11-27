// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBpsDataPerInterval(v *DescribeVsDomainBpsDataResponseBodyBpsDataPerInterval) *DescribeVsDomainBpsDataResponseBody
	GetBpsDataPerInterval() *DescribeVsDomainBpsDataResponseBodyBpsDataPerInterval
	SetDataInterval(v string) *DescribeVsDomainBpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVsDomainBpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsDomainBpsDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeVsDomainBpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVsDomainBpsDataResponseBody
	GetStartTime() *string
}

type DescribeVsDomainBpsDataResponseBody struct {
	BpsDataPerInterval *DescribeVsDomainBpsDataResponseBodyBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" type:"Struct"`
	// example:
	//
	// 2100
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 2021-10-01T07:10:48Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2021-09-18T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVsDomainBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainBpsDataResponseBody) GetBpsDataPerInterval() *DescribeVsDomainBpsDataResponseBodyBpsDataPerInterval {
	return s.BpsDataPerInterval
}

func (s *DescribeVsDomainBpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVsDomainBpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainBpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsDomainBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsDomainBpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsDomainBpsDataResponseBody) SetBpsDataPerInterval(v *DescribeVsDomainBpsDataResponseBodyBpsDataPerInterval) *DescribeVsDomainBpsDataResponseBody {
	s.BpsDataPerInterval = v
	return s
}

func (s *DescribeVsDomainBpsDataResponseBody) SetDataInterval(v string) *DescribeVsDomainBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVsDomainBpsDataResponseBody) SetDomainName(v string) *DescribeVsDomainBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainBpsDataResponseBody) SetEndTime(v string) *DescribeVsDomainBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVsDomainBpsDataResponseBody) SetRequestId(v string) *DescribeVsDomainBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsDomainBpsDataResponseBody) SetStartTime(v string) *DescribeVsDomainBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVsDomainBpsDataResponseBody) Validate() error {
	if s.BpsDataPerInterval != nil {
		if err := s.BpsDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVsDomainBpsDataResponseBodyBpsDataPerInterval struct {
	DataModule []*DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVsDomainBpsDataResponseBodyBpsDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainBpsDataResponseBodyBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainBpsDataResponseBodyBpsDataPerInterval) GetDataModule() []*DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeVsDomainBpsDataResponseBodyBpsDataPerInterval) SetDataModule(v []*DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) *DescribeVsDomainBpsDataResponseBodyBpsDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeVsDomainBpsDataResponseBodyBpsDataPerInterval) Validate() error {
	if s.DataModule != nil {
		for _, item := range s.DataModule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule struct {
	// example:
	//
	// 1000
	BpsValue *string `json:"BpsValue,omitempty" xml:"BpsValue,omitempty"`
	// example:
	//
	// 2015-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetBpsValue() *string {
	return s.BpsValue
}

func (s *DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetBpsValue(v string) *DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.BpsValue = &v
	return s
}

func (s *DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVsDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
