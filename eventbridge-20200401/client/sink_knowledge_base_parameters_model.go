// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSinkKnowledgeBaseParameters interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *SinkKnowledgeBaseParameters
	GetCatalog() *string
	SetKnowledgeBaseName(v string) *SinkKnowledgeBaseParameters
	GetKnowledgeBaseName() *string
	SetNamespace(v string) *SinkKnowledgeBaseParameters
	GetNamespace() *string
}

type SinkKnowledgeBaseParameters struct {
	// The data catalog to which the target knowledge base belongs. This parameter, together with Namespace and KnowledgeBaseName, uniquely identifies the knowledge base. You can call ListCatalogs to obtain this value.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The name of the target knowledge base, which is unique within the namespace. You can call ListKnowledgeBases to obtain this value.
	//
	// example:
	//
	// my-knowledge-base
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitempty" xml:"KnowledgeBaseName,omitempty"`
	// The namespace to which the target knowledge base belongs. The namespace must belong to the specified data catalog. You can call ListNamespaces to obtain this value.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s SinkKnowledgeBaseParameters) String() string {
	return dara.Prettify(s)
}

func (s SinkKnowledgeBaseParameters) GoString() string {
	return s.String()
}

func (s *SinkKnowledgeBaseParameters) GetCatalog() *string {
	return s.Catalog
}

func (s *SinkKnowledgeBaseParameters) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *SinkKnowledgeBaseParameters) GetNamespace() *string {
	return s.Namespace
}

func (s *SinkKnowledgeBaseParameters) SetCatalog(v string) *SinkKnowledgeBaseParameters {
	s.Catalog = &v
	return s
}

func (s *SinkKnowledgeBaseParameters) SetKnowledgeBaseName(v string) *SinkKnowledgeBaseParameters {
	s.KnowledgeBaseName = &v
	return s
}

func (s *SinkKnowledgeBaseParameters) SetNamespace(v string) *SinkKnowledgeBaseParameters {
	s.Namespace = &v
	return s
}

func (s *SinkKnowledgeBaseParameters) Validate() error {
	return dara.Validate(s)
}
