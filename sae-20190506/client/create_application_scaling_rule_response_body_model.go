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
	// The HTTP status code or the error code. Valid values:
	//
	// 	- **2xx**: The request was successful.
	//
	// 	- **3xx**: The request was redirected.
	//
	// 	- **4xx**: The request failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *CreateApplicationScalingRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The status code. Value values:
	//
	// 	- If the request was successful, **ErrorCode*	- is not returned.
	//
	// 	- If the request failed, **ErrorCode*	- is returned. For more information, see **Error codes*	- in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned. The following limits are imposed on the ID:
	//
	// 	- If the request was successful, **success*	- is returned.
	//
	// 	- An error code is returned when a request failed.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the application instances were restarted. Valid values:
	//
	// 	- **true**: The application instances were restarted.
	//
	// 	- **false**: The application instances failed to be restarted.
	//
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
	// null
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// null null
	//
	// example:
	//
	// 1616642248938
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EnableIdle *bool  `json:"EnableIdle,omitempty" xml:"EnableIdle,omitempty"`
	// null
	//
	// example:
	//
	// 1641882854484
	LastDisableTime *int64 `json:"LastDisableTime,omitempty" xml:"LastDisableTime,omitempty"`
	// The details of the metric-based auto scaling policy.
	Metric *CreateApplicationScalingRuleResponseBodyDataMetric `json:"Metric,omitempty" xml:"Metric,omitempty" type:"Struct"`
	// null null
	//
	// 	- **null**
	//
	// 	- **null**
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
	// null null
	//
	// 	- **null**
	//
	// 	- **metric**: a metric-based auto scaling policy.
	//
	// 	- **mix**: a hybrid auto scaling policy.
	//
	// example:
	//
	// timing
	ScaleRuleType *string `json:"ScaleRuleType,omitempty" xml:"ScaleRuleType,omitempty"`
	// The details of the scheduled auto scaling policy.
	Timer *CreateApplicationScalingRuleResponseBodyDataTimer `json:"Timer,omitempty" xml:"Timer,omitempty" type:"Struct"`
	// null null
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
	// The maximum number of Elastic Compute Service (ECS) instances supported by the node pool.
	//
	// example:
	//
	// 3
	MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// The metrics that are used to trigger the auto scaling policy.
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
	// The limit on the metric.
	//
	// 	- The limit on the CPU utilization. Unit: percentage.
	//
	// 	- The limit on the memory usage. Unit: percentage.
	//
	// 	- The limit on the queries per second (QPS). Unit: seconds.
	//
	// 	- The limit on the response time. Unit: milliseconds.
	//
	// 	- The limit on the average number of active TCP connections per second.
	//
	// 	- The limit on the QPS of the Internet-facing SLB instance.
	//
	// 	- The limit on the response time of the Internet-facing SLB instance. Unit: milliseconds.
	//
	// 	- The limit on the QPS of the internal-facing SLB instance.
	//
	// 	- The limit on the response time of the internal-facing SLB instance. Unit: milliseconds.
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
	// 	- **QPS**: the average QPS within 1 minute per Java application instance.
	//
	// 	- **RT**: the average response time of all API operations within 1 minute in the Java application.
	//
	// 	- **tcpActiveConn**: the average number of active TCP connections within 30 seconds per instance.
	//
	// 	- **SLB_QPS**: the average QPS of the Internet-facing SLB instance within 15 seconds per instance.
	//
	// 	- **SLB_RT**: the average response time of the Internet-facing SLB instance within 15 seconds.
	//
	// 	- **INTRANET_SLB_QPS**: the average QPS of the internal-facing SLB instance within 15 seconds per instance.
	//
	// 	- **INTRANET_SLB_RT**: the average response time of the internal-facing SLB instance within 15 seconds.
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
	// The Logstore that stores the SLB access logs.
	//
	// example:
	//
	// test
	SlbLogstore *string `json:"SlbLogstore,omitempty" xml:"SlbLogstore,omitempty"`
	// The project that stores the SLB access logs.
	//
	// example:
	//
	// test
	SlbProject *string `json:"SlbProject,omitempty" xml:"SlbProject,omitempty"`
	// The port number of the SLB instance.
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
	// The start date of the validity period of the scheduled auto scaling policy.
	//
	// 	- **null*	- (default): If you set **BeginDate*	- and **EndDate*	- to null, the scheduled auto scaling policy can always be triggered.
	//
	// 	- If the two parameters are set to specific dates, the scheduled auto scaling policy can be triggered during the period between the two dates. For example, if **BeginDate*	- is set to 2021-03-25 and **EndDate*	- is set to 2021-04-25, the auto scaling policy is valid for one month.
	//
	// example:
	//
	// 2021-03-25
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// The end date of the validity period of the scheduled auto scaling policy.
	//
	// 	- **null*	- (default): If you set **BeginDate*	- and **EndDate*	- to null, the scheduled auto scaling policy can always be triggered.
	//
	// 	- If the two parameters are set to specific dates, the scheduled auto scaling policy can be triggered during the period between the two dates. For example, if **BeginDate*	- is set to 2021-03-25 and **EndDate*	- is set to 2021-04-25, the auto scaling policy is valid for one month.
	//
	// example:
	//
	// 2021-04-25
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The days on which the scheduled auto scaling policy takes effect. Valid values:
	//
	// 	- **\\	- \\	- \\***: The scheduled auto scaling policy is executed at a specified point in time every day.
	//
	// 	- **\\	- \\	- Fri,Mon**: The scheduled auto scaling policy is executed at a specified point in time on one or more days every week. The time must be in GMT+8. Valid values:
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
	// 	- **1,2,3,28,31 \\	- \\***: The scheduled auto scaling policy is executed at a specified point in time on one or more dates of each month. Valid values: 1 to 31. If a month does not have the 31st day, the auto scaling policy is executed on the specified days other than the 31st day.
	//
	// example:
	//
	// 	- 	- *
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The points in time at which the auto scaling policy is triggered within one day.
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
	// 10
	MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// The minimum number of instances.
	//
	// example:
	//
	// 5
	MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	// The expected number of instances.
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
