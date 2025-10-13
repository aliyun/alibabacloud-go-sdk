// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeApplicationScalingRuleResponseBody
	GetCode() *string
	SetData(v *DescribeApplicationScalingRuleResponseBodyData) *DescribeApplicationScalingRuleResponseBody
	GetData() *DescribeApplicationScalingRuleResponseBodyData
	SetErrorCode(v string) *DescribeApplicationScalingRuleResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeApplicationScalingRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeApplicationScalingRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeApplicationScalingRuleResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeApplicationScalingRuleResponseBody
	GetTraceId() *string
}

type DescribeApplicationScalingRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data      *DescribeApplicationScalingRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 73404D3D-EE4F-4CB2-B197-5C46F6A1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. The ID is used to query the details of a request.
	//
	// example:
	//
	// 0b57ff7e16243300839193068e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeApplicationScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeApplicationScalingRuleResponseBody) GetData() *DescribeApplicationScalingRuleResponseBodyData {
	return s.Data
}

func (s *DescribeApplicationScalingRuleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeApplicationScalingRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeApplicationScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationScalingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeApplicationScalingRuleResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeApplicationScalingRuleResponseBody) SetCode(v string) *DescribeApplicationScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBody) SetData(v *DescribeApplicationScalingRuleResponseBodyData) *DescribeApplicationScalingRuleResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBody) SetErrorCode(v string) *DescribeApplicationScalingRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBody) SetMessage(v string) *DescribeApplicationScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBody) SetRequestId(v string) *DescribeApplicationScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBody) SetSuccess(v bool) *DescribeApplicationScalingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBody) SetTraceId(v string) *DescribeApplicationScalingRuleResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationScalingRuleResponseBodyData struct {
	// The ID of the application.
	//
	// example:
	//
	// a0d2e04c-159d-40a8-b240-d2f2c263****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the auto scaling policy was created. Unit: milliseconds.
	//
	// example:
	//
	// 1624329843790
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the auto scaling policy was last disabled.
	//
	// example:
	//
	// 1641882854484
	LastDisableTime *int64 `json:"LastDisableTime,omitempty" xml:"LastDisableTime,omitempty"`
	// The details of the metric-based auto scaling policy.
	Metric *DescribeApplicationScalingRuleResponseBodyDataMetric `json:"Metric,omitempty" xml:"Metric,omitempty" type:"Struct"`
	// The ratio of the minimum number of available instances to the current number of instances. Valid values:
	//
	// 	- **-1*	- (default value): The minimum number of available instances is not determined based on this parameter.
	//
	// 	- **0 to 100**: The minimum number of available instances is calculated by using the following formula: Number of existing instances × Value of MinReadyInstanceRatio × 100%. The calculation result is rounded up to the nearest integer. For example, if the number of existing instances is 5 and MinReadyInstanceRatio is set to 50, the minimum number of available instances is 3.
	//
	// >  If the **MinReadyInstanceRatio*	- and **MinReadyInstanceRatio*	- parameters are configured and the **MinReadyInstanceRatio*	- parameter is set to a number from 0 to 100, the value of the MinReadyInstanceRatio parameter takes precedence. For example, if the **MinReadyInstances*	- parameter is set to **5**, and the **MinReadyInstanceRatio*	- parameter is set to **50**, the minimum number of available instances is set to the nearest integer rounded up from the calculated result of the following formula: Nmber of existing instances × **50**.
	//
	// example:
	//
	// -1
	MinReadyInstanceRatio *int32 `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	// The minimum number of available instances. Valid values:
	//
	// 	- If you set the value to **0**, business is interrupted when the application is updated.
	//
	// 	- If you set this property to -1, the system calculates a recommended value as the minimum number of available instances by using the following formula: Recommended value = Number of existing instances × 25%. The calculation result is rounded up to the nearest integer. For example, if the number of existing instances is 5, the recommended value is calculated by using the following formula: 5 × 25% = 1.25. In this case, the minimum number of available instances is 2.
	//
	// >  To ensure business continuity, make sure that at least one instance is available during application deployment and rollback.
	//
	// example:
	//
	// 1
	MinReadyInstances *int32 `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
	// Indicates whether the auto scaling policy is enabled. Valid values:
	//
	// 	- **true**: enabled
	//
	// 	- **false**: disabled
	//
	// example:
	//
	// true
	ScaleRuleEnabled *bool `json:"ScaleRuleEnabled,omitempty" xml:"ScaleRuleEnabled,omitempty"`
	// The name of the auto scaling policy.
	//
	// example:
	//
	// test
	ScaleRuleName *string `json:"ScaleRuleName,omitempty" xml:"ScaleRuleName,omitempty"`
	// The type of the auto scaling policy. Valid values:
	//
	// 	- **timing**: the scheduled auto scaling policy.
	//
	// 	- **metric**: the metric-based auto scaling policy.
	//
	// 	- **mix**: the hybrid auto scaling policy.
	//
	// example:
	//
	// timing
	ScaleRuleType *string `json:"ScaleRuleType,omitempty" xml:"ScaleRuleType,omitempty"`
	// The details of the scheduled auto scaling policy.
	Timer *DescribeApplicationScalingRuleResponseBodyDataTimer `json:"Timer,omitempty" xml:"Timer,omitempty" type:"Struct"`
	// The time when the auto scaling policy was updated. Unit: milliseconds.
	//
	// example:
	//
	// 1624330075827
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeApplicationScalingRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *DescribeApplicationScalingRuleResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeApplicationScalingRuleResponseBodyData) GetLastDisableTime() *int64 {
	return s.LastDisableTime
}

func (s *DescribeApplicationScalingRuleResponseBodyData) GetMetric() *DescribeApplicationScalingRuleResponseBodyDataMetric {
	return s.Metric
}

func (s *DescribeApplicationScalingRuleResponseBodyData) GetMinReadyInstanceRatio() *int32 {
	return s.MinReadyInstanceRatio
}

func (s *DescribeApplicationScalingRuleResponseBodyData) GetMinReadyInstances() *int32 {
	return s.MinReadyInstances
}

func (s *DescribeApplicationScalingRuleResponseBodyData) GetScaleRuleEnabled() *bool {
	return s.ScaleRuleEnabled
}

func (s *DescribeApplicationScalingRuleResponseBodyData) GetScaleRuleName() *string {
	return s.ScaleRuleName
}

func (s *DescribeApplicationScalingRuleResponseBodyData) GetScaleRuleType() *string {
	return s.ScaleRuleType
}

func (s *DescribeApplicationScalingRuleResponseBodyData) GetTimer() *DescribeApplicationScalingRuleResponseBodyDataTimer {
	return s.Timer
}

func (s *DescribeApplicationScalingRuleResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetAppId(v string) *DescribeApplicationScalingRuleResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetCreateTime(v int64) *DescribeApplicationScalingRuleResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetLastDisableTime(v int64) *DescribeApplicationScalingRuleResponseBodyData {
	s.LastDisableTime = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetMetric(v *DescribeApplicationScalingRuleResponseBodyDataMetric) *DescribeApplicationScalingRuleResponseBodyData {
	s.Metric = v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetMinReadyInstanceRatio(v int32) *DescribeApplicationScalingRuleResponseBodyData {
	s.MinReadyInstanceRatio = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetMinReadyInstances(v int32) *DescribeApplicationScalingRuleResponseBodyData {
	s.MinReadyInstances = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetScaleRuleEnabled(v bool) *DescribeApplicationScalingRuleResponseBodyData {
	s.ScaleRuleEnabled = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetScaleRuleName(v string) *DescribeApplicationScalingRuleResponseBodyData {
	s.ScaleRuleName = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetScaleRuleType(v string) *DescribeApplicationScalingRuleResponseBodyData {
	s.ScaleRuleType = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetTimer(v *DescribeApplicationScalingRuleResponseBodyDataTimer) *DescribeApplicationScalingRuleResponseBodyData {
	s.Timer = v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyData) SetUpdateTime(v int64) *DescribeApplicationScalingRuleResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyData) Validate() error {
	if s.Metric != nil {
		if err := s.Metric.Validate(); err != nil {
			return err
		}
	}
	if s.Timer != nil {
		if err := s.Timer.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationScalingRuleResponseBodyDataMetric struct {
	// The maximum number of instances.
	//
	// example:
	//
	// 3
	MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// The list of metrics that are used to trigger the auto scaling policy.
	Metrics []*DescribeApplicationScalingRuleResponseBodyDataMetricMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The execution status of the metric-based auto scaling policy.
	MetricsStatus *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus `json:"MetricsStatus,omitempty" xml:"MetricsStatus,omitempty" type:"Struct"`
	// The minimum number of instances.
	//
	// example:
	//
	// 1
	MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	// Rules that determine the application scale-in.
	ScaleDownRules *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules `json:"ScaleDownRules,omitempty" xml:"ScaleDownRules,omitempty" type:"Struct"`
	// Rules that determine the application scale-out.
	ScaleUpRules *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules `json:"ScaleUpRules,omitempty" xml:"ScaleUpRules,omitempty" type:"Struct"`
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetric) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetric) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetric) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetric) GetMetrics() []*DescribeApplicationScalingRuleResponseBodyDataMetricMetrics {
	return s.Metrics
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetric) GetMetricsStatus() *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus {
	return s.MetricsStatus
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetric) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetric) GetScaleDownRules() *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules {
	return s.ScaleDownRules
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetric) GetScaleUpRules() *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules {
	return s.ScaleUpRules
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetric) SetMaxReplicas(v int32) *DescribeApplicationScalingRuleResponseBodyDataMetric {
	s.MaxReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetric) SetMetrics(v []*DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) *DescribeApplicationScalingRuleResponseBodyDataMetric {
	s.Metrics = v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetric) SetMetricsStatus(v *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) *DescribeApplicationScalingRuleResponseBodyDataMetric {
	s.MetricsStatus = v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetric) SetMinReplicas(v int32) *DescribeApplicationScalingRuleResponseBodyDataMetric {
	s.MinReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetric) SetScaleDownRules(v *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules) *DescribeApplicationScalingRuleResponseBodyDataMetric {
	s.ScaleDownRules = v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetric) SetScaleUpRules(v *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules) *DescribeApplicationScalingRuleResponseBodyDataMetric {
	s.ScaleUpRules = v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetric) Validate() error {
	if s.Metrics != nil {
		for _, item := range s.Metrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MetricsStatus != nil {
		if err := s.MetricsStatus.Validate(); err != nil {
			return err
		}
	}
	if s.ScaleDownRules != nil {
		if err := s.ScaleDownRules.Validate(); err != nil {
			return err
		}
	}
	if s.ScaleUpRules != nil {
		if err := s.ScaleUpRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationScalingRuleResponseBodyDataMetricMetrics struct {
	// The limit on the metric.
	//
	// 	- The limit on the CPU utilization. Unit: percentage.
	//
	// 	- The limit on the memory usage. Unit: percentage.
	//
	// 	- The limit on the average number of active TCP connections per second.
	//
	// 	- The limit on the QPS of the Internet-facing SLB instance.
	//
	// 	- The limit on the response time of the Internet-facing SLB instance. Unit: milliseconds.
	//
	// example:
	//
	// 20
	MetricTargetAverageUtilization *int32 `json:"MetricTargetAverageUtilization,omitempty" xml:"MetricTargetAverageUtilization,omitempty"`
	// The metric that is used to trigger the auto scaling policy. Valid values:
	//
	// 	- **CPU**: the CPU utilization.
	//
	// 	- **MEMORY**: the memory usage.
	//
	// 	- **tcpActiveConn**: the average number of active TCP connections for an instance in 30 seconds.
	//
	// 	- **SLB_QPS**: the average QPS of the Internet-facing SLB instance associated with an application instance in 15 seconds.
	//
	// 	- **SLB_RT**: the average response time of the Internet-facing SLB instance in 15 seconds.
	//
	// example:
	//
	// CPU
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// example:
	//
	// lb-xxx
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// example:
	//
	// test
	SlbLogstore *string `json:"SlbLogstore,omitempty" xml:"SlbLogstore,omitempty"`
	// example:
	//
	// test
	SlbProject *string `json:"SlbProject,omitempty" xml:"SlbProject,omitempty"`
	// example:
	//
	// 80
	Vport *string `json:"Vport,omitempty" xml:"Vport,omitempty"`
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) GetMetricTargetAverageUtilization() *int32 {
	return s.MetricTargetAverageUtilization
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) GetSlbId() *string {
	return s.SlbId
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) GetSlbLogstore() *string {
	return s.SlbLogstore
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) GetSlbProject() *string {
	return s.SlbProject
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) GetVport() *string {
	return s.Vport
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) SetMetricTargetAverageUtilization(v int32) *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.MetricTargetAverageUtilization = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) SetMetricType(v string) *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.MetricType = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) SetSlbId(v string) *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.SlbId = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) SetSlbLogstore(v string) *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.SlbLogstore = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) SetSlbProject(v string) *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.SlbProject = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) SetVport(v string) *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.Vport = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetrics) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus struct {
	// The metrics that is used to trigger the current auto scaling policy.
	CurrentMetrics []*DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics `json:"CurrentMetrics,omitempty" xml:"CurrentMetrics,omitempty" type:"Repeated"`
	// The current number of instances.
	//
	// example:
	//
	// 2
	CurrentReplicas *int64 `json:"CurrentReplicas,omitempty" xml:"CurrentReplicas,omitempty"`
	// The expected number of instances.
	//
	// example:
	//
	// 2
	DesiredReplicas *int64 `json:"DesiredReplicas,omitempty" xml:"DesiredReplicas,omitempty"`
	// The time when the auto scaling policy was last triggered.
	//
	// example:
	//
	// 2022-01-11T08:14:32Z
	LastScaleTime *string `json:"LastScaleTime,omitempty" xml:"LastScaleTime,omitempty"`
	// The metrics that are used to trigger the auto scaling policy next time.
	NextScaleMetrics []*DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics `json:"NextScaleMetrics,omitempty" xml:"NextScaleMetrics,omitempty" type:"Repeated"`
	// The duration for which the metric-based auto scaling policy takes effect next time.
	//
	// example:
	//
	// 3
	NextScaleTimePeriod *int32 `json:"NextScaleTimePeriod,omitempty" xml:"NextScaleTimePeriod,omitempty"`
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) GetCurrentMetrics() []*DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics {
	return s.CurrentMetrics
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) GetCurrentReplicas() *int64 {
	return s.CurrentReplicas
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) GetDesiredReplicas() *int64 {
	return s.DesiredReplicas
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) GetLastScaleTime() *string {
	return s.LastScaleTime
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) GetNextScaleMetrics() []*DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics {
	return s.NextScaleMetrics
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) GetNextScaleTimePeriod() *int32 {
	return s.NextScaleTimePeriod
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) SetCurrentMetrics(v []*DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus {
	s.CurrentMetrics = v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) SetCurrentReplicas(v int64) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus {
	s.CurrentReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) SetDesiredReplicas(v int64) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus {
	s.DesiredReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) SetLastScaleTime(v string) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus {
	s.LastScaleTime = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) SetNextScaleMetrics(v []*DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus {
	s.NextScaleMetrics = v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) SetNextScaleTimePeriod(v int32) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus {
	s.NextScaleTimePeriod = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus) Validate() error {
	if s.CurrentMetrics != nil {
		for _, item := range s.CurrentMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NextScaleMetrics != nil {
		for _, item := range s.NextScaleMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics struct {
	// The current value of the metric.
	//
	// example:
	//
	// 0
	CurrentValue *int64 `json:"CurrentValue,omitempty" xml:"CurrentValue,omitempty"`
	// The name of the metric.
	//
	// 	- **cpu**: the CPU utilization.
	//
	// 	- **memory**: the memory usage.
	//
	// 	- **tcpActiveConn**: the number of active TCP connections.
	//
	// 	- **slb_incall_qps**: the QPS of the Internet-facing SLB instance.
	//
	// 	- **slb_incall_rt**: the response time of the Internet-facing SLB instance.
	//
	// example:
	//
	// cpu
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the data. This parameter corresponds to the metric.
	//
	// 	- **Resource**: used when the metric is the **CPU utilization*	- or **memory usage**.
	//
	// 	- **Pods**: used when the metric is the **number of active TCP connections**.
	//
	// 	- **External**: used when the metric is about the **SLB*	- instance or from **Application Real-Time Monitoring Service (ARMS)**.
	//
	// example:
	//
	// Resource
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics) GetCurrentValue() *int64 {
	return s.CurrentValue
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics) GetType() *string {
	return s.Type
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics) SetCurrentValue(v int64) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics {
	s.CurrentValue = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics) SetName(v string) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics {
	s.Name = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics) SetType(v string) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics {
	s.Type = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics struct {
	// The name of the metric.
	//
	// 	- **cpu**: the CPU utilization.
	//
	// 	- **memory**: the memory usage.
	//
	// 	- **tcpActiveConn**: the number of active TCP connections.
	//
	// 	- **slb_incall_qps**: the QPS of the Internet-facing SLB instance.
	//
	// 	- **slb_incall_rt**: the response time of the Internet-facing SLB instance.
	//
	// example:
	//
	// cpu
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The metric value as a percentage that triggers the application scale-in next time.
	//
	// example:
	//
	// 10
	NextScaleInAverageUtilization *int32 `json:"NextScaleInAverageUtilization,omitempty" xml:"NextScaleInAverageUtilization,omitempty"`
	// The metric value as a percentage that triggers the application scale-out next time.
	//
	// example:
	//
	// 21
	NextScaleOutAverageUtilization *int32 `json:"NextScaleOutAverageUtilization,omitempty" xml:"NextScaleOutAverageUtilization,omitempty"`
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics) GetNextScaleInAverageUtilization() *int32 {
	return s.NextScaleInAverageUtilization
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics) GetNextScaleOutAverageUtilization() *int32 {
	return s.NextScaleOutAverageUtilization
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics) SetName(v string) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics {
	s.Name = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics) SetNextScaleInAverageUtilization(v int32) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics {
	s.NextScaleInAverageUtilization = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics) SetNextScaleOutAverageUtilization(v int32) *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics {
	s.NextScaleOutAverageUtilization = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules struct {
	// Indicates whether the application scale-in is disabled. Valid values:
	//
	// 	- **true**: disabled.
	//
	// 	- **false**: enabled.
	//
	// >  When this parameter is set to true, the application instances are never reduced. This prevents risks to your business in peak hours. By default, this parameter is set to false.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The cooldown time of the scale-in. Valid values: 0 to 3600. Unit: seconds. Default value: 0.
	//
	// example:
	//
	// 300
	StabilizationWindowSeconds *int64 `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
	// The step size for the scale-in. The maximum number of instances that can be reduced within a specific period of time.
	//
	// example:
	//
	// 100
	Step *int64 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules) GetStabilizationWindowSeconds() *int64 {
	return s.StabilizationWindowSeconds
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules) GetStep() *int64 {
	return s.Step
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules) SetDisabled(v bool) *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules {
	s.Disabled = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules) SetStabilizationWindowSeconds(v int64) *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules {
	s.StabilizationWindowSeconds = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules) SetStep(v int64) *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules {
	s.Step = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules struct {
	// Indicates whether the application scale-in is disabled. Valid values:
	//
	// 	- **true**: The application scale-in is disabled.
	//
	// 	- **false**: The application scale-in is enabled.
	//
	// >  When this parameter is set to true, the application instances are never reduced. This prevents risks to your business in peak hours. By default, this parameter is set to false.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The cooldown time of the scale-out. Valid values: 0 to 3600. Unit: seconds. Default value: 0.
	//
	// example:
	//
	// 300
	StabilizationWindowSeconds *int64 `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
	// The step size for the scale-out. The maximum number of instances that can be added within a specific period of time.
	//
	// example:
	//
	// 100
	Step *int64 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules) GetStabilizationWindowSeconds() *int64 {
	return s.StabilizationWindowSeconds
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules) GetStep() *int64 {
	return s.Step
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules) SetDisabled(v bool) *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules {
	s.Disabled = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules) SetStabilizationWindowSeconds(v int64) *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules {
	s.StabilizationWindowSeconds = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules) SetStep(v int64) *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules {
	s.Step = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataMetricScaleUpRules) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationScalingRuleResponseBodyDataTimer struct {
	// The start date of the validity period of the scheduled auto scaling policy. Valid values:
	//
	// 	- If both the **BeginDate*	- and **EndDate*	- parameters are set to **null**, the auto scaling policy can always be triggered. The default value for these parameters is null.
	//
	// 	- If the two parameters are set to specific dates, the scheduled auto scaling policy can be triggered during the period between the two dates. For example, if **BeginDate*	- is **2021-03-25*	- and **EndDate*	- is **2021-04-25**, the auto scaling policy is valid for one month.
	//
	// example:
	//
	// 2021-03-25
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// The end date of the validity period of the scheduled auto scaling policy. Valid values:
	//
	// 	- If both the **BeginDate*	- and **EndDate*	- parameters are set to **null**, the auto scaling policy can always be triggered. The default value for these parameters is null.
	//
	// 	- If the two parameters are set to specific dates, the scheduled auto scaling policy can be triggered during the period between the two dates. For example, if **BeginDate*	- is **2021-03-25*	- and **EndDate*	- is **2021-04-25**, the auto scaling policy is valid for one month.
	//
	// example:
	//
	// 2021-04-25
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The days on which the scheduled auto scaling policy takes effect. Valid values:
	//
	// 	- **\\	- \\	- \\***: The scheduled auto scaling policy takes effect at a specified time every day.
	//
	// 	- **\\	- \\	- Fri,Mon**: The scheduled auto scaling policy takes effect at a specified time on one or multiple days of a week. The specified time is in the GMT+8 time zone. Valid values:
	//
	//     	- **Sun**: Sunday
	//
	//     	- **Mon**: Monday
	//
	//     	- **Tue**: Tuesday
	//
	//     	- **Wed**: Wednesday
	//
	//     	- **Thu**: Thursday
	//
	//     	- **Fri**: Friday
	//
	//     	- **Sat**: Saturday
	//
	// 	- **1,2,3,28,31 \\	- \\***: The scheduled auto scaling policy takes effect at a specified time on one or multiple days of a month. Valid values: 1 to 31. If the month does not have a 31st day, the auto scaling policy takes effect on the specified days other than the 31st day.
	//
	// example:
	//
	// 	- 	- *
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The points in time when the auto scaling policy is triggered within one day.
	Schedules []*DescribeApplicationScalingRuleResponseBodyDataTimerSchedules `json:"Schedules,omitempty" xml:"Schedules,omitempty" type:"Repeated"`
}

func (s DescribeApplicationScalingRuleResponseBodyDataTimer) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyDataTimer) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimer) GetBeginDate() *string {
	return s.BeginDate
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimer) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimer) GetPeriod() *string {
	return s.Period
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimer) GetSchedules() []*DescribeApplicationScalingRuleResponseBodyDataTimerSchedules {
	return s.Schedules
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimer) SetBeginDate(v string) *DescribeApplicationScalingRuleResponseBodyDataTimer {
	s.BeginDate = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimer) SetEndDate(v string) *DescribeApplicationScalingRuleResponseBodyDataTimer {
	s.EndDate = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimer) SetPeriod(v string) *DescribeApplicationScalingRuleResponseBodyDataTimer {
	s.Period = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimer) SetSchedules(v []*DescribeApplicationScalingRuleResponseBodyDataTimerSchedules) *DescribeApplicationScalingRuleResponseBodyDataTimer {
	s.Schedules = v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimer) Validate() error {
	if s.Schedules != nil {
		for _, item := range s.Schedules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationScalingRuleResponseBodyDataTimerSchedules struct {
	// The point in time. Format: **Hour:Minute**.
	//
	// example:
	//
	// 08:00
	AtTime *string `json:"AtTime,omitempty" xml:"AtTime,omitempty"`
	// example:
	//
	// 10
	MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// example:
	//
	// 1
	MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	// The expected number of instances.
	//
	// example:
	//
	// 2
	TargetReplicas *int32 `json:"TargetReplicas,omitempty" xml:"TargetReplicas,omitempty"`
}

func (s DescribeApplicationScalingRuleResponseBodyDataTimerSchedules) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponseBodyDataTimerSchedules) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimerSchedules) GetAtTime() *string {
	return s.AtTime
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimerSchedules) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimerSchedules) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimerSchedules) GetTargetReplicas() *int32 {
	return s.TargetReplicas
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimerSchedules) SetAtTime(v string) *DescribeApplicationScalingRuleResponseBodyDataTimerSchedules {
	s.AtTime = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimerSchedules) SetMaxReplicas(v int32) *DescribeApplicationScalingRuleResponseBodyDataTimerSchedules {
	s.MaxReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimerSchedules) SetMinReplicas(v int32) *DescribeApplicationScalingRuleResponseBodyDataTimerSchedules {
	s.MinReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimerSchedules) SetTargetReplicas(v int32) *DescribeApplicationScalingRuleResponseBodyDataTimerSchedules {
	s.TargetReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponseBodyDataTimerSchedules) Validate() error {
	return dara.Validate(s)
}
