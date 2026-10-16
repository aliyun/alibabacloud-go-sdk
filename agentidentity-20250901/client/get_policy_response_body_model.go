// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v *GetPolicyResponseBodyPolicy) *GetPolicyResponseBody
	GetPolicy() *GetPolicyResponseBodyPolicy
	SetRequestId(v string) *GetPolicyResponseBody
	GetRequestId() *string
}

type GetPolicyResponseBody struct {
	// The policy details.
	Policy *GetPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The unique identifier of the request.
	//
	// example:
	//
	// 2A48EB1D-D645-5758-91AF-EDF8E36E257B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBody) GetPolicy() *GetPolicyResponseBodyPolicy {
	return s.Policy
}

func (s *GetPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPolicyResponseBody) SetPolicy(v *GetPolicyResponseBodyPolicy) *GetPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *GetPolicyResponseBody) SetRequestId(v string) *GetPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolicyResponseBody) Validate() error {
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPolicyResponseBodyPolicy struct {
	// The creation time in UTC ISO 8601 format.
	//
	// example:
	//
	// 2026-05-08T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The policy definition content.
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
	// The name of the policy. The name must be unique within the associated policy set.
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
	// The update time in UTC ISO 8601 format.
	//
	// example:
	//
	// 2026-05-08T06:19:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetPolicyResponseBodyPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBodyPolicy) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPolicyResponseBodyPolicy) GetDefinition() *Definition {
	return s.Definition
}

func (s *GetPolicyResponseBodyPolicy) GetDescription() *string {
	return s.Description
}

func (s *GetPolicyResponseBodyPolicy) GetPolicyArn() *string {
	return s.PolicyArn
}

func (s *GetPolicyResponseBodyPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *GetPolicyResponseBodyPolicy) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *GetPolicyResponseBodyPolicy) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetPolicyResponseBodyPolicy) SetCreateTime(v string) *GetPolicyResponseBodyPolicy {
	s.CreateTime = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetDefinition(v *Definition) *GetPolicyResponseBodyPolicy {
	s.Definition = v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetDescription(v string) *GetPolicyResponseBodyPolicy {
	s.Description = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetPolicyArn(v string) *GetPolicyResponseBodyPolicy {
	s.PolicyArn = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetPolicyName(v string) *GetPolicyResponseBodyPolicy {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetPolicySetName(v string) *GetPolicyResponseBodyPolicy {
	s.PolicySetName = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetUpdateTime(v string) *GetPolicyResponseBodyPolicy {
	s.UpdateTime = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) Validate() error {
	if s.Definition != nil {
		if err := s.Definition.Validate(); err != nil {
			return err
		}
	}
	return nil
}
