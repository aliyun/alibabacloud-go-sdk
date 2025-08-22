// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainUsageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArea(v string) *DescribeDcdnDomainUsageDataResponseBody
	GetArea() *string
	SetDataInterval(v string) *DescribeDcdnDomainUsageDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainUsageDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainUsageDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeDcdnDomainUsageDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainUsageDataResponseBody
	GetStartTime() *string
	SetType(v string) *DescribeDcdnDomainUsageDataResponseBody
	GetType() *string
	SetUsageDataPerInterval(v *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerInterval) *DescribeDcdnDomainUsageDataResponseBody
	GetUsageDataPerInterval() *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerInterval
}

type DescribeDcdnDomainUsageDataResponseBody struct {
	// The billable region where the usage data was collected.
	//
	// example:
	//
	// CN
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The time interval between the data entries returned. Unit: seconds.
	//
	// example:
	//
	// /
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The accelerated domain name that was queried.
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
	// The type of the returned data.
	//
	// example:
	//
	// bps
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The traffic that was collected at each interval.
	UsageDataPerInterval *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerInterval `json:"UsageDataPerInterval,omitempty" xml:"UsageDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainUsageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUsageDataResponseBody) GetArea() *string {
	return s.Area
}

func (s *DescribeDcdnDomainUsageDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainUsageDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainUsageDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainUsageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainUsageDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainUsageDataResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeDcdnDomainUsageDataResponseBody) GetUsageDataPerInterval() *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerInterval {
	return s.UsageDataPerInterval
}

func (s *DescribeDcdnDomainUsageDataResponseBody) SetArea(v string) *DescribeDcdnDomainUsageDataResponseBody {
	s.Area = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainUsageDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainUsageDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainUsageDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainUsageDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBody) SetType(v string) *DescribeDcdnDomainUsageDataResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBody) SetUsageDataPerInterval(v *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerInterval) *DescribeDcdnDomainUsageDataResponseBody {
	s.UsageDataPerInterval = v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainUsageDataResponseBodyUsageDataPerInterval struct {
	DataModule []*DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainUsageDataResponseBodyUsageDataPerInterval) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainUsageDataResponseBodyUsageDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerInterval) GetDataModule() []*DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	return s.DataModule
}

func (s *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerInterval) SetDataModule(v []*DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerInterval {
	s.DataModule = v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerInterval) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule struct {
	// The time of the peak bandwidth value if the **Field*	- parameter in the request is set to **bps**. Otherwise, this parameter returns the same value as the **TimeStamp*	- parameter.
	//
	// example:
	//
	// 2015-12-10T21:30:00Z
	PeakTime *string `json:"PeakTime,omitempty" xml:"PeakTime,omitempty"`
	// The data usage in a specific scenario.
	//
	// >  This parameter indicates the data usage in a specific scenario. If no special billable item is specified, ignore this parameter.
	//
	// example:
	//
	// 4233041**
	SpecialValue *string `json:"SpecialValue,omitempty" xml:"SpecialValue,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2015-12-10T21:30:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The usage.
	//
	// example:
	//
	// 4233041**
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) GetPeakTime() *string {
	return s.PeakTime
}

func (s *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) GetSpecialValue() *string {
	return s.SpecialValue
}

func (s *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) GetValue() *string {
	return s.Value
}

func (s *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetPeakTime(v string) *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.PeakTime = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetSpecialValue(v string) *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.SpecialValue = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetValue(v string) *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.Value = &v
	return s
}

func (s *DescribeDcdnDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) Validate() error {
	return dara.Validate(s)
}
