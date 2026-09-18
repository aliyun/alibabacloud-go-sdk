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
	SetLimit(v int32) *ListLumaCatalogsRequest
	GetLimit() *int32
	SetNextToken(v string) *ListLumaCatalogsRequest
	GetNextToken() *string
}

type ListLumaCatalogsRequest struct {
	// The name of the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
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

func (s ListLumaCatalogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLumaCatalogsRequest) GoString() string {
	return s.String()
}

func (s *ListLumaCatalogsRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *ListLumaCatalogsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListLumaCatalogsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLumaCatalogsRequest) SetAgentName(v string) *ListLumaCatalogsRequest {
	s.AgentName = &v
	return s
}

func (s *ListLumaCatalogsRequest) SetLimit(v int32) *ListLumaCatalogsRequest {
	s.Limit = &v
	return s
}

func (s *ListLumaCatalogsRequest) SetNextToken(v string) *ListLumaCatalogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListLumaCatalogsRequest) Validate() error {
	return dara.Validate(s)
}
