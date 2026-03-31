// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesForGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicies(v *ListPoliciesForGroupResponseBodyPolicies) *ListPoliciesForGroupResponseBody
	GetPolicies() *ListPoliciesForGroupResponseBodyPolicies
	SetRequestId(v string) *ListPoliciesForGroupResponseBody
	GetRequestId() *string
}

type ListPoliciesForGroupResponseBody struct {
	Policies *ListPoliciesForGroupResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7B8A4E7D-6CFF-471D-84DF-195A7A241ECB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPoliciesForGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesForGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesForGroupResponseBody) GetPolicies() *ListPoliciesForGroupResponseBodyPolicies {
	return s.Policies
}

func (s *ListPoliciesForGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPoliciesForGroupResponseBody) SetPolicies(v *ListPoliciesForGroupResponseBodyPolicies) *ListPoliciesForGroupResponseBody {
	s.Policies = v
	return s
}

func (s *ListPoliciesForGroupResponseBody) SetRequestId(v string) *ListPoliciesForGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPoliciesForGroupResponseBody) Validate() error {
	if s.Policies != nil {
		if err := s.Policies.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPoliciesForGroupResponseBodyPolicies struct {
	Policy []*ListPoliciesForGroupResponseBodyPoliciesPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Repeated"`
}

func (s ListPoliciesForGroupResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesForGroupResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListPoliciesForGroupResponseBodyPolicies) GetPolicy() []*ListPoliciesForGroupResponseBodyPoliciesPolicy {
	return s.Policy
}

func (s *ListPoliciesForGroupResponseBodyPolicies) SetPolicy(v []*ListPoliciesForGroupResponseBodyPoliciesPolicy) *ListPoliciesForGroupResponseBodyPolicies {
	s.Policy = v
	return s
}

func (s *ListPoliciesForGroupResponseBodyPolicies) Validate() error {
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

type ListPoliciesForGroupResponseBodyPoliciesPolicy struct {
	AttachDate     *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
	DefaultVersion *string `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PolicyName     *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType     *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListPoliciesForGroupResponseBodyPoliciesPolicy) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesForGroupResponseBodyPoliciesPolicy) GoString() string {
	return s.String()
}

func (s *ListPoliciesForGroupResponseBodyPoliciesPolicy) GetAttachDate() *string {
	return s.AttachDate
}

func (s *ListPoliciesForGroupResponseBodyPoliciesPolicy) GetDefaultVersion() *string {
	return s.DefaultVersion
}

func (s *ListPoliciesForGroupResponseBodyPoliciesPolicy) GetDescription() *string {
	return s.Description
}

func (s *ListPoliciesForGroupResponseBodyPoliciesPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListPoliciesForGroupResponseBodyPoliciesPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListPoliciesForGroupResponseBodyPoliciesPolicy) SetAttachDate(v string) *ListPoliciesForGroupResponseBodyPoliciesPolicy {
	s.AttachDate = &v
	return s
}

func (s *ListPoliciesForGroupResponseBodyPoliciesPolicy) SetDefaultVersion(v string) *ListPoliciesForGroupResponseBodyPoliciesPolicy {
	s.DefaultVersion = &v
	return s
}

func (s *ListPoliciesForGroupResponseBodyPoliciesPolicy) SetDescription(v string) *ListPoliciesForGroupResponseBodyPoliciesPolicy {
	s.Description = &v
	return s
}

func (s *ListPoliciesForGroupResponseBodyPoliciesPolicy) SetPolicyName(v string) *ListPoliciesForGroupResponseBodyPoliciesPolicy {
	s.PolicyName = &v
	return s
}

func (s *ListPoliciesForGroupResponseBodyPoliciesPolicy) SetPolicyType(v string) *ListPoliciesForGroupResponseBodyPoliciesPolicy {
	s.PolicyType = &v
	return s
}

func (s *ListPoliciesForGroupResponseBodyPoliciesPolicy) Validate() error {
	return dara.Validate(s)
}
