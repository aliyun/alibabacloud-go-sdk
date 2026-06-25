// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeMetricListResponseBody
	GetCount() *int32
	SetMetricTotalModel(v []*DescribeMetricListResponseBodyMetricTotalModel) *DescribeMetricListResponseBody
	GetMetricTotalModel() []*DescribeMetricListResponseBodyMetricTotalModel
	SetNextToken(v string) *DescribeMetricListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeMetricListResponseBody
	GetRequestId() *string
}

type DescribeMetricListResponseBody struct {
	// The number of records.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The instance monitoring data.
	MetricTotalModel []*DescribeMetricListResponseBodyMetricTotalModel `json:"MetricTotalModel,omitempty" xml:"MetricTotalModel,omitempty" type:"Repeated"`
	// The token for the next query. If not all results are returned, this parameter is not empty. Use the returned token to query the next page.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMetricListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricListResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeMetricListResponseBody) GetMetricTotalModel() []*DescribeMetricListResponseBodyMetricTotalModel {
	return s.MetricTotalModel
}

func (s *DescribeMetricListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMetricListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetricListResponseBody) SetCount(v int32) *DescribeMetricListResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetMetricTotalModel(v []*DescribeMetricListResponseBodyMetricTotalModel) *DescribeMetricListResponseBody {
	s.MetricTotalModel = v
	return s
}

func (s *DescribeMetricListResponseBody) SetNextToken(v string) *DescribeMetricListResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetRequestId(v string) *DescribeMetricListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricListResponseBody) Validate() error {
	if s.MetricTotalModel != nil {
		for _, item := range s.MetricTotalModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetricListResponseBodyMetricTotalModel struct {
	// For instance-level monitoring, the Cloud Phone instance ID is returned.
	//
	// <props="china">For matrix-level monitoring, the matrix ID is returned.
	//
	// example:
	//
	// acp-fkuit0cmyru4p****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// np-5hh4a31emkt6u****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The monitoring data for metrics.
	MetricModelList []*DescribeMetricListResponseBodyMetricTotalModelMetricModelList `json:"MetricModelList,omitempty" xml:"MetricModelList,omitempty" type:"Repeated"`
}

func (s DescribeMetricListResponseBodyMetricTotalModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricListResponseBodyMetricTotalModel) GoString() string {
	return s.String()
}

func (s *DescribeMetricListResponseBodyMetricTotalModel) GetAndroidInstanceId() *string {
	return s.AndroidInstanceId
}

func (s *DescribeMetricListResponseBodyMetricTotalModel) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMetricListResponseBodyMetricTotalModel) GetMetricModelList() []*DescribeMetricListResponseBodyMetricTotalModelMetricModelList {
	return s.MetricModelList
}

func (s *DescribeMetricListResponseBodyMetricTotalModel) SetAndroidInstanceId(v string) *DescribeMetricListResponseBodyMetricTotalModel {
	s.AndroidInstanceId = &v
	return s
}

func (s *DescribeMetricListResponseBodyMetricTotalModel) SetInstanceId(v string) *DescribeMetricListResponseBodyMetricTotalModel {
	s.InstanceId = &v
	return s
}

func (s *DescribeMetricListResponseBodyMetricTotalModel) SetMetricModelList(v []*DescribeMetricListResponseBodyMetricTotalModelMetricModelList) *DescribeMetricListResponseBodyMetricTotalModel {
	s.MetricModelList = v
	return s
}

func (s *DescribeMetricListResponseBodyMetricTotalModel) Validate() error {
	if s.MetricModelList != nil {
		for _, item := range s.MetricModelList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetricListResponseBodyMetricTotalModelMetricModelList struct {
	// A list of monitoring data points.
	DataPoints []*DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints `json:"DataPoints,omitempty" xml:"DataPoints,omitempty" type:"Repeated"`
	// The name of the metric.
	//
	// example:
	//
	// cpu_utilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The process name.
	//
	// example:
	//
	// Required
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
}

func (s DescribeMetricListResponseBodyMetricTotalModelMetricModelList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricListResponseBodyMetricTotalModelMetricModelList) GoString() string {
	return s.String()
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelList) GetDataPoints() []*DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints {
	return s.DataPoints
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelList) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelList) GetProcessName() *string {
	return s.ProcessName
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelList) SetDataPoints(v []*DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints) *DescribeMetricListResponseBodyMetricTotalModelMetricModelList {
	s.DataPoints = v
	return s
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelList) SetMetricName(v string) *DescribeMetricListResponseBodyMetricTotalModelMetricModelList {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelList) SetProcessName(v string) *DescribeMetricListResponseBodyMetricTotalModelMetricModelList {
	s.ProcessName = &v
	return s
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelList) Validate() error {
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

type DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints struct {
	// The average value. Unit: %.
	//
	// > The unit is bps when the metric is \\`bandwidth_package_usage_out\\` or \\`bandwidth_package_usage_in\\`.
	//
	// example:
	//
	// 99.52
	Average *float64 `json:"Average,omitempty" xml:"Average,omitempty"`
	// The GPU ID.
	//
	// example:
	//
	// cpn-1t1bxvp9az2pk****-gpu-0
	GpuId *string `json:"GpuId,omitempty" xml:"GpuId,omitempty"`
	// The maximum value. Unit: %.
	//
	// > The unit is bps when the metric is \\`bandwidth_package_usage_out\\` or \\`bandwidth_package_usage_in\\`.
	//
	// example:
	//
	// 100
	Maximum *float64 `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	// The minimum value. Unit: %.
	//
	// > The unit is bps when the metric is \\`bandwidth_package_usage_out\\` or \\`bandwidth_package_usage_in\\`.
	//
	// example:
	//
	// 93.1
	Minimum *float64 `json:"Minimum,omitempty" xml:"Minimum,omitempty"`
	// The timestamp of the statistics. Unit: milliseconds.
	//
	// example:
	//
	// 1548777660000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The data value. Unit: %.
	//
	// > The unit is bps when the metric is \\`bandwidth_package_usage_out\\` or \\`bandwidth_package_usage_in\\`.
	//
	// example:
	//
	// tf-testacc-oos-parameter
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints) GoString() string {
	return s.String()
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints) GetAverage() *float64 {
	return s.Average
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints) GetGpuId() *string {
	return s.GpuId
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints) GetMaximum() *float64 {
	return s.Maximum
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints) GetMinimum() *float64 {
	return s.Minimum
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints) GetValue() *float64 {
	return s.Value
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints) SetAverage(v float64) *DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Average = &v
	return s
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints) SetGpuId(v string) *DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.GpuId = &v
	return s
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints) SetMaximum(v float64) *DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Maximum = &v
	return s
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints) SetMinimum(v float64) *DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Minimum = &v
	return s
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints) SetTimestamp(v int64) *DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Timestamp = &v
	return s
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints) SetValue(v float64) *DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Value = &v
	return s
}

func (s *DescribeMetricListResponseBodyMetricTotalModelMetricModelListDataPoints) Validate() error {
	return dara.Validate(s)
}
