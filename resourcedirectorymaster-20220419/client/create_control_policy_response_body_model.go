// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetControlPolicy(v *CreateControlPolicyResponseBodyControlPolicy) *CreateControlPolicyResponseBody
	GetControlPolicy() *CreateControlPolicyResponseBodyControlPolicy
	SetRequestId(v string) *CreateControlPolicyResponseBody
	GetRequestId() *string
}

type CreateControlPolicyResponseBody struct {
	// The details of the access control policy.
	ControlPolicy *CreateControlPolicyResponseBodyControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 776B05B3-A0B0-464B-A191-F4E1119A94B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateControlPolicyResponseBody) GetControlPolicy() *CreateControlPolicyResponseBodyControlPolicy {
	return s.ControlPolicy
}

func (s *CreateControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateControlPolicyResponseBody) SetControlPolicy(v *CreateControlPolicyResponseBodyControlPolicy) *CreateControlPolicyResponseBody {
	s.ControlPolicy = v
	return s
}

func (s *CreateControlPolicyResponseBody) SetRequestId(v string) *CreateControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateControlPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateControlPolicyResponseBodyControlPolicy struct {
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
	// The effective scope of the access control policy.
	//
	// The value RAM indicates that the access control policy takes effect only for RAM users and RAM roles.
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
	// ExampleControlPolicy
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
	// 2021-03-18T09:24:19Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s CreateControlPolicyResponseBodyControlPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateControlPolicyResponseBodyControlPolicy) GoString() string {
	return s.String()
}

func (s *CreateControlPolicyResponseBodyControlPolicy) GetAttachmentCount() *string {
	return s.AttachmentCount
}

func (s *CreateControlPolicyResponseBodyControlPolicy) GetCreateDate() *string {
	return s.CreateDate
}

func (s *CreateControlPolicyResponseBodyControlPolicy) GetDescription() *string {
	return s.Description
}

func (s *CreateControlPolicyResponseBodyControlPolicy) GetEffectScope() *string {
	return s.EffectScope
}

func (s *CreateControlPolicyResponseBodyControlPolicy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreateControlPolicyResponseBodyControlPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreateControlPolicyResponseBodyControlPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateControlPolicyResponseBodyControlPolicy) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetAttachmentCount(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetCreateDate(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.CreateDate = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetDescription(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.Description = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetEffectScope(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.EffectScope = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetPolicyId(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.PolicyId = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetPolicyName(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.PolicyName = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetPolicyType(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.PolicyType = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) SetUpdateDate(v string) *CreateControlPolicyResponseBodyControlPolicy {
	s.UpdateDate = &v
	return s
}

func (s *CreateControlPolicyResponseBodyControlPolicy) Validate() error {
	return dara.Validate(s)
}
