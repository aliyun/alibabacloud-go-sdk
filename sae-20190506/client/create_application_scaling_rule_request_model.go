// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateApplicationScalingRuleRequest
	GetAppId() *string
	SetEnableIdle(v bool) *CreateApplicationScalingRuleRequest
	GetEnableIdle() *bool
	SetMinReadyInstanceRatio(v int32) *CreateApplicationScalingRuleRequest
	GetMinReadyInstanceRatio() *int32
	SetMinReadyInstances(v int32) *CreateApplicationScalingRuleRequest
	GetMinReadyInstances() *int32
	SetScalingRuleEnable(v bool) *CreateApplicationScalingRuleRequest
	GetScalingRuleEnable() *bool
	SetScalingRuleMetric(v string) *CreateApplicationScalingRuleRequest
	GetScalingRuleMetric() *string
	SetScalingRuleName(v string) *CreateApplicationScalingRuleRequest
	GetScalingRuleName() *string
	SetScalingRuleTimer(v string) *CreateApplicationScalingRuleRequest
	GetScalingRuleTimer() *string
	SetScalingRuleType(v string) *CreateApplicationScalingRuleRequest
	GetScalingRuleType() *string
}

type CreateApplicationScalingRuleRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether to enable the idle mode.
	EnableIdle *bool `json:"EnableIdle,omitempty" xml:"EnableIdle,omitempty"`
	// The percentage of the minimum number of surviving instances. Valid values:
	//
	// - **-1**: the initial value, which indicates that the percentage is not used.
	//
	// - **0~100**: the percentage, rounded up. For example, if you set this parameter to 50 and the current number of instances is 5, the minimum number of surviving instances is 3.
	//
	// > When both **MinReadyInstances*	- and **MinReadyInstanceRatio*	- are specified and the value of **MinReadyInstanceRatio*	- is not **-1**, the **MinReadyInstanceRatio*	- parameter prevails. For example, if **MinReadyInstances*	- is set to **5*	- and **MinReadyInstanceRatio*	- is set to **50**, the value **50*	- is used to calculate the minimum number of surviving instances.
	//
	// example:
	//
	// -1
	MinReadyInstanceRatio *int32 `json:"MinReadyInstanceRatio,omitempty" xml:"MinReadyInstanceRatio,omitempty"`
	// The minimum number of surviving instances. Valid values:
	//
	// - If you set this parameter to **0**, services are interrupted during the application upgrade.
	//
	// - If you set this parameter to **-1**, the system recommended value is used, which is 25% of the current number of instances. For example, if the current number of instances is 5, 5 × 25% = 1.25. After rounding up, the minimum number of surviving instances is 2.
	//
	// > We recommend that you set the minimum number of surviving instances for each rolling deployment to a value greater than or equal to 1 to ensure service continuity.
	//
	// example:
	//
	// 3
	MinReadyInstances *int32 `json:"MinReadyInstances,omitempty" xml:"MinReadyInstances,omitempty"`
	// Specifies whether to enable the elastic scaling policy. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// true
	ScalingRuleEnable *bool `json:"ScalingRuleEnable,omitempty" xml:"ScalingRuleEnable,omitempty"`
	// The configuration of the monitoring metric-based elastic policy. This parameter is required when you select the monitoring metric-based elastic policy settings.
	//
	// Metric description:
	//
	// - **maxReplicas**: the maximum number of instances for the application.
	//
	// - **minReplicas**: the minimum number of instances for the application.
	//
	// - **metricType**: the trigger condition for the monitoring metric. Valid values:
	//
	//     - **CPU**: CPU utilization.
	//
	//     - **MEMORY**: memory utilization.
	//
	//     - **QPS**: the average queries per second (QPS) of a single instance of a Java application within 1 minute.
	//
	//     - **RT**: the average response time (RT) of all service interfaces of a Java application within 1 minute.
	//
	//     - **tcpActiveConn**: the average number of active TCP connections of a single instance within 30 seconds.
	//
	//     - **SLB_QPS**: the average public-facing Server Load Balancer (SLB) QPS of a single instance within 15 seconds.
	//
	//     - **SLB_RT**: the average public-Internet-facing SLB response time within 15 seconds.
	//
	//     - **INTRANET_SLB_QPS**: the average private network SLB QPS of a single instance within 15 seconds.
	//
	//     - **INTRANET_SLB_RT**: the average private network SLB response time within 15 seconds.
	//
	// - **metricTargetAverageUtilization**: the target value of the monitoring metric based on the **metricType**.
	//
	//     - Target CPU utilization, in percentage.
	//
	//     - Target memory utilization, in percentage.
	//
	//     - QPS.
	//
	//     - Response time, in milliseconds.
	//
	//     - Average number of active TCP connections, in connections per second.
	//
	//     - Public-Internet-facing SLB QPS.
	//
	//     - Public-Internet-facing SLB response time, in milliseconds.
	//
	//     - Private network SLB QPS.
	//
	//     - Private network SLB response time, in milliseconds.
	//
	// - **slbId**: the SLB instance ID.
	//
	// - **slbProject**: the Simple Log Service project.
	//
	// - **slbLogstore**: the Simple Log Service Logstore.
	//
	// - **vport**: the SLB listener port. HTTP and HTTPS protocols are supported.
	//
	// - **scaleUpRules**: the scale-out rules for the application.
	//
	// - **scaleDownRules**: the scale-in rules for the application.
	//
	// - **step**: the step size for scaling out or scaling in. The maximum number of instances that can be scaled out or in per unit of time.
	//
	// - **disabled**: specifies whether to disable scale-in. If enabled, the application instances are never scaled in, which effectively prevents service risks caused by scale-in during traffic peaks. Valid values:
	//
	//     - **true**: enabled.
	//
	//     - **false**: disabled. This is the default value.
	//
	// - **stabilizationWindowSeconds**: the cooldown period for scaling out or scaling in. Valid values: [0, 3600\\]. Unit: seconds. Default value: 0.
	//
	// > You can configure one or more monitoring metrics. If multiple monitoring metrics are configured, the application scales out as long as any metric is greater than or equal to the target value. The number of instances after scale-out does not exceed the maximum number of instances for the application. If all metrics are less than the target value, the application scales in. The number of instances after scale-in is not less than the minimum number of instances for the application.
	//
	// example:
	//
	// {"maxReplicas":3,"minReplicas":1,"metrics":[{"metricType":"CPU","metricTargetAverageUtilization":20},{"metricType":"MEMORY","metricTargetAverageUtilization":30},{"metricType":"tcpActiveConn","metricTargetAverageUtilization":20},{"metricType":"SLB_QPS","MetricTargetAverageUtilization":25,"slbId":"lb-xxx","slbProject":"aliyun-fc-cn-hangzhou-d95881d9-5d3c-5f26-a6b8-************","slbLogstore":"function-log","vport":"80"},{"metricType":"SLB_RT","MetricTargetAverageUtilization":35,"slbId":"lb-xxx","slbProject":"aliyun-fc-cn-hangzhou-d95881d9-5d3c-5f26-a6b8-************","slbLogstore":"function-log","vport":"80"}],"scaleUpRules":{"step":"100","disabled":false,"stabilizationWindowSeconds":0},"scaleDownRules":{"step":"100","disabled":false,"stabilizationWindowSeconds":300}}
	ScalingRuleMetric *string `json:"ScalingRuleMetric,omitempty" xml:"ScalingRuleMetric,omitempty"`
	// The custom name of the elastic scaling policy. The policy name must be unique within the application, start with a lowercase letter, and contain only lowercase letters, digits, and hyphens (-). It can be up to 32 characters in length.
	//
	// > After the elastic scaling policy is created, the policy name cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// timer-0800-2100
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	// The configuration of the scheduled elastic policy. This parameter is required when you select a scheduled elastic policy or use an SDK for configuration.
	//
	// Metric description:
	//
	// - **beginDate*	- and **endDate**: **beginDate*	- specifies the start date, and **endDate*	- specifies the end date. These parameters are used to configure the validity period of the scheduled elastic scaling policy. Valid values:
	//
	//     - If both values are **null**, the policy is executed on a long-term basis. This is the default value.
	//
	//     - If specific dates are specified, for example, **beginDate*	- is **2021-03-25*	- and **endDate*	- is **2021-04-25**, the validity period is one month.
	//
	// - **period**: the epoch for executing the scheduled elastic scaling policy. Valid values:
	//
	//     - **\\	- \\	- \\***: executes the scheduled policy at the specified time every day.
	//
	//     - **\\	- \\	- Fri,Mon**: executes the scheduled policy at the specified time on the specified days of every week. Multiple days are supported. The time zone is GMT+8. Valid values:
	//
	//         - **Sun**: Sunday
	//
	//         - **Mon**: Monday
	//
	//         - **Tue**: Tuesday
	//
	//         - **Wed**: Wednesday
	//
	//         - **Thu**: Thursday
	//
	//         - **Fri**: Friday
	//
	//         - **Sat**: Saturday
	//
	//     - **1,2,3,28,31 \\	- \\***: executes the scheduled policy at the specified time on the specified days of every month. Multiple days are supported. Valid values: [1,31\\]. If the current month does not have the 31st day, the scheduled policy is skipped on that day.
	//
	// - **schedules**: the trigger time of the elastic scaling policy and the number of application instances to maintain during the time period. Up to 20 time points are supported. Metric description:
	//
	//     - **atTime**: the trigger time point. The format is **HH:mm**, for example, **08:00**.
	//
	//     - **targetReplicas**: the number of instances for the application, or the minimum number of surviving instances for each deployment. Valid values: [1,50\\].
	//
	//         > We recommend that you set the minimum number of surviving instances for each rolling deployment to a value greater than or equal to **1*	- to ensure that business is not interrupted. If you set this parameter to **0**, business will be interrupted during the application upgrade procedure.
	//
	// example:
	//
	// {"beginDate":null,"endDate":null,"period":"	- 	- *","schedules":[{"atTime":"08:00","targetReplicas":10},{"atTime":"20:00","targetReplicas":3}]}
	ScalingRuleTimer *string `json:"ScalingRuleTimer,omitempty" xml:"ScalingRuleTimer,omitempty"`
	// The type of the elastic scaling policy. Valid values:
	//
	// - **timing**: scheduled elastic scaling.
	//
	// - **metric**: monitoring metric-based elastic scaling.
	//
	// - **mix**: hybrid elastic scaling.
	//
	// >
	//
	// > - If you select scheduled elastic scaling, the ScalingRuleTimer parameter is required.
	//
	// > - If you select monitoring metric-based elastic scaling, the ScalingRuleMetric parameter is required.
	//
	// > - If you select hybrid elastic scaling, the ScalingRuleMetric parameter is required. You can also specify the ScalingRuleTimer and ScalingRuleMetric parameters as needed.
	//
	// This parameter is required.
	//
	// example:
	//
	// timing
	ScalingRuleType *string `json:"ScalingRuleType,omitempty" xml:"ScalingRuleType,omitempty"`
}

func (s CreateApplicationScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateApplicationScalingRuleRequest) GetEnableIdle() *bool {
	return s.EnableIdle
}

func (s *CreateApplicationScalingRuleRequest) GetMinReadyInstanceRatio() *int32 {
	return s.MinReadyInstanceRatio
}

func (s *CreateApplicationScalingRuleRequest) GetMinReadyInstances() *int32 {
	return s.MinReadyInstances
}

func (s *CreateApplicationScalingRuleRequest) GetScalingRuleEnable() *bool {
	return s.ScalingRuleEnable
}

func (s *CreateApplicationScalingRuleRequest) GetScalingRuleMetric() *string {
	return s.ScalingRuleMetric
}

func (s *CreateApplicationScalingRuleRequest) GetScalingRuleName() *string {
	return s.ScalingRuleName
}

func (s *CreateApplicationScalingRuleRequest) GetScalingRuleTimer() *string {
	return s.ScalingRuleTimer
}

func (s *CreateApplicationScalingRuleRequest) GetScalingRuleType() *string {
	return s.ScalingRuleType
}

func (s *CreateApplicationScalingRuleRequest) SetAppId(v string) *CreateApplicationScalingRuleRequest {
	s.AppId = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetEnableIdle(v bool) *CreateApplicationScalingRuleRequest {
	s.EnableIdle = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetMinReadyInstanceRatio(v int32) *CreateApplicationScalingRuleRequest {
	s.MinReadyInstanceRatio = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetMinReadyInstances(v int32) *CreateApplicationScalingRuleRequest {
	s.MinReadyInstances = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetScalingRuleEnable(v bool) *CreateApplicationScalingRuleRequest {
	s.ScalingRuleEnable = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetScalingRuleMetric(v string) *CreateApplicationScalingRuleRequest {
	s.ScalingRuleMetric = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetScalingRuleName(v string) *CreateApplicationScalingRuleRequest {
	s.ScalingRuleName = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetScalingRuleTimer(v string) *CreateApplicationScalingRuleRequest {
	s.ScalingRuleTimer = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) SetScalingRuleType(v string) *CreateApplicationScalingRuleRequest {
	s.ScalingRuleType = &v
	return s
}

func (s *CreateApplicationScalingRuleRequest) Validate() error {
	return dara.Validate(s)
}
