// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIMBotMetadata interface {
	dara.Model
	String() string
	GoString() string
	SetAgentRuntimeEndpointId(v string) *IMBotMetadata
	GetAgentRuntimeEndpointId() *string
	SetAgentRuntimeEndpointUrl(v string) *IMBotMetadata
	GetAgentRuntimeEndpointUrl() *string
	SetAgentRuntimeId(v string) *IMBotMetadata
	GetAgentRuntimeId() *string
	SetCustomFunctionMeta(v string) *IMBotMetadata
	GetCustomFunctionMeta() *string
	SetProtocolType(v string) *IMBotMetadata
	GetProtocolType() *string
	SetRole(v string) *IMBotMetadata
	GetRole() *string
}

type IMBotMetadata struct {
	// 可选
	//
	// if can be null:
	// true
	AgentRuntimeEndpointId *string `json:"agentRuntimeEndpointId,omitempty" xml:"agentRuntimeEndpointId,omitempty"`
	// 标准模式下必填：下游 Agent 可调用的 URL；custom 模式可省略
	//
	// if can be null:
	// true
	AgentRuntimeEndpointUrl *string `json:"agentRuntimeEndpointUrl,omitempty" xml:"agentRuntimeEndpointUrl,omitempty"`
	// 绑定的 Agent Runtime UUID
	//
	// if can be null:
	// true
	AgentRuntimeId *string `json:"agentRuntimeId,omitempty" xml:"agentRuntimeId,omitempty"`
	// 自定义函数元信息；可选；与后端 IMBotMetadata 持久化字段一致
	//
	// if can be null:
	// true
	CustomFunctionMeta *string `json:"customFunctionMeta,omitempty" xml:"customFunctionMeta,omitempty"`
	// 标准模式下必填，与 ValidateAgentRuntimeProtocolTypeValue 一致（大小写敏感）；custom 模式可省略
	//
	// if can be null:
	// true
	ProtocolType *string `json:"protocolType,omitempty" xml:"protocolType,omitempty"`
	// 不写入单 Bot 持久化 JSON；用于首次/更新时设置租户级 FC RAM 执行角色 ARN（acs:ram::...）；FC 启用且非 custom 时按服务逻辑校验
	//
	// if can be null:
	// true
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s IMBotMetadata) String() string {
	return dara.Prettify(s)
}

func (s IMBotMetadata) GoString() string {
	return s.String()
}

func (s *IMBotMetadata) GetAgentRuntimeEndpointId() *string {
	return s.AgentRuntimeEndpointId
}

func (s *IMBotMetadata) GetAgentRuntimeEndpointUrl() *string {
	return s.AgentRuntimeEndpointUrl
}

func (s *IMBotMetadata) GetAgentRuntimeId() *string {
	return s.AgentRuntimeId
}

func (s *IMBotMetadata) GetCustomFunctionMeta() *string {
	return s.CustomFunctionMeta
}

func (s *IMBotMetadata) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *IMBotMetadata) GetRole() *string {
	return s.Role
}

func (s *IMBotMetadata) SetAgentRuntimeEndpointId(v string) *IMBotMetadata {
	s.AgentRuntimeEndpointId = &v
	return s
}

func (s *IMBotMetadata) SetAgentRuntimeEndpointUrl(v string) *IMBotMetadata {
	s.AgentRuntimeEndpointUrl = &v
	return s
}

func (s *IMBotMetadata) SetAgentRuntimeId(v string) *IMBotMetadata {
	s.AgentRuntimeId = &v
	return s
}

func (s *IMBotMetadata) SetCustomFunctionMeta(v string) *IMBotMetadata {
	s.CustomFunctionMeta = &v
	return s
}

func (s *IMBotMetadata) SetProtocolType(v string) *IMBotMetadata {
	s.ProtocolType = &v
	return s
}

func (s *IMBotMetadata) SetRole(v string) *IMBotMetadata {
	s.Role = &v
	return s
}

func (s *IMBotMetadata) Validate() error {
	return dara.Validate(s)
}
