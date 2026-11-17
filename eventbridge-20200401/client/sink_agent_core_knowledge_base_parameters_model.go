// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSinkAgentCoreKnowledgeBaseParameters interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *SinkAgentCoreKnowledgeBaseParameters
	GetKnowledgeBaseId() *string
	SetRoleName(v string) *SinkAgentCoreKnowledgeBaseParameters
	GetRoleName() *string
	SetWorkspaceId(v string) *SinkAgentCoreKnowledgeBaseParameters
	GetWorkspaceId() *string
}

type SinkAgentCoreKnowledgeBaseParameters struct {
	// The AgentCore knowledge base ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// kb-123
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The AgentCore knowledge base role name.
	//
	// This parameter is required.
	//
	// example:
	//
	// kbRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The AgentCore workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// kws-123
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SinkAgentCoreKnowledgeBaseParameters) String() string {
	return dara.Prettify(s)
}

func (s SinkAgentCoreKnowledgeBaseParameters) GoString() string {
	return s.String()
}

func (s *SinkAgentCoreKnowledgeBaseParameters) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *SinkAgentCoreKnowledgeBaseParameters) GetRoleName() *string {
	return s.RoleName
}

func (s *SinkAgentCoreKnowledgeBaseParameters) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SinkAgentCoreKnowledgeBaseParameters) SetKnowledgeBaseId(v string) *SinkAgentCoreKnowledgeBaseParameters {
	s.KnowledgeBaseId = &v
	return s
}

func (s *SinkAgentCoreKnowledgeBaseParameters) SetRoleName(v string) *SinkAgentCoreKnowledgeBaseParameters {
	s.RoleName = &v
	return s
}

func (s *SinkAgentCoreKnowledgeBaseParameters) SetWorkspaceId(v string) *SinkAgentCoreKnowledgeBaseParameters {
	s.WorkspaceId = &v
	return s
}

func (s *SinkAgentCoreKnowledgeBaseParameters) Validate() error {
	return dara.Validate(s)
}
