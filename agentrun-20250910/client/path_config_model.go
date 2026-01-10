// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPathConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAgentRuntimeEndpointName(v string) *PathConfig
	GetAgentRuntimeEndpointName() *string
	SetMethods(v []*string) *PathConfig
	GetMethods() []*string
	SetPath(v string) *PathConfig
	GetPath() *string
	SetResourceName(v string) *PathConfig
	GetResourceName() *string
	SetResourceType(v string) *PathConfig
	GetResourceType() *string
}

type PathConfig struct {
	// agent runtime 版本（仅当 resourceType 为 runtime 时有效）
	AgentRuntimeEndpointName *string `json:"agentRuntimeEndpointName,omitempty" xml:"agentRuntimeEndpointName,omitempty"`
	// 支持的方法有：HEAD, GET, POST, PUT, DELETE, PATCH, OPTIONS
	//
	// example:
	//
	// [\"GET\"]
	Methods []*string `json:"methods,omitempty" xml:"methods,omitempty" type:"Repeated"`
	// 此条路由规则对应的请求路径。
	//
	// example:
	//
	// /login
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 资源名称
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// 资源类型（和凭证关联资源类型一致）
	//
	// example:
	//
	// runtime
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s PathConfig) String() string {
	return dara.Prettify(s)
}

func (s PathConfig) GoString() string {
	return s.String()
}

func (s *PathConfig) GetAgentRuntimeEndpointName() *string {
	return s.AgentRuntimeEndpointName
}

func (s *PathConfig) GetMethods() []*string {
	return s.Methods
}

func (s *PathConfig) GetPath() *string {
	return s.Path
}

func (s *PathConfig) GetResourceName() *string {
	return s.ResourceName
}

func (s *PathConfig) GetResourceType() *string {
	return s.ResourceType
}

func (s *PathConfig) SetAgentRuntimeEndpointName(v string) *PathConfig {
	s.AgentRuntimeEndpointName = &v
	return s
}

func (s *PathConfig) SetMethods(v []*string) *PathConfig {
	s.Methods = v
	return s
}

func (s *PathConfig) SetPath(v string) *PathConfig {
	s.Path = &v
	return s
}

func (s *PathConfig) SetResourceName(v string) *PathConfig {
	s.ResourceName = &v
	return s
}

func (s *PathConfig) SetResourceType(v string) *PathConfig {
	s.ResourceType = &v
	return s
}

func (s *PathConfig) Validate() error {
	return dara.Validate(s)
}
