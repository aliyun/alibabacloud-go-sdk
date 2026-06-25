// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricTopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeMetricTopResponseBody
	GetCount() *int32
	SetMetricTotalModel(v []*DescribeMetricTopResponseBodyMetricTotalModel) *DescribeMetricTopResponseBody
	GetMetricTotalModel() []*DescribeMetricTopResponseBodyMetricTotalModel
	SetNextToken(v string) *DescribeMetricTopResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeMetricTopResponseBody
	GetRequestId() *string
}

type DescribeMetricTopResponseBody struct {
	// The number of entries returned.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The instance monitoring data.
	MetricTotalModel []*DescribeMetricTopResponseBodyMetricTotalModel `json:"MetricTotalModel,omitempty" xml:"MetricTotalModel,omitempty" type:"Repeated"`
	// The token to retrieve the next page of results. If the results are not fully returned in a single query, this parameter is not empty. You can use this token in a subsequent query to retrieve the next page of results.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU+SQXzm0H9mu/FiSc****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4610632D-D661-5982-B3D7-5D3FD183F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMetricTopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricTopResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricTopResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeMetricTopResponseBody) GetMetricTotalModel() []*DescribeMetricTopResponseBodyMetricTotalModel {
	return s.MetricTotalModel
}

func (s *DescribeMetricTopResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMetricTopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetricTopResponseBody) SetCount(v int32) *DescribeMetricTopResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeMetricTopResponseBody) SetMetricTotalModel(v []*DescribeMetricTopResponseBodyMetricTotalModel) *DescribeMetricTopResponseBody {
	s.MetricTotalModel = v
	return s
}

func (s *DescribeMetricTopResponseBody) SetNextToken(v string) *DescribeMetricTopResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeMetricTopResponseBody) SetRequestId(v string) *DescribeMetricTopResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricTopResponseBody) Validate() error {
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

type DescribeMetricTopResponseBodyMetricTotalModel struct {
	// For instance-level monitoring, this is the cloud phone instance ID.
	//
	// <props="china">For matrix-level monitoring, this is the matrix ID.
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
	// The monitoring data of metrics.
	MetricModelList []*DescribeMetricTopResponseBodyMetricTotalModelMetricModelList `json:"MetricModelList,omitempty" xml:"MetricModelList,omitempty" type:"Repeated"`
}

func (s DescribeMetricTopResponseBodyMetricTotalModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricTopResponseBodyMetricTotalModel) GoString() string {
	return s.String()
}

func (s *DescribeMetricTopResponseBodyMetricTotalModel) GetAndroidInstanceId() *string {
	return s.AndroidInstanceId
}

func (s *DescribeMetricTopResponseBodyMetricTotalModel) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMetricTopResponseBodyMetricTotalModel) GetMetricModelList() []*DescribeMetricTopResponseBodyMetricTotalModelMetricModelList {
	return s.MetricModelList
}

func (s *DescribeMetricTopResponseBodyMetricTotalModel) SetAndroidInstanceId(v string) *DescribeMetricTopResponseBodyMetricTotalModel {
	s.AndroidInstanceId = &v
	return s
}

func (s *DescribeMetricTopResponseBodyMetricTotalModel) SetInstanceId(v string) *DescribeMetricTopResponseBodyMetricTotalModel {
	s.InstanceId = &v
	return s
}

func (s *DescribeMetricTopResponseBodyMetricTotalModel) SetMetricModelList(v []*DescribeMetricTopResponseBodyMetricTotalModelMetricModelList) *DescribeMetricTopResponseBodyMetricTotalModel {
	s.MetricModelList = v
	return s
}

func (s *DescribeMetricTopResponseBodyMetricTotalModel) Validate() error {
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

type DescribeMetricTopResponseBodyMetricTotalModelMetricModelList struct {
	// The monitoring data points.
	DataPoints []*DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints `json:"DataPoints,omitempty" xml:"DataPoints,omitempty" type:"Repeated"`
	// The name of the metric.
	//
	// example:
	//
	// instance_in_traffic
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
}

func (s DescribeMetricTopResponseBodyMetricTotalModelMetricModelList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricTopResponseBodyMetricTotalModelMetricModelList) GoString() string {
	return s.String()
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelList) GetDataPoints() []*DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints {
	return s.DataPoints
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelList) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelList) SetDataPoints(v []*DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints) *DescribeMetricTopResponseBodyMetricTotalModelMetricModelList {
	s.DataPoints = v
	return s
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelList) SetMetricName(v string) *DescribeMetricTopResponseBodyMetricTotalModelMetricModelList {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelList) Validate() error {
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

type DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints struct {
	// The average value. Unit: %.
	//
	// > The unit is bps when the metric is \\`instance_in_traffic\\` or \\`instance_out_traffic\\`.
	//
	// example:
	//
	// 99.52
	Average *float64 `json:"Average,omitempty" xml:"Average,omitempty"`
	// The ID.
	//
	// example:
	//
	// acp-fkuit0cmyru4p****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The maximum value. Unit: %.
	//
	// > The unit is bps when the metric is \\`instance_in_traffic\\` or \\`instance_out_traffic\\`.
	//
	// example:
	//
	// 100
	Maximum *float64 `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	// The minimum value.
	//
	// > The unit is bps when the metric is \\`instance_in_traffic\\` or \\`instance_out_traffic\\`.
	//
	// example:
	//
	// 93.1
	Minimum *float64 `json:"Minimum,omitempty" xml:"Minimum,omitempty"`
	// The name.
	//
	// example:
	//
	// cloudPhone-****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The properties.
	//
	// > A JSON string that contains the extension properties.
	//
	// example:
	//
	// {"IP":"192.168.*.***"}
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The timestamp of the data point. Unit: milliseconds.
	//
	// example:
	//
	// 1548777660000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints) GoString() string {
	return s.String()
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints) GetAverage() *float64 {
	return s.Average
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints) GetId() *string {
	return s.Id
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints) GetMaximum() *float64 {
	return s.Maximum
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints) GetMinimum() *float64 {
	return s.Minimum
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints) GetName() *string {
	return s.Name
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints) GetProperties() *string {
	return s.Properties
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints) SetAverage(v float64) *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Average = &v
	return s
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints) SetId(v string) *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Id = &v
	return s
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints) SetMaximum(v float64) *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Maximum = &v
	return s
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints) SetMinimum(v float64) *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Minimum = &v
	return s
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints) SetName(v string) *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Name = &v
	return s
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints) SetProperties(v string) *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Properties = &v
	return s
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints) SetTimestamp(v int64) *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints {
	s.Timestamp = &v
	return s
}

func (s *DescribeMetricTopResponseBodyMetricTotalModelMetricModelListDataPoints) Validate() error {
	return dara.Validate(s)
}
