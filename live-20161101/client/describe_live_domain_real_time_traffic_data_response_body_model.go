// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainRealTimeTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeLiveDomainRealTimeTrafficDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeLiveDomainRealTimeTrafficDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainRealTimeTrafficDataResponseBody
	GetEndTime() *string
	SetRealTimeTrafficDataPerInterval(v *DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) *DescribeLiveDomainRealTimeTrafficDataResponseBody
	GetRealTimeTrafficDataPerInterval() *DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval
	SetRequestId(v string) *DescribeLiveDomainRealTimeTrafficDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLiveDomainRealTimeTrafficDataResponseBody
	GetStartTime() *string
}

type DescribeLiveDomainRealTimeTrafficDataResponseBody struct {
	// The time interval between the entries returned. Unit: seconds
	//
	// example:
	//
	// 60
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The streaming domain.
	//
	// example:
	//
	// example.com,example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range during which the data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-12-10T15:01:00Z
	EndTime                        *string                                                                          `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RealTimeTrafficDataPerInterval *DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval `json:"RealTimeTrafficDataPerInterval,omitempty" xml:"RealTimeTrafficDataPerInterval,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A666D44F-19D6-490E-97CF-1A64AB962C57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range during which the data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-12-10T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainRealTimeTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealTimeTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBody) GetRealTimeTrafficDataPerInterval() *DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval {
	return s.RealTimeTrafficDataPerInterval
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBody) SetDataInterval(v string) *DescribeLiveDomainRealTimeTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBody) SetDomainName(v string) *DescribeLiveDomainRealTimeTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBody) SetEndTime(v string) *DescribeLiveDomainRealTimeTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBody) SetRealTimeTrafficDataPerInterval(v *DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) *DescribeLiveDomainRealTimeTrafficDataResponseBody {
	s.RealTimeTrafficDataPerInterval = v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBody) SetRequestId(v string) *DescribeLiveDomainRealTimeTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBody) SetStartTime(v string) *DescribeLiveDomainRealTimeTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBody) Validate() error {
	if s.RealTimeTrafficDataPerInterval != nil {
		if err := s.RealTimeTrafficDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval struct {
	DataModule []*DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) GetDataModule() []*DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) SetDataModule(v []*DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) *DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerInterval) Validate() error {
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

type DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeLiveDomainRealTimeTrafficDataResponseBodyRealTimeTrafficDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
