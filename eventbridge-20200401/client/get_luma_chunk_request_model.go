// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaChunkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *GetLumaChunkRequest
	GetAgentName() *string
	SetCatalog(v string) *GetLumaChunkRequest
	GetCatalog() *string
	SetChunkSeq(v int32) *GetLumaChunkRequest
	GetChunkSeq() *int32
	SetDocumentId(v string) *GetLumaChunkRequest
	GetDocumentId() *string
	SetKnowledgeBaseName(v string) *GetLumaChunkRequest
	GetKnowledgeBaseName() *string
	SetNamespace(v string) *GetLumaChunkRequest
	GetNamespace() *string
}

type GetLumaChunkRequest struct {
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
	// The sequence number of the text chunk within the document, starting from 0. You can call ListLumaChunks to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ChunkSeq *int32 `json:"ChunkSeq,omitempty" xml:"ChunkSeq,omitempty"`
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

func (s GetLumaChunkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLumaChunkRequest) GoString() string {
	return s.String()
}

func (s *GetLumaChunkRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *GetLumaChunkRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *GetLumaChunkRequest) GetChunkSeq() *int32 {
	return s.ChunkSeq
}

func (s *GetLumaChunkRequest) GetDocumentId() *string {
	return s.DocumentId
}

func (s *GetLumaChunkRequest) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *GetLumaChunkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetLumaChunkRequest) SetAgentName(v string) *GetLumaChunkRequest {
	s.AgentName = &v
	return s
}

func (s *GetLumaChunkRequest) SetCatalog(v string) *GetLumaChunkRequest {
	s.Catalog = &v
	return s
}

func (s *GetLumaChunkRequest) SetChunkSeq(v int32) *GetLumaChunkRequest {
	s.ChunkSeq = &v
	return s
}

func (s *GetLumaChunkRequest) SetDocumentId(v string) *GetLumaChunkRequest {
	s.DocumentId = &v
	return s
}

func (s *GetLumaChunkRequest) SetKnowledgeBaseName(v string) *GetLumaChunkRequest {
	s.KnowledgeBaseName = &v
	return s
}

func (s *GetLumaChunkRequest) SetNamespace(v string) *GetLumaChunkRequest {
	s.Namespace = &v
	return s
}

func (s *GetLumaChunkRequest) Validate() error {
	return dara.Validate(s)
}
