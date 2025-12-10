// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetControlPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetControlPolicy(v *GetControlPolicyResponseBodyControlPolicy) *GetControlPolicyResponseBody
	GetControlPolicy() *GetControlPolicyResponseBodyControlPolicy
	SetRequestId(v string) *GetControlPolicyResponseBody
	GetRequestId() *string
}

type GetControlPolicyResponseBody struct {
	// The details of the access control policy.
	ControlPolicy *GetControlPolicyResponseBodyControlPolicy `json:"ControlPolicy,omitempty" xml:"ControlPolicy,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// AB769936-CDFA-4D52-8CE2-A3581800044A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetControlPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetControlPolicyResponseBody) GetControlPolicy() *GetControlPolicyResponseBodyControlPolicy {
	return s.ControlPolicy
}

func (s *GetControlPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetControlPolicyResponseBody) SetControlPolicy(v *GetControlPolicyResponseBodyControlPolicy) *GetControlPolicyResponseBody {
	s.ControlPolicy = v
	return s
}

func (s *GetControlPolicyResponseBody) SetRequestId(v string) *GetControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetControlPolicyResponseBody) Validate() error {
	if s.ControlPolicy != nil {
		if err := s.ControlPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetControlPolicyResponseBodyControlPolicy struct {
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
	// 2021-03-18T08:51:33Z
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
	// The document of the access control policy.
	//
	// example:
	//
	// {\\"Version\\":\\"1\\",\\"Statement\\":[{\\"Effect\\":\\"Deny\\",\\"Action\\":[\\"ram:UpdateRole\\",\\"ram:DeleteRole\\",\\"ram:AttachPolicyToRole\\",\\"ram:DetachPolicyFromRole\\"],\\"Resource\\":\\"acs:ram:*:*:role/ResourceDirectoryAccountAccessRole\\"}]}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The ID of the access control policy.
	//
	// example:
	//
	// cp-SImPt8GCEwiq****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the access control policy.
	//
	// example:
	//
	// test
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
	// 2021-03-18T08:51:33Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetControlPolicyResponseBodyControlPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetControlPolicyResponseBodyControlPolicy) GoString() string {
	return s.String()
}

func (s *GetControlPolicyResponseBodyControlPolicy) GetAttachmentCount() *string {
	return s.AttachmentCount
}

func (s *GetControlPolicyResponseBodyControlPolicy) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetControlPolicyResponseBodyControlPolicy) GetDescription() *string {
	return s.Description
}

func (s *GetControlPolicyResponseBodyControlPolicy) GetEffectScope() *string {
	return s.EffectScope
}

func (s *GetControlPolicyResponseBodyControlPolicy) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *GetControlPolicyResponseBodyControlPolicy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetControlPolicyResponseBodyControlPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *GetControlPolicyResponseBodyControlPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetControlPolicyResponseBodyControlPolicy) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetAttachmentCount(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.AttachmentCount = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetCreateDate(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.CreateDate = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetDescription(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.Description = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetEffectScope(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.EffectScope = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetPolicyDocument(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.PolicyDocument = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetPolicyId(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.PolicyId = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetPolicyName(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.PolicyName = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetPolicyType(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.PolicyType = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) SetUpdateDate(v string) *GetControlPolicyResponseBodyControlPolicy {
	s.UpdateDate = &v
	return s
}

func (s *GetControlPolicyResponseBodyControlPolicy) Validate() error {
	return dara.Validate(s)
}
