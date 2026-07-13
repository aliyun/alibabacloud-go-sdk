// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAtiAgentRegisterInfoMarketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SearchAtiAgentRegisterInfoMarketRequest
	GetClientToken() *string
	SetKeyword(v string) *SearchAtiAgentRegisterInfoMarketRequest
	GetKeyword() *string
	SetMaxResults(v int32) *SearchAtiAgentRegisterInfoMarketRequest
	GetMaxResults() *int32
	SetNextToken(v string) *SearchAtiAgentRegisterInfoMarketRequest
	GetNextToken() *string
	SetPageNumber(v int32) *SearchAtiAgentRegisterInfoMarketRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchAtiAgentRegisterInfoMarketRequest
	GetPageSize() *int32
	SetProtocol(v string) *SearchAtiAgentRegisterInfoMarketRequest
	GetProtocol() *string
	SetStatus(v string) *SearchAtiAgentRegisterInfoMarketRequest
	GetStatus() *string
	SetTrustLevel(v string) *SearchAtiAgentRegisterInfoMarketRequest
	GetTrustLevel() *string
}

type SearchAtiAgentRegisterInfoMarketRequest struct {
	// Ensures the idempotency of the request. Generate a unique parameter value from your client to ensure uniqueness across different requests. ClientToken supports only ASCII characters and cannot exceed 64 characters in length.
	//
	// If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may differ for each API request.
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cC.....
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The keyword for searching. Matches against agent name, domain name, and description.
	//
	// example:
	//
	// example.com
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of entries per batch query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query.
	//
	// example:
	//
	// 4698691
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The current page number. Minimum value: **1**. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for the paged query. Settings: maximum value: 100. Default value: 20. This parameter controls paging behavior.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The communication protocol that the agent endpoint follows, which determines how callers interact with the agent.
	//
	// Valid values:
	//
	// - MCP: Model Context Protocol, an agent tool calling protocol developed by Anthropic.
	//
	// - A2A: Agent-to-Agent Protocol, a cross-agent communication protocol developed by Google.
	//
	// - OpenAPI: Standard RESTful API specification (Swagger/OpenAPI).
	//
	// Other agents or clients use this protocol identifier to determine how to communicate with the agent. For example, MCP uses the MCP SDK, A2A uses the A2A SDK, and OpenAPI uses standard HTTP requests.
	//
	// example:
	//
	// mcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// Queries agents based on the agent status.
	//
	// example:
	//
	// 活跃
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Queries agents based on the trust level.
	//
	// example:
	//
	// 基础认证
	TrustLevel *string `json:"TrustLevel,omitempty" xml:"TrustLevel,omitempty"`
}

func (s SearchAtiAgentRegisterInfoMarketRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchAtiAgentRegisterInfoMarketRequest) GoString() string {
	return s.String()
}

func (s *SearchAtiAgentRegisterInfoMarketRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SearchAtiAgentRegisterInfoMarketRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *SearchAtiAgentRegisterInfoMarketRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchAtiAgentRegisterInfoMarketRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchAtiAgentRegisterInfoMarketRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchAtiAgentRegisterInfoMarketRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchAtiAgentRegisterInfoMarketRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *SearchAtiAgentRegisterInfoMarketRequest) GetStatus() *string {
	return s.Status
}

func (s *SearchAtiAgentRegisterInfoMarketRequest) GetTrustLevel() *string {
	return s.TrustLevel
}

func (s *SearchAtiAgentRegisterInfoMarketRequest) SetClientToken(v string) *SearchAtiAgentRegisterInfoMarketRequest {
	s.ClientToken = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketRequest) SetKeyword(v string) *SearchAtiAgentRegisterInfoMarketRequest {
	s.Keyword = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketRequest) SetMaxResults(v int32) *SearchAtiAgentRegisterInfoMarketRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketRequest) SetNextToken(v string) *SearchAtiAgentRegisterInfoMarketRequest {
	s.NextToken = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketRequest) SetPageNumber(v int32) *SearchAtiAgentRegisterInfoMarketRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketRequest) SetPageSize(v int32) *SearchAtiAgentRegisterInfoMarketRequest {
	s.PageSize = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketRequest) SetProtocol(v string) *SearchAtiAgentRegisterInfoMarketRequest {
	s.Protocol = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketRequest) SetStatus(v string) *SearchAtiAgentRegisterInfoMarketRequest {
	s.Status = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketRequest) SetTrustLevel(v string) *SearchAtiAgentRegisterInfoMarketRequest {
	s.TrustLevel = &v
	return s
}

func (s *SearchAtiAgentRegisterInfoMarketRequest) Validate() error {
	return dara.Validate(s)
}
