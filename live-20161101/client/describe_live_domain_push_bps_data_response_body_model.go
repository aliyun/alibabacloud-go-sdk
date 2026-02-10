// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainPushBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBpsDataPerInterval(v *DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerInterval) *DescribeLiveDomainPushBpsDataResponseBody
	GetBpsDataPerInterval() *DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerInterval
	SetDataInterval(v string) *DescribeLiveDomainPushBpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeLiveDomainPushBpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainPushBpsDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeLiveDomainPushBpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLiveDomainPushBpsDataResponseBody
	GetStartTime() *string
}

type DescribeLiveDomainPushBpsDataResponseBody struct {
	BpsDataPerInterval *DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" type:"Struct"`
	// The time granularity of the query.
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
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainPushBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPushBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPushBpsDataResponseBody) GetBpsDataPerInterval() *DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerInterval {
	return s.BpsDataPerInterval
}

func (s *DescribeLiveDomainPushBpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeLiveDomainPushBpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainPushBpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainPushBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainPushBpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainPushBpsDataResponseBody) SetBpsDataPerInterval(v *DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerInterval) *DescribeLiveDomainPushBpsDataResponseBody {
	s.BpsDataPerInterval = v
	return s
}

func (s *DescribeLiveDomainPushBpsDataResponseBody) SetDataInterval(v string) *DescribeLiveDomainPushBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataResponseBody) SetDomainName(v string) *DescribeLiveDomainPushBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataResponseBody) SetEndTime(v string) *DescribeLiveDomainPushBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataResponseBody) SetRequestId(v string) *DescribeLiveDomainPushBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataResponseBody) SetStartTime(v string) *DescribeLiveDomainPushBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataResponseBody) Validate() error {
	if s.BpsDataPerInterval != nil {
		if err := s.BpsDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerInterval struct {
	DataModule []*DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerInterval) GetDataModule() []*DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerInterval) SetDataModule(v []*DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerIntervalDataModule) *DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerInterval) Validate() error {
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

type DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerIntervalDataModule struct {
	BpsValue  *string `json:"BpsValue,omitempty" xml:"BpsValue,omitempty"`
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerIntervalDataModule) GetBpsValue() *string {
	return s.BpsValue
}

func (s *DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerIntervalDataModule) SetBpsValue(v string) *DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.BpsValue = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainPushBpsDataResponseBodyBpsDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
