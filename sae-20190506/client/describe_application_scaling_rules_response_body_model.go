// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationScalingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeApplicationScalingRulesResponseBody
	GetCode() *string
	SetData(v *DescribeApplicationScalingRulesResponseBodyData) *DescribeApplicationScalingRulesResponseBody
	GetData() *DescribeApplicationScalingRulesResponseBodyData
	SetErrorCode(v string) *DescribeApplicationScalingRulesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeApplicationScalingRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeApplicationScalingRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeApplicationScalingRulesResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeApplicationScalingRulesResponseBody
	GetTraceId() *string
}

type DescribeApplicationScalingRulesResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// - **2xx**: The request was successful.
	//
	// - **3xx**: The request was redirected.
	//
	// - **4xx**: The request was invalid.
	//
	// - **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeApplicationScalingRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. This parameter is returned only when the request fails.
	//
	// -
	//
	// - For more information, see the **Error codes*	- section of this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The response message. Valid values:
	//
	// - Returns **success*	- if the request is successful.
	//
	// - Returns an error message if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to query the details of a request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeApplicationScalingRulesResponseBody) GetData() *DescribeApplicationScalingRulesResponseBodyData {
	return s.Data
}

func (s *DescribeApplicationScalingRulesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeApplicationScalingRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeApplicationScalingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationScalingRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeApplicationScalingRulesResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeApplicationScalingRulesResponseBody) SetCode(v string) *DescribeApplicationScalingRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBody) SetData(v *DescribeApplicationScalingRulesResponseBodyData) *DescribeApplicationScalingRulesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBody) SetErrorCode(v string) *DescribeApplicationScalingRulesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBody) SetMessage(v string) *DescribeApplicationScalingRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBody) SetRequestId(v string) *DescribeApplicationScalingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBody) SetSuccess(v bool) *DescribeApplicationScalingRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBody) SetTraceId(v string) *DescribeApplicationScalingRulesResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationScalingRulesResponseBodyData struct {
	// A list of auto scaling policies for the application.
	ApplicationScalingRules []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules `json:"ApplicationScalingRules,omitempty" xml:"ApplicationScalingRules,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of auto scaling policies for the application.
	//
	// example:
	//
	// 3
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyData) GetApplicationScalingRules() []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	return s.ApplicationScalingRules
}

func (s *DescribeApplicationScalingRulesResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeApplicationScalingRulesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApplicationScalingRulesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *DescribeApplicationScalingRulesResponseBodyData) SetApplicationScalingRules(v []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) *DescribeApplicationScalingRulesResponseBodyData {
	s.ApplicationScalingRules = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyData) SetCurrentPage(v int32) *DescribeApplicationScalingRulesResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyData) SetPageSize(v int32) *DescribeApplicationScalingRulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyData) SetTotalSize(v int32) *DescribeApplicationScalingRulesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyData) Validate() error {
	if s.ApplicationScalingRules != nil {
		for _, item := range s.ApplicationScalingRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules struct {
	// The application ID.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The timestamp of the policy\\"s creation, in milliseconds.
	//
	// example:
	//
	// 1616642248938
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The timestamp of when the policy was last disabled.
	//
	// example:
	//
	// 1641882854484
	LastDisableTime *int64 `json:"LastDisableTime,omitempty" xml:"LastDisableTime,omitempty"`
	// The metric-based scaling policy.
	Metric *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric `json:"Metric,omitempty" xml:"Metric,omitempty" type:"Struct"`
	// The minimum number of available instances, specified as a percentage. Valid values:
	//
	// - **-1**: Indicates that this parameter is not used.
	//
	// - **0 to 100**: a percentage that is rounded up to the nearest integer. For example, if you set this parameter to 50% and you have five instances, the minimum number of available instances is 3.
	//
	// > If you specify both **MinReadyInstances*	- and **MinReadyInstanceRatio**, the value of **MinReadyInstanceRatio*	- takes precedence, unless it is set to **-1**.
	//
	// example:
	//
	// -1
	MinReadyInstanceRatio *int32 `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	// The minimum number of available instances. Valid values:
	//
	// - If you set this parameter to **0**, the application may be interrupted during an upgrade.
	//
	// - If you set this parameter to **-1**, a recommended value is used, which is 25% of the current number of instances, rounded up to the nearest integer. For example, if an application has five instances, the minimum number of available instances is 2 (5 \\	- 25% = 1.25, rounded up).
	//
	// > To ensure business continuity during a rolling deployment, we recommend that you set this parameter to a value greater than or equal to 1.
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
	// - **timing**: A scheduled scaling policy.
	//
	// - **metric**: A metric-based scaling policy.
	//
	// - **mix**: A hybrid scaling policy.
	//
	// example:
	//
	// timing
	ScaleRuleType *string `json:"ScaleRuleType,omitempty" xml:"ScaleRuleType,omitempty"`
	// The scheduled scaling policy.
	Timer *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer `json:"Timer,omitempty" xml:"Timer,omitempty" type:"Struct"`
	// The timestamp of the last policy update, in milliseconds.
	//
	// example:
	//
	// 1616642248938
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) GetAppId() *string {
	return s.AppId
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) GetLastDisableTime() *int64 {
	return s.LastDisableTime
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) GetMetric() *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric {
	return s.Metric
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) GetMinReadyInstanceRatio() *int32 {
	return s.MinReadyInstanceRatio
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) GetMinReadyInstances() *int32 {
	return s.MinReadyInstances
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) GetScaleRuleEnabled() *bool {
	return s.ScaleRuleEnabled
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) GetScaleRuleName() *string {
	return s.ScaleRuleName
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) GetScaleRuleType() *string {
	return s.ScaleRuleType
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) GetTimer() *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer {
	return s.Timer
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetAppId(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetCreateTime(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.CreateTime = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetLastDisableTime(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.LastDisableTime = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetMetric(v *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.Metric = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetMinReadyInstanceRatio(v int32) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.MinReadyInstanceRatio = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetMinReadyInstances(v int32) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.MinReadyInstances = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetScaleRuleEnabled(v bool) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.ScaleRuleEnabled = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetScaleRuleName(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.ScaleRuleName = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetScaleRuleType(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.ScaleRuleType = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetTimer(v *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.Timer = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) SetUpdateTime(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules {
	s.UpdateTime = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules) Validate() error {
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

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric struct {
	// The maximum number of instances.
	//
	// example:
	//
	// 3
	MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// The source of the metrics.
	MetricSource *string `json:"MetricSource,omitempty" xml:"MetricSource,omitempty"`
	// The metric-based conditions that trigger scaling.
	Metrics []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The status of the metric-based scaling policy.
	MetricsStatus *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus `json:"MetricsStatus,omitempty" xml:"MetricsStatus,omitempty" type:"Struct"`
	// The minimum number of instances.
	//
	// example:
	//
	// 1
	MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	// The Prometheus metrics.
	PrometheusMetrics []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricPrometheusMetrics `json:"PrometheusMetrics,omitempty" xml:"PrometheusMetrics,omitempty" type:"Repeated"`
	// The Prometheus token.
	PrometheusToken *string `json:"PrometheusToken,omitempty" xml:"PrometheusToken,omitempty"`
	// The endpoint of the Prometheus service.
	PrometheusUrl *string `json:"PrometheusUrl,omitempty" xml:"PrometheusUrl,omitempty"`
	// Configuration for scale-in events.
	ScaleDownRules *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules `json:"ScaleDownRules,omitempty" xml:"ScaleDownRules,omitempty" type:"Struct"`
	// Configuration for scale-out events.
	ScaleUpRules *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules `json:"ScaleUpRules,omitempty" xml:"ScaleUpRules,omitempty" type:"Struct"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) GetMetricSource() *string {
	return s.MetricSource
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) GetMetrics() []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics {
	return s.Metrics
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) GetMetricsStatus() *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus {
	return s.MetricsStatus
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) GetPrometheusMetrics() []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricPrometheusMetrics {
	return s.PrometheusMetrics
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) GetPrometheusToken() *string {
	return s.PrometheusToken
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) GetPrometheusUrl() *string {
	return s.PrometheusUrl
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) GetScaleDownRules() *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules {
	return s.ScaleDownRules
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) GetScaleUpRules() *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules {
	return s.ScaleUpRules
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) SetMaxReplicas(v int32) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric {
	s.MaxReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) SetMetricSource(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric {
	s.MetricSource = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) SetMetrics(v []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric {
	s.Metrics = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) SetMetricsStatus(v *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric {
	s.MetricsStatus = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) SetMinReplicas(v int32) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric {
	s.MinReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) SetPrometheusMetrics(v []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricPrometheusMetrics) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric {
	s.PrometheusMetrics = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) SetPrometheusToken(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric {
	s.PrometheusToken = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) SetPrometheusUrl(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric {
	s.PrometheusUrl = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) SetScaleDownRules(v *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric {
	s.ScaleDownRules = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) SetScaleUpRules(v *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric {
	s.ScaleUpRules = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric) Validate() error {
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
	if s.PrometheusMetrics != nil {
		for _, item := range s.PrometheusMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics struct {
	// The target value for the metric. The unit varies based on the value of `MetricType`.
	//
	// - Target CPU usage, in percent.
	//
	// - Target memory usage, in percent.
	//
	// - Target QPS, in queries per second.
	//
	// - Target response time, in milliseconds.
	//
	// - Target number of active TCP connections.
	//
	// - Target QPS of a public-facing SLB instance, in queries per second.
	//
	// - Target response time of a public-facing SLB instance, in milliseconds.
	//
	// - Target QPS of a private SLB instance, in queries per second.
	//
	// - Target response time of a private SLB instance, in milliseconds.
	//
	// example:
	//
	// 20
	MetricTargetAverageUtilization *int32 `json:"MetricTargetAverageUtilization,omitempty" xml:"MetricTargetAverageUtilization,omitempty"`
	// The metric used to trigger the auto scaling policy. Valid values:
	//
	// - **CPU**: CPU usage.
	//
	// - **MEMORY**: memory usage.
	//
	// - **QPS**: Average queries per second (QPS) per instance over a 1-minute period. This metric applies to Java applications only.
	//
	// - **RT**: Average response time of all service interfaces in a Java application over a 1-minute period.
	//
	// - **tcpActiveConn**: Average number of active TCP connections per instance over a 30-second period.
	//
	// - **SLB_QPS**: Average QPS per instance for a public-facing SLB instance over a 15-second period.
	//
	// - **SLB_RT**: Average response time of a public-facing SLB instance over a 15-second period.
	//
	// - **INTRANET_SLB_QPS**: Average QPS per instance for a private SLB instance over a 15-second period.
	//
	// - **INTRANET_SLB_RT**: Average response time of a private SLB instance over a 15-second period.
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
	// The Logstore in Log Service that stores SLB access logs.
	//
	// example:
	//
	// test
	SlbLogstore *string `json:"SlbLogstore,omitempty" xml:"SlbLogstore,omitempty"`
	// The project in Log Service that stores SLB access logs.
	//
	// example:
	//
	// test
	SlbProject *string `json:"SlbProject,omitempty" xml:"SlbProject,omitempty"`
	// The monitored port of the SLB instance.
	//
	// example:
	//
	// 80
	Vport *string `json:"Vport,omitempty" xml:"Vport,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) GetMetricTargetAverageUtilization() *int32 {
	return s.MetricTargetAverageUtilization
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) GetSlbId() *string {
	return s.SlbId
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) GetSlbLogstore() *string {
	return s.SlbLogstore
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) GetSlbProject() *string {
	return s.SlbProject
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) GetVport() *string {
	return s.Vport
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) SetMetricTargetAverageUtilization(v int32) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics {
	s.MetricTargetAverageUtilization = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) SetMetricType(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics {
	s.MetricType = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) SetSlbId(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics {
	s.SlbId = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) SetSlbLogstore(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics {
	s.SlbLogstore = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) SetSlbProject(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics {
	s.SlbProject = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) SetVport(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics {
	s.Vport = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus struct {
	// A list of the current metrics for scaling.
	CurrentMetrics []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics `json:"CurrentMetrics,omitempty" xml:"CurrentMetrics,omitempty" type:"Repeated"`
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
	// The maximum number of instances.
	//
	// example:
	//
	// 3
	MaxReplicas *int64 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// The minimum number of instances.
	//
	// example:
	//
	// 1
	MinReplicas *int64 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	// A list of metrics for the next scaling activity.
	NextScaleMetrics []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics `json:"NextScaleMetrics,omitempty" xml:"NextScaleMetrics,omitempty" type:"Repeated"`
	// The next period for metric-based scaling.
	//
	// example:
	//
	// 3
	NextScaleTimePeriod *int32 `json:"NextScaleTimePeriod,omitempty" xml:"NextScaleTimePeriod,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) GetCurrentMetrics() []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics {
	return s.CurrentMetrics
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) GetCurrentReplicas() *int64 {
	return s.CurrentReplicas
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) GetDesiredReplicas() *int64 {
	return s.DesiredReplicas
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) GetLastScaleTime() *string {
	return s.LastScaleTime
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) GetMaxReplicas() *int64 {
	return s.MaxReplicas
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) GetMinReplicas() *int64 {
	return s.MinReplicas
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) GetNextScaleMetrics() []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics {
	return s.NextScaleMetrics
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) GetNextScaleTimePeriod() *int32 {
	return s.NextScaleTimePeriod
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) SetCurrentMetrics(v []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus {
	s.CurrentMetrics = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) SetCurrentReplicas(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus {
	s.CurrentReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) SetDesiredReplicas(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus {
	s.DesiredReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) SetLastScaleTime(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus {
	s.LastScaleTime = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) SetMaxReplicas(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus {
	s.MaxReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) SetMinReplicas(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus {
	s.MinReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) SetNextScaleMetrics(v []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus {
	s.NextScaleMetrics = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) SetNextScaleTimePeriod(v int32) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus {
	s.NextScaleTimePeriod = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus) Validate() error {
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

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics struct {
	// The current value.
	//
	// example:
	//
	// 0
	CurrentValue *int64 `json:"CurrentValue,omitempty" xml:"CurrentValue,omitempty"`
	// The name of the trigger condition.
	//
	// - **cpu**: CPU usage.
	//
	// - **memory**: memory usage.
	//
	// - **arms_incall_qps_v2**: QPS of a Java application.
	//
	// - **arms_incall_rt**: Response time of a Java application.
	//
	// - **tcpActiveConn**: The number of active TCP connections.
	//
	// - **slb_incall_qps**: QPS of a public-facing SLB instance.
	//
	// - **slb_incall_rt**: Response time of a public-facing SLB instance.
	//
	// - **intranet_slb_incall_qps**: QPS of a private SLB instance.
	//
	// - **intranet_slb_incall_rt**: Response time of a private SLB instance.
	//
	// example:
	//
	// cpu
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The data type. This parameter is associated with the specified metric.
	//
	// - **Resource**: The metric value for **cpu*	- or **memory**.
	//
	// - **Pods**: The metric value for **tcpActiveConn**.
	//
	// - **External**: The metric value for **arms*	- or **slb**.
	//
	// example:
	//
	// Resource
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics) GetCurrentValue() *int64 {
	return s.CurrentValue
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics) GetType() *string {
	return s.Type
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics) SetCurrentValue(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics {
	s.CurrentValue = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics) SetName(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics {
	s.Name = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics) SetType(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics {
	s.Type = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics struct {
	// The name of the trigger condition.
	//
	// - **cpu**: CPU usage.
	//
	// - **memory**: memory usage.
	//
	// - **arms_incall_qps_v2**: QPS of a Java application.
	//
	// - **arms_incall_rt**: Response time of a Java application.
	//
	// - **tcpActiveConn**: The number of active TCP connections.
	//
	// - **slb_incall_qps**: QPS of a public-facing SLB instance.
	//
	// - **slb_incall_rt**: Response time of a public-facing SLB instance.
	//
	// - **intranet_slb_incall_qps**: QPS of a private SLB instance.
	//
	// - **intranet_slb_incall_rt**: Response time of a private SLB instance.
	//
	// example:
	//
	// cpu
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The metric value that triggers the next scale-in event. The value is a percentage.
	//
	// example:
	//
	// 10
	NextScaleInAverageUtilization *int32 `json:"NextScaleInAverageUtilization,omitempty" xml:"NextScaleInAverageUtilization,omitempty"`
	// The metric value that triggers the next scale-out event. The value is a percentage.
	//
	// example:
	//
	// 21
	NextScaleOutAverageUtilization *int32 `json:"NextScaleOutAverageUtilization,omitempty" xml:"NextScaleOutAverageUtilization,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics) GetNextScaleInAverageUtilization() *int32 {
	return s.NextScaleInAverageUtilization
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics) GetNextScaleOutAverageUtilization() *int32 {
	return s.NextScaleOutAverageUtilization
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics) SetName(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics {
	s.Name = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics) SetNextScaleInAverageUtilization(v int32) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics {
	s.NextScaleInAverageUtilization = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics) SetNextScaleOutAverageUtilization(v int32) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics {
	s.NextScaleOutAverageUtilization = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricPrometheusMetrics struct {
	// The Prometheus query.
	PrometheusQuery *string `json:"PrometheusQuery,omitempty" xml:"PrometheusQuery,omitempty"`
	// The target value for the Prometheus query that triggers a scaling event.
	TargetMetricValue *string `json:"TargetMetricValue,omitempty" xml:"TargetMetricValue,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricPrometheusMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricPrometheusMetrics) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricPrometheusMetrics) GetPrometheusQuery() *string {
	return s.PrometheusQuery
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricPrometheusMetrics) GetTargetMetricValue() *string {
	return s.TargetMetricValue
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricPrometheusMetrics) SetPrometheusQuery(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricPrometheusMetrics {
	s.PrometheusQuery = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricPrometheusMetrics) SetTargetMetricValue(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricPrometheusMetrics {
	s.TargetMetricValue = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricPrometheusMetrics) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules struct {
	// Specifies whether to disable scale-in. Valid values:
	//
	// - **true**: Disables scale-in.
	//
	// - **false**: Enables scale-in.
	//
	// > Setting this to `true` prevents the application from scaling in, which can be useful to avoid service disruptions from unexpected capacity reduction during peak hours. Default: `false`.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The cooldown time for scale-in events, in seconds. During this period, no further scaling events are triggered. The value must be an integer from 0 to 3,600. The default value is 0.
	//
	// example:
	//
	// 300
	StabilizationWindowSeconds *int64 `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
	// The number of instances to remove in a single scale-in event.
	//
	// example:
	//
	// 100
	Step *int64 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules) GetStabilizationWindowSeconds() *int64 {
	return s.StabilizationWindowSeconds
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules) GetStep() *int64 {
	return s.Step
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules) SetDisabled(v bool) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules {
	s.Disabled = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules) SetStabilizationWindowSeconds(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules {
	s.StabilizationWindowSeconds = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules) SetStep(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules {
	s.Step = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules struct {
	// Specifies whether to disable scale-out. Valid values:
	//
	// - **true**: Disables scale-out.
	//
	// - **false**: Enables scale-out.
	//
	// > If this parameter is set to `true`, application instances are never scaled out. This can be useful to freeze the application capacity during specific events. By default, this parameter is set to `false`.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The cooldown time for scale-out events, in seconds. During this period, no further scaling events are triggered. The value must be an integer from 0 to 3,600. The default value is 0.
	//
	// example:
	//
	// 300
	StabilizationWindowSeconds *int64 `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
	// The number of instances to add in a single scale-out event.
	//
	// example:
	//
	// 100
	Step *int64 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules) GetStabilizationWindowSeconds() *int64 {
	return s.StabilizationWindowSeconds
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules) GetStep() *int64 {
	return s.Step
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules) SetDisabled(v bool) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules {
	s.Disabled = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules) SetStabilizationWindowSeconds(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules {
	s.StabilizationWindowSeconds = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules) SetStep(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules {
	s.Step = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleUpRules) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer struct {
	// The start date of the short-term scheduled scaling policy. The following rules apply:
	//
	// - If **BeginDate*	- and **EndDate*	- are not specified, the policy is long-term by default.
	//
	// - If you specify a `BeginDate` and an `EndDate`, the policy is short-term and applies only within that date range.
	//
	// example:
	//
	// 2021-03-25
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// The end date of the short-term scheduled scaling policy. The following rules apply:
	//
	// - If **BeginDate*	- and **EndDate*	- are not specified, the policy is long-term by default.
	//
	// - If you specify a `BeginDate` and an `EndDate`, the policy is short-term and applies only within that date range.
	//
	// example:
	//
	// 2021-04-25
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The days on which the scheduled scaling policy runs. Valid values:
	//
	// - **\\	- \\	- \\***: The policy is executed at a specified time every day.
	//
	// - **\\	- \\	- Fri,Mon**: Executes the policy on specified days of the week. The time zone is GMT+8. Valid days are listed below:
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
	// - **1,2,3,28,31 \\	- \\***: Executes the policy on specified days of the month (1-31). If a specified day does not exist in a given month (e.g., the 31st), the policy does not run on that day.
	//
	// example:
	//
	// 	- 	- *
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The daily trigger schedule for the policy.
	Schedules []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules `json:"Schedules,omitempty" xml:"Schedules,omitempty" type:"Repeated"`
	// The time zone.
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) GetBeginDate() *string {
	return s.BeginDate
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) GetPeriod() *string {
	return s.Period
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) GetSchedules() []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules {
	return s.Schedules
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) GetTimeZone() *string {
	return s.TimeZone
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) SetBeginDate(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer {
	s.BeginDate = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) SetEndDate(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer {
	s.EndDate = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) SetPeriod(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer {
	s.Period = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) SetSchedules(v []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer {
	s.Schedules = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) SetTimeZone(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer {
	s.TimeZone = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer) Validate() error {
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

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules struct {
	// The trigger time in `HH:mm` format.
	//
	// example:
	//
	// 08:00
	AtTime *string `json:"AtTime,omitempty" xml:"AtTime,omitempty"`
	// The maximum number of instances.
	//
	// example:
	//
	// 50
	MaxReplicas *int64 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// The minimum number of instances.
	//
	// example:
	//
	// 1
	MinReplicas *int64 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	// The target number of instances.
	//
	// example:
	//
	// 3
	TargetReplicas *int32 `json:"TargetReplicas,omitempty" xml:"TargetReplicas,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules) GetAtTime() *string {
	return s.AtTime
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules) GetMaxReplicas() *int64 {
	return s.MaxReplicas
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules) GetMinReplicas() *int64 {
	return s.MinReplicas
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules) GetTargetReplicas() *int32 {
	return s.TargetReplicas
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules) SetAtTime(v string) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules {
	s.AtTime = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules) SetMaxReplicas(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules {
	s.MaxReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules) SetMinReplicas(v int64) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules {
	s.MinReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules) SetTargetReplicas(v int32) *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules {
	s.TargetReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules) Validate() error {
	return dara.Validate(s)
}
