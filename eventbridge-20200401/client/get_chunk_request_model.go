// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChunkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *GetChunkRequest
	GetCatalog() *string
	SetChunkSeq(v int32) *GetChunkRequest
	GetChunkSeq() *int32
	SetDocumentId(v string) *GetChunkRequest
	GetDocumentId() *string
	SetKnowledgeBaseName(v string) *GetChunkRequest
	GetKnowledgeBaseName() *string
	SetNamespace(v string) *GetChunkRequest
	GetNamespace() *string
}

type GetChunkRequest struct {
	// The data catalog to which the knowledge base belongs. This parameter, together with Namespace and KnowledgeBaseName, uniquely identifies a knowledge base. You can call ListCatalogs to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The sequence number of the chunk within the document, starting from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	ChunkSeq *int32 `json:"ChunkSeq,omitempty" xml:"ChunkSeq,omitempty"`
	// The document ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// doc-bp1xxxxxxxxxxxx
	DocumentId *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// The name of the knowledge base. The name is unique within the namespace and is determined at creation time. It cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-knowledge-base
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitempty" xml:"KnowledgeBaseName,omitempty"`
	// The namespace to which the knowledge base belongs. The namespace must belong to the specified data catalog. This parameter, together with Catalog and KnowledgeBaseName, uniquely identifies a knowledge base. You can call ListNamespaces to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetChunkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChunkRequest) GoString() string {
	return s.String()
}

func (s *GetChunkRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *GetChunkRequest) GetChunkSeq() *int32 {
	return s.ChunkSeq
}

func (s *GetChunkRequest) GetDocumentId() *string {
	return s.DocumentId
}

func (s *GetChunkRequest) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *GetChunkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetChunkRequest) SetCatalog(v string) *GetChunkRequest {
	s.Catalog = &v
	return s
}

func (s *GetChunkRequest) SetChunkSeq(v int32) *GetChunkRequest {
	s.ChunkSeq = &v
	return s
}

func (s *GetChunkRequest) SetDocumentId(v string) *GetChunkRequest {
	s.DocumentId = &v
	return s
}

func (s *GetChunkRequest) SetKnowledgeBaseName(v string) *GetChunkRequest {
	s.KnowledgeBaseName = &v
	return s
}

func (s *GetChunkRequest) SetNamespace(v string) *GetChunkRequest {
	s.Namespace = &v
	return s
}

func (s *GetChunkRequest) Validate() error {
	return dara.Validate(s)
}
