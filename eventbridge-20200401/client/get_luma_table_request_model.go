// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *GetLumaTableRequest
	GetAgentName() *string
	SetCatalog(v string) *GetLumaTableRequest
	GetCatalog() *string
	SetName(v string) *GetLumaTableRequest
	GetName() *string
	SetNamespace(v string) *GetLumaTableRequest
	GetNamespace() *string
}

type GetLumaTableRequest struct {
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
	// The name of the event table bound to the Agent. You can call ListLumaTables to obtain the event table name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_table
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the namespace bound to the Agent. You can call ListLumaNamespaces to obtain the namespace name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetLumaTableRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLumaTableRequest) GoString() string {
	return s.String()
}

func (s *GetLumaTableRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *GetLumaTableRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *GetLumaTableRequest) GetName() *string {
	return s.Name
}

func (s *GetLumaTableRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetLumaTableRequest) SetAgentName(v string) *GetLumaTableRequest {
	s.AgentName = &v
	return s
}

func (s *GetLumaTableRequest) SetCatalog(v string) *GetLumaTableRequest {
	s.Catalog = &v
	return s
}

func (s *GetLumaTableRequest) SetName(v string) *GetLumaTableRequest {
	s.Name = &v
	return s
}

func (s *GetLumaTableRequest) SetNamespace(v string) *GetLumaTableRequest {
	s.Namespace = &v
	return s
}

func (s *GetLumaTableRequest) Validate() error {
	return dara.Validate(s)
}
