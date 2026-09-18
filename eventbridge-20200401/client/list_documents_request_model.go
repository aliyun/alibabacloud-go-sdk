// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDocumentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *ListDocumentsRequest
	GetCatalog() *string
	SetFileNamePrefix(v string) *ListDocumentsRequest
	GetFileNamePrefix() *string
	SetKnowledgeBaseName(v string) *ListDocumentsRequest
	GetKnowledgeBaseName() *string
	SetMaxResults(v int32) *ListDocumentsRequest
	GetMaxResults() *int32
	SetNamespace(v string) *ListDocumentsRequest
	GetNamespace() *string
	SetNextToken(v string) *ListDocumentsRequest
	GetNextToken() *string
	SetStatus(v string) *ListDocumentsRequest
	GetStatus() *string
}

type ListDocumentsRequest struct {
	// The data catalog to which the knowledge base belongs. This parameter, together with Namespace and KnowledgeBaseName, uniquely identifies a knowledge base. You can call ListCatalogs to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// Optional. Filters documents by file name prefix (prefix match, with the same semantics as the NamePrefix parameter of ListEventStreamings). If this parameter is not specified or is set to an empty string, no filtering is applied. Maximum length: 255 characters.
	//
	// example:
	//
	// AfterSalesPolicy
	FileNamePrefix *string `json:"FileNamePrefix,omitempty" xml:"FileNamePrefix,omitempty"`
	// The name of the knowledge base. The name is unique within a namespace and is determined at creation time. It cannot be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-knowledge-base
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitempty" xml:"KnowledgeBaseName,omitempty"`
	// The maximum number of results to return per page. If this parameter is not specified or is set to 0, the default value of 20 is used. Maximum value: 100.
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
	// Not required for the first query. For subsequent queries, use the NextToken value returned in the previous response. An empty value indicates that no more pages are available.
	//
	// example:
	//
	// ca1eb85f5d99c7d6a97e6****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Optional. Filters documents by exact status match. Valid values:
	//
	// - UPLOADING: uploading
	//
	// - PENDING: pending processing
	//
	// - PROCESSING: processing
	//
	// - COMPLETED: completed
	//
	// - FAILED: failed
	//
	// - DELETING: deleting
	//
	// If this parameter is not specified, no filtering is applied.
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDocumentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDocumentsRequest) GoString() string {
	return s.String()
}

func (s *ListDocumentsRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *ListDocumentsRequest) GetFileNamePrefix() *string {
	return s.FileNamePrefix
}

func (s *ListDocumentsRequest) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *ListDocumentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDocumentsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListDocumentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDocumentsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDocumentsRequest) SetCatalog(v string) *ListDocumentsRequest {
	s.Catalog = &v
	return s
}

func (s *ListDocumentsRequest) SetFileNamePrefix(v string) *ListDocumentsRequest {
	s.FileNamePrefix = &v
	return s
}

func (s *ListDocumentsRequest) SetKnowledgeBaseName(v string) *ListDocumentsRequest {
	s.KnowledgeBaseName = &v
	return s
}

func (s *ListDocumentsRequest) SetMaxResults(v int32) *ListDocumentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDocumentsRequest) SetNamespace(v string) *ListDocumentsRequest {
	s.Namespace = &v
	return s
}

func (s *ListDocumentsRequest) SetNextToken(v string) *ListDocumentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDocumentsRequest) SetStatus(v string) *ListDocumentsRequest {
	s.Status = &v
	return s
}

func (s *ListDocumentsRequest) Validate() error {
	return dara.Validate(s)
}
