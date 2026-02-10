// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainUsageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArea(v string) *DescribeDomainUsageDataResponseBody
	GetArea() *string
	SetDataInterval(v string) *DescribeDomainUsageDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainUsageDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainUsageDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDomainUsageDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainUsageDataResponseBody
	GetStartTime() *string
	SetType(v string) *DescribeDomainUsageDataResponseBody
	GetType() *string
	SetUsageDataPerInterval(v *DescribeDomainUsageDataResponseBodyUsageDataPerInterval) *DescribeDomainUsageDataResponseBody
	GetUsageDataPerInterval() *DescribeDomainUsageDataResponseBodyUsageDataPerInterval
}

type DescribeDomainUsageDataResponseBody struct {
	// The billable region where the resource usage data was generated.
	//
	// example:
	//
	// CN
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The time interval between the returned entries. Unit: seconds.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range for which the resource usage data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-12-10T21:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B955107D-E658-4E77-B913-E0AC3D31693E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range for which the resource usage data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-12-10T20:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the resource usage data.
	//
	// example:
	//
	// all
	Type                 *string                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	UsageDataPerInterval *DescribeDomainUsageDataResponseBodyUsageDataPerInterval `json:"UsageDataPerInterval,omitempty" xml:"UsageDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDomainUsageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainUsageDataResponseBody) GetArea() *string {
	return s.Area
}

func (s *DescribeDomainUsageDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainUsageDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainUsageDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainUsageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainUsageDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainUsageDataResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeDomainUsageDataResponseBody) GetUsageDataPerInterval() *DescribeDomainUsageDataResponseBodyUsageDataPerInterval {
	return s.UsageDataPerInterval
}

func (s *DescribeDomainUsageDataResponseBody) SetArea(v string) *DescribeDomainUsageDataResponseBody {
	s.Area = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBody) SetDataInterval(v string) *DescribeDomainUsageDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBody) SetDomainName(v string) *DescribeDomainUsageDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBody) SetEndTime(v string) *DescribeDomainUsageDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBody) SetRequestId(v string) *DescribeDomainUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBody) SetStartTime(v string) *DescribeDomainUsageDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBody) SetType(v string) *DescribeDomainUsageDataResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBody) SetUsageDataPerInterval(v *DescribeDomainUsageDataResponseBodyUsageDataPerInterval) *DescribeDomainUsageDataResponseBody {
	s.UsageDataPerInterval = v
	return s
}

func (s *DescribeDomainUsageDataResponseBody) Validate() error {
	if s.UsageDataPerInterval != nil {
		if err := s.UsageDataPerInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainUsageDataResponseBodyUsageDataPerInterval struct {
	DataModule []*DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDomainUsageDataResponseBodyUsageDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainUsageDataResponseBodyUsageDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDomainUsageDataResponseBodyUsageDataPerInterval) GetDataModule() []*DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDomainUsageDataResponseBodyUsageDataPerInterval) SetDataModule(v []*DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) *DescribeDomainUsageDataResponseBodyUsageDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDomainUsageDataResponseBodyUsageDataPerInterval) Validate() error {
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

type DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetValue(v string) *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
