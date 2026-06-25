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
	// The HTTP status code or a POP error code. Valid values:
	//
	// - **2xx**: The operation is successful.
	//
	// - **3xx**: A redirection is required.
	//
	// - **4xx**: A request error occurred.
	//
	// - **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeApplicationScalingRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// - If the request is successful, the **ErrorCode*	- field is not returned.
	//
	// - If the request fails, the **ErrorCode*	- field is returned. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information. Valid values:
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If the request fails, a specific error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 73404D3D-EE4F-4CB2-B197-5C46F6A1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the application instance was successfully restarted.
	//
	// - **true**: The restart succeeded.
	//
	// - **false**: The restart failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID. Use this ID to query the details of a request.
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
	// The application ID.
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
	// The metric-based scaling policy.
	Metric *DescribeApplicationScalingRuleResponseBodyDataMetric `json:"Metric,omitempty" xml:"Metric,omitempty" type:"Struct"`
	// The percentage of the minimum number of ready instances. Valid values:
	//
	// - **-1**: an initial value, which indicates that a percentage is not used.
	//
	// - **0 to 100**: a percentage that is rounded up. For example, if you set this parameter to 50% and the current number of instances is 5, the minimum number of ready instances is 3.
	//
	// > If you specify both MinReadyInstances and MinReadyInstanceRatio, and the value of **MinReadyInstanceRatio*	- is not **-1**, the value of **MinReadyInstanceRatio*	- prevails. For example, if **MinReadyInstances*	- is set to **5*	- and **MinReadyInstanceRatio*	- is set to **50**, the value **50*	- is used to calculate the minimum number of ready instances.
	//
	// example:
	//
	// -1
	MinReadyInstanceRatio *int32 `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	// The minimum number of ready instances. Valid values:
	//
	// - If you set this parameter to **0**, the application is interrupted during an upgrade.
	//
	// - If you set this parameter to -1, the system uses a recommended value for the minimum number of ready instances. The value is 25% of the current number of instances. For example, if the current number of instances is 5, the minimum number of ready instances is 2 after 5 × 25% = 1.25 is rounded up.
	//
	// > Set the minimum number of ready instances to a value greater than or equal to 1 for each rolling deployment to ensure business continuity.
	//
	// example:
	//
	// 1
	MinReadyInstances *int32 `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
	// Indicates whether the auto scaling policy is enabled. Valid values:
	//
	// - **true**: The policy is enabled.
	//
	// - **false**: The policy is disabled.
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
	// - **timing**: scheduled scaling.
	//
	// - **metric**: metric-based scaling.
	//
	// - **mix**: hybrid scaling.
	//
	// example:
	//
	// timing
	ScaleRuleType *string `json:"ScaleRuleType,omitempty" xml:"ScaleRuleType,omitempty"`
	// The scheduled scaling policy.
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
	// The list of metric-based scaling policies.
	Metrics []*DescribeApplicationScalingRuleResponseBodyDataMetricMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The status of the metric-based scaling policy.
	MetricsStatus *DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatus `json:"MetricsStatus,omitempty" xml:"MetricsStatus,omitempty" type:"Struct"`
	// The minimum number of instances.
	//
	// example:
	//
	// 1
	MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	// The scale-in rules.
	ScaleDownRules *DescribeApplicationScalingRuleResponseBodyDataMetricScaleDownRules `json:"ScaleDownRules,omitempty" xml:"ScaleDownRules,omitempty" type:"Struct"`
	// The scale-out rules.
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
	// The target value of the metric.
	//
	// - The target CPU utilization. Unit: percent.
	//
	// - The target memory usage. Unit: percent.
	//
	// - The number of queries per second (QPS).
	//
	// - The response time. Unit: milliseconds.
	//
	// - The average number of active TCP connections per second.
	//
	// - The QPS of a public-facing SLB instance.
	//
	// - The response time of a public-facing SLB instance. Unit: milliseconds.
	//
	// - The QPS of a private SLB instance.
	//
	// - The response time of a private SLB instance. Unit: milliseconds.
	//
	// example:
	//
	// 20
	MetricTargetAverageUtilization *int32 `json:"MetricTargetAverageUtilization,omitempty" xml:"MetricTargetAverageUtilization,omitempty"`
	// The metric that is used to trigger the auto scaling policy. Valid values:
	//
	// - **CPU**: CPU utilization.
	//
	// - **MEMORY**: memory usage.
	//
	// - **QPS**: the average QPS of a single instance of a Java application in one minute.
	//
	// - **RT**: the average response time (RT) of all service interfaces of a Java application in one minute.
	//
	// - **tcpActiveConn**: the average number of active TCP connections of a single instance in 30 seconds.
	//
	// - **SLB_QPS**: the average QPS of a single instance for a public-facing SLB instance in 15 seconds.
	//
	// - **SLB_RT**: the average RT of a public-facing SLB instance in 15 seconds.
	//
	// - **INTRANET_SLB_QPS**: the average QPS of a single instance for a private SLB instance in 15 seconds.
	//
	// - **INTRANET_SLB_RT**: the average RT of a private SLB instance in 15 seconds.
	//
	// example:
	//
	// CPU
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The ID of the SLB instance.
	//
	// example:
	//
	// lb-xxx
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The SLB access log Logstore.
	//
	// example:
	//
	// test
	SlbLogstore *string `json:"SlbLogstore,omitempty" xml:"SlbLogstore,omitempty"`
	// The SLB access log Project.
	//
	// example:
	//
	// test
	SlbProject *string `json:"SlbProject,omitempty" xml:"SlbProject,omitempty"`
	// The port of the SLB instance.
	//
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
	// The data of the current metric-based scaling.
	CurrentMetrics []*DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusCurrentMetrics `json:"CurrentMetrics,omitempty" xml:"CurrentMetrics,omitempty" type:"Repeated"`
	// The current number of instances.
	//
	// example:
	//
	// 2
	CurrentReplicas *int64 `json:"CurrentReplicas,omitempty" xml:"CurrentReplicas,omitempty"`
	// The target number of instances.
	//
	// example:
	//
	// 2
	DesiredReplicas *int64 `json:"DesiredReplicas,omitempty" xml:"DesiredReplicas,omitempty"`
	// The time of the last scaling activity.
	//
	// example:
	//
	// 2022-01-11T08:14:32Z
	LastScaleTime *string `json:"LastScaleTime,omitempty" xml:"LastScaleTime,omitempty"`
	// The list of metrics for the next scaling activity.
	NextScaleMetrics []*DescribeApplicationScalingRuleResponseBodyDataMetricMetricsStatusNextScaleMetrics `json:"NextScaleMetrics,omitempty" xml:"NextScaleMetrics,omitempty" type:"Repeated"`
	// The period of the next metric-based scaling.
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
	// The current value.
	//
	// example:
	//
	// 0
	CurrentValue *int64 `json:"CurrentValue,omitempty" xml:"CurrentValue,omitempty"`
	// The name of the metric.
	//
	// - **cpu**: CPU utilization.
	//
	// - **memory**: memory usage.
	//
	// - **arms_incall_qps**: the average QPS of a single instance of a Java application in one minute.
	//
	// - **arms_incall_rt**: the average RT of all service interfaces of a Java application in one minute.
	//
	// - **tcpActiveConn**: the number of active TCP connections.
	//
	// - **slb_incall_qps**: the QPS of a public-facing SLB instance.
	//
	// - **slb_incall_rt**: the RT of a public-facing SLB instance.
	//
	// - **intranet_slb_incall_qps**: the QPS of a private SLB instance.
	//
	// - **intranet_slb_incall_rt**: the RT of a private SLB instance.
	//
	// example:
	//
	// cpu
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the metric. This parameter is associated with the monitoring metric.
	//
	// - **Resource**: the metric value of **cpu*	- or **memory**.
	//
	// - **Pods**: the metric value of **tcpActiveConn**.
	//
	// - **External**: the metric value of **arms*	- or **slb**.
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
	// - **cpu**: CPU utilization.
	//
	// - **memory**: memory usage.
	//
	// - **arms_incall_qps**: the average QPS of a single instance of a Java application in one minute.
	//
	// - **arms_incall_rt**: the average RT of all service interfaces of a Java application in one minute.
	//
	// - **tcpActiveConn**: the number of active TCP connections.
	//
	// - **slb_incall_qps**: the QPS of a public-facing SLB instance.
	//
	// - **slb_incall_rt**: the RT of a public-facing SLB instance.
	//
	// - **intranet_slb_incall_qps**: the QPS of a private SLB instance.
	//
	// - **intranet_slb_incall_rt**: the RT of a private SLB instance.
	//
	// example:
	//
	// cpu
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The metric threshold for the next scale-in. The value is a percentage.
	//
	// example:
	//
	// 10
	NextScaleInAverageUtilization *int32 `json:"NextScaleInAverageUtilization,omitempty" xml:"NextScaleInAverageUtilization,omitempty"`
	// The metric threshold for the next scale-out. The value is a percentage.
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
	// Indicates whether scale-in is disabled. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// > If you enable this feature, the application is never scaled in. This prevents business risks that are caused by scale-ins during peak hours. By default, this feature is disabled.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The cooldown period for scale-ins. The value can be an integer from 0 to 3600. Unit: seconds. Default value: 0.
	//
	// example:
	//
	// 300
	StabilizationWindowSeconds *int64 `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
	// The scaling step size for scale-ins. The maximum number of instances that can be removed at a time.
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
	// Indicates whether scale-in is disabled. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// > If you enable this feature, the application is never scaled in. This prevents business risks that are caused by scale-ins during peak hours. By default, this feature is disabled.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The cooldown period for scale-outs. The value can be an integer from 0 to 3600. Unit: seconds. Default value: 0.
	//
	// example:
	//
	// 300
	StabilizationWindowSeconds *int64 `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
	// The scaling step size for scale-outs. The maximum number of instances that can be added at a time.
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
	// The start date of a short-term scheduled scaling policy. The following list describes the valid values:
	//
	// - If you leave both **BeginDate*	- and **EndDate*	- empty, the policy is a long-term policy. This is the default value.
	//
	// - If you specify a date, for example, you set **BeginDate*	- to **2021-03-25*	- and **EndDate*	- to **2021-04-25**, the policy is effective for one month.
	//
	// example:
	//
	// 2021-03-25
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// The end date of a short-term scheduled scaling policy. The following list describes the valid values:
	//
	// - If you leave both **BeginDate*	- and **EndDate*	- empty, the policy is a long-term policy. This is the default value.
	//
	// - If you specify a date, for example, you set **BeginDate*	- to **2021-03-25*	- and **EndDate*	- to **2021-04-25**, the policy is effective for one month.
	//
	// example:
	//
	// 2021-04-25
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The period in which the scheduled scaling policy is executed. Valid values:
	//
	// - **\\	- \\	- \\***: The policy is executed at a specified time every day.
	//
	// - **\\	- \\	- Fri,Mon**: The policy is executed at a specified time on one or more days of a week. You can select multiple days. The time is in GMT+8. Valid values:
	//
	//   - **Sun**: Sunday
	//
	//   - **Mon**: Monday
	//
	//   - **Tue**: Tuesday
	//
	//   - **Wed**: Wednesday
	//
	//   - **Thu**: Thursday
	//
	//   - **Fri**: Friday
	//
	//   - **Sat**: Saturday
	//
	// - **1,2,3,28,31 \\	- \\***: The policy is executed at a specified time on one or more days of a month. You can select multiple days. The value can be an integer from 1 to 31. If a month does not have a 31st day, the policy is not executed on that day.
	//
	// example:
	//
	// 	- 	- *
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The points in time when the auto scaling policy is triggered within a day.
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
	// The point in time. Format: **HH:mm**.
	//
	// example:
	//
	// 08:00
	AtTime *string `json:"AtTime,omitempty" xml:"AtTime,omitempty"`
	// The maximum number of instances.
	//
	// example:
	//
	// 10
	MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// The minimum number of instances.
	//
	// example:
	//
	// 1
	MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	// The target number of instances.
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
