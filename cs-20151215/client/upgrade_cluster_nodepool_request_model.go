// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeClusterNodepoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoreWarningCheck(v bool) *UpgradeClusterNodepoolRequest
	GetIgnoreWarningCheck() *bool
	SetImageId(v string) *UpgradeClusterNodepoolRequest
	GetImageId() *string
	SetKubernetesVersion(v string) *UpgradeClusterNodepoolRequest
	GetKubernetesVersion() *string
	SetNodeNames(v []*string) *UpgradeClusterNodepoolRequest
	GetNodeNames() []*string
	SetRollingPolicy(v *UpgradeClusterNodepoolRequestRollingPolicy) *UpgradeClusterNodepoolRequest
	GetRollingPolicy() *UpgradeClusterNodepoolRequestRollingPolicy
	SetRuntimeType(v string) *UpgradeClusterNodepoolRequest
	GetRuntimeType() *string
	SetRuntimeVersion(v string) *UpgradeClusterNodepoolRequest
	GetRuntimeVersion() *string
	SetUseReplace(v bool) *UpgradeClusterNodepoolRequest
	GetUseReplace() *bool
}

type UpgradeClusterNodepoolRequest struct {
	// Specifies whether to ignore warning-level pre-checks.
	IgnoreWarningCheck *bool `json:"ignore_warning_check,omitempty" xml:"ignore_warning_check,omitempty"`
	// The node system image ID.
	//
	// example:
	//
	// aliyun_3_x64_20G_container_optimized_20241226.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The node Kubernetes version. You can call the [DescribeKubernetesVersionMetadata](https://help.aliyun.com/document_detail/2667899.html) operation to obtain the current cluster version information from the `KubernetesVersion` field.
	//
	// example:
	//
	// 1.32.1-aliyun.1
	KubernetesVersion *string `json:"kubernetes_version,omitempty" xml:"kubernetes_version,omitempty"`
	// The list of specified nodes to upgrade. If not specified, all nodes in the node pool are upgraded by default.
	NodeNames []*string `json:"node_names,omitempty" xml:"node_names,omitempty" type:"Repeated"`
	// The rolling update configuration.
	RollingPolicy *UpgradeClusterNodepoolRequestRollingPolicy `json:"rolling_policy,omitempty" xml:"rolling_policy,omitempty" type:"Struct"`
	// The runtime type. You can call the [DescribeKubernetesVersionMetadata](https://help.aliyun.com/document_detail/2667899.html) operation to obtain the runtime information from the runtime field.
	//
	// example:
	//
	// containerd
	RuntimeType *string `json:"runtime_type,omitempty" xml:"runtime_type,omitempty"`
	// The node runtime version. You can call the [DescribeKubernetesVersionMetadata](https://help.aliyun.com/document_detail/2667899.html) operation to obtain the runtime version information from the runtime field.
	//
	// example:
	//
	// 1.6.36
	RuntimeVersion *string `json:"runtime_version,omitempty" xml:"runtime_version,omitempty"`
	// Specifies whether to use disk replacement for the upgrade. Valid values:
	//
	// - true: Uses disk replacement. The node pool is upgraded by replacing the system cloud disks of the nodes. ACK reinitializes the nodes based on the current configurations of the node pool, such as the node logon method, labels, taints, operating system image, and runtime version.
	//
	// - false: Does not use disk replacement.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	UseReplace *bool `json:"use_replace,omitempty" xml:"use_replace,omitempty"`
}

func (s UpgradeClusterNodepoolRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClusterNodepoolRequest) GoString() string {
	return s.String()
}

func (s *UpgradeClusterNodepoolRequest) GetIgnoreWarningCheck() *bool {
	return s.IgnoreWarningCheck
}

func (s *UpgradeClusterNodepoolRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpgradeClusterNodepoolRequest) GetKubernetesVersion() *string {
	return s.KubernetesVersion
}

func (s *UpgradeClusterNodepoolRequest) GetNodeNames() []*string {
	return s.NodeNames
}

func (s *UpgradeClusterNodepoolRequest) GetRollingPolicy() *UpgradeClusterNodepoolRequestRollingPolicy {
	return s.RollingPolicy
}

func (s *UpgradeClusterNodepoolRequest) GetRuntimeType() *string {
	return s.RuntimeType
}

func (s *UpgradeClusterNodepoolRequest) GetRuntimeVersion() *string {
	return s.RuntimeVersion
}

func (s *UpgradeClusterNodepoolRequest) GetUseReplace() *bool {
	return s.UseReplace
}

func (s *UpgradeClusterNodepoolRequest) SetIgnoreWarningCheck(v bool) *UpgradeClusterNodepoolRequest {
	s.IgnoreWarningCheck = &v
	return s
}

func (s *UpgradeClusterNodepoolRequest) SetImageId(v string) *UpgradeClusterNodepoolRequest {
	s.ImageId = &v
	return s
}

func (s *UpgradeClusterNodepoolRequest) SetKubernetesVersion(v string) *UpgradeClusterNodepoolRequest {
	s.KubernetesVersion = &v
	return s
}

func (s *UpgradeClusterNodepoolRequest) SetNodeNames(v []*string) *UpgradeClusterNodepoolRequest {
	s.NodeNames = v
	return s
}

func (s *UpgradeClusterNodepoolRequest) SetRollingPolicy(v *UpgradeClusterNodepoolRequestRollingPolicy) *UpgradeClusterNodepoolRequest {
	s.RollingPolicy = v
	return s
}

func (s *UpgradeClusterNodepoolRequest) SetRuntimeType(v string) *UpgradeClusterNodepoolRequest {
	s.RuntimeType = &v
	return s
}

func (s *UpgradeClusterNodepoolRequest) SetRuntimeVersion(v string) *UpgradeClusterNodepoolRequest {
	s.RuntimeVersion = &v
	return s
}

func (s *UpgradeClusterNodepoolRequest) SetUseReplace(v bool) *UpgradeClusterNodepoolRequest {
	s.UseReplace = &v
	return s
}

func (s *UpgradeClusterNodepoolRequest) Validate() error {
	if s.RollingPolicy != nil {
		if err := s.RollingPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpgradeClusterNodepoolRequestRollingPolicy struct {
	// The upgrade interval between batches. This parameter takes effect only when the pause policy is set to `NotPause`.
	//
	// Valid values: 5 to 120. Unit: minutes.
	//
	// You can set this parameter to 0, which indicates that no interval is configured.
	//
	// example:
	//
	// 5
	BatchInterval *int32 `json:"batch_interval,omitempty" xml:"batch_interval,omitempty"`
	// The maximum number of nodes allowed to fail during the rolling update. Default value: 0, which indicates that the task fails if any node fails. If the value is greater than 0, the task fails and stops when the cumulative number of failed nodes exceeds this value.
	//
	// example:
	//
	// 0
	MaxFailedNodes *int32 `json:"max_failed_nodes,omitempty" xml:"max_failed_nodes,omitempty"`
	// The node updates in the node pool are performed in batches. This parameter specifies the maximum number of nodes that can be updated in parallel in each batch.
	//
	// Valid values: 1 to 10.
	//
	// Default value: 10.
	//
	// example:
	//
	// 2
	MaxParallelism *int32 `json:"max_parallelism,omitempty" xml:"max_parallelism,omitempty"`
	// The automatic pause policy during the node upgrade process. Valid values:
	//
	// - FirstBatch: Pauses after the first batch is completed.
	//
	// - EveryBatch: Pauses after each batch is completed.
	//
	// - NotPause: Does not pause.
	//
	// example:
	//
	// NotPause
	PausePolicy *string `json:"pause_policy,omitempty" xml:"pause_policy,omitempty"`
}

func (s UpgradeClusterNodepoolRequestRollingPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClusterNodepoolRequestRollingPolicy) GoString() string {
	return s.String()
}

func (s *UpgradeClusterNodepoolRequestRollingPolicy) GetBatchInterval() *int32 {
	return s.BatchInterval
}

func (s *UpgradeClusterNodepoolRequestRollingPolicy) GetMaxFailedNodes() *int32 {
	return s.MaxFailedNodes
}

func (s *UpgradeClusterNodepoolRequestRollingPolicy) GetMaxParallelism() *int32 {
	return s.MaxParallelism
}

func (s *UpgradeClusterNodepoolRequestRollingPolicy) GetPausePolicy() *string {
	return s.PausePolicy
}

func (s *UpgradeClusterNodepoolRequestRollingPolicy) SetBatchInterval(v int32) *UpgradeClusterNodepoolRequestRollingPolicy {
	s.BatchInterval = &v
	return s
}

func (s *UpgradeClusterNodepoolRequestRollingPolicy) SetMaxFailedNodes(v int32) *UpgradeClusterNodepoolRequestRollingPolicy {
	s.MaxFailedNodes = &v
	return s
}

func (s *UpgradeClusterNodepoolRequestRollingPolicy) SetMaxParallelism(v int32) *UpgradeClusterNodepoolRequestRollingPolicy {
	s.MaxParallelism = &v
	return s
}

func (s *UpgradeClusterNodepoolRequestRollingPolicy) SetPausePolicy(v string) *UpgradeClusterNodepoolRequestRollingPolicy {
	s.PausePolicy = &v
	return s
}

func (s *UpgradeClusterNodepoolRequestRollingPolicy) Validate() error {
	return dara.Validate(s)
}
