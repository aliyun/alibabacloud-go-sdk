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
	// This parameter is required.
	Components    []*CreateNodePoolComponentInstancesRequestComponents  `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	NodeNames     []*string                                             `json:"node_names,omitempty" xml:"node_names,omitempty" type:"Repeated"`
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
	Config *CreateNodePoolComponentInstancesRequestComponentsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// kubelet
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
	// example:
	//
	// {"cpuManagerPolicy":"static"}
	CustomConfig map[string]interface{} `json:"custom_config,omitempty" xml:"custom_config,omitempty"`
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

func (s *CreateNodePoolComponentInstancesRequestComponentsConfig) SetCustomConfig(v map[string]interface{}) *CreateNodePoolComponentInstancesRequestComponentsConfig {
	s.CustomConfig = v
	return s
}

func (s *CreateNodePoolComponentInstancesRequestComponentsConfig) Validate() error {
	return dara.Validate(s)
}

type CreateNodePoolComponentInstancesRequestRollingPolicy struct {
	// example:
	//
	// 10
	BatchInterval *int64 `json:"batch_interval,omitempty" xml:"batch_interval,omitempty"`
	// example:
	//
	// 0
	MaxFailedNodes *int64 `json:"max_failed_nodes,omitempty" xml:"max_failed_nodes,omitempty"`
	// example:
	//
	// 1
	MaxParallelism *int64 `json:"max_parallelism,omitempty" xml:"max_parallelism,omitempty"`
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
