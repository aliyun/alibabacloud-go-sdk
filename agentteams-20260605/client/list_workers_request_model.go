// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentType(v string) *ListWorkersRequest
	GetAgentType() *string
	SetCredential(v string) *ListWorkersRequest
	GetCredential() *string
	SetGroup(v *ListWorkersRequestGroup) *ListWorkersRequest
	GetGroup() *ListWorkersRequestGroup
	SetInstanceId(v string) *ListWorkersRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListWorkersRequest
	GetMaxResults() *int32
	SetMcp(v string) *ListWorkersRequest
	GetMcp() *string
	SetModelName(v string) *ListWorkersRequest
	GetModelName() *string
	SetModelProvider(v string) *ListWorkersRequest
	GetModelProvider() *string
	SetNameLike(v string) *ListWorkersRequest
	GetNameLike() *string
	SetNextToken(v string) *ListWorkersRequest
	GetNextToken() *string
	SetTemplate(v *ListWorkersRequestTemplate) *ListWorkersRequest
	GetTemplate() *ListWorkersRequestTemplate
	SetVersionCode(v string) *ListWorkersRequest
	GetVersionCode() *string
}

type ListWorkersRequest struct {
	AgentType  *string                  `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	Credential *string                  `json:"Credential,omitempty" xml:"Credential,omitempty"`
	Group      *ListWorkersRequestGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// This parameter is required.
	InstanceId    *string                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults    *int32                      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Mcp           *string                     `json:"Mcp,omitempty" xml:"Mcp,omitempty"`
	ModelName     *string                     `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelProvider *string                     `json:"ModelProvider,omitempty" xml:"ModelProvider,omitempty"`
	NameLike      *string                     `json:"NameLike,omitempty" xml:"NameLike,omitempty"`
	NextToken     *string                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Template      *ListWorkersRequestTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
	VersionCode   *string                     `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s ListWorkersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkersRequest) GoString() string {
	return s.String()
}

func (s *ListWorkersRequest) GetAgentType() *string {
	return s.AgentType
}

func (s *ListWorkersRequest) GetCredential() *string {
	return s.Credential
}

func (s *ListWorkersRequest) GetGroup() *ListWorkersRequestGroup {
	return s.Group
}

func (s *ListWorkersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListWorkersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkersRequest) GetMcp() *string {
	return s.Mcp
}

func (s *ListWorkersRequest) GetModelName() *string {
	return s.ModelName
}

func (s *ListWorkersRequest) GetModelProvider() *string {
	return s.ModelProvider
}

func (s *ListWorkersRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListWorkersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkersRequest) GetTemplate() *ListWorkersRequestTemplate {
	return s.Template
}

func (s *ListWorkersRequest) GetVersionCode() *string {
	return s.VersionCode
}

func (s *ListWorkersRequest) SetAgentType(v string) *ListWorkersRequest {
	s.AgentType = &v
	return s
}

func (s *ListWorkersRequest) SetCredential(v string) *ListWorkersRequest {
	s.Credential = &v
	return s
}

func (s *ListWorkersRequest) SetGroup(v *ListWorkersRequestGroup) *ListWorkersRequest {
	s.Group = v
	return s
}

func (s *ListWorkersRequest) SetInstanceId(v string) *ListWorkersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListWorkersRequest) SetMaxResults(v int32) *ListWorkersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkersRequest) SetMcp(v string) *ListWorkersRequest {
	s.Mcp = &v
	return s
}

func (s *ListWorkersRequest) SetModelName(v string) *ListWorkersRequest {
	s.ModelName = &v
	return s
}

func (s *ListWorkersRequest) SetModelProvider(v string) *ListWorkersRequest {
	s.ModelProvider = &v
	return s
}

func (s *ListWorkersRequest) SetNameLike(v string) *ListWorkersRequest {
	s.NameLike = &v
	return s
}

func (s *ListWorkersRequest) SetNextToken(v string) *ListWorkersRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkersRequest) SetTemplate(v *ListWorkersRequestTemplate) *ListWorkersRequest {
	s.Template = v
	return s
}

func (s *ListWorkersRequest) SetVersionCode(v string) *ListWorkersRequest {
	s.VersionCode = &v
	return s
}

func (s *ListWorkersRequest) Validate() error {
	if s.Group != nil {
		if err := s.Group.Validate(); err != nil {
			return err
		}
	}
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkersRequestGroup struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListWorkersRequestGroup) String() string {
	return dara.Prettify(s)
}

func (s ListWorkersRequestGroup) GoString() string {
	return s.String()
}

func (s *ListWorkersRequestGroup) GetName() *string {
	return s.Name
}

func (s *ListWorkersRequestGroup) GetRole() *string {
	return s.Role
}

func (s *ListWorkersRequestGroup) GetType() *string {
	return s.Type
}

func (s *ListWorkersRequestGroup) SetName(v string) *ListWorkersRequestGroup {
	s.Name = &v
	return s
}

func (s *ListWorkersRequestGroup) SetRole(v string) *ListWorkersRequestGroup {
	s.Role = &v
	return s
}

func (s *ListWorkersRequestGroup) SetType(v string) *ListWorkersRequestGroup {
	s.Type = &v
	return s
}

func (s *ListWorkersRequestGroup) Validate() error {
	return dara.Validate(s)
}

type ListWorkersRequestTemplate struct {
	Label   *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListWorkersRequestTemplate) String() string {
	return dara.Prettify(s)
}

func (s ListWorkersRequestTemplate) GoString() string {
	return s.String()
}

func (s *ListWorkersRequestTemplate) GetLabel() *string {
	return s.Label
}

func (s *ListWorkersRequestTemplate) GetName() *string {
	return s.Name
}

func (s *ListWorkersRequestTemplate) GetVersion() *string {
	return s.Version
}

func (s *ListWorkersRequestTemplate) SetLabel(v string) *ListWorkersRequestTemplate {
	s.Label = &v
	return s
}

func (s *ListWorkersRequestTemplate) SetName(v string) *ListWorkersRequestTemplate {
	s.Name = &v
	return s
}

func (s *ListWorkersRequestTemplate) SetVersion(v string) *ListWorkersRequestTemplate {
	s.Version = &v
	return s
}

func (s *ListWorkersRequestTemplate) Validate() error {
	return dara.Validate(s)
}
