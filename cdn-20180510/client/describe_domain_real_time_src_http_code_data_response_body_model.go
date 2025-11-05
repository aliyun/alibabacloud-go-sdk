// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeSrcHttpCodeDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBody
	GetEndTime() *string
	SetRealTimeSrcHttpCodeData(v *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) *DescribeDomainRealTimeSrcHttpCodeDataResponseBody
	GetRealTimeSrcHttpCodeData() *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData
	SetRequestId(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBody
	GetStartTime() *string
}

type DescribeDomainRealTimeSrcHttpCodeDataResponseBody struct {
	// The time interval between the data entries returned. Unit: seconds.
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
	// 2019-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The proportions of HTTP status codes at each time interval.
	RealTimeSrcHttpCodeData *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData `json:"RealTimeSrcHttpCodeData,omitempty" xml:"RealTimeSrcHttpCodeData,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// BC858082-736F-4A25-867B-E5B67C85ACF7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2019-11-30T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) GetRealTimeSrcHttpCodeData() *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData {
	return s.RealTimeSrcHttpCodeData
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) SetDomainName(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) SetEndTime(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) SetRealTimeSrcHttpCodeData(v *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) *DescribeDomainRealTimeSrcHttpCodeDataResponseBody {
	s.RealTimeSrcHttpCodeData = v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) SetRequestId(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) SetStartTime(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBody) Validate() error {
	if s.RealTimeSrcHttpCodeData != nil {
		if err := s.RealTimeSrcHttpCodeData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData struct {
	UsageData []*DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) GetUsageData() []*DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData {
	return s.UsageData
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) SetUsageData(v []*DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData {
	s.UsageData = v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) Validate() error {
	if s.UsageData != nil {
		for _, item := range s.UsageData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData struct {
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2015-11-30T05:40:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The proportions of the HTTP status codes.
	Value *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) GetValue() *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue {
	return s.Value
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) SetTimeStamp(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) SetValue(v *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData {
	s.Value = v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) Validate() error {
	if s.Value != nil {
		if err := s.Value.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue struct {
	RealTimeSrcCodeProportionData []*DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData `json:"RealTimeSrcCodeProportionData,omitempty" xml:"RealTimeSrcCodeProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) GetRealTimeSrcCodeProportionData() []*DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData {
	return s.RealTimeSrcCodeProportionData
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) SetRealTimeSrcCodeProportionData(v []*DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue {
	s.RealTimeSrcCodeProportionData = v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) Validate() error {
	if s.RealTimeSrcCodeProportionData != nil {
		for _, item := range s.RealTimeSrcCodeProportionData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The count of each HTTP status code.
	//
	// example:
	//
	// 100
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The proportion of the HTTP status code.
	//
	// example:
	//
	// 0.62015503875969
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) GetCode() *string {
	return s.Code
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) GetCount() *string {
	return s.Count
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) GetProportion() *string {
	return s.Proportion
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) SetCode(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData {
	s.Code = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) SetCount(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData {
	s.Count = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) SetProportion(v string) *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) Validate() error {
	return dara.Validate(s)
}
