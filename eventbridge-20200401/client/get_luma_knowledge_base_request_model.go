// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *GetLumaKnowledgeBaseRequest
	GetAgentName() *string
	SetCatalog(v string) *GetLumaKnowledgeBaseRequest
	GetCatalog() *string
	SetKnowledgeBaseName(v string) *GetLumaKnowledgeBaseRequest
	GetKnowledgeBaseName() *string
	SetNamespace(v string) *GetLumaKnowledgeBaseRequest
	GetNamespace() *string
}

type GetLumaKnowledgeBaseRequest struct {
	// The name of the Agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The name of the data catalog bound to the Agent. You can call ListLumaCatalogs to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The name of the knowledge base bound to the Agent. You can call ListLumaKnowledgeBases to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-knowledge-base
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitempty" xml:"KnowledgeBaseName,omitempty"`
	// The name of the namespace bound to the Agent. You can call ListLumaNamespaces to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetLumaKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLumaKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *GetLumaKnowledgeBaseRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *GetLumaKnowledgeBaseRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *GetLumaKnowledgeBaseRequest) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *GetLumaKnowledgeBaseRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetLumaKnowledgeBaseRequest) SetAgentName(v string) *GetLumaKnowledgeBaseRequest {
	s.AgentName = &v
	return s
}

func (s *GetLumaKnowledgeBaseRequest) SetCatalog(v string) *GetLumaKnowledgeBaseRequest {
	s.Catalog = &v
	return s
}

func (s *GetLumaKnowledgeBaseRequest) SetKnowledgeBaseName(v string) *GetLumaKnowledgeBaseRequest {
	s.KnowledgeBaseName = &v
	return s
}

func (s *GetLumaKnowledgeBaseRequest) SetNamespace(v string) *GetLumaKnowledgeBaseRequest {
	s.Namespace = &v
	return s
}

func (s *GetLumaKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
