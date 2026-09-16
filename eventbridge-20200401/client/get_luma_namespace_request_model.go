// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *GetLumaNamespaceRequest
	GetAgentName() *string
	SetCatalog(v string) *GetLumaNamespaceRequest
	GetCatalog() *string
	SetName(v string) *GetLumaNamespaceRequest
	GetName() *string
}

type GetLumaNamespaceRequest struct {
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
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetLumaNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLumaNamespaceRequest) GoString() string {
	return s.String()
}

func (s *GetLumaNamespaceRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *GetLumaNamespaceRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *GetLumaNamespaceRequest) GetName() *string {
	return s.Name
}

func (s *GetLumaNamespaceRequest) SetAgentName(v string) *GetLumaNamespaceRequest {
	s.AgentName = &v
	return s
}

func (s *GetLumaNamespaceRequest) SetCatalog(v string) *GetLumaNamespaceRequest {
	s.Catalog = &v
	return s
}

func (s *GetLumaNamespaceRequest) SetName(v string) *GetLumaNamespaceRequest {
	s.Name = &v
	return s
}

func (s *GetLumaNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
