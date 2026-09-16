// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *ListLumaTablesRequest
	GetAgentName() *string
	SetCatalog(v string) *ListLumaTablesRequest
	GetCatalog() *string
	SetLimit(v int32) *ListLumaTablesRequest
	GetLimit() *int32
	SetNamespace(v string) *ListLumaTablesRequest
	GetNamespace() *string
	SetNextToken(v string) *ListLumaTablesRequest
	GetNextToken() *string
}

type ListLumaTablesRequest struct {
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
	// The maximum number of entries to return per page. Valid values: 1 to 100. Default value: 100. Each entry requires a back-to-origin metadata query, so this value also limits the number of back-to-origin queries per call.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The name of the namespace bound to the agent. You can call ListLumaNamespaces to obtain the namespace name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The starting token for a paging query. Leave this parameter empty or set it to "0" for the first query. For subsequent pages, use the NextToken value returned in the previous response.
	//
	// example:
	//
	// 0
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListLumaTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLumaTablesRequest) GoString() string {
	return s.String()
}

func (s *ListLumaTablesRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *ListLumaTablesRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *ListLumaTablesRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListLumaTablesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListLumaTablesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLumaTablesRequest) SetAgentName(v string) *ListLumaTablesRequest {
	s.AgentName = &v
	return s
}

func (s *ListLumaTablesRequest) SetCatalog(v string) *ListLumaTablesRequest {
	s.Catalog = &v
	return s
}

func (s *ListLumaTablesRequest) SetLimit(v int32) *ListLumaTablesRequest {
	s.Limit = &v
	return s
}

func (s *ListLumaTablesRequest) SetNamespace(v string) *ListLumaTablesRequest {
	s.Namespace = &v
	return s
}

func (s *ListLumaTablesRequest) SetNextToken(v string) *ListLumaTablesRequest {
	s.NextToken = &v
	return s
}

func (s *ListLumaTablesRequest) Validate() error {
	return dara.Validate(s)
}
