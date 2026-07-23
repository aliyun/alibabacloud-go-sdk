// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataAgentMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDataAgentMetricsResponseBodyData) *DescribeDataAgentMetricsResponseBody
	GetData() *DescribeDataAgentMetricsResponseBodyData
	SetErrorCode(v string) *DescribeDataAgentMetricsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DescribeDataAgentMetricsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DescribeDataAgentMetricsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDataAgentMetricsResponseBody
	GetSuccess() *bool
}

type DescribeDataAgentMetricsResponseBody struct {
	// The response struct.
	Data *DescribeDataAgentMetricsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned when the request is abnormal.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDataAgentMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentMetricsResponseBody) GetData() *DescribeDataAgentMetricsResponseBodyData {
	return s.Data
}

func (s *DescribeDataAgentMetricsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeDataAgentMetricsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDataAgentMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataAgentMetricsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDataAgentMetricsResponseBody) SetData(v *DescribeDataAgentMetricsResponseBodyData) *DescribeDataAgentMetricsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDataAgentMetricsResponseBody) SetErrorCode(v string) *DescribeDataAgentMetricsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDataAgentMetricsResponseBody) SetErrorMessage(v string) *DescribeDataAgentMetricsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDataAgentMetricsResponseBody) SetRequestId(v string) *DescribeDataAgentMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataAgentMetricsResponseBody) SetSuccess(v bool) *DescribeDataAgentMetricsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDataAgentMetricsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataAgentMetricsResponseBodyData struct {
	// The end time of the query range.
	//
	// example:
	//
	// 1782836200000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The metric type.
	//
	// example:
	//
	// basic
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The list of metrics.
	Metrics []*DescribeDataAgentMetricsResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The start time of the query range.
	//
	// example:
	//
	// 1782835200000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDataAgentMetricsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentMetricsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentMetricsResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDataAgentMetricsResponseBodyData) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeDataAgentMetricsResponseBodyData) GetMetrics() []*DescribeDataAgentMetricsResponseBodyDataMetrics {
	return s.Metrics
}

func (s *DescribeDataAgentMetricsResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDataAgentMetricsResponseBodyData) SetEndTime(v int64) *DescribeDataAgentMetricsResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeDataAgentMetricsResponseBodyData) SetMetricType(v string) *DescribeDataAgentMetricsResponseBodyData {
	s.MetricType = &v
	return s
}

func (s *DescribeDataAgentMetricsResponseBodyData) SetMetrics(v []*DescribeDataAgentMetricsResponseBodyDataMetrics) *DescribeDataAgentMetricsResponseBodyData {
	s.Metrics = v
	return s
}

func (s *DescribeDataAgentMetricsResponseBodyData) SetStartTime(v int64) *DescribeDataAgentMetricsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeDataAgentMetricsResponseBodyData) Validate() error {
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

type DescribeDataAgentMetricsResponseBodyDataMetrics struct {
	// The error message returned when the call fails.
	//
	// example:
	//
	// Timeout
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The metric name.
	//
	// example:
	//
	// data_agent_session_per_user
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The metric value.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDataAgentMetricsResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentMetricsResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentMetricsResponseBodyDataMetrics) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDataAgentMetricsResponseBodyDataMetrics) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeDataAgentMetricsResponseBodyDataMetrics) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDataAgentMetricsResponseBodyDataMetrics) GetValue() *string {
	return s.Value
}

func (s *DescribeDataAgentMetricsResponseBodyDataMetrics) SetErrorMessage(v string) *DescribeDataAgentMetricsResponseBodyDataMetrics {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDataAgentMetricsResponseBodyDataMetrics) SetMetricName(v string) *DescribeDataAgentMetricsResponseBodyDataMetrics {
	s.MetricName = &v
	return s
}

func (s *DescribeDataAgentMetricsResponseBodyDataMetrics) SetSuccess(v bool) *DescribeDataAgentMetricsResponseBodyDataMetrics {
	s.Success = &v
	return s
}

func (s *DescribeDataAgentMetricsResponseBodyDataMetrics) SetValue(v string) *DescribeDataAgentMetricsResponseBodyDataMetrics {
	s.Value = &v
	return s
}

func (s *DescribeDataAgentMetricsResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}
