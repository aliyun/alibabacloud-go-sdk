// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *CreateDataAgentKnowledgeBaseRequest
	GetDMSUnit() *string
	SetDescription(v string) *CreateDataAgentKnowledgeBaseRequest
	GetDescription() *string
	SetFromKbUuid(v string) *CreateDataAgentKnowledgeBaseRequest
	GetFromKbUuid() *string
	SetName(v string) *CreateDataAgentKnowledgeBaseRequest
	GetName() *string
	SetWorkspaceId(v string) *CreateDataAgentKnowledgeBaseRequest
	GetWorkspaceId() *string
}

type CreateDataAgentKnowledgeBaseRequest struct {
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// example:
	//
	// KnowledgeBaseTest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// kb-HZ-ra99********ss0xp1bku
	FromKbUuid *string `json:"FromKbUuid,omitempty" xml:"FromKbUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// KnowledgeBaseTest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8i5tw7omgaax*********n909jid
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDataAgentKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *CreateDataAgentKnowledgeBaseRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *CreateDataAgentKnowledgeBaseRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataAgentKnowledgeBaseRequest) GetFromKbUuid() *string {
	return s.FromKbUuid
}

func (s *CreateDataAgentKnowledgeBaseRequest) GetName() *string {
	return s.Name
}

func (s *CreateDataAgentKnowledgeBaseRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDataAgentKnowledgeBaseRequest) SetDMSUnit(v string) *CreateDataAgentKnowledgeBaseRequest {
	s.DMSUnit = &v
	return s
}

func (s *CreateDataAgentKnowledgeBaseRequest) SetDescription(v string) *CreateDataAgentKnowledgeBaseRequest {
	s.Description = &v
	return s
}

func (s *CreateDataAgentKnowledgeBaseRequest) SetFromKbUuid(v string) *CreateDataAgentKnowledgeBaseRequest {
	s.FromKbUuid = &v
	return s
}

func (s *CreateDataAgentKnowledgeBaseRequest) SetName(v string) *CreateDataAgentKnowledgeBaseRequest {
	s.Name = &v
	return s
}

func (s *CreateDataAgentKnowledgeBaseRequest) SetWorkspaceId(v string) *CreateDataAgentKnowledgeBaseRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDataAgentKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
