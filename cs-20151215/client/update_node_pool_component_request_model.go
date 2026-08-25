// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodePoolComponentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *UpdateNodePoolComponentRequestConfig) *UpdateNodePoolComponentRequest
	GetConfig() *UpdateNodePoolComponentRequestConfig
	SetDisableRolling(v bool) *UpdateNodePoolComponentRequest
	GetDisableRolling() *bool
	SetName(v string) *UpdateNodePoolComponentRequest
	GetName() *string
	SetNodeNames(v []*string) *UpdateNodePoolComponentRequest
	GetNodeNames() []*string
	SetRollingPolicy(v *UpdateNodePoolComponentRequestRollingPolicy) *UpdateNodePoolComponentRequest
	GetRollingPolicy() *UpdateNodePoolComponentRequestRollingPolicy
	SetVersion(v string) *UpdateNodePoolComponentRequest
	GetVersion() *string
}

type UpdateNodePoolComponentRequest struct {
	// The configuration of the node component.
	Config *UpdateNodePoolComponentRequestConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// Specifies whether to disable log rotation. Default value: false. Updating the baseline configuration triggers log rotation on nodes.
	DisableRolling *bool `json:"disableRolling,omitempty" xml:"disableRolling,omitempty"`
	// The name of the node component.
	//
	// example:
	//
	// kubelet
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The list of nodes for log rotation. By default, all nodes are included.
	NodeNames []*string `json:"nodeNames,omitempty" xml:"nodeNames,omitempty" type:"Repeated"`
	// The log rotation configuration.
	RollingPolicy *UpdateNodePoolComponentRequestRollingPolicy `json:"rollingPolicy,omitempty" xml:"rollingPolicy,omitempty" type:"Struct"`
	// The version of the node component.
	//
	// example:
	//
	// 1.28.9-aliyun.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpdateNodePoolComponentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodePoolComponentRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodePoolComponentRequest) GetConfig() *UpdateNodePoolComponentRequestConfig {
	return s.Config
}

func (s *UpdateNodePoolComponentRequest) GetDisableRolling() *bool {
	return s.DisableRolling
}

func (s *UpdateNodePoolComponentRequest) GetName() *string {
	return s.Name
}

func (s *UpdateNodePoolComponentRequest) GetNodeNames() []*string {
	return s.NodeNames
}

func (s *UpdateNodePoolComponentRequest) GetRollingPolicy() *UpdateNodePoolComponentRequestRollingPolicy {
	return s.RollingPolicy
}

func (s *UpdateNodePoolComponentRequest) GetVersion() *string {
	return s.Version
}

func (s *UpdateNodePoolComponentRequest) SetConfig(v *UpdateNodePoolComponentRequestConfig) *UpdateNodePoolComponentRequest {
	s.Config = v
	return s
}

func (s *UpdateNodePoolComponentRequest) SetDisableRolling(v bool) *UpdateNodePoolComponentRequest {
	s.DisableRolling = &v
	return s
}

func (s *UpdateNodePoolComponentRequest) SetName(v string) *UpdateNodePoolComponentRequest {
	s.Name = &v
	return s
}

func (s *UpdateNodePoolComponentRequest) SetNodeNames(v []*string) *UpdateNodePoolComponentRequest {
	s.NodeNames = v
	return s
}

func (s *UpdateNodePoolComponentRequest) SetRollingPolicy(v *UpdateNodePoolComponentRequestRollingPolicy) *UpdateNodePoolComponentRequest {
	s.RollingPolicy = v
	return s
}

func (s *UpdateNodePoolComponentRequest) SetVersion(v string) *UpdateNodePoolComponentRequest {
	s.Version = &v
	return s
}

func (s *UpdateNodePoolComponentRequest) Validate() error {
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

type UpdateNodePoolComponentRequestConfig struct {
	// The custom configuration of the component.
	//
	// example:
	//
	// {"cpuManagerPolicy":"static"}
	CustomConfig map[string]interface{} `json:"customConfig,omitempty" xml:"customConfig,omitempty"`
	// The environment variables of the node component.
	Envs []*UpdateNodePoolComponentRequestConfigEnvs `json:"envs,omitempty" xml:"envs,omitempty" type:"Repeated"`
}

func (s UpdateNodePoolComponentRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodePoolComponentRequestConfig) GoString() string {
	return s.String()
}

func (s *UpdateNodePoolComponentRequestConfig) GetCustomConfig() map[string]interface{} {
	return s.CustomConfig
}

func (s *UpdateNodePoolComponentRequestConfig) GetEnvs() []*UpdateNodePoolComponentRequestConfigEnvs {
	return s.Envs
}

func (s *UpdateNodePoolComponentRequestConfig) SetCustomConfig(v map[string]interface{}) *UpdateNodePoolComponentRequestConfig {
	s.CustomConfig = v
	return s
}

func (s *UpdateNodePoolComponentRequestConfig) SetEnvs(v []*UpdateNodePoolComponentRequestConfigEnvs) *UpdateNodePoolComponentRequestConfig {
	s.Envs = v
	return s
}

func (s *UpdateNodePoolComponentRequestConfig) Validate() error {
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

type UpdateNodePoolComponentRequestConfigEnvs struct {
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

func (s UpdateNodePoolComponentRequestConfigEnvs) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodePoolComponentRequestConfigEnvs) GoString() string {
	return s.String()
}

func (s *UpdateNodePoolComponentRequestConfigEnvs) GetName() *string {
	return s.Name
}

func (s *UpdateNodePoolComponentRequestConfigEnvs) GetValue() *string {
	return s.Value
}

func (s *UpdateNodePoolComponentRequestConfigEnvs) SetName(v string) *UpdateNodePoolComponentRequestConfigEnvs {
	s.Name = &v
	return s
}

func (s *UpdateNodePoolComponentRequestConfigEnvs) SetValue(v string) *UpdateNodePoolComponentRequestConfigEnvs {
	s.Value = &v
	return s
}

func (s *UpdateNodePoolComponentRequestConfigEnvs) Validate() error {
	return dara.Validate(s)
}

type UpdateNodePoolComponentRequestRollingPolicy struct {
	// The upgrade interval between batches. Unit: seconds.
	//
	// example:
	//
	// 0
	BatchInterval *int64 `json:"batchInterval,omitempty" xml:"batchInterval,omitempty"`
	// The maximum number of nodes that can fail during the rolling update. Default value: 0, which means the task fails if any node fails. If the value is greater than 0, the task fails and stops when the cumulative number of failed nodes exceeds this value.
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
	// The automatic pause policy during node upgrade.
	//
	// example:
	//
	// NotPause
	PausePolicy *string `json:"pausePolicy,omitempty" xml:"pausePolicy,omitempty"`
}

func (s UpdateNodePoolComponentRequestRollingPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodePoolComponentRequestRollingPolicy) GoString() string {
	return s.String()
}

func (s *UpdateNodePoolComponentRequestRollingPolicy) GetBatchInterval() *int64 {
	return s.BatchInterval
}

func (s *UpdateNodePoolComponentRequestRollingPolicy) GetMaxFailedNodes() *int64 {
	return s.MaxFailedNodes
}

func (s *UpdateNodePoolComponentRequestRollingPolicy) GetMaxParallelism() *int64 {
	return s.MaxParallelism
}

func (s *UpdateNodePoolComponentRequestRollingPolicy) GetPausePolicy() *string {
	return s.PausePolicy
}

func (s *UpdateNodePoolComponentRequestRollingPolicy) SetBatchInterval(v int64) *UpdateNodePoolComponentRequestRollingPolicy {
	s.BatchInterval = &v
	return s
}

func (s *UpdateNodePoolComponentRequestRollingPolicy) SetMaxFailedNodes(v int64) *UpdateNodePoolComponentRequestRollingPolicy {
	s.MaxFailedNodes = &v
	return s
}

func (s *UpdateNodePoolComponentRequestRollingPolicy) SetMaxParallelism(v int64) *UpdateNodePoolComponentRequestRollingPolicy {
	s.MaxParallelism = &v
	return s
}

func (s *UpdateNodePoolComponentRequestRollingPolicy) SetPausePolicy(v string) *UpdateNodePoolComponentRequestRollingPolicy {
	s.PausePolicy = &v
	return s
}

func (s *UpdateNodePoolComponentRequestRollingPolicy) Validate() error {
	return dara.Validate(s)
}
