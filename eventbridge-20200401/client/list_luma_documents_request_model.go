// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaDocumentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *ListLumaDocumentsRequest
	GetAgentName() *string
	SetCatalog(v string) *ListLumaDocumentsRequest
	GetCatalog() *string
	SetFileNamePrefix(v string) *ListLumaDocumentsRequest
	GetFileNamePrefix() *string
	SetKnowledgeBaseName(v string) *ListLumaDocumentsRequest
	GetKnowledgeBaseName() *string
	SetMaxResults(v int32) *ListLumaDocumentsRequest
	GetMaxResults() *int32
	SetNamespace(v string) *ListLumaDocumentsRequest
	GetNamespace() *string
	SetNextToken(v string) *ListLumaDocumentsRequest
	GetNextToken() *string
	SetStatus(v string) *ListLumaDocumentsRequest
	GetStatus() *string
}

type ListLumaDocumentsRequest struct {
	// The name of the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The name of the data catalog bound to the agent. You can call ListLumaCatalogs to obtain the catalog name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The file name prefix used to filter documents.
	//
	// example:
	//
	// manual-
	FileNamePrefix *string `json:"FileNamePrefix,omitempty" xml:"FileNamePrefix,omitempty"`
	// The name of the knowledge base bound to the agent. You can call ListLumaKnowledgeBases to obtain the knowledge base name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-knowledge-base
	KnowledgeBaseName *string `json:"KnowledgeBaseName,omitempty" xml:"KnowledgeBaseName,omitempty"`
	// The maximum number of records to return. Valid values: 1 to 100. If this parameter is not specified, the server uses a default value.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the namespace bound to the agent. You can call ListLumaNamespaces to obtain the namespace name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The pagination token. Do not specify this parameter for the first request. For subsequent requests, use the NextToken value returned in the previous response. This value is an opaque string. Do not parse it.
	//
	// example:
	//
	// ca1eb85f5d99c7d6a97e6****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The processing status used to filter documents. Valid values: Pending, Processing, Ready, and Failed.
	//
	// example:
	//
	// Ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLumaDocumentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLumaDocumentsRequest) GoString() string {
	return s.String()
}

func (s *ListLumaDocumentsRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *ListLumaDocumentsRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *ListLumaDocumentsRequest) GetFileNamePrefix() *string {
	return s.FileNamePrefix
}

func (s *ListLumaDocumentsRequest) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *ListLumaDocumentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListLumaDocumentsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListLumaDocumentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLumaDocumentsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListLumaDocumentsRequest) SetAgentName(v string) *ListLumaDocumentsRequest {
	s.AgentName = &v
	return s
}

func (s *ListLumaDocumentsRequest) SetCatalog(v string) *ListLumaDocumentsRequest {
	s.Catalog = &v
	return s
}

func (s *ListLumaDocumentsRequest) SetFileNamePrefix(v string) *ListLumaDocumentsRequest {
	s.FileNamePrefix = &v
	return s
}

func (s *ListLumaDocumentsRequest) SetKnowledgeBaseName(v string) *ListLumaDocumentsRequest {
	s.KnowledgeBaseName = &v
	return s
}

func (s *ListLumaDocumentsRequest) SetMaxResults(v int32) *ListLumaDocumentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListLumaDocumentsRequest) SetNamespace(v string) *ListLumaDocumentsRequest {
	s.Namespace = &v
	return s
}

func (s *ListLumaDocumentsRequest) SetNextToken(v string) *ListLumaDocumentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListLumaDocumentsRequest) SetStatus(v string) *ListLumaDocumentsRequest {
	s.Status = &v
	return s
}

func (s *ListLumaDocumentsRequest) Validate() error {
	return dara.Validate(s)
}
