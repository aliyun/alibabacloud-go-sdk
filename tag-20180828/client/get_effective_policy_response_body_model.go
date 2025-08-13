// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEffectivePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEffectivePolicy(v string) *GetEffectivePolicyResponseBody
	GetEffectivePolicy() *string
	SetPolicyAttachments(v []*GetEffectivePolicyResponseBodyPolicyAttachments) *GetEffectivePolicyResponseBody
	GetPolicyAttachments() []*GetEffectivePolicyResponseBodyPolicyAttachments
	SetRequestId(v string) *GetEffectivePolicyResponseBody
	GetRequestId() *string
}

type GetEffectivePolicyResponseBody struct {
	// The effective tag policy.
	//
	// example:
	//
	// {\\"tags\\":{\\"costcenter\\":{\\"tag_value\\":[\\"Beijing\\",\\"Shanghai\\"],\\"tag_key\\":\\"CostCenter\\"}}}
	EffectivePolicy   *string                                            `json:"EffectivePolicy,omitempty" xml:"EffectivePolicy,omitempty"`
	PolicyAttachments []*GetEffectivePolicyResponseBodyPolicyAttachments `json:"PolicyAttachments,omitempty" xml:"PolicyAttachments,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// BB532282-94F5-5F56-877F-32D5E2A04F3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEffectivePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEffectivePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetEffectivePolicyResponseBody) GetEffectivePolicy() *string {
	return s.EffectivePolicy
}

func (s *GetEffectivePolicyResponseBody) GetPolicyAttachments() []*GetEffectivePolicyResponseBodyPolicyAttachments {
	return s.PolicyAttachments
}

func (s *GetEffectivePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEffectivePolicyResponseBody) SetEffectivePolicy(v string) *GetEffectivePolicyResponseBody {
	s.EffectivePolicy = &v
	return s
}

func (s *GetEffectivePolicyResponseBody) SetPolicyAttachments(v []*GetEffectivePolicyResponseBodyPolicyAttachments) *GetEffectivePolicyResponseBody {
	s.PolicyAttachments = v
	return s
}

func (s *GetEffectivePolicyResponseBody) SetRequestId(v string) *GetEffectivePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEffectivePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetEffectivePolicyResponseBodyPolicyAttachments struct {
	PolicyList []*GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList `json:"PolicyList,omitempty" xml:"PolicyList,omitempty" type:"Repeated"`
	PolicyType *string                                                      `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	TagKey     *string                                                      `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s GetEffectivePolicyResponseBodyPolicyAttachments) String() string {
	return dara.Prettify(s)
}

func (s GetEffectivePolicyResponseBodyPolicyAttachments) GoString() string {
	return s.String()
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachments) GetPolicyList() []*GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList {
	return s.PolicyList
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachments) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachments) GetTagKey() *string {
	return s.TagKey
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachments) SetPolicyList(v []*GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList) *GetEffectivePolicyResponseBodyPolicyAttachments {
	s.PolicyList = v
	return s
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachments) SetPolicyType(v string) *GetEffectivePolicyResponseBodyPolicyAttachments {
	s.PolicyType = &v
	return s
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachments) SetTagKey(v string) *GetEffectivePolicyResponseBodyPolicyAttachments {
	s.TagKey = &v
	return s
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachments) Validate() error {
	return dara.Validate(s)
}

type GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList struct {
	AttachSeq  *int32  `json:"AttachSeq,omitempty" xml:"AttachSeq,omitempty"`
	AttachTime *string `json:"AttachTime,omitempty" xml:"AttachTime,omitempty"`
	PolicyId   *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	TargetId   *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList) String() string {
	return dara.Prettify(s)
}

func (s GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList) GoString() string {
	return s.String()
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList) GetAttachSeq() *int32 {
	return s.AttachSeq
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList) GetAttachTime() *string {
	return s.AttachTime
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList) GetPolicyName() *string {
	return s.PolicyName
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList) GetTargetId() *string {
	return s.TargetId
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList) GetTargetType() *string {
	return s.TargetType
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList) SetAttachSeq(v int32) *GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList {
	s.AttachSeq = &v
	return s
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList) SetAttachTime(v string) *GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList {
	s.AttachTime = &v
	return s
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList) SetPolicyId(v string) *GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList {
	s.PolicyId = &v
	return s
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList) SetPolicyName(v string) *GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList {
	s.PolicyName = &v
	return s
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList) SetTargetId(v string) *GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList {
	s.TargetId = &v
	return s
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList) SetTargetType(v string) *GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList {
	s.TargetType = &v
	return s
}

func (s *GetEffectivePolicyResponseBodyPolicyAttachmentsPolicyList) Validate() error {
	return dara.Validate(s)
}
