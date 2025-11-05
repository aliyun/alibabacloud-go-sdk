// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeSrcBpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainRealTimeSrcBpsDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainRealTimeSrcBpsDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainRealTimeSrcBpsDataResponseBody
	GetEndTime() *string
	SetRealTimeSrcBpsDataPerInterval(v *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) *DescribeDomainRealTimeSrcBpsDataResponseBody
	GetRealTimeSrcBpsDataPerInterval() *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval
	SetRequestId(v string) *DescribeDomainRealTimeSrcBpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainRealTimeSrcBpsDataResponseBody
	GetStartTime() *string
}

type DescribeDomainRealTimeSrcBpsDataResponseBody struct {
	// The time interval between the data entries returned. Unit: seconds.
	//
	// The time granularity varies with the maximum time range per query. Valid values: 60 (1 minute), 300 (5 minutes), and 3600(1 hour). For more information, see **Usage notes**.
	//
	// example:
	//
	// 60
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
	// 2019-12-10T20:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The origin bandwidth data at each interval.
	RealTimeSrcBpsDataPerInterval *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval `json:"RealTimeSrcBpsDataPerInterval,omitempty" xml:"RealTimeSrcBpsDataPerInterval,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7CBCD6AD-B016-42E5-AE0B-B3731DE8F755
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2019-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainRealTimeSrcBpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeSrcBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBody) GetRealTimeSrcBpsDataPerInterval() *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval {
	return s.RealTimeSrcBpsDataPerInterval
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBody) SetDataInterval(v string) *DescribeDomainRealTimeSrcBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBody) SetDomainName(v string) *DescribeDomainRealTimeSrcBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBody) SetEndTime(v string) *DescribeDomainRealTimeSrcBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBody) SetRealTimeSrcBpsDataPerInterval(v *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) *DescribeDomainRealTimeSrcBpsDataResponseBody {
	s.RealTimeSrcBpsDataPerInterval = v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBody) SetRequestId(v string) *DescribeDomainRealTimeSrcBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBody) SetStartTime(v string) *DescribeDomainRealTimeSrcBpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBody) Validate() error {
	if s.RealTimeSrcBpsDataPerInterval != nil {
		if err := s.RealTimeSrcBpsDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval struct {
	DataModule []*DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) GetDataModule() []*DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) SetDataModule(v []*DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerInterval) Validate() error {
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

type DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule struct {
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2019-12-10T20:01:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The bandwidth during back-to-origin routing. Unit: bit/s.
	//
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) SetValue(v string) *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainRealTimeSrcBpsDataResponseBodyRealTimeSrcBpsDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
