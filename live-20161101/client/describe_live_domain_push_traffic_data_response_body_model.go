// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainPushTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeLiveDomainPushTrafficDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeLiveDomainPushTrafficDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainPushTrafficDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeLiveDomainPushTrafficDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLiveDomainPushTrafficDataResponseBody
	GetStartTime() *string
	SetTrafficDataPerInterval(v *DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerInterval) *DescribeLiveDomainPushTrafficDataResponseBody
	GetTrafficDataPerInterval() *DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerInterval
}

type DescribeLiveDomainPushTrafficDataResponseBody struct {
	// The time granularity.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The ingest domain.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range during which the data was queried.
	//
	// example:
	//
	// 2017-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3C6CCEC4-6B88-4D4A-93E4-D47B3D92CF8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range during which the data was queried.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	StartTime              *string                                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TrafficDataPerInterval *DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerInterval `json:"TrafficDataPerInterval,omitempty" xml:"TrafficDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeLiveDomainPushTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPushTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPushTrafficDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeLiveDomainPushTrafficDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainPushTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainPushTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainPushTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainPushTrafficDataResponseBody) GetTrafficDataPerInterval() *DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerInterval {
	return s.TrafficDataPerInterval
}

func (s *DescribeLiveDomainPushTrafficDataResponseBody) SetDataInterval(v string) *DescribeLiveDomainPushTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataResponseBody) SetDomainName(v string) *DescribeLiveDomainPushTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataResponseBody) SetEndTime(v string) *DescribeLiveDomainPushTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataResponseBody) SetRequestId(v string) *DescribeLiveDomainPushTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataResponseBody) SetStartTime(v string) *DescribeLiveDomainPushTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataResponseBody) SetTrafficDataPerInterval(v *DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerInterval) *DescribeLiveDomainPushTrafficDataResponseBody {
	s.TrafficDataPerInterval = v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataResponseBody) Validate() error {
	if s.TrafficDataPerInterval != nil {
		if err := s.TrafficDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerInterval struct {
	DataModule []*DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerInterval) GetDataModule() []*DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerInterval) SetDataModule(v []*DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerIntervalDataModule) *DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerInterval) Validate() error {
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

type DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerIntervalDataModule struct {
	TimeStamp    *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	TrafficValue *string `json:"TrafficValue,omitempty" xml:"TrafficValue,omitempty"`
}

func (s DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GetTrafficValue() *string {
	return s.TrafficValue
}

func (s *DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTrafficValue(v string) *DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TrafficValue = &v
	return s
}

func (s *DescribeLiveDomainPushTrafficDataResponseBodyTrafficDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
