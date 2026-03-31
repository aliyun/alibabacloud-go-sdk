// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicies(v *ListPoliciesForUserResponseBodyPolicies) *ListPoliciesForUserResponseBody
	GetPolicies() *ListPoliciesForUserResponseBodyPolicies
	SetRequestId(v string) *ListPoliciesForUserResponseBody
	GetRequestId() *string
}

type ListPoliciesForUserResponseBody struct {
	Policies *ListPoliciesForUserResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPoliciesForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesForUserResponseBody) GetPolicies() *ListPoliciesForUserResponseBodyPolicies {
	return s.Policies
}

func (s *ListPoliciesForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPoliciesForUserResponseBody) SetPolicies(v *ListPoliciesForUserResponseBodyPolicies) *ListPoliciesForUserResponseBody {
	s.Policies = v
	return s
}

func (s *ListPoliciesForUserResponseBody) SetRequestId(v string) *ListPoliciesForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPoliciesForUserResponseBody) Validate() error {
	if s.Policies != nil {
		if err := s.Policies.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPoliciesForUserResponseBodyPolicies struct {
	Policy []*ListPoliciesForUserResponseBodyPoliciesPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Repeated"`
}

func (s ListPoliciesForUserResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesForUserResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListPoliciesForUserResponseBodyPolicies) GetPolicy() []*ListPoliciesForUserResponseBodyPoliciesPolicy {
	return s.Policy
}

func (s *ListPoliciesForUserResponseBodyPolicies) SetPolicy(v []*ListPoliciesForUserResponseBodyPoliciesPolicy) *ListPoliciesForUserResponseBodyPolicies {
	s.Policy = v
	return s
}

func (s *ListPoliciesForUserResponseBodyPolicies) Validate() error {
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

type ListPoliciesForUserResponseBodyPoliciesPolicy struct {
	AttachDate     *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType     *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListPoliciesForUserResponseBodyPoliciesPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesForUserResponseBodyPoliciesPolicy) GoString() string {
	return s.String()
}

func (s *ListPoliciesForUserResponseBodyPoliciesPolicy) GetAttachDate() *string {
	return s.AttachDate
}

func (s *ListPoliciesForUserResponseBodyPoliciesPolicy) GetDefaultVersion() *string {
	return s.DefaultVersion
}

func (s *ListPoliciesForUserResponseBodyPoliciesPolicy) GetDescription() *string {
	return s.Description
}

func (s *ListPoliciesForUserResponseBodyPoliciesPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListPoliciesForUserResponseBodyPoliciesPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListPoliciesForUserResponseBodyPoliciesPolicy) SetAttachDate(v string) *ListPoliciesForUserResponseBodyPoliciesPolicy {
	s.AttachDate = &v
	return s
}

func (s *ListPoliciesForUserResponseBodyPoliciesPolicy) SetDefaultVersion(v string) *ListPoliciesForUserResponseBodyPoliciesPolicy {
	s.DefaultVersion = &v
	return s
}

func (s *ListPoliciesForUserResponseBodyPoliciesPolicy) SetDescription(v string) *ListPoliciesForUserResponseBodyPoliciesPolicy {
	s.Description = &v
	return s
}

func (s *ListPoliciesForUserResponseBodyPoliciesPolicy) SetPolicyName(v string) *ListPoliciesForUserResponseBodyPoliciesPolicy {
	s.PolicyName = &v
	return s
}

func (s *ListPoliciesForUserResponseBodyPoliciesPolicy) SetPolicyType(v string) *ListPoliciesForUserResponseBodyPoliciesPolicy {
	s.PolicyType = &v
	return s
}

func (s *ListPoliciesForUserResponseBodyPoliciesPolicy) Validate() error {
	return dara.Validate(s)
}
