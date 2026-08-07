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
	// if can be null:
	// true
	AgentAccess         *AgentInfoAgentAccess `json:"agentAccess,omitempty" xml:"agentAccess,omitempty" type:"Struct"`
	AgentId             *string               `json:"agentId,omitempty" xml:"agentId,omitempty"`
	AgentType           *string               `json:"agentType,omitempty" xml:"agentType,omitempty"`
	AllowedCapabilities []*string             `json:"allowedCapabilities,omitempty" xml:"allowedCapabilities,omitempty" type:"Repeated"`
	CreateTimestamp     *int64                `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	Description         *string               `json:"description,omitempty" xml:"description,omitempty"`
	GatewayId           *string               `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// if can be null:
	// true
	ModelAccess     *AgentInfoModelAccess `json:"modelAccess,omitempty" xml:"modelAccess,omitempty" type:"Struct"`
	Name            *string               `json:"name,omitempty" xml:"name,omitempty"`
	ResourceGroupId *string               `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Status          *string               `json:"status,omitempty" xml:"status,omitempty"`
	UpdateTimestamp *int64                `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
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
	ConsumerIds []*string `json:"consumerIds,omitempty" xml:"consumerIds,omitempty" type:"Repeated"`
	ModelApiId  *string   `json:"modelApiId,omitempty" xml:"modelApiId,omitempty"`
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
