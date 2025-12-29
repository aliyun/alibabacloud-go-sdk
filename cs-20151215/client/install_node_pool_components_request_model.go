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
	Components    []*InstallNodePoolComponentsRequestComponents  `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	NodeNames     []*string                                      `json:"nodeNames,omitempty" xml:"nodeNames,omitempty" type:"Repeated"`
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
	Config *InstallNodePoolComponentsRequestComponentsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// kubelet
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
	CustomConfig map[string]*string `json:"customConfig,omitempty" xml:"customConfig,omitempty"`
}

func (s InstallNodePoolComponentsRequestComponentsConfig) String() string {
	return dara.Prettify(s)
}

func (s InstallNodePoolComponentsRequestComponentsConfig) GoString() string {
	return s.String()
}

func (s *InstallNodePoolComponentsRequestComponentsConfig) GetCustomConfig() map[string]*string {
	return s.CustomConfig
}

func (s *InstallNodePoolComponentsRequestComponentsConfig) SetCustomConfig(v map[string]*string) *InstallNodePoolComponentsRequestComponentsConfig {
	s.CustomConfig = v
	return s
}

func (s *InstallNodePoolComponentsRequestComponentsConfig) Validate() error {
	return dara.Validate(s)
}

type InstallNodePoolComponentsRequestRollingPolicy struct {
	// example:
	//
	// 0
	BatchInterval *int64 `json:"batchInterval,omitempty" xml:"batchInterval,omitempty"`
	// example:
	//
	// 1
	MaxParallelism *int64 `json:"maxParallelism,omitempty" xml:"maxParallelism,omitempty"`
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
