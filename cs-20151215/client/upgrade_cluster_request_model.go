// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentName(v string) *UpgradeClusterRequest
	GetComponentName() *string
	SetMasterOnly(v bool) *UpgradeClusterRequest
	GetMasterOnly() *bool
	SetNextVersion(v string) *UpgradeClusterRequest
	GetNextVersion() *string
	SetRollingPolicy(v *UpgradeClusterRequestRollingPolicy) *UpgradeClusterRequest
	GetRollingPolicy() *UpgradeClusterRequestRollingPolicy
	SetVersion(v string) *UpgradeClusterRequest
	GetVersion() *string
}

type UpgradeClusterRequest struct {
	// Deprecated
	//
	// This parameter is deprecated. No need to pass values.
	//
	// example:
	//
	// k8s
	ComponentName *string `json:"component_name,omitempty" xml:"component_name,omitempty"`
	// Specifies whether to upgrade only master nodes. Valid values:
	//
	// 	- true: upgrades master nodes only.
	//
	// 	- false: upgrades both master and worker nodes.
	//
	// example:
	//
	// true
	MasterOnly *bool `json:"master_only,omitempty" xml:"master_only,omitempty"`
	// The target Kubernetes version for cluster upgrade.
	//
	// example:
	//
	// 1.16.9-aliyun.1
	NextVersion *string `json:"next_version,omitempty" xml:"next_version,omitempty"`
	// Deprecated
	//
	// The rolling update configuration.
	RollingPolicy *UpgradeClusterRequestRollingPolicy `json:"rolling_policy,omitempty" xml:"rolling_policy,omitempty" type:"Struct"`
	// Deprecated
	//
	// This parameter is deprecated. Use next_version to specify the upgrade target Kubernetes version.
	//
	// example:
	//
	// 1.14.8-aliyun.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpgradeClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClusterRequest) GoString() string {
	return s.String()
}

func (s *UpgradeClusterRequest) GetComponentName() *string {
	return s.ComponentName
}

func (s *UpgradeClusterRequest) GetMasterOnly() *bool {
	return s.MasterOnly
}

func (s *UpgradeClusterRequest) GetNextVersion() *string {
	return s.NextVersion
}

func (s *UpgradeClusterRequest) GetRollingPolicy() *UpgradeClusterRequestRollingPolicy {
	return s.RollingPolicy
}

func (s *UpgradeClusterRequest) GetVersion() *string {
	return s.Version
}

func (s *UpgradeClusterRequest) SetComponentName(v string) *UpgradeClusterRequest {
	s.ComponentName = &v
	return s
}

func (s *UpgradeClusterRequest) SetMasterOnly(v bool) *UpgradeClusterRequest {
	s.MasterOnly = &v
	return s
}

func (s *UpgradeClusterRequest) SetNextVersion(v string) *UpgradeClusterRequest {
	s.NextVersion = &v
	return s
}

func (s *UpgradeClusterRequest) SetRollingPolicy(v *UpgradeClusterRequestRollingPolicy) *UpgradeClusterRequest {
	s.RollingPolicy = v
	return s
}

func (s *UpgradeClusterRequest) SetVersion(v string) *UpgradeClusterRequest {
	s.Version = &v
	return s
}

func (s *UpgradeClusterRequest) Validate() error {
	if s.RollingPolicy != nil {
		if err := s.RollingPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpgradeClusterRequestRollingPolicy struct {
	// Deprecated
	//
	// The maximum number of nodes concurrently upgraded per batch.
	//
	// example:
	//
	// 3
	MaxParallelism *int32 `json:"max_parallelism,omitempty" xml:"max_parallelism,omitempty"`
}

func (s UpgradeClusterRequestRollingPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClusterRequestRollingPolicy) GoString() string {
	return s.String()
}

func (s *UpgradeClusterRequestRollingPolicy) GetMaxParallelism() *int32 {
	return s.MaxParallelism
}

func (s *UpgradeClusterRequestRollingPolicy) SetMaxParallelism(v int32) *UpgradeClusterRequestRollingPolicy {
	s.MaxParallelism = &v
	return s
}

func (s *UpgradeClusterRequestRollingPolicy) Validate() error {
	return dara.Validate(s)
}
