// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *GetLumaDocumentRequest
	GetAgentName() *string
	SetCatalog(v string) *GetLumaDocumentRequest
	GetCatalog() *string
	SetDocumentId(v string) *GetLumaDocumentRequest
	GetDocumentId() *string
	SetKnowledgeBaseName(v string) *GetLumaDocumentRequest
	GetKnowledgeBaseName() *string
	SetNamespace(v string) *GetLumaDocumentRequest
	GetNamespace() *string
}

type GetLumaDocumentRequest struct {
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
	// The unique identifier of the document. You can call ListLumaDocuments to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// doc-a1b2c3d4
	DocumentId *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
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

func (s GetLumaDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLumaDocumentRequest) GoString() string {
	return s.String()
}

func (s *GetLumaDocumentRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *GetLumaDocumentRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *GetLumaDocumentRequest) GetDocumentId() *string {
	return s.DocumentId
}

func (s *GetLumaDocumentRequest) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *GetLumaDocumentRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetLumaDocumentRequest) SetAgentName(v string) *GetLumaDocumentRequest {
	s.AgentName = &v
	return s
}

func (s *GetLumaDocumentRequest) SetCatalog(v string) *GetLumaDocumentRequest {
	s.Catalog = &v
	return s
}

func (s *GetLumaDocumentRequest) SetDocumentId(v string) *GetLumaDocumentRequest {
	s.DocumentId = &v
	return s
}

func (s *GetLumaDocumentRequest) SetKnowledgeBaseName(v string) *GetLumaDocumentRequest {
	s.KnowledgeBaseName = &v
	return s
}

func (s *GetLumaDocumentRequest) SetNamespace(v string) *GetLumaDocumentRequest {
	s.Namespace = &v
	return s
}

func (s *GetLumaDocumentRequest) Validate() error {
	return dara.Validate(s)
}
