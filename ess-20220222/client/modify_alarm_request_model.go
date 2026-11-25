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
	// The list of unique identifiers of the scaling rules associated with the event-triggered task.
	AlarmActions []*string `json:"AlarmActions,omitempty" xml:"AlarmActions,omitempty" type:"Repeated"`
	// The ID of the event-triggered task.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp1hvbnmkl10vll5****_f95ce797-dc2e-4bad-9618-14fee7d1****
	AlarmTaskId *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	// The comparison operator between the statistical value and the threshold of the metric. The comparison operator indicates the relationship in which the statistical value and the threshold can satisfy the condition. Valid values:
	//
	// - The statistical value is greater than or equal to the threshold. Value: `>=`.
	//
	// - The statistical value is less than or equal to the threshold. Value: `<=`.
	//
	// - The statistical value is greater than the threshold. Value: `>`.
	//
	// - The statistical value is less than the threshold. Value: `<`.
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
	Dimensions []*ModifyAlarmRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// Specifies the effective period of the event-triggered task.
	//
	// This parameter follows cron expressions. The default format is `X X X X X ?`, where:
	//
	// - X: a placeholder for a field, which sequentially represents seconds, minutes, hours, day, and month. X can be a specific value or a special character with logical meaning. For more information about the valid values of X, see [Cron expression](https://help.aliyun.com/document_detail/25907.html).
	//
	// - ?: specifies no value.
	//
	// > The value of this parameter is **in the UTC+8 time zone by default**. You can add time zone information `TZ=+yy` before the cron expression to specify a time zone, where y represents the numeric value of the time zone. For example, `TZ=+00 	- 	- 1-2 	- 	- ?` indicates that the event-triggered task takes effect between 01:00 and 02:59 every day in the UTC+0 time zone.
	//
	// Examples and descriptions:
	//
	// -  `	- 	- 	- 	- 	- ?  `: takes effect at all times.
	//
	// -  `	- 	- 17-18 	- 	- ? `: takes effect between 17:00 and 18:59 every day in the UTC+8 time zone.
	//
	// - `TZ=+00 	- 	- 1-2 	- 	- ?`: takes effect between 01:00 and 02:59 every day in the UTC+0 time zone.
	//
	// example:
	//
	// TZ=+00 	- 	- 1-2 	- 	- ?
	Effective *string `json:"Effective,omitempty" xml:"Effective,omitempty"`
	// The number of times that the threshold expression must be met before the scaling rule is triggered. For example, the average CPU utilization must be greater than or equal to 80% in three consecutive statistical results.
	//
	// example:
	//
	// 3
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The monitoring metrics of the multi-metric event-triggered task.
	Expressions []*ModifyAlarmRequestExpressions `json:"Expressions,omitempty" xml:"Expressions,omitempty" type:"Repeated"`
	// The relationship between multiple monitoring metrics conditions. Valid values:
	//
	// - `&&`: An alert is triggered only when all monitoring metrics meet the alert conditions. An alert is triggered only when all expressions in the multi-metric event-triggered task evaluate to `true`.
	//
	// - `||`: An alert is triggered when any monitoring metric meets the alert condition.
	//
	// Default value: `&&`.
	//
	// example:
	//
	// &&
	ExpressionsLogicOperator *string `json:"ExpressionsLogicOperator,omitempty" xml:"ExpressionsLogicOperator,omitempty"`
	// The ID of the CloudMonitor application group to which the custom metric belongs. This parameter is required only when MetricType is set to custom.
	//
	// example:
	//
	// 4055401
	GroupId *int32 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the metric. Valid values vary based on the monitoring type of the event-triggered task:
	//
	// - If the event-triggered task is of the custom monitoring type (MetricType=custom), valid values are your custom metrics.
	//
	// - If the event-triggered task is of the system monitoring type (MetricType=system), valid values:
	//
	//     - CpuUtilization: (ECS) CPU utilization (%).
	//
	//     - IntranetTx: (ECS) outbound traffic over the internal network (KByte/min).
	//
	//     - IntranetRx: (ECS) inbound traffic over the internal network (KByte/min).
	//
	//     - VpcInternetTx: (ECS) outbound traffic over the Internet (KByte/min).
	//
	//     - VpcInternetRx: (ECS) inbound traffic over the Internet (KByte/min).
	//
	//     - SystemDiskReadBps: (ECS) system cloud disk read BPS (Byte/s).
	//
	//     - SystemDiskWriteBps: (ECS) system cloud disk write BPS (Byte/s).
	//
	//     - SystemDiskReadOps: (ECS) system cloud disk read IOPS (count/s).
	//
	//     - SystemDiskWriteOps: (ECS) system cloud disk write IOPS (count/s).
	//
	//     - CpuUtilizationAgent: (Agent) CPU utilization (%).
	//
	//     - GpuMemoryFreeUtilizationAgent: (Agent) GPU memory idle rate (%).
	//
	//     - GpuMemoryUtilizationAgent: (Agent) GPU memory utilization (%).
	//
	//     - MemoryUtilization: (Agent) memory utilization (%).
	//
	//     - LoadAverage: (Agent) average system load.
	//
	//     - TcpConnection: (Agent) total number of TCP connections.
	//
	//     - TcpConnection: (Agent) number of established TCP connections.
	//
	//     - PackagesNetOut: (Agent) number of packets sent by the NIC.
	//
	//     - PackagesNetIn: (Agent) number of packets received by the NIC.
	//
	//     - EciPodCpuUtilization: CPU utilization (%).
	//
	//     - EciPodMemoryUtilization: memory utilization (%).
	//
	// For more information, see [Event-triggered tasks for system monitoring](https://help.aliyun.com/document_detail/74854.html).
	//
	// example:
	//
	// MemoryUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The type of the metric. Valid values:
	//
	// - system: uses CloudMonitor system metrics.
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
	// alarmtask****
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The statistical period for the metric data, in seconds. Valid values:
	//
	// - 15.
	//
	// - 60.
	//
	// - 120.
	//
	// - 300.
	//
	// - 900.
	//
	// > A statistical period of 15 seconds is applicable only to metrics collected by the CloudMonitor agent (metrics that start with Agent) when monitoring ECS-type scaling groups. For all other metrics, the statistical period can be set to 60, 120, 300, or 900 seconds.
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
	// The method used to aggregate the metric data. Valid values:
	//
	// - Average: average value.
	//
	// - Minimum: minimum value.
	//
	// - Maximum: maximum value.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold of the monitoring metrics in a multi-metric event-triggered task. The scaling rule is executed when the threshold expression is met the specified number of times.
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
	// The key of the dimension information associated with the metric. Valid values vary based on the monitoring type of the event-triggered node:
	//
	// - If the event-triggered task is of the custom monitoring type (MetricType=custom), valid values are not restricted. You can specify custom values.
	//
	//
	// - If the event-triggered task is of the system monitoring type (MetricType=system), valid values:
	//
	//     - user_id: your account ID.
	//
	//     - scaling_group: the scaling group to be monitored.
	//
	//     - device: the type of the network interface controller (NIC) device.
	//
	//     - state: the state of the TCP connection.
	//
	//     - rulePool: the server group specified for the ALB QPS monitoring metrics.
	//
	// example:
	//
	// device
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The value of the dimension information associated with the metric. Valid values are determined by Dimension.N.DimensionKey:
	//
	// - If Dimension.N.DimensionKey is set to a custom value, which means the event-triggered task is of the custom monitoring type (MetricType=custom), valid values are not restricted. You can specify custom values.
	//
	// - If Dimension.N.DimensionKey is set to a system-specified value, which means the event-triggered task is of the system monitoring type (MetricType=system), valid values:
	//
	//     - user_id: automatically populated by the system.
	//
	//     - scaling_group: automatically populated by the system.
	//
	//     - device:
	//
	//         - eth0: for classic network instances, eth0 represents the internal NIC. For VPC-connected instances, only the eth0 NIC exists.
	//
	//         - eth1: for classic network instances, eth1 represents the public NIC.
	//
	//     - state:
	//
	//         - TCP_TOTAL: the total number of TCP connections.
	//
	//         - ESTABLISHED: the number of established TCP connections.
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
	// The comparison operator between the statistical value and the threshold of the metric. The comparison operator indicates the relationship in which the statistical value and the threshold can satisfy the condition. Valid values:
	//
	// - The statistical value is greater than or equal to the threshold. Value: `>=`.
	//
	// - The statistical value is less than or equal to the threshold. Value: `<=`.
	//
	// - The statistical value is greater than the threshold. Value: `>`.
	//
	// - The statistical value is less than the threshold. Value: `<`.
	//
	// Default value: >=.
	//
	// example:
	//
	// >=
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The name of the monitoring metric in the multi-metric event-triggered node. Valid values vary based on the monitoring type of the event-triggered task:
	//
	// - If the event-triggered task is of the custom monitoring type (MetricType=custom), valid values are your custom monitoring metrics.
	//
	// - If the event-triggered task is of the system monitoring type (MetricType=system), valid values:
	//
	//     - CpuUtilization: (ECS) CPU utilization (%).
	//
	//     - IntranetTx: (ECS) outbound traffic over the internal network (KByte/min).
	//
	//     - IntranetRx: (ECS) inbound traffic over the internal network (KByte/min).
	//
	//     - VpcInternetTx: (ECS) outbound traffic over the Internet (KByte/min).
	//
	//     - VpcInternetRx: (ECS) inbound traffic over the Internet (KByte/min).
	//
	//     - SystemDiskReadBps: (ECS) system cloud disk read BPS (Byte/s).
	//
	//     - SystemDiskWriteBps: (ECS) system cloud disk write BPS (Byte/s).
	//
	//     - SystemDiskReadOps: (ECS) system cloud disk read IOPS (count/s).
	//
	//     - SystemDiskWriteOps: (ECS) system cloud disk write IOPS (count/s).
	//
	//     - CpuUtilizationAgent: (Agent) CPU utilization (%).
	//
	//     - GpuUtilizationAgent: (Agent) GPU utilization (%).
	//
	//     - GpuMemoryFreeUtilizationAgent: (Agent) GPU memory idle rate (%).
	//
	//     - GpuMemoryUtilizationAgent: (Agent) GPU memory utilization (%).
	//
	//     - MemoryUtilization: (Agent) memory utilization (%).
	//
	//     - LoadAverage: (Agent) average system load.
	//
	//     - TcpConnection: (Agent) total number of TCP connections.
	//
	//     - TcpConnection: (Agent) number of established TCP connections.
	//
	//     - PackagesNetOut: (Agent) number of packets sent by the network interface controller (NIC).
	//
	//     - PackagesNetIn: (Agent) number of packets received by the network interface controller (NIC).
	//
	//     - EciPodCpuUtilization: CPU utilization (%).
	//
	//     - EciPodMemoryUtilization: memory utilization (%).
	//
	// For more information, see [Event-triggered tasks for system monitoring](https://help.aliyun.com/document_detail/74854.html).
	//
	// example:
	//
	// CpuUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The statistical period for the metric data in the multi-metric event-triggered task, in seconds. Valid values:
	//
	// - 15.
	//
	// - 60.
	//
	// - 120.
	//
	// - 300.
	//
	// - 900.
	//
	// > A statistical period of 15 seconds is applicable only to metrics collected by the CloudMonitor agent (metrics that start with Agent) when monitoring ECS-type scaling groups. For all other metrics, the statistical period can be set to 60, 120, 300, or 900 seconds.
	//
	// Default value: 300.
	//
	// example:
	//
	// 300
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The method used to aggregate the metric data in the multi-metric event-triggered task. Valid values:
	//
	// - Average: average value.
	//
	// - Minimum: minimum value.
	//
	// - Maximum: maximum value.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold of the monitoring metrics in a multi-metric event-triggered task. The scaling rule is executed when the threshold expression is met the specified number of times.
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
