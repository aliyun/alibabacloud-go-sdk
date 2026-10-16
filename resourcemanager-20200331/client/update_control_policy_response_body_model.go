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
	// The details of the control policy.
	ControlPolicy *UpdateControlPolicyResponseBodyControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" type:"Struct"`
	// The request ID.
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
	// The number of times the control policy is referenced.
	//
	// example:
	//
	// 0
	AttachmentCount *string `json:"AttachmentCount,omitempty" xml:"AttachmentCount,omitempty"`
	// The time when the control policy was created. The time is displayed in UTC in RFC 3339 format, indicated by Z.
	//
	// example:
	//
	// 2021-03-18T09:24:19Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the control policy.
	//
	// example:
	//
	// ExampleControlPolicy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The effective scope of the control policy. Valid values:
	//
	// - All: The control policy takes effect for Alibaba Cloud accounts, Resource Access Management (RAM) users, and RAM roles.
	//
	// - RAM: The control policy takes effect only for RAM users and RAM roles.
	//
	// example:
	//
	// RAM
	EffectScope *string `json:"EffectScope,omitempty" xml:"EffectScope,omitempty"`
	// The ID of the control policy.
	//
	// example:
	//
	// cp-jExXAqIYkwHN****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The Policy Name of the control policy.
	//
	// example:
	//
	// NewControlPolicy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the control policy. Valid values:
	//
	// - System: system control policy.
	//
	// - Custom: custom control policy.
	//
	// example:
	//
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The time when the control policy was updated. The time is displayed in UTC in RFC 3339 format, indicated by Z.
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
