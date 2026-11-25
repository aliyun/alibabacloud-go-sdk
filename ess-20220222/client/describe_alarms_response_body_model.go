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
	// The list of event-triggered tasks.
	AlarmList []*DescribeAlarmsResponseBodyAlarmList `json:"AlarmList,omitempty" xml:"AlarmList,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
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
	if s.AlarmList != nil {
		for _, item := range s.AlarmList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAlarmsResponseBodyAlarmList struct {
	// The list of unique identifiers of the scaling rules associated with the event-triggered task.
	AlarmActions []*string `json:"AlarmActions,omitempty" xml:"AlarmActions,omitempty" type:"Repeated"`
	// The ID of the event-triggered task.
	//
	// example:
	//
	// asg-bp1hvbnmkl10vll5****_f95ce797-dc2e-4bad-9618-14fee7d1****
	AlarmTaskId *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	// The comparison operator used to compare the metric value with the threshold. Valid values:
	//
	// - >=: The metric value is greater than or equal to the threshold.
	//
	// - <=: The metric value is less than or equal to the threshold.
	//
	// - >: The metric value is greater than the threshold.
	//
	// - <: The metric value is less than the threshold.
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
	Dimensions []*DescribeAlarmsResponseBodyAlarmListDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The effective period of the event-triggered node.
	//
	// example:
	//
	// 	- 	- 	- 	- 	- ?
	Effective *string `json:"Effective,omitempty" xml:"Effective,omitempty"`
	// Indicates whether the event-triggered task is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The number of times that the threshold condition must be met before the scaling rule is triggered. For example, if this parameter is set to 3, the scaling rule is triggered only when the average CPU utilization is greater than or equal to 80% for three consecutive times.
	//
	// example:
	//
	// 3
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The information about the multi-metric alert rule.
	Expressions []*DescribeAlarmsResponseBodyAlarmListExpressions `json:"Expressions,omitempty" xml:"Expressions,omitempty" type:"Repeated"`
	// The logical operator that specifies the relationship among multiple monitoring alert conditions. Valid values:
	//
	// - `&&`: An alert is triggered only when all monitoring metrics meet the alert conditions. That is, an alert is triggered only when all expressions in the multi-metric alert rule evaluate to `true`.
	//
	// - `||`: An alert is triggered when any of the monitoring metrics meets the alert condition.
	//
	// example:
	//
	// &&
	ExpressionsLogicOperator *string `json:"ExpressionsLogicOperator,omitempty" xml:"ExpressionsLogicOperator,omitempty"`
	// The monitoring metrics settings of Hybrid Cloud Monitoring. For more information about how to configure the settings, see [Use formulas to customize target tracking scaling rules](https://help.aliyun.com/document_detail/2852162.html).
	HybridMetrics []*DescribeAlarmsResponseBodyAlarmListHybridMetrics `json:"HybridMetrics,omitempty" xml:"HybridMetrics,omitempty" type:"Repeated"`
	// The ID of the Hybrid Cloud Monitoring metrics repository.
	//
	// For more information about how to manage metrics repositories, see [Manage metrics repositories](https://help.aliyun.com/document_detail/217606.html).
	//
	// example:
	//
	// aliyun-test
	HybridMonitorNamespace *string `json:"HybridMonitorNamespace,omitempty" xml:"HybridMonitorNamespace,omitempty"`
	// The name of the metric. Valid values:
	//
	// - CpuUtilization: the CPU utilization of an ECS instance. Unit: %.
	//
	// - ConcurrentConnections: the number of concurrent connections to an ECS instance.
	//
	// - IntranetTx: the outbound internal traffic of an ECS instance. Unit: Kbit/min.
	//
	// - IntranetRx: the inbound internal traffic of an ECS instance. Unit: Kbit/min.
	//
	// - VpcInternetTx: the outbound Internet traffic of an ECS instance. Unit: Kbit/min.
	//
	// - VpcInternetRx: the inbound Internet traffic of an ECS instance. Unit: Kbit/min.
	//
	// - SystemDiskReadBps: the number of bytes read from the system cloud disk per second. Unit: Byte/s.
	//
	// - SystemDiskWriteBps: the number of bytes written to the system cloud disk per second. Unit: Byte/s.
	//
	// - SystemDiskReadOps: the number of read operations on the system cloud disk per second. Unit: IOPS.
	//
	// - SystemDiskWriteOps: the number of write operations on the system cloud disk per second. Unit: IOPS.
	//
	// - CpuUtilizationAgent: the CPU utilization collected by the CloudMonitor agent. Unit: %.
	//
	// - GpuUtilizationAgent: the GPU utilization collected by the CloudMonitor agent. Unit: %.
	//
	// - GpuMemoryFreeUtilizationAgent: the percentage of idle GPU memory collected by the CloudMonitor agent. Unit: %.
	//
	// - GpuMemoryUtilizationAgent: the GPU memory utilization collected by the CloudMonitor agent. Unit: %.
	//
	// - MemoryUtilization: the memory utilization collected by the CloudMonitor agent. Unit: %.
	//
	// - LoadAverage: the average system load collected by the CloudMonitor agent.
	//
	// - TcpConnection: the total number of TCP connections collected by the CloudMonitor agent.
	//
	// - TcpConnection: the number of established TCP connections collected by the CloudMonitor agent.
	//
	// - PackagesNetOut: the number of packets sent by the internal NIC per second collected by the CloudMonitor agent. Unit: packets/s.
	//
	// - PackagesNetIn: the number of packets received by the internal NIC per second collected by the CloudMonitor agent. Unit: packets/s.
	//
	// - PackagesNetOut: the number of packets sent by the public NIC per second collected by the CloudMonitor agent. Unit: packets/s.
	//
	// - PackagesNetIn: the number of packets received by the public NIC per second collected by the CloudMonitor agent. Unit: packets/s.
	//
	// - EciPodCpuUtilization: the CPU utilization. Unit: %.
	//
	// - EciPodMemoryUtilization: the memory utilization. Unit: %.
	//
	// - LoadBalancerRealServerAverageQps: the average QPS per backend server in an ALB server group. Unit: queries/s.
	//
	// For more information, see [Event-triggered tasks for system monitoring](https://help.aliyun.com/document_detail/74854.html).
	//
	// example:
	//
	// CpuUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The type of the metric. Valid values:
	//
	// - system: system metrics of CloudMonitor.
	//
	// - custom: custom metrics reported to CloudMonitor.
	//
	// - hybrid: metrics of Hybrid Cloud Monitoring.
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
	// The period at which metric data is collected. Unit: seconds. Valid values:
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
	// > A statistical period of 15 seconds is supported only for scaling groups that contain ECS instances.
	//
	// example:
	//
	// 300
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The PromQL query statement for Hybrid Cloud Monitoring.
	//
	// example:
	//
	// (avg(last_over_time(AliyunMnsnew_ActiveMessages{region=\\"cn-hangzhou\\",userId=\\"123456****\\",queue=\\"testQueue\\"}[900s])) by (userId))/(avg(last_over_time(AliyunEss_RunningInstanceCount{instanceId=\\"asg-bp1****\\"}[900s])) by (userId) != 0)
	PromQL *string `json:"PromQL,omitempty" xml:"PromQL,omitempty"`
	// The ID of the scaling group with which the event-triggered task is associated.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The state of the event-triggered task. Valid values:
	//
	// - ALARM: Alerting. The alert condition is met.
	//
	// - OK: Normal. The alert condition is not met.
	//
	// - INSUFFICIENT_DATA: Insufficient data. The data is insufficient to determine whether the alert condition is met.
	//
	// example:
	//
	// ALARM
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The statistical method used for metric data. Valid values:
	//
	// - Average: the average value.
	//
	// - Minimum: the minimum value.
	//
	// - Maximum: the maximum value.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold of the monitoring metrics. The scaling rule is triggered when the threshold expression is met for the specified number of times.
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
	if s.HybridMetrics != nil {
		for _, item := range s.HybridMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAlarmsResponseBodyAlarmListDimensions struct {
	// The key of the dimension associated with the metric. Valid values:
	//
	// - user_id: the ID of your account.
	//
	// - scaling_group: the monitored scaling group.
	//
	// - device: the type of the NIC device.
	//
	// - state: the state of the TCP connection.
	//
	// example:
	//
	// device
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The value of the dimension associated with the metric. The valid values vary based on the dimension key.
	//
	// - `user_id`: automatically populated by the system.
	//
	// - `scaling_group`: automatically populated by the system.
	//
	// - Valid values for `device`:
	//
	//     - eth0: For classic network instances, eth0 represents the internal NIC. For VPC-connected instances, only the eth0 NIC exists.
	//
	//     - eth1: For classic network instances, eth1 represents the public NIC.
	//
	// - Valid values for `state`:
	//
	//     - TCP_TOTAL: the total number of TCP connections.
	//
	//     - ESTABLISHED: the number of established TCP connections.
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
	// The comparison operator used to compare the metric value with the threshold in the multi-metric alert rule. Valid values:
	//
	// - >=: The metric value is greater than or equal to the threshold.
	//
	// - <=: The metric value is less than or equal to the threshold.
	//
	// - >: The metric value is greater than the threshold.
	//
	// - <: The metric value is less than the threshold.
	//
	// example:
	//
	// >=
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The name of the monitoring metrics in the multi-metric alert rule. Valid values:
	//
	// - CpuUtilization: the CPU utilization of an ECS instance. Unit: %.
	//
	// - ConcurrentConnections: the number of concurrent connections to an ECS instance.
	//
	// - IntranetTx: the outbound internal network traffic of an ECS instance. Unit: Kbit/min.
	//
	// - IntranetRx: the inbound internal network traffic of an ECS instance. Unit: Kbit/min.
	//
	// - VpcInternetTx: the outbound Internet traffic of an ECS instance. Unit: Kbit/min.
	//
	// - VpcInternetRx: the inbound Internet traffic of an ECS instance. Unit: Kbit/min.
	//
	// - SystemDiskReadBps: the number of bytes read from the system cloud disk per second. Unit: Byte/s.
	//
	// - SystemDiskWriteBps: the number of bytes written to the system cloud disk per second. Unit: Byte/s.
	//
	// - SystemDiskReadOps: the number of read operations on the system cloud disk per second. Unit: IOPS.
	//
	// - SystemDiskWriteOps: the number of write operations on the system cloud disk per second. Unit: IOPS.
	//
	// - CpuUtilizationAgent: the CPU utilization collected by the CloudMonitor agent. Unit: %.
	//
	// - GpuUtilizationAgent: the GPU utilization collected by the CloudMonitor agent. Unit: %.
	//
	// - GpuMemoryFreeUtilizationAgent: the percentage of idle GPU memory collected by the CloudMonitor agent. Unit: %.
	//
	// - GpuMemoryUtilizationAgent: the GPU memory utilization collected by the CloudMonitor agent. Unit: %.
	//
	// - MemoryUtilization: the memory utilization collected by the CloudMonitor agent. Unit: %.
	//
	// - LoadAverage: the average system load collected by the CloudMonitor agent.
	//
	// - TcpConnection: the total number of TCP connections collected by the CloudMonitor agent.
	//
	// - TcpConnection: the number of established TCP connections collected by the CloudMonitor agent.
	//
	// - PackagesNetOut: the number of packets sent by the internal network interface controller (NIC) per second collected by the CloudMonitor agent. Unit: packets/s.
	//
	// - PackagesNetIn: the number of packets received by the internal network interface controller (NIC) per second collected by the CloudMonitor agent. Unit: packets/s.
	//
	// - PackagesNetOut: the number of packets sent by the public network interface controller (NIC) per second collected by the CloudMonitor agent. Unit: packets/s.
	//
	// - PackagesNetIn: the number of packets received by the public network interface controller (NIC) per second collected by the CloudMonitor agent. Unit: packets/s.
	//
	// - EciPodCpuUtilization: the CPU utilization. Unit: %.
	//
	// - EciPodMemoryUtilization: the memory utilization. Unit: %.
	//
	// - LoadBalancerRealServerAverageQps: the average QPS per backend server in an ALB server group. Unit: queries/s.
	//
	// For more information, see [Event-triggered nodes for system monitoring](https://help.aliyun.com/document_detail/74854.html).
	//
	// example:
	//
	// CpuUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The period at which metric data is collected in the multi-metric alert rule. Unit: seconds. Valid values:
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
	// > A statistical period of 15 seconds is supported only when you monitor a scaling group that contains ECS instances and use metrics collected by the CloudMonitor agent (metrics whose names start with Agent).
	//
	// example:
	//
	// 900
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The statistical method used for metric data in the multi-metric alert rule. Valid values:
	//
	// - Average: the average value.
	//
	// - Minimum: the minimum value.
	//
	// - Maximum: the maximum value.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold of the monitoring metrics in the multi-metric alert rule. The scaling rule is triggered when the threshold expression is met for the specified number of times.
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
	// The dimensions of the metric. This is used to specify the resources monitored by the metric.
	Dimensions []*DescribeAlarmsResponseBodyAlarmListHybridMetricsDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The metric calculation expression for multiple Hybrid Cloud Monitoring metrics. The calculation result of the expression is used to trigger scaling activities. The calculation expression must comply with the Reverse Polish Notation (RPN) specification, and only the `+`, `-`, `*`, and `/` operators are supported.
	//
	// example:
	//
	// (a+b)/2
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The reference ID of the metric in the metric calculation expression.
	//
	// example:
	//
	// a
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the monitoring metrics in the Hybrid Cloud Monitoring metrics repository.
	//
	// example:
	//
	// AliyunSmq_NumberOfMessagesVisible
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The statistical method of the metric. Valid values:
	//
	// - Average: the average value of all data points in the specified period.
	//
	// - Minimum: the minimum value of all data points in the specified period.
	//
	// - Maximum: the maximum value of all data points in the specified period.
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
	if s.Dimensions != nil {
		for _, item := range s.Dimensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAlarmsResponseBodyAlarmListHybridMetricsDimensions struct {
	// The key of the metric dimension.
	//
	// example:
	//
	// queue
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The value of the metric dimension.
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
