// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrFirewallV2RoutePolicyListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTrFirewallV2RoutePolicyListResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeTrFirewallV2RoutePolicyListResponseBody
	GetTotalCount() *string
	SetTrFirewallRoutePolicies(v []*DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) *DescribeTrFirewallV2RoutePolicyListResponseBody
	GetTrFirewallRoutePolicies() []*DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies
}

type DescribeTrFirewallV2RoutePolicyListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 95EB5F3A-67FE-5780-92BD-5ECBA772AB7E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The routing policies.
	TrFirewallRoutePolicies []*DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies `json:"TrFirewallRoutePolicies,omitempty" xml:"TrFirewallRoutePolicies,omitempty" type:"Repeated"`
}

func (s DescribeTrFirewallV2RoutePolicyListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallV2RoutePolicyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBody) GetTrFirewallRoutePolicies() []*DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies {
	return s.TrFirewallRoutePolicies
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBody) SetRequestId(v string) *DescribeTrFirewallV2RoutePolicyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBody) SetTotalCount(v string) *DescribeTrFirewallV2RoutePolicyListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBody) SetTrFirewallRoutePolicies(v []*DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) *DescribeTrFirewallV2RoutePolicyListResponseBody {
	s.TrFirewallRoutePolicies = v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBody) Validate() error {
	if s.TrFirewallRoutePolicies != nil {
		for _, item := range s.TrFirewallRoutePolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies struct {
	// The secondary traffic redirection instances.
	DestCandidateList []*DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList `json:"DestCandidateList,omitempty" xml:"DestCandidateList,omitempty" type:"Repeated"`
	// The description of the routing policy.
	//
	// example:
	//
	// test
	PolicyDescription *string `json:"PolicyDescription,omitempty" xml:"PolicyDescription,omitempty"`
	// The name of the routing policy.
	//
	// example:
	//
	// TEST_VPC_FW
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The status of the routing policy. Valid values:
	//
	// 	- creating: The policy is being created.
	//
	// 	- deleting: The policy is being deleted.
	//
	// 	- opening: The policy is being enabled.
	//
	// 	- opened: The policy is enabled.
	//
	// 	- closing: The policy is being disabled.
	//
	// 	- closed: The policy is disabled.
	//
	// example:
	//
	// opened
	PolicyStatus *string `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty"`
	// The type of the traffic redirection scenario of the VPC firewall. Valid values:
	//
	// 	- **fullmesh**: interconnected instances
	//
	// 	- **one_to_one**: instance to instance
	//
	// 	- **end_to_end**: instance to instances
	//
	// example:
	//
	// fullmesh
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The primary traffic redirection instances.
	SrcCandidateList []*DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList `json:"SrcCandidateList,omitempty" xml:"SrcCandidateList,omitempty" type:"Repeated"`
	// The ID of the routing policy.
	//
	// example:
	//
	// policy-7b66257c14e141fb****
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
}

func (s DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) GetDestCandidateList() []*DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList {
	return s.DestCandidateList
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) GetPolicyDescription() *string {
	return s.PolicyDescription
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) GetPolicyStatus() *string {
	return s.PolicyStatus
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) GetSrcCandidateList() []*DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList {
	return s.SrcCandidateList
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) GetTrFirewallRoutePolicyId() *string {
	return s.TrFirewallRoutePolicyId
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) SetDestCandidateList(v []*DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies {
	s.DestCandidateList = v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) SetPolicyDescription(v string) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies {
	s.PolicyDescription = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) SetPolicyName(v string) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies {
	s.PolicyName = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) SetPolicyStatus(v string) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies {
	s.PolicyStatus = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) SetPolicyType(v string) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies {
	s.PolicyType = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) SetSrcCandidateList(v []*DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies {
	s.SrcCandidateList = v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) SetTrFirewallRoutePolicyId(v string) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies {
	s.TrFirewallRoutePolicyId = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) Validate() error {
	if s.DestCandidateList != nil {
		for _, item := range s.DestCandidateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SrcCandidateList != nil {
		for _, item := range s.SrcCandidateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList struct {
	// The ID of the secondary traffic redirection instance.
	//
	// example:
	//
	// vpc-2ze9epancaw8t4sha****
	CandidateId *string `json:"CandidateId,omitempty" xml:"CandidateId,omitempty"`
	// The type of the secondary traffic redirection instance.
	//
	// example:
	//
	// VPC
	CandidateType *string `json:"CandidateType,omitempty" xml:"CandidateType,omitempty"`
}

func (s DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList) GetCandidateId() *string {
	return s.CandidateId
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList) GetCandidateType() *string {
	return s.CandidateType
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList) SetCandidateId(v string) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList {
	s.CandidateId = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList) SetCandidateType(v string) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList {
	s.CandidateType = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList) Validate() error {
	return dara.Validate(s)
}

type DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList struct {
	// The ID of the primary traffic redirection instance.
	//
	// example:
	//
	// vpc-2ze9epancaw8t4sha****
	CandidateId *string `json:"CandidateId,omitempty" xml:"CandidateId,omitempty"`
	// The type of the primary traffic redirection instance.
	//
	// example:
	//
	// VPC
	CandidateType *string `json:"CandidateType,omitempty" xml:"CandidateType,omitempty"`
}

func (s DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList) GetCandidateId() *string {
	return s.CandidateId
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList) GetCandidateType() *string {
	return s.CandidateType
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList) SetCandidateId(v string) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList {
	s.CandidateId = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList) SetCandidateType(v string) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList {
	s.CandidateType = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList) Validate() error {
	return dara.Validate(s)
}
