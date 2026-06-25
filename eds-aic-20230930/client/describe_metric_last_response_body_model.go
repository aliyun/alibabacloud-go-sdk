// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricLastResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeMetricLastResponseBody
	GetCount() *int32
	SetMetricTotalModel(v []*DescribeMetricLastResponseBodyMetricTotalModel) *DescribeMetricLastResponseBody
	GetMetricTotalModel() []*DescribeMetricLastResponseBodyMetricTotalModel
	SetNextToken(v string) *DescribeMetricLastResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeMetricLastResponseBody
	GetRequestId() *string
}

type DescribeMetricLastResponseBody struct {
	// The number of records.
	//
	// example:
	//
	// 100
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The instance monitoring data.
	MetricTotalModel []*DescribeMetricLastResponseBodyMetricTotalModel `json:"MetricTotalModel,omitempty" xml:"MetricTotalModel,omitempty" type:"Repeated"`
	// The token for the next query. If not all results are returned in a single query, the returned NextToken is not empty. You can use the returned NextToken in the next query to query the next page.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2B9E6946-0E2A-5D2B-B275-361DF81F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMetricLastResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricLastResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricLastResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeMetricLastResponseBody) GetMetricTotalModel() []*DescribeMetricLastResponseBodyMetricTotalModel {
	return s.MetricTotalModel
}

func (s *DescribeMetricLastResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMetricLastResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetricLastResponseBody) SetCount(v int32) *DescribeMetricLastResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeMetricLastResponseBody) SetMetricTotalModel(v []*DescribeMetricLastResponseBodyMetricTotalModel) *DescribeMetricLastResponseBody {
	s.MetricTotalModel = v
	return s
}

func (s *DescribeMetricLastResponseBody) SetNextToken(v string) *DescribeMetricLastResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeMetricLastResponseBody) SetRequestId(v string) *DescribeMetricLastResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricLastResponseBody) Validate() error {
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

type DescribeMetricLastResponseBodyMetricTotalModel struct {
	// For instance-level monitoring, the Cloud Phone instance ID is returned.
	//
	// <props="china">For matrix-level monitoring, the matrix ID is returned.
	//
	// example:
	//
	// acp-fkuit0cmyru4p****
	AndroidInstanceId *string `json:"AndroidInstanceId,omitempty" xml:"AndroidInstanceId,omitempty"`
	// The monitoring data for metrics.
	MetricModelList []*DescribeMetricLastResponseBodyMetricTotalModelMetricModelList `json:"MetricModelList,omitempty" xml:"MetricModelList,omitempty" type:"Repeated"`
}

func (s DescribeMetricLastResponseBodyMetricTotalModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricLastResponseBodyMetricTotalModel) GoString() string {
	return s.String()
}

func (s *DescribeMetricLastResponseBodyMetricTotalModel) GetAndroidInstanceId() *string {
	return s.AndroidInstanceId
}

func (s *DescribeMetricLastResponseBodyMetricTotalModel) GetMetricModelList() []*DescribeMetricLastResponseBodyMetricTotalModelMetricModelList {
	return s.MetricModelList
}

func (s *DescribeMetricLastResponseBodyMetricTotalModel) SetAndroidInstanceId(v string) *DescribeMetricLastResponseBodyMetricTotalModel {
	s.AndroidInstanceId = &v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModel) SetMetricModelList(v []*DescribeMetricLastResponseBodyMetricTotalModelMetricModelList) *DescribeMetricLastResponseBodyMetricTotalModel {
	s.MetricModelList = v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModel) Validate() error {
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

type DescribeMetricLastResponseBodyMetricTotalModelMetricModelList struct {
	// The monitoring data points.
	DataPoints []*DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints `json:"DataPoints,omitempty" xml:"DataPoints,omitempty" type:"Repeated"`
	// The name of the metric.
	//
	// example:
	//
	// cpu_utilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The process monitoring metrics.
	//
	// > This parameter is returned when the metric is \\`process_cpu\\` or \\`process_memory\\`.
	ProcessLastInfos []*DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos `json:"ProcessLastInfos,omitempty" xml:"ProcessLastInfos,omitempty" type:"Repeated"`
}

func (s DescribeMetricLastResponseBodyMetricTotalModelMetricModelList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricLastResponseBodyMetricTotalModelMetricModelList) GoString() string {
	return s.String()
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelList) GetDataPoints() []*DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints {
	return s.DataPoints
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelList) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelList) GetProcessLastInfos() []*DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos {
	return s.ProcessLastInfos
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelList) SetDataPoints(v []*DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelList {
	s.DataPoints = v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelList) SetMetricName(v string) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelList {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelList) SetProcessLastInfos(v []*DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelList {
	s.ProcessLastInfos = v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelList) Validate() error {
	if s.DataPoints != nil {
		for _, item := range s.DataPoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ProcessLastInfos != nil {
		for _, item := range s.ProcessLastInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints struct {
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
}

func (s DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) GoString() string {
	return s.String()
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) GetAverage() *float64 {
	return s.Average
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) GetGpuId() *string {
	return s.GpuId
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) GetMaximum() *float64 {
	return s.Maximum
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) GetMinimum() *float64 {
	return s.Minimum
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) SetAverage(v float64) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Average = &v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) SetGpuId(v string) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.GpuId = &v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) SetMaximum(v float64) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Maximum = &v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) SetMinimum(v float64) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Minimum = &v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) SetTimestamp(v int64) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Timestamp = &v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListDataPoints) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos struct {
	// The CPU utilization. Unit: %.
	//
	// example:
	//
	// 50
	CpuUsage *float64 `json:"CpuUsage,omitempty" xml:"CpuUsage,omitempty"`
	// The memory usage. Unit: %.
	//
	// example:
	//
	// 50
	MemoryUsage *float64 `json:"MemoryUsage,omitempty" xml:"MemoryUsage,omitempty"`
	// The process name.
	//
	// example:
	//
	// com.offerup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of process IDs.
	ProcessIds []*int32 `json:"ProcessIds,omitempty" xml:"ProcessIds,omitempty" type:"Repeated"`
	// The timestamp of the statistics. Unit: milliseconds.
	//
	// example:
	//
	// 1548777660000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) GoString() string {
	return s.String()
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) GetCpuUsage() *float64 {
	return s.CpuUsage
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) GetMemoryUsage() *float64 {
	return s.MemoryUsage
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) GetName() *string {
	return s.Name
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) GetProcessIds() []*int32 {
	return s.ProcessIds
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) SetCpuUsage(v float64) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos {
	s.CpuUsage = &v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) SetMemoryUsage(v float64) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos {
	s.MemoryUsage = &v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) SetName(v string) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos {
	s.Name = &v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) SetProcessIds(v []*int32) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos {
	s.ProcessIds = v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) SetTimestamp(v int64) *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos {
	s.Timestamp = &v
	return s
}

func (s *DescribeMetricLastResponseBodyMetricTotalModelMetricModelListProcessLastInfos) Validate() error {
	return dara.Validate(s)
}
