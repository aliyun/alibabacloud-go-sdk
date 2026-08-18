// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodePoolComponentInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComponentInstances(v []*ListNodePoolComponentInstancesResponseBodyComponentInstances) *ListNodePoolComponentInstancesResponseBody
	GetComponentInstances() []*ListNodePoolComponentInstancesResponseBodyComponentInstances
	SetMaxResults(v int32) *ListNodePoolComponentInstancesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodePoolComponentInstancesResponseBody
	GetNextToken() *string
}

type ListNodePoolComponentInstancesResponseBody struct {
	ComponentInstances []*ListNodePoolComponentInstancesResponseBodyComponentInstances `json:"component_instances,omitempty" xml:"component_instances,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"max_results,omitempty" xml:"max_results,omitempty"`
	// example:
	//
	// 5c0a1c0f91c14c6****
	NextToken *string `json:"next_token,omitempty" xml:"next_token,omitempty"`
}

func (s ListNodePoolComponentInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodePoolComponentInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodePoolComponentInstancesResponseBody) GetComponentInstances() []*ListNodePoolComponentInstancesResponseBodyComponentInstances {
	return s.ComponentInstances
}

func (s *ListNodePoolComponentInstancesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodePoolComponentInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodePoolComponentInstancesResponseBody) SetComponentInstances(v []*ListNodePoolComponentInstancesResponseBodyComponentInstances) *ListNodePoolComponentInstancesResponseBody {
	s.ComponentInstances = v
	return s
}

func (s *ListNodePoolComponentInstancesResponseBody) SetMaxResults(v int32) *ListNodePoolComponentInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNodePoolComponentInstancesResponseBody) SetNextToken(v string) *ListNodePoolComponentInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNodePoolComponentInstancesResponseBody) Validate() error {
	if s.ComponentInstances != nil {
		for _, item := range s.ComponentInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNodePoolComponentInstancesResponseBodyComponentInstances struct {
	Config *ListNodePoolComponentInstancesResponseBodyComponentInstancesConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// 1
	ConfigRevision *string `json:"config_revision,omitempty" xml:"config_revision,omitempty"`
	// example:
	//
	// kubelet
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// active
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// 1.33.3-aliyun.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListNodePoolComponentInstancesResponseBodyComponentInstances) String() string {
	return dara.Prettify(s)
}

func (s ListNodePoolComponentInstancesResponseBodyComponentInstances) GoString() string {
	return s.String()
}

func (s *ListNodePoolComponentInstancesResponseBodyComponentInstances) GetConfig() *ListNodePoolComponentInstancesResponseBodyComponentInstancesConfig {
	return s.Config
}

func (s *ListNodePoolComponentInstancesResponseBodyComponentInstances) GetConfigRevision() *string {
	return s.ConfigRevision
}

func (s *ListNodePoolComponentInstancesResponseBodyComponentInstances) GetName() *string {
	return s.Name
}

func (s *ListNodePoolComponentInstancesResponseBodyComponentInstances) GetState() *string {
	return s.State
}

func (s *ListNodePoolComponentInstancesResponseBodyComponentInstances) GetVersion() *string {
	return s.Version
}

func (s *ListNodePoolComponentInstancesResponseBodyComponentInstances) SetConfig(v *ListNodePoolComponentInstancesResponseBodyComponentInstancesConfig) *ListNodePoolComponentInstancesResponseBodyComponentInstances {
	s.Config = v
	return s
}

func (s *ListNodePoolComponentInstancesResponseBodyComponentInstances) SetConfigRevision(v string) *ListNodePoolComponentInstancesResponseBodyComponentInstances {
	s.ConfigRevision = &v
	return s
}

func (s *ListNodePoolComponentInstancesResponseBodyComponentInstances) SetName(v string) *ListNodePoolComponentInstancesResponseBodyComponentInstances {
	s.Name = &v
	return s
}

func (s *ListNodePoolComponentInstancesResponseBodyComponentInstances) SetState(v string) *ListNodePoolComponentInstancesResponseBodyComponentInstances {
	s.State = &v
	return s
}

func (s *ListNodePoolComponentInstancesResponseBodyComponentInstances) SetVersion(v string) *ListNodePoolComponentInstancesResponseBodyComponentInstances {
	s.Version = &v
	return s
}

func (s *ListNodePoolComponentInstancesResponseBodyComponentInstances) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNodePoolComponentInstancesResponseBodyComponentInstancesConfig struct {
	// example:
	//
	// {"cpuManagerPolicy":"static"}
	CustomConfig map[string]interface{} `json:"custom_config,omitempty" xml:"custom_config,omitempty"`
}

func (s ListNodePoolComponentInstancesResponseBodyComponentInstancesConfig) String() string {
	return dara.Prettify(s)
}

func (s ListNodePoolComponentInstancesResponseBodyComponentInstancesConfig) GoString() string {
	return s.String()
}

func (s *ListNodePoolComponentInstancesResponseBodyComponentInstancesConfig) GetCustomConfig() map[string]interface{} {
	return s.CustomConfig
}

func (s *ListNodePoolComponentInstancesResponseBodyComponentInstancesConfig) SetCustomConfig(v map[string]interface{}) *ListNodePoolComponentInstancesResponseBodyComponentInstancesConfig {
	s.CustomConfig = v
	return s
}

func (s *ListNodePoolComponentInstancesResponseBodyComponentInstancesConfig) Validate() error {
	return dara.Validate(s)
}
