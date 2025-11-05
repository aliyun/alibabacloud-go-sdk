// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainHttpCodeDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainHttpCodeDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainHttpCodeDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainHttpCodeDataResponseBody
	GetEndTime() *string
	SetHttpCodeData(v *DescribeDomainHttpCodeDataResponseBodyHttpCodeData) *DescribeDomainHttpCodeDataResponseBody
	GetHttpCodeData() *DescribeDomainHttpCodeDataResponseBodyHttpCodeData
	SetRequestId(v string) *DescribeDomainHttpCodeDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainHttpCodeDataResponseBody
	GetStartTime() *string
}

type DescribeDomainHttpCodeDataResponseBody struct {
	// The time interval.
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
	// 2021-06-29T05:45:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The proportions of HTTP status codes at each time interval.
	HttpCodeData *DescribeDomainHttpCodeDataResponseBodyHttpCodeData `json:"HttpCodeData,omitempty" xml:"HttpCodeData,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// BC858082-736F-4A25-867B-E5B67C85ACF7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range during which data was queried.
	//
	// example:
	//
	// 2021-06-29T05:30:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainHttpCodeDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainHttpCodeDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainHttpCodeDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainHttpCodeDataResponseBody) GetHttpCodeData() *DescribeDomainHttpCodeDataResponseBodyHttpCodeData {
	return s.HttpCodeData
}

func (s *DescribeDomainHttpCodeDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainHttpCodeDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeDomainHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBody) SetDomainName(v string) *DescribeDomainHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBody) SetEndTime(v string) *DescribeDomainHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBody) SetHttpCodeData(v *DescribeDomainHttpCodeDataResponseBodyHttpCodeData) *DescribeDomainHttpCodeDataResponseBody {
	s.HttpCodeData = v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBody) SetRequestId(v string) *DescribeDomainHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBody) SetStartTime(v string) *DescribeDomainHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBody) Validate() error {
	if s.HttpCodeData != nil {
		if err := s.HttpCodeData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainHttpCodeDataResponseBodyHttpCodeData struct {
	UsageData []*DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeData) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeData) GetUsageData() []*DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData {
	return s.UsageData
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeData) SetUsageData(v []*DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) *DescribeDomainHttpCodeDataResponseBodyHttpCodeData {
	s.UsageData = v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeData) Validate() error {
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

type DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData struct {
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2021-06-29T05:40:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The information about the HTTP status codes.
	Value *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) GetValue() *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue {
	return s.Value
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) SetTimeStamp(v string) *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) SetValue(v *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData {
	s.Value = v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) Validate() error {
	if s.Value != nil {
		if err := s.Value.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue struct {
	CodeProportionData []*DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData `json:"CodeProportionData,omitempty" xml:"CodeProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) GetCodeProportionData() []*DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	return s.CodeProportionData
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) SetCodeProportionData(v []*DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue {
	s.CodeProportionData = v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) Validate() error {
	if s.CodeProportionData != nil {
		for _, item := range s.CodeProportionData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData struct {
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
	// 300
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The proportion of the HTTP status code.
	//
	// example:
	//
	// 66.046511627907
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) GetCode() *string {
	return s.Code
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) GetCount() *string {
	return s.Count
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) GetProportion() *string {
	return s.Proportion
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) SetCode(v string) *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	s.Code = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) SetCount(v string) *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	s.Count = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) SetProportion(v string) *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) Validate() error {
	return dara.Validate(s)
}
