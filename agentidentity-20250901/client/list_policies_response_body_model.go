// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPoliciesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListPoliciesResponseBody
	GetNextToken() *string
	SetPolicies(v []*ListPoliciesResponseBodyPolicies) *ListPoliciesResponseBody
	GetPolicies() []*ListPoliciesResponseBodyPolicies
	SetRequestId(v string) *ListPoliciesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPoliciesResponseBody
	GetTotalCount() *int32
}

type ListPoliciesResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Policies  []*ListPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// example:
	//
	// 2A48EB1D-D645-5758-91AF-EDF8E36E257B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPoliciesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPoliciesResponseBody) GetPolicies() []*ListPoliciesResponseBodyPolicies {
	return s.Policies
}

func (s *ListPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPoliciesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPoliciesResponseBody) SetMaxResults(v int32) *ListPoliciesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPoliciesResponseBody) SetNextToken(v string) *ListPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPoliciesResponseBody) SetPolicies(v []*ListPoliciesResponseBodyPolicies) *ListPoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *ListPoliciesResponseBody) SetRequestId(v string) *ListPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPoliciesResponseBody) SetTotalCount(v int32) *ListPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPoliciesResponseBody) Validate() error {
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPoliciesResponseBodyPolicies struct {
	// example:
	//
	// 2026-05-08T06:19:17Z
	CreateTime *string     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Definition *Definition `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// example:
	//
	// example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:policyset/default-policy-set/policy/rate-limit-policy
	PolicyArn *string `json:"PolicyArn,omitempty" xml:"PolicyArn,omitempty"`
	// example:
	//
	// rate-limit-policy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// example:
	//
	// default-policy-set
	PolicySetName *string `json:"PolicySetName,omitempty" xml:"PolicySetName,omitempty"`
	// example:
	//
	// 2026-05-08T06:19:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListPoliciesResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPolicies) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPoliciesResponseBodyPolicies) GetDefinition() *Definition {
	return s.Definition
}

func (s *ListPoliciesResponseBodyPolicies) GetDescription() *string {
	return s.Description
}

func (s *ListPoliciesResponseBodyPolicies) GetPolicyArn() *string {
	return s.PolicyArn
}

func (s *ListPoliciesResponseBodyPolicies) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListPoliciesResponseBodyPolicies) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *ListPoliciesResponseBodyPolicies) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListPoliciesResponseBodyPolicies) SetCreateTime(v string) *ListPoliciesResponseBodyPolicies {
	s.CreateTime = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) SetDefinition(v *Definition) *ListPoliciesResponseBodyPolicies {
	s.Definition = v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) SetDescription(v string) *ListPoliciesResponseBodyPolicies {
	s.Description = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) SetPolicyArn(v string) *ListPoliciesResponseBodyPolicies {
	s.PolicyArn = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) SetPolicyName(v string) *ListPoliciesResponseBodyPolicies {
	s.PolicyName = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) SetPolicySetName(v string) *ListPoliciesResponseBodyPolicies {
	s.PolicySetName = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) SetUpdateTime(v string) *ListPoliciesResponseBodyPolicies {
	s.UpdateTime = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) Validate() error {
	if s.Definition != nil {
		if err := s.Definition.Validate(); err != nil {
			return err
		}
	}
	return nil
}
