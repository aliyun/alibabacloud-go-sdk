// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceAutoScalerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBehavior(v map[string]interface{}) *DescribeServiceAutoScalerResponseBody
	GetBehavior() map[string]interface{}
	SetCurrentMetrics(v []*DescribeServiceAutoScalerResponseBodyCurrentMetrics) *DescribeServiceAutoScalerResponseBody
	GetCurrentMetrics() []*DescribeServiceAutoScalerResponseBodyCurrentMetrics
	SetMaxReplica(v int32) *DescribeServiceAutoScalerResponseBody
	GetMaxReplica() *int32
	SetMinReplica(v int32) *DescribeServiceAutoScalerResponseBody
	GetMinReplica() *int32
	SetRequestId(v string) *DescribeServiceAutoScalerResponseBody
	GetRequestId() *string
	SetScaleStrategies(v []*DescribeServiceAutoScalerResponseBodyScaleStrategies) *DescribeServiceAutoScalerResponseBody
	GetScaleStrategies() []*DescribeServiceAutoScalerResponseBodyScaleStrategies
	SetServiceName(v string) *DescribeServiceAutoScalerResponseBody
	GetServiceName() *string
}

type DescribeServiceAutoScalerResponseBody struct {
	// The additional information about the Autoscaler policy, such as the interval of triggering Autoscaler.
	//
	// example:
	//
	// {
	//
	//   "behavior": {
	//
	//     "scaleDown": {
	//
	//       "stabilizationWindowSeconds": 150
	//
	//     }
	//
	//   }
	//
	// }
	Behavior map[string]interface{} `json:"Behavior,omitempty" xml:"Behavior,omitempty"`
	// The metrics.
	CurrentMetrics []*DescribeServiceAutoScalerResponseBodyCurrentMetrics `json:"CurrentMetrics,omitempty" xml:"CurrentMetrics,omitempty" type:"Repeated"`
	// The maximum number of instances in the service.
	//
	// example:
	//
	// 8
	MaxReplica *int32 `json:"MaxReplica,omitempty" xml:"MaxReplica,omitempty"`
	// The minimum number of instances in the service.
	//
	// example:
	//
	// 3
	MinReplica *int32 `json:"MinReplica,omitempty" xml:"MinReplica,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The auto scaling policies.
	ScaleStrategies []*DescribeServiceAutoScalerResponseBodyScaleStrategies `json:"ScaleStrategies,omitempty" xml:"ScaleStrategies,omitempty" type:"Repeated"`
	// The service name.
	//
	// example:
	//
	// foo
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeServiceAutoScalerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceAutoScalerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceAutoScalerResponseBody) GetBehavior() map[string]interface{} {
	return s.Behavior
}

func (s *DescribeServiceAutoScalerResponseBody) GetCurrentMetrics() []*DescribeServiceAutoScalerResponseBodyCurrentMetrics {
	return s.CurrentMetrics
}

func (s *DescribeServiceAutoScalerResponseBody) GetMaxReplica() *int32 {
	return s.MaxReplica
}

func (s *DescribeServiceAutoScalerResponseBody) GetMinReplica() *int32 {
	return s.MinReplica
}

func (s *DescribeServiceAutoScalerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceAutoScalerResponseBody) GetScaleStrategies() []*DescribeServiceAutoScalerResponseBodyScaleStrategies {
	return s.ScaleStrategies
}

func (s *DescribeServiceAutoScalerResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeServiceAutoScalerResponseBody) SetBehavior(v map[string]interface{}) *DescribeServiceAutoScalerResponseBody {
	s.Behavior = v
	return s
}

func (s *DescribeServiceAutoScalerResponseBody) SetCurrentMetrics(v []*DescribeServiceAutoScalerResponseBodyCurrentMetrics) *DescribeServiceAutoScalerResponseBody {
	s.CurrentMetrics = v
	return s
}

func (s *DescribeServiceAutoScalerResponseBody) SetMaxReplica(v int32) *DescribeServiceAutoScalerResponseBody {
	s.MaxReplica = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBody) SetMinReplica(v int32) *DescribeServiceAutoScalerResponseBody {
	s.MinReplica = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBody) SetRequestId(v string) *DescribeServiceAutoScalerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBody) SetScaleStrategies(v []*DescribeServiceAutoScalerResponseBodyScaleStrategies) *DescribeServiceAutoScalerResponseBody {
	s.ScaleStrategies = v
	return s
}

func (s *DescribeServiceAutoScalerResponseBody) SetServiceName(v string) *DescribeServiceAutoScalerResponseBody {
	s.ServiceName = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceAutoScalerResponseBodyCurrentMetrics struct {
	// The metric name. Valid values:
	//
	// 	- QPS
	//
	// 	- CPU
	//
	// example:
	//
	// qps
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// The service for which the metric is specified.
	//
	// example:
	//
	// demo_svc
	Service *string `json:"service,omitempty" xml:"service,omitempty"`
	// The metric value.
	//
	// example:
	//
	// 10
	Value *float32 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeServiceAutoScalerResponseBodyCurrentMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceAutoScalerResponseBodyCurrentMetrics) GoString() string {
	return s.String()
}

func (s *DescribeServiceAutoScalerResponseBodyCurrentMetrics) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeServiceAutoScalerResponseBodyCurrentMetrics) GetService() *string {
	return s.Service
}

func (s *DescribeServiceAutoScalerResponseBodyCurrentMetrics) GetValue() *float32 {
	return s.Value
}

func (s *DescribeServiceAutoScalerResponseBodyCurrentMetrics) SetMetricName(v string) *DescribeServiceAutoScalerResponseBodyCurrentMetrics {
	s.MetricName = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBodyCurrentMetrics) SetService(v string) *DescribeServiceAutoScalerResponseBodyCurrentMetrics {
	s.Service = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBodyCurrentMetrics) SetValue(v float32) *DescribeServiceAutoScalerResponseBodyCurrentMetrics {
	s.Value = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBodyCurrentMetrics) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceAutoScalerResponseBodyScaleStrategies struct {
	// The metric name. Valid values:
	//
	// 	- QPS: the queries per second (QPS) for an individual instance.
	//
	// 	- CPU: the CPU utilization.
	//
	// example:
	//
	// QPS
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// The service for which the metric is specified. If you do not set this parameter, the current service is specified by default.
	//
	// example:
	//
	// demo_svc
	Service *string `json:"service,omitempty" xml:"service,omitempty"`
	// The threshold of the metric that triggers auto scaling.
	//
	// 	- If you set metricName to QPS, scale-out is triggered when the average QPS for a single instance is greater than this threshold.
	//
	// 	- If you set metricName to CPU, scale-out is triggered when the average CPU utilization for a single instance is greater than this threshold.
	//
	// example:
	//
	// 10
	Threshold *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s DescribeServiceAutoScalerResponseBodyScaleStrategies) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceAutoScalerResponseBodyScaleStrategies) GoString() string {
	return s.String()
}

func (s *DescribeServiceAutoScalerResponseBodyScaleStrategies) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeServiceAutoScalerResponseBodyScaleStrategies) GetService() *string {
	return s.Service
}

func (s *DescribeServiceAutoScalerResponseBodyScaleStrategies) GetThreshold() *float32 {
	return s.Threshold
}

func (s *DescribeServiceAutoScalerResponseBodyScaleStrategies) SetMetricName(v string) *DescribeServiceAutoScalerResponseBodyScaleStrategies {
	s.MetricName = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBodyScaleStrategies) SetService(v string) *DescribeServiceAutoScalerResponseBodyScaleStrategies {
	s.Service = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBodyScaleStrategies) SetThreshold(v float32) *DescribeServiceAutoScalerResponseBodyScaleStrategies {
	s.Threshold = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBodyScaleStrategies) Validate() error {
	return dara.Validate(s)
}
