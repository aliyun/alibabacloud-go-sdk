// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaCatalogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *GetLumaCatalogRequest
	GetAgentName() *string
	SetName(v string) *GetLumaCatalogRequest
	GetName() *string
}

type GetLumaCatalogRequest struct {
	// The name of the Agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The name of the data catalog bound to the Agent. You can call ListLumaCatalogs to obtain the name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetLumaCatalogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLumaCatalogRequest) GoString() string {
	return s.String()
}

func (s *GetLumaCatalogRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *GetLumaCatalogRequest) GetName() *string {
	return s.Name
}

func (s *GetLumaCatalogRequest) SetAgentName(v string) *GetLumaCatalogRequest {
	s.AgentName = &v
	return s
}

func (s *GetLumaCatalogRequest) SetName(v string) *GetLumaCatalogRequest {
	s.Name = &v
	return s
}

func (s *GetLumaCatalogRequest) Validate() error {
	return dara.Validate(s)
}
