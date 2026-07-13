// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProtectionPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListProtectionPoliciesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProtectionPoliciesRequest
	GetNextToken() *string
	SetProtectionPolicyId(v string) *ListProtectionPoliciesRequest
	GetProtectionPolicyId() *string
	SetProtectionPolicyRegionId(v string) *ListProtectionPoliciesRequest
	GetProtectionPolicyRegionId() *string
}

type ListProtectionPoliciesRequest struct {
	// The maximum number of results to return.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The paging token.
	//
	// example:
	//
	// cae**********699
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The protection policy ID.
	//
	// example:
	//
	// p-123***7890
	ProtectionPolicyId *string `json:"ProtectionPolicyId,omitempty" xml:"ProtectionPolicyId,omitempty"`
	// The region ID of the protection policy.
	//
	// example:
	//
	// cn-hangzhou
	ProtectionPolicyRegionId *string `json:"ProtectionPolicyRegionId,omitempty" xml:"ProtectionPolicyRegionId,omitempty"`
}

func (s ListProtectionPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProtectionPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListProtectionPoliciesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProtectionPoliciesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProtectionPoliciesRequest) GetProtectionPolicyId() *string {
	return s.ProtectionPolicyId
}

func (s *ListProtectionPoliciesRequest) GetProtectionPolicyRegionId() *string {
	return s.ProtectionPolicyRegionId
}

func (s *ListProtectionPoliciesRequest) SetMaxResults(v int32) *ListProtectionPoliciesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProtectionPoliciesRequest) SetNextToken(v string) *ListProtectionPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListProtectionPoliciesRequest) SetProtectionPolicyId(v string) *ListProtectionPoliciesRequest {
	s.ProtectionPolicyId = &v
	return s
}

func (s *ListProtectionPoliciesRequest) SetProtectionPolicyRegionId(v string) *ListProtectionPoliciesRequest {
	s.ProtectionPolicyRegionId = &v
	return s
}

func (s *ListProtectionPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
