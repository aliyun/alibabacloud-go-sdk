// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoscalingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoolDownDuration(v string) *CreateAutoscalingConfigRequest
	GetCoolDownDuration() *string
	SetDaemonsetEvictionForNodes(v bool) *CreateAutoscalingConfigRequest
	GetDaemonsetEvictionForNodes() *bool
	SetExpander(v string) *CreateAutoscalingConfigRequest
	GetExpander() *string
	SetGpuUtilizationThreshold(v string) *CreateAutoscalingConfigRequest
	GetGpuUtilizationThreshold() *string
	SetMaxGracefulTerminationSec(v int32) *CreateAutoscalingConfigRequest
	GetMaxGracefulTerminationSec() *int32
	SetMinReplicaCount(v int32) *CreateAutoscalingConfigRequest
	GetMinReplicaCount() *int32
	SetPriorities(v map[string][]*string) *CreateAutoscalingConfigRequest
	GetPriorities() map[string][]*string
	SetRecycleNodeDeletionEnabled(v bool) *CreateAutoscalingConfigRequest
	GetRecycleNodeDeletionEnabled() *bool
	SetScaleDownEnabled(v bool) *CreateAutoscalingConfigRequest
	GetScaleDownEnabled() *bool
	SetScaleUpFromZero(v bool) *CreateAutoscalingConfigRequest
	GetScaleUpFromZero() *bool
	SetScalerType(v string) *CreateAutoscalingConfigRequest
	GetScalerType() *string
	SetScanInterval(v string) *CreateAutoscalingConfigRequest
	GetScanInterval() *string
	SetSkipNodesWithLocalStorage(v bool) *CreateAutoscalingConfigRequest
	GetSkipNodesWithLocalStorage() *bool
	SetSkipNodesWithSystemPods(v bool) *CreateAutoscalingConfigRequest
	GetSkipNodesWithSystemPods() *bool
	SetUnneededDuration(v string) *CreateAutoscalingConfigRequest
	GetUnneededDuration() *string
	SetUtilizationThreshold(v string) *CreateAutoscalingConfigRequest
	GetUtilizationThreshold() *string
}

type CreateAutoscalingConfigRequest struct {
	// The scale-down trigger delay. The time interval between detecting a scale-down need (reaching the scale-down threshold) and actually performing the scale-down operation (reducing the number of Pods).
	//
	// Valid values: [1,60]. Unit: minutes.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	CoolDownDuration *string `json:"cool_down_duration,omitempty" xml:"cool_down_duration,omitempty"`
	// Specifies whether cluster-autoscaler evicts DaemonSet Pods on nodes during scale-down. Valid values:
	//
	// - `true`: DaemonSet Pods are evicted.
	//
	// - `false`: DaemonSet Pods are not evicted.
	//
	// example:
	//
	// false
	DaemonsetEvictionForNodes *bool `json:"daemonset_eviction_for_nodes,omitempty" xml:"daemonset_eviction_for_nodes,omitempty"`
	// The node pool scale-out order policy. Valid values:
	//
	// - `least-waste`: The default policy. If multiple node pools are available for scale-out, the node pool with the least resource waste is selected.
	//
	// - `random`: The random policy. If multiple node pools are available for scale-out, a random node pool is selected.
	//
	// - `priority`: The priority policy. If multiple node pools are available for scale-out, the node pool with the highest priority is selected based on the custom scaling group order you defined. Node pool priorities are defined by the `priorities` parameter.
	//
	// example:
	//
	// least-waste
	Expander *string `json:"expander,omitempty" xml:"expander,omitempty"`
	// The GPU scale-down threshold. The ratio of requested resources to total resources on a node.
	//
	// A GPU node can be scaled down only when this ratio falls below the configured threshold, meaning the CPU, memory, and GPU utilization of the node are all below the GPU scale-down threshold.
	//
	// Valid values: [0.1~1].
	//
	// Default value: 0.3, which indicates 30%.
	//
	// example:
	//
	// 0.3
	GpuUtilizationThreshold *string `json:"gpu_utilization_threshold,omitempty" xml:"gpu_utilization_threshold,omitempty"`
	// The timeout period that cluster-autoscaler waits for Pod termination during node draining in scale-down scenarios.
	//
	// Unit: seconds.
	//
	// Default value: 14400.
	//
	// example:
	//
	// 14400
	MaxGracefulTerminationSec *int32 `json:"max_graceful_termination_sec,omitempty" xml:"max_graceful_termination_sec,omitempty"`
	// The minimum number of Pods allowed in each ReplicaSet before a node can be scaled down.
	//
	// Default value: 0.
	//
	// example:
	//
	// 0
	MinReplicaCount *int32 `json:"min_replica_count,omitempty" xml:"min_replica_count,omitempty"`
	// The priority configuration for automatic scaling. After you create a node pool with auto scaling enabled, you can choose whether to configure a priority policy and priority settings by using [Enable node auto scaling](https://help.aliyun.com/document_detail/119099.html) to assign priorities to the scaling groups of specified auto scaling node pools.
	//
	// Valid values: [1, 100]. The value must be a positive integer. A larger value indicates a higher priority.
	Priorities map[string][]*string `json:"priorities,omitempty" xml:"priorities,omitempty"`
	// Specifies whether to delete the corresponding Kubernetes Node object after a node is successfully scaled down in swift mode. For more information about swift mode, see [Scaling modes](https://help.aliyun.com/document_detail/119099.html). Default value: false. Valid values:
	//
	// - `true`: The Kubernetes Node object is deleted after the node is stopped in swift mode. Setting this parameter to true is not recommended because it may cause Kubernetes object data inconsistency.
	//
	// - `false`: The Kubernetes Node object is retained after the node is stopped in swift mode.
	//
	// example:
	//
	// false
	RecycleNodeDeletionEnabled *bool `json:"recycle_node_deletion_enabled,omitempty" xml:"recycle_node_deletion_enabled,omitempty"`
	// Specifies whether to allow node scale-down. Valid values:
	//
	// - `true`: Scale-down is allowed.
	//
	// - `false`: Scale-down is not allowed.
	//
	// example:
	//
	// true
	ScaleDownEnabled *bool `json:"scale_down_enabled,omitempty" xml:"scale_down_enabled,omitempty"`
	// Specifies whether cluster-autoscaler performs scale-out when the number of Ready nodes in the cluster is 0. Default value: true. Valid values:
	//
	// - `true`: Scale-out is performed.
	//
	// - `false`: Scale-out is not performed.
	//
	// example:
	//
	// true
	ScaleUpFromZero *bool `json:"scale_up_from_zero,omitempty" xml:"scale_up_from_zero,omitempty"`
	// The type of the auto scaling component. For clusters of version 1.24 and later, the default value is goatscaler. For earlier versions, the default value is cluster-autoscaler. Valid values:
	//
	// - `goatscaler`: instant scaling.
	//
	// - `cluster-autoscaler`: automatic scaling.
	//
	// example:
	//
	// goatscaler
	ScalerType *string `json:"scaler_type,omitempty" xml:"scaler_type,omitempty"`
	// The scaling sensitivity, which adjusts the interval at which the system evaluates scaling decisions.
	//
	// Valid values: 15, 30, 60, 120, 180, and 300. Unit: seconds.
	//
	// Default value: 60.
	//
	// example:
	//
	// 30
	ScanInterval *string `json:"scan_interval,omitempty" xml:"scan_interval,omitempty"`
	// Specifies whether cluster-autoscaler skips scaling down nodes that run Pods with local storage (such as EmptyDir or HostPath). Valid values:
	//
	// - `true`: Nodes are not scaled down.
	//
	// - `false`: Nodes are scaled down.
	//
	// example:
	//
	// false
	SkipNodesWithLocalStorage *bool `json:"skip_nodes_with_local_storage,omitempty" xml:"skip_nodes_with_local_storage,omitempty"`
	// Specifies whether cluster-autoscaler skips scaling down nodes that run Pods in the kube-system namespace. This feature does not apply to DaemonSet Pods or Mirror Pods. Valid values:
	//
	// - `true`: Nodes are not scaled down.
	//
	// - `false`: Nodes are scaled down.
	//
	// example:
	//
	// true
	SkipNodesWithSystemPods *bool `json:"skip_nodes_with_system_pods,omitempty" xml:"skip_nodes_with_system_pods,omitempty"`
	// The cool-down period. The time interval after the most recent scale-out during which the auto scaling component does not perform scale-down operations. Nodes added during scale-out can only be evaluated for scale-down after the cool-down period expires.
	//
	// Valid values: [1,60]. Unit: minutes.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	UnneededDuration *string `json:"unneeded_duration,omitempty" xml:"unneeded_duration,omitempty"`
	// The scale-down threshold. The ratio of requested resources to total resources on a node.
	//
	// A node can be scaled down only when this ratio falls below the configured threshold, meaning both the CPU and memory resources utilization of the node are below the scale-down threshold.
	//
	// Valid values: [0.1~1].
	//
	// Default value: 0.5, which indicates 50%.
	//
	// example:
	//
	// 0.5
	UtilizationThreshold *string `json:"utilization_threshold,omitempty" xml:"utilization_threshold,omitempty"`
}

func (s CreateAutoscalingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoscalingConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoscalingConfigRequest) GetCoolDownDuration() *string {
	return s.CoolDownDuration
}

func (s *CreateAutoscalingConfigRequest) GetDaemonsetEvictionForNodes() *bool {
	return s.DaemonsetEvictionForNodes
}

func (s *CreateAutoscalingConfigRequest) GetExpander() *string {
	return s.Expander
}

func (s *CreateAutoscalingConfigRequest) GetGpuUtilizationThreshold() *string {
	return s.GpuUtilizationThreshold
}

func (s *CreateAutoscalingConfigRequest) GetMaxGracefulTerminationSec() *int32 {
	return s.MaxGracefulTerminationSec
}

func (s *CreateAutoscalingConfigRequest) GetMinReplicaCount() *int32 {
	return s.MinReplicaCount
}

func (s *CreateAutoscalingConfigRequest) GetPriorities() map[string][]*string {
	return s.Priorities
}

func (s *CreateAutoscalingConfigRequest) GetRecycleNodeDeletionEnabled() *bool {
	return s.RecycleNodeDeletionEnabled
}

func (s *CreateAutoscalingConfigRequest) GetScaleDownEnabled() *bool {
	return s.ScaleDownEnabled
}

func (s *CreateAutoscalingConfigRequest) GetScaleUpFromZero() *bool {
	return s.ScaleUpFromZero
}

func (s *CreateAutoscalingConfigRequest) GetScalerType() *string {
	return s.ScalerType
}

func (s *CreateAutoscalingConfigRequest) GetScanInterval() *string {
	return s.ScanInterval
}

func (s *CreateAutoscalingConfigRequest) GetSkipNodesWithLocalStorage() *bool {
	return s.SkipNodesWithLocalStorage
}

func (s *CreateAutoscalingConfigRequest) GetSkipNodesWithSystemPods() *bool {
	return s.SkipNodesWithSystemPods
}

func (s *CreateAutoscalingConfigRequest) GetUnneededDuration() *string {
	return s.UnneededDuration
}

func (s *CreateAutoscalingConfigRequest) GetUtilizationThreshold() *string {
	return s.UtilizationThreshold
}

func (s *CreateAutoscalingConfigRequest) SetCoolDownDuration(v string) *CreateAutoscalingConfigRequest {
	s.CoolDownDuration = &v
	return s
}

func (s *CreateAutoscalingConfigRequest) SetDaemonsetEvictionForNodes(v bool) *CreateAutoscalingConfigRequest {
	s.DaemonsetEvictionForNodes = &v
	return s
}

func (s *CreateAutoscalingConfigRequest) SetExpander(v string) *CreateAutoscalingConfigRequest {
	s.Expander = &v
	return s
}

func (s *CreateAutoscalingConfigRequest) SetGpuUtilizationThreshold(v string) *CreateAutoscalingConfigRequest {
	s.GpuUtilizationThreshold = &v
	return s
}

func (s *CreateAutoscalingConfigRequest) SetMaxGracefulTerminationSec(v int32) *CreateAutoscalingConfigRequest {
	s.MaxGracefulTerminationSec = &v
	return s
}

func (s *CreateAutoscalingConfigRequest) SetMinReplicaCount(v int32) *CreateAutoscalingConfigRequest {
	s.MinReplicaCount = &v
	return s
}

func (s *CreateAutoscalingConfigRequest) SetPriorities(v map[string][]*string) *CreateAutoscalingConfigRequest {
	s.Priorities = v
	return s
}

func (s *CreateAutoscalingConfigRequest) SetRecycleNodeDeletionEnabled(v bool) *CreateAutoscalingConfigRequest {
	s.RecycleNodeDeletionEnabled = &v
	return s
}

func (s *CreateAutoscalingConfigRequest) SetScaleDownEnabled(v bool) *CreateAutoscalingConfigRequest {
	s.ScaleDownEnabled = &v
	return s
}

func (s *CreateAutoscalingConfigRequest) SetScaleUpFromZero(v bool) *CreateAutoscalingConfigRequest {
	s.ScaleUpFromZero = &v
	return s
}

func (s *CreateAutoscalingConfigRequest) SetScalerType(v string) *CreateAutoscalingConfigRequest {
	s.ScalerType = &v
	return s
}

func (s *CreateAutoscalingConfigRequest) SetScanInterval(v string) *CreateAutoscalingConfigRequest {
	s.ScanInterval = &v
	return s
}

func (s *CreateAutoscalingConfigRequest) SetSkipNodesWithLocalStorage(v bool) *CreateAutoscalingConfigRequest {
	s.SkipNodesWithLocalStorage = &v
	return s
}

func (s *CreateAutoscalingConfigRequest) SetSkipNodesWithSystemPods(v bool) *CreateAutoscalingConfigRequest {
	s.SkipNodesWithSystemPods = &v
	return s
}

func (s *CreateAutoscalingConfigRequest) SetUnneededDuration(v string) *CreateAutoscalingConfigRequest {
	s.UnneededDuration = &v
	return s
}

func (s *CreateAutoscalingConfigRequest) SetUtilizationThreshold(v string) *CreateAutoscalingConfigRequest {
	s.UtilizationThreshold = &v
	return s
}

func (s *CreateAutoscalingConfigRequest) Validate() error {
	return dara.Validate(s)
}
