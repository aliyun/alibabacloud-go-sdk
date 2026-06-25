// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateApplicationScalingRuleResponseBody
	GetCode() *string
	SetData(v *CreateApplicationScalingRuleResponseBodyData) *CreateApplicationScalingRuleResponseBody
	GetData() *CreateApplicationScalingRuleResponseBodyData
	SetErrorCode(v string) *CreateApplicationScalingRuleResponseBody
	GetErrorCode() *string
	SetMessage(v string) *CreateApplicationScalingRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateApplicationScalingRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateApplicationScalingRuleResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CreateApplicationScalingRuleResponseBody
	GetTraceId() *string
}

type CreateApplicationScalingRuleResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// - **2xx**: The request was successful.
	//
	// - **3xx**: The request was redirected.
	//
	// - **4xx**: A client error occurred.
	//
	// - **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateApplicationScalingRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// - This parameter is not returned if the request is successful.
	//
	// - An error code is returned if the request fails. For more information, see the **Error codes*	- section in this topic.
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
	// - **true**: The API call was successful.
	//
	// - **false**: The API call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID. You can use this ID to query the details of a request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateApplicationScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateApplicationScalingRuleResponseBody) GetData() *CreateApplicationScalingRuleResponseBodyData {
	return s.Data
}

func (s *CreateApplicationScalingRuleResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateApplicationScalingRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateApplicationScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationScalingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateApplicationScalingRuleResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateApplicationScalingRuleResponseBody) SetCode(v string) *CreateApplicationScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBody) SetData(v *CreateApplicationScalingRuleResponseBodyData) *CreateApplicationScalingRuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateApplicationScalingRuleResponseBody) SetErrorCode(v string) *CreateApplicationScalingRuleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBody) SetMessage(v string) *CreateApplicationScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBody) SetRequestId(v string) *CreateApplicationScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBody) SetSuccess(v bool) *CreateApplicationScalingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBody) SetTraceId(v string) *CreateApplicationScalingRuleResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateApplicationScalingRuleResponseBodyData struct {
	// The ID of the application.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the policy was created, in milliseconds.
	//
	// example:
	//
	// 1616642248938
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether idle mode is enabled.
	EnableIdle *bool `json:"EnableIdle,omitempty" xml:"EnableIdle,omitempty"`
	// The time when the auto scaling policy was last disabled, in milliseconds.
	//
	// example:
	//
	// 1641882854484
	LastDisableTime *int64 `json:"LastDisableTime,omitempty" xml:"LastDisableTime,omitempty"`
	// The configurations for the metric-based auto scaling policy.
	Metric *CreateApplicationScalingRuleResponseBodyDataMetric `json:"Metric,omitempty" xml:"Metric,omitempty" type:"Struct"`
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
	// - **timing**: scheduled auto scaling.
	//
	// - **metric**: metric-based auto scaling.
	//
	// - **mix**: mixed auto scaling.
	//
	// example:
	//
	// timing
	ScaleRuleType *string `json:"ScaleRuleType,omitempty" xml:"ScaleRuleType,omitempty"`
	// The configurations for the scheduled auto scaling policy.
	Timer *CreateApplicationScalingRuleResponseBodyDataTimer `json:"Timer,omitempty" xml:"Timer,omitempty" type:"Struct"`
	// The time when the policy was last updated, in milliseconds.
	//
	// example:
	//
	// 1616642248938
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateApplicationScalingRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *CreateApplicationScalingRuleResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateApplicationScalingRuleResponseBodyData) GetEnableIdle() *bool {
	return s.EnableIdle
}

func (s *CreateApplicationScalingRuleResponseBodyData) GetLastDisableTime() *int64 {
	return s.LastDisableTime
}

func (s *CreateApplicationScalingRuleResponseBodyData) GetMetric() *CreateApplicationScalingRuleResponseBodyDataMetric {
	return s.Metric
}

func (s *CreateApplicationScalingRuleResponseBodyData) GetScaleRuleEnabled() *bool {
	return s.ScaleRuleEnabled
}

func (s *CreateApplicationScalingRuleResponseBodyData) GetScaleRuleName() *string {
	return s.ScaleRuleName
}

func (s *CreateApplicationScalingRuleResponseBodyData) GetScaleRuleType() *string {
	return s.ScaleRuleType
}

func (s *CreateApplicationScalingRuleResponseBodyData) GetTimer() *CreateApplicationScalingRuleResponseBodyDataTimer {
	return s.Timer
}

func (s *CreateApplicationScalingRuleResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *CreateApplicationScalingRuleResponseBodyData) SetAppId(v string) *CreateApplicationScalingRuleResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyData) SetCreateTime(v int64) *CreateApplicationScalingRuleResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyData) SetEnableIdle(v bool) *CreateApplicationScalingRuleResponseBodyData {
	s.EnableIdle = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyData) SetLastDisableTime(v int64) *CreateApplicationScalingRuleResponseBodyData {
	s.LastDisableTime = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyData) SetMetric(v *CreateApplicationScalingRuleResponseBodyDataMetric) *CreateApplicationScalingRuleResponseBodyData {
	s.Metric = v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyData) SetScaleRuleEnabled(v bool) *CreateApplicationScalingRuleResponseBodyData {
	s.ScaleRuleEnabled = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyData) SetScaleRuleName(v string) *CreateApplicationScalingRuleResponseBodyData {
	s.ScaleRuleName = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyData) SetScaleRuleType(v string) *CreateApplicationScalingRuleResponseBodyData {
	s.ScaleRuleType = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyData) SetTimer(v *CreateApplicationScalingRuleResponseBodyDataTimer) *CreateApplicationScalingRuleResponseBodyData {
	s.Timer = v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyData) SetUpdateTime(v int64) *CreateApplicationScalingRuleResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyData) Validate() error {
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

type CreateApplicationScalingRuleResponseBodyDataMetric struct {
	// The maximum number of instances.
	//
	// example:
	//
	// 3
	MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// An array of objects that define the metrics for the metric-based auto scaling policy.
	Metrics []*CreateApplicationScalingRuleResponseBodyDataMetricMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The minimum number of instances.
	//
	// example:
	//
	// 1
	MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
}

func (s CreateApplicationScalingRuleResponseBodyDataMetric) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyDataMetric) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetric) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetric) GetMetrics() []*CreateApplicationScalingRuleResponseBodyDataMetricMetrics {
	return s.Metrics
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetric) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetric) SetMaxReplicas(v int32) *CreateApplicationScalingRuleResponseBodyDataMetric {
	s.MaxReplicas = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetric) SetMetrics(v []*CreateApplicationScalingRuleResponseBodyDataMetricMetrics) *CreateApplicationScalingRuleResponseBodyDataMetric {
	s.Metrics = v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetric) SetMinReplicas(v int32) *CreateApplicationScalingRuleResponseBodyDataMetric {
	s.MinReplicas = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetric) Validate() error {
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

type CreateApplicationScalingRuleResponseBodyDataMetricMetrics struct {
	// The target value for the metric.
	//
	// - The target CPU utilization, in percentage.
	//
	// - The target memory utilization, in percentage.
	//
	// - The target QPS.
	//
	// - The target response time, in milliseconds.
	//
	// - The target average number of active TCP connections per second.
	//
	// - The target QPS for a public-facing SLB instance.
	//
	// - The target response time for a public-facing SLB instance, in milliseconds.
	//
	// - The target QPS for an internal-facing SLB instance.
	//
	// - The target response time for an internal-facing SLB instance, in milliseconds.
	//
	// example:
	//
	// 20
	MetricTargetAverageUtilization *int32 `json:"MetricTargetAverageUtilization,omitempty" xml:"MetricTargetAverageUtilization,omitempty"`
	// The type of the metric that triggers the auto scaling policy. Valid values:
	//
	// - **CPU**: CPU utilization.
	//
	// - **MEMORY**: memory utilization.
	//
	// - **QPS**: The average QPS per instance over 1 minute for a Java application.
	//
	// - **RT**: The average response time across all service endpoints over 1 minute for a Java application.
	//
	// - **tcpActiveConn**: The average number of TCP active connections per instance over 30 seconds.
	//
	// - **SLB_QPS**: The average QPS per instance for a public SLB instance, measured over 15 seconds.
	//
	// - **SLB_RT**: The average response time for a public SLB instance, measured over 15 seconds.
	//
	// - **INTRANET_SLB_QPS**: The average QPS per instance for an internal-facing SLB instance, measured over 15 seconds.
	//
	// - **INTRANET_SLB_RT**: The average response time for an internal-facing SLB instance, measured over 15 seconds.
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
	// The Log Service Logstore for SLB access logs.
	//
	// example:
	//
	// test
	SlbLogstore *string `json:"SlbLogstore,omitempty" xml:"SlbLogstore,omitempty"`
	// The Log Service project for SLB access logs.
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

func (s CreateApplicationScalingRuleResponseBodyDataMetricMetrics) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyDataMetricMetrics) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetricMetrics) GetMetricTargetAverageUtilization() *int32 {
	return s.MetricTargetAverageUtilization
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetricMetrics) GetMetricType() *string {
	return s.MetricType
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetricMetrics) GetSlbId() *string {
	return s.SlbId
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetricMetrics) GetSlbLogstore() *string {
	return s.SlbLogstore
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetricMetrics) GetSlbProject() *string {
	return s.SlbProject
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetricMetrics) GetVport() *string {
	return s.Vport
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetricMetrics) SetMetricTargetAverageUtilization(v int32) *CreateApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.MetricTargetAverageUtilization = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetricMetrics) SetMetricType(v string) *CreateApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.MetricType = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetricMetrics) SetSlbId(v string) *CreateApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.SlbId = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetricMetrics) SetSlbLogstore(v string) *CreateApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.SlbLogstore = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetricMetrics) SetSlbProject(v string) *CreateApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.SlbProject = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetricMetrics) SetVport(v string) *CreateApplicationScalingRuleResponseBodyDataMetricMetrics {
	s.Vport = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataMetricMetrics) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationScalingRuleResponseBodyDataTimer struct {
	// The start date of the scheduled auto scaling policy.
	//
	// - If both **BeginDate*	- and **EndDate*	- are **null**, the policy is a long-term policy. This is the default.
	//
	// - For example, if you set **BeginDate*	- to 2021-03-25 and **EndDate*	- to 2021-04-25, the policy is active for one month.
	//
	// example:
	//
	// 2021-03-25
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// The end date of the scheduled auto scaling policy.
	//
	// - If both **BeginDate*	- and **EndDate*	- are **null**, the policy is a long-term policy. This is the default.
	//
	// - For example, if you set **BeginDate*	- to 2021-03-25 and **EndDate*	- to 2021-04-25, the policy is active for one month.
	//
	// example:
	//
	// 2021-04-25
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The recurrence pattern for the scheduled auto scaling policy. Valid values:
	//
	// - **\\	- \\	- \\***: The policy is executed at a specified time every day.
	//
	// - **\\	- \\	- Fri,Mon**: The policy is executed at a specified time on specific days of the week. You can select multiple days. The time is in the GMT+8 time zone. Valid values:
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
	// - **1,2,3,28,31 \\	- \\***: The policy is executed at a specified time on specific days of a month. You can select multiple days. The value can be from 1 to 31. If a specified day does not exist in a given month (for example, the 31st), the policy is not executed on that day.
	//
	// example:
	//
	// 	- 	- *
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The trigger points for the scheduled auto scaling policy.
	Schedules []*CreateApplicationScalingRuleResponseBodyDataTimerSchedules `json:"Schedules,omitempty" xml:"Schedules,omitempty" type:"Repeated"`
}

func (s CreateApplicationScalingRuleResponseBodyDataTimer) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyDataTimer) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimer) GetBeginDate() *string {
	return s.BeginDate
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimer) GetEndDate() *string {
	return s.EndDate
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimer) GetPeriod() *string {
	return s.Period
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimer) GetSchedules() []*CreateApplicationScalingRuleResponseBodyDataTimerSchedules {
	return s.Schedules
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimer) SetBeginDate(v string) *CreateApplicationScalingRuleResponseBodyDataTimer {
	s.BeginDate = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimer) SetEndDate(v string) *CreateApplicationScalingRuleResponseBodyDataTimer {
	s.EndDate = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimer) SetPeriod(v string) *CreateApplicationScalingRuleResponseBodyDataTimer {
	s.Period = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimer) SetSchedules(v []*CreateApplicationScalingRuleResponseBodyDataTimerSchedules) *CreateApplicationScalingRuleResponseBodyDataTimer {
	s.Schedules = v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimer) Validate() error {
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

type CreateApplicationScalingRuleResponseBodyDataTimerSchedules struct {
	// The trigger time. The format is **HH:mm**.
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
	// 5
	MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	// The target number of instances.
	//
	// example:
	//
	// 3
	TargetReplicas *int32 `json:"TargetReplicas,omitempty" xml:"TargetReplicas,omitempty"`
}

func (s CreateApplicationScalingRuleResponseBodyDataTimerSchedules) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyDataTimerSchedules) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimerSchedules) GetAtTime() *string {
	return s.AtTime
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimerSchedules) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimerSchedules) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimerSchedules) GetTargetReplicas() *int32 {
	return s.TargetReplicas
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimerSchedules) SetAtTime(v string) *CreateApplicationScalingRuleResponseBodyDataTimerSchedules {
	s.AtTime = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimerSchedules) SetMaxReplicas(v int32) *CreateApplicationScalingRuleResponseBodyDataTimerSchedules {
	s.MaxReplicas = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimerSchedules) SetMinReplicas(v int32) *CreateApplicationScalingRuleResponseBodyDataTimerSchedules {
	s.MinReplicas = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimerSchedules) SetTargetReplicas(v int32) *CreateApplicationScalingRuleResponseBodyDataTimerSchedules {
	s.TargetReplicas = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyDataTimerSchedules) Validate() error {
	return dara.Validate(s)
}
