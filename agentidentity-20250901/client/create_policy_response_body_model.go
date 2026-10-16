// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v *CreatePolicyResponseBodyPolicy) *CreatePolicyResponseBody
	GetPolicy() *CreatePolicyResponseBodyPolicy
	SetRequestId(v string) *CreatePolicyResponseBody
	GetRequestId() *string
}

type CreatePolicyResponseBody struct {
	// The policy details.
	Policy *CreatePolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The unique identifier of the request.
	//
	// example:
	//
	// 2A48EB1D-D645-5758-91AF-EDF8E36E257B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponseBody) GetPolicy() *CreatePolicyResponseBodyPolicy {
	return s.Policy
}

func (s *CreatePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolicyResponseBody) SetPolicy(v *CreatePolicyResponseBodyPolicy) *CreatePolicyResponseBody {
	s.Policy = v
	return s
}

func (s *CreatePolicyResponseBody) SetRequestId(v string) *CreatePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyResponseBody) Validate() error {
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePolicyResponseBodyPolicy struct {
	// The time when the policy was created, in UTC ISO 8601 format.
	//
	// example:
	//
	// 2026-05-08T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The policy definition.
	Definition *Definition `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The description.
	//
	// example:
	//
	// example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ARN of the policy.
	//
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:policyset/default-policy-set/policy/rate-limit-policy
	PolicyArn *string `json:"PolicyArn,omitempty" xml:"PolicyArn,omitempty"`
	// The name of the policy. The name must be unique within the policy set to which the policy belongs.
	//
	// example:
	//
	// rate-limit-policy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The name of the policy set. The name must be unique within the current Alibaba Cloud account.
	//
	// example:
	//
	// default-policy-set
	PolicySetName *string `json:"PolicySetName,omitempty" xml:"PolicySetName,omitempty"`
}

func (s CreatePolicyResponseBodyPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponseBodyPolicy) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreatePolicyResponseBodyPolicy) GetDefinition() *Definition {
	return s.Definition
}

func (s *CreatePolicyResponseBodyPolicy) GetDescription() *string {
	return s.Description
}

func (s *CreatePolicyResponseBodyPolicy) GetPolicyArn() *string {
	return s.PolicyArn
}

func (s *CreatePolicyResponseBodyPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreatePolicyResponseBodyPolicy) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *CreatePolicyResponseBodyPolicy) SetCreateTime(v string) *CreatePolicyResponseBodyPolicy {
	s.CreateTime = &v
	return s
}

func (s *CreatePolicyResponseBodyPolicy) SetDefinition(v *Definition) *CreatePolicyResponseBodyPolicy {
	s.Definition = v
	return s
}

func (s *CreatePolicyResponseBodyPolicy) SetDescription(v string) *CreatePolicyResponseBodyPolicy {
	s.Description = &v
	return s
}

func (s *CreatePolicyResponseBodyPolicy) SetPolicyArn(v string) *CreatePolicyResponseBodyPolicy {
	s.PolicyArn = &v
	return s
}

func (s *CreatePolicyResponseBodyPolicy) SetPolicyName(v string) *CreatePolicyResponseBodyPolicy {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyResponseBodyPolicy) SetPolicySetName(v string) *CreatePolicyResponseBodyPolicy {
	s.PolicySetName = &v
	return s
}

func (s *CreatePolicyResponseBodyPolicy) Validate() error {
	if s.Definition != nil {
		if err := s.Definition.Validate(); err != nil {
			return err
		}
	}
	return nil
}
