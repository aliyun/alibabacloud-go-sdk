// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicySetAttachedGatewaysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayType(v string) *ListPolicySetAttachedGatewaysRequest
	GetGatewayType() *string
	SetMaxResults(v int32) *ListPolicySetAttachedGatewaysRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPolicySetAttachedGatewaysRequest
	GetNextToken() *string
	SetPolicySetName(v string) *ListPolicySetAttachedGatewaysRequest
	GetPolicySetName() *string
}

type ListPolicySetAttachedGatewaysRequest struct {
	// The gateway type. Valid values:
	//
	// - AIGateway: Alibaba Cloud AI Gateway
	//
	// - AgentRunMCP: AgentRun MCP tool
	//
	// Default value: AIGateway.
	//
	// example:
	//
	// AIGateway
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	// The maximum number of results to return per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used for pagination. You do not need to specify this parameter for the first call. For subsequent calls, pass the NextToken value returned from the previous call.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the policy set. The name must be unique within the current Alibaba Cloud account.
	//
	// example:
	//
	// default-policy-set
	PolicySetName *string `json:"PolicySetName,omitempty" xml:"PolicySetName,omitempty"`
}

func (s ListPolicySetAttachedGatewaysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPolicySetAttachedGatewaysRequest) GoString() string {
	return s.String()
}

func (s *ListPolicySetAttachedGatewaysRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListPolicySetAttachedGatewaysRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPolicySetAttachedGatewaysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPolicySetAttachedGatewaysRequest) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *ListPolicySetAttachedGatewaysRequest) SetGatewayType(v string) *ListPolicySetAttachedGatewaysRequest {
	s.GatewayType = &v
	return s
}

func (s *ListPolicySetAttachedGatewaysRequest) SetMaxResults(v int32) *ListPolicySetAttachedGatewaysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPolicySetAttachedGatewaysRequest) SetNextToken(v string) *ListPolicySetAttachedGatewaysRequest {
	s.NextToken = &v
	return s
}

func (s *ListPolicySetAttachedGatewaysRequest) SetPolicySetName(v string) *ListPolicySetAttachedGatewaysRequest {
	s.PolicySetName = &v
	return s
}

func (s *ListPolicySetAttachedGatewaysRequest) Validate() error {
	return dara.Validate(s)
}
