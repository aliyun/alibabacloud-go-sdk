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
	// The configuration details for the component update.
	Config *UpdateNodePoolComponentRequestConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// Specifies whether to disable rolling updates. Default: false. If set to false, nodes will be rolled automatically to apply the new configuration.
	DisableRolling *bool `json:"disableRolling,omitempty" xml:"disableRolling,omitempty"`
	// The name of the node component.
	//
	// example:
	//
	// kubelet
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// A list of specific nodes to be rolled. If not specified, all nodes in the node pool will be updated.
	NodeNames []*string `json:"nodeNames,omitempty" xml:"nodeNames,omitempty" type:"Repeated"`
	// The rolling update configuration.
	RollingPolicy *UpdateNodePoolComponentRequestRollingPolicy `json:"rollingPolicy,omitempty" xml:"rollingPolicy,omitempty" type:"Struct"`
	// The version of the node component to be updated to.
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
	// Custom configuration parameters for the component.
	CustomConfig map[string]*string `json:"customConfig,omitempty" xml:"customConfig,omitempty"`
}

func (s UpdateNodePoolComponentRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodePoolComponentRequestConfig) GoString() string {
	return s.String()
}

func (s *UpdateNodePoolComponentRequestConfig) GetCustomConfig() map[string]*string {
	return s.CustomConfig
}

func (s *UpdateNodePoolComponentRequestConfig) SetCustomConfig(v map[string]*string) *UpdateNodePoolComponentRequestConfig {
	s.CustomConfig = v
	return s
}

func (s *UpdateNodePoolComponentRequestConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateNodePoolComponentRequestRollingPolicy struct {
	// The time interval between update batches, in seconds.
	//
	// example:
	//
	// 0
	BatchInterval *int64 `json:"batchInterval,omitempty" xml:"batchInterval,omitempty"`
	// The maximum number of nodes that can be updated concurrently. Default: 1.
	//
	// example:
	//
	// 1
	MaxParallelism *int64 `json:"maxParallelism,omitempty" xml:"maxParallelism,omitempty"`
	// The automatic pause strategy during the update process. Valid values: NotPause, FirstBatchPause, EveryBatchPause.
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
