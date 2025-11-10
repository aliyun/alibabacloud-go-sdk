// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetControlPolicy(v *UpdateControlPolicyResponseBodyControlPolicy) *UpdateControlPolicyResponseBody
	GetControlPolicy() *UpdateControlPolicyResponseBodyControlPolicy
	SetRequestId(v string) *UpdateControlPolicyResponseBody
	GetRequestId() *string
}

type UpdateControlPolicyResponseBody struct {
	// The details of the access control policy.
	ControlPolicy *UpdateControlPolicyResponseBodyControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 2DFCE4C9-04A9-4C83-BB14-FE791275EC53
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateControlPolicyResponseBody) GetControlPolicy() *UpdateControlPolicyResponseBodyControlPolicy {
	return s.ControlPolicy
}

func (s *UpdateControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateControlPolicyResponseBody) SetControlPolicy(v *UpdateControlPolicyResponseBodyControlPolicy) *UpdateControlPolicyResponseBody {
	s.ControlPolicy = v
	return s
}

func (s *UpdateControlPolicyResponseBody) SetRequestId(v string) *UpdateControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateControlPolicyResponseBody) Validate() error {
	if s.ControlPolicy != nil {
		if err := s.ControlPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateControlPolicyResponseBodyControlPolicy struct {
	// The number of times that the access control policy is referenced.
	//
	// example:
	//
	// 0
	AttachmentCount *string `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	// The time when the access control policy was created.
	//
	// example:
	//
	// 2021-03-18T09:24:19Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the access control policy.
	//
	// example:
	//
	// ExampleControlPolicy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The effective scope of the access control policy. Valid values:
	//
	// 	- All: The access control policy is in effect for Alibaba Cloud accounts, RAM users, and RAM roles.
	//
	// 	- RAM: The access control policy is in effect only for RAM users and RAM roles.
	//
	// example:
	//
	// RAM
	EffectScope *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	// The ID of the access control policy.
	//
	// example:
	//
	// cp-jExXAqIYkwHN****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the access control policy.
	//
	// example:
	//
	// NewControlPolicy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the access control policy. Valid values:
	//
	// 	- System: system access control policy
	//
	// 	- Custom: custom access control policy
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The time when the access control policy was updated.
	//
	// example:
	//
	// 2021-03-18T10:04:55Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s UpdateControlPolicyResponseBodyControlPolicy) String() string {
	return dara.Prettify(s)
}

func (s UpdateControlPolicyResponseBodyControlPolicy) GoString() string {
	return s.String()
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) GetAttachmentCount() *string {
	return s.AttachmentCount
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) GetCreateDate() *string {
	return s.CreateDate
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) GetDescription() *string {
	return s.Description
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) GetEffectScope() *string {
	return s.EffectScope
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetAttachmentCount(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetCreateDate(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.CreateDate = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetDescription(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.Description = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetEffectScope(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.EffectScope = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetPolicyId(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.PolicyId = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetPolicyName(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.PolicyName = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetPolicyType(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.PolicyType = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) SetUpdateDate(v string) *UpdateControlPolicyResponseBodyControlPolicy {
	s.UpdateDate = &v
	return s
}

func (s *UpdateControlPolicyResponseBodyControlPolicy) Validate() error {
	return dara.Validate(s)
}
