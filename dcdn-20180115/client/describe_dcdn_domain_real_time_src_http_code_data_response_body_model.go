// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody
	GetEndTime() *string
	SetRealTimeSrcHttpCodeData(v *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody
	GetRealTimeSrcHttpCodeData() *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData
	SetRequestId(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody
	GetStartTime() *string
}

type DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody struct {
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
	// example.com,example.org
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range during which data was queried.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-11-30T05:40:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The proportions of HTTP status codes at each time interval.
	RealTimeSrcHttpCodeData *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData `json:"RealTimeSrcHttpCodeData,omitempty" xml:"RealTimeSrcHttpCodeData,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// BC858082-736F-4A25-867B-E5B67C85ACF7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range during which data was queried.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-11-30T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) GetRealTimeSrcHttpCodeData() *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData {
	return s.RealTimeSrcHttpCodeData
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) SetRealTimeSrcHttpCodeData(v *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody {
	s.RealTimeSrcHttpCodeData = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData struct {
	UsageData []*DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) GetUsageData() []*DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData {
	return s.UsageData
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) SetUsageData(v []*DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData {
	s.UsageData = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeData) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData struct {
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2019-11-30T05:40:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The proportions of the HTTP status codes.
	Value *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) GetValue() *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue {
	return s.Value
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) SetTimeStamp(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) SetValue(v *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData {
	s.Value = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageData) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue struct {
	RealTimeSrcCodeProportionData []*DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData `json:"RealTimeSrcCodeProportionData,omitempty" xml:"RealTimeSrcCodeProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) GetRealTimeSrcCodeProportionData() []*DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData {
	return s.RealTimeSrcCodeProportionData
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) SetRealTimeSrcCodeProportionData(v []*DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue {
	s.RealTimeSrcCodeProportionData = v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValue) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of HTTP status codes.
	//
	// example:
	//
	// 2
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The proportion of the HTTP status code.
	//
	// example:
	//
	// 66.046511627907
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) GetCode() *string {
	return s.Code
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) GetCount() *string {
	return s.Count
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) GetProportion() *string {
	return s.Proportion
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) SetCode(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData {
	s.Code = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) SetCount(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData {
	s.Count = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) SetProportion(v string) *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponseBodyRealTimeSrcHttpCodeDataUsageDataValueRealTimeSrcCodeProportionData) Validate() error {
	return dara.Validate(s)
}
