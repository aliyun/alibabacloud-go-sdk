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
}

type ListLumaNamespacesRequest struct {
	// The name of the Agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The name of the data catalog bound to the Agent. You can call ListLumaCatalogs to obtain the catalog name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
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

func (s *ListLumaNamespacesRequest) SetAgentName(v string) *ListLumaNamespacesRequest {
	s.AgentName = &v
	return s
}

func (s *ListLumaNamespacesRequest) SetCatalog(v string) *ListLumaNamespacesRequest {
	s.Catalog = &v
	return s
}

func (s *ListLumaNamespacesRequest) Validate() error {
	return dara.Validate(s)
}
