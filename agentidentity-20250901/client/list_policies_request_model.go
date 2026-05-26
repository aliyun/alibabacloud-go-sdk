// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPoliciesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPoliciesRequest
	GetNextToken() *string
	SetPolicySetName(v string) *ListPoliciesRequest
	GetPolicySetName() *string
}

type ListPoliciesRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// default-policy-set
	PolicySetName *string `json:"PolicySetName,omitempty" xml:"PolicySetName,omitempty"`
}

func (s ListPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPoliciesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPoliciesRequest) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *ListPoliciesRequest) SetMaxResults(v int32) *ListPoliciesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPoliciesRequest) SetNextToken(v string) *ListPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPoliciesRequest) SetPolicySetName(v string) *ListPoliciesRequest {
	s.PolicySetName = &v
	return s
}

func (s *ListPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
