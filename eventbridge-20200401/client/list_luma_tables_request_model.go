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
	SetNamespace(v string) *ListLumaTablesRequest
	GetNamespace() *string
}

type ListLumaTablesRequest struct {
	// The name of the Agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The name of the data catalog bound to the Agent. You can call ListLumaCatalogs to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The name of the namespace bound to the Agent. You can call ListLumaNamespaces to obtain this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
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

func (s *ListLumaTablesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListLumaTablesRequest) SetAgentName(v string) *ListLumaTablesRequest {
	s.AgentName = &v
	return s
}

func (s *ListLumaTablesRequest) SetCatalog(v string) *ListLumaTablesRequest {
	s.Catalog = &v
	return s
}

func (s *ListLumaTablesRequest) SetNamespace(v string) *ListLumaTablesRequest {
	s.Namespace = &v
	return s
}

func (s *ListLumaTablesRequest) Validate() error {
	return dara.Validate(s)
}
