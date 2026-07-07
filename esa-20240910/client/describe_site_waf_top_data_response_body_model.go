// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteWafTopDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeSiteWafTopDataResponseBodyData) *DescribeSiteWafTopDataResponseBody
	GetData() []*DescribeSiteWafTopDataResponseBodyData
	SetEndTime(v string) *DescribeSiteWafTopDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeSiteWafTopDataResponseBody
	GetRequestId() *string
	SetSamplingRate(v float32) *DescribeSiteWafTopDataResponseBody
	GetSamplingRate() *float32
	SetStartTime(v string) *DescribeSiteWafTopDataResponseBody
	GetStartTime() *string
}

type DescribeSiteWafTopDataResponseBody struct {
	// The returned data.
	Data []*DescribeSiteWafTopDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The end of the time range for the returned data.
	//
	// The time is in ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is in UTC+0.
	//
	// example:
	//
	// 2023-04-09T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 64041D4F-B615-5DEB-AC94-F01EE433****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sampling rate. Unit: %.
	//
	// example:
	//
	// 100
	SamplingRate *float32 `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
	//
	// example:
	//
	// 2023-04-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSiteWafTopDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteWafTopDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSiteWafTopDataResponseBody) GetData() []*DescribeSiteWafTopDataResponseBodyData {
	return s.Data
}

func (s *DescribeSiteWafTopDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSiteWafTopDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSiteWafTopDataResponseBody) GetSamplingRate() *float32 {
	return s.SamplingRate
}

func (s *DescribeSiteWafTopDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSiteWafTopDataResponseBody) SetData(v []*DescribeSiteWafTopDataResponseBodyData) *DescribeSiteWafTopDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSiteWafTopDataResponseBody) SetEndTime(v string) *DescribeSiteWafTopDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeSiteWafTopDataResponseBody) SetRequestId(v string) *DescribeSiteWafTopDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSiteWafTopDataResponseBody) SetSamplingRate(v float32) *DescribeSiteWafTopDataResponseBody {
	s.SamplingRate = &v
	return s
}

func (s *DescribeSiteWafTopDataResponseBody) SetStartTime(v string) *DescribeSiteWafTopDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeSiteWafTopDataResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSiteWafTopDataResponseBodyData struct {
	// The returned data.
	DetailData []*DescribeSiteWafTopDataResponseBodyDataDetailData `json:"DetailData,omitempty" xml:"DetailData,omitempty" type:"Repeated"`
	// The query dimension.
	//
	// example:
	//
	// ALL
	DimensionName *string `json:"DimensionName,omitempty" xml:"DimensionName,omitempty"`
	// The query metric value.
	//
	// example:
	//
	// Requests
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
}

func (s DescribeSiteWafTopDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteWafTopDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSiteWafTopDataResponseBodyData) GetDetailData() []*DescribeSiteWafTopDataResponseBodyDataDetailData {
	return s.DetailData
}

func (s *DescribeSiteWafTopDataResponseBodyData) GetDimensionName() *string {
	return s.DimensionName
}

func (s *DescribeSiteWafTopDataResponseBodyData) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeSiteWafTopDataResponseBodyData) SetDetailData(v []*DescribeSiteWafTopDataResponseBodyDataDetailData) *DescribeSiteWafTopDataResponseBodyData {
	s.DetailData = v
	return s
}

func (s *DescribeSiteWafTopDataResponseBodyData) SetDimensionName(v string) *DescribeSiteWafTopDataResponseBodyData {
	s.DimensionName = &v
	return s
}

func (s *DescribeSiteWafTopDataResponseBodyData) SetFieldName(v string) *DescribeSiteWafTopDataResponseBodyData {
	s.FieldName = &v
	return s
}

func (s *DescribeSiteWafTopDataResponseBodyData) Validate() error {
	if s.DetailData != nil {
		for _, item := range s.DetailData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSiteWafTopDataResponseBodyDataDetailData struct {
	// The query dimension value.
	//
	// example:
	//
	// ALL
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
	// The value.
	//
	// example:
	//
	// 123
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSiteWafTopDataResponseBodyDataDetailData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteWafTopDataResponseBodyDataDetailData) GoString() string {
	return s.String()
}

func (s *DescribeSiteWafTopDataResponseBodyDataDetailData) GetDimensionValue() *string {
	return s.DimensionValue
}

func (s *DescribeSiteWafTopDataResponseBodyDataDetailData) GetValue() interface{} {
	return s.Value
}

func (s *DescribeSiteWafTopDataResponseBodyDataDetailData) SetDimensionValue(v string) *DescribeSiteWafTopDataResponseBodyDataDetailData {
	s.DimensionValue = &v
	return s
}

func (s *DescribeSiteWafTopDataResponseBodyDataDetailData) SetValue(v interface{}) *DescribeSiteWafTopDataResponseBodyDataDetailData {
	s.Value = v
	return s
}

func (s *DescribeSiteWafTopDataResponseBodyDataDetailData) Validate() error {
	return dara.Validate(s)
}
