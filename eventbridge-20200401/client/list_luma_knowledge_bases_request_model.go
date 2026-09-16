// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaKnowledgeBasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *ListLumaKnowledgeBasesRequest
	GetAgentName() *string
	SetCatalog(v string) *ListLumaKnowledgeBasesRequest
	GetCatalog() *string
	SetNamespace(v string) *ListLumaKnowledgeBasesRequest
	GetNamespace() *string
}

type ListLumaKnowledgeBasesRequest struct {
	// The agent name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The name of the data catalog bound to the agent. You can call ListLumaCatalogs to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The name of the namespace bound to the agent. You can call ListLumaNamespaces to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s ListLumaKnowledgeBasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLumaKnowledgeBasesRequest) GoString() string {
	return s.String()
}

func (s *ListLumaKnowledgeBasesRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *ListLumaKnowledgeBasesRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *ListLumaKnowledgeBasesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListLumaKnowledgeBasesRequest) SetAgentName(v string) *ListLumaKnowledgeBasesRequest {
	s.AgentName = &v
	return s
}

func (s *ListLumaKnowledgeBasesRequest) SetCatalog(v string) *ListLumaKnowledgeBasesRequest {
	s.Catalog = &v
	return s
}

func (s *ListLumaKnowledgeBasesRequest) SetNamespace(v string) *ListLumaKnowledgeBasesRequest {
	s.Namespace = &v
	return s
}

func (s *ListLumaKnowledgeBasesRequest) Validate() error {
	return dara.Validate(s)
}
