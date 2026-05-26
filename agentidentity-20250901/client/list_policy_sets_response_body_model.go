// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicySetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPolicySetsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListPolicySetsResponseBody
	GetNextToken() *string
	SetPolicySets(v []*ListPolicySetsResponseBodyPolicySets) *ListPolicySetsResponseBody
	GetPolicySets() []*ListPolicySetsResponseBodyPolicySets
	SetRequestId(v string) *ListPolicySetsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPolicySetsResponseBody
	GetTotalCount() *int32
}

type ListPolicySetsResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAdDWBF2
	NextToken  *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PolicySets []*ListPolicySetsResponseBodyPolicySets `json:"PolicySets,omitempty" xml:"PolicySets,omitempty" type:"Repeated"`
	// example:
	//
	// 2A48EB1D-D645-5758-91AF-EDF8E36E257B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPolicySetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPolicySetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicySetsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPolicySetsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPolicySetsResponseBody) GetPolicySets() []*ListPolicySetsResponseBodyPolicySets {
	return s.PolicySets
}

func (s *ListPolicySetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPolicySetsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPolicySetsResponseBody) SetMaxResults(v int32) *ListPolicySetsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPolicySetsResponseBody) SetNextToken(v string) *ListPolicySetsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPolicySetsResponseBody) SetPolicySets(v []*ListPolicySetsResponseBodyPolicySets) *ListPolicySetsResponseBody {
	s.PolicySets = v
	return s
}

func (s *ListPolicySetsResponseBody) SetRequestId(v string) *ListPolicySetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolicySetsResponseBody) SetTotalCount(v int32) *ListPolicySetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPolicySetsResponseBody) Validate() error {
	if s.PolicySets != nil {
		for _, item := range s.PolicySets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPolicySetsResponseBodyPolicySets struct {
	// example:
	//
	// 2026-05-08T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:policyset/default-policy-set
	PolicySetArn *string `json:"PolicySetArn,omitempty" xml:"PolicySetArn,omitempty"`
	// example:
	//
	// default-policy-set
	PolicySetName *string `json:"PolicySetName,omitempty" xml:"PolicySetName,omitempty"`
	// example:
	//
	// 2026-05-08T06:19:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListPolicySetsResponseBodyPolicySets) String() string {
	return dara.Prettify(s)
}

func (s ListPolicySetsResponseBodyPolicySets) GoString() string {
	return s.String()
}

func (s *ListPolicySetsResponseBodyPolicySets) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPolicySetsResponseBodyPolicySets) GetDescription() *string {
	return s.Description
}

func (s *ListPolicySetsResponseBodyPolicySets) GetPolicySetArn() *string {
	return s.PolicySetArn
}

func (s *ListPolicySetsResponseBodyPolicySets) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *ListPolicySetsResponseBodyPolicySets) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListPolicySetsResponseBodyPolicySets) SetCreateTime(v string) *ListPolicySetsResponseBodyPolicySets {
	s.CreateTime = &v
	return s
}

func (s *ListPolicySetsResponseBodyPolicySets) SetDescription(v string) *ListPolicySetsResponseBodyPolicySets {
	s.Description = &v
	return s
}

func (s *ListPolicySetsResponseBodyPolicySets) SetPolicySetArn(v string) *ListPolicySetsResponseBodyPolicySets {
	s.PolicySetArn = &v
	return s
}

func (s *ListPolicySetsResponseBodyPolicySets) SetPolicySetName(v string) *ListPolicySetsResponseBodyPolicySets {
	s.PolicySetName = &v
	return s
}

func (s *ListPolicySetsResponseBodyPolicySets) SetUpdateTime(v string) *ListPolicySetsResponseBodyPolicySets {
	s.UpdateTime = &v
	return s
}

func (s *ListPolicySetsResponseBodyPolicySets) Validate() error {
	return dara.Validate(s)
}
