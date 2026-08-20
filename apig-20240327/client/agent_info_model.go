// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAgentAccess(v *AgentInfoAgentAccess) *AgentInfo
	GetAgentAccess() *AgentInfoAgentAccess
	SetAgentId(v string) *AgentInfo
	GetAgentId() *string
	SetAgentType(v string) *AgentInfo
	GetAgentType() *string
	SetAllowedCapabilities(v []*string) *AgentInfo
	GetAllowedCapabilities() []*string
	SetCreateTimestamp(v int64) *AgentInfo
	GetCreateTimestamp() *int64
	SetDescription(v string) *AgentInfo
	GetDescription() *string
	SetGatewayId(v string) *AgentInfo
	GetGatewayId() *string
	SetModelAccess(v *AgentInfoModelAccess) *AgentInfo
	GetModelAccess() *AgentInfoModelAccess
	SetName(v string) *AgentInfo
	GetName() *string
	SetResourceGroupId(v string) *AgentInfo
	GetResourceGroupId() *string
	SetStatus(v string) *AgentInfo
	GetStatus() *string
	SetUpdateTimestamp(v int64) *AgentInfo
	GetUpdateTimestamp() *int64
}

type AgentInfo struct {
	// The associated resource information for the Agent access capability. Returns null if the Agent access capability is not configured.
	//
	// if can be null:
	// true
	AgentAccess *AgentInfoAgentAccess `json:"agentAccess,omitempty" xml:"agentAccess,omitempty" type:"Struct"`
	// Agent ID。
	//
	// example:
	//
	// agent-1
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The Agent type. DashScope (Bailian) allows only Agent access. Dify allows both Agent access and model access. ClaudeCode allows only model access. Custom allows both Agent access and model access.
	//
	// example:
	//
	// Custom
	AgentType *string `json:"agentType,omitempty" xml:"agentType,omitempty"`
	// The list of capabilities that the current Agent type allows to be configured. This field does not indicate that the capabilities are already configured. To determine whether a capability is configured, check whether agentAccess or modelAccess is null.
	AllowedCapabilities []*string `json:"allowedCapabilities,omitempty" xml:"allowedCapabilities,omitempty" type:"Repeated"`
	// The Agent creation time, in Unix millisecond timestamp.
	//
	// example:
	//
	// 1755129600000
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// The Agent description.
	//
	// example:
	//
	// custom agent
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The gateway ID to which the Agent belongs. When reading the associated API deployment configuration, select the configuration whose gatewayId matches this value.
	//
	// example:
	//
	// gateway-1
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The associated resource information for the model access capability. Returns null if the model access capability is not configured.
	//
	// if can be null:
	// true
	ModelAccess *AgentInfoModelAccess `json:"modelAccess,omitempty" xml:"modelAccess,omitempty" type:"Struct"`
	// The Agent name.
	//
	// example:
	//
	// my-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The resource group ID in which the Agent is saved.
	//
	// example:
	//
	// rg-1
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The Agent status. An Agent that is successfully created and queryable always returns Ready. Internal creation or compensation states are not returned externally.
	//
	// example:
	//
	// Ready
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The Agent last update time, in Unix millisecond timestamp.
	//
	// example:
	//
	// 1755129600000
	UpdateTimestamp *int64 `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s AgentInfo) String() string {
	return dara.Prettify(s)
}

func (s AgentInfo) GoString() string {
	return s.String()
}

func (s *AgentInfo) GetAgentAccess() *AgentInfoAgentAccess {
	return s.AgentAccess
}

func (s *AgentInfo) GetAgentId() *string {
	return s.AgentId
}

func (s *AgentInfo) GetAgentType() *string {
	return s.AgentType
}

func (s *AgentInfo) GetAllowedCapabilities() []*string {
	return s.AllowedCapabilities
}

func (s *AgentInfo) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *AgentInfo) GetDescription() *string {
	return s.Description
}

func (s *AgentInfo) GetGatewayId() *string {
	return s.GatewayId
}

func (s *AgentInfo) GetModelAccess() *AgentInfoModelAccess {
	return s.ModelAccess
}

func (s *AgentInfo) GetName() *string {
	return s.Name
}

func (s *AgentInfo) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AgentInfo) GetStatus() *string {
	return s.Status
}

func (s *AgentInfo) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *AgentInfo) SetAgentAccess(v *AgentInfoAgentAccess) *AgentInfo {
	s.AgentAccess = v
	return s
}

func (s *AgentInfo) SetAgentId(v string) *AgentInfo {
	s.AgentId = &v
	return s
}

func (s *AgentInfo) SetAgentType(v string) *AgentInfo {
	s.AgentType = &v
	return s
}

func (s *AgentInfo) SetAllowedCapabilities(v []*string) *AgentInfo {
	s.AllowedCapabilities = v
	return s
}

func (s *AgentInfo) SetCreateTimestamp(v int64) *AgentInfo {
	s.CreateTimestamp = &v
	return s
}

func (s *AgentInfo) SetDescription(v string) *AgentInfo {
	s.Description = &v
	return s
}

func (s *AgentInfo) SetGatewayId(v string) *AgentInfo {
	s.GatewayId = &v
	return s
}

func (s *AgentInfo) SetModelAccess(v *AgentInfoModelAccess) *AgentInfo {
	s.ModelAccess = v
	return s
}

func (s *AgentInfo) SetName(v string) *AgentInfo {
	s.Name = &v
	return s
}

func (s *AgentInfo) SetResourceGroupId(v string) *AgentInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *AgentInfo) SetStatus(v string) *AgentInfo {
	s.Status = &v
	return s
}

func (s *AgentInfo) SetUpdateTimestamp(v int64) *AgentInfo {
	s.UpdateTimestamp = &v
	return s
}

func (s *AgentInfo) Validate() error {
	if s.AgentAccess != nil {
		if err := s.AgentAccess.Validate(); err != nil {
			return err
		}
	}
	if s.ModelAccess != nil {
		if err := s.ModelAccess.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AgentInfoAgentAccess struct {
	// The HTTP API ID associated with the Agent access capability. The frontend uses this ID to call existing HTTP API, route, consumer authorization, policy, and plugin query interfaces.
	//
	// example:
	//
	// { "httpApiId": "api-abc123" }
	HttpApiId *string `json:"httpApiId,omitempty" xml:"httpApiId,omitempty"`
}

func (s AgentInfoAgentAccess) String() string {
	return dara.Prettify(s)
}

func (s AgentInfoAgentAccess) GoString() string {
	return s.String()
}

func (s *AgentInfoAgentAccess) GetHttpApiId() *string {
	return s.HttpApiId
}

func (s *AgentInfoAgentAccess) SetHttpApiId(v string) *AgentInfoAgentAccess {
	s.HttpApiId = &v
	return s
}

func (s *AgentInfoAgentAccess) Validate() error {
	return dara.Validate(s)
}

type AgentInfoModelAccess struct {
	// The list of consumer identity bindings maintained by the Agent domain. The Model API ID and the consumer IDs in this list together identify the Agent identity and take effect on all routes of the Model API. Consumer details and their Model API authorization details can be obtained through existing Consumer API and consumer authorization query interfaces.
	ConsumerIds []*string `json:"consumerIds,omitempty" xml:"consumerIds,omitempty" type:"Repeated"`
	// The Model API ID associated with the model access capability. Model Access does not distinguish routes. The frontend uses this ID to query the Model API basic information and all routes.
	//
	// example:
	//
	// { "modelApiId": "model-api-1" }
	ModelApiId *string `json:"modelApiId,omitempty" xml:"modelApiId,omitempty"`
}

func (s AgentInfoModelAccess) String() string {
	return dara.Prettify(s)
}

func (s AgentInfoModelAccess) GoString() string {
	return s.String()
}

func (s *AgentInfoModelAccess) GetConsumerIds() []*string {
	return s.ConsumerIds
}

func (s *AgentInfoModelAccess) GetModelApiId() *string {
	return s.ModelApiId
}

func (s *AgentInfoModelAccess) SetConsumerIds(v []*string) *AgentInfoModelAccess {
	s.ConsumerIds = v
	return s
}

func (s *AgentInfoModelAccess) SetModelApiId(v string) *AgentInfoModelAccess {
	s.ModelApiId = &v
	return s
}

func (s *AgentInfoModelAccess) Validate() error {
	return dara.Validate(s)
}
