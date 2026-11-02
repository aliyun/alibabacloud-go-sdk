// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePolicyShrinkRequest
	GetDescription() *string
	SetPolicyDocument(v string) *CreatePolicyShrinkRequest
	GetPolicyDocument() *string
	SetPolicyName(v string) *CreatePolicyShrinkRequest
	GetPolicyName() *string
	SetTagShrink(v string) *CreatePolicyShrinkRequest
	GetTagShrink() *string
}

type CreatePolicyShrinkRequest struct {
	// The description of the access policy.
	//
	// The description must be 1 to 1,024 characters in length.
	//
	// example:
	//
	// View ECS instances in a specified region
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Access policy document.
	//
	// Access policy document must be 1 to 6,144 characters in length.
	//
	// For more information about policy document elements and examples, see [Policy document elements](https://help.aliyun.com/document_detail/93738.html) and [Policy document examples](https://help.aliyun.com/document_detail/210969.html).
	//
	// example:
	//
	// {"Statement": [{"Effect": "Allow","Action": "ecs:Describe*","Resource": "acs:ecs:cn-qingdao:*:instance/*"}],"Version": "1"}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// Policy Name of the access policy.
	//
	// Policy Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	//
	// example:
	//
	// View-ECS-instances-in-a-specific-region
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreatePolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePolicyShrinkRequest) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *CreatePolicyShrinkRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreatePolicyShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *CreatePolicyShrinkRequest) SetDescription(v string) *CreatePolicyShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetPolicyDocument(v string) *CreatePolicyShrinkRequest {
	s.PolicyDocument = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetPolicyName(v string) *CreatePolicyShrinkRequest {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetTagShrink(v string) *CreatePolicyShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *CreatePolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
