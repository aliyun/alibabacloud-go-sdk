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
}

type DescribeAtiAgentRegisterInfoMarketRequest struct {
	// example:
	//
	// example.com
	AgentHost *string `json:"AgentHost,omitempty" xml:"AgentHost,omitempty"`
	// example:
	//
	// 5.0.1
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cC.....
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 4698691
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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

func (s *DescribeAtiAgentRegisterInfoMarketRequest) Validate() error {
	return dara.Validate(s)
}
