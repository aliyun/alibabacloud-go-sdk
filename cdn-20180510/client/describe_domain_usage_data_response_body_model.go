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
	// The ID of the billable region where the data was collected.
	//
	// example:
	//
	// CN
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The time interval between the data entries returned. Unit: seconds.
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
	// 2015-12-10T22:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B955107D-E658-4E77-B913-E0AC3D31693E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range during which data was queried.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of content.
	//
	// example:
	//
	// static
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The resource usage that was collected at each interval.
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
	// The time of the peak bandwidth value if the **Field*	- parameter in the request is set to **bps**. Otherwise, this parameter returns the same value as the **TimeStamp*	- parameter.
	//
	// example:
	//
	// 2015-12-10T21:30:00Z
	PeakTime *string `json:"PeakTime,omitempty" xml:"PeakTime,omitempty"`
	// The resource usage in a specific scenario.
	//
	// > SpecialValue indicates the data usage in a specific scenario. If no special billable item is specified, ignore this parameter.
	//
	// example:
	//
	// 423304182
	SpecialValue *string `json:"SpecialValue,omitempty" xml:"SpecialValue,omitempty"`
	// The timestamp of the data returned.
	//
	// > **TimeStamp*	- indicates the timestamp of the data returned at each interval.
	//
	// example:
	//
	// 2015-12-10T21:30:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The amount of resource usage.
	//
	// example:
	//
	// 423304182
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) GetPeakTime() *string {
	return s.PeakTime
}

func (s *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) GetSpecialValue() *string {
	return s.SpecialValue
}

func (s *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetPeakTime(v string) *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.PeakTime = &v
	return s
}

func (s *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetSpecialValue(v string) *DescribeDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.SpecialValue = &v
	return s
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
