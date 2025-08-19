// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListControlPolicyAttachmentsForTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetControlPolicyAttachments(v *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments) *ListControlPolicyAttachmentsForTargetResponseBody
	GetControlPolicyAttachments() *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments
	SetRequestId(v string) *ListControlPolicyAttachmentsForTargetResponseBody
	GetRequestId() *string
}

type ListControlPolicyAttachmentsForTargetResponseBody struct {
	// The information about the attached access control policies.
	ControlPolicyAttachments *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments `json:"ControlPolicyAttachments,omitempty" xml:"ControlPolicyAttachments,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C276B600-7B7A-49E8-938C-E16CFA955A82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListControlPolicyAttachmentsForTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListControlPolicyAttachmentsForTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ListControlPolicyAttachmentsForTargetResponseBody) GetControlPolicyAttachments() *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments {
	return s.ControlPolicyAttachments
}

func (s *ListControlPolicyAttachmentsForTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListControlPolicyAttachmentsForTargetResponseBody) SetControlPolicyAttachments(v *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments) *ListControlPolicyAttachmentsForTargetResponseBody {
	s.ControlPolicyAttachments = v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBody) SetRequestId(v string) *ListControlPolicyAttachmentsForTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments struct {
	ControlPolicyAttachment []*ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment `json:"ControlPolicyAttachment,omitempty" xml:"ControlPolicyAttachment,omitempty" type:"Repeated"`
}

func (s ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments) String() string {
	return dara.Prettify(s)
}

func (s ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments) GoString() string {
	return s.String()
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments) GetControlPolicyAttachment() []*ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment {
	return s.ControlPolicyAttachment
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments) SetControlPolicyAttachment(v []*ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments {
	s.ControlPolicyAttachment = v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachments) Validate() error {
	return dara.Validate(s)
}

type ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment struct {
	// The time when the access control policy was attached.
	//
	// example:
	//
	// 2021-03-19T02:56:24Z
	AttachDate *string `json:"AttachDate,omitempty" xml:"AttachDate,omitempty"`
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
}

func (s ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) String() string {
	return dara.Prettify(s)
}

func (s ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) GoString() string {
	return s.String()
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) GetAttachDate() *string {
	return s.AttachDate
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) GetDescription() *string {
	return s.Description
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) GetEffectScope() *string {
	return s.EffectScope
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) SetAttachDate(v string) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment {
	s.AttachDate = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) SetDescription(v string) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment {
	s.Description = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) SetEffectScope(v string) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment {
	s.EffectScope = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) SetPolicyId(v string) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment {
	s.PolicyId = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) SetPolicyName(v string) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment {
	s.PolicyName = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) SetPolicyType(v string) *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment {
	s.PolicyType = &v
	return s
}

func (s *ListControlPolicyAttachmentsForTargetResponseBodyControlPolicyAttachmentsControlPolicyAttachment) Validate() error {
	return dara.Validate(s)
}
