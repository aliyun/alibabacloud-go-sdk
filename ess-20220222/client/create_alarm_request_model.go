// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmActions(v []*string) *CreateAlarmRequest
	GetAlarmActions() []*string
	SetComparisonOperator(v string) *CreateAlarmRequest
	GetComparisonOperator() *string
	SetDescription(v string) *CreateAlarmRequest
	GetDescription() *string
	SetDimensions(v []*CreateAlarmRequestDimensions) *CreateAlarmRequest
	GetDimensions() []*CreateAlarmRequestDimensions
	SetEffective(v string) *CreateAlarmRequest
	GetEffective() *string
	SetEvaluationCount(v int32) *CreateAlarmRequest
	GetEvaluationCount() *int32
	SetExpressions(v []*CreateAlarmRequestExpressions) *CreateAlarmRequest
	GetExpressions() []*CreateAlarmRequestExpressions
	SetExpressionsLogicOperator(v string) *CreateAlarmRequest
	GetExpressionsLogicOperator() *string
	SetGroupId(v int32) *CreateAlarmRequest
	GetGroupId() *int32
	SetMetricName(v string) *CreateAlarmRequest
	GetMetricName() *string
	SetMetricType(v string) *CreateAlarmRequest
	GetMetricType() *string
	SetName(v string) *CreateAlarmRequest
	GetName() *string
	SetOwnerId(v int64) *CreateAlarmRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *CreateAlarmRequest
	GetPeriod() *int32
	SetRegionId(v string) *CreateAlarmRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateAlarmRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *CreateAlarmRequest
	GetScalingGroupId() *string
	SetStatistics(v string) *CreateAlarmRequest
	GetStatistics() *string
	SetThreshold(v float32) *CreateAlarmRequest
	GetThreshold() *float32
}

type CreateAlarmRequest struct {
	// The list of unique identifiers of the scaling rules associated with the event-triggered task.
	AlarmActions []*string `json:"AlarmActions,omitempty" xml:"AlarmActions,omitempty" type:"Repeated"`
	// The comparison operator used to compare the aggregated metric value with the threshold. This operator specifies the condition under which the threshold is met. Valid values:
	//
	// - >=: The aggregated metric value is greater than or equal to the threshold.
	//
	// - <=: The aggregated metric value is less than or equal to the threshold.
	//
	// - >: The aggregated metric value is greater than the threshold.
	//
	// - <: The aggregated metric value is less than the threshold.
	//
	// Default value: >=.
	//
	// example:
	//
	// >=
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The description of the event-triggered task.
	//
	// example:
	//
	// Test alarm task.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The dimension information associated with the metric.
	Dimensions []*CreateAlarmRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// Specifies the effective period of the event-triggered task. By default, the task takes effect at all times.
	//
	// This parameter follows the cron expression format. The default format is `X X X X X ?`, which is described as follows:
	//
	// - X: a placeholder for a field, representing seconds, minutes, hours, day of month, and month in sequence. X can be a specific value or a special character with logical meaning. For more information about the valid values of X, see [Cron expressions](https://help.aliyun.com/document_detail/25907.html).
	//
	// - ?: indicates that no value is specified.
	//
	// > The specified value of this parameter **defaults to the UTC+8 time zone**. You can add the time zone information `TZ=+yy` before the cron expression to specify a time zone, where y represents the time zone offset. For example, `TZ=+00 	- 	- 1-2 	- 	- ?` indicates that the event-triggered task takes effect between 01:00 and 02:59 every day in the UTC+0 time zone.
	//
	// Examples and their meanings:
	//
	// - `	- 	- 	- 	- 	- ?`: takes effect at all times.
	//
	// - `	- 	- 17-18 	- 	- ?`: takes effect between 17:00 and 18:59 every day in the UTC+8 time zone.
	//
	// - `TZ=+00 	- 	- 1-2 	- 	- ?`: takes effect between 01:00 and 02:59 every day in the UTC+0 time zone.
	//
	// example:
	//
	// TZ=+00 	- 	- 1-2 	- 	- ?
	Effective *string `json:"Effective,omitempty" xml:"Effective,omitempty"`
	// The number of times that the threshold expression must be met to trigger a scaling rule. For example, if this parameter is set to 3, the scaling rule is triggered when the average CPU utilization is greater than or equal to 80% for three consecutive times.
	//
	// Default value: 3.
	//
	// example:
	//
	// 3
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The information about the multi-metric alert rule.
	Expressions []*CreateAlarmRequestExpressions `json:"Expressions,omitempty" xml:"Expressions,omitempty" type:"Repeated"`
	// The logical operator between multiple metric alert conditions. Valid values:
	//
	// - `&&`: An alert is triggered only when all monitoring metrics meet the alert conditions. That is, an alert is triggered only when all expressions in the multi-metric alert rule evaluate to `true`.
	//
	// - `||`: An alert is triggered if any one of the metrics meets the alert condition.
	//
	// Default value: `&&`.
	//
	// example:
	//
	// &&
	ExpressionsLogicOperator *string `json:"ExpressionsLogicOperator,omitempty" xml:"ExpressionsLogicOperator,omitempty"`
	// The ID of the CloudMonitor application group to which the custom metric belongs. This parameter is required only when the metric type is custom.
	//
	// example:
	//
	// 4055401
	GroupId *int32 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the metric. The valid values vary based on the monitoring type of the event-triggered task:
	//
	// - If the event-triggered task is of the custom monitoring type (MetricType=custom), the valid values are your custom metrics.
	//
	// - If the event-triggered task is of the system monitoring type (MetricType=system), valid values:
	//
	//     - CpuUtilization: the CPU utilization of an ECS instance (%).
	//
	//     - ConcurrentConnections: the number of concurrent connections to an ECS instance.
	//
	//     - IntranetTx: the outbound internal network traffic of an ECS instance (Kb/min).
	//
	//     - IntranetRx: the inbound internal network traffic of an ECS instance (Kb/min).
	//
	//     - VpcInternetTx: the outbound Internet traffic of an ECS instance (Kb/min).
	//
	//     - VpcInternetRx: the inbound Internet traffic of an ECS instance (Kb/min).
	//
	//     - SystemDiskReadBps: the read BPS of the system cloud disk (Byte/s).
	//
	//     - SystemDiskWriteBps: the write BPS of the system cloud disk (Byte/s).
	//
	//     - SystemDiskReadOps: the read IOPS of the system cloud disk (operations/s).
	//
	//     - SystemDiskWriteOps: the write IOPS of the system cloud disk (operations/s).
	//
	//     - CpuUtilizationAgent: the CPU utilization monitored by the agent (%).
	//
	//     - GpuUtilizationAgent: the GPU utilization monitored by the agent (%).
	//
	//     - GpuMemoryFreeUtilizationAgent: the percentage of idle GPU memory monitored by the agent (%).
	//
	//     - GpuMemoryUtilizationAgent: the GPU memory utilization monitored by the agent (%).
	//
	//     - MemoryUtilization: the memory utilization monitored by the agent (%).
	//
	//     - LoadAverage: the average system load monitored by the agent.
	//
	//     - TcpConnection: the total number of TCP connections monitored by the agent.
	//
	//
	//
	//     - TcpConnection: the number of established TCP connections monitored by the agent.
	//
	//     - PackagesNetOut: the number of packets sent by the internal NIC per second monitored by the agent (packets/s).
	//
	//     - PackagesNetIn: the number of packets received by the internal NIC per second monitored by the agent (packets/s).
	//
	//
	//
	//     - PackagesNetOut: the number of packets sent by the public NIC per second monitored by the agent (packets/s).
	//
	//     - PackagesNetIn: the number of packets received by the public NIC per second monitored by the agent (packets/s).
	//
	//     - EciPodCpuUtilization: the CPU utilization (%).
	//
	//     - EciPodMemoryUtilization: the memory utilization (%).
	//
	//     - LoadBalancerRealServerAverageQps: the average QPS per backend server in an ALB server group (queries/s).
	//
	// For more information, see [System monitoring event-triggered tasks](https://help.aliyun.com/document_detail/74854.html).
	//
	// example:
	//
	// CpuUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The type of the metric. Valid values:
	//
	// - system: uses system metrics of CloudMonitor.
	//
	// - custom: uses custom metrics reported to CloudMonitor.
	//
	// example:
	//
	// system
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The name of the event-triggered task.
	//
	// example:
	//
	// TestAlarmTask
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The epoch for aggregating monitoring metric data. Unit: seconds. Valid values:
	//
	// - 15
	//
	// - 60
	//
	// - 120
	//
	// - 300
	//
	// - 900
	//
	// > Only scaling groups of the ECS instance type support a 15-second statistical period.
	//
	// Default value: 300.
	//
	// example:
	//
	// 300
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group associated with the event-triggered task.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The method for aggregating metric data. Valid values:
	//
	// - Average: the average value.
	//
	// - Minimum: the minimum value.
	//
	// - Maximum: the maximum value.
	//
	// Default value: Average.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold of the monitoring metrics. When the monitoring threshold expression is met for a specified number of times, it triggers the execution of a scaling rule.
	//
	// example:
	//
	// 80.0
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlarmRequest) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequest) GetAlarmActions() []*string {
	return s.AlarmActions
}

func (s *CreateAlarmRequest) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *CreateAlarmRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAlarmRequest) GetDimensions() []*CreateAlarmRequestDimensions {
	return s.Dimensions
}

func (s *CreateAlarmRequest) GetEffective() *string {
	return s.Effective
}

func (s *CreateAlarmRequest) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *CreateAlarmRequest) GetExpressions() []*CreateAlarmRequestExpressions {
	return s.Expressions
}

func (s *CreateAlarmRequest) GetExpressionsLogicOperator() *string {
	return s.ExpressionsLogicOperator
}

func (s *CreateAlarmRequest) GetGroupId() *int32 {
	return s.GroupId
}

func (s *CreateAlarmRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *CreateAlarmRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *CreateAlarmRequest) GetName() *string {
	return s.Name
}

func (s *CreateAlarmRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAlarmRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateAlarmRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAlarmRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAlarmRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *CreateAlarmRequest) GetStatistics() *string {
	return s.Statistics
}

func (s *CreateAlarmRequest) GetThreshold() *float32 {
	return s.Threshold
}

func (s *CreateAlarmRequest) SetAlarmActions(v []*string) *CreateAlarmRequest {
	s.AlarmActions = v
	return s
}

func (s *CreateAlarmRequest) SetComparisonOperator(v string) *CreateAlarmRequest {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateAlarmRequest) SetDescription(v string) *CreateAlarmRequest {
	s.Description = &v
	return s
}

func (s *CreateAlarmRequest) SetDimensions(v []*CreateAlarmRequestDimensions) *CreateAlarmRequest {
	s.Dimensions = v
	return s
}

func (s *CreateAlarmRequest) SetEffective(v string) *CreateAlarmRequest {
	s.Effective = &v
	return s
}

func (s *CreateAlarmRequest) SetEvaluationCount(v int32) *CreateAlarmRequest {
	s.EvaluationCount = &v
	return s
}

func (s *CreateAlarmRequest) SetExpressions(v []*CreateAlarmRequestExpressions) *CreateAlarmRequest {
	s.Expressions = v
	return s
}

func (s *CreateAlarmRequest) SetExpressionsLogicOperator(v string) *CreateAlarmRequest {
	s.ExpressionsLogicOperator = &v
	return s
}

func (s *CreateAlarmRequest) SetGroupId(v int32) *CreateAlarmRequest {
	s.GroupId = &v
	return s
}

func (s *CreateAlarmRequest) SetMetricName(v string) *CreateAlarmRequest {
	s.MetricName = &v
	return s
}

func (s *CreateAlarmRequest) SetMetricType(v string) *CreateAlarmRequest {
	s.MetricType = &v
	return s
}

func (s *CreateAlarmRequest) SetName(v string) *CreateAlarmRequest {
	s.Name = &v
	return s
}

func (s *CreateAlarmRequest) SetOwnerId(v int64) *CreateAlarmRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAlarmRequest) SetPeriod(v int32) *CreateAlarmRequest {
	s.Period = &v
	return s
}

func (s *CreateAlarmRequest) SetRegionId(v string) *CreateAlarmRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAlarmRequest) SetResourceOwnerAccount(v string) *CreateAlarmRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAlarmRequest) SetScalingGroupId(v string) *CreateAlarmRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *CreateAlarmRequest) SetStatistics(v string) *CreateAlarmRequest {
	s.Statistics = &v
	return s
}

func (s *CreateAlarmRequest) SetThreshold(v float32) *CreateAlarmRequest {
	s.Threshold = &v
	return s
}

func (s *CreateAlarmRequest) Validate() error {
	if s.Dimensions != nil {
		for _, item := range s.Dimensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Expressions != nil {
		for _, item := range s.Expressions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAlarmRequestDimensions struct {
	// The key of the dimension information associated with the metric. Valid values vary based on the monitoring type of the event-triggered task:
	//
	// - If the event-triggered task is of the custom monitoring type (MetricType=custom), the valid values are not limited and can be specified by you.
	//
	// - If the event-triggered task is of the system monitoring type (MetricType=system), valid values:
	//
	//     - user_id: the ID of your account.
	//
	//     - scaling_group: the monitored scaling group.
	//
	//     - device: the type of the network interface controller (NIC) device.
	//
	//     - state: the state of the TCP connection.
	//
	//     - rulePool: the ALB server group specified for the QPS monitoring metrics.
	//
	// example:
	//
	// device
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The value of the dimension information associated with the metric. Valid values vary based on Dimensions.DimensionKey:
	//
	// - If Dimensions.DimensionKey is custom, which means the event-triggered task is of the custom monitoring type (MetricType=custom), the valid values are not limited and can be specified by you.
	//
	// - If Dimensions.DimensionKey is within the system-specified range, which means the event-triggered task is of the system monitoring type (MetricType=system), valid values:
	//
	//     - user_id: automatically populated by the system.
	//
	//     - scaling_group: automatically populated by the system.
	//
	//     - device:
	//
	//         - eth0: For classic network instances, eth0 represents the internal NIC. For VPC-connected instances, only the eth0 NIC exists.
	//
	//         - eth1: For classic network instances, eth1 represents the public NIC.
	//
	//     - state:
	//
	//         - TCP_TOTAL: the total number of TCP connections.
	//
	//         - ESTABLISHED: the number of established TCP connections.
	//
	//     - rulePool: the ID of the ALB server group. Example: sgp-xxxxx.
	//
	// example:
	//
	// eth0
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
}

func (s CreateAlarmRequestDimensions) String() string {
	return dara.Prettify(s)
}

func (s CreateAlarmRequestDimensions) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestDimensions) GetDimensionKey() *string {
	return s.DimensionKey
}

func (s *CreateAlarmRequestDimensions) GetDimensionValue() *string {
	return s.DimensionValue
}

func (s *CreateAlarmRequestDimensions) SetDimensionKey(v string) *CreateAlarmRequestDimensions {
	s.DimensionKey = &v
	return s
}

func (s *CreateAlarmRequestDimensions) SetDimensionValue(v string) *CreateAlarmRequestDimensions {
	s.DimensionValue = &v
	return s
}

func (s *CreateAlarmRequestDimensions) Validate() error {
	return dara.Validate(s)
}

type CreateAlarmRequestExpressions struct {
	// The comparison operator used to compare the aggregated metric value with the threshold for the multi-metric alert rule. This operator specifies the condition under which the threshold is met. Valid values:
	//
	// - >=: The aggregated metric value is greater than or equal to the threshold.
	//
	// - <=: The aggregated metric value is less than or equal to the threshold.
	//
	// - >: The aggregated metric value is greater than the threshold.
	//
	// - <: The aggregated metric value is less than the threshold.
	//
	// Default value: >=.
	//
	// example:
	//
	// >=
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The name of the monitoring metrics for the multi-metric alert rule. Valid values vary based on the monitoring type of the event-triggered task:
	//
	// - If the event-triggered task is of the custom monitoring type (MetricType=custom), the valid values are your custom metrics.
	//
	// - If the event-triggered task is of the system monitoring type (MetricType=system), valid values:
	//
	//
	//
	//     - CpuUtilization: the CPU utilization of an ECS instance (%).
	//
	//     - ConcurrentConnections: the number of concurrent connections to an ECS instance.
	//
	//     - IntranetTx: the outbound internal network traffic of an ECS instance (Kb/min).
	//
	//     - IntranetRx: the inbound internal network traffic of an ECS instance (Kb/min).
	//
	//     - VpcInternetTx: the outbound Internet traffic of an ECS instance (Kb/min).
	//
	//     - VpcInternetRx: the inbound Internet traffic of an ECS instance (Kb/min).
	//
	//     - SystemDiskReadBps: the read BPS of the system cloud disk (Byte/s).
	//
	//     - SystemDiskWriteBps: the write BPS of the system cloud disk (Byte/s).
	//
	//     - SystemDiskReadOps: the read IOPS of the system cloud disk (operations/s).
	//
	//     - SystemDiskWriteOps: the write IOPS of the system cloud disk (operations/s).
	//
	//     - CpuUtilizationAgent: the CPU utilization monitored by the agent (%).
	//
	//     - GpuUtilizationAgent: the GPU utilization monitored by the agent (%).
	//
	//     - GpuMemoryFreeUtilizationAgent: the percentage of idle GPU memory monitored by the agent (%).
	//
	//     - GpuMemoryUtilizationAgent: the GPU memory utilization monitored by the agent (%).
	//
	//     - MemoryUtilization: the memory utilization monitored by the agent (%).
	//
	//     - LoadAverage: the average system load monitored by the agent.
	//
	//     - TcpConnection: the total number of TCP connections monitored by the agent.
	//
	//
	//
	//     - TcpConnection: the number of established TCP connections monitored by the agent.
	//
	//     - PackagesNetOut: the number of packets sent by the internal network interface controller (NIC) per second monitored by the agent (packets/s).
	//
	//     - PackagesNetIn: the number of packets received by the internal network interface controller (NIC) per second monitored by the agent (packets/s).
	//
	//     - PackagesNetOut: the number of packets sent by the public network interface controller (NIC) per second monitored by the agent (packets/s).
	//
	//     - PackagesNetIn: the number of packets received by the public network interface controller (NIC) per second monitored by the agent (packets/s).
	//
	//
	//
	//     - EciPodCpuUtilization: the CPU utilization (%).
	//
	//     - EciPodMemoryUtilization: the memory utilization (%).
	//
	//     - LoadBalancerRealServerAverageQps: the average QPS per backend server in an ALB server group (queries/s).
	//
	// For more information, see [System monitoring event-triggered tasks](https://help.aliyun.com/document_detail/74854.html).
	//
	// example:
	//
	// CpuUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The epoch for aggregating monitoring metric data in the multi-metric alert rule. Unit: seconds. Valid values:
	//
	// - 15
	//
	// - 60
	//
	// - 120
	//
	// - 300
	//
	// - 900
	//
	// > Only scaling groups of the ECS instance type support a 15-second statistical period.
	//
	// Default value: 300.
	//
	// example:
	//
	// 900
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The method for aggregating metric data in the multi-metric alert rule. Valid values:
	//
	// - Average: the average value.
	//
	// - Minimum: the minimum value.
	//
	// - Maximum: the maximum value.
	//
	// Default value: Average.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold of the monitoring metrics in the multi-metric alert rule. When the monitoring threshold expression is met for a specified number of times, it triggers the execution of a scaling rule.
	//
	// example:
	//
	// 40.0
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateAlarmRequestExpressions) String() string {
	return dara.Prettify(s)
}

func (s CreateAlarmRequestExpressions) GoString() string {
	return s.String()
}

func (s *CreateAlarmRequestExpressions) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *CreateAlarmRequestExpressions) GetMetricName() *string {
	return s.MetricName
}

func (s *CreateAlarmRequestExpressions) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateAlarmRequestExpressions) GetStatistics() *string {
	return s.Statistics
}

func (s *CreateAlarmRequestExpressions) GetThreshold() *float32 {
	return s.Threshold
}

func (s *CreateAlarmRequestExpressions) SetComparisonOperator(v string) *CreateAlarmRequestExpressions {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateAlarmRequestExpressions) SetMetricName(v string) *CreateAlarmRequestExpressions {
	s.MetricName = &v
	return s
}

func (s *CreateAlarmRequestExpressions) SetPeriod(v int32) *CreateAlarmRequestExpressions {
	s.Period = &v
	return s
}

func (s *CreateAlarmRequestExpressions) SetStatistics(v string) *CreateAlarmRequestExpressions {
	s.Statistics = &v
	return s
}

func (s *CreateAlarmRequestExpressions) SetThreshold(v float32) *CreateAlarmRequestExpressions {
	s.Threshold = &v
	return s
}

func (s *CreateAlarmRequestExpressions) Validate() error {
	return dara.Validate(s)
}
