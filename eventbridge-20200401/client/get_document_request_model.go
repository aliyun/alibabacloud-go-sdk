// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *GetDocumentRequest
	GetCatalog() *string
	SetDocumentId(v string) *GetDocumentRequest
	GetDocumentId() *string
	SetKnowledgeBaseName(v string) *GetDocumentRequest
	GetKnowledgeBaseName() *string
	SetNamespace(v string) *GetDocumentRequest
	GetNamespace() *string
}

type GetDocumentRequest struct {
	// The data catalog to which the knowledge base belongs. This parameter, together with Namespace and KnowledgeBaseName, uniquely identifies a knowledge base. You can call ListCatalogs to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
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

func (s GetDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *GetDocumentRequest) GetDocumentId() *string {
	return s.DocumentId
}

func (s *GetDocumentRequest) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *GetDocumentRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetDocumentRequest) SetCatalog(v string) *GetDocumentRequest {
	s.Catalog = &v
	return s
}

func (s *GetDocumentRequest) SetDocumentId(v string) *GetDocumentRequest {
	s.DocumentId = &v
	return s
}

func (s *GetDocumentRequest) SetKnowledgeBaseName(v string) *GetDocumentRequest {
	s.KnowledgeBaseName = &v
	return s
}

func (s *GetDocumentRequest) SetNamespace(v string) *GetDocumentRequest {
	s.Namespace = &v
	return s
}

func (s *GetDocumentRequest) Validate() error {
	return dara.Validate(s)
}
