// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainRealTimeHttpCodeDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponseBody
	GetDataInterval() *string
	SetDomainName(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponseBody
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponseBody
	GetEndTime() *string
	SetRealTimeHttpCodeData(v *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) *DescribeLiveDomainRealTimeHttpCodeDataResponseBody
	GetRealTimeHttpCodeData() *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData
	SetRequestId(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponseBody
	GetStartTime() *string
}

type DescribeLiveDomainRealTimeHttpCodeDataResponseBody struct {
	// The time interval between the entries returned. Unit: seconds Default value: 60.
	//
	// example:
	//
	// 60
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The streaming domain.
	//
	// example:
	//
	// example.com,example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range during which the data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-11-30T05:40:00Z
	EndTime              *string                                                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RealTimeHttpCodeData *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData `json:"RealTimeHttpCodeData,omitempty" xml:"RealTimeHttpCodeData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BC858082-736F-4A25-867B-E5B67C85ACF7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range during which the data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-11-30T05:33:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBody) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBody) GetRealTimeHttpCodeData() *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData {
	return s.RealTimeHttpCodeData
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBody) SetDataInterval(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBody) SetDomainName(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBody) SetEndTime(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBody) SetRealTimeHttpCodeData(v *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) *DescribeLiveDomainRealTimeHttpCodeDataResponseBody {
	s.RealTimeHttpCodeData = v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBody) SetRequestId(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBody) SetStartTime(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBody) Validate() error {
	if s.RealTimeHttpCodeData != nil {
		if err := s.RealTimeHttpCodeData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData struct {
	UsageData []*DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData `json:"UsageData,omitempty" xml:"UsageData,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) GetUsageData() []*DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData {
	return s.UsageData
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) SetUsageData(v []*DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData {
	s.UsageData = v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeData) Validate() error {
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

type DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData struct {
	TimeStamp *string                                                                               `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Struct"`
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) GetValue() *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue {
	return s.Value
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) SetTimeStamp(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) SetValue(v *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData {
	s.Value = v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageData) Validate() error {
	if s.Value != nil {
		if err := s.Value.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue struct {
	RealTimeCodeProportionData []*DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData `json:"RealTimeCodeProportionData,omitempty" xml:"RealTimeCodeProportionData,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) GetRealTimeCodeProportionData() []*DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	return s.RealTimeCodeProportionData
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) SetRealTimeCodeProportionData(v []*DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue {
	s.RealTimeCodeProportionData = v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValue) Validate() error {
	if s.RealTimeCodeProportionData != nil {
		for _, item := range s.RealTimeCodeProportionData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Count      *string `json:"Count,omitempty" xml:"Count,omitempty"`
	Proportion *string `json:"Proportion,omitempty" xml:"Proportion,omitempty"`
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) GetCode() *string {
	return s.Code
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) GetCount() *string {
	return s.Count
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) GetProportion() *string {
	return s.Proportion
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) SetCode(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	s.Code = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) SetCount(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	s.Count = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) SetProportion(v string) *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData {
	s.Proportion = &v
	return s
}

func (s *DescribeLiveDomainRealTimeHttpCodeDataResponseBodyRealTimeHttpCodeDataUsageDataValueRealTimeCodeProportionData) Validate() error {
	return dara.Validate(s)
}
