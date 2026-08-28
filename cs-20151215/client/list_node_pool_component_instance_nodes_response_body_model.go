// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodePoolComponentInstanceNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListNodePoolComponentInstanceNodesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodePoolComponentInstanceNodesResponseBody
	GetNextToken() *string
	SetNodeList(v []*ListNodePoolComponentInstanceNodesResponseBodyNodeList) *ListNodePoolComponentInstanceNodesResponseBody
	GetNodeList() []*ListNodePoolComponentInstanceNodesResponseBodyNodeList
	SetTotalCount(v int32) *ListNodePoolComponentInstanceNodesResponseBody
	GetTotalCount() *int32
}

type ListNodePoolComponentInstanceNodesResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"max_results,omitempty" xml:"max_results,omitempty"`
	// example:
	//
	// 5c0a1c0f91c14c6****
	NextToken *string                                                   `json:"next_token,omitempty" xml:"next_token,omitempty"`
	NodeList  []*ListNodePoolComponentInstanceNodesResponseBodyNodeList `json:"node_list,omitempty" xml:"node_list,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s ListNodePoolComponentInstanceNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodePoolComponentInstanceNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodePoolComponentInstanceNodesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodePoolComponentInstanceNodesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodePoolComponentInstanceNodesResponseBody) GetNodeList() []*ListNodePoolComponentInstanceNodesResponseBodyNodeList {
	return s.NodeList
}

func (s *ListNodePoolComponentInstanceNodesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNodePoolComponentInstanceNodesResponseBody) SetMaxResults(v int32) *ListNodePoolComponentInstanceNodesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNodePoolComponentInstanceNodesResponseBody) SetNextToken(v string) *ListNodePoolComponentInstanceNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNodePoolComponentInstanceNodesResponseBody) SetNodeList(v []*ListNodePoolComponentInstanceNodesResponseBodyNodeList) *ListNodePoolComponentInstanceNodesResponseBody {
	s.NodeList = v
	return s
}

func (s *ListNodePoolComponentInstanceNodesResponseBody) SetTotalCount(v int32) *ListNodePoolComponentInstanceNodesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNodePoolComponentInstanceNodesResponseBody) Validate() error {
	if s.NodeList != nil {
		for _, item := range s.NodeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNodePoolComponentInstanceNodesResponseBodyNodeList struct {
	Component *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponent `json:"component,omitempty" xml:"component,omitempty" type:"Struct"`
	// example:
	//
	// i-bp1xxxxx
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// example:
	//
	// cn-hangzhou.10.91.xx.xx
	NodeName *string `json:"node_name,omitempty" xml:"node_name,omitempty"`
}

func (s ListNodePoolComponentInstanceNodesResponseBodyNodeList) String() string {
	return dara.Prettify(s)
}

func (s ListNodePoolComponentInstanceNodesResponseBodyNodeList) GoString() string {
	return s.String()
}

func (s *ListNodePoolComponentInstanceNodesResponseBodyNodeList) GetComponent() *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponent {
	return s.Component
}

func (s *ListNodePoolComponentInstanceNodesResponseBodyNodeList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListNodePoolComponentInstanceNodesResponseBodyNodeList) GetNodeName() *string {
	return s.NodeName
}

func (s *ListNodePoolComponentInstanceNodesResponseBodyNodeList) SetComponent(v *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponent) *ListNodePoolComponentInstanceNodesResponseBodyNodeList {
	s.Component = v
	return s
}

func (s *ListNodePoolComponentInstanceNodesResponseBodyNodeList) SetInstanceId(v string) *ListNodePoolComponentInstanceNodesResponseBodyNodeList {
	s.InstanceId = &v
	return s
}

func (s *ListNodePoolComponentInstanceNodesResponseBodyNodeList) SetNodeName(v string) *ListNodePoolComponentInstanceNodesResponseBodyNodeList {
	s.NodeName = &v
	return s
}

func (s *ListNodePoolComponentInstanceNodesResponseBodyNodeList) Validate() error {
	if s.Component != nil {
		if err := s.Component.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNodePoolComponentInstanceNodesResponseBodyNodeListComponent struct {
	Config *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponentConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
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
	// 1.28.9-aliyun.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListNodePoolComponentInstanceNodesResponseBodyNodeListComponent) String() string {
	return dara.Prettify(s)
}

func (s ListNodePoolComponentInstanceNodesResponseBodyNodeListComponent) GoString() string {
	return s.String()
}

func (s *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponent) GetConfig() *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponentConfig {
	return s.Config
}

func (s *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponent) GetConfigRevision() *string {
	return s.ConfigRevision
}

func (s *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponent) GetName() *string {
	return s.Name
}

func (s *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponent) GetVersion() *string {
	return s.Version
}

func (s *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponent) SetConfig(v *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponentConfig) *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponent {
	s.Config = v
	return s
}

func (s *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponent) SetConfigRevision(v string) *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponent {
	s.ConfigRevision = &v
	return s
}

func (s *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponent) SetName(v string) *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponent {
	s.Name = &v
	return s
}

func (s *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponent) SetVersion(v string) *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponent {
	s.Version = &v
	return s
}

func (s *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponent) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNodePoolComponentInstanceNodesResponseBodyNodeListComponentConfig struct {
	CustomConfig map[string]interface{} `json:"custom_config,omitempty" xml:"custom_config,omitempty"`
}

func (s ListNodePoolComponentInstanceNodesResponseBodyNodeListComponentConfig) String() string {
	return dara.Prettify(s)
}

func (s ListNodePoolComponentInstanceNodesResponseBodyNodeListComponentConfig) GoString() string {
	return s.String()
}

func (s *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponentConfig) GetCustomConfig() map[string]interface{} {
	return s.CustomConfig
}

func (s *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponentConfig) SetCustomConfig(v map[string]interface{}) *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponentConfig {
	s.CustomConfig = v
	return s
}

func (s *ListNodePoolComponentInstanceNodesResponseBodyNodeListComponentConfig) Validate() error {
	return dara.Validate(s)
}
