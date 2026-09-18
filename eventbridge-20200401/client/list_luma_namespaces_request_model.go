// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaNamespacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *ListLumaNamespacesRequest
	GetAgentName() *string
	SetCatalog(v string) *ListLumaNamespacesRequest
	GetCatalog() *string
	SetLimit(v int32) *ListLumaNamespacesRequest
	GetLimit() *int32
	SetNextToken(v string) *ListLumaNamespacesRequest
	GetNextToken() *string
}

type ListLumaNamespacesRequest struct {
	// The name of the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The name of the data catalog bound to the agent. You can call the ListLumaCatalogs operation to obtain the catalog name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The maximum number of entries to return per page. Valid values: 1 to 100. Default value: 100. Each entry requires a back-to-origin metadata query, so this value also limits the number of back-to-origin requests per call.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The token for the paging query. Leave this parameter empty or set it to "0" for the first query. For subsequent pages, use the NextToken value returned in the previous response.
	//
	// example:
	//
	// 0
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListLumaNamespacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLumaNamespacesRequest) GoString() string {
	return s.String()
}

func (s *ListLumaNamespacesRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *ListLumaNamespacesRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *ListLumaNamespacesRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListLumaNamespacesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLumaNamespacesRequest) SetAgentName(v string) *ListLumaNamespacesRequest {
	s.AgentName = &v
	return s
}

func (s *ListLumaNamespacesRequest) SetCatalog(v string) *ListLumaNamespacesRequest {
	s.Catalog = &v
	return s
}

func (s *ListLumaNamespacesRequest) SetLimit(v int32) *ListLumaNamespacesRequest {
	s.Limit = &v
	return s
}

func (s *ListLumaNamespacesRequest) SetNextToken(v string) *ListLumaNamespacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListLumaNamespacesRequest) Validate() error {
	return dara.Validate(s)
}
