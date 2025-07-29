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
	// The waiting time before the auto scaling feature performs a scale-in activity. It is an interval between the time when the scale-in threshold is reached and the time when the scale-in activity (reducing the number of pods) starts. Unit: minutes. Default value: 10.
	//
	// example:
	//
	// 10 m
	CoolDownDuration *string `json:"cool_down_duration,omitempty" xml:"cool_down_duration,omitempty"`
	// Specifies whether to evict pods created by DaemonSets when the cluster autoscaler performs a scale-in activity. Valid values:
	//
	// 	- `true`: evicts DaemonSet pods.
	//
	// 	- `false`: does not evict DaemonSet pods.
	//
	// example:
	//
	// false
	DaemonsetEvictionForNodes *bool `json:"daemonset_eviction_for_nodes,omitempty" xml:"daemonset_eviction_for_nodes,omitempty"`
	// The node pool scale-out policy. Valid values:
	//
	// 	- `least-waste`: the default policy. If multiple node pools meet the requirement, this policy selects the node pool that will have the least idle resources after the scale-out activity is completed.
	//
	// 	- `random`: the random policy. If multiple node pools meet the requirement, this policy selects a random node pool for the scale-out activity.
	//
	// 	- `priority`: the priority-based policy If multiple node pools meet the requirement, this policy selects the node pool with the highest priority for the scale-out activity. The priority setting is stored in the ConfigMap named `cluster-autoscaler-priority-expander` in the kube-system namespace. When a scale-out activity is triggered, the policy obtains the node pool priorities from the ConfigMap based on the node pool IDs and then selects the node pool with the highest priority for the scale-out activity.
	//
	// example:
	//
	// least-waste
	Expander *string `json:"expander,omitempty" xml:"expander,omitempty"`
	// The scale-in threshold of GPU utilization. This threshold specifies the ratio of the GPU resources that are requested by pods to the total GPU resources on the node.
	//
	// A scale-in activity is performed only when the CPU utilization, memory utilization, and GPU utilization of a GPU-accelerated node are lower than the scale-in threshold of GPU utilization.
	//
	// example:
	//
	// 0.5
	GpuUtilizationThreshold *string `json:"gpu_utilization_threshold,omitempty" xml:"gpu_utilization_threshold,omitempty"`
	// The maximum amount of time to wait for pods on a node to terminate during a scale-in activity. Unit: seconds.
	//
	// example:
	//
	// 14400s
	MaxGracefulTerminationSec *int32 `json:"max_graceful_termination_sec,omitempty" xml:"max_graceful_termination_sec,omitempty"`
	// The minimum number of pods allowed in each ReplicaSet before a scale-in activity is performed.
	//
	// example:
	//
	// 0
	MinReplicaCount *int32 `json:"min_replica_count,omitempty" xml:"min_replica_count,omitempty"`
	// Auto-scaling priority configuration. After creating a node pool with elasticity enabled, you can choose whether to configure a priority strategy and priority settings through [Enabling Node Auto-scaling](https://help.aliyun.com/document_detail/119099.html). This allows you to set priorities for the specified auto-scaling node pool scaling group. The priority value range is [1, 100] and must be a positive integer.
	Priorities map[string][]*string `json:"priorities,omitempty" xml:"priorities,omitempty"`
	// Specifies whether to delete the corresponding Kubernetes node objects after nodes are removed in swift mode. For more information about the swift mode, see [Scaling mode](https://help.aliyun.com/document_detail/119099.html). Default value: false Valid values:
	//
	// 	- `true`: deletes the corresponding Kubernetes node objects after nodes are removed in swift mode. We recommend that you do not set the value to true because data inconsistency may occur in Kubernetes objects.
	//
	// 	- `false`: retains the corresponding Kubernetes node objects after nodes are removed in swift mode.
	//
	// example:
	//
	// false
	RecycleNodeDeletionEnabled *bool `json:"recycle_node_deletion_enabled,omitempty" xml:"recycle_node_deletion_enabled,omitempty"`
	// Specifies whether to allow node scale-in activities. Valid values:
	//
	// 	- `true`: allows node scale-in activities.
	//
	// 	- `false`: does not allow node scale-in activities.
	//
	// example:
	//
	// true
	ScaleDownEnabled *bool `json:"scale_down_enabled,omitempty" xml:"scale_down_enabled,omitempty"`
	// Specifies whether the cluster autoscaler performs a scale-out activity when the number of ready nodes in the cluster is 0. Default value: true. Valid values:
	//
	// 	- `true`: performs a scale-out activity.
	//
	// 	- `false`: does not perform a scale-out activity.
	//
	// example:
	//
	// true
	ScaleUpFromZero *bool `json:"scale_up_from_zero,omitempty" xml:"scale_up_from_zero,omitempty"`
	// Elastic component type, default is goatscaler for cluster version 1.24 and above, and cluster-autoscaler below that. Values:
	//
	// - `goatscaler`: Instant elasticity.
	//
	// - `cluster-autoscaler`: Auto-scaling.
	//
	// example:
	//
	// goatscaler
	ScalerType *string `json:"scaler_type,omitempty" xml:"scaler_type,omitempty"`
	// The interval at which the system scans for events that trigger scaling activities. Unit: seconds. Default value: 60.
	//
	// example:
	//
	// 30s
	ScanInterval *string `json:"scan_interval,omitempty" xml:"scan_interval,omitempty"`
	// Specifies whether the cluster autoscaler scales in nodes that host pods mounted with local volumes, such as EmptyDir or HostPath volumes. Valid values:
	//
	// 	- `true`: does not allow the cluster autoscaler to scale in these nodes.
	//
	// 	- `false`: allows the cluster autoscaler to scale in these nodes.
	//
	// example:
	//
	// false
	SkipNodesWithLocalStorage *bool `json:"skip_nodes_with_local_storage,omitempty" xml:"skip_nodes_with_local_storage,omitempty"`
	// Specifies whether the cluster autoscaler scales in nodes that host pods in the kube-system namespace. This parameter does not take effect on pods created by DaemonSets and mirror pods. Valid values:
	//
	// 	- `true`: does not allow the cluster autoscaler to scale in these nodes.
	//
	// 	- `false`: allows the cluster autoscaler to scale in these nodes.
	//
	// example:
	//
	// true
	SkipNodesWithSystemPods *bool `json:"skip_nodes_with_system_pods,omitempty" xml:"skip_nodes_with_system_pods,omitempty"`
	// The cooldown period. After the autoscaler performs a scale-out activity, the autoscaler waits a cooldown period before it can perform a scale-in activity. Newly added nodes can be removed in scale-in activities only after the cooldown period ends. Unit: minutes.
	//
	// example:
	//
	// 10 m
	UnneededDuration *string `json:"unneeded_duration,omitempty" xml:"unneeded_duration,omitempty"`
	// The scale-in threshold. This threshold specifies the ratio of the resources that are requested by pods to the total resources on the node.
	//
	// A scale-in activity is performed only when the CPU utilization and memory utilization of a node are lower than the scale-in threshold.
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
