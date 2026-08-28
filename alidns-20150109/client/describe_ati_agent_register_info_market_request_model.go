// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAtiAgentRegisterInfoMarketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentHost(v string) *DescribeAtiAgentRegisterInfoMarketRequest
	GetAgentHost() *string
	SetAgentVersion(v string) *DescribeAtiAgentRegisterInfoMarketRequest
	GetAgentVersion() *string
	SetClientToken(v string) *DescribeAtiAgentRegisterInfoMarketRequest
	GetClientToken() *string
	SetMaxResults(v int32) *DescribeAtiAgentRegisterInfoMarketRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAtiAgentRegisterInfoMarketRequest
	GetNextToken() *string
	SetTrustLevel(v string) *DescribeAtiAgentRegisterInfoMarketRequest
	GetTrustLevel() *string
}

type DescribeAtiAgentRegisterInfoMarketRequest struct {
	// The endpoint domain name through which the agent provides services externally.
	//
	// example:
	//
	// example.com
	AgentHost *string `json:"AgentHost,omitempty" xml:"AgentHost,omitempty"`
	// The agent version.
	//
	// example:
	//
	// 5.0.1
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cC.....
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The maximum number of entries to return in this request.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query.
	//
	// example:
	//
	// 4698691
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TrustLevel *string `json:"TrustLevel,omitempty" xml:"TrustLevel,omitempty"`
}

func (s DescribeAtiAgentRegisterInfoMarketRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiAgentRegisterInfoMarketRequest) GoString() string {
	return s.String()
}

func (s *DescribeAtiAgentRegisterInfoMarketRequest) GetAgentHost() *string {
	return s.AgentHost
}

func (s *DescribeAtiAgentRegisterInfoMarketRequest) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *DescribeAtiAgentRegisterInfoMarketRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeAtiAgentRegisterInfoMarketRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAtiAgentRegisterInfoMarketRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAtiAgentRegisterInfoMarketRequest) GetTrustLevel() *string {
	return s.TrustLevel
}

func (s *DescribeAtiAgentRegisterInfoMarketRequest) SetAgentHost(v string) *DescribeAtiAgentRegisterInfoMarketRequest {
	s.AgentHost = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketRequest) SetAgentVersion(v string) *DescribeAtiAgentRegisterInfoMarketRequest {
	s.AgentVersion = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketRequest) SetClientToken(v string) *DescribeAtiAgentRegisterInfoMarketRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketRequest) SetMaxResults(v int32) *DescribeAtiAgentRegisterInfoMarketRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketRequest) SetNextToken(v string) *DescribeAtiAgentRegisterInfoMarketRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketRequest) SetTrustLevel(v string) *DescribeAtiAgentRegisterInfoMarketRequest {
	s.TrustLevel = &v
	return s
}

func (s *DescribeAtiAgentRegisterInfoMarketRequest) Validate() error {
	return dara.Validate(s)
}
