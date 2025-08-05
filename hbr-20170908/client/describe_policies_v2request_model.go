// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePoliciesV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribePoliciesV2Request
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePoliciesV2Request
	GetNextToken() *string
	SetPolicyId(v string) *DescribePoliciesV2Request
	GetPolicyId() *string
}

type DescribePoliciesV2Request struct {
	// The number of results for each query.
	//
	// Valid values: 10 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to obtain the next page of backup policies.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the backup policy.
	//
	// example:
	//
	// po-000************2l6
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DescribePoliciesV2Request) String() string {
	return dara.Prettify(s)
}

func (s DescribePoliciesV2Request) GoString() string {
	return s.String()
}

func (s *DescribePoliciesV2Request) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePoliciesV2Request) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePoliciesV2Request) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribePoliciesV2Request) SetMaxResults(v int32) *DescribePoliciesV2Request {
	s.MaxResults = &v
	return s
}

func (s *DescribePoliciesV2Request) SetNextToken(v string) *DescribePoliciesV2Request {
	s.NextToken = &v
	return s
}

func (s *DescribePoliciesV2Request) SetPolicyId(v string) *DescribePoliciesV2Request {
	s.PolicyId = &v
	return s
}

func (s *DescribePoliciesV2Request) Validate() error {
	return dara.Validate(s)
}
