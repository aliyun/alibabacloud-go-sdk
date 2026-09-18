// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *ListKnowledgeBasesRequest
	GetCatalog() *string
	SetMaxResults(v int32) *ListKnowledgeBasesRequest
	GetMaxResults() *int32
	SetNamespace(v string) *ListKnowledgeBasesRequest
	GetNamespace() *string
	SetNextToken(v string) *ListKnowledgeBasesRequest
	GetNextToken() *string
}

type ListKnowledgeBasesRequest struct {
	// The data catalog to which the knowledge base belongs. This parameter, together with Namespace and KnowledgeBaseName, uniquely identifies a knowledge base. You can call ListCatalogs to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The maximum number of entries to return per page. If this parameter is not specified or is set to 0, the default value 20 is used. Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The namespace to which the knowledge base belongs. The namespace must belong to the specified data catalog. This parameter, together with Catalog and KnowledgeBaseName, uniquely identifies a knowledge base. You can call ListNamespaces to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The pagination token. You do not need to specify this parameter for the first request. For subsequent requests, use the NextToken value returned in the previous response. An empty value indicates that no more pages are available.
	//
	// example:
	//
	// ca1eb85f5d99c7d6a97e6****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListKnowledgeBasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBasesRequest) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBasesRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *ListKnowledgeBasesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeBasesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListKnowledgeBasesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeBasesRequest) SetCatalog(v string) *ListKnowledgeBasesRequest {
	s.Catalog = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetMaxResults(v int32) *ListKnowledgeBasesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetNamespace(v string) *ListKnowledgeBasesRequest {
	s.Namespace = &v
	return s
}

func (s *ListKnowledgeBasesRequest) SetNextToken(v string) *ListKnowledgeBasesRequest {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeBasesRequest) Validate() error {
	return dara.Validate(s)
}
