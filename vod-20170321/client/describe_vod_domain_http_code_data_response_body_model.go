// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainHttpCodeDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeVodDomainHttpCodeDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeVodDomainHttpCodeDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainHttpCodeDataResponseBody
	GetEndTime() *string
	SetHttpCodeData(v *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeData) *DescribeVodDomainHttpCodeDataResponseBody
	GetHttpCodeData() *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeData
	SetRequestId(v string) *DescribeVodDomainHttpCodeDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodDomainHttpCodeDataResponseBody
	GetStartTime() *string
}

type DescribeVodDomainHttpCodeDataResponseBody struct {
	DataInterval *string                                                `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName   *string                                                `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime      *string                                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	HttpCodeData *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeData `json:"HttpCodeData,omitempty" xml:"HttpCodeData,omitempty" type:"Struct"`
	RequestId    *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime    *string                                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainHttpCodeDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainHttpCodeDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeVodDomainHttpCodeDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainHttpCodeDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainHttpCodeDataResponseBody) GetHttpCodeData() *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeData {
	return s.HttpCodeData
}

func (s *DescribeVodDomainHttpCodeDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainHttpCodeDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeVodDomainHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainHttpCodeDataResponseBody) SetDomainName(v string) *DescribeVodDomainHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainHttpCodeDataResponseBody) SetEndTime(v string) *DescribeVodDomainHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainHttpCodeDataResponseBody) SetHttpCodeData(v *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeData) *DescribeVodDomainHttpCodeDataResponseBody {
	s.HttpCodeData = v
	return s
}

func (s *DescribeVodDomainHttpCodeDataResponseBody) SetRequestId(v string) *DescribeVodDomainHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainHttpCodeDataResponseBody) SetStartTime(v string) *DescribeVodDomainHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainHttpCodeDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainHttpCodeDataResponseBodyHttpCodeData struct {
	UsageData []*DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainHttpCodeDataResponseBodyHttpCodeData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainHttpCodeDataResponseBodyHttpCodeData) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeData) GetUsageData() []*DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageData {
	return s.UsageData
}

func (s *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeData) SetUsageData(v []*DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeData {
	s.UsageData = v
	return s
}

func (s *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeData) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageData struct {
	TimeStamp *string                                                              `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) GetValue() *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue {
	return s.Value
}

func (s *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) SetTimeStamp(v string) *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) SetValue(v *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageData {
	s.Value = v
	return s
}

func (s *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageData) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue struct {
	CodeProportionData []*DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData `json:"CodeProportionData,omitempty" xml:"CodeProportionData,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) GetCodeProportionData() []*DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	return s.CodeProportionData
}

func (s *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) SetCodeProportionData(v []*DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue {
	s.CodeProportionData = v
	return s
}

func (s *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValue) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Count      *string `json:"Count,omitempty" xml:"Count,omitempty"`
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
}

func (s DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) GetCode() *string {
	return s.Code
}

func (s *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) GetCount() *string {
	return s.Count
}

func (s *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) GetProportion() *string {
	return s.Proportion
}

func (s *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) SetCode(v string) *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	s.Code = &v
	return s
}

func (s *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) SetCount(v string) *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	s.Count = &v
	return s
}

func (s *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) SetProportion(v string) *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeVodDomainHttpCodeDataResponseBodyHttpCodeDataUsageDataValueCodeProportionData) Validate() error {
	return dara.Validate(s)
}
