// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *DeleteKnowledgeBaseRequest
	GetCatalog() *string
	SetKnowledgeBaseName(v string) *DeleteKnowledgeBaseRequest
	GetKnowledgeBaseName() *string
	SetNamespace(v string) *DeleteKnowledgeBaseRequest
	GetNamespace() *string
}

type DeleteKnowledgeBaseRequest struct {
	// The data catalog to which the knowledge base belongs. This parameter is used together with Namespace and KnowledgeBaseName to locate the knowledge base. You can call ListCatalogs to obtain the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The name of the knowledge base. The name must be unique within the namespace and cannot be modified after creation.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-knowledge-base
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitempty" xml:"KnowledgeBaseName,omitempty"`
	// The namespace to which the knowledge base belongs. The namespace must belong to the specified data catalog. This parameter is used together with Catalog and KnowledgeBaseName to locate the knowledge base. You can call ListNamespaces to obtain the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DeleteKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteKnowledgeBaseRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *DeleteKnowledgeBaseRequest) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *DeleteKnowledgeBaseRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteKnowledgeBaseRequest) SetCatalog(v string) *DeleteKnowledgeBaseRequest {
	s.Catalog = &v
	return s
}

func (s *DeleteKnowledgeBaseRequest) SetKnowledgeBaseName(v string) *DeleteKnowledgeBaseRequest {
	s.KnowledgeBaseName = &v
	return s
}

func (s *DeleteKnowledgeBaseRequest) SetNamespace(v string) *DeleteKnowledgeBaseRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
