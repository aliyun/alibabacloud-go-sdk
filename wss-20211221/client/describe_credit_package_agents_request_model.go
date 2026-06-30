// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreditPackageAgentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIds(v []*string) *DescribeCreditPackageAgentsRequest
	GetAgentIds() []*string
	SetAgentType(v string) *DescribeCreditPackageAgentsRequest
	GetAgentType() *string
	SetBizType(v string) *DescribeCreditPackageAgentsRequest
	GetBizType() *string
	SetMaxResults(v int32) *DescribeCreditPackageAgentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeCreditPackageAgentsRequest
	GetNextToken() *string
}

type DescribeCreditPackageAgentsRequest struct {
	// An array of agent IDs to query. Example: `["agent-1","agent-2"]`.
	AgentIds []*string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
	// The agent type. Valid values: `CREDIT_PACKAGE`, `JVS_CLAW`, `OPEN_CLAW`, and `JVS_COPILOT`.
	//
	// example:
	//
	// JVS_COPILOT
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The business type.
	//
	// example:
	//
	// BUSINESS
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The maximum number of results to return per page.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to retrieve the next page of results. Obtain this value from the `NextToken` parameter of the previous response. For the first request, set this parameter to an empty string.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeCreditPackageAgentsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditPackageAgentsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCreditPackageAgentsRequest) GetAgentIds() []*string {
	return s.AgentIds
}

func (s *DescribeCreditPackageAgentsRequest) GetAgentType() *string {
	return s.AgentType
}

func (s *DescribeCreditPackageAgentsRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribeCreditPackageAgentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCreditPackageAgentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCreditPackageAgentsRequest) SetAgentIds(v []*string) *DescribeCreditPackageAgentsRequest {
	s.AgentIds = v
	return s
}

func (s *DescribeCreditPackageAgentsRequest) SetAgentType(v string) *DescribeCreditPackageAgentsRequest {
	s.AgentType = &v
	return s
}

func (s *DescribeCreditPackageAgentsRequest) SetBizType(v string) *DescribeCreditPackageAgentsRequest {
	s.BizType = &v
	return s
}

func (s *DescribeCreditPackageAgentsRequest) SetMaxResults(v int32) *DescribeCreditPackageAgentsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeCreditPackageAgentsRequest) SetNextToken(v string) *DescribeCreditPackageAgentsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCreditPackageAgentsRequest) Validate() error {
	return dara.Validate(s)
}
