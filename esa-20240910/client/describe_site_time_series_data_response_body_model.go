// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteTimeSeriesDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeSiteTimeSeriesDataResponseBodyData) *DescribeSiteTimeSeriesDataResponseBody
	GetData() []*DescribeSiteTimeSeriesDataResponseBodyData
	SetEndTime(v string) *DescribeSiteTimeSeriesDataResponseBody
	GetEndTime() *string
	SetInterval(v int64) *DescribeSiteTimeSeriesDataResponseBody
	GetInterval() *int64
	SetRequestId(v string) *DescribeSiteTimeSeriesDataResponseBody
	GetRequestId() *string
	SetSamplingRate(v float32) *DescribeSiteTimeSeriesDataResponseBody
	GetSamplingRate() *float32
	SetStartTime(v string) *DescribeSiteTimeSeriesDataResponseBody
	GetStartTime() *string
	SetSummarizedData(v []*DescribeSiteTimeSeriesDataResponseBodySummarizedData) *DescribeSiteTimeSeriesDataResponseBody
	GetSummarizedData() []*DescribeSiteTimeSeriesDataResponseBodySummarizedData
}

type DescribeSiteTimeSeriesDataResponseBody struct {
	// Returned data.
	Data []*DescribeSiteTimeSeriesDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The end time for fetching the data.
	//
	// The date format follows ISO8601 notation and uses UTC+0, formatted as yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2023-04-09T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The granularity of the data, in seconds.
	//
	// example:
	//
	// 300
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sampling rate, in %.
	//
	// example:
	//
	// 100
	SamplingRate *float32 `json:"SamplingRate,omitempty" xml:"SamplingRate,omitempty"`
	// The start time for fetching the data.
	//
	// The date format follows ISO8601 notation and uses UTC+0, formatted as yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2023-04-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Aggregated query data.
	SummarizedData []*DescribeSiteTimeSeriesDataResponseBodySummarizedData `json:"SummarizedData,omitempty" xml:"SummarizedData,omitempty" type:"Repeated"`
}

func (s DescribeSiteTimeSeriesDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteTimeSeriesDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSiteTimeSeriesDataResponseBody) GetData() []*DescribeSiteTimeSeriesDataResponseBodyData {
	return s.Data
}

func (s *DescribeSiteTimeSeriesDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSiteTimeSeriesDataResponseBody) GetInterval() *int64 {
	return s.Interval
}

func (s *DescribeSiteTimeSeriesDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSiteTimeSeriesDataResponseBody) GetSamplingRate() *float32 {
	return s.SamplingRate
}

func (s *DescribeSiteTimeSeriesDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSiteTimeSeriesDataResponseBody) GetSummarizedData() []*DescribeSiteTimeSeriesDataResponseBodySummarizedData {
	return s.SummarizedData
}

func (s *DescribeSiteTimeSeriesDataResponseBody) SetData(v []*DescribeSiteTimeSeriesDataResponseBodyData) *DescribeSiteTimeSeriesDataResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponseBody) SetEndTime(v string) *DescribeSiteTimeSeriesDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponseBody) SetInterval(v int64) *DescribeSiteTimeSeriesDataResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponseBody) SetRequestId(v string) *DescribeSiteTimeSeriesDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponseBody) SetSamplingRate(v float32) *DescribeSiteTimeSeriesDataResponseBody {
	s.SamplingRate = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponseBody) SetStartTime(v string) *DescribeSiteTimeSeriesDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponseBody) SetSummarizedData(v []*DescribeSiteTimeSeriesDataResponseBodySummarizedData) *DescribeSiteTimeSeriesDataResponseBody {
	s.SummarizedData = v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponseBody) Validate() error {
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

type DescribeSiteTimeSeriesDataResponseBodyData struct {
	// Returned data.
	DetailData []*DescribeSiteTimeSeriesDataResponseBodyDataDetailData `json:"DetailData,omitempty" xml:"DetailData,omitempty" type:"Repeated"`
	// Query dimension.
	//
	// example:
	//
	// ALL
	DimensionName *string `json:"DimensionName,omitempty" xml:"DimensionName,omitempty"`
	// Query dimension value.
	//
	// example:
	//
	// ALL
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
	// Query metric value.
	//
	// example:
	//
	// Traffic
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
}

func (s DescribeSiteTimeSeriesDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteTimeSeriesDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSiteTimeSeriesDataResponseBodyData) GetDetailData() []*DescribeSiteTimeSeriesDataResponseBodyDataDetailData {
	return s.DetailData
}

func (s *DescribeSiteTimeSeriesDataResponseBodyData) GetDimensionName() *string {
	return s.DimensionName
}

func (s *DescribeSiteTimeSeriesDataResponseBodyData) GetDimensionValue() *string {
	return s.DimensionValue
}

func (s *DescribeSiteTimeSeriesDataResponseBodyData) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeSiteTimeSeriesDataResponseBodyData) SetDetailData(v []*DescribeSiteTimeSeriesDataResponseBodyDataDetailData) *DescribeSiteTimeSeriesDataResponseBodyData {
	s.DetailData = v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponseBodyData) SetDimensionName(v string) *DescribeSiteTimeSeriesDataResponseBodyData {
	s.DimensionName = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponseBodyData) SetDimensionValue(v string) *DescribeSiteTimeSeriesDataResponseBodyData {
	s.DimensionValue = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponseBodyData) SetFieldName(v string) *DescribeSiteTimeSeriesDataResponseBodyData {
	s.FieldName = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponseBodyData) Validate() error {
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

type DescribeSiteTimeSeriesDataResponseBodyDataDetailData struct {
	// Start timestamp of the time slice.
	//
	// example:
	//
	// 2023-04-08T16:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// Value.
	//
	// example:
	//
	// 123
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSiteTimeSeriesDataResponseBodyDataDetailData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteTimeSeriesDataResponseBodyDataDetailData) GoString() string {
	return s.String()
}

func (s *DescribeSiteTimeSeriesDataResponseBodyDataDetailData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeSiteTimeSeriesDataResponseBodyDataDetailData) GetValue() interface{} {
	return s.Value
}

func (s *DescribeSiteTimeSeriesDataResponseBodyDataDetailData) SetTimeStamp(v string) *DescribeSiteTimeSeriesDataResponseBodyDataDetailData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponseBodyDataDetailData) SetValue(v interface{}) *DescribeSiteTimeSeriesDataResponseBodyDataDetailData {
	s.Value = v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponseBodyDataDetailData) Validate() error {
	return dara.Validate(s)
}

type DescribeSiteTimeSeriesDataResponseBodySummarizedData struct {
	// The aggregation method used.
	//
	// example:
	//
	// sum
	AggMethod *string `json:"AggMethod,omitempty" xml:"AggMethod,omitempty"`
	// The dimension of the aggregated data being queried.
	//
	// example:
	//
	// ALL
	DimensionName *string `json:"DimensionName,omitempty" xml:"DimensionName,omitempty"`
	// The value of the aggregated dimension being queried.
	//
	// example:
	//
	// ALL
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
	// The value of the aggregated metric being queried.
	//
	// example:
	//
	// Traffic
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The aggregated value.
	//
	// example:
	//
	// 12345
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSiteTimeSeriesDataResponseBodySummarizedData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteTimeSeriesDataResponseBodySummarizedData) GoString() string {
	return s.String()
}

func (s *DescribeSiteTimeSeriesDataResponseBodySummarizedData) GetAggMethod() *string {
	return s.AggMethod
}

func (s *DescribeSiteTimeSeriesDataResponseBodySummarizedData) GetDimensionName() *string {
	return s.DimensionName
}

func (s *DescribeSiteTimeSeriesDataResponseBodySummarizedData) GetDimensionValue() *string {
	return s.DimensionValue
}

func (s *DescribeSiteTimeSeriesDataResponseBodySummarizedData) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeSiteTimeSeriesDataResponseBodySummarizedData) GetValue() interface{} {
	return s.Value
}

func (s *DescribeSiteTimeSeriesDataResponseBodySummarizedData) SetAggMethod(v string) *DescribeSiteTimeSeriesDataResponseBodySummarizedData {
	s.AggMethod = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponseBodySummarizedData) SetDimensionName(v string) *DescribeSiteTimeSeriesDataResponseBodySummarizedData {
	s.DimensionName = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponseBodySummarizedData) SetDimensionValue(v string) *DescribeSiteTimeSeriesDataResponseBodySummarizedData {
	s.DimensionValue = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponseBodySummarizedData) SetFieldName(v string) *DescribeSiteTimeSeriesDataResponseBodySummarizedData {
	s.FieldName = &v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponseBodySummarizedData) SetValue(v interface{}) *DescribeSiteTimeSeriesDataResponseBodySummarizedData {
	s.Value = v
	return s
}

func (s *DescribeSiteTimeSeriesDataResponseBodySummarizedData) Validate() error {
	return dara.Validate(s)
}
