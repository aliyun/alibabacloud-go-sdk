// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSrcHttpCodeDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDomainSrcHttpCodeDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDomainSrcHttpCodeDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDomainSrcHttpCodeDataResponseBody
	GetEndTime() *string
	SetHttpCodeData(v *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeData) *DescribeDomainSrcHttpCodeDataResponseBody
	GetHttpCodeData() *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeData
	SetRequestId(v string) *DescribeDomainSrcHttpCodeDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainSrcHttpCodeDataResponseBody
	GetStartTime() *string
}

type DescribeDomainSrcHttpCodeDataResponseBody struct {
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
	// example.com,example.org
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2015-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The proportions of HTTP status codes at each time interval.
	HttpCodeData *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeData `json:"HttpCodeData,omitempty" xml:"HttpCodeData,omitempty" type:"Struct"`
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
	// 2015-11-30T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainSrcHttpCodeDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcHttpCodeDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDomainSrcHttpCodeDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainSrcHttpCodeDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainSrcHttpCodeDataResponseBody) GetHttpCodeData() *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeData {
	return s.HttpCodeData
}

func (s *DescribeDomainSrcHttpCodeDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainSrcHttpCodeDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainSrcHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeDomainSrcHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBody) SetDomainName(v string) *DescribeDomainSrcHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBody) SetEndTime(v string) *DescribeDomainSrcHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBody) SetHttpCodeData(v *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeData) *DescribeDomainSrcHttpCodeDataResponseBody {
	s.HttpCodeData = v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBody) SetRequestId(v string) *DescribeDomainSrcHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBody) SetStartTime(v string) *DescribeDomainSrcHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBody) Validate() error {
	if s.HttpCodeData != nil {
		if err := s.HttpCodeData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeData struct {
	UsageData []*DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeData) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeData) GetUsageData() []*DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData {
	return s.UsageData
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeData) SetUsageData(v []*DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData) *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeData {
	s.UsageData = v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeData) Validate() error {
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

type DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData struct {
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2015-11-30T05:30:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The proportions of the HTTP status codes.
	Value *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData) GetValue() *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValue {
	return s.Value
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData) SetTimeStamp(v string) *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData) SetValue(v *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData {
	s.Value = v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageData) Validate() error {
	if s.Value != nil {
		if err := s.Value.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValue struct {
	CodeProportionData []*DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData `json:"CodeProportionData,omitempty" xml:"CodeProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) GetCodeProportionData() []*DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	return s.CodeProportionData
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) SetCodeProportionData(v []*DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValue {
	s.CodeProportionData = v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) Validate() error {
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

type DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 2300
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The proportion of the HTTP status code.
	//
	// example:
	//
	// 67.1458998935037
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
}

func (s DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) GetCode() *string {
	return s.Code
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) GetCount() *string {
	return s.Count
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) GetProportion() *string {
	return s.Proportion
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) SetCode(v string) *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	s.Code = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) SetCount(v string) *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	s.Count = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) SetProportion(v string) *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) Validate() error {
	return dara.Validate(s)
}
