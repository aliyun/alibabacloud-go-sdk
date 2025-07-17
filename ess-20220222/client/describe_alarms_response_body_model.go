// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlarmsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmList(v []*DescribeAlarmsResponseBodyAlarmList) *DescribeAlarmsResponseBody
	GetAlarmList() []*DescribeAlarmsResponseBodyAlarmList
	SetPageNumber(v int32) *DescribeAlarmsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAlarmsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAlarmsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAlarmsResponseBody
	GetTotalCount() *int32
}

type DescribeAlarmsResponseBody struct {
	// The event-triggered tasks.
	AlarmList []*DescribeAlarmsResponseBodyAlarmList `json:"AlarmList,omitempty" xml:"AlarmList,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 871C7C53-34A4-45AA-8C14-4B72FA6A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of event-triggered tasks.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAlarmsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBody) GetAlarmList() []*DescribeAlarmsResponseBodyAlarmList {
	return s.AlarmList
}

func (s *DescribeAlarmsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAlarmsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAlarmsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlarmsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAlarmsResponseBody) SetAlarmList(v []*DescribeAlarmsResponseBodyAlarmList) *DescribeAlarmsResponseBody {
	s.AlarmList = v
	return s
}

func (s *DescribeAlarmsResponseBody) SetPageNumber(v int32) *DescribeAlarmsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetPageSize(v int32) *DescribeAlarmsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetRequestId(v string) *DescribeAlarmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlarmsResponseBody) SetTotalCount(v int32) *DescribeAlarmsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAlarmsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAlarmsResponseBodyAlarmList struct {
	// The unique identifiers of the scaling rules that are associated with the event-triggered task.
	AlarmActions []*string `json:"AlarmActions,omitempty" xml:"AlarmActions,omitempty" type:"Repeated"`
	// The ID of the event-triggered task.
	//
	// example:
	//
	// asg-bp1hvbnmkl10vll5****_f95ce797-dc2e-4bad-9618-14fee7d1****
	AlarmTaskId *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	// The operator that is used to compare the metric value and the metric threshold.
	//
	// 	- Valid value if the metric value is greater than or equal to the threshold: >=.
	//
	// 	- Valid value if the metric value is less than or equal to the threshold: <=.
	//
	// 	- Valid value if the metric value is greater than the threshold: >.
	//
	// 	- Valid value if the metric value is less than the threshold: <.
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
	Dimensions []*DescribeAlarmsResponseBodyAlarmListDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The effective period of the event-triggered task.
	//
	// example:
	//
	// Test
	Effective *string `json:"Effective,omitempty" xml:"Effective,omitempty"`
	// Indicates whether the event-triggered task feature is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The number of consecutive times that the threshold must be reached before a scaling rule is executed. For example, if you set this parameter to 3, the average CPU utilization must reach or exceed 80% three times in a row before a scaling rule is executed.
	//
	// example:
	//
	// 3
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The alert conditions of the multi-metric alert rule.
	Expressions []*DescribeAlarmsResponseBodyAlarmListExpressions `json:"Expressions,omitempty" xml:"Expressions,omitempty" type:"Repeated"`
	// The relationship between the trigger conditions that are specified in the multi-metric alert rule. Valid values:
	//
	// 	- `&&`: An alert is triggered only if all metrics in the multi-metric alert rule meet their trigger conditions. In this case, an alert is triggered only if the results of all trigger conditions that are specified in the multi-metric alert rule are `true`.
	//
	// 	- `||`: An alert is triggered only if one of the metrics in the multi-metric alert rule meets its trigger condition.
	//
	// example:
	//
	// &&
	ExpressionsLogicOperator *string `json:"ExpressionsLogicOperator,omitempty" xml:"ExpressionsLogicOperator,omitempty"`
	// The Hybrid Cloud Monitoring metrics.
	HybridMetrics []*DescribeAlarmsResponseBodyAlarmListHybridMetrics `json:"HybridMetrics,omitempty" xml:"HybridMetrics,omitempty" type:"Repeated"`
	// The ID of the Hybrid Cloud Monitoring namespace.
	//
	// For information about how to manage Hybrid Cloud Monitoring namespaces, see [Manage namespaces](https://help.aliyun.com/document_detail/217606.html).
	//
	// example:
	//
	// aliyun-test
	HybridMonitorNamespace *string `json:"HybridMonitorNamespace,omitempty" xml:"HybridMonitorNamespace,omitempty"`
	// The metric name. Valid values:
	//
	// 	- CpuUtilization: the CPU utilization of an Elastic Compute Service (ECS) instance. Unit: %.
	//
	// 	- ConcurrentConnections: the number of current connections to an ECS instance.
	//
	// 	- IntranetTx: the outbound traffic over an internal network. Unit: KB/min.
	//
	// 	- IntranetRx: the inbound traffic over an internal network. Unit: KB/min.
	//
	// 	- VpcInternetTx: the outbound traffic over a virtual private cloud (VPC). Unit: KB/min.
	//
	// 	- VpcInternetRx: the inbound traffic over a VPC. Unit: KB/min.
	//
	// 	- SystemDiskReadBps: the number of bytes read from the system disk per second.
	//
	// 	- SystemDiskWriteBps: the number of bytes written to the system disk per second.
	//
	// 	- SystemDiskReadOps: the read IOPS of the system disk. Unit: counts/s.
	//
	// 	- SystemDiskWriteOps: the write IOPS of the system disk. Unit: counts/s.
	//
	// 	- CpuUtilizationAgent: the CPU utilization. Unit: %.
	//
	// 	- GpuUtilizationAgent: the GPU utilization. Unit: %.
	//
	// 	- GpuMemoryFreeUtilizationAgent: the idle GPU memory usage. Unit: %.
	//
	// 	- GpuMemoryUtilizationAgent: the GPU memory usage. Unit: %.
	//
	// 	- MemoryUtilization: the memory usage. Unit: %.
	//
	// 	- LoadAverage: the average system load.
	//
	// 	- TcpConnection: the total number of TCP connections.
	//
	// 	- TcpConnection: the number of established TCP connections.
	//
	// 	- PackagesNetOut: the number of packets sent by the internal NIC. Unit: counts/s.
	//
	// 	- PackagesNetIn: the number of packets received by the internal NIC. Unit: counts/s.
	//
	// 	- PackagesNetOut: the number of packets sent by the public NIC. Unit: counts/s.
	//
	// 	- PackagesNetIn: the number of packets received by the public NIC. Unit: counts/s.
	//
	// 	- EciPodCpuUtilization: the CPU utilization. Unit: %.
	//
	// 	- EciPodMemoryUtilization: the memory usage. Unit: %.
	//
	// 	- LoadBalancerRealServerAverageQps: the queries per second (QPS) of an instance.
	//
	// For more information, see [Event-triggered tasks of the system monitoring type](https://help.aliyun.com/document_detail/74854.html).
	//
	// example:
	//
	// CpuUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The type of the metric. Valid values:
	//
	// 	- system: system metrics of CloudMonitor
	//
	// 	- custom: custom metrics that are reported to CloudMonitor.
	//
	// 	- hybrid: metrics of Hybrid Cloud Monitoring.
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
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// >  You can set the value of this parameter to 15 Seconds only for scaling groups of the ECS type.
	//
	// example:
	//
	// 300
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The PromQL statement of Hybrid Cloud Monitoring.
	//
	// example:
	//
	// (avg(last_over_time(AliyunMnsnew_ActiveMessages{region=\\"cn-hangzhou\\",userId=\\"123456****\\",queue=\\"testQueue\\"}[900s])) by (userId))/(avg(last_over_time(AliyunEss_RunningInstanceCount{instanceId=\\"asg-bp1****\\"}[900s])) by (userId) != 0)
	PromQL *string `json:"PromQL,omitempty" xml:"PromQL,omitempty"`
	// The ID of the scaling group to which the event-triggered task is associated.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The status of the event-triggered task. Valid values:
	//
	// 	- ALARM: The alert condition is met and an alert is triggered.
	//
	// 	- OK: The alert condition is not met.
	//
	// 	- INSUFFICIENT_DATA: Auto Scaling cannot determine whether the alert condition is met due to insufficient data.
	//
	// example:
	//
	// ALARM
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The method that is used to aggregate the metric data. Valid values:
	//
	// 	- Average: the average value
	//
	// 	- Minimum: the minimum value
	//
	// 	- Maximum: the maximum value
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold of the metric. If the threshold is reached the specified number of times within the statistical period, a scaling rule is executed.
	//
	// example:
	//
	// 80.0
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeAlarmsResponseBodyAlarmList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmsResponseBodyAlarmList) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetAlarmActions() []*string {
	return s.AlarmActions
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetAlarmTaskId() *string {
	return s.AlarmTaskId
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetDescription() *string {
	return s.Description
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetDimensions() []*DescribeAlarmsResponseBodyAlarmListDimensions {
	return s.Dimensions
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetEffective() *string {
	return s.Effective
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetExpressions() []*DescribeAlarmsResponseBodyAlarmListExpressions {
	return s.Expressions
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetExpressionsLogicOperator() *string {
	return s.ExpressionsLogicOperator
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetHybridMetrics() []*DescribeAlarmsResponseBodyAlarmListHybridMetrics {
	return s.HybridMetrics
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetHybridMonitorNamespace() *string {
	return s.HybridMonitorNamespace
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetName() *string {
	return s.Name
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetPromQL() *string {
	return s.PromQL
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetState() *string {
	return s.State
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeAlarmsResponseBodyAlarmList) GetThreshold() *float32 {
	return s.Threshold
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetAlarmActions(v []*string) *DescribeAlarmsResponseBodyAlarmList {
	s.AlarmActions = v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetAlarmTaskId(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.AlarmTaskId = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetComparisonOperator(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetDescription(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.Description = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetDimensions(v []*DescribeAlarmsResponseBodyAlarmListDimensions) *DescribeAlarmsResponseBodyAlarmList {
	s.Dimensions = v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetEffective(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.Effective = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetEnable(v bool) *DescribeAlarmsResponseBodyAlarmList {
	s.Enable = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetEvaluationCount(v int32) *DescribeAlarmsResponseBodyAlarmList {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetExpressions(v []*DescribeAlarmsResponseBodyAlarmListExpressions) *DescribeAlarmsResponseBodyAlarmList {
	s.Expressions = v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetExpressionsLogicOperator(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.ExpressionsLogicOperator = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetHybridMetrics(v []*DescribeAlarmsResponseBodyAlarmListHybridMetrics) *DescribeAlarmsResponseBodyAlarmList {
	s.HybridMetrics = v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetHybridMonitorNamespace(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.HybridMonitorNamespace = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetMetricName(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.MetricName = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetMetricType(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.MetricType = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetName(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.Name = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetPeriod(v int32) *DescribeAlarmsResponseBodyAlarmList {
	s.Period = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetPromQL(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.PromQL = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetScalingGroupId(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetState(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.State = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetStatistics(v string) *DescribeAlarmsResponseBodyAlarmList {
	s.Statistics = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) SetThreshold(v float32) *DescribeAlarmsResponseBodyAlarmList {
	s.Threshold = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmList) Validate() error {
	return dara.Validate(s)
}

type DescribeAlarmsResponseBodyAlarmListDimensions struct {
	// The dimension key of the metric. Valid values:
	//
	// 	- user_id: the ID of your Alibaba Cloud account.
	//
	// 	- scaling_group: the scaling group that is monitored by the event-triggered task.
	//
	// 	- device: the NIC type.
	//
	// 	- state: the status of the TCP connection.
	//
	// example:
	//
	// device
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The dimension value of the metric. The value of DimensionValue varies based on the value of DimensionKey.
	//
	// 	- If you set DimensionKey to `user_id`, the system specifies the value of DimensionValue.
	//
	// 	- If you set DimensionKey to `scaling_group`, the system specifies the value of DimensionValue.
	//
	// 	- If you set DimensionKey to `device`, you can set DimensionValue to eth0 or eth1.
	//
	//     	- For instances of the classic network type, eth0 indicates the internal NIC. Only one eth0 NIC exists on each instance that resides in VPCs.
	//
	//     	- For instances of the classic network type, eth1 indicates the public NIC.
	//
	// 	- If you set DimensionKey to `state`, you can set DimensionValue to TCP_TOTAL or ESTABLISHED.
	//
	//     	- TCP_TOTAL indicates the total number of TCP connections.
	//
	//     	- ESTABLISHED indicates the number of TCP connections that are established.
	//
	// example:
	//
	// eth0
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
}

func (s DescribeAlarmsResponseBodyAlarmListDimensions) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmsResponseBodyAlarmListDimensions) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyAlarmListDimensions) GetDimensionKey() *string {
	return s.DimensionKey
}

func (s *DescribeAlarmsResponseBodyAlarmListDimensions) GetDimensionValue() *string {
	return s.DimensionValue
}

func (s *DescribeAlarmsResponseBodyAlarmListDimensions) SetDimensionKey(v string) *DescribeAlarmsResponseBodyAlarmListDimensions {
	s.DimensionKey = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListDimensions) SetDimensionValue(v string) *DescribeAlarmsResponseBodyAlarmListDimensions {
	s.DimensionValue = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListDimensions) Validate() error {
	return dara.Validate(s)
}

type DescribeAlarmsResponseBodyAlarmListExpressions struct {
	// The operator that is used to compare the metric value and the threshold.
	//
	// 	- Valid value if the metric value is greater than or equal to the threshold: >=.
	//
	// 	- Valid value if the metric value is less than or equal to the threshold: <=.
	//
	// 	- Valid value if the metric value is greater than the threshold: >.
	//
	// 	- Valid value if the metric value is less than the threshold: <.
	//
	// example:
	//
	// >=
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The name of the metric that is specified in the multi-metric alert rule. Valid values:
	//
	// 	- CpuUtilization: the CPU utilization of an ECS instance. Unit: %.
	//
	// 	- ConcurrentConnections: the number of current connections to an ECS instance.
	//
	// 	- IntranetTx: the outbound traffic over an internal network. Unit: KB/min.
	//
	// 	- IntranetRx: the inbound traffic over an internal network. Unit: KB/min.
	//
	// 	- VpcInternetTx: the outbound traffic over a VPC. Unit: KB/min.
	//
	// 	- VpcInternetRx: the inbound traffic over a VPC. Unit: KB/min.
	//
	// 	- SystemDiskReadBps: the number of bytes read from the system disk per second.
	//
	// 	- SystemDiskWriteBps: the number of bytes written to the system disk per second.
	//
	// 	- SystemDiskReadOps: the read IOPS of the system disk. Unit: counts/s.
	//
	// 	- SystemDiskWriteOps: the write IOPS of the system disk. Unit: counts/s.
	//
	// 	- CpuUtilizationAgent: the CPU utilization. Unit: %.
	//
	// 	- GpuUtilizationAgent: the GPU utilization. Unit: %.
	//
	// 	- GpuMemoryFreeUtilizationAgent: the idle GPU memory usage. Unit: %.
	//
	// 	- GpuMemoryUtilizationAgent: the GPU memory usage. Unit: %.
	//
	// 	- MemoryUtilization: the memory usage. Unit: %.
	//
	// 	- LoadAverage: the average system load.
	//
	// 	- TcpConnection: the total number of TCP connections.
	//
	// 	- TcpConnection: the number of established TCP connections.
	//
	// 	- PackagesNetOut: the number of packets sent by the internal NIC. Unit: counts/s.
	//
	// 	- PackagesNetIn: the number of packets received by the internal NIC. Unit: counts/s.
	//
	// 	- PackagesNetOut: the number of packets sent by the public NIC. Unit: counts/s.
	//
	// 	- PackagesNetIn: the number of packets received by the public NIC. Unit: counts/s.
	//
	// 	- EciPodCpuUtilization: the CPU utilization. Unit: %.
	//
	// 	- EciPodMemoryUtilization: the memory usage. Unit: %.
	//
	// 	- LoadBalancerRealServerAverageQps: the QPS of an instance.
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
	// >  If your scaling group is of the ECS type and the event-triggered task that is associated with your scaling group monitors CloudMonitor metrics, you can set Period to 15. In most cases, the name of a CloudMonitor metric contains Agent.
	//
	// example:
	//
	// 900
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The method that is used to aggregate statistics about the metrics in the multi-metric alert rule. Valid values:
	//
	// 	- Average: the average value
	//
	// 	- Minimum: the minimum value
	//
	// 	- Maximum: the maximum value
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold of the metric value. If the threshold is reached the specified number of times within the specified period, a scaling rule is executed.
	//
	// example:
	//
	// 40.0
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeAlarmsResponseBodyAlarmListExpressions) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmsResponseBodyAlarmListExpressions) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyAlarmListExpressions) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeAlarmsResponseBodyAlarmListExpressions) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeAlarmsResponseBodyAlarmListExpressions) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeAlarmsResponseBodyAlarmListExpressions) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeAlarmsResponseBodyAlarmListExpressions) GetThreshold() *float32 {
	return s.Threshold
}

func (s *DescribeAlarmsResponseBodyAlarmListExpressions) SetComparisonOperator(v string) *DescribeAlarmsResponseBodyAlarmListExpressions {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListExpressions) SetMetricName(v string) *DescribeAlarmsResponseBodyAlarmListExpressions {
	s.MetricName = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListExpressions) SetPeriod(v int32) *DescribeAlarmsResponseBodyAlarmListExpressions {
	s.Period = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListExpressions) SetStatistics(v string) *DescribeAlarmsResponseBodyAlarmListExpressions {
	s.Statistics = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListExpressions) SetThreshold(v float32) *DescribeAlarmsResponseBodyAlarmListExpressions {
	s.Threshold = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListExpressions) Validate() error {
	return dara.Validate(s)
}

type DescribeAlarmsResponseBodyAlarmListHybridMetrics struct {
	// The metric dimensions. This parameter is used to specify the monitored resources.
	Dimensions []*DescribeAlarmsResponseBodyAlarmListHybridMetricsDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The metric expression that consists of multiple Hybrid Cloud Monitoring metrics. It calculates a result used to trigger scaling events.
	//
	// The expression is written in Reverse Polish Notation (RPN) format and supports only the following operators: `+, -, *, /`.
	//
	// example:
	//
	// (a+b)/2
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The reference ID of the metric in the metric expression.
	//
	// example:
	//
	// a
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the Hybrid Cloud Monitoring metric.
	//
	// example:
	//
	// AliyunSmq_NumberOfMessagesVisible
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The statistical method of the metric value. Valid values:
	//
	// 	- Average: The average value of all metric values within a specified interval is calculated.
	//
	// 	- Minimum: The minimum value of all metric values within a specified interval is calculated.
	//
	// 	- Maximum: The maximum value of all metric values within a specified interval is calculated.
	//
	// example:
	//
	// Average
	Statistic *string `json:"Statistic,omitempty" xml:"Statistic,omitempty"`
}

func (s DescribeAlarmsResponseBodyAlarmListHybridMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmsResponseBodyAlarmListHybridMetrics) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyAlarmListHybridMetrics) GetDimensions() []*DescribeAlarmsResponseBodyAlarmListHybridMetricsDimensions {
	return s.Dimensions
}

func (s *DescribeAlarmsResponseBodyAlarmListHybridMetrics) GetExpression() *string {
	return s.Expression
}

func (s *DescribeAlarmsResponseBodyAlarmListHybridMetrics) GetId() *string {
	return s.Id
}

func (s *DescribeAlarmsResponseBodyAlarmListHybridMetrics) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeAlarmsResponseBodyAlarmListHybridMetrics) GetStatistic() *string {
	return s.Statistic
}

func (s *DescribeAlarmsResponseBodyAlarmListHybridMetrics) SetDimensions(v []*DescribeAlarmsResponseBodyAlarmListHybridMetricsDimensions) *DescribeAlarmsResponseBodyAlarmListHybridMetrics {
	s.Dimensions = v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListHybridMetrics) SetExpression(v string) *DescribeAlarmsResponseBodyAlarmListHybridMetrics {
	s.Expression = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListHybridMetrics) SetId(v string) *DescribeAlarmsResponseBodyAlarmListHybridMetrics {
	s.Id = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListHybridMetrics) SetMetricName(v string) *DescribeAlarmsResponseBodyAlarmListHybridMetrics {
	s.MetricName = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListHybridMetrics) SetStatistic(v string) *DescribeAlarmsResponseBodyAlarmListHybridMetrics {
	s.Statistic = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListHybridMetrics) Validate() error {
	return dara.Validate(s)
}

type DescribeAlarmsResponseBodyAlarmListHybridMetricsDimensions struct {
	// The key of the metric dimension.
	//
	// example:
	//
	// queue
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The key of the metric dimension.
	//
	// example:
	//
	// testQueue
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
}

func (s DescribeAlarmsResponseBodyAlarmListHybridMetricsDimensions) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmsResponseBodyAlarmListHybridMetricsDimensions) GoString() string {
	return s.String()
}

func (s *DescribeAlarmsResponseBodyAlarmListHybridMetricsDimensions) GetDimensionKey() *string {
	return s.DimensionKey
}

func (s *DescribeAlarmsResponseBodyAlarmListHybridMetricsDimensions) GetDimensionValue() *string {
	return s.DimensionValue
}

func (s *DescribeAlarmsResponseBodyAlarmListHybridMetricsDimensions) SetDimensionKey(v string) *DescribeAlarmsResponseBodyAlarmListHybridMetricsDimensions {
	s.DimensionKey = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListHybridMetricsDimensions) SetDimensionValue(v string) *DescribeAlarmsResponseBodyAlarmListHybridMetricsDimensions {
	s.DimensionValue = &v
	return s
}

func (s *DescribeAlarmsResponseBodyAlarmListHybridMetricsDimensions) Validate() error {
	return dara.Validate(s)
}
