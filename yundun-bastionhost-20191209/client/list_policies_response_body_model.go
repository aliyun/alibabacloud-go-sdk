// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicies(v []*ListPoliciesResponseBodyPolicies) *ListPoliciesResponseBody
	GetPolicies() []*ListPoliciesResponseBodyPolicies
	SetRequestId(v string) *ListPoliciesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListPoliciesResponseBody
	GetTotalCount() *int64
}

type ListPoliciesResponseBody struct {
	// The control policies.
	Policies []*ListPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of control policies that are returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBody) GetPolicies() []*ListPoliciesResponseBodyPolicies {
	return s.Policies
}

func (s *ListPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPoliciesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListPoliciesResponseBody) SetPolicies(v []*ListPoliciesResponseBodyPolicies) *ListPoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *ListPoliciesResponseBody) SetRequestId(v string) *ListPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPoliciesResponseBody) SetTotalCount(v int64) *ListPoliciesResponseBody {
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
	// The remarks of the control policy.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The control policy ID.
	//
	// example:
	//
	// 2
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the control policy.
	//
	// example:
	//
	// test
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The priority of the control policy. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s ListPoliciesResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPolicies) GetComment() *string {
	return s.Comment
}

func (s *ListPoliciesResponseBodyPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListPoliciesResponseBodyPolicies) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListPoliciesResponseBodyPolicies) GetPriority() *int64 {
	return s.Priority
}

func (s *ListPoliciesResponseBodyPolicies) SetComment(v string) *ListPoliciesResponseBodyPolicies {
	s.Comment = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) SetPolicyId(v string) *ListPoliciesResponseBodyPolicies {
	s.PolicyId = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) SetPolicyName(v string) *ListPoliciesResponseBodyPolicies {
	s.PolicyName = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) SetPriority(v int64) *ListPoliciesResponseBodyPolicies {
	s.Priority = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) Validate() error {
	return dara.Validate(s)
}
