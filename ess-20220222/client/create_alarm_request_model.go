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
	// The list of unique identifiers of the scaling rules that are associated with the event-triggered task.
	AlarmActions []*string `json:"AlarmActions,omitempty" xml:"AlarmActions,omitempty" type:"Repeated"`
	// The operator that you want to use to compare the metric value and the threshold. Valid values:
	//
	// 	- If the metric value is greater than or equal to the threshold, set the value to >=.
	//
	// 	- If the metric value is less than or equal to the metric threshold, set the value to <=.
	//
	// 	- If the metric value is greater than the metric threshold, set the value to >.
	//
	// 	- If the metric value is less than the metric threshold, set the value to <.
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
	// The metric dimensions.
	Dimensions []*CreateAlarmRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The effective period of the event-triggered task. By default, the event-triggered task is in effect all the time.
	//
	// This parameter follows the cron expression format. The default format is `X X X X X ?`. In the format:
	//
	// 	- X: a placeholder for a field, which represents seconds, minutes, hours, days, and months in sequence. X can be a definite value or a special character that has logical meaning. For information about the valid values of X, see [Cron expression](https://help.aliyun.com/document_detail/25907.html).
	//
	// 	- ?: No value is specified.
	//
	// > By default, this parameter value is specified in **UTC+8**. You can specify the time zone in the `TZ=+yy` format before a cron expression. y indicates the time zone. For example, `TZ=+00 	- 	- 1-2 	- 	- ?` specifies that the event-triggered task is in effect between 01:00 and 02:59 (UTC+0) every day.
	//
	// Sample values:
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
	// The number of consecutive times that the threshold must be reached before a scaling rule is executed. For example, if you set this parameter to 3, the average CPU utilization must reach or exceed 80% three times in a row before the scaling rule is executed.
	//
	// Default value: 3.
	//
	// example:
	//
	// 3
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The information about the multi-metric alert rules.
	Expressions []*CreateAlarmRequestExpressions `json:"Expressions,omitempty" xml:"Expressions,omitempty" type:"Repeated"`
	// The relationship between the trigger conditions in the multi-metric alert rule. Valid values:
	//
	// 	- `&&`: An alert is triggered only if all metrics in the multi-metric alert rule meet the trigger conditions. In this case, an alert is triggered only if the results of all trigger conditions that are specified in the multi-metric alert rule are `true`.
	//
	// 	- `||`: An alert is triggered if one of the metrics in the multi-metric alert rule meets the trigger conditions.
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
	// The metric name. The valid values of this parameter vary based on the metric type.
	//
	// 	- If you set MetricType to custom, the valid values are the metrics that you have.
	//
	// 	- If you set MetricType to system, this parameter has the following valid values:
	//
	//     	- CpuUtilization: the CPU utilization. Unit: %.
	//
	//     	- ConcurrentConnections: the number of concurrent connections.
	//
	//     	- IntranetTx: the outbound traffic over an internal network. Unit: KB/min.
	//
	//     	- IntranetRx: the inbound traffic over an internal network. Unit: KB/min.
	//
	//     	- VpcInternetTx: the outbound traffic over a virtual private cloud (VPC). Unit: KB/min.
	//
	//     	- VpcInternetRx: the inbound traffic over a VPC. Unit: KB/min.
	//
	//     	- SystemDiskReadBps: the number of bytes read from the system disk per second.
	//
	//     	- SystemDiskWriteBps: the number of bytes written to the system disk per second.
	//
	//     	- SystemDiskReadOps: the read IOPS of the system disk. Unit: counts/s.
	//
	//     	- SystemDiskWriteOps: the write IOPS of the system disk. Unit: counts/s.
	//
	//     	- CpuUtilizationAgent: the CPU utilization. Unit: %.
	//
	//     	- GpuUtilizationAgent: the GPU utilization. Unit: %.
	//
	//     	- GpuMemoryFreeUtilizationAgent: the idle GPU memory usage. Unit: %.
	//
	//     	- GpuMemoryUtilizationAgent: the GPU memory usage. Unit: %.
	//
	//     	- MemoryUtilization: the memory usage. Unit: %.
	//
	//     	- LoadAverage: the average system load.
	//
	//     	- TcpConnection: the total number of TCP connections.
	//
	//     	- TcpConnection: the number of established TCP connections.
	//
	//     	- PackagesNetOut: the number of packets sent by the internal network interface controller (NIC). Unit: counts/s.
	//
	//     	- PackagesNetIn: the number of packets received by the internal NIC. Unit: counts/s.
	//
	//     	- PackagesNetOut: the number of packets sent by the public NIC. Unit: counts/s.
	//
	//     	- PackagesNetIn: the number of packets received by the public NIC. Unit: counts/s.
	//
	//     	- EciPodCpuUtilization: the CPU utilization. Unit: %.
	//
	//     	- EciPodMemoryUtilization: the memory usage. Unit: %.
	//
	//     	- LoadBalancerRealServerAverageQps: the queries per second (QPS) of an instance.
	//
	// For more information, see [Event-triggered tasks of the system monitoring type](https://help.aliyun.com/document_detail/74854.html).
	//
	// example:
	//
	// CpuUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The metric type. Valid values:
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
	// TestAlarmTask
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The statistical period of the metric data. Unit: seconds. Valid values:
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
	// >  You can set this parameter to 15 seconds only for scaling groups of the ECS type.
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
	// The scaling group ID of the event-triggered task.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The statistical method of the metric data. Valid values:
	//
	// 	- Average: calculates the average value of the metric data.
	//
	// 	- Minimum: calculates the minimum value of the metric data.
	//
	// 	- Maximum: calculates the maximum value of the metric data.
	//
	// Default value: Average.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold of the metric value. If the threshold is reached the specified number of times within the specified period, a scaling rule is executed.
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
	return dara.Validate(s)
}

type CreateAlarmRequestDimensions struct {
	// The dimension key of the metric. The valid values vary based on the metric type.
	//
	// 	- If you set MetricType to custom, you can specify this parameter based on your business requirements.
	//
	// 	- If you set MetricType to system, this parameter has the following valid values:
	//
	//     	- user_id: the ID of your Alibaba Cloud account.
	//
	//     	- scaling_group: the scaling group that you want to monitor by using the event-triggered task.
	//
	//     	- device: the NIC type.
	//
	//     	- state: the status of the TCP connection.
	//
	// example:
	//
	// device
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The dimension value of the metric. The valid values of this parameter vary based on the value of Dimensions.DimensionKey.
	//
	// 	- If you set MetricType to custom, you can specify this parameter based on your business requirements.
	//
	// 	- If you set MetricType to system, this parameter has the following valid values:
	//
	//     	- user_id: The system specifies the value.
	//
	//     	- scaling_group: The system specifies the value.
	//
	//     	- device: You can set this parameter to eth0 or eth1.
	//
	//         	- For instances of the classic network type, eth0 specifies the internal NIC. Only one eth0 NIC exists on each instance that resides in VPCs.
	//
	//         	- For instances of the classic network type, eth1 specifies the public NIC.
	//
	//     	- state: You can set this parameter to TCP_TOTAL or ESTABLISHED.
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
	// The operator that you want to use to compare the metric value and the threshold in the multi-metric alert rule. Valid values:
	//
	// 	- If the metric value is greater than or equal to the threshold, set the value to >=.
	//
	// 	- If the metric value is less than or equal to the metric threshold, set the value to <=.
	//
	// 	- If the metric value is greater than the metric threshold, set the value to >.
	//
	// 	- If the metric value is less than the metric threshold, set the value to <.
	//
	// Default value: >=.
	//
	// example:
	//
	// >=
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The names of the metrics in the multi-metric alert rule. The valid values of this parameter vary based on the metric type.
	//
	// 	- If you set MetricType to custom, the valid values are the metrics that you have.
	//
	// 	- If you set MetricType to system, this parameter has the following valid values:
	//
	//     	- CpuUtilization: the CPU utilization. Unit: %.
	//
	//     	- ConcurrentConnections: the number of concurrent connections.
	//
	//     	- IntranetTx: the outbound traffic over an internal network. Unit: KB/min.
	//
	//     	- IntranetRx: the inbound traffic over an internal network. Unit: KB/min.
	//
	//     	- VpcInternetTx: the outbound traffic over a VPC. Unit: KB/min.
	//
	//     	- VpcInternetRx: the inbound traffic over a VPC. Unit: KB/min.
	//
	//     	- SystemDiskReadBps: the number of bytes read from the system disk per second.
	//
	//     	- SystemDiskWriteBps: the number of bytes written to the system disk per second.
	//
	//     	- SystemDiskReadOps: the read IOPS of the system disk. Unit: counts/s.
	//
	//     	- SystemDiskWriteOps: the write IOPS of the system disk. Unit: counts/s.
	//
	//     	- CpuUtilizationAgent: the CPU utilization. Unit: %.
	//
	//     	- GpuUtilizationAgent: the GPU utilization. Unit: %.
	//
	//     	- GpuMemoryFreeUtilizationAgent: the idle GPU memory usage. Unit: %.
	//
	//     	- GpuMemoryUtilizationAgent: the GPU memory usage. Unit: %.
	//
	//     	- MemoryUtilization: the memory usage. Unit: %.
	//
	//     	- LoadAverage: the average system load.
	//
	//     	- TcpConnection: the total number of TCP connections.
	//
	//     	- TcpConnection: the number of established TCP connections.
	//
	//     	- PackagesNetOut: the number of packets sent by the internal NIC. Unit: counts/s.
	//
	//     	- PackagesNetIn: the number of packets received by the internal NIC. Unit: counts/s.
	//
	//     	- PackagesNetOut: the number of packets sent by the public NIC. Unit: counts/s.
	//
	//     	- PackagesNetIn: the number of packets received by the public NIC. Unit: counts/s.
	//
	//     	- EciPodCpuUtilization: the CPU utilization. Unit: %.
	//
	//     	- EciPodMemoryUtilization: the memory usage. Unit: %.
	//
	//     	- LoadBalancerRealServerAverageQps: the QPS of an instance.
	//
	// For more information, see [Event-triggered tasks of the system monitoring type](https://help.aliyun.com/document_detail/74854.html).
	//
	// example:
	//
	// CpuUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The statistical period of the metric data in the multi-metric alert rule. Unit: seconds. Valid values:
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
	// >  You can set this parameter to 15 seconds only for scaling groups of the ECS type.
	//
	// Default value: 300.
	//
	// example:
	//
	// 900
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The method that you want to use to aggregate the metric data in the multi-metric alert rule. Valid values:
	//
	// 	- Average: the average value.
	//
	// 	- Minimum: the minimum value
	//
	// 	- Maximum: the maximum value
	//
	// Default value: Average.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold of the metric value in the multi-metric alert rule. If the threshold is reached the specified number of times within the statistical period, a scaling rule is executed.
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
