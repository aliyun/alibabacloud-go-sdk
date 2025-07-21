// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppFromTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActualParameters(v string) *CreateAppFromTemplateRequest
	GetActualParameters() *string
	SetAgentId(v string) *CreateAppFromTemplateRequest
	GetAgentId() *string
	SetConnectionsContent(v string) *CreateAppFromTemplateRequest
	GetConnectionsContent() *string
	SetDatabasesContent(v string) *CreateAppFromTemplateRequest
	GetDatabasesContent() *string
	SetDescription(v string) *CreateAppFromTemplateRequest
	GetDescription() *string
	SetFrom(v string) *CreateAppFromTemplateRequest
	GetFrom() *string
	SetIcon(v string) *CreateAppFromTemplateRequest
	GetIcon() *string
	SetName(v string) *CreateAppFromTemplateRequest
	GetName() *string
	SetTemplateId(v string) *CreateAppFromTemplateRequest
	GetTemplateId() *string
	SetType(v string) *CreateAppFromTemplateRequest
	GetType() *string
	SetVariablesContent(v string) *CreateAppFromTemplateRequest
	GetVariablesContent() *string
	SetWorkspaceId(v string) *CreateAppFromTemplateRequest
	GetWorkspaceId() *string
}

type CreateAppFromTemplateRequest struct {
	// example:
	//
	// [{"type":"bailianapp","name":"rag","actualParameter":"11"}]
	ActualParameters   *string `json:"ActualParameters,omitempty" xml:"ActualParameters,omitempty"`
	AgentId            *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	ConnectionsContent *string `json:"ConnectionsContent,omitempty" xml:"ConnectionsContent,omitempty"`
	DatabasesContent   *string `json:"DatabasesContent,omitempty" xml:"DatabasesContent,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	From               *string `json:"From,omitempty" xml:"From,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -1
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4c892729-9950-4353-8146-33542b869e01
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// Web
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VariablesContent *string `json:"VariablesContent,omitempty" xml:"VariablesContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1731664463*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateAppFromTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppFromTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateAppFromTemplateRequest) GetActualParameters() *string {
	return s.ActualParameters
}

func (s *CreateAppFromTemplateRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateAppFromTemplateRequest) GetConnectionsContent() *string {
	return s.ConnectionsContent
}

func (s *CreateAppFromTemplateRequest) GetDatabasesContent() *string {
	return s.DatabasesContent
}

func (s *CreateAppFromTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAppFromTemplateRequest) GetFrom() *string {
	return s.From
}

func (s *CreateAppFromTemplateRequest) GetIcon() *string {
	return s.Icon
}

func (s *CreateAppFromTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateAppFromTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateAppFromTemplateRequest) GetType() *string {
	return s.Type
}

func (s *CreateAppFromTemplateRequest) GetVariablesContent() *string {
	return s.VariablesContent
}

func (s *CreateAppFromTemplateRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateAppFromTemplateRequest) SetActualParameters(v string) *CreateAppFromTemplateRequest {
	s.ActualParameters = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetAgentId(v string) *CreateAppFromTemplateRequest {
	s.AgentId = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetConnectionsContent(v string) *CreateAppFromTemplateRequest {
	s.ConnectionsContent = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetDatabasesContent(v string) *CreateAppFromTemplateRequest {
	s.DatabasesContent = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetDescription(v string) *CreateAppFromTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetFrom(v string) *CreateAppFromTemplateRequest {
	s.From = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetIcon(v string) *CreateAppFromTemplateRequest {
	s.Icon = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetName(v string) *CreateAppFromTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetTemplateId(v string) *CreateAppFromTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetType(v string) *CreateAppFromTemplateRequest {
	s.Type = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetVariablesContent(v string) *CreateAppFromTemplateRequest {
	s.VariablesContent = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetWorkspaceId(v string) *CreateAppFromTemplateRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateAppFromTemplateRequest) Validate() error {
	return dara.Validate(s)
}
