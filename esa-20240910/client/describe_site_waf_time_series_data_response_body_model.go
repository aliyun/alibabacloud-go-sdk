// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteWafTimeSeriesDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeSiteWafTimeSeriesDataResponseBodyData) *DescribeSiteWafTimeSeriesDataResponseBody
	GetData() []*DescribeSiteWafTimeSeriesDataResponseBodyData
	SetEndTime(v string) *DescribeSiteWafTimeSeriesDataResponseBody
	GetEndTime() *string
	SetInterval(v int64) *DescribeSiteWafTimeSeriesDataResponseBody
	GetInterval() *int64
	SetRequestId(v string) *DescribeSiteWafTimeSeriesDataResponseBody
	GetRequestId() *string
	SetSamplingRate(v float32) *DescribeSiteWafTimeSeriesDataResponseBody
	GetSamplingRate() *float32
	SetStartTime(v string) *DescribeSiteWafTimeSeriesDataResponseBody
	GetStartTime() *string
	SetSummarizedData(v []*DescribeSiteWafTimeSeriesDataResponseBodySummarizedData) *DescribeSiteWafTimeSeriesDataResponseBody
	GetSummarizedData() []*DescribeSiteWafTimeSeriesDataResponseBodySummarizedData
}

type DescribeSiteWafTimeSeriesDataResponseBody struct {
	// The returned data.
	Data []*DescribeSiteWafTimeSeriesDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The end of the time range for the returned data.
	//
	// The time is in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is in UTC+0.
	//
	// example:
	//
	// 2023-04-09T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The data granularity. Unit: seconds.
	//
	// example:
	//
	// 300
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 93652946-2687-5428-8254-533B1E6A***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sampling rate. Unit: %.
	//
	// example:
	//
	// 100
	SamplingRate *float32 `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
	//
	// example:
	//
	// 2023-04-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The summarized data.
	SummarizedData []*DescribeSiteWafTimeSeriesDataResponseBodySummarizedData `json:"SummarizedData,omitempty" xml:"SummarizedData,omitempty" type:"Repeated"`
}

func (s DescribeSiteWafTimeSeriesDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteWafTimeSeriesDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSiteWafTimeSeriesDataResponseBody) GetData() []*DescribeSiteWafTimeSeriesDataResponseBodyData {
	return s.Data
}

func (s *DescribeSiteWafTimeSeriesDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSiteWafTimeSeriesDataResponseBody) GetInterval() *int64 {
	return s.Interval
}

func (s *DescribeSiteWafTimeSeriesDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSiteWafTimeSeriesDataResponseBody) GetSamplingRate() *float32 {
	return s.SamplingRate
}

func (s *DescribeSiteWafTimeSeriesDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSiteWafTimeSeriesDataResponseBody) GetSummarizedData() []*DescribeSiteWafTimeSeriesDataResponseBodySummarizedData {
	return s.SummarizedData
}

func (s *DescribeSiteWafTimeSeriesDataResponseBody) SetData(v []*DescribeSiteWafTimeSeriesDataResponseBodyData) *DescribeSiteWafTimeSeriesDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponseBody) SetEndTime(v string) *DescribeSiteWafTimeSeriesDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponseBody) SetInterval(v int64) *DescribeSiteWafTimeSeriesDataResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponseBody) SetRequestId(v string) *DescribeSiteWafTimeSeriesDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponseBody) SetSamplingRate(v float32) *DescribeSiteWafTimeSeriesDataResponseBody {
	s.SamplingRate = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponseBody) SetStartTime(v string) *DescribeSiteWafTimeSeriesDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponseBody) SetSummarizedData(v []*DescribeSiteWafTimeSeriesDataResponseBodySummarizedData) *DescribeSiteWafTimeSeriesDataResponseBody {
	s.SummarizedData = v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SummarizedData != nil {
		for _, item := range s.SummarizedData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSiteWafTimeSeriesDataResponseBodyData struct {
	// The returned data.
	DetailData []*DescribeSiteWafTimeSeriesDataResponseBodyDataDetailData `json:"DetailData,omitempty" xml:"DetailData,omitempty" type:"Repeated"`
	// The query dimensions.
	//
	// example:
	//
	// ALL
	DimensionName *string `json:"DimensionName,omitempty" xml:"DimensionName,omitempty"`
	// The dimension value.
	//
	// example:
	//
	// ALL
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
	// The metric name.
	//
	// example:
	//
	// Requests
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
}

func (s DescribeSiteWafTimeSeriesDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteWafTimeSeriesDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodyData) GetDetailData() []*DescribeSiteWafTimeSeriesDataResponseBodyDataDetailData {
	return s.DetailData
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodyData) GetDimensionName() *string {
	return s.DimensionName
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodyData) GetDimensionValue() *string {
	return s.DimensionValue
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodyData) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodyData) SetDetailData(v []*DescribeSiteWafTimeSeriesDataResponseBodyDataDetailData) *DescribeSiteWafTimeSeriesDataResponseBodyData {
	s.DetailData = v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodyData) SetDimensionName(v string) *DescribeSiteWafTimeSeriesDataResponseBodyData {
	s.DimensionName = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodyData) SetDimensionValue(v string) *DescribeSiteWafTimeSeriesDataResponseBodyData {
	s.DimensionValue = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodyData) SetFieldName(v string) *DescribeSiteWafTimeSeriesDataResponseBodyData {
	s.FieldName = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodyData) Validate() error {
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

type DescribeSiteWafTimeSeriesDataResponseBodyDataDetailData struct {
	// The start time of the time slice.
	//
	// The time is in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is in UTC+0.
	//
	// example:
	//
	// 2023-04-08T16:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The value.
	//
	// example:
	//
	// 123
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSiteWafTimeSeriesDataResponseBodyDataDetailData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteWafTimeSeriesDataResponseBodyDataDetailData) GoString() string {
	return s.String()
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodyDataDetailData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodyDataDetailData) GetValue() interface{} {
	return s.Value
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodyDataDetailData) SetTimeStamp(v string) *DescribeSiteWafTimeSeriesDataResponseBodyDataDetailData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodyDataDetailData) SetValue(v interface{}) *DescribeSiteWafTimeSeriesDataResponseBodyDataDetailData {
	s.Value = v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodyDataDetailData) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteWafTimeSeriesDataResponseBodySummarizedData struct {
	// The aggregation method.
	//
	// example:
	//
	// sum
	AggMethod *string `json:"AggMethod,omitempty" xml:"AggMethod,omitempty"`
	// The summarized dimension name.
	//
	// example:
	//
	// ALL
	DimensionName *string `json:"DimensionName,omitempty" xml:"DimensionName,omitempty"`
	// The summarized dimension value.
	//
	// example:
	//
	// ALL
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
	// The summarized metric name.
	//
	// example:
	//
	// Requests
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The aggregated value.
	//
	// example:
	//
	// 123456
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSiteWafTimeSeriesDataResponseBodySummarizedData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteWafTimeSeriesDataResponseBodySummarizedData) GoString() string {
	return s.String()
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodySummarizedData) GetAggMethod() *string {
	return s.AggMethod
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodySummarizedData) GetDimensionName() *string {
	return s.DimensionName
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodySummarizedData) GetDimensionValue() *string {
	return s.DimensionValue
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodySummarizedData) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodySummarizedData) GetValue() interface{} {
	return s.Value
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodySummarizedData) SetAggMethod(v string) *DescribeSiteWafTimeSeriesDataResponseBodySummarizedData {
	s.AggMethod = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodySummarizedData) SetDimensionName(v string) *DescribeSiteWafTimeSeriesDataResponseBodySummarizedData {
	s.DimensionName = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodySummarizedData) SetDimensionValue(v string) *DescribeSiteWafTimeSeriesDataResponseBodySummarizedData {
	s.DimensionValue = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodySummarizedData) SetFieldName(v string) *DescribeSiteWafTimeSeriesDataResponseBodySummarizedData {
	s.FieldName = &v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodySummarizedData) SetValue(v interface{}) *DescribeSiteWafTimeSeriesDataResponseBodySummarizedData {
	s.Value = v
	return s
}

func (s *DescribeSiteWafTimeSeriesDataResponseBodySummarizedData) Validate() error {
	return dara.Validate(s)
}
