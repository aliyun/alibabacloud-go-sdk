// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeClusterNodepoolRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// The ID of the OS image used by the nodes.
	//
	// example:
	//
	// aliyun_2_1903_x64_20G_alibase_20200529.vhd
	ImageId *string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// The Kubernetes version used by the nodes. You can call the [DescribeKubernetesVersionMetadata](https://help.aliyun.com/document_detail/2667899.html) operation and get the Kubernetes version of the current cluster in the current_version field.
	//
	// example:
	//
	// 1.22.15-aliyun.1
	KubernetesVersion *string `json:"kubernetes_version,omitempty" xml:"kubernetes_version,omitempty"`
	// The nodes you want to update. If you do not specify this parameter, all nodes in the node pool are updated by default.
	NodeNames []*string `json:"node_names,omitempty" xml:"node_names,omitempty" type:"Repeated"`
	// The rolling update configuration.
	RollingPolicy *UpgradeClusterNodepoolRequestRollingPolicy `json:"rolling_policy,omitempty" xml:"rolling_policy,omitempty" type:"Struct"`
	// The runtime type. You can call the [DescribeKubernetesVersionMetadata](https://help.aliyun.com/document_detail/2667899.html) operation and get the runtime information in the runtime field.
	//
	// example:
	//
	// containerd
	RuntimeType *string `json:"runtime_type,omitempty" xml:"runtime_type,omitempty"`
	// The version of the container runtime used by the nodes. You can call the [DescribeKubernetesVersionMetadata](https://help.aliyun.com/document_detail/2667899.html) operation and get the runtime version in the runtime field.
	//
	// example:
	//
	// 1.5.10
	RuntimeVersion *string `json:"runtime_version,omitempty" xml:"runtime_version,omitempty"`
	// Specifies whether to perform the update by replacing the system disk. Valid values:
	//
	// 	- true: replaces the system disk.
	//
	// 	- false: does not replace the system disk.
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
	// The update interval between batches takes effect only when the pause policy is set to NotPause. Unit: minutes. Valid values: 5 to 120.
	//
	// example:
	//
	// 5 minutes
	BatchInterval *int32 `json:"batch_interval,omitempty" xml:"batch_interval,omitempty"`
	// The maximum number of nodes per batch.
	//
	// example:
	//
	// 3
	MaxParallelism *int32 `json:"max_parallelism,omitempty" xml:"max_parallelism,omitempty"`
	// The policy used to pause the update. Valid values:
	//
	// 	- FirstBatch: pauses after the first batch is updated.
	//
	// 	- EveryBatch: pauses after each batch is updated.
	//
	// 	- NotPause: does not pause.
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
