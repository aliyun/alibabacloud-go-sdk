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
	SetOsConfig(v *ModifyNodePoolNodeConfigRequestOsConfig) *ModifyNodePoolNodeConfigRequest
	GetOsConfig() *ModifyNodePoolNodeConfigRequestOsConfig
	SetRollingPolicy(v *ModifyNodePoolNodeConfigRequestRollingPolicy) *ModifyNodePoolNodeConfigRequest
	GetRollingPolicy() *ModifyNodePoolNodeConfigRequestRollingPolicy
}

type ModifyNodePoolNodeConfigRequest struct {
	// The containerd runtime configuration.
	ContainerdConfig *ContainerdConfig `json:"containerd_config,omitempty" xml:"containerd_config,omitempty"`
	// The kubelet configurations.
	KubeletConfig *KubeletConfig `json:"kubelet_config,omitempty" xml:"kubelet_config,omitempty"`
	// The OS configuration.
	OsConfig *ModifyNodePoolNodeConfigRequestOsConfig `json:"os_config,omitempty" xml:"os_config,omitempty" type:"Struct"`
	// The rolling policy configuration.
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
	Hugepage *Hugepage `json:"hugepage,omitempty" xml:"hugepage,omitempty"`
	// The sysctl configuration.
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
	// The maximum number of unavailable nodes.
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

func (s *ModifyNodePoolNodeConfigRequestRollingPolicy) GetMaxParallelism() *int64 {
	return s.MaxParallelism
}

func (s *ModifyNodePoolNodeConfigRequestRollingPolicy) SetMaxParallelism(v int64) *ModifyNodePoolNodeConfigRequestRollingPolicy {
	s.MaxParallelism = &v
	return s
}

func (s *ModifyNodePoolNodeConfigRequestRollingPolicy) Validate() error {
	return dara.Validate(s)
}
