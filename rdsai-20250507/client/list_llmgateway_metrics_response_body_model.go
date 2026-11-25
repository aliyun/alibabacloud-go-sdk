// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLLMGatewayMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGranularity(v string) *ListLLMGatewayMetricsResponseBody
	GetGranularity() *string
	SetMetricItems(v []*ListLLMGatewayMetricsResponseBodyMetricItems) *ListLLMGatewayMetricsResponseBody
	GetMetricItems() []*ListLLMGatewayMetricsResponseBodyMetricItems
	SetMetricType(v string) *ListLLMGatewayMetricsResponseBody
	GetMetricType() *string
	SetRequestId(v string) *ListLLMGatewayMetricsResponseBody
	GetRequestId() *string
}

type ListLLMGatewayMetricsResponseBody struct {
	// example:
	//
	// minute
	Granularity *string                                         `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	MetricItems []*ListLLMGatewayMetricsResponseBodyMetricItems `json:"MetricItems,omitempty" xml:"MetricItems,omitempty" type:"Repeated"`
	// example:
	//
	// tokenPerMinute
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// example:
	//
	// ABCDEF77-41B5-5F9E-81E8-93DFDD379024
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLLMGatewayMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLLMGatewayMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLLMGatewayMetricsResponseBody) GetGranularity() *string {
	return s.Granularity
}

func (s *ListLLMGatewayMetricsResponseBody) GetMetricItems() []*ListLLMGatewayMetricsResponseBodyMetricItems {
	return s.MetricItems
}

func (s *ListLLMGatewayMetricsResponseBody) GetMetricType() *string {
	return s.MetricType
}

func (s *ListLLMGatewayMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLLMGatewayMetricsResponseBody) SetGranularity(v string) *ListLLMGatewayMetricsResponseBody {
	s.Granularity = &v
	return s
}

func (s *ListLLMGatewayMetricsResponseBody) SetMetricItems(v []*ListLLMGatewayMetricsResponseBodyMetricItems) *ListLLMGatewayMetricsResponseBody {
	s.MetricItems = v
	return s
}

func (s *ListLLMGatewayMetricsResponseBody) SetMetricType(v string) *ListLLMGatewayMetricsResponseBody {
	s.MetricType = &v
	return s
}

func (s *ListLLMGatewayMetricsResponseBody) SetRequestId(v string) *ListLLMGatewayMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLLMGatewayMetricsResponseBody) Validate() error {
	if s.MetricItems != nil {
		for _, item := range s.MetricItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLLMGatewayMetricsResponseBodyMetricItems struct {
	DataPoints []*ListLLMGatewayMetricsResponseBodyMetricItemsDataPoints `json:"DataPoints,omitempty" xml:"DataPoints,omitempty" type:"Repeated"`
	// example:
	//
	// inputTokens
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// tokens
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s ListLLMGatewayMetricsResponseBodyMetricItems) String() string {
	return dara.Prettify(s)
}

func (s ListLLMGatewayMetricsResponseBodyMetricItems) GoString() string {
	return s.String()
}

func (s *ListLLMGatewayMetricsResponseBodyMetricItems) GetDataPoints() []*ListLLMGatewayMetricsResponseBodyMetricItemsDataPoints {
	return s.DataPoints
}

func (s *ListLLMGatewayMetricsResponseBodyMetricItems) GetKey() *string {
	return s.Key
}

func (s *ListLLMGatewayMetricsResponseBodyMetricItems) GetModel() *string {
	return s.Model
}

func (s *ListLLMGatewayMetricsResponseBodyMetricItems) GetUnit() *string {
	return s.Unit
}

func (s *ListLLMGatewayMetricsResponseBodyMetricItems) SetDataPoints(v []*ListLLMGatewayMetricsResponseBodyMetricItemsDataPoints) *ListLLMGatewayMetricsResponseBodyMetricItems {
	s.DataPoints = v
	return s
}

func (s *ListLLMGatewayMetricsResponseBodyMetricItems) SetKey(v string) *ListLLMGatewayMetricsResponseBodyMetricItems {
	s.Key = &v
	return s
}

func (s *ListLLMGatewayMetricsResponseBodyMetricItems) SetModel(v string) *ListLLMGatewayMetricsResponseBodyMetricItems {
	s.Model = &v
	return s
}

func (s *ListLLMGatewayMetricsResponseBodyMetricItems) SetUnit(v string) *ListLLMGatewayMetricsResponseBodyMetricItems {
	s.Unit = &v
	return s
}

func (s *ListLLMGatewayMetricsResponseBodyMetricItems) Validate() error {
	if s.DataPoints != nil {
		for _, item := range s.DataPoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLLMGatewayMetricsResponseBodyMetricItemsDataPoints struct {
	// example:
	//
	// 1779799200000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 16022
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListLLMGatewayMetricsResponseBodyMetricItemsDataPoints) String() string {
	return dara.Prettify(s)
}

func (s ListLLMGatewayMetricsResponseBodyMetricItemsDataPoints) GoString() string {
	return s.String()
}

func (s *ListLLMGatewayMetricsResponseBodyMetricItemsDataPoints) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ListLLMGatewayMetricsResponseBodyMetricItemsDataPoints) GetValue() *string {
	return s.Value
}

func (s *ListLLMGatewayMetricsResponseBodyMetricItemsDataPoints) SetTimestamp(v int64) *ListLLMGatewayMetricsResponseBodyMetricItemsDataPoints {
	s.Timestamp = &v
	return s
}

func (s *ListLLMGatewayMetricsResponseBodyMetricItemsDataPoints) SetValue(v string) *ListLLMGatewayMetricsResponseBodyMetricItemsDataPoints {
	s.Value = &v
	return s
}

func (s *ListLLMGatewayMetricsResponseBodyMetricItemsDataPoints) Validate() error {
	return dara.Validate(s)
}
