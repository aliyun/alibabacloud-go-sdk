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
	// The HTTP status code or a POP error code.
	//
	// - **2xx**: The request was successful.
	//
	// - **3xx**: The request was redirected.
	//
	// - **4xx**: A client-side error occurred.
	//
	// - **5xx**: A server-side error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *UpdateApplicationScalingRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - This parameter is returned only if the request fails.
	//
	// - For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The response message.
	//
	// - **success*	- is returned if the request is successful.
	//
	// - An error message is returned if the request fails.
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
	// Specifies whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to query call details.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApplicationScalingRuleResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the scaling policy was created, in milliseconds.
	//
	// example:
	//
	// 1616642248938
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Specifies whether to enable idle mode.
	EnableIdle *bool `json:"EnableIdle,omitempty" xml:"EnableIdle,omitempty"`
	// The time when the scaling policy was last disabled, in milliseconds.
	//
	// example:
	//
	// 1641882854484
	LastDisableTime *int64 `json:"LastDisableTime,omitempty" xml:"LastDisableTime,omitempty"`
	// The configuration for metric-based scaling.
	Metric *UpdateApplicationScalingRuleResponseBodyDataMetric `json:"Metric,omitempty" xml:"Metric,omitempty" type:"Struct"`
	// Specifies whether the scaling policy is enabled. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	ScaleRuleEnabled *bool `json:"ScaleRuleEnabled,omitempty" xml:"ScaleRuleEnabled,omitempty"`
	// The name of the scaling policy.
	//
	// example:
	//
	// test
	ScaleRuleName *string `json:"ScaleRuleName,omitempty" xml:"ScaleRuleName,omitempty"`
	// The type of the scaling policy. Valid values:
	//
	// - **timing**: scheduled scaling
	//
	// - **metric**: metric-based scaling
	//
	// - **mix**: hybrid scaling
	//
	// example:
	//
	// timing
	ScaleRuleType *string `json:"ScaleRuleType,omitempty" xml:"ScaleRuleType,omitempty"`
	// The configuration for scheduled scaling.
	Timer *UpdateApplicationScalingRuleResponseBodyDataTimer `json:"Timer,omitempty" xml:"Timer,omitempty" type:"Struct"`
	// The time when the scaling policy was updated, in milliseconds.
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

type UpdateApplicationScalingRuleResponseBodyDataMetric struct {
	// The maximum number of instances.
	//
	// example:
	//
	// 3
	MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// The metrics that trigger scaling actions.
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

type UpdateApplicationScalingRuleResponseBodyDataMetricMetrics struct {
	// The target value for the specified metric. The unit varies based on the metric type.
	//
	// - Target CPU utilization, in percentage.
	//
	// - Target memory usage, in percentage.
	//
	// - Target queries per second (QPS).
	//
	// - Target response time, in milliseconds.
	//
	// - The target number of active TCP connections.
	//
	// - Target QPS for the public-facing SLB instance.
	//
	// - Target response time of the public-facing SLB instance, in milliseconds.
	//
	// - Target QPS for the internal SLB instance.
	//
	// - Target response time of the internal SLB instance, in milliseconds.
	//
	// example:
	//
	// 20
	MetricTargetAverageUtilization *int32 `json:"MetricTargetAverageUtilization,omitempty" xml:"MetricTargetAverageUtilization,omitempty"`
	// The metric that triggers the scaling policy. Valid values:
	//
	// - **CPU**: CPU utilization.
	//
	// - **MEMORY**: memory usage.
	//
	// - **QPS**: The average queries per second (QPS) per instance over the last minute. This applies only to Java applications.
	//
	// - **RT**: The average response time (RT) of all service interfaces in the application over the last minute. This applies only to Java applications.
	//
	// - **tcpActiveConn**: The average number of active TCP connections per instance over the last 30 seconds.
	//
	// - **SLB_QPS**: The average QPS from the public-facing SLB, per instance, over the last 15 seconds.
	//
	// - **SLB_RT**: The average response time of a public-facing SLB instance over the last 15 seconds.
	//
	// - **INTRANET_SLB_QPS**: The average QPS from the internal SLB, per instance, over the last 15 seconds.
	//
	// - **INTRANET_SLB_RT**: The average response time of an internal SLB instance over the last 15 seconds.
	//
	// example:
	//
	// CPU
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The SLB instance ID.
	//
	// example:
	//
	// lb-xxx
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The name of the Logstore for SLB access logs.
	//
	// example:
	//
	// test
	SlbLogstore *string `json:"SlbLogstore,omitempty" xml:"SlbLogstore,omitempty"`
	// The name of the Log Service Project for SLB access logs.
	//
	// example:
	//
	// test
	SlbProject *string `json:"SlbProject,omitempty" xml:"SlbProject,omitempty"`
	// The SLB port.
	//
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
	// The start date of the short-term scheduled scaling policy.
	//
	// - If **BeginDate*	- and **EndDate*	- are both set to **null**, the policy is long-term by default.
	//
	// - If you specify a date range, for example, **BeginDate*	- is set to 2021-03-25 and **EndDate*	- is set to 2021-04-25, the policy is effective for one month.
	//
	// example:
	//
	// 2021-03-25
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// The end date of the short-term scheduled scaling policy.
	//
	// - If **BeginDate*	- and **EndDate*	- are both set to **null**, the policy is long-term by default.
	//
	// - If you specify a date range, for example, **BeginDate*	- is set to 2021-03-25 and **EndDate*	- is set to 2021-04-25, the policy is effective for one month.
	//
	// example:
	//
	// 2021-04-25
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The recurrence schedule for the scaling policy.
	//
	// - **\\	- \\	- \\***: The policy runs at a specified time every day.
	//
	// - **\\	- \\	- Fri,Mon**: The policy runs at a specified time on specific days of a week. You can select multiple days. The time is in the GMT+8 time zone. Valid values:
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
	// - **1,2,3,28,31 \\	- \\***: The policy runs at a specified time on specific days of a month. You can select multiple days. If a month does not have a specific day, such as the 31st, the policy skips that day.
	//
	// example:
	//
	// 	- 	- *
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The schedules for the scaling policy.
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

type UpdateApplicationScalingRuleResponseBodyDataTimerSchedules struct {
	// The time at which the scaling action is triggered. Format: **HH:mm**.
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
