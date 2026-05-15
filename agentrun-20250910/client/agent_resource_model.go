// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentResource interface {
	dara.Model
	String() string
	GoString() string
	SetFlow(v *Flow) *AgentResource
	GetFlow() *Flow
	SetLatestVersion(v *AgentResourceLatestVersion) *AgentResource
	GetLatestVersion() *AgentResourceLatestVersion
	SetResourceType(v string) *AgentResource
	GetResourceType() *string
	SetRuntime(v *AgentRuntime) *AgentResource
	GetRuntime() *AgentRuntime
}

type AgentResource struct {
	// 当 resourceType 为 Flow 时，此字段包含完整的工作流对象，包括所有配置和状态信息；当 resourceType 为 AgentRuntime 时，此字段为空
	Flow *Flow `json:"flow,omitempty" xml:"flow,omitempty"`
	// 资源的最新发布版本摘要，包含版本号、描述和创建时间。如果资源没有已发布版本，则此字段为空
	LatestVersion *AgentResourceLatestVersion `json:"latestVersion,omitempty" xml:"latestVersion,omitempty" type:"Struct"`
	// 资源类型标识符，用于区分资源是智能体运行时（AgentRuntime）还是工作流（Flow）
	//
	// example:
	//
	// AgentRuntime
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// 当 resourceType 为 AgentRuntime 时，此字段包含完整的智能体运行时对象，包括所有配置和状态信息；当 resourceType 为 Flow 时，此字段为空
	Runtime *AgentRuntime `json:"runtime,omitempty" xml:"runtime,omitempty"`
}

func (s AgentResource) String() string {
	return dara.Prettify(s)
}

func (s AgentResource) GoString() string {
	return s.String()
}

func (s *AgentResource) GetFlow() *Flow {
	return s.Flow
}

func (s *AgentResource) GetLatestVersion() *AgentResourceLatestVersion {
	return s.LatestVersion
}

func (s *AgentResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *AgentResource) GetRuntime() *AgentRuntime {
	return s.Runtime
}

func (s *AgentResource) SetFlow(v *Flow) *AgentResource {
	s.Flow = v
	return s
}

func (s *AgentResource) SetLatestVersion(v *AgentResourceLatestVersion) *AgentResource {
	s.LatestVersion = v
	return s
}

func (s *AgentResource) SetResourceType(v string) *AgentResource {
	s.ResourceType = &v
	return s
}

func (s *AgentResource) SetRuntime(v *AgentRuntime) *AgentResource {
	s.Runtime = v
	return s
}

func (s *AgentResource) Validate() error {
	if s.Flow != nil {
		if err := s.Flow.Validate(); err != nil {
			return err
		}
	}
	if s.LatestVersion != nil {
		if err := s.LatestVersion.Validate(); err != nil {
			return err
		}
	}
	if s.Runtime != nil {
		if err := s.Runtime.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AgentResourceLatestVersion struct {
	// 版本发布时间，采用ISO 8601格式
	//
	// example:
	//
	// 2025-03-15T08:30:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 该版本的描述信息
	//
	// example:
	//
	// 修复了消息处理的并发问题
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 版本发布者（仅 AgentRuntime 类型资源返回）
	//
	// example:
	//
	// admin
	Publisher *string `json:"publisher,omitempty" xml:"publisher,omitempty"`
	// 最新发布的版本号
	//
	// example:
	//
	// 3
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s AgentResourceLatestVersion) String() string {
	return dara.Prettify(s)
}

func (s AgentResourceLatestVersion) GoString() string {
	return s.String()
}

func (s *AgentResourceLatestVersion) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *AgentResourceLatestVersion) GetDescription() *string {
	return s.Description
}

func (s *AgentResourceLatestVersion) GetPublisher() *string {
	return s.Publisher
}

func (s *AgentResourceLatestVersion) GetVersion() *string {
	return s.Version
}

func (s *AgentResourceLatestVersion) SetCreatedAt(v string) *AgentResourceLatestVersion {
	s.CreatedAt = &v
	return s
}

func (s *AgentResourceLatestVersion) SetDescription(v string) *AgentResourceLatestVersion {
	s.Description = &v
	return s
}

func (s *AgentResourceLatestVersion) SetPublisher(v string) *AgentResourceLatestVersion {
	s.Publisher = &v
	return s
}

func (s *AgentResourceLatestVersion) SetVersion(v string) *AgentResourceLatestVersion {
	s.Version = &v
	return s
}

func (s *AgentResourceLatestVersion) Validate() error {
	return dara.Validate(s)
}
