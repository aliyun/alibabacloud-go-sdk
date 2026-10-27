// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodePoolComponentInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponents(v []*CreateNodePoolComponentInstancesRequestComponents) *CreateNodePoolComponentInstancesRequest
	GetComponents() []*CreateNodePoolComponentInstancesRequestComponents
	SetNodeNames(v []*string) *CreateNodePoolComponentInstancesRequest
	GetNodeNames() []*string
	SetRollingPolicy(v *CreateNodePoolComponentInstancesRequestRollingPolicy) *CreateNodePoolComponentInstancesRequest
	GetRollingPolicy() *CreateNodePoolComponentInstancesRequestRollingPolicy
}

type CreateNodePoolComponentInstancesRequest struct {
	// The list of node components.
	//
	// This parameter is required.
	Components []*CreateNodePoolComponentInstancesRequestComponents `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	// The list of nodes specified for the rolling update. Default value: all nodes.
	NodeNames []*string `json:"node_names,omitempty" xml:"node_names,omitempty" type:"Repeated"`
	// The rolling update configuration.
	RollingPolicy *CreateNodePoolComponentInstancesRequestRollingPolicy `json:"rolling_policy,omitempty" xml:"rolling_policy,omitempty" type:"Struct"`
}

func (s CreateNodePoolComponentInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNodePoolComponentInstancesRequest) GoString() string {
	return s.String()
}

func (s *CreateNodePoolComponentInstancesRequest) GetComponents() []*CreateNodePoolComponentInstancesRequestComponents {
	return s.Components
}

func (s *CreateNodePoolComponentInstancesRequest) GetNodeNames() []*string {
	return s.NodeNames
}

func (s *CreateNodePoolComponentInstancesRequest) GetRollingPolicy() *CreateNodePoolComponentInstancesRequestRollingPolicy {
	return s.RollingPolicy
}

func (s *CreateNodePoolComponentInstancesRequest) SetComponents(v []*CreateNodePoolComponentInstancesRequestComponents) *CreateNodePoolComponentInstancesRequest {
	s.Components = v
	return s
}

func (s *CreateNodePoolComponentInstancesRequest) SetNodeNames(v []*string) *CreateNodePoolComponentInstancesRequest {
	s.NodeNames = v
	return s
}

func (s *CreateNodePoolComponentInstancesRequest) SetRollingPolicy(v *CreateNodePoolComponentInstancesRequestRollingPolicy) *CreateNodePoolComponentInstancesRequest {
	s.RollingPolicy = v
	return s
}

func (s *CreateNodePoolComponentInstancesRequest) Validate() error {
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

type CreateNodePoolComponentInstancesRequestComponents struct {
	// The configuration of the component.
	Config *CreateNodePoolComponentInstancesRequestComponentsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The name of the component.
	//
	// This parameter is required.
	//
	// example:
	//
	// kubelet
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version of the component.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.33.3-aliyun.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateNodePoolComponentInstancesRequestComponents) String() string {
	return dara.Prettify(s)
}

func (s CreateNodePoolComponentInstancesRequestComponents) GoString() string {
	return s.String()
}

func (s *CreateNodePoolComponentInstancesRequestComponents) GetConfig() *CreateNodePoolComponentInstancesRequestComponentsConfig {
	return s.Config
}

func (s *CreateNodePoolComponentInstancesRequestComponents) GetName() *string {
	return s.Name
}

func (s *CreateNodePoolComponentInstancesRequestComponents) GetVersion() *string {
	return s.Version
}

func (s *CreateNodePoolComponentInstancesRequestComponents) SetConfig(v *CreateNodePoolComponentInstancesRequestComponentsConfig) *CreateNodePoolComponentInstancesRequestComponents {
	s.Config = v
	return s
}

func (s *CreateNodePoolComponentInstancesRequestComponents) SetName(v string) *CreateNodePoolComponentInstancesRequestComponents {
	s.Name = &v
	return s
}

func (s *CreateNodePoolComponentInstancesRequestComponents) SetVersion(v string) *CreateNodePoolComponentInstancesRequestComponents {
	s.Version = &v
	return s
}

func (s *CreateNodePoolComponentInstancesRequestComponents) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateNodePoolComponentInstancesRequestComponentsConfig struct {
	// The custom configuration of the component.
	//
	// example:
	//
	// {"cpuManagerPolicy":"static"}
	CustomConfig map[string]interface{} `json:"custom_config,omitempty" xml:"custom_config,omitempty"`
	// The environment variables of the node component.
	Envs []*CreateNodePoolComponentInstancesRequestComponentsConfigEnvs `json:"envs,omitempty" xml:"envs,omitempty" type:"Repeated"`
}

func (s CreateNodePoolComponentInstancesRequestComponentsConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateNodePoolComponentInstancesRequestComponentsConfig) GoString() string {
	return s.String()
}

func (s *CreateNodePoolComponentInstancesRequestComponentsConfig) GetCustomConfig() map[string]interface{} {
	return s.CustomConfig
}

func (s *CreateNodePoolComponentInstancesRequestComponentsConfig) GetEnvs() []*CreateNodePoolComponentInstancesRequestComponentsConfigEnvs {
	return s.Envs
}

func (s *CreateNodePoolComponentInstancesRequestComponentsConfig) SetCustomConfig(v map[string]interface{}) *CreateNodePoolComponentInstancesRequestComponentsConfig {
	s.CustomConfig = v
	return s
}

func (s *CreateNodePoolComponentInstancesRequestComponentsConfig) SetEnvs(v []*CreateNodePoolComponentInstancesRequestComponentsConfigEnvs) *CreateNodePoolComponentInstancesRequestComponentsConfig {
	s.Envs = v
	return s
}

func (s *CreateNodePoolComponentInstancesRequestComponentsConfig) Validate() error {
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

type CreateNodePoolComponentInstancesRequestComponentsConfigEnvs struct {
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

func (s CreateNodePoolComponentInstancesRequestComponentsConfigEnvs) String() string {
	return dara.Prettify(s)
}

func (s CreateNodePoolComponentInstancesRequestComponentsConfigEnvs) GoString() string {
	return s.String()
}

func (s *CreateNodePoolComponentInstancesRequestComponentsConfigEnvs) GetName() *string {
	return s.Name
}

func (s *CreateNodePoolComponentInstancesRequestComponentsConfigEnvs) GetValue() *string {
	return s.Value
}

func (s *CreateNodePoolComponentInstancesRequestComponentsConfigEnvs) SetName(v string) *CreateNodePoolComponentInstancesRequestComponentsConfigEnvs {
	s.Name = &v
	return s
}

func (s *CreateNodePoolComponentInstancesRequestComponentsConfigEnvs) SetValue(v string) *CreateNodePoolComponentInstancesRequestComponentsConfigEnvs {
	s.Value = &v
	return s
}

func (s *CreateNodePoolComponentInstancesRequestComponentsConfigEnvs) Validate() error {
	return dara.Validate(s)
}

type CreateNodePoolComponentInstancesRequestRollingPolicy struct {
	// The upgrade interval between batches, in seconds. Default value: 0.
	//
	// example:
	//
	// 10
	BatchInterval *int64 `json:"batch_interval,omitempty" xml:"batch_interval,omitempty"`
	// The maximum number of nodes allowed to fail during the rolling update. Default value: 0. A value of 0 indicates that the task fails if any node fails. A value greater than 0 indicates that the task fails and stops if the cumulative number of failed nodes exceeds this value.
	//
	// example:
	//
	// 0
	MaxFailedNodes *int64 `json:"max_failed_nodes,omitempty" xml:"max_failed_nodes,omitempty"`
	// The maximum number of parallel updates per batch. Default value: 1.
	//
	// example:
	//
	// 1
	MaxParallelism *int64 `json:"max_parallelism,omitempty" xml:"max_parallelism,omitempty"`
	// The auto-pause policy during the node upgrade process. Valid values: NotPause, FirstBatch, and EveryBatch. Default value: NotPause.
	//
	// example:
	//
	// NotPause
	PausePolicy *string `json:"pause_policy,omitempty" xml:"pause_policy,omitempty"`
}

func (s CreateNodePoolComponentInstancesRequestRollingPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateNodePoolComponentInstancesRequestRollingPolicy) GoString() string {
	return s.String()
}

func (s *CreateNodePoolComponentInstancesRequestRollingPolicy) GetBatchInterval() *int64 {
	return s.BatchInterval
}

func (s *CreateNodePoolComponentInstancesRequestRollingPolicy) GetMaxFailedNodes() *int64 {
	return s.MaxFailedNodes
}

func (s *CreateNodePoolComponentInstancesRequestRollingPolicy) GetMaxParallelism() *int64 {
	return s.MaxParallelism
}

func (s *CreateNodePoolComponentInstancesRequestRollingPolicy) GetPausePolicy() *string {
	return s.PausePolicy
}

func (s *CreateNodePoolComponentInstancesRequestRollingPolicy) SetBatchInterval(v int64) *CreateNodePoolComponentInstancesRequestRollingPolicy {
	s.BatchInterval = &v
	return s
}

func (s *CreateNodePoolComponentInstancesRequestRollingPolicy) SetMaxFailedNodes(v int64) *CreateNodePoolComponentInstancesRequestRollingPolicy {
	s.MaxFailedNodes = &v
	return s
}

func (s *CreateNodePoolComponentInstancesRequestRollingPolicy) SetMaxParallelism(v int64) *CreateNodePoolComponentInstancesRequestRollingPolicy {
	s.MaxParallelism = &v
	return s
}

func (s *CreateNodePoolComponentInstancesRequestRollingPolicy) SetPausePolicy(v string) *CreateNodePoolComponentInstancesRequestRollingPolicy {
	s.PausePolicy = &v
	return s
}

func (s *CreateNodePoolComponentInstancesRequestRollingPolicy) Validate() error {
	return dara.Validate(s)
}
