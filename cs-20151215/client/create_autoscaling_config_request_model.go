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
	// The delay for triggering scale-in. This is the time interval between detecting a scale-in requirement (reaching the scale-in threshold) and actually executing the scale-in operation (reducing the number of pods). Valid values: 1 to 60. Unit: minutes. Default value: 10.
	//
	// example:
	//
	// 10
	CoolDownDuration *string `json:"cool_down_duration,omitempty" xml:"cool_down_duration,omitempty"`
	// Specifies whether cluster-autoscaler evicts DaemonSet pods on nodes during scale-in. Valid values:
	//
	// - true: The pods are evicted.
	//
	// - false: The pods are not evicted.
	//
	// example:
	//
	// false
	DaemonsetEvictionForNodes *bool `json:"daemonset_eviction_for_nodes,omitempty" xml:"daemonset_eviction_for_nodes,omitempty"`
	// The node pool scale-out order policy. Valid values:
	//
	// - least-waste: the default policy. If multiple node pools can be scaled out, the system selects the node pool that wastes the least resources.
	//
	// - random: the random policy. If multiple node pools can be scaled out, the system randomly selects a node pool.
	//
	// - priority: the priority policy. If multiple node pools can be scaled out, the system selects the node pool with the highest priority based on the custom scaling group order. The node pool priority is defined by the priorities parameter.
	//
	// example:
	//
	// least-waste
	Expander *string `json:"expander,omitempty" xml:"expander,omitempty"`
	// The GPU scale-in threshold, which is the ratio of requested resources to total resources on a node. A GPU node can be scaled in only when this ratio is lower than the configured threshold, meaning that the CPU, memory, and GPU utilization of the node are all lower than the GPU scale-in threshold. Valid values: 0.1 to 1. Default value: 0.3, which indicates 30%.
	//
	// example:
	//
	// 0.3
	GpuUtilizationThreshold *string `json:"gpu_utilization_threshold,omitempty" xml:"gpu_utilization_threshold,omitempty"`
	// The timeout period for cluster-autoscaler to wait for pod termination during node draining in a scale-in scenario. Unit: seconds. Default value: 14400.
	//
	// example:
	//
	// 14400
	MaxGracefulTerminationSec *int32 `json:"max_graceful_termination_sec,omitempty" xml:"max_graceful_termination_sec,omitempty"`
	// The minimum number of pods allowed in each ReplicaSet before a node is scaled in. Default value: 0.
	//
	// example:
	//
	// 0
	MinReplicaCount *int32 `json:"min_replica_count,omitempty" xml:"min_replica_count,omitempty"`
	// The auto scaling priority configuration. After you create a node pool with auto scaling enabled, you can choose whether to configure the priority policy and priority configuration by referring to [Enable auto scaling for nodes](https://help.aliyun.com/document_detail/119099.html). This allows you to configure priorities for the scaling groups of specified auto scaling node pools. Valid values: 1 to 100. The value must be a positive integer. A larger value indicates a higher priority.
	Priorities map[string][]*string `json:"priorities,omitempty" xml:"priorities,omitempty"`
	// Specifies whether to delete the corresponding Kubernetes node objects after nodes are successfully scaled in fast mode. For more information about fast mode, see [Scaling modes](https://help.aliyun.com/document_detail/119099.html). Default value: false. Valid values:
	//
	// - true: Kubernetes node objects are also deleted after nodes are stopped in fast mode. Do not set this parameter to true because it may cause data inconsistency of Kubernetes objects.
	//
	// - false: Kubernetes node objects are retained after nodes are stopped in fast mode.
	//
	// example:
	//
	// false
	RecycleNodeDeletionEnabled *bool `json:"recycle_node_deletion_enabled,omitempty" xml:"recycle_node_deletion_enabled,omitempty"`
	// Specifies whether to allow node scale-in. Valid values:
	//
	// - true: Scale-in is allowed.
	//
	// - false: Scale-in is prohibited.
	//
	// example:
	//
	// true
	ScaleDownEnabled *bool `json:"scale_down_enabled,omitempty" xml:"scale_down_enabled,omitempty"`
	// Specifies whether cluster-autoscaler performs scale-out when the number of ready nodes in the cluster is 0. Default value: true. Valid values:
	//
	// - true: Scale-out is performed.
	//
	// - false: Scale-out is not performed.
	//
	// example:
	//
	// true
	ScaleUpFromZero *bool `json:"scale_up_from_zero,omitempty" xml:"scale_up_from_zero,omitempty"`
	// The auto scaling component type. For clusters of version 1.24 and later, the default value is goatscaler. For clusters of earlier versions, the default value is cluster-autoscaler. Valid values:
	//
	// - goatscaler: instant elasticity.
	//
	// - cluster-autoscaler: auto scaling.
	//
	// example:
	//
	// goatscaler
	ScalerType *string `json:"scaler_type,omitempty" xml:"scaler_type,omitempty"`
	// The auto scaling sensitivity, which is used to adjust the interval at which the system evaluates scaling. Valid values: 15, 30, 60, 120, 180, and 300. Unit: seconds. Default value: 60.
	//
	// example:
	//
	// 30
	ScanInterval *string `json:"scan_interval,omitempty" xml:"scan_interval,omitempty"`
	// Specifies whether cluster-autoscaler scales in nodes that contain pods with local storage, such as EmptyDir or HostPath. Valid values:
	//
	// - true: The nodes are not scaled in.
	//
	// - false: The nodes are scaled in.
	//
	// example:
	//
	// false
	SkipNodesWithLocalStorage *bool `json:"skip_nodes_with_local_storage,omitempty" xml:"skip_nodes_with_local_storage,omitempty"`
	// Specifies whether cluster-autoscaler scales in nodes that run pods in the kube-system namespace. This feature does not take effect for DaemonSet pods and mirror pods. Valid values:
	//
	// - true: The nodes are not scaled in.
	//
	// - false: The nodes are scaled in.
	//
	// example:
	//
	// true
	SkipNodesWithSystemPods *bool `json:"skip_nodes_with_system_pods,omitempty" xml:"skip_nodes_with_system_pods,omitempty"`
	// The cool-down period. This is the time interval during which the auto scaling component does not perform scale-in after the most recent scale-out is completed. Scaled-out nodes can be evaluated for scale-in only after the cool-down period ends. Valid values: 1 to 60. Unit: minutes. Default value: 10.
	//
	// example:
	//
	// 10
	UnneededDuration *string `json:"unneeded_duration,omitempty" xml:"unneeded_duration,omitempty"`
	// The scale-in threshold, which is the ratio of requested resources to total resources on a node. A node can be scaled in only when this ratio is lower than the configured threshold, meaning that the utilization of both CPU and memory resources on the node is lower than the scale-in threshold. Valid values: 0.1 to 1. Default value: 0.5, which indicates 50%.
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
