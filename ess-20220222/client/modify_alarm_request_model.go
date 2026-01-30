// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmActions(v []*string) *ModifyAlarmRequest
	GetAlarmActions() []*string
	SetAlarmTaskId(v string) *ModifyAlarmRequest
	GetAlarmTaskId() *string
	SetComparisonOperator(v string) *ModifyAlarmRequest
	GetComparisonOperator() *string
	SetDescription(v string) *ModifyAlarmRequest
	GetDescription() *string
	SetDimensions(v []*ModifyAlarmRequestDimensions) *ModifyAlarmRequest
	GetDimensions() []*ModifyAlarmRequestDimensions
	SetEffective(v string) *ModifyAlarmRequest
	GetEffective() *string
	SetEvaluationCount(v int32) *ModifyAlarmRequest
	GetEvaluationCount() *int32
	SetExpressions(v []*ModifyAlarmRequestExpressions) *ModifyAlarmRequest
	GetExpressions() []*ModifyAlarmRequestExpressions
	SetExpressionsLogicOperator(v string) *ModifyAlarmRequest
	GetExpressionsLogicOperator() *string
	SetGroupId(v int32) *ModifyAlarmRequest
	GetGroupId() *int32
	SetMetricName(v string) *ModifyAlarmRequest
	GetMetricName() *string
	SetMetricType(v string) *ModifyAlarmRequest
	GetMetricType() *string
	SetName(v string) *ModifyAlarmRequest
	GetName() *string
	SetOwnerId(v int64) *ModifyAlarmRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *ModifyAlarmRequest
	GetPeriod() *int32
	SetRegionId(v string) *ModifyAlarmRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyAlarmRequest
	GetResourceOwnerAccount() *string
	SetStatistics(v string) *ModifyAlarmRequest
	GetStatistics() *string
	SetThreshold(v float32) *ModifyAlarmRequest
	GetThreshold() *float32
}

type ModifyAlarmRequest struct {
	// The list of unique identifiers of the scaling rules that are associated with the event-triggered task.
	AlarmActions []*string `json:"AlarmActions,omitempty" xml:"AlarmActions,omitempty" type:"Repeated"`
	// The ID of the event-triggered task.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp1hvbnmkl10vll5****_f95ce797-dc2e-4bad-9618-14fee7d1****
	AlarmTaskId *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	// The operator that is used to compare the metric value and the threshold. Valid values:
	//
	// 	- If the metric value is greater than or equal to the threshold, set the value to `>=`.
	//
	// 	- If the metric value is less than or equal to the threshold, set the value to `<=`.
	//
	// 	- If the metric value is greater than the threshold, set the value to `>`.
	//
	// 	- If the metric value is less than the threshold, set the value to `<`.
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
	// The metric dimensions.
	Dimensions []*ModifyAlarmRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The effective period of the event-triggered task.
	//
	// This parameter follows the cron expression format. The default format is `X X X X X ?`, in which:
	//
	// 	- X: a placeholder for a field, which represents seconds, minutes, hours, days, and months in sequence. X can be a definite value or a special character that has logical meaning. For information about the valid values of X, see [Cron expression](https://help.aliyun.com/document_detail/25907.html).
	//
	// 	- ?: No value is specified.
	//
	// > By default, the value of this parameter is specified in **UTC+8**. You can specify the time zone in the `TZ=+yy` format before a cron expression. y indicates the time zone. For example, `TZ=+00 	- 	- 1-2 	- 	- ?` specifies that the event-triggered task is in effect between 01:00 and 02:59 (UTC+0) every day.
	//
	// Examples:
	//
	// 	- ` 	- 	- 	- 	- 	- ?  `: The event-triggered task is in effect all the time.
	//
	// 	- ` 	- 	- 17-18 	- 	- ?  `: The event-triggered task is in effect between 17:00 and 18:59 (UTC+8) every day.
	//
	// 	- `TZ=+00 	- 	- 1-2 	- 	- ?`: The event-triggered task is in effect between 01:00 and 02:59 (UTC+0) every day.
	//
	// example:
	//
	// TZ=+00 	- 	- 1-2 	- 	- ?
	Effective *string `json:"Effective,omitempty" xml:"Effective,omitempty"`
	// The number of times that the threshold must be reached before a scaling rule can be executed. For example, if you set this parameter to 3, the average CPU utilization must reach or exceed 80% three times in a row before a scaling rule is triggered.
	//
	// example:
	//
	// 3
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The expressions that are specified in the multi-metric alert rule.
	Expressions []*ModifyAlarmRequestExpressions `json:"Expressions,omitempty" xml:"Expressions,omitempty" type:"Repeated"`
	// The relationship between the trigger conditions in the multi-metric alert rule. Valid values:
	//
	// 	- `&&`: An alert is triggered only if all metrics in the multi-metric alert rule meet the trigger conditions. In this case, an alert is triggered only if the results of all trigger conditions that are specified in the multi-metric alert rule are `true`.
	//
	// 	- `||`: An alert is triggered if one of the metrics in the multi-metric alert rule meets the trigger condition.
	//
	// Default value: `&&`.
	//
	// example:
	//
	// &&
	ExpressionsLogicOperator *string `json:"ExpressionsLogicOperator,omitempty" xml:"ExpressionsLogicOperator,omitempty"`
	// The ID of the application group to which the custom metric belongs. If you set the MetricType parameter to custom, you must specify this parameter.
	//
	// example:
	//
	// 4055401
	GroupId *int32 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the metric. The valid values vary based on the metric type.
	//
	// 	- If you set the MetricType parameter to custom, the valid values are your custom metrics.
	//
	// 	- If you set the MetricType parameter to system, the MetricName parameter has the following valid values:
	//
	//     	- CpuUtilization: the CPU utilization of an ECS instance. Unit: %.
	//
	//     	- IntranetTx: the outbound traffic over the internal network from an ECS instance. Unit: KB/min.
	//
	//     	- IntranetRx: the inbound traffic over the Internet to an ECS instance that resides in a virtual private cloud (VPC). Unit: KB/min.
	//
	//     	- VpcInternetTx: the outbound traffic over the Internet from an ECS instance that resides in a VPC. Unit: KB/min.
	//
	//     	- VpcInternetRx: the inbound traffic over the Internet to an ECS instance that resides in a VPC. Unit: KB/min.
	//
	//     	- SystemDiskReadBps: the number of bytes read from the system disk used by an ECS instance per second.
	//
	//     	- SystemDiskWriteBps: the number of bytes written to the system disk used by an ECS instance per second.
	//
	//     	- SystemDiskReadOps: the number of read operations on the system disk used by an ECS instance per second.
	//
	//     	- SystemDiskWriteOps: the number of write operations on the system disk used by an ECS instance per second.
	//
	//     	- CpuUtilizationAgent: the CPU utilization of an agent. Unit: %.
	//
	//     	- GpuMemoryFreeUtilizationAgent: the percentage of idle GPU memory of an agent.
	//
	//     	- GpuMemoryUtilizationAgent: the GPU memory usage of an agent. Unit: %.
	//
	//     	- MemoryUtilization: the memory usage of an agent. Unit: %.
	//
	//     	- LoadAverage: the average system load of an agent.
	//
	//     	- TcpConnection: the total number of TCP connections of an agent.
	//
	//     	- TcpConnection: the number of established TCP connections of an agent.
	//
	//     	- PackagesNetOut: the number of packets that are sent by the internal network interface controller (NIC) used by an agent.
	//
	//     	- PackagesNetIn: the number of packets that are received by the internal NIC used by an agent.
	//
	//     	- EciPodCpuUtilization: the CPU utilization of an elastic container instance. Unit: %.
	//
	//     	- EciPodMemoryUtilization: the memory usage of an elastic container instance. Unit: %.
	//
	// For more information, see [Event-triggered task for system monitoring](https://help.aliyun.com/document_detail/74854.html).
	//
	// example:
	//
	// MemoryUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The metric type. Valid Values:
	//
	// 	- system: a system metric of CloudMonitor.
	//
	// 	- custom: a custom metric that is reported to CloudMonitor.
	//
	// example:
	//
	// system
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The name of the event-triggered task.
	//
	// example:
	//
	// alarmtask****
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The period of time during which statistics about the metric is collected. Unit: seconds. Valid values:
	//
	// 	- 15
	//
	// 	- 60
	//
	// 	- 120
	//
	// 	- 300
	//
	// 	- 900
	//
	// > If your scaling group is of the ECS type and uses CloudMonitor metrics, you can set the Period parameter to 15. In other cases, you can set the Period parameter to 60, 120, 300, or 900. In most cases, the name of a CloudMonitor metric contains Agent.
	//
	// example:
	//
	// 300
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the event-triggered task.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The method that is used to aggregate statistics for the metric. Valid values:
	//
	// 	- Average
	//
	// 	- Minimum
	//
	// 	- Maximum
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The thresholds of the metric values in the multi-metric alert rule. If the thresholds are reached the specified number of times within the specified period, a scaling rule is executed.
	//
	// example:
	//
	// 80
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ModifyAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAlarmRequest) GoString() string {
	return s.String()
}

func (s *ModifyAlarmRequest) GetAlarmActions() []*string {
	return s.AlarmActions
}

func (s *ModifyAlarmRequest) GetAlarmTaskId() *string {
	return s.AlarmTaskId
}

func (s *ModifyAlarmRequest) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *ModifyAlarmRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyAlarmRequest) GetDimensions() []*ModifyAlarmRequestDimensions {
	return s.Dimensions
}

func (s *ModifyAlarmRequest) GetEffective() *string {
	return s.Effective
}

func (s *ModifyAlarmRequest) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *ModifyAlarmRequest) GetExpressions() []*ModifyAlarmRequestExpressions {
	return s.Expressions
}

func (s *ModifyAlarmRequest) GetExpressionsLogicOperator() *string {
	return s.ExpressionsLogicOperator
}

func (s *ModifyAlarmRequest) GetGroupId() *int32 {
	return s.GroupId
}

func (s *ModifyAlarmRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *ModifyAlarmRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *ModifyAlarmRequest) GetName() *string {
	return s.Name
}

func (s *ModifyAlarmRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAlarmRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *ModifyAlarmRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAlarmRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAlarmRequest) GetStatistics() *string {
	return s.Statistics
}

func (s *ModifyAlarmRequest) GetThreshold() *float32 {
	return s.Threshold
}

func (s *ModifyAlarmRequest) SetAlarmActions(v []*string) *ModifyAlarmRequest {
	s.AlarmActions = v
	return s
}

func (s *ModifyAlarmRequest) SetAlarmTaskId(v string) *ModifyAlarmRequest {
	s.AlarmTaskId = &v
	return s
}

func (s *ModifyAlarmRequest) SetComparisonOperator(v string) *ModifyAlarmRequest {
	s.ComparisonOperator = &v
	return s
}

func (s *ModifyAlarmRequest) SetDescription(v string) *ModifyAlarmRequest {
	s.Description = &v
	return s
}

func (s *ModifyAlarmRequest) SetDimensions(v []*ModifyAlarmRequestDimensions) *ModifyAlarmRequest {
	s.Dimensions = v
	return s
}

func (s *ModifyAlarmRequest) SetEffective(v string) *ModifyAlarmRequest {
	s.Effective = &v
	return s
}

func (s *ModifyAlarmRequest) SetEvaluationCount(v int32) *ModifyAlarmRequest {
	s.EvaluationCount = &v
	return s
}

func (s *ModifyAlarmRequest) SetExpressions(v []*ModifyAlarmRequestExpressions) *ModifyAlarmRequest {
	s.Expressions = v
	return s
}

func (s *ModifyAlarmRequest) SetExpressionsLogicOperator(v string) *ModifyAlarmRequest {
	s.ExpressionsLogicOperator = &v
	return s
}

func (s *ModifyAlarmRequest) SetGroupId(v int32) *ModifyAlarmRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyAlarmRequest) SetMetricName(v string) *ModifyAlarmRequest {
	s.MetricName = &v
	return s
}

func (s *ModifyAlarmRequest) SetMetricType(v string) *ModifyAlarmRequest {
	s.MetricType = &v
	return s
}

func (s *ModifyAlarmRequest) SetName(v string) *ModifyAlarmRequest {
	s.Name = &v
	return s
}

func (s *ModifyAlarmRequest) SetOwnerId(v int64) *ModifyAlarmRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAlarmRequest) SetPeriod(v int32) *ModifyAlarmRequest {
	s.Period = &v
	return s
}

func (s *ModifyAlarmRequest) SetRegionId(v string) *ModifyAlarmRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAlarmRequest) SetResourceOwnerAccount(v string) *ModifyAlarmRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAlarmRequest) SetStatistics(v string) *ModifyAlarmRequest {
	s.Statistics = &v
	return s
}

func (s *ModifyAlarmRequest) SetThreshold(v float32) *ModifyAlarmRequest {
	s.Threshold = &v
	return s
}

func (s *ModifyAlarmRequest) Validate() error {
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

type ModifyAlarmRequestDimensions struct {
	// The key of dimension N that you want to associate with the metric. The valid values of Dimension.N.DimensionKey vary based on the value of MetricType.
	//
	// 	- If you set MetricType to custom, you can specify this parameter based on your business requirements.
	//
	// 	- If you set MetricType to system, this parameter has the following valid values:
	//
	//     	- user_id: the ID of your Alibaba Cloud account.
	//
	//     	- scaling_group: the scaling group that is monitored by the event-triggered task.
	//
	//     	- device: the NIC type.
	//
	//     	- state: the status of the TCP connection.
	//
	//     	- rulePool: the specified server group for the ALB qps metric.
	//
	// example:
	//
	// device
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The dimension value of the metric. The valid values of this parameter vary based on the value of Dimensions.DimensionKey.
	//
	// 	- If you set MetricType to custom, you can specify this parameter based on your business requirements.
	//
	// 	- If you set MetricType to system, Dimension.N.DimensionValue has the following valid values:
	//
	//     	- user_id: The system specifies the value.
	//
	//     	- scaling_group: The system specifies the value.
	//
	//     	- device:
	//
	//         	- eth0: For classic network instances, eth0 indicates the internal network network interface controller. Only one eth0 NIC exists on each instance that resides in VPCs.
	//
	//         	- eth1: For classic network instances, eth1 represents the Internet network interface controller.
	//
	//     	- state:
	//
	//         	- TCP_TOTAL specifies the total number of TCP connections.
	//
	//         	- ESTABLISHED specifies the number of TCP connections that are established.
	//
	// example:
	//
	// eth0
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
}

func (s ModifyAlarmRequestDimensions) String() string {
	return dara.Prettify(s)
}

func (s ModifyAlarmRequestDimensions) GoString() string {
	return s.String()
}

func (s *ModifyAlarmRequestDimensions) GetDimensionKey() *string {
	return s.DimensionKey
}

func (s *ModifyAlarmRequestDimensions) GetDimensionValue() *string {
	return s.DimensionValue
}

func (s *ModifyAlarmRequestDimensions) SetDimensionKey(v string) *ModifyAlarmRequestDimensions {
	s.DimensionKey = &v
	return s
}

func (s *ModifyAlarmRequestDimensions) SetDimensionValue(v string) *ModifyAlarmRequestDimensions {
	s.DimensionValue = &v
	return s
}

func (s *ModifyAlarmRequestDimensions) Validate() error {
	return dara.Validate(s)
}

type ModifyAlarmRequestExpressions struct {
	// The operator that is used to compare the metric value and the threshold. Valid values:
	//
	// 	- If the metric value is greater than or equal to the threshold, set the value to `>=`.
	//
	// 	- If the metric value is less than or equal to the threshold, set the value to `<=`.
	//
	// 	- If the metric value is greater than the threshold, set the value to `>`.
	//
	// 	- If the metric value is less than the threshold, set the value to `<`.
	//
	// Default value: >=.
	//
	// example:
	//
	// >=
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The name of the metric that is specified in the multi-metric alert rule. The valid values vary based on the metric type.
	//
	// 	- If you set the MetricType parameter to custom, the valid values are your custom metrics.
	//
	// 	- If you set the MetricType parameter to system, the MetricName parameter has the following valid values:
	//
	//     	- CpuUtilization: the CPU utilization of an ECS instance. Unit: %.
	//
	//     	- IntranetTx: the outbound traffic over the internal network from an ECS instance. Unit: KB/min.
	//
	//     	- IntranetRx: the inbound traffic over the Internet to an ECS instance that resides in a VPC. Unit: KB/min.
	//
	//     	- VpcInternetTx: the outbound traffic over the Internet from an ECS instance that resides in a VPC. Unit: KB/min.
	//
	//     	- VpcInternetRx: the inbound traffic over the Internet to an ECS instance that resides in a VPC. Unit: KB/min.
	//
	//     	- SystemDiskReadBps: the number of bytes read from the system disk used by an ECS instance per second.
	//
	//     	- SystemDiskWriteBps: the number of bytes written to the system disk used by an ECS instance per second.
	//
	//     	- SystemDiskReadOps: the number of read operations on the system disk used by an ECS instance per second.
	//
	//     	- SystemDiskWriteOps: the number of write operations on the system disk used by an ECS instance per second.
	//
	//     	- CpuUtilizationAgent: the CPU utilization of an agent. Unit: %.
	//
	//     	- GpuUtilizationAgent: the GPU utilization of an agent. Unit: %.
	//
	//     	- GpuMemoryFreeUtilizationAgent: the percentage of idle GPU memory of an agent.
	//
	//     	- GpuMemoryUtilizationAgent: the GPU memory usage of an agent. Unit: %.
	//
	//     	- MemoryUtilization: the memory usage of an agent. Unit: %.
	//
	//     	- LoadAverage: the average system load of an agent.
	//
	//     	- TcpConnection: the total number of TCP connections of an agent.
	//
	//     	- TcpConnection: the number of established TCP connections of an agent.
	//
	//     	- PackagesNetOut: the number of packets that are sent by the internal NIC used by an agent.
	//
	//     	- PackagesNetIn: the number of packets that are received by the internal NIC used by an agent.
	//
	//     	- EciPodCpuUtilization: the CPU utilization of an elastic container instance. Unit: %.
	//
	//     	- EciPodMemoryUtilization: the memory usage of an elastic container instance. Unit: %.
	//
	// For more information, see [Event-triggered task for system monitoring](https://help.aliyun.com/document_detail/74854.html).
	//
	// example:
	//
	// CpuUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The period of time during which statistics about the metrics in the multi-metric alert rule is collected. Unit: seconds. Valid values:
	//
	// 	- 15
	//
	// 	- 60
	//
	// 	- 120
	//
	// 	- 300
	//
	// 	- 900
	//
	// > If your scaling group is of the ECS type and the event-triggered task associated with your scaling group monitors CloudMonitor metrics, you can set the Period parameter to 15. In other cases, you can set the Period parameter to 60, 120, 300, or 900. In most cases, the name of a CloudMonitor metric contains Agent.
	//
	// Default value: 300.
	//
	// example:
	//
	// 300
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The method that is used to aggregate statistics about the metrics that are specified in the multi-metric alert rule. Valid values:
	//
	// 	- Average
	//
	// 	- Minimum
	//
	// 	- Maximum
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The thresholds of the metric values in the multi-metric alert rule. If the thresholds are reached the specified number of times within the specified period, a scaling rule is executed.
	//
	// example:
	//
	// 40.0
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ModifyAlarmRequestExpressions) String() string {
	return dara.Prettify(s)
}

func (s ModifyAlarmRequestExpressions) GoString() string {
	return s.String()
}

func (s *ModifyAlarmRequestExpressions) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *ModifyAlarmRequestExpressions) GetMetricName() *string {
	return s.MetricName
}

func (s *ModifyAlarmRequestExpressions) GetPeriod() *int32 {
	return s.Period
}

func (s *ModifyAlarmRequestExpressions) GetStatistics() *string {
	return s.Statistics
}

func (s *ModifyAlarmRequestExpressions) GetThreshold() *float32 {
	return s.Threshold
}

func (s *ModifyAlarmRequestExpressions) SetComparisonOperator(v string) *ModifyAlarmRequestExpressions {
	s.ComparisonOperator = &v
	return s
}

func (s *ModifyAlarmRequestExpressions) SetMetricName(v string) *ModifyAlarmRequestExpressions {
	s.MetricName = &v
	return s
}

func (s *ModifyAlarmRequestExpressions) SetPeriod(v int32) *ModifyAlarmRequestExpressions {
	s.Period = &v
	return s
}

func (s *ModifyAlarmRequestExpressions) SetStatistics(v string) *ModifyAlarmRequestExpressions {
	s.Statistics = &v
	return s
}

func (s *ModifyAlarmRequestExpressions) SetThreshold(v float32) *ModifyAlarmRequestExpressions {
	s.Threshold = &v
	return s
}

func (s *ModifyAlarmRequestExpressions) Validate() error {
	return dara.Validate(s)
}
