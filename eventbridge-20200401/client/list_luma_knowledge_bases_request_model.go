// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaKnowledgeBasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *ListLumaKnowledgeBasesRequest
	GetAgentName() *string
	SetCatalog(v string) *ListLumaKnowledgeBasesRequest
	GetCatalog() *string
	SetMaxResults(v int32) *ListLumaKnowledgeBasesRequest
	GetMaxResults() *int32
	SetNamespace(v string) *ListLumaKnowledgeBasesRequest
	GetNamespace() *string
	SetNextToken(v string) *ListLumaKnowledgeBasesRequest
	GetNextToken() *string
}

type ListLumaKnowledgeBasesRequest struct {
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
	// The maximum number of entries to return. Valid values: 1 to 100. If you do not specify this parameter, the server uses the default value of 100. Each entry requires a back-to-origin metadata query, so this value also limits the number of back-to-origin requests per call.
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
	// The pagination token. Do not specify this parameter for the first request. For subsequent requests, set this parameter to the NextToken value returned in the previous response. This value is an opaque string. Do not parse it.
	//
	// example:
	//
	// ca1eb85f5d99c7d6a97e6****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListLumaKnowledgeBasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLumaKnowledgeBasesRequest) GoString() string {
	return s.String()
}

func (s *ListLumaKnowledgeBasesRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *ListLumaKnowledgeBasesRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *ListLumaKnowledgeBasesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListLumaKnowledgeBasesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListLumaKnowledgeBasesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLumaKnowledgeBasesRequest) SetAgentName(v string) *ListLumaKnowledgeBasesRequest {
	s.AgentName = &v
	return s
}

func (s *ListLumaKnowledgeBasesRequest) SetCatalog(v string) *ListLumaKnowledgeBasesRequest {
	s.Catalog = &v
	return s
}

func (s *ListLumaKnowledgeBasesRequest) SetMaxResults(v int32) *ListLumaKnowledgeBasesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListLumaKnowledgeBasesRequest) SetNamespace(v string) *ListLumaKnowledgeBasesRequest {
	s.Namespace = &v
	return s
}

func (s *ListLumaKnowledgeBasesRequest) SetNextToken(v string) *ListLumaKnowledgeBasesRequest {
	s.NextToken = &v
	return s
}

func (s *ListLumaKnowledgeBasesRequest) Validate() error {
	return dara.Validate(s)
}
