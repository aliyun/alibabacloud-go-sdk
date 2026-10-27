// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodePoolComponentInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *UpdateNodePoolComponentInstanceRequestConfig) *UpdateNodePoolComponentInstanceRequest
	GetConfig() *UpdateNodePoolComponentInstanceRequestConfig
	SetDisableRolling(v bool) *UpdateNodePoolComponentInstanceRequest
	GetDisableRolling() *bool
	SetNodeNames(v []*string) *UpdateNodePoolComponentInstanceRequest
	GetNodeNames() []*string
	SetRollingPolicy(v *UpdateNodePoolComponentInstanceRequestRollingPolicy) *UpdateNodePoolComponentInstanceRequest
	GetRollingPolicy() *UpdateNodePoolComponentInstanceRequestRollingPolicy
	SetVersion(v string) *UpdateNodePoolComponentInstanceRequest
	GetVersion() *string
}

type UpdateNodePoolComponentInstanceRequest struct {
	// The node component configuration.
	Config *UpdateNodePoolComponentInstanceRequestConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// Specifies whether to disable rolling updates. Default value: false. Updating the baseline configuration triggers a rolling update of the nodes.
	//
	// example:
	//
	// false
	DisableRolling *bool `json:"disable_rolling,omitempty" xml:"disable_rolling,omitempty"`
	// The list of nodes to be updated in a rolling manner. Default value: all nodes.
	NodeNames []*string `json:"node_names,omitempty" xml:"node_names,omitempty" type:"Repeated"`
	// The rolling update configuration.
	RollingPolicy *UpdateNodePoolComponentInstanceRequestRollingPolicy `json:"rolling_policy,omitempty" xml:"rolling_policy,omitempty" type:"Struct"`
	// The node component version.
	//
	// example:
	//
	// 1.33.3-aliyun.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpdateNodePoolComponentInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodePoolComponentInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodePoolComponentInstanceRequest) GetConfig() *UpdateNodePoolComponentInstanceRequestConfig {
	return s.Config
}

func (s *UpdateNodePoolComponentInstanceRequest) GetDisableRolling() *bool {
	return s.DisableRolling
}

func (s *UpdateNodePoolComponentInstanceRequest) GetNodeNames() []*string {
	return s.NodeNames
}

func (s *UpdateNodePoolComponentInstanceRequest) GetRollingPolicy() *UpdateNodePoolComponentInstanceRequestRollingPolicy {
	return s.RollingPolicy
}

func (s *UpdateNodePoolComponentInstanceRequest) GetVersion() *string {
	return s.Version
}

func (s *UpdateNodePoolComponentInstanceRequest) SetConfig(v *UpdateNodePoolComponentInstanceRequestConfig) *UpdateNodePoolComponentInstanceRequest {
	s.Config = v
	return s
}

func (s *UpdateNodePoolComponentInstanceRequest) SetDisableRolling(v bool) *UpdateNodePoolComponentInstanceRequest {
	s.DisableRolling = &v
	return s
}

func (s *UpdateNodePoolComponentInstanceRequest) SetNodeNames(v []*string) *UpdateNodePoolComponentInstanceRequest {
	s.NodeNames = v
	return s
}

func (s *UpdateNodePoolComponentInstanceRequest) SetRollingPolicy(v *UpdateNodePoolComponentInstanceRequestRollingPolicy) *UpdateNodePoolComponentInstanceRequest {
	s.RollingPolicy = v
	return s
}

func (s *UpdateNodePoolComponentInstanceRequest) SetVersion(v string) *UpdateNodePoolComponentInstanceRequest {
	s.Version = &v
	return s
}

func (s *UpdateNodePoolComponentInstanceRequest) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
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

type UpdateNodePoolComponentInstanceRequestConfig struct {
	// The custom configuration of the node component.
	//
	// example:
	//
	// {"cpuManagerPolicy":"static"}
	CustomConfig map[string]interface{} `json:"custom_config,omitempty" xml:"custom_config,omitempty"`
	// The environment variables of the node component.
	Envs []*UpdateNodePoolComponentInstanceRequestConfigEnvs `json:"envs,omitempty" xml:"envs,omitempty" type:"Repeated"`
}

func (s UpdateNodePoolComponentInstanceRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodePoolComponentInstanceRequestConfig) GoString() string {
	return s.String()
}

func (s *UpdateNodePoolComponentInstanceRequestConfig) GetCustomConfig() map[string]interface{} {
	return s.CustomConfig
}

func (s *UpdateNodePoolComponentInstanceRequestConfig) GetEnvs() []*UpdateNodePoolComponentInstanceRequestConfigEnvs {
	return s.Envs
}

func (s *UpdateNodePoolComponentInstanceRequestConfig) SetCustomConfig(v map[string]interface{}) *UpdateNodePoolComponentInstanceRequestConfig {
	s.CustomConfig = v
	return s
}

func (s *UpdateNodePoolComponentInstanceRequestConfig) SetEnvs(v []*UpdateNodePoolComponentInstanceRequestConfigEnvs) *UpdateNodePoolComponentInstanceRequestConfig {
	s.Envs = v
	return s
}

func (s *UpdateNodePoolComponentInstanceRequestConfig) Validate() error {
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

type UpdateNodePoolComponentInstanceRequestConfigEnvs struct {
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

func (s UpdateNodePoolComponentInstanceRequestConfigEnvs) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodePoolComponentInstanceRequestConfigEnvs) GoString() string {
	return s.String()
}

func (s *UpdateNodePoolComponentInstanceRequestConfigEnvs) GetName() *string {
	return s.Name
}

func (s *UpdateNodePoolComponentInstanceRequestConfigEnvs) GetValue() *string {
	return s.Value
}

func (s *UpdateNodePoolComponentInstanceRequestConfigEnvs) SetName(v string) *UpdateNodePoolComponentInstanceRequestConfigEnvs {
	s.Name = &v
	return s
}

func (s *UpdateNodePoolComponentInstanceRequestConfigEnvs) SetValue(v string) *UpdateNodePoolComponentInstanceRequestConfigEnvs {
	s.Value = &v
	return s
}

func (s *UpdateNodePoolComponentInstanceRequestConfigEnvs) Validate() error {
	return dara.Validate(s)
}

type UpdateNodePoolComponentInstanceRequestRollingPolicy struct {
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
	// The maximum number of parallel nodes per batch. Default value: 1.
	//
	// example:
	//
	// 1
	MaxParallelism *int64 `json:"max_parallelism,omitempty" xml:"max_parallelism,omitempty"`
	// The automatic pause policy during node upgrades. Valid values: NotPause, FirstBatch, and EveryBatch. Default value: NotPause.
	//
	// example:
	//
	// NotPause
	PausePolicy *string `json:"pause_policy,omitempty" xml:"pause_policy,omitempty"`
}

func (s UpdateNodePoolComponentInstanceRequestRollingPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodePoolComponentInstanceRequestRollingPolicy) GoString() string {
	return s.String()
}

func (s *UpdateNodePoolComponentInstanceRequestRollingPolicy) GetBatchInterval() *int64 {
	return s.BatchInterval
}

func (s *UpdateNodePoolComponentInstanceRequestRollingPolicy) GetMaxFailedNodes() *int64 {
	return s.MaxFailedNodes
}

func (s *UpdateNodePoolComponentInstanceRequestRollingPolicy) GetMaxParallelism() *int64 {
	return s.MaxParallelism
}

func (s *UpdateNodePoolComponentInstanceRequestRollingPolicy) GetPausePolicy() *string {
	return s.PausePolicy
}

func (s *UpdateNodePoolComponentInstanceRequestRollingPolicy) SetBatchInterval(v int64) *UpdateNodePoolComponentInstanceRequestRollingPolicy {
	s.BatchInterval = &v
	return s
}

func (s *UpdateNodePoolComponentInstanceRequestRollingPolicy) SetMaxFailedNodes(v int64) *UpdateNodePoolComponentInstanceRequestRollingPolicy {
	s.MaxFailedNodes = &v
	return s
}

func (s *UpdateNodePoolComponentInstanceRequestRollingPolicy) SetMaxParallelism(v int64) *UpdateNodePoolComponentInstanceRequestRollingPolicy {
	s.MaxParallelism = &v
	return s
}

func (s *UpdateNodePoolComponentInstanceRequestRollingPolicy) SetPausePolicy(v string) *UpdateNodePoolComponentInstanceRequestRollingPolicy {
	s.PausePolicy = &v
	return s
}

func (s *UpdateNodePoolComponentInstanceRequestRollingPolicy) Validate() error {
	return dara.Validate(s)
}
