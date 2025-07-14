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
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data      *DescribeApplicationScalingRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode *string                                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. The ID is used to query the details of a request.
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
	return dara.Validate(s)
}

type DescribeApplicationScalingRulesResponseBodyData struct {
	// The auto scaling policies of the application.
	ApplicationScalingRules []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules `json:"ApplicationScalingRules,omitempty" xml:"ApplicationScalingRules,omitempty" type:"Repeated"`
	// The number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of auto scaling policies.
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
	return dara.Validate(s)
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRules struct {
	// The ID of the application.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the auto scaling policy was created. Unit: milliseconds.
	//
	// example:
	//
	// 1616642248938
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the auto scaling policy was last disabled.
	//
	// example:
	//
	// 1641882854484
	LastDisableTime *int64 `json:"LastDisableTime,omitempty" xml:"LastDisableTime,omitempty"`
	// The details of the metric-based auto scaling policy.
	Metric *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric `json:"Metric,omitempty" xml:"Metric,omitempty" type:"Struct"`
	// example:
	//
	// -1
	MinReadyInstanceRatio *int32 `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
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
	Timer *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimer `json:"Timer,omitempty" xml:"Timer,omitempty" type:"Struct"`
	// The time when the auto scaling policy was updated. Unit: milliseconds.
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
	return dara.Validate(s)
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetric struct {
	// The maximum number of instances.
	//
	// example:
	//
	// 3
	MaxReplicas  *int32  `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	MetricSource *string `json:"MetricSource,omitempty" xml:"MetricSource,omitempty"`
	// The list of metrics that are used to trigger the auto scaling policy.
	Metrics []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The execution status of the metric-based auto scaling policy.
	MetricsStatus *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatus `json:"MetricsStatus,omitempty" xml:"MetricsStatus,omitempty" type:"Struct"`
	// The minimum number of instances.
	//
	// example:
	//
	// 1
	MinReplicas       *int32                                                                                           `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	PrometheusMetrics []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricPrometheusMetrics `json:"PrometheusMetrics,omitempty" xml:"PrometheusMetrics,omitempty" type:"Repeated"`
	PrometheusToken   *string                                                                                          `json:"PrometheusToken,omitempty" xml:"PrometheusToken,omitempty"`
	PrometheusUrl     *string                                                                                          `json:"PrometheusUrl,omitempty" xml:"PrometheusUrl,omitempty"`
	// Rules that determine the application scale-in.
	ScaleDownRules *DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricScaleDownRules `json:"ScaleDownRules,omitempty" xml:"ScaleDownRules,omitempty" type:"Struct"`
	// Rules that determine the application scale-out.
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
	return dara.Validate(s)
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetrics struct {
	// The limit on the metric.
	//
	// 	- The limit on the CPU utilization. Unit: percentage.
	//
	// 	- The limit on the memory usage. Unit: percentage.
	//
	// 	- The limit on the average number of active TCP connections per second.
	//
	// 	- The limit on the queries per second (QPS) of the Internet-facing Server Load Balancer (SLB) instance.
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
	// 	- **tcpActiveConn**: the average number of active TCP connections per second of an application instance in 30 seconds.
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
	// The metrics that are used to trigger the auto scaling policy this time.
	CurrentMetrics []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics `json:"CurrentMetrics,omitempty" xml:"CurrentMetrics,omitempty" type:"Repeated"`
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
	// The metrics that are used to trigger the auto scaling policy next time.
	NextScaleMetrics []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusNextScaleMetrics `json:"NextScaleMetrics,omitempty" xml:"NextScaleMetrics,omitempty" type:"Repeated"`
	// The duration for which the metric-based auto scaling policy takes effect next time.
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
	return dara.Validate(s)
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesMetricMetricsStatusCurrentMetrics struct {
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
	PrometheusQuery   *string `json:"PrometheusQuery,omitempty" xml:"PrometheusQuery,omitempty"`
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
	// Indicates whether the application scale-in was disabled. Valid values:
	//
	// 	- **true**: The application scale-in was disabled.
	//
	// 	- **false**: The application scale-in was enabled.
	//
	// >  When this parameter is set to true, the application instances will never be reduced. This prevents risks to your business in peak hours. By default, this parameter is set to false.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The cooldown time of the scale-in. Valid values: 0 to 3600. Unit: seconds. The default value is 0.
	//
	// example:
	//
	// 300
	StabilizationWindowSeconds *int64 `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
	// The step size for the scale-in. The maximum number of instances that can be reduced in a unit of time.
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
	// Indicates whether the application scale-in was disabled. Valid values:
	//
	// 	- **true**: The application scale-in was disabled.
	//
	// 	- **false**: The application scale-in was enabled.
	//
	// >  When this parameter is set to true, the application instances will never be reduced. This prevents risks to your business in peak hours. By default, this parameter is set to false.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The cooldown time of the scale-out. Valid values: 0 to 3600. Unit: seconds. The default value is 0.
	//
	// example:
	//
	// 300
	StabilizationWindowSeconds *int64 `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
	// The step size for the scale-out. The maximum number of instances that can be added in a unit of time.
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
	// The start date of the validity period of the scheduled auto scaling policy. Valid values:
	//
	// 	- If both the **BeginDate*	- and **EndDate*	- parameters are set to **null**, the auto scaling policy can always be triggered. The default value for these parameters is null.
	//
	// 	- If the two parameters are set to specific dates, the scheduled auto scaling policy can be triggered during the period between the two dates. For example, if **BeginDate*	- is 2021-03-25 and **EndDate*	- is 2021-04-25, the auto scaling policy is valid for one month.
	//
	// example:
	//
	// 2021-03-25
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// The end date of the validity period of the scheduled auto scaling policy. Valid values:
	//
	// 	- If both the **BeginDate*	- and **EndDate*	- parameters are set to **null**, the auto scaling policy can always be triggered. The default value for these parameters is null.
	//
	// 	- If the two parameters are set to specific dates, the scheduled auto scaling policy can be triggered during the period between the two dates. For example, if **BeginDate*	- is 2021-03-25 and **EndDate*	- is 2021-04-25, the auto scaling policy is valid for one month.
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
	Schedules []*DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules `json:"Schedules,omitempty" xml:"Schedules,omitempty" type:"Repeated"`
	TimeZone  *string                                                                                 `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
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
	return dara.Validate(s)
}

type DescribeApplicationScalingRulesResponseBodyDataApplicationScalingRulesTimerSchedules struct {
	// The point in time. Format: **Hour:Minute**.
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
	// The expected number of instances.
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
