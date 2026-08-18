// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodePoolNodeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContainerdConfig(v *ContainerdConfig) *ModifyNodePoolNodeConfigRequest
	GetContainerdConfig() *ContainerdConfig
	SetKubeletConfig(v *KubeletConfig) *ModifyNodePoolNodeConfigRequest
	GetKubeletConfig() *KubeletConfig
	SetNodeNames(v []*string) *ModifyNodePoolNodeConfigRequest
	GetNodeNames() []*string
	SetOsConfig(v *ModifyNodePoolNodeConfigRequestOsConfig) *ModifyNodePoolNodeConfigRequest
	GetOsConfig() *ModifyNodePoolNodeConfigRequestOsConfig
	SetRollingPolicy(v *ModifyNodePoolNodeConfigRequestRollingPolicy) *ModifyNodePoolNodeConfigRequest
	GetRollingPolicy() *ModifyNodePoolNodeConfigRequestRollingPolicy
}

type ModifyNodePoolNodeConfigRequest struct {
	// The containerd runtime configuration.
	ContainerdConfig *ContainerdConfig `json:"containerd_config,omitempty" xml:"containerd_config,omitempty"`
	// The kubelet parameter settings.
	KubeletConfig *KubeletConfig `json:"kubelet_config,omitempty" xml:"kubelet_config,omitempty"`
	// The list of nodes to upgrade.
	NodeNames []*string `json:"node_names,omitempty" xml:"node_names,omitempty" type:"Repeated"`
	// The operating system parameter settings.
	OsConfig *ModifyNodePoolNodeConfigRequestOsConfig `json:"os_config,omitempty" xml:"os_config,omitempty" type:"Struct"`
	// The rolling update configuration.
	RollingPolicy *ModifyNodePoolNodeConfigRequestRollingPolicy `json:"rolling_policy,omitempty" xml:"rolling_policy,omitempty" type:"Struct"`
}

func (s ModifyNodePoolNodeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodePoolNodeConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolNodeConfigRequest) GetContainerdConfig() *ContainerdConfig {
	return s.ContainerdConfig
}

func (s *ModifyNodePoolNodeConfigRequest) GetKubeletConfig() *KubeletConfig {
	return s.KubeletConfig
}

func (s *ModifyNodePoolNodeConfigRequest) GetNodeNames() []*string {
	return s.NodeNames
}

func (s *ModifyNodePoolNodeConfigRequest) GetOsConfig() *ModifyNodePoolNodeConfigRequestOsConfig {
	return s.OsConfig
}

func (s *ModifyNodePoolNodeConfigRequest) GetRollingPolicy() *ModifyNodePoolNodeConfigRequestRollingPolicy {
	return s.RollingPolicy
}

func (s *ModifyNodePoolNodeConfigRequest) SetContainerdConfig(v *ContainerdConfig) *ModifyNodePoolNodeConfigRequest {
	s.ContainerdConfig = v
	return s
}

func (s *ModifyNodePoolNodeConfigRequest) SetKubeletConfig(v *KubeletConfig) *ModifyNodePoolNodeConfigRequest {
	s.KubeletConfig = v
	return s
}

func (s *ModifyNodePoolNodeConfigRequest) SetNodeNames(v []*string) *ModifyNodePoolNodeConfigRequest {
	s.NodeNames = v
	return s
}

func (s *ModifyNodePoolNodeConfigRequest) SetOsConfig(v *ModifyNodePoolNodeConfigRequestOsConfig) *ModifyNodePoolNodeConfigRequest {
	s.OsConfig = v
	return s
}

func (s *ModifyNodePoolNodeConfigRequest) SetRollingPolicy(v *ModifyNodePoolNodeConfigRequestRollingPolicy) *ModifyNodePoolNodeConfigRequest {
	s.RollingPolicy = v
	return s
}

func (s *ModifyNodePoolNodeConfigRequest) Validate() error {
	if s.ContainerdConfig != nil {
		if err := s.ContainerdConfig.Validate(); err != nil {
			return err
		}
	}
	if s.KubeletConfig != nil {
		if err := s.KubeletConfig.Validate(); err != nil {
			return err
		}
	}
	if s.OsConfig != nil {
		if err := s.OsConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RollingPolicy != nil {
		if err := s.RollingPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyNodePoolNodeConfigRequestOsConfig struct {
	// The hugepage configuration.
	Hugepage *Hugepage `json:"hugepage,omitempty" xml:"hugepage,omitempty"`
	// The custom sysctl parameter settings.
	Sysctl map[string]interface{} `json:"sysctl,omitempty" xml:"sysctl,omitempty"`
}

func (s ModifyNodePoolNodeConfigRequestOsConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodePoolNodeConfigRequestOsConfig) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolNodeConfigRequestOsConfig) GetHugepage() *Hugepage {
	return s.Hugepage
}

func (s *ModifyNodePoolNodeConfigRequestOsConfig) GetSysctl() map[string]interface{} {
	return s.Sysctl
}

func (s *ModifyNodePoolNodeConfigRequestOsConfig) SetHugepage(v *Hugepage) *ModifyNodePoolNodeConfigRequestOsConfig {
	s.Hugepage = v
	return s
}

func (s *ModifyNodePoolNodeConfigRequestOsConfig) SetSysctl(v map[string]interface{}) *ModifyNodePoolNodeConfigRequestOsConfig {
	s.Sysctl = v
	return s
}

func (s *ModifyNodePoolNodeConfigRequestOsConfig) Validate() error {
	if s.Hugepage != nil {
		if err := s.Hugepage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyNodePoolNodeConfigRequestRollingPolicy struct {
	// The maximum number of nodes that are allowed to fail during the rolling update. Default value: 0, which indicates that the task fails if any node fails. If the value is greater than 0, the task fails and stops when the cumulative number of failed nodes exceeds this value.
	//
	// example:
	//
	// 0
	MaxFailedNodes *int64 `json:"max_failed_nodes,omitempty" xml:"max_failed_nodes,omitempty"`
	// The node updates in the node pool are performed in batches. This parameter specifies the maximum number of nodes that can be updated in parallel per batch.
	//
	// Valid values: [1,10].
	//
	// Default value: 10.
	//
	// example:
	//
	// 3
	MaxParallelism *int64 `json:"max_parallelism,omitempty" xml:"max_parallelism,omitempty"`
}

func (s ModifyNodePoolNodeConfigRequestRollingPolicy) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodePoolNodeConfigRequestRollingPolicy) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolNodeConfigRequestRollingPolicy) GetMaxFailedNodes() *int64 {
	return s.MaxFailedNodes
}

func (s *ModifyNodePoolNodeConfigRequestRollingPolicy) GetMaxParallelism() *int64 {
	return s.MaxParallelism
}

func (s *ModifyNodePoolNodeConfigRequestRollingPolicy) SetMaxFailedNodes(v int64) *ModifyNodePoolNodeConfigRequestRollingPolicy {
	s.MaxFailedNodes = &v
	return s
}

func (s *ModifyNodePoolNodeConfigRequestRollingPolicy) SetMaxParallelism(v int64) *ModifyNodePoolNodeConfigRequestRollingPolicy {
	s.MaxParallelism = &v
	return s
}

func (s *ModifyNodePoolNodeConfigRequestRollingPolicy) Validate() error {
	return dara.Validate(s)
}
