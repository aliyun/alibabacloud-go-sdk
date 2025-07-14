// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateApplicationScalingRuleResponseBody
	GetCode() *string
	SetData(v *UpdateApplicationScalingRuleResponseBodyData) *UpdateApplicationScalingRuleResponseBody
	GetData() *UpdateApplicationScalingRuleResponseBodyData
	SetErrorCode(v string) *UpdateApplicationScalingRuleResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateApplicationScalingRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateApplicationScalingRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateApplicationScalingRuleResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateApplicationScalingRuleResponseBody
	GetTraceId() *string
}

type UpdateApplicationScalingRuleResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *UpdateApplicationScalingRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned. Take note of the following rules:
	//
	// 	- If the call is successful, **ErrorCode*	- is not returned.
	//
	// 	- If the call fails, **ErrorCode*	- is returned. For more information, see the "**Error codes**" section in this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message. Take note of the following rules:
	//
	// 	- If the call is successful, **success*	- is returned.
	//
	// 	- If the call fails, an error code is returned.
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
	// Specifies whether the instances are successfully restarted. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of the request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateApplicationScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateApplicationScalingRuleResponseBody) GetData() *UpdateApplicationScalingRuleResponseBodyData {
	return s.Data
}

func (s *UpdateApplicationScalingRuleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateApplicationScalingRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateApplicationScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationScalingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateApplicationScalingRuleResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateApplicationScalingRuleResponseBody) SetCode(v string) *UpdateApplicationScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBody) SetData(v *UpdateApplicationScalingRuleResponseBodyData) *UpdateApplicationScalingRuleResponseBody {
	s.Data = v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBody) SetErrorCode(v string) *UpdateApplicationScalingRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBody) SetMessage(v string) *UpdateApplicationScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBody) SetRequestId(v string) *UpdateApplicationScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBody) SetSuccess(v bool) *UpdateApplicationScalingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBody) SetTraceId(v string) *UpdateApplicationScalingRuleResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationScalingRuleResponseBodyData struct {
	// The application ID.
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
	EnableIdle *bool  `json:"EnableIdle,omitempty" xml:"EnableIdle,omitempty"`
	// The time when the auto scaling policy was last disabled.
	//
	// example:
	//
	// 1641882854484
	LastDisableTime *int64 `json:"LastDisableTime,omitempty" xml:"LastDisableTime,omitempty"`
	// The details of the metric-based auto scaling policy.
	Metric *UpdateApplicationScalingRuleResponseBodyDataMetric `json:"Metric,omitempty" xml:"Metric,omitempty" type:"Struct"`
	// Specifies whether to enable the auto scaling policy. Valid values:
	//
	// 	- **true**: The auto scaling policy is enabled.
	//
	// 	- **false**: The auto scaling policy is disabled.
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
	// 	- **timing**: a scheduled auto scaling policy
	//
	// 	- **metric**: a metric-based auto scaling policy
	//
	// 	- **mix**: a hybrid auto scaling policy
	//
	// example:
	//
	// timing
	ScaleRuleType *string `json:"ScaleRuleType,omitempty" xml:"ScaleRuleType,omitempty"`
	// The details of the scheduled auto scaling policy.
	Timer *UpdateApplicationScalingRuleResponseBodyDataTimer `json:"Timer,omitempty" xml:"Timer,omitempty" type:"Struct"`
	// The time when the auto scaling policy was updated. Unit: milliseconds.
	//
	// example:
	//
	// 1616642248938
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s UpdateApplicationScalingRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *UpdateApplicationScalingRuleResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *UpdateApplicationScalingRuleResponseBodyData) GetEnableIdle() *bool {
	return s.EnableIdle
}

func (s *UpdateApplicationScalingRuleResponseBodyData) GetLastDisableTime() *int64 {
	return s.LastDisableTime
}

func (s *UpdateApplicationScalingRuleResponseBodyData) GetMetric() *UpdateApplicationScalingRuleResponseBodyDataMetric {
	return s.Metric
}

func (s *UpdateApplicationScalingRuleResponseBodyData) GetScaleRuleEnabled() *bool {
	return s.ScaleRuleEnabled
}

func (s *UpdateApplicationScalingRuleResponseBodyData) GetScaleRuleName() *string {
	return s.ScaleRuleName
}

func (s *UpdateApplicationScalingRuleResponseBodyData) GetScaleRuleType() *string {
	return s.ScaleRuleType
}

func (s *UpdateApplicationScalingRuleResponseBodyData) GetTimer() *UpdateApplicationScalingRuleResponseBodyDataTimer {
	return s.Timer
}

func (s *UpdateApplicationScalingRuleResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *UpdateApplicationScalingRuleResponseBodyData) SetAppId(v string) *UpdateApplicationScalingRuleResponseBodyData {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyData) SetCreateTime(v int64) *UpdateApplicationScalingRuleResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyData) SetEnableIdle(v bool) *UpdateApplicationScalingRuleResponseBodyData {
	s.EnableIdle = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyData) SetLastDisableTime(v int64) *UpdateApplicationScalingRuleResponseBodyData {
	s.LastDisableTime = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyData) SetMetric(v *UpdateApplicationScalingRuleResponseBodyDataMetric) *UpdateApplicationScalingRuleResponseBodyData {
	s.Metric = v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyData) SetScaleRuleEnabled(v bool) *UpdateApplicationScalingRuleResponseBodyData {
	s.ScaleRuleEnabled = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyData) SetScaleRuleName(v string) *UpdateApplicationScalingRuleResponseBodyData {
	s.ScaleRuleName = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyData) SetScaleRuleType(v string) *UpdateApplicationScalingRuleResponseBodyData {
	s.ScaleRuleType = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyData) SetTimer(v *UpdateApplicationScalingRuleResponseBodyDataTimer) *UpdateApplicationScalingRuleResponseBodyData {
	s.Timer = v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyData) SetUpdateTime(v int64) *UpdateApplicationScalingRuleResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationScalingRuleResponseBodyDataMetric struct {
	// The maximum number of instances.
	//
	// example:
	//
	// 3
	MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// The metrics that are used to trigger the auto scaling policy.
	Metrics []*UpdateApplicationScalingRuleResponseBodyDataMetricMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The minimum number of instances.
	//
	// example:
	//
	// 1
	MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
}

func (s UpdateApplicationScalingRuleResponseBodyDataMetric) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyDataMetric) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetric) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetric) GetMetrics() []*UpdateApplicationScalingRuleResponseBodyDataMetricMetrics {
	return s.Metrics
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetric) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetric) SetMaxReplicas(v int32) *UpdateApplicationScalingRuleResponseBodyDataMetric {
	s.MaxReplicas = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetric) SetMetrics(v []*UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) *UpdateApplicationScalingRuleResponseBodyDataMetric {
	s.Metrics = v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetric) SetMinReplicas(v int32) *UpdateApplicationScalingRuleResponseBodyDataMetric {
	s.MinReplicas = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetric) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationScalingRuleResponseBodyDataMetricMetrics struct {
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
	// 	- **tcpActiveConn**: the average number of active TCP connections of an application instance within 30 seconds.
	//
	// 	- **SLB_QPS**: the average QPS of the Internet-facing SLB instance associated with an application instance within 15 seconds.
	//
	// 	- **SLB_RT**: the average response time of the Internet-facing SLB instance within 15 seconds.
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

func (s UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) GetMetricTargetAverageUtilization() *int32 {
	return s.MetricTargetAverageUtilization
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) GetMetricType() *string {
	return s.MetricType
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) GetSlbId() *string {
	return s.SlbId
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) GetSlbLogstore() *string {
	return s.SlbLogstore
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) GetSlbProject() *string {
	return s.SlbProject
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) GetVport() *string {
	return s.Vport
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) SetMetricTargetAverageUtilization(v int32) *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.MetricTargetAverageUtilization = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) SetMetricType(v string) *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.MetricType = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) SetSlbId(v string) *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.SlbId = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) SetSlbLogstore(v string) *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.SlbLogstore = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) SetSlbProject(v string) *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.SlbProject = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) SetVport(v string) *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.Vport = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataMetricMetrics) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationScalingRuleResponseBodyDataTimer struct {
	// The start date of the validity period of the scheduled auto scaling policy. Parameter description:
	//
	// 	- If **BeginDate*	- and **EndDate*	- are set to **null**, the auto scaling policy is a long-term policy. Default values of the beginDate and endDate parameters: null.
	//
	// 	- If the two parameters are set to specific dates, the scheduled auto scaling policy can be triggered during the period between the two dates. For example, if **BeginDate*	- is set to 2021-03-25 and **EndDate*	- is set to 2021-04-25, the auto scaling policy is valid for one month.
	//
	// example:
	//
	// 2021-03-25
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// The end date of the validity period of the scheduled auto scaling policy. Take note of the following rules:
	//
	// 	- If **BeginDate*	- and **EndDate*	- are set to **null**, the auto scaling policy is a long-term policy. Default values of the beginDate and endDate parameters: null.
	//
	// 	- If the two parameters are set to specific dates, the scheduled auto scaling policy can be triggered during the period between the two dates. For example, if **BeginDate*	- is set to 2021-03-25 and **EndDate*	- is set to 2021-04-25, the auto scaling policy is valid for one month.
	//
	// example:
	//
	// 2021-04-25
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The frequency at which the scheduled auto scaling policy is executed. Valid values:
	//
	// 	- **\\	- \\	- \\***: The scheduled auto scaling policy is executed at a specified point in time every day.
	//
	// 	- **\\	- \\	- Fri,Mon**: The scheduled auto scaling policy is executed at a specified point in time on one or more days of each week. GMT+8 is used. Valid values:
	//
	//     	- **Sun**
	//
	//     	- **Mon**
	//
	//     	- **Tue**
	//
	//     	- **Wed**
	//
	//     	- **Thu**
	//
	//     	- **Fri**
	//
	//     	- **Sat**
	//
	// 	- **1,2,3,28,31 \\	- \\***: The scheduled auto scaling policy is executed at a specified point in time on one or more days of each month. Valid values: 1 to 31. If the month does not have a 31st day, the auto scaling policy is executed on the specified days other than the 31st day.
	//
	// example:
	//
	// 	- 	- *
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The points in time at which the auto scaling policy is triggered within one day.
	Schedules []*UpdateApplicationScalingRuleResponseBodyDataTimerSchedules `json:"Schedules,omitempty" xml:"Schedules,omitempty" type:"Repeated"`
}

func (s UpdateApplicationScalingRuleResponseBodyDataTimer) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyDataTimer) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimer) GetBeginDate() *string {
	return s.BeginDate
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimer) GetEndDate() *string {
	return s.EndDate
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimer) GetPeriod() *string {
	return s.Period
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimer) GetSchedules() []*UpdateApplicationScalingRuleResponseBodyDataTimerSchedules {
	return s.Schedules
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimer) SetBeginDate(v string) *UpdateApplicationScalingRuleResponseBodyDataTimer {
	s.BeginDate = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimer) SetEndDate(v string) *UpdateApplicationScalingRuleResponseBodyDataTimer {
	s.EndDate = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimer) SetPeriod(v string) *UpdateApplicationScalingRuleResponseBodyDataTimer {
	s.Period = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimer) SetSchedules(v []*UpdateApplicationScalingRuleResponseBodyDataTimerSchedules) *UpdateApplicationScalingRuleResponseBodyDataTimer {
	s.Schedules = v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimer) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationScalingRuleResponseBodyDataTimerSchedules struct {
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
	// 3
	TargetReplicas *int32 `json:"TargetReplicas,omitempty" xml:"TargetReplicas,omitempty"`
}

func (s UpdateApplicationScalingRuleResponseBodyDataTimerSchedules) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyDataTimerSchedules) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimerSchedules) GetAtTime() *string {
	return s.AtTime
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimerSchedules) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimerSchedules) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimerSchedules) GetTargetReplicas() *int32 {
	return s.TargetReplicas
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimerSchedules) SetAtTime(v string) *UpdateApplicationScalingRuleResponseBodyDataTimerSchedules {
	s.AtTime = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimerSchedules) SetMaxReplicas(v int32) *UpdateApplicationScalingRuleResponseBodyDataTimerSchedules {
	s.MaxReplicas = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimerSchedules) SetMinReplicas(v int32) *UpdateApplicationScalingRuleResponseBodyDataTimerSchedules {
	s.MinReplicas = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimerSchedules) SetTargetReplicas(v int32) *UpdateApplicationScalingRuleResponseBodyDataTimerSchedules {
	s.TargetReplicas = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyDataTimerSchedules) Validate() error {
	return dara.Validate(s)
}
