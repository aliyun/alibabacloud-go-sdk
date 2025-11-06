// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupMetricDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetricData(v *ListResourceGroupMetricDataResponseBodyMetricData) *ListResourceGroupMetricDataResponseBody
	GetMetricData() *ListResourceGroupMetricDataResponseBodyMetricData
	SetRequestId(v string) *ListResourceGroupMetricDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListResourceGroupMetricDataResponseBody
	GetSuccess() *bool
}

type ListResourceGroupMetricDataResponseBody struct {
	MetricData *ListResourceGroupMetricDataResponseBodyMetricData `json:"MetricData,omitempty" xml:"MetricData,omitempty" type:"Struct"`
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListResourceGroupMetricDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceGroupMetricDataResponseBody) GetMetricData() *ListResourceGroupMetricDataResponseBodyMetricData {
	return s.MetricData
}

func (s *ListResourceGroupMetricDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceGroupMetricDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListResourceGroupMetricDataResponseBody) SetMetricData(v *ListResourceGroupMetricDataResponseBodyMetricData) *ListResourceGroupMetricDataResponseBody {
	s.MetricData = v
	return s
}

func (s *ListResourceGroupMetricDataResponseBody) SetRequestId(v string) *ListResourceGroupMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceGroupMetricDataResponseBody) SetSuccess(v bool) *ListResourceGroupMetricDataResponseBody {
	s.Success = &v
	return s
}

func (s *ListResourceGroupMetricDataResponseBody) Validate() error {
	if s.MetricData != nil {
		if err := s.MetricData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourceGroupMetricDataResponseBodyMetricData struct {
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// CUSpec
	MetricName *string                                                     `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Metrics    []*ListResourceGroupMetricDataResponseBodyMetricDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// example:
	//
	// tSBOXZcAmk+akxRkwRuXnGQEsIDODyd5ulPqgytNTbLp4bhb7fuvz13FXtm87Kfl
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListResourceGroupMetricDataResponseBodyMetricData) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupMetricDataResponseBodyMetricData) GoString() string {
	return s.String()
}

func (s *ListResourceGroupMetricDataResponseBodyMetricData) GetId() *string {
	return s.Id
}

func (s *ListResourceGroupMetricDataResponseBodyMetricData) GetMetricName() *string {
	return s.MetricName
}

func (s *ListResourceGroupMetricDataResponseBodyMetricData) GetMetrics() []*ListResourceGroupMetricDataResponseBodyMetricDataMetrics {
	return s.Metrics
}

func (s *ListResourceGroupMetricDataResponseBodyMetricData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceGroupMetricDataResponseBodyMetricData) SetId(v string) *ListResourceGroupMetricDataResponseBodyMetricData {
	s.Id = &v
	return s
}

func (s *ListResourceGroupMetricDataResponseBodyMetricData) SetMetricName(v string) *ListResourceGroupMetricDataResponseBodyMetricData {
	s.MetricName = &v
	return s
}

func (s *ListResourceGroupMetricDataResponseBodyMetricData) SetMetrics(v []*ListResourceGroupMetricDataResponseBodyMetricDataMetrics) *ListResourceGroupMetricDataResponseBodyMetricData {
	s.Metrics = v
	return s
}

func (s *ListResourceGroupMetricDataResponseBodyMetricData) SetNextToken(v string) *ListResourceGroupMetricDataResponseBodyMetricData {
	s.NextToken = &v
	return s
}

func (s *ListResourceGroupMetricDataResponseBodyMetricData) Validate() error {
	if s.Metrics != nil {
		for _, item := range s.Metrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceGroupMetricDataResponseBodyMetricDataMetrics struct {
	// example:
	//
	// 1761184929633
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 1.0
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourceGroupMetricDataResponseBodyMetricDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupMetricDataResponseBodyMetricDataMetrics) GoString() string {
	return s.String()
}

func (s *ListResourceGroupMetricDataResponseBodyMetricDataMetrics) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ListResourceGroupMetricDataResponseBodyMetricDataMetrics) GetValue() *float64 {
	return s.Value
}

func (s *ListResourceGroupMetricDataResponseBodyMetricDataMetrics) SetTimestamp(v int64) *ListResourceGroupMetricDataResponseBodyMetricDataMetrics {
	s.Timestamp = &v
	return s
}

func (s *ListResourceGroupMetricDataResponseBodyMetricDataMetrics) SetValue(v float64) *ListResourceGroupMetricDataResponseBodyMetricDataMetrics {
	s.Value = &v
	return s
}

func (s *ListResourceGroupMetricDataResponseBodyMetricDataMetrics) Validate() error {
	return dara.Validate(s)
}
