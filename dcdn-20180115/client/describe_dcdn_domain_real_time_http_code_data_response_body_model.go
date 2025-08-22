// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeHttpCodeDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody
	GetEndTime() *string
	SetRealTimeHttpCodeData(v *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody
	GetRealTimeHttpCodeData() *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData
	SetRequestId(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody
	GetStartTime() *string
}

type DescribeDcdnDomainRealTimeHttpCodeDataResponseBody struct {
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
	// 2017-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The proportions of the HTTP status codes.
	RealTimeHttpCodeData *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData `json:"RealTimeHttpCodeData,omitempty" xml:"RealTimeHttpCodeData,omitempty" type:"Struct"`
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
	// 2017-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) GetRealTimeHttpCodeData() *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData {
	return s.RealTimeHttpCodeData
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) SetDomainName(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) SetEndTime(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) SetRealTimeHttpCodeData(v *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody {
	s.RealTimeHttpCodeData = v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) SetStartTime(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData struct {
	UsageData []*DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) GetUsageData() []*DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData {
	return s.UsageData
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) SetUsageData(v []*DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData {
	s.UsageData = v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData struct {
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2017-12-10T20:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The information about the HTTP status code returned.
	Value *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) GetValue() *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue {
	return s.Value
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) SetTimeStamp(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) SetValue(v *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData {
	s.Value = v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue struct {
	RealTimeCodeProportionData []*DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData `json:"RealTimeCodeProportionData,omitempty" xml:"RealTimeCodeProportionData,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) GetRealTimeCodeProportionData() []*DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	return s.RealTimeCodeProportionData
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) SetRealTimeCodeProportionData(v []*DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue {
	s.RealTimeCodeProportionData = v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData struct {
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
	// 1
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The proportion of the HTTP status code.
	//
	// example:
	//
	// 0.01155980271270037
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) GetCode() *string {
	return s.Code
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) GetCount() *string {
	return s.Count
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) GetProportion() *string {
	return s.Proportion
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) SetCode(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	s.Code = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) SetCount(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	s.Count = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) SetProportion(v string) *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) Validate() error {
	return dara.Validate(s)
}
