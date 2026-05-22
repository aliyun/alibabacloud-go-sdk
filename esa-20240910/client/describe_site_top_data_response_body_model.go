// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteTopDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeSiteTopDataResponseBodyData) *DescribeSiteTopDataResponseBody
	GetData() []*DescribeSiteTopDataResponseBodyData
	SetEndTime(v string) *DescribeSiteTopDataResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeSiteTopDataResponseBody
	GetRequestId() *string
	SetSamplingRate(v float32) *DescribeSiteTopDataResponseBody
	GetSamplingRate() *float32
	SetStartTime(v string) *DescribeSiteTopDataResponseBody
	GetStartTime() *string
}

type DescribeSiteTopDataResponseBody struct {
	Data         []*DescribeSiteTopDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	EndTime      *string                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SamplingRate *float32                               `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
	StartTime    *string                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSiteTopDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteTopDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSiteTopDataResponseBody) GetData() []*DescribeSiteTopDataResponseBodyData {
	return s.Data
}

func (s *DescribeSiteTopDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSiteTopDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSiteTopDataResponseBody) GetSamplingRate() *float32 {
	return s.SamplingRate
}

func (s *DescribeSiteTopDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSiteTopDataResponseBody) SetData(v []*DescribeSiteTopDataResponseBodyData) *DescribeSiteTopDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSiteTopDataResponseBody) SetEndTime(v string) *DescribeSiteTopDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeSiteTopDataResponseBody) SetRequestId(v string) *DescribeSiteTopDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSiteTopDataResponseBody) SetSamplingRate(v float32) *DescribeSiteTopDataResponseBody {
	s.SamplingRate = &v
	return s
}

func (s *DescribeSiteTopDataResponseBody) SetStartTime(v string) *DescribeSiteTopDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeSiteTopDataResponseBody) Validate() error {
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

type DescribeSiteTopDataResponseBodyData struct {
	DetailData    []*DescribeSiteTopDataResponseBodyDataDetailData `json:"DetailData,omitempty" xml:"DetailData,omitempty" type:"Repeated"`
	DimensionName *string                                          `json:"DimensionName,omitempty" xml:"DimensionName,omitempty"`
	FieldName     *string                                          `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
}

func (s DescribeSiteTopDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteTopDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSiteTopDataResponseBodyData) GetDetailData() []*DescribeSiteTopDataResponseBodyDataDetailData {
	return s.DetailData
}

func (s *DescribeSiteTopDataResponseBodyData) GetDimensionName() *string {
	return s.DimensionName
}

func (s *DescribeSiteTopDataResponseBodyData) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeSiteTopDataResponseBodyData) SetDetailData(v []*DescribeSiteTopDataResponseBodyDataDetailData) *DescribeSiteTopDataResponseBodyData {
	s.DetailData = v
	return s
}

func (s *DescribeSiteTopDataResponseBodyData) SetDimensionName(v string) *DescribeSiteTopDataResponseBodyData {
	s.DimensionName = &v
	return s
}

func (s *DescribeSiteTopDataResponseBodyData) SetFieldName(v string) *DescribeSiteTopDataResponseBodyData {
	s.FieldName = &v
	return s
}

func (s *DescribeSiteTopDataResponseBodyData) Validate() error {
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

type DescribeSiteTopDataResponseBodyDataDetailData struct {
	DimensionValue *string     `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
	Value          interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSiteTopDataResponseBodyDataDetailData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteTopDataResponseBodyDataDetailData) GoString() string {
	return s.String()
}

func (s *DescribeSiteTopDataResponseBodyDataDetailData) GetDimensionValue() *string {
	return s.DimensionValue
}

func (s *DescribeSiteTopDataResponseBodyDataDetailData) GetValue() interface{} {
	return s.Value
}

func (s *DescribeSiteTopDataResponseBodyDataDetailData) SetDimensionValue(v string) *DescribeSiteTopDataResponseBodyDataDetailData {
	s.DimensionValue = &v
	return s
}

func (s *DescribeSiteTopDataResponseBodyDataDetailData) SetValue(v interface{}) *DescribeSiteTopDataResponseBodyDataDetailData {
	s.Value = v
	return s
}

func (s *DescribeSiteTopDataResponseBodyDataDetailData) Validate() error {
	return dara.Validate(s)
}
