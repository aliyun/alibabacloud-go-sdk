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
	// The cool-down duration for scale-in events. This is the time interval from when the system detects a node is eligible for a scale-in to when the scale-in operation is executed.
	//
	// Valid values: 1 to 60. Unit: minutes.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	CoolDownDuration *string `json:"cool_down_duration,omitempty" xml:"cool_down_duration,omitempty"`
	// Specifies whether `cluster-autoscaler` evicts DaemonSet Pods from nodes during a scale-in event. Valid values:
	//
	// - `true`: Perform eviction.
	//
	// - `false`: Do not perform eviction.
	//
	// example:
	//
	// false
	DaemonsetEvictionForNodes *bool `json:"daemonset_eviction_for_nodes,omitempty" xml:"daemonset_eviction_for_nodes,omitempty"`
	// The strategy for selecting a node pool for a scale-out when multiple node pools are available. Valid values:
	//
	// - `least-waste`: The default strategy. The scaler selects the node pool that will have the least idle resources after a scale-out.
	//
	// - `random`: The scaler selects a random node pool from the list of eligible node pools.
	//
	// - `priority`: The scaler selects the node pool that has the highest priority. You must configure the priority of each scaling group by using the `priorities` parameter.
	//
	// example:
	//
	// least-waste
	Expander *string `json:"expander,omitempty" xml:"expander,omitempty"`
	// The GPU utilization threshold for a scale-in on GPU nodes, which is the ratio of requested resources to total allocatable resources on a node.
	//
	// A GPU node is eligible for a scale-in only if its CPU, memory, and GPU utilization all fall below this threshold.
	//
	// Valid values: [0.1, 1].
	//
	// Default value: 0.3 (30%).
	//
	// example:
	//
	// 0.3
	GpuUtilizationThreshold *string `json:"gpu_utilization_threshold,omitempty" xml:"gpu_utilization_threshold,omitempty"`
	// The maximum duration in seconds that `cluster-autoscaler` waits for Pods to terminate during a node drain for a scale-in event.
	//
	// Unit: seconds.
	//
	// Default value: 14400.
	//
	// example:
	//
	// 14400
	MaxGracefulTerminationSec *int32 `json:"max_graceful_termination_sec,omitempty" xml:"max_graceful_termination_sec,omitempty"`
	// The minimum number of Pods that must remain for any ReplicaSet after a scale-in operation. Nodes will not be scaled-in if doing so would violate this minimum.
	//
	// Default value: 0.
	//
	// example:
	//
	// 0
	MinReplicaCount *int32 `json:"min_replica_count,omitempty" xml:"min_replica_count,omitempty"`
	// Configures the priorities for scaling groups. This is used when the `expander` strategy is set to `priority`. After you create a node pool and enable autoscaling for it, you can configure the priority of its associated scaling group. For more information, see [Enable node autoscaling](https://help.aliyun.com/document_detail/119099.html).
	//
	// The priority must be a positive integer from 1 to 100. A larger value indicates a higher priority.
	Priorities map[string][]*string `json:"priorities,omitempty" xml:"priorities,omitempty"`
	// Specifies whether to delete the Kubernetes Node object after a node is successfully scaled-in using fast scaling mode. For more information, see [Scaling modes](https://help.aliyun.com/document_detail/119099.html). Default value: false. Valid values:
	//
	// - `true`: The Node object is deleted after the instance is stopped. We do not recommend this setting because it can cause data inconsistencies in Kubernetes.
	//
	// - `false`: The Node object is retained after the instance is stopped.
	//
	// example:
	//
	// false
	RecycleNodeDeletionEnabled *bool `json:"recycle_node_deletion_enabled,omitempty" xml:"recycle_node_deletion_enabled,omitempty"`
	// Specifies whether to allow node scale-in operations. Valid values:
	//
	// - `true`: Allows scale-in operations.
	//
	// - `false`: Disables scale-in operations.
	//
	// example:
	//
	// true
	ScaleDownEnabled *bool `json:"scale_down_enabled,omitempty" xml:"scale_down_enabled,omitempty"`
	// Controls whether `cluster-autoscaler` performs a scale-out operation when there are no ready nodes in the cluster. Default value: true. Valid values:
	//
	// - `true`: A scale-out operation is performed.
	//
	// - `false`: No scale-out operation is performed.
	//
	// example:
	//
	// true
	ScaleUpFromZero *bool `json:"scale_up_from_zero,omitempty" xml:"scale_up_from_zero,omitempty"`
	// The type of scaler to use. In clusters that run Kubernetes 1.24 or later, the default is goatscaler. In clusters that run an earlier version, the default is cluster-autoscaler. Valid values:
	//
	// - `goatscaler`: The proprietary scaler for fast scaling.
	//
	// - `cluster-autoscaler`: The standard Kubernetes cluster autoscaler.
	//
	// example:
	//
	// goatscaler
	ScalerType *string `json:"scaler_type,omitempty" xml:"scaler_type,omitempty"`
	// The frequency at which the system checks for scaling conditions.
	//
	// Valid values: 15, 30, 60, 120, 180, and 300. Unit: seconds.
	//
	// Default value: 60.
	//
	// example:
	//
	// 30
	ScanInterval *string `json:"scan_interval,omitempty" xml:"scan_interval,omitempty"`
	// Controls whether `cluster-autoscaler` can scale-in nodes that run Pods using local storage (for example, with `emptyDir` or `hostPath` volumes). Valid values:
	//
	// - `true`: Prevents these nodes from being scaled-in.
	//
	// - `false`: Allows these nodes to be scaled-in.
	//
	// example:
	//
	// false
	SkipNodesWithLocalStorage *bool `json:"skip_nodes_with_local_storage,omitempty" xml:"skip_nodes_with_local_storage,omitempty"`
	// Controls whether `cluster-autoscaler` can scale-in nodes that run Pods from the `kube-system` namespace. This setting does not affect DaemonSet or mirror Pods. Valid values:
	//
	// - `true`: Prevents these nodes from being scaled-in.
	//
	// - `false`: Allows these nodes to be scaled-in.
	//
	// example:
	//
	// true
	SkipNodesWithSystemPods *bool `json:"skip_nodes_with_system_pods,omitempty" xml:"skip_nodes_with_system_pods,omitempty"`
	// The stabilization window. This is the period after a scale-out event during which the scaler does not perform scale-in operations.
	//
	// Valid values: 1 to 60. Unit: minutes.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	UnneededDuration *string `json:"unneeded_duration,omitempty" xml:"unneeded_duration,omitempty"`
	// The utilization threshold for a scale-in, which is the ratio of requested resources to the total allocatable resources on a node.
	//
	// A node is eligible for a scale-in only when both its CPU and memory utilization fall below this threshold.
	//
	// Valid values: [0.1, 1].
	//
	// Default value: 0.5 (50%).
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
