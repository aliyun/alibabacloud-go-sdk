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
	SetCompatibleProtocol(v string) *PathConfig
	GetCompatibleProtocol() *string
	SetFlowEndpointName(v string) *PathConfig
	GetFlowEndpointName() *string
	SetMethods(v []*string) *PathConfig
	GetMethods() []*string
	SetPath(v string) *PathConfig
	GetPath() *string
	SetRemoveBasePathOnForward(v bool) *PathConfig
	GetRemoveBasePathOnForward() *bool
	SetResourceName(v string) *PathConfig
	GetResourceName() *string
	SetResourceType(v string) *PathConfig
	GetResourceType() *string
}

type PathConfig struct {
	// The agent runtime version. This parameter takes effect only when `resourceType` is `runtime`.
	AgentRuntimeEndpointName *string `json:"agentRuntimeEndpointName,omitempty" xml:"agentRuntimeEndpointName,omitempty"`
	// The compatible protocol, used to convert the backend response format. This parameter is required only when `resourceType` is `flow`. Valid values: `native` indicates an FnF native call; `openai`, `dify-workflow`, and `dify-chatflow` map to their corresponding compatible APIs.
	//
	// example:
	//
	// native
	CompatibleProtocol *string `json:"compatibleProtocol,omitempty" xml:"compatibleProtocol,omitempty"`
	// The Flow version/alias. This parameter takes effect only when `resourceType` is `flow`. Default value: `Default`.
	//
	// example:
	//
	// Default
	FlowEndpointName *string `json:"flowEndpointName,omitempty" xml:"flowEndpointName,omitempty"`
	// Supported methods: HEAD, GET, POST, PUT, DELETE, PATCH, and OPTIONS.
	//
	// example:
	//
	// [\\"GET\\"]
	Methods []*string `json:"methods" xml:"methods" type:"Repeated"`
	// The path for this routing rule.
	//
	// example:
	//
	// /login
	Path                    *string `json:"path,omitempty" xml:"path,omitempty"`
	RemoveBasePathOnForward *bool   `json:"removeBasePathOnForward,omitempty" xml:"removeBasePathOnForward,omitempty"`
	// The resource name.
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// The resource type. This type must match the one associated with the credential.
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

func (s *PathConfig) GetCompatibleProtocol() *string {
	return s.CompatibleProtocol
}

func (s *PathConfig) GetFlowEndpointName() *string {
	return s.FlowEndpointName
}

func (s *PathConfig) GetMethods() []*string {
	return s.Methods
}

func (s *PathConfig) GetPath() *string {
	return s.Path
}

func (s *PathConfig) GetRemoveBasePathOnForward() *bool {
	return s.RemoveBasePathOnForward
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

func (s *PathConfig) SetCompatibleProtocol(v string) *PathConfig {
	s.CompatibleProtocol = &v
	return s
}

func (s *PathConfig) SetFlowEndpointName(v string) *PathConfig {
	s.FlowEndpointName = &v
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

func (s *PathConfig) SetRemoveBasePathOnForward(v bool) *PathConfig {
	s.RemoveBasePathOnForward = &v
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
