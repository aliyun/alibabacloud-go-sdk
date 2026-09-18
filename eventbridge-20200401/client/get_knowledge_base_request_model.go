// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *GetKnowledgeBaseRequest
	GetCatalog() *string
	SetKnowledgeBaseName(v string) *GetKnowledgeBaseRequest
	GetKnowledgeBaseName() *string
	SetNamespace(v string) *GetKnowledgeBaseRequest
	GetNamespace() *string
}

type GetKnowledgeBaseRequest struct {
	// The data catalog to which the knowledge base belongs. This parameter, together with Namespace and KnowledgeBaseName, uniquely identifies a knowledge base. You can call ListCatalogs to obtain the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The name of the knowledge base. The name is unique within a namespace and is determined at creation time. It cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-knowledge-base
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitempty" xml:"KnowledgeBaseName,omitempty"`
	// The namespace to which the knowledge base belongs. The namespace must belong to the specified data catalog. This parameter, together with Catalog and KnowledgeBaseName, uniquely identifies a knowledge base. You can call ListNamespaces to obtain the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBaseRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *GetKnowledgeBaseRequest) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *GetKnowledgeBaseRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetKnowledgeBaseRequest) SetCatalog(v string) *GetKnowledgeBaseRequest {
	s.Catalog = &v
	return s
}

func (s *GetKnowledgeBaseRequest) SetKnowledgeBaseName(v string) *GetKnowledgeBaseRequest {
	s.KnowledgeBaseName = &v
	return s
}

func (s *GetKnowledgeBaseRequest) SetNamespace(v string) *GetKnowledgeBaseRequest {
	s.Namespace = &v
	return s
}

func (s *GetKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
