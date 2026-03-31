// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesForRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicies(v *ListPoliciesForRoleResponseBodyPolicies) *ListPoliciesForRoleResponseBody
	GetPolicies() *ListPoliciesForRoleResponseBodyPolicies
	SetRequestId(v string) *ListPoliciesForRoleResponseBody
	GetRequestId() *string
}

type ListPoliciesForRoleResponseBody struct {
	Policies *ListPoliciesForRoleResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPoliciesForRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesForRoleResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesForRoleResponseBody) GetPolicies() *ListPoliciesForRoleResponseBodyPolicies {
	return s.Policies
}

func (s *ListPoliciesForRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPoliciesForRoleResponseBody) SetPolicies(v *ListPoliciesForRoleResponseBodyPolicies) *ListPoliciesForRoleResponseBody {
	s.Policies = v
	return s
}

func (s *ListPoliciesForRoleResponseBody) SetRequestId(v string) *ListPoliciesForRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPoliciesForRoleResponseBody) Validate() error {
	if s.Policies != nil {
		if err := s.Policies.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPoliciesForRoleResponseBodyPolicies struct {
	Policy []*ListPoliciesForRoleResponseBodyPoliciesPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Repeated"`
}

func (s ListPoliciesForRoleResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesForRoleResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListPoliciesForRoleResponseBodyPolicies) GetPolicy() []*ListPoliciesForRoleResponseBodyPoliciesPolicy {
	return s.Policy
}

func (s *ListPoliciesForRoleResponseBodyPolicies) SetPolicy(v []*ListPoliciesForRoleResponseBodyPoliciesPolicy) *ListPoliciesForRoleResponseBodyPolicies {
	s.Policy = v
	return s
}

func (s *ListPoliciesForRoleResponseBodyPolicies) Validate() error {
	if s.Policy != nil {
		for _, item := range s.Policy {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPoliciesForRoleResponseBodyPoliciesPolicy struct {
	AttachDate     *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType     *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListPoliciesForRoleResponseBodyPoliciesPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesForRoleResponseBodyPoliciesPolicy) GoString() string {
	return s.String()
}

func (s *ListPoliciesForRoleResponseBodyPoliciesPolicy) GetAttachDate() *string {
	return s.AttachDate
}

func (s *ListPoliciesForRoleResponseBodyPoliciesPolicy) GetDefaultVersion() *string {
	return s.DefaultVersion
}

func (s *ListPoliciesForRoleResponseBodyPoliciesPolicy) GetDescription() *string {
	return s.Description
}

func (s *ListPoliciesForRoleResponseBodyPoliciesPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListPoliciesForRoleResponseBodyPoliciesPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListPoliciesForRoleResponseBodyPoliciesPolicy) SetAttachDate(v string) *ListPoliciesForRoleResponseBodyPoliciesPolicy {
	s.AttachDate = &v
	return s
}

func (s *ListPoliciesForRoleResponseBodyPoliciesPolicy) SetDefaultVersion(v string) *ListPoliciesForRoleResponseBodyPoliciesPolicy {
	s.DefaultVersion = &v
	return s
}

func (s *ListPoliciesForRoleResponseBodyPoliciesPolicy) SetDescription(v string) *ListPoliciesForRoleResponseBodyPoliciesPolicy {
	s.Description = &v
	return s
}

func (s *ListPoliciesForRoleResponseBodyPoliciesPolicy) SetPolicyName(v string) *ListPoliciesForRoleResponseBodyPoliciesPolicy {
	s.PolicyName = &v
	return s
}

func (s *ListPoliciesForRoleResponseBodyPoliciesPolicy) SetPolicyType(v string) *ListPoliciesForRoleResponseBodyPoliciesPolicy {
	s.PolicyType = &v
	return s
}

func (s *ListPoliciesForRoleResponseBodyPoliciesPolicy) Validate() error {
	return dara.Validate(s)
}
