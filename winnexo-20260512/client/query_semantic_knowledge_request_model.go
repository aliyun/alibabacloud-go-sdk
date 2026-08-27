// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySemanticKnowledgeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *QuerySemanticKnowledgeRequest
	GetAgentName() *string
	SetGraphName(v string) *QuerySemanticKnowledgeRequest
	GetGraphName() *string
	SetQuery(v string) *QuerySemanticKnowledgeRequest
	GetQuery() *string
	SetTenantId(v string) *QuerySemanticKnowledgeRequest
	GetTenantId() *string
}

type QuerySemanticKnowledgeRequest struct {
	// The digital human name. Call listAuthorizedAgents first to retrieve the list of USE permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// SampleDigitalHuman
	AgentName *string `json:"agentName,omitempty" xml:"agentName,omitempty"`
	// The knowledge graph name. Call listGraphs first to retrieve available graphs.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The natural language query question.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass it explicitly with --tenant-id.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21577
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QuerySemanticKnowledgeRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySemanticKnowledgeRequest) GoString() string {
	return s.String()
}

func (s *QuerySemanticKnowledgeRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *QuerySemanticKnowledgeRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *QuerySemanticKnowledgeRequest) GetQuery() *string {
	return s.Query
}

func (s *QuerySemanticKnowledgeRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QuerySemanticKnowledgeRequest) SetAgentName(v string) *QuerySemanticKnowledgeRequest {
	s.AgentName = &v
	return s
}

func (s *QuerySemanticKnowledgeRequest) SetGraphName(v string) *QuerySemanticKnowledgeRequest {
	s.GraphName = &v
	return s
}

func (s *QuerySemanticKnowledgeRequest) SetQuery(v string) *QuerySemanticKnowledgeRequest {
	s.Query = &v
	return s
}

func (s *QuerySemanticKnowledgeRequest) SetTenantId(v string) *QuerySemanticKnowledgeRequest {
	s.TenantId = &v
	return s
}

func (s *QuerySemanticKnowledgeRequest) Validate() error {
	return dara.Validate(s)
}
