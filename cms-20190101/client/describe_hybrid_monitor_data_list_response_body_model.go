// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridMonitorDataListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeHybridMonitorDataListResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeHybridMonitorDataListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeHybridMonitorDataListResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeHybridMonitorDataListResponseBody
	GetSuccess() *string
	SetTimeSeries(v []*DescribeHybridMonitorDataListResponseBodyTimeSeries) *DescribeHybridMonitorDataListResponseBody
	GetTimeSeries() []*DescribeHybridMonitorDataListResponseBodyTimeSeries
}

type DescribeHybridMonitorDataListResponseBody struct {
	// The response code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// end timestamp must not be before start time.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C240412F-3F5F-50E2-ACEC-DE808EF9C4BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The returned monitoring data.
	TimeSeries []*DescribeHybridMonitorDataListResponseBodyTimeSeries `json:"TimeSeries,omitempty" xml:"TimeSeries,omitempty" type:"Repeated"`
}

func (s DescribeHybridMonitorDataListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorDataListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorDataListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHybridMonitorDataListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHybridMonitorDataListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridMonitorDataListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeHybridMonitorDataListResponseBody) GetTimeSeries() []*DescribeHybridMonitorDataListResponseBodyTimeSeries {
	return s.TimeSeries
}

func (s *DescribeHybridMonitorDataListResponseBody) SetCode(v string) *DescribeHybridMonitorDataListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHybridMonitorDataListResponseBody) SetMessage(v string) *DescribeHybridMonitorDataListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHybridMonitorDataListResponseBody) SetRequestId(v string) *DescribeHybridMonitorDataListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridMonitorDataListResponseBody) SetSuccess(v string) *DescribeHybridMonitorDataListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHybridMonitorDataListResponseBody) SetTimeSeries(v []*DescribeHybridMonitorDataListResponseBodyTimeSeries) *DescribeHybridMonitorDataListResponseBody {
	s.TimeSeries = v
	return s
}

func (s *DescribeHybridMonitorDataListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridMonitorDataListResponseBodyTimeSeries struct {
	// The tags of the time dimension.
	Labels []*DescribeHybridMonitorDataListResponseBodyTimeSeriesLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The metric name.
	//
	// example:
	//
	// AliyunEcs_cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The metric values that are collected at different timestamps.
	Values []*DescribeHybridMonitorDataListResponseBodyTimeSeriesValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeHybridMonitorDataListResponseBodyTimeSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorDataListResponseBodyTimeSeries) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorDataListResponseBodyTimeSeries) GetLabels() []*DescribeHybridMonitorDataListResponseBodyTimeSeriesLabels {
	return s.Labels
}

func (s *DescribeHybridMonitorDataListResponseBodyTimeSeries) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeHybridMonitorDataListResponseBodyTimeSeries) GetValues() []*DescribeHybridMonitorDataListResponseBodyTimeSeriesValues {
	return s.Values
}

func (s *DescribeHybridMonitorDataListResponseBodyTimeSeries) SetLabels(v []*DescribeHybridMonitorDataListResponseBodyTimeSeriesLabels) *DescribeHybridMonitorDataListResponseBodyTimeSeries {
	s.Labels = v
	return s
}

func (s *DescribeHybridMonitorDataListResponseBodyTimeSeries) SetMetricName(v string) *DescribeHybridMonitorDataListResponseBodyTimeSeries {
	s.MetricName = &v
	return s
}

func (s *DescribeHybridMonitorDataListResponseBodyTimeSeries) SetValues(v []*DescribeHybridMonitorDataListResponseBodyTimeSeriesValues) *DescribeHybridMonitorDataListResponseBodyTimeSeries {
	s.Values = v
	return s
}

func (s *DescribeHybridMonitorDataListResponseBodyTimeSeries) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridMonitorDataListResponseBodyTimeSeriesLabels struct {
	// The tag key.
	//
	// example:
	//
	// instanceId
	K *string `json:"K,omitempty" xml:"K,omitempty"`
	// The tag value.
	//
	// example:
	//
	// i-rj99xc6cptkk64ml****
	V *string `json:"V,omitempty" xml:"V,omitempty"`
}

func (s DescribeHybridMonitorDataListResponseBodyTimeSeriesLabels) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorDataListResponseBodyTimeSeriesLabels) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorDataListResponseBodyTimeSeriesLabels) GetK() *string {
	return s.K
}

func (s *DescribeHybridMonitorDataListResponseBodyTimeSeriesLabels) GetV() *string {
	return s.V
}

func (s *DescribeHybridMonitorDataListResponseBodyTimeSeriesLabels) SetK(v string) *DescribeHybridMonitorDataListResponseBodyTimeSeriesLabels {
	s.K = &v
	return s
}

func (s *DescribeHybridMonitorDataListResponseBodyTimeSeriesLabels) SetV(v string) *DescribeHybridMonitorDataListResponseBodyTimeSeriesLabels {
	s.V = &v
	return s
}

func (s *DescribeHybridMonitorDataListResponseBodyTimeSeriesLabels) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridMonitorDataListResponseBodyTimeSeriesValues struct {
	// The timestamp that indicates the time when the metric value is collected.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 1653804865
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
	// The metric value.
	//
	// example:
	//
	// 0.13
	V *string `json:"V,omitempty" xml:"V,omitempty"`
}

func (s DescribeHybridMonitorDataListResponseBodyTimeSeriesValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorDataListResponseBodyTimeSeriesValues) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorDataListResponseBodyTimeSeriesValues) GetTs() *string {
	return s.Ts
}

func (s *DescribeHybridMonitorDataListResponseBodyTimeSeriesValues) GetV() *string {
	return s.V
}

func (s *DescribeHybridMonitorDataListResponseBodyTimeSeriesValues) SetTs(v string) *DescribeHybridMonitorDataListResponseBodyTimeSeriesValues {
	s.Ts = &v
	return s
}

func (s *DescribeHybridMonitorDataListResponseBodyTimeSeriesValues) SetV(v string) *DescribeHybridMonitorDataListResponseBodyTimeSeriesValues {
	s.V = &v
	return s
}

func (s *DescribeHybridMonitorDataListResponseBodyTimeSeriesValues) Validate() error {
	return dara.Validate(s)
}
