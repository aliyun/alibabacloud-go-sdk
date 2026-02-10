// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainRealTimeBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeLiveDomainRealTimeBpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeLiveDomainRealTimeBpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainRealTimeBpsDataResponseBody
	GetEndTime() *string
	SetRealTimeBpsDataPerInterval(v *DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerInterval) *DescribeLiveDomainRealTimeBpsDataResponseBody
	GetRealTimeBpsDataPerInterval() *DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerInterval
	SetRequestId(v string) *DescribeLiveDomainRealTimeBpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLiveDomainRealTimeBpsDataResponseBody
	GetStartTime() *string
}

type DescribeLiveDomainRealTimeBpsDataResponseBody struct {
	// The interval at which data was queried.
	//
	// example:
	//
	// 60
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The streaming domain name or names that was queried.
	//
	// example:
	//
	// example.com,example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range in which data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-11-30T05:40:00Z
	EndTime                    *string                                                                  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RealTimeBpsDataPerInterval *DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerInterval `json:"RealTimeBpsDataPerInterval,omitempty" xml:"RealTimeBpsDataPerInterval,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// BC858082-736F-4A25-867B-E5B67C85ACF7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range in which data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-11-30T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainRealTimeBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealTimeBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBody) GetRealTimeBpsDataPerInterval() *DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerInterval {
	return s.RealTimeBpsDataPerInterval
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBody) SetDataInterval(v string) *DescribeLiveDomainRealTimeBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBody) SetDomainName(v string) *DescribeLiveDomainRealTimeBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBody) SetEndTime(v string) *DescribeLiveDomainRealTimeBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBody) SetRealTimeBpsDataPerInterval(v *DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerInterval) *DescribeLiveDomainRealTimeBpsDataResponseBody {
	s.RealTimeBpsDataPerInterval = v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBody) SetRequestId(v string) *DescribeLiveDomainRealTimeBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBody) SetStartTime(v string) *DescribeLiveDomainRealTimeBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBody) Validate() error {
	if s.RealTimeBpsDataPerInterval != nil {
		if err := s.RealTimeBpsDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerInterval struct {
	DataModule []*DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerInterval) GetDataModule() []*DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerInterval) SetDataModule(v []*DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerIntervalDataModule) *DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerInterval) Validate() error {
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

type DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerIntervalDataModule) SetValue(v string) *DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeLiveDomainRealTimeBpsDataResponseBodyRealTimeBpsDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
