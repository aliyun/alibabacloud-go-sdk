// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaCatalogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *ListLumaCatalogsRequest
	GetAgentName() *string
}

type ListLumaCatalogsRequest struct {
	// The name of the Agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
}

func (s ListLumaCatalogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLumaCatalogsRequest) GoString() string {
	return s.String()
}

func (s *ListLumaCatalogsRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *ListLumaCatalogsRequest) SetAgentName(v string) *ListLumaCatalogsRequest {
	s.AgentName = &v
	return s
}

func (s *ListLumaCatalogsRequest) Validate() error {
	return dara.Validate(s)
}
