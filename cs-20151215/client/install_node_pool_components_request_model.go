// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallNodePoolComponentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponents(v []*InstallNodePoolComponentsRequestComponents) *InstallNodePoolComponentsRequest
	GetComponents() []*InstallNodePoolComponentsRequestComponents
	SetNodeNames(v []*string) *InstallNodePoolComponentsRequest
	GetNodeNames() []*string
	SetRollingPolicy(v *InstallNodePoolComponentsRequestRollingPolicy) *InstallNodePoolComponentsRequest
	GetRollingPolicy() *InstallNodePoolComponentsRequestRollingPolicy
}

type InstallNodePoolComponentsRequest struct {
	// The list of node components.
	Components []*InstallNodePoolComponentsRequestComponents `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	// The list of node names for the rolling operation. By default, all nodes are included.
	NodeNames []*string `json:"nodeNames,omitempty" xml:"nodeNames,omitempty" type:"Repeated"`
	// The rolling policy configuration.
	RollingPolicy *InstallNodePoolComponentsRequestRollingPolicy `json:"rollingPolicy,omitempty" xml:"rollingPolicy,omitempty" type:"Struct"`
}

func (s InstallNodePoolComponentsRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallNodePoolComponentsRequest) GoString() string {
	return s.String()
}

func (s *InstallNodePoolComponentsRequest) GetComponents() []*InstallNodePoolComponentsRequestComponents {
	return s.Components
}

func (s *InstallNodePoolComponentsRequest) GetNodeNames() []*string {
	return s.NodeNames
}

func (s *InstallNodePoolComponentsRequest) GetRollingPolicy() *InstallNodePoolComponentsRequestRollingPolicy {
	return s.RollingPolicy
}

func (s *InstallNodePoolComponentsRequest) SetComponents(v []*InstallNodePoolComponentsRequestComponents) *InstallNodePoolComponentsRequest {
	s.Components = v
	return s
}

func (s *InstallNodePoolComponentsRequest) SetNodeNames(v []*string) *InstallNodePoolComponentsRequest {
	s.NodeNames = v
	return s
}

func (s *InstallNodePoolComponentsRequest) SetRollingPolicy(v *InstallNodePoolComponentsRequestRollingPolicy) *InstallNodePoolComponentsRequest {
	s.RollingPolicy = v
	return s
}

func (s *InstallNodePoolComponentsRequest) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RollingPolicy != nil {
		if err := s.RollingPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstallNodePoolComponentsRequestComponents struct {
	// The component configuration.
	Config *InstallNodePoolComponentsRequestComponentsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The component name.
	//
	// This parameter is required.
	//
	// example:
	//
	// kubelet
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The component version.
	//
	// example:
	//
	// 1.28.9-aliyun.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s InstallNodePoolComponentsRequestComponents) String() string {
	return dara.Prettify(s)
}

func (s InstallNodePoolComponentsRequestComponents) GoString() string {
	return s.String()
}

func (s *InstallNodePoolComponentsRequestComponents) GetConfig() *InstallNodePoolComponentsRequestComponentsConfig {
	return s.Config
}

func (s *InstallNodePoolComponentsRequestComponents) GetName() *string {
	return s.Name
}

func (s *InstallNodePoolComponentsRequestComponents) GetVersion() *string {
	return s.Version
}

func (s *InstallNodePoolComponentsRequestComponents) SetConfig(v *InstallNodePoolComponentsRequestComponentsConfig) *InstallNodePoolComponentsRequestComponents {
	s.Config = v
	return s
}

func (s *InstallNodePoolComponentsRequestComponents) SetName(v string) *InstallNodePoolComponentsRequestComponents {
	s.Name = &v
	return s
}

func (s *InstallNodePoolComponentsRequestComponents) SetVersion(v string) *InstallNodePoolComponentsRequestComponents {
	s.Version = &v
	return s
}

func (s *InstallNodePoolComponentsRequestComponents) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstallNodePoolComponentsRequestComponentsConfig struct {
	// The custom configuration of the component.
	//
	// example:
	//
	// {"cpuManagerPolicy":"static"}
	CustomConfig map[string]interface{} `json:"customConfig,omitempty" xml:"customConfig,omitempty"`
	// The environment variables of the node component.
	Envs []*InstallNodePoolComponentsRequestComponentsConfigEnvs `json:"envs,omitempty" xml:"envs,omitempty" type:"Repeated"`
}

func (s InstallNodePoolComponentsRequestComponentsConfig) String() string {
	return dara.Prettify(s)
}

func (s InstallNodePoolComponentsRequestComponentsConfig) GoString() string {
	return s.String()
}

func (s *InstallNodePoolComponentsRequestComponentsConfig) GetCustomConfig() map[string]interface{} {
	return s.CustomConfig
}

func (s *InstallNodePoolComponentsRequestComponentsConfig) GetEnvs() []*InstallNodePoolComponentsRequestComponentsConfigEnvs {
	return s.Envs
}

func (s *InstallNodePoolComponentsRequestComponentsConfig) SetCustomConfig(v map[string]interface{}) *InstallNodePoolComponentsRequestComponentsConfig {
	s.CustomConfig = v
	return s
}

func (s *InstallNodePoolComponentsRequestComponentsConfig) SetEnvs(v []*InstallNodePoolComponentsRequestComponentsConfigEnvs) *InstallNodePoolComponentsRequestComponentsConfig {
	s.Envs = v
	return s
}

func (s *InstallNodePoolComponentsRequestComponentsConfig) Validate() error {
	if s.Envs != nil {
		for _, item := range s.Envs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InstallNodePoolComponentsRequestComponentsConfigEnvs struct {
	// The name of the environment variable.
	//
	// example:
	//
	// LOG_LEVEL
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the environment variable.
	//
	// example:
	//
	// info
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s InstallNodePoolComponentsRequestComponentsConfigEnvs) String() string {
	return dara.Prettify(s)
}

func (s InstallNodePoolComponentsRequestComponentsConfigEnvs) GoString() string {
	return s.String()
}

func (s *InstallNodePoolComponentsRequestComponentsConfigEnvs) GetName() *string {
	return s.Name
}

func (s *InstallNodePoolComponentsRequestComponentsConfigEnvs) GetValue() *string {
	return s.Value
}

func (s *InstallNodePoolComponentsRequestComponentsConfigEnvs) SetName(v string) *InstallNodePoolComponentsRequestComponentsConfigEnvs {
	s.Name = &v
	return s
}

func (s *InstallNodePoolComponentsRequestComponentsConfigEnvs) SetValue(v string) *InstallNodePoolComponentsRequestComponentsConfigEnvs {
	s.Value = &v
	return s
}

func (s *InstallNodePoolComponentsRequestComponentsConfigEnvs) Validate() error {
	return dara.Validate(s)
}

type InstallNodePoolComponentsRequestRollingPolicy struct {
	// The upgrade interval between batches. Unit: seconds.
	//
	// example:
	//
	// 0
	BatchInterval *int64 `json:"batchInterval,omitempty" xml:"batchInterval,omitempty"`
	// The maximum number of nodes that are allowed to fail during the rolling process. Default value: 0, which indicates that the task is considered failed if any node fails. If the value is greater than 0, the task is considered failed and stops when the cumulative number of failed nodes exceeds this value.
	//
	// example:
	//
	// 0
	MaxFailedNodes *int64 `json:"maxFailedNodes,omitempty" xml:"maxFailedNodes,omitempty"`
	// The maximum number of parallel operations per batch. Default value: 1.
	//
	// example:
	//
	// 1
	MaxParallelism *int64 `json:"maxParallelism,omitempty" xml:"maxParallelism,omitempty"`
	// The automatic pause policy during the node upgrade process.
	//
	// example:
	//
	// NotPause
	PausePolicy *string `json:"pausePolicy,omitempty" xml:"pausePolicy,omitempty"`
}

func (s InstallNodePoolComponentsRequestRollingPolicy) String() string {
	return dara.Prettify(s)
}

func (s InstallNodePoolComponentsRequestRollingPolicy) GoString() string {
	return s.String()
}

func (s *InstallNodePoolComponentsRequestRollingPolicy) GetBatchInterval() *int64 {
	return s.BatchInterval
}

func (s *InstallNodePoolComponentsRequestRollingPolicy) GetMaxFailedNodes() *int64 {
	return s.MaxFailedNodes
}

func (s *InstallNodePoolComponentsRequestRollingPolicy) GetMaxParallelism() *int64 {
	return s.MaxParallelism
}

func (s *InstallNodePoolComponentsRequestRollingPolicy) GetPausePolicy() *string {
	return s.PausePolicy
}

func (s *InstallNodePoolComponentsRequestRollingPolicy) SetBatchInterval(v int64) *InstallNodePoolComponentsRequestRollingPolicy {
	s.BatchInterval = &v
	return s
}

func (s *InstallNodePoolComponentsRequestRollingPolicy) SetMaxFailedNodes(v int64) *InstallNodePoolComponentsRequestRollingPolicy {
	s.MaxFailedNodes = &v
	return s
}

func (s *InstallNodePoolComponentsRequestRollingPolicy) SetMaxParallelism(v int64) *InstallNodePoolComponentsRequestRollingPolicy {
	s.MaxParallelism = &v
	return s
}

func (s *InstallNodePoolComponentsRequestRollingPolicy) SetPausePolicy(v string) *InstallNodePoolComponentsRequestRollingPolicy {
	s.PausePolicy = &v
	return s
}

func (s *InstallNodePoolComponentsRequestRollingPolicy) Validate() error {
	return dara.Validate(s)
}
